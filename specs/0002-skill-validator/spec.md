---
spec: "0002"
title: Skill description validator (geremmyas lint)
family: multi-assistant
phase: 1
status: Approved
owner: ""
depends_on: []
origin: docs/prds/2026-06-22-multi-assistant-framework.md
---

# Spec: Skill description validator (`geremmyas lint`)

## Context & Motivation

Skill discovery depends on good `description` metadata. Assistants that rely on a
markdown skill index (Claude, OpenCode, and the Codex target from spec 0001)
select a skill from its description alone, so a weak description silently breaks
discovery. The repository already documents quality expectations in the
`skill-authoring` skill (directory name and `name:` must match; the description
must state what the skill does and when to use it), but nothing enforces them.

This spec adds a `geremmyas lint` command that validates skill files in the
canonical `project/.github/skills/` tree against objective rules, runnable
locally and in CI, so weak skills fail fast.

## Requirements

### Functional

- [ ] Add a `lint` subcommand that scans canonical skills and reports violations.
- [ ] Validate each skill's frontmatter `description`:
  - [ ] contains a "use when" trigger phrase;
  - [ ] contains a negative-scope phrase (e.g. "do NOT use" / "not for");
  - [ ] is at most 1024 characters;
  - [ ] contains no angle-bracket markup (`<` / `>`).
- [ ] Validate that frontmatter `name` matches the skill's directory name.
- [ ] Validate that the `SKILL.md` body is at most 500 lines.
- [ ] Exit non-zero when any skill fails; exit zero when all pass.
- [ ] Print a clear per-skill report listing each failing rule and the offending
      skill path.

### Non-Functional

- [ ] Platform: macOS and Linux only.
- [ ] No new third-party dependencies; reuse the existing frontmatter parser.
- [ ] Runs in CI as a required check.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Isolated logic, pure functions, single module, no I/O | each rule passes/fails on crafted description and body inputs |
| **integration** | Cross-module flows, CLI + filesystem | `lint` over a temp skills tree returns the right exit code and report |
| **both** | Feature spans layers | rule engine (unit) + command wiring (integration) |

Default: **both** — unit tests for the rule engine, integration tests for the
command over a fixture skills tree.

## Acceptance Criteria

- [ ] Given a skill whose description lacks a "use when" phrase, when `lint` runs,
      then it reports the missing-trigger rule and exits non-zero.
- [ ] Given a skill whose description lacks a negative-scope phrase, when `lint`
      runs, then it reports the missing-negative-scope rule and exits non-zero.
- [ ] Given a description longer than 1024 characters, when `lint` runs, then it
      reports the length rule and exits non-zero.
- [ ] Given a description containing `<` or `>`, when `lint` runs, then it reports
      the markup rule and exits non-zero.
- [ ] Given a skill whose `name` differs from its directory, when `lint` runs,
      then it reports the name-mismatch rule and exits non-zero.
- [ ] Given a `SKILL.md` body longer than 500 lines, when `lint` runs, then it
      reports the body-length rule and exits non-zero.
- [ ] Given a skills tree where every skill is valid, when `lint` runs, then it
      prints a success summary and exits zero.

## Edge Cases

- A skill directory without a `SKILL.md`: report it as a violation, not a crash.
- Empty or missing `description`: treated as failing the trigger and
  negative-scope rules.
- Description exactly 1024 characters: passes (boundary inclusive); 1025 fails.
- Body exactly 500 lines: passes; 501 fails.
- Frontmatter `name` missing: report name-mismatch.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Source of truth scanned | `project/.github/skills/` canonical tree | Root mirrors via symlinks; canonical avoids double-counting |
| Rule basis | Repository's `skill-authoring` expectations + objective limits | Enforces existing documented conventions |
| Failure mode | Non-zero exit with per-rule report | Usable as a CI gate |

## Open questions

- **Negative-scope detection.** Accept a small set of phrases ("do not use", "do
  NOT use", "not for", "don't use") case-insensitively; confirm the accepted set
  at review.
- **Scope of scan.** Lint canonical skills only, or also user-prompt skills under
  `user/`? Proposed: canonical `project/.github/skills/` only for this spec.

## Out of Scope

- Auto-fixing or rewriting descriptions.
- A skill scaffolding generator (discarded in the PRD).
- Security scanning of skill contents.
- Validating instructions or agents (skills only for this spec).
