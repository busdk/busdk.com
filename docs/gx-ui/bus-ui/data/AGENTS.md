# AGENTS.md

Published-safe guidance for Bus UI Data display documentation.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Data Display Pages

1. Data display is a component family. Keep `data/index.html` as the family hub
   and keep one child page per data component.
2. The Data display family page should use the same Shells-style structure:
   public API table, explicit child-page links, and no redundant back-link
   block.
3. Data child pages must include the shared GX/UI side nav and mark their own
   component entry as current.
4. Keep Data display nested in the Bus UI Library side nav with its child
   components listed below it.
5. Avoid old wording that describes child pages as detached concept pages.
