# Testing Phaser games

Inspect the installed Phaser, test runner, bundler, DOM emulator, and browser-test configuration.

## Test boundaries

- Put simulation and rule modules outside Scene classes so they run in the normal TypeScript test runner.
- Represent input as tick-indexed actions rather than synthesizing DOM events for rule tests.
- Wrap Scene, physics, camera, audio, and storage dependencies behind small adapters.
- Use real browser automation for canvas, WebGL, focus, pointer, gamepad, audio unlock, resizing, and asset loading.

## Lifecycle and browser cases

Test Scene restart, shutdown, sleep, wake, pause, and destroy. Assert that listeners, colliders, timers, tweens, and DOM elements do not multiply. Test hidden-tab recovery and abnormal deltas.

## Visual harness

Fix canvas size, scale mode, renderer, DPR, seed, camera, and animation frame. Expose a deterministic ready signal before capture. Prefer debug overlays for origins, colliders, and tile boundaries.

## Export smoke test

Serve the production bundle over HTTP. Verify base paths, hashed assets, dynamic imports, service workers, browser console errors, resize behavior, and a short playable route.

## Official references

- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
- [Game loop](https://docs.phaser.io/phaser/concepts/game)
- [Scale Manager](https://docs.phaser.io/phaser/concepts/scale-manager)
- [Loader](https://docs.phaser.io/phaser/concepts/loader)
