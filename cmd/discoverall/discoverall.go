// Example program to display all devices discovered on the local network.
package main

import (
	"fmt"
	"log"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/ssdp"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	devices, err := goupnp.DiscoverDevices(ssdp.SSDPAll)
	if err != nil {
		return err
	}
	for _, device := range devices {
		fmt.Printf("Location: %v\n", device.Location)
		fmt.Printf("USN: %v\n", device.USN)
		if device.Err != nil {
			fmt.Printf("  Error: %v\n", device.Err)
		} else {
			fmt.Printf("  Root v%d.%d @ %s\n",
				device.Root.SpecVersion.Major, device.Root.SpecVersion.Minor,
				device.Root.URLBaseStr)
			fmt.Printf("  Type: %s\n", device.Root.Device.DeviceType)
			fmt.Printf("  Friendly name: %s\n", device.Root.Device.FriendlyName)
			fmt.Printf("  Num devices: %d\n", len(device.Root.Device.Devices))
		}
	}
	return nil
}
