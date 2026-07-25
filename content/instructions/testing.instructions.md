---
description: "Use when writing or reviewing tests. Covers general test design independent of test framework."
applyTo: "**/*.test.*,**/*.spec.*,**/tests/**,**/test/**,**/*_test.*"
---

# Testing Conventions

- Verify observable behavior, not private implementation details.
- Keep each test focused on one behavior or acceptance criterion.
- Use names that describe the scenario and expected outcome.
- Prefer Arrange-Act-Assert or Given-When-Then structure.
- Keep tests isolated; no shared mutable state between test cases.
- Use the existing framework, fixtures, helpers, and naming style.
- Mock only system boundaries such as network, filesystem, time, and external
  services.
- Do not change tests just to make implementation pass. If a test seems wrong,
  revisit the spec or acceptance criteria first.
