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
	"net/url"

	"github.com/huin/goupnp/v2/errkind"
	"github.com/huin/goupnp/v2/httpu"
	"github.com/huin/goupnp/v2/metadata"
	"github.com/huin/goupnp/v2/ssdp"
)

// MaybeRootDevice contains either a RootDevice or an error.
type MaybeRootDevice struct {
	// Set iff Err == nil.
	Root *metadata.RootDevice

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
			maybe.Err = errkind.Wrap(errkind.BadData, err, "response location")
			continue
		}
		maybe.Location = loc
		if root, err := metadata.RequestRootDevice(ctx, loc); err != nil {
			maybe.Err = err
		} else {
			maybe.Root = root
		}
	}

	return results, nil
}
