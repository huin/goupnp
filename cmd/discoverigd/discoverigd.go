// Serves as a simple example/test of discovering UPnP devices on the local
// network.
package main

import (
	"fmt"

	"github.com/huin/goupnp"
)

var spaces = "                                     "

type indentLevel int

func (i indentLevel) String() string {
	return spaces[:i]
}

func displayDevice(indent indentLevel, device *goupnp.Device) {
	fmt.Println(indent.String(), device)
	for _, srv := range device.Services {
		fmt.Println((indent + 1).String(), srv, srv.SCPDURL.URL.String(), srv.ControlURL.URL.String())
		fmt.Println(goupnp.ServiceTypeWANPPPConnection, srv.ServiceType)
		if srv.ServiceType == goupnp.ServiceTypeWANPPPConnection {
			results, err := goupnp.PerformSoapAction(goupnp.ServiceTypeWANPPPConnection, "GetExternalIPAddress", &srv.ControlURL.URL, nil)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(results)
			}
		}
	}
	for i := range device.Devices {
		displayDevice(indent+1, &device.Devices[i])
	}
}

func main() {
	if results, err := goupnp.DiscoverDevices(goupnp.DeviceTypeInternetGatewayDevice); err != nil {
		fmt.Println("Error discovering InternetGatewayDevice with UPnP: ", err)
	} else {
		fmt.Printf("Discovered %d InternetGatewayDevices:\n", len(results))
		for _, maybeRootDevice := range results {
			if maybeRootDevice.Err != nil {
				fmt.Println(maybeRootDevice.Err)
			} else {
				displayDevice(0, &maybeRootDevice.Root.Device)
			}
		}
	}
}
