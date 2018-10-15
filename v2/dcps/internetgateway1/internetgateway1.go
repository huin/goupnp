// Package internetgateway1 is a client for UPnP Device Control Protocol
// Internet Gateway Device v1.
//
// This DCP is documented in detail at: http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v1-Device.pdf
//
// Typically, use one of the New* functions to create clients for services.
package internetgateway1

// ***********************************************************
// GENERATED FILE - DO NOT EDIT BY HAND. See README.md
// ***********************************************************

import (
	"context"
	"net/url"
	"time"

	"github.com/huin/goupnp/v2/discover"
	"github.com/huin/goupnp/v2/metadata"
	"github.com/huin/goupnp/v2/soap"
	"github.com/huin/goupnp/v2/ssdp"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const (
	URN_LANDevice_1           = "urn:schemas-upnp-org:device:LANDevice:1"
	URN_WANConnectionDevice_1 = "urn:schemas-upnp-org:device:WANConnectionDevice:1"
	URN_WANDevice_1           = "urn:schemas-upnp-org:device:WANDevice:1"
)

// Service URNs:
const (
	URN_LANHostConfigManagement_1  = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"
	URN_Layer3Forwarding_1         = "urn:schemas-upnp-org:service:Layer3Forwarding:1"
	URN_WANCableLinkConfig_1       = "urn:schemas-upnp-org:service:WANCableLinkConfig:1"
	URN_WANCommonInterfaceConfig_1 = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"
	URN_WANDSLLinkConfig_1         = "urn:schemas-upnp-org:service:WANDSLLinkConfig:1"
	URN_WANEthernetLinkConfig_1    = "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1"
	URN_WANIPConnection_1          = "urn:schemas-upnp-org:service:WANIPConnection:1"
	URN_WANPOTSLinkConfig_1        = "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1"
	URN_WANPPPConnection_1         = "urn:schemas-upnp-org:service:WANPPPConnection:1"
)

// LANHostConfigManagement1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:LANHostConfigManagement:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type LANHostConfigManagement1 struct {
	discover.ServiceClient
}

// NewLANHostConfigManagement1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewLANHostConfigManagement1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*LANHostConfigManagement1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_LANHostConfigManagement_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newLANHostConfigManagement1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewLANHostConfigManagement1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewLANHostConfigManagement1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*LANHostConfigManagement1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_LANHostConfigManagement_1,
	)
	if err != nil {
		return nil, err
	}
	return newLANHostConfigManagement1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewLANHostConfigManagement1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewLANHostConfigManagement1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*LANHostConfigManagement1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_LANHostConfigManagement_1,
	)
	if err != nil {
		return nil, err
	}
	return newLANHostConfigManagement1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newLANHostConfigManagement1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*LANHostConfigManagement1 {
	clients := make([]*LANHostConfigManagement1, len(genericClients))
	for i := range genericClients {
		clients[i] = &LANHostConfigManagement1{genericClients[i]}
	}
	return clients
}

// SetDHCPServerConfigurable implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetDHCPServerConfigurable(
	ctx context.Context,
	NewDHCPServerConfigurable bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDHCPServerConfigurable string `xml:"NewDHCPServerConfigurable"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDHCPServerConfigurable, err = soap.MarshalBoolean(
		NewDHCPServerConfigurable,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetDHCPServerConfigurable",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDHCPServerConfigurable implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetDHCPServerConfigurable(
	ctx context.Context,
) (
	NewDHCPServerConfigurable bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDHCPServerConfigurable string `xml:"NewDHCPServerConfigurable"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetDHCPServerConfigurable",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDHCPServerConfigurable, err = soap.UnmarshalBoolean(
		response.NewDHCPServerConfigurable,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetDHCPRelay implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetDHCPRelay(
	ctx context.Context,
	NewDHCPRelay bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDHCPRelay string `xml:"NewDHCPRelay"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDHCPRelay, err = soap.MarshalBoolean(
		NewDHCPRelay,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetDHCPRelay",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDHCPRelay implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetDHCPRelay(
	ctx context.Context,
) (
	NewDHCPRelay bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDHCPRelay string `xml:"NewDHCPRelay"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetDHCPRelay",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDHCPRelay, err = soap.UnmarshalBoolean(
		response.NewDHCPRelay,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetSubnetMask implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetSubnetMask(
	ctx context.Context,
	NewSubnetMask string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewSubnetMask string `xml:"NewSubnetMask"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewSubnetMask, err = soap.MarshalString(
		NewSubnetMask,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetSubnetMask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetSubnetMask implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetSubnetMask(
	ctx context.Context,
) (
	NewSubnetMask string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewSubnetMask string `xml:"NewSubnetMask"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetSubnetMask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewSubnetMask, err = soap.UnmarshalString(
		response.NewSubnetMask,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetIPRouter implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetIPRouter(
	ctx context.Context,
	NewIPRouters string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewIPRouters string `xml:"NewIPRouters"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewIPRouters, err = soap.MarshalString(
		NewIPRouters,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetIPRouter",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeleteIPRouter implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) DeleteIPRouter(
	ctx context.Context,
	NewIPRouters string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewIPRouters string `xml:"NewIPRouters"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewIPRouters, err = soap.MarshalString(
		NewIPRouters,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"DeleteIPRouter",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetIPRoutersList implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetIPRoutersList(
	ctx context.Context,
) (
	NewIPRouters string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewIPRouters string `xml:"NewIPRouters"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetIPRoutersList",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewIPRouters, err = soap.UnmarshalString(
		response.NewIPRouters,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetDomainName implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetDomainName(
	ctx context.Context,
	NewDomainName string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDomainName string `xml:"NewDomainName"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDomainName, err = soap.MarshalString(
		NewDomainName,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetDomainName",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDomainName implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetDomainName(
	ctx context.Context,
) (
	NewDomainName string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDomainName string `xml:"NewDomainName"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetDomainName",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDomainName, err = soap.UnmarshalString(
		response.NewDomainName,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetAddressRange implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetAddressRange(
	ctx context.Context,
	NewMinAddress string,
	NewMaxAddress string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewMinAddress string `xml:"NewMinAddress"`
		NewMaxAddress string `xml:"NewMaxAddress"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewMinAddress, err = soap.MarshalString(
		NewMinAddress,
	); err != nil {
		return
	}
	if request.NewMaxAddress, err = soap.MarshalString(
		NewMaxAddress,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetAddressRange",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetAddressRange implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetAddressRange(
	ctx context.Context,
) (
	NewMinAddress string,
	NewMaxAddress string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewMinAddress string `xml:"NewMinAddress"`
		NewMaxAddress string `xml:"NewMaxAddress"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetAddressRange",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewMinAddress, err = soap.UnmarshalString(
		response.NewMinAddress,
	); err != nil {
		return
	}
	if NewMaxAddress, err = soap.UnmarshalString(
		response.NewMaxAddress,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetReservedAddress implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetReservedAddress(
	ctx context.Context,
	NewReservedAddresses string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewReservedAddresses string `xml:"NewReservedAddresses"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewReservedAddresses, err = soap.MarshalString(
		NewReservedAddresses,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetReservedAddress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeleteReservedAddress implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) DeleteReservedAddress(
	ctx context.Context,
	NewReservedAddresses string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewReservedAddresses string `xml:"NewReservedAddresses"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewReservedAddresses, err = soap.MarshalString(
		NewReservedAddresses,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"DeleteReservedAddress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetReservedAddresses implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetReservedAddresses(
	ctx context.Context,
) (
	NewReservedAddresses string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewReservedAddresses string `xml:"NewReservedAddresses"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetReservedAddresses",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewReservedAddresses, err = soap.UnmarshalString(
		response.NewReservedAddresses,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetDNSServer implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) SetDNSServer(
	ctx context.Context,
	NewDNSServers string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDNSServers string `xml:"NewDNSServers"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDNSServers, err = soap.MarshalString(
		NewDNSServers,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"SetDNSServer",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeleteDNSServer implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) DeleteDNSServer(
	ctx context.Context,
	NewDNSServers string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDNSServers string `xml:"NewDNSServers"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDNSServers, err = soap.MarshalString(
		NewDNSServers,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"DeleteDNSServer",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDNSServers implements a UPnP action of the same name.
func (client *LANHostConfigManagement1) GetDNSServers(
	ctx context.Context,
) (
	NewDNSServers string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDNSServers string `xml:"NewDNSServers"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_LANHostConfigManagement_1,
		"GetDNSServers",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDNSServers, err = soap.UnmarshalString(
		response.NewDNSServers,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Layer3Forwarding1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:Layer3Forwarding:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type Layer3Forwarding1 struct {
	discover.ServiceClient
}

// NewLayer3Forwarding1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewLayer3Forwarding1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*Layer3Forwarding1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_Layer3Forwarding_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newLayer3Forwarding1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewLayer3Forwarding1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewLayer3Forwarding1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*Layer3Forwarding1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_Layer3Forwarding_1,
	)
	if err != nil {
		return nil, err
	}
	return newLayer3Forwarding1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewLayer3Forwarding1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewLayer3Forwarding1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*Layer3Forwarding1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_Layer3Forwarding_1,
	)
	if err != nil {
		return nil, err
	}
	return newLayer3Forwarding1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newLayer3Forwarding1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*Layer3Forwarding1 {
	clients := make([]*Layer3Forwarding1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Layer3Forwarding1{genericClients[i]}
	}
	return clients
}

// SetDefaultConnectionService implements a UPnP action of the same name.
func (client *Layer3Forwarding1) SetDefaultConnectionService(
	ctx context.Context,
	NewDefaultConnectionService string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDefaultConnectionService string `xml:"NewDefaultConnectionService"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDefaultConnectionService, err = soap.MarshalString(
		NewDefaultConnectionService,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_Layer3Forwarding_1,
		"SetDefaultConnectionService",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDefaultConnectionService implements a UPnP action of the same name.
func (client *Layer3Forwarding1) GetDefaultConnectionService(
	ctx context.Context,
) (
	NewDefaultConnectionService string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDefaultConnectionService string `xml:"NewDefaultConnectionService"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_Layer3Forwarding_1,
		"GetDefaultConnectionService",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDefaultConnectionService, err = soap.UnmarshalString(
		response.NewDefaultConnectionService,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANCableLinkConfig1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANCableLinkConfig:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANCableLinkConfig1 struct {
	discover.ServiceClient
}

// NewWANCableLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANCableLinkConfig1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANCableLinkConfig1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANCableLinkConfig_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANCableLinkConfig1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANCableLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANCableLinkConfig1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANCableLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANCableLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANCableLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANCableLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANCableLinkConfig1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANCableLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANCableLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANCableLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANCableLinkConfig1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANCableLinkConfig1 {
	clients := make([]*WANCableLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANCableLinkConfig1{genericClients[i]}
	}
	return clients
}

// GetCableLinkConfigInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewCableLinkConfigState: allowed values: notReady, dsSyncComplete, usParamAcquired, rangingComplete, ipComplete, todEstablished, paramTransferComplete, registrationComplete, operational, accessDenied
//
// * NewLinkType: allowed values: Ethernet
func (client *WANCableLinkConfig1) GetCableLinkConfigInfo(
	ctx context.Context,
) (
	NewCableLinkConfigState string,
	NewLinkType string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewCableLinkConfigState string `xml:"NewCableLinkConfigState"`
		NewLinkType             string `xml:"NewLinkType"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetCableLinkConfigInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewCableLinkConfigState, err = soap.UnmarshalString(
		response.NewCableLinkConfigState,
	); err != nil {
		return
	}
	if NewLinkType, err = soap.UnmarshalString(
		response.NewLinkType,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDownstreamFrequency implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetDownstreamFrequency(
	ctx context.Context,
) (
	NewDownstreamFrequency uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDownstreamFrequency string `xml:"NewDownstreamFrequency"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetDownstreamFrequency",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDownstreamFrequency, err = soap.UnmarshalUI4(
		response.NewDownstreamFrequency,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDownstreamModulation implements a UPnP action of the same name.
//
// Return values:
//
// * NewDownstreamModulation: allowed values: 64QAM, 256QAM
func (client *WANCableLinkConfig1) GetDownstreamModulation(
	ctx context.Context,
) (
	NewDownstreamModulation string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDownstreamModulation string `xml:"NewDownstreamModulation"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetDownstreamModulation",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDownstreamModulation, err = soap.UnmarshalString(
		response.NewDownstreamModulation,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetUpstreamFrequency implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetUpstreamFrequency(
	ctx context.Context,
) (
	NewUpstreamFrequency uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewUpstreamFrequency string `xml:"NewUpstreamFrequency"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetUpstreamFrequency",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewUpstreamFrequency, err = soap.UnmarshalUI4(
		response.NewUpstreamFrequency,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetUpstreamModulation implements a UPnP action of the same name.
//
// Return values:
//
// * NewUpstreamModulation: allowed values: QPSK, 16QAM
func (client *WANCableLinkConfig1) GetUpstreamModulation(
	ctx context.Context,
) (
	NewUpstreamModulation string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewUpstreamModulation string `xml:"NewUpstreamModulation"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetUpstreamModulation",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewUpstreamModulation, err = soap.UnmarshalString(
		response.NewUpstreamModulation,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetUpstreamChannelID implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetUpstreamChannelID(
	ctx context.Context,
) (
	NewUpstreamChannelID uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewUpstreamChannelID string `xml:"NewUpstreamChannelID"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetUpstreamChannelID",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewUpstreamChannelID, err = soap.UnmarshalUI4(
		response.NewUpstreamChannelID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetUpstreamPowerLevel implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetUpstreamPowerLevel(
	ctx context.Context,
) (
	NewUpstreamPowerLevel uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewUpstreamPowerLevel string `xml:"NewUpstreamPowerLevel"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetUpstreamPowerLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewUpstreamPowerLevel, err = soap.UnmarshalUI4(
		response.NewUpstreamPowerLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetBPIEncryptionEnabled implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetBPIEncryptionEnabled(
	ctx context.Context,
) (
	NewBPIEncryptionEnabled bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewBPIEncryptionEnabled string `xml:"NewBPIEncryptionEnabled"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetBPIEncryptionEnabled",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewBPIEncryptionEnabled, err = soap.UnmarshalBoolean(
		response.NewBPIEncryptionEnabled,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetConfigFile implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetConfigFile(
	ctx context.Context,
) (
	NewConfigFile string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewConfigFile string `xml:"NewConfigFile"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetConfigFile",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewConfigFile, err = soap.UnmarshalString(
		response.NewConfigFile,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTFTPServer implements a UPnP action of the same name.
func (client *WANCableLinkConfig1) GetTFTPServer(
	ctx context.Context,
) (
	NewTFTPServer string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewTFTPServer string `xml:"NewTFTPServer"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCableLinkConfig_1,
		"GetTFTPServer",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewTFTPServer, err = soap.UnmarshalString(
		response.NewTFTPServer,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANCommonInterfaceConfig1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANCommonInterfaceConfig1 struct {
	discover.ServiceClient
}

// NewWANCommonInterfaceConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANCommonInterfaceConfig1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANCommonInterfaceConfig1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANCommonInterfaceConfig1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANCommonInterfaceConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANCommonInterfaceConfig1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANCommonInterfaceConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANCommonInterfaceConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANCommonInterfaceConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANCommonInterfaceConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANCommonInterfaceConfig1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANCommonInterfaceConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANCommonInterfaceConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANCommonInterfaceConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANCommonInterfaceConfig1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANCommonInterfaceConfig1 {
	clients := make([]*WANCommonInterfaceConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANCommonInterfaceConfig1{genericClients[i]}
	}
	return clients
}

// SetEnabledForInternet implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) SetEnabledForInternet(
	ctx context.Context,
	NewEnabledForInternet bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewEnabledForInternet string `xml:"NewEnabledForInternet"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewEnabledForInternet, err = soap.MarshalBoolean(
		NewEnabledForInternet,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"SetEnabledForInternet",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetEnabledForInternet implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetEnabledForInternet(
	ctx context.Context,
) (
	NewEnabledForInternet bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewEnabledForInternet string `xml:"NewEnabledForInternet"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetEnabledForInternet",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewEnabledForInternet, err = soap.UnmarshalBoolean(
		response.NewEnabledForInternet,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetCommonLinkProperties implements a UPnP action of the same name.
//
// Return values:
//
// * NewWANAccessType: allowed values: DSL, POTS, Cable, Ethernet
//
// * NewPhysicalLinkStatus: allowed values: Up, Down
func (client *WANCommonInterfaceConfig1) GetCommonLinkProperties(
	ctx context.Context,
) (
	NewWANAccessType string,
	NewLayer1UpstreamMaxBitRate uint32,
	NewLayer1DownstreamMaxBitRate uint32,
	NewPhysicalLinkStatus string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewWANAccessType              string `xml:"NewWANAccessType"`
		NewLayer1UpstreamMaxBitRate   string `xml:"NewLayer1UpstreamMaxBitRate"`
		NewLayer1DownstreamMaxBitRate string `xml:"NewLayer1DownstreamMaxBitRate"`
		NewPhysicalLinkStatus         string `xml:"NewPhysicalLinkStatus"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetCommonLinkProperties",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewWANAccessType, err = soap.UnmarshalString(
		response.NewWANAccessType,
	); err != nil {
		return
	}
	if NewLayer1UpstreamMaxBitRate, err = soap.UnmarshalUI4(
		response.NewLayer1UpstreamMaxBitRate,
	); err != nil {
		return
	}
	if NewLayer1DownstreamMaxBitRate, err = soap.UnmarshalUI4(
		response.NewLayer1DownstreamMaxBitRate,
	); err != nil {
		return
	}
	if NewPhysicalLinkStatus, err = soap.UnmarshalString(
		response.NewPhysicalLinkStatus,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetWANAccessProvider implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetWANAccessProvider(
	ctx context.Context,
) (
	NewWANAccessProvider string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewWANAccessProvider string `xml:"NewWANAccessProvider"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetWANAccessProvider",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewWANAccessProvider, err = soap.UnmarshalString(
		response.NewWANAccessProvider,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetMaximumActiveConnections implements a UPnP action of the same name.
//
// Return values:
//
// * NewMaximumActiveConnections: allowed value range: minimum=1, step=1
func (client *WANCommonInterfaceConfig1) GetMaximumActiveConnections(
	ctx context.Context,
) (
	NewMaximumActiveConnections uint16,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewMaximumActiveConnections string `xml:"NewMaximumActiveConnections"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetMaximumActiveConnections",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewMaximumActiveConnections, err = soap.UnmarshalUI2(
		response.NewMaximumActiveConnections,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTotalBytesSent implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetTotalBytesSent(
	ctx context.Context,
) (
	NewTotalBytesSent uint64,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewTotalBytesSent string `xml:"NewTotalBytesSent"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetTotalBytesSent",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewTotalBytesSent, err = soap.UnmarshalUI8(
		response.NewTotalBytesSent,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTotalBytesReceived implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetTotalBytesReceived(
	ctx context.Context,
) (
	NewTotalBytesReceived uint64,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewTotalBytesReceived string `xml:"NewTotalBytesReceived"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetTotalBytesReceived",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewTotalBytesReceived, err = soap.UnmarshalUI8(
		response.NewTotalBytesReceived,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTotalPacketsSent implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetTotalPacketsSent(
	ctx context.Context,
) (
	NewTotalPacketsSent uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewTotalPacketsSent string `xml:"NewTotalPacketsSent"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetTotalPacketsSent",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewTotalPacketsSent, err = soap.UnmarshalUI4(
		response.NewTotalPacketsSent,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTotalPacketsReceived implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetTotalPacketsReceived(
	ctx context.Context,
) (
	NewTotalPacketsReceived uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewTotalPacketsReceived string `xml:"NewTotalPacketsReceived"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetTotalPacketsReceived",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewTotalPacketsReceived, err = soap.UnmarshalUI4(
		response.NewTotalPacketsReceived,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetActiveConnection implements a UPnP action of the same name.
func (client *WANCommonInterfaceConfig1) GetActiveConnection(
	ctx context.Context,
	NewActiveConnectionIndex uint16,
) (
	NewActiveConnDeviceContainer string,
	NewActiveConnectionServiceID string,
	err error,
) {
	// Request structure.
	request := &struct {
		NewActiveConnectionIndex string `xml:"NewActiveConnectionIndex"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewActiveConnectionIndex, err = soap.MarshalUI2(
		NewActiveConnectionIndex,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewActiveConnDeviceContainer string `xml:"NewActiveConnDeviceContainer"`
		NewActiveConnectionServiceID string `xml:"NewActiveConnectionServiceID"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANCommonInterfaceConfig_1,
		"GetActiveConnection",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewActiveConnDeviceContainer, err = soap.UnmarshalString(
		response.NewActiveConnDeviceContainer,
	); err != nil {
		return
	}
	if NewActiveConnectionServiceID, err = soap.UnmarshalString(
		response.NewActiveConnectionServiceID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANDSLLinkConfig1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANDSLLinkConfig:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANDSLLinkConfig1 struct {
	discover.ServiceClient
}

// NewWANDSLLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANDSLLinkConfig1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANDSLLinkConfig1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANDSLLinkConfig_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANDSLLinkConfig1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANDSLLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANDSLLinkConfig1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANDSLLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANDSLLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANDSLLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANDSLLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANDSLLinkConfig1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANDSLLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANDSLLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANDSLLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANDSLLinkConfig1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANDSLLinkConfig1 {
	clients := make([]*WANDSLLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANDSLLinkConfig1{genericClients[i]}
	}
	return clients
}

// SetDSLLinkType implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) SetDSLLinkType(
	ctx context.Context,
	NewLinkType string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewLinkType string `xml:"NewLinkType"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewLinkType, err = soap.MarshalString(
		NewLinkType,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"SetDSLLinkType",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDSLLinkInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewLinkStatus: allowed values: Up, Down
func (client *WANDSLLinkConfig1) GetDSLLinkInfo(
	ctx context.Context,
) (
	NewLinkType string,
	NewLinkStatus string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewLinkType   string `xml:"NewLinkType"`
		NewLinkStatus string `xml:"NewLinkStatus"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"GetDSLLinkInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewLinkType, err = soap.UnmarshalString(
		response.NewLinkType,
	); err != nil {
		return
	}
	if NewLinkStatus, err = soap.UnmarshalString(
		response.NewLinkStatus,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetAutoConfig implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) GetAutoConfig(
	ctx context.Context,
) (
	NewAutoConfig bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewAutoConfig string `xml:"NewAutoConfig"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"GetAutoConfig",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewAutoConfig, err = soap.UnmarshalBoolean(
		response.NewAutoConfig,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetModulationType implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) GetModulationType(
	ctx context.Context,
) (
	NewModulationType string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewModulationType string `xml:"NewModulationType"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"GetModulationType",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewModulationType, err = soap.UnmarshalString(
		response.NewModulationType,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetDestinationAddress implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) SetDestinationAddress(
	ctx context.Context,
	NewDestinationAddress string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewDestinationAddress string `xml:"NewDestinationAddress"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewDestinationAddress, err = soap.MarshalString(
		NewDestinationAddress,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"SetDestinationAddress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetDestinationAddress implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) GetDestinationAddress(
	ctx context.Context,
) (
	NewDestinationAddress string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDestinationAddress string `xml:"NewDestinationAddress"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"GetDestinationAddress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDestinationAddress, err = soap.UnmarshalString(
		response.NewDestinationAddress,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetATMEncapsulation implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) SetATMEncapsulation(
	ctx context.Context,
	NewATMEncapsulation string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewATMEncapsulation string `xml:"NewATMEncapsulation"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewATMEncapsulation, err = soap.MarshalString(
		NewATMEncapsulation,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"SetATMEncapsulation",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetATMEncapsulation implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) GetATMEncapsulation(
	ctx context.Context,
) (
	NewATMEncapsulation string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewATMEncapsulation string `xml:"NewATMEncapsulation"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"GetATMEncapsulation",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewATMEncapsulation, err = soap.UnmarshalString(
		response.NewATMEncapsulation,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetFCSPreserved implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) SetFCSPreserved(
	ctx context.Context,
	NewFCSPreserved bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewFCSPreserved string `xml:"NewFCSPreserved"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewFCSPreserved, err = soap.MarshalBoolean(
		NewFCSPreserved,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"SetFCSPreserved",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetFCSPreserved implements a UPnP action of the same name.
func (client *WANDSLLinkConfig1) GetFCSPreserved(
	ctx context.Context,
) (
	NewFCSPreserved bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewFCSPreserved string `xml:"NewFCSPreserved"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANDSLLinkConfig_1,
		"GetFCSPreserved",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewFCSPreserved, err = soap.UnmarshalBoolean(
		response.NewFCSPreserved,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANEthernetLinkConfig1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANEthernetLinkConfig1 struct {
	discover.ServiceClient
}

// NewWANEthernetLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANEthernetLinkConfig1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANEthernetLinkConfig1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANEthernetLinkConfig_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANEthernetLinkConfig1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANEthernetLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANEthernetLinkConfig1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANEthernetLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANEthernetLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANEthernetLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANEthernetLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANEthernetLinkConfig1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANEthernetLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANEthernetLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANEthernetLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANEthernetLinkConfig1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANEthernetLinkConfig1 {
	clients := make([]*WANEthernetLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANEthernetLinkConfig1{genericClients[i]}
	}
	return clients
}

// GetEthernetLinkStatus implements a UPnP action of the same name.
//
// Return values:
//
// * NewEthernetLinkStatus: allowed values: Up, Down
func (client *WANEthernetLinkConfig1) GetEthernetLinkStatus(
	ctx context.Context,
) (
	NewEthernetLinkStatus string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewEthernetLinkStatus string `xml:"NewEthernetLinkStatus"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANEthernetLinkConfig_1,
		"GetEthernetLinkStatus",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewEthernetLinkStatus, err = soap.UnmarshalString(
		response.NewEthernetLinkStatus,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANIPConnection1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANIPConnection:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANIPConnection1 struct {
	discover.ServiceClient
}

// NewWANIPConnection1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANIPConnection1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANIPConnection1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANIPConnection_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANIPConnection1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANIPConnection1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANIPConnection1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANIPConnection1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANIPConnection_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANIPConnection1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANIPConnection1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANIPConnection1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANIPConnection1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANIPConnection_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANIPConnection1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANIPConnection1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANIPConnection1 {
	clients := make([]*WANIPConnection1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANIPConnection1{genericClients[i]}
	}
	return clients
}

// SetConnectionType implements a UPnP action of the same name.
func (client *WANIPConnection1) SetConnectionType(
	ctx context.Context,
	NewConnectionType string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewConnectionType string `xml:"NewConnectionType"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewConnectionType, err = soap.MarshalString(
		NewConnectionType,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"SetConnectionType",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetConnectionTypeInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewPossibleConnectionTypes: allowed values: Unconfigured, IP_Routed, IP_Bridged
func (client *WANIPConnection1) GetConnectionTypeInfo(
	ctx context.Context,
) (
	NewConnectionType string,
	NewPossibleConnectionTypes string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewConnectionType          string `xml:"NewConnectionType"`
		NewPossibleConnectionTypes string `xml:"NewPossibleConnectionTypes"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetConnectionTypeInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewConnectionType, err = soap.UnmarshalString(
		response.NewConnectionType,
	); err != nil {
		return
	}
	if NewPossibleConnectionTypes, err = soap.UnmarshalString(
		response.NewPossibleConnectionTypes,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// RequestConnection implements a UPnP action of the same name.
func (client *WANIPConnection1) RequestConnection(
	ctx context.Context,
) (
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"RequestConnection",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// RequestTermination implements a UPnP action of the same name.
func (client *WANIPConnection1) RequestTermination(
	ctx context.Context,
) (
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"RequestTermination",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// ForceTermination implements a UPnP action of the same name.
func (client *WANIPConnection1) ForceTermination(
	ctx context.Context,
) (
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"ForceTermination",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetAutoDisconnectTime implements a UPnP action of the same name.
func (client *WANIPConnection1) SetAutoDisconnectTime(
	ctx context.Context,
	NewAutoDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewAutoDisconnectTime string `xml:"NewAutoDisconnectTime"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewAutoDisconnectTime, err = soap.MarshalUI4(
		NewAutoDisconnectTime,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"SetAutoDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetIdleDisconnectTime implements a UPnP action of the same name.
func (client *WANIPConnection1) SetIdleDisconnectTime(
	ctx context.Context,
	NewIdleDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewIdleDisconnectTime string `xml:"NewIdleDisconnectTime"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewIdleDisconnectTime, err = soap.MarshalUI4(
		NewIdleDisconnectTime,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"SetIdleDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetWarnDisconnectDelay implements a UPnP action of the same name.
func (client *WANIPConnection1) SetWarnDisconnectDelay(
	ctx context.Context,
	NewWarnDisconnectDelay uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewWarnDisconnectDelay string `xml:"NewWarnDisconnectDelay"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewWarnDisconnectDelay, err = soap.MarshalUI4(
		NewWarnDisconnectDelay,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"SetWarnDisconnectDelay",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetStatusInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewConnectionStatus: allowed values: Unconfigured, Connected, Disconnected
//
// * NewLastConnectionError: allowed values: ERROR_NONE
func (client *WANIPConnection1) GetStatusInfo(
	ctx context.Context,
) (
	NewConnectionStatus string,
	NewLastConnectionError string,
	NewUptime uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewConnectionStatus    string `xml:"NewConnectionStatus"`
		NewLastConnectionError string `xml:"NewLastConnectionError"`
		NewUptime              string `xml:"NewUptime"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetStatusInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewConnectionStatus, err = soap.UnmarshalString(
		response.NewConnectionStatus,
	); err != nil {
		return
	}
	if NewLastConnectionError, err = soap.UnmarshalString(
		response.NewLastConnectionError,
	); err != nil {
		return
	}
	if NewUptime, err = soap.UnmarshalUI4(
		response.NewUptime,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetAutoDisconnectTime implements a UPnP action of the same name.
func (client *WANIPConnection1) GetAutoDisconnectTime(
	ctx context.Context,
) (
	NewAutoDisconnectTime uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewAutoDisconnectTime string `xml:"NewAutoDisconnectTime"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetAutoDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewAutoDisconnectTime, err = soap.UnmarshalUI4(
		response.NewAutoDisconnectTime,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetIdleDisconnectTime implements a UPnP action of the same name.
func (client *WANIPConnection1) GetIdleDisconnectTime(
	ctx context.Context,
) (
	NewIdleDisconnectTime uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewIdleDisconnectTime string `xml:"NewIdleDisconnectTime"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetIdleDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewIdleDisconnectTime, err = soap.UnmarshalUI4(
		response.NewIdleDisconnectTime,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetWarnDisconnectDelay implements a UPnP action of the same name.
func (client *WANIPConnection1) GetWarnDisconnectDelay(
	ctx context.Context,
) (
	NewWarnDisconnectDelay uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewWarnDisconnectDelay string `xml:"NewWarnDisconnectDelay"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetWarnDisconnectDelay",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewWarnDisconnectDelay, err = soap.UnmarshalUI4(
		response.NewWarnDisconnectDelay,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetNATRSIPStatus implements a UPnP action of the same name.
func (client *WANIPConnection1) GetNATRSIPStatus(
	ctx context.Context,
) (
	NewRSIPAvailable bool,
	NewNATEnabled bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewRSIPAvailable string `xml:"NewRSIPAvailable"`
		NewNATEnabled    string `xml:"NewNATEnabled"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetNATRSIPStatus",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewRSIPAvailable, err = soap.UnmarshalBoolean(
		response.NewRSIPAvailable,
	); err != nil {
		return
	}
	if NewNATEnabled, err = soap.UnmarshalBoolean(
		response.NewNATEnabled,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetGenericPortMappingEntry implements a UPnP action of the same name.
//
// Return values:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANIPConnection1) GetGenericPortMappingEntry(
	ctx context.Context,
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
	request := &struct {
		NewPortMappingIndex string `xml:"NewPortMappingIndex"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewPortMappingIndex, err = soap.MarshalUI2(
		NewPortMappingIndex,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewRemoteHost             string `xml:"NewRemoteHost"`
		NewExternalPort           string `xml:"NewExternalPort"`
		NewProtocol               string `xml:"NewProtocol"`
		NewInternalPort           string `xml:"NewInternalPort"`
		NewInternalClient         string `xml:"NewInternalClient"`
		NewEnabled                string `xml:"NewEnabled"`
		NewPortMappingDescription string `xml:"NewPortMappingDescription"`
		NewLeaseDuration          string `xml:"NewLeaseDuration"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetGenericPortMappingEntry",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewRemoteHost, err = soap.UnmarshalString(
		response.NewRemoteHost,
	); err != nil {
		return
	}
	if NewExternalPort, err = soap.UnmarshalUI2(
		response.NewExternalPort,
	); err != nil {
		return
	}
	if NewProtocol, err = soap.UnmarshalString(
		response.NewProtocol,
	); err != nil {
		return
	}
	if NewInternalPort, err = soap.UnmarshalUI2(
		response.NewInternalPort,
	); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(
		response.NewInternalClient,
	); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(
		response.NewEnabled,
	); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(
		response.NewPortMappingDescription,
	); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUI4(
		response.NewLeaseDuration,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSpecificPortMappingEntry implements a UPnP action of the same name.
//
// Parameters:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANIPConnection1) GetSpecificPortMappingEntry(
	ctx context.Context,
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
	request := &struct {
		NewRemoteHost   string `xml:"NewRemoteHost"`
		NewExternalPort string `xml:"NewExternalPort"`
		NewProtocol     string `xml:"NewProtocol"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewRemoteHost, err = soap.MarshalString(
		NewRemoteHost,
	); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUI2(
		NewExternalPort,
	); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(
		NewProtocol,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewInternalPort           string `xml:"NewInternalPort"`
		NewInternalClient         string `xml:"NewInternalClient"`
		NewEnabled                string `xml:"NewEnabled"`
		NewPortMappingDescription string `xml:"NewPortMappingDescription"`
		NewLeaseDuration          string `xml:"NewLeaseDuration"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetSpecificPortMappingEntry",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewInternalPort, err = soap.UnmarshalUI2(
		response.NewInternalPort,
	); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(
		response.NewInternalClient,
	); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(
		response.NewEnabled,
	); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(
		response.NewPortMappingDescription,
	); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUI4(
		response.NewLeaseDuration,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// AddPortMapping implements a UPnP action of the same name.
//
// Parameters:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANIPConnection1) AddPortMapping(
	ctx context.Context,
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
	request := &struct {
		NewRemoteHost             string `xml:"NewRemoteHost"`
		NewExternalPort           string `xml:"NewExternalPort"`
		NewProtocol               string `xml:"NewProtocol"`
		NewInternalPort           string `xml:"NewInternalPort"`
		NewInternalClient         string `xml:"NewInternalClient"`
		NewEnabled                string `xml:"NewEnabled"`
		NewPortMappingDescription string `xml:"NewPortMappingDescription"`
		NewLeaseDuration          string `xml:"NewLeaseDuration"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewRemoteHost, err = soap.MarshalString(
		NewRemoteHost,
	); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUI2(
		NewExternalPort,
	); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(
		NewProtocol,
	); err != nil {
		return
	}
	if request.NewInternalPort, err = soap.MarshalUI2(
		NewInternalPort,
	); err != nil {
		return
	}
	if request.NewInternalClient, err = soap.MarshalString(
		NewInternalClient,
	); err != nil {
		return
	}
	if request.NewEnabled, err = soap.MarshalBoolean(
		NewEnabled,
	); err != nil {
		return
	}
	if request.NewPortMappingDescription, err = soap.MarshalString(
		NewPortMappingDescription,
	); err != nil {
		return
	}
	if request.NewLeaseDuration, err = soap.MarshalUI4(
		NewLeaseDuration,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"AddPortMapping",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeletePortMapping implements a UPnP action of the same name.
//
// Parameters:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANIPConnection1) DeletePortMapping(
	ctx context.Context,
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewRemoteHost   string `xml:"NewRemoteHost"`
		NewExternalPort string `xml:"NewExternalPort"`
		NewProtocol     string `xml:"NewProtocol"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewRemoteHost, err = soap.MarshalString(
		NewRemoteHost,
	); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUI2(
		NewExternalPort,
	); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(
		NewProtocol,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"DeletePortMapping",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetExternalIPAddress implements a UPnP action of the same name.
func (client *WANIPConnection1) GetExternalIPAddress(
	ctx context.Context,
) (
	NewExternalIPAddress string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewExternalIPAddress string `xml:"NewExternalIPAddress"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANIPConnection_1,
		"GetExternalIPAddress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewExternalIPAddress, err = soap.UnmarshalString(
		response.NewExternalIPAddress,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANPOTSLinkConfig1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANPOTSLinkConfig1 struct {
	discover.ServiceClient
}

// NewWANPOTSLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANPOTSLinkConfig1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANPOTSLinkConfig1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANPOTSLinkConfig_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANPOTSLinkConfig1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANPOTSLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANPOTSLinkConfig1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANPOTSLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANPOTSLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANPOTSLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANPOTSLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANPOTSLinkConfig1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANPOTSLinkConfig1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANPOTSLinkConfig_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANPOTSLinkConfig1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANPOTSLinkConfig1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANPOTSLinkConfig1 {
	clients := make([]*WANPOTSLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANPOTSLinkConfig1{genericClients[i]}
	}
	return clients
}

// SetISPInfo implements a UPnP action of the same name.
//
// Parameters:
//
// * NewLinkType: allowed values: PPP_Dialup
func (client *WANPOTSLinkConfig1) SetISPInfo(
	ctx context.Context,
	NewISPPhoneNumber string,
	NewISPInfo string,
	NewLinkType string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewISPPhoneNumber string `xml:"NewISPPhoneNumber"`
		NewISPInfo        string `xml:"NewISPInfo"`
		NewLinkType       string `xml:"NewLinkType"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewISPPhoneNumber, err = soap.MarshalString(
		NewISPPhoneNumber,
	); err != nil {
		return
	}
	if request.NewISPInfo, err = soap.MarshalString(
		NewISPInfo,
	); err != nil {
		return
	}
	if request.NewLinkType, err = soap.MarshalString(
		NewLinkType,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"SetISPInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetCallRetryInfo implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) SetCallRetryInfo(
	ctx context.Context,
	NewNumberOfRetries uint32,
	NewDelayBetweenRetries uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewNumberOfRetries     string `xml:"NewNumberOfRetries"`
		NewDelayBetweenRetries string `xml:"NewDelayBetweenRetries"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewNumberOfRetries, err = soap.MarshalUI4(
		NewNumberOfRetries,
	); err != nil {
		return
	}
	if request.NewDelayBetweenRetries, err = soap.MarshalUI4(
		NewDelayBetweenRetries,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"SetCallRetryInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetISPInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewLinkType: allowed values: PPP_Dialup
func (client *WANPOTSLinkConfig1) GetISPInfo(
	ctx context.Context,
) (
	NewISPPhoneNumber string,
	NewISPInfo string,
	NewLinkType string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewISPPhoneNumber string `xml:"NewISPPhoneNumber"`
		NewISPInfo        string `xml:"NewISPInfo"`
		NewLinkType       string `xml:"NewLinkType"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetISPInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewISPPhoneNumber, err = soap.UnmarshalString(
		response.NewISPPhoneNumber,
	); err != nil {
		return
	}
	if NewISPInfo, err = soap.UnmarshalString(
		response.NewISPInfo,
	); err != nil {
		return
	}
	if NewLinkType, err = soap.UnmarshalString(
		response.NewLinkType,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetCallRetryInfo implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) GetCallRetryInfo(
	ctx context.Context,
) (
	NewNumberOfRetries uint32,
	NewDelayBetweenRetries uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewNumberOfRetries     string `xml:"NewNumberOfRetries"`
		NewDelayBetweenRetries string `xml:"NewDelayBetweenRetries"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetCallRetryInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewNumberOfRetries, err = soap.UnmarshalUI4(
		response.NewNumberOfRetries,
	); err != nil {
		return
	}
	if NewDelayBetweenRetries, err = soap.UnmarshalUI4(
		response.NewDelayBetweenRetries,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetFclass implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) GetFclass(
	ctx context.Context,
) (
	NewFclass string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewFclass string `xml:"NewFclass"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetFclass",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewFclass, err = soap.UnmarshalString(
		response.NewFclass,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDataModulationSupported implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) GetDataModulationSupported(
	ctx context.Context,
) (
	NewDataModulationSupported string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDataModulationSupported string `xml:"NewDataModulationSupported"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetDataModulationSupported",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDataModulationSupported, err = soap.UnmarshalString(
		response.NewDataModulationSupported,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDataProtocol implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) GetDataProtocol(
	ctx context.Context,
) (
	NewDataProtocol string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDataProtocol string `xml:"NewDataProtocol"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetDataProtocol",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDataProtocol, err = soap.UnmarshalString(
		response.NewDataProtocol,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDataCompression implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) GetDataCompression(
	ctx context.Context,
) (
	NewDataCompression string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewDataCompression string `xml:"NewDataCompression"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetDataCompression",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewDataCompression, err = soap.UnmarshalString(
		response.NewDataCompression,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPlusVTRCommandSupported implements a UPnP action of the same name.
func (client *WANPOTSLinkConfig1) GetPlusVTRCommandSupported(
	ctx context.Context,
) (
	NewPlusVTRCommandSupported bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewPlusVTRCommandSupported string `xml:"NewPlusVTRCommandSupported"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPOTSLinkConfig_1,
		"GetPlusVTRCommandSupported",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewPlusVTRCommandSupported, err = soap.UnmarshalBoolean(
		response.NewPlusVTRCommandSupported,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// WANPPPConnection1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:WANPPPConnection:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type WANPPPConnection1 struct {
	discover.ServiceClient
}

// NewWANPPPConnection1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANPPPConnection1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*WANPPPConnection1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_WANPPPConnection_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newWANPPPConnection1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewWANPPPConnection1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANPPPConnection1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*WANPPPConnection1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_WANPPPConnection_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANPPPConnection1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewWANPPPConnection1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANPPPConnection1ClientsFromRootDevice(
	rootDevice *metadata.RootDevice,
	loc *url.URL,
) (
	[]*WANPPPConnection1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_WANPPPConnection_1,
	)
	if err != nil {
		return nil, err
	}
	return newWANPPPConnection1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newWANPPPConnection1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*WANPPPConnection1 {
	clients := make([]*WANPPPConnection1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANPPPConnection1{genericClients[i]}
	}
	return clients
}

// SetConnectionType implements a UPnP action of the same name.
func (client *WANPPPConnection1) SetConnectionType(
	ctx context.Context,
	NewConnectionType string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewConnectionType string `xml:"NewConnectionType"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewConnectionType, err = soap.MarshalString(
		NewConnectionType,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"SetConnectionType",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetConnectionTypeInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewPossibleConnectionTypes: allowed values: Unconfigured, IP_Routed, DHCP_Spoofed, PPPoE_Bridged, PPTP_Relay, L2TP_Relay, PPPoE_Relay
func (client *WANPPPConnection1) GetConnectionTypeInfo(
	ctx context.Context,
) (
	NewConnectionType string,
	NewPossibleConnectionTypes string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewConnectionType          string `xml:"NewConnectionType"`
		NewPossibleConnectionTypes string `xml:"NewPossibleConnectionTypes"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetConnectionTypeInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewConnectionType, err = soap.UnmarshalString(
		response.NewConnectionType,
	); err != nil {
		return
	}
	if NewPossibleConnectionTypes, err = soap.UnmarshalString(
		response.NewPossibleConnectionTypes,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ConfigureConnection implements a UPnP action of the same name.
func (client *WANPPPConnection1) ConfigureConnection(
	ctx context.Context,
	NewUserName string,
	NewPassword string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewUserName string `xml:"NewUserName"`
		NewPassword string `xml:"NewPassword"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewUserName, err = soap.MarshalString(
		NewUserName,
	); err != nil {
		return
	}
	if request.NewPassword, err = soap.MarshalString(
		NewPassword,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"ConfigureConnection",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// RequestConnection implements a UPnP action of the same name.
func (client *WANPPPConnection1) RequestConnection(
	ctx context.Context,
) (
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"RequestConnection",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// RequestTermination implements a UPnP action of the same name.
func (client *WANPPPConnection1) RequestTermination(
	ctx context.Context,
) (
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"RequestTermination",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// ForceTermination implements a UPnP action of the same name.
func (client *WANPPPConnection1) ForceTermination(
	ctx context.Context,
) (
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"ForceTermination",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetAutoDisconnectTime implements a UPnP action of the same name.
func (client *WANPPPConnection1) SetAutoDisconnectTime(
	ctx context.Context,
	NewAutoDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewAutoDisconnectTime string `xml:"NewAutoDisconnectTime"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewAutoDisconnectTime, err = soap.MarshalUI4(
		NewAutoDisconnectTime,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"SetAutoDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetIdleDisconnectTime implements a UPnP action of the same name.
func (client *WANPPPConnection1) SetIdleDisconnectTime(
	ctx context.Context,
	NewIdleDisconnectTime uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewIdleDisconnectTime string `xml:"NewIdleDisconnectTime"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewIdleDisconnectTime, err = soap.MarshalUI4(
		NewIdleDisconnectTime,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"SetIdleDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetWarnDisconnectDelay implements a UPnP action of the same name.
func (client *WANPPPConnection1) SetWarnDisconnectDelay(
	ctx context.Context,
	NewWarnDisconnectDelay uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewWarnDisconnectDelay string `xml:"NewWarnDisconnectDelay"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewWarnDisconnectDelay, err = soap.MarshalUI4(
		NewWarnDisconnectDelay,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"SetWarnDisconnectDelay",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetStatusInfo implements a UPnP action of the same name.
//
// Return values:
//
// * NewConnectionStatus: allowed values: Unconfigured, Connected, Disconnected
//
// * NewLastConnectionError: allowed values: ERROR_NONE
func (client *WANPPPConnection1) GetStatusInfo(
	ctx context.Context,
) (
	NewConnectionStatus string,
	NewLastConnectionError string,
	NewUptime uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewConnectionStatus    string `xml:"NewConnectionStatus"`
		NewLastConnectionError string `xml:"NewLastConnectionError"`
		NewUptime              string `xml:"NewUptime"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetStatusInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewConnectionStatus, err = soap.UnmarshalString(
		response.NewConnectionStatus,
	); err != nil {
		return
	}
	if NewLastConnectionError, err = soap.UnmarshalString(
		response.NewLastConnectionError,
	); err != nil {
		return
	}
	if NewUptime, err = soap.UnmarshalUI4(
		response.NewUptime,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetLinkLayerMaxBitRates implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetLinkLayerMaxBitRates(
	ctx context.Context,
) (
	NewUpstreamMaxBitRate uint32,
	NewDownstreamMaxBitRate uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewUpstreamMaxBitRate   string `xml:"NewUpstreamMaxBitRate"`
		NewDownstreamMaxBitRate string `xml:"NewDownstreamMaxBitRate"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetLinkLayerMaxBitRates",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewUpstreamMaxBitRate, err = soap.UnmarshalUI4(
		response.NewUpstreamMaxBitRate,
	); err != nil {
		return
	}
	if NewDownstreamMaxBitRate, err = soap.UnmarshalUI4(
		response.NewDownstreamMaxBitRate,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPPPEncryptionProtocol implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetPPPEncryptionProtocol(
	ctx context.Context,
) (
	NewPPPEncryptionProtocol string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewPPPEncryptionProtocol string `xml:"NewPPPEncryptionProtocol"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetPPPEncryptionProtocol",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewPPPEncryptionProtocol, err = soap.UnmarshalString(
		response.NewPPPEncryptionProtocol,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPPPCompressionProtocol implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetPPPCompressionProtocol(
	ctx context.Context,
) (
	NewPPPCompressionProtocol string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewPPPCompressionProtocol string `xml:"NewPPPCompressionProtocol"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetPPPCompressionProtocol",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewPPPCompressionProtocol, err = soap.UnmarshalString(
		response.NewPPPCompressionProtocol,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPPPAuthenticationProtocol implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetPPPAuthenticationProtocol(
	ctx context.Context,
) (
	NewPPPAuthenticationProtocol string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewPPPAuthenticationProtocol string `xml:"NewPPPAuthenticationProtocol"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetPPPAuthenticationProtocol",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewPPPAuthenticationProtocol, err = soap.UnmarshalString(
		response.NewPPPAuthenticationProtocol,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetUserName implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetUserName(
	ctx context.Context,
) (
	NewUserName string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewUserName string `xml:"NewUserName"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetUserName",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewUserName, err = soap.UnmarshalString(
		response.NewUserName,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPassword implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetPassword(
	ctx context.Context,
) (
	NewPassword string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewPassword string `xml:"NewPassword"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetPassword",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewPassword, err = soap.UnmarshalString(
		response.NewPassword,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetAutoDisconnectTime implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetAutoDisconnectTime(
	ctx context.Context,
) (
	NewAutoDisconnectTime uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewAutoDisconnectTime string `xml:"NewAutoDisconnectTime"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetAutoDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewAutoDisconnectTime, err = soap.UnmarshalUI4(
		response.NewAutoDisconnectTime,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetIdleDisconnectTime implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetIdleDisconnectTime(
	ctx context.Context,
) (
	NewIdleDisconnectTime uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewIdleDisconnectTime string `xml:"NewIdleDisconnectTime"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetIdleDisconnectTime",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewIdleDisconnectTime, err = soap.UnmarshalUI4(
		response.NewIdleDisconnectTime,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetWarnDisconnectDelay implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetWarnDisconnectDelay(
	ctx context.Context,
) (
	NewWarnDisconnectDelay uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewWarnDisconnectDelay string `xml:"NewWarnDisconnectDelay"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetWarnDisconnectDelay",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewWarnDisconnectDelay, err = soap.UnmarshalUI4(
		response.NewWarnDisconnectDelay,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetNATRSIPStatus implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetNATRSIPStatus(
	ctx context.Context,
) (
	NewRSIPAvailable bool,
	NewNATEnabled bool,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewRSIPAvailable string `xml:"NewRSIPAvailable"`
		NewNATEnabled    string `xml:"NewNATEnabled"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetNATRSIPStatus",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewRSIPAvailable, err = soap.UnmarshalBoolean(
		response.NewRSIPAvailable,
	); err != nil {
		return
	}
	if NewNATEnabled, err = soap.UnmarshalBoolean(
		response.NewNATEnabled,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetGenericPortMappingEntry implements a UPnP action of the same name.
//
// Return values:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANPPPConnection1) GetGenericPortMappingEntry(
	ctx context.Context,
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
	request := &struct {
		NewPortMappingIndex string `xml:"NewPortMappingIndex"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewPortMappingIndex, err = soap.MarshalUI2(
		NewPortMappingIndex,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewRemoteHost             string `xml:"NewRemoteHost"`
		NewExternalPort           string `xml:"NewExternalPort"`
		NewProtocol               string `xml:"NewProtocol"`
		NewInternalPort           string `xml:"NewInternalPort"`
		NewInternalClient         string `xml:"NewInternalClient"`
		NewEnabled                string `xml:"NewEnabled"`
		NewPortMappingDescription string `xml:"NewPortMappingDescription"`
		NewLeaseDuration          string `xml:"NewLeaseDuration"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetGenericPortMappingEntry",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewRemoteHost, err = soap.UnmarshalString(
		response.NewRemoteHost,
	); err != nil {
		return
	}
	if NewExternalPort, err = soap.UnmarshalUI2(
		response.NewExternalPort,
	); err != nil {
		return
	}
	if NewProtocol, err = soap.UnmarshalString(
		response.NewProtocol,
	); err != nil {
		return
	}
	if NewInternalPort, err = soap.UnmarshalUI2(
		response.NewInternalPort,
	); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(
		response.NewInternalClient,
	); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(
		response.NewEnabled,
	); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(
		response.NewPortMappingDescription,
	); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUI4(
		response.NewLeaseDuration,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSpecificPortMappingEntry implements a UPnP action of the same name.
//
// Parameters:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANPPPConnection1) GetSpecificPortMappingEntry(
	ctx context.Context,
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
	request := &struct {
		NewRemoteHost   string `xml:"NewRemoteHost"`
		NewExternalPort string `xml:"NewExternalPort"`
		NewProtocol     string `xml:"NewProtocol"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewRemoteHost, err = soap.MarshalString(
		NewRemoteHost,
	); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUI2(
		NewExternalPort,
	); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(
		NewProtocol,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewInternalPort           string `xml:"NewInternalPort"`
		NewInternalClient         string `xml:"NewInternalClient"`
		NewEnabled                string `xml:"NewEnabled"`
		NewPortMappingDescription string `xml:"NewPortMappingDescription"`
		NewLeaseDuration          string `xml:"NewLeaseDuration"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetSpecificPortMappingEntry",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewInternalPort, err = soap.UnmarshalUI2(
		response.NewInternalPort,
	); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(
		response.NewInternalClient,
	); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(
		response.NewEnabled,
	); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(
		response.NewPortMappingDescription,
	); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUI4(
		response.NewLeaseDuration,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// AddPortMapping implements a UPnP action of the same name.
//
// Parameters:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANPPPConnection1) AddPortMapping(
	ctx context.Context,
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
	request := &struct {
		NewRemoteHost             string `xml:"NewRemoteHost"`
		NewExternalPort           string `xml:"NewExternalPort"`
		NewProtocol               string `xml:"NewProtocol"`
		NewInternalPort           string `xml:"NewInternalPort"`
		NewInternalClient         string `xml:"NewInternalClient"`
		NewEnabled                string `xml:"NewEnabled"`
		NewPortMappingDescription string `xml:"NewPortMappingDescription"`
		NewLeaseDuration          string `xml:"NewLeaseDuration"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewRemoteHost, err = soap.MarshalString(
		NewRemoteHost,
	); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUI2(
		NewExternalPort,
	); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(
		NewProtocol,
	); err != nil {
		return
	}
	if request.NewInternalPort, err = soap.MarshalUI2(
		NewInternalPort,
	); err != nil {
		return
	}
	if request.NewInternalClient, err = soap.MarshalString(
		NewInternalClient,
	); err != nil {
		return
	}
	if request.NewEnabled, err = soap.MarshalBoolean(
		NewEnabled,
	); err != nil {
		return
	}
	if request.NewPortMappingDescription, err = soap.MarshalString(
		NewPortMappingDescription,
	); err != nil {
		return
	}
	if request.NewLeaseDuration, err = soap.MarshalUI4(
		NewLeaseDuration,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"AddPortMapping",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeletePortMapping implements a UPnP action of the same name.
//
// Parameters:
//
// * NewProtocol: allowed values: TCP, UDP
func (client *WANPPPConnection1) DeletePortMapping(
	ctx context.Context,
	NewRemoteHost string,
	NewExternalPort uint16,
	NewProtocol string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		NewRemoteHost   string `xml:"NewRemoteHost"`
		NewExternalPort string `xml:"NewExternalPort"`
		NewProtocol     string `xml:"NewProtocol"`
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.NewRemoteHost, err = soap.MarshalString(
		NewRemoteHost,
	); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUI2(
		NewExternalPort,
	); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(
		NewProtocol,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"DeletePortMapping",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetExternalIPAddress implements a UPnP action of the same name.
func (client *WANPPPConnection1) GetExternalIPAddress(
	ctx context.Context,
) (
	NewExternalIPAddress string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewExternalIPAddress string `xml:"NewExternalIPAddress"`
	}{}

	// Perform the SOAP call.
	if err = client.Client.PerformAction(
		ctx,
		URN_WANPPPConnection_1,
		"GetExternalIPAddress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewExternalIPAddress, err = soap.UnmarshalString(
		response.NewExternalIPAddress,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}
