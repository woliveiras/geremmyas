---
spec: "0004"
title: Skill catalog tiering
family: multi-assistant
phase: 2
status: Implemented
owner: ""
depends_on: []
origin: Direct user request (catalog review)
---

# Spec: Skill catalog tiering

## Context & Motivation

The catalog distributes ~40 skills in a single flat tier. Every pack installs
both globally and per-project, so any repo carries skills it never uses while
the full skill listing is injected into the assistant context on every turn.
The coherent SDD workflow (requirements-interview through git-commit, plus the
8 guardrails) is the catalog's main asset, but it is diluted by stack-specific
helpers, version-pinned framework recipes, and academic-writing skills mixed
into the same flat tier.

A note on scope of the fix: the host always injects the description of every
installed skill, and that injection is not something geremmyas controls. The
only lever geremmyas owns is **which skills get installed** in a repo. So this
spec reduces the injected set by installing fewer, scoped packs, not by trying
to suppress injection of installed skills.

This spec introduces an explicit **tier/scope** for each pack so a repo only
carries what it needs, and consolidates redundant skills. See
[docs/decisions/0001-tier-skill-catalog.md](../../docs/decisions/0001-tier-skill-catalog.md).

## Requirements

### Functional

- [ ] Add a `tier` field to every pack in `catalog/packs.json` with one of:
      `core`, `stack`.
- [ ] `core`: general SDLC workflow packs (`core`, `sdd`): the SDD pipeline,
      the 8 guardrails, `bugfix-loop`, `git-commit`, and `skill-authoring`.
- [ ] `stack`: everything else, opt-in per project and never forced into a repo
      that does not use it (framework recipes, CI-per-language,
      `terraform-change`, `gcloud-operation`, `supabase-workflow`,
      `postgres-query-review`, `chromadb-rag-workflow`, `langgraph-agent-design`,
      `llm-integration-review`, `rust-release`, `premortem`, `research`,
      `blog`).
- [ ] Classify `research` and `blog` as `stack` packs (opt-in, not default).
      `research` = `scientific-paper`, `scientific-case-study-research`, and
      `paper-review`; `blog` = `text-review`. Both already exist as packs.
- [ ] Add `paper-review` to the `research` pack by vendoring its skill into
      `project/.github/skills/paper-review/`.
- [ ] Remove caveman from geremmyas entirely (it is not distributed as a pack
      and its ecosystem of hooks, statusline, and MCP shrink cannot be
      reproduced by copying markdown). Recommend the upstream installer in
      `README.md`, and drop the caveman brevity directive from the distributed
      `AGENTS.md` and maintainer instructions.
- [ ] `geremmyas lint` validates that every pack declares a valid `tier`.
- [ ] Install/sync respects tier: `stack` packs are not auto-included with
      `core`; `lint`/`doctor` reject an unknown or missing tier.
- [ ] Update `docs/creating-packs.md`, `README.md`, and `specs/README.md` to
      document tiers and the new default-install set.

### Non-Functional

- [ ] No new third-party dependencies.
- [ ] Backward-compatible migration path: a missing `tier` fails lint with a
      clear message rather than silently defaulting.
- [ ] Reduce the default per-repo skill count from ~52 to ~18-22.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Catalog parsing, tier validation, install filtering | `packs.json` tier field parses; lint rejects missing/invalid tier; `research` contains `paper-review` |
| **integration** | End-to-end install behavior | `geremmyas init --packs core` installs only core; stack pack opt-in adds only its skills |

Default: **unit** for schema/validation/dispatch, plus **integration** for the
install-scope behavior.

## Acceptance Criteria

- [ ] Given `catalog/packs.json`, when parsed, then every pack has a `tier` of
      `core` or `stack`.
- [ ] Given a pack with no `tier`, when `geremmyas lint` runs, then it exits
      non-zero with a message naming the offending pack.
- [ ] Given `geremmyas init` with only core selected, when sync runs, then only
      `core`-tier skills are installed and stack skills are absent.
- [ ] Given the catalog, when parsed, then no caveman pack or skill is present,
      and `README.md` links the upstream caveman installer.
- [ ] Given the `research` pack, when opted in, then `paper-review` installs
      alongside the scientific skills.
- [ ] Given the research/blog skills, when no `research`/`blog` pack is opted
      in, then they are not installed in a default project.
- [ ] Given the docs, when tiers ship, then `creating-packs.md`, `README.md`,
      and `specs/README.md` describe the tier model and default set.

## Edge Cases

- A skill that fits two tiers: the canonical skill is `core`; no skill appears
  in two tiers.
- Existing consumer installs predating tiers: lint flags the missing field;
  `sync` does not silently re-tier without the field present.
- `ci-workflow` (general) stays `core`; per-language CI setups are `stack`.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Tier model | `core` / `stack` | Maps to "installed everywhere" vs "opt-in per stack"; the `personal` tier was dropped for lack of distributed members |
| Missing tier | Hard lint failure | Avoids silent misclassification on migration |
| Caveman | Removed from geremmyas; recommend upstream in README | External ecosystem (hooks, statusline, MCP shrink) that geremmyas's copy-markdown model cannot reproduce |
| Research skills | Opt-in `research` stack pack incl. `paper-review` | Academic writing is out of the default SDLC path |
| Lazy injection | Rejected | The host injects every installed skill; geremmyas cannot suppress injection of an installed skill, only avoid installing it |

## Out of Scope

- Rewriting the content of any individual skill (only tier metadata and
  relocation change here).
- Changing the SDD workflow orchestration in `AGENTS.md`.
- Removing skills outright, except caveman which is removed from geremmyas and
  delegated to its upstream installer.
- Multi-catalog or multi-repo splits (rejected in ADR 0001 as Option 3).
- Runtime/online catalog updates.
- Suppressing host injection of installed skills (a lazy router was considered
  and rejected: the host scans the whole skills root, so hiding a skill means
  not installing it as a skill at all).
