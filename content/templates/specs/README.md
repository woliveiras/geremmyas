# Specs

Specs are executable contracts: they describe **what** will be built, **why**,
and the **acceptance criteria** before any production code. Proposals and audits
may live under `docs/proposals/`; specs are the implementation contracts.

Maintain this file by hand as the **index** of every spec in the repository.
Agents must update it whenever a spec changes lifecycle status.

After the first `geremmyas sync`, this file is **preserved** on later syncs
unless you pass `--force` (same policy as `AGENTS.md`).

## Conventions

### Feature folder structure

Each spec lives in a numbered folder:

```text
specs/
├── README.md                 # this index (families, status, numbering)
└── NNNN-<slug>/
    ├── spec.md               # WHAT + WHY + acceptance criteria
    ├── plan.md               # HOW (approach, decisions, tradeoffs)
    └── tasks.md              # executable checklist (vertical slices)
```

`plan.md` and `tasks.md` are always created with `spec.md` for new features.

### Progress (`tasks.md`)

Checkbox state is the only progress signal for resuming work:

- `- [ ]` pending
- `- [~]` in progress
- `- [x]` done

Agents must update checkboxes while implementing. Do not create separate
handoff files. After reading this index for status, open the feature folder
and continue from the in-progress or next pending task.

### Numbering

- Numbers are **global**, sequential, and **immutable** (`0001`, `0002`, …).
- Assign the next free number across all of `specs/` (see below).
- Deprecated specs keep their number; numbers are never recycled.
- Folder name: `NNNN-<slug>/` (four digits, lowercase kebab-case slug).

### How to allocate a number for a new spec

1. Scan `specs/` for the highest existing `NNNN` prefix.
2. **New product family** (independent goal): reserve the next round block
   (e.g. highest is `0041` → start the family at `0050` or `0100`).
3. **Extend an existing family**: use the next free number inside that family's
   reserved block.
4. **Block full**: take the next global number and set `family:` in `spec.md`
   frontmatter so the index still groups it correctly.

### Reserved blocks by family

Each family typically reserves 10 numbers per phase (use 50 or 100 for large
programs). Edit this table when you add a family or reserve a new block.

| Family | Reserved block |
| --- | --- |
| _Example — Foundation_ | 0001–0009 |
| _Example — APIs_ | 0010–0019 |
| _Reserved for next families_ | 0100+ |

## Families

A **family** is a coordinated set of specs that share one product goal (for
example: onboarding, billing, platform tooling).

### When to create a new family vs extend an existing one

**Create a new family when:**

- The product goal is independent of existing families.
- You expect more than ~3 related specs.
- It has a distinct owner or deadline.

**Extend an existing family when:**

- The spec fits an existing phase naturally.
- It is a small addition (1–2 specs) to the same goal.

### Optional `family` header in `spec.md`

Use YAML frontmatter so the index can group specs even if the number sits outside
the family's reserved block:

```yaml
---
spec: "0042"
title: Short human title
family: example-foundation
phase: 1
status: Draft
owner: ""
depends_on: []
origin: ""
---
```

## Status lifecycle

Update `status` in frontmatter and the index tables together.

| Status | Meaning |
| --- | --- |
| **Draft** | Artifacts are being written or revised |
| **Ready** | Automatic readiness checks pass; implementation may start |
| **In Progress** | At least one implementation task is active |
| **Verified** | Acceptance criteria and fresh verification pass |

Historical specs may retain `In Review`, `Approved`, `Implemented`, `Completed`,
or `Deprecated`. These legacy statuses remain valid audit history; do not bulk
migrate them. Use the four-state lifecycle above for new specs and specs that
are deliberately migrated during active work.

### Ready

A spec is **Ready** when:

1. Acceptance criteria are testable (not vague).
2. Contracts (API, schema, UI, CLI) are concrete enough to test.
3. Dependencies are identified.
4. Tasks are vertical, sequenced, and have verification commands.
5. Material assumptions and error paths are documented.
6. No explicit read-only, plan-only, or no-edits session override applies.

Readiness is automatic. Implementation and feature tests start after these
checks pass, without a conversational approval pause.

### In Progress

A spec is **In Progress** when the first implementation task is marked `[~]`.
If an in-scope assumption changes, update and revalidate the artifacts, then
continue while the objective remains unchanged.

### Verified

A spec is **Verified** when:

1. All acceptance criteria pass with fresh automated evidence where available.
2. Every finished task is `[x]` and no stale `[~]` remains.
3. Residual manual checks are documented without blocking local completion.
4. Operational docs and planned work are reconciled when applicable.
5. Independent review has no actionable finding.
6. `spec.md` records verification evidence and links available commits.

Failed evidence keeps or returns the spec to **In Progress**. An unresolved material decision that appears later returns it to **Draft**.

Escalate before implementation only when blast radius is high and either
rollback is unproven or critical harness evidence is missing. Explicit session
overrides and production authority boundaries always remain in force.

---

## Index

Specs are grouped by **family**. Each family has its own product goal and
reserved number block (see table above). Add a subsection per family; split
tables by **phase** when helpful.

### Family: _&lt;family-name&gt;_

- **Blocks:** 0001–0009 _(edit)_
- **Origin:** `docs/prds/...` or `docs/proposals/...` _(optional link)_
- **Goal:** One sentence describing the product outcome.

#### Phase 0 — _phase title_

| Spec | Title | Status | Depends on / Origin |
| --- | --- | --- | --- |
| [0001](0001-example-slug/spec.md) | Example spec title | Draft | — |

#### Phase 1 — _phase title_

| Spec | Title | Status | Depends on / Origin |
| --- | --- | --- | --- |
| | | | |

---

<!-- Copy the "Family:" subsection above for each family. Keep rows sorted by spec number. -->
