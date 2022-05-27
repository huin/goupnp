package zipread

import (
	"archive/zip"
	"bytes"
	"io"
	"io/fs"
	"os"
)

type SizedReaderAt struct {
	R    io.ReaderAt
	Size int64
}

func NewSizedReaderFromOsFile(f *os.File) (*SizedReaderAt, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return &SizedReaderAt{R: f, Size: stat.Size()}, nil
}

func NewSizedReaderFromReader(r io.Reader) (*SizedReaderAt, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	contents := bytes.NewReader(data)
	return &SizedReaderAt{R: contents, Size: int64(len(data))}, nil
}

type ZipRead struct {
	*zip.Reader
}

func New(r *SizedReaderAt) (*ZipRead, error) {
	zr, err := zip.NewReader(r.R, r.Size)
	if err != nil {
		return nil, err
	}
	return &ZipRead{zr}, nil
}

func FromOsFile(f *os.File) (*ZipRead, error) {
	r, err := NewSizedReaderFromOsFile(f)
	if err != nil {
		return nil, err
	}
	return New(r)
}

func FromFsFile(f fs.File) (*ZipRead, error) {
	r, err := NewSizedReaderFromReader(f)
	if err != nil {
		return nil, err
	}
	return New(r)
}

func (zr *ZipRead) OpenZip(path string) (*ZipRead, error) {
	f, err := zr.Reader.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return FromFsFile(f)
}
