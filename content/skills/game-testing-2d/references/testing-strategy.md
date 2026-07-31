# Testing strategy

## Evidence ladder

Use the lowest layer that proves the contract:

1. pure functions and data validation;
2. deterministic simulation over explicit ticks;
3. engine scene or physics integration;
4. rendered runtime or browser automation;
5. exported artifact smoke test;
6. human play, visual, audio, accessibility, and feel review.

Higher layers complement lower ones; they do not erase the need for stable focused tests.

## Deterministic harness

Control:

- seed and PRNG algorithm;
- starting state and entity order;
- input actions by tick;
- step duration and number of steps;
- clock, timers, and pause state;
- asynchronous asset or scene completion;
- locale, viewport, and device profile where relevant.

Record a compact trace containing tick, inputs, state transitions, collisions, emitted events, and a final canonical hash.

## Useful contracts

- Movement reaches expected positions within tolerances.
- A state transition happens once under its guard.
- Attack windows hit valid targets once and reject invalid targets.
- Collision and navigation layers agree with semantic roles.
- Animation names, frame counts, and event markers exist.
- A seed generates connected, valid content.
- Save data round-trips and migrates.
- Scene restart does not duplicate listeners or retained objects.

## Property and fuzz testing

Generate bounded inputs and assert invariants such as finite positions, nonnegative health, reachable exits, valid IDs, and no entity occupying forbidden cells. Save the smallest failing seed.

## Visual checks

Use screenshots for stable rendering contracts, not for rules already testable as data. Fix logical resolution, DPR, fonts, renderer, seed, camera, and animation frame. Review intentional changes rather than blindly updating baselines.

## Flake triage

Log seed, test order, elapsed ticks, environment, engine version, and artifact hash. Remove real-time waits, implicit ordering, shared mutable globals, and unawaited resource loading before increasing retries.

## Manual acceptance

Keep checklists for:

- control feel and input devices;
- readability and animation timing;
- audio balance and autoplay;
- camera motion and photosensitivity settings;
- target hardware performance;
- browser or export-specific behavior.
