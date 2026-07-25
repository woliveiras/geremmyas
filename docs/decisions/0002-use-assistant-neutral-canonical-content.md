---
status: implemented
date: 2026-07-25
---

# Use assistant-neutral canonical content with target adapters

## Context and Problem Statement

Geremmyas supports several coding assistants, but 88 of its 91 catalog entries
use Copilot's `.github/*` layout as both source taxonomy and default consumer
destination. How should one canonical catalog remain reusable without treating
one assistant's filesystem as the domain model?

## Decision Drivers

* Shared skills, instructions, contracts, and templates need one source of truth.
* A selected target must receive only its native artifacts.
* Existing Copilot consumers need compatible output paths.
* Target removal must not delete user-modified or unowned files.

## Considered Options

* Keep `project/.github` canonical and continue generating other targets from it.
* Rename the source tree but continue inferring artifact meaning from consumer
  destination paths.
* Use neutral typed content with explicit target adapters.

## Decision Outcome

Chosen option: "Use neutral typed content with explicit target adapters",
because it preserves a single source of truth while separating artifact meaning
from assistant-native paths. Shared sources move to `content/`, assistant-only
sources live under `targets/<target>/`, and catalog entries declare their kind.

### Consequences

* Good, because Codex-only installation no longer requires Copilot directories.
* Good, because target routing can be tested from artifact type and selected
  targets instead of path prefixes.
* Good, because Copilot can retain its current `.github/*` destinations.
* Bad, because the catalog, embedded paths, tests, documentation, and
  maintainer dogfooding links must migrate together.
* Bad, because project target removal needs ownership metadata and conservative
  legacy adoption.

### Confirmation

Catalog validation rejects missing or unknown artifact kinds. Integration tests
cover clean Copilot-only, Codex-only, and mixed installs. Reconciliation tests
prove that only unchanged owned files are removed and that modified, unowned,
or symlinked files are preserved.

## Pros and Cons of the Options

### Keep `project/.github` canonical

* Good, because it requires no source migration.
* Bad, because Copilot remains the implicit baseline and target selection cannot
  prevent Copilot artifacts from being copied first.

### Rename sources but keep destination-based inference

* Good, because contributor-facing paths look neutral.
* Bad, because the catalog and runtime still use `.github/*` to decide artifact
  meaning, so the architectural coupling remains.

### Use neutral typed content with target adapters

* Good, because source taxonomy, artifact meaning, and assistant destinations
  become independent.
* Bad, because it is the largest migration and requires compatibility tests for
  every supported target.

## More Information

* [Spec 0007](../../specs/0007-assistant-neutral-content/spec.md)
* [Multi-assistant PRD](../prds/2026-06-22-multi-assistant-framework.md)
