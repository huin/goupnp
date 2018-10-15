goupnp is a [UPnP](https://en.wikipedia.org/wiki/Universal_Plug_and_Play)
client library for the [Go programming language](https://golang.org/).

## Installation

Run `go get -u github.com/huin/goupnp/v2`.

## Documentation

### Supported DCPs

You probably want to start with one of these.

* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/dcps/av1?status.svg)
  av1](https://godoc.org/github.com/huin/goupnp/v2/dcps/av1) Client for UPnP
  Device Control Protocol MediaServer v1 and MediaRenderer v1.
* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/dcps/internetgateway1?status.svg)
  internetgateway1](https://godoc.org/github.com/huin/goupnp/v2/dcps/internetgateway1)
  Client for UPnP Device Control Protocol Internet Gateway Device v1.
* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/dcps/internetgateway2?status.svg)
  internetgateway2](https://godoc.org/github.com/huin/goupnp/v2/dcps/internetgateway2)
  Client for UPnP Device Control Protocol Internet Gateway Device v2.

### Core components

The core components were originally intended as internal, but have seen use
outside of goupnp.

* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/discover?status.svg)
  (discover)](https://godoc.org/github.com/huin/goupnp/v2/discover) core
  library; contains datastructures and utilities typically used by the
  implemented DCPs.
* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/ssdp?status.svg)
  ssdp](https://godoc.org/github.com/huin/goupnp/v2/ssdp) SSDP client
  implementation ([Simple Service Discovery
  Protocol](https://en.wikipedia.org/wiki/Simple_Service_Discovery_Protocol));
  used to discover UPnP services on a network.
* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/httpu?status.svg)
  httpu](https://godoc.org/github.com/huin/goupnp/v2/httpu)
  [HTTPU](https://en.wikipedia.org/wiki/HTTPU) implementation, underlies SSDP.
* [![GoDoc](https://godoc.org/github.com/huin/goupnp/v2/soap?status.svg)
  soap](https://godoc.org/github.com/huin/goupnp/v2/soap) SOAP client
  implementation ([Simple Object Access
  Protocol](https://en.wikipedia.org/wiki/SOAP)); used to communicate with
  discovered services.

## Regenerating generated source code for DCPs

1. Build code generator:

	`go get -u github.com/huin/goupnp/v2/cmd/goupnpdcpgen`

2. Regenerate the code:

	`go generate ./...`

3. Run tests, and commit the generated code.

## Supporting additional UPnP devices and services

Supporting additional services is, in the trivial case, simply a matter of
adding the service to the `dcpMetadata` whitelist in
`cmd/goupnpdcpgen/registry.go`, regenerating the source code (see above), and
committing that source code.

However, it would be helpful if anyone needing such a service could test the
service against the service they have, and then reporting any trouble
encountered as an [issue on this
project](https://github.com/huin/goupnp/v2/issues/new). If it just works, then
please report at least minimal working functionality as an issue, and
optionally contribute the metadata upstream.
