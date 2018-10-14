// Package errkind contains the error kinds produced by goupnp.
package errkind

import (
	"fmt"

	errors "gopkg.in/src-d/go-errors.v1"
)

var (
	// BadData matches errors due to bad read or received data.
	BadData = errors.NewKind("%s")

	// FileContext provides context for a file that an error was encountered in.
	FileContext = errors.NewKind("in file %q")

	// HTTP matches errors from the HTTP protocol.
	HTTP = errors.NewKind("%s")

	// InvalidArgument matches errors for invalid arguments or states when
	// using the API.
	InvalidArgument = errors.NewKind("%s")

	// Network matches network errors.
	Network = errors.NewKind("%s")

	// NotFound matches errors for a value that is not found.
	NotFound = errors.NewKind("%s")

	// SOAP matches errors from the SOAP protocol.
	SOAP = errors.NewKind("%s")

	// SSDP matches errors from the SSDP protocol.
	SSDP = errors.NewKind("%s")
)

// New creates an error of the given kind, with the given format and values.
func New(kind *errors.Kind, format string, values ...interface{}) error {
	return kind.New(fmt.Sprintf(format, values...))
}

// New creates an error of the given kind, wrapping the given error, with the
// given format and values.
func Wrap(kind *errors.Kind, err error, format string, values ...interface{}) error {
	return kind.Wrap(err, fmt.Sprintf(format, values...))
}

// NewUnexpectedHTTPStatus returns an error of kind HTTP for an unexpected HTTP
// response code.
func NewUnexpectedHTTPStatus(statusCode int, status string) error {
	return New(
		HTTP, "bad HTTP response status %d %s",
		statusCode, status,
	)
}
