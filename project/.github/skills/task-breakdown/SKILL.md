---
name: task-breakdown
description: "Break a PRD, spec, or plan into vertical tasks in tasks.md. Use when: creating tasks, breaking down work, converting a plan to tasks, planning agent-ready implementation steps, or preparing AFK/HITL task structure."
---

# Task Breakdown

Create or update `tasks.md` from a PRD, spec, or `plan.md`.

## Process

1. Read the source artifact the user provides. If none is provided, search for
   the relevant PRD, spec, or `plan.md`.
2. Read existing `tasks.md` if present and preserve completed or unrelated work.
3. Break the work into vertical slices: each task should deliver a narrow,
   verifiable behavior across the layers it needs.
4. Mark each task as:
   - `AFK`: an agent can execute with local context and verification.
   - `HITL`: a human decision, review, or external access is required.
5. For each task, include dependencies, desired behavior, acceptance criteria,
   test expectations, and verification commands.
6. When a task is ready for AFK execution, add a `brief` field pointing to the
   expected local agent brief path, or route to `agent-brief` to create it.

## Task Shape

Each task should include:

- title
- type: `AFK` or `HITL`
- priority
- blocked-by list
- summary
- desired behavior
- acceptance criteria
- verification
- brief: path to `agent-briefs/<task-slug>.md` when AFK-ready

## Rules

- Prefer thin vertical slices over horizontal layer tasks.
- Avoid tasks like "create schema", "create endpoint", or "create UI" unless
  they are only setup work for a separately verifiable behavior.
- Keep local files as the primary workflow. Do not turn work into GitHub Issues
  unless the user explicitly asks.
- Use `afk-task-triage` when existing tasks need classification, splitting, or
  readiness review before AFK execution.
- Use the project's domain vocabulary from `GLOSSARY.md` or `CONTEXT.md` when
  either exists.
