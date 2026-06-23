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

1. Ask the user for the feature name, brief description, family (if any), phase,
   and whether this is a new feature, expansion, or documenting an existing
   decision.
2. Read `specs/README.md` and list existing `specs/NNNN-*` folders.
3. Allocate the next spec number and create or update the feature folder.
4. Fill in [spec template](./assets/spec-template.md) with correct frontmatter
   (`spec`, `title`, `family`, `phase`, `status: Draft`) and save as `spec.md`.
5. Write `plan.md` with implementation sequencing and dependencies.
6. Write initial `tasks.md` using `task-breakdown` conventions (vertical
   slices, checkboxes, `test-type` per task).
7. **Update `specs/README.md`:** add or update the row in the correct family/
   phase table (Spec link, Title, Status `Draft`, Depends on / Origin).
   Reserve or extend family blocks in the numbering table when needed.
8. **Approval gate:** Present a summary of `spec.md`, `plan.md`, `tasks.md`, and
   the index row to the user and **stop**. Do not generate tests or production
   code until the user explicitly approves the spec.
9. After user approval, set `status: Approved` in `spec.md` frontmatter and
   update the Status column in `specs/README.md` to `Approved`.

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
- Do not implement or write feature tests before spec approval.
- When a spec reaches **Implemented**, update frontmatter and `specs/README.md`
  (status + links to PRs/commits in `spec.md`).
