# geremmyas

[![CI](https://github.com/woliveiras/geremmyas/actions/workflows/ci.yml/badge.svg)](https://github.com/woliveiras/geremmyas/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![ShellCheck](https://img.shields.io/badge/shell-ShellCheck-brightgreen)](https://www.shellcheck.net/)

Opinionated GitHub Copilot configurations for VS Code.

Agents, skills, instructions, hooks, prompts, and a pack-based CLI named
`geremmyas`.

<p align="center">
  <img src=".github/assets/geremmyas-logo.png" alt="Geremmyas Logo" width="200"/>
</p>

## Why?

Every project needs the same Copilot setup: language-specific instructions, code review prompts, testing conventions, security guardrails. Instead of copying files between repos, install once and apply everywhere.

**What you get:**

- **Pack-based project installs** with `geremmyas.yml`, so each repository gets only the instructions and skills it needs
- **Instruction files** auto-applied by file glob for languages, frameworks, testing, and security
- **AGENTS.md** project contract for agent workflows, artifact locations, and operating rules
- **4 agents** for spec-driven development: write specs → generate tests → implement code → update docs
- **Workflow and utility skills** for specs, tests, docs, migrations, ADRs, state management patterns, and commit messages
- **Command guardrails** that block `git push --force`, `rm -rf /`, `terraform destroy`, and other dangerous commands
- **Prompt templates** for code review, refactoring, test generation, and SDD workflow

```
geremmyas/
├── install.sh / uninstall.sh          # Install scripts
├── user/
│   ├── copilot-instructions.md        # Global bootstrap that points agents to local AGENTS.md
│   └── prompts/                       # Global prompts (review, refactor, test)
└── project/
    ├── AGENTS.md                      # Project-level operating contract for agents
    ├── mise.toml                      # Tool version management (keeps env consistent)
    └── .github/
        ├── copilot-instructions.md    # Project-level Copilot instructions
        ├── instructions/              # Auto-applied language, framework, and convention files
        ├── agents/                    # 4 agents + 6 design references
        │   └── references/            # Deep modules, interface design, etc.
        ├── skills/                    # Workflow and utility skills with asset templates
        └── hooks/                     # Command guardrails (BLOCK/ASK rules)

Global install (geremmyas global):
  → ~/.agents/skills/           (user-level skills)
  → ~/.copilot/instructions/    (user-level instructions)
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

Create a config in a repository:

```bash
geremmyas init
```

Install the declared packs:

```bash
geremmyas sync
```

List available packs:

```bash
geremmyas list
```

Example `geremmyas.yml`:

```yaml
version: 1
packs:
  - core
  - sdd
  - python-api
  - data-postgres
```

Optional writing, research, and demo packs:

| Pack | Use when |
|------|----------|
| `blog` | Reviewing and rewriting technical blog posts while preserving the author's voice |
| `brag-me` | Preparing local demo brag decks from merged PRs, commits, cloud metrics, issue trackers, and manual notes |
| `premortem` | Stress-testing plans, decisions, and launches by assuming failure and working backward |
| `research` | Writing, reviewing, and planning scientific papers, SLRs, peer reviews, and empirical case studies |

Add them to a repository:

```bash
geremmyas add blog research brag-me
geremmyas sync
```

Use `brag-me` from Copilot or another compatible agent to create or refresh a
demo under `me/brag-me/YYYY-MM-DD-highlight/`. The skill fills
`YYYY-MM-DD-brag.md` from available evidence such as merged PRs, git commits,
GCP data, Sentry issues, and manual notes, then generates an offline
Reveal.js-style `index.html` deck that opens directly in a browser.

Use `geremmyas add <pack>` and `geremmyas remove <pack>` to update the config.
Run `geremmyas doctor` to validate the catalog and local config.

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

Install packs to your VS Code user-level directory so they apply across all projects:

```bash
geremmyas global sdd python-ai infra-terraform blog research
```

Or use interactive selection:

```bash
geremmyas global
```

Global packs are installed to:
- **Skills**: `~/.agents/skills/`
- **Instructions**: `~/.copilot/instructions/`

These are the standard VS Code user-level paths, shared across all workspaces.

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


### User-level Templates

The `user/` directory contains optional prompt and instruction templates. The
binary installer does not copy these files globally. Repository setup is driven
by `geremmyas.yml` and `geremmyas sync`.

| File | Purpose |
|------|---------|
| `user/copilot-instructions.md` | Bootstrap template: find and follow local `AGENTS.md` |
| `user/prompts/review.prompt.md` | Structured code review checklist |
| `user/prompts/refactor.prompt.md` | Refactor preserving behavior |
| `user/prompts/test.prompt.md` | Generate unit tests matching project patterns |
| `user/prompts/sdd.prompt.md` | Full SDD cycle: spec -> test -> implement -> review -> docs |

### Project-level


Installed to the project root and `.github/`.

| File | Purpose |
|------|---------|
| `AGENTS.md` | Project operating contract for agents: workflows, artifact paths, skill routing, verification rules |
| `.github/copilot-instructions.md` | Project overview, conventions, directory structure, and build/test commands |

`AGENTS.md` is the source of truth for agent behavior in a repository. It should
reference skills instead of duplicating their full procedures.
`.github/copilot-instructions.md` remains useful for project facts and
Copilot-wide context.

#### Instructions (`.github/instructions/`)

Instructions are short, auto-applied rules selected by `applyTo` globs. Use them
for conventions that should be present whenever Copilot edits a matching file.
Use skills for explicit workflows, and `assets/` or `references/` for long
examples and recipes.

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

#### Agents (`.github/agents/`)

Agents are narrow roles with a specific output contract. They should route to
skills for reusable workflows instead of duplicating skill procedures.

| Agent | Role |
|-------|------|
| `spec-writer` | Explores requirements and routes to `requirements-interview`, `generate-spec`, and `task-breakdown` conventions. |
| `explorer` | Read-only codebase mapper with a structured project summary. |
| `reviewer` | Spec-driven reviewer that checks specs, tests, and code alignment. |
| `architect` | Multi-design evaluation: explores → candidates → 3 parallel sub-designs → recommends. Saves ADRs or implementation plans. |

Agents reference design heuristics in `.github/agents/references/` (deep modules, interface design, complexity signals, dependency categories, pragmatic heuristics, seam finding).

#### When to Use What: Code Review

There are three review surfaces — each for a different context:

| Surface | How to invoke | Best for |
|---------|--------------|----------|
| `@reviewer` agent | `@reviewer review this change` | **Spec-driven review** — verifies code against specs and tests, checks architecture with deep-module heuristics. Use when specs exist. |
| `/review` prompt | Type `/review` in Copilot Chat | **Quick general review** — security, readability, correctness checklist. Use for fast feedback on any code, no specs needed. |
| Built-in `/review` | `/review` in Copilot CLI | **Diff-based review** — analyzes staged/branch changes automatically. Use for pre-commit or pre-PR checks. |

#### Skills (`.github/skills/`)

Skills are explicit capabilities. Workflow skills guide multi-step work; utility
skills provide focused technical recipes or generated artifacts.

Workflow and artifact skills install broadly. Stack-specific recipe skills such
as `validate-with-zod`, `model-state-with-xstate`, `manage-state-with-zustand`,
and `migrate-react-router` install only when the matching dependency is detected;
their core rules live in the matching instruction files.

Recommended organization for future skills:

- `engineering`: PRD, specs, planning, tasks, TDD, bugfix, review, architecture, docs.
- `productivity`: skill authoring, concise communication, alignment before execution.
- `personal`: workflows specific to your setup that should not install by default.
- `utils`: rare tools, migrations, setup helpers, and guardrails.

| Skill | Purpose |
|-------|---------|
| `requirements-interview` | Explore code and clarify requirements before PRD/spec work |
| `generate-spec` | Fill a structured spec from direct input (no interview) |
| `task-breakdown` | Convert PRD, specs, or `plan.md` into vertical tasks with checkboxes |
| `generate-tests-from-spec` | Generate unit or integration tests from approved spec criteria |
| `vertical-tdd` | Implement one behavior at a time with red-green-refactor |
| `bugfix-loop` | Reproduce, diagnose, regression-test, and fix bugs |
| `update-docs` | Update `docs/` after implementing a feature |
| `git-commit` | Review staged changes and create a Conventional Commit with confirmation |
| `generate-glossary` | Extract domain terminology into `GLOSSARY.md` |
| `generate-adr` | Record an Architectural Decision in MADR 4.0 format |
| `skill-authoring` | Create or revise Copilot skills using this repo's conventions |
| `terraform-change` | Plan and review Terraform changes with approval gates |
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

The `sdd` pack installs `specs/README.md` as the repository index: **families**,
reserved number blocks, **status** lifecycle (`Draft` → `Approved` →
`Implemented`), and per-family tables (Spec, Title, Status, Depends on).

Each spec folder uses a global number:

```text
specs/README.md
specs/NNNN-<feature-slug>/spec.md    # YAML frontmatter: spec, family, phase, status
specs/NNNN-<feature-slug>/plan.md
specs/NNNN-<feature-slug>/tasks.md
```

Agents update `specs/README.md` when creating, approving, or completing a spec.

#### Interactive workflow and progress

Work is interactive: the human approves specs and bugfix proposals before
implementation. Every feature uses a folder with **all three** artifacts (see
above).

`tasks.md` uses checkboxes for progress (`[ ]` pending, `[~]` in progress,
`[x]` done). Agents must keep checkboxes current while working. Each task
includes a `test-type` (`unit`, `integration`, or `both`).

Use `specs/README.md` for status across specs. When resuming work, read the
feature folder (`spec.md`, `plan.md`, `tasks.md`) and continue from the
in-progress or next pending task.

At the start of `requirements-interview`, the agent asks whether it may create
git commits or the developer handles commits.

Do not create GitHub Issues, labels, or issue-state workflows unless explicitly
requested.

#### Domain Vocabulary

Agents and skills should read domain vocabulary before writing PRDs, specs,
tests, reviews, bugfix documents, ADRs, or user-facing copy.

`GLOSSARY.md` is the default vocabulary artifact for new projects. `CONTEXT.md`
is also supported for repositories that already use that convention. If both
exist, read both; treat `GLOSSARY.md` as the canonical term list and
`CONTEXT.md` as broader domain context unless the project says otherwise. If
they conflict, ask before changing either file.

Absence of both files should not block work. Create or update vocabulary only
when real ambiguity, inconsistent naming, or overloaded domain language appears.

#### Hooks (`.github/hooks/`)

Command guardrails that intercept dangerous terminal commands:

- **BLOCK**: `git push --force`, `rm -rf /`, `terraform destroy`, secret leaks
- **ASK**: `git push`, `sudo`, `DROP TABLE`, `pip install`, `npm install -g`

Rules are configurable in `guardrails-rules.txt`.

### mise.toml

Template for [mise](https://mise.jdx.dev/) tool version management. Uncomment the tools your project uses.

## Spec Driven Development (SDD) Workflow

This project is built around **Spec Driven Development** — specs and tests are the source of truth, code adapts to them.

### The Cycle

```
┌─────────┐     ┌──────────┐     ┌────────────┐     ┌──────────┐     ┌──────┐
│  1.Spec │────▶│ 2.Tests  │────▶│ 3.Implement│────▶│ 4.Review │────▶│5.Docs│
└─────────┘     └──────────┘     └────────────┘     └──────────┘     └──────┘
 @spec-writer    generate-tests-   Write code        @reviewer        update-docs
 or              from-spec skill   to pass tests     agent            skill
 generate-spec                     (never edit                        (if API/arch
 skill                             tests)                             changed)
```

### Workflow by Change Type

For new features:

1. Use `requirements-interview` to clarify the product and technical shape.
2. Write or update a PRD.
3. Use `@spec-writer` or `generate-spec` to create testable specs.
4. Use `task-breakdown` to produce vertical tasks in `tasks.md`.
5. Use `vertical-tdd` or `generate-tests-from-spec` depending on whether you are implementing now or only generating tests.
6. Use `@reviewer` for spec-driven review.
7. Use `update-docs` when public API, architecture, setup, or configuration changed.

For existing features:

1. Use `requirements-interview` to decide whether the product flow changes.
2. If the product flow changes, update the PRD before writing specs.
3. If the product flow does not change, write targeted specs and continue through tasks, tests, implementation, review, and docs.

For bugs:

1. Use `bugfix-loop` to document the bug and build a reproduction loop.
2. Write a regression test at the correct boundary.
3. Fix the code and rerun the original reproduction loop.
4. Write a postmortem only when the bug was an outage.

### Quick Start — `/sdd` Prompt

The fastest way to run the full cycle is the `/sdd` global prompt. In Copilot Chat:

```
/sdd Add user authentication with JWT
```

The prompt orchestrates each step in order with explicit gates — it won't advance without your approval.

### Step by Step (Manual)

You can also run each step individually:

#### 1. Write a Spec

```
@spec-writer I need a feature for user authentication with JWT
```

The agent interviews you (one question at a time), then produces a structured spec in `specs/`. For large features, it auto-detects scope and proposes vertical slices.

If you already know the requirements and don't need an interview:

```
Use the generate-spec skill to create a spec for JWT auth
```

#### 2. Generate Tests

```
Use the generate-tests-from-spec skill for specs/user-auth.md
```

Each acceptance criterion from the spec becomes at least one test. Tests must fail initially (red phase).

#### 3. Implement

Write code to make the tests pass. The golden rule: **never modify the tests**. If a test seems wrong, revisit the spec first.

#### 4. Review

```
@reviewer review the user-auth implementation
```

The reviewer checks alignment between spec → tests → code. It verifies every acceptance criterion has a test, every test has matching code, and flags architecture issues using deep-module heuristics.

For quick reviews without specs, use the `/review` prompt instead.

#### 5. Update Docs

```
Use the update-docs skill for the user-auth feature
```

Only needed when public API, architecture, or setup changed. Skip for internal-only changes.

### Architecture Decisions

When the feature involves significant design choices, use these before or alongside the SDD cycle:

| Need | Tool | Output |
|------|------|--------|
| Explore architecture opportunities | `@architect` agent | ADR or implementation plan with evaluated designs |
| Record a quick decision | `generate-adr` skill | ADR in `docs/decisions/` (MADR 4.0) |

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

See [CONTRIBUTING.md](CONTRIBUTING.md) for full guidelines.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

## License

MIT
