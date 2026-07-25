---
description: "Use when creating, editing, or reviewing Zustand v5 stores. Covers store structure, middleware ordering, immer mutations, persistence, and selector patterns."
applyTo: "**/store/**/*.ts, **/stores/**/*.ts"
---

# Zustand v5 Conventions

- Use Zustand for shared client state that survives component unmounts but is
  not server state, URL state, or local component state.
- Do not use Zustand for server cache; use TanStack Query. Do not use it for
  URL state; use router search params.
- Split state and actions into separate TypeScript interfaces, then compose the
  store type.
- Export the hook, not the raw store object, unless the project has an explicit
  non-React integration boundary.
- Middleware order: `devtools` outermost, then `persist`, then `immer`
  innermost.
- Use `partialize` for persisted stores; never persist functions, derived
  values, or sensitive data.
- Use unique persistence keys per store.
- Select the minimum state slice in components to avoid broad rerenders.
- If XState owns transitions, let Zustand mirror state for rendering; do not put
  XState transition logic inside the store.
- For middleware variants, Immer recipes, and XState sync patterns, use the
  `manage-state-with-zustand` skill.
