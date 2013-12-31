package example_test

import (
	"fmt"
	"os"

	"github.com/huin/goupnp/dcps/internetgateway1"
)

// Discovering a internet gateway devices on the local network, and asking each
// of them for their external IP address.
func Example_getExternalIPAddress() {
	// import (
	//   "fmt"
	//   "os"
	//   "github.com/huin/goupnp/dcps/internetgateway1"
	// )
	fmt.Println("Running")
	clients, errors, err := internetgateway1.NewWANPPPConnection1Clients()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error discovering service with UPnP: ", err)
		return
	}

	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "Error discovering %d services:\n", len(errors))
		for _, err := range errors {
			fmt.Println("  ", err)
		}
	}

	fmt.Fprintf(os.Stderr, "Successfully discovered %d services:\n", len(clients))
	for _, client := range clients {
		device := &client.RootDevice.Device

		fmt.Fprintln(os.Stderr, "  Device:", device.FriendlyName)
		if addr, err := client.GetExternalIPAddress(); err != nil {
			fmt.Fprintf(os.Stderr, "    Failed to get external IP address: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "    External IP address: %v\n", addr)
		}
	}
	fmt.Println("Complete")

	// Output:
	// Running
	// Complete
}
