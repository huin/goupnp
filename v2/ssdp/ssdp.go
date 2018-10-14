// package ssdp implements a client for the Simple Service Discovery Protocol,
// typically used to discover available UPnP devices on the network.
//
// See also: https://en.wikipedia.org/wiki/Simple_Service_Discovery_Protocol
package ssdp

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/huin/goupnp/v2/httpu"
)

const (
	ssdpDiscover   = `"ssdp:discover"`
	ntsAlive       = `ssdp:alive`
	ntsByebye      = `ssdp:byebye`
	ntsUpdate      = `ssdp:update`
	ssdpUDP4Addr   = "239.255.255.250:1900"
	ssdpSearchPort = 1900
	methodSearch   = "M-SEARCH"
	methodNotify   = "NOTIFY"

	// SSDPAll is a value for searchTarget that searches for all devices and
	// services.
	SSDPAll = "ssdp:all"

	// UPNPRootDevice is a value for searchTarget that searches for all root
	// devices.
	UPNPRootDevice = "upnp:rootdevice"
)

// RawSearch performs a fairly raw SSDP search request, and returns the unique
// response(s) that it receives. Each response has the requested searchTarget, a
// USN, and a valid location.
func RawSearch(
	ctx context.Context,
	hc HTTPUClient,
	searchTarget string,
	options ...SearchOption,
) ([]*http.Response, error) {
	ss := &searchSettings{
		waitFor: 3 * time.Second,
	}
	ss.applyOptions(options)

	seenUsns := make(map[string]bool)
	var responses []*http.Response
	req := &http.Request{
		Method: methodSearch,
		// TODO: Support both IPv4 and IPv6.
		Host: ssdpUDP4Addr,
		URL:  &url.URL{Opaque: "*"},
		Header: http.Header{
			// Putting headers in here avoids them being title-cased.
			// (The UPnP discovery protocol uses case-sensitive headers)
			"HOST": []string{ssdpUDP4Addr},
			"MX": []string{
				strconv.FormatInt(int64(ss.waitFor/time.Second), 10),
			},
			"MAN": []string{ssdpDiscover},
			"ST":  []string{searchTarget},
		},
	}
	req = req.WithContext(ctx)
	ss.hrs = append(ss.hrs, httpu.WaitFor(ss.waitFor))
	allResponses, err := hc.Do(req, ss.hrs...)
	if err != nil {
		return nil, err
	}

	isExactSearch := searchTarget != SSDPAll && searchTarget != UPNPRootDevice

	for _, response := range allResponses {
		if response.StatusCode != 200 {
			log.Printf(
				"ssdp: got response status code %q in search response",
				response.Status,
			)
			continue
		}
		st := response.Header.Get("ST")
		if isExactSearch && st != searchTarget {
			continue
		}
		location, err := response.Location()
		if err != nil {
			log.Printf(
				"ssdp: no usable location in search response (discarding): %v",
				err,
			)
			continue
		}
		usn := response.Header.Get("USN")
		if usn == "" {
			log.Printf(
				"ssdp: empty/missing USN in search response "+
					"(using location instead): %v",
				err,
			)
			usn = location.String()
		}
		if _, alreadySeen := seenUsns[usn]; !alreadySeen {
			seenUsns[usn] = true
			responses = append(responses, response)
		}
	}

	return responses, nil
}

type searchSettings struct {
	hrs     []httpu.RequestOption
	waitFor time.Duration
}

func (ss *searchSettings) applyOptions(options []SearchOption) {
	for _, o := range options {
		o(ss)
	}
}

type SearchOption func(*searchSettings)

// NumSends controls how many redundant requests to send.
func NumSends(numSends int) SearchOption {
	return func(ss *searchSettings) {
		ss.hrs = append(ss.hrs, httpu.NumSends(numSends))
	}
}

// WaitFor controls how long to wait for HTTPU responses.
func WaitFor(d time.Duration) SearchOption {
	return func(ss *searchSettings) {
		ss.waitFor = d
	}
}

// HTTPUClient is the interface required for SSDP searches. Typically
// implemented by *httpu.HTTPUClient.
type HTTPUClient interface {
	Do(
		req *http.Request,
		options ...httpu.RequestOption,
	) ([]*http.Response, error)
}
