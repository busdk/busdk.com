# PLAN.md

# Current work

- [x] Align the `busdk.com` Bus Engine OS iframe/manifest surface with the current
      browser-hosted release manifest surface from `https://dev.hg.fi/beos/`:
      keep the GitHub Pages direct-launch fallback, consume profile-specific
      manifest fields in an explicit live-release check, and record the
      remaining hosting boundary separately from website code.
- [x] Extend the `busdk.com` Bus Engine OS live-release supervision check so it
      validates the profile page and browser runtime advertised by the manifest,
      not just the top-level release headers and manifest JSON.
- [x] Add a public `busdk.com/engine/` hosting check that verifies the deployed
      parent page still advertises the current BEO release URL and reports
      whether iframe activation is blocked by missing COOP/COEP headers or can
      proceed under header-capable hosting.
- [x] Add an opt-in public parent-page gate that fails when a staging or
      header-capable Bus Engine product page is expected to activate the
      `/beos` iframe but is still missing COOP/COEP headers.
- [x] Add an opt-in public parent-page gate that fails when a staging or
      current deployment is expected to expose the live profile-manifest link
      but the deployed `busdk.com/engine/` page is still behind local source.
- [x] Add one named public deployment gate target that requires both the
      COOP/COEP iframe-eligible parent page and the live profile-manifest link,
      so staging or future hosting can be validated without remembering both
      opt-in environment variables.
- [x] Extend the public deployment gate so it also validates whether the
      published `/beos` release can be embedded by the parent origin, catching
      the cross-origin CORP case where parent COOP/COEP alone is not enough.
- [x] Make the normal public parent-page check report release embeddability as
      an informational status line, so the next hosting issue remains visible
      even while the current GitHub Pages parent still fails earlier COOP/COEP
      activation.
- [x] Align local goal/status artifacts and the Engine source checker with the
      updated `iframe, landing, and manifest` ownership wording, including a
      minimal landing-surface assertion for the Bus Engine product page.
- [x] Extend the public parent-page check to assert the deployed Engine landing
      page still exposes the Bus Engine OS browser-hosted release section, while
      leaving the newer profile-manifest link under the existing publish-drift
      gate until deployed.
- [x] Make the public parent-page check print a landing status line after the
      deployed Engine landing release section and CTA assertions pass, so
      iframe, landing, and manifest state are all visible in command output.
- [x] Add a narrow static-host header configuration for the Bus Engine product
      page plus a local check, so a header-capable deployment has the exact
      COOP/COEP parent-page policy needed for the published `/beos` iframe
      without moving Engine-owned artifacts into `busdk.com`.
- [x] Add a local `wasm-virtual-server` bundle-surface check that verifies the
      repository static bundle manifest and boot harness still expose the
      graphical display/profile fields expected by the current browser release
      supervision lane.
- [x] Document the Bus Engine OS iframe and manifest supervision targets in
      `README.md` so operators can distinguish offline website checks, live
      release-host checks, public parent-page checks, and the transitional
      static exporter.
- [x] Add a compact root `STATUS_UPDATE.md` for the current Bus Engine OS
      browser-release website lane, covering the two open PLAN items, live
      iframe/header status, hosting-versus-code blockers, and dispatch split
      without secrets.
- [x] Add a local check for `STATUS_UPDATE.md` so the manager-facing status
      artifact stays aligned with the two open PLAN items, current live
      iframe/header state, manifest-shape status, and dispatch boundary.
- [x] Add a compact root `GOAL.md` for the current Bus Engine OS browser-release
      website lane, and include it in the local manager-status consistency
      check so the goal, status, and two open PLAN items cannot drift apart.
- [x] Make the networked Bus Engine OS supervision checks configurable for
      staging or future release hosts without editing checker source.
- [x] Make the live Bus Engine OS release supervision check report whether the
      published profile manifest has explicit profile metadata or only the
      current path-implied `virtual-server` shape, so the remaining
      manifest-shape blocker is visible in command output.
- [x] Add one aggregate Bus Engine OS website supervision target that runs the
      local bundle, local Engine overview, live release-host, and deployed
      public parent-page checks in the intended order.
- [x] Expose Makefile variables for Bus Engine OS release URL, profile path,
      and public Engine parent URL so the aggregate supervision target can run
      against staging or future hosting without shell-specific environment
      setup.
- [x] Make the local Engine overview source checker honor the same Bus Engine
      OS release URL and profile path variables as the live/public supervision
      checks.
- [x] Extend the live release-host supervision check to compare key published
      artifact response sizes against the browser-hosted manifest without
      downloading large VM payloads.
- [x] Extend the live release-host supervision check from required runtime
      roles to every file listed in the browser-hosted manifest, including
      license, notice, and source-material surfaces.
- [x] Add bounded fetch timeouts to the networked Bus Engine OS supervision
      checks so a stalled release host or parent page fails with a clear
      diagnostic instead of hanging the aggregate target.
- [x] Make the repository-local Bus Engine OS iframe harness select runtime
      display, device, resolution, keyboard, and guest metadata from a named
      manifest profile instead of only the top-level manifest fields.
- [x] Add a compact manifest-driven profile selector/status to the
      repository-local Bus Engine OS iframe harness so profile choice is visible
      in the UI and planned profiles cannot boot against the current
      virtual-server artifacts by accident.
- [x] Add a COOP/COEP-aware live iframe slot to the Bus Engine overview that
      targets the published `https://dev.hg.fi/beos/` release when the parent
      page is cross-origin isolated, while preserving the direct launch preview
      fallback on GitHub Pages.
- [x] Switch the Bus Engine overview release surface from the transitional
      repository-local bundle path to the published browser-hosted release at
      `https://dev.hg.fi/beos/`, keep the copy aligned with the RISC-V 64
      virtual-server proof, and verify the static site still passes local
      quality checks.
- [x] Replace the GitHub Pages-hosted live iframe with a direct launch preview
      after browser proof showed the iframe cannot be cross-origin isolated
      from the current `busdk.com/engine/` parent page.
- [x] Rename the Bus Engine OS QEMU/WASM virtual-server website surface to
      `wasm-virtual-server`: the live iframe app, static bundle source path,
      generated public path default, and current evidence now use the
      product-shaped name; obsolete lab-era wording is limited to compatibility
      aliases outside the active website surface.
- [x] Move browser-hosted Bus Engine OS publishing out of `busdk.com` and onto
      the Engine-owned development host, currently `https://dev.hg.fi/beos/`:
      DoD is a
      deployable browser client and QEMU/WASM artifact bundle served from the
      Engine host with required COOP/COEP/CORP headers, release/source material
      files, manifest, and smoke-tested boot path; `busdk.com` keeps a direct
      launch preview and product copy that points at the hosted Bus Engine OS
      virtual server URL until the parent product page can be served with the
      COOP/COEP headers required for a true live iframe.
- [ ] Update the Bus Engine OS iframe from terminal-only boot output to an
      interactive graphics surface: use the QEMU/WASM SDL display path, keep
      serial diagnostics visible, focus a canvas for keyboard input, pass the
      display device and resolution from the manifest, and keep the screenshot
      fallback for unsupported browsers or artifacts.
      - Status: website code now has a guarded iframe slot, direct fallback,
        local graphical bundle harness, canvas focus path, serial diagnostics,
        manifest-driven display fields, profile selector/status, optional
        `docs/_headers`, and checks for the published `display=wasm` /
        `displayDevice=virtio-gpu-pci` release.
      - Blocker: the current deployed `https://busdk.com/engine/` parent page
        still reports `iframe_state=fallback-required` because GitHub Pages
        does not serve the required COOP/COEP headers. A header-capable host or
        proxy must publish the parent page, then the staging/public check should
        run with `BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1`.
      - Dispatch: website lane owns parent-page markup, fallback behavior,
        deploy-header config, and supervision checks. Engine/BEO owns the
        hosted `/beos` release artifacts, runtime, and any guest/runtime fixes
        needed for graphical keyboard proof.
- [ ] Add a browser-hosted OS manifest shape that can describe `virtual-server`
      and `virtual-desktop` profiles separately enough for the iframe to choose
      a graphical or console-oriented runtime without hard-coding the page.
      - Status: the repository-local static bundle manifest has separate
        `virtual-server` and planned `virtual-desktop` profile blocks, and the
        local iframe harness reads profile runtime/guest fields from the
        manifest. The live `/beos/virtual-server/browser-hosted-manifest.json`
        currently passes website supervision as `profile_path=virtual-server`
        with `manifest_profile_shape=path-implied`; the fetched manifest is
        still the older `generated_at=2026-07-06T13:10:16Z` export and does
        not expose `profile`, `id`, `name`, or `profiles[]` identity fields.
      - Blocker: BEO/BusDK now has explicit browser-hosted manifest profile
        identity/gates, but the rebuilt/exported `/beos` artifact is not yet
        visible at the live route. Website checks accept the current release
        shape but report the path-implied status so the remaining export and
        release-host boundary is visible.
      - Next gate: `make engine-beos-release-profile-gate` must pass once the
        rebuilt/exported `/beos` artifact is expected to expose explicit
        profile metadata.
      - Dispatch: website lane owns consuming and checking explicit profile
        metadata when it appears. Engine/BEO owns rebuilding/exporting the
        hosted `/beos` manifest with explicit profile metadata and future
        `virtual-desktop` artifacts.
- [x] Build a publishable browser-hosted Bus Engine OS static bundle with one command
      that accepts an output directory, writes the iframe page, CSS, boot
      script, manifest, preview image, QEMU/WASM runtime, Bus Engine OS guest
      artifacts, firmware, generated license indexes, third-party notices, and
      source-material payloads for shipped license-obligation packages into
      that directory, and keeps the product overview iframe integration
      independent of local artifact staging.
- [x] Embed Bus Engine OS on the Bus Engine product website as a live
      screenshot-like QEMU/WASM preview: use the proven 64-bit QEMU WASM
      runtime and Bus Engine OS guest artifacts or a documented artifact
      reference, add a static-site component under `docs/engine/` that boots in
      browser when supported, keep a screenshot fallback, document required
      COOP/COEP hosting headers, resolve the public GPL corresponding-source
      delivery shape for the QEMU runtime, run website and browser smoke
      checks, and commit the website work.
- [x] Update Bus Engine product pages for the accepted Bus Engine OS
      virtual-server profile, supported macOS arm64 and Linux 64-bit host
      paths, GUI profile development status, and clearer technical usage
      guidance for evaluators.
- [x] Rewrite the Bus Engine commercial pages around the Founding Development Preview, Bus Engine OS as the rolling Linux distribution, Codex App Server architecture, separate AI/model fees, current implementation status, target support boundaries, and customer-only source/compliance wording.
- [x] Rewrite Bus Engine image wording so Bus Engine OS is the product focus and Debian images are described as an optional compatibility input without contrast-formula copy.
- [x] Rewrite the Bus Engine overview to explain that it is a Linux distribution for virtualized Bus workloads, show `bus engine` runtime examples, and describe artifact/profile-based base image, Bus Engine OS, kernel, and component version selection.
- [x] Fix shared WASM-rendered BusDK product navigation so relative top and side nav bases resolve against the current page URL, covering `/engine/`, `/services/`, and other product pages.
- [x] Clarify that Bus Engine OS is the Bus-built Linux distribution and Debian cloud images are optional compatibility inputs.
- [x] Update Bus Engine OS promote examples so `--workspace <workspace>` is no longer shown as required.
- [x] Clarify that customers can buy commercial support to harden Bus Engine OS toward production readiness for specific use cases.
- [x] Remove old lab-era wording as visible Bus Engine product wording; use
      Bus Engine OS QEMU/WASM virtual-server wording for the active static
      surface.

# Current evidence

The Bus Engine product line now has a published site under `docs/engine/`, with
overview, modules/capabilities, documentation, and contact navigation. The
product index links to Bus Engine as one line, the shared BusDK nav data
includes the `engine` product id, the nav regression test covers the product
links, and the Bus UI demo WASM navigation asset has been rebuilt. The overview
positions Bus Engine as the BusDK product for AI-powered Linux system
engineering.

The browser-hosted Bus Engine OS preview currently has one transitional
static-bundle command in this website repository:
`make engine-wasm-os-static`, backed by
`scripts/write-engine-wasm-os-static.sh OUTPUT_DIR`. It writes the iframe page,
CSS, boot script, manifest, preview image, QEMU/WASM runtime, guest artifacts,
firmware, `README.txt`, and an `iframe.html` embed snippet into the requested
directory. The generated iframe path defaults to
`/engine/wasm-virtual-server/` and can be set with
`BUS_ENGINE_WASM_OS_PUBLIC_PATH`. The target architecture is an
Engine-owned host, such as `https://engine.busdk.com/`, with a header-capable
Bus Engine product page consuming the hosted Bus Engine OS virtual server
through an iframe. The current GitHub Pages-hosted `busdk.com/engine/` page
uses a direct launch preview until the parent page can send COOP/COEP headers.
The transitional website-hosted bundle should be described as Bus Engine OS
running as a virtual server on QEMU/WASM.

The Bus Engine overview page now points at the published development release at
`https://dev.hg.fi/beos/` instead of the transitional repository-local
`wasm-virtual-server` path. A 2026-07-06 public Chrome probe showed the direct
release is cross-origin isolated with `SharedArrayBuffer`, but the GitHub
Pages-served `busdk.com/engine/` parent is not cross-origin isolated, and the
framed release fails with `Error: cross-origin isolation is required for pthread
WebAssembly`. The overview therefore uses a direct launch preview until the
parent product page moves behind COOP/COEP-capable hosting or proxying.

The 2026-07-06 public release at `https://dev.hg.fi/beos/` boots the RISC-V 64
`virtual-server` profile to `event=ready state=multi-user` in the default
browser smoke at `420424ms`. The Engine overview now includes an iframe element
for that published release and assigns its `src` only when the parent page is
already cross-origin isolated; under GitHub Pages it keeps the direct launch
preview visible.

The website repository now has an explicit live-release supervision check:
`make engine-beos-release-check` runs
`scripts/check-engine-beos-release.mjs`, validates the public
`https://dev.hg.fi/beos/` COOP/COEP/CORP headers, reads
`virtual-server/browser-hosted-manifest.json`, and verifies the current
browser-hosted manifest shape for `riscv64`, `display=wasm`,
`displayDevice=virtio-gpu-pci`, QEMU JavaScript/WASM, kernel, rootfs, browser
runtime, checksums, sizes, and default runtime parameters. On 2026-07-07 the
check passed against the live release with 14 manifest files. The Bus Engine
overview links to the profile manifest as part of the public release surface.
The same check now reports the live profile identity surface separately:
`profile_path=virtual-server` and, for the current published manifest,
`manifest_profile_shape=path-implied`. This keeps the active website checks
green for the current release while making the remaining Engine-owned manifest
evolution visible when operators ask whether `virtual-server` and
`virtual-desktop` are explicitly described by the published manifest.
The check now also fetches `virtual-server/`, verifies the profile page is
served with isolation headers, confirms the page has the graphical canvas,
serial diagnostics, serial input, generated `display=wasm` and
`displayDevice=virtio-gpu-pci` defaults, and verifies the manifest-advertised
browser runtime is reachable with isolation headers and contains the expected
runtime hooks for isolation checks, QEMU argument generation, virtio GPU
display, and serial chardev input.
The live release check also compares required published artifact sizes against
the manifest: QEMU JavaScript/WASM, kernel, rootfs, browser runtime, and index
surface. Large payloads are checked with HEAD `Content-Length`; the profile
index is checked from the already-fetched body because the current host does
not send `Content-Length` for `virtual-server/index.html` HEAD.
The check now walks every file listed in the live browser-hosted manifest,
including license, notice, and source-material surfaces. Small files whose
HEAD response omits `Content-Length` are fetched and compared by body size;
large VM payloads still require manifest-matching HEAD sizes. On 2026-07-07
the live check passed with `published_files=14`.
The normal local `make quality` target now starts with
`scripts/check-engine-wasm-bundle-surface.mjs` and
`scripts/check-engine-beos-surface.mjs`. The bundle check verifies the
repository-local `docs/engine/wasm-virtual-server/` manifest keeps the
`virtual-server` and `virtual-desktop` profiles, graphical SDL display fields,
800x600 resolution, keyboard-enabled canvas, screenshot fallback, serial
diagnostics, and browser isolation checks. The Engine overview check verifies
that the page still has the guarded iframe slot, direct-launch fallback,
manifest link, cross-origin-isolation permission request, and
`SharedArrayBuffer` guard for the current release URL before the broader GX/UI
documentation checks run.
The repository now includes an optional `docs/_headers` file for static hosts
that understand Netlify/Cloudflare Pages-style header configuration. It applies
`Cross-Origin-Opener-Policy: same-origin` and
`Cross-Origin-Embedder-Policy: require-corp` only to `/engine/` and
`/engine/*`, leaving the Engine-owned `/beos` release artifacts outside
`busdk.com`. `scripts/check-engine-hosting-headers-config.mjs` verifies that
scope and policy, and `make engine-beos-check` now runs it before the live
release-host and deployed public parent-page checks. This is deployment input
for a header-capable host or proxy; it does not change GitHub Pages behavior by
itself.
The repository-local `docs/engine/wasm-virtual-server/` iframe harness now
renders a manifest-populated profile selector and selected-profile status. The
selector reads `manifest.profiles`, reloads the page with `?profile=...`, and
keeps boot disabled when the selected profile is not `current-artifact`, so the
planned `virtual-desktop` profile stays visible without booting against the
current `virtual-server` artifact set. `make engine-wasm-bundle-surface-check`
now verifies the profile selector, profile status, manifest-driven profile
population, and planned-profile boot gate.
The public deployed-page check `make engine-beos-public-page-check` fetches
`https://busdk.com/engine/`, verifies that the live parent page still targets
`https://dev.hg.fi/beos/` through the guarded iframe and direct-launch
fallback, and reports the current parent iframe state. On 2026-07-07 it passed
with `iframe_state=fallback-required`, meaning the release URL is present but
the live GitHub Pages parent still lacks the isolation headers required for
automatic iframe activation. The new profile-manifest link is present in the
local website source and reports `manifest_link=not-yet-deployed` until the
website changes are published.
The public parent-page checker also has an opt-in deployment gate:
`BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1`. In normal mode it continues to report
the current GitHub Pages fallback state. In require mode it fails if the parent
page lacks COOP/COEP; on 2026-07-07 that negative check failed against
`https://busdk.com/engine/` with the expected diagnostic that the page is not
iframe-eligible. Use this mode for staging or future header-capable hosting
where the live iframe is expected to activate.
The public parent-page checker also has an opt-in publish-drift gate:
`BUS_ENGINE_REQUIRE_MANIFEST_LINK=1`. In normal mode it continues to report
`manifest_link=not-yet-deployed` for the current public page. In require mode
it fails if the parent page does not link the expected profile manifest; on
2026-07-07 that negative check failed against `https://busdk.com/engine/` with
the expected diagnostic naming
`https://dev.hg.fi/beos/virtual-server/browser-hosted-manifest.json`.
`make engine-beos-public-deploy-gate` combines both parent-page requirements
for staging or future public hosting by requiring COOP/COEP iframe eligibility
and the live profile-manifest link. On 2026-07-07 it failed against the current
`https://busdk.com/engine/` page with the expected COOP/COEP diagnostic.
The public deploy gate now also fetches the release URL in require-iframe mode
and checks whether a cross-origin parent can embed it. A one-off local loopback
probe with an isolated parent on one origin and a release response on another
origin sending `Cross-Origin-Resource-Policy: same-origin` failed with the
expected diagnostic: the release needs `CORP: cross-origin` or a same-origin
parent/proxy path.
`README.md` now documents the offline `make quality` Engine checks, the
networked `make engine-beos-release-check` and
`make engine-beos-public-page-check` supervision targets, the expected
`fallback-required` GitHub Pages state, and keeps the transitional
`make engine-wasm-os-static` exporter clearly separate from the published
release-host path.
The networked checkers can be pointed at staging or future release hosts with
`BUS_ENGINE_BEOS_RELEASE_URL`; profile paths default to `virtual-server/` and
can be changed with `BUS_ENGINE_BEOS_PROFILE_PATH`. The public parent-page
checker also accepts an optional Engine page URL argument for non-production
parent pages.
`make engine-beos-check` is now the aggregate supervision target for this
surface. It runs the repository-local bundle check, the local Engine overview
check, the live release-host check, and the deployed public parent-page check
in that order.
The aggregate and networked Make targets accept `BUS_ENGINE_BEOS_RELEASE_URL`,
`BUS_ENGINE_BEOS_PROFILE_PATH`, and `BUS_ENGINE_PUBLIC_PAGE_URL` as Make
variables, so staging or future host checks can be run with a single
script-friendly command while keeping the default production values unchanged.
The local Engine overview source checker now uses the same
`BUS_ENGINE_BEOS_RELEASE_URL` and `BUS_ENGINE_BEOS_PROFILE_PATH` values as the
live release and public parent-page checks, so a future host migration can be
verified through one consistent Make-variable path.
The networked release and public parent-page checkers now apply a bounded
per-request timeout controlled by `BUS_ENGINE_CHECK_TIMEOUT_MS`, defaulting to
30000 ms. Timeout failures report the URL and timeout value instead of letting
the aggregate supervision target hang indefinitely.
The repository-local `docs/engine/wasm-virtual-server/manifest.json` now gives
`virtual-server` and `virtual-desktop` separate profile blocks with guest and
runtime metadata. The local iframe boot harness selects the active profile from
`?profile=` or the default guest profile, then reads display mode, display
device, resolution, keyboard support, expected guest text, and kernel append
from that profile with top-level compatibility fallbacks. The
`engine-wasm-bundle-surface-check` target now verifies that profile-aware
contract.

Buyer-facing source access copy now uses Git access wording instead of naming a
specific Git hosting provider. Historical blog references to the public site
and hosting remain unchanged.

Bus Engine product pages now foreground the story that the product provides a
source-configurable Linux system operated by an AI Linux engineer.

Bus Engine product pages avoid self-referential product-taxonomy wording and
use direct buyer-facing sentences instead.

The Bus Engine pages pass the new `bus-lint` HTML public marketing rubric with
`--agent codex`.

Bus Engine pricing now uses a product-local Founding Technical Preview page:
a limited founding offer, source access, unlimited test systems, monthly and
one-time founding options, account/support boundaries, no production SLA, no
AI credits, and clear separation between commercial account entitlements and
software license terms.

The module-owned Bus UI catalog now drives `docs/gx-ui/` coverage. The
deterministic `scripts/check-gx-ui-component-pages.sh` check verifies every
implemented catalog entry has a GX/UI reference page, and every implemented
component entry has a live demo hook plus the shared demo loader.

The Bus UI `make component-test-coverage` target verifies implemented catalog
symbols still exist as exported Go API and implemented component entries have
package-local test evidence for preferred public symbols. It also verifies
that exported public Bus UI APIs are either cataloged for GX/UI docs or
classified as non-element runtime, DTO, infrastructure, test, or tooling API.

The BusDK downstream adopter audit reports UI/GX adopters, fails on forbidden
production `pkg/uikit` use, and currently reports zero local UI candidate files
after moving the last generic DOM helper behavior into shared `bus-ui`.

Bus Engine pricing now embeds the live Stripe pricing table on
`docs/engine/pricing/index.html` and no longer repeats stale static renewal or
founding-price rows outside the purchase flow.

Bus Engine public copy now presents the first release as a limited paid
Founding Technical Preview for source-access operators. The binary-only Runtime
plan is deferred until Public Beta.

The Bus Engine overview now describes Bus Engine OS as the Bus-built Linux
distribution managed by AI agents. Debian cloud images are described as an
optional compatibility input for preview workflows.

Bus Engine runtime image and Debian compatibility details now live on a
dedicated Images page linked from the Engine overview, top navigation, side
navigation, and generated Bus UI navigation data.

Bus Engine public pages now remove the 1 July 2026 preview-start language,
describe the accepted `virtual-server` Bus Engine OS profile, name macOS arm64
and Linux 64-bit operator host paths, call out `virtual-gui` as an in-progress
GUI profile, and show the normal build/promote/start/status/SSH command flow.
The promote examples now use `bus engine os artifact promote-engine` as the
default command, with workspace selection treated as optional.
The server workload image row now describes package, service, boot-expectation,
and test configuration as part of the current accepted `virtual-server`
blueprint workflow.
The Bus Engine overview now says corresponding source for covered customer
release binaries is provided through the customer release area at no extra
charge.
The production-readiness copy now keeps the preview boundary while saying
commercial engineering and support can be purchased to harden, validate, and
maintain Bus Engine OS for a specific customer use case.

Bus Engine Stripe Products and product pages now describe the license split:
monthly access includes the Bus Engine product-line codebase under FSL with
MIT or Apache 2.0 conversion after two years; the one-time option includes the
current Bus Engine product-line codebase under MIT or Apache 2.0 at purchase
plus one year of FSL-licensed updates; third-party software keeps its own
licenses; and FSL applies only to Bus-related code licensed by us.
