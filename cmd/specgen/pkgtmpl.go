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

const ({{range .DeviceTypes}}
	{{.Const}} = "{{.URN}}"
{{end}})

const ({{range .ServiceTypes}}
	{{.Const}} = "{{.URN}}"
{{end}})

{{range .Services}}
{{$srv := .}}
{{$srvIdent := printf "%s%s" .Name .Version}}

// {{$srvIdent}} is a client for UPnP SOAP service with URN "{{.URN}}".
type {{$srvIdent}} struct {
	SOAPClient soap.SOAPClient
}

{{range .SCPD.Actions}}{{/* loops over *SCPDWithURN values */}}

// {{.Name}} action.
// Arguments:
//{{range .Arguments}}{{if .IsInput}}
// * {{.Name}}: {{$v := $srv.SCPD.GetStateVariable .RelatedStateVariable}}
{{if $v}}// (related state variable: {{$v.Name}})
// - {{if $v.AllowedValueRange}}allowed range: {{$v.AllowedValueRange.Minimum}} to {{$v.AllowedValueRange.Maximum}}{{end}}
// - {{if $v.AllowedValues}}allowed values:
// {{range $i, $val := $v.AllowedValues}}{{if $i}}|{{end}}{{$val}}{{end}}{{end}}
//{{else}}
// (unknown){{end}}
//{{end}}{{end}}
//
// Return values:
//{{range .Arguments}}{{if .IsOutput}}
// * {{.Name}}: {{$v := $srv.SCPD.GetStateVariable .RelatedStateVariable}}
{{if $v}}// (related state variable: {{$v.Name}})
// - {{if $v.AllowedValueRange}}allowed range: {{$v.AllowedValueRange.Minimum}} to {{$v.AllowedValueRange.Maximum}}{{end}}
// - {{if $v.AllowedValues}}allowed values:
// {{range $i, $val := $v.AllowedValues}}{{if $i}}|{{end}}{{$val}}{{end}}{{end}}
//{{else}}
// (unknown){{end}}
//{{end}}{{end}}
func (client *{{$srvIdent}}) {{.Name}}({{range .Arguments}}{{if .IsInput}}
	{{$argWrap := $srv.Argument .}}{{$argWrap.AsParameter}},{{end}}{{end}}
) ({{range .Arguments}}{{if .IsOutput}}
	{{$argWrap := $srv.Argument .}}{{$argWrap.AsParameter}},{{end}}{{end}}
	err error,
) {
	// Request structure.
	var request struct {{"{"}}{{range .Arguments}}{{if .IsInput}}{{.Name}} string
{{end}}{{end}}}
	// BEGIN Marshal arguments into request.
{{range .Arguments}}{{if .IsInput}}{{$argWrap := $srv.Argument .}}
	if request.{{.Name}}, err = {{$argWrap.Marshal}}; err != nil {
		return
	}
{{end}}{{end}}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {{"{"}}{{range .Arguments}}{{if .IsOutput}}{{.Name}} string
{{end}}{{end}}}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction({{$srv.URNParts.Const}}, "{{.Name}}", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.
{{range .Arguments}}{{if .IsOutput}}{{$argWrap := $srv.Argument .}}
	if {{.Name}}, err = {{$argWrap.Unmarshal "response"}}; err != nil {
		return
	}
{{end}}{{end}}
	// END Unmarshal arguments from response.
	return
}
{{end}}{{/* range .SCPD.Actions */}}
{{end}}{{/* range .Services */}}
`))
