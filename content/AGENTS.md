# AGENTS.md

This is the operating contract. The nearest
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
- Feature artifacts advance through `Draft` -> `Ready` -> `In Progress` ->
  `Verified`. A machine-ready package is established by automatic checks, not
  human approval.
- Honor explicit session overrides: read-only, plan-only, no-edits, or
  no-commits.
- Every bug needs `docs/bugfixes/YYYY-MM-DD-<slug>.md`, a reproduction, an
  evidence-backed fix proposal, and a regression test that fails before the fix.
- Create postmortems only for production outages and ADRs only for complex,
  hard-to-reverse decisions.
- Preserve user work and never revert unrelated changes.
- After each coherent task-owned slice passes fresh tests, docs reconciliation,
  and independent review, the primary agent creates an atomic local Conventional
  Commit by default, even while the feature remains `In Progress`. Do not push;
  commit authority excludes history rewrite, merge, tag, release, publication,
  and production changes.
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

1. Use `requirements-interview` to inspect existing behavior and resolve
   ambiguity. Update the PRD first when product flow changes.
2. Use `generate-spec` to create or update the three feature artifacts.
3. Validate the artifacts and mark them `Ready` when their automatic readiness
   checks pass. Continue with `vertical-tdd` one behavior at a time and set the
   feature to `In Progress` when implementation starts.
4. If implementation changes an in-scope assumption, update the artifacts,
   rerun their readiness checks, and continue while the objective is unchanged.
5. Escalate before implementation only when blast radius is high and either
   rollback is unproven or critical harness evidence is missing. After
   acceptance criteria, tasks, docs, fresh verification, and independent review
   are reconciled with no actionable finding, mark it `Verified`.
6. Use `update-docs` when API, architecture, setup, or config changed.
7. Commit each task-owned slice after its evidence passes; do not wait for the
   whole feature to reach `Verified`.

### Bugs

Use `bugfix-loop`. Reproduce before production edits, rank hypotheses, document
the proposed fix, and prove the regression test fails. Then apply the smallest
root-cause fix, rerun the regression test, original reproduction, and nearby
suite, remove temporary instrumentation, and record the actual cause. Continue
autonomously unless the objective materially expands, production or external
authority is required, or the same blocker survives three evidence-driven
cycles. Record subjective or unavailable checks as residual evidence instead of
blocking local completion.

## Agent Routing

- `explorer`: expensive read-only mapping across many files.
- `spec-writer`: unclear requirements that need isolated exploration.
- `reviewer`: implementation review against a `Ready` spec.
- `architect`: material architecture options after ordinary exploration.

Work inline for a small query or narrow edit. Delegate independent, read-heavy
work when the returned summary will be smaller than the exploration. Never
delegate shared-state edits or redo a subagent's completed exploration inline.
The primary agent alone stages and commits integrated work.

## Completion

Before claiming completion:

1. Run focused tests and the nearest relevant suite.
2. Confirm acceptance criteria, error paths, and required regression coverage.
3. Remove temporary logs, harnesses, and instrumentation.
4. Update `tasks.md`, reconcile `plan.md`, and update spec/index status.
5. Run `git status --short` and explain remaining changes.

Shell guardrails live in `.github/hooks/` for Copilot and generated Cursor hooks.
