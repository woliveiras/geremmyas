# Tasks: Autonomous agent workflows

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Specify autonomous agent workflows** (test-type: unit)
  - blocked-by: none
  - summary: Update the multi-assistant PRD and add spec 0008, plan, tasks, and
    index entry from the workflow-gate audit and resolved user decisions.
  - desired behavior: Autonomy, evidence, commit defaults, delegation, dependency
    provenance, escalation, and production boundaries are durable and consistent.
  - acceptance criteria: The PRD, spec, plan, tasks, and index agree; release
    workflow changes are explicitly deferred.
  - verification: Independent artifact review and `git diff --check`.
  - commit: `docs: specify autonomous agent workflows`

- [x] **Make feature delivery autonomous** (test-type: both)
  - blocked-by: Specify autonomous agent workflows
  - summary: Replace human feature approval with machine readiness across the
    contract, requirements, specification, TDD, prompts, templates, docs, and a
    durable full-catalog gate migration inventory.
  - desired behavior: An unambiguous feature proceeds from durable artifacts to
    implementation and verification without routine questions.
  - acceptance criteria: Contract tests cover exact lifecycle transitions,
    readiness, in-scope revalidation, blast-radius and rollback-based escalation,
    explicit narrow overrides, a full-catalog unclassified-gate scan, and
    consistent target materialization.
  - verification: Focused workflow/generator tests, `go test ./internal/cli`,
    `go run ./cmd/geremmyas lint`, and target materialization smoke tests.
  - evidence: Contract and scanner regressions, five-target materialization,
    full Go suite, build, doctor catalog check, context diagnostics, lint,
    `git diff --check`, and independent review with no remaining finding.
  - docs: README lifecycle, guardrails framework, template, and prompt guidance.
  - commit: `feat: make feature delivery autonomous`

- [x] **Run bugfixes as an autonomous evidence loop** (test-type: both)
  - blocked-by: Make feature delivery autonomous
  - summary: Remove bugfix-proposal approval while preserving reproduction,
    ranked hypotheses, red regression, verified fix, cause, and cleanup.
  - desired behavior: A reproducible bug is fixed and documented without human
    interruption unless scope, production authority, or a persistent blocker
    requires escalation.
  - acceptance criteria: Canonical and materialized workflows contain the full
    evidence loop and no pre-fix human approval gate.
  - verification: Focused policy tests, `go test ./internal/cli`, and
    `go run ./cmd/geremmyas lint`.
  - evidence: Canonical loop and five-target materialization tests, full Go
    suite, lint, `git diff --check`, and independent review with no findings.
  - docs: Bugfix lifecycle and residual-risk reporting.
  - commit: `feat: automate the bugfix evidence loop`

- [x] **Commit verified slices by default** (test-type: both)
  - blocked-by: Make feature delivery autonomous
  - summary: Remove the initial commit question and per-file/message approval;
    require task-owned atomic staging, diff-derived Conventional Commits, and
    explicit no-commit overrides.
  - desired behavior: Verified local work produces an auditable atomic history by
    default without granting push or history-rewrite authority.
  - acceptance criteria: Workflow tests cover default commit, opt-out, dirty
    worktree preservation, message/diff alignment, separate push authority, and
    tests/docs in the same slice.
  - verification: Focused contract tests, `go test ./internal/cli`, lint, and an
    independent agent review of cached-diff/commit-plan alignment without pushing.
  - evidence: Canonical and five-target policy tests, mixed-hunk/index safety
    contract checks, full Go suite, lint, `git diff --check`, and independent
    review with no findings.
  - docs: Commit authority, atomicity checklist, and override examples.
  - commit: `feat: commit verified slices by default`

- [x] **Orchestrate specialist subagents autonomously** (test-type: both)
  - blocked-by: Make feature delivery autonomous
  - summary: Permit proactive specialist exploration, implementation, testing,
    security, performance, documentation, review, and audit with isolated edit
    ownership and primary-owned Git integration.
  - desired behavior: Independent work runs in parallel when safe; findings are
    repaired and re-reviewed without routing routine decisions to the user.
  - acceptance criteria: Agent contracts cover ownership, supported delegation
    fallback, review waves, three-cycle blocker escalation, no concurrent Git,
    and preservation of unrelated work.
  - verification: Agent contract tests, target materialization, lint, and
    `go test ./internal/cli`.
  - docs: Delegation matrix, ownership rules, review loop, and escalation model.
  - commit: `feat: orchestrate specialist subagents`

- [x] **Enforce contextual authority boundaries** (test-type: both)
  - blocked-by: Make feature delivery autonomous, Commit verified slices by default
  - summary: Allow catalogued dependencies and disposable-environment operations;
    gate new uncatalogued direct dependencies and production mutations; deny
    dangerous out-of-scope commands instead of prompting repeatedly.
  - desired behavior: The harness remains autonomous locally while provenance,
    secrets, destructive commands, push, and production stay protected.
  - acceptance criteria: Policy and hook tests distinguish catalogued/new direct
    dependencies, local/test/ambiguous/production targets, safe alternatives,
    hard denials, commit, push, and publication authority.
  - verification: Focused guardrail tests, ShellCheck, `go test ./internal/cli`,
    lint, and generated Copilot/Cursor hook smokes.
  - evidence: Full Go suite, build, catalog/lint checks, ShellCheck, executable
    Copilot/Cursor hook variants, fail-closed parsing tests, and independent
    re-review with no findings.
  - docs: Dependency provenance and environment authority matrix.
  - commit: `feat: enforce contextual agent authority`

- [x] **Verify and document the autonomous rollout** (test-type: integration)
  - blocked-by: Run bugfixes as an autonomous evidence loop, Commit verified
    slices by default, Orchestrate specialist subagents autonomously, Enforce
    contextual authority boundaries
  - summary: Run the complete verification and materialization matrix, reconcile
    workflow artifacts, document migration, and close the spec without changing
    release workflows or pushing.
  - desired behavior: Every supported target receives one consistent autonomous
    workflow with fresh evidence and an auditable local history.
  - acceptance criteria: All commands in `plan.md` pass; the complete gate
    inventory has no unclassified conversational gate; every applicable pack and
    target is covered; spec, plan, tasks, index, README, and guardrails docs agree;
    no stale `[~]` remains; release follow-up is recorded separately.
  - verification: Full plan matrix, `git diff --check`, `git status --short`, and
    `git log --oneline` atomic-history review.
  - evidence: All 44 resolved packs materialized as 614 files across Codex,
    Copilot, Cursor, Claude Code, and OpenCode; full Go suite, build, lint,
    doctor, context, ShellCheck, hook smokes, atomic-history review, and final
    independent audit passed with no findings.
  - docs: Migration note, final evidence, and deferred release follow-up.
  - commit: `docs: complete autonomous workflow rollout`
