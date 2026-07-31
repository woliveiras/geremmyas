# AI architecture

## Choose a decision model

- Use a finite-state machine for a small set of mutually exclusive modes.
- Use hierarchical states when modes share substates or interruption rules.
- Use a behavior tree for reusable ordered tasks and fallback logic.
- Use utility scoring for competing continuous priorities.
- Combine models only at explicit boundaries.

## Perception

Model observations such as:

- distance and direction;
- line of sight and occlusion;
- hearing or recent disturbance;
- damage source;
- navigation reachability;
- ally signals;
- last-known target position.

Specify sampling cadence, filters, forgetting, and uncertainty. Make debug overlays show what the AI actually perceives.

## Decision and memory

Keep mutable blackboard data scoped to one agent unless deliberately shared. Record the selected action and reason. Add hysteresis or minimum commitment time to prevent rapid oscillation.

## Navigation and steering

Separate high-level target selection, path calculation, local steering, and final physics motion. Handle unreachable goals, stale paths, moving targets, doors, dynamic obstacles, and recovery from being stuck.

## Combat fairness

Define telegraph, reaction window, range, cooldown, active window, recovery, target lock, interruption, and repeated-hit behavior. Make difficulty adjust readable parameters rather than secretly breaking rules.

## Groups and encounters

Budget simultaneous attackers, spawn visibility, protected cells, retreat space, and reinforcement cadence. Avoid all agents making identical choices on the same tick.

## Determinism and tests

Inject time and randomness. Test decisions from synthetic observations, then integrate navigation and physics. Capture seed, tick, perceptions, chosen action, path status, and transition reason.

## Debugging

Display current state, target, last-known position, path, perception shapes, cooldowns, and failure reason. Ensure debug drawing can be disabled without changing AI results.
