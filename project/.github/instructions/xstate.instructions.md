---
description: "Use when creating, editing, or reviewing XState v5 state machines. Covers createMachine vs setup(), onDone/onError event typing, actor patterns, and testing with createActor."
applyTo: "**/*.machine.ts, **/*.machine.test.ts"
---

# XState v5 Conventions

- Use XState when the feature has multiple exclusive states, guarded
  transitions, side effects, parallel states, or complex event handling. Use a
  simpler store for toggles and preferences.
- Use `createMachine(config, implementations)` when invoked actors need
  `onDone` or `onError` handling. Keep `setup().createMachine()` for pure
  event-driven machines or when its typing fits the machine.
- Declare machine `types` at the top of the config for context and events.
- Use `fromPromise` for async operations and `fromCallback` for event-based or
  subscription-style actors.
- Pass runtime dependencies through actor `input`; do not close over mutable
  module state.
- In React, expose one actor reference through context and use `useSelector` to
  subscribe to the smallest needed state slice.
- In tests, create machines with `createActor`, send events, assert snapshots or
  side effects, and always stop the actor.
- For full React integration and actor recipes, use the
  `model-state-with-xstate` skill.
