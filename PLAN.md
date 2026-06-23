# PLAN.md

# Current work

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

# Current evidence

The Bus Engine product line now has a published site under `docs/engine/`, with
overview, modules/capabilities, documentation, and contact navigation. The
product index links to Bus Engine as one line, the shared BusDK nav data
includes the `engine` product id, the nav regression test covers the product
links, and the Bus UI demo WASM navigation asset has been rebuilt. The overview
positions Bus Engine as the BusDK product for AI-powered Linux system
engineering.

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
25-50 named operators, source access, unlimited test systems, monthly and
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
