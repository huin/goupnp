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
		fmt.Println((indent + 1).String(), srv)
	}
	for _, subdev := range device.Devices {
		displayDevice(indent+1, subdev)
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
				displayDevice(0, maybeRootDevice.Root.Device)
			}
		}
	}
}
