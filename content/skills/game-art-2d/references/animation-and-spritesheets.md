# Animation and Spritesheets

Use this reference for sprite cycles, combat actions, environmental loops,
effects, and atlas preparation.

## Use an anchor-first workflow

1. Approve one in-game source frame.
2. Lock its silhouette, proportions, palette, materials, view, and anchor.
3. Define the action's timing and key poses.
4. Generate or draw the whole ordered strip in one request.
5. Slice raw slots into individual frames.
6. Normalize all frames with one shared scale and anchor.
7. Review a contact sheet and the animation in motion.
8. Pack only the approved frames.

Generating frames independently increases character and style drift. Use it
only when the user accepts manual cleanup.

## Plan motion before frames

Describe gameplay timing, not only pose names:

- anticipation;
- action;
- contact or peak;
- recoil or follow-through;
- recovery;
- return to idle.

For loops, ensure the last frame transitions cleanly to the first. For attacks,
identify the exact contact frame so gameplay timing can remain code-owned.

Use fewer strong frames before adding in-betweens. A readable four-frame action
is more useful than a smooth but ambiguous twelve-frame action.

## Lock the strip contract

Specify:

- exact frame count;
- row and column order;
- source slot dimensions;
- final cell dimensions;
- facing direction;
- shared baseline or pivot;
- whether overflow is allowed for weapons or effects;
- loop, hold, and one-shot behavior;
- intended frame duration or timing table;
- whether frame 1 must remain byte-for-byte identical to the anchor.

Keep characters, weapons, shadows, and effects on separate layers or sheets when
the runtime must recolor, swap, time, or reuse them independently.

## Prompt a full strip

```text
Runtime asset: ordered <state> animation strip
Anchor image: preserve this exact character identity and rendering style
Layout: exactly <N> frames, one row, left to right, equal empty slots
Action beats: <frame-by-frame pose intent>
Invariants: same view, facing, proportions, palette, materials, outfit, line
weight, pixel density, and ground contact
Canvas: generous padding, removable flat chroma background
Avoid: labels, dividers, scenery, camera changes, duplicated limbs, cropped
motion, per-frame rescaling, shadows unless specified
```

Do not ask an image model to render timing labels inside the sheet.

## Normalize frames

After generation:

1. remove the background and validate alpha;
2. slice cells without trimming away intentional overflow prematurely;
3. inspect each cell for duplicate or missing poses;
4. calculate one scale from the largest visible frame;
5. apply that scale to every frame;
6. align a shared anchor, usually bottom-center;
7. preserve equal cell dimensions;
8. pack and write frame metadata;
9. render a contact sheet;
10. validate the cycle in-engine.

Use `slice_spritesheet.py`, `pack_spritesheet.py`, and
`make_contact_sheet.py` for the deterministic steps.

## Evaluate motion

Check:

- silhouette changes communicate the action without interior detail;
- feet or contact points do not slide unintentionally;
- volume and limb count remain stable;
- frame-to-frame scale does not pulse;
- equipment follows a plausible arc;
- attack contact aligns with gameplay events;
- hurt and defeat poses do not resemble active or idle states;
- loop seams do not pop;
- VFX do not obscure essential hit or telegraph information;
- the motion reads at actual game speed and display scale.

Do not infer animation quality from a static sheet alone.

## Pixel animation

For pixel art:

- move important masses by deliberate whole-pixel increments;
- keep cluster shapes coherent between frames;
- avoid automatic subpixel interpolation;
- preserve a stable outline and palette;
- use smears, impact frames, or selective deformation intentionally;
- export and display with nearest-neighbor sampling.

Do not add more frames merely to imitate high-frame-rate animation. Match the
frame count and cadence to the chosen pixel-art language.
