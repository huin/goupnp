// specgen generates Go code from the UPnP specification files.
//
// The specification is available for download from:
// http://upnp.org/resources/upnpresources.zip
package main

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/scpd"
	"github.com/huin/goutil/codegen"
)

// flags
var (
	specFilename = flag.String("spec", "", "Path to the specification file.")
	outDir       = flag.String("out-dir", "", "Path to the output directory.")
	enableGofmt  = flag.Bool("gofmt", true, "Pass the output through gofmt. "+
		"Disable if debugging code output problems.")
)

var (
	deviceURNPrefix  = "urn:schemas-upnp-org:device:"
	serviceURNPrefix = "urn:schemas-upnp-org:service:"
)

func main() {
	flag.Parse()

	if *specFilename == "" {
		log.Fatal("--spec is required")
	}
	if *outDir == "" {
		log.Fatal("--out-dir is required")
	}

	specArchive, err := openZipfile(*specFilename)
	if err != nil {
		log.Fatalf("Error opening %s: %v", *specFilename)
	}
	defer specArchive.Close()

	dcpsCol := newDcpsCollection(map[string]string{
		"Internet Gateway_1": "internetgateway1",
		"Internet Gateway_2": "internetgateway2",
	})
	for _, f := range globFiles("standardizeddcps/*/*.zip", specArchive.Reader) {
		dirName := strings.TrimPrefix(f.Name, "standardizeddcps/")
		slashIndex := strings.Index(dirName, "/")
		if slashIndex == -1 {
			// Should not happen.
			log.Printf("Could not find / in %q", dirName)
			return
		}
		dirName = dirName[:slashIndex]

		dcp := dcpsCol.dcpsForDir(dirName)
		if dcp == nil {
			log.Printf("No alias defined for directory %q: skipping", dirName)
			continue
		}

		dcp.processZipFile(f)
	}

	for _, dcp := range dcpsCol.dcpsByAlias {
		if err := dcp.writePackage(*outDir); err != nil {
			log.Printf("Error writing package %q: %v", dcp.Name, err)
		}
	}
}

type dcpsCollection struct {
	dcpsAliasByDir map[string]string
	dcpsByAlias    map[string]*DCP
}

func newDcpsCollection(dcpsAliasByDir map[string]string) *dcpsCollection {
	c := &dcpsCollection{
		dcpsAliasByDir: dcpsAliasByDir,
		dcpsByAlias:    make(map[string]*DCP),
	}
	for _, alias := range dcpsAliasByDir {
		c.dcpsByAlias[alias] = newDCP(alias)
	}
	return c
}

func (c dcpsCollection) dcpsForDir(dirName string) *DCP {
	alias, ok := c.dcpsAliasByDir[dirName]
	if !ok {
		return nil
	}
	return c.dcpsByAlias[alias]
}

// DCP collects together information about a UPnP Device Control Protocol.
type DCP struct {
	Name         string
	DeviceTypes  map[string]*URNParts
	ServiceTypes map[string]*URNParts
	Services     []SCPDWithURN
}

func newDCP(name string) *DCP {
	return &DCP{
		Name:         name,
		DeviceTypes:  make(map[string]*URNParts),
		ServiceTypes: make(map[string]*URNParts),
	}
}

func (dcp *DCP) processZipFile(file *zip.File) {
	archive, err := openChildZip(file)
	if err != nil {
		log.Println("Error reading child zip file:", err)
		return
	}
	for _, deviceFile := range globFiles("*/device/*.xml", archive) {
		dcp.processDeviceFile(deviceFile)
	}
	for _, scpdFile := range globFiles("*/service/*.xml", archive) {
		dcp.processSCPDFile(scpdFile)
	}
}

func (dcp *DCP) processDeviceFile(file *zip.File) {
	var device goupnp.Device
	if err := unmarshalXmlFile(file, &device); err != nil {
		log.Printf("Error decoding device XML from file %q: %v", file.Name, err)
		return
	}
	device.VisitDevices(func(d *goupnp.Device) {
		t := strings.TrimSpace(d.DeviceType)
		if t != "" {
			u, err := extractURNParts(t, deviceURNPrefix)
			if err != nil {
				log.Println(err)
				return
			}
			dcp.DeviceTypes[t] = u
		}
	})
	device.VisitServices(func(s *goupnp.Service) {
		u, err := extractURNParts(s.ServiceType, serviceURNPrefix)
		if err != nil {
			log.Println(err)
			return
		}
		dcp.ServiceTypes[s.ServiceType] = u
	})
}

func (dcp *DCP) writePackage(outDir string) error {
	packageDirname := filepath.Join(outDir, dcp.Name)
	err := os.MkdirAll(packageDirname, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	packageFilename := filepath.Join(packageDirname, dcp.Name+".go")
	packageFile, err := os.Create(packageFilename)
	if err != nil {
		return err
	}
	var output io.WriteCloser = packageFile
	if *enableGofmt {
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

func (dcp *DCP) processSCPDFile(file *zip.File) {
	scpd := new(scpd.SCPD)
	if err := unmarshalXmlFile(file, scpd); err != nil {
		log.Printf("Error decoding SCPD XML from file %q: %v", file.Name, err)
		return
	}
	scpd.Clean()
	urnParts, err := urnPartsFromSCPDFilename(file.Name)
	if err != nil {
		log.Printf("Could not recognize SCPD filename %q: %v", file.Name, err)
		return
	}
	dcp.Services = append(dcp.Services, SCPDWithURN{
		URNParts: urnParts,
		SCPD:     scpd,
	})
}

type SCPDWithURN struct {
	*URNParts
	SCPD *scpd.SCPD
}

func (s *SCPDWithURN) WrapArgument(arg scpd.Argument) (*argumentWrapper, error) {
	relVar := s.SCPD.GetStateVariable(arg.RelatedStateVariable)
	if relVar == nil {
		return nil, fmt.Errorf("no such state variable: %q, for argument %q", arg.RelatedStateVariable, arg.Name)
	}
	cnv, ok := typeConvs[relVar.DataType.Name]
	if !ok {
		return nil, fmt.Errorf("unknown data type: %q, for state variable %q, for argument %q", relVar.DataType.Type, arg.RelatedStateVariable, arg.Name)
	}
	return &argumentWrapper{
		Argument: arg,
		relVar:   relVar,
		conv:     cnv,
	}, nil
}

type argumentWrapper struct {
	scpd.Argument
	relVar *scpd.StateVariable
	conv   conv
}

func (arg *argumentWrapper) AsParameter() string {
	return fmt.Sprintf("%s %s", arg.Name, arg.conv.ExtType)
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
	return fmt.Sprintf("soap.Marshal%s(%s)", arg.conv.FuncSuffix, arg.Name)
}

func (arg *argumentWrapper) Unmarshal(objVar string) string {
	return fmt.Sprintf("soap.Unmarshal%s(%s.%s)", arg.conv.FuncSuffix, objVar, arg.Name)
}

type conv struct {
	FuncSuffix string
	ExtType    string
}

// typeConvs maps from a SOAP type (e.g "fixed.14.4") to the function name
// suffix inside the soap module (e.g "Fixed14_4") and the Go type.
var typeConvs = map[string]conv{
	"ui1":         conv{"Ui1", "uint8"},
	"ui2":         conv{"Ui2", "uint16"},
	"ui4":         conv{"Ui4", "uint32"},
	"i1":          conv{"I1", "int8"},
	"i2":          conv{"I2", "int16"},
	"i4":          conv{"I4", "int32"},
	"int":         conv{"Int", "int64"},
	"r4":          conv{"R4", "float32"},
	"r8":          conv{"R8", "float64"},
	"number":      conv{"R8", "float64"}, // Alias for r8.
	"fixed.14.4":  conv{"Fixed14_4", "float64"},
	"float":       conv{"R8", "float64"},
	"char":        conv{"Char", "rune"},
	"string":      conv{"String", "string"},
	"date":        conv{"Date", "time.Time"},
	"dateTime":    conv{"DateTime", "time.Time"},
	"dateTime.tz": conv{"DateTimeTz", "time.Time"},
	"time":        conv{"TimeOfDay", "soap.TimeOfDay"},
	"time.tz":     conv{"TimeOfDayTz", "soap.TimeOfDay"},
	"boolean":     conv{"Boolean", "bool"},
	"bin.base64":  conv{"BinBase64", "[]byte"},
	"bin.hex":     conv{"BinHex", "[]byte"},
}

type closeableZipReader struct {
	io.Closer
	*zip.Reader
}

func openZipfile(filename string) (*closeableZipReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	archive, err := zip.NewReader(file, fi.Size())
	if err != nil {
		return nil, err
	}
	return &closeableZipReader{
		Closer: file,
		Reader: archive,
	}, nil
}

// openChildZip opens a zip file within another zip file.
func openChildZip(file *zip.File) (*zip.Reader, error) {
	zipFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()

	zipBytes, err := ioutil.ReadAll(zipFile)
	if err != nil {
		return nil, err
	}

	return zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
}

func globFiles(pattern string, archive *zip.Reader) []*zip.File {
	var files []*zip.File
	for _, f := range archive.File {
		if matched, err := path.Match(pattern, f.Name); err != nil {
			// This shouldn't happen - all patterns are hard-coded, errors in them
			// are a programming error.
			panic(err)
		} else if matched {
			files = append(files, f)
		}
	}
	return files
}

func unmarshalXmlFile(file *zip.File, data interface{}) error {
	r, err := file.Open()
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(r)
	r.Close()
	return decoder.Decode(data)
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

var scpdFilenameRe = regexp.MustCompile(
	`.*/([a-zA-Z0-9]+)([0-9]+)\.xml`)

func urnPartsFromSCPDFilename(filename string) (*URNParts, error) {
	parts := scpdFilenameRe.FindStringSubmatch(filename)
	if len(parts) != 3 {
		return nil, fmt.Errorf("SCPD filename %q does not have expected number of parts", filename)
	}
	name, version := parts[1], parts[2]
	return &URNParts{
		URN:     serviceURNPrefix + name + ":" + version,
		Name:    name,
		Version: version,
	}, nil
}
