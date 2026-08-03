# Tasks: Lazy workflow harness

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [ ] **Inventory global harness state** (test-type: both)
  - blocked-by: none
  - summary: Add a shared ownership-aware inventory and expose `global list` in
    human and JSON forms.
  - desired behavior: Users can distinguish managed, modified, missing,
    obsolete, adoptable, unowned, symlinked, and external state without writes.
  - acceptance criteria: Stable output includes manifest packs/targets/version,
    path/hash/state records, missing/corrupt behavior, and observed plugin roots.
  - verification: `go test ./internal/cli -run 'Global(List|Inventory)' -count=1`

- [ ] **Clear managed global state safely** (test-type: both)
  - blocked-by: Inventory global harness state
  - summary: Add dry-run and target-aware clear on top of manifest reconciliation.
  - desired behavior: Only eligible owned paths are removed; every preserved path
    has a reason and the manifest remains atomic and valid.
  - acceptance criteria: Tests cover dry-run byte identity, target subsets,
    modified/unowned/symlink paths, adoptable files, force boundaries, missing,
    corrupt, and incompatible manifests.
  - verification: `go test ./internal/cli -run 'Global(Clear|Manifest|Reconcile)' -count=1`

- [ ] **Migrate workflow names and packs** (test-type: both)
  - blocked-by: Clear managed global state safely
  - summary: Rename the seven public skills and replace SDD selection with
    `coding`, `quality`, and `base`.
  - desired behavior: `coding` exposes four active-development skills and `base`
    exposes seven technology-neutral skills with no aliases or agents.
  - acceptance criteria: Catalog, lint, project/global target outputs, gates,
    prompts, documentation, and old managed-path reconciliation use the new names.
  - verification: `go test ./internal/cli -run 'Catalog|Pack|Sync|Global|Workflow' -count=1 && go run ./cmd/geremmyas lint`

- [ ] **Consolidate lazy documentation workflows** (test-type: unit)
  - blocked-by: Migrate workflow names and packs
  - summary: Merge update-docs, glossary, ADR/MADR, and RFC guidance into `docs`
    with a compact router, references, and templates.
  - desired behavior: One public skill loads only the procedure and template for
    the requested document type; spec artifacts remain under `spec`.
  - acceptance criteria: Lint and policy tests enforce trigger scope, routing,
    resolved local links, template presence, and context budgets.
  - verification: `go test ./internal/cli -run 'Docs|Lint|Workflow' -count=1 && go run ./cmd/geremmyas lint`

- [ ] **Replace bundled agents with dynamic delegation contracts** (test-type: both)
  - blocked-by: Migrate workflow names and packs, Consolidate lazy documentation workflows
  - summary: Remove the ten canonical agent profiles and catalog selection,
    preserve reusable procedures in skills/references, and keep generic adapters.
  - desired behavior: No target receives bundled custom agents; review and
    specialist work use bounded runtime-created subagents when supported.
  - acceptance criteria: Five-target tests assert absence of generated agent
    files, adapter unit coverage remains, and stale managed agents reconcile
    conservatively.
  - verification: `go test ./internal/cli -run 'Agent|Workflow|Sync|Global' -count=1`

- [ ] **Enforce lazy routing and measure the reduced surface** (test-type: both)
  - blocked-by: Migrate workflow names and packs, Replace bundled agents with dynamic delegation contracts
  - summary: Tighten AGENTS/prompt routing, context diagnostics, budgets, and
    isolated baselines for no-state, global, coding, and base scenarios.
  - desired behavior: Simple work starts with zero workflow skills, ambiguity
    starts with at most refine, and closing capabilities stay phase-local.
  - acceptance criteria: Structural routing tests pass and temporary environments
    produce distinguishable deterministic context reports without real-home access.
  - verification: `go test ./internal/cli -run 'Context|Workflow|Budget' -count=1 && go run ./cmd/geremmyas context`

- [ ] **Complete full verification and migration documentation** (test-type: integration)
  - blocked-by: all previous tasks
  - summary: Reconcile public docs and feature artifacts, run every supported
    target in temporary roots, and close all acceptance criteria.
  - desired behavior: The new harness is safe to install, inspect, clear, and use
    with fresh evidence and no unexplained worktree changes.
  - acceptance criteria: Full Go tests, lint, doctor, temporary-home integration,
    target matrix, diff check, independent review, and status reconciliation pass.
  - verification: `go test ./... -count=1 && go run ./cmd/geremmyas lint && go run ./cmd/geremmyas doctor && git diff --check`
