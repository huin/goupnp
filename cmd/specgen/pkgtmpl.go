package main

import (
	"text/template"
)

var packageTmpl = template.Must(template.New("package").Parse(`package {{.Name}}

import "github.com/huin/goupnp/soap"

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

{{$reqType := printf "_%s_%s_Request" $srvIdent .Name}}
{{$respType := printf "_%s_%s_Response" $srvIdent .Name}}

// {{$reqType}} is the XML structure for the input arguments for action {{.Name}}.
type {{$reqType}} struct {{"{"}}{{range .Arguments}}{{if .IsInput}}
	{{.Name}} {{$srv.SCPD.GoKindNameForVariable .RelatedStateVariable "string"}}
{{end}}{{end}}}

// {{$respType}} is the XML structure for the output arguments for action {{.Name}}.
type {{$respType}} struct {{"{"}}{{range .Arguments}}{{if .IsOutput}}
	{{.Name}} {{$srv.SCPD.GoKindNameForVariable .RelatedStateVariable "string"}}
{{end}}{{end}}}

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
	{{.Name}} {{$srv.SCPD.GoKindNameForVariable .RelatedStateVariable "string"}},
{{end}}{{end}}) ({{range .Arguments}}{{if .IsOutput}}
	{{.Name}} {{$srv.SCPD.GoKindNameForVariable .RelatedStateVariable "string"}},
{{end}}{{end}} err error) {
	request := {{$reqType}}{
{{range .Arguments}}{{if .IsInput}}
	{{.Name}}: {{.Name}},
{{end}}{{end}}
	}
	var response {{$respType}}
	err = client.SOAPClient.PerformAction({{$srv.URNParts.Const}}, "{{.Name}}", &request, &response)
	if err != nil {
		return
	}
{{range .Arguments}}{{if .IsOutput}}
	{{.Name}} = response.{{.Name}}
{{end}}{{end}}
	return
}

{{end}}{{/* range .SCPD.Actions */}}
{{end}}{{/* range .Services */}}
`))
