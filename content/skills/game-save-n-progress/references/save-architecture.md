# Save architecture

## Separate data lifetimes

- **Settings:** audio, display, accessibility, language, input mappings.
- **Profile:** unlocks, achievements, currencies, long-term progression.
- **Run:** current procedural seed, inventory, health, encounter state.
- **Checkpoint:** minimal restore point plus world deltas.
- **Cache:** reproducible or disposable data that does not belong in a save.

Save each lifetime independently when their write cadence and recovery needs differ.

## Schema envelope

Include:

- schema version;
- game or content version when relevant;
- stable profile or slot ID;
- written timestamp for diagnostics, not ordering truth by itself;
- payload;
- optional checksum for corruption detection;
- generator and catalog versions for procedural content.

## Stable data

Persist semantic IDs, scalar state, bounded collections, and world deltas. Avoid coordinates tied to repackable atlases, localized strings, display names, node paths, and engine resources unless they are explicit stable contracts.

## Migrations

Use sequential migrations:

```text
read -> parse -> validate envelope -> migrate vN to vN+1 -> validate current -> apply
```

Keep migration fixtures for every supported historical version. Reject future versions without overwriting them.

## Write safety

Prefer:

1. serialize and validate in memory;
2. write a temporary candidate;
3. flush or close;
4. preserve a backup if policy requires;
5. replace the active save;
6. verify the new active file when feasible.

Browser storage may not offer true atomic rename; design the adapter with generations or active-pointer records if interruption safety matters.

## Autosave

Save at stable gameplay boundaries, debounce repeated triggers, serialize writes, and surface failure without blocking the main loop. Do not save half-applied transactions.

## Tests

Cover round trips, migrations, defaults, unknown IDs, duplicates, numeric bounds, malformed text, truncation, interrupted writes, full quota, unavailable storage, deletion, and repeated load/save cycles.

## Cloud boundary

Treat cloud sync as a separate replication system. Define identity, encryption, offline behavior, conflict detection, resolution, deletion, and provider failure before integrating it.
