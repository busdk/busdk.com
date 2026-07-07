# PLAN.md

# Current work

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
- [ ] Add a browser-hosted OS manifest shape that can describe `virtual-server`
      and `virtual-desktop` profiles separately enough for the iframe to choose
      a graphical or console-oriented runtime without hard-coding the page.
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
