// Serves as a simple example/test of discovering UPnP devices on the local
// network.
package main

import (
	"fmt"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
)

func main() {
	results, err := goupnp.DiscoverDevices(internetgateway1.URN_WANPPPConnection_1)
	if err != nil {
		fmt.Println("Error discovering InternetGatewayDevice with UPnP: ", err)
		return
	}

	fmt.Printf("Discovered %d InternetGatewayDevices:\n", len(results))
	for _, maybeRootDevice := range results {
		if maybeRootDevice.Err != nil {
			fmt.Println(maybeRootDevice.Err)
			continue
		}

		device := &maybeRootDevice.Root.Device

		fmt.Println("Device ", device.FriendlyName)
		wanPPPSrvs := device.FindService(internetgateway1.URN_WANPPPConnection_1)
		if len(wanPPPSrvs) < 1 {
			fmt.Printf("Could not find expected service on device %s\n", device.FriendlyName)
			continue
		}

		for _, srv := range wanPPPSrvs {
			client := internetgateway1.WANPPPConnection1{srv.NewSOAPClient()}
			if addr, err := client.GetExternalIPAddress(); err != nil {
				fmt.Printf("Failed to get external IP address: %v\n", err)
			} else {
				fmt.Printf("External IP address: %v\n", addr)
			}
		}
	}
}
