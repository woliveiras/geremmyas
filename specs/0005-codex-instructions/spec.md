---
spec: "0005"
title: Codex global instruction distribution
family: multi-assistant
phase: 3
status: Implemented
owner: ""
depends_on: ["0001"]
origin: Direct user request (Codex not loading instructions)
---

# Spec: Codex global instruction distribution

## Context & Motivation

On a Codex-only machine, `geremmyas global` does not make the catalog's
instructions reach Codex, and the generated Codex `AGENTS.md` is written to a
path Codex never reads.

Two concrete faults were confirmed on the maintainer's machine:

1. **Wrong global path.** `geremmyas global --targets codex` writes
   `~/.config/codex/AGENTS.md`, but Codex.app reads `~/.codex/AGENTS.md`
   (`CODEX_HOME=~/.codex`, per its `config.toml`). The generated file is never
   read; Codex loads a small hand-written `~/.codex/AGENTS.md` instead.
2. **Instructions never reach Codex.** The `*.instructions.md` files are copied
   only to `~/.copilot/instructions/` and only when the `copilot` target is
   active. That folder is Copilot-only; Codex has no `applyTo` instruction
   folder. Even the generated Codex `AGENTS.md` drops instruction content
   entirely (`buildIDEAgentsDoc` collects `artifacts.instructions` but never
   writes them).

Skills already work because Codex scans `~/.agents/skills/`. Instructions need a
Codex-visible channel. The chosen approach (Option A) is **on-demand
reference**: mirror the instruction files into a Codex-owned directory and have
`AGENTS.md` index them by `applyTo` glob, matching the existing skills pattern
("read the file when the task matches"). This keeps `AGENTS.md` lean and
preserves per-file-type scoping, unlike inlining every instruction.

The maintainer authors everything in geremmyas and only runs `geremmyas` on the
target machine, so the generated `~/.codex/AGENTS.md` may be overwritten freely
(no hand-written rules to preserve).

## Requirements

### Functional

- [ ] Fix the Codex global `AGENTS.md` path from `~/.config/codex/AGENTS.md` to
      `$HOME/.codex/AGENTS.md` (Codex's `CODEX_HOME`) in `global_paths.go` and
      the Codex generator.
- [ ] `geremmyas global` copies the catalog instructions unconditionally to
      `~/.copilot/instructions/` (Copilot) regardless of which target is
      selected.
- [ ] `geremmyas global --targets codex` mirrors the same instruction files
      into a Codex-owned directory `$HOME/.codex/instructions/`.
- [ ] The generated Codex `AGENTS.md` adds an "Instructions" section that lists
      each instruction file with its `applyTo` glob and an on-demand pointer to
      the Codex-owned path (`~/.codex/instructions/<file>` for global scope,
      `.github/instructions/<file>` for project scope).
- [ ] The Codex documents never reference `~/.copilot/instructions/`.
- [ ] The project-scope Codex `AGENTS.md` references `.github/instructions/`
      (already populated by `sync`), so no extra project copy is needed.

### Non-Functional

- [ ] No new third-party dependencies.
- [ ] The generated `~/.codex/AGENTS.md` is overwritten by default; `--force`
      semantics stay unchanged for customized files.
- [ ] Re-running `global` is idempotent (same output, no duplication).
- [ ] The stale `~/.config/codex/AGENTS.md` is no longer produced.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Generator content and path selection | Global Codex doc path is `.codex/AGENTS.md`; the doc contains an Instructions section listing each `applyTo` and the Codex-owned pointer; no `~/.copilot` reference |
| **integration** | End-to-end global install into a temp `$HOME` | `global --targets codex` writes `<home>/.codex/AGENTS.md` and `<home>/.codex/instructions/*.instructions.md`; `~/.copilot/instructions/` also populated |

Default: **unit** for generator/path, plus **integration** for the global
install layout.

## Acceptance Criteria

- [ ] Given `geremmyas global --targets codex`, when it runs, then the Codex
      `AGENTS.md` is written to `$HOME/.codex/AGENTS.md` and not to
      `$HOME/.config/codex/AGENTS.md`.
- [ ] Given the generated global Codex `AGENTS.md`, when inspected, then it
      contains an Instructions section that lists every pack instruction with
      its `applyTo` glob and a pointer to `~/.codex/instructions/<file>`.
- [ ] Given `geremmyas global --targets codex`, when it runs, then every pack
      instruction file exists under `$HOME/.codex/instructions/`.
- [ ] Given any `geremmyas global` run, when it completes, then the catalog
      instructions are present under `~/.copilot/instructions/` regardless of
      the selected target.
- [ ] Given the generated Codex documents, when inspected, then they never
      reference `~/.copilot/instructions/`.
- [ ] Given a project-scope Codex `AGENTS.md`, when inspected, then its
      Instructions section points to `.github/instructions/<file>`.

## Edge Cases

- An instruction file without an `applyTo` glob: list it under an "all files"
  or unscoped label rather than omitting it.
- A pack with no instructions: the Codex doc omits the Instructions section
  (same as the existing skills/agents sections).
- Re-running `global` over an existing `~/.codex/instructions/`: files are
  overwritten in place, no stale duplicates introduced by this change.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Codex home path | `$HOME/.codex` | Codex.app sets `CODEX_HOME=~/.codex`; `~/.config/codex` is never read by Codex |
| Instruction delivery | On-demand reference (Option A) | Keeps `AGENTS.md` lean and preserves `applyTo` scoping; inlining all instructions would be always-on and heavy |
| Codex-owned mirror | `$HOME/.codex/instructions/` | Codex must not read the Copilot-only `~/.copilot/instructions/` path |
| Copy scope | Instructions copied to both Copilot and Codex locations | Per user request, independent of which assistant is installed |
| Overwrite `~/.codex/AGENTS.md` | Yes | Maintainer authors everything in geremmyas; no hand-written rules to preserve |

## Out of Scope

- Inlining full instruction bodies into `AGENTS.md` (Option B).
- Changing how skills are distributed or discovered.
- Any `applyTo` evaluation engine inside Codex; scoping is advisory text the
  model follows when reading files.
- Merging with a pre-existing hand-written `~/.codex/AGENTS.md`.
