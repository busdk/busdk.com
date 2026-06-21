# PLAN.md

# Current work

- [ ] Complete the final GX/UI reusable-element completion audit by comparing
  the module-owned catalog against exported public reusable UI APIs, deciding
  whether any intentionally exported helpers should be added to the catalog or
  explicitly classified as non-element runtime/infrastructure API.

# Current evidence

The module-owned Bus UI catalog now drives `docs/gx-ui/` coverage. The
deterministic `scripts/check-gx-ui-component-pages.sh` check verifies every
implemented catalog entry has a GX/UI reference page, and every implemented
component entry has a live demo hook plus the shared demo loader.

The Bus UI `make component-test-coverage` target verifies implemented catalog
symbols still exist as exported Go API and implemented component entries have
package-local test evidence for preferred public symbols.

The BusDK downstream adopter audit reports UI/GX adopters, fails on forbidden
production `pkg/uikit` use, and currently reports zero local UI candidate files
after moving the last generic DOM helper behavior into shared `bus-ui`.
