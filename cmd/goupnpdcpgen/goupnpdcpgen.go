// Command to generate DCP package source from the XML specification.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var (
	deviceURNPrefix  = "urn:schemas-upnp-org:device:"
	serviceURNPrefix = "urn:schemas-upnp-org:service:"
)

func main() {
	var (
		dcpName  = flag.String("dcp_name", "", "Name of the DCP to generate.")
		specsDir = flag.String("specs_dir", ".", "Path to the specification storage directory. "+
			"This is used to find (and download if not present) the specification ZIP files.")
		useGofmt = flag.Bool("gofmt", true, "Pass the generated code through gofmt. "+
			"Disable this if debugging code generation and needing to see the generated code "+
			"prior to being passed through gofmt.")
		codeTmplFile = flag.String("code_tmpl_file", "", "Path to Go template to generate code from.")
	)
	flag.Parse()

	if err := run(*dcpName, *specsDir, *useGofmt, *codeTmplFile); err != nil {
		log.Fatal(err)
	}
}

func run(dcpName, specsDir string, useGofmt bool, codeTmplFile string) error {
	codeTmpl, err := template.New(filepath.Base(codeTmplFile)).Funcs(template.FuncMap{
		"base": filepath.Base,
	}).ParseFiles(codeTmplFile)
	if err != nil {
		return fmt.Errorf("error parsing template from file: %w", err)
	}

	if err := os.MkdirAll(specsDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create specs-dir %q: %v", specsDir, err)
	}

	for _, metadata := range dcpMetadata {
		if metadata.Name != dcpName {
			continue
		}

		dcp := newDCP(metadata)

		err := metadata.Src.process(".", metadata.Name, dcp)
		if err != nil {
			return fmt.Errorf("error processing spec %s: %v", metadata.Name, err)
		}

		filename := filepath.Base(metadata.Name) + ".go"
		if err := dcp.writeCode(filename, codeTmpl); err != nil {
			return fmt.Errorf("error writing package %q: %v", dcp.Metadata.Name, err)
		}

		if !useGofmt {
			return nil
		}

		return gofmt(filename)
	}

	return fmt.Errorf("could not find DCP with name %q", dcpName)
}

func gofmt(filename string) error {
	cmd := exec.Command("gofmt", "-w", filename)
	return cmd.Run()
}
