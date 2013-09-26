// This file contains XML structures for communicating with UPnP devices.

package goupnp

import (
	"encoding/xml"
)

const (
	deviceXmlNs = "urn:schemas-upnp-org:device-1-0"
)

type xmlRootDevice struct {
	Name xml.Name `xml:"root`
	SpecVersion xmlSpecVersion `xml:"specVersion"`
	URLBase string `xml:"URLBase"`
	Device xmlDevice `xml:"device"`
}

type xmlSpecVersion struct {
	Major int32 `xml:"major"`
	Minor int32 `xml:"minor"`
}

type xmlDevice struct {
	DeviceType string `xml:"deviceType"`
	FriendlyName string `xml:"friendlyName"`
	Manufacturer string `xml:"manufacturer"`
	ManufacturerURL string `xml:"manufacturerURL"`
	ModelDescription string `xml:"modelDescription"`
	ModelName string `xml:"modelName"`
	ModelNumber string `xml:"modelNumber"`
	ModelURL string `xml:"modelURL"`
	SerialNumber string `xml:"serialNumber"`
	UDN string `xml:"UDN"`
	UPC string `xml:"UPC,omitempty"`
	Icons []xmlIcon `xml:"iconList>icon,omitempty"`
	Services []xmlService `xml:"serviceList>service,omitempty"`
	Devices []xmlDevice `xml:"deviceList>device,omitempty"`

	// Extra observed elements:
	PresentationURL string `xml:"presentationURL"`
}

type xmlIcon struct {
	Mimetype string `xml:"mimetype"`
	Width int32 `xml:"width"`
	Height int32 `xml:"height"`
	Depth int32 `xml:"depth"`
	URL string `xml:"url"`
}

type xmlService struct {
	ServiceType string `xml:"serviceType"`
	ServiceId string `xml:"serviceId"`
	SCPDURL string `xml:"SCPDURL"`
	ControlURL string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
}
