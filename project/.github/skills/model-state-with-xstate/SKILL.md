---
name: model-state-with-xstate
description: "XState v5 recipes for React integration, testing with createActor, and actor patterns. Use when: xstate react, testing state machine, xstate actors, machine provide, useSelector, useActorRef."
---

# Model State with XState

Recipes for common XState v5 patterns. For core rules (createMachine vs setup, types),
see the xstate instruction file which auto-loads for `*.machine.ts` files.

## React Integration

### Provider pattern — create actor once, share via context

```tsx
import { createActorContext } from "@xstate/react"

const MyMachineContext = createActorContext(myMachine)

// Provider
function App() {
  return (
    <MyMachineContext.Provider>
      <MyFeature />
    </MyMachineContext.Provider>
  )
}

// Consumer — select minimum slice
function MyFeature() {
  const isLoading = MyMachineContext.useSelector((s) => s.matches("loading"))
  const data = MyMachineContext.useSelector((s) => s.context.data)
  const actorRef = MyMachineContext.useActorRef()

  return <button onClick={() => actorRef.send({ type: "RETRY" })}>Retry</button>
}
```

### Manual approach (without createActorContext)

```tsx
import { useActorRef, useSelector } from "@xstate/react"

function MyFeature() {
  const actorRef = useActorRef(myMachine)
  const isLoading = useSelector(actorRef, (s) => s.matches("loading"))
  const error = useSelector(actorRef, (s) => s.context.error)

  actorRef.send({ type: "RETRY" })
}
```

Expose `actorRef` via React context so consumers don't recreate the machine.

## Actor Recipes

### fromPromise — async operations

```ts
const fetchDataActor = fromPromise(async ({ input }: { input: { id: string } }) => {
  const response = await fetch(`/api/items/${input.id}`)
  if (!response.ok) throw new Error("Fetch failed")
  return response.json() as Promise<Item>
})
```

### fromCallback — event-based / DOM listeners

```ts
const resizeActor = fromCallback(({ sendBack }) => {
  const handler = () => sendBack({ type: "RESIZE", width: window.innerWidth })
  window.addEventListener("resize", handler)
  return () => window.removeEventListener("resize", handler)
})
```

Pass runtime dependencies via `input`, never close over mutable state:

```ts
invoke: {
  src: "fetchData",
  input: ({ context }) => ({ id: context.itemId }),
}
```

## Testing Patterns

### Basic test with createActor

```ts
const actor = createActor(myMachine)
actor.start()
actor.send({ type: "SUBMIT" })
expect(actor.getSnapshot().value).toBe("submitting")
actor.stop()
```

### Override actors in tests with machine.provide()

```ts
const testMachine = myMachine.provide({
  actors: {
    fetchData: fromPromise(async () => mockResult),
  },
  guards: {
    isValid: () => true,
  },
})
const actor = createActor(testMachine)
```

### Async assertions — avoid polling on state name

`vi.waitFor(() => snapshot.matches("X"))` can miss transient states.

Wait for observable side effects, then check context:

```ts
await vi.waitFor(() => expect(mockFn).toHaveBeenCalled())
await vi.waitFor(() => {
  expect(actor.getSnapshot().context.data).not.toBeNull()
})
```

Always call `actor.stop()` after each test.
