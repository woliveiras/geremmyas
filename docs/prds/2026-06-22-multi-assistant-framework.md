# PRD: geremmyas as a personal multi-assistant framework

- Date: 2026-06-22
- Status: Draft
- Owner: woliveiras

## Summary

geremmyas today distributes AI-assistant configuration (instructions, skills,
agents, hooks, templates) from a single canonical source under `project/`, with
per-IDE generation. It treats GitHub Copilot as the native target and generates
derived artifacts for Cursor, Claude Code, and OpenCode.

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

## Goals

- Make Codex a supported generation target, in parity with the existing OpenCode
  target (an `AGENTS.md`-style document plus an on-demand skill index), for both
  project scope and global scope.
- Add an automated quality check (`geremmyas lint`) for skill description
  metadata, runnable locally and in CI.
- Preserve the existing strengths: single canonical source with per-IDE
  generation, portable `AGENTS.md` contract, pack model with dependencies, shell
  guardrails, and SDD approval gates.

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

## Scope decisions

| Item | Decision |
| --- | --- |
| Codex target | In scope |
| Skill description validator (`geremmyas lint`) | In scope |
| Auto-trigger bootstrap (session-start) | Deferred |
| Claude as full plugin | Deferred |
| Content hash + registry | Discarded |
| Install lockfile | Discarded |
| Skill generator | Discarded |
| Marketplace / cohesive versioning | Discarded |

## Working principles (invariants)

These constrain how features in this PRD are built and are not themselves
deliverables of this PRD:

- **Tests first.** Before changing production code, the agent verifies that tests
  cover the expected behavior; if absent, it writes them first (red), confirms
  they fail for the right reason, then changes code. Focus on well-designed unit
  and integration tests.
- **Side-by-side work.** The human follows along; no fire-and-forget autonomous
  builds. Subagents are used only for read-only investigation and for reviewing a
  diff against the tests, returning to the main thread.
- **Single canonical source.** `project/` stays the source of truth; all targets
  are generated from it.
- **Platform.** macOS and Linux only.

## Success criteria

- `geremmyas init/sync/global --targets codex` produces a Codex document plus a
  skill index in the correct locations (project and global), idempotently, behind
  the `geremmyas:generated` marker.
- `geremmyas lint` flags skills whose descriptions lack "use when" triggers or
  negative scope, exceed the description length limit, contain disallowed markup,
  whose `name` does not match the folder, or whose body exceeds the line limit;
  it passes clean skills and is wired into CI.
- No regression in existing targets, packs, guardrails, or SDD gates.

## Linked specs

- `specs/0001-codex-target/` — Codex generation target.
- `specs/0002-skill-validator/` — `geremmyas lint` skill description validator.
