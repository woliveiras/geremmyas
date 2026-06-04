---
spec: "0002"
title: Dashboard overview HTML generation
family: platform
phase: 0
status: Implemented
owner: ""
depends_on: [1]
origin: "docs/prds/2026-06-04-dashboard.md"
---

# Spec: Dashboard overview HTML generation

## Context & Motivation

With parsed spec data (0001), the next step is producing visible output: a
static HTML overview page and a compact auto-generated `specs/README.md`.

The overview page is the primary entry point for humans (devs, PMs, POs). The
auto-generated README is the primary entry point for coding agents. Together
they replace the hand-maintained `specs/README.md` index that breaks at scale.

This spec also updates the `sdd` pack template to document frontmatter fields,
ensuring new projects produce dashboard-compatible specs from day one.

## Requirements

### Functional

- [ ] Register `dashboard` command in CLI switch with `--output` flag
- [ ] Generate `index.html` overview page with:
  - Global summary bar (total specs, implemented, in-progress, pending counts)
  - Family cards showing: name, goal, spec count, status breakdown (color-coded
    badges), active phase, progress percentage
  - Click family card links to family detail page
- [ ] Generate one family detail page per family:
  - Specs grouped by phase
  - Phase progress bars (done/total)
  - Spec rows: number, title, status badge, task progress bar
  - Click spec links to spec detail page
- [ ] Generate one spec detail page per spec:
  - Rendered markdown body (spec.md content → HTML)
  - Sidebar: status, family, phase, owner, depends_on links, origin link
  - Task progress from tasks.md (bar + counts)
  - Links to plan.md and tasks.md rendered pages (if they exist)
- [ ] Auto-generate `specs/README.md` in compact format:
  - Header: `<!-- geremmyas:generated -->` marker
  - One line per spec: `- NNNN Title — Status`
  - Grouped by family with `## Family Name (NNNN–NNNN)` headings
  - Sorted by spec number within each family
- [ ] Embed all templates and CSS via `go:embed` in the binary
- [ ] Default output directory: `.geremmyas/dashboard/`
- [ ] Update `sdd` pack template (`project/templates/specs/README.md`) to
  document required frontmatter fields for dashboard compatibility

### Non-Functional

- [ ] Generated HTML works offline — no CDN, no external requests
- [ ] Dark/light theme via `prefers-color-scheme` CSS media query
- [ ] Responsive layout for screen sharing in dailies (min 1024px)
- [ ] Generation of 100 specs completes in under 2 seconds (excluding git)

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Template rendering, README generation | render family card, format README line |
| **integration** | End-to-end: parse → generate → verify HTML output | generate dashboard from temp spec tree, check files exist |

Default: **integration** for the command, **unit** for individual renderers.

## Acceptance Criteria

- [ ] Given a project with 3 families and 15 specs, when `geremmyas dashboard`
  runs, then `.geremmyas/dashboard/index.html` exists and contains one card per
  family
- [ ] Given a spec with status "Implemented", when rendered on overview, then
  the status badge uses a green color class
- [ ] Given `--output ./site`, when dashboard runs, then HTML files are written
  to `./site/` instead of default
- [ ] Given a spec folder with spec.md, plan.md, and tasks.md, when dashboard
  runs, then the spec detail page renders all three with navigation links
- [ ] Given a spec without plan.md, when dashboard runs, then the spec detail
  page omits the plan link without error
- [ ] Given 5 specs across 2 families, when dashboard runs, then
  `specs/README.md` is overwritten with compact format containing exactly 5
  spec lines grouped under 2 family headings
- [ ] Given `specs/README.md` already has custom content, when dashboard runs,
  then the file is replaced (content is auto-generated, not preserved)
- [ ] Given the generated `specs/README.md`, when an agent reads it, then each
  spec line contains number, title, and current status
- [ ] Given `index.html` opened in a browser with `prefers-color-scheme: dark`,
  then the page renders with dark background and light text
- [ ] Given a spec with `depends_on: [1, 3]`, when rendered on detail page,
  then dependencies appear as clickable links to specs 0001 and 0003

## Edge Cases

- Family with zero non-deprecated specs — show card with "0 active specs" label
- Spec body contains HTML entities or special characters — escaped properly
- Very long spec title — truncated with ellipsis on overview cards
- `--output` path does not exist — create directory tree automatically
- Running dashboard in repo with no `specs/` directory — generate empty
  dashboard with "No specs found" message

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Template engine | Go `html/template` | Stdlib, no dependency; sufficient for static pages |
| CSS framework | Custom minimal CSS | Pico/Classless adds dependency; hand-written is < 500 lines for this scope |
| Markdown → HTML | Use goldmark (Go markdown parser) or simple regex | goldmark is widely used in Go ecosystem; evaluate at implementation |
| README format | Flat list, not tables | Tables are harder for agents to parse incrementally; one-line format is grep-friendly |
| Embed strategy | `dashboard_assets/` directory with `go:embed` | Separate from existing catalog/project/user embeds |

## Out of Scope

- Board/kanban view (spec 0003)
- Git date extraction and metrics (spec 0004)
- Live server and watch mode (spec 0005)
- PRD/bugfix linking and dependency graph (spec 0006)
- Chart.js or any charting (spec 0004)
