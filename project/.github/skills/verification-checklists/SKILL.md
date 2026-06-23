---
name: verification-checklists
description: "HARD GATE: Mandates fresh verification evidence before task completion. Use when: post-implementation, before marking work done. Do not use: before code is written, for planning, or for code review requests."
---


# Verification Checklists

<HARD-GATE>
Do NOT claim a task is complete without fresh verification evidence.
No "should work now" without running it.
No "pretty sure it's fixed" without output and review.
Confidence ≠ Evidence. Compile success ≠ Runtime proof.
</HARD-GATE>

## When to Use

- After implementing any behavior per vertical-tdd
- Before marking a task `[x]` complete
- Before moving to the next task
- To prevent phantom completion claims

## When NOT to Use

- During design phase (use approval-gates-before-implementation)
- For reading code only (verification requires execution)
- When blocked on external dependencies
- During planning or speculation

## Red Flags — STOP and Verify

| Excuse | Reality | Evidence Needed |
|--------|---------|-----------------|
| "Should work now" | Opinion | Run test → show output |
| "I'm confident it's fixed" | Confidence ≠ Evidence | Fresh test run with new output |
| "Code looks right" | Static review is not verification | Run behavior → observe result |
| "All tests pass" (didn't rerun) | Stale information | Re-run tests NOW, capture output |
| "Fixed it like I did before" | That context is gone | Fresh reproduction + fix + test |
| "No errors in the build" | Build ≠ behavior | Actually execute the fixed code |
| "Just one line changed" | Risky assumption | Verify the one-line impact |

## Verification Template (per task)

```
TASK: [name]

[1] RUN COMMAND
   $ go test -run TestX -v

[2] ACTUAL OUTPUT
   === RUN   TestX
   --- PASS: TestX (0.02s)
   PASS
   ok      example.com/pkg    0.025s

[3] VERIFICATION
   ✓ Test passes
   ✓ No regression in other tests (ran full suite)
   ✓ Output matches acceptance criteria
   ✓ Verified on clean checkout (no state from previous runs)

[4] DECISION
   VERIFIED: Task complete, moving to next.
   -or-
   NOT VERIFIED: [reason], investigating...
```

## Verification Evidence Types

**Mandatory** (not sufficient: code review alone)
- [ ] Test output (log, terminal screenshot, test result)
- [ ] Execution output (before-after for behavior changes)
- [ ] Fresh run (not cached, not from 30 mins ago)

**Supporting** (use with mandatory)
- [ ] Code review of changes
- [ ] Manual testing screenshot (if UI)
- [ ] Performance measurement (if perf task)
- [ ] Regression test run

**Not sufficient** (insufficient alone)
- [ ] "Code looks correct"
- [ ] "Compiled successfully"
- [ ] "Should work in theory"
- [ ] "Peer said it's fine"

## Anti-Patterns

**Phantom Completion**
- Implementing code → Declaring task done without running it → Bug escapes
- Applying fix → Assuming test passes → Test actually fails
- Refactoring → "Should be equivalent" without running → Breaks downstream

**Stale Verification**
- "I ran the test 30 mins ago" → Re-run NOW
- "The CI passed" → CI is not your verification; run locally or observe CI again
- "Last week I verified this pattern" → Different codebase, different context

**Weak Verification**
- Running only the happy path → Also verify edge cases from spec
- One test run → Run full suite to catch regressions
- "Manual testing showed it works" without logs → Capture reproducible output

**Self-Deception**
- Skipping verification because "I'm experienced" → Experience is not evidence
- "I know this works because I wrote it" → Verification is independent of author
- "PR review will catch issues" → PR review is not your verification gate

---

**Core Principle**: Verification is not optional. Before you claim done, you must have fresh output proving the behavior works. This prevents 90% of escaped bugs and agent hallucinations.
