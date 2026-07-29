# Pop Art Style Guide

Use this pack only when the user or project selects pop art. Do not mix it into
the default pixel-art direction implicitly.

## Visual signature

Use bold graphic silhouettes, flat high-contrast color regions, assertive
outlines, selective halftone texture, and energetic composition. Keep the
result as original game art rather than an imitation of a named artist,
character, comic panel, or trademarked visual system.

## Resolution and scale

Choose enough resolution for clean contours and halftone control. For runtime
sprites, simplify shapes aggressively before reducing them. For UI panels,
portraits, and backgrounds, preserve safe areas for code-owned text and
controls.

## Shape and line

- use large simple masses and readable gesture;
- apply strong outer contours and selective interior lines;
- vary line weight intentionally to create hierarchy;
- use bursts, diagonals, and cropped shapes only when they do not harm gameplay
  readability;
- keep repeated assets consistent in contour weight and shape exaggeration.

## Palette and value

- use a compact set of saturated primaries or project-defined accents;
- pair vivid colors with stable dark and light anchors;
- reserve the most intense contrast for gameplay-critical subjects;
- avoid making every surface equally saturated;
- keep semantic colors consistent across enemies, hazards, rewards, and UI.

## Light and material

Favor graphic value blocks over realistic rendering. Use halftone or hatch
texture selectively for shadow, depth, or emphasis. Keep texture scale
consistent and prevent it from producing shimmer when sprites move or scale.

## Characters and objects

Exaggerate pose, silhouette, tools, expressions, and impact shapes. Preserve
identity through contour and color blocking rather than fine material detail.
Keep props readable without surrounding caption boxes or decorative copy.

## Environments

Use simplified depth planes and controlled pattern fields. Keep the gameplay
plane quieter than focal characters and effects. Isolate decorative bursts,
speed lines, or panels so the runtime can time or disable them.

## Animation

Favor strong key poses, held frames, snap, stretch, graphic impact shapes, and
limited smear frames. Keep gameplay timing in code. Avoid dense halftone motion
that flickers or obscures hit information.

## Prompt block

```text
Style: original pop-art-inspired 2D game art, bold graphic silhouette, strong
controlled outlines, flat high-contrast color blocks, compact saturated
palette, selective halftone shadow texture, energetic but runtime-readable
composition, clean isolated subject, no words, no speech balloons, no logos, no
watermark, no imitation of a named artist or protected comic property.
```

## Quality checklist

- silhouette reads at runtime size;
- outline weight is consistent across the family;
- saturation and contrast preserve gameplay hierarchy;
- halftone scale does not shimmer;
- graphic accents do not resemble UI or interaction markers accidentally;
- no generated text, logos, captions, or protected visual identifiers remain;
- effects and backgrounds leave critical play space readable.
