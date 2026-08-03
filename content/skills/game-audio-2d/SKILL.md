---
name: game-audio-2d
description: "Use when 2D game music, ambience, sound effects, Web Audio, buses, volume, crossfades, spatial sound, or voice limits change. Do not use for visual feedback, accessibility-only cues, art, or release packaging."
---

# Game Audio 2D

Treat audio as a routed gameplay system with explicit ownership, priorities, and platform behavior.

## Establish the audio contract

1. Inspect engine version, target platforms, current mixer or Sound Manager, asset formats, loading strategy, and settings persistence.
2. Classify sounds as music, ambience, gameplay SFX, voice, or UI.
3. Define trigger, ownership, concurrency, priority, loop, fade, spatial, pause, and cleanup behavior.
4. Set a loudness and headroom policy appropriate to the project.

## Load references

- Read [audio-system.md](references/audio-system.md) for categories, playback policies, mixing, variation, and validation.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-save-n-progress` for settings persistence and `$game-feel-2d` for impact timing.

## Integrate one audio family

1. Register stable semantic audio IDs.
2. Route through a category or bus.
3. Apply concurrency and priority rules.
4. Trigger from authoritative gameplay or UI events.
5. Handle pause, scene transitions, interruption, and object destruction.
6. Provide user volume and mute controls without destroying mix ratios.
7. Verify the actual target platform.

## Verify

- Test missing files, locked audio, rapid repeated triggers, scene reload, pause, focus loss, and device changes where applicable.
- Check loops for clicks and unintended gaps.
- Monitor clipping, masking, excessive simultaneity, and memory.
- Test with music, effects, and UI active together.
- Keep listening tests distinct from automated routing and lifecycle tests.

## Guardrails

- Do not play an unrestricted new instance for every overlap callback.
- Do not tie essential information to sound alone.
- Do not let global music survive scenes accidentally.
- Do not normalize every sound to maximum amplitude.
- Do not assume editor or desktop audio behavior matches web exports.
