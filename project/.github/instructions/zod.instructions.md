---
description: "Use when creating, editing, or reviewing Zod v4 schemas. Covers schema definition, type inference, validation patterns, common v4 API differences, and runtime validation."
applyTo: "**/schemas/**/*.ts, **/lib/api/**/*.ts"
---

# Zod v4 Conventions

- Import with `import * as z from "zod"`.
- Treat schemas as the source of truth for external contracts. Infer TypeScript
  types with `z.infer<typeof Schema>` instead of maintaining matching
  interfaces.
- Keep schema files organized by domain and re-export schemas/types from a
  small index when that matches the project style.
- Use `safeParse` at runtime boundaries and handle both success and failure.
- In Zod v4, call `z.record(keySchema, valueSchema)` with both key and value
  schemas.
- In Zod v4, use `z.prettifyError()` for user-facing error strings instead of
  v3-style error formatting helpers.
- Use the v4 `error` option for custom messages.
- Use `optional`, `nullable`, and `nullish` deliberately; they encode different
  runtime contracts.
- For API, form, and localStorage validation recipes, use the
  `validate-with-zod` skill.
