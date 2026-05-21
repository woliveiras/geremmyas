---
name: agent-brief
description: "Create a durable local AFK agent brief from a spec, plan, or task. Use when: preparing work for another agent, creating AFK execution briefs, turning a task into a self-contained implementation contract, or documenting what an agent should do without human context."
---

# Agent Brief

Create a local brief that an AFK agent can execute without relying on the
current conversation.

## Location

Save briefs under the feature folder:

```text
specs/YYYY-MM-DD-<feature-slug>/agent-briefs/<task-slug>.md
```

If the work is repo-level and has no feature folder, save under:

```text
docs/agent-briefs/YYYY-MM-DD-<task-slug>.md
```

Create directories lazily. Use lowercase kebab-case slugs.

## Process

1. Read the source artifacts: spec, `plan.md`, `tasks.md`, bugfix document, ADR,
   or docs referenced by the user.
2. Re-read `GLOSSARY.md` or `CONTEXT.md` when either exists.
3. Confirm the task is `AFK`: clear desired behavior, testable acceptance
   criteria, local verification, no open human decision, no destructive
   operation, no required external credentials.
4. If the task is not AFK, write why it is `HITL` instead of creating an AFK
   brief.
5. Write the brief with durable context: behavior, contracts, interfaces,
   verification, scope, and risks.
6. Link the brief from the relevant task entry when editing `tasks.md`.

## Brief Template

````markdown
# Agent Brief: <task title>

**Category:** feature | bugfix | refactor | docs | infra
**Type:** AFK
**Source artifacts:** spec, plan, task, bugfix, ADR, or docs paths

## Summary

One or two sentences describing the work.

## Current Behavior

What happens now, or the current baseline the work builds on.

## Desired Behavior

What should be true when the agent finishes.

## Key Interfaces and Contracts

- Type, function, API, config, command, or data contract that matters

## Acceptance Criteria

- [ ] Specific, independently verifiable criterion

## Verification

```bash
<focused command>
```

## Out of Scope

- Adjacent work the agent must not do

## Risks and Constraints

- Known constraint, dependency, or risk

## Expected Final Report

- Files changed
- Tests or checks run
- Behavior verified
- Follow-ups or blockers
````

## Rules

- Do not use GitHub Issues or labels unless the user explicitly asks.
- Do not reference line numbers.
- Avoid fragile file paths. Prefer contracts, types, commands, behavior, and
  source artifact paths.
- Do not include secrets, credentials, production data, or private user data.
- Do not create AFK briefs for tasks that require human judgment, external
  access, destructive operations, or ambiguous requirements.

## Output

- Path to the brief
- Whether the task is AFK or HITL
- Missing information, if the task is not ready
