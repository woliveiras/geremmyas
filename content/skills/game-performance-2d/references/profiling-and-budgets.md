# Profiling and budgets

## Define budgets

Record:

- target frame time and percentile;
- CPU and GPU split;
- maximum decoded texture and audio memory;
- startup and level-load targets;
- download size;
- active entity, particle, light, and audio counts;
- draw calls or canvas operations;
- supported minimum hardware.

At 60 FPS the entire frame has about 16.67 ms; engine, browser, OS, and game systems share it.

## Classify the bottleneck

- **CPU simulation:** scripts, physics, AI, pathfinding, allocation.
- **CPU render submission:** traversal, batching, text, state changes.
- **GPU:** fill rate, overdraw, shaders, lights, large render targets.
- **Memory:** textures, audio decode, duplicate resources, retained scenes.
- **Loading:** network, decompression, imports, shader compilation, scene construction.
- **Frame pacing:** GC, synchronous I/O, asset creation, compilation, background-tab recovery.

## Measurement protocol

Use a fixed build, seed, route, viewport, duration, warm-up, and capture window. Record median and slow percentiles. Keep before-and-after traces.

## Optimization order

1. Remove unnecessary work.
2. Reduce update frequency or scope.
3. Cull invisible or distant work.
4. Batch or combine compatible work.
5. Cache stable results.
6. Reduce content cost with an explicit quality decision.
7. Pool only measured high-churn objects.
8. Parallelize only safe, sufficiently large tasks.
9. Use lower-level engine APIs only after higher-level options are exhausted.

## Regression checks

Verify visual output, collision, AI cadence, audio, input latency, loading, memory after repeated transitions, and behavior on low-end targets.

## Reporting

Include baseline, test environment, profiler evidence, change, new measurement, variance, tradeoffs, and remaining bottlenecks.
