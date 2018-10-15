package metadata

import (
	"fmt"
	"net/url"
)

// Service is a service provided by a UPnP Device.
type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceID   string `xml:"serviceId"`
	// SCPDURL is the URL of the SOAP service description.
	SCPDURL URLField `xml:"SCPDURL"`
	// Control URL is the URL of the SOAP endpoint.
	ControlURL  URLField `xml:"controlURL"`
	EventSubURL URLField `xml:"eventSubURL"`
}

// setURLBase sets the URLBase for the Service.
func (srv *Service) setURLBase(urlBase *url.URL) {
	srv.SCPDURL.setURLBase(urlBase)
	srv.ControlURL.setURLBase(urlBase)
	srv.EventSubURL.setURLBase(urlBase)
}

func (srv *Service) String() string {
	return fmt.Sprintf("Service ID %s : %s", srv.ServiceID, srv.ServiceType)
}
