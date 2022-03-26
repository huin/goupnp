// Package client provides a basic SOAP client.
package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/huin/goupnp/v2alpha/soap/envelope"
)

var _ HttpClient = &http.Client{}

// HttpClient defines the interface required of an HTTP client. It is a subset of *http.Client.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Option is the type for optional configuration of a Client.
type Option func(*options)

// WithHTTPClient specifies an *http.Client to use instead of
// http.DefaultClient.
func WithHTTPClient(httpClient HttpClient) Option {
	return func(o *options) {
		o.httpClient = httpClient
	}
}

type options struct {
	httpClient HttpClient
}

// Client is a SOAP client, attached to a specific SOAP endpoint.
// the zero value is not usable, use NewClient() to create an instance.
type Client struct {
	httpClient            HttpClient
	endpointURL           string
	maxErrorResponseBytes int
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

// PerformAction makes a SOAP request, with the given `argsIn` as input
// arguments, and `argsOut` to capture the output arguments into.
// `serviceType` is the SOAP service type URN, `actionName` is the action to
// call.
//
// This is a convenience for calling `c.Do` without creating `*Action` values.
func (c *Client) PerformAction(
	ctx context.Context, serviceType, actionName string,
	argsIn, argsOut any,
) error {
	actionIn := envelope.NewSendAction(serviceType, actionName, argsIn)
	actionOut := &envelope.Action{Args: argsOut}
	return c.Do(ctx, actionIn, actionOut)
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
		return fmt.Errorf("SOAP request got HTTP %s (%d)",
			resp.Status, resp.StatusCode)
	}

	return ParseResponseAction(resp, actionOut)
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
		return fmt.Errorf("encoding envelope: %w", err)
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
		return errors.New("missing response body")
	}

	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if err := envelope.Read(buf, actionOut); err != nil {
		if _, ok := err.(*envelope.Fault); ok {
			// Parsed cleanly, got SOAP fault.
			return err
		}
		// Parsing problem, provide some information for context.
		dispLen := buf.Len()
		truncMessage := ""
		if dispLen > 1024 {
			dispLen = 1024
			truncMessage = fmt.Sprintf("first %d bytes: ", dispLen)
		}
		return fmt.Errorf(
			"parsing response body (%s%q): %w",
			truncMessage, buf.Bytes()[:dispLen],
			err,
		)
	}

	return nil
}
