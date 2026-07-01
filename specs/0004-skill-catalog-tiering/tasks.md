# Tasks: Skill catalog tiering

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Tier field in pack schema** (test-type: unit)
  - blocked-by: none
  - summary: Add a `tier` field (`core` | `stack`) to the pack struct and JSON
    parser.
  - desired behavior: Packs expose a typed `tier`; parsing a pack without a
    tier is detectable.
  - acceptance criteria: Unit test parses a pack with each valid tier and a
    parser/validation path flags a missing tier.
  - verification: `go test ./internal/cli`

- [x] **Lint validates tier** (test-type: unit)
  - blocked-by: Tier field in pack schema
  - summary: `geremmyas lint` fails when a pack omits `tier` or uses an invalid
    value.
  - desired behavior: Lint exits non-zero naming the offending pack.
  - acceptance criteria: Lint test covers missing tier and invalid value.
  - verification: `go test ./internal/cli`

- [x] **Classify existing packs** (test-type: unit)
  - blocked-by: Lint validates tier
  - summary: Assign `core` / `stack` to every pack in `catalog/packs.json`
    (`core`, `sdd` = `core`; everything else = `stack`).
  - desired behavior: Catalog passes lint; the default set is the core tier
    only.
  - acceptance criteria: Catalog test asserts the core set and that no pack is
    untiered.
  - verification: `go test ./internal/cli && ./geremmyas lint`

- [x] **Tier-aware install** (test-type: integration)
  - blocked-by: Classify existing packs
  - summary: Install/sync respects tier so a core-only install excludes stack
    packs; the interactive picker labels each pack's tier.
  - desired behavior: Core install excludes stack skills.
  - acceptance criteria: Integration test installs core-only and asserts stack
    skills are absent.
  - verification: `go test ./internal/cli`

- [x] **Add paper-review to research** (test-type: unit)
  - blocked-by: Classify existing packs
  - summary: Vendor the `paper-review` skill into
    `project/.github/skills/paper-review/` and add it to the `research` pack;
    classify `research` and `blog` as `stack`.
  - desired behavior: Opting into `research` installs `paper-review` plus the
    scientific skills; neither pack is in a default install.
  - acceptance criteria: Catalog test asserts `research` contains
    `paper-review` and that `research`/`blog` are absent from the default set.
  - verification: `go test ./internal/cli`

- [x] **Remove caveman from geremmyas** (test-type: unit)
  - blocked-by: Classify existing packs
  - summary: Delete the caveman brevity directive from distributed `AGENTS.md`
    and maintainer instructions, remove `.agents/skills/caveman*` and the
    caveman cursor rule, and recommend the upstream installer in `README.md`.
  - desired behavior: No caveman content ships from geremmyas; the README points
    to the upstream project.
  - acceptance criteria: A repo-wide check finds no caveman pack/skill and the
    README links the upstream installer.
  - verification: `go test ./internal/cli`; manual grep for `caveman`

- [x] **Documentation** (test-type: unit)
  - blocked-by: Tier-aware install
  - summary: Document tiers and the default set in `creating-packs.md`,
    `README.md`, and `specs/README.md`.
  - desired behavior: Readers understand `core` / `stack` and what
    installs by default.
  - acceptance criteria: Docs describe the tier model and migration note.
  - verification: Manual doc review; `./geremmyas doctor`
