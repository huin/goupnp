// goupnp is an implementation of a client for UPnP devices.
package goupnp

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// Non-exhaustive set of UPnP service types.
const (
	ServiceTypeLayer3Forwarding         = "urn:schemas-upnp-org:service:Layer3Forwarding:1"
	ServiceTypeWANCommonInterfaceConfig = "urn:schemas-upnp-org:WANCommonInterfaceConfig:1"
	// WANPPPConnection is typically useful with regard to the external IP and
	// port forwarding.
	// http://upnp.org/specs/gw/UPnP-gw-WANPPPConnection-v1-Service.pdf
	ServiceTypeWANPPPConnection = "urn:schemas-upnp-org:WANPPPConnection:1"
)

// Non-exhaustive set of UPnP device types.
const (
	// Device type for InternetGatewayDevice.
	// http://upnp.org/specs/gw/upnp-gw-internetgatewaydevice-v1-device.pdf
	DeviceTypeInternetGatewayDevice = "urn:schemas-upnp-org:device:InternetGatewayDevice:1"
)

type ContextError struct {
	Context string
	Err     error
}

func (err ContextError) Error() string {
	return fmt.Sprintf("%s: %v", err.Context, err.Err)
}

type MaybeRootDevice struct {
	Root *RootDevice
	Err  error
}

// DiscoverDevices attempts to find targets of the given type. searchTarget is
// typically a value from a DeviceType* constant. An error is returned for
// errors while attempting to send the query. An error or RootDevice is
// returned for each discovered service.
func DiscoverDevices(searchTarget string) ([]MaybeRootDevice, error) {
	httpu, err := NewHTTPUClient()
	if err != nil {
		return nil, err
	}
	defer httpu.Close()
	responses, err := SSDPRawSearch(httpu, string(searchTarget), 2, 3)
	if err != nil {
		return nil, err
	}

	results := make([]MaybeRootDevice, len(responses))
	for i, response := range responses {
		maybe := &results[i]
		loc, err := response.Location()
		if err != nil {
			maybe.Err = ContextError{"unexpected bad location from search", err}
			continue
		}
		locStr := loc.String()
		root := new(RootDevice)
		if err := requestXml(locStr, DeviceXMLNamespace, root); err != nil {
			maybe.Err = ContextError{fmt.Sprintf("error requesting root device details from %q", locStr), err}
			continue
		}
		urlBase, err := url.Parse(root.URLBaseStr)
		if err != nil {
			maybe.Err = ContextError{fmt.Sprintf("error parsing URLBase %q from %q: %v", root.URLBaseStr, locStr), err}
			continue
		}
		root.SetURLBase(urlBase)
		maybe.Root = root
	}

	return results, nil
}

func requestXml(url string, defaultSpace string, doc interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("goupnp: got response status %s from %q",
			resp.Status, url)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.DefaultSpace = defaultSpace

	return decoder.Decode(doc)
}
