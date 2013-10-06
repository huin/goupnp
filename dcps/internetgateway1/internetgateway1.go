package internetgateway1

import "github.com/huin/goupnp/soap"

const (
	URN_LANDevice_1 = "urn:schemas-upnp-org:device:LANDevice:1"

	URN_WANConnectionDevice_1 = "urn:schemas-upnp-org:device:WANConnectionDevice:1"

	URN_WANDevice_1 = "urn:schemas-upnp-org:device:WANDevice:1"
)

const (
	URN_LANHostConfigManagement_1 = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"

	URN_Layer3Forwarding_1 = "urn:schemas-upnp-org:service:Layer3Forwarding:1"

	URN_WANCableLinkConfig_1 = "urn:schemas-upnp-org:service:WANCableLinkConfig:1"

	URN_WANCommonInterfaceConfig_1 = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"

	URN_WANDSLLinkConfig_1 = "urn:schemas-upnp-org:service:WANDSLLinkConfig:1"

	URN_WANEthernetLinkConfig_1 = "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1"

	URN_WANIPConnection_1 = "urn:schemas-upnp-org:service:WANIPConnection:1"

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
type _LANHostConfigManagement1_SetDHCPServerConfigurable_Response struct {}

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
	NewDHCPServerConfigurable string,
) ( err error) {
	request := _LANHostConfigManagement1_SetDHCPServerConfigurable_Request{

	NewDHCPServerConfigurable: NewDHCPServerConfigurable,

	}
	var response _LANHostConfigManagement1_SetDHCPServerConfigurable_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPServerConfigurable", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetDHCPServerConfigurable_Request is the XML structure for the input arguments for action GetDHCPServerConfigurable.
type _LANHostConfigManagement1_GetDHCPServerConfigurable_Request struct {}

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
	NewDHCPServerConfigurable string,
 err error) {
	request := _LANHostConfigManagement1_GetDHCPServerConfigurable_Request{

	}
	var response _LANHostConfigManagement1_GetDHCPServerConfigurable_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPServerConfigurable", &request, &response)
	if err != nil {
		return
	}

	NewDHCPServerConfigurable = response.NewDHCPServerConfigurable

	return
}






// _LANHostConfigManagement1_SetDHCPRelay_Request is the XML structure for the input arguments for action SetDHCPRelay.
type _LANHostConfigManagement1_SetDHCPRelay_Request struct {
	NewDHCPRelay string
}

// _LANHostConfigManagement1_SetDHCPRelay_Response is the XML structure for the output arguments for action SetDHCPRelay.
type _LANHostConfigManagement1_SetDHCPRelay_Response struct {}

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
	NewDHCPRelay string,
) ( err error) {
	request := _LANHostConfigManagement1_SetDHCPRelay_Request{

	NewDHCPRelay: NewDHCPRelay,

	}
	var response _LANHostConfigManagement1_SetDHCPRelay_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPRelay", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetDHCPRelay_Request is the XML structure for the input arguments for action GetDHCPRelay.
type _LANHostConfigManagement1_GetDHCPRelay_Request struct {}

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
	NewDHCPRelay string,
 err error) {
	request := _LANHostConfigManagement1_GetDHCPRelay_Request{

	}
	var response _LANHostConfigManagement1_GetDHCPRelay_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPRelay", &request, &response)
	if err != nil {
		return
	}

	NewDHCPRelay = response.NewDHCPRelay

	return
}






// _LANHostConfigManagement1_SetSubnetMask_Request is the XML structure for the input arguments for action SetSubnetMask.
type _LANHostConfigManagement1_SetSubnetMask_Request struct {
	NewSubnetMask string
}

// _LANHostConfigManagement1_SetSubnetMask_Response is the XML structure for the output arguments for action SetSubnetMask.
type _LANHostConfigManagement1_SetSubnetMask_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_SetSubnetMask_Request{

	NewSubnetMask: NewSubnetMask,

	}
	var response _LANHostConfigManagement1_SetSubnetMask_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetSubnetMask", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetSubnetMask_Request is the XML structure for the input arguments for action GetSubnetMask.
type _LANHostConfigManagement1_GetSubnetMask_Request struct {}

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
 err error) {
	request := _LANHostConfigManagement1_GetSubnetMask_Request{

	}
	var response _LANHostConfigManagement1_GetSubnetMask_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetSubnetMask", &request, &response)
	if err != nil {
		return
	}

	NewSubnetMask = response.NewSubnetMask

	return
}






// _LANHostConfigManagement1_SetIPRouter_Request is the XML structure for the input arguments for action SetIPRouter.
type _LANHostConfigManagement1_SetIPRouter_Request struct {
	NewIPRouters string
}

// _LANHostConfigManagement1_SetIPRouter_Response is the XML structure for the output arguments for action SetIPRouter.
type _LANHostConfigManagement1_SetIPRouter_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_SetIPRouter_Request{

	NewIPRouters: NewIPRouters,

	}
	var response _LANHostConfigManagement1_SetIPRouter_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetIPRouter", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_DeleteIPRouter_Request is the XML structure for the input arguments for action DeleteIPRouter.
type _LANHostConfigManagement1_DeleteIPRouter_Request struct {
	NewIPRouters string
}

// _LANHostConfigManagement1_DeleteIPRouter_Response is the XML structure for the output arguments for action DeleteIPRouter.
type _LANHostConfigManagement1_DeleteIPRouter_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_DeleteIPRouter_Request{

	NewIPRouters: NewIPRouters,

	}
	var response _LANHostConfigManagement1_DeleteIPRouter_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteIPRouter", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetIPRoutersList_Request is the XML structure for the input arguments for action GetIPRoutersList.
type _LANHostConfigManagement1_GetIPRoutersList_Request struct {}

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
 err error) {
	request := _LANHostConfigManagement1_GetIPRoutersList_Request{

	}
	var response _LANHostConfigManagement1_GetIPRoutersList_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetIPRoutersList", &request, &response)
	if err != nil {
		return
	}

	NewIPRouters = response.NewIPRouters

	return
}






// _LANHostConfigManagement1_SetDomainName_Request is the XML structure for the input arguments for action SetDomainName.
type _LANHostConfigManagement1_SetDomainName_Request struct {
	NewDomainName string
}

// _LANHostConfigManagement1_SetDomainName_Response is the XML structure for the output arguments for action SetDomainName.
type _LANHostConfigManagement1_SetDomainName_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_SetDomainName_Request{

	NewDomainName: NewDomainName,

	}
	var response _LANHostConfigManagement1_SetDomainName_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDomainName", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetDomainName_Request is the XML structure for the input arguments for action GetDomainName.
type _LANHostConfigManagement1_GetDomainName_Request struct {}

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
 err error) {
	request := _LANHostConfigManagement1_GetDomainName_Request{

	}
	var response _LANHostConfigManagement1_GetDomainName_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDomainName", &request, &response)
	if err != nil {
		return
	}

	NewDomainName = response.NewDomainName

	return
}






// _LANHostConfigManagement1_SetAddressRange_Request is the XML structure for the input arguments for action SetAddressRange.
type _LANHostConfigManagement1_SetAddressRange_Request struct {
	NewMinAddress string

	NewMaxAddress string
}

// _LANHostConfigManagement1_SetAddressRange_Response is the XML structure for the output arguments for action SetAddressRange.
type _LANHostConfigManagement1_SetAddressRange_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_SetAddressRange_Request{

	NewMinAddress: NewMinAddress,

	NewMaxAddress: NewMaxAddress,

	}
	var response _LANHostConfigManagement1_SetAddressRange_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetAddressRange", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetAddressRange_Request is the XML structure for the input arguments for action GetAddressRange.
type _LANHostConfigManagement1_GetAddressRange_Request struct {}

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
 err error) {
	request := _LANHostConfigManagement1_GetAddressRange_Request{

	}
	var response _LANHostConfigManagement1_GetAddressRange_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetAddressRange", &request, &response)
	if err != nil {
		return
	}

	NewMinAddress = response.NewMinAddress

	NewMaxAddress = response.NewMaxAddress

	return
}






// _LANHostConfigManagement1_SetReservedAddress_Request is the XML structure for the input arguments for action SetReservedAddress.
type _LANHostConfigManagement1_SetReservedAddress_Request struct {
	NewReservedAddresses string
}

// _LANHostConfigManagement1_SetReservedAddress_Response is the XML structure for the output arguments for action SetReservedAddress.
type _LANHostConfigManagement1_SetReservedAddress_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_SetReservedAddress_Request{

	NewReservedAddresses: NewReservedAddresses,

	}
	var response _LANHostConfigManagement1_SetReservedAddress_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetReservedAddress", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_DeleteReservedAddress_Request is the XML structure for the input arguments for action DeleteReservedAddress.
type _LANHostConfigManagement1_DeleteReservedAddress_Request struct {
	NewReservedAddresses string
}

// _LANHostConfigManagement1_DeleteReservedAddress_Response is the XML structure for the output arguments for action DeleteReservedAddress.
type _LANHostConfigManagement1_DeleteReservedAddress_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_DeleteReservedAddress_Request{

	NewReservedAddresses: NewReservedAddresses,

	}
	var response _LANHostConfigManagement1_DeleteReservedAddress_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteReservedAddress", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetReservedAddresses_Request is the XML structure for the input arguments for action GetReservedAddresses.
type _LANHostConfigManagement1_GetReservedAddresses_Request struct {}

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
 err error) {
	request := _LANHostConfigManagement1_GetReservedAddresses_Request{

	}
	var response _LANHostConfigManagement1_GetReservedAddresses_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetReservedAddresses", &request, &response)
	if err != nil {
		return
	}

	NewReservedAddresses = response.NewReservedAddresses

	return
}






// _LANHostConfigManagement1_SetDNSServer_Request is the XML structure for the input arguments for action SetDNSServer.
type _LANHostConfigManagement1_SetDNSServer_Request struct {
	NewDNSServers string
}

// _LANHostConfigManagement1_SetDNSServer_Response is the XML structure for the output arguments for action SetDNSServer.
type _LANHostConfigManagement1_SetDNSServer_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_SetDNSServer_Request{

	NewDNSServers: NewDNSServers,

	}
	var response _LANHostConfigManagement1_SetDNSServer_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDNSServer", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_DeleteDNSServer_Request is the XML structure for the input arguments for action DeleteDNSServer.
type _LANHostConfigManagement1_DeleteDNSServer_Request struct {
	NewDNSServers string
}

// _LANHostConfigManagement1_DeleteDNSServer_Response is the XML structure for the output arguments for action DeleteDNSServer.
type _LANHostConfigManagement1_DeleteDNSServer_Response struct {}

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
) ( err error) {
	request := _LANHostConfigManagement1_DeleteDNSServer_Request{

	NewDNSServers: NewDNSServers,

	}
	var response _LANHostConfigManagement1_DeleteDNSServer_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteDNSServer", &request, &response)
	if err != nil {
		return
	}

	return
}






// _LANHostConfigManagement1_GetDNSServers_Request is the XML structure for the input arguments for action GetDNSServers.
type _LANHostConfigManagement1_GetDNSServers_Request struct {}

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
 err error) {
	request := _LANHostConfigManagement1_GetDNSServers_Request{

	}
	var response _LANHostConfigManagement1_GetDNSServers_Response
	err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDNSServers", &request, &response)
	if err != nil {
		return
	}

	NewDNSServers = response.NewDNSServers

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
type _Layer3Forwarding1_SetDefaultConnectionService_Response struct {}

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
) ( err error) {
	request := _Layer3Forwarding1_SetDefaultConnectionService_Request{

	NewDefaultConnectionService: NewDefaultConnectionService,

	}
	var response _Layer3Forwarding1_SetDefaultConnectionService_Response
	err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "SetDefaultConnectionService", &request, &response)
	if err != nil {
		return
	}

	return
}






// _Layer3Forwarding1_GetDefaultConnectionService_Request is the XML structure for the input arguments for action GetDefaultConnectionService.
type _Layer3Forwarding1_GetDefaultConnectionService_Request struct {}

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
 err error) {
	request := _Layer3Forwarding1_GetDefaultConnectionService_Request{

	}
	var response _Layer3Forwarding1_GetDefaultConnectionService_Response
	err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "GetDefaultConnectionService", &request, &response)
	if err != nil {
		return
	}

	NewDefaultConnectionService = response.NewDefaultConnectionService

	return
}






// WANCableLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCableLinkConfig:1".
type WANCableLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}






// _WANCableLinkConfig1_GetCableLinkConfigInfo_Request is the XML structure for the input arguments for action GetCableLinkConfigInfo.
type _WANCableLinkConfig1_GetCableLinkConfigInfo_Request struct {}

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
 err error) {
	request := _WANCableLinkConfig1_GetCableLinkConfigInfo_Request{

	}
	var response _WANCableLinkConfig1_GetCableLinkConfigInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetCableLinkConfigInfo", &request, &response)
	if err != nil {
		return
	}

	NewCableLinkConfigState = response.NewCableLinkConfigState

	NewLinkType = response.NewLinkType

	return
}






// _WANCableLinkConfig1_GetDownstreamFrequency_Request is the XML structure for the input arguments for action GetDownstreamFrequency.
type _WANCableLinkConfig1_GetDownstreamFrequency_Request struct {}

// _WANCableLinkConfig1_GetDownstreamFrequency_Response is the XML structure for the output arguments for action GetDownstreamFrequency.
type _WANCableLinkConfig1_GetDownstreamFrequency_Response struct {
	NewDownstreamFrequency uint32
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
 err error) {
	request := _WANCableLinkConfig1_GetDownstreamFrequency_Request{

	}
	var response _WANCableLinkConfig1_GetDownstreamFrequency_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamFrequency", &request, &response)
	if err != nil {
		return
	}

	NewDownstreamFrequency = response.NewDownstreamFrequency

	return
}






// _WANCableLinkConfig1_GetDownstreamModulation_Request is the XML structure for the input arguments for action GetDownstreamModulation.
type _WANCableLinkConfig1_GetDownstreamModulation_Request struct {}

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
 err error) {
	request := _WANCableLinkConfig1_GetDownstreamModulation_Request{

	}
	var response _WANCableLinkConfig1_GetDownstreamModulation_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamModulation", &request, &response)
	if err != nil {
		return
	}

	NewDownstreamModulation = response.NewDownstreamModulation

	return
}






// _WANCableLinkConfig1_GetUpstreamFrequency_Request is the XML structure for the input arguments for action GetUpstreamFrequency.
type _WANCableLinkConfig1_GetUpstreamFrequency_Request struct {}

// _WANCableLinkConfig1_GetUpstreamFrequency_Response is the XML structure for the output arguments for action GetUpstreamFrequency.
type _WANCableLinkConfig1_GetUpstreamFrequency_Response struct {
	NewUpstreamFrequency uint32
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
 err error) {
	request := _WANCableLinkConfig1_GetUpstreamFrequency_Request{

	}
	var response _WANCableLinkConfig1_GetUpstreamFrequency_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamFrequency", &request, &response)
	if err != nil {
		return
	}

	NewUpstreamFrequency = response.NewUpstreamFrequency

	return
}






// _WANCableLinkConfig1_GetUpstreamModulation_Request is the XML structure for the input arguments for action GetUpstreamModulation.
type _WANCableLinkConfig1_GetUpstreamModulation_Request struct {}

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
 err error) {
	request := _WANCableLinkConfig1_GetUpstreamModulation_Request{

	}
	var response _WANCableLinkConfig1_GetUpstreamModulation_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamModulation", &request, &response)
	if err != nil {
		return
	}

	NewUpstreamModulation = response.NewUpstreamModulation

	return
}






// _WANCableLinkConfig1_GetUpstreamChannelID_Request is the XML structure for the input arguments for action GetUpstreamChannelID.
type _WANCableLinkConfig1_GetUpstreamChannelID_Request struct {}

// _WANCableLinkConfig1_GetUpstreamChannelID_Response is the XML structure for the output arguments for action GetUpstreamChannelID.
type _WANCableLinkConfig1_GetUpstreamChannelID_Response struct {
	NewUpstreamChannelID uint32
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
 err error) {
	request := _WANCableLinkConfig1_GetUpstreamChannelID_Request{

	}
	var response _WANCableLinkConfig1_GetUpstreamChannelID_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamChannelID", &request, &response)
	if err != nil {
		return
	}

	NewUpstreamChannelID = response.NewUpstreamChannelID

	return
}






// _WANCableLinkConfig1_GetUpstreamPowerLevel_Request is the XML structure for the input arguments for action GetUpstreamPowerLevel.
type _WANCableLinkConfig1_GetUpstreamPowerLevel_Request struct {}

// _WANCableLinkConfig1_GetUpstreamPowerLevel_Response is the XML structure for the output arguments for action GetUpstreamPowerLevel.
type _WANCableLinkConfig1_GetUpstreamPowerLevel_Response struct {
	NewUpstreamPowerLevel uint32
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
 err error) {
	request := _WANCableLinkConfig1_GetUpstreamPowerLevel_Request{

	}
	var response _WANCableLinkConfig1_GetUpstreamPowerLevel_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamPowerLevel", &request, &response)
	if err != nil {
		return
	}

	NewUpstreamPowerLevel = response.NewUpstreamPowerLevel

	return
}






// _WANCableLinkConfig1_GetBPIEncryptionEnabled_Request is the XML structure for the input arguments for action GetBPIEncryptionEnabled.
type _WANCableLinkConfig1_GetBPIEncryptionEnabled_Request struct {}

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
	NewBPIEncryptionEnabled string,
 err error) {
	request := _WANCableLinkConfig1_GetBPIEncryptionEnabled_Request{

	}
	var response _WANCableLinkConfig1_GetBPIEncryptionEnabled_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetBPIEncryptionEnabled", &request, &response)
	if err != nil {
		return
	}

	NewBPIEncryptionEnabled = response.NewBPIEncryptionEnabled

	return
}






// _WANCableLinkConfig1_GetConfigFile_Request is the XML structure for the input arguments for action GetConfigFile.
type _WANCableLinkConfig1_GetConfigFile_Request struct {}

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
 err error) {
	request := _WANCableLinkConfig1_GetConfigFile_Request{

	}
	var response _WANCableLinkConfig1_GetConfigFile_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetConfigFile", &request, &response)
	if err != nil {
		return
	}

	NewConfigFile = response.NewConfigFile

	return
}






// _WANCableLinkConfig1_GetTFTPServer_Request is the XML structure for the input arguments for action GetTFTPServer.
type _WANCableLinkConfig1_GetTFTPServer_Request struct {}

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
 err error) {
	request := _WANCableLinkConfig1_GetTFTPServer_Request{

	}
	var response _WANCableLinkConfig1_GetTFTPServer_Response
	err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetTFTPServer", &request, &response)
	if err != nil {
		return
	}

	NewTFTPServer = response.NewTFTPServer

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
type _WANCommonInterfaceConfig1_SetEnabledForInternet_Response struct {}

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
	NewEnabledForInternet string,
) ( err error) {
	request := _WANCommonInterfaceConfig1_SetEnabledForInternet_Request{

	NewEnabledForInternet: NewEnabledForInternet,

	}
	var response _WANCommonInterfaceConfig1_SetEnabledForInternet_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "SetEnabledForInternet", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANCommonInterfaceConfig1_GetEnabledForInternet_Request is the XML structure for the input arguments for action GetEnabledForInternet.
type _WANCommonInterfaceConfig1_GetEnabledForInternet_Request struct {}

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
	NewEnabledForInternet string,
 err error) {
	request := _WANCommonInterfaceConfig1_GetEnabledForInternet_Request{

	}
	var response _WANCommonInterfaceConfig1_GetEnabledForInternet_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetEnabledForInternet", &request, &response)
	if err != nil {
		return
	}

	NewEnabledForInternet = response.NewEnabledForInternet

	return
}






// _WANCommonInterfaceConfig1_GetCommonLinkProperties_Request is the XML structure for the input arguments for action GetCommonLinkProperties.
type _WANCommonInterfaceConfig1_GetCommonLinkProperties_Request struct {}

// _WANCommonInterfaceConfig1_GetCommonLinkProperties_Response is the XML structure for the output arguments for action GetCommonLinkProperties.
type _WANCommonInterfaceConfig1_GetCommonLinkProperties_Response struct {
	NewWANAccessType string

	NewLayer1UpstreamMaxBitRate uint32

	NewLayer1DownstreamMaxBitRate uint32

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
 err error) {
	request := _WANCommonInterfaceConfig1_GetCommonLinkProperties_Request{

	}
	var response _WANCommonInterfaceConfig1_GetCommonLinkProperties_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetCommonLinkProperties", &request, &response)
	if err != nil {
		return
	}

	NewWANAccessType = response.NewWANAccessType

	NewLayer1UpstreamMaxBitRate = response.NewLayer1UpstreamMaxBitRate

	NewLayer1DownstreamMaxBitRate = response.NewLayer1DownstreamMaxBitRate

	NewPhysicalLinkStatus = response.NewPhysicalLinkStatus

	return
}






// _WANCommonInterfaceConfig1_GetWANAccessProvider_Request is the XML structure for the input arguments for action GetWANAccessProvider.
type _WANCommonInterfaceConfig1_GetWANAccessProvider_Request struct {}

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
 err error) {
	request := _WANCommonInterfaceConfig1_GetWANAccessProvider_Request{

	}
	var response _WANCommonInterfaceConfig1_GetWANAccessProvider_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetWANAccessProvider", &request, &response)
	if err != nil {
		return
	}

	NewWANAccessProvider = response.NewWANAccessProvider

	return
}






// _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Request is the XML structure for the input arguments for action GetMaximumActiveConnections.
type _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Request struct {}

// _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Response is the XML structure for the output arguments for action GetMaximumActiveConnections.
type _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Response struct {
	NewMaximumActiveConnections uint16
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
 err error) {
	request := _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Request{

	}
	var response _WANCommonInterfaceConfig1_GetMaximumActiveConnections_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetMaximumActiveConnections", &request, &response)
	if err != nil {
		return
	}

	NewMaximumActiveConnections = response.NewMaximumActiveConnections

	return
}






// _WANCommonInterfaceConfig1_GetTotalBytesSent_Request is the XML structure for the input arguments for action GetTotalBytesSent.
type _WANCommonInterfaceConfig1_GetTotalBytesSent_Request struct {}

// _WANCommonInterfaceConfig1_GetTotalBytesSent_Response is the XML structure for the output arguments for action GetTotalBytesSent.
type _WANCommonInterfaceConfig1_GetTotalBytesSent_Response struct {
	NewTotalBytesSent uint32
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
 err error) {
	request := _WANCommonInterfaceConfig1_GetTotalBytesSent_Request{

	}
	var response _WANCommonInterfaceConfig1_GetTotalBytesSent_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesSent", &request, &response)
	if err != nil {
		return
	}

	NewTotalBytesSent = response.NewTotalBytesSent

	return
}






// _WANCommonInterfaceConfig1_GetTotalBytesReceived_Request is the XML structure for the input arguments for action GetTotalBytesReceived.
type _WANCommonInterfaceConfig1_GetTotalBytesReceived_Request struct {}

// _WANCommonInterfaceConfig1_GetTotalBytesReceived_Response is the XML structure for the output arguments for action GetTotalBytesReceived.
type _WANCommonInterfaceConfig1_GetTotalBytesReceived_Response struct {
	NewTotalBytesReceived uint32
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
 err error) {
	request := _WANCommonInterfaceConfig1_GetTotalBytesReceived_Request{

	}
	var response _WANCommonInterfaceConfig1_GetTotalBytesReceived_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesReceived", &request, &response)
	if err != nil {
		return
	}

	NewTotalBytesReceived = response.NewTotalBytesReceived

	return
}






// _WANCommonInterfaceConfig1_GetTotalPacketsSent_Request is the XML structure for the input arguments for action GetTotalPacketsSent.
type _WANCommonInterfaceConfig1_GetTotalPacketsSent_Request struct {}

// _WANCommonInterfaceConfig1_GetTotalPacketsSent_Response is the XML structure for the output arguments for action GetTotalPacketsSent.
type _WANCommonInterfaceConfig1_GetTotalPacketsSent_Response struct {
	NewTotalPacketsSent uint32
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
 err error) {
	request := _WANCommonInterfaceConfig1_GetTotalPacketsSent_Request{

	}
	var response _WANCommonInterfaceConfig1_GetTotalPacketsSent_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsSent", &request, &response)
	if err != nil {
		return
	}

	NewTotalPacketsSent = response.NewTotalPacketsSent

	return
}






// _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Request is the XML structure for the input arguments for action GetTotalPacketsReceived.
type _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Request struct {}

// _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Response is the XML structure for the output arguments for action GetTotalPacketsReceived.
type _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Response struct {
	NewTotalPacketsReceived uint32
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
 err error) {
	request := _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Request{

	}
	var response _WANCommonInterfaceConfig1_GetTotalPacketsReceived_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsReceived", &request, &response)
	if err != nil {
		return
	}

	NewTotalPacketsReceived = response.NewTotalPacketsReceived

	return
}






// _WANCommonInterfaceConfig1_GetActiveConnection_Request is the XML structure for the input arguments for action GetActiveConnection.
type _WANCommonInterfaceConfig1_GetActiveConnection_Request struct {
	NewActiveConnectionIndex uint16
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
 err error) {
	request := _WANCommonInterfaceConfig1_GetActiveConnection_Request{

	NewActiveConnectionIndex: NewActiveConnectionIndex,

	}
	var response _WANCommonInterfaceConfig1_GetActiveConnection_Response
	err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetActiveConnection", &request, &response)
	if err != nil {
		return
	}

	NewActiveConnDeviceContainer = response.NewActiveConnDeviceContainer

	NewActiveConnectionServiceID = response.NewActiveConnectionServiceID

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
type _WANDSLLinkConfig1_SetDSLLinkType_Response struct {}

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
) ( err error) {
	request := _WANDSLLinkConfig1_SetDSLLinkType_Request{

	NewLinkType: NewLinkType,

	}
	var response _WANDSLLinkConfig1_SetDSLLinkType_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDSLLinkType", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANDSLLinkConfig1_GetDSLLinkInfo_Request is the XML structure for the input arguments for action GetDSLLinkInfo.
type _WANDSLLinkConfig1_GetDSLLinkInfo_Request struct {}

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
 err error) {
	request := _WANDSLLinkConfig1_GetDSLLinkInfo_Request{

	}
	var response _WANDSLLinkConfig1_GetDSLLinkInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDSLLinkInfo", &request, &response)
	if err != nil {
		return
	}

	NewLinkType = response.NewLinkType

	NewLinkStatus = response.NewLinkStatus

	return
}






// _WANDSLLinkConfig1_GetAutoConfig_Request is the XML structure for the input arguments for action GetAutoConfig.
type _WANDSLLinkConfig1_GetAutoConfig_Request struct {}

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
	NewAutoConfig string,
 err error) {
	request := _WANDSLLinkConfig1_GetAutoConfig_Request{

	}
	var response _WANDSLLinkConfig1_GetAutoConfig_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetAutoConfig", &request, &response)
	if err != nil {
		return
	}

	NewAutoConfig = response.NewAutoConfig

	return
}






// _WANDSLLinkConfig1_GetModulationType_Request is the XML structure for the input arguments for action GetModulationType.
type _WANDSLLinkConfig1_GetModulationType_Request struct {}

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
 err error) {
	request := _WANDSLLinkConfig1_GetModulationType_Request{

	}
	var response _WANDSLLinkConfig1_GetModulationType_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetModulationType", &request, &response)
	if err != nil {
		return
	}

	NewModulationType = response.NewModulationType

	return
}






// _WANDSLLinkConfig1_SetDestinationAddress_Request is the XML structure for the input arguments for action SetDestinationAddress.
type _WANDSLLinkConfig1_SetDestinationAddress_Request struct {
	NewDestinationAddress string
}

// _WANDSLLinkConfig1_SetDestinationAddress_Response is the XML structure for the output arguments for action SetDestinationAddress.
type _WANDSLLinkConfig1_SetDestinationAddress_Response struct {}

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
) ( err error) {
	request := _WANDSLLinkConfig1_SetDestinationAddress_Request{

	NewDestinationAddress: NewDestinationAddress,

	}
	var response _WANDSLLinkConfig1_SetDestinationAddress_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDestinationAddress", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANDSLLinkConfig1_GetDestinationAddress_Request is the XML structure for the input arguments for action GetDestinationAddress.
type _WANDSLLinkConfig1_GetDestinationAddress_Request struct {}

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
 err error) {
	request := _WANDSLLinkConfig1_GetDestinationAddress_Request{

	}
	var response _WANDSLLinkConfig1_GetDestinationAddress_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDestinationAddress", &request, &response)
	if err != nil {
		return
	}

	NewDestinationAddress = response.NewDestinationAddress

	return
}






// _WANDSLLinkConfig1_SetATMEncapsulation_Request is the XML structure for the input arguments for action SetATMEncapsulation.
type _WANDSLLinkConfig1_SetATMEncapsulation_Request struct {
	NewATMEncapsulation string
}

// _WANDSLLinkConfig1_SetATMEncapsulation_Response is the XML structure for the output arguments for action SetATMEncapsulation.
type _WANDSLLinkConfig1_SetATMEncapsulation_Response struct {}

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
) ( err error) {
	request := _WANDSLLinkConfig1_SetATMEncapsulation_Request{

	NewATMEncapsulation: NewATMEncapsulation,

	}
	var response _WANDSLLinkConfig1_SetATMEncapsulation_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetATMEncapsulation", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANDSLLinkConfig1_GetATMEncapsulation_Request is the XML structure for the input arguments for action GetATMEncapsulation.
type _WANDSLLinkConfig1_GetATMEncapsulation_Request struct {}

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
 err error) {
	request := _WANDSLLinkConfig1_GetATMEncapsulation_Request{

	}
	var response _WANDSLLinkConfig1_GetATMEncapsulation_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetATMEncapsulation", &request, &response)
	if err != nil {
		return
	}

	NewATMEncapsulation = response.NewATMEncapsulation

	return
}






// _WANDSLLinkConfig1_SetFCSPreserved_Request is the XML structure for the input arguments for action SetFCSPreserved.
type _WANDSLLinkConfig1_SetFCSPreserved_Request struct {
	NewFCSPreserved string
}

// _WANDSLLinkConfig1_SetFCSPreserved_Response is the XML structure for the output arguments for action SetFCSPreserved.
type _WANDSLLinkConfig1_SetFCSPreserved_Response struct {}

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
	NewFCSPreserved string,
) ( err error) {
	request := _WANDSLLinkConfig1_SetFCSPreserved_Request{

	NewFCSPreserved: NewFCSPreserved,

	}
	var response _WANDSLLinkConfig1_SetFCSPreserved_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetFCSPreserved", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANDSLLinkConfig1_GetFCSPreserved_Request is the XML structure for the input arguments for action GetFCSPreserved.
type _WANDSLLinkConfig1_GetFCSPreserved_Request struct {}

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
	NewFCSPreserved string,
 err error) {
	request := _WANDSLLinkConfig1_GetFCSPreserved_Request{

	}
	var response _WANDSLLinkConfig1_GetFCSPreserved_Response
	err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetFCSPreserved", &request, &response)
	if err != nil {
		return
	}

	NewFCSPreserved = response.NewFCSPreserved

	return
}






// WANEthernetLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1".
type WANEthernetLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}






// _WANEthernetLinkConfig1_GetEthernetLinkStatus_Request is the XML structure for the input arguments for action GetEthernetLinkStatus.
type _WANEthernetLinkConfig1_GetEthernetLinkStatus_Request struct {}

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
 err error) {
	request := _WANEthernetLinkConfig1_GetEthernetLinkStatus_Request{

	}
	var response _WANEthernetLinkConfig1_GetEthernetLinkStatus_Response
	err = client.SOAPClient.PerformAction(URN_WANEthernetLinkConfig_1, "GetEthernetLinkStatus", &request, &response)
	if err != nil {
		return
	}

	NewEthernetLinkStatus = response.NewEthernetLinkStatus

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
type _WANIPConnection1_SetConnectionType_Response struct {}

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
) ( err error) {
	request := _WANIPConnection1_SetConnectionType_Request{

	NewConnectionType: NewConnectionType,

	}
	var response _WANIPConnection1_SetConnectionType_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetConnectionType", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_GetConnectionTypeInfo_Request is the XML structure for the input arguments for action GetConnectionTypeInfo.
type _WANIPConnection1_GetConnectionTypeInfo_Request struct {}

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
 err error) {
	request := _WANIPConnection1_GetConnectionTypeInfo_Request{

	}
	var response _WANIPConnection1_GetConnectionTypeInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetConnectionTypeInfo", &request, &response)
	if err != nil {
		return
	}

	NewConnectionType = response.NewConnectionType

	NewPossibleConnectionTypes = response.NewPossibleConnectionTypes

	return
}






// _WANIPConnection1_RequestConnection_Request is the XML structure for the input arguments for action RequestConnection.
type _WANIPConnection1_RequestConnection_Request struct {}

// _WANIPConnection1_RequestConnection_Response is the XML structure for the output arguments for action RequestConnection.
type _WANIPConnection1_RequestConnection_Response struct {}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) RequestConnection() ( err error) {
	request := _WANIPConnection1_RequestConnection_Request{

	}
	var response _WANIPConnection1_RequestConnection_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestConnection", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_RequestTermination_Request is the XML structure for the input arguments for action RequestTermination.
type _WANIPConnection1_RequestTermination_Request struct {}

// _WANIPConnection1_RequestTermination_Response is the XML structure for the output arguments for action RequestTermination.
type _WANIPConnection1_RequestTermination_Response struct {}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) RequestTermination() ( err error) {
	request := _WANIPConnection1_RequestTermination_Request{

	}
	var response _WANIPConnection1_RequestTermination_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestTermination", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_ForceTermination_Request is the XML structure for the input arguments for action ForceTermination.
type _WANIPConnection1_ForceTermination_Request struct {}

// _WANIPConnection1_ForceTermination_Response is the XML structure for the output arguments for action ForceTermination.
type _WANIPConnection1_ForceTermination_Response struct {}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANIPConnection1) ForceTermination() ( err error) {
	request := _WANIPConnection1_ForceTermination_Request{

	}
	var response _WANIPConnection1_ForceTermination_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "ForceTermination", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_SetAutoDisconnectTime_Request is the XML structure for the input arguments for action SetAutoDisconnectTime.
type _WANIPConnection1_SetAutoDisconnectTime_Request struct {
	NewAutoDisconnectTime uint32
}

// _WANIPConnection1_SetAutoDisconnectTime_Response is the XML structure for the output arguments for action SetAutoDisconnectTime.
type _WANIPConnection1_SetAutoDisconnectTime_Response struct {}

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
) ( err error) {
	request := _WANIPConnection1_SetAutoDisconnectTime_Request{

	NewAutoDisconnectTime: NewAutoDisconnectTime,

	}
	var response _WANIPConnection1_SetAutoDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetAutoDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_SetIdleDisconnectTime_Request is the XML structure for the input arguments for action SetIdleDisconnectTime.
type _WANIPConnection1_SetIdleDisconnectTime_Request struct {
	NewIdleDisconnectTime uint32
}

// _WANIPConnection1_SetIdleDisconnectTime_Response is the XML structure for the output arguments for action SetIdleDisconnectTime.
type _WANIPConnection1_SetIdleDisconnectTime_Response struct {}

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
) ( err error) {
	request := _WANIPConnection1_SetIdleDisconnectTime_Request{

	NewIdleDisconnectTime: NewIdleDisconnectTime,

	}
	var response _WANIPConnection1_SetIdleDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetIdleDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_SetWarnDisconnectDelay_Request is the XML structure for the input arguments for action SetWarnDisconnectDelay.
type _WANIPConnection1_SetWarnDisconnectDelay_Request struct {
	NewWarnDisconnectDelay uint32
}

// _WANIPConnection1_SetWarnDisconnectDelay_Response is the XML structure for the output arguments for action SetWarnDisconnectDelay.
type _WANIPConnection1_SetWarnDisconnectDelay_Response struct {}

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
) ( err error) {
	request := _WANIPConnection1_SetWarnDisconnectDelay_Request{

	NewWarnDisconnectDelay: NewWarnDisconnectDelay,

	}
	var response _WANIPConnection1_SetWarnDisconnectDelay_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetWarnDisconnectDelay", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_GetStatusInfo_Request is the XML structure for the input arguments for action GetStatusInfo.
type _WANIPConnection1_GetStatusInfo_Request struct {}

// _WANIPConnection1_GetStatusInfo_Response is the XML structure for the output arguments for action GetStatusInfo.
type _WANIPConnection1_GetStatusInfo_Response struct {
	NewConnectionStatus string

	NewLastConnectionError string

	NewUptime uint32
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
 err error) {
	request := _WANIPConnection1_GetStatusInfo_Request{

	}
	var response _WANIPConnection1_GetStatusInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetStatusInfo", &request, &response)
	if err != nil {
		return
	}

	NewConnectionStatus = response.NewConnectionStatus

	NewLastConnectionError = response.NewLastConnectionError

	NewUptime = response.NewUptime

	return
}






// _WANIPConnection1_GetAutoDisconnectTime_Request is the XML structure for the input arguments for action GetAutoDisconnectTime.
type _WANIPConnection1_GetAutoDisconnectTime_Request struct {}

// _WANIPConnection1_GetAutoDisconnectTime_Response is the XML structure for the output arguments for action GetAutoDisconnectTime.
type _WANIPConnection1_GetAutoDisconnectTime_Response struct {
	NewAutoDisconnectTime uint32
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
 err error) {
	request := _WANIPConnection1_GetAutoDisconnectTime_Request{

	}
	var response _WANIPConnection1_GetAutoDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetAutoDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	NewAutoDisconnectTime = response.NewAutoDisconnectTime

	return
}






// _WANIPConnection1_GetIdleDisconnectTime_Request is the XML structure for the input arguments for action GetIdleDisconnectTime.
type _WANIPConnection1_GetIdleDisconnectTime_Request struct {}

// _WANIPConnection1_GetIdleDisconnectTime_Response is the XML structure for the output arguments for action GetIdleDisconnectTime.
type _WANIPConnection1_GetIdleDisconnectTime_Response struct {
	NewIdleDisconnectTime uint32
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
 err error) {
	request := _WANIPConnection1_GetIdleDisconnectTime_Request{

	}
	var response _WANIPConnection1_GetIdleDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetIdleDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	NewIdleDisconnectTime = response.NewIdleDisconnectTime

	return
}






// _WANIPConnection1_GetWarnDisconnectDelay_Request is the XML structure for the input arguments for action GetWarnDisconnectDelay.
type _WANIPConnection1_GetWarnDisconnectDelay_Request struct {}

// _WANIPConnection1_GetWarnDisconnectDelay_Response is the XML structure for the output arguments for action GetWarnDisconnectDelay.
type _WANIPConnection1_GetWarnDisconnectDelay_Response struct {
	NewWarnDisconnectDelay uint32
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
 err error) {
	request := _WANIPConnection1_GetWarnDisconnectDelay_Request{

	}
	var response _WANIPConnection1_GetWarnDisconnectDelay_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetWarnDisconnectDelay", &request, &response)
	if err != nil {
		return
	}

	NewWarnDisconnectDelay = response.NewWarnDisconnectDelay

	return
}






// _WANIPConnection1_GetNATRSIPStatus_Request is the XML structure for the input arguments for action GetNATRSIPStatus.
type _WANIPConnection1_GetNATRSIPStatus_Request struct {}

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
	NewRSIPAvailable string,

	NewNATEnabled string,
 err error) {
	request := _WANIPConnection1_GetNATRSIPStatus_Request{

	}
	var response _WANIPConnection1_GetNATRSIPStatus_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetNATRSIPStatus", &request, &response)
	if err != nil {
		return
	}

	NewRSIPAvailable = response.NewRSIPAvailable

	NewNATEnabled = response.NewNATEnabled

	return
}






// _WANIPConnection1_GetGenericPortMappingEntry_Request is the XML structure for the input arguments for action GetGenericPortMappingEntry.
type _WANIPConnection1_GetGenericPortMappingEntry_Request struct {
	NewPortMappingIndex uint16
}

// _WANIPConnection1_GetGenericPortMappingEntry_Response is the XML structure for the output arguments for action GetGenericPortMappingEntry.
type _WANIPConnection1_GetGenericPortMappingEntry_Response struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string

	NewInternalPort uint16

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration uint32
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

	NewEnabled string,

	NewPortMappingDescription string,

	NewLeaseDuration uint32,
 err error) {
	request := _WANIPConnection1_GetGenericPortMappingEntry_Request{

	NewPortMappingIndex: NewPortMappingIndex,

	}
	var response _WANIPConnection1_GetGenericPortMappingEntry_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetGenericPortMappingEntry", &request, &response)
	if err != nil {
		return
	}

	NewRemoteHost = response.NewRemoteHost

	NewExternalPort = response.NewExternalPort

	NewProtocol = response.NewProtocol

	NewInternalPort = response.NewInternalPort

	NewInternalClient = response.NewInternalClient

	NewEnabled = response.NewEnabled

	NewPortMappingDescription = response.NewPortMappingDescription

	NewLeaseDuration = response.NewLeaseDuration

	return
}






// _WANIPConnection1_GetSpecificPortMappingEntry_Request is the XML structure for the input arguments for action GetSpecificPortMappingEntry.
type _WANIPConnection1_GetSpecificPortMappingEntry_Request struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string
}

// _WANIPConnection1_GetSpecificPortMappingEntry_Response is the XML structure for the output arguments for action GetSpecificPortMappingEntry.
type _WANIPConnection1_GetSpecificPortMappingEntry_Response struct {
	NewInternalPort uint16

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration uint32
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

	NewEnabled string,

	NewPortMappingDescription string,

	NewLeaseDuration uint32,
 err error) {
	request := _WANIPConnection1_GetSpecificPortMappingEntry_Request{

	NewRemoteHost: NewRemoteHost,

	NewExternalPort: NewExternalPort,

	NewProtocol: NewProtocol,

	}
	var response _WANIPConnection1_GetSpecificPortMappingEntry_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetSpecificPortMappingEntry", &request, &response)
	if err != nil {
		return
	}

	NewInternalPort = response.NewInternalPort

	NewInternalClient = response.NewInternalClient

	NewEnabled = response.NewEnabled

	NewPortMappingDescription = response.NewPortMappingDescription

	NewLeaseDuration = response.NewLeaseDuration

	return
}






// _WANIPConnection1_AddPortMapping_Request is the XML structure for the input arguments for action AddPortMapping.
type _WANIPConnection1_AddPortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string

	NewInternalPort uint16

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration uint32
}

// _WANIPConnection1_AddPortMapping_Response is the XML structure for the output arguments for action AddPortMapping.
type _WANIPConnection1_AddPortMapping_Response struct {}

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

	NewEnabled string,

	NewPortMappingDescription string,

	NewLeaseDuration uint32,
) ( err error) {
	request := _WANIPConnection1_AddPortMapping_Request{

	NewRemoteHost: NewRemoteHost,

	NewExternalPort: NewExternalPort,

	NewProtocol: NewProtocol,

	NewInternalPort: NewInternalPort,

	NewInternalClient: NewInternalClient,

	NewEnabled: NewEnabled,

	NewPortMappingDescription: NewPortMappingDescription,

	NewLeaseDuration: NewLeaseDuration,

	}
	var response _WANIPConnection1_AddPortMapping_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "AddPortMapping", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_DeletePortMapping_Request is the XML structure for the input arguments for action DeletePortMapping.
type _WANIPConnection1_DeletePortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string
}

// _WANIPConnection1_DeletePortMapping_Response is the XML structure for the output arguments for action DeletePortMapping.
type _WANIPConnection1_DeletePortMapping_Response struct {}

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
) ( err error) {
	request := _WANIPConnection1_DeletePortMapping_Request{

	NewRemoteHost: NewRemoteHost,

	NewExternalPort: NewExternalPort,

	NewProtocol: NewProtocol,

	}
	var response _WANIPConnection1_DeletePortMapping_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "DeletePortMapping", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANIPConnection1_GetExternalIPAddress_Request is the XML structure for the input arguments for action GetExternalIPAddress.
type _WANIPConnection1_GetExternalIPAddress_Request struct {}

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
 err error) {
	request := _WANIPConnection1_GetExternalIPAddress_Request{

	}
	var response _WANIPConnection1_GetExternalIPAddress_Response
	err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetExternalIPAddress", &request, &response)
	if err != nil {
		return
	}

	NewExternalIPAddress = response.NewExternalIPAddress

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
type _WANPOTSLinkConfig1_SetISPInfo_Response struct {}

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
) ( err error) {
	request := _WANPOTSLinkConfig1_SetISPInfo_Request{

	NewISPPhoneNumber: NewISPPhoneNumber,

	NewISPInfo: NewISPInfo,

	NewLinkType: NewLinkType,

	}
	var response _WANPOTSLinkConfig1_SetISPInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetISPInfo", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPOTSLinkConfig1_SetCallRetryInfo_Request is the XML structure for the input arguments for action SetCallRetryInfo.
type _WANPOTSLinkConfig1_SetCallRetryInfo_Request struct {
	NewNumberOfRetries uint32

	NewDelayBetweenRetries uint32
}

// _WANPOTSLinkConfig1_SetCallRetryInfo_Response is the XML structure for the output arguments for action SetCallRetryInfo.
type _WANPOTSLinkConfig1_SetCallRetryInfo_Response struct {}

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
) ( err error) {
	request := _WANPOTSLinkConfig1_SetCallRetryInfo_Request{

	NewNumberOfRetries: NewNumberOfRetries,

	NewDelayBetweenRetries: NewDelayBetweenRetries,

	}
	var response _WANPOTSLinkConfig1_SetCallRetryInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetCallRetryInfo", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPOTSLinkConfig1_GetISPInfo_Request is the XML structure for the input arguments for action GetISPInfo.
type _WANPOTSLinkConfig1_GetISPInfo_Request struct {}

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
 err error) {
	request := _WANPOTSLinkConfig1_GetISPInfo_Request{

	}
	var response _WANPOTSLinkConfig1_GetISPInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetISPInfo", &request, &response)
	if err != nil {
		return
	}

	NewISPPhoneNumber = response.NewISPPhoneNumber

	NewISPInfo = response.NewISPInfo

	NewLinkType = response.NewLinkType

	return
}






// _WANPOTSLinkConfig1_GetCallRetryInfo_Request is the XML structure for the input arguments for action GetCallRetryInfo.
type _WANPOTSLinkConfig1_GetCallRetryInfo_Request struct {}

// _WANPOTSLinkConfig1_GetCallRetryInfo_Response is the XML structure for the output arguments for action GetCallRetryInfo.
type _WANPOTSLinkConfig1_GetCallRetryInfo_Response struct {
	NewNumberOfRetries uint32

	NewDelayBetweenRetries uint32
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
 err error) {
	request := _WANPOTSLinkConfig1_GetCallRetryInfo_Request{

	}
	var response _WANPOTSLinkConfig1_GetCallRetryInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetCallRetryInfo", &request, &response)
	if err != nil {
		return
	}

	NewNumberOfRetries = response.NewNumberOfRetries

	NewDelayBetweenRetries = response.NewDelayBetweenRetries

	return
}






// _WANPOTSLinkConfig1_GetFclass_Request is the XML structure for the input arguments for action GetFclass.
type _WANPOTSLinkConfig1_GetFclass_Request struct {}

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
 err error) {
	request := _WANPOTSLinkConfig1_GetFclass_Request{

	}
	var response _WANPOTSLinkConfig1_GetFclass_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetFclass", &request, &response)
	if err != nil {
		return
	}

	NewFclass = response.NewFclass

	return
}






// _WANPOTSLinkConfig1_GetDataModulationSupported_Request is the XML structure for the input arguments for action GetDataModulationSupported.
type _WANPOTSLinkConfig1_GetDataModulationSupported_Request struct {}

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
 err error) {
	request := _WANPOTSLinkConfig1_GetDataModulationSupported_Request{

	}
	var response _WANPOTSLinkConfig1_GetDataModulationSupported_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataModulationSupported", &request, &response)
	if err != nil {
		return
	}

	NewDataModulationSupported = response.NewDataModulationSupported

	return
}






// _WANPOTSLinkConfig1_GetDataProtocol_Request is the XML structure for the input arguments for action GetDataProtocol.
type _WANPOTSLinkConfig1_GetDataProtocol_Request struct {}

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
 err error) {
	request := _WANPOTSLinkConfig1_GetDataProtocol_Request{

	}
	var response _WANPOTSLinkConfig1_GetDataProtocol_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataProtocol", &request, &response)
	if err != nil {
		return
	}

	NewDataProtocol = response.NewDataProtocol

	return
}






// _WANPOTSLinkConfig1_GetDataCompression_Request is the XML structure for the input arguments for action GetDataCompression.
type _WANPOTSLinkConfig1_GetDataCompression_Request struct {}

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
 err error) {
	request := _WANPOTSLinkConfig1_GetDataCompression_Request{

	}
	var response _WANPOTSLinkConfig1_GetDataCompression_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataCompression", &request, &response)
	if err != nil {
		return
	}

	NewDataCompression = response.NewDataCompression

	return
}






// _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Request is the XML structure for the input arguments for action GetPlusVTRCommandSupported.
type _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Request struct {}

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
	NewPlusVTRCommandSupported string,
 err error) {
	request := _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Request{

	}
	var response _WANPOTSLinkConfig1_GetPlusVTRCommandSupported_Response
	err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetPlusVTRCommandSupported", &request, &response)
	if err != nil {
		return
	}

	NewPlusVTRCommandSupported = response.NewPlusVTRCommandSupported

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
type _WANPPPConnection1_SetConnectionType_Response struct {}

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
) ( err error) {
	request := _WANPPPConnection1_SetConnectionType_Request{

	NewConnectionType: NewConnectionType,

	}
	var response _WANPPPConnection1_SetConnectionType_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetConnectionType", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_GetConnectionTypeInfo_Request is the XML structure for the input arguments for action GetConnectionTypeInfo.
type _WANPPPConnection1_GetConnectionTypeInfo_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetConnectionTypeInfo_Request{

	}
	var response _WANPPPConnection1_GetConnectionTypeInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetConnectionTypeInfo", &request, &response)
	if err != nil {
		return
	}

	NewConnectionType = response.NewConnectionType

	NewPossibleConnectionTypes = response.NewPossibleConnectionTypes

	return
}






// _WANPPPConnection1_ConfigureConnection_Request is the XML structure for the input arguments for action ConfigureConnection.
type _WANPPPConnection1_ConfigureConnection_Request struct {
	NewUserName string

	NewPassword string
}

// _WANPPPConnection1_ConfigureConnection_Response is the XML structure for the output arguments for action ConfigureConnection.
type _WANPPPConnection1_ConfigureConnection_Response struct {}

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
) ( err error) {
	request := _WANPPPConnection1_ConfigureConnection_Request{

	NewUserName: NewUserName,

	NewPassword: NewPassword,

	}
	var response _WANPPPConnection1_ConfigureConnection_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ConfigureConnection", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_RequestConnection_Request is the XML structure for the input arguments for action RequestConnection.
type _WANPPPConnection1_RequestConnection_Request struct {}

// _WANPPPConnection1_RequestConnection_Response is the XML structure for the output arguments for action RequestConnection.
type _WANPPPConnection1_RequestConnection_Response struct {}

// RequestConnection action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) RequestConnection() ( err error) {
	request := _WANPPPConnection1_RequestConnection_Request{

	}
	var response _WANPPPConnection1_RequestConnection_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestConnection", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_RequestTermination_Request is the XML structure for the input arguments for action RequestTermination.
type _WANPPPConnection1_RequestTermination_Request struct {}

// _WANPPPConnection1_RequestTermination_Response is the XML structure for the output arguments for action RequestTermination.
type _WANPPPConnection1_RequestTermination_Response struct {}

// RequestTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) RequestTermination() ( err error) {
	request := _WANPPPConnection1_RequestTermination_Request{

	}
	var response _WANPPPConnection1_RequestTermination_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestTermination", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_ForceTermination_Request is the XML structure for the input arguments for action ForceTermination.
type _WANPPPConnection1_ForceTermination_Request struct {}

// _WANPPPConnection1_ForceTermination_Response is the XML structure for the output arguments for action ForceTermination.
type _WANPPPConnection1_ForceTermination_Response struct {}

// ForceTermination action.
// Arguments:
//
//
// Return values:
//
func (client *WANPPPConnection1) ForceTermination() ( err error) {
	request := _WANPPPConnection1_ForceTermination_Request{

	}
	var response _WANPPPConnection1_ForceTermination_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ForceTermination", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_SetAutoDisconnectTime_Request is the XML structure for the input arguments for action SetAutoDisconnectTime.
type _WANPPPConnection1_SetAutoDisconnectTime_Request struct {
	NewAutoDisconnectTime uint32
}

// _WANPPPConnection1_SetAutoDisconnectTime_Response is the XML structure for the output arguments for action SetAutoDisconnectTime.
type _WANPPPConnection1_SetAutoDisconnectTime_Response struct {}

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
) ( err error) {
	request := _WANPPPConnection1_SetAutoDisconnectTime_Request{

	NewAutoDisconnectTime: NewAutoDisconnectTime,

	}
	var response _WANPPPConnection1_SetAutoDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetAutoDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_SetIdleDisconnectTime_Request is the XML structure for the input arguments for action SetIdleDisconnectTime.
type _WANPPPConnection1_SetIdleDisconnectTime_Request struct {
	NewIdleDisconnectTime uint32
}

// _WANPPPConnection1_SetIdleDisconnectTime_Response is the XML structure for the output arguments for action SetIdleDisconnectTime.
type _WANPPPConnection1_SetIdleDisconnectTime_Response struct {}

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
) ( err error) {
	request := _WANPPPConnection1_SetIdleDisconnectTime_Request{

	NewIdleDisconnectTime: NewIdleDisconnectTime,

	}
	var response _WANPPPConnection1_SetIdleDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetIdleDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_SetWarnDisconnectDelay_Request is the XML structure for the input arguments for action SetWarnDisconnectDelay.
type _WANPPPConnection1_SetWarnDisconnectDelay_Request struct {
	NewWarnDisconnectDelay uint32
}

// _WANPPPConnection1_SetWarnDisconnectDelay_Response is the XML structure for the output arguments for action SetWarnDisconnectDelay.
type _WANPPPConnection1_SetWarnDisconnectDelay_Response struct {}

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
) ( err error) {
	request := _WANPPPConnection1_SetWarnDisconnectDelay_Request{

	NewWarnDisconnectDelay: NewWarnDisconnectDelay,

	}
	var response _WANPPPConnection1_SetWarnDisconnectDelay_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetWarnDisconnectDelay", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_GetStatusInfo_Request is the XML structure for the input arguments for action GetStatusInfo.
type _WANPPPConnection1_GetStatusInfo_Request struct {}

// _WANPPPConnection1_GetStatusInfo_Response is the XML structure for the output arguments for action GetStatusInfo.
type _WANPPPConnection1_GetStatusInfo_Response struct {
	NewConnectionStatus string

	NewLastConnectionError string

	NewUptime uint32
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
 err error) {
	request := _WANPPPConnection1_GetStatusInfo_Request{

	}
	var response _WANPPPConnection1_GetStatusInfo_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetStatusInfo", &request, &response)
	if err != nil {
		return
	}

	NewConnectionStatus = response.NewConnectionStatus

	NewLastConnectionError = response.NewLastConnectionError

	NewUptime = response.NewUptime

	return
}






// _WANPPPConnection1_GetLinkLayerMaxBitRates_Request is the XML structure for the input arguments for action GetLinkLayerMaxBitRates.
type _WANPPPConnection1_GetLinkLayerMaxBitRates_Request struct {}

// _WANPPPConnection1_GetLinkLayerMaxBitRates_Response is the XML structure for the output arguments for action GetLinkLayerMaxBitRates.
type _WANPPPConnection1_GetLinkLayerMaxBitRates_Response struct {
	NewUpstreamMaxBitRate uint32

	NewDownstreamMaxBitRate uint32
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
 err error) {
	request := _WANPPPConnection1_GetLinkLayerMaxBitRates_Request{

	}
	var response _WANPPPConnection1_GetLinkLayerMaxBitRates_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetLinkLayerMaxBitRates", &request, &response)
	if err != nil {
		return
	}

	NewUpstreamMaxBitRate = response.NewUpstreamMaxBitRate

	NewDownstreamMaxBitRate = response.NewDownstreamMaxBitRate

	return
}






// _WANPPPConnection1_GetPPPEncryptionProtocol_Request is the XML structure for the input arguments for action GetPPPEncryptionProtocol.
type _WANPPPConnection1_GetPPPEncryptionProtocol_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetPPPEncryptionProtocol_Request{

	}
	var response _WANPPPConnection1_GetPPPEncryptionProtocol_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPEncryptionProtocol", &request, &response)
	if err != nil {
		return
	}

	NewPPPEncryptionProtocol = response.NewPPPEncryptionProtocol

	return
}






// _WANPPPConnection1_GetPPPCompressionProtocol_Request is the XML structure for the input arguments for action GetPPPCompressionProtocol.
type _WANPPPConnection1_GetPPPCompressionProtocol_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetPPPCompressionProtocol_Request{

	}
	var response _WANPPPConnection1_GetPPPCompressionProtocol_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPCompressionProtocol", &request, &response)
	if err != nil {
		return
	}

	NewPPPCompressionProtocol = response.NewPPPCompressionProtocol

	return
}






// _WANPPPConnection1_GetPPPAuthenticationProtocol_Request is the XML structure for the input arguments for action GetPPPAuthenticationProtocol.
type _WANPPPConnection1_GetPPPAuthenticationProtocol_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetPPPAuthenticationProtocol_Request{

	}
	var response _WANPPPConnection1_GetPPPAuthenticationProtocol_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPAuthenticationProtocol", &request, &response)
	if err != nil {
		return
	}

	NewPPPAuthenticationProtocol = response.NewPPPAuthenticationProtocol

	return
}






// _WANPPPConnection1_GetUserName_Request is the XML structure for the input arguments for action GetUserName.
type _WANPPPConnection1_GetUserName_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetUserName_Request{

	}
	var response _WANPPPConnection1_GetUserName_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetUserName", &request, &response)
	if err != nil {
		return
	}

	NewUserName = response.NewUserName

	return
}






// _WANPPPConnection1_GetPassword_Request is the XML structure for the input arguments for action GetPassword.
type _WANPPPConnection1_GetPassword_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetPassword_Request{

	}
	var response _WANPPPConnection1_GetPassword_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPassword", &request, &response)
	if err != nil {
		return
	}

	NewPassword = response.NewPassword

	return
}






// _WANPPPConnection1_GetAutoDisconnectTime_Request is the XML structure for the input arguments for action GetAutoDisconnectTime.
type _WANPPPConnection1_GetAutoDisconnectTime_Request struct {}

// _WANPPPConnection1_GetAutoDisconnectTime_Response is the XML structure for the output arguments for action GetAutoDisconnectTime.
type _WANPPPConnection1_GetAutoDisconnectTime_Response struct {
	NewAutoDisconnectTime uint32
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
 err error) {
	request := _WANPPPConnection1_GetAutoDisconnectTime_Request{

	}
	var response _WANPPPConnection1_GetAutoDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetAutoDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	NewAutoDisconnectTime = response.NewAutoDisconnectTime

	return
}






// _WANPPPConnection1_GetIdleDisconnectTime_Request is the XML structure for the input arguments for action GetIdleDisconnectTime.
type _WANPPPConnection1_GetIdleDisconnectTime_Request struct {}

// _WANPPPConnection1_GetIdleDisconnectTime_Response is the XML structure for the output arguments for action GetIdleDisconnectTime.
type _WANPPPConnection1_GetIdleDisconnectTime_Response struct {
	NewIdleDisconnectTime uint32
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
 err error) {
	request := _WANPPPConnection1_GetIdleDisconnectTime_Request{

	}
	var response _WANPPPConnection1_GetIdleDisconnectTime_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetIdleDisconnectTime", &request, &response)
	if err != nil {
		return
	}

	NewIdleDisconnectTime = response.NewIdleDisconnectTime

	return
}






// _WANPPPConnection1_GetWarnDisconnectDelay_Request is the XML structure for the input arguments for action GetWarnDisconnectDelay.
type _WANPPPConnection1_GetWarnDisconnectDelay_Request struct {}

// _WANPPPConnection1_GetWarnDisconnectDelay_Response is the XML structure for the output arguments for action GetWarnDisconnectDelay.
type _WANPPPConnection1_GetWarnDisconnectDelay_Response struct {
	NewWarnDisconnectDelay uint32
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
 err error) {
	request := _WANPPPConnection1_GetWarnDisconnectDelay_Request{

	}
	var response _WANPPPConnection1_GetWarnDisconnectDelay_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetWarnDisconnectDelay", &request, &response)
	if err != nil {
		return
	}

	NewWarnDisconnectDelay = response.NewWarnDisconnectDelay

	return
}






// _WANPPPConnection1_GetNATRSIPStatus_Request is the XML structure for the input arguments for action GetNATRSIPStatus.
type _WANPPPConnection1_GetNATRSIPStatus_Request struct {}

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
	NewRSIPAvailable string,

	NewNATEnabled string,
 err error) {
	request := _WANPPPConnection1_GetNATRSIPStatus_Request{

	}
	var response _WANPPPConnection1_GetNATRSIPStatus_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetNATRSIPStatus", &request, &response)
	if err != nil {
		return
	}

	NewRSIPAvailable = response.NewRSIPAvailable

	NewNATEnabled = response.NewNATEnabled

	return
}






// _WANPPPConnection1_GetGenericPortMappingEntry_Request is the XML structure for the input arguments for action GetGenericPortMappingEntry.
type _WANPPPConnection1_GetGenericPortMappingEntry_Request struct {
	NewPortMappingIndex uint16
}

// _WANPPPConnection1_GetGenericPortMappingEntry_Response is the XML structure for the output arguments for action GetGenericPortMappingEntry.
type _WANPPPConnection1_GetGenericPortMappingEntry_Response struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string

	NewInternalPort uint16

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration uint32
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

	NewEnabled string,

	NewPortMappingDescription string,

	NewLeaseDuration uint32,
 err error) {
	request := _WANPPPConnection1_GetGenericPortMappingEntry_Request{

	NewPortMappingIndex: NewPortMappingIndex,

	}
	var response _WANPPPConnection1_GetGenericPortMappingEntry_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetGenericPortMappingEntry", &request, &response)
	if err != nil {
		return
	}

	NewRemoteHost = response.NewRemoteHost

	NewExternalPort = response.NewExternalPort

	NewProtocol = response.NewProtocol

	NewInternalPort = response.NewInternalPort

	NewInternalClient = response.NewInternalClient

	NewEnabled = response.NewEnabled

	NewPortMappingDescription = response.NewPortMappingDescription

	NewLeaseDuration = response.NewLeaseDuration

	return
}






// _WANPPPConnection1_GetSpecificPortMappingEntry_Request is the XML structure for the input arguments for action GetSpecificPortMappingEntry.
type _WANPPPConnection1_GetSpecificPortMappingEntry_Request struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string
}

// _WANPPPConnection1_GetSpecificPortMappingEntry_Response is the XML structure for the output arguments for action GetSpecificPortMappingEntry.
type _WANPPPConnection1_GetSpecificPortMappingEntry_Response struct {
	NewInternalPort uint16

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration uint32
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

	NewEnabled string,

	NewPortMappingDescription string,

	NewLeaseDuration uint32,
 err error) {
	request := _WANPPPConnection1_GetSpecificPortMappingEntry_Request{

	NewRemoteHost: NewRemoteHost,

	NewExternalPort: NewExternalPort,

	NewProtocol: NewProtocol,

	}
	var response _WANPPPConnection1_GetSpecificPortMappingEntry_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetSpecificPortMappingEntry", &request, &response)
	if err != nil {
		return
	}

	NewInternalPort = response.NewInternalPort

	NewInternalClient = response.NewInternalClient

	NewEnabled = response.NewEnabled

	NewPortMappingDescription = response.NewPortMappingDescription

	NewLeaseDuration = response.NewLeaseDuration

	return
}






// _WANPPPConnection1_AddPortMapping_Request is the XML structure for the input arguments for action AddPortMapping.
type _WANPPPConnection1_AddPortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string

	NewInternalPort uint16

	NewInternalClient string

	NewEnabled string

	NewPortMappingDescription string

	NewLeaseDuration uint32
}

// _WANPPPConnection1_AddPortMapping_Response is the XML structure for the output arguments for action AddPortMapping.
type _WANPPPConnection1_AddPortMapping_Response struct {}

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

	NewEnabled string,

	NewPortMappingDescription string,

	NewLeaseDuration uint32,
) ( err error) {
	request := _WANPPPConnection1_AddPortMapping_Request{

	NewRemoteHost: NewRemoteHost,

	NewExternalPort: NewExternalPort,

	NewProtocol: NewProtocol,

	NewInternalPort: NewInternalPort,

	NewInternalClient: NewInternalClient,

	NewEnabled: NewEnabled,

	NewPortMappingDescription: NewPortMappingDescription,

	NewLeaseDuration: NewLeaseDuration,

	}
	var response _WANPPPConnection1_AddPortMapping_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "AddPortMapping", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_DeletePortMapping_Request is the XML structure for the input arguments for action DeletePortMapping.
type _WANPPPConnection1_DeletePortMapping_Request struct {
	NewRemoteHost string

	NewExternalPort uint16

	NewProtocol string
}

// _WANPPPConnection1_DeletePortMapping_Response is the XML structure for the output arguments for action DeletePortMapping.
type _WANPPPConnection1_DeletePortMapping_Response struct {}

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
) ( err error) {
	request := _WANPPPConnection1_DeletePortMapping_Request{

	NewRemoteHost: NewRemoteHost,

	NewExternalPort: NewExternalPort,

	NewProtocol: NewProtocol,

	}
	var response _WANPPPConnection1_DeletePortMapping_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "DeletePortMapping", &request, &response)
	if err != nil {
		return
	}

	return
}






// _WANPPPConnection1_GetExternalIPAddress_Request is the XML structure for the input arguments for action GetExternalIPAddress.
type _WANPPPConnection1_GetExternalIPAddress_Request struct {}

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
 err error) {
	request := _WANPPPConnection1_GetExternalIPAddress_Request{

	}
	var response _WANPPPConnection1_GetExternalIPAddress_Response
	err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetExternalIPAddress", &request, &response)
	if err != nil {
		return
	}

	NewExternalIPAddress = response.NewExternalIPAddress

	return
}



