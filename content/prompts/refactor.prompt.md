---
description: "Refactor code while preserving behavior. Use when you want to improve code structure without changing what it does."
---

Refactor the provided code. Rules:

1. **Preserve behavior exactly** — the output, side effects, and error handling must remain identical
2. **Improve readability** — clearer names, simpler control flow, fewer nesting levels
3. **Reduce duplication** — but only when the duplicated code changes for the same reasons (avoid premature DRY)
4. **Simplify interfaces** — fewer parameters, smaller surface area
5. **Keep changes minimal** — do not reorganize files or refactor beyond what's needed

Show the refactored code and explain each change in one sentence.

If existing tests cover this code, verify they still pass conceptually. Do not modify tests.
