---
status: Implemented
date: 2026-06-30
---

# Tier the skill catalog into core and stack scopes

## Context and Problem Statement

The catalog ships ~40 skills as a single flat tier, all installed both globally
and per-project by the same packs. In a real SDLC most skills are dormant for
any given repo, yet every skill description is injected into the assistant
context on every turn, and the global/project duplication is inherent to the
current model. How should the catalog be organized so that a repo only carries
the skills it actually uses, without losing the coherent SDD workflow?

## Decision Drivers

* Per-turn context budget: the full skill listing is injected every turn, so
  catalog size has a fixed token cost in every repo.
* Signal-to-noise for skill selection: too many near-overlapping skills dilute
  model attention and create redundant nudges.
* Separation of concerns: general workflow skills are useful everywhere; stack
  skills only matter in repos using that stack.
* Maintenance: version-pinned framework recipes and niche skills rot and cost
  curation effort.

## Considered Options

* **Option 1 — Status quo.** Keep a flat catalog; every pack installs globally
  and per-project.
* **Option 2 — Tier the catalog** into `core` (general SDLC workflow, installed
  everywhere) and `stack` (opt-in per project by technology), consolidate
  redundant skills, remove caveman in favor of its upstream installer, and fold
  `paper-review` into the opt-in `research` pack.
* **Option 3 — Split into multiple catalogs/repos** (one per domain) consumed
  independently.

## Decision Outcome

Chosen option: "Option 2 — Tier the catalog", because it directly addresses the
per-turn context cost and signal dilution drivers while preserving the coherent
SDD + guardrails core that is the catalog's main asset. It avoids the
distribution and discovery overhead of Option 3 and removes the duplication and
dormancy problems that Option 1 leaves in place.

### Consequences

* Good, because a typical repo carries ~18-22 skills instead of ~52, shrinking
  the injected listing and sharpening selection.
* Good, because global vs project scope becomes intentional (`core` everywhere,
  stack helpers per-project), removing accidental duplication.
* Bad, because it is a breaking change to pack layout and names; consumers must
  re-run install/sync and may need to opt back into stack packs.
* Bad, because tier metadata adds a small schema and validation surface to
  `catalog/packs.json`.
* Neutral, because caveman leaves the catalog entirely (recommended via its
  upstream installer) and `paper-review` joins the `research` stack pack. A
  `personal` tier was considered for global-only tooling but dropped for lack
  of distributed members.

A lazy-injection variant (an `auto_inject` flag plus a generated `skill-router`
so non-core skill descriptions are not injected every turn) was considered and
rejected. The host scans the whole skills root and injects every installed
skill; geremmyas cannot suppress that. Hiding a skill would mean storing it
outside any scanned skills root, which strips its native skill behavior
(auto-trigger). The robust lever geremmyas owns is which skills get installed,
which tiering already provides.

### Confirmation

Catalog tests (`internal/cli/catalog_test.go`), `geremmyas lint`, and
`geremmyas doctor` enforce the new tier field and install behavior. The
companion spec `specs/0004-skill-catalog-tiering/` defines the testable
acceptance criteria.

## Pros and Cons of the Options

### Option 1 — Status quo

* Good, because no migration and no consumer breakage.
* Bad, because every repo pays the full catalog context cost and carries dormant
  skills.
* Bad, because global/project duplication and redundant skill clusters remain.

### Option 2 — Tier the catalog

* Good, because scope is explicit and the injected set is minimal per repo.
* Good, because the SDD workflow core stays intact and coherent.
* Bad, because it is a breaking re-tiering that consumers must adopt.

### Option 3 — Split into multiple catalogs/repos

* Good, because maximal separation and independent versioning.
* Bad, because it fragments distribution, dependency resolution, and discovery
  for a single-binary tool.
* Bad, because it is the largest change for the same context-budget benefit as
  Option 2.

## More Information

Companion spec: [specs/0004-skill-catalog-tiering/spec.md](../../specs/0004-skill-catalog-tiering/spec.md).
Related: [docs/creating-packs.md](../creating-packs.md), `catalog/packs.json`.
