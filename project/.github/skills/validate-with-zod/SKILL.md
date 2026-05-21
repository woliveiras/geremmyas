---
name: validate-with-zod
description: "Generate Zod validation patterns for common use cases. Use when: validating API responses, form validation, localStorage parsing, building schema wrappers, Zod recipes, validation patterns."
---

# Validate with Zod

Generate Zod validation code for common use cases. Use the project's existing
Zod conventions (see `zod.instructions.md`) and adapt patterns to the specific need.

## When to Use

- Setting up API response validation
- Adding form validation to a component
- Parsing data from localStorage or external sources
- Building reusable schema wrappers

## Patterns

### API Client Validation

Validate responses at the boundary with a generic fetch wrapper:

```ts
async function fetchAPI<T>(
  path: string,
  options?: RequestInit,
  schema?: z.ZodSchema<T>,
): Promise<T> {
  const data = await res.json()
  if (schema) return schema.parse(data) // throws ZodError if shape is wrong
  return data as T                       // backwards-compatible for unvalidated endpoints
}

// Usage
fetchAPI("/user/me", undefined, UserSchema)
```

### Generic Response Wrappers

```ts
export function PaginatedResponseSchema<T extends z.ZodTypeAny>(itemSchema: T) {
  return z.object({
    records: z.array(itemSchema),
    totalRecords: z.number(),
  })
}
```

### Form Validation

```ts
const loginSchema = z.object({
  username: z.string().min(1, { error: "Required" }).transform(s => s.trim()),
  password: z.string().min(1, { error: "Required" }),
})

const result = loginSchema.safeParse({ username, password })
if (!result.success) {
  setError(z.prettifyError(result.error))
  return
}
await login(result.data.username, result.data.password)
```

### localStorage / External Data Parsing

```ts
const keySchema = z.string().min(1)

function getStoredKey(): string | null {
  const raw = localStorage.getItem("my-key")
  const result = keySchema.safeParse(raw)
  return result.success ? result.data : null
}
```

### Discriminated Unions

```ts
const ResultSchema = z.discriminatedUnion("type", [
  z.object({ type: z.literal("success"), data: z.string() }),
  z.object({ type: z.literal("error"),   message: z.string() }),
])
```

## Procedure

1. Ask the user which pattern they need (or detect from context)
2. Read existing schemas in `**/schemas/**` to match naming and style
3. Generate the validation code following the patterns above
4. Place schemas in the project's schema directory
5. Infer types from schemas — never create parallel type definitions
