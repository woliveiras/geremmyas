# AGENTS.md

This is the operating contract for coding agents in this repository. The nearest
project-local `AGENTS.md` overrides global defaults.

## Instruction Order

1. Read this contract and the active assistant's project overview.
2. Read `GLOSSARY.md` before specs, tests, reviews, bugfix documents, ADRs, or
   user-facing copy when that file exists.
3. Load technology instructions only for files being edited.
4. Load a skill only when its trigger matches the current phase.
5. Use a custom agent only for an isolated role that materially benefits from a
   separate context.

## Invariants

- Understand the relevant request, spec, task, or bugfix document before code.
- Never change tests merely to make them pass; reconcile behavior with the spec.
- Every feature needs `spec.md`, `plan.md`, and `tasks.md` before implementation.
- After creating or materially changing those artifacts, present them and stop
  until the user explicitly approves the spec.
- Every bug needs `docs/bugfixes/YYYY-MM-DD-<slug>.md`, a reproduction, an
  approved fix proposal, and a regression test that fails before the fix.
- Create postmortems only for production outages and ADRs only for complex,
  hard-to-reverse decisions.
- Preserve user work and never revert unrelated changes.
- Do not commit, amend, or push without explicit permission and confirmation.
- Keep `tasks.md` current: `[~]` while active and `[x]` only after verification.

## Artifacts

- PRD: `docs/prds/YYYY-MM-DD-<feature-slug>.md`
- Feature: `specs/NNNN-<feature-slug>/{spec,plan,tasks}.md`
- Spec index: `specs/README.md`
- Bugfix: `docs/bugfixes/YYYY-MM-DD-<bug-slug>.md`
- Postmortem: `docs/postmortems/YYYY-MM-DD-<incident-slug>.md`
- ADR: `docs/decisions/NNNN-title-with-dashes.md`

Use local dates, lowercase kebab-case slugs, and global four-digit spec numbers.
Maintain `specs/README.md` whenever a spec is created or changes status.

## Work Routing

### Features and expansions

1. Use `requirements-interview` to inspect existing behavior, resolve ambiguity,
   and record commit permission. Update the PRD first when product flow changes.
2. Use `generate-spec` to create or update the three feature artifacts.
3. Stop at the approval gate. After approval, use `vertical-tdd` one behavior at
   a time, then `update-docs` when API, architecture, setup, or config changed.

### Bugs

Use `bugfix-loop`. Reproduce before production edits, rank hypotheses, document
the proposed fix, and stop for approval. Then add the regression test, apply the
fix, rerun the original reproduction, and remove temporary instrumentation.

### Explicit capabilities

- `generate-glossary`: establish domain vocabulary.
- `generate-adr`: record an accepted, durable architecture decision.
- `verification-checklists`: require fresh evidence before completion.
- `code-review-requesting`: prepare a verified change for review.
- `git-commit`: inspect and commit only explicitly approved files.

Stack-specific skills are opt-in. Use them only when the repository installs the
matching pack and the task crosses that technology boundary.

## Agent Routing

- `explorer`: expensive read-only mapping across many files.
- `spec-writer`: unclear requirements that need isolated exploration.
- `reviewer`: implementation review against an approved spec.
- `architect`: material architecture options after ordinary exploration.

Work inline for a small query or narrow edit. Delegate independent, read-heavy
work when the returned summary will be smaller than the exploration. Never
delegate shared-state edits or redo a subagent's completed exploration inline.

## Completion

Before claiming completion:

1. Run focused tests and the nearest relevant suite.
2. Confirm acceptance criteria, error paths, and required regression coverage.
3. Remove temporary logs, harnesses, and instrumentation.
4. Update `tasks.md`, reconcile `plan.md`, and update spec/index status.
5. Run `git status --short` and explain remaining changes.

Shell guardrails live in `.github/hooks/` for Copilot and generated Cursor hooks.
