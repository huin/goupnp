package main

// dcp contains extra metadata to use when generating dcp source files.
type dcpMetadata struct {
	// What to name the Go dcp package.
	Name string
	// Official name for the dcp.
	OfficialName string
	// Optional - URL for further documentation about the dcp.
	DocURL string
	// Where to download the XML spec from.
	XMLSpecURL string
	// Any special-case functions to run against the dcp before writing it out.
	Hacks []dcpHackFn
}

var dcpMetadataRegistry = []dcpMetadata{
	{
		Name:         "internetgateway1",
		OfficialName: "Internet Gateway Device v1",
		DocURL:       "http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v1-Device.pdf",
		XMLSpecURL:   "http://upnp.org/specs/gw/UPnP-gw-IGD-TestFiles-20010921.zip",
		Hacks:        []dcpHackFn{totalBytesHack},
	},
	{
		Name:         "internetgateway2",
		OfficialName: "Internet Gateway Device v2",
		DocURL:       "http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v2-Device.pdf",
		XMLSpecURL:   "http://upnp.org/specs/gw/UPnP-gw-IGD-Testfiles-20110224.zip",
		Hacks: []dcpHackFn{
			func(dcp *dcp) error {
				missingURN := "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1"
				if _, ok := dcp.ServiceTypes[missingURN]; ok {
					return nil
				}
				urnParts, err := extractURNParts(missingURN, serviceURNPrefix)
				if err != nil {
					return err
				}
				dcp.ServiceTypes[missingURN] = urnParts
				return nil
			}, totalBytesHack,
		},
	},
	{
		Name:         "av1",
		OfficialName: "MediaServer v1 and MediaRenderer v1",
		DocURL:       "http://upnp.org/specs/av/av1/",
		XMLSpecURL:   "http://upnp.org/specs/av/UPnP-av-TestFiles-20070927.zip",
	},
}

func totalBytesHack(dcp *dcp) error {
	for _, service := range dcp.Services {
		if service.URNParts.URN == "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1" {
			variables := service.SCPD.StateVariables
			for key, variable := range variables {
				varName := variable.Name
				if varName == "TotalBytesSent" || varName == "TotalBytesReceived" {
					// Fix size of total bytes which is by default ui4 or
					// maximum 4 GiB.
					variable.DataType.Name = "ui8"
					variables[key] = variable
				}
			}

			break
		}
	}

	return nil
}

type dcpHackFn func(*dcp) error
