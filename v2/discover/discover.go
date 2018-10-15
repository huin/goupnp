// Package discover contains code to discover UPnP devices.
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
