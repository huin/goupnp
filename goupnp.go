package goupnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"io"
	"os"
	"net/http"
	"net/url"
)

const (
	ssdpUDP4Addr = "239.255.255.250:1900"

	methodSearch = "M-SEARCH"
	// Search Target for InternetGatewayDevice.
	stIgd  = "urn:schemas-upnp-org:device:InternetGatewayDevice:1"
	hdrMan = `"ssdp:discover"`
)

// DiscoverIGD attempts to find Internet Gateway Devices.
//
// TODO: Fix implementation to discover multiple. Currently it will find a
// maximum of one.
func DiscoverIGD() ([]IGD, error) {
	hc := http.Client{
		Transport: udpRoundTripper{},
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return errors.New("goupnp: unexpected HTTP redirect")
		},
		Jar: nil,
	}

	request := http.Request{
		Method: methodSearch,
		// TODO: Support both IPv4 and IPv6.
		Host: ssdpUDP4Addr,
		URL:  &url.URL{Opaque: "*"},
		Header: http.Header{
			// Putting headers in here avoids them being title-cased.
			// (The UPnP discovery protocol uses case-sensitive headers)
			"HOST": []string{ssdpUDP4Addr},
			"MX":   []string{"2"}, // TODO: Variable max wait time.
			"MAN":  []string{hdrMan},
			"ST":   []string{stIgd},
		},
	}

	response, err := hc.Do(&request)
	if err != nil {
		return nil, err
	}

	// Any errors past this point are simply "no result found". We log the
	// errors, but report no results. In a future version of this implementation,
	// multiple *good* results can be returned.

	if response.StatusCode != 200 {
		log.Printf("goupnp: response code %d %q from UPnP discovery",
			response.StatusCode, response.Status)
		return nil, nil
	}
	if st := response.Header.Get("ST"); st != stIgd {
		log.Printf("goupnp: got unexpected search target result %q", st)
		return nil, nil
	}

	location, err := response.Location()
	if err != nil {
		log.Printf("goupnp: missing location in response")
		return nil, nil
	}

	igd, err := requestIgd(location.String())
	if err != nil {
		log.Printf("goupnp: error requesting IGD: %v", err)
		return nil, nil
	}

	return []IGD{igd}, nil
}

// IGD defines the interface for an Internet Gateway Device.
type IGD interface {
}

type igd struct {
	serviceUrl string
}

func requestIgd(serviceUrl string) (IGD, error) {
	resp, err := http.Get(serviceUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(io.TeeReader(resp.Body, os.Stdout))
	decoder.DefaultSpace = deviceXmlNs
	var root xmlRootDevice
	if err = decoder.Decode(&root); err != nil {
		return nil, err
	}
	log.Printf("%+v", root)

	return igd{serviceUrl}, nil
}

func (device *igd) String() string {
	return fmt.Sprintf("goupnp.IGD @ %s", device.serviceUrl)
}
