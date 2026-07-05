# busdk.com

Static BusDK commercial website.

## Bus Engine OS In The Browser

Target architecture: a header-capable Bus Engine product page embeds a Bus
Engine OS virtual server running in the browser. The current development
release is published from `https://dev.hg.fi/beos/` with the QEMU/WASM runtime,
Bus Engine OS guest artifacts, release files, isolation headers, and browser
boot client outside this website repository. The current GitHub Pages-hosted
`busdk.com/engine/` page uses a direct launch preview because GitHub Pages does
not provide the COOP/COEP response headers required for a working pthread
WebAssembly iframe.

The command below is the older transitional exporter for local website-bundle
experiments:

```sh
make engine-wasm-os-static
```

The default output directory is `tmp/engine-wasm-os-static`. Override it with:

```sh
make engine-wasm-os-static ENGINE_WASM_OS_STATIC_DIR=/path/to/output
```

The command writes the browser-hosted OS page, iframe snippet, preview asset,
QEMU/WASM runtime, Bus Engine OS guest artifacts, firmware, generated license
indexes, third-party notices, and required source-material payloads into that
directory. Keep new runtime/client ownership out of this website repository;
website work should consume the published release URL instead of copying those
runtime artifacts into this repository. Re-enable a live iframe only from a
parent page served with COOP/COEP headers.

The release license index is scoped to the shipped Bus Engine OS package
manifests. `source-materials/` is intentionally limited to QEMU source
materials and source inputs for shipped package recipes whose effective
recorded license expression requires source delivery, such as GPL, LGPL, AGPL,
MPL, CDDL, EPL, CPL, GFDL, and EUPL style obligations. Dual-license packages
with a non-copyleft option do not copy source archives unless another required
license term still creates a source-delivery obligation. Permissive packages
remain listed in the license and notice indexes when shipped, but their source
archives are not copied.

The transitional exporter uses the `virtual-server` `x86_64` Bus Engine OS
package manifests by default. Set these only for local experiments with a
different proven guest image:

```sh
BUS_ENGINE_WASM_OS_PROFILE=virtual-server \
BUS_ENGINE_WASM_OS_TARGET_ARCH=x86_64 \
BUS_ENGINE_WASM_OS_QEMU_ARTIFACTS=/path/to/qemu-wasm \
BUS_ENGINE_WASM_OS_GUEST_ARTIFACTS=/path/to/guest-artifacts \
BUS_ENGINE_WASM_OS_QEMU_SOURCE=/path/to/qemu \
BUS_ENGINE_WASM_OS_ENGINE_OS_DIR=/path/to/bus-engine-os \
BUS_ENGINE_WASM_OS_SOURCES_CACHE=/path/to/verified-sources \
make engine-wasm-os-static ENGINE_WASM_OS_STATIC_DIR=/path/to/output
```

The generated directory must be served with COOP/COEP headers for browsers that
require `SharedArrayBuffer` for the QEMU/WASM runtime.
