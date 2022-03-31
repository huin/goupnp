// Package envelope is responsible for encoding and decoding SOAP envelopes.
package envelope

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// ErrFault can be used as a target with errors.Is.
var ErrFault error = errors.New("xml fault")

// FaultDetail carries XML-encoded application-specific Fault details.
type FaultDetail struct {
	Raw []byte `xml:",innerxml"`
}

// Fault implements error, and contains SOAP fault information.
type Fault struct {
	Code   string      `xml:"faultcode"`
	String string      `xml:"faultstring"`
	Actor  string      `xml:"faultactor"`
	Detail FaultDetail `xml:"detail"`
}

func (fe *Fault) Error() string {
	return fmt.Sprintf("SOAP fault code=%s: %s", fe.Code, fe.String)
}

func (fe *Fault) Is(target error) bool {
	return target == ErrFault
}

// Action wraps a SOAP action to be read or written as part of a SOAP envelope.
type Action struct {
	// XMLName specifies the XML element namespace (URI) and name. Together
	// these identify the SOAP action.
	XMLName xml.Name
	// Args is an arbitrary struct containing fields for encoding or decoding
	// arguments. See https://pkg.go.dev/encoding/xml@go1.17.1#Marshal and
	// https://pkg.go.dev/encoding/xml@go1.17.1#Unmarshal for details on
	// annotating fields in the structure.
	Args any
}

// NewSendAction creates a SOAP action for receiving arguments.
func NewRecvAction(args any) *Action {
	return &Action{Args: args}
}

// NewSendAction creates a SOAP action for sending with the given namespace URL,
// action name, and arguments.
func NewSendAction(serviceType, actionName string, args any) *Action {
	return &Action{
		XMLName: xml.Name{Space: serviceType, Local: actionName},
		Args:    args,
	}
}

var _ xml.Marshaler = &Action{}

// MarshalXML implements `xml.Marshaller`.
//
// This is an implementation detail that allows packing elements inside the
// action element from the struct in `a.Args`.
func (a *Action) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	v := reflect.Indirect(reflect.ValueOf(a.Args))
	t := v.Type()
	elemName := xml.Name{Space: "", Local: "u:" + a.XMLName.Local}
	startElement := xml.StartElement{
		Name: elemName,
		Attr: []xml.Attr{{
			Name:  xml.Name{Space: "", Local: "xmlns:u"},
			Value: a.XMLName.Space,
		}},
	}
	switch t.Kind() {
	case reflect.Struct:
		// Hardcodes the XML namespace. See comment in Write() for context.
		return e.EncodeElement(a.Args, startElement)
	case reflect.Map:
		if err := e.EncodeToken(startElement); err != nil {
			return err
		}
		kt := t.Key()
		if kt.Kind() != reflect.String {
			return fmt.Errorf(
				"SOAP action wants string as map key in args: %w",
				&xml.UnsupportedTypeError{Type: kt})
		}
		iter := v.MapRange()
		for iter.Next() {
			k := iter.Key()
			// TODO: does this support string newtypes? convert?
			ks := k.Interface().(string)
			v := iter.Value()
			ke := xml.StartElement{Name: xml.Name{Local: ks}}
			if err := e.EncodeElement(v.Interface(), ke); err != nil {
				return fmt.Errorf(
					"SOAP action error while encoding arg %q: %w", ks, err)
			}
		}
		return e.EncodeToken(xml.EndElement{Name: elemName})
	default:
		return fmt.Errorf(
			"SOAP action does not support type as args: %w",
			&xml.UnsupportedTypeError{Type: t})
	}
}

var _ xml.Unmarshaler = &Action{}

// UnmarshalXML implements `xml.Unmarshaller`.
//
// This is an implementation detail that allows unpacking elements inside the
// action element into the struct in `a.Args`.
func (a *Action) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.XMLName = start.Name
	argsValue := reflect.Indirect(reflect.ValueOf(a.Args))
	argsType := argsValue.Type()
	switch argsType.Kind() {
	case reflect.Struct:
		return d.DecodeElement(a.Args, &start)
	case reflect.Map:
		keyType := argsType.Key()
		if keyType.Kind() != reflect.String {
			return fmt.Errorf(
				"SOAP action wants string as map key in args: %w",
				&xml.UnsupportedTypeError{Type: keyType})
		}
		valueType := argsType.Elem()
		if valueType.Kind() == reflect.Interface {
			return fmt.Errorf(
				"SOAP action wants a concrete type as map value in args: %w",
				&xml.UnsupportedTypeError{Type: valueType})
		}
		for {
			untypedToken, err := d.Token()
			if err != nil {
				return err
			}
			switch token := untypedToken.(type) {
			case xml.EndElement:
				return nil
			case xml.StartElement:
				if len(token.Attr) > 0 {
					return fmt.Errorf(
						"SOAP action arg does not support attributes, got %v",
						token.Attr)
				}
				if token.Name.Space != "" {
					return fmt.Errorf(
						"SOAP action arg does not support non-empty namespace, got %q",
						token.Name.Space)
				}
				key := token.Name.Local
				value := reflect.New(valueType)
				if err := d.DecodeElement(value.Interface(), &token); err != nil {
					return fmt.Errorf(
						"SOAP action arg %q errored while decoding: %w", key, err)
				}
				argsValue.SetMapIndex(reflect.ValueOf(key), reflect.Indirect(value))
			case xml.Comment:
			case xml.ProcInst:
				return fmt.Errorf(
					"SOAP action args contained unexpected token %v",
					untypedToken)
			case xml.Directive:
				return fmt.Errorf(
					"SOAP action args contained unexpected token %v",
					untypedToken)
			case xml.CharData:
				cd := string(token)
				if len(strings.TrimSpace(cd)) > 0 {
					return fmt.Errorf(
						"SOAP action args contained stray text: %q", cd)
				}
			default:
				return fmt.Errorf(
					"SOAP action found unknown XML token type: %T", untypedToken)
			}
		}
	default:
		return fmt.Errorf(
			"SOAP action does not support type as args: %w",
			&xml.UnsupportedTypeError{Type: argsType})
	}
}

// Various "constant" bytes used in the written envelope.
var (
	envOpen  = []byte(xml.Header + `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body>`)
	envClose = []byte(`</s:Body></s:Envelope>`)
)

// Write marshals a SOAP envelope to the writer. Errors can be from the writer
// or XML encoding.
func Write(w io.Writer, action *Action) error {
	// Experiments with one router have shown that it 500s for requests where
	// the outer default xmlns is set to the SOAP namespace, and then
	// reassigning the default namespace within that to the service namespace.
	// Most of the code in this function is hand-coding the outer XML to
	// workaround this.
	// Resolving https://github.com/golang/go/issues/9519 might remove the need
	// for this workaround.

	_, err := w.Write(envOpen)
	if err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	err = enc.Encode(action)
	if err != nil {
		return err
	}
	err = enc.Flush()
	_, err = w.Write(envClose)
	return err
}

// Read unmarshals a SOAP envelope from the reader. Errors can either be from
// the reader, XML decoding, or a *Fault.
func Read(r io.Reader, action *Action) error {
	env := envelope{
		Body: body{
			Action: action,
		},
	}

	dec := xml.NewDecoder(r)
	err := dec.Decode(&env)
	if err != nil {
		return err
	}

	if env.Body.Fault != nil {
		return env.Body.Fault
	}

	return nil
}

type envelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
	Body          body     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type body struct {
	Fault  *Fault  `xml:"Fault"`
	Action *Action `xml:",any"`
}
