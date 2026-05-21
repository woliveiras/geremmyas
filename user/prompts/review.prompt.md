---
description: "Quick code review checklist. Use for fast feedback without specs. For spec-driven review, use the @reviewer agent instead."
---

Review the provided code changes using this checklist:

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

Be specific. Point to exact lines. Suggest concrete alternatives, not vague improvements.
