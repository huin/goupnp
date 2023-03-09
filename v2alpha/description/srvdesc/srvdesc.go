// Package srvdesc contains data structures that represent an SCPD at a higher level than XML.
package srvdesc

import (
	"errors"
	"fmt"
	"sort"

	"github.com/huin/goupnp/v2alpha/description/xmlsrvdesc"
)

var (
	ErrBadDescription         = errors.New("bad XML description")
	ErrMissingDefinition      = errors.New("missing definition")
	ErrUnsupportedDescription = errors.New("unsupported XML description")
)

// SCPD is the top level service description.
type SCPD struct {
	ActionByName   map[string]*Action
	VariableByName map[string]*StateVariable
}

// FromXML creates an SCPD from XML data.
//
// It assumes that xmlDesc.Clean() has been called.
func FromXML(xmlDesc *xmlsrvdesc.SCPD) (*SCPD, error) {
	scpd := &SCPD{
		ActionByName:   make(map[string]*Action, len(xmlDesc.Actions)),
		VariableByName: make(map[string]*StateVariable, len(xmlDesc.StateVariables)),
	}
	stateVariables := scpd.VariableByName
	for _, xmlSV := range xmlDesc.StateVariables {
		sv, err := stateVariableFromXML(xmlSV)
		if err != nil {
			return nil, fmt.Errorf("processing state variable %q: %w", xmlSV.Name, err)
		}
		if _, exists := stateVariables[sv.Name]; exists {
			return nil, fmt.Errorf("%w: multiple state variables with name %q",
				ErrBadDescription, sv.Name)
		}
		stateVariables[sv.Name] = sv
	}
	actions := scpd.ActionByName
	for _, xmlAction := range xmlDesc.Actions {
		action, err := actionFromXML(xmlAction, scpd)
		if err != nil {
			return nil, fmt.Errorf("processing action %q: %w", xmlAction.Name, err)
		}
		if _, exists := actions[action.Name]; exists {
			return nil, fmt.Errorf("%w: multiple actions with name %q",
				ErrBadDescription, action.Name)
		}
		actions[action.Name] = action
	}
	return scpd, nil
}

// SortedActions returns the actions, in order of name.
func (scpd *SCPD) SortedActions() []*Action {
	actions := make([]*Action, 0, len(scpd.ActionByName))
	for _, a := range scpd.ActionByName {
		actions = append(actions, a)
	}
	sort.Slice(actions, func(i, j int) bool {
		return actions[i].Name < actions[j].Name
	})
	return actions
}

// Action describes a single UPnP SOAP action.
type Action struct {
	SCPD    *SCPD
	Name    string
	InArgs  []*Argument
	OutArgs []*Argument
}

// actionFromXML creates an Action from the given XML description.
func actionFromXML(xmlAction *xmlsrvdesc.Action, scpd *SCPD) (*Action, error) {
	if xmlAction.Name == "" {
		return nil, fmt.Errorf("%w: empty action name", ErrBadDescription)
	}
	action := &Action{
		SCPD: scpd,
		Name: xmlAction.Name,
	}
	var inArgs []*Argument
	var outArgs []*Argument
	for _, xmlArg := range xmlAction.Arguments {
		arg, err := argumentFromXML(xmlArg, action)
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
				ErrBadDescription, xmlArg.Name, xmlArg.Direction)
		}
	}
	action.InArgs = inArgs
	action.OutArgs = outArgs
	return action, nil
}

// Argument description data.
type Argument struct {
	Action                   *Action
	Name                     string
	RelatedStateVariableName string
}

// argumentFromXML creates an Argument from the XML description.
func argumentFromXML(xmlArg *xmlsrvdesc.Argument, action *Action) (*Argument, error) {
	if xmlArg.Name == "" {
		return nil, fmt.Errorf("%w: empty argument name", ErrBadDescription)
	}
	if xmlArg.RelatedStateVariable == "" {
		return nil, fmt.Errorf("%w: empty related state variable", ErrBadDescription)
	}
	return &Argument{
		Action:                   action,
		Name:                     xmlArg.Name,
		RelatedStateVariableName: xmlArg.RelatedStateVariable,
	}, nil
}

func (arg *Argument) RelatedStateVariable() (*StateVariable, error) {
	if v, ok := arg.Action.SCPD.VariableByName[arg.RelatedStateVariableName]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("%w: state variable %q", ErrMissingDefinition, arg.RelatedStateVariableName)
}

// StateVariable description data.
type StateVariable struct {
	Name     string
	DataType string

	AllowedValues []string
}

func stateVariableFromXML(xmlSV *xmlsrvdesc.StateVariable) (*StateVariable, error) {
	if xmlSV.Name == "" {
		return nil, fmt.Errorf("%w: empty state variable name", ErrBadDescription)
	}
	if xmlSV.DataType.Type != "" {
		return nil, fmt.Errorf("%w: unsupported data type %q",
			ErrUnsupportedDescription, xmlSV.DataType.Type)
	}
	if xmlSV.DataType.Name != "string" && len(xmlSV.AllowedValues) > 0 {
		return nil, fmt.Errorf("%w: allowedValueList is currently unsupported for type %q",
			ErrUnsupportedDescription, xmlSV.DataType.Name)
	}
	return &StateVariable{
		Name:          xmlSV.Name,
		DataType:      xmlSV.DataType.Name,
		AllowedValues: xmlSV.AllowedValues,
	}, nil
}
