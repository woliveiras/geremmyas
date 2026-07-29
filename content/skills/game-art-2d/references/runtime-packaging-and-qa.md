# Runtime Packaging and QA

Use this reference when preparing assets for an engine, auditing technical
quality, or claiming that art is shippable.

## Keep source stages distinct

Maintain three conceptual stages even if the project uses different folders:

- **Raw**: generated, scanned, imported, or artist-authored sources.
- **Working**: separated layers, frames, masks, cleanup, and variants.
- **Final**: selected runtime assets and metadata.

Do not make runtime code depend on files in raw or temporary generation
directories.

## Choose formats intentionally

- Use PNG for alpha sprites, pixel art, and lossless intermediate work.
- Use lossless WebP only after checking engine support and nearest-neighbor
  behavior.
- Use lossy WebP or JPEG for large opaque backgrounds only after visual
  comparison at runtime size.
- Avoid JPEG for pixel art, masks, UI, sprites, or sharp alpha-adjacent detail.
- Preserve editable source formats outside the runtime bundle when they exist.

## Validate dimensions and sampling

Check:

- exact image dimensions;
- atlas and GPU texture limits;
- tile and frame divisibility;
- transparent padding and extruded atlas borders;
- power-of-two requirements only when the target runtime actually needs them;
- logical size versus displayed size;
- nearest-neighbor sampling for pixel art;
- no fractional transforms that blur logical pixels;
- consistent pivots and baselines.

Use `asset_qa.py` before integration.

## Validate alpha

For cutout assets:

- confirm an alpha channel exists;
- confirm background pixels are actually transparent;
- inspect bright and dark checkerboards;
- remove green or magenta edge spill;
- preserve intentional holes and interior transparency;
- avoid semi-transparent halos around hard-edged pixel art;
- check premultiplied-alpha expectations in the engine.

When using the `imagegen` chroma-key helper, inspect the result rather than
assuming a successful command produced clean edges.

## Validate sprites and animation

Check:

- every expected frame exists exactly once;
- cells share dimensions;
- frame order and naming match runtime metadata;
- anchors do not drift;
- frames do not pulse in scale;
- loops and one-shots use correct timing;
- event frames align with code-owned gameplay events;
- mirrored directions preserve intended asymmetry;
- atlases include adequate padding to prevent texture bleeding.

Use a contact sheet for curation and a running animation for final judgment.

## Validate tiles, maps, and levels

Check:

- tile boundaries and transitions;
- no seams under camera movement and scaling;
- collision and navigation data remain separate and correct;
- decoration does not change gameplay affordances;
- occlusion layers do not hide critical actors;
- repeated textures and landmarks look intentional;
- seeded maps reproduce the same layout and asset selections when required.

## Validate performance

Measure or inspect:

- decoded texture memory, not only compressed file size;
- atlas count and dimensions;
- draw calls and texture swaps;
- animation frame count and active effects;
- overdraw from large transparent quads;
- loading and caching behavior;
- mobile or low-end limits when in scope.

Do not optimize by degrading every asset uniformly. Reduce unused resolution,
duplicate frames, unnecessary alpha, and oversized transparent bounds first.

## Validate accessibility and comfort

- do not encode critical state by color alone;
- test foreground/background contrast in real scenes;
- provide reduced motion, reduced shake, or reduced flash paths when applicable;
- keep rapid flashes within the project's safety policy;
- ensure UI art supports focus, large text, and high-contrast states;
- avoid visual noise that conceals hazards or targets.

## Evidence for completion

Report:

- final asset and metadata paths;
- target dimensions, formats, grid, and pivots;
- commands used for deterministic processing;
- automatic checks and their results;
- engine or browser scene used for runtime inspection;
- unresolved manual art direction or animation judgment.

Generation success is not runtime acceptance. State clearly when in-engine,
device, accessibility, or human art review remains open.
