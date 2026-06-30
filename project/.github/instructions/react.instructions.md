---
description: "Use when writing or reviewing React components. Covers functional components, hooks, state management boundaries, and composition."
applyTo: "**/*.tsx,**/*.jsx"
---

# React Conventions

## Components
- Functional components with hooks — no class components
- Named exports (no default exports)
- Explicit `Props` interface for every component
- Keep components focused — split large components into smaller ones

## State Management
- **Server state**: TanStack Query (see `tanstack-query.instructions.md`)
- **UI state**: `useState` for local, Zustand for shared (see `zustand.instructions.md`)
- **URL state**: React Router params and loaders (see `react-router.instructions.md`)
- **Complex flows**: XState (see `xstate.instructions.md`)

## Hooks
- Custom hooks for data fetching: `useBooks()`, `useCreateBook()`
- Wrap TanStack Query calls in custom hooks — never call `useQuery` directly in components
- Invalidate queries after mutations

## Accessibility
- Semantic HTML elements (`<button>`, `<nav>`, `<main>`, not `<div onClick>`)
- `aria-*` attributes where semantic HTML is insufficient
- Keyboard navigation support
- Proper heading hierarchy (`h1` → `h2` → `h3`)

## Patterns
- Colocate related files: `BookCard.tsx` + `BookCard.test.tsx` + `useBookData.ts`
- Feature-sliced structure for large apps: `features/featureName/{api,components,hooks,types}/`
- API calls through a shared client (axios or fetch wrapper)
