# geremmyas

[![CI](https://github.com/woliveiras/geremmyas/actions/workflows/ci.yml/badge.svg)](https://github.com/woliveiras/geremmyas/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![ShellCheck](https://img.shields.io/badge/shell-ShellCheck-brightgreen)](https://www.shellcheck.net/)

Skills, instructions, contracts, hooks, prompts, and a pack-based CLI for coding assistants

<p align="center">
  <img src=".github/assets/geremmyas-logo.svg" alt="Geremmyas logo" width="200"/>
</p>

## Why?

Every project needs the same coding-assistant setup: language-specific
instructions, review workflows, testing conventions, and security guardrails.
Instead of copying files between repositories or tying shared content to one
assistant, install once and materialize it for the selected targets.

**What you get:**

- **Pack-based project installs** with `geremmyas.yml`, so each repository gets only the instructions and skills it needs
- **Instruction files** auto-applied by file glob for languages, frameworks, testing, and security
- **AGENTS.md** project contract for agent workflows, artifact locations, and operating rules
- **Guardrails Framework** with phase-local skills, lazy references, authority boundaries, and executable command hooks
- **Dynamic delegation contracts** for isolated review or specialist work, without permanent agent personas
- **Workflow and utility skills** for specs, tests, docs, migrations, ADRs, state management patterns, and commit messages
- **Command guardrails** that block `git push --force`, `rm -rf /`, `terraform destroy`, and other dangerous commands
- **Prompt templates** for code review, refactoring, test generation, and the technology-neutral base workflow

## Documentation

| Document | Description |
| --- | --- |
| [docs/architecture.md](docs/architecture.md) | Embed FS, pack resolution, sync preserve/overwrite, global vs project |
| [docs/creating-packs.md](docs/creating-packs.md) | How to add packs, skills, and instructions |
| [docs/guardrails-framework.md](docs/guardrails-framework.md) | Error-prevention system: gates, decisions, anti-patterns, quality workflows |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Setup, conventions, and PR flow |
| [AGENTS.md](AGENTS.md) | Agent operating contract for synced projects |

This maintainer repo also uses [`specs/README.md`](specs/README.md) for platform
features. Shared sources live under `content/`; assistant-specific adapters
live under `targets/`. Root files and selected `.github` paths are symlinks used
for dogfooding. Do not run `geremmyas project` here.

```
geremmyas/
├── install.sh / uninstall.sh          # Install scripts
├── catalog/packs.json                 # Semantic artifact catalog
├── content/                           # Assistant-neutral canonical sources
│   ├── AGENTS.md                      # Project operating contract
│   ├── instructions/                  # Language and framework guidance
│   ├── skills/                        # Workflow and utility skills
│   ├── guardrails/                    # Portable command policy
│   ├── prompts/                       # Review, refactor, test, and base prompts
│   └── templates/                     # Project artifact templates
└── targets/
    └── copilot/                       # Copilot-only instructions and hooks

Global install (`geremmyas global [--targets ...] <pack>...`):
  → ~/.agents/skills/             (targets with portable skills)
  → ~/.copilot/instructions/      (copilot target)
  → ~/.codex/instructions/        (codex target)
  → ~/.cursor/rules/              (cursor target)
  → ~/.claude/CLAUDE.md           (claude-code target)
  → ~/.config/opencode/AGENTS.md  (opencode target)
```

## Install

Install or update the `geremmyas` binary:

```bash
curl -fsSL https://raw.githubusercontent.com/woliveiras/geremmyas/main/install.sh | bash
```

The installer downloads the latest release binary to `~/.local/bin/geremmyas`.
When run from a local checkout, it can also build the binary with Go if a release
asset is not available yet.

Use `XDG_BIN_HOME` to choose another install directory:

```bash
curl -fsSL https://raw.githubusercontent.com/woliveiras/geremmyas/main/install.sh | XDG_BIN_HOME="$HOME/bin" bash
```

From a local checkout, use `GEREMMYAS_INSTALL_SOURCE=checkout` to build from source:

```bash
XDG_BIN_HOME="$HOME/bin" ./install.sh
```

## Update

Run the same installer again:

```bash
curl -fsSL https://raw.githubusercontent.com/woliveiras/geremmyas/main/install.sh | bash
```

Or from a local checkout:

```bash
./install.sh update
```

## Uninstall

Remove the binary:

```bash
curl -fsSL https://raw.githubusercontent.com/woliveiras/geremmyas/main/install.sh | bash -s -- uninstall
```

Or from a local checkout:

```bash
./install.sh uninstall
```

## Usage

### Quick start

```bash
geremmyas init              # creates geremmyas.yml (default: core, coding)
geremmyas sync              # install packs from config into cwd
geremmyas list              # all catalog packs
geremmyas version           # installed binary version
geremmyas doctor            # validate embed + config
```

Example `geremmyas.yml`:

```yaml
version: 1
packs:
  - core
  - coding
  - python-api
  - data-postgres
targets:
  - copilot
  - cursor
  - claude-code
```

Default `targets` when omitted: `copilot` only. Add `cursor`, `claude-code`, or
`opencode` to generate IDE-specific files from the same packs (see
[docs/architecture.md](docs/architecture.md#multi-ide-targets)).

Optional writing, research, and demo packs:

| Pack | Use when |
|------|----------|
| `blog` | Reviewing and rewriting technical blog posts while preserving the author's voice |
| `premortem` | Stress-testing plans, decisions, and launches by assuming failure and working backward |
| `research` | Writing, reviewing, and planning scientific papers, SLRs, peer reviews, and empirical case studies |

Add them to a repository:

```bash
geremmyas add blog research
geremmyas sync
```

### Caveman (recommended)

I strongly recommend installing *JuliusBrussee/caveman*.

```bash
curl -fsSL https://raw.githubusercontent.com/JuliusBrussee/caveman/main/install.sh | bash
```

See [JuliusBrussee/caveman](https://github.com/JuliusBrussee/caveman) for details.

### CLI reference

| Command | Purpose |
| --- | --- |
| `version` | Print the installed Geremmyas binary version |
| `list` | Print all packs (`name` + description) |
| `init [--packs a,b] [--targets copilot,cursor,...] [--force]` | Create `geremmyas.yml`; interactive TUI if no `--packs` and TTY |
| `sync [--force] [--targets ...]` | Sync packs + run IDE generators from config |
| `add <pack>...` | Append packs to config only (**does not** sync) |
| `remove <pack>...` | Remove packs from config only (**does not** delete synced files) |
| `project [--force] [--targets ...] <pack>...` | `add` + `sync` in one step; interactive pack picker available |
| `global [--targets ...] [--force] <pack>...` | Reconcile the complete user-level pack/target state |
| `global list [--targets ...] [--include-adoptable] [--json]` | Inspect managed and observed global harness state without writing |
| `global clear [--targets ...] [--include-adoptable] [--dry-run] [--force] [--json]` | Plan or safely remove global state owned by Geremmyas |
| `context [--root path] [--json]` | Report approximate context cost, selected packs, and skill ownership by source |
| `doctor` | Validate catalog sources and local `geremmyas.yml` |

**Defaults:** non-interactive `init` writes `core` and `coding`.

The default remains complete through compact fallbacks in `AGENTS.md`: its
Completion and Agent Routing sections cover fresh evidence, bounded review,
minimal docs reconciliation, and local commits without requiring another
discoverable skill. Add `quality` (or choose `base`) when you want the detailed
`verify`, review-contract, `docs`, and `git-commit` procedures available lazily.

**Sync output:** `installed`, `updated`, `preserved`, `skipped` — see
[docs/architecture.md](docs/architecture.md).

Project sync reconciles files through `.geremmyas/project-manifest.json`.
Modified, unowned, and symlinked paths are preserved; obsolete files are
removed only when they remain unchanged and owned.

Global inventory distinguishes current, modified, missing, obsolete, unowned,
symlinked, non-regular, unreadable, and adoptable files. Adoption requires an
exact canonical hash or a trusted Geremmyas generated marker; external plugin
and runtime caches are observations only and never become owned.

`global clear --dry-run` renders the same removal plan without changing files or
the manifest. A target-scoped clear retains shared skills while another target
still uses them. Modified owned files require `--force`; symlinks, non-regular
paths, unowned files, plugins, and runtime caches are never removed. Use
`--include-adoptable` only for a total clear of exact canonical legacy files.
Marker-only generated legacy files additionally require `--force`, because the
marker cannot prove that the generated body was never customized.

**Generated IDE files** (marker `geremmyas:generated`): `.cursor/rules/*.mdc`,
`.cursor/hooks.json`, `CLAUDE.md`, `.opencode/AGENTS.md`, `.codex/AGENTS.md`. Re-sync
updates them; custom edits preserved unless `--force`.

**Global install** routes semantic artifacts by target. Portable skills use
`~/.agents/skills/`; Copilot and Codex instructions are written only when their
respective targets are selected.

Use `geremmyas add <pack>` then `geremmyas sync`, or use `project` to do both.

Install packs into the current project and update `geremmyas.yml` in one step:

```bash
geremmyas project python-api data-postgres
```

Or use interactive selection:

```bash
geremmyas project
```

`geremmyas project` preserves customizable project files by default, including
`AGENTS.md`, `specs/README.md`, `mise.toml`, `.github/copilot-instructions.md`,
and guardrail hooks. Use `--force` to overwrite those files during sync:

```bash
geremmyas project --force core
```

### Global Install

Install packs to user-level directories so they apply across all projects:

```bash
geremmyas global base python-ai infra-terraform blog research
geremmyas global --targets copilot,cursor base
geremmyas global --targets claude-code,opencode coding
geremmyas global --targets codex coding python-ai
```

Each invocation is the complete desired global state, not an additive install.
Geremmyas records owned files in
`${XDG_STATE_HOME:-~/.local/state}/geremmyas/global-manifest.json`. When packs or
targets are removed, unchanged owned files are deleted; modified and unowned
files are preserved. The first managed run adopts only files whose content
exactly matches the current embedded catalog, so unknown legacy and third-party
skills are never removed automatically.

Inspect the current state without changing files or the manifest:

```bash
geremmyas global list
geremmyas global list --targets codex
geremmyas global list --include-adoptable
geremmyas global list --json
```

The inventory distinguishes current, modified, missing, obsolete, unowned,
symlinked, non-regular, unreadable, and adoptable paths. System skills and
plugin caches are reported as external observations, never as Geremmyas-owned
files. `--targets` filters the report only; it does not change installed state.

Preview or apply conservative cleanup:

```bash
geremmyas global clear --dry-run
geremmyas global clear --targets codex --dry-run
geremmyas global clear
```

Or use interactive selection:

```bash
geremmyas global
```

Default target is `copilot` for backward compatibility. With `--targets`:

| Target | Output |
| --- | --- |
| `copilot` | `~/.agents/skills/`, `~/.copilot/instructions/` |
| `cursor` | above skills + `~/.cursor/rules/`, `~/.cursor/hooks.json` |
| `claude-code` | `~/.claude/CLAUDE.md` |
| `opencode` | `~/.config/opencode/AGENTS.md` |
| `codex` | `~/.codex/AGENTS.md` + `~/.codex/instructions/` (instruction index) |

Instructions are routed only to the selected assistants: Copilot uses
`~/.copilot/instructions/`, while Codex uses `~/.codex/instructions/`. Codex
also reads `~/.codex/AGENTS.md` (its `CODEX_HOME`), which indexes each
instruction by `applyTo`.
The file is a compact bootstrap: it defers to the nearest project `AGENTS.md`
and does not duplicate Codex's native skill discovery or unsupported agent
roles.
Earlier versions wrote `~/.config/codex/AGENTS.md`; delete that stale file once
after upgrading.

After upgrading from an append-only global install, run `geremmyas global list
--include-adoptable` first, then reconcile with the complete pack and target set
you want. The migration catalog recognizes exact hashes of the renamed workflow
skills and removed agent profiles; modified legacy files remain unowned or
preserved. Former top-level policy skills now live under their owning workflows
as listed in [Guardrails Framework](docs/guardrails-framework.md#internal-references).

Generated global files preserve user edits unless you pass `--force`. A corrupt
or unsupported manifest blocks the run before any global files are written;
move the manifest aside only after reviewing the installed global directories.

### Context report

Run `geremmyas context` to compare the embedded catalog, Copilot and portable
project roots, `~/.agents/skills`, Codex system skills, and Codex plugin cache.
The report shows the project/global manifest selections and separates top-level
and nested `SKILL.md` files, frontmatter bytes, approximate tokens, and
ownership. Use `--root` to inspect another project without changing directory
and `--json` for repeatable baseline comparisons, including per-skill body and
support-file costs. Human output shows `coding`, `quality`, and `base` upper
bounds: discovery metadata, all selected top-level bodies, and support content
that should load only on demand. Token counts use `(bytes + 3) / 4`; they are a
stable comparison metric, not an exact model tokenizer result.

`context` measures installed/discoverable text, not what a host actually loaded,
when it compacted, model token billing, or wall-clock latency. Measure those
runtime effects with an external controlled A/B: same assistant/version/model,
repository snapshot, prompt suite, cache state, and repetitions; compare an
isolated empty home with project `coding`, project `base`, and an explicit
global installation. Record median and tail latency plus runtime-reported input,
cache, tool, and compaction events. Geremmyas deliberately does not call an LLM.

External system and plugin roots are observed only. `unowned` means Geremmyas
has no authority to remove that skill; `modified` means a manifest-owned global
skill no longer matches the hash Geremmyas installed. Plugin-cache counts are an
upper bound that can include inactive or older versions; the Codex host decides
which cached plugins are active in a session.

### Migration from the former SDD pack

This release intentionally makes a clean break instead of installing aliases:

| Previous capability | Replacement |
| --- | --- |
| `sdd` pack | `coding` for the efficient default, or `base` for coding + closing workflows |
| `requirements-interview` | `refine` |
| `generate-spec` | `spec` |
| `vertical-tdd` | `tdd` |
| `bugfix-loop` | `bugfix` |
| `verification-checklists` | `verify` |
| `update-docs`, `generate-glossary`, `generate-adr` | `docs`, selecting one lazy mode reference |
| `code-review-requesting` | `verify/references/review-contract.md`, used by a runtime-created subagent |
| bundled custom agents | bounded dynamic subagents created only when useful |

There are no compatibility skill directories because they would restore
discovery cost and ambiguous triggers. Project/global reconciliation removes
manifest-owned old paths that still match their installed hash. The embedded
migration hash catalog also recognizes intact pre-manifest portable skills and
agents; modified or unknown files remain untouched. Inspect global state before
cleanup:

```bash
geremmyas global list --include-adoptable
geremmyas global clear --dry-run --include-adoptable
```

### Pack catalog

Run `geremmyas list` for the live list. Dependencies are resolved automatically
(for example `supabase` pulls in `data-postgres`).

| Group | Packs | Notes |
| --- | --- | --- |
| **Baseline** | `core`, `coding`, `quality`, `base` | Default: `core` + `coding` with compact closing fallbacks; `base` adds detailed quality workflows |
| **Workflow helpers** | `decision-support`, `skill-maintenance` | Optional decision and catalog-maintainer skills |
| **Writing & research** | `blog`, `research`, `premortem` | Optional content workflows |
| **Games** | `game-core`, `game-ui`, `game-systems`, `game-performance`, `game-audio`, `game-art`, `game-delivery`, `game-dev` | Select routine 2D domains independently; `game-dev` is the explicit complete metapack |
| **TypeScript / Node** | `typescript-base`, `typescript-ci`, `node-api`, `nestjs`, `fastify` | `nestjs` / `fastify` need `node-api` |
| **React** | `react-web`, `react-router`, `react-state`, `react-data`, `tailwind` | Most depend on `react-web` → `typescript-base` |
| **Python** | `python-base`, `python-api`, `python-ai`, `python-ci`, `python-sqlite` | `python-ci` needs `infra-ci` |
| **Go** | `go-base`, `go-api`, `go-sqlite`, `go-devtools`, `go-ci` | `go-sqlite` needs `data-sqlite` |
| **Rust** | `rust-base`, `rust-ci`, `rust-release` | |
| **Data** | `data-postgres`, `data-sqlite`, `data-chromadb`, `supabase`, `node-sqlite` | `supabase` depends on `data-postgres` |
| **Infra** | `infra-docker`, `infra-ci`, `infra-gcp`, `infra-terraform` | |
| **Mobile** | `android`, `android-ci` | |
| **Other** | `astro` | Astro/MDX instructions |

Packs only install what you list in `geremmyas.yml` (plus transitive `depends`).
There is no automatic detection of `package.json` or `go.mod` in the target repo.

For a game project, prefer the smallest stable domain set instead of installing
`game-dev` automatically. For example, a gameplay-heavy Phaser project can use:

```yaml
packs:
  - base
  - typescript-base
  - game-core
  - game-ui
  - game-systems
```

Add `game-performance`, `game-audio`, `game-art`, or `game-delivery` when those
domains become active. `game-dev` depends on every focused game pack and remains
available when advertising all eleven skills is an intentional trade-off.
`game-art-2d` remains a backward-compatible pack name for `game-art`; it does
not create a duplicate skill. Smaller packs reduce the discovery surface, but
actual skill activation still depends on the assistant runtime.

You can also choose project vs global during interactive init:

```bash
geremmyas init
```

### mise (Environment Consistency)

Every project sync includes a `mise.toml` template for consistent tool versions.
After sync, activate it:

```bash
mise trust
mise install
```

This ensures all contributors use the same Go, Node, Python (or whatever tools
the project needs) versions without manual setup.

## What's Included

### Prompts and target templates

Portable prompt source templates live in `content/prompts/`. They are canonical
examples, not catalog artifacts: `geremmyas sync` and `geremmyas global` do not
install them. Copy one into an assistant's native prompt location deliberately
when that target supports custom prompts. Assistant-specific bootstrap and
instruction templates live under `targets/<assistant>/`.

| File | Purpose |
|------|---------|
| `targets/copilot/global-instructions.md` | Copilot bootstrap: find and follow local `AGENTS.md` |
| `content/prompts/review.prompt.md` | Structured code review checklist |
| `content/prompts/refactor.prompt.md` | Refactor preserving behavior |
| `content/prompts/test.prompt.md` | Generate unit tests matching project patterns |
| `content/prompts/base.prompt.md` | Full technology-neutral cycle: refine -> spec -> test -> implement -> verify -> review -> docs |

### Project outputs

The selected targets determine output paths. Shared artifacts remain at the
project root; assistant-specific content is emitted only for that target.

| Target | Principal outputs |
|------|---------|
| Shared | `AGENTS.md`, `mise.toml`, project templates |
| Copilot | `.github/skills`, `.github/instructions`, `.github/hooks` |
| Codex | `.agents/skills`, `.codex/instructions`, `.codex/AGENTS.md` |
| Cursor | `.agents/skills`, `.cursor/rules`, `.cursor/hooks.json` |
| Claude Code | `.agents/skills`, `.claude/instructions`, `CLAUDE.md` |
| OpenCode | `.agents/skills`, `.opencode/instructions`, `.opencode/AGENTS.md` |

`AGENTS.md` is the source of truth for agent behavior in a repository. It should
reference skills instead of duplicating their full procedures.
`.github/copilot-instructions.md` remains useful for project facts and
Copilot-wide context.

#### Instructions (`content/instructions/`)

Instructions are short, auto-applied rules selected by `applyTo` globs. Use them
for conventions that should be present whenever an assistant edits a matching
file. Target planners place the same source in the assistant's native
destination. Use skills for explicit workflows, and `assets/` or `references/`
for long examples and recipes.

`geremmyas lint` protects the default context budget: descriptions are limited
to 240 characters, skill bodies to 250 lines, the composed `base` workflow to 7 public
skills, and `content/AGENTS.md` to 700 words. Nested support files must use
descriptive names instead of `SKILL.md`. See
[Creating packs, skills, and instructions](docs/creating-packs.md#context-budgets).

| File | Applies To | Focus |
|------|-----------|-------|
| `typescript.instructions.md` | `**/*.ts, **/*.tsx` | Strict mode, interfaces, named exports |
| `nestjs.instructions.md` | NestJS modules, controllers, providers, and lifecycle files | Modules, DI, DTO validation, guards, interceptors, filters |
| `fastify.instructions.md` | Fastify routes, plugins, and server files | Plugins, JSON Schema contracts, hooks, logging, errors |
| `node-sqlite.instructions.md` | Node SQLite database/repository files | `node:sqlite`, `better-sqlite3`, prepared statements, transactions |
| `go.instructions.md` | `**/*.go` | Error wrapping, table-driven tests, context |
| `echo.instructions.md` | Echo handlers, routes, middleware, and server files | Handlers, middleware, context, centralized errors, graceful shutdown |
| `go-sqlite.instructions.md` | Go SQLite database/repository files | `database/sql`, `modernc.org/sqlite`, DSN pragmas, connection limits |
| `go-embed.instructions.md` | Go files in projects using `//go:embed` | `embed.FS`, package-relative paths, `fs.Sub`, read-only assets |
| `air.instructions.md` | Air config and local Go dev containers | Hot reload config, excludes, disposable binaries, dev-only usage |
| `python.instructions.md` | `**/*.py` | Python language-level conventions |
| `fastapi.instructions.md` | FastAPI route and API files | Routers, dependency injection, request/response models |
| `pydantic.instructions.md` | Pydantic schemas, models, DTOs, and settings | Pydantic v2 validation, serialization, settings, boundaries |
| `langchain.instructions.md` | LangChain chains, agents, retrievers, and RAG files | Runnables, tools, retrieval, structured outputs, tracing |
| `langgraph.instructions.md` | LangGraph graph, workflow, and agent files | State schemas, nodes, checkpoints, interrupts, resume |
| `llm-service.instructions.md` | LLM service and agent integration files | Provider boundary, structured outputs, retries, limits, observability |
| `python-sqlite.instructions.md` | Python SQLite database/repository files | `sqlite3`, SQLAlchemy SQLite, parameter binding, connection lifecycle |
| `postgres.instructions.md` | PostgreSQL SQL, migrations, repositories, and database files | Constraints, indexes, transactions, query plans, connection pools |
| `chromadb.instructions.md` | ChromaDB, RAG, retriever, and embedding files | Clients, collections, embedding functions, metadata filters, backups |
| `supabase.instructions.md` | Supabase clients, migrations, RLS, functions, and config | RLS, API keys, policies, migrations, generated types, Edge Functions |
| `kotlin.instructions.md` | `**/*.kt` | MVVM, Hilt, Room, Compose |
| `android-sqlite.instructions.md` | Android Room, DAO, entity, migration, and repository files | Room, DAOs, migrations, async queries, migration tests |
| `react.instructions.md` | `**/*.tsx, **/*.jsx` | TanStack Query, feature-sliced, a11y |
| `astro-mdx.instructions.md` | `**/*.mdx, **/*.astro` | Frontmatter, no H1, code fences |
| `testing.instructions.md` | test files and test folders | General test design independent of framework |
| `e2e-testing.instructions.md` | E2E test files and config | User journeys, stable selectors, reliable verification |
| `integration-testing.instructions.md` | integration test files and folders | Module boundaries and controlled external dependencies |
| `api-security.instructions.md` | API handlers, controllers, routes, middleware | API input, authorization, logging, abuse controls |
| `web-security.instructions.md` | browser UI and route components | XSS, redirects, client storage, client/server validation |
| `android-security.instructions.md` | Android Kotlin and manifest files | Storage, permissions, intents, networking |
| `docker.instructions.md` | Dockerfiles and `.dockerignore` files | Multi-stage builds, pinned images, non-root users, no baked secrets |
| `docker-compose.instructions.md` | Docker Compose files | Healthchecks, networks, volumes, env files, local orchestration |
| `github-actions.instructions.md` | GitHub Actions workflows and actions | Permissions, OIDC, pinned actions, concurrency, artifacts |
| `gcp.instructions.md` | GCP CLI scripts, Cloud Build, and deploy scripts | Explicit project/account, ADC vs gcloud auth, impersonation |
| `terraform.instructions.md` | Terraform `.tf` and `.tfvars` files | fmt, validate, remote state, lockfile, modules, imports, moved blocks |
| `react-router.instructions.md` | Route modules | React Router v7 loaders, actions, typegen |
| `sqlite.instructions.md` | SQL, database, storage, repository, and migration files | SQLite schema, pragmas, transactions, indexes, migrations |
| `tailwind.instructions.md` | Component TSX files | Tailwind CSS v4 utilities and pitfalls |
| `tanstack-query.instructions.md` | hooks, API, query files | TanStack Query v5 hooks and keys |
| `xstate.instructions.md` | `*.machine.ts` | XState v5 machines and actors |
| `zod.instructions.md` | schemas and API files | Zod v4 schemas and parsing |
| `zustand.instructions.md` | store files | Zustand v5 stores and middleware |

#### Runtime delegation

Geremmyas no longer distributes permanent custom-agent profiles. Capable
runtimes create bounded subagents only when isolation, specialization, or
parallelism materially helps; trivial work remains with the primary agent. Each
delegation supplies objective, scope, ownership, evidence, unknowns, output, and
authority. The primary owns integration and Git.

Independent review is a lazy reference under
`verify/references/review-contract.md`, not a discoverable skill or permanent
persona. Targets without subagents execute the same contract inline and report
that isolated context was unavailable. Generic target-adapter support remains
available for future custom agents that provide a concrete permission, model,
tool, MCP, or isolation benefit.

#### When to Use What: Code Review

There are three review surfaces, each for a different context:

| Surface | How to invoke | Best for |
|---------|--------------|----------|
| `verify` review contract | Runtime subagent after fresh verification, or directly for an explicit read-only review | **Spec-driven review** against acceptance criteria, tests, code, and risk boundaries. |
| `/review` prompt | Type `/review` in Copilot Chat | **Quick general review** — security, readability, correctness checklist. Use for fast feedback on any code, no specs needed. |
| Built-in `/review` | `/review` in Copilot CLI | **Diff-based review** — analyzes staged/branch changes automatically. Use for pre-commit or pre-PR checks. |

#### Skills (`content/skills/`)

Skills are explicit capabilities. Workflow skills guide multi-step work; utility
skills provide focused technical recipes or generated artifacts.

Active workflow skills ship with `coding`; closing capabilities ship with
`quality`; `base` composes both. Stack-specific recipe skills (for example
`validate-with-zod`, `model-state-with-xstate`, `migrate-react-router`) install
when you add the matching pack (`react-data`, `react-router`, etc.); core rules
live in the paired instruction files.

Use a top-level skill only for a capability users invoke directly. Composition
steps, checklists, examples, review contracts, and policy belong in the owning
skill's `references/`.

| Skill | Purpose |
|-------|---------|
| `refine` | Resolve material ambiguity before specification; unload when clear |
| `spec` | Create or update the numbered spec, plan, and tasks package |
| `tdd` | Implement one observable behavior at a time with red-green-refactor |
| `bugfix` | Reproduce, diagnose, regression-test, and fix bugs |
| `docs` | Route lazily to project docs, glossary, ADR/MADR, or RFC support |
| `git-commit` | Create task-owned atomic Conventional Commits from verified diffs |
| `verify` | Require fresh execution evidence and expose the lazy independent-review contract |
| `decision-framework` | Evaluate material decisions (`decision-support` pack) |
| `skill-authoring` | Create or revise skills (`skill-maintenance` pack) |
| `game-art-2d` | Create and integrate 2D runtime art (`game-art`; legacy pack name `game-art-2d`) |
| `gameplay-programming-2d` | Build movement, combat, interaction, physics, and scene flow (`game-core`) |
| `game-testing-2d` | Test deterministic simulation, engine integration, and exported behavior (`game-core`) |
| `game-feel-2d` | Tune controls, camera, impact, and moment-to-moment feedback (`game-ui`) |
| `game-ai-2d` | Build enemy perception, decisions, navigation, and encounters (`game-systems`) |
| `game-performance-2d` | Profile and optimize measured 2D game bottlenecks (`game-performance`) |
| `procedural-generation-2d` | Generate deterministic, validated maps and content (`game-systems`) |
| `game-save-n-progress` | Implement versioned saves, settings, and progression (`game-systems`) |
| `game-audio-2d` | Integrate routed, bounded, and platform-aware game audio (`game-audio`) |
| `game-ui-accessibility` | Build responsive UI, localization presentation, focus, touch, and accessibility (`game-ui`) |
| `game-build-and-release` | Produce and verify Phaser bundles and Godot exports (`game-delivery`) |
| `typescript-ci-setup` | TypeScript CI pipeline (`typescript-ci` pack) |
| `python-ci-setup` | Python CI pipeline (`python-ci` pack) |
| `go-ci-setup` | Go CI pipeline (`go-ci` pack) |
| `rust-ci-setup` | Rust CI pipeline (`rust-ci` pack) |
| `rust-release` | Rust crate/binary release engineering (`rust-release` pack) |
| `android-ci-setup` | Android CI pipeline (`android-ci` pack) |
| `terraform-change` | Plan and execute context-aware Terraform changes |
| `gcloud-operation` | Prepare safe Google Cloud CLI operations with explicit project/account context |
| `ci-workflow` | Create, review, or debug GitHub Actions CI/CD workflows |
| `llm-integration-review` | Design or review production LLM service integrations |
| `langgraph-agent-design` | Design LangGraph agents around state, checkpoints, tools, and HITL |
| `supabase-workflow` | Plan Supabase schema, RLS, Auth, Storage, and Edge Function changes |
| `postgres-query-review` | Review PostgreSQL queries, migrations, indexes, and plans |
| `chromadb-rag-workflow` | Design or review ChromaDB-backed RAG ingestion and retrieval |
| `validate-with-zod` | Zod validation recipes (API clients, forms, localStorage) |
| `migrate-react-router` | Step-by-step guide for React Router v6 → v7 migration |
| `model-state-with-xstate` | XState v5 recipes: React integration, actors, testing |
| `manage-state-with-zustand` | Zustand v5 recipes: middleware setup, immer, XState sync |
| `text-review` | Rewrite technical blog drafts while preserving voice and facts (`blog` pack) |
| `scientific-paper` | Create, review, critique, and validate scientific papers (`research` pack) |
| `scientific-case-study-research` | Design empirical SE case studies with protocol, triangulation, and validity checks (`research` pack) |
| `premortem` | Run a premortem on plans, decisions, or launches — assumes failure and works backward to find blind spots (`premortem` pack) |

#### Specs index (`specs/README.md`)

The `coding` pack installs `specs/README.md` as the repository index: **families**,
reserved number blocks, **status** lifecycle (`Draft` → `Ready` → `In Progress`
→ `Verified`), and per-family tables (Spec, Title, Status, Depends on).

Each spec folder uses a global number:

```text
specs/README.md
specs/NNNN-<feature-slug>/spec.md    # YAML frontmatter: spec, family, phase, status
specs/NNNN-<feature-slug>/plan.md
specs/NNNN-<feature-slug>/tasks.md
```

Agents update `specs/README.md` whenever a spec changes lifecycle status.

#### Autonomous workflow and progress

Every feature uses a folder with **all three** artifacts (see above). Agents
advance a package to `Ready` automatically when its acceptance criteria are
testable, contracts and dependencies are known, verification commands exist,
and no material product decision remains unresolved. Bugfixes run as an
autonomous reproduce-red-fix-green-document loop. Verified slices create atomic
local Conventional Commits by default; an explicit no-commits, read-only,
plan-only, or no-edits instruction overrides that behavior.

`tasks.md` uses checkboxes for progress (`[ ]` pending, `[~]` in progress,
`[x]` done). Agents must keep checkboxes current while working. Each task
includes a `test-type` (`unit`, `integration`, or `both`).

Use `specs/README.md` for status across specs. When resuming work, read the
feature folder (`spec.md`, `plan.md`, `tasks.md`) and continue from the
in-progress or next pending task.

Local commit authority never implies push, amend, rebase, merge, tag, release,
publication, or production deployment.

Do not create GitHub Issues, labels, or issue-state workflows unless explicitly
requested.

#### Domain Vocabulary

Agents and skills should read domain vocabulary before writing PRDs, specs,
tests, reviews, bugfix documents, ADRs, or user-facing copy.

`GLOSSARY.md` is the default vocabulary artifact for new projects. `CONTEXT.md`
is also supported for repositories that already use that convention. If both
exist, read both; treat `GLOSSARY.md` as the canonical term list and
`CONTEXT.md` as broader domain context unless the project says otherwise. If
they conflict, use repository evidence and `GLOSSARY.md` precedence, reconcile
both, and record the decision.

Absence of both files should not block work. Create or update vocabulary only
when real ambiguity, inconsistent naming, or overloaded domain language appears.

#### Copilot hooks (`.github/hooks/`)

Command guardrails that intercept dangerous terminal commands:

- **BLOCK**: force-push, destructive Git resets, broad root/home deletion,
  `terraform destroy`, and secret leaks
- **ASK**: `git push` and privileged `sudo` commands

Existing or catalogued dependencies and verified local/disposable/test targets
are autonomous. A new uncatalogued direct dependency needs provenance,
maintenance, security, and license evidence. It also needs build-versus-buy
analysis and explicit user choice before installation. Every production mutation, deploy, release,
publication, or policy change needs explicit user authorization.
Guarded Terraform, `gcloud`, and `psql` mutations require a verified
non-production marker such as `GEREMMYAS_TARGET=test`; otherwise the hook treats
the target as protected.

Rules are configurable in `guardrails-rules.txt`.

### mise.toml

Template for [mise](https://mise.jdx.dev/) tool version management. Uncomment the tools your project uses.

## Lazy coding workflow

The workflow loads only the capability needed by the current phase. A simple,
well-scoped task needs no workflow skill; ambiguity starts with `refine`, a
feature advances through `spec` and `tdd`, and completion loads `verify`.

### The Cycle

```
refine? ──▶ spec ──▶ tdd (red/green/refactor) ──▶ verify ──▶ review? ──▶ docs?
 ambiguity  durable   one behavior per cycle       evidence   isolated   matching
 only       artifacts                                      subagent     mode only
```

### Workflow by Change Type

For new features:

1. Inspect the request and repository directly; load `refine` only when a
   material requirement remains ambiguous.
2. Use `spec` to create the durable `spec.md`, `plan.md`, and `tasks.md` package.
3. Use `tdd` for one vertical behavior at a time.
4. Load `verify` only when fresh completion evidence is due.
5. When independent review is useful, have the runtime create a bounded
   subagent using `verify/references/review-contract.md`.
6. Load `docs` only when public API, architecture, setup, configuration,
   vocabulary, ADR, or RFC work is actually needed.

For existing features:

1. Use `refine` only if repository evidence does not establish whether the
   product flow changes.
2. If the product flow changes, update the PRD before writing specs.
3. If the product flow does not change, write targeted specs and continue through tasks, tests, implementation, review, and docs.

For bugs:

1. Use `bugfix` to document the symptom, impact, reproduction, and ranked
   hypotheses.
2. Write the regression test at the correct boundary and record its command,
   output, and expected failing reason before changing production code.
3. Apply the smallest root-cause fix, then rerun the regression test, original
   reproduction, and nearest relevant suite.
4. Record the actual cause, remove temporary instrumentation, and document any
   residual evidence that automation could not establish.
5. Write a postmortem only when the bug was an outage.

### Using the base prompt source

After deliberately installing `content/prompts/base.prompt.md` as a Copilot
custom prompt, invoke it in Copilot Chat:

```
/base Add user authentication with JWT
```

The prompt advances through machine-readiness, red, green, review, documentation,
and verification gates without routine approval pauses. It asks only when a
material product decision remains unresolved or an authority boundary applies.

The migration inventory lives in `catalog/workflow-gates.json`. Repository lint
scans catalogued workflow content, prompts, target adapters, and public workflow
docs, and fails when it finds an unclassified conversational pause.

### Step by Step (Manual)

You can also run each step individually:

#### 1. Refine only if needed, then write a spec

```
Use the refine skill to resolve the remaining ambiguity in JWT authentication
```

The agent asks only questions whose answers cannot be established safely from
the repository. Once the behavior is clear, it unloads that phase and records
the result in durable artifacts.

If you already know the requirements and don't need an interview:

```
Use the spec skill to create a spec for JWT auth
```

#### 2. Generate Tests

```
Use the tdd skill for the next behavior in specs/user-auth/tasks.md
```

Each acceptance criterion from the spec becomes at least one test. Tests must fail initially (red phase).

#### 3. Implement

Write code to make the tests pass. The golden rule: **never modify the tests**. If a test seems wrong, revisit the spec first.

#### 4. Review

```
Create an independent read-only subagent using the verify review contract
```

The bounded reviewer checks spec → tests → code alignment and returns findings,
evidence, unknowns, and a verdict. It is created by the assistant runtime rather
than installed as a permanent custom agent.

For quick reviews without specs, use the `/review` prompt instead.

#### 5. Update Docs

```
Use the docs skill in project-docs mode for the user-auth feature
```

Only needed when public API, architecture, or setup changed. Skip for internal-only changes.

### Architecture Decisions

When the feature involves significant design choices, use these before or alongside the SDD cycle:

| Need | Tool | Output |
|------|------|--------|
| Explore architecture opportunities | Dynamic bounded subagent when isolation adds value | Evaluated options and evidence |
| Record a durable decision | `docs` skill in ADR mode | ADR in `docs/decisions/` (MADR 4.0) |
| Develop a proposal | `docs` skill in RFC mode | Repository-local RFC |

## Customization

All synced files are meant to be edited. After running `geremmyas sync`:

1. **Review `AGENTS.md`** — adjust artifact paths, workflow rules, and skill routing for the repository
2. **Edit `.github/copilot-instructions.md`** — fill in project name, description, directory structure, and build commands
3. **Add more packs** — run `geremmyas add <pack>` and `geremmyas sync`
4. **Tune guardrails** — add or remove rules in `guardrails-rules.txt`

## Contributing

```bash
mise trust && mise install   # setup Go toolchain
go test ./...                # run tests
go build ./cmd/geremmyas     # build binary
```

See [CONTRIBUTING.md](CONTRIBUTING.md) and [docs/creating-packs.md](docs/creating-packs.md)
for contributor guidelines.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

## License

MIT
