# AGENTS.md

Published-safe guidance for Bus UI Assistant documentation.

## Publication Safety

This file is inside the public docs tree. It must not contain secrets, private
paths, tokens, worker ids, internal task ids, customer data, or process-only
notes.

## Assistant Pages

1. Assistant is a Bus UI family. Keep Assistant and Terminal as separate
   subject areas, but represent Assistant inside the shared Bus UI Library side
   nav.
2. The Assistant hub should link to focused child pages for assistant shell,
   panels, composer, messages, approvals, model selection, thread lists, and
   related assistant components.
3. Assistant child pages should be nested under Assistant in the shared side
   nav and mark their own entry as current.
4. Do not reintroduce a standalone "Assistant UI" side-nav group when the same
   pages can live under the shared Bus UI Library nav.
5. Remove stale generated fragments when editing Assistant pages.
