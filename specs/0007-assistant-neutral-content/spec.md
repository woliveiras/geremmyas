---
spec: "0007"
title: Assistant-neutral canonical content
family: multi-assistant
phase: 5
status: Approved
owner: ""
depends_on: ["0001", "0005", "0006"]
origin: Direct user request following repository coupling audit
---

# Spec: Assistant-neutral canonical content

## Context & Motivation

Geremmyas supports Copilot, Codex, Cursor, Claude Code, and OpenCode, but its
canonical model still uses Copilot paths. Of the 91 catalog entries at the time
of this spec, 88 use `project/.github/*` as both source taxonomy and consumer
destination. The CLI also determines whether an entry is a skill, instruction,
agent, or hook by inspecting its `.github/*` target.

This makes Copilot the implicit baseline. Project sync copies pack entries
before target generators run, so a clean `--targets codex` install receives
Copilot directories. Global installation also copies instructions to
`~/.copilot/instructions` regardless of the selected target.

The canonical content must describe artifact meaning rather than one
assistant's filesystem. Target adapters will materialize assistant-native
outputs. Existing Copilot behavior remains the default for backward
compatibility.

## Requirements

### Functional

- [ ] Move canonical contracts, skills, instructions, agents, guardrail rules,
      prompts, and templates from `project/` and `user/` into a neutral
      `content/` tree.
- [ ] Keep Copilot-only project instructions and hook adapters under
      `targets/copilot/`.
- [ ] Replace `.github/*` prefix inference with explicit artifact kinds in the
      catalog model.
- [ ] Materialize shared artifacts once and assistant-native artifacts only for
      selected targets.
- [ ] For a Codex project target, keep the repository `AGENTS.md`, install
      project skills under `.agents/skills`, and make on-demand instruction
      content reachable from the generated Codex index.
- [ ] Preserve the existing Copilot destinations when `copilot` is selected:
      `.github/copilot-instructions.md`, `.github/instructions`,
      `.github/skills`, `.github/agents`, and `.github/hooks`.
- [ ] Generate the union of required outputs for multiple selected targets
      without duplicate writes or order-dependent results.
- [ ] Stop installing global Copilot instructions unless `copilot` is selected.
      Continue installing shared skills when required by any selected target and
      mirror Codex instructions only for `codex`.
- [ ] Reconcile obsolete project artifacts conservatively using ownership and
      content hashes. Remove only unchanged Geremmyas-owned files, preserve
      modified or unowned files, and never follow symlinks during cleanup.
- [ ] Migrate pre-manifest installs conservatively by adopting exact matches to
      known legacy Geremmyas content. Report preserved legacy files that cannot
      be proven safe to remove.
- [ ] Replace maintainer dogfooding symlinks that expose shared content through
      `.github/*` with target-specific links or generated adapters.
- [ ] Update CLI help, catalog descriptions, diagnostics, documentation, and
      maintainer guidance to describe coding assistants rather than Copilot as
      the framework baseline.

### Non-Functional

- [ ] Existing default configuration remains `copilot` for backward
      compatibility.
- [ ] Existing Copilot output content and overwrite protection remain
      compatible when `copilot` is selected.
- [ ] Project reconciliation must not delete modified, unowned, or symlinked
      files.
- [ ] No new third-party Go dependencies.
- [ ] macOS and Linux remain the supported platforms.
- [ ] Each large implementation slice includes its tests and documentation in
      a separate Conventional Commit.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Artifact typing, destination planning, path safety, ownership decisions | A skill maps to Codex and Copilot destinations without inspecting `.github` |
| **integration** | CLI plus embedded filesystem plus temporary project/home | Codex-only sync omits `.github`; mixed targets produce both output families |
| **both** | Migration and reconciliation combine policy with filesystem changes | Modified legacy files survive while unchanged owned artifacts are removed |

Default: unit tests for artifact and destination policy, integration tests for
project/global materialization, and both for migration safety.

## Acceptance Criteria

- [ ] Given the embedded catalog, when it is validated, then canonical sources
      use `content/` or `targets/<target>/` and every installable entry has an
      explicit artifact kind.
- [ ] Given a clean project configured only for `codex`, when `geremmyas sync`
      runs, then `AGENTS.md`, `.agents/skills`, Codex instruction/index output,
      and shared project files are present, while Copilot instructions, agents,
      and hooks are absent.
- [ ] Given a clean project configured only for `copilot`, when sync runs, then
      the current `.github/*` Copilot output contract is preserved.
- [ ] Given a project configured for `codex` and `copilot`, when sync runs, then
      both native output sets are present and shared content is written once.
- [ ] Given a global install configured only for `codex`, when it runs, then no
      files are written under `~/.copilot/instructions`.
- [ ] Given a previous target-owned file whose content still matches the
      ownership manifest, when that target is removed and sync runs, then the
      file is removed and empty managed directories are cleaned safely.
- [ ] Given a previous target-owned file modified by the user, when its target
      is removed, then the file remains and is reported as preserved.
- [ ] Given an unowned file or symlink under a managed-looking directory, when
      reconciliation runs, then it is not followed, overwritten, or deleted.
- [ ] Given an installation created before the project manifest exists, when
      migration runs, then exact known Geremmyas files may be adopted, while
      unknown or changed files remain untouched.
- [ ] Given `geremmyas context`, help, README, and architecture documentation,
      when inspected, then neutral canonical paths and target-specific outputs
      are described consistently.
- [ ] Given the completed migration, when the full verification matrix runs,
      then catalog validation, lint, unit/integration tests, build, doctor, and
      clean-install smoke tests pass.

## Edge Cases

- A pack contains a directory entry with nested assets and references.
- Two packs request the same artifact or destination through dependencies.
- A target is removed while another target still requires the same shared file.
- An old install has no ownership manifest and contains customized Copilot
  instructions or guardrail rules.
- A managed file is replaced by a symlink after installation.
- The ownership manifest is missing, corrupt, or has an unsupported version.
- A generated assistant file exists without the `geremmyas:generated` marker.
- The maintainer repository contains GitHub workflows and templates that are
  platform files, not assistant content.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Canonical root | `content/` | Content is used at project and global scope, so `project/` is misleading |
| Target-specific source | `targets/<target>/` | Assistant-native templates and adapters remain explicit without defining shared taxonomy |
| Artifact classification | Explicit catalog kind | Destination paths must not define domain meaning |
| Codex project skills | `.agents/skills` | This is Codex's documented repository skill discovery path |
| Copilot compatibility | Preserve `.github/*` outputs | Existing consumers and default behavior continue working |
| Project cleanup | Ownership manifest plus unchanged hash | Target removal becomes effective without deleting user work |
| Legacy migration | Adopt exact known matches only | Older installs have no ownership metadata |
| Root `.github` platform files | Keep workflows, templates, and assets | They belong to GitHub hosting, not Copilot configuration |
| Default target | Keep `copilot` | Avoid an unrelated breaking default change |

## Out of Scope

- Renaming historical specs, changelog entries, bugfix documents, or release
  links that accurately describe past work.
- Removing GitHub workflows, issue templates, pull request templates, or
  repository assets.
- Adding new assistant targets.
- Changing the pack dependency model or public pack names.
- Exact context tokenization or plugin management.
- Pushes, pull requests, rebases, merges, or release publication.
