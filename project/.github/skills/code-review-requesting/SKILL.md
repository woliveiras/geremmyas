---
name: code-review-requesting
description: "Structured code review workflow ensuring context and clarity. Use when: implementation complete, tests verified, before merge request. Do not use: during active implementation, before verification checklist, or for design reviews."
---


# Code Review Requesting

## When to Use

- After implementation complete and verified
- Before submitting PR or sharing code
- To prepare context for effective review
- When you need sign-off before merge

## When NOT to Use

- During development (use inline feedback)
- Before spec approval (review spec first, not code)
- For trivial changes (self-review may be enough)
- When reviewer is not the decision maker

## Pre-Review Checklist

**Code Quality**
- [ ] No commented-out code
- [ ] No debug `console.log` / `println` calls
- [ ] No temporary variables or obvious hacks
- [ ] Tested and verification evidence collected
- [ ] No merge conflicts, rebased on main

**Documentation**
- [ ] Complex logic has comments explaining "why"
- [ ] Public APIs have docstrings
- [ ] Change impact documented if non-obvious
- [ ] Tests updated for changed behavior

**Testing**
- [ ] New tests added for new behavior
- [ ] Regression tests pass
- [ ] Edge cases covered per spec
- [ ] Can run full suite locally without failures

**Context Preparation**

Frame the review:

```
## What Changed?
- Implemented [feature from spec]
- Refactored [component] for [reason: perf/clarity/design]
- Fixed [bug: issue title]

## Why This Approach?
- Chose [X] because [design decision from spec]
- Rejected [Y] because [tradeoff]
- See spec NNNN-slug/spec.md for full context

## Testing
- [X] New tests: 5 tests covering acceptance criteria
- [X] Regression: full suite passes, 0 regressions
- [X] Manual: verified scenario A, B, C

## Files Changed
- core/handler.ts: Updated to [behavior]
- tests/handler.test.ts: Added 5 new tests
- docs/README.md: Updated usage examples

## Review Focus
- Does implementation match spec acceptance criteria?
- Edge case handling: payment declined, network timeout, empty payload
- Performance: N+1 queries, memory leaks
- Security: input validation, error surfaces

## Known Unknowns
- Database query performance under load (no prod data to test)
- Mobile Safari testing (iOS 15 only, not older versions)
```

## Review Request Template

**For PR/MR:**

```
## Title
feat(auth): add email verification flow

## Description
Implemented email verification per spec 0003-auth-flow.
User receives code, submits, backend validates.

## Spec Reference
See specs/0003-auth-flow/spec.md for full acceptance criteria.

## Changes
- Added AuthService.sendVerificationEmail()
- Added VerificationController with POST /verify endpoint
- Added verificationCode table migration
- 8 new tests, all passing

## Testing
✅ Unit tests: 8/8 pass
✅ Integration: email mock tested
✅ Manual: tested code expiry, invalid codes, happy path

## Verification
See verification checklist in #comments

## Reviewer Checklist
- [ ] Code matches spec acceptance criteria
- [ ] No security regressions (validation, SQL injection, etc.)
- [ ] Tests cover happy + sad paths
- [ ] Error handling appropriate
```

## Common Reviewer Questions (Answer Preemptively)

**Design**
- "Why this structure?" → Reference spec and decision (in code comments)
- "Why not [alternative]?" → Document tradeoff upfront

**Testing**
- "What about [edge case]?" → List covered cases in review request
- "Tested on [platform]?" → State explicitly: "Tested on Node 22, not 20"

**Performance**
- "Any performance impact?" → Include benchmarks if relevant
- "Database queries?" → Include query count or EXPLAIN output

**Security**
- "Input validation?" → Show validation code or link
- "Error messages leak data?" → Document what's sanitized

## Red Flags — Fix Before Requesting Review

| Issue | Fix |
|-------|-----|
| No test additions or changes | Reconsider if tests are needed per spec |
| "Code is self-documenting" | Add comments explaining non-obvious decisions |
| Spec reference missing | Include link to spec.md |
| "Tested locally, should be fine" | Show actual test output |
| Whitespace-heavy changes | Use `git diff --ignore-all-space` to focus |
| No description, "pls review" | Write the structured description above |
| Feedback loop unclear | Specify: "Approve if OK, else suggest fixes" |

## Review Cycle Management

**After Requesting Review**

1. **Respond to feedback**
   - Don't dismiss suggestions without reasoning
   - If disagree, explain why per spec/design
   - If agree, apply and re-verify

2. **Re-request after changes**
   - Don't ghost reviewers
   - Link to updated verification evidence
   - Summarize fixes applied

3. **Accept review decision**
   - Approval means code is good AND spec-compliant
   - Rejection means fix, don't rationalize
   - Only override on explicit escalation

---

**Key Principle**: Effective review saves rework. Pre-frame context, answer common questions upfront, show verification evidence. Make reviewer's job easy, decisions fast.
