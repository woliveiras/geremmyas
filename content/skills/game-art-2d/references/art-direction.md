# Art Direction and Coherence

Use this reference when establishing, extending, or auditing the visual language
of a 2D game.

## Build a visual system

Define decisions that can be repeated across hundreds of assets:

- **Player fantasy**: the emotion and identity the art must reinforce.
- **Camera**: side view, top-down, front, three-quarter, or isometric.
- **Logical scale**: character height, tile size, item icon size, and UI scale.
- **Shape grammar**: dominant geometric, organic, soft, sharp, heavy, or fragile
  forms.
- **Value hierarchy**: which elements remain legible in grayscale and under
  effects.
- **Palette roles**: world neutrals, friendly accents, hostile accents,
  interactable accents, hazards, and UI feedback.
- **Outline system**: none, selective, internal, external, colored, or
  value-shifted.
- **Material language**: how metal, cloth, stone, vegetation, energy, glass, and
  damage are represented.
- **Lighting model**: baked direction, ambient-only, rim light, emissive rules,
  and runtime lighting assumptions.
- **Detail budget**: amount and frequency of texture at each asset size.
- **Motion language**: restrained, elastic, heavy, mechanical, snappy, or
  theatrical.

Prefer measurable rules such as "characters use 24-32 logical pixels of height"
over mood-only statements such as "retro but modern."

## Create anchor assets

Choose a small set of approved assets to serve as visual truth:

1. one protagonist or frequently seen unit;
2. one ordinary enemy;
3. one interactive prop or item;
4. one environment slice;
5. one representative effect or UI state.

Use anchors as references for every later generation. Do not rely on prose alone
to preserve identity or rendering style.

## Design for gameplay readability

Assign visual hierarchy before decoration:

- reserve the strongest contrast for gameplay-critical subjects;
- distinguish friend, enemy, hazard, reward, and decoration without requiring
  text;
- keep floor or background values quieter than characters and interactables;
- preserve enough negative space around hitboxes and interaction points;
- ensure color is not the only signal;
- test at target size, in motion, and against real backgrounds.

## Maintain family consistency

For every asset family, lock:

- view and perspective;
- baseline or pivot;
- pixel density or brush density;
- light direction;
- outline weight;
- palette family;
- material treatment;
- proportion range;
- amount of wear, noise, and ornament.

Describe each new member as a delta from its family anchor. Regenerate or edit
the smallest inconsistent region instead of reinventing the asset.

## Audit an existing style

Create an inventory and compare representative assets. Look for:

- mixed camera angles or projections;
- inconsistent pixel density or resolution;
- value ranges that make characters disappear into backgrounds;
- incompatible outlines or edge softness;
- arbitrary palette growth;
- repeated silhouettes across different gameplay roles;
- lighting that changes direction between assets;
- materials rendered with unrelated texture languages;
- UI art that belongs to a different visual world.

Separate technical defects from subjective direction choices. Recommend a
canonical rule and list which assets would need migration before changing them.

## Write prompts from rules

Translate art direction into observable prompt language:

```text
Style system: <medium and rendering rules>
Shape language: <dominant silhouettes and proportions>
Palette roles: <base, accent, danger, emissive>
Lighting: <direction, contrast, baked/runtime boundary>
Detail budget: <density at target size>
Consistency anchors: <reference image roles>
Gameplay read: <what must be noticed first>
Avoid: <specific drift and presentation artifacts>
```

Do not pad prompts with incompatible style adjectives. Every adjective must
change an observable property.
