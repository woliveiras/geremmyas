# Guardrails Framework

Error-prevention system for AI coding agents. Eight skills that form hard gates, decision frameworks, anti-pattern detection, and quality workflows.

**Problem solved**: AI agents rationalize, skip gates, claim work complete without proof, implement features without approval, miss language-specific concerns.

**Solution**: Mandatory gates + fresh verification + anti-pattern detection.

---

## Framework Overview

### 1. Hard Gates (Blocks Work)

#### `approval-gates-before-implementation`

**When**: Before writing any production code for a feature.

**Gate**: Spec + plan + tasks must exist AND user must explicitly approve.

**Blocks these rationalizations**:
- "I can see what's needed, I'll implement while designing"
- "It's so simple I don't need a spec"
- "Self-approval is fine, I understand the requirement"
- "I'll refactor to match spec later"

**Usage**:
```
1. Create spec.md + plan.md + tasks.md with generate-spec
2. Present spec to user → STOP
3. Wait for explicit "Looks good" or "Approve"
4. Only then: vertical-tdd, regression-testing, code-review-requesting
```

#### `verification-checklists`

**When**: After implementation, before claiming any task is complete.

**Gate**: [Run] → [Output captured] → [Verified: X] evidence required.

**Blocks these rationalizations**:
- "I think it works"
- "Tests passed in my local, should be fine"
- "I didn't break anything obvious"

**Template**:
```
- [Run] `npm test -- LoginFlow`
- [Output] PASS: 42 tests, all green
- [Verified] Login with expired session: returns 401 ✓
- [Decision] TASK COMPLETE
```

---

### 2. Decision Frameworks (Prevent Hasty Choices)

#### `decision-framework`

**When**: Before choosing architecture, major tools, algorithms, or process changes.

**Mandate**: Context → Options → Decision → Reversibility + Bias check.

**Catches**:
- Confirmation bias (only exploring one option)
- Sunk cost (continuing bad path because "already invested")
- Authority bias (taking recommendation without evaluation)
- Emotional decisions (picking "cool" over "right")

**Pattern**:
```
DECISION: Replace sqlite3 with PostgreSQL

CONTEXT: Multi-tenant SaaS, 50k users, high concurrency

OPTIONS:
1. PostgreSQL (better concurrency, managed backups)
2. Neon serverless (PostgreSQL + scale-to-zero)
3. CockroachDB (distributed, overkill for now)

REVERSIBILITY: High (data migration tool exists)

TRADEOFFS:
- Latency: +5ms cold start (Neon) vs <1ms (self-hosted)
- Cost: $50/mo (Neon) vs $200/mo (Postgres managed)

DECISION: PostgreSQL via Neon (best of scaling + cost)

RISKS CHECKED:
- Vendor lock? Minimal (standard SQL)
- Migration path? Well-tested
- Sunk cost? Not present (new feature)
```

#### `subagent-selection`

**When**: Deciding whether to do work inline or delegate to a specialist agent.

**Matrix**:
- Code exploration + question answering? → `Explore` subagent
- Architecture improvement needed? → `architect` agent
- Spec writing from unclear requirements? → `spec-writer` agent
- Code review against spec? → `reviewer` agent
- Simple 1-2 file edits? → Inline

**Cost-benefit check**:
```
Task: Find all GraphQL resolvers that call external APIs

Time inline: 15 min (grep + manual review)
Subagent: 2 min (Explore with context)
Context cost: Small

→ USE SUBAGENT
```

---

### 3. Detection & Blocking (Catch Anti-Patterns)

#### `agent-rationalization-blocking`

**When**: Pre-implementation, pre-completion checkpoint.

**Identifies these excuses**:

| Excuse | Reality | Fix |
|--------|---------|-----|
| "It's simple/obvious" | Requirements not verified | Get approval from spec |
| "Probably fine" | Untested edge case | Regression test required |
| "I'm experienced here" | Overconfidence bias | Test anyway |
| "Tests are overkill" | Skipping verification gate | Add minimal test |
| "I'll refactor later" | Technical debt accumulates | Do it now |
| "No one will notice" | Produces silent bugs | Add test + code review |

**Usage**: Read before implementation starts. If you catch yourself thinking any of these, use the "Fix" column.

#### `abort-criteria`

**When**: Monitoring a task that seems stuck.

**Abort signals** (STOP when 2+ triggered):
1. **Time Budget Exceeded**: Planned 2 hours, now at 4 hours
2. **Circular Debugging**: Tried 5 approaches, back to square one
3. **Scope Creep**: Task bloated 3x original size
4. **Unknown Unknowns**: "I don't know what I don't know" repeating
5. **Test Failures**: Can't make tests pass, can't understand why
6. **Spec Conflict**: Task contradicts approved spec
7. **Architectural Mismatch**: Task doesn't fit system design
8. **Wrong Person**: Task needs domain expert, you're not it

**Action on abort**:
```
STOP. Do not push.
Document: What blocked? What was tried?
Escalate: Ask user for clarification or help.
```

---

### 4. Quality Workflows (Mandatory for Code & Bugs)

#### `regression-testing`

**When**: Creating a fix for any bug. BEFORE applying the fix.

**Mandatory**: One test that FAILS before fix, PASSES after fix.

**Multi-language**:

**Go**:
```go
func TestLoginWithExpiredSessionReturnsUnauthorized(t *testing.T) {
    session := createSessionWithExpiry(time.Now().Add(-1 * time.Hour))
    resp := httpClient.Post("/api/login", loginRequest{sessionID: session.ID})
    assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
```

**Python**:
```python
def test_login_with_expired_session_returns_unauthorized():
    session = create_session_with_expiry(datetime.now() - timedelta(hours=1))
    with pytest.raises(UnauthorizedError):
        login(session.id)
```

**JavaScript**:
```typescript
test("LoginWithExpiredSessionReturnsUnauthorized", async () => {
  const session = createSessionWithExpiry(Date.now() - 3600000);
  await expect(login(session.id)).rejects.toThrow("unauthorized");
});
```

**Verification**:
```bash
1. [Run] test BEFORE fix → FAILS ✓
2. Apply fix
3. [Run] test AFTER fix → PASSES ✓
4. [Run] full suite → no regressions ✓
```

#### `code-review-requesting`

**When**: Implementation complete, ready for review.

**Pre-review checklist**:
- [ ] Spec requirements all implemented
- [ ] Tests passing (unit + integration)
- [ ] No merge conflicts
- [ ] Docs updated
- [ ] No temporary debug code

**Review request**:
```markdown
## What changed

Fixed login returning 500 for expired sessions.

## Why

Expired session tokens weren't validated before processing.
Now returns 401 (Unauthorized) with clear error.

## How to test

1. Create session with past expiry
2. POST /api/login with that session_id
3. Expect 401 + "session expired" error

## Regression test

LoginWithExpiredSessionReturnsUnauthorized (new)
```

---

## Integration Example: Bug Fix Workflow

**Scenario**: Login endpoint returns 500 for expired sessions (should return 401).

```
Step 1: USE bugfix-loop SKILL
├─ Reproduce: Create expired session, call login → 500 error
├─ Hypothesize: Session validation missing
└─ [HARD GATE] Present hypothesis → wait for approval

Step 2: USE regression-testing SKILL
├─ Write test that FAILS with current code
└─ [HARD GATE] Verify test fails before touching production code

Step 3: Apply fix

Step 4: USE verification-checklists SKILL
├─ [Run] regression test → PASSES
├─ [Run] full suite → no regressions
└─ [Decision] Fix verified

Step 5: USE code-review-requesting SKILL
├─ Pre-review checklist (docs, tests, conflicts)
└─ Request review with context

Step 6: USE git-commit SKILL (if user granted permission)
```

---

## Integration Example: Feature Implementation

**Scenario**: Build a dashboard with charts and filters.

```
Step 1: USE requirements-interview SKILL
├─ Clarify: What data? Which charts? Filter criteria?
├─ Ask: May I commit? (Store answer for session)
└─ [HARD GATE] Collect explicit commit permission

Step 2: USE generate-spec SKILL
├─ Create specs/NNNN-dashboard/spec.md
├─ Create specs/NNNN-dashboard/plan.md
├─ Create specs/NNNN-dashboard/tasks.md
└─ [HARD GATE] Present spec → wait for approval

Step 3: FOR EACH TASK, USE vertical-tdd SKILL
├─ [Run] Write failing test
├─ [Green] Implement minimum to pass
├─ [Refactor] Clean code
└─ USE verification-checklists → proof of passing test

Step 4: BEFORE CODE REVIEW, check decision-framework SKILL
├─ Did I make any major tool/arch choices?
├─ Are they documented with trade-offs?
└─ Continue

Step 5: USE code-review-requesting SKILL
├─ Gather context, code quality checks
└─ Submit with clear narrative

Step 6: USE update-docs SKILL
└─ Sync architecture / setup docs

Step 7: USE git-commit SKILL (if permission granted)
```

---

## When Guardrails Save Time

| Scenario | Without Guardrails | With Guardrails |
|----------|-------------------|-----------------|
| Implement without spec | Build wrong thing (1-2 days wasted) | Stop early, clarify, save days |
| Skip regression test | Bug reappears in production | Caught in CI, fixed same day |
| Implement 3 options, pick emotionally | Choose wrong, revert later | Decision-framework forces evaluation |
| Task seems endless | Keep trying, no progress signal | abort-criteria says STOP, escalate |
| Code review goes back/forth | Multiple rounds of "forgot X" | Pre-review checklist prevents surprises |
| Copy-paste logic error | Silent bug, hard to reproduce | Specific regression test catches it |

---

## Checklist: Load Guardrails in Your Session

When starting any task:
- [ ] Load `approval-gates-before-implementation` before design
- [ ] Load `decision-framework` before major choices
- [ ] Load `agent-rationalization-blocking` as pre-implementation check
- [ ] Load `regression-testing` before any bug fix
- [ ] Load `verification-checklists` before claiming "done"
- [ ] Load `code-review-requesting` before review submission
- [ ] Load `abort-criteria` if a task feels stuck
- [ ] Load `subagent-selection` if unsure whether to delegate

**Result**: 8 gates that prevent 80% of agent errors (self-approval, skipped verification, missed regressions, overconfidence).
