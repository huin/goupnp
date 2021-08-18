package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/scpd"
	"github.com/huin/goupnp/soap"
)

// DCP collects together information about a UPnP Device Control Protocol.
type DCP struct {
	Metadata     DCPMetadata
	DocURLs      []string
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

func (dcp *DCP) Reset() {
	dcp.DocURLs = nil
	dcp.DeviceTypes = make(map[string]*URNParts)
	dcp.ServiceTypes = make(map[string]*URNParts)
}

func (dcp *DCP) processZipFile(archive []*zip.File, devices, services []string) error {
	var f int
	for _, devicesGlob := range devices {
		for _, deviceFile := range globFiles(devicesGlob, archive) {
			if err := dcp.processDeviceFile(deviceFile); err != nil {
				return err
			}
			f++
		}
	}
	for _, scpdsGlob := range services {
		for _, scpdFile := range globFiles(scpdsGlob, archive) {
			if err := dcp.processSCPDFile(scpdFile); err != nil {
				return err
			}
			f++
		}
	}
	if f < 1 {
		return fmt.Errorf("no sdcp/device found in %q and %q", devices, services)
	}
	return nil
}

func (dcp *DCP) processDeviceFile(file *zip.File) error {
	var device goupnp.Device
	if err := unmarshalXmlFile(file, &device); err != nil {
		return fmt.Errorf("error decoding device XML from file %q: %v", file.Name, err)
	}
	var mainErr error
	device.VisitDevices(func(d *goupnp.Device) {
		t := strings.TrimSpace(d.DeviceType)
		if t != "" {
			u, err := extractURNParts(t, deviceURNPrefix)
			if err != nil {
				mainErr = err
			}
			dcp.DeviceTypes[t] = u
		}
	})
	device.VisitServices(func(s *goupnp.Service) {
		u, err := extractURNParts(s.ServiceType, serviceURNPrefix)
		if err != nil {
			mainErr = err
		}
		dcp.ServiceTypes[s.ServiceType] = u
	})
	return mainErr
}

func (dcp *DCP) writeCode(outFile string, codeTmpl *template.Template) error {
	packageFile, err := os.Create(outFile)
	if err != nil {
		return err
	}
	var output io.WriteCloser = packageFile
	if err = codeTmpl.Execute(output, dcp); err != nil {
		output.Close()
		return err
	}
	return output.Close()
}

func (dcp *DCP) OrderedServices() []SCPDWithURN {
	services := append([]SCPDWithURN{}, dcp.Services...)
	sort.SliceStable(services, func(i, j int) bool {
		return services[i].URNParts.URN < services[j].URNParts.URN
	})
	return services
}

func orderedURNParts(urnMap map[string]*URNParts) []*URNParts {
	urns := make([]*URNParts, 0, len(urnMap))
	for _, urn := range urnMap {
		urns = append(urns, urn)
	}
	sort.SliceStable(urns, func(i, j int) bool {
		return urns[i].URN < urns[j].URN
	})
	return urns
}

func (dcp *DCP) OrderedDeviceTypes() []*URNParts {
	return orderedURNParts(dcp.DeviceTypes)
}

func (dcp *DCP) OrderedServiceTypes() []*URNParts {
	return orderedURNParts(dcp.ServiceTypes)
}

func (dcp *DCP) processSCPDFile(file *zip.File) error {
	scpd := new(scpd.SCPD)
	if err := unmarshalXmlFile(file, scpd); err != nil {
		return fmt.Errorf("error decoding SCPD XML from file %q: %v", file.Name, err)
	}
	scpd.Clean()
	urnParts, err := urnPartsFromSCPDFilename(file.Name)
	if err != nil {
		return fmt.Errorf("could not recognize SCPD filename %q: %v", file.Name, err)
	}
	dcp.Services = append(dcp.Services, SCPDWithURN{
		URNParts: urnParts,
		SCPD:     scpd,
	})
	return nil
}

type SCPDWithURN struct {
	*URNParts
	SCPD *scpd.SCPD
}

func (s *SCPDWithURN) WrapArguments(args []*scpd.Argument) (argumentWrapperList, error) {
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

func (s *SCPDWithURN) wrapArgument(arg *scpd.Argument) (*argumentWrapper, error) {
	relVar := s.SCPD.GetStateVariable(arg.RelatedStateVariable)
	if relVar == nil {
		return nil, fmt.Errorf("no such state variable: %q, for argument %q", arg.RelatedStateVariable, arg.Name)
	}
	cnv, ok := soap.TypeDataMap[relVar.DataType.Name]
	if !ok {
		return nil, fmt.Errorf("unknown data type: %q, for state variable %q, for argument %q", relVar.DataType.Type, arg.RelatedStateVariable, arg.Name)
	}
	return &argumentWrapper{
		Argument:   *arg,
		relVar:     relVar,
		typeDataV1: cnv,
	}, nil
}

type argumentWrapper struct {
	scpd.Argument
	relVar     *scpd.StateVariable
	typeDataV1 soap.TypeData
}

func (arg *argumentWrapper) AsParameter() string {
	return fmt.Sprintf("%s %s", arg.Name, arg.typeDataV1.GoTypeName())
}

func (arg *argumentWrapper) HasDoc() bool {
	rng := arg.relVar.AllowedValueRange
	return ((rng != nil && (rng.Minimum != "" || rng.Maximum != "" || rng.Step != "")) ||
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

func (arg *argumentWrapper) MarshalV1() string {
	return fmt.Sprintf("soap.%s(%s)", arg.typeDataV1.MarshalFunc(), arg.Name)
}

func (arg *argumentWrapper) UnmarshalV1(objVar string) string {
	return fmt.Sprintf("soap.%s(%s.%s)", arg.typeDataV1.UnmarshalFunc(), objVar, arg.Name)
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
		return nil, fmt.Errorf("%q does not have expected prefix %q", urn, expectedPrefix)
	}
	parts := strings.SplitN(strings.TrimPrefix(urn, expectedPrefix), ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("%q does not have a name and version", urn)
	}
	name, version := parts[0], parts[1]
	return &URNParts{urn, name, version}, nil
}
