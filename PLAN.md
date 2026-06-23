# PLAN.md

# Current work

- [x] Apply the new `bus-lint` HTML rubric findings to the Bus Engine modules page so its capability overview uses buyer-facing language and links the stated evaluator path.

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
