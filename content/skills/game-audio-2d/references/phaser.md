# Phaser audio integration

Inspect installed Phaser audio APIs, browser targets, codec assets, loading, Scene ownership, and current user-gesture handling.

## Patterns

- Treat the Sound Manager as global and give every loop an explicit owner.
- Unlock or resume audio after an accepted user gesture according to browser policy.
- Load compatible source alternatives based on target support.
- Route semantic categories through project-owned gain or volume policy.
- Stop or transfer scene-owned loops during shutdown.
- Limit repeated SFX and clean completed sound instances.
- Test tab focus, pause, mobile browsers, Bluetooth or device changes, and production hosting.

## Web constraints

Browser autoplay, background throttling, codec support, and Web Audio availability can change runtime behavior. Use installed Phaser docs and real target browsers as authority.

## Official references

- [Audio](https://docs.phaser.io/phaser/concepts/audio)
- [Loader](https://docs.phaser.io/phaser/concepts/loader)
- [Device](https://docs.phaser.io/phaser/concepts/device)
- [Scenes](https://docs.phaser.io/phaser/concepts/scenes)
