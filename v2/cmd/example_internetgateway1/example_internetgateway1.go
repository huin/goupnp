package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huin/goupnp/v2/dcps/internetgateway1"
)

func main() {
	ctx := context.Background()

	clients, errors, err := internetgateway1.NewWANPPPConnection1Clients(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got %d errors finding servers and %d successfully discovered.\n",
		len(errors), len(clients))
	for i, e := range errors {
		fmt.Printf("Error finding server #%d: %v\n", i+1, e)
	}

	for _, c := range clients {
		dev := &c.ServiceClient.RootDevice.Device
		srv := c.ServiceClient.Service
		fmt.Println(dev.FriendlyName, " :: ", srv.String())
		scpd, err := srv.RequestSCPD(ctx)
		if err != nil {
			fmt.Printf("  Error requesting service SCPD: %v\n", err)
		} else {
			fmt.Println("  Available actions:")
			for _, action := range scpd.Actions {
				fmt.Printf("  * %s\n", action.Name)
				for _, arg := range action.Arguments {
					var varDesc string
					if stateVar := scpd.GetStateVariable(arg.RelatedStateVariable); stateVar != nil {
						varDesc = fmt.Sprintf(" (%s)", stateVar.DataType.Name)
					}
					fmt.Printf("    * [%s] %s%s\n", arg.Direction, arg.Name, varDesc)
				}
			}
		}

		if scpd == nil || scpd.GetAction("GetExternalIPAddress") != nil {
			ip, err := c.GetExternalIPAddress(ctx)
			fmt.Println("GetExternalIPAddress: ", ip, err)
		}

		if scpd == nil || scpd.GetAction("GetStatusInfo") != nil {
			status, lastErr, uptime, err := c.GetStatusInfo(ctx)
			fmt.Println("GetStatusInfo: ", status, lastErr, uptime, err)
		}

		if scpd == nil || scpd.GetAction("GetIdleDisconnectTime") != nil {
			idleTime, err := c.GetIdleDisconnectTime(ctx)
			fmt.Println("GetIdleDisconnectTime: ", idleTime, err)
		}

		if scpd == nil || scpd.GetAction("AddPortMapping") != nil {
			err := c.AddPortMapping(ctx, "", 5000, "TCP", 5001, "192.168.1.2", true, "Test port mapping", 0)
			fmt.Println("AddPortMapping: ", err)
		}
		if scpd == nil || scpd.GetAction("DeletePortMapping") != nil {
			err := c.DeletePortMapping(ctx, "", 5000, "TCP")
			fmt.Println("DeletePortMapping: ", err)
		}
	}
}
