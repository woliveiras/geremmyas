# Specs Index

Index of every spec in this repository. Maintain this file by hand when
creating, approving, or completing a spec.

Each spec lives in a numbered folder `specs/NNNN-<slug>/` with `spec.md`,
`plan.md`, and `tasks.md`.

## Numbering

| Block | Family | Reserved |
| --- | --- | --- |
| 0001–0099 | multi-assistant | Personal multi-assistant framework |

## Specs

### multi-assistant

PRD: [docs/prds/2026-06-22-multi-assistant-framework.md](../docs/prds/2026-06-22-multi-assistant-framework.md)

| Spec | Title | Phase | Status | Depends on / Origin |
| --- | --- | --- | --- | --- |
| [0001](0001-codex-target/spec.md) | Codex generation target | 1 | Completed | PRD multi-assistant |
| [0002](0002-skill-validator/spec.md) | Skill description validator (`geremmyas lint`) | 1 | Approved | PRD multi-assistant |
| [0003](0003-version-command/spec.md) | Geremmyas version command | 1 | Implemented | Direct user request |
| [0004](0004-skill-catalog-tiering/spec.md) | Skill catalog tiering | 2 | Implemented | Direct user request (catalog review) |
| [0005](0005-codex-instructions/spec.md) | Codex global instruction distribution | 3 | Implemented | Direct user request (Codex not loading instructions) |
| [0006](0006-context-efficient-workflows/spec.md) | Context-efficient agent workflows | 4 | Approved | Audit of global skills, agents, and context usage; depends on 0004 and 0005 |

## Decisions

| ADR | Title | Status |
| --- | --- | --- |
| [0001](../docs/decisions/0001-tier-skill-catalog.md) | Tier the skill catalog into core, stack, and personal scopes | Implemented |
