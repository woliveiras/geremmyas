# Shared gameplay foundations

## Model the behavior

Record:

- authoritative state and valid transitions;
- inputs or stimuli and their consumption rules;
- time domain and units;
- collision, navigation, and targeting assumptions;
- animation and audio cues emitted by the simulation;
- cancellation, interruption, death, pause, and reload behavior.

Prefer an explicit transition table when three or more states interact.

## Separate responsibilities

- **Input:** convert keyboard, pointer, touch, or controller signals into intentions.
- **Simulation:** update rules and authoritative state.
- **Physics:** resolve movement, overlaps, and queries.
- **Presentation:** display animation, camera, particles, UI, and audio.
- **Persistence:** serialize durable data only.

Presentation may observe gameplay events but must not become the sole source of damage, inventory, or progression truth.

## Treat time deliberately

- Use one documented unit internally.
- Use a fixed physics step for behavior that depends on stable collision or deterministic stepping.
- Use frame delta for presentation that should follow rendered time.
- Define whether timers follow gameplay time, real time, or UI time.
- Clamp or subdivide pathological deltas after tab suspension or debugger pauses.
- Test at multiple simulated frame rates.

## State machines

Give each state:

- entry and exit effects;
- accepted inputs;
- transition guards and priorities;
- update ownership;
- interruption rules;
- cleanup requirements.

Avoid boolean soups such as `is_attacking`, `is_hurt`, `is_dead`, and `is_dashing` when the combinations are mutually exclusive.

## Movement and collision

- Distinguish desired velocity from resolved motion.
- Keep collider dimensions stable across cosmetic frames.
- Define one-way platforms, slopes, moving platforms, ladders, and teleportation explicitly.
- Use layers and masks by semantic role.
- Decide whether triggers are edge-triggered or continuous.
- Revalidate collision after scaling, pivot, or tile-size changes.

## Combat

Separate:

- attack intent and cooldown;
- active hitbox windows;
- target filtering;
- damage calculation;
- invulnerability and repeated-hit policy;
- knockback and hit stun;
- visual and audio feedback.

Assign stable attack or event IDs so repeated overlap callbacks cannot apply the same hit twice unless intended.

## Data and identity

Use stable semantic IDs for items, abilities, enemies, maps, and checkpoints. Keep authored definitions immutable at runtime; store mutable instance state separately. Do not persist engine object references.

## Completion contract

Deliver:

- changed behavior and ownership boundary;
- invariants and edge cases;
- automated verification;
- manual runtime checks still required;
- compatibility or migration notes.
