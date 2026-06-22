# Plan: Skill description validator (`geremmyas lint`)

Spec: [spec.md](./spec.md)

## Approach

Add a pure rule engine that takes a parsed skill (frontmatter + body) and returns
a list of violations, then wire a `lint` subcommand that walks the canonical
skills tree, runs the engine, prints a report, and sets the exit code. The rule
engine is unit-tested in isolation; the command is integration-tested over a
fixture skills tree.

## Touch points

- `internal/cli/lint.go` (new) — rule engine (`lintDescription`, `lintName`,
  `lintBody`) and `runLint` that scans skills, aggregates violations, prints the
  report, and returns an error when any skill fails.
- `internal/cli/cli.go` — add `case "lint"` to the command switch and a usage
  line; map a failing lint to a non-zero exit.
- Reuse the existing frontmatter parser (`parseMarkdownFrontmatter`) and the
  embedded/skill lookup helpers used by the generators.
- CI workflow under `.github/workflows/` — add a `geremmyas lint` step.

## Rule definitions

- description: non-empty; contains a "use when" phrase; contains a negative-scope
  phrase; length <= 1024; no `<`/`>`.
- name: equals the skill's directory base name.
- body: line count <= 500.
- structure: `SKILL.md` exists for each skill directory.

## Sequencing

1. Rule engine for description rules — unit slice.
2. Name-match and body-length rules — unit slice.
3. `lint` command scanning the canonical tree with report + exit code —
   integration slice over a fixture tree.
4. Missing-`SKILL.md` and boundary cases — unit/integration slice.
5. CI wiring — workflow step.

## Dependencies

- None external. Reuses the frontmatter parser and skill discovery helpers.

## Risks

- Negative-scope phrase matching could be too strict or too loose; keep the
  accepted phrase set small and documented (spec open question).
- Existing repository skills may currently fail the new rules; running the engine
  over the current tree during implementation surfaces and fixes those, but that
  cleanup is tracked separately if the volume is large.

## Verification

- `go test ./internal/cli/ -run Lint`
- `go build -o geremmyas ./cmd/geremmyas && ./geremmyas lint`
- CI: the lint step fails the build on any violation.
