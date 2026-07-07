# busdk.com

Static BusDK commercial website.

## Bus Engine OS In The Browser

Target architecture: a header-capable Bus Engine product page embeds a Bus
Engine OS virtual server running in the browser. The current development
release is published from `https://dev.hg.fi/beos/` with the QEMU/WASM runtime,
Bus Engine OS guest artifacts, release files, isolation headers, source
material, and browser boot client outside this website repository. The Engine
overview includes a live iframe slot for that published URL, but only assigns
the iframe `src` when the parent page is already cross-origin isolated. The
current GitHub Pages-hosted `busdk.com/engine/` page therefore keeps the direct
launch preview because GitHub Pages does not provide the COOP/COEP response
headers required for a working pthread WebAssembly iframe.

`docs/_headers` contains the narrow COOP/COEP policy for hosts that understand
Netlify/Cloudflare Pages-style `_headers` files:

```txt
/engine/
  Cross-Origin-Opener-Policy: same-origin
  Cross-Origin-Embedder-Policy: require-corp

/engine/*
  Cross-Origin-Opener-Policy: same-origin
  Cross-Origin-Embedder-Policy: require-corp
```

That file is deployment input for a header-capable static host or proxy. It
does not change GitHub Pages behavior by itself, and it does not move the
Engine-owned `/beos` release artifacts into this website repository.
When the parent page and `/beos` release are on different origins, the release
host must also allow cross-origin embedding, for example with
`Cross-Origin-Resource-Policy: cross-origin`, or the release must be served
through a same-origin parent/proxy path. Parent COOP/COEP headers alone are not
enough for a cross-origin iframe if the release response restricts embedding.

Use the normal quality target for offline source checks before changing the
Engine iframe or local bundle surface:

```sh
make quality
```

The first two checks in that target are specific to this browser-hosted OS
surface. `scripts/check-engine-wasm-bundle-surface.mjs` verifies the
repository-local `docs/engine/wasm-virtual-server/` manifest, canvas, serial
diagnostics, screenshot fallback, keyboard focus, and browser-isolation guard.
`scripts/check-engine-beos-surface.mjs` verifies the Engine landing page keeps
the browser-release copy, guarded iframe slot, direct-launch fallback, current
release URL, manifest link, and `SharedArrayBuffer` gate. The broader target
also runs GX/UI
documentation checks, which require the expected sibling module checkouts.
`scripts/check-engine-hosting-headers-config.mjs` verifies the optional
`docs/_headers` policy stays scoped to `/engine/` and `/engine/*` with the
required parent-page COOP/COEP values.
`scripts/check-engine-status-update.mjs` verifies `GOAL.md` and
`STATUS_UPDATE.md` still match the two active PLAN items and current
live-status markers.

Use these explicit networked checks when supervising deployed or published
state:

```sh
make engine-beos-check
make engine-beos-release-check
make engine-beos-public-page-check
make engine-beos-public-deploy-gate
make engine-status-update-check
```

`make engine-beos-check` is the full website supervision pass for this surface:
it runs the local bundle check, local Engine overview check, live release-host
check, header-capable deployment-config check, manager status check, and
deployed public parent-page check in that order.

`make engine-status-update-check` checks only the local manager-facing goal and
status artifacts. It fails if `GOAL.md` or `STATUS_UPDATE.md` stops naming the
two open PLAN items, current live release URL, path-implied manifest status,
public fallback state, or dispatch boundary.

`make engine-beos-public-deploy-gate` is the fail-closed parent-page deployment
gate for staging or future public hosting. It requires both a COOP/COEP
iframe-eligible parent page, a release response that is embeddable from the
parent origin, and the live profile-manifest link.

`make engine-beos-release-check` fetches the public Bus Engine OS release host,
its `virtual-server/` profile page, profile manifest, generated index, and
browser runtime script. It verifies the isolation headers, `riscv64`
browser-hosted manifest, `display=wasm`, `displayDevice=virtio-gpu-pci`, QEMU
runtime files, kernel, rootfs, graphical canvas, serial diagnostics, serial
input, and runtime hooks used by the browser client. The check also prints the
profile path and manifest profile-shape status. The current published
`virtual-server` release is accepted as `manifest_profile_shape=path-implied`;
future release manifests can make that explicit through a top-level profile id
or a `profiles[]` array without changing the website checker.

`make engine-beos-public-page-check` fetches the deployed
`https://busdk.com/engine/` parent page. It verifies that the public page still
points at `https://dev.hg.fi/beos/`, keeps the guarded iframe and direct
fallback, and reports `iframe_state=iframe-eligible` only when the parent page
itself is served with the COOP/COEP headers required for automatic iframe
activation. It also reports whether the release response is embeddable from
the parent origin. `iframe_state=fallback-required` and
`release_embedding=cross-origin-blocked-corp-same-origin` are the expected
current GitHub Pages state.

For staging or a header-capable deployment that is expected to activate the
iframe, require the parent-page headers explicitly. For a deployment that is
also expected to include the current local profile-manifest link, require that
link too:

```sh
make engine-beos-public-page-check \
  BUS_ENGINE_PUBLIC_PAGE_URL=https://example.test/engine/ \
  BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1 \
  BUS_ENGINE_REQUIRE_MANIFEST_LINK=1
```

That mode fails if the parent page still lacks COOP/COEP. Leave
`BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE` and `BUS_ENGINE_REQUIRE_MANIFEST_LINK`
unset or `0` when checking the current GitHub Pages deployment, where fallback
and `manifest_link=not-yet-deployed` reporting are the expected state.

Use the named deployment gate when both requirements are expected:

```sh
make engine-beos-public-deploy-gate \
  BUS_ENGINE_PUBLIC_PAGE_URL=https://example.test/engine/
```

For staging or a future release host, pass Make variables instead of editing
the scripts:

```sh
make engine-beos-check \
  BUS_ENGINE_BEOS_RELEASE_URL=https://example.test/beos/ \
  BUS_ENGINE_BEOS_PROFILE_PATH=virtual-server/ \
  BUS_ENGINE_PUBLIC_PAGE_URL=https://example.test/engine/ \
  BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1 \
  BUS_ENGINE_REQUIRE_MANIFEST_LINK=1 \
  BUS_ENGINE_CHECK_TIMEOUT_MS=30000
```

Set `BUS_ENGINE_BEOS_PROFILE_PATH` only when the profile path is not
`virtual-server/`. The local Engine overview source check, live release check,
and public parent-page check all use those values. The underlying scripts also
accept equivalent environment variables, and `check-engine-public-page.mjs`
accepts an optional Engine page URL argument when checking a non-production
parent page. `BUS_ENGINE_CHECK_TIMEOUT_MS` controls each network request
timeout for the live release and public parent-page checks.

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
runtime artifacts into this repository. The published overview can activate
its live iframe from a parent page served with COOP/COEP headers; otherwise it
keeps the launch preview.

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

The local iframe harness reads the active profile from `?profile=...` and falls
back to the manifest's default guest profile. The repository-local manifest
keeps separate `virtual-server` and planned `virtual-desktop` profile blocks so
display mode, display device, resolution, keyboard behavior, and guest boot
expectations can be selected from the manifest instead of hard-coded into the
page.
