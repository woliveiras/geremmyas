# Plan: Assistant-neutral canonical content

Spec: [spec.md](./spec.md)

Status: Approved

## Approach

Separate artifact meaning from assistant destinations in five large changes.
First introduce a typed catalog and destination planner while preserving current
output. Then move embedded sources into neutral and target-specific trees.
Next make project sync target-aware with conservative ownership reconciliation.
After that, correct global target routing. Finish by replacing maintainer
dogfooding links and updating the public documentation.

Each large change gets focused tests, relevant documentation, and its own
Conventional Commit. The specification and accepted architecture record are
committed separately before implementation.

## Sequencing

1. After approval, mark the spec Approved, record the accepted architecture,
   and commit the specification artifacts.
2. Add explicit artifact kinds and a deterministic target destination planner,
   initially preserving existing output behavior.
3. Move canonical embedded sources to `content/` and Copilot adapters to
   `targets/copilot/`; update catalog validation, lint, diagnostics, and embeds.
4. Make project sync materialize only selected targets and reconcile obsolete
   owned artifacts safely, including conservative legacy adoption.
5. Make global installation route instructions and generated content strictly
   by selected target.
6. Replace maintainer symlinks and terminology, keeping GitHub platform files
   under the root `.github/`.
7. Run the full release verification and reconcile spec, plan, tasks, index,
   documentation, and migration notes.

## Main Touch Points

- `catalog/packs.json`, `assets.go`, and `internal/cli/catalog.go` for typed
  artifact metadata and neutral embedded roots.
- New focused artifact planning and project manifest modules under
  `internal/cli/`.
- `internal/cli/sync.go`, `targets.go`, and `generate*.go` for target-aware
  project outputs.
- `internal/cli/global.go`, `global_manifest.go`, and `global_paths.go` for
  target-correct global routing.
- `internal/cli/context.go`, `lint.go`, and their tests for neutral source
  diagnostics.
- `content/`, `targets/copilot/`, root dogfooding links, README, architecture,
  pack authoring guidance, and maintainer instructions.

## Dependencies

- Builds on Codex target support from specs 0001 and 0005.
- Reuses the hash and ownership safety model introduced by spec 0006.
- Uses the existing Go standard library only.

## Risks

- Moving 88 catalog entries can create silent omissions. Catalog validation and
  clean-install integration tests must cover every pack entry.
- Target-aware cleanup can remove files. Removal is limited to unchanged,
  manifest-owned regular files; symlinks, modified files, and unowned files are
  preserved.
- Mixed-target installs can duplicate shared skills. Destination planning must
  deduplicate before filesystem writes and produce stable ordering.
- Existing consumers may rely on Copilot output paths. Those paths remain
  unchanged whenever `copilot` is selected.
- The maintainer repository needs both GitHub platform files and assistant
  dogfooding. Only assistant-content symlinks move; workflows, issue templates,
  pull request templates, and assets remain.

## Commit Boundaries

1. `docs: specify assistant-neutral content model`
2. `refactor: type assistant content artifacts`
3. `refactor: move canonical assistant content`
4. `feat: materialize project content by target`
5. `fix: route global content by target`
6. `docs: generalize assistant framework terminology`
7. `docs: complete assistant-neutral content rollout`

## Verification

- Focused unit tests for catalog kinds and destination planning.
- Focused project sync tests for Codex-only, Copilot-only, and mixed targets.
- Migration tests for unchanged, modified, unowned, corrupt-manifest, and
  symlink cases.
- Global tests proving Codex-only installs do not touch Copilot paths.
- `go test ./internal/cli`
- `go test ./...`
- `go build -o geremmyas ./cmd/geremmyas`
- `./geremmyas lint`
- `./geremmyas doctor`
- `./geremmyas context`
- Clean temporary project smoke tests for each supported target and a mixed
  Codex/Copilot configuration.
- `git diff --check`, `git status --short`, and scoped `git log` review.
