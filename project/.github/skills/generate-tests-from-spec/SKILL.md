---
name: generate-tests-from-spec
description: "Generate tests from a spec's acceptance criteria. Use when: you have an approved spec and need tests, TDD from spec. Do not use: before spec approval, for test planning without specs."
---


# Generate Tests from Spec

Generate tests that cover every acceptance criterion in an **approved** spec.

## When to Use

- The spec is written and the user has explicitly approved it
- You're starting the TDD cycle (tests before code)
- You want to verify spec coverage with tests

Do not run this skill before spec approval (see `AGENTS.md`).

## Test-Type Decision

1. Read the spec's **Test Strategy** section and each task's `test-type` in
   `tasks.md` when present.
2. Choose the test boundary:

| test-type | Write |
| --- | --- |
| **unit** | Fast, isolated tests; mock only system boundaries if needed |
| **integration** | Tests across modules, real DB/API/filesystem per project patterns |
| **both** | At least one unit and one integration test for the criterion or task |

Signals for **integration**: multiple modules, HTTP/CLI entrypoints, database,
external services, acceptance criteria describing end-to-end behavior.

Signals for **unit**: pure logic, single function/module, no I/O.

3. Follow `.github/instructions/testing.instructions.md` and
   `integration-testing.instructions.md` when they apply to edited files.

## Procedure

1. Read `specs/NNNN-<slug>/spec.md` (or path the user gives).
2. Read `tasks.md` in the same feature folder for per-task `test-type`.
3. Extract all acceptance criteria (Given/When/Then items).
4. For each criterion, generate at least one test of the chosen type(s):
   - Happy path
   - Edge cases from the spec
   - Error cases from the spec
5. Follow **existing test patterns** in the project (framework, naming, fixtures).
6. If no existing tests are found, ask the user which framework to use.
7. Colocate test files with the code they exercise when that is the project norm.

## Rules

- One test per behavior — each test verifies one acceptance criterion or slice
- Test names must read as documentation
- Tests must be self-contained — no shared mutable state between tests
- Do NOT implement production code — only tests (red phase)
- Tests should fail initially for the expected reason
- Do not generate tests before spec approval

## Output

For each acceptance criterion:

```
Criterion: "Given X, when Y, then Z"
test-type: unit | integration | both
Test: test_file.ext → test function name
```

Then write the actual test code.
