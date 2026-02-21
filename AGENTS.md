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

## Documentation Rules

1. Keep end-user docs concise, task-oriented, and readable.
2. Prefer short paragraphs and avoid repetitive wording.
3. Use lists/tables only when they materially improve clarity.
4. When mentioning the `bus` GitHub repository, inline-link to `https://github.com/busdk/bus`.

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
