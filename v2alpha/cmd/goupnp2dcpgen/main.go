package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/huin/goupnp/v2alpha/cmd/goupnp2dcpgen/zipread"
	"github.com/huin/goupnp/v2alpha/description/scpd"
	"github.com/huin/goupnp/v2alpha/description/xmlscpd"
)

var (
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
	for _, m := range manifests {
		if err := processDCP(upnpresources, m); err != nil {
			return fmt.Errorf("processing DCP %s: %w", m.Path, err)
		}
	}
	return nil
}

var manifests = []*DCPSpecManifest{
	{
		Path: "standardizeddcps/Internet Gateway_2/UPnP-gw-IGD-TestFiles-20101210.zip",
		Services: map[string]string{
			"LANHostConfigManagement:1": "xml data files/service/LANHostConfigManagement1.xml",
			"WANPPPConnection:1":        "xml data files/service/WANPPPConnection1.xml",
		},
	},
}

func processDCP(
	upnpresources *zipread.ZipRead,
	manifest *DCPSpecManifest,
) error {
	dcpSpecData, err := upnpresources.OpenZip(manifest.Path)
	if err != nil {
		return err
	}
	for name, path := range manifest.Services {
		if err := processService(dcpSpecData, name, path); err != nil {
			return fmt.Errorf("processing service %s: %w", name, err)
		}
	}
	return nil
}

func processService(
	dcpSpecData *zipread.ZipRead,
	name string,
	path string,
) error {
	fmt.Printf("%s\n", name)
	f, err := dcpSpecData.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	d := xml.NewDecoder(f)

	xmlSCPD := &xmlscpd.SCPD{}
	if err := d.Decode(xmlSCPD); err != nil {
		return err
	}
	xmlSCPD.Clean()

	for _, action := range xmlSCPD.Actions {
		fmt.Printf("* %s()\n", action.Name)
		for _, arg := range action.Arguments {
			direction := "?"
			if arg.Direction == "in" {
				direction = "<-"
			} else if arg.Direction == "out" {
				direction = "->"
			}
			fmt.Printf("  %s %s %s\n", direction, arg.Name, arg.RelatedStateVariable)
		}
	}

	_, err := scpd.FromXML(xmlSCPD)
	if err != nil {
		return err
	}
	return nil
}

type DCPSpecManifest struct {
	// Path is the file path within upnpresources.zip to the DCP spec ZIP file.
	Path string
	// Services maps from a service name (e.g. "FooBar:1") to a path within the DCP spec ZIP file
	// (e.g. "xml data files/service/FooBar1.xml").
	Services map[string]string
}
