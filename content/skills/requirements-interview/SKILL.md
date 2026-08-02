---
name: requirements-interview
description: "Explore code and resolve product or technical ambiguity before specification. Use when starting or expanding a feature. Do not use when requirements are already clear."
---


# Requirements Interview

Reach shared understanding before writing PRDs, specs, plans, or tasks.
For the full stop conditions, see
[readiness and authority boundaries](./references/approval-gates.md).

## Process

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
7. Record explicit session overrides such as read-only, plan-only, no-edits, or
   no-commits. These overrides take precedence over the autonomous workflow.
8. When the objective is clear, let `generate-spec` create and validate the
   durable artifacts without adding a conversational approval pause.

## Rules

- Local atomic commits are the default after a verified slice. Omit
  conversational permission and file, hunk, or message selection prompts.
- A read-only, plan-only, no-edits, or no-commits override disables commits for
  the session. Report changed or proposed files at the commit boundary. Commit
  authority never implies push or production authority.
- Do not ask questions that code or existing docs can answer.
- Do not write a spec until the key branches of the decision tree are resolved.
- Infer decisions from repository conventions when the choice is reversible and
  record the inference in the relevant artifact.
- Escalate a feature before implementation only when blast radius is high and
  at least one of these is true: rollback is unproven, or critical harness
  evidence is missing. Product ambiguity that materially changes the objective
  still requires clarification.
- Do not create an ADR unless the decision is complex, hard to reverse,
  surprising without context, and based on a real trade-off.
