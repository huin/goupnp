package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
)

func acquireFile(specFilename string, xmlSpecURL string) error {
	if f, err := os.Open(specFilename); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		f.Close()
		return nil
	}

	tmpFilename := specFilename + ".download"
	if err := downloadFile(tmpFilename, xmlSpecURL); err != nil {
		return err
	}

	return os.Rename(tmpFilename, specFilename)
}

func downloadFile(filename, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not download %q from %q: %s",
			filename, url, resp.Status)
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		w.Close()
		return err
	}

	return w.Close()
}

func globFiles(pattern string, archive *zip.ReadCloser) []*zip.File {
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
	defer r.Close()
	return decoder.Decode(data)
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
