# geremmyas architecture

geremmyas is a Go CLI that ships Copilot agent content as **packs**. At build
time, `catalog/`, `project/`, and `user/` are embedded into the binary. At
runtime, the CLI copies selected pack files into a target repository or into
user-level VS Code paths.

## Repository layout

```text
cmd/geremmyas/          CLI entrypoint
internal/cli/dashboard/ Dashboard parser, renderer, git metrics, serve/watch
internal/cli/           Commands, catalog, config, sync, global install
catalog/packs.json      Pack manifest (names, depends, source → target)
project/                Canonical files synced into consumer repos
user/                   Embedded prompts/bootstrap (not copied by project sync)
assets.go               //go:embed catalog/** project/** user/**
```

### Maintainer repo vs consumer repo

This repository **dogfoods** content without running `geremmyas project`:

| Path at repo root | How it maps |
| --- | --- |
| `AGENTS.md` | Symlink → `project/AGENTS.md` |
| `.github/agents`, `hooks`, `instructions`, `skills` | Symlinks → `project/.github/…` |
| `.github/copilot-instructions.md` | Symlink → `project/.github/copilot-instructions.geremmyas.md` (maintainer only) |
| `project/.github/copilot-instructions.md` | Generic template; synced by `core` pack to consumers |

Edit under `project/`; the root symlinks stay in sync. Do **not** run
`geremmyas project` here — sync would replace symlinks with plain copies and
fork the canonical tree.

Consumer repos use `geremmyas.yml` + `geremmyas sync` (or `geremmyas project`)
to copy from the embedded filesystem, not from your local checkout paths.

## Embedded filesystem

```go
//go:embed catalog/** project/** user/**
var EmbeddedFiles embed.FS
```

Implications:

- Pack `source` paths in `catalog/packs.json` must exist under those trees.
- Changing `project/` or `catalog/` requires **rebuilding** the binary for local
  testing (`go build ./cmd/geremmyas`).
- Released binaries only include what was embedded at release build time.

Run `geremmyas doctor` to verify every pack source path is present in the embed.

## Configuration

`geremmyas.yml` in the target repository:

```yaml
version: 1
packs:
  - core
  - sdd
targets:
  - copilot
  - cursor
  - claude-code
```

- Default packs for non-interactive `init`: `core`, `sdd`.
- Default targets when omitted: `copilot` only (backward compatible).
- Valid targets: `copilot`, `cursor`, `claude-code`, `opencode`.
- `add` / `remove` only edit packs; they do **not** sync files.
- `sync` and `project` read the config, resolve dependencies, then run pack sync
  and IDE generators.

## Multi-IDE targets

The canonical source is always `project/` (embedded at build time). Targets
select which IDE-specific outputs `sync` / `project` generate:

| Target | Output | Source |
| --- | --- | --- |
| `copilot` | `.github/skills/`, `.github/agents/`, `.github/instructions/`, hooks | Pack file copy (existing behavior) |
| `cursor` | `.cursor/rules/*.mdc`, `.cursor/hooks.json` | Instructions, skills, agents, hooks |
| `claude-code` | `CLAUDE.md` | `AGENTS.md` + skill/agent index |
| `opencode` | `.opencode/AGENTS.md` | Same as Claude Code |

`AGENTS.md` is portable — Cursor, Claude Code, and OpenCode all read it. Targets
add IDE-native formats on top.

Generated files include a `geremmyas:generated` marker. Re-sync updates them;
custom edits are preserved unless you pass `--force`.

Override targets per run: `geremmyas sync --targets copilot,cursor`.

## Pack resolution

`catalog/packs.json` lists packs with optional `depends`. When you request
`nestjs`, the CLI resolves the chain (for example `nestjs` → `node-api` →
`typescript-base`) and installs all required packs in dependency order.

Duplicate `target` paths across packs: the first copy wins; later copies count
as `skipped`.

## Project sync (`sync`, `project`)

For each pack file entry:

- **File** — copy `source` → `target` under the current working directory.
- **Directory** — walk all files recursively (skills, hooks, agents).

### Sync summary counters

| Counter | Meaning |
| --- | --- |
| `installed` | New file on disk |
| `updated` | Existing file overwritten |
| `preserved` | Customizable file left unchanged (content differed, no `--force`) |
| `skipped` | Unchanged content, or duplicate target already copied |

### Customizable targets

Unless `--force`, these paths are **not** overwritten when local content differs:

- `AGENTS.md`
- `specs/README.md`
- `mise.toml`
- `.github/copilot-instructions.md`
- `.github/hooks/guardrails-rules.txt`

Everything else (skills, instructions, agents, etc.) is updated when the embed
differs.

`geremmyas project` = update `geremmyas.yml` + run sync. Interactive `project`
can ask once whether to force-overwrite customizable files.

## Global install (`global`)

`geremmyas global [--targets ...] [--force] <pack>...` installs packs to
user-level paths. Default target is `copilot` (backward compatible).

| Target | File copy | Generated output |
| --- | --- | --- |
| `copilot` | skills → `~/.agents/skills/`, instructions → `~/.copilot/instructions/` | — |
| `cursor` | skills → `~/.agents/skills/` (skill rules need them) | `~/.cursor/rules/*.mdc`, `~/.cursor/hooks.json` |
| `claude-code` | — | `~/.claude/CLAUDE.md` |
| `opencode` | — | `~/.config/opencode/AGENTS.md` |

Examples:

```bash
geremmyas global core sdd                              # copilot only (default)
geremmyas global --targets copilot,cursor core sdd     # VS Code + Cursor user rules
geremmyas global --targets cursor sdd                  # Cursor rules + skills, no Copilot instructions
geremmyas global --targets claude-code,opencode sdd    # IDE docs only
```

Global Cursor skill rules reference `~/.agents/skills/<name>/SKILL.md`. Global
agent rules embed full agent content (agents are not copied to a separate global
path).

Not installed globally: `AGENTS.md`, `mise.toml`, project-level agents/hooks,
`copilot-instructions.md`, `specs/README.md`, templates.

Global **file copies** always overwrite. Generated files follow the same
preserve/overwrite rules as project sync (`geremmyas:generated` marker, `--force`
to overwrite customized files).

## Dashboard (`dashboard`)

`geremmyas dashboard` scans `specs/`, `docs/prds/`, and `docs/bugfixes/`, then
writes a static site under `.geremmyas/dashboard/` (default). Implementation
lives in `internal/cli/dashboard/` with embedded templates under
`internal/cli/dashboard/dashboard_assets/`.

Pipeline: **parse** → optional **git dates** (`.geremmyas-cache/gitdates.json`)
→ **metrics** → **render HTML** → overwrite **`specs/README.md`** (compact index).

| Flag | Effect |
| --- | --- |
| `--output DIR` | Output directory (default `.geremmyas/dashboard`) |
| `--no-git` | Skip git log; metrics page shows unavailable state |
| `--no-cache` | Full git rescan |
| `--serve` | Serve output on `127.0.0.1` (default port 8080) |
| `--watch` | Re-run full pipeline on `specs/` / `docs/` changes (implies `--serve`) |

`sync` still **preserves** hand-written `specs/README.md`; `dashboard`
**replaces** it by design (see PRD). Recommend `.geremmyas-cache/` in
`.gitignore`.

## `user/` directory

Embedded but **not** installed by `sync` / `project`. Contains optional global
bootstrap (`user/copilot-instructions.md`) and prompts (`user/prompts/*.md`).
Copy or reference these manually if you use them outside pack sync.

## CI and releases

- **CI** (`ci.yml`): tests on changes to Go and catalog paths.
- **Release** (`release.yml`): release-please bumps version, tags, builds
  cross-platform binaries, uploads to GitHub Releases.
- **geremmyas.yml** workflow: build matrix on PR/push (subset of paths).

Breaking releases: use Conventional Commits with `feat!:` or `BREAKING CHANGE:`
in the commit body; release-please bumps the major version.

## Mental model

```mermaid
flowchart LR
  subgraph build [Build time]
    catalog[catalog/packs.json]
    project[project/**]
    user[user/**]
    embed[embed.FS in binary]
    catalog --> embed
    project --> embed
    user --> embed
  end
  subgraph runtime [Runtime]
    CLI[geremmyas CLI]
    yml[geremmyas.yml]
    repo[Consumer repo]
    global["~/.agents/skills\n~/.copilot/instructions\n~/.cursor/rules\n~/.claude/CLAUDE.md"]
  end
  embed --> CLI
  yml --> CLI
  CLI -->|sync / project| repo
  CLI -->|global| global
```
