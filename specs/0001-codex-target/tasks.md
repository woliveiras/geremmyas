# Tasks: Codex generation target

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Recognize `codex` as a valid target** (test-type: unit)
  - blocked-by: none
  - summary: Add `TargetCodex` constant, register in `validTargets`, include in
    the validation error message.
  - desired behavior: Target normalization/validation accepts `codex` and rejects
    unknown values while listing `codex` as valid.
  - acceptance criteria: Unknown target error lists `codex`; `codex` survives
    `normalizeTargets`.
  - verification: `go test ./internal/cli/ -run Target`

- [x] **Generate Codex document at project scope** (test-type: integration)
  - blocked-by: Recognize `codex` as a valid target
  - summary: Add `generateCodex`/`generateCodexAt` using `buildIDEAgentsDoc` and
    dispatch in `generate.go`; project path `.codex/AGENTS.md`.
  - desired behavior: `init`/`sync --targets codex` writes the Codex document with
    the `AGENTS.md` body plus the skill index, behind the generated marker.
  - acceptance criteria: File exists at the project Codex path; contains the
    AGENTS body and a "Skills (on demand)" section; re-running sync is idempotent.
  - verification: `go test ./internal/cli/ -run 'Sync.*Codex|Codex'`

- [ ] **Generate Codex document at global scope** (test-type: integration)
  - blocked-by: Generate Codex document at project scope
  - summary: Add the global Codex destination in `global_paths.go` and dispatch
    for global scope; global path `~/.codex/AGENTS.md`.
  - desired behavior: `global --targets codex` writes the Codex document under the
    global Codex location.
  - acceptance criteria: File exists at `~/.codex/AGENTS.md` (resolved test home);
    contains AGENTS body and skill index.
  - verification: `go test ./internal/cli/ -run 'Global.*Codex|Codex'`

- [ ] **Surface `codex` in CLI help, usage, and destination summary** (test-type: integration)
  - blocked-by: Generate Codex document at project scope
  - summary: Add `codex` to the `--targets` flag help, the usage line, and the
    summary/destination listing.
  - desired behavior: Help and summary output mention `codex` and its
    destination path.
  - acceptance criteria: Help/usage text includes `codex`; summary lists the
    Codex destination when the target is selected.
  - verification: `go test ./internal/cli/ -run 'Help|Summary|Doctor'`

- [ ] **Verify mixed-target and `--force` behavior** (test-type: integration)
  - blocked-by: Generate Codex document at project scope
  - summary: Confirm `--targets codex,copilot` produces both outputs and that
    `--force` overwrite semantics match other targets.
  - desired behavior: Codex and Copilot outputs coexist; customized Codex file is
    only overwritten with `--force`.
  - acceptance criteria: Both outputs present; customized file preserved without
    `--force`, overwritten with `--force`.
  - verification: `go test ./internal/cli/...`
