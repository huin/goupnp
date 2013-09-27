package goupnp

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

const (
	// Search Target for InternetGatewayDevice.
	SearchTargetIGD = "urn:schemas-upnp-org:device:InternetGatewayDevice:1"
)

// DiscoverIGD attempts to find Internet Gateway Devices.
//
// TODO: Fix implementation to discover multiple. Currently it will find a
// maximum of one.
func DiscoverIGD() ([]*IGD, error) {
	httpu, err := NewHTTPUClient()
	if err != nil {
		return nil, err
	}
	responses, err := SSDPRawSearch(httpu, SearchTargetIGD, 2, 3)

	results := make([]*IGD, 0, len(responses))
	for _, response := range responses {
		loc, err := response.Location()
		if err != nil {
			log.Printf("goupnp: unexpected bad location from search: %v", err)
			continue
		}
		igd, err := requestIgd(loc.String())
		if err != nil {
			log.Printf("goupnp: error requesting IGD: %v", err)
			continue
		}
		results = append(results, igd)
	}

	return results, nil
}

// IGD defines the interface for an Internet Gateway Device.
type IGD struct {
	xml xmlRootDevice
}

func requestIgd(serviceUrl string) (*IGD, error) {
	resp, err := http.Get(serviceUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("goupnp: got response status %s from IGD at %q",
			resp.Status, serviceUrl)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.DefaultSpace = deviceXmlNs
	var xml xmlRootDevice
	if err = decoder.Decode(&xml); err != nil {
		return nil, err
	}
	log.Printf("%+v", xml)

	return &IGD{xml}, nil
}

func (igd *IGD) Device() *Device {
	return &Device{
		igd.xml.URLBase,
		igd.xml.Device,
	}
}

func (igd *IGD) String() string {
	return fmt.Sprintf("IGD{UDN: %q friendlyName: %q}",
		igd.xml.Device.UDN, igd.xml.Device.FriendlyName)
}

type Device struct {
	urlBase string
	xml     xmlDevice
}

func (device *Device) String() string {
	return fmt.Sprintf("Device{friendlyName: %q}", device.xml.FriendlyName)
}

func (device *Device) Devices() []*Device {
	devices := make([]*Device, len(device.xml.Devices))
	for i, childXml := range device.xml.Devices {
		devices[i] = &Device{
			device.urlBase,
			childXml,
		}
	}
	return devices
}

func (device *Device) Services() []*Service {
	srvs := make([]*Service, len(device.xml.Services))
	for i, childXml := range device.xml.Services {
		srvs[i] = &Service{
			device.urlBase,
			childXml,
		}
	}
	return srvs
}

type Service struct {
	urlBase string
	xml     xmlService
}

func (srv *Service) String() string {
	return fmt.Sprintf("Service{serviceType: %q}", srv.xml.ServiceType)
}
