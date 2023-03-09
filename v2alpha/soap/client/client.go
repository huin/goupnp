// Package client provides a basic SOAP client.
package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/huin/goupnp/v2alpha/soap"
	"github.com/huin/goupnp/v2alpha/soap/envelope"
)

var (
	// ErrSOAP can be used with errors.Is.
	ErrSOAP = errors.New("SOAP error")
)

// SOAPError describes an error from this package, potentially including a
// lower-level cause.
type SOAPError struct {
	// description describes the error from the SOAP perspective.
	description string
	// cause may be nil.
	cause error
}

func (se *SOAPError) Error() string {
	b := &strings.Builder{}
	b.WriteString("SOAP error")
	if se.description != "" {
		b.WriteString(": ")
		b.WriteString(se.description)
	}
	if se.cause != nil {
		b.WriteString(": ")
		b.WriteString(se.cause.Error())
	}
	return b.String()
}

func (se *SOAPError) Is(target error) bool {
	return target == ErrSOAP
}

func (se *SOAPError) Unwrap() error {
	return se.cause
}

var _ HTTPClient = &http.Client{}

// HTTPClient defines the interface required of an HTTP client. It is a subset of *http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Option is the type for optional configuration of a Client.
type Option func(*options)

// WithHTTPClient specifies an *http.Client to use instead of
// http.DefaultClient.
func WithHTTPClient(httpClient HTTPClient) Option {
	return func(o *options) {
		o.httpClient = httpClient
	}
}

type options struct {
	httpClient HTTPClient
}

// Client is a SOAP client, attached to a specific SOAP endpoint.
// the zero value is not usable, use NewClient() to create an instance.
type Client struct {
	httpClient  HTTPClient
	endpointURL string
}

// New creates a new SOAP client, which will POST its requests to the
// given URL.
func New(endpointURL string, opts ...Option) *Client {
	co := options{
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(&co)
	}
	return &Client{
		httpClient:  co.httpClient,
		endpointURL: endpointURL,
	}
}

// PerformAction makes a SOAP request, with the given action values to provide
// arguments (`args`) and capture the `reply` into.
func (c *Client) Do(
	ctx context.Context,
	actionIn, actionOut *envelope.Action,
) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpointURL, nil)
	if err != nil {
		return err
	}
	if err := SetRequestAction(req, actionIn); err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &SOAPError{
			description: fmt.Sprintf("SOAP request got HTTP %s (%d)",
				resp.Status, resp.StatusCode),
		}
	}

	return ParseResponseAction(resp, actionOut)
}

// PerformAction makes a SOAP request, with the given action.
//
// This is a convenience for calling `Client.Do` without creating
// `*envelope.Action` values.
func PerformAction(
	ctx context.Context, c *Client,
	action soap.Action,
) error {
	actionIn := envelope.NewSendAction(
		action.ServiceType(), action.ActionName(), action.RefRequest())
	actionOut := &envelope.Action{Args: action.RefResponse()}
	return c.Do(ctx, actionIn, actionOut)
}

// SetRequestAction updates fields in `req` with the given SOAP action.
// Specifically it sets Body, ContentLength, Method, and the SOAPACTION and
// CONTENT-TYPE headers.
func SetRequestAction(
	req *http.Request,
	actionIn *envelope.Action,
) error {
	buf := &bytes.Buffer{}
	err := envelope.Write(buf, actionIn)
	if err != nil {
		return &SOAPError{
			description: "encoding envelope",
			cause:       err,
		}
	}

	req.Body = io.NopCloser(buf)
	req.ContentLength = int64(buf.Len())
	req.Method = http.MethodPost
	req.Header["SOAPACTION"] = []string{fmt.Sprintf(
		`"%s#%s"`, actionIn.XMLName.Space, actionIn.XMLName.Local)}
	req.Header["CONTENT-TYPE"] = []string{`text/xml; charset="utf-8"`}

	return nil
}

// ParseResponse extracts a parsed action from an HTTP response.
// The caller is responsible for calling resp.Body.Close(), but this function
// will consume the entire response body.
func ParseResponseAction(
	resp *http.Response,
	actionOut *envelope.Action,
) error {
	if resp.Body == nil {
		return &SOAPError{description: "missing HTTP response body"}
	}

	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return &SOAPError{
			description: "reading HTTP response body",
			cause:       err,
		}
	}

	if err := envelope.Read(buf, actionOut); err != nil {
		if errors.Is(err, envelope.ErrFault) {
			// Parsed cleanly, got SOAP fault.
			return &SOAPError{
				description: "SOAP fault",
				cause:       err,
			}
		}
		// Parsing problem, provide some information for context.
		dispLen := buf.Len()
		truncMessage := ""
		if dispLen > 1024 {
			dispLen = 1024
			truncMessage = fmt.Sprintf("first %d bytes (total %d bytes): ", dispLen, buf.Len())
		}
		return &SOAPError{
			description: fmt.Sprintf(
				"parsing SOAP response from HTTP body (%s%q)",
				truncMessage, buf.Bytes()[:dispLen],
			),
			cause: err,
		}
	}

	return nil
}
