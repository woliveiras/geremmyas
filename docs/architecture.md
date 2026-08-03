# geremmyas architecture

geremmyas is a Go CLI that distributes coding-assistant content as packs. The
catalog describes semantic artifacts; target adapters decide where each
artifact is materialized for Copilot, Codex, Cursor, Claude Code, or OpenCode.

## Repository layout

```text
cmd/geremmyas/          CLI entrypoint
internal/cli/           Commands, catalog, planning, sync, and global install
catalog/packs.json      Pack manifest (kind, source, target-neutral path)
content/                Canonical assistant-neutral content
targets/                Assistant-specific adapters and templates
assets.go               //go:embed catalog/** content/** targets/**
```

`content/` owns portable contracts, skills, instructions, prompts,
templates, guardrails, and tooling. `targets/<assistant>/` exists only when an
artifact or runtime integration is genuinely assistant-specific. This prevents
the source tree from treating `.github/` as the universal content model.

### Maintainer repo vs consumer repo

This repository dogfoods the same sources without running `geremmyas project`:

| Root path | Canonical source |
| --- | --- |
| `AGENTS.md` | `content/AGENTS.md` |
| `.github/instructions`, `.github/skills` | `content/` |
| `.github/hooks` | `targets/copilot/hooks/` |
| `.github/copilot-instructions.md` | `targets/copilot/maintainer-instructions.md` |

GitHub platform files such as workflows and issue templates remain in
`.github/`; their location is a GitHub convention, not assistant coupling.
Edit the canonical source rather than its root symlink. Do not run project sync
in this repository.

## Embedded filesystem and catalog

```go
//go:embed catalog/** content/** targets/**
var EmbeddedFiles embed.FS
```

Each catalog entry has:

- `kind`: semantic type such as `skill`, `instruction`, `agent`, `hook`, or
  `contract`;
- `source`: embedded source path;
- `path`: target-neutral relative name used by the materializer.

There is no consumer destination in the catalog. Rebuild the binary after
changing embedded content, and run `geremmyas doctor` to validate sources and
artifact metadata.

Pack dependencies are resolved in dependency order. Duplicate semantic
artifacts are deduplicated before target planning.

## Configuration

```yaml
version: 1
packs:
  - core
  - coding
targets:
  - copilot
  - codex
```

- Default packs for non-interactive `init`: `core`, `coding`.
- `AGENTS.md` supplies compact completion, review, docs, and commit fallbacks
  when `quality` is absent. `base` adds the detailed lazy procedures.
- Default target when omitted: `copilot`, for compatibility.
- Valid targets: `copilot`, `codex`, `cursor`, `claude-code`, `opencode`.
- `add` and `remove` edit configuration only.
- `sync` and `project` resolve packs and materialize the selected target union.

## Project materialization

Portable sources can appear at different destinations:

| Target | Principal project outputs |
| --- | --- |
| Copilot | `.github/skills`, `.github/instructions`, `.github/hooks` |
| Codex | `.agents/skills`, `.codex/instructions`, `.codex/AGENTS.md` |
| Cursor | `.agents/skills`, `.cursor/rules`, `.cursor/hooks.json` |
| Claude Code | `.agents/skills`, `.claude/instructions`, `CLAUDE.md` |
| OpenCode | `.agents/skills`, `.opencode/instructions`, `.opencode/AGENTS.md` |

`AGENTS.md`, `mise.toml`, and selected templates are shared project outputs.
Mixed target selection produces the union, so Copilot paths appear only when
Copilot is selected.

Generated adapter files include `geremmyas:generated`. Customized generated
files are preserved unless `--force` is used.

### Project ownership and reconciliation

Project state is recorded in `.geremmyas/project-manifest.json`. It stores the
selected packs, targets, destination paths, and hashes written by Geremmyas.
Each sync computes the complete desired state and removes obsolete files only
when they are still unchanged and owned.

Modified, unowned, and symlinked files are preserved. A missing manifest adopts
only exact matches from the current catalog or `catalog/legacy-artifacts.json`.
That migration catalog contains paths and hashes, not installable content, so
renamed skills and removed portable agent profiles can be retired without
remaining discoverable. A corrupt or unsupported manifest stops the operation
before files are copied. Destination traversal is symlink-safe and will not
write through a selected target path.

### Sync summary

| Counter | Meaning |
| --- | --- |
| `installed` | New file |
| `updated` | Owned or forced file replaced |
| `preserved` | Local content intentionally left unchanged |
| `skipped` | Content already current or duplicate output |

`geremmyas project` updates `geremmyas.yml` and then performs the same sync.

## Global materialization

`geremmyas global [--targets ...] [--force] <pack>...` treats each invocation
as the complete desired global state:

| Target | Principal global outputs |
| --- | --- |
| Copilot | `~/.agents/skills`, `~/.copilot/instructions` |
| Codex | `~/.agents/skills`, `~/.codex/instructions`, `~/.codex/AGENTS.md` |
| Cursor | `~/.agents/skills`, `~/.cursor/rules`, `~/.cursor/hooks.json` |
| Claude Code | `~/.agents/skills`, `~/.claude/CLAUDE.md` |
| OpenCode | `~/.agents/skills`, `~/.config/opencode/AGENTS.md` |

Copilot instructions are written only when Copilot is selected. Codex
instructions are written only when Codex is selected. Mixed selection produces
their union.

Global ownership is recorded at
`${XDG_STATE_HOME:-$HOME/.local/state}/geremmyas/global-manifest.json`.
Reconciliation follows the same conservative rules as project sync: only
unchanged owned files are removed, unknown files are never claimed, and corrupt
manifests fail before mutation.

`geremmyas global list` is the read-only view of that state. It inspects each
manifest path with `Lstat`, rejects paths outside managed roots, does not follow
symlink components, and reports current, modified, missing, obsolete, unowned,
symlink, non-regular, unreadable, or exact canonical adoptable state. Human and
JSON output share the same inventory model. Target filters affect presentation
only. Codex system skills and plugin caches are observed separately and never
acquire ownership.

Historical manifests whose pack names no longer resolve remain readable for
migration: the report sets `catalog_resolved` to false, includes a
`catalog_error`, and does not infer that any owned path is obsolete. Invalid
manifest versions, targets, paths, and normalized path collisions fail closed.

`geremmyas global clear` consumes the same inventory through a plan/apply split.
Dry-run stops after planning. Apply rechecks path type, symlink components, and
content hash immediately before every removal. Target-scoped clearing removes
target-native files but keeps shared skills until no installed target remains.
Modified owned regular files are preserved unless `--force` is explicit;
symlinked, non-regular, unreadable, unowned, and external paths remain protected
even with force. `--include-adoptable` recognizes exact canonical hashes and
generated markers, but marker-only candidates additionally require `--force`
because their original generated bytes cannot be reconstructed. Shared
adoptables require a total, unfiltered clear.

Mutating global commands acquire an OS advisory lock through
`global-manifest.lock` beside the manifest and hold it across snapshot
validation, filesystem changes, and atomic manifest replacement. The operating
system releases the lock with the process descriptor, including after a crash;
the persistent lock file itself does not mean the lock is held. Concurrent
Geremmyas mutations fail before changing files, while inventory and dry-run
remain lock-free and read-only.

## Context diagnostics

`geremmyas context [--root path] [--json]` inventories embedded catalog skills,
Copilot and portable project roots, `~/.agents/skills`, Codex system skills, and
the Codex plugin cache. It reports project and global manifest selections so
isolated no-state, `coding`, and `base` baselines remain distinguishable.
Filesystem walks do not follow symlinks. Managed paths come from ownership
manifests; system and plugin paths are observed only.

Frontmatter cost uses `(bytes + 3) / 4` as a deterministic comparison metric,
not an exact model tokenizer result. JSON includes each catalog skill's
discovery, body, and support-file upper bounds; human output aggregates those
figures for `coding`, `quality`, and `base`. Support cost is potential lazy
content, not a claim that the runtime loaded every reference.

Actual activation, compaction, provider token accounting, and latency remain
runtime behavior. They require a controlled external A/B with the same
assistant, model, repository snapshot, prompt suite, cache conditions, and
repetitions. The CLI does not add an LLM client or infer those values from files.

## CI and releases

- `ci.yml` runs Go tests, structural checks, skill lint, and shell checks.
- `release.yml` uses release-please and builds cross-platform binaries.
- Breaking changes use `feat!:` or a `BREAKING CHANGE:` footer.

## Mental model

```mermaid
flowchart LR
  subgraph build [Build time]
    catalog[catalog/packs.json]
    content[content/**]
    adapters[targets/**]
    embed[embed.FS]
    catalog --> embed
    content --> embed
    adapters --> embed
  end
  subgraph runtime [Runtime]
    cli[geremmyas CLI]
    config[geremmyas.yml]
    planner[target-aware planner]
    project[consumer project]
    global[user-level directories]
  end
  embed --> cli
  config --> planner
  cli --> planner
  planner -->|sync / project| project
  planner -->|global| global
```
