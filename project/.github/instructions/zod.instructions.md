---
description: "Use when creating, editing, or reviewing Zod v4 schemas. Covers schema definition, type inference, validation patterns, common v4 API differences, and runtime validation."
applyTo: "**/schemas/**/*.ts, **/lib/api/**/*.ts"
---

# Zod v4 Conventions

## Types are inferred — never duplicated

```ts
import * as z from "zod"

export const UserSchema = z.object({
  id: z.number(),
  username: z.string(),
  role: z.string(),
})

// ✅ Infer type from schema
export type User = z.infer<typeof UserSchema>

// ❌ Never define a separate matching interface
export interface User { id: number; username: string; role: string }
```

## Import style

```ts
import * as z from "zod"
```

## Zod v4 — key differences from v3

### `z.record` requires two arguments
```ts
// ✅ v4
z.record(z.string(), z.unknown())

// ❌ v3-only (fails silently or errors in v4)
z.record(z.unknown())
```

### Error formatting
```ts
// ✅ v4
const result = schema.safeParse(data)
if (!result.success) {
  const message = z.prettifyError(result.error) // user-facing string
}

// ❌ v3 method (removed in v4)
result.error.format()
```

### Custom error messages use `error` key
```ts
// ✅ v4
z.string().min(1, { error: "Required" })

// ⚠️ v3 style (deprecated in v4)
z.string().min(1, { message: "Required" })
```

## Common field modifiers

```ts
z.string().optional()   // string | undefined
z.string().nullable()   // string | null
z.string().nullish()    // string | null | undefined
```

## Schema file layout

One schema file per domain, types inferred from schemas, exported via barrel:

```
src/lib/schemas/
├── user.schema.ts
├── book.schema.ts
└── index.ts        # re-exports all schemas and types
```

Do not maintain a parallel `types/` folder — schemas are the single source of truth for types.

For validation recipes (API clients, forms, localStorage), use the `validate-with-zod` skill.
