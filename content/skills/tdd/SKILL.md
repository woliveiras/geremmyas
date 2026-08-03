---
name: tdd
description: "Implement one observable behavior at a time with red-green-refactor. Use when executing a Ready feature task. Do not use for test-only planning, bulk implementation, or bugfix diagnosis."
---


# TDD

Use one behavior as the unit of progress. Run when the feature is `Ready`: its
spec, plan, tasks, acceptance criteria, and test strategy pass the repository's
machine-readable readiness checks. Human approval is not a feature gate.

## Process

1. Read `spec.md`, `plan.md`, and `tasks.md` in the feature folder.
2. Pick the next task with `- [ ]` or continue `- [~]` in progress. Mark the
   task `- [~]` when starting and `- [x]` when done.
3. Read the task's **test-type** (`unit`, `integration`, or `both`) and the
   spec's Test Strategy. If missing, decide using the same rules as
[test generation reference](./references/test-generation.md).
4. Write one test that verifies observable behavior through the appropriate
   boundary (unit: public API of a module; integration: cross-module or I/O).
5. Run the test and confirm it fails for the expected reason.
6. Implement the minimum production code to pass that test.
7. Run the focused test and the nearest relevant suite.
8. Repeat for the next behavior or task.
9. Refactor only when tests are green; rerun tests after each refactor.
10. Use the lifecycle `Draft` -> `Ready` -> `In Progress` -> `Verified`. Set
    `In Progress` when implementation begins.
11. Before `Verified` or any handoff, mark every finished task `[x]`, clear stale
    `[~]`, reconcile required docs and plan items, and obtain independent review
    with no actionable finding.
12. Set `Verified` in `spec.md` and update `specs/README.md` only when every
    acceptance criterion has fresh evidence and step 11 is complete. Merge,
    commit, push, and deployment are separate authority dimensions and do not
    change feature lifecycle state.

## Rules

- Do not write all tests first and then all implementation.
- Do not test private functions, call order, or internal collaborators.
- Mock only system boundaries (network, time, filesystem, external services).
- If a test appears wrong, revisit the spec before changing the test.
- When the user only wants tests from acceptance criteria without implementation,
  follow [test generation](./references/test-generation.md).
- When progress becomes circular or scope changes materially, follow
  [abort criteria](./references/abort-criteria.md).
- Escalate to the user only for unresolved material product ambiguity, or when
  blast radius is high and at least one evidence gap remains: rollback is
  unproven, or critical harness evidence is missing. Otherwise document the
  decision, use specialist review when useful, and continue through the
  executable gates.
- Update `tasks.md` checkboxes as each task starts (`[~]`) and finishes (`[x]`).
  Stale checkboxes block reliable resumption from `specs/README.md` and the
  feature folder.
