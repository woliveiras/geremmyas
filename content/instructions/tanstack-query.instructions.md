---
description: "Use when writing or reviewing TanStack Query v5 hooks, queries, or mutations. Focuses on v5 API differences and project-safe hook patterns."
applyTo: "**/hooks/**/*.ts, **/hooks/**/*.tsx, **/api/**/*.ts, **/queries/**/*.ts"
---

# TanStack Query v5 Conventions

- Use the single options object API:
  `useQuery({ queryKey, queryFn })` and `useMutation({ mutationFn })`.
- Do not use removed v4 overloads like `useQuery(["books"], fetchBooks)`.
- Do not put `onSuccess`, `onError`, or `onSettled` on `useQuery`; react to
  `data`, `error`, and status flags in the component instead.
- `onSuccess`, `onError`, and `onSettled` remain valid on `useMutation`.
- Prefer custom hooks around queries and mutations; components should not repeat
  raw query setup.
- Query keys are arrays with stable, descriptive segments.
- Invalidate by query key after successful mutations.
- Use Suspense-specific hooks (`useSuspenseQuery`,
  `useSuspenseInfiniteQuery`, `useSuspenseQueries`) when Suspense is desired.

## v5 Renames

| v4 name | v5 replacement |
| --- | --- |
| `cacheTime` | `gcTime` |
| `keepPreviousData` | `placeholderData: keepPreviousData` |
| `isLoading` first-load meaning | `isPending` plus `isFetching` |
| `isInitialLoading` | `isLoading` (isPending + isFetching) |
