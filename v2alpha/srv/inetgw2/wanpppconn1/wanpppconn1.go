// Package wanpppconn1 provides types for the "urn:schemas-upnp-org:service:WANPPPConnection:1" service.
package wanpppconn1

import (
  pkg1 "github.com/huin/goupnp/v2alpha/soap"
  pkg2 "github.com/huin/goupnp/v2alpha/soap/types"
)

const ServiceType = "urn:schemas-upnp-org:service:WANPPPConnection:1"

// AddPortMapping provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type AddPortMapping struct{
  Request AddPortMappingRequest
  Response AddPortMappingResponse
}

var _ pkg1.Action = &AddPortMapping{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *AddPortMapping) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *AddPortMapping) ActionName() string { return "AddPortMapping" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *AddPortMapping) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *AddPortMapping) RefResponse() any { return &a.Response }

// AddPortMappingRequest contains the "in" args for the "AddPortMapping" action.
type AddPortMappingRequest struct{
  NewRemoteHost string
  NewExternalPort pkg2.UI2
  NewProtocol string
  NewInternalPort pkg2.UI2
  NewInternalClient string
  NewEnabled pkg2.Boolean
  NewPortMappingDescription string
  NewLeaseDuration pkg2.UI4
}

// AddPortMappingResponse contains the "out" args for the "AddPortMapping" action.
type AddPortMappingResponse struct{}

// ConfigureConnection provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type ConfigureConnection struct{
  Request ConfigureConnectionRequest
  Response ConfigureConnectionResponse
}

var _ pkg1.Action = &ConfigureConnection{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ConfigureConnection) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ConfigureConnection) ActionName() string { return "ConfigureConnection" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ConfigureConnection) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ConfigureConnection) RefResponse() any { return &a.Response }

// ConfigureConnectionRequest contains the "in" args for the "ConfigureConnection" action.
type ConfigureConnectionRequest struct{
  NewUserName string
  NewPassword string
}

// ConfigureConnectionResponse contains the "out" args for the "ConfigureConnection" action.
type ConfigureConnectionResponse struct{}

// DeletePortMapping provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type DeletePortMapping struct{
  Request DeletePortMappingRequest
  Response DeletePortMappingResponse
}

var _ pkg1.Action = &DeletePortMapping{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeletePortMapping) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeletePortMapping) ActionName() string { return "DeletePortMapping" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeletePortMapping) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *DeletePortMapping) RefResponse() any { return &a.Response }

// DeletePortMappingRequest contains the "in" args for the "DeletePortMapping" action.
type DeletePortMappingRequest struct{
  NewRemoteHost string
  NewExternalPort pkg2.UI2
  NewProtocol string
}

// DeletePortMappingResponse contains the "out" args for the "DeletePortMapping" action.
type DeletePortMappingResponse struct{}

// ForceTermination provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type ForceTermination struct{
  Request ForceTerminationRequest
  Response ForceTerminationResponse
}

var _ pkg1.Action = &ForceTermination{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ForceTermination) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ForceTermination) ActionName() string { return "ForceTermination" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ForceTermination) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *ForceTermination) RefResponse() any { return &a.Response }

// ForceTerminationRequest contains the "in" args for the "ForceTermination" action.
type ForceTerminationRequest struct{}

// ForceTerminationResponse contains the "out" args for the "ForceTermination" action.
type ForceTerminationResponse struct{}

// GetAutoDisconnectTime provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetAutoDisconnectTime struct{
  Request GetAutoDisconnectTimeRequest
  Response GetAutoDisconnectTimeResponse
}

var _ pkg1.Action = &GetAutoDisconnectTime{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAutoDisconnectTime) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAutoDisconnectTime) ActionName() string { return "GetAutoDisconnectTime" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAutoDisconnectTime) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetAutoDisconnectTime) RefResponse() any { return &a.Response }

// GetAutoDisconnectTimeRequest contains the "in" args for the "GetAutoDisconnectTime" action.
type GetAutoDisconnectTimeRequest struct{}

// GetAutoDisconnectTimeResponse contains the "out" args for the "GetAutoDisconnectTime" action.
type GetAutoDisconnectTimeResponse struct{
  NewAutoDisconnectTime pkg2.UI4
}

// GetConnectionTypeInfo provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetConnectionTypeInfo struct{
  Request GetConnectionTypeInfoRequest
  Response GetConnectionTypeInfoResponse
}

var _ pkg1.Action = &GetConnectionTypeInfo{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetConnectionTypeInfo) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetConnectionTypeInfo) ActionName() string { return "GetConnectionTypeInfo" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetConnectionTypeInfo) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetConnectionTypeInfo) RefResponse() any { return &a.Response }

// GetConnectionTypeInfoRequest contains the "in" args for the "GetConnectionTypeInfo" action.
type GetConnectionTypeInfoRequest struct{}

// GetConnectionTypeInfoResponse contains the "out" args for the "GetConnectionTypeInfo" action.
type GetConnectionTypeInfoResponse struct{
  NewConnectionType string
  NewPossibleConnectionTypes string
}

// GetExternalIPAddress provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetExternalIPAddress struct{
  Request GetExternalIPAddressRequest
  Response GetExternalIPAddressResponse
}

var _ pkg1.Action = &GetExternalIPAddress{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetExternalIPAddress) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetExternalIPAddress) ActionName() string { return "GetExternalIPAddress" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetExternalIPAddress) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetExternalIPAddress) RefResponse() any { return &a.Response }

// GetExternalIPAddressRequest contains the "in" args for the "GetExternalIPAddress" action.
type GetExternalIPAddressRequest struct{}

// GetExternalIPAddressResponse contains the "out" args for the "GetExternalIPAddress" action.
type GetExternalIPAddressResponse struct{
  NewExternalIPAddress string
}

// GetGenericPortMappingEntry provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetGenericPortMappingEntry struct{
  Request GetGenericPortMappingEntryRequest
  Response GetGenericPortMappingEntryResponse
}

var _ pkg1.Action = &GetGenericPortMappingEntry{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetGenericPortMappingEntry) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetGenericPortMappingEntry) ActionName() string { return "GetGenericPortMappingEntry" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetGenericPortMappingEntry) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetGenericPortMappingEntry) RefResponse() any { return &a.Response }

// GetGenericPortMappingEntryRequest contains the "in" args for the "GetGenericPortMappingEntry" action.
type GetGenericPortMappingEntryRequest struct{
  NewPortMappingIndex pkg2.UI2
}

// GetGenericPortMappingEntryResponse contains the "out" args for the "GetGenericPortMappingEntry" action.
type GetGenericPortMappingEntryResponse struct{
  NewRemoteHost string
  NewExternalPort pkg2.UI2
  NewProtocol string
  NewInternalPort pkg2.UI2
  NewInternalClient string
  NewEnabled pkg2.Boolean
  NewPortMappingDescription string
  NewLeaseDuration pkg2.UI4
}

// GetIdleDisconnectTime provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetIdleDisconnectTime struct{
  Request GetIdleDisconnectTimeRequest
  Response GetIdleDisconnectTimeResponse
}

var _ pkg1.Action = &GetIdleDisconnectTime{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIdleDisconnectTime) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIdleDisconnectTime) ActionName() string { return "GetIdleDisconnectTime" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIdleDisconnectTime) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetIdleDisconnectTime) RefResponse() any { return &a.Response }

// GetIdleDisconnectTimeRequest contains the "in" args for the "GetIdleDisconnectTime" action.
type GetIdleDisconnectTimeRequest struct{}

// GetIdleDisconnectTimeResponse contains the "out" args for the "GetIdleDisconnectTime" action.
type GetIdleDisconnectTimeResponse struct{
  NewIdleDisconnectTime pkg2.UI4
}

// GetLinkLayerMaxBitRates provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetLinkLayerMaxBitRates struct{
  Request GetLinkLayerMaxBitRatesRequest
  Response GetLinkLayerMaxBitRatesResponse
}

var _ pkg1.Action = &GetLinkLayerMaxBitRates{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetLinkLayerMaxBitRates) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetLinkLayerMaxBitRates) ActionName() string { return "GetLinkLayerMaxBitRates" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetLinkLayerMaxBitRates) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetLinkLayerMaxBitRates) RefResponse() any { return &a.Response }

// GetLinkLayerMaxBitRatesRequest contains the "in" args for the "GetLinkLayerMaxBitRates" action.
type GetLinkLayerMaxBitRatesRequest struct{}

// GetLinkLayerMaxBitRatesResponse contains the "out" args for the "GetLinkLayerMaxBitRates" action.
type GetLinkLayerMaxBitRatesResponse struct{
  NewUpstreamMaxBitRate pkg2.UI4
  NewDownstreamMaxBitRate pkg2.UI4
}

// GetNATRSIPStatus provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetNATRSIPStatus struct{
  Request GetNATRSIPStatusRequest
  Response GetNATRSIPStatusResponse
}

var _ pkg1.Action = &GetNATRSIPStatus{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetNATRSIPStatus) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetNATRSIPStatus) ActionName() string { return "GetNATRSIPStatus" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetNATRSIPStatus) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetNATRSIPStatus) RefResponse() any { return &a.Response }

// GetNATRSIPStatusRequest contains the "in" args for the "GetNATRSIPStatus" action.
type GetNATRSIPStatusRequest struct{}

// GetNATRSIPStatusResponse contains the "out" args for the "GetNATRSIPStatus" action.
type GetNATRSIPStatusResponse struct{
  NewRSIPAvailable pkg2.Boolean
  NewNATEnabled pkg2.Boolean
}

// GetPPPAuthenticationProtocol provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetPPPAuthenticationProtocol struct{
  Request GetPPPAuthenticationProtocolRequest
  Response GetPPPAuthenticationProtocolResponse
}

var _ pkg1.Action = &GetPPPAuthenticationProtocol{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPAuthenticationProtocol) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPAuthenticationProtocol) ActionName() string { return "GetPPPAuthenticationProtocol" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPAuthenticationProtocol) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPAuthenticationProtocol) RefResponse() any { return &a.Response }

// GetPPPAuthenticationProtocolRequest contains the "in" args for the "GetPPPAuthenticationProtocol" action.
type GetPPPAuthenticationProtocolRequest struct{}

// GetPPPAuthenticationProtocolResponse contains the "out" args for the "GetPPPAuthenticationProtocol" action.
type GetPPPAuthenticationProtocolResponse struct{
  NewPPPAuthenticationProtocol string
}

// GetPPPCompressionProtocol provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetPPPCompressionProtocol struct{
  Request GetPPPCompressionProtocolRequest
  Response GetPPPCompressionProtocolResponse
}

var _ pkg1.Action = &GetPPPCompressionProtocol{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPCompressionProtocol) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPCompressionProtocol) ActionName() string { return "GetPPPCompressionProtocol" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPCompressionProtocol) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPCompressionProtocol) RefResponse() any { return &a.Response }

// GetPPPCompressionProtocolRequest contains the "in" args for the "GetPPPCompressionProtocol" action.
type GetPPPCompressionProtocolRequest struct{}

// GetPPPCompressionProtocolResponse contains the "out" args for the "GetPPPCompressionProtocol" action.
type GetPPPCompressionProtocolResponse struct{
  NewPPPCompressionProtocol string
}

// GetPPPEncryptionProtocol provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetPPPEncryptionProtocol struct{
  Request GetPPPEncryptionProtocolRequest
  Response GetPPPEncryptionProtocolResponse
}

var _ pkg1.Action = &GetPPPEncryptionProtocol{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPEncryptionProtocol) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPEncryptionProtocol) ActionName() string { return "GetPPPEncryptionProtocol" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPEncryptionProtocol) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPPPEncryptionProtocol) RefResponse() any { return &a.Response }

// GetPPPEncryptionProtocolRequest contains the "in" args for the "GetPPPEncryptionProtocol" action.
type GetPPPEncryptionProtocolRequest struct{}

// GetPPPEncryptionProtocolResponse contains the "out" args for the "GetPPPEncryptionProtocol" action.
type GetPPPEncryptionProtocolResponse struct{
  NewPPPEncryptionProtocol string
}

// GetPassword provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetPassword struct{
  Request GetPasswordRequest
  Response GetPasswordResponse
}

var _ pkg1.Action = &GetPassword{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPassword) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPassword) ActionName() string { return "GetPassword" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPassword) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetPassword) RefResponse() any { return &a.Response }

// GetPasswordRequest contains the "in" args for the "GetPassword" action.
type GetPasswordRequest struct{}

// GetPasswordResponse contains the "out" args for the "GetPassword" action.
type GetPasswordResponse struct{
  NewPassword string
}

// GetSpecificPortMappingEntry provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetSpecificPortMappingEntry struct{
  Request GetSpecificPortMappingEntryRequest
  Response GetSpecificPortMappingEntryResponse
}

var _ pkg1.Action = &GetSpecificPortMappingEntry{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSpecificPortMappingEntry) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSpecificPortMappingEntry) ActionName() string { return "GetSpecificPortMappingEntry" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSpecificPortMappingEntry) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetSpecificPortMappingEntry) RefResponse() any { return &a.Response }

// GetSpecificPortMappingEntryRequest contains the "in" args for the "GetSpecificPortMappingEntry" action.
type GetSpecificPortMappingEntryRequest struct{
  NewRemoteHost string
  NewExternalPort pkg2.UI2
  NewProtocol string
}

// GetSpecificPortMappingEntryResponse contains the "out" args for the "GetSpecificPortMappingEntry" action.
type GetSpecificPortMappingEntryResponse struct{
  NewInternalPort pkg2.UI2
  NewInternalClient string
  NewEnabled pkg2.Boolean
  NewPortMappingDescription string
  NewLeaseDuration pkg2.UI4
}

// GetStatusInfo provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetStatusInfo struct{
  Request GetStatusInfoRequest
  Response GetStatusInfoResponse
}

var _ pkg1.Action = &GetStatusInfo{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetStatusInfo) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetStatusInfo) ActionName() string { return "GetStatusInfo" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetStatusInfo) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetStatusInfo) RefResponse() any { return &a.Response }

// GetStatusInfoRequest contains the "in" args for the "GetStatusInfo" action.
type GetStatusInfoRequest struct{}

// GetStatusInfoResponse contains the "out" args for the "GetStatusInfo" action.
type GetStatusInfoResponse struct{
  NewConnectionStatus string
  NewLastConnectionError string
  NewUptime pkg2.UI4
}

// GetUserName provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetUserName struct{
  Request GetUserNameRequest
  Response GetUserNameResponse
}

var _ pkg1.Action = &GetUserName{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetUserName) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetUserName) ActionName() string { return "GetUserName" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetUserName) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetUserName) RefResponse() any { return &a.Response }

// GetUserNameRequest contains the "in" args for the "GetUserName" action.
type GetUserNameRequest struct{}

// GetUserNameResponse contains the "out" args for the "GetUserName" action.
type GetUserNameResponse struct{
  NewUserName string
}

// GetWarnDisconnectDelay provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type GetWarnDisconnectDelay struct{
  Request GetWarnDisconnectDelayRequest
  Response GetWarnDisconnectDelayResponse
}

var _ pkg1.Action = &GetWarnDisconnectDelay{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetWarnDisconnectDelay) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetWarnDisconnectDelay) ActionName() string { return "GetWarnDisconnectDelay" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetWarnDisconnectDelay) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *GetWarnDisconnectDelay) RefResponse() any { return &a.Response }

// GetWarnDisconnectDelayRequest contains the "in" args for the "GetWarnDisconnectDelay" action.
type GetWarnDisconnectDelayRequest struct{}

// GetWarnDisconnectDelayResponse contains the "out" args for the "GetWarnDisconnectDelay" action.
type GetWarnDisconnectDelayResponse struct{
  NewWarnDisconnectDelay pkg2.UI4
}

// RequestConnection provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type RequestConnection struct{
  Request RequestConnectionRequest
  Response RequestConnectionResponse
}

var _ pkg1.Action = &RequestConnection{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestConnection) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestConnection) ActionName() string { return "RequestConnection" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestConnection) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestConnection) RefResponse() any { return &a.Response }

// RequestConnectionRequest contains the "in" args for the "RequestConnection" action.
type RequestConnectionRequest struct{}

// RequestConnectionResponse contains the "out" args for the "RequestConnection" action.
type RequestConnectionResponse struct{}

// RequestTermination provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type RequestTermination struct{
  Request RequestTerminationRequest
  Response RequestTerminationResponse
}

var _ pkg1.Action = &RequestTermination{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestTermination) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestTermination) ActionName() string { return "RequestTermination" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestTermination) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *RequestTermination) RefResponse() any { return &a.Response }

// RequestTerminationRequest contains the "in" args for the "RequestTermination" action.
type RequestTerminationRequest struct{}

// RequestTerminationResponse contains the "out" args for the "RequestTermination" action.
type RequestTerminationResponse struct{}

// SetAutoDisconnectTime provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetAutoDisconnectTime struct{
  Request SetAutoDisconnectTimeRequest
  Response SetAutoDisconnectTimeResponse
}

var _ pkg1.Action = &SetAutoDisconnectTime{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAutoDisconnectTime) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAutoDisconnectTime) ActionName() string { return "SetAutoDisconnectTime" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAutoDisconnectTime) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetAutoDisconnectTime) RefResponse() any { return &a.Response }

// SetAutoDisconnectTimeRequest contains the "in" args for the "SetAutoDisconnectTime" action.
type SetAutoDisconnectTimeRequest struct{
  NewAutoDisconnectTime pkg2.UI4
}

// SetAutoDisconnectTimeResponse contains the "out" args for the "SetAutoDisconnectTime" action.
type SetAutoDisconnectTimeResponse struct{}

// SetConnectionType provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetConnectionType struct{
  Request SetConnectionTypeRequest
  Response SetConnectionTypeResponse
}

var _ pkg1.Action = &SetConnectionType{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetConnectionType) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetConnectionType) ActionName() string { return "SetConnectionType" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetConnectionType) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetConnectionType) RefResponse() any { return &a.Response }

// SetConnectionTypeRequest contains the "in" args for the "SetConnectionType" action.
type SetConnectionTypeRequest struct{
  NewConnectionType string
}

// SetConnectionTypeResponse contains the "out" args for the "SetConnectionType" action.
type SetConnectionTypeResponse struct{}

// SetIdleDisconnectTime provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetIdleDisconnectTime struct{
  Request SetIdleDisconnectTimeRequest
  Response SetIdleDisconnectTimeResponse
}

var _ pkg1.Action = &SetIdleDisconnectTime{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIdleDisconnectTime) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIdleDisconnectTime) ActionName() string { return "SetIdleDisconnectTime" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIdleDisconnectTime) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetIdleDisconnectTime) RefResponse() any { return &a.Response }

// SetIdleDisconnectTimeRequest contains the "in" args for the "SetIdleDisconnectTime" action.
type SetIdleDisconnectTimeRequest struct{
  NewIdleDisconnectTime pkg2.UI4
}

// SetIdleDisconnectTimeResponse contains the "out" args for the "SetIdleDisconnectTime" action.
type SetIdleDisconnectTimeResponse struct{}

// SetWarnDisconnectDelay provides request and response for the action.
//
// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action, self-describing the SOAP action.
type SetWarnDisconnectDelay struct{
  Request SetWarnDisconnectDelayRequest
  Response SetWarnDisconnectDelayResponse
}

var _ pkg1.Action = &SetWarnDisconnectDelay{}

// ServiceType implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetWarnDisconnectDelay) ServiceType() string { return ServiceType }
// ActionName implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetWarnDisconnectDelay) ActionName() string { return "SetWarnDisconnectDelay" }
// RefRequest implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetWarnDisconnectDelay) RefRequest() any { return &a.Request }
// RefResponse implements "github.com/huin/goupnp/v2alpha/soap".Action.
func (a *SetWarnDisconnectDelay) RefResponse() any { return &a.Response }

// SetWarnDisconnectDelayRequest contains the "in" args for the "SetWarnDisconnectDelay" action.
type SetWarnDisconnectDelayRequest struct{
  NewWarnDisconnectDelay pkg2.UI4
}

// SetWarnDisconnectDelayResponse contains the "out" args for the "SetWarnDisconnectDelay" action.
type SetWarnDisconnectDelayResponse struct{}
