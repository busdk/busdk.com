# AGENTS.md

## Highest-Priority Rule: Precise Language And Precise Reasoning

Precise language is part of precise reasoning. Before acting or reporting, name
the exact object, action, scope, evidence, and uncertainty. State what changed
and what did not change. Never use a broader claim than the evidence supports.
If an exact, unambiguous sentence cannot be written, inspect the evidence or
ask for clarification before proceeding.


Published-safe guidance for the GX/UI documentation subtree.

## Scope

Apply this file to every page under `docs/gx-ui/`. More specific
`AGENTS.md` files in child directories add local rules for their subtree.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Unified GX/UI Documentation

1. Keep the GX/UI docs visually and structurally unified. Pages in the same
   docs family should share the same side navigation pattern, page frame,
   heading scale, and link style.
2. Treat the Shells section as the reference shape for Bus UI component
   families: one family hub, one focused leaf page per public component, and
   a nested side-nav entry for every child page.
3. Every GX/UI subtree page should render inside the docs layout with
   `.gx-doc-layout`, `.gx-side-nav`, and `.gx-doc-content` unless the page is a
   top-level landing page with an intentional different layout.
4. Use relative same-site links with explicit `index.html` targets inside this
   subtree. The pages must work from a local static server and from the public
   site.
5. Exactly one side-nav link should mark the current page with
   `aria-current="page"`. Leaf pages should mark the leaf entry, not only the
   family hub.
6. Do not create duplicate mini-navigation groups for a family when that
   family can be represented in the shared GX/UI side nav.
7. Keep hero and page headings at documentation scale. Do not let GX/UI docs
   pages inherit oversized marketing hero typography.
8. Public docs pages and public guidance files must follow the same
   publication-safety boundary as this file.

## Content Quality

1. Page copy should name the reader-visible API or authoring concept directly.
   Avoid vague placeholder language and generated copy fragments.
2. API tables and examples must match the currently documented public surface.
   If a public helper exists, name it directly; if it is only a lower-level
   implementation detail, do not pretend it is the reader-facing API.
3. Every Bus UI component or component-family page that shows a Go API example
   must also show a `.gx source` example for the same visible use case.
   `.gx source` examples should look like TSX-style Go markup: intrinsic tags,
   nested children, `<Text value={...}></Text>`, and uppercase local component
   tags with lower-camel props. Do not satisfy this requirement with a pure Go
   wrapper around `ui.*`, `gx.Element` slices, or generated-Go render-boundary
   code.
4. When a page changes navigation or layout, verify local links for the edited
   subtree and inspect at least one representative rendered page when practical.
