---
name: generate-spec
description: "Generate a structured spec from direct input. Use when: you already know what you want and just need the formatted template. Do not use: for exploration and interviews (use spec-writer agent)."
---


# Generate Spec

Generate a structured feature specification and the mandatory companion
artifacts `plan.md` and `tasks.md` in a numbered feature folder. Keep
`specs/README.md` in sync as the specs index.

## When to Use

- You have a clear idea of what to build
- You want a formatted spec without an interview
- You're documenting an existing decision

For unclear requirements that need exploration, use the `spec-writer` agent instead.

## File Location and Naming

Always use a numbered feature folder with all three artifacts:

```text
specs/README.md                      # index (families, status, numbering)
specs/NNNN-<feature-slug>/spec.md
specs/NNNN-<feature-slug>/plan.md
specs/NNNN-<feature-slug>/tasks.md
```

Allocate `NNNN` using the rules in `specs/README.md` (next global number or next
slot in the family's reserved block). Example:

```text
specs/0042-jwt-authentication/spec.md
specs/0042-jwt-authentication/plan.md
specs/0042-jwt-authentication/tasks.md
```

If `specs/README.md` is missing, create it from the geremmyas template or copy
the structure documented in `AGENTS.md`.

## Procedure

1. Infer the feature name, brief description, family (if any), phase, and change
   type from the request and repository. Escalate one focused question only for
   a missing material choice that would change the outcome.
2. Read `specs/README.md` and list existing `specs/NNNN-*` folders.
3. Allocate the next spec number and create or update the feature folder.
4. Fill in [spec template](./assets/spec-template.md) with correct frontmatter
   (`spec`, `title`, `family`, `phase`, `status: Draft`) and save as `spec.md`.
5. Write `plan.md` with implementation sequencing and dependencies.
6. Write initial `tasks.md` using the bundled
   [task breakdown](./references/task-breakdown.md) conventions (vertical
   slices, checkboxes, `test-type` per task).
7. **Update `specs/README.md`:** add or update the row in the correct family/
   phase table (Spec link, Title, Status `Draft`, Depends on / Origin).
   Reserve or extend family blocks in the numbering table when needed.
8. Run the automatic readiness checks from
   [readiness and authority boundaries](../requirements-interview/references/approval-gates.md).
   Correct failures and record relevant decisions or evidence in the artifacts.
9. Set `status: Ready` in `spec.md` frontmatter and update the index to `Ready`.
   Summarize the artifacts for auditability, then continue to implementation
   unless an explicit read-only, plan-only, or no-edits override applies.
10. Set the feature to `In Progress` when implementation starts. If an in-scope
    assumption changes, update the artifacts and rerun readiness automatically.
    Set it to `Verified` only after every acceptance criterion has fresh
    evidence, every finished task is `[x]`, no stale `[~]` remains, required
    docs and plan items are reconciled, and independent review has no actionable
    finding.

## Output

Use the template from [assets/spec-template.md](./assets/spec-template.md).
Ensure every acceptance criterion is testable (maps to at least one test).
Define test strategy in the spec (unit vs integration vs both).

## Rules

- Do not leave placeholder text in saved artifacts (including `spec: "0000"`).
- Keep acceptance criteria in Given/When/Then form when possible.
- Use `GLOSSARY.md` vocabulary when it exists.
- Put implementation sequencing in `plan.md`, not in the spec body.
- Put task list and progress in `tasks.md`, not only in the spec.
- Put accepted architectural decisions in an ADR when the bar is met (complex and
  hard to reverse), not only in the spec.
- Do not implement or write feature tests before the spec reaches **Ready**.
- Use only `Draft`, `Ready`, `In Progress`, and `Verified` for a new or actively
  migrated feature lifecycle. Existing historical states remain valid without a
  bulk migration. Update frontmatter and `specs/README.md` together at each
  transition.
- Before escalating a feature, confirm that blast radius is high and at least
  one of these is true: rollback is unproven, or critical harness evidence is
  missing.
- When a spec reaches **Verified**, record the fresh verification evidence and
  links to any commits that delivered the work.

## Write like a human

Write prose a practitioner would actually write, not generic assistant output. Keep the spec template and required formatting, but keep these AI writing tells out of the prose:

- No em dashes and no curly quotes: use commas, parentheses, or new sentences, and straight quotes (" ').
- Cut filler vocabulary: "delve", "leverage", "utilize" (use "use"), "robust", "seamless", "crucial", "pivotal", "testament", "underscore", "showcase", "foster", "landscape" (as an abstract noun).
- Drop significance padding ("stands as a testament to", "plays a pivotal role", "reflects a broader shift") and trailing "-ing" analysis clauses ("...highlighting its importance").
- Prefer plain "is"/"has" over "serves as"/"boasts". Attribute every claim to a named source or delete it.
- Skip forced tricolons and decorative emoji. State the point first; cut sentences that only announce a topic or restate the previous one.
