# Phaser gameplay integration

Inspect `package.json`, lockfile, Phaser type definitions, scene registration, physics config, and project lifecycle before using an API. Treat installed types as authoritative.

## Boundaries

- Use Scenes as lifecycle and composition boundaries, not as an excuse to place all rules in one class.
- Keep pure gameplay rules in TypeScript modules that can run without WebGL, DOM, or a Phaser Scene.
- Use Scene events, timers, input, cameras, and physics through narrow adapters.
- Create global animations and asset keys once; remove scene-owned listeners on shutdown.
- Distinguish the game registry, scene data, and per-object data by lifetime.

## Time and physics

- Phaser callbacks commonly expose milliseconds; normalize units at module boundaries.
- Choose Arcade Physics for simple body-based 2D motion and Matter only when its shape and constraint model is required.
- Do not depend on callback order without verifying it in the installed version.
- After browser suspension, prevent one giant delta from advancing the simulation unchecked.

## Input

Map keyboard, pointer, touch, and gamepad input to actions such as `move`, `jump`, and `attack`. Capture just-pressed and held semantics separately. Respect browser focus and scene pause.

## Lifecycle checks

Test start, sleep, wake, pause, resume, shutdown, restart, and destroy. Watch for duplicate subscriptions, timers, colliders, animation listeners, and global registry state.

## Official references

- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
- [Input](https://docs.phaser.io/phaser/concepts/input)
- [Data Manager](https://docs.phaser.io/phaser/concepts/data-manager)
- [Time](https://docs.phaser.io/phaser/concepts/time)
- [Physics](https://docs.phaser.io/phaser/concepts/physics)
