package main

import (
	"text/template"
)

var packageTmpl = template.Must(template.New("package").Parse(`package {{.Name}}

import (
	"time"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/soap"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const ({{range .DeviceTypes}}
	{{.Const}} = "{{.URN}}"{{end}}
)

// Service URNs:
const ({{range .ServiceTypes}}
	{{.Const}} = "{{.URN}}"{{end}}
)

{{range .Services}}
{{$srv := .}}
{{$srvIdent := printf "%s%s" .Name .Version}}

// {{$srvIdent}} is a client for UPnP SOAP service with URN "{{.URN}}". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type {{$srvIdent}} struct {
	goupnp.ServiceClient
}

// New{{$srvIdent}}Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func New{{$srvIdent}}Clients() (clients []*{{$srvIdent}}, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClients({{$srv.Const}}); err != nil {
		return
	}
	clients = make([]*{{$srvIdent}}, len(genericClients))
	for i := range genericClients {
		clients[i] = &{{$srvIdent}}{genericClients[i]}
	}
	return
}

{{range .SCPD.Actions}}{{/* loops over *SCPDWithURN values */}}

{{$inargs := .InputArguments}}{{$outargs := .OutputArguments}}
// {{if $inargs}}Arguments:{{range $inargs}}{{$argWrap := $srv.WrapArgument .}}
//
// * {{.Name}}: {{$argWrap.Document}}{{end}}{{end}}
//
// {{if $outargs}}Return values:{{range $outargs}}{{$argWrap := $srv.WrapArgument .}}
//
// * {{.Name}}: {{$argWrap.Document}}{{end}}{{end}}
func (client *{{$srvIdent}}) {{.Name}}({{range $inargs}}{{/*
*/}}{{$argWrap := $srv.WrapArgument .}}{{$argWrap.AsParameter}}, {{end}}{{/*
*/}}) ({{range $outargs}}{{/*
*/}}{{$argWrap := $srv.WrapArgument .}}{{$argWrap.AsParameter}}, {{end}} err error) {
	// Request structure.
	request := {{if $inargs}}&{{template "argstruct" $inargs}}{{"{}"}}{{else}}{{"interface{}(nil)"}}{{end}}
	// BEGIN Marshal arguments into request.
{{range $inargs}}{{$argWrap := $srv.WrapArgument .}}
	if request.{{.Name}}, err = {{$argWrap.Marshal}}; err != nil {
		return
	}{{end}}
	// END Marshal arguments into request.

	// Response structure.
	response := {{if $outargs}}&{{template "argstruct" $outargs}}{{"{}"}}{{else}}{{"interface{}(nil)"}}{{end}}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction({{$srv.URNParts.Const}}, "{{.Name}}", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.
{{range $outargs}}{{$argWrap := $srv.WrapArgument .}}
	if {{.Name}}, err = {{$argWrap.Unmarshal "response"}}; err != nil {
		return
	}{{end}}
	// END Unmarshal arguments from response.
	return
}
{{end}}{{/* range .SCPD.Actions */}}
{{end}}{{/* range .Services */}}

{{define "argstruct"}}struct {{"{"}}{{range .}}
{{.Name}} string
{{end}}{{"}"}}{{end}}
`))
