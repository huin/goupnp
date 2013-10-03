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
		fmt.Printf("%sService %s\n", indent+1, srv.ServiceType)
	}
	for i := range device.Devices {
		displayDevice(indent+1, &device.Devices[i])
	}
}

func main() {
	if results, err := goupnp.DiscoverDevices(goupnp.ServiceTypeWANPPPConnection); err != nil {
		fmt.Println("Error discovering InternetGatewayDevice with UPnP: ", err)
	} else {
		fmt.Printf("Discovered %d InternetGatewayDevices:\n", len(results))
		for _, maybeRootDevice := range results {
			if maybeRootDevice.Err != nil {
				fmt.Println(maybeRootDevice.Err)
				continue
			}

			device := &maybeRootDevice.Root.Device

			displayDevice(0, device)
			wanPPPSrvs := device.FindService(goupnp.ServiceTypeWANPPPConnection)
			if len(wanPPPSrvs) < 1 {
				fmt.Printf("Could not find expected service on device %s\n", device.FriendlyName)
				continue
			} else if len(wanPPPSrvs) > 1 {
				fmt.Printf("Got more than one expected service on device %s\n", device.FriendlyName)
			}
			srv := wanPPPSrvs[0]
			results, err := goupnp.PerformSoapAction(goupnp.ServiceTypeWANPPPConnection, "GetExternalIPAddress", &srv.ControlURL.URL, nil)
			if err != nil {
				fmt.Printf("Failed to GetExternalIPAddress from %s: %v\n", device.FriendlyName, err)
				continue
			}
			fmt.Printf("Got GetExternalIPAddress result from %s: %v\n", device.FriendlyName, results)
		}
	}
}
