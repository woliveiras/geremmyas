# Godot AI integration

Inspect the installed Godot version, navigation map and layers, physics layers, scene structure, agent nodes, and avoidance settings.

## Patterns

- Keep decision data in scripts or Resources and scene composition in nodes.
- Use physics queries and navigation APIs through explicit agent adapters.
- Separate NavigationAgent target/path state from the character’s final physics movement.
- Connect navigation and action signals once and clean them up with node lifetime.
- Rate-limit path changes and detect stale or unreachable goals.
- Use `_draw`, gizmos, or dedicated debug nodes for perception and paths without affecting behavior.

## Navigation cautions

Navigation changes may synchronize asynchronously. Verify map readiness and installed-version semantics before querying. Treat avoidance as local steering, not proof that a global path exists.

## Official references

- [2D navigation overview](https://docs.godotengine.org/en/stable/tutorials/navigation/navigation_introduction_2d.html)
- [Using NavigationAgents](https://docs.godotengine.org/en/stable/tutorials/navigation/navigation_using_navigationagents.html)
- [Physics](https://docs.godotengine.org/en/stable/tutorials/physics/index.html)
