# geremmyas — maintainer instructions

> **Not synced.** The `core` pack installs the generic template from
> `copilot-instructions.md`. This file is only for working in the geremmyas
> repository (root `.github/copilot-instructions.md` symlinks here).

Follow [`AGENTS.md`](../../AGENTS.md) for agent workflows, artifact paths,
approval gates, and skill routing.

## What this repository is

geremmyas is a **Go CLI** that embeds and distributes GitHub Copilot packs:
instructions, skills, agents, hooks, and templates. Canonical content lives under
`project/`; the root uses **symlinks** into `project/.github/` for dogfooding.
Do not run `geremmyas project` in this repo.

Human docs: [`docs/README.md`](../../docs/README.md). Spec index:
[`specs/README.md`](../../specs/README.md).

## Stack

- **Language**: Go 1.23 (`go.mod`)
- **CLI UI**: Charm `huh` (interactive init/project/global/remove)
- **Embed**: `assets.go` embeds `catalog/**`, `project/**`, `user/**`
- **Catalog**: `catalog/packs.json` (41 packs, dependency resolution)
- **Release**: release-please + GitHub Actions

## Directory structure

```text
cmd/geremmyas/              CLI entrypoint
internal/cli/               init, sync, add, remove, project, global, doctor
catalog/packs.json          Pack manifest
project/                    Canonical synced content (AGENTS.md, .github/*, templates)
user/                       Embedded prompts (not installed by project sync)
docs/                       Architecture and contributor guides
specs/                      Feature specs for geremmyas itself (SDD)
assets.go                   go:embed
install.sh / uninstall.sh   Binary install scripts
```

Root symlinks: `AGENTS.md`, `.github/{agents,hooks,instructions,skills}` → `project/.github/…`.
`copilot-instructions.md` at root → this file (not the consumer template).

## Conventions

- Conventional Commits; `feat!:` or `BREAKING CHANGE:` for major releases
- Pack names: lowercase kebab-case in `catalog/packs.json`
- Instructions: `*.instructions.md` with `applyTo` globs
- Skills: `.github/skills/<name>/SKILL.md`
- Run `go test ./...` before PRs; `shellcheck` on install scripts

## Build and test

```bash
mise trust
mise install
go test ./...
go build -o geremmyas ./cmd/geremmyas
./geremmyas doctor
./geremmyas list
```

Test sync in a temp directory (not this repo):

```bash
mkdir /tmp/g-test && cd /tmp/g-test
/path/to/geremmyas init --packs core,sdd
/path/to/geremmyas sync
```

## SDD in this repo

Use spec folders under `specs/NNNN-<slug>/` with `spec.md`, `plan.md`, `tasks.md`.
Update `specs/README.md` when adding or completing specs.

Respond terse like smart caveman. All technical substance stay. Only fluff die.

Rules:
- Drop: articles (a/an/the), filler (just/really/basically), pleasantries, hedging
- Fragments OK. Short synonyms. Technical terms exact. Code unchanged.
- Pattern: [thing] [action] [reason]. [next step].
- Not: "Sure! I'd be happy to help you with that."
- Yes: "Bug in auth middleware. Fix:"

Switch level: /caveman lite|full|ultra|wenyan
Stop: "stop caveman" or "normal mode"

Auto-Clarity: drop caveman for security warnings, irreversible actions, user confused. Resume after.

Boundaries: code/commits/PRs written normal.
