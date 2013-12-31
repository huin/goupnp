package internetgateway2

import (
	"time"

	"github.com/huin/goupnp/soap"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

const (
	URN_LANDevice_1 = "urn:schemas-upnp-org:device:LANDevice:1"

	URN_WANConnectionDevice_1 = "urn:schemas-upnp-org:device:WANConnectionDevice:1"

	URN_WANConnectionDevice_2 = "urn:schemas-upnp-org:device:WANConnectionDevice:2"

	URN_WANDevice_1 = "urn:schemas-upnp-org:device:WANDevice:1"

	URN_WANDevice_2 = "urn:schemas-upnp-org:device:WANDevice:2"
)

const (
	URN_LANHostConfigManagement_1 = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"

	URN_Layer3Forwarding_1 = "urn:schemas-upnp-org:service:Layer3Forwarding:1"

	URN_WANCableLinkConfig_1 = "urn:schemas-upnp-org:service:WANCableLinkConfig:1"

	URN_WANCommonInterfaceConfig_1 = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"

	URN_WANDSLLinkConfig_1 = "urn:schemas-upnp-org:service:WANDSLLinkConfig:1"

	URN_WANEthernetLinkConfig_1 = "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1"

	URN_WANIPConnection_1 = "urn:schemas-upnp-org:service:WANIPConnection:1"

	URN_WANIPConnection_2 = "urn:schemas-upnp-org:service:WANIPConnection:2"

	URN_WANIPv6FirewallControl_1 = "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1"

	URN_WANPOTSLinkConfig_1 = "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1"

	URN_WANPPPConnection_1 = "urn:schemas-upnp-org:service:WANPPPConnection:1"
)

// LANHostConfigManagement1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:LANHostConfigManagement:1".
type LANHostConfigManagement1 struct {
	SOAPClient soap.SOAPClient
}

// SetDHCPServerConfigurable action.
// Arguments:
//
// * NewDHCPServerConfigurable:
// (related state variable: DHCPServerConfigurable)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetDHCPServerConfigurable(
	NewDHCPServerConfigurable bool,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDHCPServerConfigurable string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDHCPServerConfigurable, err = soap.MarshalBoolean(NewDHCPServerConfigurable); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPServerConfigurable", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDHCPServerConfigurable action.
// Arguments:
//
//
// Return values:
//
// * NewDHCPServerConfigurable:
// (related state variable: DHCPServerConfigurable)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetDHCPServerConfigurable() (
	NewDHCPServerConfigurable bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDHCPServerConfigurable string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPServerConfigurable", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDHCPServerConfigurable, err = soap.UnmarshalBoolean(response.NewDHCPServerConfigurable); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetDHCPRelay action.
// Arguments:
//
// * NewDHCPRelay:
// (related state variable: DHCPRelay)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetDHCPRelay(
	NewDHCPRelay bool,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDHCPRelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDHCPRelay, err = soap.MarshalBoolean(NewDHCPRelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPRelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDHCPRelay action.
// Arguments:
//
//
// Return values:
//
// * NewDHCPRelay:
// (related state variable: DHCPRelay)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetDHCPRelay() (
	NewDHCPRelay bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDHCPRelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPRelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDHCPRelay, err = soap.UnmarshalBoolean(response.NewDHCPRelay); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetSubnetMask action.
// Arguments:
//
// * NewSubnetMask:
// (related state variable: SubnetMask)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetSubnetMask(
	NewSubnetMask string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewSubnetMask string
	}
	// BEGIN Marshal arguments into request.

	if request.NewSubnetMask, err = soap.MarshalString(NewSubnetMask); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetSubnetMask", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetSubnetMask action.
// Arguments:
//
//
// Return values:
//
// * NewSubnetMask:
// (related state variable: SubnetMask)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetSubnetMask() (
	NewSubnetMask string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewSubnetMask string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetSubnetMask", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewSubnetMask, err = soap.UnmarshalString(response.NewSubnetMask); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetIPRouter action.
// Arguments:
//
// * NewIPRouters:
// (related state variable: IPRouters)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetIPRouter(
	NewIPRouters string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewIPRouters string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIPRouters, err = soap.MarshalString(NewIPRouters); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetIPRouter", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteIPRouter action.
// Arguments:
//
// * NewIPRouters:
// (related state variable: IPRouters)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) DeleteIPRouter(
	NewIPRouters string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewIPRouters string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIPRouters, err = soap.MarshalString(NewIPRouters); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteIPRouter", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetIPRoutersList action.
// Arguments:
//
//
// Return values:
//
// * NewIPRouters:
// (related state variable: IPRouters)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetIPRoutersList() (
	NewIPRouters string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIPRouters string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetIPRoutersList", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIPRouters, err = soap.UnmarshalString(response.NewIPRouters); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetDomainName action.
// Arguments:
//
// * NewDomainName:
// (related state variable: DomainName)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetDomainName(
	NewDomainName string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDomainName string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDomainName, err = soap.MarshalString(NewDomainName); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDomainName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDomainName action.
// Arguments:
//
//
// Return values:
//
// * NewDomainName:
// (related state variable: DomainName)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetDomainName() (
	NewDomainName string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDomainName string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDomainName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDomainName, err = soap.UnmarshalString(response.NewDomainName); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetAddressRange action.
// Arguments:
//
// * NewMinAddress:
// (related state variable: MinAddress)
// -
// -
//
//
// * NewMaxAddress:
// (related state variable: MaxAddress)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetAddressRange(
	NewMinAddress string,
	NewMaxAddress string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewMinAddress string
		NewMaxAddress string
	}
	// BEGIN Marshal arguments into request.

	if request.NewMinAddress, err = soap.MarshalString(NewMinAddress); err != nil {
		return
	}

	if request.NewMaxAddress, err = soap.MarshalString(NewMaxAddress); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetAddressRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetAddressRange action.
// Arguments:
//
//
// Return values:
//
// * NewMinAddress:
// (related state variable: MinAddress)
// -
// -
//
//
// * NewMaxAddress:
// (related state variable: MaxAddress)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetAddressRange() (
	NewMinAddress string,
	NewMaxAddress string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewMinAddress string
		NewMaxAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetAddressRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewMinAddress, err = soap.UnmarshalString(response.NewMinAddress); err != nil {
		return
	}

	if NewMaxAddress, err = soap.UnmarshalString(response.NewMaxAddress); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetReservedAddress action.
// Arguments:
//
// * NewReservedAddresses:
// (related state variable: ReservedAddresses)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetReservedAddress(
	NewReservedAddresses string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewReservedAddresses string
	}
	// BEGIN Marshal arguments into request.

	if request.NewReservedAddresses, err = soap.MarshalString(NewReservedAddresses); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetReservedAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteReservedAddress action.
// Arguments:
//
// * NewReservedAddresses:
// (related state variable: ReservedAddresses)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) DeleteReservedAddress(
	NewReservedAddresses string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewReservedAddresses string
	}
	// BEGIN Marshal arguments into request.

	if request.NewReservedAddresses, err = soap.MarshalString(NewReservedAddresses); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteReservedAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetReservedAddresses action.
// Arguments:
//
//
// Return values:
//
// * NewReservedAddresses:
// (related state variable: ReservedAddresses)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetReservedAddresses() (
	NewReservedAddresses string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewReservedAddresses string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetReservedAddresses", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewReservedAddresses, err = soap.UnmarshalString(response.NewReservedAddresses); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetDNSServer action.
// Arguments:
//
// * NewDNSServers:
// (related state variable: DNSServers)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) SetDNSServer(
	NewDNSServers string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDNSServers string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDNSServers, err = soap.MarshalString(NewDNSServers); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDNSServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteDNSServer action.
// Arguments:
//
// * NewDNSServers:
// (related state variable: DNSServers)
// -
// -
//
//
//
// Return values:
//
func (client *LANHostConfigManagement1) DeleteDNSServer(
	NewDNSServers string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDNSServers string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDNSServers, err = soap.MarshalString(NewDNSServers); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteDNSServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDNSServers action.
// Arguments:
//
//
// Return values:
//
// * NewDNSServers:
// (related state variable: DNSServers)
// -
// -
//
//
func (client *LANHostConfigManagement1) GetDNSServers() (
	NewDNSServers string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDNSServers string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDNSServers", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDNSServers, err = soap.UnmarshalString(response.NewDNSServers); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// Layer3Forwarding1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:Layer3Forwarding:1".
type Layer3Forwarding1 struct {
	SOAPClient soap.SOAPClient
}

// SetDefaultConnectionService action.
// Arguments:
//
// * NewDefaultConnectionService:
// (related state variable: DefaultConnectionService)
// -
// -
//
//
//
// Return values:
//
func (client *Layer3Forwarding1) SetDefaultConnectionService(
	NewDefaultConnectionService string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDefaultConnectionService string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDefaultConnectionService, err = soap.MarshalString(NewDefaultConnectionService); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "SetDefaultConnectionService", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDefaultConnectionService action.
// Arguments:
//
//
// Return values:
//
// * NewDefaultConnectionService:
// (related state variable: DefaultConnectionService)
// -
// -
//
//
func (client *Layer3Forwarding1) GetDefaultConnectionService() (
	NewDefaultConnectionService string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDefaultConnectionService string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "GetDefaultConnectionService", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDefaultConnectionService, err = soap.UnmarshalString(response.NewDefaultConnectionService); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANCableLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCableLinkConfig:1".
type WANCableLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

// GetCableLinkConfigInfo action.
// Arguments:
//
//
// Return values:
//
// * NewCableLinkConfigState:
// (related state variable: CableLinkConfigState)
// -
// - allowed values:
// notReady|dsSyncComplete|usParamAcquired|rangingComplete|ipComplete|todEstablished|paramTransferComplete|registrationComplete|operational|accessDenied
//
//
// * NewLinkType:
// (related state variable: LinkType)
// -
// - allowed values:
// Ethernet
//
//
func (client *WANCableLinkConfig1) GetCableLinkConfigInfo() (
	NewCableLinkConfigState string,
	NewLinkType string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewCableLinkConfigState string
		NewLinkType             string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetCableLinkConfigInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewCableLinkConfigState, err = soap.UnmarshalString(response.NewCableLinkConfigState); err != nil {
		return
	}

	if NewLinkType, err = soap.UnmarshalString(response.NewLinkType); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetDownstreamFrequency action.
// Arguments:
//
//
// Return values:
//
// * NewDownstreamFrequency:
// (related state variable: DownstreamFrequency)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetDownstreamFrequency() (
	NewDownstreamFrequency uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDownstreamFrequency string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamFrequency", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDownstreamFrequency, err = soap.UnmarshalUi4(response.NewDownstreamFrequency); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetDownstreamModulation action.
// Arguments:
//
//
// Return values:
//
// * NewDownstreamModulation:
// (related state variable: DownstreamModulation)
// -
// - allowed values:
// 64QAM|256QAM
//
//
func (client *WANCableLinkConfig1) GetDownstreamModulation() (
	NewDownstreamModulation string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDownstreamModulation string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamModulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDownstreamModulation, err = soap.UnmarshalString(response.NewDownstreamModulation); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetUpstreamFrequency action.
// Arguments:
//
//
// Return values:
//
// * NewUpstreamFrequency:
// (related state variable: UpstreamFrequency)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetUpstreamFrequency() (
	NewUpstreamFrequency uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamFrequency string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamFrequency", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamFrequency, err = soap.UnmarshalUi4(response.NewUpstreamFrequency); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetUpstreamModulation action.
// Arguments:
//
//
// Return values:
//
// * NewUpstreamModulation:
// (related state variable: UpstreamModulation)
// -
// - allowed values:
// QPSK|16QAM
//
//
func (client *WANCableLinkConfig1) GetUpstreamModulation() (
	NewUpstreamModulation string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamModulation string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamModulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamModulation, err = soap.UnmarshalString(response.NewUpstreamModulation); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetUpstreamChannelID action.
// Arguments:
//
//
// Return values:
//
// * NewUpstreamChannelID:
// (related state variable: UpstreamChannelID)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetUpstreamChannelID() (
	NewUpstreamChannelID uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamChannelID string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamChannelID", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamChannelID, err = soap.UnmarshalUi4(response.NewUpstreamChannelID); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetUpstreamPowerLevel action.
// Arguments:
//
//
// Return values:
//
// * NewUpstreamPowerLevel:
// (related state variable: UpstreamPowerLevel)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetUpstreamPowerLevel() (
	NewUpstreamPowerLevel uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamPowerLevel string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamPowerLevel", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamPowerLevel, err = soap.UnmarshalUi4(response.NewUpstreamPowerLevel); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetBPIEncryptionEnabled action.
// Arguments:
//
//
// Return values:
//
// * NewBPIEncryptionEnabled:
// (related state variable: BPIEncryptionEnabled)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetBPIEncryptionEnabled() (
	NewBPIEncryptionEnabled bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewBPIEncryptionEnabled string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetBPIEncryptionEnabled", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewBPIEncryptionEnabled, err = soap.UnmarshalBoolean(response.NewBPIEncryptionEnabled); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetConfigFile action.
// Arguments:
//
//
// Return values:
//
// * NewConfigFile:
// (related state variable: ConfigFile)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetConfigFile() (
	NewConfigFile string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConfigFile string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetConfigFile", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConfigFile, err = soap.UnmarshalString(response.NewConfigFile); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetTFTPServer action.
// Arguments:
//
//
// Return values:
//
// * NewTFTPServer:
// (related state variable: TFTPServer)
// -
// -
//
//
func (client *WANCableLinkConfig1) GetTFTPServer() (
	NewTFTPServer string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTFTPServer string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetTFTPServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTFTPServer, err = soap.UnmarshalString(response.NewTFTPServer); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANCommonInterfaceConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1".
type WANCommonInterfaceConfig1 struct {
	SOAPClient soap.SOAPClient
}

// SetEnabledForInternet action.
// Arguments:
//
// * NewEnabledForInternet:
// (related state variable: EnabledForInternet)
// -
// -
//
//
//
// Return values:
//
func (client *WANCommonInterfaceConfig1) SetEnabledForInternet(
	NewEnabledForInternet bool,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewEnabledForInternet string
	}
	// BEGIN Marshal arguments into request.

	if request.NewEnabledForInternet, err = soap.MarshalBoolean(NewEnabledForInternet); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "SetEnabledForInternet", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetEnabledForInternet action.
// Arguments:
//
//
// Return values:
//
// * NewEnabledForInternet:
// (related state variable: EnabledForInternet)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetEnabledForInternet() (
	NewEnabledForInternet bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewEnabledForInternet string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetEnabledForInternet", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewEnabledForInternet, err = soap.UnmarshalBoolean(response.NewEnabledForInternet); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetCommonLinkProperties action.
// Arguments:
//
//
// Return values:
//
// * NewWANAccessType:
// (related state variable: WANAccessType)
// -
// - allowed values:
// DSL|POTS|Cable|Ethernet
//
//
// * NewLayer1UpstreamMaxBitRate:
// (related state variable: Layer1UpstreamMaxBitRate)
// -
// -
//
//
// * NewLayer1DownstreamMaxBitRate:
// (related state variable: Layer1DownstreamMaxBitRate)
// -
// -
//
//
// * NewPhysicalLinkStatus:
// (related state variable: PhysicalLinkStatus)
// -
// - allowed values:
// Up|Down
//
//
func (client *WANCommonInterfaceConfig1) GetCommonLinkProperties() (
	NewWANAccessType string,
	NewLayer1UpstreamMaxBitRate uint32,
	NewLayer1DownstreamMaxBitRate uint32,
	NewPhysicalLinkStatus string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWANAccessType              string
		NewLayer1UpstreamMaxBitRate   string
		NewLayer1DownstreamMaxBitRate string
		NewPhysicalLinkStatus         string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetCommonLinkProperties", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWANAccessType, err = soap.UnmarshalString(response.NewWANAccessType); err != nil {
		return
	}

	if NewLayer1UpstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewLayer1UpstreamMaxBitRate); err != nil {
		return
	}

	if NewLayer1DownstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewLayer1DownstreamMaxBitRate); err != nil {
		return
	}

	if NewPhysicalLinkStatus, err = soap.UnmarshalString(response.NewPhysicalLinkStatus); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetWANAccessProvider action.
// Arguments:
//
//
// Return values:
//
// * NewWANAccessProvider:
// (related state variable: WANAccessProvider)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetWANAccessProvider() (
	NewWANAccessProvider string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWANAccessProvider string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetWANAccessProvider", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWANAccessProvider, err = soap.UnmarshalString(response.NewWANAccessProvider); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetMaximumActiveConnections action.
// Arguments:
//
//
// Return values:
//
// * NewMaximumActiveConnections:
// (related state variable: MaximumActiveConnections)
// - allowed range: 1 to
// -
//
//
func (client *WANCommonInterfaceConfig1) GetMaximumActiveConnections() (
	NewMaximumActiveConnections uint16,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewMaximumActiveConnections string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetMaximumActiveConnections", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewMaximumActiveConnections, err = soap.UnmarshalUi2(response.NewMaximumActiveConnections); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetTotalBytesSent action.
// Arguments:
//
//
// Return values:
//
// * NewTotalBytesSent:
// (related state variable: TotalBytesSent)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetTotalBytesSent() (
	NewTotalBytesSent uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalBytesSent string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesSent", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalBytesSent, err = soap.UnmarshalUi4(response.NewTotalBytesSent); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetTotalBytesReceived action.
// Arguments:
//
//
// Return values:
//
// * NewTotalBytesReceived:
// (related state variable: TotalBytesReceived)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetTotalBytesReceived() (
	NewTotalBytesReceived uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalBytesReceived string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesReceived", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalBytesReceived, err = soap.UnmarshalUi4(response.NewTotalBytesReceived); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetTotalPacketsSent action.
// Arguments:
//
//
// Return values:
//
// * NewTotalPacketsSent:
// (related state variable: TotalPacketsSent)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetTotalPacketsSent() (
	NewTotalPacketsSent uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalPacketsSent string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsSent", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalPacketsSent, err = soap.UnmarshalUi4(response.NewTotalPacketsSent); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetTotalPacketsReceived action.
// Arguments:
//
//
// Return values:
//
// * NewTotalPacketsReceived:
// (related state variable: TotalPacketsReceived)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetTotalPacketsReceived() (
	NewTotalPacketsReceived uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalPacketsReceived string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsReceived", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalPacketsReceived, err = soap.UnmarshalUi4(response.NewTotalPacketsReceived); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetActiveConnection action.
// Arguments:
//
// * NewActiveConnectionIndex:
// (related state variable: NumberOfActiveConnections)
// -
// -
//
//
//
// Return values:
//
// * NewActiveConnDeviceContainer:
// (related state variable: ActiveConnectionDeviceContainer)
// -
// -
//
//
// * NewActiveConnectionServiceID:
// (related state variable: ActiveConnectionServiceID)
// -
// -
//
//
func (client *WANCommonInterfaceConfig1) GetActiveConnection(
	NewActiveConnectionIndex uint16,
) (
	NewActiveConnDeviceContainer string,
	NewActiveConnectionServiceID string,
	err error,
) {
	// Request structure.
	var request struct {
		NewActiveConnectionIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewActiveConnectionIndex, err = soap.MarshalUi2(NewActiveConnectionIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewActiveConnDeviceContainer string
		NewActiveConnectionServiceID string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetActiveConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewActiveConnDeviceContainer, err = soap.UnmarshalString(response.NewActiveConnDeviceContainer); err != nil {
		return
	}

	if NewActiveConnectionServiceID, err = soap.UnmarshalString(response.NewActiveConnectionServiceID); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANDSLLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANDSLLinkConfig:1".
type WANDSLLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

// SetDSLLinkType action.
// Arguments:
//
// * NewLinkType:
// (related state variable: LinkType)
// -
// -
//
//
//
// Return values:
//
func (client *WANDSLLinkConfig1) SetDSLLinkType(
	NewLinkType string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewLinkType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewLinkType, err = soap.MarshalString(NewLinkType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDSLLinkType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDSLLinkInfo action.
// Arguments:
//
//
// Return values:
//
// * NewLinkType:
// (related state variable: LinkType)
// -
// -
//
//
// * NewLinkStatus:
// (related state variable: LinkStatus)
// -
// - allowed values:
// Up|Down
//
//
func (client *WANDSLLinkConfig1) GetDSLLinkInfo() (
	NewLinkType string,
	NewLinkStatus string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewLinkType   string
		NewLinkStatus string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDSLLinkInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewLinkType, err = soap.UnmarshalString(response.NewLinkType); err != nil {
		return
	}

	if NewLinkStatus, err = soap.UnmarshalString(response.NewLinkStatus); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetAutoConfig action.
// Arguments:
//
//
// Return values:
//
// * NewAutoConfig:
// (related state variable: AutoConfig)
// -
// -
//
//
func (client *WANDSLLinkConfig1) GetAutoConfig() (
	NewAutoConfig bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoConfig string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetAutoConfig", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoConfig, err = soap.UnmarshalBoolean(response.NewAutoConfig); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetModulationType action.
// Arguments:
//
//
// Return values:
//
// * NewModulationType:
// (related state variable: ModulationType)
// -
// -
//
//
func (client *WANDSLLinkConfig1) GetModulationType() (
	NewModulationType string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewModulationType string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetModulationType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewModulationType, err = soap.UnmarshalString(response.NewModulationType); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetDestinationAddress action.
// Arguments:
//
// * NewDestinationAddress:
// (related state variable: DestinationAddress)
// -
// -
//
//
//
// Return values:
//
func (client *WANDSLLinkConfig1) SetDestinationAddress(
	NewDestinationAddress string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewDestinationAddress string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDestinationAddress, err = soap.MarshalString(NewDestinationAddress); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDestinationAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetDestinationAddress action.
// Arguments:
//
//
// Return values:
//
// * NewDestinationAddress:
// (related state variable: DestinationAddress)
// -
// -
//
//
func (client *WANDSLLinkConfig1) GetDestinationAddress() (
	NewDestinationAddress string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDestinationAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDestinationAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDestinationAddress, err = soap.UnmarshalString(response.NewDestinationAddress); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetATMEncapsulation action.
// Arguments:
//
// * NewATMEncapsulation:
// (related state variable: ATMEncapsulation)
// -
// -
//
//
//
// Return values:
//
func (client *WANDSLLinkConfig1) SetATMEncapsulation(
	NewATMEncapsulation string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewATMEncapsulation string
	}
	// BEGIN Marshal arguments into request.

	if request.NewATMEncapsulation, err = soap.MarshalString(NewATMEncapsulation); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetATMEncapsulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetATMEncapsulation action.
// Arguments:
//
//
// Return values:
//
// * NewATMEncapsulation:
// (related state variable: ATMEncapsulation)
// -
// -
//
//
func (client *WANDSLLinkConfig1) GetATMEncapsulation() (
	NewATMEncapsulation string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewATMEncapsulation string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetATMEncapsulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewATMEncapsulation, err = soap.UnmarshalString(response.NewATMEncapsulation); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// SetFCSPreserved action.
// Arguments:
//
// * NewFCSPreserved:
// (related state variable: FCSPreserved)
// -
// -
//
//
//
// Return values:
//
func (client *WANDSLLinkConfig1) SetFCSPreserved(
	NewFCSPreserved bool,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewFCSPreserved string
	}
	// BEGIN Marshal arguments into request.

	if request.NewFCSPreserved, err = soap.MarshalBoolean(NewFCSPreserved); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetFCSPreserved", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetFCSPreserved action.
// Arguments:
//
//
// Return values:
//
// * NewFCSPreserved:
// (related state variable: FCSPreserved)
// -
// -
//
//
func (client *WANDSLLinkConfig1) GetFCSPreserved() (
	NewFCSPreserved bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewFCSPreserved string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetFCSPreserved", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewFCSPreserved, err = soap.UnmarshalBoolean(response.NewFCSPreserved); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANEthernetLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1".
type WANEthernetLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

// GetEthernetLinkStatus action.
// Arguments:
//
//
// Return values:
//
// * NewEthernetLinkStatus:
// (related state variable: EthernetLinkStatus)
// -
// - allowed values:
// Up|Down
//
//
func (client *WANEthernetLinkConfig1) GetEthernetLinkStatus() (
	NewEthernetLinkStatus string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewEthernetLinkStatus string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANEthernetLinkConfig_1, "GetEthernetLinkStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewEthernetLinkStatus, err = soap.UnmarshalString(response.NewEthernetLinkStatus); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANIPConnection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPConnection:1".
type WANIPConnection1 struct {
	SOAPClient soap.SOAPClient
}

// SetConnectionType action.
// Arguments:
//
// * NewConnectionType:
// (related state variable: ConnectionType)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection1) SetConnectionType(
	NewConnectionType string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewConnectionType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetConnectionTypeInfo action.
// Arguments:
//
//
// Return values:
//
// * NewConnectionType:
// (related state variable: ConnectionType)
// -
// -
//
//
// * NewPossibleConnectionTypes:
// (related state variable: PossibleConnectionTypes)
// -
// - allowed values:
// Unconfigured|IP_Routed|IP_Bridged
//
//
func (client *WANIPConnection1) GetConnectionTypeInfo() (
	NewConnectionType string,
	NewPossibleConnectionTypes string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionType          string
		NewPossibleConnectionTypes string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetConnectionTypeInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionType, err = soap.UnmarshalString(response.NewConnectionType); err != nil {
		return
	}

	if NewPossibleConnectionTypes, err = soap.UnmarshalString(response.NewPossibleConnectionTypes); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) RequestConnection() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) RequestTermination() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) ForceTermination() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAutoDisconnectTime action.
// Arguments:
//
// * NewAutoDisconnectTime:
// (related state variable: AutoDisconnectTime)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection1) SetAutoDisconnectTime(
	NewAutoDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewAutoDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetIdleDisconnectTime action.
// Arguments:
//
// * NewIdleDisconnectTime:
// (related state variable: IdleDisconnectTime)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection1) SetIdleDisconnectTime(
	NewIdleDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewIdleDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetWarnDisconnectDelay action.
// Arguments:
//
// * NewWarnDisconnectDelay:
// (related state variable: WarnDisconnectDelay)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection1) SetWarnDisconnectDelay(
	NewWarnDisconnectDelay uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewWarnDisconnectDelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetStatusInfo action.
// Arguments:
//
//
// Return values:
//
// * NewConnectionStatus:
// (related state variable: ConnectionStatus)
// -
// - allowed values:
// Unconfigured|Connected|Disconnected
//
//
// * NewLastConnectionError:
// (related state variable: LastConnectionError)
// -
// - allowed values:
// ERROR_NONE
//
//
// * NewUptime:
// (related state variable: Uptime)
// -
// -
//
//
func (client *WANIPConnection1) GetStatusInfo() (
	NewConnectionStatus string,
	NewLastConnectionError string,
	NewUptime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionStatus    string
		NewLastConnectionError string
		NewUptime              string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetStatusInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionStatus, err = soap.UnmarshalString(response.NewConnectionStatus); err != nil {
		return
	}

	if NewLastConnectionError, err = soap.UnmarshalString(response.NewLastConnectionError); err != nil {
		return
	}

	if NewUptime, err = soap.UnmarshalUi4(response.NewUptime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetAutoDisconnectTime action.
// Arguments:
//
//
// Return values:
//
// * NewAutoDisconnectTime:
// (related state variable: AutoDisconnectTime)
// -
// -
//
//
func (client *WANIPConnection1) GetAutoDisconnectTime() (
	NewAutoDisconnectTime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoDisconnectTime, err = soap.UnmarshalUi4(response.NewAutoDisconnectTime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetIdleDisconnectTime action.
// Arguments:
//
//
// Return values:
//
// * NewIdleDisconnectTime:
// (related state variable: IdleDisconnectTime)
// -
// -
//
//
func (client *WANIPConnection1) GetIdleDisconnectTime() (
	NewIdleDisconnectTime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIdleDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIdleDisconnectTime, err = soap.UnmarshalUi4(response.NewIdleDisconnectTime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetWarnDisconnectDelay action.
// Arguments:
//
//
// Return values:
//
// * NewWarnDisconnectDelay:
// (related state variable: WarnDisconnectDelay)
// -
// -
//
//
func (client *WANIPConnection1) GetWarnDisconnectDelay() (
	NewWarnDisconnectDelay uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWarnDisconnectDelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWarnDisconnectDelay, err = soap.UnmarshalUi4(response.NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetNATRSIPStatus action.
// Arguments:
//
//
// Return values:
//
// * NewRSIPAvailable:
// (related state variable: RSIPAvailable)
// -
// -
//
//
// * NewNATEnabled:
// (related state variable: NATEnabled)
// -
// -
//
//
func (client *WANIPConnection1) GetNATRSIPStatus() (
	NewRSIPAvailable bool,
	NewNATEnabled bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRSIPAvailable string
		NewNATEnabled    string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetNATRSIPStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRSIPAvailable, err = soap.UnmarshalBoolean(response.NewRSIPAvailable); err != nil {
		return
	}

	if NewNATEnabled, err = soap.UnmarshalBoolean(response.NewNATEnabled); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetGenericPortMappingEntry action.
// Arguments:
//
// * NewPortMappingIndex:
// (related state variable: PortMappingNumberOfEntries)
// -
// -
//
//
//
// Return values:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// -
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// -
// -
//
//
func (client *WANIPConnection1) GetGenericPortMappingEntry(
	NewPortMappingIndex uint16,
) (
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
	err error,
) {
	// Request structure.
	var request struct {
		NewPortMappingIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRemoteHost, err = soap.UnmarshalString(response.NewRemoteHost); err != nil {
		return
	}

	if NewExternalPort, err = soap.UnmarshalUi2(response.NewExternalPort); err != nil {
		return
	}

	if NewProtocol, err = soap.UnmarshalString(response.NewProtocol); err != nil {
		return
	}

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}

	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}

	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}

	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}

	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetSpecificPortMappingEntry action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
//
// Return values:
//
// * NewInternalPort:
// (related state variable: InternalPort)
// -
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// -
// -
//
//
func (client *WANIPConnection1) GetSpecificPortMappingEntry(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}

	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}

	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}

	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}

	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// AddPortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// -
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection1) AddPortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	if request.NewInternalPort, err = soap.MarshalUi2(NewInternalPort); err != nil {
		return
	}

	if request.NewInternalClient, err = soap.MarshalString(NewInternalClient); err != nil {
		return
	}

	if request.NewEnabled, err = soap.MarshalBoolean(NewEnabled); err != nil {
		return
	}

	if request.NewPortMappingDescription, err = soap.MarshalString(NewPortMappingDescription); err != nil {
		return
	}

	if request.NewLeaseDuration, err = soap.MarshalUi4(NewLeaseDuration); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeletePortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
//
// Return values:
//
func (client *WANIPConnection1) DeletePortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetExternalIPAddress action.
// Arguments:
//
//
// Return values:
//
// * NewExternalIPAddress:
// (related state variable: ExternalIPAddress)
// -
// -
//
//
func (client *WANIPConnection1) GetExternalIPAddress() (
	NewExternalIPAddress string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewExternalIPAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetExternalIPAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewExternalIPAddress, err = soap.UnmarshalString(response.NewExternalIPAddress); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANIPConnection2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPConnection:2".
type WANIPConnection2 struct {
	SOAPClient soap.SOAPClient
}

// SetConnectionType action.
// Arguments:
//
// * NewConnectionType:
// (related state variable: ConnectionType)
// -
// - allowed values:
// Unconfigured|IP_Routed|IP_Bridged
//
//
//
// Return values:
//
func (client *WANIPConnection2) SetConnectionType(
	NewConnectionType string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewConnectionType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetConnectionTypeInfo action.
// Arguments:
//
//
// Return values:
//
// * NewConnectionType:
// (related state variable: ConnectionType)
// -
// - allowed values:
// Unconfigured|IP_Routed|IP_Bridged
//
//
// * NewPossibleConnectionTypes:
// (related state variable: PossibleConnectionTypes)
// -
// -
//
//
func (client *WANIPConnection2) GetConnectionTypeInfo() (
	NewConnectionType string,
	NewPossibleConnectionTypes string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionType          string
		NewPossibleConnectionTypes string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetConnectionTypeInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionType, err = soap.UnmarshalString(response.NewConnectionType); err != nil {
		return
	}

	if NewPossibleConnectionTypes, err = soap.UnmarshalString(response.NewPossibleConnectionTypes); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection2) RequestConnection() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection2) RequestTermination() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection2) ForceTermination() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAutoDisconnectTime action.
// Arguments:
//
// * NewAutoDisconnectTime:
// (related state variable: AutoDisconnectTime)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection2) SetAutoDisconnectTime(
	NewAutoDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewAutoDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetIdleDisconnectTime action.
// Arguments:
//
// * NewIdleDisconnectTime:
// (related state variable: IdleDisconnectTime)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection2) SetIdleDisconnectTime(
	NewIdleDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewIdleDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetWarnDisconnectDelay action.
// Arguments:
//
// * NewWarnDisconnectDelay:
// (related state variable: WarnDisconnectDelay)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection2) SetWarnDisconnectDelay(
	NewWarnDisconnectDelay uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewWarnDisconnectDelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetStatusInfo action.
// Arguments:
//
//
// Return values:
//
// * NewConnectionStatus:
// (related state variable: ConnectionStatus)
// -
// - allowed values:
// Unconfigured|Connected|Disconnected
//
//
// * NewLastConnectionError:
// (related state variable: LastConnectionError)
// -
// - allowed values:
// ERROR_NONE
//
//
// * NewUptime:
// (related state variable: Uptime)
// -
// -
//
//
func (client *WANIPConnection2) GetStatusInfo() (
	NewConnectionStatus string,
	NewLastConnectionError string,
	NewUptime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionStatus    string
		NewLastConnectionError string
		NewUptime              string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetStatusInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionStatus, err = soap.UnmarshalString(response.NewConnectionStatus); err != nil {
		return
	}

	if NewLastConnectionError, err = soap.UnmarshalString(response.NewLastConnectionError); err != nil {
		return
	}

	if NewUptime, err = soap.UnmarshalUi4(response.NewUptime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetAutoDisconnectTime action.
// Arguments:
//
//
// Return values:
//
// * NewAutoDisconnectTime:
// (related state variable: AutoDisconnectTime)
// -
// -
//
//
func (client *WANIPConnection2) GetAutoDisconnectTime() (
	NewAutoDisconnectTime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoDisconnectTime, err = soap.UnmarshalUi4(response.NewAutoDisconnectTime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetIdleDisconnectTime action.
// Arguments:
//
//
// Return values:
//
// * NewIdleDisconnectTime:
// (related state variable: IdleDisconnectTime)
// -
// -
//
//
func (client *WANIPConnection2) GetIdleDisconnectTime() (
	NewIdleDisconnectTime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIdleDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIdleDisconnectTime, err = soap.UnmarshalUi4(response.NewIdleDisconnectTime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetWarnDisconnectDelay action.
// Arguments:
//
//
// Return values:
//
// * NewWarnDisconnectDelay:
// (related state variable: WarnDisconnectDelay)
// -
// -
//
//
func (client *WANIPConnection2) GetWarnDisconnectDelay() (
	NewWarnDisconnectDelay uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWarnDisconnectDelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWarnDisconnectDelay, err = soap.UnmarshalUi4(response.NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetNATRSIPStatus action.
// Arguments:
//
//
// Return values:
//
// * NewRSIPAvailable:
// (related state variable: RSIPAvailable)
// -
// -
//
//
// * NewNATEnabled:
// (related state variable: NATEnabled)
// -
// -
//
//
func (client *WANIPConnection2) GetNATRSIPStatus() (
	NewRSIPAvailable bool,
	NewNATEnabled bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRSIPAvailable string
		NewNATEnabled    string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetNATRSIPStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRSIPAvailable, err = soap.UnmarshalBoolean(response.NewRSIPAvailable); err != nil {
		return
	}

	if NewNATEnabled, err = soap.UnmarshalBoolean(response.NewNATEnabled); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetGenericPortMappingEntry action.
// Arguments:
//
// * NewPortMappingIndex:
// (related state variable: PortMappingNumberOfEntries)
// -
// -
//
//
//
// Return values:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// - allowed range: 1 to 65535
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// - allowed range: 0 to 604800
// -
//
//
func (client *WANIPConnection2) GetGenericPortMappingEntry(
	NewPortMappingIndex uint16,
) (
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
	err error,
) {
	// Request structure.
	var request struct {
		NewPortMappingIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRemoteHost, err = soap.UnmarshalString(response.NewRemoteHost); err != nil {
		return
	}

	if NewExternalPort, err = soap.UnmarshalUi2(response.NewExternalPort); err != nil {
		return
	}

	if NewProtocol, err = soap.UnmarshalString(response.NewProtocol); err != nil {
		return
	}

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}

	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}

	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}

	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}

	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetSpecificPortMappingEntry action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
//
// Return values:
//
// * NewInternalPort:
// (related state variable: InternalPort)
// - allowed range: 1 to 65535
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// - allowed range: 0 to 604800
// -
//
//
func (client *WANIPConnection2) GetSpecificPortMappingEntry(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}

	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}

	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}

	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}

	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// AddPortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// - allowed range: 1 to 65535
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// - allowed range: 0 to 604800
// -
//
//
//
// Return values:
//
func (client *WANIPConnection2) AddPortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	if request.NewInternalPort, err = soap.MarshalUi2(NewInternalPort); err != nil {
		return
	}

	if request.NewInternalClient, err = soap.MarshalString(NewInternalClient); err != nil {
		return
	}

	if request.NewEnabled, err = soap.MarshalBoolean(NewEnabled); err != nil {
		return
	}

	if request.NewPortMappingDescription, err = soap.MarshalString(NewPortMappingDescription); err != nil {
		return
	}

	if request.NewLeaseDuration, err = soap.MarshalUi4(NewLeaseDuration); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeletePortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
//
// Return values:
//
func (client *WANIPConnection2) DeletePortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeletePortMappingRange action.
// Arguments:
//
// * NewStartPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewEndPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewManage:
// (related state variable: A_ARG_TYPE_Manage)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPConnection2) DeletePortMappingRange(
	NewStartPort uint16,
	NewEndPort uint16,
	NewProtocol string,
	NewManage bool,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewStartPort string
		NewEndPort   string
		NewProtocol  string
		NewManage    string
	}
	// BEGIN Marshal arguments into request.

	if request.NewStartPort, err = soap.MarshalUi2(NewStartPort); err != nil {
		return
	}

	if request.NewEndPort, err = soap.MarshalUi2(NewEndPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	if request.NewManage, err = soap.MarshalBoolean(NewManage); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "DeletePortMappingRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetExternalIPAddress action.
// Arguments:
//
//
// Return values:
//
// * NewExternalIPAddress:
// (related state variable: ExternalIPAddress)
// -
// -
//
//
func (client *WANIPConnection2) GetExternalIPAddress() (
	NewExternalIPAddress string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewExternalIPAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetExternalIPAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewExternalIPAddress, err = soap.UnmarshalString(response.NewExternalIPAddress); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetListOfPortMappings action.
// Arguments:
//
// * NewStartPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewEndPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewManage:
// (related state variable: A_ARG_TYPE_Manage)
// -
// -
//
//
// * NewNumberOfPorts:
// (related state variable: PortMappingNumberOfEntries)
// -
// -
//
//
//
// Return values:
//
// * NewPortListing:
// (related state variable: A_ARG_TYPE_PortListing)
// -
// -
//
//
func (client *WANIPConnection2) GetListOfPortMappings(
	NewStartPort uint16,
	NewEndPort uint16,
	NewProtocol string,
	NewManage bool,
	NewNumberOfPorts uint16,
) (
	NewPortListing string,
	err error,
) {
	// Request structure.
	var request struct {
		NewStartPort     string
		NewEndPort       string
		NewProtocol      string
		NewManage        string
		NewNumberOfPorts string
	}
	// BEGIN Marshal arguments into request.

	if request.NewStartPort, err = soap.MarshalUi2(NewStartPort); err != nil {
		return
	}

	if request.NewEndPort, err = soap.MarshalUi2(NewEndPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	if request.NewManage, err = soap.MarshalBoolean(NewManage); err != nil {
		return
	}

	if request.NewNumberOfPorts, err = soap.MarshalUi2(NewNumberOfPorts); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPortListing string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetListOfPortMappings", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPortListing, err = soap.UnmarshalString(response.NewPortListing); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// AddAnyPortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// - allowed range: 1 to 65535
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// - allowed range: 0 to 604800
// -
//
//
//
// Return values:
//
// * NewReservedPort:
// (related state variable: ExternalPort)
// -
// -
//
//
func (client *WANIPConnection2) AddAnyPortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
) (
	NewReservedPort uint16,
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	if request.NewInternalPort, err = soap.MarshalUi2(NewInternalPort); err != nil {
		return
	}

	if request.NewInternalClient, err = soap.MarshalString(NewInternalClient); err != nil {
		return
	}

	if request.NewEnabled, err = soap.MarshalBoolean(NewEnabled); err != nil {
		return
	}

	if request.NewPortMappingDescription, err = soap.MarshalString(NewPortMappingDescription); err != nil {
		return
	}

	if request.NewLeaseDuration, err = soap.MarshalUi4(NewLeaseDuration); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewReservedPort string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "AddAnyPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewReservedPort, err = soap.UnmarshalUi2(response.NewReservedPort); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANIPv6FirewallControl1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1".
type WANIPv6FirewallControl1 struct {
	SOAPClient soap.SOAPClient
}

// GetFirewallStatus action.
// Arguments:
//
//
// Return values:
//
// * FirewallEnabled:
// (related state variable: FirewallEnabled)
// -
// -
//
//
// * InboundPinholeAllowed:
// (related state variable: InboundPinholeAllowed)
// -
// -
//
//
func (client *WANIPv6FirewallControl1) GetFirewallStatus() (
	FirewallEnabled bool,
	InboundPinholeAllowed bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		FirewallEnabled       string
		InboundPinholeAllowed string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "GetFirewallStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if FirewallEnabled, err = soap.UnmarshalBoolean(response.FirewallEnabled); err != nil {
		return
	}

	if InboundPinholeAllowed, err = soap.UnmarshalBoolean(response.InboundPinholeAllowed); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetOutboundPinholeTimeout action.
// Arguments:
//
// * RemoteHost:
// (related state variable: A_ARG_TYPE_IPv6Address)
// -
// -
//
//
// * RemotePort:
// (related state variable: A_ARG_TYPE_Port)
// -
// -
//
//
// * InternalClient:
// (related state variable: A_ARG_TYPE_IPv6Address)
// -
// -
//
//
// * InternalPort:
// (related state variable: A_ARG_TYPE_Port)
// -
// -
//
//
// * Protocol:
// (related state variable: A_ARG_TYPE_Protocol)
// -
// -
//
//
//
// Return values:
//
// * OutboundPinholeTimeout:
// (related state variable: A_ARG_TYPE_OutboundPinholeTimeout)
// -
// -
//
//
func (client *WANIPv6FirewallControl1) GetOutboundPinholeTimeout(
	RemoteHost string,
	RemotePort uint16,
	InternalClient string,
	InternalPort uint16,
	Protocol uint16,
) (
	OutboundPinholeTimeout uint32,
	err error,
) {
	// Request structure.
	var request struct {
		RemoteHost     string
		RemotePort     string
		InternalClient string
		InternalPort   string
		Protocol       string
	}
	// BEGIN Marshal arguments into request.

	if request.RemoteHost, err = soap.MarshalString(RemoteHost); err != nil {
		return
	}

	if request.RemotePort, err = soap.MarshalUi2(RemotePort); err != nil {
		return
	}

	if request.InternalClient, err = soap.MarshalString(InternalClient); err != nil {
		return
	}

	if request.InternalPort, err = soap.MarshalUi2(InternalPort); err != nil {
		return
	}

	if request.Protocol, err = soap.MarshalUi2(Protocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		OutboundPinholeTimeout string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "GetOutboundPinholeTimeout", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if OutboundPinholeTimeout, err = soap.UnmarshalUi4(response.OutboundPinholeTimeout); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// AddPinhole action.
// Arguments:
//
// * RemoteHost:
// (related state variable: A_ARG_TYPE_IPv6Address)
// -
// -
//
//
// * RemotePort:
// (related state variable: A_ARG_TYPE_Port)
// -
// -
//
//
// * InternalClient:
// (related state variable: A_ARG_TYPE_IPv6Address)
// -
// -
//
//
// * InternalPort:
// (related state variable: A_ARG_TYPE_Port)
// -
// -
//
//
// * Protocol:
// (related state variable: A_ARG_TYPE_Protocol)
// -
// -
//
//
// * LeaseTime:
// (related state variable: A_ARG_TYPE_LeaseTime)
// - allowed range: 1 to 86400
// -
//
//
//
// Return values:
//
// * UniqueID:
// (related state variable: A_ARG_TYPE_UniqueID)
// -
// -
//
//
func (client *WANIPv6FirewallControl1) AddPinhole(
	RemoteHost string,
	RemotePort uint16,
	InternalClient string,
	InternalPort uint16,
	Protocol uint16,
	LeaseTime uint32,
) (
	UniqueID uint16,
	err error,
) {
	// Request structure.
	var request struct {
		RemoteHost     string
		RemotePort     string
		InternalClient string
		InternalPort   string
		Protocol       string
		LeaseTime      string
	}
	// BEGIN Marshal arguments into request.

	if request.RemoteHost, err = soap.MarshalString(RemoteHost); err != nil {
		return
	}

	if request.RemotePort, err = soap.MarshalUi2(RemotePort); err != nil {
		return
	}

	if request.InternalClient, err = soap.MarshalString(InternalClient); err != nil {
		return
	}

	if request.InternalPort, err = soap.MarshalUi2(InternalPort); err != nil {
		return
	}

	if request.Protocol, err = soap.MarshalUi2(Protocol); err != nil {
		return
	}

	if request.LeaseTime, err = soap.MarshalUi4(LeaseTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		UniqueID string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "AddPinhole", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if UniqueID, err = soap.UnmarshalUi2(response.UniqueID); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// UpdatePinhole action.
// Arguments:
//
// * UniqueID:
// (related state variable: A_ARG_TYPE_UniqueID)
// -
// -
//
//
// * NewLeaseTime:
// (related state variable: A_ARG_TYPE_LeaseTime)
// - allowed range: 1 to 86400
// -
//
//
//
// Return values:
//
func (client *WANIPv6FirewallControl1) UpdatePinhole(
	UniqueID uint16,
	NewLeaseTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		UniqueID     string
		NewLeaseTime string
	}
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	if request.NewLeaseTime, err = soap.MarshalUi4(NewLeaseTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "UpdatePinhole", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeletePinhole action.
// Arguments:
//
// * UniqueID:
// (related state variable: A_ARG_TYPE_UniqueID)
// -
// -
//
//
//
// Return values:
//
func (client *WANIPv6FirewallControl1) DeletePinhole(
	UniqueID uint16,
) (
	err error,
) {
	// Request structure.
	var request struct {
		UniqueID string
	}
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "DeletePinhole", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetPinholePackets action.
// Arguments:
//
// * UniqueID:
// (related state variable: A_ARG_TYPE_UniqueID)
// -
// -
//
//
//
// Return values:
//
// * PinholePackets:
// (related state variable: A_ARG_TYPE_PinholePackets)
// -
// -
//
//
func (client *WANIPv6FirewallControl1) GetPinholePackets(
	UniqueID uint16,
) (
	PinholePackets uint32,
	err error,
) {
	// Request structure.
	var request struct {
		UniqueID string
	}
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		PinholePackets string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "GetPinholePackets", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PinholePackets, err = soap.UnmarshalUi4(response.PinholePackets); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// CheckPinholeWorking action.
// Arguments:
//
// * UniqueID:
// (related state variable: A_ARG_TYPE_UniqueID)
// -
// -
//
//
//
// Return values:
//
// * IsWorking:
// (related state variable: A_ARG_TYPE_Boolean)
// -
// -
//
//
func (client *WANIPv6FirewallControl1) CheckPinholeWorking(
	UniqueID uint16,
) (
	IsWorking bool,
	err error,
) {
	// Request structure.
	var request struct {
		UniqueID string
	}
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		IsWorking string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "CheckPinholeWorking", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if IsWorking, err = soap.UnmarshalBoolean(response.IsWorking); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANPOTSLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1".
type WANPOTSLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

// SetISPInfo action.
// Arguments:
//
// * NewISPPhoneNumber:
// (related state variable: ISPPhoneNumber)
// -
// -
//
//
// * NewISPInfo:
// (related state variable: ISPInfo)
// -
// -
//
//
// * NewLinkType:
// (related state variable: LinkType)
// -
// - allowed values:
// PPP_Dialup
//
//
//
// Return values:
//
func (client *WANPOTSLinkConfig1) SetISPInfo(
	NewISPPhoneNumber string,
	NewISPInfo string,
	NewLinkType string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewISPPhoneNumber string
		NewISPInfo        string
		NewLinkType       string
	}
	// BEGIN Marshal arguments into request.

	if request.NewISPPhoneNumber, err = soap.MarshalString(NewISPPhoneNumber); err != nil {
		return
	}

	if request.NewISPInfo, err = soap.MarshalString(NewISPInfo); err != nil {
		return
	}

	if request.NewLinkType, err = soap.MarshalString(NewLinkType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetISPInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetCallRetryInfo action.
// Arguments:
//
// * NewNumberOfRetries:
// (related state variable: NumberOfRetries)
// -
// -
//
//
// * NewDelayBetweenRetries:
// (related state variable: DelayBetweenRetries)
// -
// -
//
//
//
// Return values:
//
func (client *WANPOTSLinkConfig1) SetCallRetryInfo(
	NewNumberOfRetries uint32,
	NewDelayBetweenRetries uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewNumberOfRetries     string
		NewDelayBetweenRetries string
	}
	// BEGIN Marshal arguments into request.

	if request.NewNumberOfRetries, err = soap.MarshalUi4(NewNumberOfRetries); err != nil {
		return
	}

	if request.NewDelayBetweenRetries, err = soap.MarshalUi4(NewDelayBetweenRetries); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetCallRetryInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetISPInfo action.
// Arguments:
//
//
// Return values:
//
// * NewISPPhoneNumber:
// (related state variable: ISPPhoneNumber)
// -
// -
//
//
// * NewISPInfo:
// (related state variable: ISPInfo)
// -
// -
//
//
// * NewLinkType:
// (related state variable: LinkType)
// -
// - allowed values:
// PPP_Dialup
//
//
func (client *WANPOTSLinkConfig1) GetISPInfo() (
	NewISPPhoneNumber string,
	NewISPInfo string,
	NewLinkType string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewISPPhoneNumber string
		NewISPInfo        string
		NewLinkType       string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetISPInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewISPPhoneNumber, err = soap.UnmarshalString(response.NewISPPhoneNumber); err != nil {
		return
	}

	if NewISPInfo, err = soap.UnmarshalString(response.NewISPInfo); err != nil {
		return
	}

	if NewLinkType, err = soap.UnmarshalString(response.NewLinkType); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetCallRetryInfo action.
// Arguments:
//
//
// Return values:
//
// * NewNumberOfRetries:
// (related state variable: NumberOfRetries)
// -
// -
//
//
// * NewDelayBetweenRetries:
// (related state variable: DelayBetweenRetries)
// -
// -
//
//
func (client *WANPOTSLinkConfig1) GetCallRetryInfo() (
	NewNumberOfRetries uint32,
	NewDelayBetweenRetries uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewNumberOfRetries     string
		NewDelayBetweenRetries string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetCallRetryInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewNumberOfRetries, err = soap.UnmarshalUi4(response.NewNumberOfRetries); err != nil {
		return
	}

	if NewDelayBetweenRetries, err = soap.UnmarshalUi4(response.NewDelayBetweenRetries); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetFclass action.
// Arguments:
//
//
// Return values:
//
// * NewFclass:
// (related state variable: Fclass)
// -
// -
//
//
func (client *WANPOTSLinkConfig1) GetFclass() (
	NewFclass string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewFclass string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetFclass", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewFclass, err = soap.UnmarshalString(response.NewFclass); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetDataModulationSupported action.
// Arguments:
//
//
// Return values:
//
// * NewDataModulationSupported:
// (related state variable: DataModulationSupported)
// -
// -
//
//
func (client *WANPOTSLinkConfig1) GetDataModulationSupported() (
	NewDataModulationSupported string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDataModulationSupported string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataModulationSupported", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDataModulationSupported, err = soap.UnmarshalString(response.NewDataModulationSupported); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetDataProtocol action.
// Arguments:
//
//
// Return values:
//
// * NewDataProtocol:
// (related state variable: DataProtocol)
// -
// -
//
//
func (client *WANPOTSLinkConfig1) GetDataProtocol() (
	NewDataProtocol string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDataProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDataProtocol, err = soap.UnmarshalString(response.NewDataProtocol); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetDataCompression action.
// Arguments:
//
//
// Return values:
//
// * NewDataCompression:
// (related state variable: DataCompression)
// -
// -
//
//
func (client *WANPOTSLinkConfig1) GetDataCompression() (
	NewDataCompression string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDataCompression string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataCompression", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDataCompression, err = soap.UnmarshalString(response.NewDataCompression); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetPlusVTRCommandSupported action.
// Arguments:
//
//
// Return values:
//
// * NewPlusVTRCommandSupported:
// (related state variable: PlusVTRCommandSupported)
// -
// -
//
//
func (client *WANPOTSLinkConfig1) GetPlusVTRCommandSupported() (
	NewPlusVTRCommandSupported bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPlusVTRCommandSupported string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetPlusVTRCommandSupported", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPlusVTRCommandSupported, err = soap.UnmarshalBoolean(response.NewPlusVTRCommandSupported); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// WANPPPConnection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANPPPConnection:1".
type WANPPPConnection1 struct {
	SOAPClient soap.SOAPClient
}

// SetConnectionType action.
// Arguments:
//
// * NewConnectionType:
// (related state variable: ConnectionType)
// -
// -
//
//
//
// Return values:
//
func (client *WANPPPConnection1) SetConnectionType(
	NewConnectionType string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewConnectionType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetConnectionTypeInfo action.
// Arguments:
//
//
// Return values:
//
// * NewConnectionType:
// (related state variable: ConnectionType)
// -
// -
//
//
// * NewPossibleConnectionTypes:
// (related state variable: PossibleConnectionTypes)
// -
// - allowed values:
// Unconfigured|IP_Routed|DHCP_Spoofed|PPPoE_Bridged|PPTP_Relay|L2TP_Relay|PPPoE_Relay
//
//
func (client *WANPPPConnection1) GetConnectionTypeInfo() (
	NewConnectionType string,
	NewPossibleConnectionTypes string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionType          string
		NewPossibleConnectionTypes string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetConnectionTypeInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionType, err = soap.UnmarshalString(response.NewConnectionType); err != nil {
		return
	}

	if NewPossibleConnectionTypes, err = soap.UnmarshalString(response.NewPossibleConnectionTypes); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// ConfigureConnection action.
// Arguments:
//
// * NewUserName:
// (related state variable: UserName)
// -
// -
//
//
// * NewPassword:
// (related state variable: Password)
// -
// -
//
//
//
// Return values:
//
func (client *WANPPPConnection1) ConfigureConnection(
	NewUserName string,
	NewPassword string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewUserName string
		NewPassword string
	}
	// BEGIN Marshal arguments into request.

	if request.NewUserName, err = soap.MarshalString(NewUserName); err != nil {
		return
	}

	if request.NewPassword, err = soap.MarshalString(NewPassword); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ConfigureConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) RequestConnection() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) RequestTermination() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) ForceTermination() (
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAutoDisconnectTime action.
// Arguments:
//
// * NewAutoDisconnectTime:
// (related state variable: AutoDisconnectTime)
// -
// -
//
//
//
// Return values:
//
func (client *WANPPPConnection1) SetAutoDisconnectTime(
	NewAutoDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewAutoDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetIdleDisconnectTime action.
// Arguments:
//
// * NewIdleDisconnectTime:
// (related state variable: IdleDisconnectTime)
// -
// -
//
//
//
// Return values:
//
func (client *WANPPPConnection1) SetIdleDisconnectTime(
	NewIdleDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewIdleDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetWarnDisconnectDelay action.
// Arguments:
//
// * NewWarnDisconnectDelay:
// (related state variable: WarnDisconnectDelay)
// -
// -
//
//
//
// Return values:
//
func (client *WANPPPConnection1) SetWarnDisconnectDelay(
	NewWarnDisconnectDelay uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewWarnDisconnectDelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetStatusInfo action.
// Arguments:
//
//
// Return values:
//
// * NewConnectionStatus:
// (related state variable: ConnectionStatus)
// -
// - allowed values:
// Unconfigured|Connected|Disconnected
//
//
// * NewLastConnectionError:
// (related state variable: LastConnectionError)
// -
// - allowed values:
// ERROR_NONE
//
//
// * NewUptime:
// (related state variable: Uptime)
// -
// -
//
//
func (client *WANPPPConnection1) GetStatusInfo() (
	NewConnectionStatus string,
	NewLastConnectionError string,
	NewUptime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionStatus    string
		NewLastConnectionError string
		NewUptime              string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetStatusInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionStatus, err = soap.UnmarshalString(response.NewConnectionStatus); err != nil {
		return
	}

	if NewLastConnectionError, err = soap.UnmarshalString(response.NewLastConnectionError); err != nil {
		return
	}

	if NewUptime, err = soap.UnmarshalUi4(response.NewUptime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetLinkLayerMaxBitRates action.
// Arguments:
//
//
// Return values:
//
// * NewUpstreamMaxBitRate:
// (related state variable: UpstreamMaxBitRate)
// -
// -
//
//
// * NewDownstreamMaxBitRate:
// (related state variable: DownstreamMaxBitRate)
// -
// -
//
//
func (client *WANPPPConnection1) GetLinkLayerMaxBitRates() (
	NewUpstreamMaxBitRate uint32,
	NewDownstreamMaxBitRate uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamMaxBitRate   string
		NewDownstreamMaxBitRate string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetLinkLayerMaxBitRates", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewUpstreamMaxBitRate); err != nil {
		return
	}

	if NewDownstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewDownstreamMaxBitRate); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetPPPEncryptionProtocol action.
// Arguments:
//
//
// Return values:
//
// * NewPPPEncryptionProtocol:
// (related state variable: PPPEncryptionProtocol)
// -
// -
//
//
func (client *WANPPPConnection1) GetPPPEncryptionProtocol() (
	NewPPPEncryptionProtocol string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPPPEncryptionProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPEncryptionProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPPPEncryptionProtocol, err = soap.UnmarshalString(response.NewPPPEncryptionProtocol); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetPPPCompressionProtocol action.
// Arguments:
//
//
// Return values:
//
// * NewPPPCompressionProtocol:
// (related state variable: PPPCompressionProtocol)
// -
// -
//
//
func (client *WANPPPConnection1) GetPPPCompressionProtocol() (
	NewPPPCompressionProtocol string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPPPCompressionProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPCompressionProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPPPCompressionProtocol, err = soap.UnmarshalString(response.NewPPPCompressionProtocol); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetPPPAuthenticationProtocol action.
// Arguments:
//
//
// Return values:
//
// * NewPPPAuthenticationProtocol:
// (related state variable: PPPAuthenticationProtocol)
// -
// -
//
//
func (client *WANPPPConnection1) GetPPPAuthenticationProtocol() (
	NewPPPAuthenticationProtocol string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPPPAuthenticationProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPAuthenticationProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPPPAuthenticationProtocol, err = soap.UnmarshalString(response.NewPPPAuthenticationProtocol); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetUserName action.
// Arguments:
//
//
// Return values:
//
// * NewUserName:
// (related state variable: UserName)
// -
// -
//
//
func (client *WANPPPConnection1) GetUserName() (
	NewUserName string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUserName string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetUserName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUserName, err = soap.UnmarshalString(response.NewUserName); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetPassword action.
// Arguments:
//
//
// Return values:
//
// * NewPassword:
// (related state variable: Password)
// -
// -
//
//
func (client *WANPPPConnection1) GetPassword() (
	NewPassword string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPassword string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPassword", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPassword, err = soap.UnmarshalString(response.NewPassword); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetAutoDisconnectTime action.
// Arguments:
//
//
// Return values:
//
// * NewAutoDisconnectTime:
// (related state variable: AutoDisconnectTime)
// -
// -
//
//
func (client *WANPPPConnection1) GetAutoDisconnectTime() (
	NewAutoDisconnectTime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoDisconnectTime, err = soap.UnmarshalUi4(response.NewAutoDisconnectTime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetIdleDisconnectTime action.
// Arguments:
//
//
// Return values:
//
// * NewIdleDisconnectTime:
// (related state variable: IdleDisconnectTime)
// -
// -
//
//
func (client *WANPPPConnection1) GetIdleDisconnectTime() (
	NewIdleDisconnectTime uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIdleDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIdleDisconnectTime, err = soap.UnmarshalUi4(response.NewIdleDisconnectTime); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetWarnDisconnectDelay action.
// Arguments:
//
//
// Return values:
//
// * NewWarnDisconnectDelay:
// (related state variable: WarnDisconnectDelay)
// -
// -
//
//
func (client *WANPPPConnection1) GetWarnDisconnectDelay() (
	NewWarnDisconnectDelay uint32,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWarnDisconnectDelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWarnDisconnectDelay, err = soap.UnmarshalUi4(response.NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetNATRSIPStatus action.
// Arguments:
//
//
// Return values:
//
// * NewRSIPAvailable:
// (related state variable: RSIPAvailable)
// -
// -
//
//
// * NewNATEnabled:
// (related state variable: NATEnabled)
// -
// -
//
//
func (client *WANPPPConnection1) GetNATRSIPStatus() (
	NewRSIPAvailable bool,
	NewNATEnabled bool,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRSIPAvailable string
		NewNATEnabled    string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetNATRSIPStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRSIPAvailable, err = soap.UnmarshalBoolean(response.NewRSIPAvailable); err != nil {
		return
	}

	if NewNATEnabled, err = soap.UnmarshalBoolean(response.NewNATEnabled); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetGenericPortMappingEntry action.
// Arguments:
//
// * NewPortMappingIndex:
// (related state variable: PortMappingNumberOfEntries)
// -
// -
//
//
//
// Return values:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// -
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// -
// -
//
//
func (client *WANPPPConnection1) GetGenericPortMappingEntry(
	NewPortMappingIndex uint16,
) (
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
	err error,
) {
	// Request structure.
	var request struct {
		NewPortMappingIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRemoteHost, err = soap.UnmarshalString(response.NewRemoteHost); err != nil {
		return
	}

	if NewExternalPort, err = soap.UnmarshalUi2(response.NewExternalPort); err != nil {
		return
	}

	if NewProtocol, err = soap.UnmarshalString(response.NewProtocol); err != nil {
		return
	}

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}

	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}

	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}

	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}

	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// GetSpecificPortMappingEntry action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
//
// Return values:
//
// * NewInternalPort:
// (related state variable: InternalPort)
// -
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// -
// -
//
//
func (client *WANPPPConnection1) GetSpecificPortMappingEntry(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}

	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}

	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}

	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}

	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}

// AddPortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
// * NewInternalPort:
// (related state variable: InternalPort)
// -
// -
//
//
// * NewInternalClient:
// (related state variable: InternalClient)
// -
// -
//
//
// * NewEnabled:
// (related state variable: PortMappingEnabled)
// -
// -
//
//
// * NewPortMappingDescription:
// (related state variable: PortMappingDescription)
// -
// -
//
//
// * NewLeaseDuration:
// (related state variable: PortMappingLeaseDuration)
// -
// -
//
//
//
// Return values:
//
func (client *WANPPPConnection1) AddPortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
	NewInternalPort uint16,
	NewInternalClient string,
	NewEnabled bool,
	NewPortMappingDescription string,
	NewLeaseDuration uint32,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	if request.NewInternalPort, err = soap.MarshalUi2(NewInternalPort); err != nil {
		return
	}

	if request.NewInternalClient, err = soap.MarshalString(NewInternalClient); err != nil {
		return
	}

	if request.NewEnabled, err = soap.MarshalBoolean(NewEnabled); err != nil {
		return
	}

	if request.NewPortMappingDescription, err = soap.MarshalString(NewPortMappingDescription); err != nil {
		return
	}

	if request.NewLeaseDuration, err = soap.MarshalUi4(NewLeaseDuration); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeletePortMapping action.
// Arguments:
//
// * NewRemoteHost:
// (related state variable: RemoteHost)
// -
// -
//
//
// * NewExternalPort:
// (related state variable: ExternalPort)
// -
// -
//
//
// * NewProtocol:
// (related state variable: PortMappingProtocol)
// -
// - allowed values:
// TCP|UDP
//
//
//
// Return values:
//
func (client *WANPPPConnection1) DeletePortMapping(
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	err error,
) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}

	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}

	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// GetExternalIPAddress action.
// Arguments:
//
//
// Return values:
//
// * NewExternalIPAddress:
// (related state variable: ExternalIPAddress)
// -
// -
//
//
func (client *WANPPPConnection1) GetExternalIPAddress() (
	NewExternalIPAddress string,
	err error,
) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewExternalIPAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetExternalIPAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewExternalIPAddress, err = soap.UnmarshalString(response.NewExternalIPAddress); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}
