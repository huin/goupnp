package metadata

import (
	"fmt"
	"net/url"
)

// Service is a service provided by a UPnP Device.
type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceId   string `xml:"serviceId"`
	// SCPDURL is the URL of the SOAP service description.
	SCPDURL URLField `xml:"SCPDURL"`
	// Control URL is the URL of the SOAP endpoint.
	ControlURL  URLField `xml:"controlURL"`
	EventSubURL URLField `xml:"eventSubURL"`
}

// SetURLBase sets the URLBase for the Service.
func (srv *Service) SetURLBase(urlBase *url.URL) {
	srv.SCPDURL.SetURLBase(urlBase)
	srv.ControlURL.SetURLBase(urlBase)
	srv.EventSubURL.SetURLBase(urlBase)
}

func (srv *Service) String() string {
	return fmt.Sprintf("Service ID %s : %s", srv.ServiceId, srv.ServiceType)
}
