# AGENTS.md

Scoped guidance for the `busdk.com/docs/blog/` subtree on the GitHub Pages site.

## Scope And Precedence

1. Apply this file to `busdk.com/docs/blog/` and its descendants.
2. Follow the repository-root `AGENTS.md` files first, then this file for blog-specific rules.
3. Keep behavior static-site friendly, deterministic, and easy to review in Git.

## Blog Purpose

1. The blog exists under the public BusDK marketing/documentation site.
2. Posts must work for multiple audiences at once without splitting the published copy into audience-specific lanes.
3. Explain impact, implementation, and compatibility in one shared article flow that reads naturally to any reader.
4. Every post must reinforce BusDK's core product promises:
   1. Git-native auditable history.
   2. Deterministic workflows.
   3. CLI-first control plane.
   4. Modular ecosystem.
   5. AI-ready but not AI-dependent operation.

## Required Output Per Post

1. Every post must include a clear, descriptive H1 and a 1-2 sentence ingress that answers:
   1. What changed or what is being explained.
   2. Why it matters now.
   3. What became possible, safer, clearer, or simpler for the reader.
2. Every post must include a `TL;DR` block with 3-6 bullets.
3. `TL;DR` bullets should be labeled by content such as significance, implementation, compatibility, availability, or limits, not by audience segment.
4. Every post about a feature, fix, or release must tell the reader concretely what changed in the product, not just that "work happened" or that a tracker item moved.
5. For bug-fix posts, identify the bug by visible behavior or recognizable symptom so a reader can tell "this is the thing that was fixed".
6. For feature posts, show what the new capability does in practice with a short concrete example, command, output, screenshot description, or before/after comparison.
7. Do not make tracker workflow, re-verification loops, reopen/close churn, or commit hygiene the main story of a public post unless the user explicitly wants a process post.
8. Every post must include a concrete "how to use this" section with at least one command or API example and expected output, unless the post explicitly marks the element as not applicable.
9. Every post must include compatibility guidance:
   1. State explicitly whether existing workspaces, scripts, modules, data formats, or integrations are affected.
   2. If affected, include a short migration and verification path.
   3. If not affected, say so explicitly.
10. Every feature or release-oriented post must include a human-readable change log section using these headings when relevant:
   1. `Added`
   2. `Changed`
   3. `Deprecated`
   4. `Removed`
   5. `Fixed`
   6. `Security`
11. Every post may include a small related-links block when it adds value, but avoid audience-segmented CTA lanes.
12. Every post must include related links to canonical documentation or release material when such links materially help the reader.

## Mandatory Structure

1. Use this default structure unless the post type clearly needs a tighter variant:
   1. Metadata and head tags.
   2. Hero with H1 and ingress.
   3. `TL;DR`.
   4. Prerequisites when needed.
   5. What changed and why it matters.
   6. Concrete example, before/after, or usage path.
   7. Screenshots, diagrams, or "not applicable".
   8. Change log and compatibility.
   9. Known limits or FAQ.
   10. Related links when genuinely useful.
2. Keep the page skimmable:
   1. Put the most important information first.
   2. Use descriptive section headings.
   3. Keep paragraphs short.

## Supported Post Types

1. `release-notes`
   1. Focus on what changed, who is affected, how to adopt it, and how to verify it.
   2. Always include a structured change log.
2. `deep-dive`
   1. Focus on problem, design choice, architecture, and operational impact.
   2. Include concrete examples and measured or observed practical consequences.
3. `tutorial`
   1. Focus on step-by-step execution.
   2. Include prerequisites, expected outputs, and common failure corrections.
4. `case-study`
   1. Focus on before/after, measurable outcomes, and copyable workflow pattern.
5. `announcement`
   1. Focus on why the feature exists, availability, limits, and immediate trial path.

## BusDK-Specific Writing Rules

1. Prefer claims that can be verified from:
   1. Git history.
   2. Release tags.
   3. Public docs.
   4. Public CLI behavior.
2. When a post discusses project history, use exact dates and public-facing concrete changes. Commit hashes are source material, not default published content.
3. Do not invent historical motives or roadmap promises not supported by public evidence.
4. When referencing the Bus dispatcher repository, inline-link to `https://github.com/busdk/bus`.
5. When linking to product docs, always use canonical `https://docs.busdk.com/...` URLs.
6. Do not link to private SDD pages from the public blog.
7. If a capability is pre-release, beta, preview, or still stabilizing, say so plainly.

## Metadata And SEO Requirements

1. Every blog page must include unique metadata in the HTML head:
   1. `<title>`
   2. `<meta name="description">`
   3. canonical URL
   4. Open Graph tags
   5. Twitter/X card tags
   6. JSON-LD structured data
2. Use `BlogPosting` or `Article` JSON-LD for article pages.
3. Blog index pages may use `Blog` JSON-LD.
4. Each page must have accurate publish and modified dates.
5. Use descriptive link text. Do not use vague CTA text such as "click here".
6. If the same post is published in Finnish and English on separate URLs, add hreflang alternates for both.
7. If no alternate language exists yet, omit hreflang rather than inventing placeholder links.

## Accessibility And Performance

1. Use a correct heading hierarchy.
2. All informative images must have meaningful alt text.
3. Decorative images should use empty alt text or be hidden from assistive technology.
4. Do not place essential text only inside images.
5. Keep hero media lightweight and justified by comprehension value.
6. Use responsive images and `loading="lazy"` for below-the-fold images.
7. Preserve strong color contrast, keyboard focus visibility, and mobile readability.
8. Keep pages light enough to support good Core Web Vitals targets.

## Local Validation Notes

1. In this environment, `xmllint --html --noout` is not a valid HTML5 page validator for the blog. It reports modern semantic tags such as `header`, `nav`, `main`, `section`, `article`, `aside`, `time`, and `footer` as parser errors even on otherwise valid HTML5 documents.
2. Do not use `xmllint --html --noout` as the pass/fail signal for blog page validity here.
3. Prefer practical verification for this subtree:
   1. inspect rendered output in a browser or headless browser
   2. confirm expected links, titles, and key text are present
   3. report clearly when only structural/manual verification was possible
4. When invoking the Chrome binary on macOS from shell, quote or escape the application path because `Google Chrome.app` contains a space. Example: `\"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome\" --headless ...`

## Style And Language

1. Write like a technically competent colleague:
   1. clear
   2. calm
   3. concrete
   4. non-hype
2. Avoid slang, idioms, and culturally narrow jokes so the text stays translation-friendly.
3. Prefer inclusive, gender-neutral language.
4. Keep paragraphs to one main idea each.
5. Use descriptive headings instead of generic labels whenever possible.
6. The first paragraph must always explain what changed and why it matters.
7. Explain new terms on first use.
8. When writing in Finnish, prefer clear general language over heavy nominal or passive phrasing.
9. When producing paired Finnish and English posts, keep terminology consistent across languages.
10. Do not let prompt language or process language leak into published copy. Avoid phrases that sound like internal assignment framing, such as "what git history tells", "retrospective news", "phase", "series overview", or other wording that centers the writing process instead of the reader's question.
11. Prefer reader-centered framing in titles and intros. The post should sound like a real product or project update that answers "what changed and why should I care?" rather than a commentary on how the article was assembled.
12. Blog index and article intros must be terse. Do not spend sentences explaining that the page contains posts, dates, updates, or concrete changes if the surrounding UI already shows that. Say the useful thing once, briefly.
13. Avoid repeating the same idea in adjacent headline and intro copy. If the H1 already says what changed, the intro should add the most important consequence, not restate the same sentence in longer form.
14. Cut meta-explanations about the blog itself unless they materially help navigation. Default to one short sentence of substance over two explanatory sentences.
15. On blog index pages, organize and label entries by the change itself, not by the publishing format. Dates are supporting metadata, not the main message.
16. Do not overemphasize that posts are "day-specific" or "published by day" in reader-facing copy. Use dates as quiet context only.
17. If multiple independent changes happened on the same day, prefer separate reader-facing entries or clearly change-first labels rather than making the date the organizing idea.
18. Blog index lists must be ordered newest first unless the user explicitly asks for another ordering.
19. Dates should be visually easy to scan, especially on index cards and article headers. Treat the date as clear supporting metadata, not body text lost inside paragraphs.
20. Do not include raw Git commit hashes in reader-facing blog copy by default. Name the concrete change, command, module, bug, or workflow instead. Use Git only as an internal source unless the user explicitly asks for hashes.
21. Do not publish Git verification snippets in public posts by default. Use repository history only as an editorial source unless the user explicitly asks for Git-facing material and readers actually have access to it.
22. On the index page and article headers, make the date visually distinct from body copy through layout or styling. The date should read as clear metadata, not as a sentence inside the paragraph flow.
23. If the article header already shows the date, do not repeat the date at the start of the ingress or body unless the date itself is the point of the sentence. Lead with the change, not with the calendar.
24. Do not publish audience-lane copy such as separate SMB, developer, evaluator, or "choose your path" sections in the public blog. The same article should work for all readers without segment labels.
25. Do not write "the tracker handled", "the bug was reopened and closed", "the same day was spent on", or similar process narration unless that workflow detail is itself the user-visible news. Public posts should lead with the product change, not internal handling.
26. When covering a bug fix, describe the bug in terms of what broke for the user: wrong output, missing row, incorrect report section, failing command, broken import, or similar recognizable symptom.
27. When covering a bug fix, prefer the final stable outcome over the repair journey. The reader usually needs to know what was wrong and what is correct now, not how many intermediate fixes were attempted.
28. When covering a feature or fix, include one concrete sign of the changed behavior: a command, output snippet, UI location, before/after sentence, sample file effect, or other recognizable example.
29. Avoid empty intensity words such as "concrete", "visible", "significant", or "important" unless the sentence also says concretely what changed.
30. If a sentence can be rewritten from process language into outcome language, prefer the outcome version. Example: write "profit and loss now includes prior-year correction rows correctly" instead of "prior-year-correction profit-and-loss bug was closed".
31. A post is not done if a reader who knows the product cannot identify whether the article affects their workflow. Name the affected command, module, report, screen, data shape, or symptom explicitly.

## Visual And Layout Rules

1. Preserve the visual language of the existing `busdk.com` site:
   1. static HTML
   2. lightweight CSS
   3. responsive layout
   4. no unnecessary framework additions
2. Reuse the existing BusDK design cues where practical:
   1. teal accent palette
   2. strong typography
   3. subtle gradients
   4. practical code-block presentation
3. The blog must remain readable on mobile first, then desktop.
4. Do not add heavy client-side dependencies for blog behavior.

## Review Checklist

1. Verify all commands and outputs before publishing.
2. Verify all dates, versions, and commit references against Git history.
3. Confirm compatibility statements are accurate.
4. Confirm no secrets, tokens, customer data, or private URLs appear in text or images.
5. Confirm internal links point to the exact canonical docs page the reader needs.
6. Confirm metadata matches visible content.
7. Confirm JSON-LD does not claim content that is not visible on the page.
8. Confirm screenshots are sanitized and alt text is present.
9. Confirm related links or next steps, if present, are relevant and not segmented by audience.

## Publishing Workflow

1. Choose the post type first.
2. Draft the post using the required structure.
3. Run technical review for commands, APIs, compatibility, and links.
4. Run product review for value proposition, terminology, and CTA clarity.
5. Run accessibility and media review for headings, images, alt text, and readability.
6. Run SEO review for title, description, canonical URL, structured data, and sharing metadata.
7. If bilingual, publish on separate URLs and connect them with hreflang.
8. Verify the staged HTML/CSS locally before reporting completion.

## Metrics To Support

1. Posts should be easy to measure for:
   1. page views
   2. scroll depth
   3. documentation click-through
   4. install or evaluation clicks
   5. documentation click-through from posts
2. When possible, structure links so downstream analytics can distinguish article-level navigation without segmenting readers into separate public paths.

## History-Driven Posts

1. If a post is based on project history, explicitly state the observation window with exact dates.
2. Separate observed facts from interpretation.
3. Prefer timeline sections that tie dates to visible public site changes, shipped features, released modules, fixed bugs, or other reader-meaningful milestones. Use commits and tracker movement only as editorial evidence.
4. Explain why a history item mattered in product terms first. Internal process detail is secondary and usually omitted.
5. Do not include repository-forensics or Git-verification command sequences in public history posts unless the user explicitly asks for them and the relevant history is actually available to readers.
6. When the user requests a multi-post history series, publish the articles in chronological order from oldest phase to newest phase.
7. The blog index for a history series must list the phase articles in chronological order and may keep any later overview article in a separate "series overview" section.
8. Retrospective history posts for BusDK must be day-specific by default: one concrete publication date per post, not broad date ranges such as `24.1.2026 - 7.2.2026`, unless the user explicitly asks for a range summary.
9. Each retrospective day-post must state what changed on that exact date in concrete product terms: affected module, command, report, screen, import path, release artifact, or recognizable bug symptom.
10. If the source material shows multiple edits to the same bug on one day, compress them into the final reader-relevant outcome unless the intermediate states materially changed what users saw.
11. Avoid generic strategy prose in retrospective posts. Prefer concrete statements like "on 2026-04-04 changed-scope detection started noticing dirty submodules in root `make test` and `make e2e`" over vague phase summaries.
