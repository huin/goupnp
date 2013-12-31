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

// _LANHostConfigManagement1_SetDHCPServerConfigurable_Request is the XML structure for the input arguments for action SetDHCPServerConfigurable.
type _LANHostConfigManagement1_SetDHCPServerConfigurable_Request struct {
	NewDHCPServerConfigurable string
}

// _LANHostConfigManagement1_SetDHCPServerConfigurable_Response is the XML structure for the output arguments for action SetDHCPServerConfigurable.
type _LANHostConfigManagement1_SetDHCPServerConfigurable_Response struct{}

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
	var request _LANHostConfigManagement1_SetDHCPServerConfigurable_Request
	// BEGIN Marshal arguments into request.

	if request.NewDHCPServerConfigurable, err = soap.MarshalBoolean(NewDHCPServerConfigurable); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetDHCPServerConfigurable_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPServerConfigurable", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetDHCPServerConfigurable_Request is the XML structure for the input arguments for action GetDHCPServerConfigurable.
type _LANHostConfigManagement1_GetDHCPServerConfigurable_Request struct{}

// _LANHostConfigManagement1_GetDHCPServerConfigurable_Response is the XML structure for the output arguments for action GetDHCPServerConfigurable.
type _LANHostConfigManagement1_GetDHCPServerConfigurable_Response struct {
	NewDHCPServerConfigurable string
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
	var request _LANHostConfigManagement1_GetDHCPServerConfigurable_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetDHCPServerConfigurable_Response
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

// _LANHostConfigManagement1_SetDHCPRelay_Request is the XML structure for the input arguments for action SetDHCPRelay.
type _LANHostConfigManagement1_SetDHCPRelay_Request struct {
	NewDHCPRelay string
}

// _LANHostConfigManagement1_SetDHCPRelay_Response is the XML structure for the output arguments for action SetDHCPRelay.
type _LANHostConfigManagement1_SetDHCPRelay_Response struct{}

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
	var request _LANHostConfigManagement1_SetDHCPRelay_Request
	// BEGIN Marshal arguments into request.

	if request.NewDHCPRelay, err = soap.MarshalBoolean(NewDHCPRelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetDHCPRelay_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPRelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetDHCPRelay_Request is the XML structure for the input arguments for action GetDHCPRelay.
type _LANHostConfigManagement1_GetDHCPRelay_Request struct{}

// _LANHostConfigManagement1_GetDHCPRelay_Response is the XML structure for the output arguments for action GetDHCPRelay.
type _LANHostConfigManagement1_GetDHCPRelay_Response struct {
	NewDHCPRelay string
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
	var request _LANHostConfigManagement1_GetDHCPRelay_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetDHCPRelay_Response
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

// _LANHostConfigManagement1_SetSubnetMask_Request is the XML structure for the input arguments for action SetSubnetMask.
type _LANHostConfigManagement1_SetSubnetMask_Request struct {
	NewSubnetMask string
}

// _LANHostConfigManagement1_SetSubnetMask_Response is the XML structure for the output arguments for action SetSubnetMask.
type _LANHostConfigManagement1_SetSubnetMask_Response struct{}

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
	var request _LANHostConfigManagement1_SetSubnetMask_Request
	// BEGIN Marshal arguments into request.

	if request.NewSubnetMask, err = soap.MarshalString(NewSubnetMask); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetSubnetMask_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetSubnetMask", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetSubnetMask_Request is the XML structure for the input arguments for action GetSubnetMask.
type _LANHostConfigManagement1_GetSubnetMask_Request struct{}

// _LANHostConfigManagement1_GetSubnetMask_Response is the XML structure for the output arguments for action GetSubnetMask.
type _LANHostConfigManagement1_GetSubnetMask_Response struct {
	NewSubnetMask string
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
	var request _LANHostConfigManagement1_GetSubnetMask_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetSubnetMask_Response
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

// _LANHostConfigManagement1_SetIPRouter_Request is the XML structure for the input arguments for action SetIPRouter.
type _LANHostConfigManagement1_SetIPRouter_Request struct {
	NewIPRouters string
}

// _LANHostConfigManagement1_SetIPRouter_Response is the XML structure for the output arguments for action SetIPRouter.
type _LANHostConfigManagement1_SetIPRouter_Response struct{}

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
	var request _LANHostConfigManagement1_SetIPRouter_Request
	// BEGIN Marshal arguments into request.

	if request.NewIPRouters, err = soap.MarshalString(NewIPRouters); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetIPRouter_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetIPRouter", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_DeleteIPRouter_Request is the XML structure for the input arguments for action DeleteIPRouter.
type _LANHostConfigManagement1_DeleteIPRouter_Request struct {
	NewIPRouters string
}

// _LANHostConfigManagement1_DeleteIPRouter_Response is the XML structure for the output arguments for action DeleteIPRouter.
type _LANHostConfigManagement1_DeleteIPRouter_Response struct{}

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
	var request _LANHostConfigManagement1_DeleteIPRouter_Request
	// BEGIN Marshal arguments into request.

	if request.NewIPRouters, err = soap.MarshalString(NewIPRouters); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_DeleteIPRouter_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteIPRouter", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetIPRoutersList_Request is the XML structure for the input arguments for action GetIPRoutersList.
type _LANHostConfigManagement1_GetIPRoutersList_Request struct{}

// _LANHostConfigManagement1_GetIPRoutersList_Response is the XML structure for the output arguments for action GetIPRoutersList.
type _LANHostConfigManagement1_GetIPRoutersList_Response struct {
	NewIPRouters string
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
	var request _LANHostConfigManagement1_GetIPRoutersList_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetIPRoutersList_Response
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

// _LANHostConfigManagement1_SetDomainName_Request is the XML structure for the input arguments for action SetDomainName.
type _LANHostConfigManagement1_SetDomainName_Request struct {
	NewDomainName string
}

// _LANHostConfigManagement1_SetDomainName_Response is the XML structure for the output arguments for action SetDomainName.
type _LANHostConfigManagement1_SetDomainName_Response struct{}

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
	var request _LANHostConfigManagement1_SetDomainName_Request
	// BEGIN Marshal arguments into request.

	if request.NewDomainName, err = soap.MarshalString(NewDomainName); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetDomainName_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDomainName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetDomainName_Request is the XML structure for the input arguments for action GetDomainName.
type _LANHostConfigManagement1_GetDomainName_Request struct{}

// _LANHostConfigManagement1_GetDomainName_Response is the XML structure for the output arguments for action GetDomainName.
type _LANHostConfigManagement1_GetDomainName_Response struct {
	NewDomainName string
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
	var request _LANHostConfigManagement1_GetDomainName_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetDomainName_Response
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

// _LANHostConfigManagement1_SetAddressRange_Request is the XML structure for the input arguments for action SetAddressRange.
type _LANHostConfigManagement1_SetAddressRange_Request struct {
	NewMinAddress string

	NewMaxAddress string
}

// _LANHostConfigManagement1_SetAddressRange_Response is the XML structure for the output arguments for action SetAddressRange.
type _LANHostConfigManagement1_SetAddressRange_Response struct{}

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
	var request _LANHostConfigManagement1_SetAddressRange_Request
	// BEGIN Marshal arguments into request.

	if request.NewMinAddress, err = soap.MarshalString(NewMinAddress); err != nil {
		return
	}

	if request.NewMaxAddress, err = soap.MarshalString(NewMaxAddress); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetAddressRange_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetAddressRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetAddressRange_Request is the XML structure for the input arguments for action GetAddressRange.
type _LANHostConfigManagement1_GetAddressRange_Request struct{}

// _LANHostConfigManagement1_GetAddressRange_Response is the XML structure for the output arguments for action GetAddressRange.
type _LANHostConfigManagement1_GetAddressRange_Response struct {
	NewMinAddress string

	NewMaxAddress string
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
	var request _LANHostConfigManagement1_GetAddressRange_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetAddressRange_Response
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

// _LANHostConfigManagement1_SetReservedAddress_Request is the XML structure for the input arguments for action SetReservedAddress.
type _LANHostConfigManagement1_SetReservedAddress_Request struct {
	NewReservedAddresses string
}

// _LANHostConfigManagement1_SetReservedAddress_Response is the XML structure for the output arguments for action SetReservedAddress.
type _LANHostConfigManagement1_SetReservedAddress_Response struct{}

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
	var request _LANHostConfigManagement1_SetReservedAddress_Request
	// BEGIN Marshal arguments into request.

	if request.NewReservedAddresses, err = soap.MarshalString(NewReservedAddresses); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetReservedAddress_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetReservedAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_DeleteReservedAddress_Request is the XML structure for the input arguments for action DeleteReservedAddress.
type _LANHostConfigManagement1_DeleteReservedAddress_Request struct {
	NewReservedAddresses string
}

// _LANHostConfigManagement1_DeleteReservedAddress_Response is the XML structure for the output arguments for action DeleteReservedAddress.
type _LANHostConfigManagement1_DeleteReservedAddress_Response struct{}

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
	var request _LANHostConfigManagement1_DeleteReservedAddress_Request
	// BEGIN Marshal arguments into request.

	if request.NewReservedAddresses, err = soap.MarshalString(NewReservedAddresses); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_DeleteReservedAddress_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteReservedAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetReservedAddresses_Request is the XML structure for the input arguments for action GetReservedAddresses.
type _LANHostConfigManagement1_GetReservedAddresses_Request struct{}

// _LANHostConfigManagement1_GetReservedAddresses_Response is the XML structure for the output arguments for action GetReservedAddresses.
type _LANHostConfigManagement1_GetReservedAddresses_Response struct {
	NewReservedAddresses string
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
	var request _LANHostConfigManagement1_GetReservedAddresses_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetReservedAddresses_Response
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

// _LANHostConfigManagement1_SetDNSServer_Request is the XML structure for the input arguments for action SetDNSServer.
type _LANHostConfigManagement1_SetDNSServer_Request struct {
	NewDNSServers string
}

// _LANHostConfigManagement1_SetDNSServer_Response is the XML structure for the output arguments for action SetDNSServer.
type _LANHostConfigManagement1_SetDNSServer_Response struct{}

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
	var request _LANHostConfigManagement1_SetDNSServer_Request
	// BEGIN Marshal arguments into request.

	if request.NewDNSServers, err = soap.MarshalString(NewDNSServers); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_SetDNSServer_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDNSServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_DeleteDNSServer_Request is the XML structure for the input arguments for action DeleteDNSServer.
type _LANHostConfigManagement1_DeleteDNSServer_Request struct {
	NewDNSServers string
}

// _LANHostConfigManagement1_DeleteDNSServer_Response is the XML structure for the output arguments for action DeleteDNSServer.
type _LANHostConfigManagement1_DeleteDNSServer_Response struct{}

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
	var request _LANHostConfigManagement1_DeleteDNSServer_Request
	// BEGIN Marshal arguments into request.

	if request.NewDNSServers, err = soap.MarshalString(NewDNSServers); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_DeleteDNSServer_Response
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteDNSServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _LANHostConfigManagement1_GetDNSServers_Request is the XML structure for the input arguments for action GetDNSServers.
type _LANHostConfigManagement1_GetDNSServers_Request struct{}

// _LANHostConfigManagement1_GetDNSServers_Response is the XML structure for the output arguments for action GetDNSServers.
type _LANHostConfigManagement1_GetDNSServers_Response struct {
	NewDNSServers string
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
	var request _LANHostConfigManagement1_GetDNSServers_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _LANHostConfigManagement1_GetDNSServers_Response
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

// _Layer3Forwarding1_SetDefaultConnectionService_Request is the XML structure for the input arguments for action SetDefaultConnectionService.
type _Layer3Forwarding1_SetDefaultConnectionService_Request struct {
	NewDefaultConnectionService string
}

// _Layer3Forwarding1_SetDefaultConnectionService_Response is the XML structure for the output arguments for action SetDefaultConnectionService.
type _Layer3Forwarding1_SetDefaultConnectionService_Response struct{}

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
	var request _Layer3Forwarding1_SetDefaultConnectionService_Request
	// BEGIN Marshal arguments into request.

	if request.NewDefaultConnectionService, err = soap.MarshalString(NewDefaultConnectionService); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _Layer3Forwarding1_SetDefaultConnectionService_Response
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "SetDefaultConnectionService", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _Layer3Forwarding1_GetDefaultConnectionService_Request is the XML structure for the input arguments for action GetDefaultConnectionService.
type _Layer3Forwarding1_GetDefaultConnectionService_Request struct{}

// _Layer3Forwarding1_GetDefaultConnectionService_Response is the XML structure for the output arguments for action GetDefaultConnectionService.
type _Layer3Forwarding1_GetDefaultConnectionService_Response struct {
	NewDefaultConnectionService string
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
	var request _Layer3Forwarding1_GetDefaultConnectionService_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _Layer3Forwarding1_GetDefaultConnectionService_Response
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

// _WANCableLinkConfig1_GetCableLinkConfigInfo_Request is the XML structure for the input arguments for action GetCableLinkConfigInfo.
type _WANCableLinkConfig1_GetCableLinkConfigInfo_Request struct{}

// _WANCableLinkConfig1_GetCableLinkConfigInfo_Response is the XML structure for the output arguments for action GetCableLinkConfigInfo.
type _WANCableLinkConfig1_GetCableLinkConfigInfo_Response struct {
	NewCableLinkConfigState string

	NewLinkType string
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
	var request _WANCableLinkConfig1_GetCableLinkConfigInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetCableLinkConfigInfo_Response
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

// _WANCableLinkConfig1_GetDownstreamFrequency_Request is the XML structure for the input arguments for action GetDownstreamFrequency.
type _WANCableLinkConfig1_GetDownstreamFrequency_Request struct{}

// _WANCableLinkConfig1_GetDownstreamFrequency_Response is the XML structure for the output arguments for action GetDownstreamFrequency.
type _WANCableLinkConfig1_GetDownstreamFrequency_Response struct {
	NewDownstreamFrequency string
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
	var request _WANCableLinkConfig1_GetDownstreamFrequency_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetDownstreamFrequency_Response
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

// _WANCableLinkConfig1_GetDownstreamModulation_Request is the XML structure for the input arguments for action GetDownstreamModulation.
type _WANCableLinkConfig1_GetDownstreamModulation_Request struct{}

// _WANCableLinkConfig1_GetDownstreamModulation_Response is the XML structure for the output arguments for action GetDownstreamModulation.
type _WANCableLinkConfig1_GetDownstreamModulation_Response struct {
	NewDownstreamModulation string
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
	var request _WANCableLinkConfig1_GetDownstreamModulation_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetDownstreamModulation_Response
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

// _WANCableLinkConfig1_GetUpstreamFrequency_Request is the XML structure for the input arguments for action GetUpstreamFrequency.
type _WANCableLinkConfig1_GetUpstreamFrequency_Request struct{}

// _WANCableLinkConfig1_GetUpstreamFrequency_Response is the XML structure for the output arguments for action GetUpstreamFrequency.
type _WANCableLinkConfig1_GetUpstreamFrequency_Response struct {
	NewUpstreamFrequency string
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
	var request _WANCableLinkConfig1_GetUpstreamFrequency_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetUpstreamFrequency_Response
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

// _WANCableLinkConfig1_GetUpstreamModulation_Request is the XML structure for the input arguments for action GetUpstreamModulation.
type _WANCableLinkConfig1_GetUpstreamModulation_Request struct{}

// _WANCableLinkConfig1_GetUpstreamModulation_Response is the XML structure for the output arguments for action GetUpstreamModulation.
type _WANCableLinkConfig1_GetUpstreamModulation_Response struct {
	NewUpstreamModulation string
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
	var request _WANCableLinkConfig1_GetUpstreamModulation_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetUpstreamModulation_Response
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

// _WANCableLinkConfig1_GetUpstreamChannelID_Request is the XML structure for the input arguments for action GetUpstreamChannelID.
type _WANCableLinkConfig1_GetUpstreamChannelID_Request struct{}

// _WANCableLinkConfig1_GetUpstreamChannelID_Response is the XML structure for the output arguments for action GetUpstreamChannelID.
type _WANCableLinkConfig1_GetUpstreamChannelID_Response struct {
	NewUpstreamChannelID string
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
	var request _WANCableLinkConfig1_GetUpstreamChannelID_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetUpstreamChannelID_Response
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

// _WANCableLinkConfig1_GetUpstreamPowerLevel_Request is the XML structure for the input arguments for action GetUpstreamPowerLevel.
type _WANCableLinkConfig1_GetUpstreamPowerLevel_Request struct{}

// _WANCableLinkConfig1_GetUpstreamPowerLevel_Response is the XML structure for the output arguments for action GetUpstreamPowerLevel.
type _WANCableLinkConfig1_GetUpstreamPowerLevel_Response struct {
	NewUpstreamPowerLevel string
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
	var request _WANCableLinkConfig1_GetUpstreamPowerLevel_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetUpstreamPowerLevel_Response
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

// _WANCableLinkConfig1_GetBPIEncryptionEnabled_Request is the XML structure for the input arguments for action GetBPIEncryptionEnabled.
type _WANCableLinkConfig1_GetBPIEncryptionEnabled_Request struct{}

// _WANCableLinkConfig1_GetBPIEncryptionEnabled_Response is the XML structure for the output arguments for action GetBPIEncryptionEnabled.
type _WANCableLinkConfig1_GetBPIEncryptionEnabled_Response struct {
	NewBPIEncryptionEnabled string
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
	var request _WANCableLinkConfig1_GetBPIEncryptionEnabled_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetBPIEncryptionEnabled_Response
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

// _WANCableLinkConfig1_GetConfigFile_Request is the XML structure for the input arguments for action GetConfigFile.
type _WANCableLinkConfig1_GetConfigFile_Request struct{}

// _WANCableLinkConfig1_GetConfigFile_Response is the XML structure for the output arguments for action GetConfigFile.
type _WANCableLinkConfig1_GetConfigFile_Response struct {
	NewConfigFile string
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
	var request _WANCableLinkConfig1_GetConfigFile_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetConfigFile_Response
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

// _WANCableLinkConfig1_GetTFTPServer_Request is the XML structure for the input arguments for action GetTFTPServer.
type _WANCableLinkConfig1_GetTFTPServer_Request struct{}

// _WANCableLinkConfig1_GetTFTPServer_Response is the XML structure for the output arguments for action GetTFTPServer.
type _WANCableLinkConfig1_GetTFTPServer_Response struct {
	NewTFTPServer string
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
	var request _WANCableLinkConfig1_GetTFTPServer_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCableLinkConfig1_GetTFTPServer_Response
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

// _WANCommonInterfaceConfig1_SetEnabledForInternet_Request is the XML structure for the input arguments for action SetEnabledForInternet.
type _WANCommonInterfaceConfig1_SetEnabledForInternet_Request struct {
	NewEnabledForInternet string
}

// _WANCommonInterfaceConfig1_SetEnabledForInternet_Response is the XML structure for the output arguments for action SetEnabledForInternet.
type _WANCommonInterfaceConfig1_SetEnabledForInternet_Response struct{}

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
	var request _WANCommonInterfaceConfig1_SetEnabledForInternet_Request
	// BEGIN Marshal arguments into request.

	if request.NewEnabledForInternet, err = soap.MarshalBoolean(NewEnabledForInternet); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_SetEnabledForInternet_Response
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "SetEnabledForInternet", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANCommonInterfaceConfig1_GetEnabledForInternet_Request is the XML structure for the input arguments for action GetEnabledForInternet.
type _WANCommonInterfaceConfig1_GetEnabledForInternet_Request struct{}

// _WANCommonInterfaceConfig1_GetEnabledForInternet_Response is the XML structure for the output arguments for action GetEnabledForInternet.
type _WANCommonInterfaceConfig1_GetEnabledForInternet_Response struct {
	NewEnabledForInternet string
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
	var request _WANCommonInterfaceConfig1_GetEnabledForInternet_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetEnabledForInternet_Response
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

// _WANCommonInterfaceConfig1_GetCommonLinkProperties_Request is the XML structure for the input arguments for action GetCommonLinkProperties.
type _WANCommonInterfaceConfig1_GetCommonLinkProperties_Request struct{}

// _WANCommonInterfaceConfig1_GetCommonLinkProperties_Response is the XML structure for the output arguments for action GetCommonLinkProperties.
type _WANCommonInterfaceConfig1_GetCommonLinkProperties_Response struct {
	NewWANAccessType string

	NewLayer1UpstreamMaxBitRate string

	NewLayer1DownstreamMaxBitRate string

	NewPhysicalLinkStatus string
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
	var request _WANCommonInterfaceConfig1_GetCommonLinkProperties_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetCommonLinkProperties_Response
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

// _WANCommonInterfaceConfig1_GetWANAccessProvider_Request is the XML structure for the input arguments for action GetWANAccessProvider.
type _WANCommonInterfaceConfig1_GetWANAccessProvider_Request struct{}

// _WANCommonInterfaceConfig1_GetWANAccessProvider_Response is the XML structure for the output arguments for action GetWANAccessProvider.
type _WANCommonInterfaceConfig1_GetWANAccessProvider_Response struct {
	NewWANAccessProvider string
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
	var request _WANCommonInterfaceConfig1_GetWANAccessProvider_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetWANAccessProvider_Response
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

// _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Request is the XML structure for the input arguments for action GetMaximumActiveConnections.
type _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Request struct{}

// _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Response is the XML structure for the output arguments for action GetMaximumActiveConnections.
type _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Response struct {
	NewMaximumActiveConnections string
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
	var request _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Response
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

// _WANCommonInterfaceConfig1_GetTotalBytesSent_Request is the XML structure for the input arguments for action GetTotalBytesSent.
type _WANCommonInterfaceConfig1_GetTotalBytesSent_Request struct{}

// _WANCommonInterfaceConfig1_GetTotalBytesSent_Response is the XML structure for the output arguments for action GetTotalBytesSent.
type _WANCommonInterfaceConfig1_GetTotalBytesSent_Response struct {
	NewTotalBytesSent string
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
	var request _WANCommonInterfaceConfig1_GetTotalBytesSent_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetTotalBytesSent_Response
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

// _WANCommonInterfaceConfig1_GetTotalBytesReceived_Request is the XML structure for the input arguments for action GetTotalBytesReceived.
type _WANCommonInterfaceConfig1_GetTotalBytesReceived_Request struct{}

// _WANCommonInterfaceConfig1_GetTotalBytesReceived_Response is the XML structure for the output arguments for action GetTotalBytesReceived.
type _WANCommonInterfaceConfig1_GetTotalBytesReceived_Response struct {
	NewTotalBytesReceived string
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
	var request _WANCommonInterfaceConfig1_GetTotalBytesReceived_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetTotalBytesReceived_Response
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

// _WANCommonInterfaceConfig1_GetTotalPacketsSent_Request is the XML structure for the input arguments for action GetTotalPacketsSent.
type _WANCommonInterfaceConfig1_GetTotalPacketsSent_Request struct{}

// _WANCommonInterfaceConfig1_GetTotalPacketsSent_Response is the XML structure for the output arguments for action GetTotalPacketsSent.
type _WANCommonInterfaceConfig1_GetTotalPacketsSent_Response struct {
	NewTotalPacketsSent string
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
	var request _WANCommonInterfaceConfig1_GetTotalPacketsSent_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetTotalPacketsSent_Response
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

// _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Request is the XML structure for the input arguments for action GetTotalPacketsReceived.
type _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Request struct{}

// _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Response is the XML structure for the output arguments for action GetTotalPacketsReceived.
type _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Response struct {
	NewTotalPacketsReceived string
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
	var request _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Response
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

// _WANCommonInterfaceConfig1_GetActiveConnection_Request is the XML structure for the input arguments for action GetActiveConnection.
type _WANCommonInterfaceConfig1_GetActiveConnection_Request struct {
	NewActiveConnectionIndex string
}

// _WANCommonInterfaceConfig1_GetActiveConnection_Response is the XML structure for the output arguments for action GetActiveConnection.
type _WANCommonInterfaceConfig1_GetActiveConnection_Response struct {
	NewActiveConnDeviceContainer string

	NewActiveConnectionServiceID string
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
	var request _WANCommonInterfaceConfig1_GetActiveConnection_Request
	// BEGIN Marshal arguments into request.

	if request.NewActiveConnectionIndex, err = soap.MarshalUi2(NewActiveConnectionIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANCommonInterfaceConfig1_GetActiveConnection_Response
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

// _WANDSLLinkConfig1_SetDSLLinkType_Request is the XML structure for the input arguments for action SetDSLLinkType.
type _WANDSLLinkConfig1_SetDSLLinkType_Request struct {
	NewLinkType string
}

// _WANDSLLinkConfig1_SetDSLLinkType_Response is the XML structure for the output arguments for action SetDSLLinkType.
type _WANDSLLinkConfig1_SetDSLLinkType_Response struct{}

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
	var request _WANDSLLinkConfig1_SetDSLLinkType_Request
	// BEGIN Marshal arguments into request.

	if request.NewLinkType, err = soap.MarshalString(NewLinkType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_SetDSLLinkType_Response
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDSLLinkType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANDSLLinkConfig1_GetDSLLinkInfo_Request is the XML structure for the input arguments for action GetDSLLinkInfo.
type _WANDSLLinkConfig1_GetDSLLinkInfo_Request struct{}

// _WANDSLLinkConfig1_GetDSLLinkInfo_Response is the XML structure for the output arguments for action GetDSLLinkInfo.
type _WANDSLLinkConfig1_GetDSLLinkInfo_Response struct {
	NewLinkType string

	NewLinkStatus string
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
	var request _WANDSLLinkConfig1_GetDSLLinkInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_GetDSLLinkInfo_Response
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

// _WANDSLLinkConfig1_GetAutoConfig_Request is the XML structure for the input arguments for action GetAutoConfig.
type _WANDSLLinkConfig1_GetAutoConfig_Request struct{}

// _WANDSLLinkConfig1_GetAutoConfig_Response is the XML structure for the output arguments for action GetAutoConfig.
type _WANDSLLinkConfig1_GetAutoConfig_Response struct {
	NewAutoConfig string
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
	var request _WANDSLLinkConfig1_GetAutoConfig_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_GetAutoConfig_Response
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

// _WANDSLLinkConfig1_GetModulationType_Request is the XML structure for the input arguments for action GetModulationType.
type _WANDSLLinkConfig1_GetModulationType_Request struct{}

// _WANDSLLinkConfig1_GetModulationType_Response is the XML structure for the output arguments for action GetModulationType.
type _WANDSLLinkConfig1_GetModulationType_Response struct {
	NewModulationType string
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
	var request _WANDSLLinkConfig1_GetModulationType_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_GetModulationType_Response
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

// _WANDSLLinkConfig1_SetDestinationAddress_Request is the XML structure for the input arguments for action SetDestinationAddress.
type _WANDSLLinkConfig1_SetDestinationAddress_Request struct {
	NewDestinationAddress string
}

// _WANDSLLinkConfig1_SetDestinationAddress_Response is the XML structure for the output arguments for action SetDestinationAddress.
type _WANDSLLinkConfig1_SetDestinationAddress_Response struct{}

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
	var request _WANDSLLinkConfig1_SetDestinationAddress_Request
	// BEGIN Marshal arguments into request.

	if request.NewDestinationAddress, err = soap.MarshalString(NewDestinationAddress); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_SetDestinationAddress_Response
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDestinationAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANDSLLinkConfig1_GetDestinationAddress_Request is the XML structure for the input arguments for action GetDestinationAddress.
type _WANDSLLinkConfig1_GetDestinationAddress_Request struct{}

// _WANDSLLinkConfig1_GetDestinationAddress_Response is the XML structure for the output arguments for action GetDestinationAddress.
type _WANDSLLinkConfig1_GetDestinationAddress_Response struct {
	NewDestinationAddress string
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
	var request _WANDSLLinkConfig1_GetDestinationAddress_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_GetDestinationAddress_Response
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

// _WANDSLLinkConfig1_SetATMEncapsulation_Request is the XML structure for the input arguments for action SetATMEncapsulation.
type _WANDSLLinkConfig1_SetATMEncapsulation_Request struct {
	NewATMEncapsulation string
}

// _WANDSLLinkConfig1_SetATMEncapsulation_Response is the XML structure for the output arguments for action SetATMEncapsulation.
type _WANDSLLinkConfig1_SetATMEncapsulation_Response struct{}

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
	var request _WANDSLLinkConfig1_SetATMEncapsulation_Request
	// BEGIN Marshal arguments into request.

	if request.NewATMEncapsulation, err = soap.MarshalString(NewATMEncapsulation); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_SetATMEncapsulation_Response
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetATMEncapsulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANDSLLinkConfig1_GetATMEncapsulation_Request is the XML structure for the input arguments for action GetATMEncapsulation.
type _WANDSLLinkConfig1_GetATMEncapsulation_Request struct{}

// _WANDSLLinkConfig1_GetATMEncapsulation_Response is the XML structure for the output arguments for action GetATMEncapsulation.
type _WANDSLLinkConfig1_GetATMEncapsulation_Response struct {
	NewATMEncapsulation string
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
	var request _WANDSLLinkConfig1_GetATMEncapsulation_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_GetATMEncapsulation_Response
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

// _WANDSLLinkConfig1_SetFCSPreserved_Request is the XML structure for the input arguments for action SetFCSPreserved.
type _WANDSLLinkConfig1_SetFCSPreserved_Request struct {
	NewFCSPreserved string
}

// _WANDSLLinkConfig1_SetFCSPreserved_Response is the XML structure for the output arguments for action SetFCSPreserved.
type _WANDSLLinkConfig1_SetFCSPreserved_Response struct{}

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
	var request _WANDSLLinkConfig1_SetFCSPreserved_Request
	// BEGIN Marshal arguments into request.

	if request.NewFCSPreserved, err = soap.MarshalBoolean(NewFCSPreserved); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_SetFCSPreserved_Response
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetFCSPreserved", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANDSLLinkConfig1_GetFCSPreserved_Request is the XML structure for the input arguments for action GetFCSPreserved.
type _WANDSLLinkConfig1_GetFCSPreserved_Request struct{}

// _WANDSLLinkConfig1_GetFCSPreserved_Response is the XML structure for the output arguments for action GetFCSPreserved.
type _WANDSLLinkConfig1_GetFCSPreserved_Response struct {
	NewFCSPreserved string
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
	var request _WANDSLLinkConfig1_GetFCSPreserved_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANDSLLinkConfig1_GetFCSPreserved_Response
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

// _WANEthernetLinkConfig1_GetEthernetLinkStatus_Request is the XML structure for the input arguments for action GetEthernetLinkStatus.
type _WANEthernetLinkConfig1_GetEthernetLinkStatus_Request struct{}

// _WANEthernetLinkConfig1_GetEthernetLinkStatus_Response is the XML structure for the output arguments for action GetEthernetLinkStatus.
type _WANEthernetLinkConfig1_GetEthernetLinkStatus_Response struct {
	NewEthernetLinkStatus string
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
	var request _WANEthernetLinkConfig1_GetEthernetLinkStatus_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANEthernetLinkConfig1_GetEthernetLinkStatus_Response
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

// _WANIPConnection1_SetConnectionType_Request is the XML structure for the input arguments for action SetConnectionType.
type _WANIPConnection1_SetConnectionType_Request struct {
	NewConnectionType string
}

// _WANIPConnection1_SetConnectionType_Response is the XML structure for the output arguments for action SetConnectionType.
type _WANIPConnection1_SetConnectionType_Response struct{}

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
	var request _WANIPConnection1_SetConnectionType_Request
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_SetConnectionType_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_GetConnectionTypeInfo_Request is the XML structure for the input arguments for action GetConnectionTypeInfo.
type _WANIPConnection1_GetConnectionTypeInfo_Request struct{}

// _WANIPConnection1_GetConnectionTypeInfo_Response is the XML structure for the output arguments for action GetConnectionTypeInfo.
type _WANIPConnection1_GetConnectionTypeInfo_Response struct {
	NewConnectionType string

	NewPossibleConnectionTypes string
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
	var request _WANIPConnection1_GetConnectionTypeInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetConnectionTypeInfo_Response
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

// _WANIPConnection1_RequestConnection_Request is the XML structure for the input arguments for action RequestConnection.
type _WANIPConnection1_RequestConnection_Request struct{}

// _WANIPConnection1_RequestConnection_Response is the XML structure for the output arguments for action RequestConnection.
type _WANIPConnection1_RequestConnection_Response struct{}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) RequestConnection() (
	err error,
) {
	var request _WANIPConnection1_RequestConnection_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_RequestConnection_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_RequestTermination_Request is the XML structure for the input arguments for action RequestTermination.
type _WANIPConnection1_RequestTermination_Request struct{}

// _WANIPConnection1_RequestTermination_Response is the XML structure for the output arguments for action RequestTermination.
type _WANIPConnection1_RequestTermination_Response struct{}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) RequestTermination() (
	err error,
) {
	var request _WANIPConnection1_RequestTermination_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_RequestTermination_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_ForceTermination_Request is the XML structure for the input arguments for action ForceTermination.
type _WANIPConnection1_ForceTermination_Request struct{}

// _WANIPConnection1_ForceTermination_Response is the XML structure for the output arguments for action ForceTermination.
type _WANIPConnection1_ForceTermination_Response struct{}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) ForceTermination() (
	err error,
) {
	var request _WANIPConnection1_ForceTermination_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_ForceTermination_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_SetAutoDisconnectTime_Request is the XML structure for the input arguments for action SetAutoDisconnectTime.
type _WANIPConnection1_SetAutoDisconnectTime_Request struct {
	NewAutoDisconnectTime string
}

// _WANIPConnection1_SetAutoDisconnectTime_Response is the XML structure for the output arguments for action SetAutoDisconnectTime.
type _WANIPConnection1_SetAutoDisconnectTime_Response struct{}

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
	var request _WANIPConnection1_SetAutoDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_SetAutoDisconnectTime_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_SetIdleDisconnectTime_Request is the XML structure for the input arguments for action SetIdleDisconnectTime.
type _WANIPConnection1_SetIdleDisconnectTime_Request struct {
	NewIdleDisconnectTime string
}

// _WANIPConnection1_SetIdleDisconnectTime_Response is the XML structure for the output arguments for action SetIdleDisconnectTime.
type _WANIPConnection1_SetIdleDisconnectTime_Response struct{}

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
	var request _WANIPConnection1_SetIdleDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_SetIdleDisconnectTime_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_SetWarnDisconnectDelay_Request is the XML structure for the input arguments for action SetWarnDisconnectDelay.
type _WANIPConnection1_SetWarnDisconnectDelay_Request struct {
	NewWarnDisconnectDelay string
}

// _WANIPConnection1_SetWarnDisconnectDelay_Response is the XML structure for the output arguments for action SetWarnDisconnectDelay.
type _WANIPConnection1_SetWarnDisconnectDelay_Response struct{}

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
	var request _WANIPConnection1_SetWarnDisconnectDelay_Request
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_SetWarnDisconnectDelay_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_GetStatusInfo_Request is the XML structure for the input arguments for action GetStatusInfo.
type _WANIPConnection1_GetStatusInfo_Request struct{}

// _WANIPConnection1_GetStatusInfo_Response is the XML structure for the output arguments for action GetStatusInfo.
type _WANIPConnection1_GetStatusInfo_Response struct {
	NewConnectionStatus string

	NewLastConnectionError string

	NewUptime string
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
	var request _WANIPConnection1_GetStatusInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetStatusInfo_Response
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

// _WANIPConnection1_GetAutoDisconnectTime_Request is the XML structure for the input arguments for action GetAutoDisconnectTime.
type _WANIPConnection1_GetAutoDisconnectTime_Request struct{}

// _WANIPConnection1_GetAutoDisconnectTime_Response is the XML structure for the output arguments for action GetAutoDisconnectTime.
type _WANIPConnection1_GetAutoDisconnectTime_Response struct {
	NewAutoDisconnectTime string
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
	var request _WANIPConnection1_GetAutoDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetAutoDisconnectTime_Response
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

// _WANIPConnection1_GetIdleDisconnectTime_Request is the XML structure for the input arguments for action GetIdleDisconnectTime.
type _WANIPConnection1_GetIdleDisconnectTime_Request struct{}

// _WANIPConnection1_GetIdleDisconnectTime_Response is the XML structure for the output arguments for action GetIdleDisconnectTime.
type _WANIPConnection1_GetIdleDisconnectTime_Response struct {
	NewIdleDisconnectTime string
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
	var request _WANIPConnection1_GetIdleDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetIdleDisconnectTime_Response
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

// _WANIPConnection1_GetWarnDisconnectDelay_Request is the XML structure for the input arguments for action GetWarnDisconnectDelay.
type _WANIPConnection1_GetWarnDisconnectDelay_Request struct{}

// _WANIPConnection1_GetWarnDisconnectDelay_Response is the XML structure for the output arguments for action GetWarnDisconnectDelay.
type _WANIPConnection1_GetWarnDisconnectDelay_Response struct {
	NewWarnDisconnectDelay string
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
	var request _WANIPConnection1_GetWarnDisconnectDelay_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetWarnDisconnectDelay_Response
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

// _WANIPConnection1_GetNATRSIPStatus_Request is the XML structure for the input arguments for action GetNATRSIPStatus.
type _WANIPConnection1_GetNATRSIPStatus_Request struct{}

// _WANIPConnection1_GetNATRSIPStatus_Response is the XML structure for the output arguments for action GetNATRSIPStatus.
type _WANIPConnection1_GetNATRSIPStatus_Response struct {
	NewRSIPAvailable string

	NewNATEnabled string
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
	var request _WANIPConnection1_GetNATRSIPStatus_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetNATRSIPStatus_Response
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

// _WANIPConnection1_GetGenericPortMappingEntry_Request is the XML structure for the input arguments for action GetGenericPortMappingEntry.
type _WANIPConnection1_GetGenericPortMappingEntry_Request struct {
	NewPortMappingIndex string
}

// _WANIPConnection1_GetGenericPortMappingEntry_Response is the XML structure for the output arguments for action GetGenericPortMappingEntry.
type _WANIPConnection1_GetGenericPortMappingEntry_Response struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
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
	var request _WANIPConnection1_GetGenericPortMappingEntry_Request
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetGenericPortMappingEntry_Response
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

// _WANIPConnection1_GetSpecificPortMappingEntry_Request is the XML structure for the input arguments for action GetSpecificPortMappingEntry.
type _WANIPConnection1_GetSpecificPortMappingEntry_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string
}

// _WANIPConnection1_GetSpecificPortMappingEntry_Response is the XML structure for the output arguments for action GetSpecificPortMappingEntry.
type _WANIPConnection1_GetSpecificPortMappingEntry_Response struct {
	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
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
	var request _WANIPConnection1_GetSpecificPortMappingEntry_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection1_GetSpecificPortMappingEntry_Response
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

// _WANIPConnection1_AddPortMapping_Request is the XML structure for the input arguments for action AddPortMapping.
type _WANIPConnection1_AddPortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
}

// _WANIPConnection1_AddPortMapping_Response is the XML structure for the output arguments for action AddPortMapping.
type _WANIPConnection1_AddPortMapping_Response struct{}

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
	var request _WANIPConnection1_AddPortMapping_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection1_AddPortMapping_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_DeletePortMapping_Request is the XML structure for the input arguments for action DeletePortMapping.
type _WANIPConnection1_DeletePortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string
}

// _WANIPConnection1_DeletePortMapping_Response is the XML structure for the output arguments for action DeletePortMapping.
type _WANIPConnection1_DeletePortMapping_Response struct{}

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
	var request _WANIPConnection1_DeletePortMapping_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection1_DeletePortMapping_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection1_GetExternalIPAddress_Request is the XML structure for the input arguments for action GetExternalIPAddress.
type _WANIPConnection1_GetExternalIPAddress_Request struct{}

// _WANIPConnection1_GetExternalIPAddress_Response is the XML structure for the output arguments for action GetExternalIPAddress.
type _WANIPConnection1_GetExternalIPAddress_Response struct {
	NewExternalIPAddress string
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
	var request _WANIPConnection1_GetExternalIPAddress_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection1_GetExternalIPAddress_Response
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

// _WANIPConnection2_SetConnectionType_Request is the XML structure for the input arguments for action SetConnectionType.
type _WANIPConnection2_SetConnectionType_Request struct {
	NewConnectionType string
}

// _WANIPConnection2_SetConnectionType_Response is the XML structure for the output arguments for action SetConnectionType.
type _WANIPConnection2_SetConnectionType_Response struct{}

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
	var request _WANIPConnection2_SetConnectionType_Request
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_SetConnectionType_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_GetConnectionTypeInfo_Request is the XML structure for the input arguments for action GetConnectionTypeInfo.
type _WANIPConnection2_GetConnectionTypeInfo_Request struct{}

// _WANIPConnection2_GetConnectionTypeInfo_Response is the XML structure for the output arguments for action GetConnectionTypeInfo.
type _WANIPConnection2_GetConnectionTypeInfo_Response struct {
	NewConnectionType string

	NewPossibleConnectionTypes string
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
	var request _WANIPConnection2_GetConnectionTypeInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetConnectionTypeInfo_Response
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

// _WANIPConnection2_RequestConnection_Request is the XML structure for the input arguments for action RequestConnection.
type _WANIPConnection2_RequestConnection_Request struct{}

// _WANIPConnection2_RequestConnection_Response is the XML structure for the output arguments for action RequestConnection.
type _WANIPConnection2_RequestConnection_Response struct{}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection2) RequestConnection() (
	err error,
) {
	var request _WANIPConnection2_RequestConnection_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_RequestConnection_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_RequestTermination_Request is the XML structure for the input arguments for action RequestTermination.
type _WANIPConnection2_RequestTermination_Request struct{}

// _WANIPConnection2_RequestTermination_Response is the XML structure for the output arguments for action RequestTermination.
type _WANIPConnection2_RequestTermination_Response struct{}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection2) RequestTermination() (
	err error,
) {
	var request _WANIPConnection2_RequestTermination_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_RequestTermination_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_ForceTermination_Request is the XML structure for the input arguments for action ForceTermination.
type _WANIPConnection2_ForceTermination_Request struct{}

// _WANIPConnection2_ForceTermination_Response is the XML structure for the output arguments for action ForceTermination.
type _WANIPConnection2_ForceTermination_Response struct{}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection2) ForceTermination() (
	err error,
) {
	var request _WANIPConnection2_ForceTermination_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_ForceTermination_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_SetAutoDisconnectTime_Request is the XML structure for the input arguments for action SetAutoDisconnectTime.
type _WANIPConnection2_SetAutoDisconnectTime_Request struct {
	NewAutoDisconnectTime string
}

// _WANIPConnection2_SetAutoDisconnectTime_Response is the XML structure for the output arguments for action SetAutoDisconnectTime.
type _WANIPConnection2_SetAutoDisconnectTime_Response struct{}

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
	var request _WANIPConnection2_SetAutoDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_SetAutoDisconnectTime_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_SetIdleDisconnectTime_Request is the XML structure for the input arguments for action SetIdleDisconnectTime.
type _WANIPConnection2_SetIdleDisconnectTime_Request struct {
	NewIdleDisconnectTime string
}

// _WANIPConnection2_SetIdleDisconnectTime_Response is the XML structure for the output arguments for action SetIdleDisconnectTime.
type _WANIPConnection2_SetIdleDisconnectTime_Response struct{}

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
	var request _WANIPConnection2_SetIdleDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_SetIdleDisconnectTime_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_SetWarnDisconnectDelay_Request is the XML structure for the input arguments for action SetWarnDisconnectDelay.
type _WANIPConnection2_SetWarnDisconnectDelay_Request struct {
	NewWarnDisconnectDelay string
}

// _WANIPConnection2_SetWarnDisconnectDelay_Response is the XML structure for the output arguments for action SetWarnDisconnectDelay.
type _WANIPConnection2_SetWarnDisconnectDelay_Response struct{}

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
	var request _WANIPConnection2_SetWarnDisconnectDelay_Request
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_SetWarnDisconnectDelay_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_GetStatusInfo_Request is the XML structure for the input arguments for action GetStatusInfo.
type _WANIPConnection2_GetStatusInfo_Request struct{}

// _WANIPConnection2_GetStatusInfo_Response is the XML structure for the output arguments for action GetStatusInfo.
type _WANIPConnection2_GetStatusInfo_Response struct {
	NewConnectionStatus string

	NewLastConnectionError string

	NewUptime string
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
	var request _WANIPConnection2_GetStatusInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetStatusInfo_Response
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

// _WANIPConnection2_GetAutoDisconnectTime_Request is the XML structure for the input arguments for action GetAutoDisconnectTime.
type _WANIPConnection2_GetAutoDisconnectTime_Request struct{}

// _WANIPConnection2_GetAutoDisconnectTime_Response is the XML structure for the output arguments for action GetAutoDisconnectTime.
type _WANIPConnection2_GetAutoDisconnectTime_Response struct {
	NewAutoDisconnectTime string
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
	var request _WANIPConnection2_GetAutoDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetAutoDisconnectTime_Response
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

// _WANIPConnection2_GetIdleDisconnectTime_Request is the XML structure for the input arguments for action GetIdleDisconnectTime.
type _WANIPConnection2_GetIdleDisconnectTime_Request struct{}

// _WANIPConnection2_GetIdleDisconnectTime_Response is the XML structure for the output arguments for action GetIdleDisconnectTime.
type _WANIPConnection2_GetIdleDisconnectTime_Response struct {
	NewIdleDisconnectTime string
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
	var request _WANIPConnection2_GetIdleDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetIdleDisconnectTime_Response
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

// _WANIPConnection2_GetWarnDisconnectDelay_Request is the XML structure for the input arguments for action GetWarnDisconnectDelay.
type _WANIPConnection2_GetWarnDisconnectDelay_Request struct{}

// _WANIPConnection2_GetWarnDisconnectDelay_Response is the XML structure for the output arguments for action GetWarnDisconnectDelay.
type _WANIPConnection2_GetWarnDisconnectDelay_Response struct {
	NewWarnDisconnectDelay string
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
	var request _WANIPConnection2_GetWarnDisconnectDelay_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetWarnDisconnectDelay_Response
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

// _WANIPConnection2_GetNATRSIPStatus_Request is the XML structure for the input arguments for action GetNATRSIPStatus.
type _WANIPConnection2_GetNATRSIPStatus_Request struct{}

// _WANIPConnection2_GetNATRSIPStatus_Response is the XML structure for the output arguments for action GetNATRSIPStatus.
type _WANIPConnection2_GetNATRSIPStatus_Response struct {
	NewRSIPAvailable string

	NewNATEnabled string
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
	var request _WANIPConnection2_GetNATRSIPStatus_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetNATRSIPStatus_Response
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

// _WANIPConnection2_GetGenericPortMappingEntry_Request is the XML structure for the input arguments for action GetGenericPortMappingEntry.
type _WANIPConnection2_GetGenericPortMappingEntry_Request struct {
	NewPortMappingIndex string
}

// _WANIPConnection2_GetGenericPortMappingEntry_Response is the XML structure for the output arguments for action GetGenericPortMappingEntry.
type _WANIPConnection2_GetGenericPortMappingEntry_Response struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
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
	var request _WANIPConnection2_GetGenericPortMappingEntry_Request
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetGenericPortMappingEntry_Response
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

// _WANIPConnection2_GetSpecificPortMappingEntry_Request is the XML structure for the input arguments for action GetSpecificPortMappingEntry.
type _WANIPConnection2_GetSpecificPortMappingEntry_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string
}

// _WANIPConnection2_GetSpecificPortMappingEntry_Response is the XML structure for the output arguments for action GetSpecificPortMappingEntry.
type _WANIPConnection2_GetSpecificPortMappingEntry_Response struct {
	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
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
	var request _WANIPConnection2_GetSpecificPortMappingEntry_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection2_GetSpecificPortMappingEntry_Response
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

// _WANIPConnection2_AddPortMapping_Request is the XML structure for the input arguments for action AddPortMapping.
type _WANIPConnection2_AddPortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
}

// _WANIPConnection2_AddPortMapping_Response is the XML structure for the output arguments for action AddPortMapping.
type _WANIPConnection2_AddPortMapping_Response struct{}

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
	var request _WANIPConnection2_AddPortMapping_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection2_AddPortMapping_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_DeletePortMapping_Request is the XML structure for the input arguments for action DeletePortMapping.
type _WANIPConnection2_DeletePortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string
}

// _WANIPConnection2_DeletePortMapping_Response is the XML structure for the output arguments for action DeletePortMapping.
type _WANIPConnection2_DeletePortMapping_Response struct{}

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
	var request _WANIPConnection2_DeletePortMapping_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection2_DeletePortMapping_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_DeletePortMappingRange_Request is the XML structure for the input arguments for action DeletePortMappingRange.
type _WANIPConnection2_DeletePortMappingRange_Request struct {
	NewStartPort string

	NewEndPort string

	NewProtocol string

	NewManage string
}

// _WANIPConnection2_DeletePortMappingRange_Response is the XML structure for the output arguments for action DeletePortMappingRange.
type _WANIPConnection2_DeletePortMappingRange_Response struct{}

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
	var request _WANIPConnection2_DeletePortMappingRange_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection2_DeletePortMappingRange_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "DeletePortMappingRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPConnection2_GetExternalIPAddress_Request is the XML structure for the input arguments for action GetExternalIPAddress.
type _WANIPConnection2_GetExternalIPAddress_Request struct{}

// _WANIPConnection2_GetExternalIPAddress_Response is the XML structure for the output arguments for action GetExternalIPAddress.
type _WANIPConnection2_GetExternalIPAddress_Response struct {
	NewExternalIPAddress string
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
	var request _WANIPConnection2_GetExternalIPAddress_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPConnection2_GetExternalIPAddress_Response
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

// _WANIPConnection2_GetListOfPortMappings_Request is the XML structure for the input arguments for action GetListOfPortMappings.
type _WANIPConnection2_GetListOfPortMappings_Request struct {
	NewStartPort string

	NewEndPort string

	NewProtocol string

	NewManage string

	NewNumberOfPorts string
}

// _WANIPConnection2_GetListOfPortMappings_Response is the XML structure for the output arguments for action GetListOfPortMappings.
type _WANIPConnection2_GetListOfPortMappings_Response struct {
	NewPortListing string
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
	var request _WANIPConnection2_GetListOfPortMappings_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection2_GetListOfPortMappings_Response
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

// _WANIPConnection2_AddAnyPortMapping_Request is the XML structure for the input arguments for action AddAnyPortMapping.
type _WANIPConnection2_AddAnyPortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
}

// _WANIPConnection2_AddAnyPortMapping_Response is the XML structure for the output arguments for action AddAnyPortMapping.
type _WANIPConnection2_AddAnyPortMapping_Response struct {
	NewReservedPort string
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
	var request _WANIPConnection2_AddAnyPortMapping_Request
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

	// Perform the SOAP call.
	var response _WANIPConnection2_AddAnyPortMapping_Response
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

// _WANIPv6FirewallControl1_GetFirewallStatus_Request is the XML structure for the input arguments for action GetFirewallStatus.
type _WANIPv6FirewallControl1_GetFirewallStatus_Request struct{}

// _WANIPv6FirewallControl1_GetFirewallStatus_Response is the XML structure for the output arguments for action GetFirewallStatus.
type _WANIPv6FirewallControl1_GetFirewallStatus_Response struct {
	FirewallEnabled string

	InboundPinholeAllowed string
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
	var request _WANIPv6FirewallControl1_GetFirewallStatus_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_GetFirewallStatus_Response
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

// _WANIPv6FirewallControl1_GetOutboundPinholeTimeout_Request is the XML structure for the input arguments for action GetOutboundPinholeTimeout.
type _WANIPv6FirewallControl1_GetOutboundPinholeTimeout_Request struct {
	RemoteHost string

	RemotePort string

	InternalClient string

	InternalPort string

	Protocol string
}

// _WANIPv6FirewallControl1_GetOutboundPinholeTimeout_Response is the XML structure for the output arguments for action GetOutboundPinholeTimeout.
type _WANIPv6FirewallControl1_GetOutboundPinholeTimeout_Response struct {
	OutboundPinholeTimeout string
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
	var request _WANIPv6FirewallControl1_GetOutboundPinholeTimeout_Request
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

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_GetOutboundPinholeTimeout_Response
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

// _WANIPv6FirewallControl1_AddPinhole_Request is the XML structure for the input arguments for action AddPinhole.
type _WANIPv6FirewallControl1_AddPinhole_Request struct {
	RemoteHost string

	RemotePort string

	InternalClient string

	InternalPort string

	Protocol string

	LeaseTime string
}

// _WANIPv6FirewallControl1_AddPinhole_Response is the XML structure for the output arguments for action AddPinhole.
type _WANIPv6FirewallControl1_AddPinhole_Response struct {
	UniqueID string
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
	var request _WANIPv6FirewallControl1_AddPinhole_Request
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

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_AddPinhole_Response
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

// _WANIPv6FirewallControl1_UpdatePinhole_Request is the XML structure for the input arguments for action UpdatePinhole.
type _WANIPv6FirewallControl1_UpdatePinhole_Request struct {
	UniqueID string

	NewLeaseTime string
}

// _WANIPv6FirewallControl1_UpdatePinhole_Response is the XML structure for the output arguments for action UpdatePinhole.
type _WANIPv6FirewallControl1_UpdatePinhole_Response struct{}

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
	var request _WANIPv6FirewallControl1_UpdatePinhole_Request
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	if request.NewLeaseTime, err = soap.MarshalUi4(NewLeaseTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_UpdatePinhole_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "UpdatePinhole", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPv6FirewallControl1_DeletePinhole_Request is the XML structure for the input arguments for action DeletePinhole.
type _WANIPv6FirewallControl1_DeletePinhole_Request struct {
	UniqueID string
}

// _WANIPv6FirewallControl1_DeletePinhole_Response is the XML structure for the output arguments for action DeletePinhole.
type _WANIPv6FirewallControl1_DeletePinhole_Response struct{}

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
	var request _WANIPv6FirewallControl1_DeletePinhole_Request
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_DeletePinhole_Response
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "DeletePinhole", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANIPv6FirewallControl1_GetPinholePackets_Request is the XML structure for the input arguments for action GetPinholePackets.
type _WANIPv6FirewallControl1_GetPinholePackets_Request struct {
	UniqueID string
}

// _WANIPv6FirewallControl1_GetPinholePackets_Response is the XML structure for the output arguments for action GetPinholePackets.
type _WANIPv6FirewallControl1_GetPinholePackets_Response struct {
	PinholePackets string
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
	var request _WANIPv6FirewallControl1_GetPinholePackets_Request
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_GetPinholePackets_Response
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

// _WANIPv6FirewallControl1_CheckPinholeWorking_Request is the XML structure for the input arguments for action CheckPinholeWorking.
type _WANIPv6FirewallControl1_CheckPinholeWorking_Request struct {
	UniqueID string
}

// _WANIPv6FirewallControl1_CheckPinholeWorking_Response is the XML structure for the output arguments for action CheckPinholeWorking.
type _WANIPv6FirewallControl1_CheckPinholeWorking_Response struct {
	IsWorking string
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
	var request _WANIPv6FirewallControl1_CheckPinholeWorking_Request
	// BEGIN Marshal arguments into request.

	if request.UniqueID, err = soap.MarshalUi2(UniqueID); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANIPv6FirewallControl1_CheckPinholeWorking_Response
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

// _WANPOTSLinkConfig1_SetISPInfo_Request is the XML structure for the input arguments for action SetISPInfo.
type _WANPOTSLinkConfig1_SetISPInfo_Request struct {
	NewISPPhoneNumber string

	NewISPInfo string

	NewLinkType string
}

// _WANPOTSLinkConfig1_SetISPInfo_Response is the XML structure for the output arguments for action SetISPInfo.
type _WANPOTSLinkConfig1_SetISPInfo_Response struct{}

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
	var request _WANPOTSLinkConfig1_SetISPInfo_Request
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

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_SetISPInfo_Response
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetISPInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPOTSLinkConfig1_SetCallRetryInfo_Request is the XML structure for the input arguments for action SetCallRetryInfo.
type _WANPOTSLinkConfig1_SetCallRetryInfo_Request struct {
	NewNumberOfRetries string

	NewDelayBetweenRetries string
}

// _WANPOTSLinkConfig1_SetCallRetryInfo_Response is the XML structure for the output arguments for action SetCallRetryInfo.
type _WANPOTSLinkConfig1_SetCallRetryInfo_Response struct{}

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
	var request _WANPOTSLinkConfig1_SetCallRetryInfo_Request
	// BEGIN Marshal arguments into request.

	if request.NewNumberOfRetries, err = soap.MarshalUi4(NewNumberOfRetries); err != nil {
		return
	}

	if request.NewDelayBetweenRetries, err = soap.MarshalUi4(NewDelayBetweenRetries); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_SetCallRetryInfo_Response
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetCallRetryInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPOTSLinkConfig1_GetISPInfo_Request is the XML structure for the input arguments for action GetISPInfo.
type _WANPOTSLinkConfig1_GetISPInfo_Request struct{}

// _WANPOTSLinkConfig1_GetISPInfo_Response is the XML structure for the output arguments for action GetISPInfo.
type _WANPOTSLinkConfig1_GetISPInfo_Response struct {
	NewISPPhoneNumber string

	NewISPInfo string

	NewLinkType string
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
	var request _WANPOTSLinkConfig1_GetISPInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetISPInfo_Response
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

// _WANPOTSLinkConfig1_GetCallRetryInfo_Request is the XML structure for the input arguments for action GetCallRetryInfo.
type _WANPOTSLinkConfig1_GetCallRetryInfo_Request struct{}

// _WANPOTSLinkConfig1_GetCallRetryInfo_Response is the XML structure for the output arguments for action GetCallRetryInfo.
type _WANPOTSLinkConfig1_GetCallRetryInfo_Response struct {
	NewNumberOfRetries string

	NewDelayBetweenRetries string
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
	var request _WANPOTSLinkConfig1_GetCallRetryInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetCallRetryInfo_Response
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

// _WANPOTSLinkConfig1_GetFclass_Request is the XML structure for the input arguments for action GetFclass.
type _WANPOTSLinkConfig1_GetFclass_Request struct{}

// _WANPOTSLinkConfig1_GetFclass_Response is the XML structure for the output arguments for action GetFclass.
type _WANPOTSLinkConfig1_GetFclass_Response struct {
	NewFclass string
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
	var request _WANPOTSLinkConfig1_GetFclass_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetFclass_Response
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

// _WANPOTSLinkConfig1_GetDataModulationSupported_Request is the XML structure for the input arguments for action GetDataModulationSupported.
type _WANPOTSLinkConfig1_GetDataModulationSupported_Request struct{}

// _WANPOTSLinkConfig1_GetDataModulationSupported_Response is the XML structure for the output arguments for action GetDataModulationSupported.
type _WANPOTSLinkConfig1_GetDataModulationSupported_Response struct {
	NewDataModulationSupported string
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
	var request _WANPOTSLinkConfig1_GetDataModulationSupported_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetDataModulationSupported_Response
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

// _WANPOTSLinkConfig1_GetDataProtocol_Request is the XML structure for the input arguments for action GetDataProtocol.
type _WANPOTSLinkConfig1_GetDataProtocol_Request struct{}

// _WANPOTSLinkConfig1_GetDataProtocol_Response is the XML structure for the output arguments for action GetDataProtocol.
type _WANPOTSLinkConfig1_GetDataProtocol_Response struct {
	NewDataProtocol string
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
	var request _WANPOTSLinkConfig1_GetDataProtocol_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetDataProtocol_Response
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

// _WANPOTSLinkConfig1_GetDataCompression_Request is the XML structure for the input arguments for action GetDataCompression.
type _WANPOTSLinkConfig1_GetDataCompression_Request struct{}

// _WANPOTSLinkConfig1_GetDataCompression_Response is the XML structure for the output arguments for action GetDataCompression.
type _WANPOTSLinkConfig1_GetDataCompression_Response struct {
	NewDataCompression string
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
	var request _WANPOTSLinkConfig1_GetDataCompression_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetDataCompression_Response
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

// _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Request is the XML structure for the input arguments for action GetPlusVTRCommandSupported.
type _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Request struct{}

// _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Response is the XML structure for the output arguments for action GetPlusVTRCommandSupported.
type _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Response struct {
	NewPlusVTRCommandSupported string
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
	var request _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Response
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

// _WANPPPConnection1_SetConnectionType_Request is the XML structure for the input arguments for action SetConnectionType.
type _WANPPPConnection1_SetConnectionType_Request struct {
	NewConnectionType string
}

// _WANPPPConnection1_SetConnectionType_Response is the XML structure for the output arguments for action SetConnectionType.
type _WANPPPConnection1_SetConnectionType_Response struct{}

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
	var request _WANPPPConnection1_SetConnectionType_Request
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_SetConnectionType_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_GetConnectionTypeInfo_Request is the XML structure for the input arguments for action GetConnectionTypeInfo.
type _WANPPPConnection1_GetConnectionTypeInfo_Request struct{}

// _WANPPPConnection1_GetConnectionTypeInfo_Response is the XML structure for the output arguments for action GetConnectionTypeInfo.
type _WANPPPConnection1_GetConnectionTypeInfo_Response struct {
	NewConnectionType string

	NewPossibleConnectionTypes string
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
	var request _WANPPPConnection1_GetConnectionTypeInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetConnectionTypeInfo_Response
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

// _WANPPPConnection1_ConfigureConnection_Request is the XML structure for the input arguments for action ConfigureConnection.
type _WANPPPConnection1_ConfigureConnection_Request struct {
	NewUserName string

	NewPassword string
}

// _WANPPPConnection1_ConfigureConnection_Response is the XML structure for the output arguments for action ConfigureConnection.
type _WANPPPConnection1_ConfigureConnection_Response struct{}

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
	var request _WANPPPConnection1_ConfigureConnection_Request
	// BEGIN Marshal arguments into request.

	if request.NewUserName, err = soap.MarshalString(NewUserName); err != nil {
		return
	}

	if request.NewPassword, err = soap.MarshalString(NewPassword); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_ConfigureConnection_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ConfigureConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_RequestConnection_Request is the XML structure for the input arguments for action RequestConnection.
type _WANPPPConnection1_RequestConnection_Request struct{}

// _WANPPPConnection1_RequestConnection_Response is the XML structure for the output arguments for action RequestConnection.
type _WANPPPConnection1_RequestConnection_Response struct{}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) RequestConnection() (
	err error,
) {
	var request _WANPPPConnection1_RequestConnection_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_RequestConnection_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_RequestTermination_Request is the XML structure for the input arguments for action RequestTermination.
type _WANPPPConnection1_RequestTermination_Request struct{}

// _WANPPPConnection1_RequestTermination_Response is the XML structure for the output arguments for action RequestTermination.
type _WANPPPConnection1_RequestTermination_Response struct{}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) RequestTermination() (
	err error,
) {
	var request _WANPPPConnection1_RequestTermination_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_RequestTermination_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_ForceTermination_Request is the XML structure for the input arguments for action ForceTermination.
type _WANPPPConnection1_ForceTermination_Request struct{}

// _WANPPPConnection1_ForceTermination_Response is the XML structure for the output arguments for action ForceTermination.
type _WANPPPConnection1_ForceTermination_Response struct{}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) ForceTermination() (
	err error,
) {
	var request _WANPPPConnection1_ForceTermination_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_ForceTermination_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_SetAutoDisconnectTime_Request is the XML structure for the input arguments for action SetAutoDisconnectTime.
type _WANPPPConnection1_SetAutoDisconnectTime_Request struct {
	NewAutoDisconnectTime string
}

// _WANPPPConnection1_SetAutoDisconnectTime_Response is the XML structure for the output arguments for action SetAutoDisconnectTime.
type _WANPPPConnection1_SetAutoDisconnectTime_Response struct{}

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
	var request _WANPPPConnection1_SetAutoDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_SetAutoDisconnectTime_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_SetIdleDisconnectTime_Request is the XML structure for the input arguments for action SetIdleDisconnectTime.
type _WANPPPConnection1_SetIdleDisconnectTime_Request struct {
	NewIdleDisconnectTime string
}

// _WANPPPConnection1_SetIdleDisconnectTime_Response is the XML structure for the output arguments for action SetIdleDisconnectTime.
type _WANPPPConnection1_SetIdleDisconnectTime_Response struct{}

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
	var request _WANPPPConnection1_SetIdleDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_SetIdleDisconnectTime_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_SetWarnDisconnectDelay_Request is the XML structure for the input arguments for action SetWarnDisconnectDelay.
type _WANPPPConnection1_SetWarnDisconnectDelay_Request struct {
	NewWarnDisconnectDelay string
}

// _WANPPPConnection1_SetWarnDisconnectDelay_Response is the XML structure for the output arguments for action SetWarnDisconnectDelay.
type _WANPPPConnection1_SetWarnDisconnectDelay_Response struct{}

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
	var request _WANPPPConnection1_SetWarnDisconnectDelay_Request
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_SetWarnDisconnectDelay_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_GetStatusInfo_Request is the XML structure for the input arguments for action GetStatusInfo.
type _WANPPPConnection1_GetStatusInfo_Request struct{}

// _WANPPPConnection1_GetStatusInfo_Response is the XML structure for the output arguments for action GetStatusInfo.
type _WANPPPConnection1_GetStatusInfo_Response struct {
	NewConnectionStatus string

	NewLastConnectionError string

	NewUptime string
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
	var request _WANPPPConnection1_GetStatusInfo_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetStatusInfo_Response
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

// _WANPPPConnection1_GetLinkLayerMaxBitRates_Request is the XML structure for the input arguments for action GetLinkLayerMaxBitRates.
type _WANPPPConnection1_GetLinkLayerMaxBitRates_Request struct{}

// _WANPPPConnection1_GetLinkLayerMaxBitRates_Response is the XML structure for the output arguments for action GetLinkLayerMaxBitRates.
type _WANPPPConnection1_GetLinkLayerMaxBitRates_Response struct {
	NewUpstreamMaxBitRate string

	NewDownstreamMaxBitRate string
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
	var request _WANPPPConnection1_GetLinkLayerMaxBitRates_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetLinkLayerMaxBitRates_Response
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

// _WANPPPConnection1_GetPPPEncryptionProtocol_Request is the XML structure for the input arguments for action GetPPPEncryptionProtocol.
type _WANPPPConnection1_GetPPPEncryptionProtocol_Request struct{}

// _WANPPPConnection1_GetPPPEncryptionProtocol_Response is the XML structure for the output arguments for action GetPPPEncryptionProtocol.
type _WANPPPConnection1_GetPPPEncryptionProtocol_Response struct {
	NewPPPEncryptionProtocol string
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
	var request _WANPPPConnection1_GetPPPEncryptionProtocol_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetPPPEncryptionProtocol_Response
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

// _WANPPPConnection1_GetPPPCompressionProtocol_Request is the XML structure for the input arguments for action GetPPPCompressionProtocol.
type _WANPPPConnection1_GetPPPCompressionProtocol_Request struct{}

// _WANPPPConnection1_GetPPPCompressionProtocol_Response is the XML structure for the output arguments for action GetPPPCompressionProtocol.
type _WANPPPConnection1_GetPPPCompressionProtocol_Response struct {
	NewPPPCompressionProtocol string
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
	var request _WANPPPConnection1_GetPPPCompressionProtocol_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetPPPCompressionProtocol_Response
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

// _WANPPPConnection1_GetPPPAuthenticationProtocol_Request is the XML structure for the input arguments for action GetPPPAuthenticationProtocol.
type _WANPPPConnection1_GetPPPAuthenticationProtocol_Request struct{}

// _WANPPPConnection1_GetPPPAuthenticationProtocol_Response is the XML structure for the output arguments for action GetPPPAuthenticationProtocol.
type _WANPPPConnection1_GetPPPAuthenticationProtocol_Response struct {
	NewPPPAuthenticationProtocol string
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
	var request _WANPPPConnection1_GetPPPAuthenticationProtocol_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetPPPAuthenticationProtocol_Response
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

// _WANPPPConnection1_GetUserName_Request is the XML structure for the input arguments for action GetUserName.
type _WANPPPConnection1_GetUserName_Request struct{}

// _WANPPPConnection1_GetUserName_Response is the XML structure for the output arguments for action GetUserName.
type _WANPPPConnection1_GetUserName_Response struct {
	NewUserName string
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
	var request _WANPPPConnection1_GetUserName_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetUserName_Response
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

// _WANPPPConnection1_GetPassword_Request is the XML structure for the input arguments for action GetPassword.
type _WANPPPConnection1_GetPassword_Request struct{}

// _WANPPPConnection1_GetPassword_Response is the XML structure for the output arguments for action GetPassword.
type _WANPPPConnection1_GetPassword_Response struct {
	NewPassword string
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
	var request _WANPPPConnection1_GetPassword_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetPassword_Response
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

// _WANPPPConnection1_GetAutoDisconnectTime_Request is the XML structure for the input arguments for action GetAutoDisconnectTime.
type _WANPPPConnection1_GetAutoDisconnectTime_Request struct{}

// _WANPPPConnection1_GetAutoDisconnectTime_Response is the XML structure for the output arguments for action GetAutoDisconnectTime.
type _WANPPPConnection1_GetAutoDisconnectTime_Response struct {
	NewAutoDisconnectTime string
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
	var request _WANPPPConnection1_GetAutoDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetAutoDisconnectTime_Response
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

// _WANPPPConnection1_GetIdleDisconnectTime_Request is the XML structure for the input arguments for action GetIdleDisconnectTime.
type _WANPPPConnection1_GetIdleDisconnectTime_Request struct{}

// _WANPPPConnection1_GetIdleDisconnectTime_Response is the XML structure for the output arguments for action GetIdleDisconnectTime.
type _WANPPPConnection1_GetIdleDisconnectTime_Response struct {
	NewIdleDisconnectTime string
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
	var request _WANPPPConnection1_GetIdleDisconnectTime_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetIdleDisconnectTime_Response
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

// _WANPPPConnection1_GetWarnDisconnectDelay_Request is the XML structure for the input arguments for action GetWarnDisconnectDelay.
type _WANPPPConnection1_GetWarnDisconnectDelay_Request struct{}

// _WANPPPConnection1_GetWarnDisconnectDelay_Response is the XML structure for the output arguments for action GetWarnDisconnectDelay.
type _WANPPPConnection1_GetWarnDisconnectDelay_Response struct {
	NewWarnDisconnectDelay string
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
	var request _WANPPPConnection1_GetWarnDisconnectDelay_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetWarnDisconnectDelay_Response
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

// _WANPPPConnection1_GetNATRSIPStatus_Request is the XML structure for the input arguments for action GetNATRSIPStatus.
type _WANPPPConnection1_GetNATRSIPStatus_Request struct{}

// _WANPPPConnection1_GetNATRSIPStatus_Response is the XML structure for the output arguments for action GetNATRSIPStatus.
type _WANPPPConnection1_GetNATRSIPStatus_Response struct {
	NewRSIPAvailable string

	NewNATEnabled string
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
	var request _WANPPPConnection1_GetNATRSIPStatus_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetNATRSIPStatus_Response
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

// _WANPPPConnection1_GetGenericPortMappingEntry_Request is the XML structure for the input arguments for action GetGenericPortMappingEntry.
type _WANPPPConnection1_GetGenericPortMappingEntry_Request struct {
	NewPortMappingIndex string
}

// _WANPPPConnection1_GetGenericPortMappingEntry_Response is the XML structure for the output arguments for action GetGenericPortMappingEntry.
type _WANPPPConnection1_GetGenericPortMappingEntry_Response struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
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
	var request _WANPPPConnection1_GetGenericPortMappingEntry_Request
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetGenericPortMappingEntry_Response
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

// _WANPPPConnection1_GetSpecificPortMappingEntry_Request is the XML structure for the input arguments for action GetSpecificPortMappingEntry.
type _WANPPPConnection1_GetSpecificPortMappingEntry_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string
}

// _WANPPPConnection1_GetSpecificPortMappingEntry_Response is the XML structure for the output arguments for action GetSpecificPortMappingEntry.
type _WANPPPConnection1_GetSpecificPortMappingEntry_Response struct {
	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
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
	var request _WANPPPConnection1_GetSpecificPortMappingEntry_Request
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

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetSpecificPortMappingEntry_Response
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

// _WANPPPConnection1_AddPortMapping_Request is the XML structure for the input arguments for action AddPortMapping.
type _WANPPPConnection1_AddPortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string

	NewInternalPort string

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration string
}

// _WANPPPConnection1_AddPortMapping_Response is the XML structure for the output arguments for action AddPortMapping.
type _WANPPPConnection1_AddPortMapping_Response struct{}

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
	var request _WANPPPConnection1_AddPortMapping_Request
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

	// Perform the SOAP call.
	var response _WANPPPConnection1_AddPortMapping_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_DeletePortMapping_Request is the XML structure for the input arguments for action DeletePortMapping.
type _WANPPPConnection1_DeletePortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort string

	NewProtocol string
}

// _WANPPPConnection1_DeletePortMapping_Response is the XML structure for the output arguments for action DeletePortMapping.
type _WANPPPConnection1_DeletePortMapping_Response struct{}

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
	var request _WANPPPConnection1_DeletePortMapping_Request
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

	// Perform the SOAP call.
	var response _WANPPPConnection1_DeletePortMapping_Response
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// _WANPPPConnection1_GetExternalIPAddress_Request is the XML structure for the input arguments for action GetExternalIPAddress.
type _WANPPPConnection1_GetExternalIPAddress_Request struct{}

// _WANPPPConnection1_GetExternalIPAddress_Response is the XML structure for the output arguments for action GetExternalIPAddress.
type _WANPPPConnection1_GetExternalIPAddress_Response struct {
	NewExternalIPAddress string
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
	var request _WANPPPConnection1_GetExternalIPAddress_Request
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Perform the SOAP call.
	var response _WANPPPConnection1_GetExternalIPAddress_Response
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
