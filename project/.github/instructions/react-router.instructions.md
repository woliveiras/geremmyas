---
description: "Use when writing or reviewing React Router v7 routes in framework mode. Covers loaders, actions, typegen, and breaking changes from v6."
applyTo: "**/routes/**/*.tsx, **/routes/**/*.ts, **/app/routes/**/*.tsx"
---

# React Router v7 — Framework Mode

## Route Data Pattern

Data fetching and mutations happen in `loader` and `action` functions, not in components:

```ts
// Route module exports
export async function loader({ params }: Route.LoaderArgs) {
  const book = await getBook(params.id)
  if (!book) throw new Response("Not Found", { status: 404 })
  return { book }
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData()
  await updateBook(formData)
  return { ok: true }
}

export default function BookPage({ loaderData }: Route.ComponentProps) {
  const { book } = loaderData
  return <h1>{book.title}</h1>
}
```

## Type Safety — Typegen

v7 generates route types automatically. Use `Route.*` types from the generated `+types` file:

```ts
import type { Route } from "./+types/book"

export async function loader({ params }: Route.LoaderArgs) { ... }
export default function BookPage({ loaderData }: Route.ComponentProps) { ... }
```

Never manually type loader data — typegen infers it from the loader return.

## Imports

All imports come from `react-router` (not `react-router-dom`):

```ts
import { Link, useNavigate, Form } from "react-router"
```

## File-Based Routing

Routes are defined by file structure in `app/routes/`:

```
app/routes/
├── _index.tsx          → /
├── books.tsx           → /books (layout)
├── books._index.tsx    → /books (index)
├── books.$id.tsx       → /books/:id
└── books.$id.edit.tsx  → /books/:id/edit
```

- `.` separates URL segments: `books.$id` → `/books/:id`
- `_` prefix = pathless layout: `_auth.tsx` wraps without adding URL segment
- `$` prefix = dynamic param: `$id` → `:id`

## Forms and Mutations

Use `<Form>` for mutations — submits to the route's `action`:

```tsx
import { Form } from "react-router"

<Form method="post">
  <input name="title" />
  <button type="submit">Save</button>
</Form>
```

For programmatic mutations, use `useSubmit` or `useFetcher`.

## Error Handling

Export `ErrorBoundary` from the route module:

```tsx
export function ErrorBoundary() {
  const error = useRouteError()
  if (isRouteErrorResponse(error)) {
    return <p>{error.status}: {error.statusText}</p>
  }
  return <p>Something went wrong</p>
}
```

## Navigation

- `<Link to="/books">` for declarative navigation
- `<NavLink>` for active state styling
- `useNavigate()` for programmatic navigation
- Prefer `<Link>` over `useNavigate()` when possible

For migrating from v6, use the `migrate-react-router` skill.
