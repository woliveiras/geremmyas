# Phaser persistence

Inspect target browsers or wrappers, current storage adapter, serialization library, lifecycle events, and installed Phaser version.

## Patterns

- Keep save code independent of Phaser Scenes and Game Objects.
- Use plain validated data; the game registry is runtime state, not durable storage by itself.
- Wrap `localStorage`, IndexedDB, platform SDKs, or filesystem bridges behind one asynchronous interface.
- Account for quota, privacy mode, blocked storage, eviction, multiple tabs, and JSON parse failures.
- Save on explicit stable events; browser unload callbacks are not a reliable sole strategy.
- Separate synchronous web APIs from the game loop to avoid avoidable stalls.

## Testing

Use an in-memory adapter for rules, fake failure adapters for error paths, and real browser tests for persistence, reload, storage denial, and multi-version fixtures.

## Official references

- [Data Manager](https://docs.phaser.io/phaser/concepts/data-manager)
- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
- [MDN Web Storage](https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API)
- [MDN IndexedDB](https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API)
