---
name: procedural-generation-2d
description: "Build deterministic procedural content for Phaser and Godot. Use when generating seeded maps, rooms, terrain, encounters, loot, spawns, or validated layouts. Do not use global randomness or treat structural validity as proof of fun."
---

# Procedural Generation 2D

Generate semantic game data deterministically, validate it, then map it to engine resources and visuals.

## Define the generation contract

1. Inspect engine version, map representation, tile size, collision and navigation rules, content catalogs, and save compatibility.
2. Define inputs: seed, generator version, catalog version, dimensions, difficulty, biome, and required landmarks.
3. Define outputs as semantic layers rather than atlas coordinates.
4. List hard invariants and soft quality goals separately.
5. Choose the simplest algorithm family that meets the topology.

## Load references

- Read [generation-and-validation.md](references/generation-and-validation.md) for deterministic streams, pipeline stages, algorithms, and validators.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-art-2d` for tileset topology, `$game-ai-2d` for encounter behavior, and `$game-testing-2d` for property tests.

## Build a staged generator

1. Generate topology.
2. Reserve critical route, entrance, exit, landmarks, and protected cells.
3. Validate connectivity and clearances.
4. Place gameplay content from semantic catalogs.
5. Add cosmetics with an independent random stream.
6. Emit canonical output plus diagnostics and a stable hash.
7. Fail with actionable reasons or retry under an explicit bounded policy.

## Verify

- Re-run identical inputs and compare canonical hashes.
- Run many seeds and retain the smallest failing seed.
- Verify catalog additions do not perturb unrelated topology.
- Visualize layers, connectivity, sockets, and rejected placements.
- Playtest variety, pacing, readability, and exploitability separately from structural validity.

## Guardrails

- Do not use global randomness.
- Do not let render iteration order affect generation.
- Do not encode gameplay rules in atlas coordinates.
- Do not retry forever or silently accept invalid output.
- Do not claim quality from connectivity tests alone.
