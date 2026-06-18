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
8. Blog topic discovery should add only BusDK product surfaces that are actually
   implemented in Bus and can be verified through local module docs, help text,
   tests, or safe command execution. Do not add external regulatory, market, or
   background-only article ideas unless the user explicitly asks for that kind
   of topic.

## Documentation Rules

1. Keep `busdk.com/docs/` content focused on product/commercial landing-page communication for buyers and evaluators.
2. Prefer short paragraphs and avoid repetitive wording.
3. Use lists/tables only when they materially improve clarity.
4. When mentioning the `bus` GitHub repository, inline-link to `https://github.com/busdk/bus`.
5. Product pages on `busdk.com` may include commercial/landing messaging and product-audience framing when it improves clarity for buyers/evaluators.
6. Current product positioning: present BusDK as a self-hostable platform for developing, hosting, billing, and operating AI products, with accounting and compliance as an important supported product package rather than the whole product identity.
7. When discussing deployment, distinguish managed Finnish cloud operation, dedicated/customer-controlled environments, and customer self-hosting. Contractual data-processing terms are a commercial offer and must not be described as a code feature.
8. When linking from `busdk.com/docs/` to documentation pages, always use canonical `https://docs.busdk.com/...` URLs.
9. Do not link directly to private SDD pages from `busdk.com/docs/`; you may mention that private implementation design specifications exist.
10. Within published `busdk.com/docs/` pages, keep same-site navigation links relative so they work in local file or local server development as well as on the public site. Do not use absolute `https://busdk.com/...` URLs for header/home links inside subtree pages when a correct relative link exists.
11. When the current working directory is already `busdk.com`, run `git` directly instead of using `git -C busdk.com`; otherwise Git tries to enter a nonexistent nested `busdk.com/busdk.com` path.
12. When using `rg` from this repository with a search pattern containing
    backticks, pass each pattern through a single-quoted `-e` argument; do not
    put backtick-containing alternatives inside a double-quoted shell command.
13. Before passing optional repository paths such as `package.json`, `scripts`,
    or other tool-specific files to `rg`, `sed`, or `cat`, verify they exist
    in `busdk.com`; this static site repo may not contain common package or
    script directories.
14. Do not hard-code exact source-package EUR totals or per-module pricing
    tables in `busdk.com/docs/`. Link to the generated docs pricing page for
    estimates, and keep website copy clear that final commercial prices are
    contract quotes.
15. Published product sites under `busdk.com/docs/<product>/` must stand alone:
    do not add cross-product product-family navigation inside their top nav,
    side nav, or shared rendered nav. Shared static, GX, WASM, or other
    components may be reused, but the content they render must be specific to
    the current product site.
16. Product examples should show the smallest normal command first. Do not add
    optional default flags such as `--file services.yml`, `--profile-dir
    profiles`, default env files, or default state paths unless the example is
    specifically teaching non-default paths, Docker images, dedicated state, or
    deployment packaging.

## Publication Boundary

1. `busdk.com/docs/` is a published static site tree. Do not place `AGENTS.md`
   files anywhere under that tree unless the operator explicitly requests
   published-safe local guidance for a docs subtree.
2. Keep durable agent instructions for the `busdk.com` repo in non-published
   paths such as `busdk.com/AGENTS.md`.
3. The GX/UI docs subtree intentionally contains published-safe `AGENTS.md`
   files under `busdk.com/docs/gx-ui/`. Those files must stay free of secrets,
   private paths, worker ids, internal task ids, and process-only notes.
4. If another subtree under `busdk.com/docs/` needs extra guidance, add a
   clearly named section for that subtree in this root file unless the operator
   explicitly asks for published-safe local `AGENTS.md` files.

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
15. In blog code examples, distinguish the command clearly from its textual
    output. Prefer showing the command with a shell prompt marker such as `$`
    and then show the resulting output in a separate output block when that is
    clearer than mixing them together.
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
27. Do not rely on the sentence pattern "ei vain ..., vaan ..." or similar
    empty contrast formulas as a default writing move. If there is a real
    product meaning behind the contrast, spell that meaning out directly:
    tell the reader what becomes easier, more portable, more standard, more
    deterministic, or otherwise more useful in practice.
28. Prefer reader-facing terminology over Git-internal jargon in published
    prose. Replace terms such as "dirty submodule" with clearer user-facing
    wording like "muuttunut submoduuli" unless the Git term itself is the
    point being documented.
29. Treat developer-facing BusDK tools such as `bus-dev` and `bus-run` as real
    product surfaces, not as background-only implementation details. They are
    valid blog topics when the change affects actual commands, workflows,
    tooling contracts, or bug fixes in a concrete way.
30. When reviewing history for missing articles, assume most meaningful work
    should map to some concrete changed behavior, command, workflow, or bug fix.
    Do not discard a topic merely because it is developer-facing or happened
    "in the background"; inspect what actually changed first.
31. All Bus modules that are published as binaries are valid public product
    surfaces for the blog, even when their source code is not publicly
    distributed. It is acceptable to present their commands, workflows, and
    user-visible behavior in blog posts.
32. Do not imply that a module is non-public merely because its source is not
    open. The publication boundary for blog writing is the shipped binary
    interface and public documentation, not source-code availability.
33. Do not publish local scratch or verification paths such as `busdk.com/tmp`
    in reader-facing command examples. Verify examples however needed, then
    rewrite them to reader-facing project-root paths, documented file paths,
    or other generic locations that make sense outside the writer's machine.
34. Avoid abstract, teleological prose such as "the project started to appear
    as a deliberate whole" or similar AI-sounding summary language. State the
    concrete significance of the day instead: what was created, what became
    possible, and why that moment mattered even if the product was still early.
35. Do not leak research provenance or writing workflow into public prose.
    Avoid phrases like "tässä root-repossa", "git-historian kautta", or other
    source-oriented wording unless the source itself is the public-facing
    subject. Write from the reader's point of view and keep only the product-
    level fact that matters.
36. Do not narrate the reader from the outside with phrases like "lukijan
    kannalta", "lukija huomaa", or similar editorial meta-commentary unless
    the sentence is truly about a concrete user action or outcome. State the
    product meaning directly instead.
37. When writing for a broad public audience, do not center the sentence on
    developer jargon such as "root", "Makefile", repository structure, or
    similar implementation-facing terms. Prefer general language like shared
    development environment, common tooling, or shared structure, and mention
    exact technical names only as secondary clarifications when they help.
38. When an article title or ingress changes, update the blog index card in
    `busdk.com/docs/blog/index.html` in the same change set so the listing,
    article page, and reader expectation stay in sync.
39. Before making any historical claim in a blog post, verify the actual change
    from Git history first. Do not infer behavior from commit titles, repo
    creation dates, or high-level assumptions alone.
40. Historical-article verification process:
    1. identify the exact date and candidate commit(s) with `git log --since=... --until=...`
    2. inspect the real file-level change with `git log --stat`, `git show`, or both
    3. if the claim is about code behavior, open the changed source/help/docs files and verify what actually appeared in that commit
    4. only then write the article claim, using the verified behavior and date
    5. if the code did not yet exist, do not describe the feature as existing; write instead about the actual state that was introduced
41. Repository creation, initial commits, pin bumps, and documentation commits
    are not automatically product milestones. Treat them as blog-worthy only
    after verifying what concrete reader-visible capability, command, docs
    surface, or workflow changed.
42. Retrospective blog posts must read like same-day news. Titles, ingresses,
    TL;DR blocks, body text, metadata descriptions, and blog-index cards should
    default to present-tense or immediate-news phrasing instead of later
    summary language.
43. Before considering a blog article finished, reread it once specifically for
    hidden retrospective wording such as `sai`, `toi`, `alkoi`, `julkaistiin`,
    `siirtyi`, `muuttui`, `kuvattiin`, or similar forms, and rewrite them when
    the article is meant to sound like it was published on that day.
44. Before creating a new blog article for a date that already has an article,
    first check whether the new material belongs in the existing same-day
    article instead. Prefer expanding the existing article into a richer and
    more complete same-day piece when the topics are meaningfully related,
    and create a separate article only when the subject is clearly independent
    enough to deserve its own page.
45. When a post introduces a new tool or module, use the module README,
    public docs, help text, visible source surface, tests, and any safely
    runnable command surface as research material so the post explains what the
    tool actually does, how it is used, and what kind of output or behavior the
    reader should expect.
46. When a command can safely create visible workspace state under
    `busdk.com/tmp/`, prefer running it there and quoting the real resulting
    output or files instead of relying only on README wording.
47. Tool-introduction posts must say plainly when the tool is experimental,
    unfinished, or not yet meant for production use. Do not present an early
    or exploratory tool as mature just because it ships in a package.
48. Honest product communication matters more than launch hype. Prefer
    accurate descriptions such as experimental, early, narrow, or still
    evolving whenever the docs, README, tests, or observed behavior support
    that characterization.
49. When a new tool surface is intentionally narrow, say that plainly and then
    explain the narrow scope with one concrete example. Do not describe a thin
    first release as if it were already a broad platform.
50. Avoid overstating development status with words like "active" unless the
    pace itself is the verified point. Prefer neutral wording such as "under
    development" when the tool is unfinished but current momentum is not the
    reader-relevant fact.
51. Do not use empty contrast phrases such as "this is not X but Y" unless the
    contrast adds concrete meaning for the reader. If a sentence can be removed
    without losing information, remove it.
52. When a tool is unfinished, describe that in product terms rather than
    backlog or process terms. Do not write that some capability is "later
    work" or similar. Instead explain what the tool already does usefully
    today, what broader purpose the module exists for, and which part of that
    broader purpose is not yet ready for normal use.
53. When telling the reader that a capability is still missing or incomplete,
    prefer suitability language over internal planning language. Help the
    reader decide whether the current tool is already useful for assertions,
    inspection, packaging, or another narrow task, instead of describing the
    missing part as an implementation to-do.
54. TL;DR lists are acceptable when a list is the clearest way to present the
    condensed facts. Do not flatten a good list into prose only to avoid an
    "AI-written" feel; judge the structure by clarity, not by whether it looks
55. When iterating Finnish blog prose, prioritize corrections in this order:
    1. meaning, factual accuracy, dates, amounts, obligations, and anything
       that could change the reader's understanding
    2. clear general-language norm issues such as punctuation, compounds,
       capitalization, hyphenation, and obvious inflection errors
    3. clarity, rhythm, structure, paragraphing, and information order
    4. optional polish such as wording alternatives, SEO refinements, or other
       stylistic tuning
56. Meaning and factual precision outrank style. If a stylistic rewrite risks
    changing emphasis, certainty, legal meaning, accounting meaning, or scope,
    keep the original claim narrower and mark the uncertainty in the writing
    process instead of forcing a smoother sentence.
57. Preserve the writer's voice when correcting Finnish. Do not flatten the
    prose into generic polished AI copy. Intentional warmth, directness, or
    light colloquial tone may stay when they do not reduce clarity or trust.
58. In legal, tax, accounting, money, deadline, and responsibility language,
    apply the strictest clarity standard. If wording could change who must do
    what, by when, for how much, or under which condition, fix that before any
    stylistic editing.
59. Do not force uncertain language corrections into published copy. If a term,
    inflection, brand spelling, or interpretation cannot be verified from the
    relevant source, keep the safe wording or verify first; do not guess.
60. When editing for clarity, prefer self-contained paragraphs and sections.
    Readers may land directly on a section from search, so the surrounding
    wording should not depend on hidden context or previous paragraphs.
61. Prefer descriptive anchor text and keep the page title, visible H1, and
    article promise semantically aligned. Do not use vague link text when a
    more informative phrase helps both the reader and discoverability.
62. When revising sentence flow in Finnish, watch especially for accidental
    repetition, compound-word errors, misleading word order, uneven lists,
    and wrong dash or hyphen usage. Fix these as language-quality issues, not
    as cosmetic polishing.
63. Internal editing aids such as tracked-change notation, priority labels, or
    uncertainty markers may be useful during drafting, but they must never leak
    into the published article text.
    like a list.
55. During the final editorial pass, scan explicitly for repeated sentence
    starters and cadence patterns across the feed. In particular, avoid
    overusing formulas such as "Tämä tekee ...", "Tämä auttaa ...",
    "Tämä on tärkeä ...", "Samalla ...", "Jos haluat ...",
    "Ensimmäinen ...", "Yksi ...", and "Käytännössä ..." when they function
    only as reusable rhythm instead of carrying new meaning.
56. These patterns are not forbidden absolutely, but they must not become the
    default rhythm of the blog. When they repeat across nearby articles,
    rewrite the sentence into a more direct statement of effect, capability,
    meaning, or next step.
57. Prefer varying how docs links are introduced. Do not end many articles with
    the same "Jos haluat ..." formula when a simpler direct sentence such as
    "Nykyinen komentopinta löytyy ..." or an inline semantic link is clearer.
58. Use an explicit blog-iteration workflow for every substantial blog pass:
    1. update `busdk.com/PLAN.md` first with the current pass and any newly found article ideas
    2. inspect the relevant day or tool from Git history before making claims
    3. read the related README, public docs, help text, visible source surface, tests, and safe runnable examples so the article describes the tool honestly
    4. before creating a new same-day article, check whether the content belongs in an existing same-day article and expand that first when it fits naturally
    5. update the article body and then the blog index card in the same pass
    6. do a separate editorial reread for repetition, meta language, vague claims, and stale same-day-news tense
    7. if the review uncovers more promising topics or more writing fixes than can be completed immediately, add them back to `busdk.com/PLAN.md` as open items instead of leaving them implicit
59. When re-iterating existing articles, do not rely only on prior blog text.
    Re-open the underlying tool surface when needed so the next revision can
    say something more concrete, more honest, or more useful than the previous
    draft.
60. When an unfinished tool is already useful in one narrow slice, describe
    that slice positively first, then state the missing broader scope in calm
    product language. Avoid framing that leads with absence, backlog thinking,
    or a negative disclaimer when the present tool already has a real use.

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
