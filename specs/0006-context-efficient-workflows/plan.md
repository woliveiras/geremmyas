# Plan: Context-efficient agent workflows

Spec: [spec.md](./spec.md)

Status: Implemented

## Approach

Deliver the work as six independently verifiable improvements. First make global
state ownership-aware and declarative. Then specialize Codex generation so it
does not duplicate native context. Consolidate the SDD catalog using a strict
skill/reference/agent taxonomy. Add diagnostics and budgets to prevent future
growth. Finally tighten custom-agent delegation contracts.

Every implementation task includes its tests and documentation and ends in one
Conventional Commit. The specification artifacts are committed separately after
approval.

## Sequencing

1. Add a versioned global ownership manifest and safe desired-state
   reconciliation around copies, instruction mirrors, and generated files.
2. Split Codex document rendering from targets that require embedded contracts;
   keep only the compact bootstrap and instruction pointers for Codex.
3. Consolidate the SDD catalog, relocate internal procedures to references, and
   shorten `AGENTS.md` routing without changing approval or verification policy.
4. Add the `context` command and source/ownership-aware measurements.
5. Tighten lint budgets and refactor remaining oversized skills into references.
6. Bound custom-agent inputs/outputs and make architecture fan-out conditional.
7. Run focused tests after each slice and the full release verification after
   the final slice.

## Main Touch Points

- `internal/cli/global.go`, `global_paths.go`, `cli.go`, `generate*.go` for
  ownership, reconciliation, and target-aware output.
- New focused manifest/context modules under `internal/cli/` rather than adding
  state logic to command dispatch.
- `catalog/packs.json`, `project/AGENTS.md`, `.github/skills/`, and
  `.github/agents/` for catalog and workflow consolidation.
- `internal/cli/lint.go` and catalog/lint tests for budgets.
- `README.md`, `docs/architecture.md`, `docs/creating-packs.md`, and
  `docs/guardrails-framework.md` for public behavior and maintainer guidance.

## Dependencies

- Builds on spec 0004 tiering and spec 0005 Codex global instruction delivery.
- Uses only the Go standard library (`crypto/sha256`, `encoding/json`, atomic
  rename in the state directory).

## Risks

- Pruning global files has destructive potential. Deletion is restricted to
  manifest-owned paths whose current hash equals the installed hash.
- Consolidating public skill names can break explicit invocations. Keep the
  remaining user-facing capabilities stable and document removed policy skills.
- A very small `AGENTS.md` can omit important invariants. Budgets apply after the
  mandatory approval, bugfix, commit, preservation, and verification rules are
  retained.
- Plugin metadata remains a large external source of context. Diagnostics expose
  it, but Geremmyas does not mutate it.

## Verification

- `go test ./internal/cli`
- `go test ./...`
- `go build -o geremmyas ./cmd/geremmyas`
- `./geremmyas lint`
- `./geremmyas doctor`
- `./geremmyas context`
- Temporary-home integration: expand and shrink a global pack set, modify one
  managed file, add one unowned skill, and verify removal/preservation/reporting.
- `git status --short` and one scoped commit per completed task.

## Completion

Completed on 2026-07-14. The full test, build, lint, doctor, and context matrix
passed. A temporary-home integration expanded and reduced the desired global
state, removed unchanged obsolete files, preserved a modified file, and
reported an unowned skill. The implementation was delivered in one scoped
commit per improvement.
