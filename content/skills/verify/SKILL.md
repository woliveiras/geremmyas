---
name: verify
description: "Collect fresh execution evidence before completion or handoff. Use when a changed slice is ready for a completion claim. Do not use for planning, discovery, or read-only review."
---

# Verify

Do not claim completion from confidence, inspection, compilation alone, or
stale output. Collect fresh evidence at the boundary affected by the change.

## Procedure

1. Derive the required checks from the request, acceptance criteria, changed
   files, risk, and repository commands.
2. Run the focused test or reproduction for the changed behavior.
3. Run the nearest relevant suite and static checks when practical.
4. Exercise error paths, regression coverage, and observable runtime behavior
   required by the acceptance criteria. A build alone does not prove behavior.
5. Record command, exit code, concise result, and evidence path. Keep failures
   separate from passing checks and distinguish preexisting failures.
6. Remove temporary logs, harnesses, and instrumentation.
7. Reconcile `tasks.md`, `plan.md`, required docs, and lifecycle state.
8. After evidence passes, load the
   [independent review contract](./references/review-contract.md) only when the
   change's risk or governing contract requires review.
9. If a required check cannot run, record its owner, reason, residual risk, and
   whether it blocks local completion.

## Evidence record

```text
scope: behavior or acceptance criterion
command: exact command
result: pass | fail | unavailable (exit code and concise output)
freshness: run in this verification wave
evidence: test, log, screenshot, or artifact path
residual: unavailable or subjective checks and owner
```

Use [rationalization checks](./references/rationalization.md) only when someone
proposes skipping or weakening evidence. Never treat review as a substitute for
execution evidence, or execution evidence as a substitute for required review.
