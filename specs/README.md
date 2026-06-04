# Specs (geremmyas)

Index for **geremmyas platform** work in this repository. Consumer projects
get a fuller template via the `sdd` pack (`project/templates/specs/README.md`).

Each feature uses a numbered folder with `spec.md`, `plan.md`, and `tasks.md`.
Progress lives in `tasks.md` checkboxes (`[ ]`, `[~]`, `[x]`). See
[`AGENTS.md`](../AGENTS.md).

## Families

| Family | Reserved block | Goal |
| --- | --- | --- |
| Platform | 0001–0099 | CLI, catalog, packs, docs, release tooling |

## Platform

**Origin:** [docs/prds/2026-06-04-dashboard.md](../docs/prds/2026-06-04-dashboard.md)

### Phase 0 — Foundation + Overview

| Spec | Title | Status | Depends on |
| --- | --- | --- | --- |
| [0001](0001-dashboard-parser/spec.md) | Dashboard spec parser | Draft | — |
| [0002](0002-dashboard-overview/spec.md) | Dashboard overview HTML generation | Draft | 0001 |

### Phase 1 — Board & Progress

| Spec | Title | Status | Depends on |
| --- | --- | --- | --- |
| [0003](0003-dashboard-board/spec.md) | Dashboard board view | Draft | 0002 |

### Phase 2 — Metrics

| Spec | Title | Status | Depends on |
| --- | --- | --- | --- |
| [0004](0004-dashboard-metrics/spec.md) | Dashboard git dates and metrics | Draft | 0002 |

### Phase 3 — Developer Experience

| Spec | Title | Status | Depends on |
| --- | --- | --- | --- |
| [0005](0005-dashboard-serve/spec.md) | Dashboard serve and watch mode | Draft | 0002 |

### Phase 4 — Navigation

| Spec | Title | Status | Depends on |
| --- | --- | --- | --- |
| [0006](0006-dashboard-linking/spec.md) | Dashboard artifact linking and dependency visualization | Draft | 0002, 0004 |

### Reserved blocks (future)

| Family | Reserved block |
| --- | --- |
| Packs & content | 0100–0199 |
| Multi-IDE | 0200–0299 |

## Status lifecycle

`Draft` → `In Review` → `Approved` → `Implemented` → `Deprecated`

Update this table when creating, approving, or completing a spec.
