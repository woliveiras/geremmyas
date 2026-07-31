# Generation and validation

## Determinism contract

Persist:

- master seed;
- generator algorithm version;
- catalog or content version;
- engine version if engine RNG or ordering affects results;
- dimensions and generation options;
- canonical output hash.

Use an application-owned PRNG when results must remain stable across engine or runtime upgrades.

## Independent random streams

Derive labeled streams such as:

- topology;
- rooms;
- hazards;
- enemies;
- loot;
- props;
- cosmetics.

Adding a cosmetic option must not change the critical route or encounters.

Sort candidate coordinates and semantic IDs before random selection. Never depend on hash-map, scene-tree, object, or filesystem iteration order.

## Algorithm selection

- Use graph or room grammar generation for authored pacing and routes.
- Use BSP for rectilinear room subdivision.
- Use cellular automata for cave-like fields, followed by connectivity repair.
- Use drunkard walks for organic corridors with bounded coverage.
- Use noise for continuous fields, not as a complete level-design guarantee.
- Use Wave Function Collapse for local adjacency constraints; add global connectivity validation.
- Combine algorithms only through explicit staged contracts.

## Layered output

Emit semantic layers for ground, walls, hazards, props, foreground, collision, navigation, spawns, triggers, and metadata. Resolve semantic IDs to engine-specific tiles or scenes only at the integration boundary.

## Hard validators

Check:

- entrance reaches exit;
- required landmarks exist;
- critical path length and width;
- player and enemy clearances;
- protected cells remain free;
- spawn reachability and visibility rules;
- no unknown semantic IDs;
- layer dimensions and bounds;
- collision and navigation agreement;
- bounded generation time and attempts.

## Soft metrics

Measure repetition, branching, dead ends, encounter spacing, landmark distribution, density, backtracking, and visual rhythm. Use these metrics to flag human review, not to manufacture certainty.

## Failure artifacts

Save seed, versions, options, stage, reason, canonical data, and a layer visualization. Minimize failing maps when practical.
