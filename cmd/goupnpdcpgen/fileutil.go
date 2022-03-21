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
	"strings"
)

func acquireFile(specFilename string, xmlSpecURL string) error {
	tmpFilename := specFilename + ".download"
	defer os.Remove(tmpFilename)

	if fileExists(specFilename) {
		return nil
	}

	if err := downloadFile(tmpFilename, xmlSpecURL); err != nil {
		return err
	}

	return copyFile(specFilename, tmpFilename)
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

func globFiles(pattern string, archive []*zip.File) []*zip.File {
	var files []*zip.File
	pattern = strings.ToLower(pattern)
	for _, f := range archive {
		if matched, err := path.Match(pattern, strings.ToLower(f.Name)); err != nil {
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

func copyFile(dst string, src string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	return writeFile(dst, f)
}

func writeFile(dst string, r io.ReadCloser) error {
	defer r.Close()
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func fileExists(p string) bool {
	f, err := os.Open(p)
	if err != nil {
		return !os.IsNotExist(err)
	}
	f.Close()
	return true
}
