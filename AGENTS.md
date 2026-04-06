# AGENTS.md

Scoped guidance for the `busdk.com` repository.

## Scope And Precedence

1. Apply this file to the whole `busdk.com` repository.
2. If instructions conflict, use this order:
   1. Repository identity and safety constraints.
   2. Definition of done and quality gates.
   3. Documentation and implementation conventions.
   4. Task-specific user instructions.
3. Prefer minimal, deterministic, script-friendly behavior.

## Repository Identity (`busdk.com`)

1. This repository hosts the `busdk.com` website and related static documentation assets.
2. Keep implementation focused on website/docs content for this repository only.
3. Do not introduce unrelated build systems, services, or runtime dependencies unless explicitly requested.

## Definition Of Done

1. Keep changes scoped and deterministic.
2. Update documentation/content in the same change set when behavior or messaging changes.
3. Preserve backward-compatible site behavior unless explicitly asked to change it.
4. If tests/checks exist for changed areas, run them and ensure they pass.
5. If no automated checks exist, verify impacted pages/assets directly and report what was validated.
6. When the user gives new instructions, create or update `busdk.com/PLAN.md`
   first so current work and still-open earlier work stay visible as unchecked
   items until they are actually finished.
7. Whenever research into project history, module behavior, docs, or source
   uncovers a clearly blog-worthy future topic, add it immediately to
   `busdk.com/PLAN.md` as an unchecked item with the most natural article date
   for that topic. Do not leave promising article ideas only in conversation.

## Documentation Rules

1. Keep `busdk.com/docs/` content focused on product/commercial landing-page communication for buyers and evaluators.
2. Prefer short paragraphs and avoid repetitive wording.
3. Use lists/tables only when they materially improve clarity.
4. When mentioning the `bus` GitHub repository, inline-link to `https://github.com/busdk/bus`.
5. Product pages on `busdk.com` may include commercial/landing messaging and product-audience framing when it improves clarity for buyers/evaluators.
6. When linking from `busdk.com/docs/` to documentation pages, always use canonical `https://docs.busdk.com/...` URLs.
7. Do not link directly to private SDD pages from `busdk.com/docs/`; you may mention that private implementation design specifications exist.
8. Within published `busdk.com/docs/` pages, keep same-site navigation links relative so they work in local file or local server development as well as on the public site. Do not use absolute `https://busdk.com/...` URLs for header/home links inside subtree pages when a correct relative link exists.

## Publication Boundary

1. `busdk.com/docs/` is a published static site tree. Do not place `AGENTS.md`
   files anywhere under that tree.
2. Keep durable agent instructions for the `busdk.com` repo in non-published
   paths such as `busdk.com/AGENTS.md`.
3. If a subtree under `busdk.com/docs/` needs extra guidance, add a clearly
   named section for that subtree in this root file instead of creating a new
   published `AGENTS.md`.

## Blog Rules (`busdk.com/docs/blog/`)

1. The blog is one reverse-chronological feed of standalone articles. Do not
   create overview hubs, phase pages, or summary pages unless explicitly asked.
2. Posts must focus on what changed, why it matters, and how the change shows
   up in practice. Do not foreground Git history, trackers, reopen/close loops,
   or the fact that the article is retrospective.
3. Use dates as quiet metadata only. Do not center the prose around "this day",
   "same day", or the publishing format.
4. Write one shared article flow for all readers. Do not add audience lanes,
   SMB/developer split sections, or "choose your path" copy.
5. In Finnish prose, prefer Finnish terms whenever a natural Finnish term
   exists. Keep English only for exact command names, flags, code identifiers,
   product names, or quoted UI labels.
6. Every post needs:
   1. clear H1
   2. short ingress
   3. distinct metadata row
   4. `TL;DR` that stands on its own
   5. one flowing body section after `TL;DR`
7. The body must contain concrete reader-visible substance: a real command,
   recognizable before/after behavior, output difference, report effect, UI
   effect, or another practical sign of the change.
8. Do not leave abstract claims unexplained. If the reader cannot tell what a
   phrase like "row description is preserved" means in practice, rewrite it
   with a concrete symptom or example.
9. Prefer real Bus command examples over pseudocode. Every published Bus
   command example must be verified in the exact written form, including `\`
   continuation lines and flag order.
10. Do not show command syntax alone when the reader also needs to understand
    what the command returns. Include the essential part of the output, result,
    or success condition so the example explains both invocation and meaning.
11. The output excerpt does not need to be complete, but it must contain the
    part that helps the reader recognize the behavior being discussed.
12. If output varies by workspace or data, show the stable relevant fragment or
    state the real success condition instead of inventing a fake fixed output block.
13. When a command is primarily a check or assertion surface, and its exit
    status is part of the command's core meaning, say so explicitly. For
    example, note when `0` means the assertion passed and non-zero means the
    check failed.
14. Use command wrapping with `\` only at semantically sensible boundaries, and
    keep `\` as the final character on the line with no trailing spaces.
15. Research may use Git history, `git worktree`, module source trees under
    `./bus` and `./bus-*`, end-user docs under `./docs/docs/`, private SDD docs
    under `./sdd/docs/`, help text, and tested CLI behavior.
16. Public blog copy must not reveal private module internals, proprietary
    implementation details, hidden architecture choices, or other business
    secrets.
17. When a post introduces or expands a user-facing command, report, module
    surface, or other feature that has public docs, include a semantic inline
    link to the relevant `https://docs.busdk.com/...` page at the point where
    the reader would naturally want more detail.
18. The docs link is a continuation path, not a substitute for explanation.
    The blog post must still explain the feature properly in its own text with
    concrete examples, effects, and output cues instead of reducing the post to
    a short teaser plus link.
19. Do not explain the writing strategy to the reader. Avoid sentences that
    justify why the post uses examples, why something is "opened up here", or
    why the text was structured in a certain editorial way. Keep only product
    meaning, user-visible behavior, and relevant next steps.
20. Do not mention internal verification work in reader-facing prose. Avoid
    sentences like "this command was tested as written" or other statements
    about how the article was validated. Validation belongs to the writing
    process unless the verification result itself is user-visible product news.
21. If research shows stale public docs under `./docs/docs/`, fix them in the
    same turn when practical. If research uncovers a current defect or
    misleading docs/help mismatch that should not be silently worked around,
    add a concrete root-level entry to `BUGS.md` or `FEATURE_REQUESTS.md`.
22. For Bus command checks and blog-example verification under this repo, use
    `busdk.com/tmp/` as the default writable local test workspace when possible.
23. Prefer running examples in `busdk.com/tmp/` without extra permissions over
    ad hoc temp locations elsewhere, unless a module-specific fixture is clearly
    better for the command being tested.
24. Treat blog writing as an iterative process. Draft the article, add concrete
    examples and outputs in the right places, then reread from the reader's
    point of view and remove repetition, vague claims, and sentences that do
    not add new information.
25. Do not stop after one cleanup pass. Iterate the text several times and
    explicitly check it against the existing style rules in this file before
    considering the article done.
25. Title, ingress, `TL;DR`, and body must not merely restate the same claim in
    different words. Each layer should add something new: promise, summary,
    and concrete explanation.
26. Be aggressive about removing repeated information. Prefer one precise,
    meaningful sentence over several sentences that restate the same point
    with slightly different wording.
27. Prefer reader-facing terminology over Git-internal jargon in published
    prose. Replace terms such as "dirty submodule" with clearer user-facing
    wording like "muuttunut submoduuli" unless the Git term itself is the
    point being documented.
28. Treat developer-facing BusDK tools such as `bus-dev` and `bus-run` as real
    product surfaces, not as background-only implementation details. They are
    valid blog topics when the change affects actual commands, workflows,
    tooling contracts, or bug fixes in a concrete way.
29. When reviewing history for missing articles, assume most meaningful work
    should map to some concrete changed behavior, command, workflow, or bug fix.
    Do not discard a topic merely because it is developer-facing or happened
    "in the background"; inspect what actually changed first.
30. All Bus modules that are published as binaries are valid public product
    surfaces for the blog, even when their source code is not publicly
    distributed. It is acceptable to present their commands, workflows, and
    user-visible behavior in blog posts.
31. Do not imply that a module is non-public merely because its source is not
    open. The publication boundary for blog writing is the shipped binary
    interface and public documentation, not source-code availability.
32. Do not publish local scratch or verification paths such as `busdk.com/tmp`
    in reader-facing command examples. Verify examples however needed, then
    rewrite them to reader-facing project-root paths, documented file paths,
    or other generic locations that make sense outside the writer's machine.
33. Avoid abstract, teleological prose such as "the project started to appear
    as a deliberate whole" or similar AI-sounding summary language. State the
    concrete significance of the day instead: what was created, what became
    possible, and why that moment mattered even if the product was still early.
34. Do not leak research provenance or writing workflow into public prose.
    Avoid phrases like "tässä root-repossa", "git-historian kautta", or other
    source-oriented wording unless the source itself is the public-facing
    subject. Write from the reader's point of view and keep only the product-
    level fact that matters.
35. Do not narrate the reader from the outside with phrases like "lukijan
    kannalta", "lukija huomaa", or similar editorial meta-commentary unless
    the sentence is truly about a concrete user action or outcome. State the
    product meaning directly instead.
36. When writing for a broad public audience, do not center the sentence on
    developer jargon such as "root", "Makefile", repository structure, or
    similar implementation-facing terms. Prefer general language like shared
    development environment, common tooling, or shared structure, and mention
    exact technical names only as secondary clarifications when they help.
37. When an article title or ingress changes, update the blog index card in
    `busdk.com/docs/blog/index.html` in the same change set so the listing,
    article page, and reader expectation stay in sync.
38. Before making any historical claim in a blog post, verify the actual change
    from Git history first. Do not infer behavior from commit titles, repo
    creation dates, or high-level assumptions alone.
39. Historical-article verification process:
    1. identify the exact date and candidate commit(s) with `git log --since=... --until=...`
    2. inspect the real file-level change with `git log --stat`, `git show`, or both
    3. if the claim is about code behavior, open the changed source/help/docs files and verify what actually appeared in that commit
    4. only then write the article claim, using the verified behavior and date
    5. if the code did not yet exist, do not describe the feature as existing; write instead about the actual state that was introduced
40. Repository creation, initial commits, pin bumps, and documentation commits
    are not automatically product milestones. Treat them as blog-worthy only
    after verifying what concrete reader-visible capability, command, docs
    surface, or workflow changed.
41. Retrospective blog posts must read like same-day news. Titles, ingresses,
    TL;DR blocks, body text, metadata descriptions, and blog-index cards should
    default to present-tense or immediate-news phrasing instead of later
    summary language.
42. Before considering a blog article finished, reread it once specifically for
    hidden retrospective wording such as `sai`, `toi`, `alkoi`, `julkaistiin`,
    `siirtyi`, `muuttui`, `kuvattiin`, or similar forms, and rewrite them when
    the article is meant to sound like it was published on that day.
43. Before creating a new blog article for a date that already has an article,
    first check whether the new material belongs in the existing same-day
    article instead. Prefer expanding the existing article into a richer and
    more complete same-day piece when the topics are meaningfully related,
    and create a separate article only when the subject is clearly independent
    enough to deserve its own page.

## Commit Workflow (When Asked To Commit)

1. Commit only staged changes.
2. Do not auto-stage files unless explicitly asked.
3. Use small, meaningful, imperative commit messages.
4. Never push, tag, or run remote synchronization unless explicitly asked.
5. If hooks/checks reject a commit, report the failure and apply the minimal fix before retrying.

## Deletion Safety Rule

1. Never use internal delete tools.
2. If a path is tracked, use `git rm` (or `git rm --cached` to untrack but keep file).
3. If untracked, use `rm` (`-r`/`-f` only when necessary).
4. After deletion, update references/imports/links accordingly.
5. If a target path is already absent, treat it as a warning and continue.

## Context Memory Rule

1. When durable repo-specific workflow guidance is learned, record it in this `AGENTS.md` in the same change set.
2. Keep this file scoped to the `busdk.com` subtree only.
3. Revisit and refine this file as project context evolves.

## Gitignore Rule

1. .bus MUST be tracked; never add .bus or .bus/ to .gitignore.
2. In private repositories, .bus/ must be tracked; .bus/secrets may be tracked in private repositories only and must not be tracked otherwise.
3. Runtime lock artifacts such as .bus-dev.lock may be ignored.
