---
description: "Use when creating, editing, or reviewing XState v5 state machines. Covers createMachine vs setup(), onDone/onError event typing, actor patterns, and testing with createActor."
applyTo: "**/*.machine.ts, **/*.machine.test.ts"
---

# XState v5 Conventions

## When to use

Use XState for features with **multiple exclusive states and side effects**:
- Auth flows: `checkingSession → authenticating → authenticated / unauthenticated`
- Multi-step wizards, async resource lifecycles (`idle → loading → ready / failed`)
- Anything with guarded transitions, parallel states, or complex event handling

For simple toggles or preferences, use Zustand instead.

## `createMachine` vs `setup().createMachine()`

**Use `createMachine(config, implementations)` when the machine has `onDone`/`onError`.**

`setup().createMachine()` has a known issue: actions in `setup.actions` receive the event
typed as your union — `DoneActorEvent` is not in that union, so `event.output` is
inaccessible at runtime.

**Canonical pattern — inline `assign` with `unknown` cast:**

```ts
export const myMachine = createMachine(
  {
    states: {
      loading: {
        invoke: {
          src: "fetchData",
          onDone: {
            target: "ready",
            actions: assign(({ event }: { event: unknown }) => {
              const e = event as { output: MyResult }
              return { data: e.output, error: null }
            }),
          },
          onError: {
            target: "failed",
            actions: assign(({ event }: { event: unknown }) => {
              const e = event as { error: unknown }
              return { error: e.error instanceof Error ? e.error.message : "Unknown error" }
            }),
          },
        },
      },
    },
  },
  { actors: { fetchData: fetchDataActor } },
)
```

`setup().createMachine()` is fine for pure event-driven machines with no `invoke`.

## Types

Declare `types` at the top of the machine config for full TypeScript inference:

```ts
export const myMachine = createMachine({
  types: {
    context: {} as MyContext,
    events: {} as MyEvent,
  },
})
```

## Actors

- `fromPromise` — async operations (API calls, async init)
- `fromCallback` — event-based / DOM listeners
- Pass runtime dependencies via `input`, never close over mutable state

## React — select minimum slice

```ts
const isLoading = useSelector(actorRef, (s) => s.matches("loading"))
```

Expose `actorRef` via React context so consumers don't recreate the machine.

## Testing — use `createActor`, avoid state polling

```ts
const actor = createActor(myMachine)
actor.start()
actor.send({ type: "SUBMIT" })
expect(actor.getSnapshot().value).toBe("submitting")
actor.stop()
```

Wait for side effects, not state names: `await vi.waitFor(() => expect(mockFn).toHaveBeenCalled())`

Always call `actor.stop()` after each test.

For full React integration, actor recipes, and testing patterns, use the `model-state-with-xstate` skill.
