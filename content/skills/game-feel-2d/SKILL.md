---
name: game-feel-2d
description: "Use when tuning 2D controls, jumps, input buffering, hit stop, recoil, camera shake, particles, telegraphs, or moment-to-moment feedback. Do not use to change gameplay rules, UI accessibility, performance budgets, or tests."
---

# Game Feel 2D

Improve responsiveness and feedback without changing the intended gameplay rules accidentally.

## Establish a baseline

1. Read project instructions and identify engine version, physics model, input devices, camera, target frame rate, and accessibility settings.
2. Reproduce and measure the current mechanic.
3. Separate input latency, simulation response, animation, camera, effects, and audio.
4. Define the desired feel in measurable terms such as time to apex, stop time, buffer window, hit-stop duration, and camera settling.

## Load references

- Read [feel-patterns.md](references/feel-patterns.md) for controls, motion, impact, camera, and tuning.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$gameplay-programming-2d` when authoritative rules change and `$game-audio-2d` for mix and playback behavior.

## Tune one feedback loop

1. Put tunable values in a named configuration grouped by mechanic.
2. Change one variable family at a time.
3. Preserve gameplay truth during presentation-only effects.
4. Define stacking, cancellation, pause, and slow-motion behavior.
5. Provide reduced-motion, reduced-flash, and camera-shake controls when the effect can cause discomfort.
6. Compare before and after at real speed and slow motion.

## Verify

- Test multiple frame rates and input devices.
- Check that hit stop does not duplicate damage or break timers.
- Check that camera and pixel-art transforms remain stable.
- Confirm feedback still communicates without color, sound, or shake alone where accessibility requires redundancy.
- Require human playtesting for the final feel judgment.

## Guardrails

- Do not hide sluggish rules under particles or camera shake.
- Do not mutate collision bodies for squash and stretch.
- Do not use unbounded shake, flash, rumble, or time scale.
- Do not couple damage application to a cosmetic tween callback.
