// Serves as a simple example/test of discovering UPnP devices on the local
// network.
package main

import (
	"encoding/xml"
	"fmt"

	"github.com/huin/goupnp"
)

var spaces = "                                     "

type GetExternalIPAddressRequest struct {
	XMLName xml.Name
}

type GetExternalIPAddressResponse struct {
	XMLName              xml.Name
	NewExternalIPAddress string
}

type Uint16Value struct {
	XMLName xml.Name
	Value   uint16 `xml:",chardata"`
}

type GetGenericPortMappingEntryRequest struct {
	XMLName             xml.Name
	NewPortMappingIndex Uint16Value
}

type GetGenericPortMappingEntryResponse struct {
	XMLName           xml.Name
	NewRemoteHost     string
	NewExternalPort   uint16
	NewProtocol       string
	NewInternalPort   uint16
	NewInternalClient string
	NewEnabled        string // boolean
	NewLeaseDuration  uint32
}

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

			if scdp, err := srv.RequestSCDP(); err != nil {
				fmt.Printf("Error requesting SCPD: %v\n", err)
			} else {
				fmt.Println("Available SCPD actions:")
				for _, action := range scdp.Actions {
					fmt.Println(" ", action.Name)
				}
			}

			srvClient := srv.NewSOAPClient()

			{
				inAction := GetExternalIPAddressRequest{XMLName: xml.Name{Space: goupnp.ServiceTypeWANPPPConnection, Local: "GetExternalIPAddress"}}
				var outAction GetExternalIPAddressResponse
				err := srvClient.PerformAction(goupnp.ServiceTypeWANPPPConnection, "GetExternalIPAddress", &inAction, &outAction)
				if err != nil {
					fmt.Printf("Failed to GetExternalIPAddress from %s: %v\n", device.FriendlyName, err)
					continue
				}
				fmt.Printf("Got GetExternalIPAddress result from %s: %+v\n", device.FriendlyName, outAction)
			}

			for i := uint16(0); i < 10; i++ {
				inAction := GetGenericPortMappingEntryRequest{XMLName: xml.Name{Space: goupnp.ServiceTypeWANPPPConnection, Local: "GetGenericPortMappingEntry"}, NewPortMappingIndex: Uint16Value{XMLName: xml.Name{"", "NewPortMappingIndex"}, Value: i}}
				var outAction GetGenericPortMappingEntryResponse
				err := srvClient.PerformAction(goupnp.ServiceTypeWANPPPConnection, "GetGenericPortMappingEntry", &inAction, &outAction)
				if err != nil {
					fmt.Printf("Failed to GetGenericPortMappingEntry on %s: %v\n", device.FriendlyName, err)
					continue
				}
				fmt.Printf("Got GetGenericPortMappingEntry from %s: %+v\n", device.FriendlyName, outAction)
			}
		}
	}
}
