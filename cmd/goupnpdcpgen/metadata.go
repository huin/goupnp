package main

import (
	"strings"
)

// DCP contains extra metadata to use when generating DCP source files.
type DCPMetadata struct {
	Name         string // What to name the Go DCP package.
	OfficialName string // Official name for the DCP.
	Src          dcpProvider
}

var dcpMetadata = []DCPMetadata{
	{
		Name:         "internetgateway1",
		OfficialName: "Internet Gateway Device v1",
		Src: upnpdotorg{
			DocURL:     "http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v1-Device.pdf",
			XMLSpecURL: "http://upnp.org/specs/gw/UPnP-gw-IGD-TestFiles-20010921.zip",
			Hacks: []DCPHackFn{
				fixTotalBytes("urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"),
			},
		},
	},
	{
		Name:         "internetgateway2",
		OfficialName: "Internet Gateway Device v2",
		Src: upnpdotorg{
			DocURL:     "http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v2-Device.pdf",
			XMLSpecURL: "http://upnp.org/specs/gw/UPnP-gw-IGD-Testfiles-20110224.zip",
			Hacks: []DCPHackFn{
				fixMissingURN("urn:schemas-upnp-org:service:WANIPv6FirewallControl:1"),
				fixTotalBytes("urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"),
			},
		},
	},
	{
		Name:         "av1",
		OfficialName: "MediaServer v1 and MediaRenderer v1",
		Src: upnpdotorg{
			DocURL:     "http://upnp.org/specs/av/av1/",
			XMLSpecURL: "http://upnp.org/specs/av/UPnP-av-TestFiles-20070927.zip",
		},
	},
	{
		Name:         "ocf/internetgateway1",
		OfficialName: "Internet Gateway Device v1 - Open Connectivity Foundation",
		Src: openconnectivitydotorg{
			SpecsURL:       ocfSpecsURL,
			DocPath:        "*/DeviceProtection_1/UPnP-gw-*v1*.pdf",
			XMLSpecZipPath: "*/DeviceProtection_1/UPnP-gw-IGD-TestFiles-*.zip",
			XMLServicePath: []string{"*/service/*1.xml"},
			XMLDevicePath:  []string{"*/device/*1.xml"},
			Hacks: []DCPHackFn{
				fixMissingURN("urn:schemas-upnp-org:service:DeviceProtection:1"),
				fixMissingURN("urn:schemas-upnp-org:service:WANIPv6FirewallControl:1"),
				fixTotalBytes(),
			},
		},
	},
	{
		Name:         "ocf/internetgateway2",
		OfficialName: "Internet Gateway Device v2 - Open Connectivity Foundation",
		Src: openconnectivitydotorg{
			SpecsURL:       ocfSpecsURL,
			DocPath:        "*/Internet Gateway_2/UPnP-gw-*.pdf",
			XMLSpecZipPath: "*/Internet Gateway_2/UPnP-gw-IGD-TestFiles-*.zip",
			XMLServicePath: []string{"*/service/*1.xml", "*/service/*2.xml"},
			XMLDevicePath:  []string{"*/device/*1.xml", "*/device/*2.xml"},
			Hacks: []DCPHackFn{
				fixMissingURN("urn:schemas-upnp-org:service:DeviceProtection:1"),
				fixTotalBytes(),
			},
		},
	},
}

func fixTotalBytes(malformedURNs ...string) func(dcp *DCP) error {
	malformedVariables := []string{
		"TotalBytesSent",
		"TotalBytesReceived",
	}
	return func(dcp *DCP) error {
		for _, service := range dcp.Services {
			var process bool
			for _, malformedURN := range malformedURNs {
				if service.URN == malformedURN {
					process = true
					break
				}
			}
			if process || len(malformedURNs) < 1 {
				variables := service.SCPD.StateVariables
				for key, variable := range variables {
					varName := variable.Name
					for _, malformedVariable := range malformedVariables {
						if strings.HasSuffix(varName, malformedVariable) {
							// Fix size of total bytes which is by default ui4 or maximum 4 GiB.
							variable.DataType.Name = "ui8"
							variables[key] = variable
						}
					}
				}
			}
		}
		return nil
	}
}

func fixMissingURN(missingURNs ...string) func(dcp *DCP) error {
	return func(dcp *DCP) error {
		for _, missingURN := range missingURNs {
			if _, ok := dcp.ServiceTypes[missingURN]; ok {
				continue
			}
			urnParts, err := extractURNParts(missingURN, serviceURNPrefix)
			if err != nil {
				return err
			}
			dcp.ServiceTypes[missingURN] = urnParts
		}
		return nil
	}
}

type DCPHackFn func(*DCP) error
