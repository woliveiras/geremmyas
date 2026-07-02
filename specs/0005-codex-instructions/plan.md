# Plan: Codex global instruction distribution

Spec: [spec.md](./spec.md)

Status: Implemented

## Approach

Correct the Codex global `AGENTS.md` path to `$HOME/.codex/AGENTS.md`, always
copy the catalog instructions to `~/.copilot/instructions/`, mirror them into a
Codex-owned `~/.codex/instructions/` for the `codex` target, and extend the
Codex generator to emit an on-demand Instructions index (each `applyTo` glob
plus a pointer to the Codex-owned path). Inlining full instruction bodies
(Option B) was rejected to keep `AGENTS.md` lean and preserve per-file-type
scoping.

## Touch points

- `internal/cli/global_paths.go` — change `codexAgents` from `.config/codex` to
  `.codex`; add the Codex instructions dir constant.
- `internal/cli/generate_claude.go` — `generateCodexAt` global `relPath`
  `.config/codex/AGENTS.md` → `.codex/AGENTS.md`; extend `buildIDEAgentsDoc`
  (or a Codex-specific step) to render the Instructions section from
  `artifacts.instructions`, resolving each file's `applyTo` and pointer path by
  scope.
- `internal/cli/generate.go` — for the global Codex target, mirror instruction
  files into `$HOME/.codex/instructions/`.
- `internal/cli/targets.go` — `globalCopyFlags` sets `instructions = true`
  unconditionally so `~/.copilot/instructions/` is always populated.
- `internal/cli/frontmatter.go` — reuse `parseMarkdownFrontmatter` to read
  `applyTo` from each instruction file (no change expected).
- `internal/cli/generate_test.go`, `global_test.go`, `targets` tests — cover the
  new path, the Instructions section, and the global install layout.
- `docs/architecture.md` (multi-IDE targets), `docs/creating-packs.md`,
  `README.md` — document Codex instruction delivery and the corrected path.

## Sequencing

1. Fix the global Codex `AGENTS.md` path in `global_paths.go` and
   `generateCodexAt`, with a failing unit test asserting `.codex/AGENTS.md`
   (red).
2. Add the Instructions section to the Codex generator: list each instruction's
   `applyTo` and the scope-correct pointer (`~/.codex/instructions/<file>`
   global, `.github/instructions/<file>` project). Unit-test the rendered
   content and the absence of any `~/.copilot` reference.
3. Mirror instruction files into `$HOME/.codex/instructions/` for the global
   Codex target. Integration-test the install layout in a temp `$HOME`.
4. Make `globalCopyFlags` copy instructions unconditionally; assert
   `~/.copilot/instructions/` is populated regardless of target.
5. Handle the missing-`applyTo` edge case and the no-instructions pack case.
6. Update docs.
7. Run focused tests, the full suite, `geremmyas lint`, and `geremmyas doctor`.
8. Note the stale `~/.config/codex/AGENTS.md` for manual cleanup on the
   maintainer machine (not produced by geremmyas anymore).

## Dependencies

- No external dependencies.
- Builds on spec 0001 (Codex generation target) and the existing global install
  and generator paths.

## Risks

- Over-scoping the Instructions section could bloat `AGENTS.md`; keep it a
  compact `applyTo` + pointer list, not inlined bodies.
- Path change orphans the old `~/.config/codex/AGENTS.md`; call it out so the
  maintainer removes it once.
- Codex following on-demand pointers depends on model behavior; the same
  mechanism already works for skills, which lowers this risk.
