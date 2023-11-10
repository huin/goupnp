package goupnp

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

// ValidUTF8Reader implements a Reader which reads only bytes that constitute valid UTF-8
type ValidUTF8Reader struct {
	buffer *bufio.Reader
}

// Function Read reads bytes in the byte array b. n is the number of bytes read.
func (rd ValidUTF8Reader) Read(b []byte) (n int, err error) {
	for {
		var r rune
		var size int
		r, size, err = rd.buffer.ReadRune()
		if err != nil {
			return
		}
		if r == unicode.ReplacementChar && size == 1 {
			continue
		} else if n+size < len(b) {
			utf8.EncodeRune(b[n:], r)
			n += size
		} else {
			rd.buffer.UnreadRune()
			break
		}
	}
	return
}

// NewValidUTF8Reader constructs a new ValidUTF8Reader that wraps an existing io.Reader
func NewValidUTF8Reader(rd io.Reader) ValidUTF8Reader {
	return ValidUTF8Reader{bufio.NewReader(rd)}
}
