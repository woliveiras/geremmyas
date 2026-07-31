# Phaser UI and accessibility

Inspect Phaser version, canvas scale policy, DOM container use, scene layering, input plugins, fonts, and browser targets.

## Patterns

- Use a dedicated UI Scene when its lifecycle and camera differ from gameplay.
- Keep UI state in testable TypeScript models.
- Use Phaser Game Objects for integrated canvas visuals and native DOM elements when semantic HTML, text input, or browser accessibility materially helps.
- Remember DOM elements sit outside the canvas display list and have camera and nesting limitations.
- Translate pointer, keyboard, touch, and gamepad events into semantic UI actions.
- Handle browser focus, resizing, DPR, fullscreen, and orientation changes.
- Expose accessible HTML alternatives for essential web controls when canvas alone cannot meet the requirement.

## Testing

Use browser automation for tab order, DOM semantics, focus restoration, resize, touch targets, and keyboard flow. Use screenshots for stable layout, not as the sole accessibility evidence.

## Official references

- [Input](https://docs.phaser.io/phaser/concepts/input)
- [DOM Element](https://docs.phaser.io/phaser/concepts/gameobjects/dom-element)
- [Scale Manager](https://docs.phaser.io/phaser/concepts/scale-manager)
- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
