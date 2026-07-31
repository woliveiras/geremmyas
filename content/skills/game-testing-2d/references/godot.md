# Testing Godot games

Inspect the installed Godot executable, project language, addons, test framework, import state, and CI display capabilities.

## Test boundaries

- Put pure rules in scripts or data objects that can run without a rendered scene.
- Instantiate the smallest scene that crosses the Node, signal, resource, or physics boundary.
- Step explicit frames or physics ticks where the harness supports it.
- Use stable semantic paths or injected references rather than brittle editor-only lookups.

## Lifecycle cases

Test tree entry and exit, deferred calls, `queue_free`, pause modes, scene reload, Autoload reset, and reused `Resource` instances. Await tree changes rather than sleeping.

## Headless limits

Use `--headless` for suitable CI and command-line checks, but do not claim visual renderer, audio output, controller, or platform integration acceptance from headless results.

## Export smoke test

Create the actual export preset artifact, serve web exports over HTTP, and verify resource loading, start scene, a minimal gameplay route, console output, and clean exit.

## Official references

- [Command line](https://docs.godotengine.org/en/stable/tutorials/editor/command_line_tutorial.html)
- [Running Godot in CI](https://docs.godotengine.org/en/stable/tutorials/editor/command_line_tutorial.html)
- [Exporting projects](https://docs.godotengine.org/en/stable/tutorials/export/exporting_projects.html)
