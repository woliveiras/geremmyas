---
spec: "0006"
title: Dashboard artifact linking and dependency visualization
family: platform
phase: 4
status: Approved
owner: ""
depends_on: [2, 4]
origin: "docs/prds/2026-06-04-dashboard.md"
---

# Spec: Dashboard artifact linking and dependency visualization

## Context & Motivation

Specs do not exist in isolation. They originate from PRDs, depend on other
specs, and relate to bugfix documents. Navigating these relationships in raw
markdown requires manual cross-referencing across directories.

This spec adds bidirectional linking between artifacts and a dependency
visualization to the dashboard, completing the navigation story.

## Requirements

### Functional

- [ ] PRD detail pages: render `docs/prds/*.md` as HTML pages in the dashboard
- [ ] Bugfix detail pages: render `docs/bugfixes/*.md` as HTML pages
- [ ] Spec detail page: link origin field to PRD detail page (if origin matches
  a known PRD path)
- [ ] PRD detail page: show list of specs that reference this PRD via `origin`
  field (reverse link)
- [ ] Spec detail page: `depends_on` specs rendered as clickable links
- [ ] Spec detail page: show "blocked by" (specs this depends on that are not
  Implemented) and "blocks" (specs that depend on this one)
- [ ] Family detail page: dependency list per phase showing which specs are
  blocked
- [ ] Add PRDs and Bugfixes sections to the navigation bar
- [ ] PRDs index page: list of all PRDs with date, title, linked spec count
- [ ] Bugfixes index page: list of all bugfixes with date, title

### Non-Functional

- [ ] Dependency resolution completes in O(n) where n = total specs
- [ ] Pages render correctly when no PRDs or bugfixes exist (empty state)
- [ ] Links to non-existent artifacts (e.g., origin pointing to missing PRD)
  show a "not found" indicator instead of a broken link

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Dependency graph computation, reverse link building | build adjacency list, find blocked specs |
| **integration** | End-to-end linking across generated pages | temp tree with PRDs + specs + deps → verify links in HTML |

## Acceptance Criteria

- [ ] Given spec 0005 with `origin: "docs/prds/2026-01-15-onboarding.md"`, when
  rendered, then the origin field is a clickable link to the PRD detail page
- [ ] Given a PRD referenced by specs 0001, 0002, 0005, when PRD detail page
  renders, then it shows a "Related specs" section listing all three
- [ ] Given spec 0010 with `depends_on: [1, 5]` where spec 0001 is Implemented
  and spec 0005 is Approved, when rendered, then "blocked by" shows spec 0005
  and "blocks" is empty
- [ ] Given spec 0005 that is depended on by specs 0010 and 0012, when
  rendered, then "blocks" section shows specs 0010 and 0012
- [ ] Given a family detail page with 3 specs where spec 0003 depends on 0001
  (Implemented), when rendered, then no specs are shown as blocked
- [ ] Given no `docs/prds/` directory, when dashboard runs, then PRDs nav link
  shows "No PRDs" empty state page
- [ ] Given spec with `origin: "docs/prds/nonexistent.md"`, when rendered, then
  origin shows text with "(not found)" indicator

## Edge Cases

- Circular dependencies (0001 → 0002 → 0001) — detect and warn, do not crash
- Spec depends on a spec number that does not exist — show "(missing)" label
- PRD referenced by 50+ specs — page still renders correctly (scrollable list)
- Bugfix document without frontmatter — extract date from filename, title from
  first heading
- Same PRD path referenced with different casing — normalize paths for matching

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Dependency visualization | Text list with status badges, not a graph diagram | Mermaid/D3 adds complexity; text list is sufficient for v1 |
| Reverse link computation | Build adjacency map at parse time | O(n) pass over all specs; no repeated scans |
| Circular dependency handling | Warn in output, render both links normally | Specs can still be browsed; the warning alerts the user |
| PRD matching | Exact path match on `origin` field | Simple, deterministic; user controls the path |

## Out of Scope

- Visual dependency graph (Mermaid, D3, etc.) — future enhancement
- Editing links from the dashboard
- Auto-detecting PRD↔spec relationships beyond the `origin` field
- Full-text search across all artifacts
