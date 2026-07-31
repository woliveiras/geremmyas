# Godot build and release

Inspect the installed Godot and export-template versions, `project.godot`, `export_presets.cfg`, addons, target renderer, and platform credentials.

## Version control boundary

- Track `export_presets.cfg` when it contains nonsecret release configuration.
- Do not track `.godot/export_credentials.cfg` or other credentials.
- Track source assets and import sidecars; do not track the generated `.godot` cache.

## Build

Use named export presets and the repository’s exact Godot executable. In CI use headless export where appropriate. Ensure templates match the engine version.

## Web

Keep generated companion filenames consistent. Serve with correct WASM MIME type and compression. If threads or extensions require cross-origin isolation, verify the actual host headers. Test the chosen renderer and browser matrix.

## Verify

Run the exported artifact, not only the editor project. Inspect engine and browser logs, resource loading, input, audio, save paths, window or canvas sizing, pause, and a short gameplay route.

## Official references

- [Exporting projects](https://docs.godotengine.org/en/stable/tutorials/export/exporting_projects.html)
- [Command line](https://docs.godotengine.org/en/stable/tutorials/editor/command_line_tutorial.html)
- [Exporting for the Web](https://docs.godotengine.org/en/stable/tutorials/export/exporting_for_web.html)
