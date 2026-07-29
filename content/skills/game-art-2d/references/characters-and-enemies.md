# Characters and Enemies

Use this reference for protagonists, companions, NPCs, recruits, ordinary
enemies, elites, and bosses.

## Start from gameplay

Record the subject's gameplay role before appearance:

- movement and locomotion;
- attack range and direction;
- defensive or evasive behavior;
- relative size and threat;
- required held items, equipment, or attach points;
- damage, status, and defeat states;
- viewing distance and target display size.

Make gameplay role readable in silhouette. Do not depend on small surface detail
to distinguish fast, armored, ranged, support, or boss units.

## Build the canonical design

Produce these artifacts in order when the scope needs them:

1. silhouette candidates at target size;
2. approved neutral pose in the runtime view;
3. limited turnaround only for required directions;
4. expression or damage states only when visible in game;
5. material and palette breakdown;
6. equipment or modular attachment rules;
7. animation anchors.

Keep the neutral runtime pose as the canonical anchor. Concept art may support
the design but must not replace it.

## Preserve identity

Lock high-value identity features:

- overall height and width ratio;
- head-to-body ratio;
- center of mass;
- face, visor, eyes, or primary readable marker;
- costume or chassis silhouette;
- dominant and accent colors;
- left/right asymmetry;
- equipment placement;
- outline and pixel density.

Repeat these invariants in every edit and animation prompt. Use the canonical
image as a reference.

## Design enemies as a readable family

Share two or three family traits, then vary gameplay traits:

- common material, construction method, faction marking, or energy color;
- distinct silhouette for each combat role;
- escalating scale, ornament, damage, or energy for elites and bosses;
- consistent weak-point and attack-telegraph language.

Avoid palette swaps as the only distinction unless the runtime intentionally
uses them and accessibility remains acceptable.

## Handle directions

Generate only directions the game uses:

- side-view games often need one facing plus a runtime mirror;
- top-down games may need four or eight directions;
- front-facing autobattlers may need a single combat view;
- isometric games must keep the same projection and ground angle.

Do not mirror a design with meaningful asymmetry unless the game accepts the
swap. Preserve weapon hand, damage, text, and emblem orientation intentionally.

## Plan states

Choose the smallest state set that communicates gameplay:

- idle;
- locomotion;
- primary action or attack;
- active ability;
- hurt or impact;
- disabled, destroyed, defeated, or despawn;
- selected, targeted, buffed, debuffed, or interactable overlays when needed.

Keep overlays and reusable effects separate from the base character where the
runtime can composite them.

## Prompt pattern

```text
Runtime asset: canonical <character/enemy> sprite
Gameplay role: <role and threat>
View and facing: <camera and direction>
Target logical size: <width x height>
Silhouette: <large readable masses>
Identity invariants: <proportions, face, costume/chassis, asymmetry>
Materials: <family rendering rules>
Palette: <shared family palette plus subject accents>
Pose: <neutral or named gameplay state>
Anchor: <usually bottom-center>
Background: removable flat chroma or approved opaque scene
Avoid: pose montage, labels, scenery, cast shadow, extra equipment, cropped body
```

Validate the result at target size and on at least one common bright and dark
game background.
