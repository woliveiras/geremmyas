---
name: generate-glossary
description: "Extract and formalize domain terminology into a glossary. Use when: defining domain terms, building glossaries, hardening terminology. Do not use: for general documentation, simple dictionaries."
---


# Generate Glossary

Extract domain terminology from conversation, specs, or code and formalize it
into the project's vocabulary artifact.

`GLOSSARY.md` is the default for new projects. `CONTEXT.md` is also supported
when a repository already uses that convention. If both exist, read both; treat
`GLOSSARY.md` as the canonical term list and `CONTEXT.md` as broader domain
context unless the project says otherwise. If they conflict, ask before changing
either file.

## When to Use

- Starting a new project and aligning on terminology
- Domain terms are used inconsistently across the codebase
- Multiple people use different words for the same concept
- Onboarding new team members who need a term reference

## Procedure

1. Scan the conversation, specs, or code for domain-relevant nouns, verbs, and concepts
2. Identify problems:
   - Same word used for different concepts (ambiguity)
   - Different words used for the same concept (synonyms)
   - Vague or overloaded terms
3. Propose canonical terms — be **opinionated** about term choices
4. Choose the target:
   - update `GLOSSARY.md` when it exists
   - update `CONTEXT.md` when it is the only existing vocabulary artifact
   - create `GLOSSARY.md` for new projects with neither file
5. Use the [glossary template](./assets/glossary-template.md) when creating
   `GLOSSARY.md`
6. Incorporate new terms without duplicating or contradicting existing
   definitions

## Rules

- **Be opinionated** — pick the best term, list others as "aliases to avoid"
- **Flag ambiguities explicitly** — call out conflicts with clear recommendations
- **Domain terms only** — skip generic programming concepts (array, function, endpoint)
- **Tight definitions** — one sentence max; define what it IS, not what it does
- **Show relationships** — use bold term names, express cardinality where obvious
- **Group by context** — split into tables when natural clusters emerge
- **Include example dialogue** — 3-5 exchanges showing terms used precisely

## Write like a human

Write definitions and examples a practitioner would actually write, not generic assistant output. Keep the glossary format, but keep these AI writing tells out of the prose:

- No em dashes and no curly quotes: use commas, parentheses, or new sentences, and straight quotes (" ').
- Cut filler vocabulary: "delve", "leverage", "utilize" (use "use"), "robust", "seamless", "crucial", "pivotal", "testament", "underscore", "showcase", "foster", "landscape" (as an abstract noun).
- Drop significance padding ("plays a pivotal role", "reflects a broader shift") and trailing "-ing" analysis clauses ("...highlighting its importance").
- Prefer plain "is"/"has" over "serves as"/"boasts". Define what the term is; do not editorialize about its importance.
- Skip forced tricolons and decorative emoji.
