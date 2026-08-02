---
spec: "0008"
title: Autonomous agent workflows
family: multi-assistant
phase: 6
status: Approved
owner: ""
depends_on: ["0006", "0007"]
origin: Workflow-gate audit and user autonomy policy from 2026-08-02
---

# Spec: Autonomous agent workflows

## Context & Motivation

Geremmyas currently assumes that a human follows feature and bugfix work
step-by-step. The contract, skills, agents, prompts, templates, and guardrail
documentation repeat approval gates for specifications, fixes, architecture,
staging, commit messages, and infrastructure commands. The original product PRD
also limits subagents to read-only investigation and review.

The user now trusts the engineering harness and wants local work to be autonomous.
Agents should keep specifications, plans, tasks, bugfix records, tests, reviews,
and commit history accurate enough for later audit instead of asking for routine
approval. Local commits are enabled by default. Push and every production
mutation or publication remain explicit authority boundaries.

## Requirements

### Functional

- [ ] Make autonomy the default in the project contract. Agents may specify,
      implement, test, review, document, and locally commit in-scope work without
      conversational approval gates.
- [ ] Keep `spec.md`, `plan.md`, and `tasks.md` mandatory for features, but allow
      implementation to begin when the work package is machine-ready: acceptance
      criteria are testable, contracts and dependencies are known, verification
      commands exist, and no material product decision is unresolved.
- [ ] Replace human-centered workflow states with `Draft`, `Ready`, `In Progress`,
      and `Verified`. Commit, push, merge, and release evidence are recorded
      separately and do not redefine technical verification. Existing historical
      spec statuses remain valid without bulk migration.
- [ ] Enforce lifecycle transitions consistently: `Draft` means artifacts or
      material decisions are incomplete; `Ready` requires testable acceptance
      criteria, concrete contracts, known dependencies, verification commands,
      and no unresolved material decision; `In Progress` begins when the first
      implementation task is marked `[~]`; `Verified` requires every acceptance
      criterion evidenced, every finished task `[x]`, no stale `[~]`, required
      docs reconciled, and independent review with no actionable finding. Failed
      evidence keeps or returns the spec to `In Progress`; a newly unresolved
      material decision returns it to `Draft`.
- [ ] Let agents update and revalidate work artifacts automatically when the
      approach changes inside the authorized objective. Escalate only when the
      objective materially changes or requirements become contradictory.
- [ ] Define material-risk escalation from recorded blast radius, reversibility,
      harness coverage, and rollback evidence. Always include incompatible public
      contracts, irreversible data or state migration, authentication or
      authorization boundary redesign, and changes to core runtime, toolchain,
      build, or deployment architecture. Escalate only when blast radius is high
      and rollback is unproven or critical harness evidence is missing, so normal
      in-scope changes do not become conversational gates.
- [ ] Remove the mandatory commit-permission question from requirements and SDD
      entrypoints. Explicit instructions such as read-only, plan-only, or no
      commits override the default for that session.
- [ ] Make bugfixes autonomous while preserving the bugfix document,
      reproduction, ranked hypotheses, regression test that fails before the fix,
      verified correction, actual cause, cleanup, and nearby-suite evidence.
- [ ] Preserve red, green, review, documentation, and completion gates as
      executable evidence gates. They must not require routine human approval.
- [ ] Create atomic local commits by default. Each commit contains one coherent
      task-owned slice, its tests, and required documentation; uses Conventional
      Commits; derives its message from the staged diff; and excludes unrelated
      user changes.
- [ ] Keep push, history rewriting, and production release or publication outside
      local commit authority. No workflow may infer push from commit permission.
- [ ] Permit proactive specialist subagents for specification, architecture,
      implementation, tests, security, performance, documentation, review, and
      audit. Parallel edits require explicit file, module, or worktree ownership;
      the primary agent owns integration and Git operations.
- [ ] Run independent review after a slice or verification wave, repair findings,
      and re-review automatically. Escalate after three consecutive review or
      debugging cycles fail to converge on the same blocker.
- [ ] Infer tools and patterns from the repository and installed catalog. Existing
      project dependencies and catalogued capabilities may be used autonomously.
      A new uncatalogued direct dependency requires explicit user choice after a
      provenance, maintenance, security, license, and build-versus-buy assessment.
- [ ] Allow mutations in local, disposable, and test environments when the target
      is verified and rollback or recreation is available. Require explicit user
      authorization for every production mutation, deploy, release, publication,
      or policy change.
- [ ] Replace unnecessary command confirmations with contextual allowance or a
      safe alternative. Continue to deny force-push, broad destructive deletion,
      secret exposure, and dangerous operations outside the authorized scope.
- [ ] Treat human play, accessibility, audio, visual, and feel judgments as
      residual evidence when automation cannot establish the claim. Their absence
      does not block other autonomous work or erase the documented limitation.
- [ ] Maintain a durable migration inventory of every conversational gate in
      `content/AGENTS.md`, all skill and reference trees, all agent and reference
      trees, prompts, instructions, guardrails, templates, target adapters,
      catalogued packs, README, and workflow documentation. Classify each gate as
      removed, retained authority boundary, residual evidence, or deferred release
      work. Add an automated scan that fails on any unclassified gate.
- [ ] Update public documentation and every materialized assistant target so the
      autonomy, commit, delegation, dependency, escalation, and production
      boundaries are consistent.

### Non-Functional

- [ ] Add no new third-party dependencies.
- [ ] Preserve user-authored, modified, unowned, ignored, secret, and unrelated
      files throughout autonomous work and commit selection.
- [ ] Keep existing catalog and contract size budgets enforced by `geremmyas lint`.
- [ ] Keep canonical workflow content assistant-neutral under `content/`; target
      adapters may translate capabilities but not weaken authority boundaries.
- [ ] Deliver each implementation slice as a separate Conventional Commit with
      its tests, documentation, and fresh verification evidence.
- [ ] Do not push, deploy, publish, mutate production, or change GitHub release
      workflows during this specification.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Contract, skill, agent, prompt, template, or guardrail invariant can be inspected without target materialization | Default commit policy; no human spec gate; production authority wording; subagent ownership rules |
| **integration** | A pack must materialize consistent behavior for one or more assistant targets | Temporary project installs for Codex, Copilot, Cursor, Claude Code, and OpenCode |
| **both** | Policy changes cross canonical content, catalog routing, generated contracts, and hooks | Autonomous SDD lifecycle and contextual command guardrails |

Default: unit assertions for canonical policy and focused integration tests for
pack resolution and generated target output. Run `geremmyas lint`, the full Go
suite, build, doctor, context diagnostics, ShellCheck, and `git diff --check` at
the final verification wave.

## Acceptance Criteria

- [ ] Given an unambiguous feature request, when its spec, plan, and tasks become
      machine-ready, then implementation begins without asking the user to approve
      the artifacts.
- [ ] Given a reproduced bug, when the regression test fails for the expected
      reason, then the agent applies and verifies the fix without pausing for a
      bugfix-proposal approval.
- [ ] Given an approach change inside the original objective, when artifacts and
      tests are updated, then the workflow revalidates and continues without
      reapproval.
- [ ] Given a contradictory requirement or material-risk trigger, when repository
      evidence and specialist review cannot resolve it, then the agent records the
      evidence and asks one focused human decision before continuing.
- [ ] Given a proposed high-blast-radius change to runtime, toolchain, build,
      deployment architecture, public contracts, data, or security boundaries,
      when rollback is unproven or critical harness evidence is missing, then the
      agent escalates once with the risk record; otherwise it continues through
      the autonomous evidence loop.
- [ ] Given a session without a no-commit instruction, when a verified slice is
      complete, then the agent creates an atomic Conventional Commit without
      asking for file or message approval.
- [ ] Given an explicit no-commit, read-only, or plan-only instruction, when work
      reaches the commit boundary, then no commit is created and the changed or
      proposed files are reported.
- [ ] Given a dirty worktree, when the agent commits a completed slice, then only
      task-owned files or hunks are staged and unrelated changes remain intact.
- [ ] Given independent specialist work, when it can be partitioned safely, then
      subagents may edit in parallel and the primary agent integrates their work,
      runs verification, and owns commits.
- [ ] Given review findings, when they are actionable and in scope, then the agent
      repairs and re-reviews them without human intervention; the same unresolved
      blocker across three cycles is escalated with evidence.
- [ ] Given an existing or catalogued dependency, when the task needs it, then the
      agent may use it autonomously. Given a new uncatalogued direct dependency,
      the agent stops before installation and presents provenance plus a
      build-versus-buy recommendation.
- [ ] Given a verified local or disposable test target, when an in-scope mutation
      is needed, then it proceeds autonomously. Given a production target, no
      mutation, deploy, release, publication, or policy change occurs without
      explicit authorization.
- [ ] Given a request to commit, when no push was explicitly requested, then no
      push, force-push, amend, rebase, release, or publication occurs.
- [ ] Given a subjective acceptance claim that automation cannot establish, when
      all automated evidence is complete, then the workflow records the residual
      human validation without blocking unrelated completion.
- [ ] Given every canonical workflow surface and every catalogued pack, when the
      gate inventory scan runs, then every conversational approval or confirmation
      is classified and no obsolete feature, bugfix, commit-message, architecture,
      tool-selection, or non-production gate remains unclassified.
- [ ] Given all applicable packs materialized for every supported target, when the
      generated artifacts are inspected, then they express the same autonomous
      lifecycle and authority boundaries without broken references.

## Edge Cases

- The user explicitly requests exploration only, planning only, no edits, or no
  commits: that narrower instruction overrides autonomous defaults.
- Production status is ambiguous: treat the target as protected until current
  configuration proves it is local, disposable, or test-only.
- A dirty worktree overlaps the same file or hunk: isolate the task in a worktree,
  serialize the edit, or escalate only when ownership cannot be resolved safely.
- A target lacks native subagents: keep the same specialist-review contract
  inline or through the target's supported delegation mechanism.
- The harness cannot execute a required behavior: continue all safe work, record
  the missing evidence, and do not claim that boundary is verified.
- A direct dependency already exists but needs a compatible update: treat it as
  catalogued, inspect changelog and compatibility, and use the normal harness.
- Package-manager-selected transitive dependencies are recorded and audited by
  the lockfile; only a suspicious provenance or security result reopens the gate.
- Specialist agents disagree: compare evidence against the spec and tests, use a
  third specialist when useful, and escalate only if the material conflict remains.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Default operating mode | Autonomous local delivery | Auditability and executable evidence replace routine conversational approval |
| Feature readiness | Machine-ready work package | Testable criteria and known contracts are stronger progress gates than user acknowledgement |
| Local commits | Enabled by default, explicit opt-out | Atomic history is part of the audit trail, not a separate manual ceremony |
| Commit integration | Primary agent only | Avoids index and history races while specialists work in parallel |
| Delegation | Proactive specialists with isolated ownership | Adds depth and parallelism without shared-state collisions |
| Review loop | Repair and re-review, escalate after three repeated blocker cycles | Prevents both premature escalation and endless agent loops |
| Dependency authority | Existing/catalogued autonomous; new direct dependency gated | Preserves provenance and build-versus-buy control |
| Environment authority | Local/test autonomous; production explicit | Keeps the harness fast without granting production authority |
| Dangerous commands | Deny or use a safe alternative | Safety failures should not become repetitive confirmation prompts |
| Subjective validation | Residual evidence | Human judgment remains honest without blocking automatable delivery |

## Out of Scope

- Changing `.github/workflows/release.yml`, release environment protection, or
  automatic release-asset publication. This is the next rollout.
- Push, production deployment, publication, or changes to live infrastructure.
- Bulk migration of historical spec status values.
- Adding third-party dependencies or custom subagent runtimes.
- Removing verification, regression testing, documentation, preservation, or
  secret-safety requirements.
