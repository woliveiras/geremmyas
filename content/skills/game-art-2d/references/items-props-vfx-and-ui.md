# Items, Props, VFX, and UI

Use this reference for inventory art, equipment, pickups, environmental props,
effects, portraits, badges, HUD elements, and menu illustration.

## Design items and props by function

Classify the object:

- collectible or reward;
- usable or equippable item;
- interactive world prop;
- destructible or hazard;
- decoration;
- objective or key object;
- UI-only representation.

Make interaction state readable through silhouette, value, accent, animation, or
placement. Do not rely on glow for every interactable.

For an item family, lock:

- canvas and object occupancy;
- camera angle;
- outline and pixel density;
- palette roles;
- material rendering;
- lighting direction;
- rarity or state overlays;
- background and shadow policy.

Generate one family member first and use it as the anchor.

## Keep world sprites and icons separate

World sprites optimize for scene readability and ground contact. Inventory icons
optimize for recognition inside a fixed UI frame. Derive them from the same
design, but do not assume one raster serves both jobs.

Avoid baking rarity frames, quantities, cooldowns, selection rings, or status
marks into base item art when the UI can compose them.

## Build effects around timing

Define:

- gameplay event and exact contact frame;
- duration and loop behavior;
- area, direction, and origin;
- readability against common backgrounds;
- additive, alpha, multiply, or normal blend expectations;
- whether color carries gameplay meaning;
- maximum screen coverage and flash intensity;
- reduced-motion or reduced-flash alternative.

Keep hit logic in code. Art communicates the event but does not determine the
collision or timing boundary.

For reusable effects, separate core shape, particles, trail, glow, and impact
when runtime composition improves control.

## Prompt pattern for items and props

```text
Runtime asset: <world sprite/inventory icon/interactive prop>
Function: <gameplay role>
View and logical size: <camera and dimensions>
Family anchor: <reference role>
Silhouette: <recognition at target size>
Materials and palette: <shared rules plus object delta>
Occupancy and padding: <canvas usage>
Background: removable flat chroma
Avoid: text, UI frame, scenery, cast shadow, extra objects, cropped silhouette
```

## Prompt pattern for VFX

```text
Runtime asset: ordered <effect> sprite strip
Gameplay event: <telegraph/contact/recovery>
Origin and direction: <pivot and travel>
Frame count and layout: <exact ordered slots>
Shape and color language: <meaning>
Blend expectation: <runtime blend mode>
Background: removable flat chroma
Avoid: scenery, characters, labels, arbitrary camera motion, opaque rectangular
background, effect crossing slot boundaries
```

## UI art rules

- Match the game's palette and material language without sacrificing text and
  control clarity.
- Prefer code-native layout, text, borders, and simple state changes.
- Use raster generation for portraits, decorative panels, complex emblems, or
  illustrated content, not for entire interactive screens.
- Keep normal, hover, focus, selected, pressed, disabled, warning, and success
  states explicit.
- Preserve scalable nine-slice regions when the engine supports them.
- Test at target DPI and with localization, large text, keyboard focus, and
  controller navigation when applicable.
