---
name: afk-task-triage
description: "Classify local tasks as AFK or HITL and prepare agent-ready work from specs, plans, and tasks.md. Use when: reviewing tasks for agent execution, deciding what can run without human input, splitting large tasks, or preparing AFK briefs."
---

# AFK Task Triage

Review local tasks and decide which ones can be delegated to an AFK agent.

## Process

1. Read the relevant spec, `plan.md`, `tasks.md`, bugfix document, ADRs, and
   domain vocabulary.
2. Preserve completed tasks and unrelated tasks.
3. For each open task, classify it:
   - `AFK`: clear behavior, testable acceptance criteria, local verification,
     no human decision, no destructive operation, no required external secrets.
   - `HITL`: needs product/architecture/design judgment, production access,
     credentials, sensitive data, destructive operation, or clearer scope.
4. Split large tasks until each AFK task is a thin vertical slice with an
   objective finish line.
5. Add or update task fields: `type`, `blocked-by`, `acceptance criteria`,
   `verification`, and `brief`.
6. Use `agent-brief` for each task that is ready for AFK execution.
7. Report AFK-ready tasks, HITL blockers, dependencies, and suggested next
   brief to execute.

## AFK Criteria

A task is AFK only when all are true:

- Desired behavior is specific.
- Acceptance criteria are independently verifiable.
- Verification can run locally or through documented non-destructive commands.
- Required context is in artifacts, code, docs, or repository instructions.
- The agent can complete the task without asking a human to choose direction.

## HITL Criteria

Mark a task as HITL when any are true:

- Requirements are ambiguous.
- Product, UX, or architecture choice is still open.
- Production, security-sensitive data, credentials, or destructive operations
  are required.
- There is no objective verification.
- The task is too broad to fit one focused agent session.

## Rules

- Keep this workflow local-first. Do not create GitHub Issues by default.
- Prefer splitting tasks over marking a broad task AFK.
- Do not downgrade a real human decision into an AFK task.
- Use `GLOSSARY.md` or `CONTEXT.md` vocabulary when either exists.
- Do not mark tasks as in-progress or deprecated; use checked boxes and Git
  history.

## Output

- Updated task classification
- AFK-ready task list
- HITL blocker list
- Created or recommended agent brief paths
