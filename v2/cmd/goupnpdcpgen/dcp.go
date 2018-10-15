package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/huin/goupnp/v2/errkind"
	"github.com/huin/goupnp/v2/metadata"
	"github.com/huin/goutil/codegen"
)

// DCP collects together information about a UPnP Device Control Protocol.
type DCP struct {
	Metadata     DCPMetadata
	DeviceTypes  map[string]*URNParts
	ServiceTypes map[string]*URNParts
	Services     []SCPDWithURN
}

func newDCP(metadata DCPMetadata) *DCP {
	return &DCP{
		Metadata:     metadata,
		DeviceTypes:  make(map[string]*URNParts),
		ServiceTypes: make(map[string]*URNParts),
	}
}

func (dcp *DCP) processZipFile(filename string) error {
	archive, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer archive.Close()
	for _, deviceFile := range globFiles("*/device/*.xml", archive) {
		if err := dcp.processDeviceFile(deviceFile); err != nil {
			return errkind.FileContext.Wrap(err, deviceFile)
		}
	}
	for _, scpdFile := range globFiles("*/service/*.xml", archive) {
		if err := dcp.processSCPDFile(scpdFile); err != nil {
			return errkind.FileContext.Wrap(err, scpdFile)
		}
	}
	return nil
}

func (dcp *DCP) processDeviceFile(file *zip.File) error {
	var device metadata.Device
	if err := unmarshalXmlFile(file, &device); err != nil {
		return err
	}
	var mainErr error
	device.VisitDevices(func(d *metadata.Device) {
		t := strings.TrimSpace(d.DeviceType)
		if t != "" {
			u, err := extractURNParts(t, deviceURNPrefix)
			if err != nil {
				mainErr = err
			}
			dcp.DeviceTypes[t] = u
		}
	})
	device.VisitServices(func(s *metadata.Service) {
		u, err := extractURNParts(s.ServiceType, serviceURNPrefix)
		if err != nil {
			mainErr = err
		}
		dcp.ServiceTypes[s.ServiceType] = u
	})
	return mainErr
}

func (dcp *DCP) writeCode(outFile string, useGofmt bool) error {
	packageFile, err := os.Create(outFile)
	if err != nil {
		return err
	}
	var output io.WriteCloser = packageFile
	if useGofmt {
		if output, err = codegen.NewGofmtWriteCloser(output); err != nil {
			packageFile.Close()
			return err
		}
	}
	if err = packageTmpl.Execute(output, dcp); err != nil {
		output.Close()
		return err
	}
	return output.Close()
}

func (dcp *DCP) processSCPDFile(file *zip.File) error {
	scpd := new(metadata.SCPD)
	if err := unmarshalXmlFile(file, scpd); err != nil {
		return err
	}
	scpd.Clean()
	urnParts, err := urnPartsFromSCPDFilename(file.Name)
	if err != nil {
		return err
	}
	dcp.Services = append(dcp.Services, SCPDWithURN{
		URNParts: urnParts,
		SCPD:     scpd,
	})
	return nil
}

type SCPDWithURN struct {
	*URNParts
	SCPD *metadata.SCPD
}

func (s *SCPDWithURN) WrapArguments(
	args []*metadata.Argument,
) (argumentWrapperList, error) {
	wrappedArgs := make(argumentWrapperList, len(args))
	for i, arg := range args {
		wa, err := s.wrapArgument(arg)
		if err != nil {
			return nil, err
		}
		wrappedArgs[i] = wa
	}
	return wrappedArgs, nil
}

func (s *SCPDWithURN) wrapArgument(
	arg *metadata.Argument,
) (*argumentWrapper, error) {
	relVar := s.SCPD.GetStateVariable(arg.RelatedStateVariable)
	if relVar == nil {
		return nil, errkind.New(
			errkind.BadData,
			"no such state variable: %q, for argument %q",
			arg.RelatedStateVariable, arg.Name,
		)
	}
	cnv, ok := typeConvs[relVar.DataType.Name]
	if !ok {
		return nil, errkind.New(
			errkind.BadData,
			"unknown data type: %q, for state variable %q, for argument %q",
			relVar.DataType.Type,
			arg.RelatedStateVariable,
			arg.Name,
		)
	}
	return &argumentWrapper{
		Argument: *arg,
		relVar:   relVar,
		conv:     cnv,
	}, nil
}

type argumentWrapper struct {
	metadata.Argument
	relVar *metadata.StateVariable
	conv   conv
}

func (arg *argumentWrapper) AsParameter() string {
	return fmt.Sprintf("%s %s", arg.Name, arg.conv.ExtType)
}

func (arg *argumentWrapper) HasDoc() bool {
	rng := arg.relVar.AllowedValueRange
	return ((rng != nil &&
		(rng.Minimum != "" || rng.Maximum != "" || rng.Step != "")) ||
		len(arg.relVar.AllowedValues) > 0)
}

func (arg *argumentWrapper) Document() string {
	relVar := arg.relVar
	if rng := relVar.AllowedValueRange; rng != nil {
		var parts []string
		if rng.Minimum != "" {
			parts = append(parts, fmt.Sprintf("minimum=%s", rng.Minimum))
		}
		if rng.Maximum != "" {
			parts = append(parts, fmt.Sprintf("maximum=%s", rng.Maximum))
		}
		if rng.Step != "" {
			parts = append(parts, fmt.Sprintf("step=%s", rng.Step))
		}
		return "allowed value range: " + strings.Join(parts, ", ")
	}
	if len(relVar.AllowedValues) != 0 {
		return "allowed values: " + strings.Join(relVar.AllowedValues, ", ")
	}
	return ""
}

func (arg *argumentWrapper) Marshal() string {
	return fmt.Sprintf("soap.Marshal%s(\n%s,\n)", arg.conv.FuncSuffix, arg.Name)
}

func (arg *argumentWrapper) Unmarshal(objVar string) string {
	return fmt.Sprintf("soap.Unmarshal%s(\n%s.%s,\n)",
		arg.conv.FuncSuffix, objVar, arg.Name,
	)
}

type argumentWrapperList []*argumentWrapper

func (args argumentWrapperList) HasDoc() bool {
	for _, arg := range args {
		if arg.HasDoc() {
			return true
		}
	}
	return false
}

type URNParts struct {
	URN     string
	Name    string
	Version string
}

func (u *URNParts) Const() string {
	return fmt.Sprintf("URN_%s_%s", u.Name, u.Version)
}

// extractURNParts extracts the name and version from a URN string.
func extractURNParts(urn, expectedPrefix string) (*URNParts, error) {
	if !strings.HasPrefix(urn, expectedPrefix) {
		return nil, errkind.New(
			errkind.BadData, "%q does not have expected prefix %q",
			urn, expectedPrefix,
		)
	}
	parts := strings.SplitN(strings.TrimPrefix(urn, expectedPrefix), ":", 2)
	if len(parts) != 2 {
		return nil, errkind.New(
			errkind.BadData, "%q does not have a name and version", urn,
		)
	}
	name, version := parts[0], parts[1]
	return &URNParts{urn, name, version}, nil
}
