---
name: vertical-tdd
description: "Implement one behavior at a time with a red-green-refactor loop. Use when: executing tasks.md, doing TDD, implementing a feature from specs. Do not use: for test planning, bulk implementation."
---


# Vertical TDD

Use one behavior as the unit of progress. Run only after the user has approved
the spec.

## Process

1. Read `spec.md`, `plan.md`, and `tasks.md` in the feature folder.
2. Pick the next task with `- [ ]` or continue `- [~]` in progress. Mark the
   task `- [~]` when starting and `- [x]` when done.
3. Read the task's **test-type** (`unit`, `integration`, or `both`) and the
   spec's Test Strategy. If missing, decide using the same rules as
   [test generation reference](./references/generate-tests-from-spec.md).
4. Write one test that verifies observable behavior through the appropriate
   boundary (unit: public API of a module; integration: cross-module or I/O).
5. Run the test and confirm it fails for the expected reason.
6. Implement the minimum production code to pass that test.
7. Run the focused test and the nearest relevant suite.
8. Repeat for the next behavior or task.
9. Refactor only when tests are green; rerun tests after each refactor.
10. When all acceptance criteria are done and code is merged, set `status:
    Implemented` in `spec.md` frontmatter and update the row in `specs/README.md`.
11. Before stopping work (task done, session ending, or handing off), reconcile
    status: mark every finished task `- [x]` in `tasks.md`, clear any stale
    `- [~]`, and annotate or mark done any `plan.md` item fully delivered. Do
    not declare work complete with stale checkboxes.

## Rules

- Do not write all tests first and then all implementation.
- Do not test private functions, call order, or internal collaborators.
- Mock only system boundaries (network, time, filesystem, external services).
- If a test appears wrong, revisit the spec before changing the test.
- When the user only wants tests from acceptance criteria without implementation,
  follow [test generation](./references/generate-tests-from-spec.md).
- When progress becomes circular or scope changes materially, follow
  [abort criteria](./references/abort-criteria.md).
- Update `tasks.md` checkboxes as each task starts (`[~]`) and finishes (`[x]`).
  Stale checkboxes block reliable resumption from `specs/README.md` and the
  feature folder.
