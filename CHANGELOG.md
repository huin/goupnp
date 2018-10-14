# Changes in v2

* Imports `github.com/huin/goupnp/...` change to
  `github.com/huin/goupnp/v2/...`.
* `github.com/huin/goupnp` top-level package contents moved to
  `github.com/huin/goupnp/v2/discover`.
* Many APIs now take a required `context.Context` parameter.
* Renaming:
    * `httpu.HTTPUClient` -> `httpu.Client`
    * `ssdp.SSDPRawSearch` -> `ssdp.RawSearch`
* `ssdp.RawSearch` and `httpu.HTTPUClient.Do` both take options to alter
  behavior. These options can be passed through from the appropriate `New*`
  functions in `dcps/*` packages. This should also allow for further options in
  future without further breaking API compatibility.
