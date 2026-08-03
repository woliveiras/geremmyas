---
name: gameplay-programming-2d
description: "Use when implementing 2D player movement, combat rules, abilities, interactions, physics, inventory, or scene flow in Phaser or Godot. Do not use for enemy AI, procedural generation, UI, persistence, art, audio, or release work."
---

# Gameplay Programming 2D

Build gameplay as explicit, testable simulation rules with engine code at the boundary.

## Start from project evidence

1. Read the nearest repository instructions and existing gameplay code.
2. Detect the engine, exact installed version, language, physics backend, target platforms, and test setup.
3. Identify the current ownership of input, simulation, presentation, collision, and persistence.
4. Preserve established conventions unless they cause the reported problem.
5. Define the requested behavior with measurable states, transitions, timing, and failure cases.

## Load only relevant references

- Read [shared-foundations.md](references/shared-foundations.md) for architecture, time, state, collision, combat, and contracts.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-testing-2d` for a test strategy, `$game-feel-2d` for responsiveness and feedback, and `$game-art-2d` for visual assets.

## Implement a vertical behavior

1. Express gameplay inputs as intentions, not device keys.
2. Keep authoritative state independent of sprites, animation frames, and visible bounds.
3. Choose update ownership and time domain explicitly.
4. Implement one complete behavior from input or stimulus through state change to presentation.
5. Keep physics bodies, hitboxes, hurtboxes, sockets, and navigation as gameplay data.
6. Make one-shot events idempotent and define interruption rules.
7. Clean up subscriptions, timers, tweens, nodes, and pooled objects at lifecycle boundaries.

## Verify

- Test state transitions, frame-rate independence, collision boundaries, interruptions, and repeated entry.
- Exercise low and high frame rates plus pause, resume, scene reload, and object destruction.
- Inspect the behavior in a running build at the target resolution and input devices.
- Report automated evidence separately from subjective play-feel acceptance.

## Guardrails

- Do not derive damage or collision from opaque pixels.
- Do not let animations silently own gameplay truth.
- Do not mix milliseconds and seconds or render and physics clocks.
- Do not introduce a global manager when local scene ownership is sufficient.
- Do not rewrite the architecture before proving the existing boundary cannot support the behavior.
