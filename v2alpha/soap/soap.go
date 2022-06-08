// Package soap defines basic types used by SOAP packages.
package soap

// SOAPAction defines the interface for the convenience self-describing action request/response
// struct types.
type SOAPAction interface {
	// ServiceType returns Service type, e.g. "urn:schemas-upnp-org:service:Foo:1".
	ServiceType() string
	// ActionName returns Action name, e.g. "SetBar".
	ActionName() string
	// RefRequest returns reference to the action request member.
	RefRequest() any
	// RefResponse returns reference to the action response member.
	RefResponse() any
}
