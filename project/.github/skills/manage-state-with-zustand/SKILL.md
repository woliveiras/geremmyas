---
name: manage-state-with-zustand
description: "Zustand v5 recipes for middleware setup, immer mutations, persistence, and XState sync. Use when: zustand middleware, zustand persist, zustand immer, zustand xstate, store patterns."
---

# Manage State with Zustand

Recipes for common Zustand v5 patterns. For core rules (when to use, store structure,
middleware order, selectors), see the zustand instruction file which auto-loads for store files.

## Middleware Setup

### With persistence

```ts
create<State & Actions>()(
  devtools(
    persist(
      immer((set) => ({
        // state
        items: [],
        // actions
        addItem: (item) => set((state) => { state.items.push(item) }),
        reset: () => set(() => ({ items: [] })),
      })),
      {
        name: "app-items",
        partialize: (state) => ({ items: state.items }),
      },
    ),
    { name: "ItemsStore", enabled: import.meta.env.DEV },
  ),
)
```

### Without persistence

```ts
create<State & Actions>()(
  devtools(
    immer((set) => ({
      count: 0,
      increment: () => set((state) => { state.count += 1 }),
    })),
    { name: "CounterStore", enabled: import.meta.env.DEV },
  ),
)
```

## Immer Mutation Recipes

### Update nested item

```ts
updateItem: (id, changes) =>
  set((state) => {
    const item = state.items.find((i) => i.id === id)
    if (item) Object.assign(item, changes)
  }),
```

### Replace whole state (reset)

Return a new object instead of mutating:

```ts
reset: () => set(() => ({ ...DEFAULT_STATE })),
```

### Toggle boolean

```ts
toggleDarkMode: () => set((state) => { state.darkMode = !state.darkMode }),
```

## Syncing with XState

XState owns logic (transitions, guards, side effects).
Zustand is the read cache for components.
Sync in a Provider via subscription, not inside the machine:

```tsx
useEffect(() => {
  const sub = actorRef.subscribe((snapshot) => {
    if (snapshot.matches("authenticated")) {
      useAuthStore.getState().setAuth(snapshot.context.user, snapshot.context.token)
    } else if (snapshot.matches("unauthenticated")) {
      useAuthStore.getState().clearAuth()
    }
  })
  return () => sub.unsubscribe()
}, [actorRef])
```

Key rules:
- Subscribe in a React `useEffect`, not in the Zustand store definition
- Use `getState()` to avoid stale closures
- XState drives state transitions; Zustand reflects them for React consumption
