# Phaser 2D Asset Pipeline

Use this reference when producing, integrating, or auditing art for a Phaser
game. Inspect `package.json`, the lockfile, and the installed Phaser types or
source before writing integration code. Match documentation to the installed
major and minor version.

## Contents

- Choose the Phaser texture form
- Follow the project loading boundary
- Integrate sprites and animation
- Preserve pivots and bounds
- Configure pixel art
- Package atlases safely
- Build tilemaps and levels
- Validate in Phaser
- Official sources

## Choose the Phaser texture form

Use the smallest representation that preserves the asset contract:

- **Image**: one complete texture with one base frame. Load with `load.image`.
- **Spritesheet**: uniform cells in a regular grid. Load with
  `load.spritesheet` and explicit `frameWidth`, `frameHeight`, `margin`, and
  `spacing` when applicable.
- **Atlas**: named frames with variable bounds or packed whitespace. Load the
  image and atlas data together with `load.atlas`.
- **Tilemap**: structured map data plus one or more separately loaded tileset
  images. Load the map format supported by the installed version.

Prefer a spritesheet for fixed-size pixel animation. Prefer an atlas when frame
sizes differ, stable semantic names matter, or many asset families share a
texture. Do not convert a simple four-frame sheet into an atlas without a
project or runtime reason.

## Follow the project loading boundary

Discover how the project resolves URLs before copying assets:

- static public paths;
- bundler-imported URLs;
- asset manifests or generated catalogs;
- scene-owned `preload` methods;
- a shared boot or preload scene;
- lazy loading by feature.

Preserve that boundary. Do not invent a second loader or duplicate the same
texture under multiple keys.

Use stable, unique texture keys. Phaser stores loaded textures in a global
Texture Manager, so unrelated scenes can reference the same texture and frame
definitions. Inspect `this.textures` after loading when frame parsing is in
doubt.

## Integrate sprites and animation

For a regular sheet:

1. verify total dimensions against cell size, margin, and spacing;
2. load it with the exact frame configuration;
3. inspect generated numeric frame indexes;
4. create the animation once through the project's Animation Manager boundary;
5. use `generateFrameNumbers` for numeric spritesheet frames;
6. set loop, frame rate, repeat, and per-frame timing from the animation
   contract;
7. keep hit, footstep, spawn, and contact events in gameplay code.

For an atlas:

1. preserve stable semantic frame names;
2. use `generateFrameNames` or an explicit ordered frame list;
3. retain source size and trim metadata;
4. preserve custom pivots only when the project consumes them deliberately;
5. disable rotation during packing when it complicates pixel-art review or
   tooling.

Do not infer animation order from filesystem or object-key iteration. Store it
explicitly.

## Preserve pivots and bounds

Phaser Game Objects use an origin, commonly normalized from `0` to `1`. Map the
art contract's bottom-center anchor to the project's origin convention and
verify it on every frame.

Keep these concepts separate:

- source frame rectangle;
- trimmed visible rectangle;
- original untrimmed source size;
- visual origin or custom pivot;
- physics body or hitbox;
- gameplay contact point.

Do not generate physics bodies from alpha bounds when gameplay needs stable
collision. Do not let a large VFX frame shift the character origin.

## Configure pixel art

For an all-pixel-art game, evaluate Phaser's `pixelArt` game configuration. In
supported versions it selects nearest-neighbor texture filtering and disables
antialiasing while enabling pixel rounding. If the project mixes pixel and
smooth art, set the required filter on individual textures instead of changing
the whole renderer casually.

Also verify:

- integer display scaling where the design requires it;
- integer camera or object positions at the final transform boundary;
- no CSS scaling that introduces blur;
- no fractional atlas frames;
- no unintentional smoothing after post-processing or texture replacement;
- Canvas and WebGL output when both renderers are supported.

Do not assume nearest filtering fixes inconsistent logical pixel sizes inside
the source image.

## Package atlases safely

When packing:

- keep frame names stable across repacks;
- add enough padding or extrusion to prevent neighboring-frame bleeding;
- retain transparent source bounds when the pivot depends on them;
- avoid oversized empty bounds and unnecessary texture pages;
- check the browser/device maximum texture size through the renderer when atlas
  size is close to the project budget;
- verify the actual decoded texture footprint, not only PNG or WebP size.

Repacking must not silently renumber frames consumed by content or animation
data.

## Build tilemaps and levels

Keep map data, tileset images, and gameplay semantics explicit. Phaser can parse
formats such as Tiled JSON, CSV, or arrays, but exact support and limitations
depend on the installed version.

For Tiled workflows:

- keep tileset names consistent between Tiled data and Phaser integration;
- verify tile width, height, margin, spacing, and first global ID;
- verify whether the installed Phaser parser supports external tileset files
  before choosing `.tsx` references;
- preserve tile and object-layer custom properties intentionally;
- keep collision, navigation, spawn, encounter, and trigger semantics in map
  data or domain content, not decorative pixels;
- preserve deterministic layout and content IDs independently from atlas
  coordinates;
- test transformed, flipped, and rotated tiles if the map uses them.

Use separate layers for background, terrain, gameplay metadata, props, and
foreground when their lifecycle or ordering differs.

## Validate in Phaser

Before calling the asset integrated:

- confirm texture keys and frame names/indexes exist;
- inspect `__MISSING` texture occurrences and loader warnings;
- compare parsed frame bounds with the asset manifest;
- verify origin, scale, flip, tint, depth, scroll factor, and blend mode;
- play every animation at actual timing and confirm event frames;
- move the camera across tile boundaries to find seams;
- test common viewport sizes and device pixel ratios;
- confirm disposal or scene transitions do not remove shared textures early;
- capture a runtime screenshot at the actual logical resolution.

## Official sources

- Loader: https://docs.phaser.io/phaser/concepts/loader
- Textures and frames: https://docs.phaser.io/phaser/concepts/textures
- Animations: https://docs.phaser.io/phaser/concepts/animations
- Phaser API: https://docs.phaser.io/api-documentation/

Use the version selector or installed source when the project version differs
from the documentation page.
