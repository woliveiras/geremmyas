# Plan: Skill catalog tiering

Spec: [spec.md](./spec.md)

Status: Draft

## Approach

Add a `tier` field to the pack schema in `catalog/packs.json`, classify every
existing pack, add `paper-review` to the opt-in `research` stack pack, remove
caveman from geremmyas (recommend the upstream installer in the README), and
teach `lint` and the install/sync paths to respect tiers. Documentation is
updated to describe the tier model and the reduced default install set.

A lazy router (suppressing host injection of installed skills) was considered
and rejected: the host scans the whole skills root and injects every installed
skill, so the only lever geremmyas owns is which skills get installed.

## Touch points

- `catalog/packs.json` — add `tier` to every pack; add `paper-review` to the
  `research` pack.
- `project/.github/skills/paper-review/` — vendor the paper-review skill
  (source: `~/.agents/skills/paper-review/`).
- `internal/cli/catalog.go` — parse and expose `tier`.
- `internal/cli/lint.go` — validate presence and value of `tier`.
- `internal/cli/sync.go` / install path — reject `personal` for project scope;
  filter by tier on install.
- `internal/cli/catalog_test.go`, `lint_test.go`, `sync_test.go` — cover the new
  schema, validation, and install behavior.
- `project/AGENTS.md`, `project/.github/copilot-instructions.geremmyas.md`,
  `project/.github/skills/subagent-selection/SKILL.md` — drop caveman
  references; `.agents/skills/caveman*` and `.cursor/rules/caveman.mdc` — remove
  the maintainer's local caveman install.
- `docs/creating-packs.md`, `README.md`, `specs/README.md` — document tiers and
  recommend the upstream caveman installer.

## Sequencing

1. Add `tier` to the pack struct and parser with a failing test for a missing
   tier (red).
2. Classify every existing pack in `packs.json` as `core` / `stack` (`core`,
   `sdd` = `core`; everything else = `stack`).
3. Add lint validation for `tier` (presence + allowed values).
4. Add install/sync tier filtering and the personal-vs-project guard.
5. Vendor `paper-review` into `project/.github/skills/` and add it to the
   `research` pack.
6. Remove caveman from the repo and add the upstream recommendation to the
   README.
7. Update docs and `specs/README.md`.
8. Run focused tests, full suite, `geremmyas lint`, and `geremmyas doctor`.

## Dependencies

- No external dependencies.
- Builds on the existing `catalog/packs.json` schema and `lint`/`doctor`
  commands.

## Risks

- Misclassifying a pack changes what consumers receive; the classification table
  in the spec is reviewed before implementation.
- Breaking change to pack layout; documented in ADR 0001 and the README
  migration note.
- Removing caveman deletes the maintainer's local install; confirm scope before
  deleting `.agents/skills/caveman*` and `.cursor/rules/caveman.mdc`.

## Verification

- `go test ./internal/cli`
- `go test ./...`
- `./geremmyas lint`
- `./geremmyas doctor`
- Manual: `geremmyas init --packs core` in a temp dir installs only core skills.
