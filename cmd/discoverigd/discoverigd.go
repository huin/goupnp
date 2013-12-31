// Serves as a simple example/test of discovering UPnP devices on the local
// network.
package main

import (
	"fmt"

	"github.com/huin/goupnp/dcps/internetgateway1"
)

func main() {
	clients, errors, err := internetgateway1.NewWANPPPConnection1Clients()
	if err != nil {
		fmt.Println("Error discovering service with UPnP: ", err)
		return
	}

	if len(errors) > 0 {
		fmt.Printf("Error discovering %d services:\n", len(errors))
		for _, err := range errors {
			fmt.Println("  ", err)
		}
	}

	fmt.Printf("Successfully discovered %d services:\n", len(clients))
	for _, client := range clients {
		device := &client.RootDevice.Device

		fmt.Println("  Device:", device.FriendlyName)
		if addr, err := client.GetExternalIPAddress(); err != nil {
			fmt.Printf("    Failed to get external IP address: %v\n", err)
		} else {
			fmt.Printf("    External IP address: %v\n", addr)
		}
	}
}
