# Godot game-feel integration

Inspect the installed Godot version, Input Map, physics ticks, camera nodes, animation ownership, process modes, and project stretch settings.

## Patterns

- Read gameplay input through named Input Map actions.
- Apply physics movement in `_physics_process`.
- Put squash, flash, particles, and presentation recoil on child visuals rather than collision roots.
- Use `AnimationPlayer`, Tween, particles, and camera nodes according to existing project conventions.
- Keep hit-stop ownership explicit; decide which nodes, timers, audio, and UI continue processing while gameplay pauses.
- Restore time scale and process modes on interruption, death, scene change, and editor reload.

## Pixel-art checks

Verify viewport scaling, camera smoothing, pixel snap, integer zoom, and final transforms together. Nearest filtering alone does not prevent shimmer.

## Official references

- [Input handling](https://docs.godotengine.org/en/stable/tutorials/inputs/index.html)
- [2D movement](https://docs.godotengine.org/en/stable/tutorials/2d/2d_movement.html)
- [Animation](https://docs.godotengine.org/en/stable/tutorials/animation/index.html)
- [2D](https://docs.godotengine.org/en/stable/tutorials/2d/index.html)
