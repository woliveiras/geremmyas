# AGENTS.md

This file is the operating contract for coding agents in this repository.
Follow it before using generic global defaults.

## Instruction Order

1. Treat the nearest project-local `AGENTS.md` as the active operating contract.
2. Use `.github/copilot-instructions.md` as supporting context when you need
   project overview, commands, or repository conventions.
3. Load `.github/instructions/*.instructions.md` only when their `applyTo`
   patterns match the files being edited.
4. Use `.github/skills/*/SKILL.md` for explicit workflows.
5. Use `.github/agents/*.agent.md` only when a named role is useful.

If a global instruction conflicts with this file, this file wins.

## Core Rules

- Do not change production code before understanding the relevant spec, task,
  bugfix document, or user request.
- Do not modify tests just to make them pass. If a test appears wrong, revisit
  the spec or acceptance criteria first.
- Do not write production code for a feature that has a PRD but no spec. The PRD
  frames the problem; the spec defines testable behavior. Both must exist before
  implementation.
- Always create a bugfix document in `docs/bugfixes/` for every bug. No
  exceptions.
- Always write a regression test for every bug fix. No exceptions.
- Do not create postmortems unless the bug caused a production outage.
- Do not create ADRs unless the decision is complex AND hard to reverse.
- Do not create `CONSTITUTION.md` by default. Use this file for operational
  rules unless the user explicitly asks for a separate constitution.
- Do not commit, amend, or push without explicit user confirmation and without
  commit permission granted during `requirements-interview`.
- Preserve user changes. Never revert unrelated work.
- Keep each feature's `tasks.md` checkbox state current during work: mark `- [~]`
  when starting a task and `- [x]` when it is done. Do not leave stale
  checkboxes.

## Approval Gates

These gates are mandatory. Do not skip them when using skills or agents
directly (not only when using the SDD prompt).

- **Spec gate:** After creating or updating `spec.md`, `plan.md`, and
  `tasks.md`, present the spec to the user and **stop**. Do not write production
  code or tests for the feature until the user explicitly approves the spec.
- **Bugfix gate:** After documenting reproduction and a proposed fix in the
  bugfix document, present the proposal to the user and **stop**. Do not apply
  the fix until the user explicitly approves.

## Domain Vocabulary

Read `GLOSSARY.md` before writing PRDs, specs, tests, reviews,
bugfix documents, ADRs, or user-facing copy.

Absence of this file does not block work.

## Artifact Locations

Use these paths unless the project already documents a different convention:

| Artifact | Default path |
| --- | --- |
| Specs index | `specs/README.md` (families, status tables, numbering) |
| PRD | `docs/prds/YYYY-MM-DD-<feature-slug>.md` |
| Feature folder | `specs/NNNN-<feature-slug>/` |
| Feature spec | `specs/NNNN-<feature-slug>/spec.md` |
| Feature plan | `specs/NNNN-<feature-slug>/plan.md` |
| Feature tasks | `specs/NNNN-<feature-slug>/tasks.md` |
| Repo-level tasks | `tasks.md` |
| Bugfix document | `docs/bugfixes/YYYY-MM-DD-<bug-slug>.md` |
| Postmortem | `docs/postmortems/YYYY-MM-DD-<incident-slug>.md` |
| ADR | `docs/decisions/NNNN-title-with-dashes.md` |

Use the local date when creating timestamped PRDs, bugfixes, and postmortems.
Spec folders use a **global four-digit number** plus slug: `NNNN-<feature-slug>`.
Slugs are lowercase kebab-case and describe the user-visible capability.

Read and maintain `specs/README.md` when creating, approving, or completing a
spec: update family tables, status, and reserved blocks as needed.

For every new feature, always create the feature folder with **all three**
artifacts: `spec.md`, `plan.md`, and `tasks.md`. Do not implement until all
three exist and the user has approved the spec.

## Workflows

### New Features

1. Use `requirements-interview` to explore code and clarify requirements. At
   the start, ask whether the agent may create git commits or the developer
   handles commits. Store the answer for the session.
2. Decide whether a PRD is needed (product behavior framing) or a spec alone
   is enough (behavior already clear).
3. If a PRD is needed: write or update the PRD, then write the spec from the
   PRD. **Do not write production code until the spec exists.**
4. Always create the feature folder with `spec.md`, `plan.md`, and `tasks.md`
   (use `generate-spec`, `task-breakdown`, and/or `spec-writer` as needed).
5. **Approval gate:** Present the spec (and plan/tasks summary) to the user.
   Stop and wait for explicit approval before implementation or test generation
   for the feature.
6. For each task in `tasks.md`, choose test type (`unit`, `integration`, or
   `both`) from the spec's test strategy and task scope.
7. Use `vertical-tdd` to implement one behavior at a time (red-green-refactor).
8. Use `reviewer` for spec-driven review.
9. Use `update-docs` when API, architecture, setup, or configuration changed.
10. Use `git-commit` only after verification, explicit confirmation, and only
    if the user granted commit permission in step 1. Otherwise report changed
    files and leave committing to the developer.

### Existing Features

1. Decide whether the product flow changes.
2. If the product flow changes, update the PRD first, then update the spec.
3. If the product flow does not change, write or update targeted specs in the
   feature folder (`spec.md`, `plan.md`, `tasks.md` as needed).
4. **Approval gate** applies when the spec changes materially.
5. Continue through tasks, tests, implementation, review, and docs.

### Bugs

1. Use `bugfix-loop`.
2. Always save the bugfix document under `docs/bugfixes/`.
3. Build a reproduction loop before changing production code.
4. Document hypotheses and proposed fix; **approval gate:** present the
   proposal and stop until the user approves.
5. Add or update a regression test at the correct boundary (mandatory).
6. Apply the fix and rerun the original reproduction.
7. Write a postmortem only when the bug was a production outage.

## Progress and resumption

When switching sessions, tools, or pausing work:

- Read `specs/README.md` for spec status across the repository.
- Open the feature folder `specs/NNNN-<slug>/` for durable truth on one feature.
- Use `tasks.md` checkboxes as the progress signal: `- [ ]` pending, `- [~]`
  in progress, `- [x]` done.
- Update checkboxes as work moves; do not create separate handoff documents.
- Resume by reading `spec.md`, `plan.md`, and `tasks.md`, then continue from
  the `- [~]` task or the first `- [ ]` after completed work.

## Skill Routing

Use these skills instead of reimplementing their procedures inline:

- `requirements-interview`: clarify product and technical requirements; commit
  permission at start.
- `generate-spec`: write spec, plan, and tasks in a feature folder.
- `task-breakdown`: create or update vertical tasks in `tasks.md`.
- `generate-tests-from-spec`: generate tests from acceptance criteria (after
  spec approval).
- `vertical-tdd`: implement one behavior per red-green-refactor cycle.
- `bugfix-loop`: investigate and fix bugs with reproduction and regression.
- `generate-glossary`: create or update domain vocabulary.
- `generate-adr`: record durable architecture decisions (bar: complex and hard
  to reverse).
- `update-docs`: sync documentation after implementation.
- `git-commit`: inspect staged changes and create a commit with confirmation.

Do not create GitHub Issues, labels, or issue state workflows unless the user
explicitly asks.

Use matching `.github/instructions/*.instructions.md` for edits in a single
technology, and any workflow skills installed from geremmyas packs (for example
Terraform, GCP, CI, LLM, Supabase, Postgres, ChromaDB) when the task crosses
files, needs sequencing, or has approval or verification gates.

## Agent Routing

- Use `spec-writer` when requirements are unclear and a spec is needed.
- Use `reviewer` for spec-driven review.
- Use `architect` for architecture exploration with multiple design options.
- Use `explorer` for read-only project mapping.

## Verification

Before saying work is complete:

- Run the focused tests for the changed behavior.
- Run the nearest relevant suite when practical.
- Report any verification that could not be run.
- Check that temporary logs, harnesses, and instrumentation were removed.
- Check `git status --short` and explain remaining changes.
