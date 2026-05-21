# copilot-configs

[![CI](https://github.com/woliveiras/copilot-configs/actions/workflows/ci.yml/badge.svg)](https://github.com/woliveiras/copilot-configs/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![ShellCheck](https://img.shields.io/badge/shell-ShellCheck-brightgreen)](https://www.shellcheck.net/)

**Dotfiles for AI** — opinionated GitHub Copilot configurations for VS Code. Agents, skills, instructions, hooks, and prompts. One command to install.

## Why?

Every project needs the same Copilot setup: language-specific instructions, code review prompts, testing conventions, security guardrails. Instead of copying files between repos, install once and apply everywhere.

**What you get:**

- **Instruction files** auto-applied by file glob for languages, frameworks, testing, and security
- **AGENTS.md** project contract for agent workflows, artifact locations, and operating rules
- **4 agents** for spec-driven development: write specs → generate tests → implement code → update docs
- **Workflow and utility skills** for specs, tests, docs, migrations, ADRs, state management patterns, and commit messages
- **Command guardrails** that block `git push --force`, `rm -rf /`, `terraform destroy`, and other dangerous commands
- **4 global prompts** for code review, refactoring, test generation, and SDD workflow

```
copilot-configs/
├── install.sh / uninstall.sh          # Install scripts
├── user/
│   ├── copilot-instructions.md        # Global bootstrap that points agents to local AGENTS.md
│   └── prompts/                       # Global prompts (review, refactor, test)
└── project/
    ├── AGENTS.md                      # Project-level operating contract for agents
    ├── mise.toml                      # Tool version management template
    └── .github/
        ├── copilot-instructions.md    # Project-level Copilot instructions
        ├── instructions/              # Auto-applied language, framework, and convention files
        ├── agents/                    # 4 agents + 6 design references
        │   └── references/            # Deep modules, interface design, etc.
        ├── skills/                    # Workflow and utility skills with asset templates
        └── hooks/                     # Command guardrails (BLOCK/ASK rules)
```

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/woliveiras/copilot-configs/main/install.sh | bash
```

This clones the repo to `~/.copilot-configs/` and copies global prompts plus a
small Copilot instruction bootstrap to your VS Code user directory.

## Update

Run the same install command again:

```bash
curl -fsSL https://raw.githubusercontent.com/woliveiras/copilot-configs/main/install.sh | bash
```

The installer pulls the latest version and updates your global prompts automatically. Files that are identical to the latest version are silently skipped.

For project-level files, re-run `--project` inside your repo:

```bash
~/.copilot-configs/install.sh --project
```

Managed files (instructions, agents, skills, hooks) are updated to the latest version. Customizable files (`AGENTS.md`, `copilot-instructions.md`, `guardrails-rules.txt`) are preserved. Use `--force` to overwrite everything, including customized files.

## Uninstall

```bash
~/.copilot-configs/uninstall.sh
```

Removes the `~/.copilot-configs/` directory and global prompts from VS Code. Project-level files (`.github/` in your repos) are not touched — remove them manually if needed.

## Project Setup

After installing, apply project-level configs to any repository:

```bash
~/.copilot-configs/install.sh --project
```

To also configure `copilot-instructions.md` placeholders interactively:

```bash
~/.copilot-configs/install.sh --project --configure
```

The `--configure` flag offers two modes:
- **Interactive** — answer questions about your project
- **Auto-detect** — detects stack, directories and build commands from manifest files

You can also run `--configure` standalone on an existing project to reconfigure:

```bash
~/.copilot-configs/install.sh --configure
```

Run from inside your project directory.

## What's Included

### Global (user-level)

Prompts are installed to `~/Library/Application Support/Code/User/prompts/`
(macOS) or `~/.config/Code/User/prompts/` (Linux).

The global `copilot-instructions.md` bootstrap is installed to the VS Code user
directory. It is intentionally small: it tells agents to look for `AGENTS.md` in
the active workspace and follow the project-local file when present.
Depending on your Copilot client/settings, user-level instruction files may need
to be referenced from the VS Code custom instructions settings. The project-local
`AGENTS.md` remains the reliable contract installed into each repository.

Skills can be installed globally as reusable capabilities, but `AGENTS.md`
should be project-level by default because it contains repository-specific
paths, workflows, commands, and artifact locations. Any global `AGENTS.md` or
global instruction should be treated as a template/default, not as the active
contract for every repository.

| File | Purpose |
|------|---------|
| `copilot-instructions.md` | Global bootstrap: find and follow local `AGENTS.md` |
| `review.prompt.md` | Structured code review checklist |
| `refactor.prompt.md` | Refactor preserving behavior |
| `test.prompt.md` | Generate unit tests matching project patterns |
| `sdd.prompt.md` | Full SDD cycle: spec → test → implement → review → docs |

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
- `productivity`: handoff, skill authoring, concise communication, alignment before execution.
- `personal`: workflows specific to your setup that should not install by default.
- `utils`: rare tools, migrations, setup helpers, and guardrails.

| Skill | Purpose |
|-------|---------|
| `requirements-interview` | Explore code and clarify requirements before PRD/spec work |
| `generate-spec` | Fill a structured spec from direct input (no interview) |
| `task-breakdown` | Convert PRD, specs, or `plan.md` into vertical tasks |
| `afk-task-triage` | Classify local tasks as AFK/HITL and prepare agent-ready work |
| `agent-brief` | Create durable local AFK briefs from specs, plans, and tasks |
| `generate-tests-from-spec` | Generate unit tests from a spec's acceptance criteria |
| `vertical-tdd` | Implement one behavior at a time with red-green-refactor |
| `bugfix-loop` | Reproduce, diagnose, regression-test, and fix bugs |
| `update-docs` | Update `docs/` after implementing a feature |
| `git-commit` | Review staged changes and create a Conventional Commit with confirmation |
| `generate-glossary` | Extract domain terminology into `GLOSSARY.md` |
| `generate-adr` | Record an Architectural Decision in MADR 4.0 format |
| `session-handoff` | Create a concise handoff for another session or agent |
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

#### Local AFK Workflow

AFK delegation is local-first. Use specs, `plan.md`, and `tasks.md` as the
source of truth instead of GitHub Issues.

1. Use `task-breakdown` to create vertical tasks with `type`, `blocked-by`,
   acceptance criteria, verification, and an optional brief path.
2. Use `afk-task-triage` to split broad tasks and classify each open task as
   `AFK` or `HITL`.
3. Use `agent-brief` for each AFK-ready task. Briefs live in
   `specs/YYYY-MM-DD-<feature-slug>/agent-briefs/<task-slug>.md`, or
   `docs/agent-briefs/YYYY-MM-DD-<task-slug>.md` for repo-level work.
4. Use `session-handoff` only when transferring conversation state. Use
   `agent-brief` when preparing an implementation contract for another agent.

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

All files are meant to be edited. After running `--project`:

1. **Review `AGENTS.md`** — adjust artifact paths, workflow rules, and skill routing for the repository
2. **Edit `.github/copilot-instructions.md`** — fill in project name, description, directory structure, and build commands
3. **Add more instructions** — copy from `~/.copilot-configs/project/.github/instructions/`
4. **Tune guardrails** — add or remove rules in `guardrails-rules.txt`

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

## License

MIT
