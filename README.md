# busdk.com

Static BusDK commercial website.

## Bus Engine Browser Lab

Create the publishable Bus Engine OS browser-lab bundle with:

```sh
make engine-wasm-os-static
```

The default output directory is `tmp/engine-wasm-os-static`. Override it with:

```sh
make engine-wasm-os-static ENGINE_WASM_OS_STATIC_DIR=/path/to/output
```

The command writes the browser-lab page, iframe snippet, preview asset,
QEMU/WASM runtime, Bus Engine OS guest artifacts, firmware, generated license
indexes, third-party notices, and required source-material payloads into that
directory.

The release license index is scoped to the shipped Bus Engine OS package
manifests. `source-materials/` is intentionally limited to QEMU source
materials and source inputs for shipped package recipes whose recorded
licenses require source delivery, such as GPL, LGPL, AGPL, MPL, CDDL, EPL,
CPL, GFDL, and EUPL style obligations. Permissive packages remain listed in
the license and notice indexes when shipped, but their source archives are not
copied.

The exporter uses the `virtual-server` `x86_64` Bus Engine OS package manifests
by default. Set these only when publishing a different proven guest image:

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
