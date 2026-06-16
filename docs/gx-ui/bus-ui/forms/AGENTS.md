# AGENTS.md

Published-safe guidance for Bus UI Forms documentation.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Forms Pages

1. Forms is a component family, not a single flat concept page.
2. Keep `forms/index.html` as the family hub and give each form piece its own
   focused child page when the public API supports it.
3. Split distinct public controls such as Input, TextInput, PasswordInput, and
   DateInput into separate child pages instead of bundling them under an
   "Input family" leaf.
4. If a form helper cannot honestly become a single-component page, say what
   public boundary it documents and keep the page title/API table truthful.
5. Forms pages must appear under the shared Bus UI Library side nav with Forms
   as a child entry and each form piece as a grandchild entry.
6. Do not reintroduce a separate Forms-only side-nav group on Forms child
   pages.
