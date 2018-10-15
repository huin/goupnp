package metadata

import (
	"context"
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/huin/goupnp/v2/errkind"
)

const (
	// SCPDXMLNamespace is the XML namespace for SCPD.
	SCPDXMLNamespace = "urn:schemas-upnp-org:service-1-0"
)

func cleanWhitespace(s *string) {
	*s = strings.TrimSpace(*s)
}

// SCPD is the SOAP service description as described by section 2.5 "Service
// description" in
// http://upnp.org/specs/arch/UPnP-arch-DeviceArchitecture-v1.1.pdf
type SCPD struct {
	XMLName        xml.Name        `xml:"scpd"`
	ConfigID       string          `xml:"configId,attr"`
	SpecVersion    SpecVersion     `xml:"specVersion"`
	Actions        []Action        `xml:"actionList>action"`
	StateVariables []StateVariable `xml:"serviceStateTable>stateVariable"`
}

// RequestSCPD requests the SCPD (SOAP actions and state variables description)
// for the service.
func RequestSCPD(ctx context.Context, loc *url.URL) (*SCPD, error) {
	s := new(SCPD)
	if err := requestXML(ctx, loc, SCPDXMLNamespace, s); err != nil {
		return nil, errkind.URLContext.Wrap(err, loc.String())
	}
	return s, nil
}

// Clean attempts to remove stray whitespace etc. in the structure. It seems
// unfortunately common for stray whitespace to be present in SCPD documents,
// this method attempts to make it easy to clean them out.
func (scpd *SCPD) Clean() {
	cleanWhitespace(&scpd.ConfigID)
	for i := range scpd.Actions {
		scpd.Actions[i].clean()
	}
	for i := range scpd.StateVariables {
		scpd.StateVariables[i].clean()
	}
}

// GetStateVariable returns a state variable of the given name, or nil if not
// found.
func (scpd *SCPD) GetStateVariable(variable string) *StateVariable {
	for i := range scpd.StateVariables {
		v := &scpd.StateVariables[i]
		if v.Name == variable {
			return v
		}
	}
	return nil
}

// GetAction returns a action of the given name, or nil if not found.
func (scpd *SCPD) GetAction(action string) *Action {
	for i := range scpd.Actions {
		a := &scpd.Actions[i]
		if a.Name == action {
			return a
		}
	}
	return nil
}

// Action describes a UPnP SOAP action.
type Action struct {
	Name      string     `xml:"name"`
	Arguments []Argument `xml:"argumentList>argument"`
}

func (action *Action) clean() {
	cleanWhitespace(&action.Name)
	for i := range action.Arguments {
		action.Arguments[i].clean()
	}
}

// InputArguments returns the set of input arguments.
func (action *Action) InputArguments() []*Argument {
	var result []*Argument
	for i := range action.Arguments {
		arg := &action.Arguments[i]
		if arg.IsInput() {
			result = append(result, arg)
		}
	}
	return result
}

// OutputArguments returns the set of output arguments.
func (action *Action) OutputArguments() []*Argument {
	var result []*Argument
	for i := range action.Arguments {
		arg := &action.Arguments[i]
		if arg.IsOutput() {
			result = append(result, arg)
		}
	}
	return result
}

// Argument describes a UPnP action argument.
type Argument struct {
	Name                 string `xml:"name"`
	Direction            string `xml:"direction"`            // in|out
	RelatedStateVariable string `xml:"relatedStateVariable"` // ?
	Retval               string `xml:"retval"`               // ?
}

func (arg *Argument) clean() {
	cleanWhitespace(&arg.Name)
	cleanWhitespace(&arg.Direction)
	cleanWhitespace(&arg.RelatedStateVariable)
	cleanWhitespace(&arg.Retval)
}

// IsInput returns true iff the argument is an input argument.
func (arg *Argument) IsInput() bool {
	return arg.Direction == "in"
}

// IsOutput returns true iff the argument is an output argument.
func (arg *Argument) IsOutput() bool {
	return arg.Direction == "out"
}

// StateVariable describes a UPnP state variable.
type StateVariable struct {
	Name              string             `xml:"name"`
	SendEvents        string             `xml:"sendEvents,attr"` // yes|no
	Multicast         string             `xml:"multicast,attr"`  // yes|no
	DataType          DataType           `xml:"dataType"`
	DefaultValue      string             `xml:"defaultValue"`
	AllowedValueRange *AllowedValueRange `xml:"allowedValueRange"`
	AllowedValues     []string           `xml:"allowedValueList>allowedValue"`
}

func (v *StateVariable) clean() {
	cleanWhitespace(&v.Name)
	cleanWhitespace(&v.SendEvents)
	cleanWhitespace(&v.Multicast)
	v.DataType.clean()
	cleanWhitespace(&v.DefaultValue)
	if v.AllowedValueRange != nil {
		v.AllowedValueRange.clean()
	}
	for i := range v.AllowedValues {
		cleanWhitespace(&v.AllowedValues[i])
	}
}

// AllowedValueRange indicates the minimum and maximum values for a variable.
type AllowedValueRange struct {
	Minimum string `xml:"minimum"`
	Maximum string `xml:"maximum"`
	Step    string `xml:"step"`
}

func (r *AllowedValueRange) clean() {
	cleanWhitespace(&r.Minimum)
	cleanWhitespace(&r.Maximum)
	cleanWhitespace(&r.Step)
}

// DataType indicates the SOAP data type for a variable.
type DataType struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

func (dt *DataType) clean() {
	cleanWhitespace(&dt.Name)
	cleanWhitespace(&dt.Type)
}
