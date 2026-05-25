---
name: brag-me
description: "Create and refresh local brag repository demo decks from merged PRs, git commits, GCP metrics, closed Sentry issues, and manual notes. Use when preparing team demos before planning, showing impact metrics, or building Reveal.js presentations under me/brag-me. Trigger: brag repository, brag-me, demo before planning, demo da planning, pull requests mergeados, Sentry issues fechadas, GCP impact, cost/performance/observability/accessibility improvements, /brag-me."
argument-hint: "Mode: create | refresh | collect | deck-review"
---

# Brag-Me

Create or update a local brag repository for team demo presentations.

The output lives in the repository where the skill runs:

```text
me/brag-me/YYYY-MM-DD-highlight/
├── YYYY-MM-DD-brag.md
└── index.html
```

`YYYY-MM-DD-brag.md` is the editable source of truth. `index.html` is the
self-contained Reveal.js presentation that can be opened directly in a browser.

## Modes

| Mode | Use when | Output |
|---|---|---|
| `create` | Starting a new demo package | New dated folder, brag source, and deck |
| `refresh` | Updating an existing demo with new evidence | Updated generated sections and deck |
| `collect` | Gathering evidence before writing slides | Filled or updated brag source |
| `deck-review` | Reviewing an existing deck before demo | Findings and focused edits |

## Procedure

1. Determine the period covered by the demo. If it is unclear, ask for the
   sprint, date range, release, or "since last planning" boundary.
2. Determine the scope: repository, GitHub owner or organization, author
   identity, GCP project or service, Sentry organization or project, and target
   audience. Infer what is obvious from the workspace and ask only for missing
   decisions.
3. Inspect existing `me/brag-me/` folders. Update the matching demo when one
   exists; otherwise create `me/brag-me/YYYY-MM-DD-highlight/`, where
   `highlight` is lowercase kebab-case and names the biggest impact.
4. Create or update `YYYY-MM-DD-brag.md` from
   [brag-template.md](./assets/brag-template.md). Preserve manual sections and
   user-written notes unless the user explicitly asks to rewrite them.
5. Collect evidence from available sources, in this priority order:
   user-provided notes, existing manual additions, local git history, merged
   GitHub pull requests, improvement commits, closed or resolved Sentry issues,
   GCP cost/performance/reliability metrics, changelogs or release notes,
   benchmark outputs, logs, screenshots, then explicit assumptions.
6. Use authenticated local tools when available, such as `git`, `gh`, `gcloud`,
   Sentry CLI or API access, local exported reports, and user-pasted data. If a
   tool is missing or unauthenticated, record the gap in the brag source and
   continue with other evidence.
7. Generate `index.html` from [reveal-template.html](./assets/reveal-template.html)
   using only information present in `YYYY-MM-DD-brag.md`. Manual additions are
   allowed to become slides.
8. Report the created or updated folder, sources checked, evidence used, manual
   sections preserved, generated files, and remaining data gaps.

## Brag Source Rules

- Do not fabricate PRs, commits, issues, costs, performance gains, reliability
  gains, accessibility impact, or customer impact.
- Mark unknown values as `TBD` with a short note explaining how to fill them.
- Include provenance for every collected item: source type, link or command or
  file path when available, date/window, and confidence.
- Keep generated evidence separate from `Manual additions` so a later refresh
  can update generated content without overwriting user context.
- Preserve `Manual additions`, caveats, screenshots, user-written narrative,
  and manually curated links during refreshes.

## Deck Rules

- The deck must be a single `index.html` that opens directly in a browser.
- Do not require a dev server, package install, build step, CDN, or network
  access for the generated deck.
- Use Reveal.js semantics and keyboard navigation. When official Reveal.js
  assets are available locally, inline the official CSS/JS and preserve the MIT
  attribution in the generated HTML. Do not fetch CDN assets silently.
- Build slides from the brag source, not hidden state.
- Use readable font sizes, strong contrast, and no color-only meaning.
- Charts must include labels, units, before/after values, and measurement
  windows.
- Screenshots need captions or alt text with the source and date.

## Recommended Evidence Commands

Use these only when the relevant tools are installed and authenticated:

```bash
git log --since=<date> --author=<author> --oneline --decorate
gh pr list --state merged --author <user> --search "merged:>=YYYY-MM-DD"
gcloud logging read <filter> --project <project> --limit 50
sentry-cli issues list --org <org> --project <project>
```

Record commands that were run, summaries of useful output, and any access
errors in `YYYY-MM-DD-brag.md`.

## Output

End with:

- Folder created or updated
- Files created or updated
- Evidence sources checked
- Manual sections preserved
- Missing access or `TBD` data
- How to open the deck
