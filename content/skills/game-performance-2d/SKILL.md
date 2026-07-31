---
name: game-performance-2d
description: "Profile and optimize 2D Phaser and Godot games against explicit budgets. Use when diagnosing FPS, stutter, loading, memory, draw calls, tilemaps, particles, or web exports. Do not use for speculative optimization."
---

# Game Performance 2D

Optimize measured bottlenecks against an explicit platform budget.

## Establish the budget and baseline

1. Identify target hardware, renderer, resolution, browser or export target, and required frame rate.
2. Reproduce the representative scene in a production-like build.
3. Capture CPU, GPU, memory, loading, network, and frame-time evidence as applicable.
4. Record engine version, build flags, content seed, viewport, and object counts.
5. Rank bottlenecks by player impact and measured cost.

## Load references

- Read [profiling-and-budgets.md](references/profiling-and-budgets.md) for measurement, bottleneck classification, and optimization order.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-build-and-release` when production build configuration or asset delivery is the boundary.

## Optimize one bottleneck

1. State the hypothesis and metric expected to change.
2. Apply the smallest architectural or content change that targets it.
3. Preserve gameplay, visual, and accessibility contracts.
4. Re-run the identical benchmark and compare percentiles, not a single FPS reading.
5. Remove instrumentation that should not ship or gate it behind debug settings.

## Verify

- Compare before and after on target-like hardware.
- Test worst-case content, not only an empty scene.
- Check memory growth across repeated scene changes.
- Verify both average frame time and spikes.
- Report tradeoffs, quality changes, and untested targets.

## Guardrails

- Do not optimize from intuition alone.
- Do not pool cheap objects without proving allocation is material.
- Do not trade deterministic behavior or correctness for small gains silently.
- Do not use editor performance as proof of exported-build performance.
- Do not move work to threads before checking API thread safety and transfer cost.
