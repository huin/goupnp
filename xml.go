// This file contains XML structures for communicating with UPnP devices.

package goupnp

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

const (
	DeviceXMLNamespace = "urn:schemas-upnp-org:device-1-0"
)

type RootDevice struct {
	URLBase     *url.URL    `xml:"-"`
	Name        xml.Name    `xml:"root`
	SpecVersion SpecVersion `xml:"specVersion"`
	URLBaseStr  string      `xml:"URLBase"`
	Device      *Device     `xml:"device"`
}

func (root *RootDevice) SetURLBase(urlBase *url.URL) {
	root.URLBase = urlBase
	root.URLBaseStr = urlBase.String()
	root.Device.SetURLBase(urlBase)
}

type SpecVersion struct {
	Major int32 `xml:"major"`
	Minor int32 `xml:"minor"`
}

type Device struct {
	URLBase          *url.URL   `xml:"-"`
	DeviceType       string     `xml:"deviceType"`
	FriendlyName     string     `xml:"friendlyName"`
	Manufacturer     string     `xml:"manufacturer"`
	ManufacturerURL  string     `xml:"manufacturerURL"`
	ModelDescription string     `xml:"modelDescription"`
	ModelName        string     `xml:"modelName"`
	ModelNumber      string     `xml:"modelNumber"`
	ModelURL         string     `xml:"modelURL"`
	SerialNumber     string     `xml:"serialNumber"`
	UDN              string     `xml:"UDN"`
	UPC              string     `xml:"UPC,omitempty"`
	Icons            []*Icon    `xml:"iconList>icon,omitempty"`
	Services         []*Service `xml:"serviceList>service,omitempty"`
	Devices          []*Device  `xml:"deviceList>device,omitempty"`

	// Extra observed elements:
	PresentationURL string `xml:"presentationURL"`
}

func (device *Device) SetURLBase(urlBase *url.URL) {
	device.URLBase = urlBase
	for _, srv := range device.Services {
		srv.SetURLBase(urlBase)
	}
	for _, child := range device.Devices {
		child.SetURLBase(urlBase)
	}
}

func (device *Device) String() string {
	return fmt.Sprintf("Device ID %s : %s (%s)", device.UDN, device.DeviceType, device.FriendlyName)
}

type Icon struct {
	Mimetype string `xml:"mimetype"`
	Width    int32  `xml:"width"`
	Height   int32  `xml:"height"`
	Depth    int32  `xml:"depth"`
	URL      string `xml:"url"`
}

type Service struct {
	URLBase     *url.URL `xml:"-"`
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	SCPDURL     string   `xml:"SCPDURL"`
	ControlURL  string   `xml:"controlURL"`
	EventSubURL string   `xml:"eventSubURL"`
}

func (srv *Service) SetURLBase(urlBase *url.URL) {
	srv.URLBase = urlBase
}

func (srv *Service) String() string {
	return fmt.Sprintf("Service ID %s : %s", srv.ServiceId, srv.ServiceType)
}
