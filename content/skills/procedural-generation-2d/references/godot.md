# Godot procedural integration

Inspect the installed Godot version, TileSet and TileMapLayer resources, scene ownership, navigation layers, physics layers, and resource formats.

## Patterns

- Keep topology and placement in deterministic scripts or data objects.
- Use stable semantic IDs until the final TileSet or scene-catalog adapter.
- Populate separate TileMapLayer nodes for ground, walls, hazards, foreground, and metadata when their ordering or lifecycle differs.
- Batch map changes where supported by the installed version.
- Create physics, navigation, and gameplay objects from semantic layers, not visible pixels.
- Keep editor tooling and runtime generation on the same versioned contract.

## Determinism

Do not assume engine RNG output remains stable across upgrades unless the project locks that version. Avoid scene-tree and dictionary iteration as ordering inputs.

## Official references

- [Using TileMaps](https://docs.godotengine.org/en/stable/tutorials/2d/using_tilemaps.html)
- [Using TileSets](https://docs.godotengine.org/en/stable/tutorials/2d/using_tilesets.html)
- [Random number generation](https://docs.godotengine.org/en/stable/tutorials/math/random_number_generation.html)
- [2D navigation](https://docs.godotengine.org/en/stable/tutorials/navigation/navigation_introduction_2d.html)
