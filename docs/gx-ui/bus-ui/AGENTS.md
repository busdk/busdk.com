# AGENTS.md

Published-safe guidance for Bus UI documentation pages.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Bus UI Family Shape

1. Bus UI families should follow the Shells pattern:
   one family `index.html` page plus one focused child page for each public
   component or component-like helper.
2. The family page introduces the family, lists the public APIs, and links to
   child pages with explicit `index.html` hrefs.
3. Child pages document one component or helper boundary. Avoid bundling
   several distinct public components into one "family" leaf unless the public
   API cannot honestly be split.
4. Family and child pages should use the shared Bus UI Library section in the
   GX/UI side nav. The family link uses `gx-side-nav-child`; child links use
   `gx-side-nav-grandchild`.
5. Keep families such as Assistant and Terminal separate when their subject
   matter differs, but merge their navigation into the shared Bus UI Library
   side nav instead of adding standalone "Assistant UI" or "Terminal UI"
   groups.
6. Do not leave old concept-page wording when the page has become a component
   family hub. Prefer "Public APIs", "Components", and concrete component
   names over "concept pages" or vague back-link sections.
7. Remove generated copy fragments when found, especially stray phrases that
   begin mid-sentence or mention "Any remaining Use the public package".

## Side Navigation

1. Keep the Bus UI side nav complete across Bus UI pages. If a family gains or
   loses child pages, update the repeated side-nav blocks consistently.
2. Each family page should mark the family entry as current. Each child page
   should mark its own child entry as current.
3. Preserve already-integrated families when editing a narrow section. Do not
   flatten Forms, Data, Assistant, Terminal, or Shells while changing another
   Bus UI subtree.
4. Prefer explicit, relative `index.html` links in side nav entries so the
   docs work in local previews and public hosting.
