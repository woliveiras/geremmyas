---
description: "Generate tests for the provided code. Use when you want unit tests that cover behavior, edge cases, and errors."
---

Generate unit tests for the provided code. Follow these rules:

1. **Match existing patterns** — use the same test framework, naming conventions, and file organization already in the project
2. **One test per behavior** — each test verifies one specific thing
3. **Descriptive names** — test names should read as documentation (e.g., "returns error when input is empty")
4. **Cover the important paths**:
   - Happy path (expected input → expected output)
   - Edge cases (empty, null, boundary values)
   - Error cases (invalid input, failures)
5. **No implementation details** — test the interface, not the internals
6. **Self-contained** — each test sets up its own data, no shared mutable state between tests

If a spec exists for this code, derive test cases from the acceptance criteria in the spec.
