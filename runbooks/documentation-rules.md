# Documentation Rules

Detailed rules for `docs/` content outside the blog; expands the root
`AGENTS.md` core.

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
17. Product-site navigation must not link directly to `mailto:` addresses.
    Route navigation to product-local pages such as Contact, Pricing, or
    Deployment instead, and reserve `mailto:` links for explicit CTA buttons in
    page content.
