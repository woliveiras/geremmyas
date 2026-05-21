---
name: vertical-tdd
description: "Implement one behavior at a time with a red-green-refactor loop. Use when: executing tasks.md, doing TDD, implementing a feature from specs, or avoiding bulk test generation."
---

# Vertical TDD

Use one behavior as the unit of progress.

## Process

1. Read the relevant spec and `tasks.md`.
2. Pick one task or one behavior from the task.
3. Write one test that verifies observable behavior through a public interface.
4. Run the test and confirm it fails for the expected reason.
5. Implement the minimum production code needed to pass that test.
6. Run the focused test and the nearest relevant suite.
7. Repeat for the next behavior.
8. Refactor only when tests are green, and rerun tests after each refactor.

## Rules

- Do not write all tests first and then all implementation.
- Do not test private functions, call order, or internal collaborators.
- Mock only system boundaries such as network, time, filesystem, or external
  services.
- If a test appears wrong, revisit the spec before changing the test.
- Use `generate-tests-from-spec` when the user only wants tests generated from
  acceptance criteria without implementation.
