# Accessible game UI

## Input and focus

- Support all input methods claimed by the game.
- Use visible focus indicators independent of hover.
- Define focus order and directional neighbors.
- Set initial focus and restore it after modal dismissal.
- Prevent gameplay input while a blocking menu owns input.
- Show device-aware prompts without changing action semantics.
- Allow remapping and resolve conflicts explicitly.

## Layout

Design against logical resolution, safe areas, aspect ratios, and text expansion. Prefer containers or constraints over fixed coordinates. Test maximum supported text size and long localized strings.

## Readability

- Use sufficient contrast and legible font sizes.
- Keep essential text as text, not pixels in an image.
- Avoid dense all-caps body copy.
- Distinguish focus, selected, disabled, danger, and success by more than hue.
- Keep critical HUD information readable over worst-case backgrounds.

## Feedback

Pair essential audio cues with visual or textual signals. Pair color with icon, shape, pattern, position, or text. Provide subtitles or captions when dialogue or audio conveys required information.

## Motion and flashes

Offer independent controls where appropriate for camera shake, flashes, rumble, motion effects, and repeated particles. Respect the setting in gameplay, UI, cutscenes, and tutorials.

## Settings

Make accessibility options reachable before demanding difficult play. Preview effects safely. Persist settings independently of a gameplay save and provide a reset path.

## Error and confirmation

Explain invalid actions, destructive choices, storage failures, and remapping conflicts. Keep confirmations actionable and return focus predictably.

## Verification matrix

Test:

- keyboard only;
- controller only;
- pointer and touch;
- reduced motion and flashes;
- mute with captions or visual cues;
- color-vision simulations;
- large text and localization expansion;
- narrow, wide, and safe-area viewports;
- pause, disconnect, focus loss, and device switching.
