---
spec: "0010"
title: Modular game-development routing
family: multi-assistant
phase: 8
status: In Progress
owner: ""
depends_on: ["0009"]
origin: Rusted Codex A/B benchmark and user-approved game-dev pack split
---

# Spec: Modular game-development routing

## Context & Motivation

The lazy workflow harness reduced the general workflow surface, but the
`game-dev` stack pack still advertises eleven specialist skills whenever it is
installed. A controlled Codex experiment in Rusted showed that the complete
pack increased discovery pressure without reliably activating the relevant
game skill for a concrete localization and Battle UI analysis.

This feature splits the family into task-oriented packs while retaining
`game-dev` as an explicit complete metapack. It also makes each skill's
frontmatter trigger more discriminating and adds a repository-owned routing
conformance corpus. The conformance test is a deterministic lexical proxy; it
does not claim to reproduce or control an assistant runtime's semantic router.

`GLOSSARY.md` is absent at specification time. Existing catalog and harness
terminology remains authoritative.

## Requirements

### Functional

- [ ] Add `game-core` with gameplay programming and game testing.
- [ ] Add `game-ui` with UI/accessibility and game-feel capabilities.
- [ ] Add `game-systems` with AI, procedural generation, and save/progression.
- [ ] Add focused `game-performance`, `game-audio`, `game-art`, and
      `game-delivery` packs.
- [ ] Retain `game-dev` as a complete metapack that depends on all focused packs
      and owns no skill files directly.
- [ ] Preserve `game-art-2d` as a pack-name compatibility alias without adding
      another discoverable skill or duplicate materialized path.
- [ ] Give all eleven game skills positive and negative trigger scope that
      distinguishes neighboring capabilities.
- [ ] Add routing conformance cases covering every game skill, including a
      localization and Battle UI case that selects `game-ui-accessibility`.
- [ ] Update public pack and skill documentation for focused installation.
- [ ] Reconcile Rusted from `game-dev` to only its routinely relevant focused
      packs and prove the obsolete managed skills are removed conservatively.

### Non-Functional

- [ ] No new third-party dependencies or LLM/runtime integration.
- [ ] Existing `game-dev` configuration remains valid and resolves all eleven
      skills exactly once.
- [ ] Pack splitting changes discovery scope, not canonical skill behavior.
- [ ] Routing tests state their lexical limitation and never present their
      result as runtime enforcement.
- [ ] All skill descriptions remain inside existing lint and context budgets.
- [ ] Rusted benchmark history and unrelated user changes remain intact.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Catalog closure and description routing invariants | focused pack membership, complete metapack, compatibility alias, lexical routing winners |
| **integration** | Materialized project reconciliation and assistant behavior | Rusted local sync, removed obsolete managed skills, Codex skill-load smoke |

## Acceptance Criteria

- [ ] Given any focused game pack, when the catalog resolves it, then only its
      declared domain skills and dependencies become discoverable.
- [ ] Given `game-dev`, when the catalog resolves it, then all eleven game
      skills materialize exactly once through focused pack dependencies.
- [ ] Given `game-art-2d`, when the catalog resolves it, then the canonical
      `game-art` pack supplies the same single skill without duplicate output.
- [ ] Given each routing corpus prompt, when description signals are scored,
      then the declared game skill is the unique winner.
- [ ] Given a request about Battle localization, canvas/DOM labels, responsive
      UI, focus, touch, contrast, or accessible names, then
      `game-ui-accessibility` is the declared routing target and adjacent game
      descriptions exclude that ownership.
- [ ] Given Rusted's routine needs, when `base`, `typescript-base`, `game-core`,
      `game-ui`, and `game-systems` are materialized, then art, audio,
      performance, and delivery skills are not advertised until selected.
- [ ] Given the Rusted routing smoke, when Codex analyzes the concrete Battle UI
      request, then the evidence shows whether it actually reads the expected
      skill; any failure is reported rather than hidden by structural tests.

## Decisions

| Decision | Choice | Reasoning |
| --- | --- | --- |
| Complete pack | Keep `game-dev` as a metapack | Existing users retain the explicit full-workflow selection |
| Focused default | Install only domain packs routinely needed by a project | Reduces discovery competition without withholding capabilities from real game tasks |
| Art compatibility | Pack alias, not skill alias | Pack metadata has no assistant discovery cost; duplicate skills would restore it |
| Routing validation | Versioned lexical corpus plus runtime smoke | Deterministic CI catches description drift while the smoke records actual runtime behavior |
| Runtime scope | No LLM client in Geremmyas | The tool remains a harness installer and diagnostics CLI |

## Out of Scope

- Predicting exact model token usage, latency, or semantic activation.
- Making a textual skill trigger into hard runtime enforcement.
- Redesigning the procedural bodies, references, scripts, or assets of the game
  skills beyond trigger wording.
- Installing art, audio, performance, or delivery packs in Rusted before a task
  actually requires them.
