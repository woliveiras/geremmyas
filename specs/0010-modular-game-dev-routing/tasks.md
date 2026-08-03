# Tasks: Modular game-development routing

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Specify focused game pack topology** (test-type: unit)
  - blocked-by: none
  - summary: Lock focused membership, complete closure, compatibility, and
    deduplication in catalog tests.
  - verification: `go test ./internal/cli -run 'Game|CatalogComposite' -count=1`

- [x] **Add routing conformance corpus** (test-type: unit)
  - blocked-by: Specify focused game pack topology
  - summary: Cover every game skill with prompt signals and require a unique
    lexical description winner without claiming runtime equivalence.
  - verification: `go test ./internal/cli -run 'GameSkillRouting' -count=1`

- [x] **Implement focused packs and discriminating triggers** (test-type: unit)
  - blocked-by: previous tasks
  - summary: Split `game-dev`, retain the art pack alias, and narrow all eleven
    frontmatter descriptions.
  - verification: `go test ./internal/cli -run 'Game|Catalog|Skill' -count=1 && go run ./cmd/geremmyas lint`

- [x] **Document focused game installation** (test-type: unit)
  - blocked-by: Implement focused packs and discriminating triggers
  - summary: Update public pack and skill references without overstating
    routing enforcement.
  - verification: `go run ./cmd/geremmyas lint && git diff --check`

- [~] **Reconcile and test Rusted** (test-type: integration)
  - blocked-by: Geremmyas verification
  - summary: Select the routine focused packs, reconcile obsolete managed
    skills, strengthen project routing, and capture a Codex smoke.
  - verification: project context, manifest/file inventory, and saved smoke evidence

- [ ] **Complete verification and evidence** (test-type: both)
  - blocked-by: all previous tasks
  - summary: Run full tests, lint, doctor, diff checks, reconcile spec status,
    and report both repository worktrees.
  - verification: `go test ./... -count=1 && go run ./cmd/geremmyas lint && go run ./cmd/geremmyas doctor && git diff --check`
