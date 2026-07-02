# Tasks: Codex global instruction distribution

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Fix Codex global AGENTS.md path** (test-type: unit)
  - blocked-by: none
  - summary: Write the global Codex `AGENTS.md` to `$HOME/.codex/AGENTS.md`
    instead of `$HOME/.config/codex/AGENTS.md`.
  - desired behavior: `global --targets codex` targets the Codex home Codex
    actually reads.
  - acceptance criteria: Unit test asserts the global Codex relPath is
    `.codex/AGENTS.md`; `global_paths.codexAgents` uses `.codex`.
  - verification: `go test ./internal/cli`

- [x] **Instructions index in Codex doc** (test-type: unit)
  - blocked-by: Fix Codex global AGENTS.md path
  - summary: Render an Instructions section in the generated Codex `AGENTS.md`
    listing each pack instruction with its `applyTo` glob and a scope-correct
    pointer.
  - desired behavior: Global scope points to `~/.codex/instructions/<file>`;
    project scope points to `.github/instructions/<file>`; no `~/.copilot`
    reference appears.
  - acceptance criteria: Unit test asserts the section lists `applyTo` + pointer
    and contains no `~/.copilot` string.
  - verification: `go test ./internal/cli`

- [x] **Mirror instructions into Codex home** (test-type: integration)
  - blocked-by: Instructions index in Codex doc
  - summary: For the global `codex` target, copy every pack instruction file to
    `$HOME/.codex/instructions/`.
  - desired behavior: Codex can read each referenced instruction on demand from
    its own directory.
  - acceptance criteria: Integration test in a temp `$HOME` asserts every pack
    instruction exists under `.codex/instructions/`.
  - verification: `go test ./internal/cli`

- [x] **Copy instructions unconditionally** (test-type: unit)
  - blocked-by: Fix Codex global AGENTS.md path
  - summary: `globalCopyFlags` sets `instructions = true` regardless of target
    so `~/.copilot/instructions/` is always populated.
  - desired behavior: Instructions land in both Copilot and Codex locations
    independent of the selected assistant.
  - acceptance criteria: Unit test asserts `instructions` is true for a
    codex-only target; integration test asserts `.copilot/instructions/` is
    populated.
  - verification: `go test ./internal/cli`

- [x] **Edge cases** (test-type: unit)
  - blocked-by: Instructions index in Codex doc
  - summary: Handle an instruction without `applyTo` (unscoped label) and a pack
    with no instructions (omit the section).
  - desired behavior: No crash or empty/dangling section; unscoped instructions
    still listed.
  - acceptance criteria: Unit tests cover the missing-`applyTo` and
    no-instructions cases.
  - verification: `go test ./internal/cli`

- [x] **Documentation** (test-type: unit)
  - blocked-by: Mirror instructions into Codex home
  - summary: Document Codex instruction delivery and the corrected path in
    `docs/architecture.md`, `docs/creating-packs.md`, and `README.md`; note the
    stale `~/.config/codex/AGENTS.md` cleanup.
  - desired behavior: Readers understand how instructions reach Codex and where
    the global doc lives.
  - acceptance criteria: Docs describe the on-demand model, the `~/.codex`
    paths, and the one-time cleanup note.
  - verification: Manual doc review; `./geremmyas doctor`
