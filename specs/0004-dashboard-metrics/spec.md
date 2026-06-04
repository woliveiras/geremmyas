---
spec: "0004"
title: Dashboard git dates and metrics
family: platform
phase: 2
status: Approved
owner: ""
depends_on: [2]
origin: "docs/prds/2026-06-04-dashboard.md"
---

# Spec: Dashboard git dates and metrics

## Context & Motivation

Teams need to answer "are we shipping faster?" and "which family is blocked?"
These questions require timestamps (when was a spec created, approved,
implemented) and computed metrics (velocity, lead time).

Git history already contains this information — commit dates for when spec files
were created and when their status changed. This spec extracts those dates,
caches them for performance, and generates a metrics page with charts.

## Requirements

### Functional

- [ ] Extract spec creation date from git history (first commit adding the
  `spec.md` file)
- [ ] Extract status transition dates from git diff history (commits where
  `status:` frontmatter value changed)
- [ ] Use batch `git log` commands (not per-file) for performance
- [ ] Cache results in `.geremmyas-cache/gitdates.json` with `last_commit` SHA
- [ ] Incremental cache update: only scan commits since `last_commit`
- [ ] `--no-git` flag skips git extraction entirely (faster, metrics page
  shows "Git data not available")
- [ ] `--no-cache` flag forces full git rescan
- [ ] Compute metrics from dates:
  - Velocity: count of specs reaching "Implemented" per week and per month
  - Average lead time: median created → implemented
  - Average review time: median created → approved
  - Average implementation time: median approved → implemented
- [ ] Exclude deprecated specs from metrics calculations
- [ ] Generate `metrics.html` page with:
  - Summary cards: avg lead time, avg review time, avg impl time, specs in
    progress count
  - Velocity bar chart (specs implemented per month)
  - Lead time trend line chart (median per month)
  - Phase breakdown stacked bar chart (specs by status per phase per family)
  - Family progress horizontal bars
- [ ] Add metrics link to the shared navigation bar
- [ ] Add timeline section to spec detail page (from 0002): created → approved
  → implemented dates

### Non-Functional

- [ ] Cold git scan (100 specs, no cache): under 5 seconds
- [ ] Warm scan (cached, incremental): under 500ms
- [ ] Charts rendered via Chart.js (vendored, embedded in binary)
- [ ] Metrics page works offline
- [ ] `.geremmyas-cache/` added to `.gitignore` recommendation in docs

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Date parsing, metric computation, cache serialization | parse git log output, compute median lead time, read/write cache JSON |
| **integration** | Git extraction in a real git repo | create temp repo with commits, extract dates, verify values |

Default: **unit** for computation, **integration** for git interaction.

## Acceptance Criteria

- [ ] Given a spec.md first committed on 2026-01-10, when git dates extracted,
  then created_at = 2026-01-10 (±1 day for timezone)
- [ ] Given a spec.md whose status changed from "Draft" to "Approved" on
  2026-01-20, when git dates extracted, then approved_at = 2026-01-20
- [ ] Given `--no-git` flag, when dashboard runs, then metrics page shows "Git
  data not available" and no git commands are executed
- [ ] Given a cached gitdates.json with last_commit=abc123, when dashboard runs
  and 2 new commits exist, then only those 2 commits are scanned
- [ ] Given `--no-cache` flag, when dashboard runs, then full git scan runs
  regardless of existing cache
- [ ] Given 10 specs implemented over 3 months, when metrics page renders, then
  velocity chart shows bars for each month with correct counts
- [ ] Given specs with lead times of 5, 10, 15, 20 days, when metrics computed,
  then average lead time shows median (12.5 days)
- [ ] Given 2 deprecated specs with dates, when metrics computed, then they are
  excluded from all calculations
- [ ] Given a spec detail page, when git dates are available, then timeline
  shows created → approved → implemented dates with duration between each
- [ ] Given cold scan of 100 specs, when timed, then completes under 5 seconds

## Edge Cases

- Spec created via `git mv` (rename) — `--follow` may not work in batch mode;
  fallback to earliest known commit for that path
- Status changed multiple times (Draft → Approved → Draft → Approved) — use
  the latest transition to each status
- Spec has no status changes in git (created directly as "Implemented") —
  created_at = implemented_at, no review/impl time
- Git not available (not a git repo, or git not installed) — same behavior as
  `--no-git`
- Cache file corrupted or invalid JSON — delete and rescan
- Commit author dates vs committer dates — use author date (`%aI`) as it
  reflects when work was done

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Git command | Shell out to `git log` | Avoids go-git dependency (~10MB); git is always available in dev environments |
| Batch strategy | One `git log -p` call for all spec.md files | O(1) git commands instead of O(n); parse output in Go |
| Chart library | Chart.js (vendored, minified) | ~200KB, well-known, good docs, no build step |
| Chart rendering | `<canvas>` with inline JSON data | Chart.js reads data from a `<script>` block in the page |
| Cache format | JSON with last_commit SHA + per-spec dates | Human-readable, debuggable, simple to parse |
| Median vs mean | Median for lead/review/impl times | Resistant to outliers (one spec taking 6 months skews mean) |

## Out of Scope

- Forecasting or projections
- Per-developer metrics
- Burndown charts (would need a "planned" count)
- Comparing metrics across repos
- Including deprecated specs in metrics (decided: excluded)
