# Tasks: Assistant-neutral canonical content

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Approve and record the assistant-neutral content model** (test-type: unit)
  - blocked-by: none
  - summary: Review the PRD update, spec, plan, tasks, index row, and accepted
    architecture record before implementation.
  - desired behavior: Scope, compatibility, safety rules, and commit boundaries
    are durable and explicitly approved.
  - acceptance criteria: Artifacts agree on target behavior, migration safety,
    neutral paths, and out-of-scope historical/platform files.
  - verification: Manual artifact review and `git diff --check`.
  - commit: `docs: specify assistant-neutral content model`

- [x] **Type artifacts and plan destinations** (test-type: unit)
  - blocked-by: Approve and record the assistant-neutral content model
  - summary: Add explicit catalog artifact kinds and a deterministic planner for
    shared and target-native destinations.
  - desired behavior: Artifact meaning no longer depends on `.github/*`
    prefixes and mixed targets produce a stable deduplicated plan.
  - acceptance criteria: Tests cover every kind, every supported target,
    duplicates, unknown kinds, unsafe paths, and stable ordering.
  - verification: `go test ./internal/cli -run 'Artifact|Destination|Catalog'`
    and `go test ./internal/cli`.
  - commit: `refactor: type assistant content artifacts`

- [x] **Move canonical assistant content** (test-type: both)
  - blocked-by: Type artifacts and plan destinations
  - summary: Move shared sources to `content/`, Copilot adapters to
    `targets/copilot/`, and update embeds, catalog, lint, diagnostics, and tests.
  - desired behavior: No shared source uses `.github` as taxonomy and all pack
    sources remain valid.
  - acceptance criteria: Catalog validation accounts for every moved entry;
    lint and context use neutral roots; no orphaned canonical content remains.
  - verification: `go test ./internal/cli`, `./geremmyas lint`, and
    `./geremmyas doctor`.
  - commit: `refactor: move canonical assistant content`

- [x] **Materialize and reconcile project targets** (test-type: both)
  - blocked-by: Move canonical assistant content
  - summary: Make project sync write only shared plus selected target outputs and
    add hash-based project ownership reconciliation with legacy adoption.
  - desired behavior: Clean Codex-only projects contain no Copilot-only files,
    target removal cleans unchanged owned files, and user content survives.
  - acceptance criteria: Integration tests cover Codex-only, Copilot-only,
    mixed targets, target removal, modified files, unowned files, symlinks,
    missing/corrupt manifests, and pre-manifest installs.
  - verification: `go test ./internal/cli -run 'Sync|Project|Manifest|Codex|Copilot'`
    and `go test ./internal/cli`.
  - commit: `feat: materialize project content by target`

- [x] **Route global content by selected target** (test-type: both)
  - blocked-by: Type artifacts and plan destinations
  - summary: Replace unconditional Copilot instruction copying with target-aware
    global destination planning and reconciliation.
  - desired behavior: Codex-only and Cursor-only global installs do not create
    Copilot instructions; mixed installs still produce the complete union.
  - acceptance criteria: Tests cover every target alone, mixed targets,
    shrinking desired state, modified outputs, and manifest ownership.
  - verification: `go test ./internal/cli -run 'Global|Target|Codex|Copilot'`
    and `go test ./internal/cli`.
  - commit: `fix: route global content by target`

- [x] **Generalize maintainer surfaces and documentation** (test-type: integration)
  - blocked-by: Move canonical assistant content, Materialize and reconcile
    project targets, Route global content by selected target
  - summary: Replace assistant-content dogfooding symlinks, update help and
    documentation terminology, and preserve GitHub platform files.
  - desired behavior: Contributors see a neutral source model and can identify
    each target adapter without mistaking `.github` for the canonical tree.
  - acceptance criteria: README, architecture, pack authoring, context output,
    CLI help, and maintainer instructions agree; workflows and repository
    templates remain under `.github`.
  - verification: Documentation link/path checks, CLI help snapshot tests,
    `git diff --check`, and `git status --short`.
  - commit: `docs: generalize assistant framework terminology`

- [x] **Complete release verification** (test-type: integration)
  - blocked-by: Materialize and reconcile project targets, Route global content
    by selected target, Generalize maintainer surfaces and documentation
  - summary: Run the complete verification matrix and reconcile all durable
    progress artifacts.
  - desired behavior: The migration is releasable with fresh evidence and one
    scoped commit per large change.
  - acceptance criteria: All checks pass; no stale `[~]` remains; spec and index
    are Implemented; plan records completion evidence; worktree is explainable.
  - verification: All commands from `plan.md`, followed by
    `git status --short` and `git log --oneline`.
  - commit: `docs: complete assistant-neutral content rollout`
