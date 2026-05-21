---
name: migrate-react-router
description: "Guide migration from React Router v6 to v7 framework mode. Use when: migrating routes, upgrading react-router, v6 to v7, converting to framework mode."
---

# Migrate React Router v6 to v7

Step-by-step guide for migrating a React Router v6 codebase to v7 framework mode.

## When to Use

- Upgrading an existing app from React Router v6 to v7
- Converting classic `<BrowserRouter>` setup to framework mode
- Systematic migration of route files

## Breaking Changes Reference

### Imports: `react-router` replaces `react-router-dom`

```ts
// ❌ v6
import { Link, useNavigate, Form } from "react-router-dom"

// ✅ v7
import { Link, useNavigate, Form } from "react-router"
```

### Removed APIs

| v6 | v7 Replacement |
|---|---|
| `json()` helper | Return plain objects from loaders |
| `defer()` | Use native promises and `<Awaited>` |
| `<BrowserRouter>` (classic mode) | Framework mode with file-based routes |
| `useLoaderData()` hook | `loaderData` prop on component |

### File-Based Routing

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

### Loader Data Access

```tsx
// ❌ v6
export default function BookPage() {
  const data = useLoaderData()
  return <h1>{data.book.title}</h1>
}

// ✅ v7
export default function BookPage({ loaderData }: Route.ComponentProps) {
  const { book } = loaderData
  return <h1>{book.title}</h1>
}
```

### Typegen

v7 generates route types automatically. Use `Route.*` types:

```ts
import type { Route } from "./+types/book"

export async function loader({ params }: Route.LoaderArgs) { ... }
export default function BookPage({ loaderData }: Route.ComponentProps) { ... }
```

## Migration Procedure

1. **Update dependencies**: `npm install react-router@latest` (remove `react-router-dom`)
2. **Fix imports**: search and replace `react-router-dom` → `react-router` across all files
3. **Remove `json()` calls**: return plain objects from loaders/actions
4. **Remove `defer()` calls**: use native promises
5. **Convert route definitions**: move from `<Route>` components to file-based routes in `app/routes/`
6. **Update components**: replace `useLoaderData()` with `loaderData` prop
7. **Add typegen**: use `Route.*` types from generated `+types` files
8. **Run tests**: verify all routes load correctly
9. **Remove old router setup**: delete `<BrowserRouter>` / `createBrowserRouter` configuration

## Rules

- Migrate one route at a time — don't batch all routes in a single commit
- Run tests after each route migration
- Keep `ErrorBoundary` exports in every route module
