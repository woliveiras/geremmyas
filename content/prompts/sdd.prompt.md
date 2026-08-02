---
description: "Run the full Spec Driven Development cycle for a feature. Orchestrates spec → test → implement → review → docs."
---

Guide me through the Spec Driven Development cycle for this feature.
Follow each step in order. Advance automatically when each evidence gate passes.

## Step 0 — Authority overrides

Record any explicit read-only, plan-only, no-edits, or no-commits override. Do
not ask about commit permission; atomic local commits are the default. This does
not authorize push, production changes, or publication.

## Step 1 — Spec

Use @spec-writer to interview me and produce a numbered folder under `specs/`
(`NNNN-<slug>/` with `spec.md`, `plan.md`, `tasks.md`) and update `specs/README.md`
index (family, status `Draft`).
If the feature needs product framing, write or update a PRD first, then the spec.

**Gate:** Mark the package `Ready` when acceptance criteria are testable,
contracts and dependencies are known, verification commands exist, and no
material product decision remains unresolved. Otherwise update it or ask one
focused question for the unresolved decision.

## Step 2 — Tests

Use the `generate-tests-from-spec` skill on the **Ready** spec. Choose unit or
integration tests per the spec Test Strategy and each task's `test-type`.
Each criterion maps to at least one test. Tests must fail initially (red).

**Gate:** Tests must be created and confirmed failing before proceeding.

## Step 3 — Implement

Use `vertical-tdd` to implement one behavior at a time. Update `tasks.md` checkboxes.
Never modify the tests — they are the source of truth.
Mark the spec `In Progress` with the first `[~]` task. Revalidate changed
in-scope assumptions and continue while the objective is unchanged.

**Gate:** All tests pass (green) before proceeding.

## Step 4 — Review

Use @reviewer to verify alignment between spec, tests, and code.
Flag any acceptance criteria without tests, any tests without matching code, or architecture issues.

**Gate:** Review issues resolved before proceeding.

## Step 5 — Docs (if needed)

Use the `update-docs` skill only if:
- Public API changed (new endpoints, parameters, responses)
- Architecture changed (new modules, boundaries, patterns)
- Setup changed (new dependencies, build steps, env vars)

Skip this step if the change is internal-only.

After each task-owned slice passes focused tests, docs reconciliation, and
independent review, continue through Step 6 before starting the next slice. The
feature remains `In Progress` while more tasks remain.

## Step 6 — Local commit per slice

After each task-owned slice passes its evidence gates, have the primary agent use
`git-commit` to create an atomic Conventional Commit containing implementation,
tests, required docs, and durable artifacts. Do not wait for the whole feature
to reach `Verified`. Do not ask for file, hunk, message, or commit confirmation.
For a Step 0 override, skip the commit and report changed or proposed files.
Never push. Return to Step 2 for the next slice.

## Final lifecycle reconciliation

After every acceptance criterion, task, doc, and independent review is
reconciled, mark the spec `Verified`. Failed evidence keeps or returns it to `In
Progress`; a newly unresolved material decision returns it to `Draft`. Commit,
push, merge, and release evidence remain separate from lifecycle state.

---

Start with Step 0, then continue automatically.
