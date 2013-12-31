package main

import (
	"text/template"
)

var packageTmpl = template.Must(template.New("package").Parse(`package {{.Name}}

import (
	"time"

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

// {{$srvIdent}} is a client for UPnP SOAP service with URN "{{.URN}}".
type {{$srvIdent}} struct {
	SOAPClient *soap.SOAPClient
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
	var request struct {{"{"}}{{range .Arguments}}{{if .IsInput}}{{.Name}} string
{{end}}{{end}}}
	// BEGIN Marshal arguments into request.
{{range $inargs}}{{$argWrap := $srv.WrapArgument .}}
	if request.{{.Name}}, err = {{$argWrap.Marshal}}; err != nil {
		return
	}{{end}}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {{"{"}}{{range $outargs}}{{.Name}} string
{{end}}}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction({{$srv.URNParts.Const}}, "{{.Name}}", &request, &response); err != nil {
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
`))
