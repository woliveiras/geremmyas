# Phaser procedural integration

Inspect Phaser and TypeScript versions, tilemap format, physics setup, asset catalogs, and whether generation runs on the main thread or a worker.

## Patterns

- Keep generation in pure TypeScript independent of Scenes and rendering.
- Return serializable semantic data.
- Convert semantic layers into Phaser Tilemaps, Game Objects, physics bodies, and navigation data in one adapter.
- Use explicit tileset names and properties when importing Tiled data.
- Version worker jobs and discard stale results after scene changes.
- Build debug scenes that color routes, regions, spawns, hazards, and validation failures.

## Browser constraints

Measure generation, data transfer, tilemap construction, texture loading, and collider creation separately. Yield or use a worker only when measurement justifies the complexity.

## Official references

- [Tilemap API](https://docs.phaser.io/api-documentation/class/tilemaps-tilemap)
- [Loader](https://docs.phaser.io/phaser/concepts/loader)
- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
