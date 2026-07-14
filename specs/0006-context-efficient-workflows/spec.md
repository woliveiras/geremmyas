---
spec: "0006"
title: Context-efficient agent workflows
family: multi-assistant
phase: 4
status: Implemented
owner: ""
depends_on: ["0004", "0005"]
origin: Audit of global skills, agents, and context usage
---

# Spec: Context-efficient agent workflows

## Context & Motivation

Geremmyas already tiers skills into core and stack packs, but installed state is
append-only: selecting fewer packs does not remove earlier files. On the audited
machine, `~/.agents/skills` contained 61 `SKILL.md` files, including stale nested
asset markers, while the current core requires 19 top-level skills. The generated
global Codex contract also embeds the complete project contract, so the same
1,400-word policy is loaded globally and locally.

The SDD pack compounds this cost by exposing policy and internal composition as
independent skills. For example, approval, regression, rationalization, abort,
verification, and subagent-selection rules overlap with `AGENTS.md` and with
orchestrating workflows. The result is a large discovery catalog and ambiguous
trigger timing.

This feature makes global installation declarative, generates target-aware
context, reduces the discoverable SDD surface, adds context diagnostics and
budgets, and gives subagents bounded contracts.

## Requirements

### Functional

- [x] Treat every `geremmyas global [--targets ...] <pack>...` invocation as the
      complete desired state for Geremmyas-managed global artifacts.
- [x] Persist an ownership manifest under
      `${XDG_STATE_HOME:-$HOME/.local/state}/geremmyas/` containing the selected
      packs, targets, destination paths, and installed content hashes.
- [x] Remove an obsolete managed file only when its current hash still matches
      the manifest. Preserve and report modified files and files without an
      ownership record. Remove empty directories only inside managed roots.
- [x] Make initial migration conservative: adopt desired files written by the
      current run and exact current-catalog matches, but never delete unknown
      legacy or third-party files without an ownership record.
- [x] Include installed, updated, removed, preserved, and unowned counts in the
      global reconciliation summary. Manifest writes must be atomic.
- [x] Generate a compact Codex global bootstrap that does not embed the project
      `AGENTS.md`, duplicate Codex's native skill catalog, or advertise unavailable
      agent pickers. Preserve the on-demand Codex instruction index.
- [x] Keep Claude Code and OpenCode generation behavior compatible where those
      targets still need an embedded contract or skill index.
- [x] Reduce the default `sdd` pack to at most 10 discoverable workflow skills.
      Move internal policies, examples, and composed procedures to references or
      concise `AGENTS.md` invariants.
- [x] Keep explicit user capabilities discoverable, including feature
      requirements/specification, bugfix, implementation, glossary, ADR, review,
      verification, and commit workflows. Move `skill-authoring` and general
      decision support to opt-in stack packs.
- [x] Rewrite the project `AGENTS.md` as a concise contract with phase-aware
      routing. It must not instruct agents to load completion-only skills early.
- [x] Add `geremmyas context` with a stable human-readable report for catalog,
      project, global, Codex system, and Codex plugin skill roots when present.
      Report top-level and nested `SKILL.md` counts, ownership state, frontmatter
      bytes, approximate tokens, and generated contract size.
- [x] Make unavailable roots non-fatal and label external plugin/system content
      as observed but not managed by Geremmyas.
- [x] Extend `geremmyas lint` to reject nested `SKILL.md` files, descriptions
      over 240 characters, skill bodies over 250 lines, an SDD pack with more
      than 10 discoverable skills, or an `AGENTS.md` contract over 700 words.
- [x] Update bundled skills to satisfy those budgets by moving long examples and
      recipes to references instead of deleting useful guidance.
- [x] Give `explorer`, `spec-writer`, `reviewer`, and `architect` explicit scope,
      evidence, unknowns, and concise-output contracts. Architecture fan-out to
      three subagents must be conditional on a material interface decision.

### Non-Functional

- [x] No new third-party Go dependencies.
- [x] Global reconciliation must never delete an unowned or modified file.
- [x] A failed install or manifest write must leave the previous manifest valid.
- [x] Context estimates are clearly labeled approximate and use a deterministic
      byte-based formula rather than a model-specific tokenizer.
- [x] Existing project sync preservation behavior remains unchanged.
- [x] Each improvement is delivered as a separate Conventional Commit, with its
      tests and required documentation included in the same commit.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Manifest parsing, hashing, pruning decisions, budgets, report aggregation, target rendering | Modified files preserved; nested skill rejected; Codex bootstrap omits project contract |
| **integration** | CLI behavior across repeated global runs and real filesystem trees | Install `sdd` plus a stack pack, rerun with `sdd`, assert only unchanged managed stack files are removed |
| **both** | Changes combine decision logic with filesystem or generated output | Atomic manifest reconciliation and `context` report over a temporary home/project |

Default: unit tests for pure policy and rendering, plus integration tests for
global reconciliation and context-root discovery.

## Acceptance Criteria

- [x] Given a managed global install containing `sdd` and `python-ai`, when
      `global sdd` runs, then unchanged Python AI artifacts are removed and the
      manifest records only the desired state.
- [x] Given an obsolete managed file modified after installation, when a smaller
      desired state is applied, then the file remains and the summary reports it
      as preserved.
- [x] Given an unowned skill in `~/.agents/skills`, when reconciliation runs,
      then it is never deleted and is reported as unowned by `geremmyas context`.
- [x] Given a failed manifest replacement, when reconciliation exits, then the
      previous valid manifest remains readable.
- [x] Given global and project Codex instructions, when the resulting context is
      inspected, then the global document does not contain the project contract
      or a duplicate skills section and the project contract remains authoritative.
- [x] Given the default `core` plus `sdd` resolution, when catalog artifacts are
      counted, then no more than 10 top-level skills are discoverable.
- [x] Given any bundled skill tree, when `geremmyas lint` runs, then nested
      `SKILL.md` files and budget violations fail with actionable paths and codes.
- [x] Given `geremmyas context` on a machine with Geremmyas, system, plugin, and
      project skills, when it runs, then each source is reported separately and
      only manifest-owned files are called managed.
- [x] Given a simple architecture exploration, when the architect role runs,
      then it does not require three subagents unless a material interface choice
      with multiple viable designs has been selected.
- [x] Given the completed feature, when documentation is read, then global
      desired-state semantics, migration safety, context budgets, skill taxonomy,
      and subagent guidance are documented.

## Edge Cases

- Missing or corrupt manifest: fail safely without pruning and provide a recovery
  message; installation may not silently replace an unreadable manifest.
- Duplicate desired destinations from pack dependencies: preserve first-wins
  resolution and record one owner entry.
- Symlinks in managed roots: never follow them during deletion or context scans.
- A removed target with a modified generated file: preserve and report it.
- First run after upgrade: unknown historical files remain until explicitly
  cleaned; diagnostics identify them without claiming ownership.
- Plugin caches may contain nested skills intentionally; report them, but apply
  nested-skill lint only to Geremmyas's canonical `project/.github/skills` tree.
- A context root does not exist: report zero or omit it without failing.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Global command semantics | Complete desired state | Prevents historical union growth and matches the user's explicit choice |
| Deletion authority | Manifest ownership plus unchanged hash | Makes pruning useful without risking user or third-party content |
| State location | XDG state directory with `$HOME/.local/state` fallback | Separates machine state from user-authored configuration |
| First-run migration | Conservative adoption, no deletion of unknown files | Existing installs predate ownership metadata |
| Codex generation | Target-specific compact bootstrap | Codex already loads local contracts and discovers global skills natively |
| Skill taxonomy | Capability = skill; internal procedure = reference; isolated role = agent | Reduces discovery cost and composition ambiguity |
| Context estimate | `(bytes + 3) / 4` approximate tokens | Deterministic, dependency-free, explicitly approximate |
| External plugins | Observe only | Geremmyas does not own plugin installation state |

## Out of Scope

- Removing or disabling Codex plugins.
- Automatically deleting unknown legacy skills during the first managed run.
- Changing project `sync` to desired-state pruning; this feature applies to
  user-level global installation only.
- Exact tokenizer-specific context accounting.
- Adding a general marketplace or remote catalog registry.
- Creating custom native subagent runtimes for assistants that do not expose
  them.
