// goupnp is an implementation of a client for various UPnP services.
//
// For most uses, it is recommended to use the code-generated packages under
// github.com/huin/goupnp/v2/dcps. Example use is shown at
// http://godoc.org/github.com/huin/goupnp/v2/example
//
// A commonly used client is internetgateway1.WANPPPConnection1:
// http://godoc.org/github.com/huin/goupnp/v2/dcps/internetgateway1#WANPPPConnection1
//
// Currently only a couple of schemas have code generated for them from the UPnP
// example XML specifications. Not all methods will work on these clients,
// because the generated stubs contain the full set of specified methods from
// the XML specifications, and the discovered services will likely support a
// subset of those methods.
package discover

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/url"
	"time"

	"github.com/huin/goupnp/v2/httpu"
	"github.com/huin/goupnp/v2/ssdp"
	"golang.org/x/net/html/charset"
	errors "gopkg.in/src-d/go-errors.v1"
)

var (
	// ErrBadResponse matches errors due to bad responses received for a search.
	ErrBadResponse = errors.NewKind("bad response data in %s")
	// ErrBadResponseStatus matches errors due to HTTP response status codes that
	// were unexpected or that indicate failure.
	ErrBadResponseStatus = errors.NewKind("bad response status %d %s")
	// ErrResponseFrom matches errors annotating where a bad response came from.
	ErrResponseFrom = errors.NewKind("from %q")
)

// MaybeRootDevice contains either a RootDevice or an error.
type MaybeRootDevice struct {
	// Set iff Err == nil.
	Root *RootDevice

	// The location the device was discovered at. This can be used with
	// DeviceByURL, assuming the device is still present. A location represents
	// the discovery of a device, regardless of if there was an error probing
	// it.
	Location *url.URL

	// Any error encountered probing a discovered device.
	Err error
}

// Devices attempts to find targets of the given type. This is typically the
// entry-point for this package. searchTarget is typically a URN in the form
// "urn:schemas-upnp-org:device:..." or "urn:schemas-upnp-org:service:...". A
// single error is returned for errors while attempting to send the query. An
// error or RootDevice is returned for each discovered RootDevice.
func Devices(
	ctx context.Context,
	searchTarget string,
	searchOpts ...ssdp.SearchOption,
) ([]MaybeRootDevice, error) {
	httpu, err := httpu.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer httpu.Close()
	responses, err := ssdp.RawSearch(
		ctx,
		httpu,
		string(searchTarget),
		searchOpts...,
	)
	if err != nil {
		return nil, err
	}

	results := make([]MaybeRootDevice, len(responses))
	for i, response := range responses {
		maybe := &results[i]
		loc, err := response.Location()
		if err != nil {
			maybe.Err = ErrBadResponse.Wrap(err, "location")
			continue
		}
		maybe.Location = loc
		if root, err := DeviceByURL(ctx, loc); err != nil {
			maybe.Err = err
		} else {
			maybe.Root = root
		}
	}

	return results, nil
}

func DeviceByURL(ctx context.Context, loc *url.URL) (*RootDevice, error) {
	locStr := loc.String()
	root := new(RootDevice)
	if err := requestXml(ctx, locStr, DeviceXMLNamespace, root); err != nil {
		return nil, ErrResponseFrom.Wrap(err, locStr)
	}
	var urlBaseStr string
	if root.URLBaseStr != "" {
		urlBaseStr = root.URLBaseStr
	} else {
		urlBaseStr = locStr
	}
	urlBase, err := url.Parse(urlBaseStr)
	if err != nil {
		return nil, ErrBadResponse.Wrap(err, "location")
	}
	root.SetURLBase(urlBase)
	return root, nil
}

func requestXml(
	ctx context.Context,
	url string,
	defaultSpace string,
	doc interface{},
) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ErrBadResponseStatus.New(resp.StatusCode, resp.Status)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.DefaultSpace = defaultSpace
	decoder.CharsetReader = charset.NewReaderLabel

	return decoder.Decode(doc)
}
