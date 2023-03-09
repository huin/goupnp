// Package lanhostcfgmgmt1 provides types for the "urn:schemas-upnp-org:service:LANHostConfigManagement:1" service.
//
// Documented at https://openconnectivity.org/wp-content/uploads/2015/11/UPnP-gw-LANHostConfigManagement-v1-Service.pdf.
package lanhostcfgmgmt1

import (
	pkg1 "github.com/huin/goupnp/v2alpha/soap"
	pkg2 "github.com/huin/goupnp/v2alpha/soap/types"
)

const ServiceType = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"

// DeleteDNSServer provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type DeleteDNSServer struct {
	Request  DeleteDNSServerRequest
	Response DeleteDNSServerResponse
}

var _ pkg1.Action = &DeleteDNSServer{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteDNSServer) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteDNSServer) ActionName() string { return "DeleteDNSServer" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteDNSServer) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteDNSServer) RefResponse() any { return &a.Response }

// DeleteDNSServerRequest contains the "in" args for the "DeleteDNSServer" action.
type DeleteDNSServerRequest struct {
	// NewDNSServers relates to state variable DNSServers.
	NewDNSServers string
}

// DeleteDNSServerResponse contains the "out" args for the "DeleteDNSServer" action.
type DeleteDNSServerResponse struct{}

// DeleteIPRouter provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type DeleteIPRouter struct {
	Request  DeleteIPRouterRequest
	Response DeleteIPRouterResponse
}

var _ pkg1.Action = &DeleteIPRouter{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteIPRouter) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteIPRouter) ActionName() string { return "DeleteIPRouter" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteIPRouter) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteIPRouter) RefResponse() any { return &a.Response }

// DeleteIPRouterRequest contains the "in" args for the "DeleteIPRouter" action.
type DeleteIPRouterRequest struct {
	// NewIPRouters relates to state variable IPRouters.
	NewIPRouters string
}

// DeleteIPRouterResponse contains the "out" args for the "DeleteIPRouter" action.
type DeleteIPRouterResponse struct{}

// DeleteReservedAddress provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type DeleteReservedAddress struct {
	Request  DeleteReservedAddressRequest
	Response DeleteReservedAddressResponse
}

var _ pkg1.Action = &DeleteReservedAddress{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteReservedAddress) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteReservedAddress) ActionName() string { return "DeleteReservedAddress" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteReservedAddress) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeleteReservedAddress) RefResponse() any { return &a.Response }

// DeleteReservedAddressRequest contains the "in" args for the "DeleteReservedAddress" action.
type DeleteReservedAddressRequest struct {
	// NewReservedAddresses relates to state variable ReservedAddresses.
	NewReservedAddresses string
}

// DeleteReservedAddressResponse contains the "out" args for the "DeleteReservedAddress" action.
type DeleteReservedAddressResponse struct{}

// GetAddressRange provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetAddressRange struct {
	Request  GetAddressRangeRequest
	Response GetAddressRangeResponse
}

var _ pkg1.Action = &GetAddressRange{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAddressRange) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAddressRange) ActionName() string { return "GetAddressRange" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAddressRange) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAddressRange) RefResponse() any { return &a.Response }

// GetAddressRangeRequest contains the "in" args for the "GetAddressRange" action.
type GetAddressRangeRequest struct{}

// GetAddressRangeResponse contains the "out" args for the "GetAddressRange" action.
type GetAddressRangeResponse struct {
	// NewMinAddress relates to state variable MinAddress.
	NewMinAddress string
	// NewMaxAddress relates to state variable MaxAddress.
	NewMaxAddress string
}

// GetDHCPRelay provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetDHCPRelay struct {
	Request  GetDHCPRelayRequest
	Response GetDHCPRelayResponse
}

var _ pkg1.Action = &GetDHCPRelay{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPRelay) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPRelay) ActionName() string { return "GetDHCPRelay" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPRelay) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPRelay) RefResponse() any { return &a.Response }

// GetDHCPRelayRequest contains the "in" args for the "GetDHCPRelay" action.
type GetDHCPRelayRequest struct{}

// GetDHCPRelayResponse contains the "out" args for the "GetDHCPRelay" action.
type GetDHCPRelayResponse struct {
	// NewDHCPRelay relates to state variable DHCPRelay.
	NewDHCPRelay pkg2.Boolean
}

// GetDHCPServerConfigurable provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetDHCPServerConfigurable struct {
	Request  GetDHCPServerConfigurableRequest
	Response GetDHCPServerConfigurableResponse
}

var _ pkg1.Action = &GetDHCPServerConfigurable{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPServerConfigurable) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPServerConfigurable) ActionName() string { return "GetDHCPServerConfigurable" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPServerConfigurable) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDHCPServerConfigurable) RefResponse() any { return &a.Response }

// GetDHCPServerConfigurableRequest contains the "in" args for the "GetDHCPServerConfigurable" action.
type GetDHCPServerConfigurableRequest struct{}

// GetDHCPServerConfigurableResponse contains the "out" args for the "GetDHCPServerConfigurable" action.
type GetDHCPServerConfigurableResponse struct {
	// NewDHCPServerConfigurable relates to state variable DHCPServerConfigurable.
	NewDHCPServerConfigurable pkg2.Boolean
}

// GetDNSServers provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetDNSServers struct {
	Request  GetDNSServersRequest
	Response GetDNSServersResponse
}

var _ pkg1.Action = &GetDNSServers{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDNSServers) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDNSServers) ActionName() string { return "GetDNSServers" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDNSServers) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDNSServers) RefResponse() any { return &a.Response }

// GetDNSServersRequest contains the "in" args for the "GetDNSServers" action.
type GetDNSServersRequest struct{}

// GetDNSServersResponse contains the "out" args for the "GetDNSServers" action.
type GetDNSServersResponse struct {
	// NewDNSServers relates to state variable DNSServers.
	NewDNSServers string
}

// GetDomainName provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetDomainName struct {
	Request  GetDomainNameRequest
	Response GetDomainNameResponse
}

var _ pkg1.Action = &GetDomainName{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDomainName) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDomainName) ActionName() string { return "GetDomainName" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDomainName) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetDomainName) RefResponse() any { return &a.Response }

// GetDomainNameRequest contains the "in" args for the "GetDomainName" action.
type GetDomainNameRequest struct{}

// GetDomainNameResponse contains the "out" args for the "GetDomainName" action.
type GetDomainNameResponse struct {
	// NewDomainName relates to state variable DomainName.
	NewDomainName string
}

// GetIPRoutersList provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetIPRoutersList struct {
	Request  GetIPRoutersListRequest
	Response GetIPRoutersListResponse
}

var _ pkg1.Action = &GetIPRoutersList{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIPRoutersList) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIPRoutersList) ActionName() string { return "GetIPRoutersList" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIPRoutersList) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIPRoutersList) RefResponse() any { return &a.Response }

// GetIPRoutersListRequest contains the "in" args for the "GetIPRoutersList" action.
type GetIPRoutersListRequest struct{}

// GetIPRoutersListResponse contains the "out" args for the "GetIPRoutersList" action.
type GetIPRoutersListResponse struct {
	// NewIPRouters relates to state variable IPRouters.
	NewIPRouters string
}

// GetReservedAddresses provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetReservedAddresses struct {
	Request  GetReservedAddressesRequest
	Response GetReservedAddressesResponse
}

var _ pkg1.Action = &GetReservedAddresses{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetReservedAddresses) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetReservedAddresses) ActionName() string { return "GetReservedAddresses" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetReservedAddresses) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetReservedAddresses) RefResponse() any { return &a.Response }

// GetReservedAddressesRequest contains the "in" args for the "GetReservedAddresses" action.
type GetReservedAddressesRequest struct{}

// GetReservedAddressesResponse contains the "out" args for the "GetReservedAddresses" action.
type GetReservedAddressesResponse struct {
	// NewReservedAddresses relates to state variable ReservedAddresses.
	NewReservedAddresses string
}

// GetSubnetMask provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetSubnetMask struct {
	Request  GetSubnetMaskRequest
	Response GetSubnetMaskResponse
}

var _ pkg1.Action = &GetSubnetMask{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSubnetMask) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSubnetMask) ActionName() string { return "GetSubnetMask" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSubnetMask) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSubnetMask) RefResponse() any { return &a.Response }

// GetSubnetMaskRequest contains the "in" args for the "GetSubnetMask" action.
type GetSubnetMaskRequest struct{}

// GetSubnetMaskResponse contains the "out" args for the "GetSubnetMask" action.
type GetSubnetMaskResponse struct {
	// NewSubnetMask relates to state variable SubnetMask.
	NewSubnetMask string
}

// SetAddressRange provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetAddressRange struct {
	Request  SetAddressRangeRequest
	Response SetAddressRangeResponse
}

var _ pkg1.Action = &SetAddressRange{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAddressRange) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAddressRange) ActionName() string { return "SetAddressRange" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAddressRange) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAddressRange) RefResponse() any { return &a.Response }

// SetAddressRangeRequest contains the "in" args for the "SetAddressRange" action.
type SetAddressRangeRequest struct {
	// NewMinAddress relates to state variable MinAddress.
	NewMinAddress string
	// NewMaxAddress relates to state variable MaxAddress.
	NewMaxAddress string
}

// SetAddressRangeResponse contains the "out" args for the "SetAddressRange" action.
type SetAddressRangeResponse struct{}

// SetDHCPRelay provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetDHCPRelay struct {
	Request  SetDHCPRelayRequest
	Response SetDHCPRelayResponse
}

var _ pkg1.Action = &SetDHCPRelay{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPRelay) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPRelay) ActionName() string { return "SetDHCPRelay" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPRelay) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPRelay) RefResponse() any { return &a.Response }

// SetDHCPRelayRequest contains the "in" args for the "SetDHCPRelay" action.
type SetDHCPRelayRequest struct {
	// NewDHCPRelay relates to state variable DHCPRelay.
	NewDHCPRelay pkg2.Boolean
}

// SetDHCPRelayResponse contains the "out" args for the "SetDHCPRelay" action.
type SetDHCPRelayResponse struct{}

// SetDHCPServerConfigurable provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetDHCPServerConfigurable struct {
	Request  SetDHCPServerConfigurableRequest
	Response SetDHCPServerConfigurableResponse
}

var _ pkg1.Action = &SetDHCPServerConfigurable{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPServerConfigurable) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPServerConfigurable) ActionName() string { return "SetDHCPServerConfigurable" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPServerConfigurable) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDHCPServerConfigurable) RefResponse() any { return &a.Response }

// SetDHCPServerConfigurableRequest contains the "in" args for the "SetDHCPServerConfigurable" action.
type SetDHCPServerConfigurableRequest struct {
	// NewDHCPServerConfigurable relates to state variable DHCPServerConfigurable.
	NewDHCPServerConfigurable pkg2.Boolean
}

// SetDHCPServerConfigurableResponse contains the "out" args for the "SetDHCPServerConfigurable" action.
type SetDHCPServerConfigurableResponse struct{}

// SetDNSServer provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetDNSServer struct {
	Request  SetDNSServerRequest
	Response SetDNSServerResponse
}

var _ pkg1.Action = &SetDNSServer{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDNSServer) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDNSServer) ActionName() string { return "SetDNSServer" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDNSServer) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDNSServer) RefResponse() any { return &a.Response }

// SetDNSServerRequest contains the "in" args for the "SetDNSServer" action.
type SetDNSServerRequest struct {
	// NewDNSServers relates to state variable DNSServers.
	NewDNSServers string
}

// SetDNSServerResponse contains the "out" args for the "SetDNSServer" action.
type SetDNSServerResponse struct{}

// SetDomainName provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetDomainName struct {
	Request  SetDomainNameRequest
	Response SetDomainNameResponse
}

var _ pkg1.Action = &SetDomainName{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDomainName) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDomainName) ActionName() string { return "SetDomainName" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDomainName) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetDomainName) RefResponse() any { return &a.Response }

// SetDomainNameRequest contains the "in" args for the "SetDomainName" action.
type SetDomainNameRequest struct {
	// NewDomainName relates to state variable DomainName.
	NewDomainName string
}

// SetDomainNameResponse contains the "out" args for the "SetDomainName" action.
type SetDomainNameResponse struct{}

// SetIPRouter provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetIPRouter struct {
	Request  SetIPRouterRequest
	Response SetIPRouterResponse
}

var _ pkg1.Action = &SetIPRouter{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIPRouter) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIPRouter) ActionName() string { return "SetIPRouter" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIPRouter) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIPRouter) RefResponse() any { return &a.Response }

// SetIPRouterRequest contains the "in" args for the "SetIPRouter" action.
type SetIPRouterRequest struct {
	// NewIPRouters relates to state variable IPRouters.
	NewIPRouters string
}

// SetIPRouterResponse contains the "out" args for the "SetIPRouter" action.
type SetIPRouterResponse struct{}

// SetReservedAddress provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetReservedAddress struct {
	Request  SetReservedAddressRequest
	Response SetReservedAddressResponse
}

var _ pkg1.Action = &SetReservedAddress{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetReservedAddress) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetReservedAddress) ActionName() string { return "SetReservedAddress" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetReservedAddress) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetReservedAddress) RefResponse() any { return &a.Response }

// SetReservedAddressRequest contains the "in" args for the "SetReservedAddress" action.
type SetReservedAddressRequest struct {
	// NewReservedAddresses relates to state variable ReservedAddresses.
	NewReservedAddresses string
}

// SetReservedAddressResponse contains the "out" args for the "SetReservedAddress" action.
type SetReservedAddressResponse struct{}

// SetSubnetMask provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetSubnetMask struct {
	Request  SetSubnetMaskRequest
	Response SetSubnetMaskResponse
}

var _ pkg1.Action = &SetSubnetMask{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetSubnetMask) ServiceType() string { return ServiceType }

// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetSubnetMask) ActionName() string { return "SetSubnetMask" }

// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetSubnetMask) RefRequest() any { return &a.Request }

// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetSubnetMask) RefResponse() any { return &a.Response }

// SetSubnetMaskRequest contains the "in" args for the "SetSubnetMask" action.
type SetSubnetMaskRequest struct {
	// NewSubnetMask relates to state variable SubnetMask.
	NewSubnetMask string
}

// SetSubnetMaskResponse contains the "out" args for the "SetSubnetMask" action.
type SetSubnetMaskResponse struct{}
