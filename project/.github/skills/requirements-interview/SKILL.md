---
name: requirements-interview
description: "Interview before PRD or spec work by exploring code first, resolving product and technical ambiguity, and deciding the right artifact. Use when: starting a feature, expanding an existing feature, clarifying requirements, or deciding whether a PRD must change."
---

# Requirements Interview

Reach shared understanding before writing PRDs, specs, plans, or tasks.

## Process

0. **Commit permission (first):** Before other questions, ask: "Do I have
   permission to create git commits for this work, or do you handle commits
   yourself?" Record the answer for the session. If the user does not grant
   permission, downstream skills must not run `git-commit`; report changed files
   instead.
1. Explore the relevant code, docs, specs, PRDs, `GLOSSARY.md`, `CONTEXT.md`,
   and ADRs before asking questions.
2. Classify the work:
   - new feature
   - existing feature expansion
   - bugfix
   - refactor
   - documentation/setup
3. For existing features, decide whether the product flow changes. If it does,
   update the PRD before writing specs. If it does not, write targeted specs.
   If a PRD exists without a spec, do not proceed to implementation until a spec
   is written from the PRD.
4. Ask one question at a time. For each question, include a recommended answer
   based on the codebase and documents.
5. Challenge vague or overloaded terms and propose canonical vocabulary.
6. Record resolved decisions in the right artifact: PRD, spec, ADR, or
   `tasks.md`.

## Rules

- Do not skip the commit permission question at the start.
- Do not ask questions that code or existing docs can answer.
- Do not write a spec until the key branches of the decision tree are resolved.
- Do not create an ADR unless the decision is complex, hard to reverse,
  surprising without context, and based on a real trade-off.
