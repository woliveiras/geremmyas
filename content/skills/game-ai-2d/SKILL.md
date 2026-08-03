---
name: game-ai-2d
description: "Use when enemy or NPC behavior needs perception, decision state machines, navigation, pathfinding, groups, encounters, or bosses in 2D Phaser or Godot. Do not use for player controls, procedural maps, UI, saves, or test strategy."
---

# Game AI 2D

Build readable AI from explicit perception, decision, navigation, action, and memory boundaries.

## Establish the AI contract

1. Inspect engine version, gameplay architecture, navigation system, collision layers, and existing enemy behavior.
2. Define goals, observable information, allowed actions, transition priority, timing, and fairness constraints.
3. Separate authored archetype data from mutable agent state.
4. Choose the simplest decision model that expresses the behavior.

## Load references

- Read [ai-architecture.md](references/ai-architecture.md) for model selection, perception, navigation, encounters, and testing.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$gameplay-programming-2d` for combat and movement integration and `$procedural-generation-2d` for spawn and map validity.

## Implement one observable behavior

1. Compute perception from explicit queries and memory.
2. Make a bounded decision at a documented cadence.
3. Request navigation or steering without letting it decide combat truth.
4. Execute an action with entry, completion, interruption, and cleanup.
5. Emit debug state and reason codes.
6. Keep expensive work distributed or budgeted.

## Verify

- Test transitions with synthetic observations.
- Test unreachable targets, lost targets, pause, death, despawn, and scene reload.
- Use fixed seeds for randomized choices.
- Stress multiple agents and measure decision and navigation cost.
- Playtest readability, telegraphing, fairness, and exploitability.

## Guardrails

- Do not give AI access to information the design marks hidden.
- Do not recompute paths or expensive perception every rendered frame without evidence.
- Do not mix animation completion with authoritative action completion silently.
- Do not add behavior-tree complexity to a small finite-state problem.
