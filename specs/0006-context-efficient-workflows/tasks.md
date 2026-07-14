# Tasks: Context-efficient agent workflows

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Specify context-efficient workflow distribution** (test-type: unit)
  - blocked-by: none
  - summary: Update the multi-assistant PRD and add spec 0006, plan, tasks, and
    index entry.
  - desired behavior: The approved scope and safety constraints are durable.
  - acceptance criteria: All four artifacts agree on desired-state semantics,
    budgets, taxonomy, commit boundaries, and out-of-scope plugin mutation.
  - verification: Manual artifact review and `git diff --check`.
  - commit: `docs: specify context-efficient agent workflows`

- [ ] **Reconcile managed global state** (test-type: both)
  - blocked-by: Specify context-efficient workflow distribution
  - summary: Add an atomic ownership manifest and make `global` reconcile the
    complete desired pack/target state.
  - desired behavior: Unchanged obsolete managed files are removed; modified or
    unowned files are preserved and reported.
  - acceptance criteria: Unit tests cover hash decisions and corrupt state;
    integration tests cover expansion, shrinking, modified files, generated
    target removal, first-run adoption, and empty managed directories.
  - verification: `go test ./internal/cli -run 'Global|Manifest|Reconcile'` and
    `go test ./internal/cli`.
  - docs: README global usage and architecture state/reconciliation sections.
  - commit: `feat: reconcile managed global installs`

- [ ] **Generate compact Codex context** (test-type: both)
  - blocked-by: Reconcile managed global state
  - summary: Use a Codex-specific document builder that omits the embedded
    project contract, native skills index, and unusable agent-role advertising.
  - desired behavior: Codex receives only a compact bootstrap and on-demand
    instruction pointers; other targets retain required indexes.
  - acceptance criteria: Generator tests prove the project contract and Skills
    section are absent from global Codex output, instructions remain reachable,
    and Claude/OpenCode behavior does not regress.
  - verification: `go test ./internal/cli -run 'Codex|Claude|OpenCode'` and
    `go test ./internal/cli`.
  - docs: Target capability table and Codex global bootstrap behavior.
  - commit: `feat: generate compact Codex context`

- [ ] **Consolidate the SDD skill catalog** (test-type: both)
  - blocked-by: Generate compact Codex context
  - summary: Keep at most 10 user-facing SDD skills, move internal guardrails and
    composition steps to references, create opt-in packs for maintainer/decision
    helpers, and replace the verbose guardrail block with phase-aware routing.
  - desired behavior: Users retain feature, bugfix, implementation, review,
    verification, ADR, glossary, and commit capabilities without policy skills
    competing for discovery.
  - acceptance criteria: Catalog tests assert the default skill set and count;
    agent/skill references resolve; no nested support file is named `SKILL.md`;
    approval, regression, preservation, and verification invariants remain.
  - verification: `go test ./internal/cli`, `go test ./...`, and
    `./geremmyas lint`.
  - docs: README skill catalog, guardrails framework, creating-packs taxonomy,
    and migration list for removed public skill names.
  - commit: `refactor: consolidate SDD workflow skills`

- [ ] **Report context usage** (test-type: both)
  - blocked-by: Reconcile managed global state, Consolidate the SDD skill catalog
  - summary: Add `geremmyas context` to inventory known skill roots, distinguish
    owned and external content, and estimate metadata/context size.
  - desired behavior: A user can identify top-level, nested, stale, modified,
    unowned, system, and plugin contributions without mutating them.
  - acceptance criteria: Deterministic report tests cover missing roots, symlinks,
    nested files, manifest ownership, external plugin roots, and token estimates.
  - verification: `go test ./internal/cli -run Context`, `go test ./internal/cli`,
    and a manual `./geremmyas context` run.
  - docs: CLI reference, interpretation guide, and ownership limitations.
  - commit: `feat: report agent context usage`

- [ ] **Enforce context budgets** (test-type: unit)
  - blocked-by: Consolidate the SDD skill catalog, Report context usage
  - summary: Add lint rules for nested skills, metadata/body limits, SDD count,
    and `AGENTS.md` size; move remaining oversized examples to references.
  - desired behavior: CI blocks catalog and contract growth beyond the approved
    budgets with actionable violations.
  - acceptance criteria: Each violation has a focused failing test and the
    canonical repository passes all new budgets.
  - verification: `go test ./internal/cli -run Lint`, `./geremmyas lint`, and
    `go test ./...`.
  - docs: Skill authoring limits and context-budget rationale.
  - commit: `feat: enforce context budgets`

- [ ] **Bound subagent workflows** (test-type: unit)
  - blocked-by: Consolidate the SDD skill catalog
  - summary: Update explorer, spec-writer, reviewer, and architect contracts with
    bounded scope and output; make architecture fan-out conditional.
  - desired behavior: Subagents isolate expensive exploration and return concise
    evidence without multiplying routine work.
  - acceptance criteria: Agent contract tests or catalog assertions verify the
    required scope/evidence/unknowns/output clauses and conditional fan-out.
  - verification: `go test ./internal/cli`, `./geremmyas lint`, and manual agent
    contract review.
  - docs: Agent routing and delegation guidance.
  - commit: `refactor: bound subagent workflows`

- [ ] **Complete release verification** (test-type: integration)
  - blocked-by: Reconcile managed global state, Generate compact Codex context,
    Consolidate the SDD skill catalog, Report context usage, Enforce context
    budgets, Bound subagent workflows
  - summary: Run the full test/build/lint/doctor/context matrix, reconcile spec
    status and task checkboxes, and document any migration caveats.
  - desired behavior: The feature is releasable with fresh evidence and a clean,
    explainable worktree.
  - acceptance criteria: All verification commands pass; spec/index are marked
    Implemented; no stale `[~]` task remains; commit history has one commit per
    improvement plus the specification commit.
  - verification: Commands listed in `plan.md`, followed by
    `git status --short` and `git log --oneline` review.
  - docs: Final migration note and changelog entry when required.
  - commit: `docs: complete context workflow rollout`
