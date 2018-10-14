// Client for UPnP Device Control Protocol MediaServer v1 and MediaRenderer v1.
//
// This DCP is documented in detail at: http://upnp.org/specs/av/av1/
//
// Typically, use one of the New* functions to create clients for services.
package av1

// ***********************************************************
// GENERATED FILE - DO NOT EDIT BY HAND. See README.md
// ***********************************************************

import (
	"context"
	"net/url"
	"time"

	"github.com/huin/goupnp/v2/discover"
	"github.com/huin/goupnp/v2/soap"
	"github.com/huin/goupnp/v2/ssdp"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const ()

// Service URNs:
const (
	URN_AVTransport_1        = "urn:schemas-upnp-org:service:AVTransport:1"
	URN_AVTransport_2        = "urn:schemas-upnp-org:service:AVTransport:2"
	URN_ConnectionManager_1  = "urn:schemas-upnp-org:service:ConnectionManager:1"
	URN_ConnectionManager_2  = "urn:schemas-upnp-org:service:ConnectionManager:2"
	URN_ContentDirectory_1   = "urn:schemas-upnp-org:service:ContentDirectory:1"
	URN_ContentDirectory_2   = "urn:schemas-upnp-org:service:ContentDirectory:2"
	URN_ContentDirectory_3   = "urn:schemas-upnp-org:service:ContentDirectory:3"
	URN_RenderingControl_1   = "urn:schemas-upnp-org:service:RenderingControl:1"
	URN_RenderingControl_2   = "urn:schemas-upnp-org:service:RenderingControl:2"
	URN_ScheduledRecording_1 = "urn:schemas-upnp-org:service:ScheduledRecording:1"
	URN_ScheduledRecording_2 = "urn:schemas-upnp-org:service:ScheduledRecording:2"
)

// AVTransport1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:AVTransport:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type AVTransport1 struct {
	discover.ServiceClient
}

// NewAVTransport1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*AVTransport1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_AVTransport_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newAVTransport1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewAVTransport1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*AVTransport1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_AVTransport_1,
	)
	if err != nil {
		return nil, err
	}
	return newAVTransport1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewAVTransport1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport1ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*AVTransport1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_AVTransport_1,
	)
	if err != nil {
		return nil, err
	}
	return newAVTransport1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newAVTransport1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*AVTransport1 {
	clients := make([]*AVTransport1, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport1{genericClients[i]}
	}
	return clients
}

// SetAVTransportURI
func (client *AVTransport1) SetAVTransportURI(
	ctx context.Context,
	InstanceID uint32,
	CurrentURI string,
	CurrentURIMetaData string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID         string
		CurrentURI         string
		CurrentURIMetaData string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.CurrentURI, err = soap.MarshalString(
		CurrentURI,
	); err != nil {
		return
	}
	if request.CurrentURIMetaData, err = soap.MarshalString(
		CurrentURIMetaData,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"SetAVTransportURI",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetNextAVTransportURI
func (client *AVTransport1) SetNextAVTransportURI(
	ctx context.Context,
	InstanceID uint32,
	NextURI string,
	NextURIMetaData string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID      string
		NextURI         string
		NextURIMetaData string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.NextURI, err = soap.MarshalString(
		NextURI,
	); err != nil {
		return
	}
	if request.NextURIMetaData, err = soap.MarshalString(
		NextURIMetaData,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"SetNextAVTransportURI",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetMediaInfo
//
// Return values:
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport1) GetMediaInfo(
	ctx context.Context,
	InstanceID uint32,
) (
	NrTracks uint32,
	MediaDuration string,
	CurrentURI string,
	CurrentURIMetaData string,
	NextURI string,
	NextURIMetaData string,
	PlayMedium string,
	RecordMedium string,
	WriteStatus string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NrTracks           string
		MediaDuration      string
		CurrentURI         string
		CurrentURIMetaData string
		NextURI            string
		NextURIMetaData    string
		PlayMedium         string
		RecordMedium       string
		WriteStatus        string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"GetMediaInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NrTracks, err = soap.UnmarshalUi4(
		response.NrTracks,
	); err != nil {
		return
	}
	if MediaDuration, err = soap.UnmarshalString(
		response.MediaDuration,
	); err != nil {
		return
	}
	if CurrentURI, err = soap.UnmarshalString(
		response.CurrentURI,
	); err != nil {
		return
	}
	if CurrentURIMetaData, err = soap.UnmarshalString(
		response.CurrentURIMetaData,
	); err != nil {
		return
	}
	if NextURI, err = soap.UnmarshalString(
		response.NextURI,
	); err != nil {
		return
	}
	if NextURIMetaData, err = soap.UnmarshalString(
		response.NextURIMetaData,
	); err != nil {
		return
	}
	if PlayMedium, err = soap.UnmarshalString(
		response.PlayMedium,
	); err != nil {
		return
	}
	if RecordMedium, err = soap.UnmarshalString(
		response.RecordMedium,
	); err != nil {
		return
	}
	if WriteStatus, err = soap.UnmarshalString(
		response.WriteStatus,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTransportInfo
//
// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PLAYING
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
func (client *AVTransport1) GetTransportInfo(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentTransportState string,
	CurrentTransportStatus string,
	CurrentSpeed string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentTransportState  string
		CurrentTransportStatus string
		CurrentSpeed           string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"GetTransportInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentTransportState, err = soap.UnmarshalString(
		response.CurrentTransportState,
	); err != nil {
		return
	}
	if CurrentTransportStatus, err = soap.UnmarshalString(
		response.CurrentTransportStatus,
	); err != nil {
		return
	}
	if CurrentSpeed, err = soap.UnmarshalString(
		response.CurrentSpeed,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPositionInfo
//
// Return values:
//
// * Track: allowed value range: minimum=0, step=1
func (client *AVTransport1) GetPositionInfo(
	ctx context.Context,
	InstanceID uint32,
) (
	Track uint32,
	TrackDuration string,
	TrackMetaData string,
	TrackURI string,
	RelTime string,
	AbsTime string,
	RelCount int32,
	AbsCount int32,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Track         string
		TrackDuration string
		TrackMetaData string
		TrackURI      string
		RelTime       string
		AbsTime       string
		RelCount      string
		AbsCount      string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"GetPositionInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Track, err = soap.UnmarshalUi4(
		response.Track,
	); err != nil {
		return
	}
	if TrackDuration, err = soap.UnmarshalString(
		response.TrackDuration,
	); err != nil {
		return
	}
	if TrackMetaData, err = soap.UnmarshalString(
		response.TrackMetaData,
	); err != nil {
		return
	}
	if TrackURI, err = soap.UnmarshalString(
		response.TrackURI,
	); err != nil {
		return
	}
	if RelTime, err = soap.UnmarshalString(
		response.RelTime,
	); err != nil {
		return
	}
	if AbsTime, err = soap.UnmarshalString(
		response.AbsTime,
	); err != nil {
		return
	}
	if RelCount, err = soap.UnmarshalI4(
		response.RelCount,
	); err != nil {
		return
	}
	if AbsCount, err = soap.UnmarshalI4(
		response.AbsCount,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDeviceCapabilities
func (client *AVTransport1) GetDeviceCapabilities(
	ctx context.Context,
	InstanceID uint32,
) (
	PlayMedia string,
	RecMedia string,
	RecQualityModes string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PlayMedia       string
		RecMedia        string
		RecQualityModes string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"GetDeviceCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PlayMedia, err = soap.UnmarshalString(
		response.PlayMedia,
	); err != nil {
		return
	}
	if RecMedia, err = soap.UnmarshalString(
		response.RecMedia,
	); err != nil {
		return
	}
	if RecQualityModes, err = soap.UnmarshalString(
		response.RecQualityModes,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTransportSettings
//
// Return values:
//
// * PlayMode: allowed values: NORMAL
func (client *AVTransport1) GetTransportSettings(
	ctx context.Context,
	InstanceID uint32,
) (
	PlayMode string,
	RecQualityMode string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PlayMode       string
		RecQualityMode string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"GetTransportSettings",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PlayMode, err = soap.UnmarshalString(
		response.PlayMode,
	); err != nil {
		return
	}
	if RecQualityMode, err = soap.UnmarshalString(
		response.RecQualityMode,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Stop
func (client *AVTransport1) Stop(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Stop",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Play
//
// Parameters:
//
// * Speed: allowed values: 1
func (client *AVTransport1) Play(
	ctx context.Context,
	InstanceID uint32,
	Speed string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Speed      string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Speed, err = soap.MarshalString(
		Speed,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Play",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Pause
func (client *AVTransport1) Pause(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Pause",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Record
func (client *AVTransport1) Record(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Record",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Seek
//
// Parameters:
//
// * Unit: allowed values: TRACK_NR
func (client *AVTransport1) Seek(
	ctx context.Context,
	InstanceID uint32,
	Unit string,
	Target string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Unit       string
		Target     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Unit, err = soap.MarshalString(
		Unit,
	); err != nil {
		return
	}
	if request.Target, err = soap.MarshalString(
		Target,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Seek",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Next
func (client *AVTransport1) Next(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Next",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Previous
func (client *AVTransport1) Previous(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"Previous",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetPlayMode
//
// Parameters:
//
// * NewPlayMode: allowed values: NORMAL
func (client *AVTransport1) SetPlayMode(
	ctx context.Context,
	InstanceID uint32,
	NewPlayMode string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID  string
		NewPlayMode string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.NewPlayMode, err = soap.MarshalString(
		NewPlayMode,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"SetPlayMode",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetRecordQualityMode
func (client *AVTransport1) SetRecordQualityMode(
	ctx context.Context,
	InstanceID uint32,
	NewRecordQualityMode string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID           string
		NewRecordQualityMode string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.NewRecordQualityMode, err = soap.MarshalString(
		NewRecordQualityMode,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"SetRecordQualityMode",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetCurrentTransportActions
func (client *AVTransport1) GetCurrentTransportActions(
	ctx context.Context,
	InstanceID uint32,
) (
	Actions string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Actions string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_1,
		"GetCurrentTransportActions",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Actions, err = soap.UnmarshalString(
		response.Actions,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// AVTransport2 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:AVTransport:2".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type AVTransport2 struct {
	discover.ServiceClient
}

// NewAVTransport2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport2Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*AVTransport2,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_AVTransport_2,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newAVTransport2ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewAVTransport2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport2ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*AVTransport2,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_AVTransport_2,
	)
	if err != nil {
		return nil, err
	}
	return newAVTransport2ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewAVTransport2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport2ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*AVTransport2,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_AVTransport_2,
	)
	if err != nil {
		return nil, err
	}
	return newAVTransport2ClientsFromGenericClients(
		genericClients,
	), nil
}

func newAVTransport2ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*AVTransport2 {
	clients := make([]*AVTransport2, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport2{genericClients[i]}
	}
	return clients
}

// SetAVTransportURI
func (client *AVTransport2) SetAVTransportURI(
	ctx context.Context,
	InstanceID uint32,
	CurrentURI string,
	CurrentURIMetaData string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID         string
		CurrentURI         string
		CurrentURIMetaData string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.CurrentURI, err = soap.MarshalString(
		CurrentURI,
	); err != nil {
		return
	}
	if request.CurrentURIMetaData, err = soap.MarshalString(
		CurrentURIMetaData,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"SetAVTransportURI",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetNextAVTransportURI
func (client *AVTransport2) SetNextAVTransportURI(
	ctx context.Context,
	InstanceID uint32,
	NextURI string,
	NextURIMetaData string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID      string
		NextURI         string
		NextURIMetaData string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.NextURI, err = soap.MarshalString(
		NextURI,
	); err != nil {
		return
	}
	if request.NextURIMetaData, err = soap.MarshalString(
		NextURIMetaData,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"SetNextAVTransportURI",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetMediaInfo
//
// Return values:
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport2) GetMediaInfo(
	ctx context.Context,
	InstanceID uint32,
) (
	NrTracks uint32,
	MediaDuration string,
	CurrentURI string,
	CurrentURIMetaData string,
	NextURI string,
	NextURIMetaData string,
	PlayMedium string,
	RecordMedium string,
	WriteStatus string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NrTracks           string
		MediaDuration      string
		CurrentURI         string
		CurrentURIMetaData string
		NextURI            string
		NextURIMetaData    string
		PlayMedium         string
		RecordMedium       string
		WriteStatus        string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetMediaInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NrTracks, err = soap.UnmarshalUi4(
		response.NrTracks,
	); err != nil {
		return
	}
	if MediaDuration, err = soap.UnmarshalString(
		response.MediaDuration,
	); err != nil {
		return
	}
	if CurrentURI, err = soap.UnmarshalString(
		response.CurrentURI,
	); err != nil {
		return
	}
	if CurrentURIMetaData, err = soap.UnmarshalString(
		response.CurrentURIMetaData,
	); err != nil {
		return
	}
	if NextURI, err = soap.UnmarshalString(
		response.NextURI,
	); err != nil {
		return
	}
	if NextURIMetaData, err = soap.UnmarshalString(
		response.NextURIMetaData,
	); err != nil {
		return
	}
	if PlayMedium, err = soap.UnmarshalString(
		response.PlayMedium,
	); err != nil {
		return
	}
	if RecordMedium, err = soap.UnmarshalString(
		response.RecordMedium,
	); err != nil {
		return
	}
	if WriteStatus, err = soap.UnmarshalString(
		response.WriteStatus,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetMediaInfo_Ext
//
// Return values:
//
// * CurrentType: allowed values: NO_MEDIA, TRACK_AWARE, TRACK_UNAWARE
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport2) GetMediaInfo_Ext(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentType string,
	NrTracks uint32,
	MediaDuration string,
	CurrentURI string,
	CurrentURIMetaData string,
	NextURI string,
	NextURIMetaData string,
	PlayMedium string,
	RecordMedium string,
	WriteStatus string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentType        string
		NrTracks           string
		MediaDuration      string
		CurrentURI         string
		CurrentURIMetaData string
		NextURI            string
		NextURIMetaData    string
		PlayMedium         string
		RecordMedium       string
		WriteStatus        string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetMediaInfo_Ext",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentType, err = soap.UnmarshalString(
		response.CurrentType,
	); err != nil {
		return
	}
	if NrTracks, err = soap.UnmarshalUi4(
		response.NrTracks,
	); err != nil {
		return
	}
	if MediaDuration, err = soap.UnmarshalString(
		response.MediaDuration,
	); err != nil {
		return
	}
	if CurrentURI, err = soap.UnmarshalString(
		response.CurrentURI,
	); err != nil {
		return
	}
	if CurrentURIMetaData, err = soap.UnmarshalString(
		response.CurrentURIMetaData,
	); err != nil {
		return
	}
	if NextURI, err = soap.UnmarshalString(
		response.NextURI,
	); err != nil {
		return
	}
	if NextURIMetaData, err = soap.UnmarshalString(
		response.NextURIMetaData,
	); err != nil {
		return
	}
	if PlayMedium, err = soap.UnmarshalString(
		response.PlayMedium,
	); err != nil {
		return
	}
	if RecordMedium, err = soap.UnmarshalString(
		response.RecordMedium,
	); err != nil {
		return
	}
	if WriteStatus, err = soap.UnmarshalString(
		response.WriteStatus,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTransportInfo
//
// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PLAYING
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
func (client *AVTransport2) GetTransportInfo(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentTransportState string,
	CurrentTransportStatus string,
	CurrentSpeed string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentTransportState  string
		CurrentTransportStatus string
		CurrentSpeed           string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetTransportInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentTransportState, err = soap.UnmarshalString(
		response.CurrentTransportState,
	); err != nil {
		return
	}
	if CurrentTransportStatus, err = soap.UnmarshalString(
		response.CurrentTransportStatus,
	); err != nil {
		return
	}
	if CurrentSpeed, err = soap.UnmarshalString(
		response.CurrentSpeed,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPositionInfo
//
// Return values:
//
// * Track: allowed value range: minimum=0, step=1
func (client *AVTransport2) GetPositionInfo(
	ctx context.Context,
	InstanceID uint32,
) (
	Track uint32,
	TrackDuration string,
	TrackMetaData string,
	TrackURI string,
	RelTime string,
	AbsTime string,
	RelCount int32,
	AbsCount int32,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Track         string
		TrackDuration string
		TrackMetaData string
		TrackURI      string
		RelTime       string
		AbsTime       string
		RelCount      string
		AbsCount      string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetPositionInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Track, err = soap.UnmarshalUi4(
		response.Track,
	); err != nil {
		return
	}
	if TrackDuration, err = soap.UnmarshalString(
		response.TrackDuration,
	); err != nil {
		return
	}
	if TrackMetaData, err = soap.UnmarshalString(
		response.TrackMetaData,
	); err != nil {
		return
	}
	if TrackURI, err = soap.UnmarshalString(
		response.TrackURI,
	); err != nil {
		return
	}
	if RelTime, err = soap.UnmarshalString(
		response.RelTime,
	); err != nil {
		return
	}
	if AbsTime, err = soap.UnmarshalString(
		response.AbsTime,
	); err != nil {
		return
	}
	if RelCount, err = soap.UnmarshalI4(
		response.RelCount,
	); err != nil {
		return
	}
	if AbsCount, err = soap.UnmarshalI4(
		response.AbsCount,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDeviceCapabilities
func (client *AVTransport2) GetDeviceCapabilities(
	ctx context.Context,
	InstanceID uint32,
) (
	PlayMedia string,
	RecMedia string,
	RecQualityModes string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PlayMedia       string
		RecMedia        string
		RecQualityModes string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetDeviceCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PlayMedia, err = soap.UnmarshalString(
		response.PlayMedia,
	); err != nil {
		return
	}
	if RecMedia, err = soap.UnmarshalString(
		response.RecMedia,
	); err != nil {
		return
	}
	if RecQualityModes, err = soap.UnmarshalString(
		response.RecQualityModes,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetTransportSettings
//
// Return values:
//
// * PlayMode: allowed values: NORMAL
func (client *AVTransport2) GetTransportSettings(
	ctx context.Context,
	InstanceID uint32,
) (
	PlayMode string,
	RecQualityMode string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PlayMode       string
		RecQualityMode string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetTransportSettings",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PlayMode, err = soap.UnmarshalString(
		response.PlayMode,
	); err != nil {
		return
	}
	if RecQualityMode, err = soap.UnmarshalString(
		response.RecQualityMode,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Stop
func (client *AVTransport2) Stop(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Stop",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Play
//
// Parameters:
//
// * Speed: allowed values: 1
func (client *AVTransport2) Play(
	ctx context.Context,
	InstanceID uint32,
	Speed string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Speed      string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Speed, err = soap.MarshalString(
		Speed,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Play",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Pause
func (client *AVTransport2) Pause(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Pause",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Record
func (client *AVTransport2) Record(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Record",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Seek
//
// Parameters:
//
// * Unit: allowed values: TRACK_NR
func (client *AVTransport2) Seek(
	ctx context.Context,
	InstanceID uint32,
	Unit string,
	Target string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Unit       string
		Target     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Unit, err = soap.MarshalString(
		Unit,
	); err != nil {
		return
	}
	if request.Target, err = soap.MarshalString(
		Target,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Seek",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Next
func (client *AVTransport2) Next(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Next",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// Previous
func (client *AVTransport2) Previous(
	ctx context.Context,
	InstanceID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"Previous",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetPlayMode
//
// Parameters:
//
// * NewPlayMode: allowed values: NORMAL
func (client *AVTransport2) SetPlayMode(
	ctx context.Context,
	InstanceID uint32,
	NewPlayMode string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID  string
		NewPlayMode string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.NewPlayMode, err = soap.MarshalString(
		NewPlayMode,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"SetPlayMode",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// SetRecordQualityMode
func (client *AVTransport2) SetRecordQualityMode(
	ctx context.Context,
	InstanceID uint32,
	NewRecordQualityMode string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID           string
		NewRecordQualityMode string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.NewRecordQualityMode, err = soap.MarshalString(
		NewRecordQualityMode,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"SetRecordQualityMode",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetCurrentTransportActions
func (client *AVTransport2) GetCurrentTransportActions(
	ctx context.Context,
	InstanceID uint32,
) (
	Actions string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Actions string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetCurrentTransportActions",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Actions, err = soap.UnmarshalString(
		response.Actions,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetDRMState
//
// Return values:
//
// * CurrentDRMState: allowed values: OK
func (client *AVTransport2) GetDRMState(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentDRMState string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentDRMState string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetDRMState",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentDRMState, err = soap.UnmarshalString(
		response.CurrentDRMState,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetStateVariables
func (client *AVTransport2) GetStateVariables(
	ctx context.Context,
	InstanceID uint32,
	StateVariableList string,
) (
	StateVariableValuePairs string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID        string
		StateVariableList string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.StateVariableList, err = soap.MarshalString(
		StateVariableList,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		StateVariableValuePairs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"GetStateVariables",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if StateVariableValuePairs, err = soap.UnmarshalString(
		response.StateVariableValuePairs,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetStateVariables
func (client *AVTransport2) SetStateVariables(
	ctx context.Context,
	InstanceID uint32,
	AVTransportUDN string,
	ServiceType string,
	ServiceId string,
	StateVariableValuePairs string,
) (
	StateVariableList string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID              string
		AVTransportUDN          string
		ServiceType             string
		ServiceId               string
		StateVariableValuePairs string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.AVTransportUDN, err = soap.MarshalString(
		AVTransportUDN,
	); err != nil {
		return
	}
	if request.ServiceType, err = soap.MarshalString(
		ServiceType,
	); err != nil {
		return
	}
	if request.ServiceId, err = soap.MarshalString(
		ServiceId,
	); err != nil {
		return
	}
	if request.StateVariableValuePairs, err = soap.MarshalString(
		StateVariableValuePairs,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		StateVariableList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_AVTransport_2,
		"SetStateVariables",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if StateVariableList, err = soap.UnmarshalString(
		response.StateVariableList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ConnectionManager1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ConnectionManager:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ConnectionManager1 struct {
	discover.ServiceClient
}

// NewConnectionManager1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ConnectionManager1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ConnectionManager_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newConnectionManager1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewConnectionManager1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ConnectionManager1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ConnectionManager_1,
	)
	if err != nil {
		return nil, err
	}
	return newConnectionManager1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewConnectionManager1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager1ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ConnectionManager1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ConnectionManager_1,
	)
	if err != nil {
		return nil, err
	}
	return newConnectionManager1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newConnectionManager1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ConnectionManager1 {
	clients := make([]*ConnectionManager1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager1{genericClients[i]}
	}
	return clients
}

// GetProtocolInfo
func (client *ConnectionManager1) GetProtocolInfo(
	ctx context.Context,
) (
	Source string,
	Sink string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Source string
		Sink   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_1,
		"GetProtocolInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Source, err = soap.UnmarshalString(
		response.Source,
	); err != nil {
		return
	}
	if Sink, err = soap.UnmarshalString(
		response.Sink,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// PrepareForConnection
//
// Parameters:
//
// * Direction: allowed values: Input, Output
func (client *ConnectionManager1) PrepareForConnection(
	ctx context.Context,
	RemoteProtocolInfo string,
	PeerConnectionManager string,
	PeerConnectionID int32,
	Direction string,
) (
	ConnectionID int32,
	AVTransportID int32,
	RcsID int32,
	err error,
) {
	// Request structure.
	request := &struct {
		RemoteProtocolInfo    string
		PeerConnectionManager string
		PeerConnectionID      string
		Direction             string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RemoteProtocolInfo, err = soap.MarshalString(
		RemoteProtocolInfo,
	); err != nil {
		return
	}
	if request.PeerConnectionManager, err = soap.MarshalString(
		PeerConnectionManager,
	); err != nil {
		return
	}
	if request.PeerConnectionID, err = soap.MarshalI4(
		PeerConnectionID,
	); err != nil {
		return
	}
	if request.Direction, err = soap.MarshalString(
		Direction,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ConnectionID  string
		AVTransportID string
		RcsID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_1,
		"PrepareForConnection",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ConnectionID, err = soap.UnmarshalI4(
		response.ConnectionID,
	); err != nil {
		return
	}
	if AVTransportID, err = soap.UnmarshalI4(
		response.AVTransportID,
	); err != nil {
		return
	}
	if RcsID, err = soap.UnmarshalI4(
		response.RcsID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ConnectionComplete
func (client *ConnectionManager1) ConnectionComplete(
	ctx context.Context,
	ConnectionID int32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ConnectionID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ConnectionID, err = soap.MarshalI4(
		ConnectionID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_1,
		"ConnectionComplete",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetCurrentConnectionIDs
func (client *ConnectionManager1) GetCurrentConnectionIDs(
	ctx context.Context,
) (
	ConnectionIDs string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ConnectionIDs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_1,
		"GetCurrentConnectionIDs",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ConnectionIDs, err = soap.UnmarshalString(
		response.ConnectionIDs,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetCurrentConnectionInfo
//
// Return values:
//
// * Direction: allowed values: Input, Output
//
// * Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
func (client *ConnectionManager1) GetCurrentConnectionInfo(
	ctx context.Context,
	ConnectionID int32,
) (
	RcsID int32,
	AVTransportID int32,
	ProtocolInfo string,
	PeerConnectionManager string,
	PeerConnectionID int32,
	Direction string,
	Status string,
	err error,
) {
	// Request structure.
	request := &struct {
		ConnectionID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ConnectionID, err = soap.MarshalI4(
		ConnectionID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RcsID                 string
		AVTransportID         string
		ProtocolInfo          string
		PeerConnectionManager string
		PeerConnectionID      string
		Direction             string
		Status                string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_1,
		"GetCurrentConnectionInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RcsID, err = soap.UnmarshalI4(
		response.RcsID,
	); err != nil {
		return
	}
	if AVTransportID, err = soap.UnmarshalI4(
		response.AVTransportID,
	); err != nil {
		return
	}
	if ProtocolInfo, err = soap.UnmarshalString(
		response.ProtocolInfo,
	); err != nil {
		return
	}
	if PeerConnectionManager, err = soap.UnmarshalString(
		response.PeerConnectionManager,
	); err != nil {
		return
	}
	if PeerConnectionID, err = soap.UnmarshalI4(
		response.PeerConnectionID,
	); err != nil {
		return
	}
	if Direction, err = soap.UnmarshalString(
		response.Direction,
	); err != nil {
		return
	}
	if Status, err = soap.UnmarshalString(
		response.Status,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ConnectionManager2 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ConnectionManager:2".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ConnectionManager2 struct {
	discover.ServiceClient
}

// NewConnectionManager2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager2Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ConnectionManager2,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ConnectionManager_2,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newConnectionManager2ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewConnectionManager2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager2ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ConnectionManager2,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ConnectionManager_2,
	)
	if err != nil {
		return nil, err
	}
	return newConnectionManager2ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewConnectionManager2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager2ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ConnectionManager2,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ConnectionManager_2,
	)
	if err != nil {
		return nil, err
	}
	return newConnectionManager2ClientsFromGenericClients(
		genericClients,
	), nil
}

func newConnectionManager2ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ConnectionManager2 {
	clients := make([]*ConnectionManager2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager2{genericClients[i]}
	}
	return clients
}

// GetProtocolInfo
func (client *ConnectionManager2) GetProtocolInfo(
	ctx context.Context,
) (
	Source string,
	Sink string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Source string
		Sink   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_2,
		"GetProtocolInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Source, err = soap.UnmarshalString(
		response.Source,
	); err != nil {
		return
	}
	if Sink, err = soap.UnmarshalString(
		response.Sink,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// PrepareForConnection
//
// Parameters:
//
// * Direction: allowed values: Input, Output
func (client *ConnectionManager2) PrepareForConnection(
	ctx context.Context,
	RemoteProtocolInfo string,
	PeerConnectionManager string,
	PeerConnectionID int32,
	Direction string,
) (
	ConnectionID int32,
	AVTransportID int32,
	RcsID int32,
	err error,
) {
	// Request structure.
	request := &struct {
		RemoteProtocolInfo    string
		PeerConnectionManager string
		PeerConnectionID      string
		Direction             string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RemoteProtocolInfo, err = soap.MarshalString(
		RemoteProtocolInfo,
	); err != nil {
		return
	}
	if request.PeerConnectionManager, err = soap.MarshalString(
		PeerConnectionManager,
	); err != nil {
		return
	}
	if request.PeerConnectionID, err = soap.MarshalI4(
		PeerConnectionID,
	); err != nil {
		return
	}
	if request.Direction, err = soap.MarshalString(
		Direction,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ConnectionID  string
		AVTransportID string
		RcsID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_2,
		"PrepareForConnection",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ConnectionID, err = soap.UnmarshalI4(
		response.ConnectionID,
	); err != nil {
		return
	}
	if AVTransportID, err = soap.UnmarshalI4(
		response.AVTransportID,
	); err != nil {
		return
	}
	if RcsID, err = soap.UnmarshalI4(
		response.RcsID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ConnectionComplete
func (client *ConnectionManager2) ConnectionComplete(
	ctx context.Context,
	ConnectionID int32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ConnectionID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ConnectionID, err = soap.MarshalI4(
		ConnectionID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_2,
		"ConnectionComplete",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetCurrentConnectionIDs
func (client *ConnectionManager2) GetCurrentConnectionIDs(
	ctx context.Context,
) (
	ConnectionIDs string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ConnectionIDs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_2,
		"GetCurrentConnectionIDs",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ConnectionIDs, err = soap.UnmarshalString(
		response.ConnectionIDs,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetCurrentConnectionInfo
//
// Return values:
//
// * Direction: allowed values: Input, Output
//
// * Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
func (client *ConnectionManager2) GetCurrentConnectionInfo(
	ctx context.Context,
	ConnectionID int32,
) (
	RcsID int32,
	AVTransportID int32,
	ProtocolInfo string,
	PeerConnectionManager string,
	PeerConnectionID int32,
	Direction string,
	Status string,
	err error,
) {
	// Request structure.
	request := &struct {
		ConnectionID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ConnectionID, err = soap.MarshalI4(
		ConnectionID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RcsID                 string
		AVTransportID         string
		ProtocolInfo          string
		PeerConnectionManager string
		PeerConnectionID      string
		Direction             string
		Status                string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ConnectionManager_2,
		"GetCurrentConnectionInfo",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RcsID, err = soap.UnmarshalI4(
		response.RcsID,
	); err != nil {
		return
	}
	if AVTransportID, err = soap.UnmarshalI4(
		response.AVTransportID,
	); err != nil {
		return
	}
	if ProtocolInfo, err = soap.UnmarshalString(
		response.ProtocolInfo,
	); err != nil {
		return
	}
	if PeerConnectionManager, err = soap.UnmarshalString(
		response.PeerConnectionManager,
	); err != nil {
		return
	}
	if PeerConnectionID, err = soap.UnmarshalI4(
		response.PeerConnectionID,
	); err != nil {
		return
	}
	if Direction, err = soap.UnmarshalString(
		response.Direction,
	); err != nil {
		return
	}
	if Status, err = soap.UnmarshalString(
		response.Status,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ContentDirectory1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ContentDirectory:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ContentDirectory1 struct {
	discover.ServiceClient
}

// NewContentDirectory1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ContentDirectory1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ContentDirectory_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newContentDirectory1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewContentDirectory1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ContentDirectory1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ContentDirectory_1,
	)
	if err != nil {
		return nil, err
	}
	return newContentDirectory1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewContentDirectory1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory1ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ContentDirectory1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ContentDirectory_1,
	)
	if err != nil {
		return nil, err
	}
	return newContentDirectory1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newContentDirectory1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ContentDirectory1 {
	clients := make([]*ContentDirectory1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory1{genericClients[i]}
	}
	return clients
}

// GetSearchCapabilities
func (client *ContentDirectory1) GetSearchCapabilities(
	ctx context.Context,
) (
	SearchCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SearchCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"GetSearchCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SearchCaps, err = soap.UnmarshalString(
		response.SearchCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSortCapabilities
func (client *ContentDirectory1) GetSortCapabilities(
	ctx context.Context,
) (
	SortCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"GetSortCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortCaps, err = soap.UnmarshalString(
		response.SortCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSystemUpdateID
func (client *ContentDirectory1) GetSystemUpdateID(
	ctx context.Context,
) (
	Id uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Id string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"GetSystemUpdateID",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Id, err = soap.UnmarshalUi4(
		response.Id,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Browse
//
// Parameters:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren
func (client *ContentDirectory1) Browse(
	ctx context.Context,
	ObjectID string,
	BrowseFlag string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID       string
		BrowseFlag     string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.BrowseFlag, err = soap.MarshalString(
		BrowseFlag,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"Browse",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Search
func (client *ContentDirectory1) Search(
	ctx context.Context,
	ContainerID string,
	SearchCriteria string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID    string
		SearchCriteria string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.SearchCriteria, err = soap.MarshalString(
		SearchCriteria,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"Search",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateObject
func (client *ContentDirectory1) CreateObject(
	ctx context.Context,
	ContainerID string,
	Elements string,
) (
	ObjectID string,
	Result string,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID string
		Elements    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.Elements, err = soap.MarshalString(
		Elements,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ObjectID string
		Result   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"CreateObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ObjectID, err = soap.UnmarshalString(
		response.ObjectID,
	); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DestroyObject
func (client *ContentDirectory1) DestroyObject(
	ctx context.Context,
	ObjectID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"DestroyObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// UpdateObject
func (client *ContentDirectory1) UpdateObject(
	ctx context.Context,
	ObjectID string,
	CurrentTagValue string,
	NewTagValue string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID        string
		CurrentTagValue string
		NewTagValue     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.CurrentTagValue, err = soap.MarshalString(
		CurrentTagValue,
	); err != nil {
		return
	}
	if request.NewTagValue, err = soap.MarshalString(
		NewTagValue,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"UpdateObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// ImportResource
func (client *ContentDirectory1) ImportResource(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (
	TransferID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.SourceURI, err = soap.MarshalURI(
		SourceURI,
	); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(
		DestinationURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"ImportResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferID, err = soap.UnmarshalUi4(
		response.TransferID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ExportResource
func (client *ContentDirectory1) ExportResource(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (
	TransferID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.SourceURI, err = soap.MarshalURI(
		SourceURI,
	); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(
		DestinationURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"ExportResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferID, err = soap.UnmarshalUi4(
		response.TransferID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// StopTransferResource
func (client *ContentDirectory1) StopTransferResource(
	ctx context.Context,
	TransferID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.TransferID, err = soap.MarshalUi4(
		TransferID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"StopTransferResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetTransferProgress
//
// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory1) GetTransferProgress(
	ctx context.Context,
	TransferID uint32,
) (
	TransferStatus string,
	TransferLength string,
	TransferTotal string,
	err error,
) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.TransferID, err = soap.MarshalUi4(
		TransferID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferStatus string
		TransferLength string
		TransferTotal  string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"GetTransferProgress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferStatus, err = soap.UnmarshalString(
		response.TransferStatus,
	); err != nil {
		return
	}
	if TransferLength, err = soap.UnmarshalString(
		response.TransferLength,
	); err != nil {
		return
	}
	if TransferTotal, err = soap.UnmarshalString(
		response.TransferTotal,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DeleteResource
func (client *ContentDirectory1) DeleteResource(
	ctx context.Context,
	ResourceURI *url.URL,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ResourceURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ResourceURI, err = soap.MarshalURI(
		ResourceURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"DeleteResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// CreateReference
func (client *ContentDirectory1) CreateReference(
	ctx context.Context,
	ContainerID string,
	ObjectID string,
) (
	NewID string,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID string
		ObjectID    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_1,
		"CreateReference",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewID, err = soap.UnmarshalString(
		response.NewID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ContentDirectory2 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ContentDirectory:2".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ContentDirectory2 struct {
	discover.ServiceClient
}

// NewContentDirectory2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory2Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ContentDirectory2,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ContentDirectory_2,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newContentDirectory2ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewContentDirectory2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory2ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ContentDirectory2,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ContentDirectory_2,
	)
	if err != nil {
		return nil, err
	}
	return newContentDirectory2ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewContentDirectory2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory2ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ContentDirectory2,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ContentDirectory_2,
	)
	if err != nil {
		return nil, err
	}
	return newContentDirectory2ClientsFromGenericClients(
		genericClients,
	), nil
}

func newContentDirectory2ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ContentDirectory2 {
	clients := make([]*ContentDirectory2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory2{genericClients[i]}
	}
	return clients
}

// GetSearchCapabilities
func (client *ContentDirectory2) GetSearchCapabilities(
	ctx context.Context,
) (
	SearchCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SearchCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"GetSearchCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SearchCaps, err = soap.UnmarshalString(
		response.SearchCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSortCapabilities
func (client *ContentDirectory2) GetSortCapabilities(
	ctx context.Context,
) (
	SortCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"GetSortCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortCaps, err = soap.UnmarshalString(
		response.SortCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSortExtensionCapabilities
func (client *ContentDirectory2) GetSortExtensionCapabilities(
	ctx context.Context,
) (
	SortExtensionCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortExtensionCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"GetSortExtensionCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortExtensionCaps, err = soap.UnmarshalString(
		response.SortExtensionCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetFeatureList
func (client *ContentDirectory2) GetFeatureList(
	ctx context.Context,
) (
	FeatureList string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		FeatureList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"GetFeatureList",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if FeatureList, err = soap.UnmarshalString(
		response.FeatureList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSystemUpdateID
func (client *ContentDirectory2) GetSystemUpdateID(
	ctx context.Context,
) (
	Id uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Id string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"GetSystemUpdateID",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Id, err = soap.UnmarshalUi4(
		response.Id,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Browse
//
// Parameters:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren
func (client *ContentDirectory2) Browse(
	ctx context.Context,
	ObjectID string,
	BrowseFlag string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID       string
		BrowseFlag     string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.BrowseFlag, err = soap.MarshalString(
		BrowseFlag,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"Browse",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Search
func (client *ContentDirectory2) Search(
	ctx context.Context,
	ContainerID string,
	SearchCriteria string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID    string
		SearchCriteria string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.SearchCriteria, err = soap.MarshalString(
		SearchCriteria,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"Search",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateObject
func (client *ContentDirectory2) CreateObject(
	ctx context.Context,
	ContainerID string,
	Elements string,
) (
	ObjectID string,
	Result string,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID string
		Elements    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.Elements, err = soap.MarshalString(
		Elements,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ObjectID string
		Result   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"CreateObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ObjectID, err = soap.UnmarshalString(
		response.ObjectID,
	); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DestroyObject
func (client *ContentDirectory2) DestroyObject(
	ctx context.Context,
	ObjectID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"DestroyObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// UpdateObject
func (client *ContentDirectory2) UpdateObject(
	ctx context.Context,
	ObjectID string,
	CurrentTagValue string,
	NewTagValue string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID        string
		CurrentTagValue string
		NewTagValue     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.CurrentTagValue, err = soap.MarshalString(
		CurrentTagValue,
	); err != nil {
		return
	}
	if request.NewTagValue, err = soap.MarshalString(
		NewTagValue,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"UpdateObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// MoveObject
func (client *ContentDirectory2) MoveObject(
	ctx context.Context,
	ObjectID string,
	NewParentID string,
) (
	NewObjectID string,
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID    string
		NewParentID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.NewParentID, err = soap.MarshalString(
		NewParentID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewObjectID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"MoveObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewObjectID, err = soap.UnmarshalString(
		response.NewObjectID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ImportResource
func (client *ContentDirectory2) ImportResource(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (
	TransferID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.SourceURI, err = soap.MarshalURI(
		SourceURI,
	); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(
		DestinationURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"ImportResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferID, err = soap.UnmarshalUi4(
		response.TransferID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ExportResource
func (client *ContentDirectory2) ExportResource(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (
	TransferID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.SourceURI, err = soap.MarshalURI(
		SourceURI,
	); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(
		DestinationURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"ExportResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferID, err = soap.UnmarshalUi4(
		response.TransferID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DeleteResource
func (client *ContentDirectory2) DeleteResource(
	ctx context.Context,
	ResourceURI *url.URL,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ResourceURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ResourceURI, err = soap.MarshalURI(
		ResourceURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"DeleteResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// StopTransferResource
func (client *ContentDirectory2) StopTransferResource(
	ctx context.Context,
	TransferID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.TransferID, err = soap.MarshalUi4(
		TransferID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"StopTransferResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetTransferProgress
//
// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory2) GetTransferProgress(
	ctx context.Context,
	TransferID uint32,
) (
	TransferStatus string,
	TransferLength string,
	TransferTotal string,
	err error,
) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.TransferID, err = soap.MarshalUi4(
		TransferID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferStatus string
		TransferLength string
		TransferTotal  string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"GetTransferProgress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferStatus, err = soap.UnmarshalString(
		response.TransferStatus,
	); err != nil {
		return
	}
	if TransferLength, err = soap.UnmarshalString(
		response.TransferLength,
	); err != nil {
		return
	}
	if TransferTotal, err = soap.UnmarshalString(
		response.TransferTotal,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateReference
func (client *ContentDirectory2) CreateReference(
	ctx context.Context,
	ContainerID string,
	ObjectID string,
) (
	NewID string,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID string
		ObjectID    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_2,
		"CreateReference",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewID, err = soap.UnmarshalString(
		response.NewID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ContentDirectory3 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ContentDirectory:3".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ContentDirectory3 struct {
	discover.ServiceClient
}

// NewContentDirectory3Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory3Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ContentDirectory3,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ContentDirectory_3,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newContentDirectory3ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewContentDirectory3ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory3ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ContentDirectory3,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ContentDirectory_3,
	)
	if err != nil {
		return nil, err
	}
	return newContentDirectory3ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewContentDirectory3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory3ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ContentDirectory3,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ContentDirectory_3,
	)
	if err != nil {
		return nil, err
	}
	return newContentDirectory3ClientsFromGenericClients(
		genericClients,
	), nil
}

func newContentDirectory3ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ContentDirectory3 {
	clients := make([]*ContentDirectory3, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory3{genericClients[i]}
	}
	return clients
}

// GetSearchCapabilities
func (client *ContentDirectory3) GetSearchCapabilities(
	ctx context.Context,
) (
	SearchCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SearchCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetSearchCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SearchCaps, err = soap.UnmarshalString(
		response.SearchCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSortCapabilities
func (client *ContentDirectory3) GetSortCapabilities(
	ctx context.Context,
) (
	SortCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetSortCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortCaps, err = soap.UnmarshalString(
		response.SortCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSortExtensionCapabilities
func (client *ContentDirectory3) GetSortExtensionCapabilities(
	ctx context.Context,
) (
	SortExtensionCaps string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortExtensionCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetSortExtensionCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortExtensionCaps, err = soap.UnmarshalString(
		response.SortExtensionCaps,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetFeatureList
func (client *ContentDirectory3) GetFeatureList(
	ctx context.Context,
) (
	FeatureList string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		FeatureList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetFeatureList",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if FeatureList, err = soap.UnmarshalString(
		response.FeatureList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetSystemUpdateID
func (client *ContentDirectory3) GetSystemUpdateID(
	ctx context.Context,
) (
	Id uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Id string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetSystemUpdateID",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Id, err = soap.UnmarshalUi4(
		response.Id,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetServiceResetToken
func (client *ContentDirectory3) GetServiceResetToken(
	ctx context.Context,
) (
	ResetToken string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ResetToken string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetServiceResetToken",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ResetToken, err = soap.UnmarshalString(
		response.ResetToken,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Browse
//
// Parameters:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren
func (client *ContentDirectory3) Browse(
	ctx context.Context,
	ObjectID string,
	BrowseFlag string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID       string
		BrowseFlag     string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.BrowseFlag, err = soap.MarshalString(
		BrowseFlag,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"Browse",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// Search
func (client *ContentDirectory3) Search(
	ctx context.Context,
	ContainerID string,
	SearchCriteria string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID    string
		SearchCriteria string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.SearchCriteria, err = soap.MarshalString(
		SearchCriteria,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"Search",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateObject
func (client *ContentDirectory3) CreateObject(
	ctx context.Context,
	ContainerID string,
	Elements string,
) (
	ObjectID string,
	Result string,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID string
		Elements    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.Elements, err = soap.MarshalString(
		Elements,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		ObjectID string
		Result   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"CreateObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if ObjectID, err = soap.UnmarshalString(
		response.ObjectID,
	); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DestroyObject
func (client *ContentDirectory3) DestroyObject(
	ctx context.Context,
	ObjectID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"DestroyObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// UpdateObject
func (client *ContentDirectory3) UpdateObject(
	ctx context.Context,
	ObjectID string,
	CurrentTagValue string,
	NewTagValue string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID        string
		CurrentTagValue string
		NewTagValue     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.CurrentTagValue, err = soap.MarshalString(
		CurrentTagValue,
	); err != nil {
		return
	}
	if request.NewTagValue, err = soap.MarshalString(
		NewTagValue,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"UpdateObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// MoveObject
func (client *ContentDirectory3) MoveObject(
	ctx context.Context,
	ObjectID string,
	NewParentID string,
) (
	NewObjectID string,
	err error,
) {
	// Request structure.
	request := &struct {
		ObjectID    string
		NewParentID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	if request.NewParentID, err = soap.MarshalString(
		NewParentID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewObjectID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"MoveObject",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewObjectID, err = soap.UnmarshalString(
		response.NewObjectID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ImportResource
func (client *ContentDirectory3) ImportResource(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (
	TransferID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.SourceURI, err = soap.MarshalURI(
		SourceURI,
	); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(
		DestinationURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"ImportResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferID, err = soap.UnmarshalUi4(
		response.TransferID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ExportResource
func (client *ContentDirectory3) ExportResource(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (
	TransferID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.SourceURI, err = soap.MarshalURI(
		SourceURI,
	); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(
		DestinationURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"ExportResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferID, err = soap.UnmarshalUi4(
		response.TransferID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DeleteResource
func (client *ContentDirectory3) DeleteResource(
	ctx context.Context,
	ResourceURI *url.URL,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		ResourceURI string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ResourceURI, err = soap.MarshalURI(
		ResourceURI,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"DeleteResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// StopTransferResource
func (client *ContentDirectory3) StopTransferResource(
	ctx context.Context,
	TransferID uint32,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.TransferID, err = soap.MarshalUi4(
		TransferID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"StopTransferResource",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetTransferProgress
//
// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory3) GetTransferProgress(
	ctx context.Context,
	TransferID uint32,
) (
	TransferStatus string,
	TransferLength string,
	TransferTotal string,
	err error,
) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.TransferID, err = soap.MarshalUi4(
		TransferID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		TransferStatus string
		TransferLength string
		TransferTotal  string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetTransferProgress",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if TransferStatus, err = soap.UnmarshalString(
		response.TransferStatus,
	); err != nil {
		return
	}
	if TransferLength, err = soap.UnmarshalString(
		response.TransferLength,
	); err != nil {
		return
	}
	if TransferTotal, err = soap.UnmarshalString(
		response.TransferTotal,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateReference
func (client *ContentDirectory3) CreateReference(
	ctx context.Context,
	ContainerID string,
	ObjectID string,
) (
	NewID string,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID string
		ObjectID    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.ObjectID, err = soap.MarshalString(
		ObjectID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		NewID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"CreateReference",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if NewID, err = soap.UnmarshalString(
		response.NewID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// FreeFormQuery
func (client *ContentDirectory3) FreeFormQuery(
	ctx context.Context,
	ContainerID string,
	CDSView uint32,
	QueryRequest string,
) (
	QueryResult string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		ContainerID  string
		CDSView      string
		QueryRequest string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.ContainerID, err = soap.MarshalString(
		ContainerID,
	); err != nil {
		return
	}
	if request.CDSView, err = soap.MarshalUi4(
		CDSView,
	); err != nil {
		return
	}
	if request.QueryRequest, err = soap.MarshalString(
		QueryRequest,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		QueryResult string
		UpdateID    string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"FreeFormQuery",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if QueryResult, err = soap.UnmarshalString(
		response.QueryResult,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetFreeFormQueryCapabilities
func (client *ContentDirectory3) GetFreeFormQueryCapabilities(
	ctx context.Context,
) (
	FFQCapabilities string,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		FFQCapabilities string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ContentDirectory_3,
		"GetFreeFormQueryCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if FFQCapabilities, err = soap.UnmarshalString(
		response.FFQCapabilities,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// RenderingControl1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:RenderingControl:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type RenderingControl1 struct {
	discover.ServiceClient
}

// NewRenderingControl1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*RenderingControl1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_RenderingControl_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newRenderingControl1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewRenderingControl1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*RenderingControl1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_RenderingControl_1,
	)
	if err != nil {
		return nil, err
	}
	return newRenderingControl1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewRenderingControl1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl1ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*RenderingControl1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_RenderingControl_1,
	)
	if err != nil {
		return nil, err
	}
	return newRenderingControl1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newRenderingControl1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*RenderingControl1 {
	clients := make([]*RenderingControl1, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl1{genericClients[i]}
	}
	return clients
}

// ListPresets
func (client *RenderingControl1) ListPresets(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentPresetNameList string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentPresetNameList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"ListPresets",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentPresetNameList, err = soap.UnmarshalString(
		response.CurrentPresetNameList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SelectPreset
//
// Parameters:
//
// * PresetName: allowed values: FactoryDefaults
func (client *RenderingControl1) SelectPreset(
	ctx context.Context,
	InstanceID uint32,
	PresetName string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		PresetName string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.PresetName, err = soap.MarshalString(
		PresetName,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SelectPreset",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetBrightness
//
// Return values:
//
// * CurrentBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetBrightness(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentBrightness uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentBrightness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetBrightness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentBrightness, err = soap.UnmarshalUi2(
		response.CurrentBrightness,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetBrightness
//
// Parameters:
//
// * DesiredBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetBrightness(
	ctx context.Context,
	InstanceID uint32,
	DesiredBrightness uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID        string
		DesiredBrightness string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredBrightness, err = soap.MarshalUi2(
		DesiredBrightness,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetBrightness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetContrast
//
// Return values:
//
// * CurrentContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetContrast(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentContrast uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentContrast string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetContrast",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentContrast, err = soap.UnmarshalUi2(
		response.CurrentContrast,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetContrast
//
// Parameters:
//
// * DesiredContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetContrast(
	ctx context.Context,
	InstanceID uint32,
	DesiredContrast uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID      string
		DesiredContrast string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredContrast, err = soap.MarshalUi2(
		DesiredContrast,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetContrast",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetSharpness
//
// Return values:
//
// * CurrentSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetSharpness(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentSharpness uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentSharpness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetSharpness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentSharpness, err = soap.UnmarshalUi2(
		response.CurrentSharpness,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetSharpness
//
// Parameters:
//
// * DesiredSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetSharpness(
	ctx context.Context,
	InstanceID uint32,
	DesiredSharpness uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID       string
		DesiredSharpness string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredSharpness, err = soap.MarshalUi2(
		DesiredSharpness,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetSharpness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRedVideoGain
func (client *RenderingControl1) GetRedVideoGain(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentRedVideoGain uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentRedVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetRedVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentRedVideoGain, err = soap.UnmarshalUi2(
		response.CurrentRedVideoGain,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetRedVideoGain
func (client *RenderingControl1) SetRedVideoGain(
	ctx context.Context,
	InstanceID uint32,
	DesiredRedVideoGain uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID          string
		DesiredRedVideoGain string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredRedVideoGain, err = soap.MarshalUi2(
		DesiredRedVideoGain,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetRedVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetGreenVideoGain
//
// Return values:
//
// * CurrentGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetGreenVideoGain(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentGreenVideoGain uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentGreenVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetGreenVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentGreenVideoGain, err = soap.UnmarshalUi2(
		response.CurrentGreenVideoGain,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetGreenVideoGain
//
// Parameters:
//
// * DesiredGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetGreenVideoGain(
	ctx context.Context,
	InstanceID uint32,
	DesiredGreenVideoGain uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID            string
		DesiredGreenVideoGain string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredGreenVideoGain, err = soap.MarshalUi2(
		DesiredGreenVideoGain,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetGreenVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetBlueVideoGain
//
// Return values:
//
// * CurrentBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetBlueVideoGain(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentBlueVideoGain uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentBlueVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetBlueVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentBlueVideoGain, err = soap.UnmarshalUi2(
		response.CurrentBlueVideoGain,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetBlueVideoGain
//
// Parameters:
//
// * DesiredBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetBlueVideoGain(
	ctx context.Context,
	InstanceID uint32,
	DesiredBlueVideoGain uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID           string
		DesiredBlueVideoGain string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredBlueVideoGain, err = soap.MarshalUi2(
		DesiredBlueVideoGain,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetBlueVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRedVideoBlackLevel
//
// Return values:
//
// * CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetRedVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentRedVideoBlackLevel uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentRedVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetRedVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentRedVideoBlackLevel, err = soap.UnmarshalUi2(
		response.CurrentRedVideoBlackLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetRedVideoBlackLevel
//
// Parameters:
//
// * DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetRedVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
	DesiredRedVideoBlackLevel uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                string
		DesiredRedVideoBlackLevel string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredRedVideoBlackLevel, err = soap.MarshalUi2(
		DesiredRedVideoBlackLevel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetRedVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetGreenVideoBlackLevel
//
// Return values:
//
// * CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetGreenVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentGreenVideoBlackLevel uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentGreenVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetGreenVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentGreenVideoBlackLevel, err = soap.UnmarshalUi2(
		response.CurrentGreenVideoBlackLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetGreenVideoBlackLevel
//
// Parameters:
//
// * DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetGreenVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
	DesiredGreenVideoBlackLevel uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                  string
		DesiredGreenVideoBlackLevel string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredGreenVideoBlackLevel, err = soap.MarshalUi2(
		DesiredGreenVideoBlackLevel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetGreenVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetBlueVideoBlackLevel
//
// Return values:
//
// * CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetBlueVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentBlueVideoBlackLevel uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentBlueVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetBlueVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentBlueVideoBlackLevel, err = soap.UnmarshalUi2(
		response.CurrentBlueVideoBlackLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetBlueVideoBlackLevel
//
// Parameters:
//
// * DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetBlueVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
	DesiredBlueVideoBlackLevel uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                 string
		DesiredBlueVideoBlackLevel string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredBlueVideoBlackLevel, err = soap.MarshalUi2(
		DesiredBlueVideoBlackLevel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetBlueVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetColorTemperature
//
// Return values:
//
// * CurrentColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetColorTemperature(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentColorTemperature uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentColorTemperature string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetColorTemperature",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentColorTemperature, err = soap.UnmarshalUi2(
		response.CurrentColorTemperature,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetColorTemperature
//
// Parameters:
//
// * DesiredColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetColorTemperature(
	ctx context.Context,
	InstanceID uint32,
	DesiredColorTemperature uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID              string
		DesiredColorTemperature string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredColorTemperature, err = soap.MarshalUi2(
		DesiredColorTemperature,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetColorTemperature",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetHorizontalKeystone
//
// Return values:
//
// * CurrentHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl1) GetHorizontalKeystone(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentHorizontalKeystone int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentHorizontalKeystone string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetHorizontalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentHorizontalKeystone, err = soap.UnmarshalI2(
		response.CurrentHorizontalKeystone,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetHorizontalKeystone
//
// Parameters:
//
// * DesiredHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl1) SetHorizontalKeystone(
	ctx context.Context,
	InstanceID uint32,
	DesiredHorizontalKeystone int16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                string
		DesiredHorizontalKeystone string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredHorizontalKeystone, err = soap.MarshalI2(
		DesiredHorizontalKeystone,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetHorizontalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVerticalKeystone
//
// Return values:
//
// * CurrentVerticalKeystone: allowed value range: step=1
func (client *RenderingControl1) GetVerticalKeystone(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentVerticalKeystone int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentVerticalKeystone string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetVerticalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentVerticalKeystone, err = soap.UnmarshalI2(
		response.CurrentVerticalKeystone,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetVerticalKeystone
//
// Parameters:
//
// * DesiredVerticalKeystone: allowed value range: step=1
func (client *RenderingControl1) SetVerticalKeystone(
	ctx context.Context,
	InstanceID uint32,
	DesiredVerticalKeystone int16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID              string
		DesiredVerticalKeystone string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredVerticalKeystone, err = soap.MarshalI2(
		DesiredVerticalKeystone,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetVerticalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetMute
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) GetMute(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentMute bool,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentMute string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetMute",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentMute, err = soap.UnmarshalBoolean(
		response.CurrentMute,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetMute
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) SetMute(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredMute bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID  string
		Channel     string
		DesiredMute string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredMute, err = soap.MarshalBoolean(
		DesiredMute,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetMute",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVolume
//
// Parameters:
//
// * Channel: allowed values: Master
//
// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetVolume(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentVolume uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentVolume string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetVolume",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentVolume, err = soap.UnmarshalUi2(
		response.CurrentVolume,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetVolume
//
// Parameters:
//
// * Channel: allowed values: Master
//
// * DesiredVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl1) SetVolume(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredVolume uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID    string
		Channel       string
		DesiredVolume string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredVolume, err = soap.MarshalUi2(
		DesiredVolume,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetVolume",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVolumeDB
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) GetVolumeDB(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentVolume int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentVolume string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetVolumeDB",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentVolume, err = soap.UnmarshalI2(
		response.CurrentVolume,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetVolumeDB
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) SetVolumeDB(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredVolume int16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID    string
		Channel       string
		DesiredVolume string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredVolume, err = soap.MarshalI2(
		DesiredVolume,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetVolumeDB",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVolumeDBRange
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) GetVolumeDBRange(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	MinValue int16,
	MaxValue int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		MinValue string
		MaxValue string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetVolumeDBRange",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if MinValue, err = soap.UnmarshalI2(
		response.MinValue,
	); err != nil {
		return
	}
	if MaxValue, err = soap.UnmarshalI2(
		response.MaxValue,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetLoudness
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) GetLoudness(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentLoudness bool,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentLoudness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"GetLoudness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentLoudness, err = soap.UnmarshalBoolean(
		response.CurrentLoudness,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetLoudness
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl1) SetLoudness(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredLoudness bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID      string
		Channel         string
		DesiredLoudness string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredLoudness, err = soap.MarshalBoolean(
		DesiredLoudness,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_1,
		"SetLoudness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// RenderingControl2 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:RenderingControl:2".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type RenderingControl2 struct {
	discover.ServiceClient
}

// NewRenderingControl2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl2Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*RenderingControl2,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_RenderingControl_2,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newRenderingControl2ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewRenderingControl2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl2ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*RenderingControl2,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_RenderingControl_2,
	)
	if err != nil {
		return nil, err
	}
	return newRenderingControl2ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewRenderingControl2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl2ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*RenderingControl2,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_RenderingControl_2,
	)
	if err != nil {
		return nil, err
	}
	return newRenderingControl2ClientsFromGenericClients(
		genericClients,
	), nil
}

func newRenderingControl2ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*RenderingControl2 {
	clients := make([]*RenderingControl2, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl2{genericClients[i]}
	}
	return clients
}

// ListPresets
func (client *RenderingControl2) ListPresets(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentPresetNameList string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentPresetNameList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"ListPresets",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentPresetNameList, err = soap.UnmarshalString(
		response.CurrentPresetNameList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SelectPreset
//
// Parameters:
//
// * PresetName: allowed values: FactoryDefaults
func (client *RenderingControl2) SelectPreset(
	ctx context.Context,
	InstanceID uint32,
	PresetName string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		PresetName string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.PresetName, err = soap.MarshalString(
		PresetName,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SelectPreset",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetBrightness
//
// Return values:
//
// * CurrentBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetBrightness(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentBrightness uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentBrightness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetBrightness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentBrightness, err = soap.UnmarshalUi2(
		response.CurrentBrightness,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetBrightness
//
// Parameters:
//
// * DesiredBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetBrightness(
	ctx context.Context,
	InstanceID uint32,
	DesiredBrightness uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID        string
		DesiredBrightness string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredBrightness, err = soap.MarshalUi2(
		DesiredBrightness,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetBrightness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetContrast
//
// Return values:
//
// * CurrentContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetContrast(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentContrast uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentContrast string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetContrast",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentContrast, err = soap.UnmarshalUi2(
		response.CurrentContrast,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetContrast
//
// Parameters:
//
// * DesiredContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetContrast(
	ctx context.Context,
	InstanceID uint32,
	DesiredContrast uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID      string
		DesiredContrast string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredContrast, err = soap.MarshalUi2(
		DesiredContrast,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetContrast",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetSharpness
//
// Return values:
//
// * CurrentSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetSharpness(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentSharpness uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentSharpness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetSharpness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentSharpness, err = soap.UnmarshalUi2(
		response.CurrentSharpness,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetSharpness
//
// Parameters:
//
// * DesiredSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetSharpness(
	ctx context.Context,
	InstanceID uint32,
	DesiredSharpness uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID       string
		DesiredSharpness string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredSharpness, err = soap.MarshalUi2(
		DesiredSharpness,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetSharpness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRedVideoGain
//
// Return values:
//
// * CurrentRedVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetRedVideoGain(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentRedVideoGain uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentRedVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetRedVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentRedVideoGain, err = soap.UnmarshalUi2(
		response.CurrentRedVideoGain,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetRedVideoGain
//
// Parameters:
//
// * DesiredRedVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetRedVideoGain(
	ctx context.Context,
	InstanceID uint32,
	DesiredRedVideoGain uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID          string
		DesiredRedVideoGain string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredRedVideoGain, err = soap.MarshalUi2(
		DesiredRedVideoGain,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetRedVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetGreenVideoGain
//
// Return values:
//
// * CurrentGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetGreenVideoGain(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentGreenVideoGain uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentGreenVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetGreenVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentGreenVideoGain, err = soap.UnmarshalUi2(
		response.CurrentGreenVideoGain,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetGreenVideoGain
//
// Parameters:
//
// * DesiredGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetGreenVideoGain(
	ctx context.Context,
	InstanceID uint32,
	DesiredGreenVideoGain uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID            string
		DesiredGreenVideoGain string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredGreenVideoGain, err = soap.MarshalUi2(
		DesiredGreenVideoGain,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetGreenVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetBlueVideoGain
//
// Return values:
//
// * CurrentBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetBlueVideoGain(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentBlueVideoGain uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentBlueVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetBlueVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentBlueVideoGain, err = soap.UnmarshalUi2(
		response.CurrentBlueVideoGain,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetBlueVideoGain
//
// Parameters:
//
// * DesiredBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetBlueVideoGain(
	ctx context.Context,
	InstanceID uint32,
	DesiredBlueVideoGain uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID           string
		DesiredBlueVideoGain string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredBlueVideoGain, err = soap.MarshalUi2(
		DesiredBlueVideoGain,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetBlueVideoGain",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRedVideoBlackLevel
//
// Return values:
//
// * CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetRedVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentRedVideoBlackLevel uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentRedVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetRedVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentRedVideoBlackLevel, err = soap.UnmarshalUi2(
		response.CurrentRedVideoBlackLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetRedVideoBlackLevel
//
// Parameters:
//
// * DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetRedVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
	DesiredRedVideoBlackLevel uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                string
		DesiredRedVideoBlackLevel string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredRedVideoBlackLevel, err = soap.MarshalUi2(
		DesiredRedVideoBlackLevel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetRedVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetGreenVideoBlackLevel
//
// Return values:
//
// * CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetGreenVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentGreenVideoBlackLevel uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentGreenVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetGreenVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentGreenVideoBlackLevel, err = soap.UnmarshalUi2(
		response.CurrentGreenVideoBlackLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetGreenVideoBlackLevel
//
// Parameters:
//
// * DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetGreenVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
	DesiredGreenVideoBlackLevel uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                  string
		DesiredGreenVideoBlackLevel string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredGreenVideoBlackLevel, err = soap.MarshalUi2(
		DesiredGreenVideoBlackLevel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetGreenVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetBlueVideoBlackLevel
//
// Return values:
//
// * CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetBlueVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentBlueVideoBlackLevel uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentBlueVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetBlueVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentBlueVideoBlackLevel, err = soap.UnmarshalUi2(
		response.CurrentBlueVideoBlackLevel,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetBlueVideoBlackLevel
//
// Parameters:
//
// * DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetBlueVideoBlackLevel(
	ctx context.Context,
	InstanceID uint32,
	DesiredBlueVideoBlackLevel uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                 string
		DesiredBlueVideoBlackLevel string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredBlueVideoBlackLevel, err = soap.MarshalUi2(
		DesiredBlueVideoBlackLevel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetBlueVideoBlackLevel",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetColorTemperature
//
// Return values:
//
// * CurrentColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetColorTemperature(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentColorTemperature uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentColorTemperature string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetColorTemperature",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentColorTemperature, err = soap.UnmarshalUi2(
		response.CurrentColorTemperature,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetColorTemperature
//
// Parameters:
//
// * DesiredColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetColorTemperature(
	ctx context.Context,
	InstanceID uint32,
	DesiredColorTemperature uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID              string
		DesiredColorTemperature string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredColorTemperature, err = soap.MarshalUi2(
		DesiredColorTemperature,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetColorTemperature",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetHorizontalKeystone
//
// Return values:
//
// * CurrentHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl2) GetHorizontalKeystone(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentHorizontalKeystone int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentHorizontalKeystone string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetHorizontalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentHorizontalKeystone, err = soap.UnmarshalI2(
		response.CurrentHorizontalKeystone,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetHorizontalKeystone
//
// Parameters:
//
// * DesiredHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl2) SetHorizontalKeystone(
	ctx context.Context,
	InstanceID uint32,
	DesiredHorizontalKeystone int16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID                string
		DesiredHorizontalKeystone string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredHorizontalKeystone, err = soap.MarshalI2(
		DesiredHorizontalKeystone,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetHorizontalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVerticalKeystone
//
// Return values:
//
// * CurrentVerticalKeystone: allowed value range: step=1
func (client *RenderingControl2) GetVerticalKeystone(
	ctx context.Context,
	InstanceID uint32,
) (
	CurrentVerticalKeystone int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentVerticalKeystone string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetVerticalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentVerticalKeystone, err = soap.UnmarshalI2(
		response.CurrentVerticalKeystone,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetVerticalKeystone
//
// Parameters:
//
// * DesiredVerticalKeystone: allowed value range: step=1
func (client *RenderingControl2) SetVerticalKeystone(
	ctx context.Context,
	InstanceID uint32,
	DesiredVerticalKeystone int16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID              string
		DesiredVerticalKeystone string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.DesiredVerticalKeystone, err = soap.MarshalI2(
		DesiredVerticalKeystone,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetVerticalKeystone",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetMute
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) GetMute(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentMute bool,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentMute string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetMute",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentMute, err = soap.UnmarshalBoolean(
		response.CurrentMute,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetMute
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) SetMute(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredMute bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID  string
		Channel     string
		DesiredMute string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredMute, err = soap.MarshalBoolean(
		DesiredMute,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetMute",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVolume
//
// Parameters:
//
// * Channel: allowed values: Master
//
// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetVolume(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentVolume uint16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentVolume string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetVolume",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentVolume, err = soap.UnmarshalUi2(
		response.CurrentVolume,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetVolume
//
// Parameters:
//
// * Channel: allowed values: Master
//
// * DesiredVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl2) SetVolume(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredVolume uint16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID    string
		Channel       string
		DesiredVolume string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredVolume, err = soap.MarshalUi2(
		DesiredVolume,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetVolume",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVolumeDB
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) GetVolumeDB(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentVolume int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentVolume string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetVolumeDB",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentVolume, err = soap.UnmarshalI2(
		response.CurrentVolume,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetVolumeDB
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) SetVolumeDB(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredVolume int16,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID    string
		Channel       string
		DesiredVolume string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredVolume, err = soap.MarshalI2(
		DesiredVolume,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetVolumeDB",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetVolumeDBRange
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) GetVolumeDBRange(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	MinValue int16,
	MaxValue int16,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		MinValue string
		MaxValue string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetVolumeDBRange",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if MinValue, err = soap.UnmarshalI2(
		response.MinValue,
	); err != nil {
		return
	}
	if MaxValue, err = soap.UnmarshalI2(
		response.MaxValue,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetLoudness
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) GetLoudness(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (
	CurrentLoudness bool,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		CurrentLoudness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetLoudness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if CurrentLoudness, err = soap.UnmarshalBoolean(
		response.CurrentLoudness,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetLoudness
//
// Parameters:
//
// * Channel: allowed values: Master
func (client *RenderingControl2) SetLoudness(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredLoudness bool,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID      string
		Channel         string
		DesiredLoudness string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(
		Channel,
	); err != nil {
		return
	}
	if request.DesiredLoudness, err = soap.MarshalBoolean(
		DesiredLoudness,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetLoudness",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetStateVariables
func (client *RenderingControl2) GetStateVariables(
	ctx context.Context,
	InstanceID uint32,
	StateVariableList string,
) (
	StateVariableValuePairs string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID        string
		StateVariableList string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.StateVariableList, err = soap.MarshalString(
		StateVariableList,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		StateVariableValuePairs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"GetStateVariables",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if StateVariableValuePairs, err = soap.UnmarshalString(
		response.StateVariableValuePairs,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// SetStateVariables
func (client *RenderingControl2) SetStateVariables(
	ctx context.Context,
	InstanceID uint32,
	RenderingControlUDN string,
	ServiceType string,
	ServiceId string,
	StateVariableValuePairs string,
) (
	StateVariableList string,
	err error,
) {
	// Request structure.
	request := &struct {
		InstanceID              string
		RenderingControlUDN     string
		ServiceType             string
		ServiceId               string
		StateVariableValuePairs string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.InstanceID, err = soap.MarshalUi4(
		InstanceID,
	); err != nil {
		return
	}
	if request.RenderingControlUDN, err = soap.MarshalString(
		RenderingControlUDN,
	); err != nil {
		return
	}
	if request.ServiceType, err = soap.MarshalString(
		ServiceType,
	); err != nil {
		return
	}
	if request.ServiceId, err = soap.MarshalString(
		ServiceId,
	); err != nil {
		return
	}
	if request.StateVariableValuePairs, err = soap.MarshalString(
		StateVariableValuePairs,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		StateVariableList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_RenderingControl_2,
		"SetStateVariables",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if StateVariableList, err = soap.UnmarshalString(
		response.StateVariableList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ScheduledRecording1 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ScheduledRecording:1".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ScheduledRecording1 struct {
	discover.ServiceClient
}

// NewScheduledRecording1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewScheduledRecording1Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ScheduledRecording1,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ScheduledRecording_1,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newScheduledRecording1ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewScheduledRecording1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewScheduledRecording1ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ScheduledRecording1,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ScheduledRecording_1,
	)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording1ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewScheduledRecording1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewScheduledRecording1ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ScheduledRecording1,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ScheduledRecording_1,
	)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording1ClientsFromGenericClients(
		genericClients,
	), nil
}

func newScheduledRecording1ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ScheduledRecording1 {
	clients := make([]*ScheduledRecording1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ScheduledRecording1{genericClients[i]}
	}
	return clients
}

// GetSortCapabilities
func (client *ScheduledRecording1) GetSortCapabilities(
	ctx context.Context,
) (
	SortCaps string,
	SortLevelCap uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortCaps     string
		SortLevelCap string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetSortCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortCaps, err = soap.UnmarshalString(
		response.SortCaps,
	); err != nil {
		return
	}
	if SortLevelCap, err = soap.UnmarshalUi4(
		response.SortLevelCap,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPropertyList
//
// Parameters:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
func (client *ScheduledRecording1) GetPropertyList(
	ctx context.Context,
	DataTypeID string,
) (
	PropertyList string,
	err error,
) {
	// Request structure.
	request := &struct {
		DataTypeID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.DataTypeID, err = soap.MarshalString(
		DataTypeID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PropertyList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetPropertyList",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PropertyList, err = soap.UnmarshalString(
		response.PropertyList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetAllowedValues
//
// Parameters:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
func (client *ScheduledRecording1) GetAllowedValues(
	ctx context.Context,
	DataTypeID string,
	Filter string,
) (
	PropertyInfo string,
	err error,
) {
	// Request structure.
	request := &struct {
		DataTypeID string
		Filter     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.DataTypeID, err = soap.MarshalString(
		DataTypeID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PropertyInfo string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetAllowedValues",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PropertyInfo, err = soap.UnmarshalString(
		response.PropertyInfo,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetStateUpdateID
func (client *ScheduledRecording1) GetStateUpdateID(
	ctx context.Context,
) (
	Id uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Id string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetStateUpdateID",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Id, err = soap.UnmarshalUi4(
		response.Id,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// BrowseRecordSchedules
func (client *ScheduledRecording1) BrowseRecordSchedules(
	ctx context.Context,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"BrowseRecordSchedules",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// BrowseRecordTasks
func (client *ScheduledRecording1) BrowseRecordTasks(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
		StartingIndex    string
		RequestedCount   string
		SortCriteria     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"BrowseRecordTasks",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateRecordSchedule
func (client *ScheduledRecording1) CreateRecordSchedule(
	ctx context.Context,
	Elements string,
) (
	RecordScheduleID string,
	Result string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		Elements string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.Elements, err = soap.MarshalString(
		Elements,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RecordScheduleID string
		Result           string
		UpdateID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"CreateRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RecordScheduleID, err = soap.UnmarshalString(
		response.RecordScheduleID,
	); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DeleteRecordSchedule
func (client *ScheduledRecording1) DeleteRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"DeleteRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordSchedule
func (client *ScheduledRecording1) GetRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
) (
	Result string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// EnableRecordSchedule
func (client *ScheduledRecording1) EnableRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"EnableRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DisableRecordSchedule
func (client *ScheduledRecording1) DisableRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"DisableRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeleteRecordTask
func (client *ScheduledRecording1) DeleteRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"DeleteRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordTask
func (client *ScheduledRecording1) GetRecordTask(
	ctx context.Context,
	RecordTaskID string,
	Filter string,
) (
	Result string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
		Filter       string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// EnableRecordTask
func (client *ScheduledRecording1) EnableRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"EnableRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DisableRecordTask
func (client *ScheduledRecording1) DisableRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"DisableRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// ResetRecordTask
func (client *ScheduledRecording1) ResetRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"ResetRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordScheduleConflicts
func (client *ScheduledRecording1) GetRecordScheduleConflicts(
	ctx context.Context,
	RecordScheduleID string,
) (
	RecordScheduleConflictIDList string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RecordScheduleConflictIDList string
		UpdateID                     string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetRecordScheduleConflicts",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RecordScheduleConflictIDList, err = soap.UnmarshalString(
		response.RecordScheduleConflictIDList,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordTaskConflicts
func (client *ScheduledRecording1) GetRecordTaskConflicts(
	ctx context.Context,
	RecordTaskID string,
) (
	RecordTaskConflictIDList string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RecordTaskConflictIDList string
		UpdateID                 string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_1,
		"GetRecordTaskConflicts",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RecordTaskConflictIDList, err = soap.UnmarshalString(
		response.RecordTaskConflictIDList,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// ScheduledRecording2 is a client for UPnP SOAP service with URN
// "urn:schemas-upnp-org:service:ScheduledRecording:2".
// See discover.ServiceClient, which contains RootDevice and Service attributes
// which are provided for informational value.
type ScheduledRecording2 struct {
	discover.ServiceClient
}

// NewScheduledRecording2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewScheduledRecording2Clients(
	ctx context.Context,
	searchOpts ...ssdp.SearchOption,
) (
	clients []*ScheduledRecording2,
	errors []error, err error,
) {
	var genericClients []discover.ServiceClient
	if genericClients, errors, err = discover.NewServiceClients(
		ctx,
		URN_ScheduledRecording_2,
		searchOpts...,
	); err != nil {
		return
	}
	clients = newScheduledRecording2ClientsFromGenericClients(
		genericClients,
	)
	return
}

// NewScheduledRecording2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewScheduledRecording2ClientsByURL(
	ctx context.Context,
	loc *url.URL,
) (
	[]*ScheduledRecording2,
	error,
) {
	genericClients, err := discover.NewServiceClientsByURL(
		ctx,
		loc,
		URN_ScheduledRecording_2,
	)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording2ClientsFromGenericClients(
		genericClients,
	), nil
}

// NewScheduledRecording2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewScheduledRecording2ClientsFromRootDevice(
	rootDevice *discover.RootDevice,
	loc *url.URL,
) (
	[]*ScheduledRecording2,
	error,
) {
	genericClients, err := discover.NewServiceClientsFromRootDevice(
		rootDevice,
		loc,
		URN_ScheduledRecording_2,
	)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording2ClientsFromGenericClients(
		genericClients,
	), nil
}

func newScheduledRecording2ClientsFromGenericClients(
	genericClients []discover.ServiceClient,
) []*ScheduledRecording2 {
	clients := make([]*ScheduledRecording2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ScheduledRecording2{genericClients[i]}
	}
	return clients
}

// GetSortCapabilities
func (client *ScheduledRecording2) GetSortCapabilities(
	ctx context.Context,
) (
	SortCaps string,
	SortLevelCap uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		SortCaps     string
		SortLevelCap string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetSortCapabilities",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if SortCaps, err = soap.UnmarshalString(
		response.SortCaps,
	); err != nil {
		return
	}
	if SortLevelCap, err = soap.UnmarshalUi4(
		response.SortLevelCap,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetPropertyList
//
// Parameters:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
func (client *ScheduledRecording2) GetPropertyList(
	ctx context.Context,
	DataTypeID string,
) (
	PropertyList string,
	err error,
) {
	// Request structure.
	request := &struct {
		DataTypeID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.DataTypeID, err = soap.MarshalString(
		DataTypeID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PropertyList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetPropertyList",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PropertyList, err = soap.UnmarshalString(
		response.PropertyList,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetAllowedValues
//
// Parameters:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
func (client *ScheduledRecording2) GetAllowedValues(
	ctx context.Context,
	DataTypeID string,
	Filter string,
) (
	PropertyInfo string,
	err error,
) {
	// Request structure.
	request := &struct {
		DataTypeID string
		Filter     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.DataTypeID, err = soap.MarshalString(
		DataTypeID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		PropertyInfo string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetAllowedValues",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if PropertyInfo, err = soap.UnmarshalString(
		response.PropertyInfo,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetStateUpdateID
func (client *ScheduledRecording2) GetStateUpdateID(
	ctx context.Context,
) (
	Id uint32,
	err error,
) {
	// Request structure.
	request := interface{}(nil)

	// BEGIN Marshal arguments into request struct.
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Id string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetStateUpdateID",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Id, err = soap.UnmarshalUi4(
		response.Id,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// BrowseRecordSchedules
func (client *ScheduledRecording2) BrowseRecordSchedules(
	ctx context.Context,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"BrowseRecordSchedules",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// BrowseRecordTasks
func (client *ScheduledRecording2) BrowseRecordTasks(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (
	Result string,
	NumberReturned uint32,
	TotalMatches uint32,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
		StartingIndex    string
		RequestedCount   string
		SortCriteria     string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(
		StartingIndex,
	); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(
		RequestedCount,
	); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(
		SortCriteria,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"BrowseRecordTasks",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(
		response.NumberReturned,
	); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(
		response.TotalMatches,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// CreateRecordSchedule
func (client *ScheduledRecording2) CreateRecordSchedule(
	ctx context.Context,
	Elements string,
) (
	RecordScheduleID string,
	Result string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		Elements string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.Elements, err = soap.MarshalString(
		Elements,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RecordScheduleID string
		Result           string
		UpdateID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"CreateRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RecordScheduleID, err = soap.UnmarshalString(
		response.RecordScheduleID,
	); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// DeleteRecordSchedule
func (client *ScheduledRecording2) DeleteRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"DeleteRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordSchedule
func (client *ScheduledRecording2) GetRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
) (
	Result string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// EnableRecordSchedule
func (client *ScheduledRecording2) EnableRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"EnableRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DisableRecordSchedule
func (client *ScheduledRecording2) DisableRecordSchedule(
	ctx context.Context,
	RecordScheduleID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"DisableRecordSchedule",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DeleteRecordTask
func (client *ScheduledRecording2) DeleteRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"DeleteRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordTask
func (client *ScheduledRecording2) GetRecordTask(
	ctx context.Context,
	RecordTaskID string,
	Filter string,
) (
	Result string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
		Filter       string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(
		Filter,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if Result, err = soap.UnmarshalString(
		response.Result,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// EnableRecordTask
func (client *ScheduledRecording2) EnableRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"EnableRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// DisableRecordTask
func (client *ScheduledRecording2) DisableRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"DisableRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// ResetRecordTask
func (client *ScheduledRecording2) ResetRecordTask(
	ctx context.Context,
	RecordTaskID string,
) (
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"ResetRecordTask",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordScheduleConflicts
func (client *ScheduledRecording2) GetRecordScheduleConflicts(
	ctx context.Context,
	RecordScheduleID string,
) (
	RecordScheduleConflictIDList string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordScheduleID, err = soap.MarshalString(
		RecordScheduleID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RecordScheduleConflictIDList string
		UpdateID                     string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetRecordScheduleConflicts",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RecordScheduleConflictIDList, err = soap.UnmarshalString(
		response.RecordScheduleConflictIDList,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}

// GetRecordTaskConflicts
func (client *ScheduledRecording2) GetRecordTaskConflicts(
	ctx context.Context,
	RecordTaskID string,
) (
	RecordTaskConflictIDList string,
	UpdateID uint32,
	err error,
) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}

	// BEGIN Marshal arguments into request struct.
	if request.RecordTaskID, err = soap.MarshalString(
		RecordTaskID,
	); err != nil {
		return
	}
	// END Marshal arguments into request struct.

	// Response structure.
	response := &struct {
		RecordTaskConflictIDList string
		UpdateID                 string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(
		ctx,
		URN_ScheduledRecording_2,
		"GetRecordTaskConflicts",
		request,
		response,
	); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response struct.
	if RecordTaskConflictIDList, err = soap.UnmarshalString(
		response.RecordTaskConflictIDList,
	); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(
		response.UpdateID,
	); err != nil {
		return
	}
	// END Unmarshal arguments from response struct.

	return
}
