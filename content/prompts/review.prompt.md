---
description: "Quick code review checklist. Use for fast feedback without specs. For spec-driven review against requirements or acceptance criteria, use the @reviewer agent instead."
---

Review the provided code changes using this quick checklist.

This prompt is for general review when there is no spec-driven workflow in
scope. If the user asks whether the implementation matches specs, acceptance
criteria, or tests generated from specs, route them to the `@reviewer` agent.

If a `GLOSSARY.md` or `CONTEXT.md` file exists, use that vocabulary when judging
names, behavior, and user-facing language.

## Correctness
- Does the code do what it claims to do?
- Are edge cases handled?
- Are error paths covered?

## Security
- Is user input validated and sanitized?
- Are queries parameterized (no string concatenation)?
- Are secrets kept out of code?

## Readability
- Is the code clear without excessive comments?
- Are names descriptive and consistent with the codebase?
- Are functions focused (single responsibility)?

## Tests
- Are there tests for the new/changed behavior?
- Do tests cover edge cases and error scenarios?
- Are test names descriptive enough to serve as documentation?
- If tests look wrong, is there a spec or acceptance criterion that proves the
  mismatch?

Be specific. Point to exact lines. Suggest concrete alternatives, not vague improvements.
