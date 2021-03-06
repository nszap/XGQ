package main

import (
	"fmt"
	"os"

	"github.com/huin/goupnp"
)

func main() {
}

type GetExternalIPAddresser interface {
	GetExternalIPAddress() (NewExternalIPAddress string, err error)
	GetServiceClient() *goupnp.ServiceClient
}

func DisplayExternalIPResults(clients []GetExternalIPAddresser, errors []error, err error) {
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
		device := &client.GetServiceClient().RootDevice.Device

		fmt.Fprintln(os.Stderr, "  Device:", device.FriendlyName)
		if addr, err := client.GetExternalIPAddress(); err != nil {
			fmt.Fprintf(os.Stderr, "    Failed to get external IP address: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "    External IP address: %v\n", addr)
		}
	}
}
