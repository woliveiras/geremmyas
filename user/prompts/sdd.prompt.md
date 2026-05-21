---
description: "Run the full Spec Driven Development cycle for a feature. Orchestrates spec → test → implement → review → docs."
---

Guide me through the Spec Driven Development cycle for this feature.
Follow each step in order. **Do not advance to the next step without my explicit approval.**

## Step 1 — Spec

Use @spec-writer to interview me and produce a spec in `specs/`.
If the feature is large, break it into vertical slices first.

**Gate:** I must approve the spec before proceeding.

## Step 2 — Tests

Use the `generate-tests-from-spec` skill to generate tests from the approved spec's acceptance criteria.
Each criterion maps to at least one test. Tests must fail initially (red).

**Gate:** Tests must be created and confirmed failing before proceeding.

## Step 3 — Implement

Write production code to make the tests pass.
Never modify the tests — they are the source of truth.

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

---

Start with Step 1 now.
