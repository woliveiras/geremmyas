# Godot performance

Profile an exported build on target-like hardware. Use the installed version’s profiler, monitors, visual profiler, debugger, and platform tools.

## Common checks

- Inspect script, physics, navigation, animation, particles, canvas draw, lights, and resource loading.
- Reduce unnecessary node processing and signal or polling work.
- Reuse imported resources and detect duplicated mutable resources.
- Measure texture memory, viewport size, overdraw, lights, and particle counts.
- Use servers or lower-level APIs only after profiling proves node overhead material.
- Check shader and resource warm-up when diagnosing stutter.
- Verify web-specific WebAssembly, WebGL, audio, and thread constraints.

## Threads

Read thread-safe API guidance before using threads. Keep scene-tree mutation on supported boundaries, minimize synchronization, and wait for worker completion cleanly.

## Official references

- [Performance](https://docs.godotengine.org/en/stable/tutorials/performance/index.html)
- [Using multiple threads](https://docs.godotengine.org/en/stable/tutorials/performance/using_multiple_threads.html)
- [Optimization using Servers](https://docs.godotengine.org/en/stable/tutorials/performance/using_servers.html)
- [Web export](https://docs.godotengine.org/en/stable/tutorials/export/exporting_for_web.html)
