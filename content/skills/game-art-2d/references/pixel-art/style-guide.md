# Pixel Art Style Guide

Use this as the default style pack when the project and user provide no other
art direction.

## Visual signature

Build images from deliberate pixel clusters at a chosen logical resolution.
Use crisp edges, controlled stair-stepping, limited ramps, and intentional
silhouettes. Avoid high-resolution paintings with a pixel filter applied.

Strict rules:

- choose a logical canvas before drawing;
- keep pixel density consistent across the asset family;
- disable anti-aliasing and smoothing in final assets and runtime display;
- use nearest-neighbor scaling by whole-number factors;
- avoid isolated noise pixels unless they communicate texture or sparkle;
- keep clusters readable at 1x logical size.

## Resolution and scale

Choose the smallest logical size that supports gameplay:

- 8x8 to 16x16: symbolic pickups, markers, and micro-icons;
- 16x16 to 32x32: simple items, tiles, and compact units;
- 32x32 to 64x64: readable characters, enemies, equipment, and effects;
- 64x64 and above: bosses, portraits, complex props, or detailed environments.

These are starting points, not universal rules. Base the scale on camera,
viewport, and existing assets. Keep character height, tile size, and UI icon
size in a shared scale system.

Generate at a larger multiple when the image tool needs a large canvas, then
reduce to the logical grid and inspect the result. Do not accept inconsistent
pseudo-pixels from the generated source.

## Shape and line

- Begin with silhouette and large clusters.
- Prefer meaningful concave and convex contour changes over noisy outlines.
- Use one outline system across the family: dark external, selective colored,
  light-side omission, or no outline.
- Keep line weight consistent with the logical resolution.
- Use selective internal lines; do not outline every material boundary.
- Avoid accidental tangents and one-pixel notches that disappear in motion.

## Palette and value

Start with 4 to 8 colors for small assets and expand only when materials or
states require it. Build hue-shifted ramps instead of adding unrelated colors.

Assign palette roles:

- deepest outline or occlusion;
- material shadows;
- material midtones;
- lit planes or highlights;
- gameplay accent;
- optional emission, hazard, or status color.

Test in grayscale. Preserve readable value separation between subject,
background, interactive elements, and effects. Avoid excessive near-duplicate
colors.

## Light and material

Lock one baked light direction for a family. Describe materials through cluster
shape and highlight behavior:

- metal: sharp value jumps and selective highlights;
- cloth: broader soft clusters with fewer specular pixels;
- stone: broken planes and restrained texture;
- glass or energy: controlled emission and interior contrast;
- damaged surfaces: larger chips and breaks before micro-noise.

Do not shade every edge with automatic gradients. Use banded or clustered
shading deliberately.

## Characters and objects

- exaggerate gameplay-relevant masses and tools;
- keep faces or primary identity markers readable at target size;
- reserve tiny detail for stable idle or portrait assets;
- use distinct silhouette families for different combat roles;
- keep ground contact and pivots stable;
- ensure held items do not merge ambiguously with limbs or body.

## Environments

- keep tile topology exact;
- reduce contrast and texture behind characters;
- use modular clusters that can repeat without obvious seams;
- align perspective and light direction across tiles;
- separate decorative overlays from collision tiles;
- use landmarks and value grouping to support navigation.

## Animation

- use whole-pixel movement for key masses;
- preserve cluster volume between frames;
- hold key poses when a low frame cadence is part of the style;
- use smears and impact frames intentionally;
- avoid subpixel interpolation and automatic frame blending;
- test foot placement and looping at runtime speed.

## Prompt block

Append only when it matches the project:

```text
Style: authentic production pixel art at <logical width x height>, deliberate
pixel clusters, crisp hard edges, no anti-aliasing, consistent pixel density,
restricted <N>-color palette with coherent hue-shifted ramps, readable
silhouette at 1x, controlled banded shading, fixed <light direction> lighting,
no text, no watermark, no poster layout, no smooth painted gradients, no
high-resolution brush texture, no noisy single-pixel scatter.
```

When generating a large source canvas, add:

```text
Represent every logical pixel as an exact uniform square block at a consistent
integer scale. Do not mix block sizes or add subpixel detail.
```

## Processing

Use `pixel_art_process.py` to create a constrained logical-resolution asset and
an enlarged nearest-neighbor preview. Treat this as cleanup and normalization,
not a substitute for a readable generated silhouette.

## Quality checklist

- logical dimensions match the contract;
- all visible marks align to the logical grid;
- no anti-aliased or semi-transparent fringe remains;
- palette count and ramps match the family;
- pixel density is consistent;
- silhouette reads at 1x;
- outline and light direction match anchors;
- texture supports material instead of adding noise;
- nearest-neighbor runtime sampling is verified;
- animation clusters remain coherent in motion.
