# Phaser AI integration

Inspect installed Phaser APIs, the chosen physics system, tilemap representation, pathfinding library if any, Scene lifecycle, and update ownership.

## Patterns

- Keep decisions and state machines in pure TypeScript.
- Adapt Phaser positions, bodies, tile data, and time into plain observations.
- Update decisions on a bounded cadence while movement can update each physics step.
- Treat third-party pathfinding APIs and grid conventions as versioned dependencies.
- Use Graphics or debug Game Objects for overlays, then destroy them with the Scene.
- Remove event listeners, timers, colliders, and path requests on shutdown or agent destruction.

## Browser constraints

Budget main-thread pathfinding and perception. For workers, pass serializable snapshots and version requests so stale results can be discarded.

## Official references

- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
- [Physics](https://docs.phaser.io/phaser/concepts/physics)
- [Tilemap API](https://docs.phaser.io/api-documentation/class/tilemaps-tilemap)
- [Time](https://docs.phaser.io/phaser/concepts/time)
