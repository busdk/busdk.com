# AGENTS.md

## Highest-Priority Rule: Precise Language And Precise Reasoning

Precise language is part of precise reasoning. Before acting or reporting, name
the exact object, action, scope, evidence, and uncertainty. State what changed
and what did not change. Never use a broader claim than the evidence supports.
If an exact, unambiguous sentence cannot be written, inspect the evidence or
ask for clarification before proceeding.


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
6. When the user gives new instructions, create or update `PLAN.md`
   first so current work and still-open earlier work stay visible as unchecked
   items until they are actually finished.
7. Whenever research into project history, module behavior, docs, or source
   uncovers a clearly blog-worthy future topic, add it immediately to
   `PLAN.md` as an unchecked item with the most natural article date
   for that topic. Do not leave promising article ideas only in conversation.
8. Blog topic discovery should add only BusDK product surfaces that are actually
   implemented in Bus and can be verified through local module docs, help text,
   tests, or safe command execution. Do not add external regulatory, market, or
   background-only article ideas unless the user explicitly asks for that kind
   of topic.

## Documentation Rules

Keep `docs/` content focused on product/commercial landing-page communication
for buyers and evaluators. Before editing anything under `docs/` outside the
blog, read `runbooks/documentation-rules.md`.


## Publication Boundary

1. `busdk.com/docs/` is a published static site tree. Do not place `AGENTS.md`
   files anywhere under that tree unless the operator explicitly requests
   published-safe local guidance for a docs subtree.
2. Keep durable agent instructions for the `busdk.com` repo in non-published
   paths such as this root `AGENTS.md` and `runbooks/`.
3. The GX/UI docs subtree intentionally contains published-safe `AGENTS.md`
   files under `busdk.com/docs/gx-ui/`. Those files must stay free of secrets,
   private paths, worker ids, internal task ids, and process-only notes.
4. If another subtree under `busdk.com/docs/` needs extra guidance, add a
   clearly named section for that subtree in this root file unless the operator
   explicitly asks for published-safe local `AGENTS.md` files.

## Blog Rules (`docs/blog/`)

The blog is one reverse-chronological feed of standalone articles. Before
writing, revising, or indexing any blog article, read `runbooks/blog-rules.md`;
it carries the full article, style, language, verification, and iteration
rules.


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
