package goupnp

import (
	"encoding/xml"
)

const (
	SCPDXMLNamespace = "urn:schemas-upnp-org:service-1-0"
)

// SCPD is the service description as described by section 2.5 "Service
// description" in
// http://upnp.org/specs/arch/UPnP-arch-DeviceArchitecture-v1.1.pdf
type SCPD struct {
	XMLName        xml.Name        `xml:"scpd"`
	ConfigId       string          `xml:"configId,attr"`
	SpecVersion    SpecVersion     `xml:"specVersion"`
	Actions        []Action        `xml:"actionList>action"`
	StateVariables []StateVariable `xml:"serviceStateTable>stateVariable"`
}

type Action struct {
	Name      string     `xml:"name"`
	Arguments []Argument `xml:"argumentList>argument"`
}

type Argument struct {
	Name                 string `xml:"name"`
	Direction            string `xml:"direction"`            // in|out
	RelatedStateVariable string `xml:"relatedStateVariable"` // ?
	Retval               string `xml:"retval"`               // ?
}

type StateVariable struct {
	Name              string            `xml:"name"`
	SendEvents        string            `xml:"sendEvents,attr"` // yes|no
	Multicast         string            `xml:"multicast,attr"`  // yes|no
	DataType          DataType          `xml:"dataType"`
	DefaultValue      string            `xml:"defaultValue"`
	AllowedValueRange AllowedValueRange `xml:"allowedValueRange"`
	AllowedValue      []string          `xml:"allowedValueList>allowedValue"`
}

type AllowedValueRange struct {
	Minimum string `xml:"minimum"`
	Maximum string `xml:"maximum"`
	Step    string `xml:"step"`
}

type DataType struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}
