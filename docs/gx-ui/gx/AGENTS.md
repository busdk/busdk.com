# AGENTS.md

## Highest-Priority Rule: Precise Language And Precise Reasoning

Precise language is part of precise reasoning. Before acting or reporting, name
the exact object, action, scope, evidence, and uncertainty. State what changed
and what did not change. Never use a broader claim than the evidence supports.
If an exact, unambiguous sentence cannot be written, inspect the evidence or
ask for clarification before proceeding.


Published-safe guidance for GX Framework documentation pages.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## GX Framework Side Nav

1. Every GX Framework subpage under this directory should include the shared
   GX/UI docs layout and side nav unless it is intentionally redesigned as a
   landing page.
2. The side nav should list the GX Framework siblings consistently, including
   Source files, Component functions, Events, Rendering, Runtime bridges,
   Testing, Props and children, Generated Go, Effects, and Nodes and render
   tree when those pages exist.
3. Each GX Framework subpage should mark its own side-nav entry with
   `aria-current="page"`.
4. Do not leave a GX Framework leaf as a bare single-column page when sibling
   pages use the docs side nav.
5. After changing the GX Framework nav, scan all GX Framework `index.html`
   pages for `.gx-side-nav`, `aria-current="page"`, and valid relative links.
