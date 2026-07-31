# Phaser performance

Profile the production bundle in the target browsers. Separate JavaScript, Phaser update, render submission, GPU, asset download, decode, and garbage collection.

## Common checks

- Inspect object counts, active physics bodies, colliders, particles, lights, text, masks, cameras, and render textures.
- Avoid per-frame allocations, array rebuilding, string formatting, and repeated pathfinding in hot loops.
- Disable or sleep offscreen systems whose rules permit it.
- Prefer texture atlases and stable asset keys where they improve batching and delivery.
- Verify Canvas and WebGL only if both are supported targets.
- Use browser Performance and Memory tooling with source maps for diagnosis.
- Test resize, DPR, hidden-tab recovery, and mobile thermal constraints.

## Workers

Use workers for sufficiently large serializable computations such as generation or pathfinding. Version jobs, discard stale results, and measure transfer overhead.

## Official references

- [Game Objects](https://docs.phaser.io/phaser/concepts/gameobjects)
- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
- [Loader](https://docs.phaser.io/phaser/concepts/loader)
- [Device](https://docs.phaser.io/phaser/concepts/device)
