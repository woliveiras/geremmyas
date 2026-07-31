# Game-feel patterns

## Input response

- Distinguish pressed, held, released, buffered, and consumed input.
- Add coyote time only around a clearly defined grounded transition.
- Buffer actions for a bounded window and consume them once.
- Define input priority when dash, attack, jump, and interaction compete.
- Keep input prompts device-aware.

## Movement

Tune with named values:

- maximum speed;
- acceleration and deceleration;
- turn acceleration;
- gravity, jump impulse, rise and fall multipliers;
- apex assistance;
- air control;
- terminal velocity;
- friction and slope behavior.

Measure rather than guess: time to full speed, time to stop, jump height, time to apex, total airtime, and reversal time.

## Impact

Layer only the signals the mechanic needs:

- anticipation and active pose;
- hit stop or short time dilation;
- knockback and hit stun;
- sprite flash or palette change;
- particles and decals;
- camera impulse;
- audio and optional rumble.

Keep damage and target selection authoritative outside this layer.

## Camera

Define:

- dead zone and look-ahead;
- follow stiffness and damping;
- room or world bounds;
- zoom rules;
- shake amplitude, frequency, falloff, and stacking;
- behavior during cutscenes, pause, death, and teleport.

For pixel art, verify integer logical transforms and rendering at supported zoom levels.

## Tuning workflow

1. Capture a baseline clip and measurements.
2. Pick a single player-visible goal.
3. Change one parameter group.
4. Compare at normal and reduced speed.
5. Test edge cases and interruptions.
6. Save named presets when comparing alternatives.

## Accessibility

Offer independent controls for shake, flashes, motion blur, rumble, and repeated rapid effects. Preserve essential information through shape, timing, text, or audio alternatives.
