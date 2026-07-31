# Audio system

## Categories and routing

Use stable categories:

- master;
- music;
- ambience;
- gameplay SFX;
- voice;
- UI.

Allow independent user control while preserving authored relative levels inside each category.

## Playback policy

For every sound family define:

- trigger and semantic ID;
- source variants;
- maximum simultaneous voices;
- priority and steal policy;
- minimum retrigger interval;
- loop and fade behavior;
- positional range and attenuation;
- pause and focus behavior;
- owning scene or persistent service;
- preload or streaming policy.

## Variation

Use small bounded pitch, volume, start-offset, or variant changes where repetition is undesirable. Keep critical rhythm, voice, and musical pitch stable. Seed randomness when audio choice must replay deterministically.

## Music and ambience

Represent music state separately from scene object lifetime. Define crossfade duration, transition bar or marker if rhythm matters, resume versus restart, and interruption priority. Layer ambience by location without leaking loops after transitions.

## Gameplay synchronization

Trigger sounds from authoritative events such as accepted jump, confirmed hit, pickup, UI action, or state transition. Animation markers may emit cues, but missing presentation must not change gameplay.

## Mixing

Leave headroom for simultaneous effects. Test busy gameplay rather than isolated files. Use ducking only with a documented priority such as voice over music. Avoid hard clipping and repeated full-scale transients.

## Accessibility

Provide captions, visual telegraphs, directional indicators, or controller feedback for essential cues. Keep mute and per-category controls reachable before gameplay.

## Validation

Automate ID existence, category routing, concurrency limits, cleanup, settings round-trip, and missing-file handling. Use human listening for balance, fatigue, intelligibility, loop quality, and device differences.
