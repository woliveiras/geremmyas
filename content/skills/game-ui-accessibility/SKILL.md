---
name: game-ui-accessibility
description: "Build and audit accessible game UI for Phaser and Godot. Use when creating HUDs, menus, focus navigation, touch controls, remapping, localization, text scaling, captions, or reduced motion. Do not use for visual-only interface reviews."
---

# Game UI Accessibility

Build interfaces that remain readable, operable, and understandable across supported displays and input methods.

## Establish the interface contract

1. Inspect engine version, target platforms, logical resolution, scale policy, input actions, localization, font pipeline, and existing accessibility settings.
2. List screens, HUD information, focus order, modal behavior, devices, and safe-area requirements.
3. Identify essential information and provide redundant cues.
4. Define supported text scale, languages, aspect ratios, and reduced-effect settings.

## Load references

- Read [accessible-ui.md](references/accessible-ui.md) for layout, navigation, text, feedback, settings, and verification.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-save-n-progress` to persist settings and `$game-art-2d` for UI asset production.

## Implement one complete flow

1. Use semantic actions rather than raw device keys.
2. Establish initial focus and deterministic navigation.
3. Make hover, focus, pressed, disabled, error, and selected states visually distinct.
4. Support cancellation and prevent input from leaking through modals.
5. Reflow or scale under supported viewport and text settings.
6. Announce essential state through more than color or sound alone.
7. Restore focus after closing overlays.

## Verify

- Complete the flow using keyboard, controller, pointer, and touch as supported.
- Test localization expansion, missing glyphs, text scale, narrow and wide screens, safe areas, and paused gameplay.
- Check contrast, flashing, shake, captions, and remapping.
- Use human accessibility review for claims automation cannot establish.

## Guardrails

- Do not reuse gameplay actions for UI focus navigation when the engine reserves UI actions.
- Do not indicate state with color alone.
- Do not trap focus or leave it on hidden controls.
- Do not bake essential text into raster assets.
- Do not make reduced motion remove essential timing information without an alternative.
