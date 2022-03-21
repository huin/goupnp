package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type dcpProvider interface {
	process(tmpdir, name string, dcp *DCP) error
}
type upnpdotorg struct {
	DocURL     string // Optional - URL for further documentation about the DCP.
	XMLSpecURL string // Where to download the XML spec from.
	// Any special-case functions to run against the DCP before writing it out.
	Hacks []DCPHackFn
}

func (u upnpdotorg) process(tmpdir, name string, dcp *DCP) error {
	dcp.DocURLs = append(dcp.DocURLs, u.DocURL)
	specFilename := filepath.Join(tmpdir, name+".zip")
	err := acquireFile(specFilename, u.XMLSpecURL)
	if err != nil {
		return fmt.Errorf("could not acquire spec for %s: %v", name, err)
	}
	archive, err := zip.OpenReader(specFilename)
	if err != nil {
		return fmt.Errorf("error reading zip file %q: %v", specFilename, err)
	}
	defer archive.Close()
	if err := dcp.processZipFile(archive.File, []string{"*/device/*.xml"}, []string{"*/service/*.xml"}); err != nil {
		return fmt.Errorf("error processing spec file %q: %v", specFilename, err)
	}
	for i, hack := range u.Hacks {
		if err := hack(dcp); err != nil {
			return fmt.Errorf("error with Hack[%d] for %s: %v", i, name, err)
		}
	}
	return nil
}

const ocfSpecsURL = "https://openconnectivity.org/upnp-specs/upnpresources.zip"

type openconnectivitydotorg struct {
	DocPath        string // Optional - Glob to the related documentation about the DCP.
	SpecsURL       string // The HTTP location of the zip archive containing all XML spec.
	XMLSpecZipPath string // Glob to the zip XML spec file within upnpresources.zip.
	// Glob to the services XML files within the ZIP matching XMLSpecZipPath.
	XMLServicePath []string
	// Glob to the devices XML files within the ZIP matching XMLSpecZipPath.
	XMLDevicePath []string
	// Any special-case functions to run against the DCP before writing it out.
	Hacks []DCPHackFn
}

func (o openconnectivitydotorg) process(tmpdir, name string, dcp *DCP) error {
	fname := filepath.Base(name)
	allSpecsFilename := filepath.Join(tmpdir, "openconnectivitydotorg_"+fname+".zip")
	err := acquireFile(allSpecsFilename, o.SpecsURL)
	if err != nil {
		return fmt.Errorf("could not acquire specs %s: %v", name, err)
	}
	allSpecsArchive, err := zip.OpenReader(allSpecsFilename)
	if err != nil {
		return fmt.Errorf("error reading zip file %q: %v", allSpecsFilename, err)
	}
	specsArchives := globFiles(o.XMLSpecZipPath, allSpecsArchive.File)
	if len(specsArchives) < 1 {
		return fmt.Errorf("zip archive %q does not contain specifications at %q", allSpecsFilename, o.XMLSpecZipPath)
	}
	for _, specArchive := range specsArchives {
		f, err := specArchive.Open()
		if err != nil {
			return fmt.Errorf("error reading zip file %q: %v", specArchive.Name, err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		archive, err := zip.NewReader(bytes.NewReader(b), specArchive.FileInfo().Size())
		if err != nil {
			return fmt.Errorf("error reading zip file %q: %v", specArchive.Name, err)
		}
		if err := dcp.processZipFile(archive.File, o.XMLDevicePath, o.XMLServicePath); err != nil {
			return fmt.Errorf("error processing spec file %q: %v", specArchive.Name, err)
		}
	}
	for i, hack := range o.Hacks {
		if err := hack(dcp); err != nil {
			return fmt.Errorf("error with Hack[%d] for %s: %v", i, name, err)
		}
	}

	for _, d := range globFiles(o.DocPath, allSpecsArchive.File) {
		dcp.DocURLs = append(dcp.DocURLs, d.Name)
	}
	return nil
}
