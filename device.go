// This file contains XML structures for communicating with UPnP devices.

package goupnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
)

const (
	DeviceXMLNamespace = "urn:schemas-upnp-org:device-1-0"
)

// RootDevice is the device description as described by section 2.3 "Device
// description" in
// http://upnp.org/specs/arch/UPnP-arch-DeviceArchitecture-v1.1.pdf
type RootDevice struct {
	Name        xml.Name    `xml:"root`
	SpecVersion SpecVersion `xml:"specVersion"`
	URLBase     url.URL     `xml:"-"`
	URLBaseStr  string      `xml:"URLBase"`
	Device      Device      `xml:"device"`
}

func (root *RootDevice) SetURLBase(urlBase *url.URL) {
	root.URLBase = *urlBase
	root.URLBaseStr = urlBase.String()
	root.Device.SetURLBase(urlBase)
}

type SpecVersion struct {
	Major int32 `xml:"major"`
	Minor int32 `xml:"minor"`
}

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

func (device *Device) SetURLBase(urlBase *url.URL) {
	device.ManufacturerURL.SetURLBase(urlBase)
	device.ModelURL.SetURLBase(urlBase)
	device.PresentationURL.SetURLBase(urlBase)
	for _, icon := range device.Icons {
		icon.SetURLBase(urlBase)
	}
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
	Mimetype string   `xml:"mimetype"`
	Width    int32    `xml:"width"`
	Height   int32    `xml:"height"`
	Depth    int32    `xml:"depth"`
	URL      URLField `xml:"url"`
}

func (icon *Icon) SetURLBase(url *url.URL) {
	icon.URL.SetURLBase(url)
}

type Service struct {
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	SCPDURL     URLField `xml:"SCPDURL"`
	ControlURL  URLField `xml:"controlURL"`
	EventSubURL URLField `xml:"eventSubURL"`
}

func (srv *Service) SetURLBase(urlBase *url.URL) {
	srv.SCPDURL.SetURLBase(urlBase)
	srv.ControlURL.SetURLBase(urlBase)
	srv.EventSubURL.SetURLBase(urlBase)
}

func (srv *Service) String() string {
	return fmt.Sprintf("Service ID %s : %s", srv.ServiceId, srv.ServiceType)
}

func (srv *Service) RequestSCDP() (*SCPD, error) {
	if !srv.SCPDURL.Ok {
		return nil, errors.New("bad/missing SCPD URL, or no URLBase has been set")
	}
	scpd := new(SCPD)
	if err := requestXml(srv.SCPDURL.URL.String(), SCPDXMLNamespace, scpd); err != nil {
		return nil, err
	}
	return scpd, nil
}

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
