# Plan: Codex generation target

Spec: [spec.md](./spec.md)

## Approach

Mirror the OpenCode target end to end. The Codex document is built with the
shared `buildIDEAgentsDoc` helper and written behind the `geremmyas:generated`
marker, for both project and global scope. The work is small and additive: a new
target constant, a generator function, dispatch wiring, global paths, and
surface text, each covered by tests.

## Touch points

- `internal/cli/targets.go` — add `TargetCodex` constant, register in
  `validTargets`, include in the validation error message.
- `internal/cli/generate_claude.go` (or a sibling generator file) — add
  `generateCodex` / `generateCodexAt` using `buildIDEAgentsDoc(scope, root,
  artifacts, "codex", "Codex AGENTS.md")`, with project path `.codex/AGENTS.md`
  and global path `~/.codex/AGENTS.md`.
- `internal/cli/generate.go` — dispatch `if hasTarget(targets, TargetCodex)`.
- `internal/cli/global_paths.go` — add the global Codex destination.
- `internal/cli/cli.go` — add `codex` to `--targets` help, the usage line, and
  the summary/destination listing.

## Sequencing

1. Target recognition (constant, validation, error message) — unit slice.
2. Project-scope generation (generator + dispatch) — integration slice.
3. Global-scope generation (global paths + dispatch) — integration slice.
4. Surface text (help, usage, summary destinations) — integration/CLI slice.
5. Idempotency and `--force` behavior verified through the shared marker path.

## Dependencies

- None external. Relies only on existing helpers (`buildIDEAgentsDoc`,
  `writeGeneratedFile`, `hasTarget`, global path resolution).

## Risks

- Global path assumption (`~/.codex/AGENTS.md`) may differ from the installed
  Codex layout; confirm at review (spec open question).
- Project-scope path choice (Option A `.codex/AGENTS.md`) is a design decision to
  confirm before implementation.

## Verification

- `go test ./...`
- `go build -o geremmyas ./cmd/geremmyas`
- Manual: `geremmyas init --targets codex` + `geremmyas sync` in a temp dir, then
  inspect the generated Codex document and re-run to confirm idempotency.
