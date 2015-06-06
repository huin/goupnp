goupnp is a UPnP client library for Go

Installation
------------

Run `go get -u github.com/huin/goupnp`.

Documentation
-------------

[![GoDoc](https://godoc.org/github.com/huin/goupnp?status.svg)](https://godoc.org/github.com/huin/goupnp)

Regenerating dcps generated source code:
----------------------------------------

1. Install gotasks: `go get -u github.com/jingweno/gotask`
2. Change to the gotasks directory: `cd gotasks`
3. Download UPnP specification data (if not done already): `wget http://upnp.org/resources/upnpresources.zip`
4. Regenerate source code: `gotask specgen -s upnpresources.zip -o ../dcps`

Supporting additional UPnP devices and services:
------------------------------------------------

Supporting additional services is, in the trivial case, simply a matter of
adding the service to the `dcpMetadataByDir` whitelist in
`gotasks/specgen_task.go`, regenerating the source code (see above), and
committing that source code.

However, it would be helpful if anyone needing such a service could test the
service against the service they have, and then reporting any trouble
encountered as an [issue on this
project](https://github.com/huin/goupnp/issues/new). If it just works, then
please report at least minimal working functionality as an issue, and
optionally contribute the metadata upstream.
