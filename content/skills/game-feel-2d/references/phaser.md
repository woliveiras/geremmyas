# Phaser game-feel integration

Inspect the installed Phaser version and current Scene, input, camera, tween, time, animation, and physics setup.

## Patterns

- Store buffer and coyote windows in the gameplay time domain.
- Use Scene clocks and tweens for presentation whose lifetime belongs to that Scene.
- Keep physics velocity authoritative; apply squash, recoil visuals, and camera effects to presentation objects or containers.
- Use camera effects or bounded custom trauma, with explicit stacking and cancellation.
- Clean up input, animation, tween, and timer listeners on Scene shutdown.
- Treat global time scale and per-Scene time scale as consequential state; restore them after hit stop.

## Pixel-art checks

Keep camera position, zoom, sprite position, and display scaling compatible with the project’s pixel policy. Camera rounding cannot repair fractional CSS canvas scaling.

## Official references

- [Input](https://docs.phaser.io/phaser/concepts/input)
- [Cameras](https://docs.phaser.io/phaser/concepts/cameras)
- [Tweens](https://docs.phaser.io/phaser/concepts/tweens)
- [Time](https://docs.phaser.io/phaser/concepts/time)
