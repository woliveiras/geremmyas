---
spec: "0001"
title: Codex generation target
family: multi-assistant
phase: 1
status: Approved
owner: ""
depends_on: []
origin: docs/prds/2026-06-22-multi-assistant-framework.md
---

# Spec: Codex generation target

## Context & Motivation

geremmyas generates per-IDE artifacts from a single canonical source under
`project/`. Copilot is native (direct file copy); Cursor, Claude Code, and
OpenCode are generated. Codex is used daily on personal projects but is not a
target, so synced skills and the `AGENTS.md` contract are not reachable from
Codex.

Codex reads `AGENTS.md` natively each session (project tree walk plus a global
`~/.codex/AGENTS.md`). This makes a Codex target low-cost: it mirrors the
existing OpenCode target, which emits an `AGENTS.md`-style document plus an
on-demand skill index. The skill index in that document is also the lightweight
discovery path for Codex (a full session-start bootstrap is out of scope, see the
PRD non-goals).

## Requirements

### Functional

- [ ] Add `codex` as a recognized target alongside `copilot`, `cursor`,
      `claude-code`, and `opencode`.
- [ ] Generate a Codex document from `project/AGENTS.md` plus the synced skill
      index and agent roles, reusing the shared IDE-agents document builder.
- [ ] The generated Codex document includes an explicit "consult skills before
      acting" directive next to the skill index, so the index is a trigger and
      not just a passive list (the Codex Nível 1 discovery path).
- [ ] Support project scope and global scope with correct output paths.
- [ ] Mark generated output with the `geremmyas:generated` marker so sync is
      idempotent and respects `--force` overwrite semantics.
- [ ] Surface `codex` in CLI help text, the `--targets` flag documentation, and
      the doctor/summary path that lists target destinations.

### Non-Functional

- [ ] Platform: macOS and Linux only.
- [ ] No new third-party dependencies; reuse existing generation helpers.
- [ ] No regression to existing targets (copilot, cursor, claude-code,
      opencode).

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Isolated logic, pure functions, single module, no I/O | target normalization/validation accepts `codex` |
| **integration** | Cross-module flows, CLI + filesystem | `init`/`sync`/`global --targets codex` writes expected files |
| **both** | Feature spans layers | — |

Default: **integration** for the generation flow (CLI + filesystem), with a
**unit** check for target validation.

## Acceptance Criteria

- [ ] Given a project initialized with `--targets codex`, when `geremmyas sync`
      runs, then a Codex document is written at the project Codex path and
      contains the `AGENTS.md` body plus a "Skills (on demand)" index.
- [ ] Given a generated Codex document, when it is inspected, then it contains an
      explicit "consult skills before acting" directive adjacent to the skill
      index (not only the list of skills).
- [ ] Given `--targets codex` at global scope, when `geremmyas global` runs, then
      the Codex document is written under the global Codex location
      (`~/.codex/...`).
- [ ] Given a Codex document already generated, when `geremmyas sync` runs again
      without changes, then the file is unchanged (idempotent) and is only
      overwritten when the user customized it and passes `--force`.
- [ ] Given `geremmyas sync --targets codex,copilot`, when it runs, then both the
      Codex document and the native Copilot output are produced without
      interfering with each other.
- [ ] Given an unknown target value, when targets are validated, then the error
      message lists `codex` among the valid targets.
- [ ] Given `geremmyas` help or summary output, when targets are listed, then
      `codex` and its destination appear.

## Edge Cases

- Skill set is empty: the document is still generated, with no skill index
  section.
- A user-customized Codex file exists without the generated marker: respect the
  existing overwrite rules (do not clobber without `--force`).
- Global home directory resolution fails: fail with a clear error, consistent
  with the other global targets.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Generation approach | Reuse `buildIDEAgentsDoc` like OpenCode/Claude | Lowest cost, parity with existing targets |
| Discovery mechanism | Skill index inside the Codex document | Codex loads `AGENTS.md` every session; index is the trigger |
| Session-start bootstrap | Out of scope | Deferred per PRD non-goals |

## Open questions

- **Where the "consult skills" directive lives.** The skill index is built by the
  shared `buildIDEAgentsDoc` helper used by OpenCode and Claude too. Option A:
  strengthen the directive in the shared builder (also improves OpenCode/Claude,
  consistent). Option B: emit the directive only on the Codex path (surgical, but
  diverges between targets). Confirm during review; Option A is preferred for
  consistency.
- **Project-scope path.** Codex reads the repository root `AGENTS.md` natively.
  Option A: write a separate `.codex/AGENTS.md` (parity with `.opencode/`).
  Option B: rely on the root `AGENTS.md` and only emit the skill index. The plan
  proposes Option A for parity and to carry the skill index without touching the
  root contract; confirm during review.
- **Global path shape.** `~/.codex/AGENTS.md` is the assumed global destination;
  confirm against the installed Codex layout.

## Out of Scope

- Auto-trigger session-start bootstrap.
- A Codex-specific plugin manifest or hooks.
- Any change to the canonical `AGENTS.md` content.
