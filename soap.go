// Definition for the SOAP structure required for UPnP's SOAP usage.

package goupnp

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	SoapEncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
)

type NameValue struct {
	Name  string
	Value string
}

// NewSoapAction creates a SoapEnvelope with the given action and arguments.
func newSoapAction(actionNamespace, actionName string, arguments []NameValue) *SoapEnvelope {
	env := &SoapEnvelope{
		EncodingStyle: SoapEncodingStyle,
		Body: SoapBody{
			Action: SoapAction{
				XMLName:   xml.Name{actionNamespace, actionName},
				Arguments: make([]SoapArgument, len(arguments)),
			},
		},
	}

	for i := range arguments {
		env.Body.Action.Arguments[i].XMLName.Local = arguments[i].Name
		env.Body.Action.Arguments[i].Value = arguments[i].Value
	}

	return env
}

// PerformSoapAction makes a SOAP request, with the given action.
func PerformSoapAction(actionNamespace, actionName string, url *url.URL, arguments []NameValue) ([]NameValue, error) {
	requestEnv := newSoapAction(actionNamespace, actionName, arguments)
	requestBytes, err := xml.Marshal(requestEnv)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(&http.Request{
		Method: "POST",
		URL:    url,
		Header: http.Header{
			"SOAPACTION":   []string{actionNamespace + "#" + actionName},
			"CONTENT-TYPE": []string{"text/xml; charset=\"utf-8\""},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(requestBytes)),
		// Set ContentLength to avoid chunked encoding - some servers might not support it.
		ContentLength: int64(len(requestBytes)),
	})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("goupnp: SOAP request got %s response from %s", response.Status, url.String())
	}

	decoder := xml.NewDecoder(response.Body)
	fmt.Println(response.Header)
	var responseEnv SoapEnvelope
	if err := decoder.Decode(&responseEnv); err != nil {
		return nil, err
	}

	results := make([]NameValue, len(responseEnv.Body.Action.Arguments))
	for i, soapArg := range responseEnv.Body.Action.Arguments {
		results[i] = NameValue{
			Name:  soapArg.XMLName.Local,
			Value: soapArg.Value,
		}
	}

	return results, nil
}

type SoapEnvelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
	Body          SoapBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type SoapBody struct {
	Action SoapAction `xml:",any"`
}

type SoapAction struct {
	XMLName   xml.Name
	Arguments []SoapArgument `xml:",any"`
}

type SoapArgument struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}
