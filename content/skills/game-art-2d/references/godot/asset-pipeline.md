# Godot Engine 2D Asset Pipeline

Use this reference when producing, importing, or auditing art for a Godot 4
project. Inspect `project.godot`, the executable version, project settings, and
existing `.tscn`, `.tres`, and `<asset>.import` files before changing the
pipeline. Use official documentation matching the installed release.

## Contents

- Respect the import pipeline
- Choose the Godot resource form
- Configure pixel art
- Integrate sprites and animation
- Store reusable resources
- Build TileSets and maps
- Handle tile seams and depth
- Validate in Godot
- Official sources

## Respect the import pipeline

Place runtime source assets inside the project and reference them through
`res://`. Godot imports supported source files automatically:

- keep the source image;
- keep the generated `<asset>.import` sidecar because it stores import
  configuration;
- do not edit or commit generated `.godot/imported/` cache artifacts;
- use ResourceLoader or normal resource references instead of addressing
  imported `.ctex` files directly;
- reimport through the editor or supported command-line workflow after changing
  import settings.

Apply shared import settings to a selected family or through project import
defaults. Do not assume replacing a PNG automatically preserves the intended
filter, compression, mipmap, or repeat settings.

## Choose the Godot resource form

- **Sprite2D**: one texture, an atlas region, or a regular sheet addressed with
  `hframes`, `vframes`, and `frame`.
- **AnimatedSprite2D + SpriteFrames**: named animations built from separate
  textures or sheet regions, with loop mode, speed, and relative frame
  durations.
- **Sprite2D + AnimationPlayer**: use when one timeline must coordinate sprite
  frames with other node properties or events.
- **AtlasTexture**: use a region from a larger texture while preserving a
  resource boundary.
- **TileSet + TileMapLayer**: use atlas or scene-backed tiles with terrain,
  physics, navigation, occlusion, and custom data as needed.

Choose one representation per asset family and follow existing scene
conventions. Do not create parallel `SpriteFrames` and `AnimationPlayer`
definitions for the same state without a coordination reason.

## Configure pixel art

Inspect project and node-level texture filtering before changing it. Set the
Canvas texture filter to nearest at the appropriate project, parent, or node
scope. Keep overrides intentional.

Also verify:

- mipmaps are disabled for pixel sprites that are displayed at integer
  magnification and should remain exact;
- lossless import is used for pixel art and hard alpha;
- repeat is enabled only for textures designed to tile;
- the camera, viewport, stretch settings, and transforms preserve the desired
  logical pixel grid;
- `Sprite2D.centered` and offsets do not place odd-sized sprites between pixels;
- pixel-snap project settings are evaluated when transforms produce deformation;
- 2D lighting resolution matches the style, because nearest texture filtering
  alone does not make lights and shadows pixelated.

Do not enable global settings without checking non-pixel UI, backgrounds, and
other texture families.

## Integrate sprites and animation

For a regular sheet:

1. verify exact cell dimensions and frame count;
2. create or update a `SpriteFrames` resource;
3. add frames in explicit order;
4. name animations by stable gameplay state;
5. set loop mode, FPS, and relative frame duration from the asset contract;
6. map bottom-center or other pivots through `centered`, `offset`, or scene-node
   placement consistently;
7. keep hit, contact, spawn, and footstep timing in code or an AnimationPlayer
   call track when the project owns events there.

Use `Sprite2D` frame coordinates or `AnimationPlayer` only when that is the
established project pattern.

Keep visual bounds, node origin, collision shape, navigation agent, weapon
socket, and effect origin separate. A trimmed frame must not move a stable
collision or gameplay anchor.

## Store reusable resources

Save reusable `SpriteFrames`, `TileSet`, materials, and other shared resources
as external `.tres` or `.res` files when several scenes consume them. Use text
resources when reviewable diffs matter and the project accepts them.

Avoid embedding duplicate copies of a resource in many scenes. Preserve stable
resource paths when content or scripts reference them.

## Build TileSets and maps

Use `TileSetAtlasSource` for regular atlas textures. Lock:

- `TileSet.tile_size`;
- atlas `texture_region_size`;
- margins and separation;
- texture padding;
- terrain sets and neighbor rules;
- alternative tiles;
- animated tile layout and timing;
- physics, navigation, occlusion, and custom data layers.

Use multiple `TileMapLayer` nodes when background, terrain, metadata, props,
foreground, collision, or navigation need separate ordering, visibility, or
lifecycle. Treat the older `TileMap` node as version-sensitive; current Godot 4
documentation deprecates it in favor of `TileMapLayer`.

Keep procedural layout deterministic:

- generate semantic tile or content IDs before atlas coordinates;
- sort candidates before seeded selection;
- keep topology, hazards, props, and cosmetic variation on separate RNG streams;
- version the generator and content catalog;
- map semantic IDs to TileSet source, atlas coordinate, and alternative ID at
  the integration boundary;
- verify identical seeds reproduce identical semantic maps.

Do not let decoration or TileSet repacking change collision or route topology.

## Handle tile seams and depth

- keep atlas cells aligned to integer regions;
- enable or preserve texture padding when required to avoid lines between
  tiles;
- validate margins and separation after every atlas layout change;
- use Y-sort only when the projection and scene structure require it;
- set texture origins and occlusion deliberately for tall tiles;
- test camera movement, zoom, stretch, and viewport scaling;
- verify terrain transitions on stress maps, not only in the TileSet preview.

For isometric art, align TileSet shape, layout, offset axis, texture origin, and
Y-sort policy across every layer.

## Validate in Godot

Before calling the asset integrated:

- confirm source and `.import` files are present while `.godot/` remains
  generated;
- reimport and check the Output panel for warnings;
- inspect filtering, compression, mipmaps, repeat, and alpha;
- verify `SpriteFrames` order, names, loops, FPS, and frame durations;
- verify centered state, offset, node position, scale, and flip;
- inspect collision, navigation, occlusion, sockets, and gameplay anchors;
- run every animation in the target scene;
- traverse TileMapLayer boundaries and terrain transitions;
- test the project's supported renderer and export target;
- capture a runtime screenshot at the actual logical resolution.

## Official sources

- Import process:
  https://docs.godotengine.org/en/stable/tutorials/assets_pipeline/import_process.html
- Sprite animation:
  https://docs.godotengine.org/en/stable/tutorials/2d/2d_sprite_animation.html
- Sprite2D:
  https://docs.godotengine.org/en/stable/classes/class_sprite2d.html
- AnimatedSprite2D:
  https://docs.godotengine.org/en/stable/classes/class_animatedsprite2d.html
- SpriteFrames:
  https://docs.godotengine.org/en/stable/classes/class_spriteframes.html
- TileMapLayer:
  https://docs.godotengine.org/en/stable/classes/class_tilemaplayer.html
- TileSet:
  https://docs.godotengine.org/en/stable/classes/class_tileset.html
- TileSetAtlasSource:
  https://docs.godotengine.org/en/stable/classes/class_tilesetatlassource.html

Switch `stable` to the installed release path when version-specific behavior
matters.
