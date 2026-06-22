# Tasks: Skill description validator (`geremmyas lint`)

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Description rule engine** (test-type: unit)
  - blocked-by: none
  - summary: Implement description checks (use-when trigger, negative scope,
    <=1024 chars, no `<`/`>`) returning structured violations.
  - desired behavior: Each rule passes on valid input and fails with a named
    violation on crafted invalid input.
  - acceptance criteria: Missing trigger, missing negative scope, >1024 chars,
    and angle-bracket markup each produce the matching violation; boundary 1024
    passes and 1025 fails.
  - verification: `go test ./internal/cli/ -run LintDescription`

- [~] **Name-match and body-length rules** (test-type: unit)
  - blocked-by: Description rule engine
  - summary: Add `name` == directory check and `SKILL.md` body <=500 lines check.
  - desired behavior: Name mismatch and body over 500 lines each produce the
    matching violation; 500 lines passes, 501 fails; missing `name` reports
    mismatch.
  - acceptance criteria: Crafted inputs map to the expected violations at the
    boundaries.
  - verification: `go test ./internal/cli/ -run 'LintName|LintBody'`

- [ ] **`lint` command over the canonical skills tree** (test-type: integration)
  - blocked-by: Name-match and body-length rules
  - summary: Add `runLint` and `case "lint"` in `cli.go`; scan
    `project/.github/skills/`, print a per-skill report, set exit code.
  - desired behavior: A tree with violations exits non-zero and lists each
    failing rule with the skill path; a clean tree exits zero with a success
    summary.
  - acceptance criteria: Fixture tree with a bad skill → non-zero + report;
    all-valid tree → zero + summary.
  - verification: `go test ./internal/cli/ -run RunLint`

- [ ] **Missing `SKILL.md` and empty description handling** (test-type: both)
  - blocked-by: `lint` command over the canonical skills tree
  - summary: Treat a skill directory without `SKILL.md` and an empty/missing
    description as violations without crashing.
  - desired behavior: Missing file → structure violation; empty description →
    trigger + negative-scope violations.
  - acceptance criteria: Both cases reported as violations; command still exits
    cleanly (non-zero) without panic.
  - verification: `go test ./internal/cli/ -run 'Lint|RunLint'`

- [ ] **Wire `geremmyas lint` into CI** (test-type: integration)
  - blocked-by: `lint` command over the canonical skills tree
  - summary: Add a lint step to the CI workflow so violations fail the build.
  - desired behavior: CI runs `geremmyas lint` and fails on any violation.
  - acceptance criteria: Workflow includes the lint step; local `./geremmyas
    lint` over the current tree passes (after fixing any surfaced skills).
  - verification: `./geremmyas lint` exits zero on the cleaned tree.
