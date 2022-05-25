// Package scpd contains data structures that represent an SCPD at a higher level than XML.
package scpd

import (
	"errors"
	"fmt"
	"sort"

	"github.com/huin/goupnp/v2alpha/description/xmlscpd"
)

var (
	BadDescriptionError         = errors.New("bad XML description")
	UnsupportedDescriptionError = errors.New("unsupported XML description")
)

// SCPD is the top level service description.
type SCPD struct {
	actionByName   map[string]*Action
	variableByName map[string]*StateVariable
}

// FromXML creates an SCPD from XML data.
//
// It assumes that xmlDesc.Clean() has been called.
func FromXML(xmlDesc *xmlscpd.SCPD) (*SCPD, error) {
	stateVariables := make(map[string]*StateVariable, len(xmlDesc.StateVariables))
	for _, xmlSV := range xmlDesc.StateVariables {
		sv, err := stateVariableFromXML(xmlSV)
		if err != nil {
			return nil, fmt.Errorf("processing state variable %q: %w", xmlSV.Name, err)
		}
		if _, exists := stateVariables[sv.name]; exists {
			return nil, fmt.Errorf("%w: multiple state variables with name %q",
				BadDescriptionError, sv.name)
		}
		stateVariables[sv.name] = sv
	}
	actions := make(map[string]*Action, len(xmlDesc.Actions))
	for _, xmlAction := range xmlDesc.Actions {
		action, err := actionFromXML(xmlAction)
		if err != nil {
			return nil, fmt.Errorf("processing action %q: %w", xmlAction.Name, err)
		}
		if _, exists := actions[action.name]; exists {
			return nil, fmt.Errorf("%w: multiple actions with name %q",
				BadDescriptionError, action.name)
		}
		actions[action.name] = action
	}
	return &SCPD{
		actionByName:   actions,
		variableByName: stateVariables,
	}, nil
}

// ActionNames returns the ordered names of each action.
func (scpd *SCPD) ActionNames() []string {
	names := make([]string, 0, len(scpd.actionByName))
	for name := range scpd.actionByName {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Action returns an action with the given name.
func (scpd *SCPD) Action(name string) *Action {
	return scpd.actionByName[name]
}

// Variable returns a state variable with the given name.
func (scpd *SCPD) Variable(name string) *StateVariable {
	return scpd.variableByName[name]
}

// Action describes a single UPnP SOAP action.
type Action struct {
	name    string
	inArgs  []*Argument
	outArgs []*Argument
}

// actionFromXML creates an Action from the given XML description.
func actionFromXML(xmlAction *xmlscpd.Action) (*Action, error) {
	if xmlAction.Name == "" {
		return nil, fmt.Errorf("%w: empty action name", BadDescriptionError)
	}
	var inArgs []*Argument
	var outArgs []*Argument
	for _, xmlArg := range xmlAction.Arguments {
		arg, err := argumentFromXML(xmlArg)
		if err != nil {
			return nil, fmt.Errorf("processing argument %q: %w", xmlArg.Name, err)
		}
		switch xmlArg.Direction {
		case "in":
			inArgs = append(inArgs, arg)
		case "out":
			outArgs = append(outArgs, arg)
		default:
			return nil, fmt.Errorf("%w: argument %q has invalid direction %q",
				BadDescriptionError, xmlArg.Name, xmlArg.Direction)
		}
	}
	return &Action{
		name:    xmlAction.Name,
		inArgs:  inArgs,
		outArgs: outArgs,
	}, nil
}

// Argument description data.
type Argument struct {
	name                 string
	relatedStateVariable string
}

// argumentFromXML creates an Argument from the XML description.
func argumentFromXML(xmlArg *xmlscpd.Argument) (*Argument, error) {
	if xmlArg.Name == "" {
		return nil, fmt.Errorf("%w: empty argument name", BadDescriptionError)
	}
	if xmlArg.RelatedStateVariable == "" {
		return nil, fmt.Errorf("%w: empty related state variable", BadDescriptionError)
	}
	return &Argument{
		name:                 xmlArg.Name,
		relatedStateVariable: xmlArg.RelatedStateVariable,
	}, nil
}

func (arg *Argument) Name() string {
	return arg.name
}

func (arg *Argument) RelatedStateVariableName() string {
	return arg.relatedStateVariable
}

// StateVariable description data.
type StateVariable struct {
	name     string
	dataType string
}

func stateVariableFromXML(xmlSV *xmlscpd.StateVariable) (*StateVariable, error) {
	if xmlSV.Name == "" {
		return nil, fmt.Errorf("%w: empty state variable name", BadDescriptionError)
	}
	if xmlSV.DataType.Type != "" {
		return nil, fmt.Errorf("%w: unsupported data type %q",
			UnsupportedDescriptionError, xmlSV.DataType.Type)
	}
	return &StateVariable{
		name:     xmlSV.Name,
		dataType: xmlSV.DataType.Name,
	}, nil
}

func (sv *StateVariable) Name() string {
	return sv.name
}

func (sv *StateVariable) DataTypeName() string {
	return sv.dataType
}
