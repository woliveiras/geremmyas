# Godot gameplay integration

Inspect `project.godot`, the installed Godot minor version, input map, physics settings, scene ownership, resources, and test harness. Do not assume APIs from another Godot minor.

## Boundaries

- Use scenes for cohesive reusable objects and composition.
- Use scripts or `Resource` definitions for rules and authored data.
- Use signals for observable events while keeping ownership and cleanup clear.
- Prefer local nodes to Autoloads; reserve Autoloads for genuinely cross-scene lifetime.
- Keep collision shapes, navigation, sockets, and hurtboxes separate from sprites.

## Time and physics

- Put physics-dependent motion in `_physics_process`.
- Put render-only presentation in `_process`.
- Use the engine’s physics motion methods rather than manually teleporting collision bodies each frame.
- Define pause processing for gameplay, UI, timers, and audio deliberately.

## Input

Define semantic actions in Input Map. Keep gameplay actions separate from built-in UI focus actions. Represent just-pressed, held, strength, and device-specific prompts explicitly.

## Lifecycle checks

Test scene instantiation, tree entry and exit, pause, reload, `queue_free`, signal disconnection, and reused external resources. Avoid mutating a shared `Resource` when per-instance state is intended.

## Official references

- [Godot 2D](https://docs.godotengine.org/en/stable/tutorials/2d/index.html)
- [Input examples](https://docs.godotengine.org/en/stable/tutorials/inputs/input_examples.html)
- [Physics](https://docs.godotengine.org/en/stable/tutorials/physics/index.html)
- [Best practices](https://docs.godotengine.org/en/stable/tutorials/best_practices/index.html)
