# GOAL.md

## Active Goal

Supervise the `busdk.com` website iframe, landing, and manifest surface for
the current Bus Engine OS browser release.

## Website-Owned Scope

- Keep the Bus Engine product page aligned with the published
  `https://dev.hg.fi/beos/` browser release.
- Keep the Bus Engine landing copy aligned with the current browser release,
  fallback state, and profile-manifest path.
- Keep the parent-page iframe activation guarded by cross-origin isolation and
  `SharedArrayBuffer` availability.
- Preserve the direct-launch fallback while the public parent page remains
  fallback-only.
- Maintain `docs/_headers` for header-capable static hosts or proxies.
- Maintain local, live-release, public-parent, and manager-status checks.
- Consume explicit profile metadata from the published virtual-server release.

## Out Of Scope For This Repository

- Do not edit Bus Engine OS, QEMU, or hosted `/beos` runtime artifacts here.
- Do not move Engine-owned release artifacts into `busdk.com`.
- Do not publish secrets, private tokens, or private implementation notes.

## Active Gates

- `make engine-beos-check` must pass.
- `make engine-status-update-check` must pass.
- `make engine-beos-release-profile-gate` must pass for the published
  virtual-server release.
- `make engine-beos-public-page-check` may report
  `iframe_state=fallback-required` for the current GitHub Pages deployment.
- `make engine-beos-public-page-check BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1`
  is the gate for staging or future header-capable parent hosting.
- `make engine-beos-public-page-check BUS_ENGINE_REQUIRE_MANIFEST_LINK=1` is
  the gate for a deployment expected to include the profile-manifest link.
- `make engine-beos-public-deploy-gate` is the combined fail-closed gate for a
  staging or future public parent page that is expected to satisfy both
  conditions and embed the release from the parent origin.

## Current Open Conditions

- Public iframe activation is not complete until a header-capable
  `busdk.com/engine/` deployment passes `BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1`
  and the release is served either from the same origin/proxy path or with
  `Cross-Origin-Resource-Policy: cross-origin`.
- Published virtual-desktop profile metadata is not complete until Engine/BEO
  publishes a virtual-desktop artifact and the website live checks can verify
  its explicit profile shape separately from `virtual-server`.

## Completion Evidence

The website lane is complete only when the two open `PLAN.md` items are closed
with command evidence, `STATUS_UPDATE.md` reports no fallback-only public
parent state for the active release, and the dispatch boundary with Engine/BEO
remains clear.
