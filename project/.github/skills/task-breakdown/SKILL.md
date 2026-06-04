---
name: task-breakdown
description: "Break a PRD, spec, or plan into vertical tasks in tasks.md. Use when: creating tasks, breaking down work, converting a plan to tasks, or tracking implementation progress with checkboxes."
---

# Task Breakdown

Create or update `tasks.md` from a PRD, spec, or `plan.md` in the feature
folder `specs/NNNN-<feature-slug>/tasks.md`.

## Process

1. Read the source artifact the user provides. If none is provided, search for
   the relevant PRD, `spec.md`, or `plan.md` in `specs/`.
2. Read existing `tasks.md` if present and preserve completed or unrelated work.
3. Break the work into vertical slices: each task should deliver a narrow,
   verifiable behavior across the layers it needs.
4. For each task, set **test-type**: `unit`, `integration`, or `both`, using the
   spec's Test Strategy and whether the task touches I/O or multiple modules.
5. List tasks with checkbox status for progress tracking:
   - `- [ ]` pending
   - `- [~]` in progress
   - `- [x]` done
6. For each task, include dependencies (blocked-by), desired behavior, acceptance
   criteria, test expectations, verification commands, and test-type.

## Task Shape

Use checkbox lines in `tasks.md` plus detail under each task when needed:

```markdown
## Tasks

- [ ] **Task title** (test-type: unit | integration | both)
  - blocked-by: ...
  - summary: ...
  - desired behavior: ...
  - acceptance criteria: ...
  - verification: `command`
```

## Rules

- Prefer thin vertical slices over horizontal layer tasks.
- Avoid tasks like "create schema", "create endpoint", or "create UI" unless
  they are only setup work for a separately verifiable behavior.
- Keep local files as the primary workflow. Do not turn work into GitHub Issues
  unless the user explicitly asks.
- Checkbox state in `tasks.md` is the source of truth for progress. Keep
  checkboxes current during implementation (`[~]` while in progress, `[x]` when
  done). A new session resumes via `specs/README.md` and the feature folder
  (`spec.md`, `plan.md`, `tasks.md`).
- Use the project's domain vocabulary from `GLOSSARY.md` or `CONTEXT.md` when
  either exists.
- Do not start implementation until the user has approved the spec (see
  `AGENTS.md` approval gates).
