// Definition for the SOAP structure required for UPnP's SOAP usage.

package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	soapEncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	soapPrefix        = `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body>`
	soapSuffix        = `</s:Body></s:Envelope>`
)

type SOAPClient struct {
	EndpointURL url.URL
	HTTPClient  http.Client
}

func NewSOAPClient(endpointURL url.URL) *SOAPClient {
	return &SOAPClient{
		EndpointURL: endpointURL,
	}
}

// PerformSOAPAction makes a SOAP request, with the given action.
func (client *SOAPClient) PerformAction(actionNamespace, actionName string, inAction interface{}, outAction interface{}) error {
	requestBytes, err := encodeRequestAction(inAction)
	if err != nil {
		return err
	}

	response, err := client.HTTPClient.Do(&http.Request{
		Method: "POST",
		URL:    &client.EndpointURL,
		Header: http.Header{
			"SOAPACTION":   []string{actionNamespace + "#" + actionName},
			"CONTENT-TYPE": []string{"text/xml; charset=\"utf-8\""},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(requestBytes)),
		// Set ContentLength to avoid chunked encoding - some servers might not support it.
		ContentLength: int64(len(requestBytes)),
	})
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("goupnp: SOAP request got HTTP %s", response.Status)
	}

	responseEnv := newSOAPEnvelope()
	decoder := xml.NewDecoder(response.Body)
	if err := decoder.Decode(responseEnv); err != nil {
		return err
	}

	if responseEnv.Body.Fault != nil {
		return responseEnv.Body.Fault
	}

	if outAction != nil {
		if err := xml.Unmarshal(responseEnv.Body.RawAction, outAction); err != nil {
			return err
		}
	}

	return nil
}

// newSOAPAction creates a soapEnvelope with the given action and arguments.
func newSOAPEnvelope() *soapEnvelope {
	return &soapEnvelope{
		EncodingStyle: soapEncodingStyle,
	}
}

// encodeRequestAction is a hacky way to create an encoded SOAP envelope
// containing the given action. Experiments with one router have shown that it
// 500s for requests where the outer default xmlns is set to the SOAP
// namespace, and then reassigning the default namespace within that to the
// service namespace. Hand-coding the outer XML to work-around this.
func encodeRequestAction(inAction interface{}) ([]byte, error) {
	requestBuf := new(bytes.Buffer)
	requestBuf.WriteString(soapPrefix)
	if inAction != nil {
		requestEnc := xml.NewEncoder(requestBuf)
		if err := requestEnc.Encode(inAction); err != nil {
			return nil, err
		}
	}
	requestBuf.WriteString(soapSuffix)
	return requestBuf.Bytes(), nil
}

type soapEnvelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
	Body          soapBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type soapBody struct {
	Fault     *SOAPFaultError `xml:"Fault"`
	RawAction []byte          `xml:",innerxml"`
}

// SOAPFaultError implements error, and contains SOAP fault information.
type SOAPFaultError struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}

func (err *SOAPFaultError) Error() string {
	return fmt.Sprintf("SOAP fault: %s", err.FaultString)
}
