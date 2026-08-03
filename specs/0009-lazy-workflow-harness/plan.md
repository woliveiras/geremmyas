# Plan: Lazy workflow harness

Spec: [spec.md](./spec.md)

Status: Ready

## Approach

Deliver the feature in independently verifiable slices. First expose global
state without mutation, then add safe clearing by reusing manifest
reconciliation. Migrate workflow names and packs only after global cleanup can
explain and remove owned legacy paths. Consolidate documentation and delegation
contracts, remove bundled agents, then measure the reduced materialized surface
in isolated homes.

## Sequencing

1. Add a shared global inventory model and `global list` human/JSON rendering.
2. Add dry-run planning and safe global clearing, including target filtering,
   adoptable legacy files, force boundaries, and manifest updates.
3. Rename the seven workflow skills and introduce `coding`, `quality`, and
   `base`, with tests for dependencies, budgets, destinations, and legacy
   reconciliation.
4. Consolidate documentation into the `docs` router and move review into the
   `verify` support tree.
5. Remove all bundled custom agents and migrate useful procedures into skills
   or non-discoverable references while keeping adapter support generic.
6. Tighten AGENTS routing, prompts, gates, diagnostics, and documentation around
   one-skill-at-entry and phase transitions.
7. Run isolated context baselines and the full five-target verification matrix.

## Main Touch Points

- `internal/cli/cli.go`, `global.go`, `global_manifest.go`, `context.go`, and
  focused tests for inventory and clearing.
- `catalog/packs.json`, `content/AGENTS.md`, `content/skills/`,
  `catalog/workflow-gates.json`, and `content/prompts/` for workflow migration.
- `content/agents/`, artifact planning, native generators, and target tests for
  removing distributed agents without deleting generic support.
- `README.md`, `docs/architecture.md`, `docs/guardrails-framework.md`, and
  migration guidance for the public interface.

## Dependencies

- Spec 0006 global manifest, context budgets, and skill taxonomy.
- Spec 0008 autonomous workflow and contextual authority boundaries.
- Go standard library only.

## Risks

- Global clearing can delete user-level files. Hashes, ownership, target scope,
  path validation, symlink rejection, and dry-run equivalence must be tested at
  the final mutation boundary.
- Skill renames are a public workflow migration. Duplicated aliases would reduce
  compatibility risk but recreate discovery cost, so migration is documented
  and managed old paths are reconciled instead.
- A broad `docs` skill can trigger too often. Its description and body need
  explicit negative scope and exactly-one-reference routing.
- Removing agents can reduce enforcement on runtimes that used native tool
  restrictions. The primary contract and review reference must state the
  residual when dynamic isolation is unavailable.
- `base` can be mistaken for all catalog content. Documentation must state that
  it excludes technology packs and runtime plugins.

## Verification

- Focused red/green tests listed in each task.
- `go test ./internal/cli -count=1`
- `go test ./... -count=1`
- `go run ./cmd/geremmyas lint`
- `go run ./cmd/geremmyas doctor`
- Temporary-home `global list`, dry-run, clear, and target-scoped reconciliation.
- Temporary-project materialization for Codex, Copilot, Cursor, Claude Code, and
  OpenCode with `coding`, `quality`, and `base`.
- `git diff --check` and `git status --short` after every slice.

## Completion

Set the spec to `Verified` only after all tasks and acceptance criteria have
fresh evidence, old managed names reconcile safely, the no-agent target matrix
passes, context baselines are recorded, and independent review has no actionable
finding.
