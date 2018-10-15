# Changes in v2

* Imports `github.com/huin/goupnp/...` change to
  `github.com/huin/goupnp/v2/...`.
* `github.com/huin/goupnp` top-level package contents split into multiple
  packages:
  * `github.com/huin/goupnp/v2/discover`.
  * `github.com/huin/goupnp/v2/metadata`.
* `github.com/huin/goupnp/scpd` content merged into
  `github.com/huin/goupnp/v2/metadata`.
* Many APIs now take a required `context.Context` parameter.
* Renaming:
    * `httpu.HTTPUClient` -> `httpu.Client`
    * `ssdp.SSDPRawSearch` -> `ssdp.RawSearch`
    * `scpd.SCPD.ConfigId` -> `metadata.SCPD.ConfigID`
    * `soap.Marshal*` and `soap.Unmarshal*` functions renamed to meet go naming
      conventions.
    * `soap.SOAP*` -> `soap.*` to avoid stutter.
* `ssdp.RawSearch` and `httpu.HTTPUClient.Do` both take options to alter
  behavior. These options can be passed through from the appropriate `New*`
  functions in `dcps/*` packages. This should also allow for further options in
  future without further breaking API compatibility.
