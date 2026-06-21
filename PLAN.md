# PLAN.md

# Current work

- [ ] Build a complete GX/UI reusable-element catalog under `docs/gx-ui/` so end users can identify, manually test, and choose the right shared component for each app surface.
- [ ] Audit `bus-ui/pkg/*` exported reusable UI components, helpers, catalogs, and testkits against existing `docs/gx-ui/` pages; record missing page coverage, stale pages, and pages without both Go API and `.gx source` examples.
- [ ] Add a deterministic docs coverage check that compares the reusable UI catalog/source inventory with `busdk.com/docs/gx-ui/` pages so missing manual-test pages are caught automatically.
- [ ] Extend automated reusable-component tests so UI apps can rely on unit-tested shared components instead of duplicating component behavior tests in each app.
- [ ] Audit Bus UI app consumers under `bus-*/internal/ui` and track local UI elements that should be replaced by reusable `bus-ui` components.
