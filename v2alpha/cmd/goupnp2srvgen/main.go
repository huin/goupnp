package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/huin/goupnp/v2alpha/cmd/goupnp2srvgen/tmplfuncs"
	"github.com/huin/goupnp/v2alpha/cmd/goupnp2srvgen/zipread"
	"github.com/huin/goupnp/v2alpha/description/srvdesc"
	"github.com/huin/goupnp/v2alpha/description/typedesc"
	"github.com/huin/goupnp/v2alpha/description/xmlsrvdesc"
	"github.com/huin/goupnp/v2alpha/soap/types"
)

var (
	srvTemplate      = flag.String("srv_template", "", "Path to srv.gotemplate.")
	upnpresourcesZip = flag.String("upnpresources_zip", "", "Path to upnpresources.zip.")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(flag.Args()) > 0 {
		return fmt.Errorf("unused arguments: %s", strings.Join(flag.Args(), " "))
	}
	if *srvTemplate == "" {
		return errors.New("-srv_template is a required flag.")
	}
	tmpl, err := template.New(filepath.Base(*srvTemplate)).Funcs(template.FuncMap{
		"args":  tmplfuncs.Args,
		"quote": strconv.Quote,
	}).ParseFiles(*srvTemplate)
	if err != nil {
		return fmt.Errorf("loading srv_template %q: %w", *srvTemplate, err)
	}

	if *upnpresourcesZip == "" {
		return errors.New("-upnpresources_zip is a required flag.")
	}
	f, err := os.Open(*upnpresourcesZip)
	if err != nil {
		return err
	}
	defer f.Close()
	upnpresources, err := zipread.FromOsFile(f)
	if err != nil {
		return err
	}

	// Use default type map for now. Addtional types could be use instead or
	// as well as necessary for extended types.
	typeMap := types.TypeMap()

	for _, m := range manifests {
		if err := processDCP(upnpresources, m, typeMap, tmpl); err != nil {
			return fmt.Errorf("processing DCP %s: %w", m.Path, err)
		}
	}
	return nil
}

var manifests = []*DCPSpecManifest{
	{
		Path: "standardizeddcps/Internet Gateway_2/UPnP-gw-IGD-TestFiles-20101210.zip",
		Services: []*ServiceManifest{
			{
				Package: "lanhostconfigmanagement1",
				Type:    "urn:schemas-upnp-org:service:LANHostConfigManagement:1",
				Path:    "xml data files/service/LANHostConfigManagement1.xml",
			},
			{
				Package: "wanpppconnection1",
				Type:    "urn:schemas-upnp-org:service:WANPPPConnection:1",
				Path:    "xml data files/service/WANPPPConnection1.xml",
			},
		},
	},
}

func processDCP(
	upnpresources *zipread.ZipRead,
	manifest *DCPSpecManifest,
	typeMap typedesc.TypeMap,
	tmpl *template.Template,
) error {
	dcpSpecData, err := upnpresources.OpenZip(manifest.Path)
	if err != nil {
		return err
	}
	for _, srvManifest := range manifest.Services {
		if err := processService(dcpSpecData, srvManifest, typeMap, tmpl); err != nil {
			return fmt.Errorf("processing service %s: %w", srvManifest.Type, err)
		}
	}
	return nil
}

func processService(
	dcpSpecData *zipread.ZipRead,
	srvManifest *ServiceManifest,
	typeMap typedesc.TypeMap,
	tmpl *template.Template,
) error {
	f, err := dcpSpecData.Open(srvManifest.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	d := xml.NewDecoder(f)

	xmlSCPD := &xmlsrvdesc.SCPD{}
	if err := d.Decode(xmlSCPD); err != nil {
		return err
	}
	xmlSCPD.Clean()

	sd, err := srvdesc.FromXML(xmlSCPD)
	if err != nil {
		return fmt.Errorf("transforming service description: %w", err)
	}

	imps, err := accumulateImports(sd, typeMap)
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "service", tmplArgs{
		Imps: imps,
		SCPD: sd,
	})
	if err != nil {
		return fmt.Errorf("executing srv_template: %w", err)
	}

	return nil
}

type DCPSpecManifest struct {
	// Path is the file path within upnpresources.zip to the DCP spec ZIP file.
	Path string
	// Services maps from a service name (e.g. "FooBar:1") to a path within the DCP spec ZIP file
	// (e.g. "xml data files/service/FooBar1.xml").
	Services []*ServiceManifest
}

type ServiceManifest struct {
	// Package is the Go package name to generate e.g. "foo1".
	Package string
	// Type is the SOAP namespace and service type that identifes the service e.g.
	// "urn:schemas-upnp-org:service:Foo:1"
	Type string
	// Path within the DCP spec ZIP file e.g. "xml data files/service/Foo1.xml".
	Path string
}

type tmplArgs struct {
	Imps *imports
	SCPD *srvdesc.SCPD
}

type imports struct {
	// Maps from a type name like "ui4" to the `alias.name` for the import.
	TypeRefByTypeName map[string]string
	// Each required import line, ordered by path.
	ImportLines []importItem
}

type importItem struct {
	Alias string
	Path  string
}

func accumulateImports(srvDesc *srvdesc.SCPD, typeMap typedesc.TypeMap) (*imports, error) {
	typeNames := make(map[string]bool)
	err := visitTypesSCPD(srvDesc, func(typeName string) {
		typeNames[typeName] = true
	})
	if err != nil {
		return nil, err
	}

	// Have sorted list of import package paths. Partly for aesthetics of generated code, but also
	// to have stable-generated aliases.
	paths := make(map[string]bool)
	for typeName := range typeNames {
		t, ok := typeMap[typeName]
		if !ok {
			return nil, fmt.Errorf("unknown type %q", typeName)
		}
		pkgPath := t.GoType.PkgPath()
		if pkgPath == "" {
			// Builtin type, ignore.
			continue
		}
		paths[pkgPath] = true
	}
	sortedPaths := make([]string, 0, len(paths))
	for path := range paths {
		sortedPaths = append(sortedPaths, path)
	}
	sort.Strings(sortedPaths)

	// Generate import aliases.
	index := 1
	aliasByPath := make(map[string]string, len(paths))
	importLines := make([]importItem, 0, len(paths))
	for _, path := range sortedPaths {
		alias := fmt.Sprintf("pkg%d", index)
		index++
		importLines = append(importLines, importItem{
			Alias: alias,
			Path:  path,
		})
		aliasByPath[path] = alias
	}

	// Populate typeRefByTypeName.
	typeRefByTypeName := make(map[string]string, len(typeNames))
	for typeName := range typeNames {
		goType := typeMap[typeName]
		pkgPath := goType.GoType.PkgPath()
		alias := aliasByPath[pkgPath]
		if alias == "" {
			// Builtin type.
			typeRefByTypeName[typeName] = goType.GoType.Name()
		} else {
			typeRefByTypeName[typeName] = fmt.Sprintf(
				"%s.%s", alias, goType.GoType.Name())
		}
	}

	return &imports{
		TypeRefByTypeName: typeRefByTypeName,
		ImportLines:       importLines,
	}, nil
}

type typeVisitor func(typeName string)

// visitTypesSCPD calls `visitor` with each data type name (e.g. "ui4") referenced
// by action arguments.`
func visitTypesSCPD(scpd *srvdesc.SCPD, visitor typeVisitor) error {
	for _, action := range scpd.ActionByName {
		if err := visitTypesAction(action, visitor); err != nil {
			return err
		}
	}
	return nil
}

func visitTypesAction(action *srvdesc.Action, visitor typeVisitor) error {
	for _, arg := range action.InArgs {
		sv, err := arg.RelatedStateVariable()
		if err != nil {
			return err
		}
		visitor(sv.DataType)
	}
	for _, arg := range action.OutArgs {
		sv, err := arg.RelatedStateVariable()
		if err != nil {
			return err
		}
		visitor(sv.DataType)
	}
	return nil
}
