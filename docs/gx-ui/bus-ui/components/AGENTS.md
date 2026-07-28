# AGENTS.md

## Highest-Priority Rule: Precise Language And Precise Reasoning

Precise language is part of precise reasoning. Before acting or reporting, name
the exact object, action, scope, evidence, and uncertainty. State what changed
and what did not change. Never use a broader claim than the evidence supports.
If an exact, unambiguous sentence cannot be written, inspect the evidence or
ask for clarification before proceeding.


Published-safe guidance for Bus UI component reference pages.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Component Examples

1. Component pages that show a Go API example must also show a `.gx` source
   example for the same component or family scenario.
2. Keep the existing Go example. The `.gx` example is an additional authoring
   view, not a replacement.
3. Label the examples clearly, for example "Go API" and ".gx source".
4. Make `.gx` examples real source syntax: uppercase component tags,
   lower-camel props, and intrinsic markup. Do not use generated-Go wrapper
   examples, `ui.*` calls, or `gx.Element` slices inside `.gx source` blocks.
5. Examples should be specific to the documented component. Avoid generic
   placeholder `.gx` snippets that do not exercise that component.
6. When a public function exists, name it directly in the page's API table and
   example. For example, an AppShell page should show `ui.AppShell` when that
   is the reader-facing API.
7. Imports, props names, helper names, and render-boundary calls in examples
   must match the current public documentation surface.

## Component Page Shape

1. Keep one focused leaf page per public component whenever the API surface
   makes that honest.
2. Component family overview pages should link to every child page and keep the
   same nested side-nav hierarchy as the child pages.
3. Do not hide component pages behind flat family-only navigation.
