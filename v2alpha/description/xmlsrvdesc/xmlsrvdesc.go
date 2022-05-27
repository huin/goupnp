// Package xmlsrvdesc contains the XML data structures used in SCPD (Service Control Protocol
// Description).
//
// Described in section 2.5 of
// https://openconnectivity.org/upnp-specs/UPnP-arch-DeviceArchitecture-v2.0-20200417.pdf.
package xmlsrvdesc

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

func cleanWhitespace(s *string) {
	*s = strings.TrimSpace(*s)
}

// SCPD is the top level XML service description.
type SCPD struct {
	XMLName        xml.Name         `xml:"scpd"`
	ConfigId       string           `xml:"configId,attr"`
	SpecVersion    SpecVersion      `xml:"specVersion"`
	Actions        []*Action        `xml:"actionList>action"`
	StateVariables []*StateVariable `xml:"serviceStateTable>stateVariable"`
}

// Clean removes stray whitespace in the structure.
//
// It's common for stray whitespace to be present in SCPD documents, this method removes them
// in-place.
func (scpd *SCPD) Clean() {
	cleanWhitespace(&scpd.ConfigId)
	for i := range scpd.Actions {
		scpd.Actions[i].Clean()
	}
	for i := range scpd.StateVariables {
		scpd.StateVariables[i].Clean()
	}
}

// SpecVersion is part of a SCPD document, describes the version of the
// specification that the data adheres to.
type SpecVersion struct {
	Major int32 `xml:"major"`
	Minor int32 `xml:"minor"`
}

// Action XML description data.
type Action struct {
	Optional  PresenceBool
	Name      string      `xml:"name"`
	Arguments []*Argument `xml:"argumentList>argument"`
}

// Clean removes stray whitespace in the structure.
func (action *Action) Clean() {
	cleanWhitespace(&action.Name)
	for i := range action.Arguments {
		action.Arguments[i].Clean()
	}
}

// Argument XML data.
type Argument struct {
	Name                 string `xml:"name"`
	Direction            string `xml:"direction"`            // in|out
	RelatedStateVariable string `xml:"relatedStateVariable"` // ?
	Retval               string `xml:"retval"`               // ?
}

// Clean removes stray whitespace in the structure.
func (arg *Argument) Clean() {
	cleanWhitespace(&arg.Name)
	cleanWhitespace(&arg.Direction)
	cleanWhitespace(&arg.RelatedStateVariable)
	cleanWhitespace(&arg.Retval)
}

// StateVariable XML data.
type StateVariable struct {
	Optional          PresenceBool
	Name              string             `xml:"name"`
	SendEvents        string             `xml:"sendEvents,attr"` // yes|no
	Multicast         string             `xml:"multicast,attr"`  // yes|no
	DataType          DataType           `xml:"dataType"`
	DefaultValue      string             `xml:"defaultValue"`
	AllowedValueRange *AllowedValueRange `xml:"allowedValueRange"`
	AllowedValues     []string           `xml:"allowedValueList>allowedValue"`
}

// Clean removes stray whitespace in the structure.
func (v *StateVariable) Clean() {
	cleanWhitespace(&v.Name)
	cleanWhitespace(&v.SendEvents)
	cleanWhitespace(&v.Multicast)
	v.DataType.Clean()
	cleanWhitespace(&v.DefaultValue)
	if v.AllowedValueRange != nil {
		v.AllowedValueRange.Clean()
	}
	for i := range v.AllowedValues {
		cleanWhitespace(&v.AllowedValues[i])
	}
}

// AllowedValueRange XML data.
type AllowedValueRange struct {
	Minimum string `xml:"minimum"`
	Maximum string `xml:"maximum"`
	Step    string `xml:"step"`
}

// Clean removes stray whitespace in the structure.
func (r *AllowedValueRange) Clean() {
	cleanWhitespace(&r.Minimum)
	cleanWhitespace(&r.Maximum)
	cleanWhitespace(&r.Step)
}

// DataType XML data.
type DataType struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// Clean removes stray whitespace in the structure.
func (dt *DataType) Clean() {
	cleanWhitespace(&dt.Name)
	cleanWhitespace(&dt.Type)
}

// PresenceBool represents an empty XML element that is true if present.
//
// Is an error if it contains any attributes or contents.
type PresenceBool bool

func (pb *PresenceBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*pb = true
	if len(start.Attr) > 0 {
		return fmt.Errorf("unexpected attributes on element %s:%s",
			start.Name.Space, start.Name.Local)
	}
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := tok.(type) {
		case xml.CharData:
			if len(bytes.TrimSpace([]byte(tok))) > 0 {
				return fmt.Errorf("unexpected char data on element %s:%s",
					start.Name.Space, start.Name.Local)
			}
		case xml.EndElement:
			return nil
		case xml.Comment:
		default:
			return fmt.Errorf("unexpected %T token on element %s:%s",
				tok, start.Name.Space, start.Name.Local)
		}
	}
}
