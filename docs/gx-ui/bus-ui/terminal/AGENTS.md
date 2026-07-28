# AGENTS.md

## Highest-Priority Rule: Precise Language And Precise Reasoning

Precise language is part of precise reasoning. Before acting or reporting, name
the exact object, action, scope, evidence, and uncertainty. State what changed
and what did not change. Never use a broader claim than the evidence supports.
If an exact, unambiguous sentence cannot be written, inspect the evidence or
ask for clarification before proceeding.


Published-safe guidance for Bus UI Terminal documentation.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Terminal Pages

1. Terminal is a Bus UI family. Keep Terminal and Assistant as separate
   subject areas, but represent Terminal inside the shared Bus UI Library side
   nav.
2. The Terminal hub should link to focused child pages for terminal sessions,
   input, output, approvals, adapters, and related terminal components.
3. Terminal child pages should be nested under Terminal in the shared side nav
   and mark their own entry as current.
4. Do not reintroduce a standalone "Terminal UI" side-nav group when the same
   pages can live under the shared Bus UI Library nav.
5. Remove stale generated fragments when editing Terminal pages.
