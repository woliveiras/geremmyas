# Maps, Tilesets, and Levels

Use this reference for environment art, modular kits, backgrounds, tilesets,
map composition, and playable levels.

## Separate the deliverables

Treat these as different artifacts:

- **Concept environment**: establishes mood, scale, landmarks, and materials.
- **Background**: a rendered layer behind gameplay.
- **Tileset or modular kit**: reusable pieces with exact dimensions and seams.
- **Map art**: the visual composition of a location.
- **Level data**: playable structure, collision, traversal, encounters, spawns,
  triggers, goals, and pacing.

Never use a single generated environment image as level data. Reconstruct
playable geometry and semantics explicitly.

## Start with the level contract

Record:

- camera and projection;
- target viewport and overscan;
- tile or module size;
- movement grid and collision precision;
- character footprint and clearance;
- supported terrain and elevation;
- path width, jump arc, or tactical range constraints;
- interaction, hazard, reward, and exit language;
- parallax and foreground layers;
- map or level file format;
- deterministic seed requirements.

For procedural levels, keep layout generation deterministic. Treat art as
content selected by layout rules, never as the source of randomness or
connectivity.

## Build environments in layers

Use explicit layers when applicable:

1. far background or sky;
2. distant architecture or terrain;
3. gameplay background;
4. terrain or floor;
5. collision and traversal metadata;
6. props and interactables;
7. characters and effects;
8. foreground occluders;
9. lighting, fog, color grade, or post-process.

Keep collision and interaction data independent from decorative art.

## Design a tileset

Define the smallest complete topology before generation:

- solid center and edges;
- outer and inner corners;
- transitions between material families;
- isolated pieces and end caps;
- slopes, stairs, ladders, doors, or platforms when used;
- damage, hazard, blocked, or alternate states;
- overlay decals that should not multiply base tiles.

Use exact logical dimensions. Preserve projection, light direction, pixel
density, and edge topology across every tile.

Validate:

- dimensions are multiples of tile size;
- transparent padding and extrusion follow engine requirements;
- edges meet without unintended seams;
- transitions cover actual neighbor combinations;
- repeating texture is not obvious at common camera distances;
- decorative variation does not change collision expectations.

## Compose a playable level

Use this sequence:

1. block out traversal and encounters with plain geometry;
2. test player scale, movement, combat space, sight lines, and pacing;
3. establish landmarks and route hierarchy;
4. assign environment modules and tiles;
5. add props without obscuring navigation;
6. add lighting, VFX, and foreground dressing;
7. validate the level with real gameplay;
8. optimize atlases, layers, draw calls, and texture memory.

For maps that represent a route graph rather than physical space, prioritize
node hierarchy, reachability, current position, available choices, hazards, and
progress over environmental realism.

## Generate environment art

For backgrounds:

```text
Runtime asset: 2D <background/parallax layer>
Projection and camera: <view>
Viewport and crop: <dimensions and safe area>
Gameplay area: <required negative space and contrast>
Layer role: <far/mid/gameplay/foreground>
Material and style system: <rules>
Lighting and palette: <world palette>
Tiling or continuation: <horizontal/vertical/none>
Avoid: characters, UI, text, fake collision edges, unintended focal clutter
```

For tiles or modules:

```text
Runtime asset: modular <tileset/environment kit>
Logical tile size: <width x height>
Topology: <required centers, edges, corners, transitions>
Projection: <orthographic/top-down/side/isometric>
Shared invariants: <palette, light direction, outline, pixel density>
Layout: exact regular grid with generous slot separation
Background: removable flat chroma
Avoid: perspective drift, labels, shadows crossing tile boundaries, incomplete
edges, arbitrary merged tiles
```

Expect to curate and reconstruct generated tile families. Do not assume a model
produced mathematically seamless or topologically complete tiles without tests.

## Validate readability

At runtime size, confirm:

- the player and enemies separate from the floor and background;
- safe ground, hazards, blocked routes, exits, and interactables are distinct;
- foreground layers do not hide critical action;
- landmarks support navigation;
- decorative noise does not imitate pickups, attacks, or UI;
- parallax does not cause motion discomfort or false depth cues;
- colors retain meaning under common color-vision deficiencies.
