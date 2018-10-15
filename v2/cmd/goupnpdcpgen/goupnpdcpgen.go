// Command to generate DCP package source from the XML specification.
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/huin/goupnp/v2/errkind"
	errors "gopkg.in/src-d/go-errors.v1"
)

var (
	deviceURNPrefix  = "urn:schemas-upnp-org:device:"
	serviceURNPrefix = "urn:schemas-upnp-org:service:"
)

func main() {
	var (
		dcpName = flag.String("dcp_name", "",
			"Name of the DCP to generate.",
		)
		specsDir = flag.String("specs_dir", ".",
			"Path to the specification storage directory. "+
				"This is used to find (and download if not present) the "+
				"specification ZIP files.",
		)
		useGofmt = flag.Bool("gofmt", true,
			"Pass the generated code through gofmt. "+
				"Disable this if debugging code generation and needing to see "+
				"the generated code prior to being passed through gofmt.",
		)
	)
	flag.Parse()

	if err := run(*dcpName, *specsDir, *useGofmt); err != nil {
		log.Fatal(err)
	}
}

func run(dcpName, specsDir string, useGofmt bool) error {
	if err := os.MkdirAll(specsDir, os.ModePerm); err != nil {
		return errors.NewKind("create specs directory").Wrap(err)
	}

	for _, d := range dcpMetadataRegistry {
		if d.Name != dcpName {
			continue
		}
		specFilename := filepath.Join(specsDir, d.Name+".zip")
		err := acquireFile(specFilename, d.XMLSpecURL)
		if err != nil {
			return errkind.FileContext.Wrap(err, d.XMLSpecURL)
		}
		dcp := newDCP(d)
		if err := dcp.processZipFile(specFilename); err != nil {
			return errkind.FileContext.Wrap(err, specFilename)
		}
		for i, hack := range d.Hacks {
			if err := hack(dcp); err != nil {
				return errors.NewKind("in Hack[%d]").Wrap(err, i)
			}
		}
		if err := dcp.writeCode(d.Name+".go", useGofmt); err != nil {
			return errors.NewKind("writing generated code").Wrap(err)
		}

		return nil
	}

	return errkind.New(errkind.NotFound, "no DCP found with name %q", dcpName)
}
