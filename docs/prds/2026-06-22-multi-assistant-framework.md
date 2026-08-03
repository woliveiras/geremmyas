# PRD: geremmyas as a personal multi-assistant framework

- Date: 2026-06-22
- Status: Draft
- Owner: woliveiras

## Summary

geremmyas distributes AI-assistant configuration (instructions, skills, agents,
hooks, templates) from assistant-neutral canonical sources under `content/`,
with assistant-specific adapters under `targets/`. Target generators materialize
the appropriate project and global artifacts for each supported assistant.

This PRD frames the evolution of geremmyas from a Copilot-focused config
distributor into a **personal multi-assistant working framework** that supports
the assistants used day to day: Copilot at work, Codex on personal projects, and
Cursor, with Claude as a future target. The goal is to make the same canonical
content usable across these assistants without losing the strengths that already
exist.

## Problem

The author uses multiple AI assistants but the framework only treats one of them
(Copilot) as a first-class target. Concretely:

1. **Codex is not a target**, despite daily use on personal projects. There is no
   generation path that makes synced skills and the `AGENTS.md` contract usable
   from Codex.
2. **Skill description quality is unchecked.** Skill discovery depends on good
   `description` metadata (clear "use when" triggers and negative scope). Nothing
   enforces this, so weak descriptions degrade discovery, especially on
   assistants that rely on a markdown skill index instead of directory scanning.
3. **Global installs accumulate context indefinitely.** Re-running
   `geremmyas global` copies the newly selected packs but does not remove files
   from earlier selections. Generated assistant documents can therefore describe
   one pack set while the globally scanned skill directory contains the union of
   every historical install.
4. **Workflow policy is duplicated across layers.** The global contract embeds
   the project contract, while guardrails and orchestration steps are advertised
   again as independently discoverable skills. This consumes context and creates
   conflicting trigger timing.
5. **Workflow approval gates interrupt routine delivery.** Feature, bugfix,
   architecture, commit, and tool-selection workflows repeatedly stop for human
   confirmation even when repository evidence, automated tests, and specialist
   review can resolve the decision. Subagents are restricted to read-heavy work,
   so the primary agent cannot use them as an autonomous engineering team.
6. **Workflow discovery is still broader than the active phase.** The default
   SDD pack advertises planning, implementation, verification, documentation,
   architecture, Git, and custom-agent roles at the same time. Short requests
   therefore pay discovery and routing costs for capabilities that are not yet
   relevant.
7. **Global state is safe to reconcile but hard to inspect or empty.** The
   ownership manifest supports desired-state updates, but users lack explicit
   inventory and clear operations that distinguish managed, modified,
   adoptable, unowned, symlinked, and external artifacts.

## Goals

- Make Codex a supported generation target, in parity with the existing OpenCode
  target (an `AGENTS.md`-style document plus an on-demand skill index), for both
  project scope and global scope.
- Add an automated quality check (`geremmyas lint`) for skill description
  metadata, runnable locally and in CI.
- Make global installation declarative and ownership-aware so the selected packs
  are the desired state, while preserving user-modified and external files.
- Keep generated assistant context target-aware: do not repeat contracts or
  native skill indexes on assistants that already discover them.
- Make the canonical content model assistant-neutral. Assistant-native paths
  such as `.github/` and `.codex/` are output destinations owned by target
  adapters, not source taxonomy.
- Make project sync honor the selected targets so a clean Codex-only install
  does not receive Copilot instructions, agents, or hooks.
- Provide context diagnostics and enforce catalog budgets so context growth is
  visible before release.
- Preserve the existing strengths: single canonical source with per-IDE
  generation, portable `AGENTS.md` contract, pack model with dependencies, shell
  safety guardrails, test-first delivery, and fresh verification evidence.
- Make local engineering autonomous by default. Agents create and maintain the
  durable work artifacts, implement, test, review, document, and create atomic
  Conventional Commits unless the user explicitly disables commits.
- Allow specialist subagents to explore, design, implement, test, audit, and
  review independently when ownership is partitioned and Git integration remains
  centralized.
- Reserve human escalation for material product ambiguity, unmitigated risk to
  production compatibility, new uncatalogued dependencies, persistent blockers,
  push, and production mutations or publication.
- Make the technology-neutral workflow available as a `base` pack composed from
  phase-focused `coding` and `quality` packs. Keep stack-specific capabilities
  opt-in and project-local.
- Keep the public workflow vocabulary short and phase-specific: `refine`,
  `spec`, `tdd`, `bugfix`, `verify`, `docs`, and `git-commit`.
- Consolidate project documentation, glossary, ADR/MADR, and RFC procedures
  behind the single `docs` skill with support files loaded on demand.
- Prefer runtime-created subagents with bounded delegation contracts. Do not
  distribute custom agents by default when they add only a persona rather than
  enforceable tools, permissions, models, or isolated context.
- Provide explicit global inventory and safe clearing commands without taking
  ownership of plugins, caches, symlinks, or third-party files.

## Non-Goals (this PRD)

- **Auto-trigger bootstrap** (a session-start mechanism that injects a "consult
  skills before acting" instruction). Deferred. The Codex skill index in
  `AGENTS.md` is the lightweight discovery path for now.
- **Claude as a full plugin** with automatic skill triggering. Deferred; the
  current generated `CLAUDE.md` index stays.
- **Content integrity via content hashing and a registry.** Discarded for
  personal use; the `geremmyas:generated` marker-based sync is sufficient.
- **Consumer install lockfile.** Discarded.
- **Skill scaffolding generator.** Discarded.
- **Marketplace / cohesive external versioning / monorepo packaging.** Discarded;
  keep the single Go binary plus packs.
- **Broadening to many assistants beyond those actually used.** Map only Codex,
  Copilot, Cursor, Claude, plus the existing OpenCode target.
- **Release workflow redesign.** Environment protection and automatic release
  behavior in `.github/workflows/` will be reviewed after the autonomous local
  workflow rollout.

## Scope decisions

| Item | Decision |
| --- | --- |
| Codex target | In scope |
| Skill description validator (`geremmyas lint`) | In scope |
| Auto-trigger bootstrap (session-start) | Deferred |
| Claude as full plugin | Deferred |
| General content registry | Discarded |
| Consumer project lockfile | Discarded |
| Global ownership manifest | In scope for safe desired-state reconciliation |
| Context diagnostics and budgets | In scope |
| Assistant-neutral canonical content | In scope |
| Target-aware project materialization | In scope |
| Autonomous local engineering workflow | In scope |
| Atomic local commits by default | In scope; explicit user opt-out |
| Proactive specialist subagents | In scope with isolated ownership |
| New uncatalogued dependencies | Require explicit user decision |
| Push and production mutation/publication | Require explicit user authorization |
| GitHub release workflow redesign | Deferred to the next rollout |
| Skill generator | Discarded |
| Marketplace / cohesive versioning | Discarded |
| Global inventory and clear | In scope with manifest and hash safety |
| Technology-neutral workflow pack | `base`, composed from `coding` and `quality` |
| Documentation capabilities | One `docs` skill with lazy references |
| Bundled custom agents | Remove from distribution; retain generic adapter support |
| Plugins and runtime caches | Observe only; never clear as Geremmyas-owned state |

## Working principles (invariants)

These constrain how features in this PRD are built and are not themselves
deliverables of this PRD:

- **Tests first.** Before changing production code, the agent verifies that tests
  cover the expected behavior; if absent, it writes them first (red), confirms
  they fail for the right reason, then changes code. Focus on well-designed unit
  and integration tests.
- **Autonomy by default.** Agents continue through specification, implementation,
  verification, documentation, review, and local commits without conversational
  approval gates. They escalate only when evidence cannot resolve a material
  decision or an authority boundary is reached.
- **Atomic history.** Each verified slice ends in one task-owned Conventional
  Commit containing its behavior, tests, and required documentation. A user can
  disable commits explicitly for a session; push is never inferred.
- **Specialist collaboration.** Subagents may work in parallel when file, module,
  or worktree ownership is explicit. The primary agent integrates findings,
  resolves review feedback, and owns Git operations.
- **Authority boundaries.** Local and disposable test environments are available
  to the harness. Push and every production mutation or publication require
  explicit authorization. Dangerous commands outside the task are denied rather
  than converted into routine confirmation prompts.
- **Dependency provenance.** Existing project or Geremmyas-catalogued
  dependencies may be used autonomously. Introducing a new uncatalogued direct
  dependency requires the user to choose between adopting it and building a
  local module after provenance, maintenance, security, and license review.
- **Single canonical source.** `content/` is the shared source of truth and
  `targets/<assistant>/` contains assistant-specific adapters. Generated consumer
  paths are outputs, not an alternate source taxonomy.
- **Platform.** macOS and Linux only.

## Success criteria

- `geremmyas init/sync/global --targets codex` produces a Codex document plus a
  skill index in the correct locations (project and global), idempotently, behind
  the `geremmyas:generated` marker.
- `geremmyas lint` flags skills whose descriptions lack "use when" triggers or
  negative scope, exceed the description length limit, contain disallowed markup,
  whose `name` does not match the folder, or whose body exceeds the line limit;
  it passes clean skills and is wired into CI.
- Re-running `geremmyas global` with a smaller pack set removes only unchanged
  files previously recorded as Geremmyas-owned and reports modified or unowned
  leftovers without deleting them.
- Codex receives a compact global bootstrap instead of a duplicate project
  contract or duplicate native skill catalog.
- `geremmyas context` reports global, project, system, and plugin skill counts,
  nested skill files, ownership state, and approximate context cost.
- The default SDD catalog stays within explicit skill-count and metadata budgets.
- Feature and bugfix workflows proceed from durable artifacts to verified local
  delivery without waiting for routine human approval.
- Local commits are atomic, use Conventional Commits, match their staged diff,
  and are created by default unless the user opts out. Push remains separate.
- Specialist subagents can perform partitioned implementation and independent
  review; repeated findings are repaired and re-reviewed before completion.
- A simple request begins with no workflow skill, and an ambiguous request begins
  with at most `refine`; closing skills are not selected before their phase.
- `geremmyas global list` reports owned and observed global state in stable human
  and JSON formats, and `geremmyas global clear --dry-run` reports the exact safe
  removal plan without writing.
- `geremmyas global clear` removes only eligible Geremmyas-owned artifacts for
  the selected targets and preserves modified, unowned, external, and symlinked
  files unless a narrower explicit safe mode applies.
- The `base` pack resolves the complete technology-neutral workflow without
  including stack packs, plugins, MCP servers, or custom agents.
- The public workflow surface contains seven skills and no bundled custom-agent
  profiles. Review and other isolated work use concise delegation contracts.
- Automated gates retain red/green tests, nearby suites, evidence capture,
  documentation reconciliation, and residual-risk reporting.
- New uncatalogued dependencies, push, and production mutations or publication
  stop at an explicit authority boundary.
- No regression in existing targets, packs, preservation rules, or verification
  gates.
- A clean install materializes only shared artifacts and the outputs required by
  the selected targets. Existing modified or unowned files are never removed by
  migration.

## Linked specs

- `specs/0001-codex-target/` — Codex generation target.
- `specs/0002-skill-validator/` — `geremmyas lint` skill description validator.
- `specs/0006-context-efficient-workflows/` — managed global state, compact
  target output, skill consolidation, diagnostics, budgets, and agent contracts.
- `specs/0007-assistant-neutral-content/` — neutral canonical content, typed
  artifacts, and target-aware project/global materialization.
- `specs/0008-autonomous-agent-workflows/` — autonomy-by-default workflows,
  specialist delegation, atomic local commits, and contextual authority gates.
- `specs/0009-lazy-workflow-harness/` — global inventory and clearing,
  phase-focused workflow packs, lazy documentation references, and dynamic
  subagent contracts.
