---
name: game-art-2d
description: "Design, generate, animate, package, and validate production art for 2D games. Use when Codex creates or audits sprites, tilesets, maps, backgrounds, VFX, or UI assets for Phaser or Godot. Do not use for 3D or marketing art."
---

# 2D Game Art Studio

Create cohesive 2D game art as runtime assets, not isolated illustrations.
Combine Codex image generation with deterministic local processing and in-engine validation.

## Operating boundaries

- Work only in 2D. Treat isometric art as 2D when the runtime consumes sprites or
  tiles; route meshes, rigs, materials, and 3D scenes elsewhere.
- Prefer the project's established art direction over this skill's defaults.
- Default to pixel art only when the project and user provide no style.
- Distinguish concept art from shippable art. Never present a concept sheet,
  poster composition, or arbitrary AI grid as a runtime-ready asset.
- Distinguish map art from level design. Keep collision, traversal, spawn,
  trigger, encounter, and navigation data explicit in the game or level format.
- Create original work. Do not copy protected characters, maps, interfaces, or
  distinctive living-artist styles. Translate references into general visual
  traits.
- Preserve existing assets. Use versioned sibling filenames unless the user
  explicitly requests replacement.

## Choose the work mode

Use one or more modes, but keep the current deliverable explicit:

1. **Direction**: define a visual language, style guide, palette, shape grammar,
   camera, scale, and asset constraints.
2. **Production**: generate or edit canonical assets, variants, animation
   frames, tiles, backgrounds, or interface art.
3. **Integration**: normalize, pack, import, configure, and validate assets in
   the target runtime.
4. **Audit**: inspect an existing asset set for coherence, technical defects,
   missing states, readability, or runtime risk.

## Core workflow

### 1. Inspect the project

Read the nearest agent instructions, art docs, engine configuration, asset
manifests, source art, runtime consumers, and naming conventions. Inspect
representative images with `view_image`. Record:

- engine and renderer;
- camera and projection;
- target viewport and display scale;
- existing art style and palette;
- expected source and shipping formats;
- sprite cell, tile, atlas, pivot, and animation conventions;
- performance, accessibility, and content constraints.

If no project exists, infer only reversible defaults and ask about a choice only
when it materially changes the result.

### 2. Classify the asset job

Classify the request before generating:

- character or enemy;
- animation or spritesheet;
- item, prop, pickup, equipment, or icon;
- effect or environmental animation;
- tileset, modular environment kit, or background;
- map composition or playable level;
- HUD, menu, portrait, badge, or other UI art.

Read only the matching task reference:

- Art direction and coherence: [art-direction.md](references/art-direction.md)
- Characters and enemies:
  [characters-and-enemies.md](references/characters-and-enemies.md)
- Animation and spritesheets:
  [animation-and-spritesheets.md](references/animation-and-spritesheets.md)
- Maps, tilesets, and levels:
  [maps-and-levels.md](references/maps-and-levels.md)
- Items, props, VFX, and UI:
  [items-props-vfx-and-ui.md](references/items-props-vfx-and-ui.md)
- Runtime packaging and QA:
  [runtime-packaging-and-qa.md](references/runtime-packaging-and-qa.md)

When integration or engine-specific QA is in scope, detect the installed engine
version and read exactly one matching reference:

- Phaser: [asset-pipeline.md](references/phaser/asset-pipeline.md)
- Godot Engine: [asset-pipeline.md](references/godot/asset-pipeline.md)

Do not assume Phaser 3 and 4 or different Godot 4 releases expose identical
APIs. Use the project's installed version and its matching official docs as the
source of truth.

### 3. Select the style pack

Read the project's style guide first. Otherwise select a bundled pack:

- Default: [pixel art](references/pixel-art/style-guide.md)
- Optional: [pop art](references/pop-art/style-guide.md)

To add a new style, follow
[style-pack-contract.md](references/style-pack-contract.md) and link it directly
from this section. Do not mix style packs unless the user requests a hybrid.

### 4. Lock the asset contract

Normalize the request into a compact contract. Include only applicable fields:

```text
Runtime job: <how the game uses it>
Engine/version: <for example Phaser 4.x or Godot 4.x>
Asset family: <character/enemy/item/tileset/background/UI/VFX>
View: <side/top-down/front/three-quarter/isometric>
Target size: <logical pixels or runtime dimensions>
Display scale: <for example 1x, 2x, or 4x nearest-neighbor>
Grid/layout: <cell size, frame count, rows, columns, tile size>
Anchor/pivot: <for example bottom-center>
Silhouette and shape language: <readability constraints>
Palette and value range: <existing palette or limit>
Material and texture language: <surface rules>
Animation states: <idle/walk/attack/hurt/etc.>
Background: <opaque, transparent, or removable chroma>
Invariants: <identity, proportions, palette, facing, equipment>
Avoid: <artifacts, unwanted elements, unsafe similarities>
Output: <source, final format, naming, destination>
```

For a family of assets, create one contract for the family and a short delta for
each member. Keep shared invariants identical.

### 5. Produce the smallest useful vertical slice

Create one representative shippable asset before producing a full set:

- one canonical character pose before animation or costume variants;
- one enemy plus its gameplay silhouette test before the full faction;
- one item icon at final display size before the full inventory;
- one terrain transition set before a full tileset;
- one playable room or map segment before a full level;
- one animation cycle before every state.

Obtain user approval when a direction choice is expensive to propagate. Reuse
the approved asset as the anchor for the rest of the family.

### 6. Generate or edit with Codex

Use the installed `imagegen` skill and its built-in `image_gen` path. Do not ask
for an API key and do not switch to its CLI/API fallback unless the user
explicitly requests or approves that path.

- Use one built-in call per distinct asset or variant.
- Use reference images to preserve style, identity, proportions, materials, and
  palette.
- Treat an existing canonical asset as an edit/reference anchor, not loose
  inspiration.
- For animation, generate a complete ordered strip from one approved anchor
  frame. Avoid independent frame generation unless the user accepts drift.
- For cutouts, follow `imagegen`'s built-in-first chroma-key and local removal
  workflow.
- Ask for no text, watermark, scenery, decorative border, or presentation
  layout unless the runtime asset actually needs it.
- Preserve the asset contract verbatim across iterations. Change one targeted
  variable at a time.

Always inspect generated output with `view_image` before accepting it.

### 7. Post-process deterministically

Use the bundled scripts for repeatable operations. They require Pillow. In
Codex Desktop, call the workspace dependency loader and use the Python path it
returns. Otherwise run a script with `uv run`; each script declares its own
Pillow dependency.

```bash
GAME_ART_SKILL="/absolute/path/to/game-art-2d"
uv run "$GAME_ART_SKILL/scripts/asset_qa.py" <asset.png>
```

- Inspect technical properties:
  `scripts/asset_qa.py`
- Convert a high-resolution source into constrained logical pixel art:
  `scripts/pixel_art_process.py`
- Slice a regular generated strip into source frames:
  `scripts/slice_spritesheet.py`
- Normalize and pack ordered frames with one shared scale and anchor:
  `scripts/pack_spritesheet.py`
- Render a labeled review sheet:
  `scripts/make_contact_sheet.py`

Use explicit output paths and keep raw, working, and final assets separate.
Never overwrite the only source image.

### 8. Integrate only the selected finals

When integration is in scope:

- copy selected finals into the project's established asset location;
- update manifests, preloaders, atlas metadata, animation definitions, tilemap
  references, or content catalogs at the narrowest boundary;
- preserve nearest-neighbor sampling for pixel art;
- keep gameplay metadata outside raster files;
- remove discarded variants and temporary generation artifacts unless the user
  asks to retain them.

Do not claim integration when only image files were generated.

### 9. Validate at runtime scale

Validate both the art and the asset contract:

- silhouette and focal point read at actual display size;
- palette, lighting, pixel density, outline, and material language match the
  family;
- identity and proportions remain stable across variants and frames;
- alpha edges are clean and no chroma fringe remains;
- cells, pivots, frame order, timing, and loops are correct;
- tiles meet without seams and transitions cover required neighbors;
- interactive objects differ from decoration by shape, value, color, or motion;
- maps preserve traversal and combat readability;
- atlases fit runtime texture and memory budgets;
- the asset is inspected in-engine when the project can run.

Report generated and modified paths, the final asset contract or prompt, local
processing performed, validation evidence, and any manual art judgment still
open.

## Default output conventions

Follow project conventions when present. Otherwise use:

```text
art/
  raw/       # generated or imported sources
  working/   # separated frames and processing intermediates
  final/     # selected source-controlled deliverables
```

Use stable lowercase kebab-case names:

```text
<family>-<subject>-<state>-<direction>-<frame-or-version>.<ext>
```

Examples:

- `robot-caretaker-idle-south-01.png`
- `enemy-scrap-drone-attack-east-sheet.png`
- `item-circuit-core-icon-v2.png`
- `tileset-foundry-floor-32.png`

Keep source and shipping formats explicit. Prefer PNG for alpha sprites and
lossless pixel art; use WebP only after verifying lossless settings and engine
support.
