---
spec: "0009"
title: Lazy workflow harness
family: multi-assistant
phase: 7
status: Verified
owner: ""
depends_on: ["0006", "0008"]
origin: Workflow efficiency audit and user-approved capability simplification
---

# Spec: Lazy workflow harness

## Context & Motivation

Geremmyas already materializes target-specific harness artifacts and reconciles
managed global files by manifest and hash. The remaining workflow surface is
broader than the active phase: the `sdd` pack advertises ten skills and ten
custom agents, globally installed artifacts are not directly listable or safely
clearable, and documentation workflows repeat routing and writing rules across
separate skills.

This feature makes the default workflow project-local and phase-focused. It adds
explicit global inventory and clearing operations, renames the public workflow
skills, consolidates documentation behind lazy references, removes bundled
custom agents, and keeps stack-specific packs separate from a technology-neutral
`base` pack.

`GLOSSARY.md` is absent at specification time. The feature therefore uses the
terminology already established by specs 0006 through 0008 and the current CLI.

## Requirements

### Functional

- [x] Add `geremmyas global list` with stable human-readable and JSON output.
- [x] Classify global artifacts as managed-current, managed-modified,
      managed-missing, obsolete, adoptable, unowned, symlink, or external.
- [x] Report Geremmyas packs, targets, manifest version, paths, hashes, and
      ownership without treating plugin caches as installed or managed content.
- [x] Add `geremmyas global clear --dry-run` that computes and reports the exact
      removal and preservation plan without filesystem or manifest writes.
- [x] Add `geremmyas global clear` with optional `--targets`,
      `--include-adoptable`, and `--force` behavior.
- [x] Reuse the existing manifest, hashing, path validation, and reconciliation
      logic rather than creating an independent deletion implementation.
- [x] Preserve modified, unowned, external, and symlinked paths by default.
      `--force` may affect only files whose Geremmyas ownership is proven.
- [x] Fail closed on corrupt or incompatible manifests. A missing manifest may
      inventory known files but may not silently claim ambiguous ownership.
- [x] Replace the public workflow names with `refine`, `spec`, `tdd`, `bugfix`,
      `verify`, `docs`, and `git-commit`.
- [x] Add `coding` (`refine`, `spec`, `tdd`, `bugfix`), `quality` (`verify`,
      `docs`, `git-commit`), and `base` (`coding` plus `quality`) packs.
- [x] Keep `base` technology-neutral. It must not select stack packs, plugins,
      MCP servers, or custom agents.
- [x] Make project-local `coding` the recommended efficient starting point and
      avoid advertising closing skills before their phase.
- [x] Consolidate update/generation of project documentation, glossary,
      ADR/MADR, and RFC artifacts in `docs`. Its main body must route to exactly
      the relevant reference and load a template only when one is needed.
- [x] Keep PRD, spec, plan, and tasks behavior under `spec`, not `docs`.
- [x] Remove all bundled custom-agent profiles and their catalog artifact while
      retaining generic target-adapter support for future agents that provide a
      concrete permission, model, tool, MCP, or isolation benefit.
- [x] Preserve independent review as a concise, non-discoverable delegation
      contract loaded from `verify` and executed by a runtime-created subagent
      when supported.
- [x] Update canonical contracts, prompts, workflow gates, target adapters,
      documentation, tests, and generated-path expectations to the new names.
- [x] Reconcile old managed skill and agent destinations on the next project or
      global materialization without deleting modified or unowned legacy files.
- [x] Extend context diagnostics or structured output enough to compare no
      Geremmyas state, global state, project `coding`, and project `base` in
      isolated temporary homes.

### Non-Functional

- [x] No new third-party dependencies.
- [x] Global clearing never follows symlinks and never deletes plugin or runtime
      cache content.
- [x] Existing `geremmyas global [flags] <pack>...` syntax remains supported.
- [x] Human and JSON inventory report the same state deterministically.
- [x] Skill directories and frontmatter names match after migration.
- [x] No duplicate compatibility skill directories are materialized because
      aliases would restore discovery cost and trigger ambiguity.
- [x] `content/AGENTS.md` remains within its 700-word budget.
- [x] Tests use temporary `HOME` and `XDG_STATE_HOME`; development never clears
      the real user profile.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Parsing, classification, pack resolution, routing, rendering | global subcommand parsing, ownership states, renamed pack contents, docs router invariants |
| **integration** | Manifest and filesystem behavior across targets | dry-run immutability, target-scoped clear, modified/unowned/symlink preservation, old-path reconciliation |
| **both** | Behavior combines policy with filesystem materialization | adoptable legacy clearing, five-target output migration, isolated context baselines |

Default: unit tests for pure inventory and catalog logic, plus temporary-home
integration tests for every clearing mode and materialization migration.

## Acceptance Criteria

- [x] Given a valid global manifest, when `global list --json` runs, then every
      managed path has a deterministic ownership state and the JSON agrees with
      the human report.
- [x] Given any global state, when `global clear --dry-run` runs, then its output
      describes removals and preservations and no file or manifest byte changes.
- [x] Given intact managed artifacts for multiple targets, when one target is
      cleared, then only that target's eligible paths are removed and the
      remaining manifest state is valid.
- [x] Given modified, unowned, plugin-cache, or symlink paths, when clear runs,
      then those paths are preserved and reported with the reason.
- [x] Given a corrupt or incompatible manifest, when list or clear runs, then the
      command fails before deletion and explains the recovery boundary.
- [x] Given `coding`, when packs resolve, then exactly `refine`, `spec`, `tdd`,
      and `bugfix` are publicly discoverable workflow skills.
- [x] Given `base`, when packs resolve, then exactly seven technology-neutral
      workflow skills are discoverable and no agent or stack-specific artifact
      is selected.
- [x] Given a documentation request, when `docs` runs, then its compact router
      selects project docs, glossary, ADR/MADR, or RFC support without requiring
      another public skill.
- [x] Given project or global materialization after an older managed install,
      when reconciliation completes, then eligible old skill/agent paths are
      removed and modified or unowned paths are preserved.
- [x] Given any supported target, when `base` is materialized, then it receives
      no bundled custom-agent profile and its primary-agent contract describes
      bounded runtime delegation instead.
- [x] Given a simple request, when routing policy is inspected, then it prescribes
      zero workflow skills; given material ambiguity, it prescribes at most
      `refine` before clarification.
- [x] Given isolated temporary homes, when context baselines are collected, then
      no-state, global, `coding`, and `base` results are distinguishable without
      reading or mutating the real home directory.

## Edge Cases

- A pack is named `list` or `clear`: reserve those names as global subcommands
  and reject catalog conflicts in lint or catalog validation.
- `--targets` contains a target absent from the manifest: report no eligible
  managed files and leave other state unchanged.
- A managed file changes between planning and deletion: recheck its hash at the
  mutation boundary and preserve it unless the explicit proven-ownership force
  rule applies.
- A historical manifest names a pack removed by this migration: inventory its
  owned paths, report unresolved catalog state, and do not infer obsolescence
  from an incomplete desired-state calculation. Clearing still relies on proven
  manifest ownership rather than the current catalog.
- A parent directory or destination component becomes a symlink: stop before
  traversal and report the protected path.
- An adoptable generated file was customized: preserve it unless explicit force
  semantics prove both origin and requested scope.
- An old skill directory is unowned because it predates manifests: report it as
  adoptable only for an exact canonical hash or trusted generated marker;
  otherwise leave it unowned.
- A runtime lacks subagents: perform the review inline and report that
  independent context isolation was unavailable.
- A user explicitly invokes an old skill name: migration documentation explains
  the replacement; duplicate compatibility skills are not installed.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Global cleanup interface | `global list` and `global clear` subcommands | Explicit operations are clearer and extensible while preserving the existing desired-state command |
| Clear authority | Ownership plus current hash by default | Reuses the existing conservative safety model |
| External plugins | Observe only | Geremmyas does not install or own runtime plugin state |
| Workflow packs | `coding`, `quality`, and technology-neutral `base` | Separates active coding phases from closing procedures and composes naturally with stack packs |
| Documentation | One public `docs` router with lazy references | Reduces discovery and duplicated writing policy without deleting specialized guidance |
| Review | Reference contract plus runtime-created subagent | Keeps independent review without permanent agent metadata |
| Bundled agents | Remove all current profiles; retain adapter capability | Current profiles mostly encode personas that capable runtimes can create on demand |
| Rename compatibility | Migrate canonical and managed paths without duplicate aliases | Physical aliases would preserve the context and trigger cost being removed |
| Dynamic latency benchmark | External runtime protocol, not an LLM client in Geremmyas | Keeps Geremmyas focused on installing and measuring the harness |

## Out of Scope

- Installing, uninstalling, disabling, or clearing plugins, MCP servers, apps,
  connectors, runtime caches, or assistant-managed system skills.
- Adding an LLM client or runtime orchestration engine to Geremmyas.
- Guaranteeing that every assistant obeys textual lazy-routing policy.
- Removing generic CLI support for future custom-agent artifacts.
- Exact model tokenizer accounting or a cross-provider latency guarantee.
- Technology-specific pack redesign beyond references required by renamed
  workflow capabilities.
