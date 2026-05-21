---
description: "Use when creating, editing, or reviewing Zustand v5 stores. Covers store structure, middleware ordering, immer mutations, persistence, and selector patterns."
applyTo: "**/store/**/*.ts, **/stores/**/*.ts"
---

# Zustand v5 Conventions

## When to use

- **Shared UI state** multiple components need to read: auth credentials, theme, preferences
- State that survives component unmounts but is not server state

Do NOT use Zustand for:
- Server state → TanStack Query
- Local component state → `useState`/`useReducer`
- URL state → router search params

## Store structure

Split state and actions into separate interfaces:

```ts
interface MyState {
  value: string
  count: number
}

interface MyActions {
  setValue: (v: string) => void
  increment: () => void
  reset: () => void
}

export const useMyStore = create<MyState & MyActions>()(/* middleware */)
```

## Middleware

**Order:** `devtools` outermost → `persist` → `immer` innermost.

Canonical setup with persistence:

```ts
create<State & Actions>()(
  devtools(
    persist(
      immer((set) => ({ /* state + actions */ })),
      {
        name: "app-<store-name>",
        partialize: (state) => ({ key: state.key }),
      },
    ),
    { name: "<StoreName>Store", enabled: import.meta.env.DEV },
  ),
)
```

Without persistence, drop the `persist()` wrapper.

## Persistence rules

- Use `partialize` — never persist derived state or functions
- If data has its own persistence (manual `localStorage`), do NOT also use `persist`
- Keys must be unique per store: `"appname-<store-name>"`

## Selectors

Select the minimum slice to avoid unnecessary re-renders:

```ts
// ✅ Granular
const isAuthenticated = useAuthStore((s) => s.isAuthenticated)

// ❌ Whole store — re-renders on any state change
const store = useAuthStore()
```

Export only the hook (`useMyStore`), not the store object itself.

For middleware variants, immer recipes, and XState sync patterns, use the `manage-state-with-zustand` skill.
