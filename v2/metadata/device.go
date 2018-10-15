// Package metadata contains structures for the XML UPnP metadata.
package metadata

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/huin/goupnp/v2/errkind"
)

const (
	DeviceXMLNamespace = "urn:schemas-upnp-org:device-1-0"
)

// RootDevice is the device description as described by section 2.3 "Device
// description" in
// http://upnp.org/specs/arch/UPnP-arch-DeviceArchitecture-v1.1.pdf
type RootDevice struct {
	XMLName     xml.Name    `xml:"root"`
	SpecVersion SpecVersion `xml:"specVersion"`
	URLBase     url.URL     `xml:"-"`
	URLBaseStr  string      `xml:"URLBase"`
	Device      Device      `xml:"device"`
}

// RequestRootDevice retrieves RootDevice XML metadata from the given URL.
func RequestRootDevice(ctx context.Context, loc *url.URL) (*RootDevice, error) {
	root := new(RootDevice)
	if err := requestXml(ctx, loc, DeviceXMLNamespace, root); err != nil {
		return nil, errkind.URLContext.Wrap(err, loc.String())
	}
	var urlBase *url.URL
	if root.URLBaseStr != "" {
		var err error
		urlBase, err = url.Parse(root.URLBaseStr)
		if err != nil {
			return nil, errkind.Wrap(errkind.BadData, err, "bad URL base")
		}
	} else {
		urlBase = loc
	}
	root.SetURLBase(urlBase)
	return root, nil
}

// SetURLBase sets the URLBase for the RootDevice and its underlying components.
func (root *RootDevice) SetURLBase(urlBase *url.URL) {
	root.URLBase = *urlBase
	root.URLBaseStr = urlBase.String()
	root.Device.SetURLBase(urlBase)
}

// SpecVersion is part of a RootDevice or SCPD document, describes the version
// of the specification that the data adheres to.
type SpecVersion struct {
	Major int32 `xml:"major"`
	Minor int32 `xml:"minor"`
}

// Device is a UPnP device. It can have child devices.
type Device struct {
	DeviceType       string    `xml:"deviceType"`
	FriendlyName     string    `xml:"friendlyName"`
	Manufacturer     string    `xml:"manufacturer"`
	ManufacturerURL  URLField  `xml:"manufacturerURL"`
	ModelDescription string    `xml:"modelDescription"`
	ModelName        string    `xml:"modelName"`
	ModelNumber      string    `xml:"modelNumber"`
	ModelURL         URLField  `xml:"modelURL"`
	SerialNumber     string    `xml:"serialNumber"`
	UDN              string    `xml:"UDN"`
	UPC              string    `xml:"UPC,omitempty"`
	Icons            []Icon    `xml:"iconList>icon,omitempty"`
	Services         []Service `xml:"serviceList>service,omitempty"`
	Devices          []Device  `xml:"deviceList>device,omitempty"`

	// Extra observed elements:
	PresentationURL URLField `xml:"presentationURL"`
}

// VisitDevices calls visitor for the device, and all its descendent devices.
func (device *Device) VisitDevices(visitor func(*Device)) {
	visitor(device)
	for i := range device.Devices {
		device.Devices[i].VisitDevices(visitor)
	}
}

// VisitServices calls visitor for all Services under the device and all its
// descendent devices.
func (device *Device) VisitServices(visitor func(*Service)) {
	device.VisitDevices(func(d *Device) {
		for i := range d.Services {
			visitor(&d.Services[i])
		}
	})
}

// FindService finds all (if any) Services under the device and its descendents
// that have the given ServiceType.
func (device *Device) FindService(serviceType string) []*Service {
	var services []*Service
	device.VisitServices(func(s *Service) {
		if s.ServiceType == serviceType {
			services = append(services, s)
		}
	})
	return services
}

// SetURLBase sets the URLBase for the Device and its underlying components.
func (device *Device) SetURLBase(urlBase *url.URL) {
	device.ManufacturerURL.SetURLBase(urlBase)
	device.ModelURL.SetURLBase(urlBase)
	device.PresentationURL.SetURLBase(urlBase)
	for i := range device.Icons {
		device.Icons[i].SetURLBase(urlBase)
	}
	for i := range device.Services {
		device.Services[i].SetURLBase(urlBase)
	}
	for i := range device.Devices {
		device.Devices[i].SetURLBase(urlBase)
	}
}

func (device *Device) String() string {
	return fmt.Sprintf("Device ID %s : %s (%s)", device.UDN, device.DeviceType, device.FriendlyName)
}

// Icon is a representative image that a device might include in its
// description.
type Icon struct {
	Mimetype string   `xml:"mimetype"`
	Width    int32    `xml:"width"`
	Height   int32    `xml:"height"`
	Depth    int32    `xml:"depth"`
	URL      URLField `xml:"url"`
}

// SetURLBase sets the URLBase for the Icon.
func (icon *Icon) SetURLBase(url *url.URL) {
	icon.URL.SetURLBase(url)
}

// URLField is a URL that is part of a device description.
type URLField struct {
	URL url.URL `xml:"-"`
	Ok  bool    `xml:"-"`
	Str string  `xml:",chardata"`
}

func (uf *URLField) SetURLBase(urlBase *url.URL) {
	refUrl, err := url.Parse(uf.Str)
	if err != nil {
		uf.URL = url.URL{}
		uf.Ok = false
		return
	}

	uf.URL = *urlBase.ResolveReference(refUrl)
	uf.Ok = true
}
