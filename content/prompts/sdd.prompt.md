---
description: "Run the full Spec Driven Development cycle for a feature. Orchestrates spec → test → implement → review → docs."
---

Guide me through the Spec Driven Development cycle for this feature.
Follow each step in order. Advance automatically when each evidence gate passes.

## Step 0 — Commit permission

Ask whether you may create git commits for this work or I handle commits myself.
Remember the answer for the rest of the cycle.

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

After acceptance criteria, tasks, docs, fresh evidence, and independent review
are reconciled, mark the spec `Verified`. Record failed evidence by keeping or
returning it to `In Progress`; return it to `Draft` for a newly unresolved
material decision. Commit, push, merge, and release evidence remain separate.

## Step 6 — Commit (optional)

Use `git-commit` only if I granted commit permission in Step 0. Otherwise summarize
changed files and leave committing to me.

---

Start with Step 0 now.
