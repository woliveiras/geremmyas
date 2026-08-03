---
name: docs
description: "Maintain project docs, glossary, ADR/MADR, or RFC artifacts through focused modes. Use when code, terminology, or accepted decisions need durable documentation. Do not use for feature specifications or unrelated prose."
---

# Docs

Route a documentation task to only the support material it needs. Do not load
every reference up front.

## Routing

| Request or change | Load |
| --- | --- |
| API, architecture, setup, configuration, onboarding | [project docs](./references/project-docs.md) |
| Domain vocabulary, ambiguous or conflicting terms | [glossary](./references/glossary.md) |
| Accepted complex and hard-to-reverse architecture decision | [ADR/MADR](./references/adr.md) |
| Design proposal that still needs discussion and alternatives | [RFC](./references/rfc.md) |

Load [writing style](./references/writing-style.md) only when producing or
rewriting prose. Load a template or asset only when creating that artifact.

## Contract

1. Inspect the user request, repository convention, existing artifact, and
   implemented behavior before choosing a mode.
2. Select one mode by default. Select more only when the request materially
   spans multiple artifact types.
3. Update the smallest relevant surface and preserve neighboring structure.
4. Keep PRDs, `spec.md`, `plan.md`, and `tasks.md` under `spec`, not this skill.
5. Do not create an ADR for routine or reversible choices. Use an RFC while a
   material proposal is unresolved; record an ADR only after acceptance.
6. Validate local links, filenames, indexes, and claims against current code.
7. Report which mode and references were loaded.
