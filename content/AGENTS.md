# AGENTS.md

This is the operating contract. The nearest project-local `AGENTS.md` wins.

## Instruction Order

1. Read this contract and the target's project overview.
2. Read `GLOSSARY.md` before domain artifacts or user copy when present.
3. Load technology instructions only for edited files.
4. Load skills only when their trigger matches the current phase.

## Lazy Routing

Classify the request before loading workflow skills:

- Simple question, lookup, trivial edit, or clear docs: load no workflow skill.
- Material ambiguity: load only `refine`, ask one evidence-backed question at a
  time, then stop loading it when the objective is clear.
- Clear feature: start with `spec`; load `tdd` only during implementation.
- Clear bug: start with `bugfix`.
- Completion: load `verify` when installed; otherwise follow Completion below.
- Review: load `verify/references/review-contract.md` when installed; otherwise
  apply Agent Routing below as the bounded contract.
- Durable docs: load `docs` when installed; otherwise update the smallest
  relevant surface directly from repository evidence.

Never preload an entire pack or future-phase skills into the task context.

## Invariants

- Understand the request or governing artifact before code. Never change tests
  merely to pass; reconcile behavior with the spec.
- Every feature needs `spec.md`, `plan.md`, and `tasks.md` before implementation.
- Lifecycle is `Draft` -> `Ready` -> `In Progress` -> `Verified`; automatic
  checks establish machine-ready state. Honor read-only, plan-only, no-edits,
  or no-commits overrides.
- Every bug needs `docs/bugfixes/YYYY-MM-DD-<slug>.md`, a reproduction, an
  evidence-backed fix proposal, and a regression test that fails before the fix.
- Postmortems require production outages; ADRs require complex,
  hard-to-reverse decisions. Preserve unrelated user work.
- An atomic local Conventional Commit by default follows each verified slice.
  The feature remains `In Progress` while tasks remain. Do not push or infer
  history rewrite, merge, tag, release, publication, or production authority.
- Existing project dependencies and catalogued capabilities are autonomous. A
  new uncatalogued direct dependency needs provenance, maintenance, security,
  license, and build-versus-buy evidence plus explicit user choice before installation.
- Mutate a verified local, disposable, or test target with rollback or recreation.
  Treat an ambiguous target as protected. Every production mutation, deploy,
  release, publication, or policy change needs explicit user authorization.
- Keep `tasks.md` current: `[~]` active, `[x]` verified.

## Artifacts

- PRD: `docs/prds/YYYY-MM-DD-<feature-slug>.md`
- Feature: `specs/NNNN-<feature-slug>/{spec,plan,tasks}.md`
- Spec index: `specs/README.md`
- Bugfix: `docs/bugfixes/YYYY-MM-DD-<bug-slug>.md`
- Postmortem: `docs/postmortems/YYYY-MM-DD-<incident-slug>.md`
- ADR: `docs/decisions/NNNN-title-with-dashes.md`

Use local dates, kebab-case, and four-digit spec numbers. Maintain the index.

## Work Routing

### Features and expansions

1. Use `refine` only for ambiguity. Update the PRD when product flow changes.
2. Use `spec` to create or update the three feature artifacts.
3. Mark machine-ready artifacts `Ready`; start `tdd` and set `In Progress`.
4. Revalidate in-scope assumptions without changing the objective.
5. Escalate only when blast radius is high and either
   rollback is unproven or critical harness evidence is missing. After
   acceptance criteria, tasks, docs, fresh verification, and independent review
   are reconciled with no actionable finding, mark it `Verified`.
6. Use `docs` when installed, or its compact fallback above; commit each slice.

### Bugs

Use `bugfix`. Reproduce before production edits, rank hypotheses, document
the proposed fix, and prove the regression test fails. Apply the smallest fix;
rerun it, the original reproduction, and nearby suite; remove temporary
instrumentation; record the actual cause. Continue unless scope or authority
changes, or a blocker survives three evidence-driven cycles. Record unavailable
checks as residual evidence.

## Agent Routing

Create a runtime subagent only when specialization, isolation, or parallelism
helps. Give editors explicit ownership; parallel edits must be disjoint. The
primary agent owns integration and Git; subagents never commit. Run independent
review after each slice. The primary repairs actionable findings and re-reviews.
Escalate after three consecutive cycles. Without subagents, work inline with the
same bounded contract. Keep trivial work inline.

## Completion

Before claiming completion:

1. Run focused tests and the nearest relevant suite.
2. Confirm acceptance criteria, error paths, and required regression coverage.
3. Remove temporary logs, harnesses, and instrumentation.
4. Update `tasks.md`, reconcile `plan.md`, and update spec/index status.
5. Run `git status --short` and explain remaining changes.

Shell guardrails live in `.github/hooks/` for Copilot and generated Cursor hooks.
