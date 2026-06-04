---
spec: "0003"
title: Dashboard board view
family: platform
phase: 1
status: Approved
owner: ""
depends_on: [2]
origin: "docs/prds/2026-06-04-dashboard.md"
---

# Spec: Dashboard board view

## Context & Motivation

The overview page (0002) shows specs grouped by family and phase. PMs and POs
also need a kanban-style board grouped by status — the standard view for
tracking work in progress.

A single global board with 100+ specs would be unreadable. Instead, each family
gets its own kanban board embedded in the family detail page. This keeps the
board scoped and useful.

## Requirements

### Functional

- [ ] Add kanban board section to each family detail page with four columns:
  Draft, In Review, Approved, Implemented
- [ ] Each spec appears as a card in its status column showing: spec number,
  title, phase badge, task progress bar (done/total)
- [ ] Deprecated specs hidden by default with a toggle to show them (fifth
  column when visible)
- [ ] Filter dropdown: by phase (within the family)
- [ ] Filters are client-side JavaScript (no server required)
- [ ] Card click navigates to the spec detail page (from 0002)
- [ ] Board is a section within the family page (not a separate page) — users
  can toggle between list view and board view

### Non-Functional

- [ ] Board renders correctly with 50 spec cards per family without layout
  breakage
- [ ] Filters respond instantly (< 100ms, client-side DOM filtering)
- [ ] No external JS dependencies — vanilla JavaScript only
- [ ] Works offline (same as 0002)

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Card grouping logic, filter data attributes | group specs by status, count per column |
| **integration** | Generate board page and verify structure | temp spec tree → dashboard → check board.html content |

## Acceptance Criteria

- [ ] Given a family with 10 specs (3 Draft, 2 In Review, 3 Approved, 2
  Implemented), when board renders, then each column contains the correct
  count of cards
- [ ] Given a spec card, when rendered, then it shows task progress bar with
  correct percentage (e.g., 4/6 tasks = 67%)
- [ ] Given the phase filter set to "Phase 1", when applied, then only specs
  in Phase 1 of that family are visible
- [ ] Given 2 deprecated specs in the family, when board loads, then they are
  hidden; when toggle is clicked, then they appear in a "Deprecated" column
- [ ] Given a spec card, when clicked, then navigates to the spec detail page
- [ ] Given a family with 50 spec cards, when board renders, then columns
  scroll vertically without horizontal overflow
- [ ] Given the family detail page, when user toggles between list and board
  view, then the view switches without page reload

## Edge Cases

- Status value not in the known set (typo in frontmatter) — place in "Draft"
  column with a warning icon
- Spec has no tasks.md — card shows "No tasks" instead of progress bar
- All specs in one status — other columns show empty state ("No specs")
- Phase filter matches zero specs — show "No matches" message
- Family with only 1 spec — board still renders correctly (single card)

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Filter implementation | CSS class toggle via JS | Simpler than re-rendering; works offline |
| Card layout | CSS Grid columns, cards as flex items | Responsive, handles variable card counts |
| Progress bar | HTML `<progress>` element + CSS | Semantic, accessible, no JS needed |
| Deprecated toggle | Checkbox input + CSS `:checked` selector | Zero JS for the toggle itself |

## Out of Scope

- Drag-and-drop between columns (this is read-only)
- Editing spec status from the board
- Metrics or charts on the board page (spec 0004)
