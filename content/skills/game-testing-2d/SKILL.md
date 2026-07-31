---
name: game-testing-2d
description: "Design and run deterministic tests for 2D Phaser and Godot games. Use when verifying simulation, movement, combat, collisions, maps, replays, scenes, browsers, or exported builds. Do not use as a substitute for required human playtesting."
---

# Game Testing 2D

Test authoritative rules at the narrowest deterministic boundary, then add engine and runtime evidence where integration matters.

## Start from the failure surface

1. Read repository instructions, the behavior contract, and current tests.
2. Detect engine version, language, runner, physics backend, target platforms, and CI limitations.
3. Reproduce the behavior before changing production code when diagnosing a bug.
4. Classify the needed evidence: pure rule, engine integration, rendered output, exported build, or human play.

## Load references

- Read [testing-strategy.md](references/testing-strategy.md) for test layers, deterministic harnesses, replays, maps, and visual checks.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Combine with `$gameplay-programming-2d` for behavior ownership and `$game-build-and-release` for artifact smoke tests.

## Build the smallest reliable test

1. Control time, random seeds, inputs, and initial state.
2. Assert externally meaningful state instead of incidental implementation details.
3. Step simulation explicitly when possible.
4. Add engine objects only when the contract crosses an engine boundary.
5. Make failures explain the frame, seed, entity, state, and expected invariant.
6. Preserve a minimal regression case for every fixed bug.

## Verify and report

- Run the focused test and nearest practical suite fresh.
- Repeat suspected flaky tests with controlled seeds and environment.
- Verify exports separately from editor or development mode.
- State what automation proved and what still requires visual, audio, controller, performance, or playtesting judgment.

## Guardrails

- Do not use arbitrary sleeps as synchronization.
- Do not call a test deterministic while uncontrolled global time or randomness remains.
- Do not replace gameplay assertions with snapshots alone.
- Do not infer an exported build works because unit tests passed.
- Do not modify a correct test merely to accommodate broken behavior.
