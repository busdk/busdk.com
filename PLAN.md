# PLAN.md

# Current work

- [x] Build a publishable Bus Engine browser-lab static bundle with one command
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
- [x] Tighten Bus Engine copy around the concrete problem it solves, how it differs from ordinary distributions/platforms, and one example workflow.
- [x] Remove internal legal-process meta commentary from Bus Engine licensing pages while preserving customer-facing source-delivery boundaries.
- [x] Remove README-referential Bus Engine runtime-image copy and make the public artifact-flow documentation direct.
- [x] Reduce Bus Engine marketing-page Codex mentions so the product story says the agent coordinates and Codex remains a linked technical runtime detail.
- [x] Link Bus Engine product-page Codex mentions to the upstream OpenAI Codex repository.
- [x] Remove static commercial price amounts from public product pages so Stripe remains the only price source.
- [x] Rewrite the Bus Engine commercial pages around the Founding Development Preview, Bus Engine OS as the rolling Linux distribution, Codex App Server architecture, separate AI/model fees, current implementation status, target support boundaries, and customer-only source/compliance wording.
- [x] Add focused Bus Engine architecture, FAQ, and licensing pages without turning the product site into implementation notes.
- [x] Update Bus Engine navigation, homepage copy, metadata, and generated product nav data so all new pages are reachable and no Engine links lose their product prefix.
- [x] Use `docs/engine/preview.jpg` as the SEO/social preview image for the `/engine/` overview page.
- [x] Use plain `0.x.0` version examples for Bus Engine preview releases instead of `-preview.1` suffixes.
- [x] Remove prescriptive Bus Engine pricing-page workflow copy and keep FSL links in licensing text, not plan headings.
- [x] Correct Bus Engine preview source-access copy so it covers only the limited Bus Engine product line, not all BusDK.
- [x] Rewrite Bus Engine image wording so Bus Engine OS is the product focus and Debian images are described as an optional compatibility input without contrast-formula copy.
- [x] Link every published product-page FSL mention to `https://fsl.software/`.
- [x] Rewrite the Bus Engine overview to explain that it is a Linux distribution for virtualized Bus workloads, show `bus engine` runtime examples, and describe artifact/profile-based base image, Bus Engine OS, kernel, and component version selection.
- [x] Fix shared WASM-rendered BusDK product navigation so relative top and side nav bases resolve against the current page URL, covering `/engine/`, `/services/`, and other product pages.
- [x] Remove the separate Bus Engine Documentation nav item and fold documentation links into the Modules page.
- [x] Update Bus Engine AI access copy to name customer-provided OpenAI Codex access and configured local LLM model support instead of ChatGPT subscription wording.
- [x] Rewrite the Bus Engine overview around the current product reality: a customizable Linux server operating-system distribution powered and maintained by Codex AI agents, with separate Codex subscription or API access required.
- [x] Fix Bus Engine generated top and side navigation bases so runtime-rendered links keep the `/engine/` product prefix.
- [x] Adjust the Bus Engine homepage card so it better matches neighboring product-card length and describes a custom operating-system outcome, not only engineering assistance.
- [x] Apply the new `bus-lint` HTML rubric findings to the Bus Engine modules page so its capability overview uses buyer-facing language and links the stated evaluator path.
- [x] Add Bus Engine pricing information using the Named Operator License model, including unlimited systems, monthly and one-time options, support/account boundaries, and product-local pricing navigation.
- [x] Balance the frontpage card copy lengths for Bus Agentic Development, Bus AI Platform, and Bus Engine without changing other product cards.
- [x] Compact the Bus Engine product site into shorter focused pages while preserving pricing, target support boundaries, and contact paths.
- [x] Clarify on the Bus Engine site that AI credits are not included and model access can use supported customer-provided ChatGPT subscriptions or other configured providers.
- [x] Embed the live Bus Engine Stripe pricing table on the product pricing page and remove stale static price rows.
- [x] Reposition Bus Engine pricing and product copy around the limited Founding Technical Preview and defer the binary-only Runtime plan until Public Beta.
- [x] Clarify that Bus Engine OS is the Bus-built Linux distribution and Debian cloud images are optional compatibility inputs.
- [x] Split Bus Engine runtime-image and Debian compatibility details from the overview into a dedicated feature subpage.
- [x] Update Bus Engine OS promote examples so `--workspace <workspace>` is no longer shown as required.
- [x] Refresh the Bus Engine server workload status so workload package and service blueprints are no longer described as future-only work.
- [x] Make Bus Engine source-delivery copy direct: covered customer-release source is provided, not merely expected.
- [x] Clarify that customers can buy commercial support to harden Bus Engine OS toward production readiness for specific use cases.
- [x] Align Bus Engine Stripe and product-page license copy: monthly FSL source converting to MIT or Apache 2.0 after two years, one-time current codebase under MIT or Apache 2.0 plus one year of FSL updates, and third-party software retaining its own licenses.

# Current evidence

The Bus Engine product line now has a published site under `docs/engine/`, with
overview, modules/capabilities, documentation, and contact navigation. The
product index links to Bus Engine as one line, the shared BusDK nav data
includes the `engine` product id, the nav regression test covers the product
links, and the Bus UI demo WASM navigation asset has been rebuilt. The overview
positions Bus Engine as the BusDK product for AI-powered Linux system
engineering.

The Bus Engine WASM OS preview now has one maintained static-bundle command:
`make engine-wasm-os-static`, backed by
`scripts/write-engine-wasm-os-static.sh OUTPUT_DIR`. It writes the iframe page,
CSS, boot script, manifest, preview image, QEMU/WASM runtime, guest artifacts,
firmware, `README.txt`, and an `iframe.html` embed snippet into the requested
directory. The generated iframe path defaults to `/engine/browser-lab/` and can
be set with `BUS_ENGINE_WASM_OS_PUBLIC_PATH`.

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
