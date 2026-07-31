# Godot audio integration

Inspect the installed Godot version, audio bus layout, imported streams, player nodes, target platforms, and web playback settings.

## Patterns

- Route playback to named buses such as Music, SFX, Voice, and UI.
- Persist user settings as linear values while converting appropriately for bus dB controls.
- Choose `AudioStreamPlayer` or `AudioStreamPlayer2D` by spatial need.
- Give persistent music and scene-local ambience explicit ownership.
- Use polyphony, player pools, or concurrency managers only according to measured needs.
- Keep bus names stable because renamed buses can break routing.
- Verify effects and playback modes on web exports; support differs from desktop.

## Version control

Track the selected runtime audio files, import sidecars, and `default_bus_layout.tres`. Do not track generated `.godot` import cache.

## Official references

- [Audio buses](https://docs.godotengine.org/en/stable/tutorials/audio/audio_buses.html)
- [Audio streams](https://docs.godotengine.org/en/stable/tutorials/audio/audio_streams.html)
- [Web export audio](https://docs.godotengine.org/en/stable/tutorials/export/exporting_for_web.html)
