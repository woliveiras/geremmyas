# geremmyas — maintainer instructions

> **Not synced.** The `core` pack installs the generic template from
> `project-instructions.md`. This file is only for working in the geremmyas
> repository (root `.github/copilot-instructions.md` symlinks here).

Follow [`AGENTS.md`](../../AGENTS.md) for agent workflows, artifact paths,
authority boundaries, and skill routing.

## What this repository is

geremmyas is a **Go CLI** that embeds and distributes coding-assistant packs:
instructions, skills, contracts, hooks, and templates. Assistant-neutral content
lives under `content/`; target adapters live under `targets/`. The root uses
**symlinks** into those trees for dogfooding.
Do not run `geremmyas project` in this repo.

Human docs: [`docs/README.md`](../../docs/README.md). Spec index:
[`specs/README.md`](../../specs/README.md).

## Stack

- **Language**: Go 1.23 (`go.mod`)
- **CLI UI**: Charm `huh` (interactive init/project/global/remove)
- **Embed**: `assets.go` embeds `catalog/**`, `content/**`, `targets/**`
- **Catalog**: `catalog/packs.json` (dependency resolution)
- **Release**: release-please + GitHub Actions

## Directory structure

```text
cmd/geremmyas/              CLI entrypoint
internal/cli/               init, sync, add, remove, project, global, doctor
catalog/packs.json          Pack manifest
content/                    Canonical assistant-neutral content
targets/                    Assistant-specific adapters and templates
docs/                       Architecture and contributor guides
specs/                      Feature specs for geremmyas itself (SDD)
assets.go                   go:embed
install.sh / uninstall.sh   Binary install scripts
```

Root `AGENTS.md`, `.github/instructions`, and `.github/skills` expose canonical
content; `.github/hooks` and `.github/copilot-instructions.md` expose the
Copilot adapter. Permanent custom agents are not bundled.

## Conventions

- Conventional Commits; `feat!:` or `BREAKING CHANGE:` for major releases
- Pack names: lowercase kebab-case in `catalog/packs.json`
- Instructions: `*.instructions.md` with `applyTo` globs
- Skills: `content/skills/<name>/SKILL.md`
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
/path/to/geremmyas init --packs core,coding
/path/to/geremmyas sync
```

## SDD in this repo

Use spec folders under `specs/NNNN-<slug>/` with `spec.md`, `plan.md`, `tasks.md`.
Update `specs/README.md` when adding or completing specs.
