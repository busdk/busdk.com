# STATUS_UPDATE.md

## Bus Engine OS Browser Release Website Lane - 2026-07-07

### Scope

This file tracks the `busdk.com` website-owned iframe, landing, and manifest
surface for the current Bus Engine OS browser release. It must not contain
secrets, private tokens, or Engine/QEMU implementation notes.

### Current Live Status

- Published release URL: `https://dev.hg.fi/beos/`
- Published profile path: `virtual-server/`
- Published profile manifest:
  `https://dev.hg.fi/beos/virtual-server/browser-hosted-manifest.json`
- Current release check: `make engine-beos-check` passes.
- Live release manifest status: `profile_path=virtual-server` and
  `manifest_profile_shape=path-implied`.
- Live release display status: `display=wasm` and
  `display_device=virtio-gpu-pci`.
- Published file status: `files=14` and `published_files=14`.
- Public parent status: `https://busdk.com/engine/` reports
  `iframe_state=fallback-required`.
- Landing status: local Engine landing source names the Bus Engine OS browser
  release and links the current profile manifest.
- Public landing status: deployed `https://busdk.com/engine/` still exposes the
  browser-hosted release section and live QEMU/WASM release CTA; the public
  checker reports `landing=browser-hosted-release`.
- Release embedding status: the current public parent/release pair reports
  `release_embedding=cross-origin-blocked-corp-same-origin`.
- Public parent deployment status: local source has the profile-manifest link;
  deployed public page still reports `manifest_link=not-yet-deployed`.
- Cross-origin embedding status: a future header-capable parent must either
  serve `/beos` through the same origin/proxy path or use a release response
  that allows cross-origin embedding, such as
  `Cross-Origin-Resource-Policy: cross-origin`.

### Open PLAN Items

1. Update the Bus Engine OS iframe from terminal-only boot output to an
   interactive graphics surface.
   - Website status: parent markup, guarded iframe activation, direct fallback,
     local graphical bundle harness, canvas focus, serial diagnostics,
     manifest-driven display fields, profile selector/status,
     `docs/_headers`, and supervision checks are in place.
   - Blocker: the deployed `busdk.com/engine/` parent page is still hosted in a
     mode that lacks COOP/COEP, so the public page remains fallback-only. A
     future cross-origin parent also needs the `/beos` release response to be
     embeddable from that parent origin, or a same-origin parent/proxy path.
   - Next gate: publish through a header-capable host or proxy and run:

     ```sh
     make engine-beos-public-page-check \
       BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1 \
       BUS_ENGINE_REQUIRE_MANIFEST_LINK=1
     ```

     The equivalent named gate is:

     ```sh
     make engine-beos-public-deploy-gate
     ```

2. Add a browser-hosted OS manifest shape that can describe `virtual-server`
   and `virtual-desktop` profiles separately.
   - Website status: the repository-local bundle manifest has separate
     `virtual-server` and planned `virtual-desktop` profile blocks, and the
     local harness consumes profile runtime and guest fields from the manifest.
   - Blocker: the published `/beos` manifest is still path-implied for the
     current `virtual-server` release and does not yet publish explicit
     `virtual-server` / `virtual-desktop` profile metadata.
   - Next gate: Engine/BEO publishes explicit profile metadata; then
     `make engine-beos-release-check` should report an explicit profile shape
     instead of `manifest_profile_shape=path-implied`.

### Hosting Versus Code

- Website code is not the current blocker for the `virtual-server` release
  surface. The website lane has checks for local source, header configuration,
  live release host, and deployed parent-page status.
- Hosting remains the public iframe blocker. GitHub Pages does not provide the
  required parent-page COOP/COEP headers for automatic iframe activation.
- Cross-origin release embedding is part of the hosting boundary: parent
  COOP/COEP alone is not enough if the release response restricts embedding to
  its own origin.
- `docs/_headers` is included for header-capable static hosts or proxies. It
  does not change GitHub Pages behavior by itself.
- Engine/BEO owns the hosted `/beos` artifacts, release manifest, runtime,
  release-host headers, explicit future profile metadata, and future
  `virtual-desktop` artifacts.

### Checks

- `make engine-beos-check`: passing.
- `make engine-beos-public-page-check`: passing in report-only mode with
  `landing=browser-hosted-release`, `iframe_state=fallback-required`, and
  `release_embedding=cross-origin-blocked-corp-same-origin`.
- `make engine-beos-public-page-check BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1`:
  expected failure against current `https://busdk.com/engine/` because the
  deployed parent page lacks COOP/COEP.
- `make engine-beos-public-page-check BUS_ENGINE_REQUIRE_MANIFEST_LINK=1`:
  expected failure against current `https://busdk.com/engine/` until the local
  profile-manifest link is published.
- `make engine-beos-public-deploy-gate`: expected failure against current
  `https://busdk.com/engine/` until both parent-page COOP/COEP headers and the
  profile-manifest link are deployed, and the release is embeddable from the
  parent origin or served through the same origin.
- `make quality`: Engine-specific checks pass first, then the broader target
  fails at the existing missing sibling dependency `../bus-update/go.mod`.

### Dispatch Plan

- Website lane:
  - keep parent-page iframe/fallback markup aligned with the published release;
  - keep `docs/_headers`, local checks, live checks, and public-page checks
    current;
  - consume explicit profile metadata when Engine/BEO publishes it;
  - rerun `BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1` and
    `BUS_ENGINE_REQUIRE_MANIFEST_LINK=1` against any staging or future
    header-capable parent deployment, or use
    `make engine-beos-public-deploy-gate` for the combined check.
- Engine/BEO lane:
  - keep `https://dev.hg.fi/beos/` artifacts and headers healthy;
  - publish explicit profile metadata for the current and future profiles;
  - own runtime/guest fixes for graphical keyboard proof and future
    `virtual-desktop` readiness.
