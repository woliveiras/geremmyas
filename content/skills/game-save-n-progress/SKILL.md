---
name: game-save-n-progress
description: "Implement durable saves and progression for Phaser and Godot. Use when building slots, settings, checkpoints, migrations, recovery, autosave, or storage adapters. Do not use for engine-object serialization or client tamper resistance."
---

# Game Save N Progress

Persist a small, versioned domain snapshot; validate and migrate it before applying it to runtime state.

## Establish the persistence contract

1. Inspect engine version, target platforms, existing save format, storage APIs, progression systems, and repository security rules.
2. Classify data as settings, profile progression, run state, checkpoint, cache, or telemetry.
3. Define stable IDs, schema version, compatibility window, autosave triggers, and recovery behavior.
4. Identify sensitive, platform-owned, or externally synchronized data.

## Load references

- Read [save-architecture.md](references/save-architecture.md) for schemas, migrations, write safety, slots, checkpoints, and tests.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-testing-2d` for migration and corruption tests and `$game-build-and-release` for platform acceptance.

## Implement the pipeline

1. Extract a plain domain snapshot from runtime state.
2. Serialize with a top-level schema version and metadata.
3. Write through a platform adapter using the safest available replace strategy.
4. Read raw data without mutating the game.
5. Parse, validate, migrate stepwise, and normalize defaults.
6. Apply only validated data to a fresh runtime state.
7. Preserve or quarantine recoverable failures according to policy.

## Verify

- Round-trip every supported save kind.
- Load oldest supported fixtures through every migration.
- Test truncated, malformed, unknown-version, and semantically invalid saves.
- Test autosave interruption, repeated writes, slot deletion, and no-space or unavailable-storage behavior where possible.
- Verify the real exported target separately.

## Guardrails

- Do not serialize scene nodes, sprites, callbacks, or engine object references.
- Do not treat a TypeScript type or GDScript annotation as runtime validation.
- Do not overwrite the only recoverable save before a new write succeeds.
- Do not promise tamper resistance for client-owned saves.
- Do not add cloud synchronization without an explicit conflict policy.
