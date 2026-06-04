---
spec: "0001"
title: Dashboard spec parser
family: platform
phase: 0
status: Draft
owner: ""
depends_on: []
origin: "docs/prds/2026-06-04-dashboard.md"
---

# Spec: Dashboard spec parser

## Context & Motivation

The `geremmyas dashboard` command needs a foundation layer that scans the
repository for spec artifacts and produces a structured in-memory data model.
Every downstream view (overview, board, metrics, serve) depends on this parser.

Today agents and humans read raw markdown. The parser extracts structured data
from YAML frontmatter and checkbox patterns so that HTML renderers and metric
calculators can work with typed Go structs instead of raw text.

## Requirements

### Functional

- [ ] Scan `specs/` for directories matching `NNNN-<slug>/`
- [ ] Parse YAML frontmatter from `spec.md` (spec, title, family, phase,
  status, owner, depends_on, origin)
- [ ] Parse markdown body from `spec.md` for HTML rendering downstream
- [ ] Detect existence and read body of `plan.md` per spec folder
- [ ] Parse `tasks.md` checkbox counts: `- [ ]` pending, `- [~]` in progress,
  `- [x]` done
- [ ] Scan `docs/prds/*.md` for frontmatter (title, date) or extract date from
  filename pattern `YYYY-MM-DD-<slug>.md`
- [ ] Scan `docs/bugfixes/*.md` with same date extraction logic
- [ ] Group specs by family and phase from frontmatter fields
- [ ] Specs without frontmatter placed in "Ungrouped" family with a warning
- [ ] Sort specs by number within each phase

### Non-Functional

- [ ] Performance: parsing 100 spec folders completes in under 1 second
- [ ] No new external dependencies — use stdlib + existing frontmatter parser
- [ ] Errors in individual spec files logged as warnings, do not abort the full
  scan

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Parse logic, data extraction | frontmatter parsing, checkbox counting, filename date extraction |
| **integration** | Full scan of a temp directory tree | scan specs/ with multiple families, plan/tasks presence |

Default: **unit** for parsers, **integration** for the directory scanner.

## Acceptance Criteria

- [ ] Given a `specs/0001-foo/spec.md` with valid frontmatter, when parsed,
  then returns SpecSummary with all fields populated
- [ ] Given a `specs/0001-foo/tasks.md` with 3 `[x]`, 1 `[~]`, 2 `[ ]`, when
  parsed, then TaskStats = {Total: 6, Done: 3, InProgress: 1, Pending: 2}
- [ ] Given a `specs/0002-bar/spec.md` without frontmatter, when parsed, then
  spec is placed in "Ungrouped" family and a warning is emitted
- [ ] Given a `specs/0003-baz/` folder with no `plan.md`, when parsed, then
  HasPlan = false
- [ ] Given `docs/prds/2026-01-15-onboarding.md` with title in frontmatter,
  when parsed, then PRD has title and date 2026-01-15
- [ ] Given `docs/prds/2026-03-20-billing.md` without frontmatter, when parsed,
  then PRD date extracted from filename, title from first heading or filename
  slug
- [ ] Given 100 spec folders on disk, when scanned, then completes in under 1
  second
- [ ] Given specs with family="onboarding" phase=0 and phase=1, when grouped,
  then returned in two separate Phase structs under one Family
- [ ] Given two specs where 0005 depends_on [0001, 0003], when parsed, then
  SpecSummary.DependsOn = [1, 3]

## Edge Cases

- `spec.md` exists but is empty (zero bytes) — skip with warning
- `tasks.md` has nested checkboxes (indented `- [x]`) — count all levels
- `depends_on` contains a string spec number ("0001") — parse to int
- `specs/` directory does not exist — return empty result, no error
- `docs/prds/` directory does not exist — return empty PRD list, no error
- Folder name does not match `NNNN-<slug>` pattern — skip with warning
- Duplicate spec numbers across folders — warn, keep first found

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Frontmatter parser | Extend existing `internal/cli/frontmatter.go` | Already handles YAML frontmatter for instruction files; avoid new dependency |
| Markdown body | Store as raw string, not parsed AST | HTML rendering is downstream (0002); parser just extracts text |
| Checkbox regex | `(?m)^\\s*- \\[([ x~])\\]` | Handles indented subtasks; matches the three states |
| PRD date source | Frontmatter `date` field → filename fallback | Frontmatter is authoritative; filename is always available |

## Out of Scope

- HTML rendering of spec content (spec 0002)
- Git date extraction (spec 0004)
- Serving parsed data over HTTP (spec 0005)
- Dependency graph computation (spec 0006)
