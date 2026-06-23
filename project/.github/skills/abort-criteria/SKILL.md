---
name: abort-criteria
description: "Decision gates for when to STOP a task. Prevents sunk-cost pushing and endless exploration. Use when a task becomes unclear, blocked, or complexity exceeds estimate."
---

# Abort Criteria

## When to Use

- When a task is taking longer than estimated
- When you hit an unexpected blocker
- When you're going in circles (same error repeatedly)
- When requirements became unclear mid-task
- When you're solving a different problem than intended
- Before sinking more time into uncertain work

## When NOT to Use

- During normal implementation (just keep going)
- For exploration phase (exploring is the point)
- For known hard work (estimated at high complexity)
- When close to done (finish > abort)

## Abort Signals

### Signal 1: Time Budget Exceeded

**Threshold**: Task estimated 2 hours, now at 4 hours

**What to do**:
- STOP
- Summarize findings so far
- Present options to user:
  - Option A: Continue (why you think you're close)
  - Option B: Reframe (the problem is smaller/different)
  - Option C: Escalate (need senior/specialist)
  - Option D: Defer (not critical path, do later)

**Red flag**: "Just need 1 more hour" said 3 times = abort signal

---

### Signal 2: Circular Debugging

**Threshold**: Same error appearing repeatedly after multiple fixes

**Pattern**:
```
Fix A → Error B
Fix B → Error C
Fix C → Back to Error A or new Error D
```

**What to do**:
- STOP the fix loop
- Step back and document:
  - What's the real problem?
  - Are fixes symptoms, not causes?
  - Is there a deeper architecture issue?
- Options:
  - Root cause analysis (bugfix-loop skill)
  - Escalate to senior
  - Revert to last known good, try different approach

**Red flag**: More than 2 cycles = underlying issue, not surface fix

---

### Signal 3: Scope Creep

**Threshold**: Task changed mid-implementation

**Examples**:
- Started: "Add login button"
- Now: "Redesign entire auth flow"
- Started: "Fix typo"
- Now: "Rewrite entire module for clarity"

**What to do**:
- STOP
- Document what changed
- Get re-approval from user for new scope
- Split into separate tasks
- Resume original task or start new one, not both

**Red flag**: "While we're in here..." = scope creep, not progress

---

### Signal 4: Unknown Unknowns

**Threshold**: Task blocked on something you don't understand

**Examples**:
- "How does this library work?" (30 min investigation, still unclear)
- "Where is this config?" (searched 20 files, not found)
- "What's the expected behavior?" (spec is ambiguous)

**What to do**:
- STOP pushing
- Escalate:
  - Ask team/user for clarification (not internet search)
  - For library questions: read source, ask in chat (don't spend >1 hour)
  - For config: ask who set it up (don't guess)
- Gather answer, resume

**Red flag**: >1 hour investigation without progress = escalate

---

### Signal 5: Test Failures Won't Resolve

**Threshold**: Changes don't fix the failing test

**Pattern**:
```
Test fails with: "X is nil"
Fix: Add null check
Test fails with: "X is nil" (still!)
Fix: Add more null checks
... continues ...
```

**What to do**:
- STOP trying to patch
- Examine the test itself:
  - Is the test wrong?
  - Is the setup wrong?
  - Am I misunderstanding the assertion?
- Run the test in isolation with debug output
- If still stuck: escalate with debug evidence

**Red flag**: Same failure after 3 different fixes = test issue

---

### Signal 6: Spec Conflict Discovered

**Threshold**: Code contradicts the spec, or spec is impossible to implement

**Examples**:
- Spec says: "Validate email" but also "Accept any string"
- Spec says: "10ms response time" but requires 3 DB queries
- Spec says: "Free tier" but implementation costs $10k/month

**What to do**:
- STOP implementation
- Document the conflict with evidence
- Present to user/spec-owner:
  - "Spec says X, but this implementation is impossible because Y"
  - "These two requirements contradict"
  - "Cost is 10x the budget"
- Get clarification or spec update before resuming

**Red flag**: Trying to force contradictory requirements = waste

---

### Signal 7: Architectural Mismatch

**Threshold**: Current architecture doesn't support the feature

**Examples**:
- Feature needs real-time sync but system is request-response only
- Feature needs persistence but current design is stateless
- Feature needs horizontal scaling but system is monolithic

**What to do**:
- STOP feature implementation
- Document the architectural gap
- Options:
  - Option A: Redesign architecture (big effort)
  - Option B: Simplify feature to fit architecture (spec change)
  - Option C: Use workaround (debt, temporary)
  - Option D: Defer feature until architecture ready
- Decide with leadership, then resume

**Red flag**: Trying to force feature into incompatible architecture = technical debt debt

---

### Signal 8: You're Not the Right Person

**Threshold**: Task needs expertise you don't have

**Examples**:
- Task requires Kubernetes experience, you've never used it
- Task is security-critical, you're not security-trained
- Task requires domain knowledge you lack

**What to do**:
- STOP and be honest
- Options:
  - Option A: Get help from expert (pair program)
  - Option B: Learn (if time allows and task not critical)
  - Option C: Escalate to expert
- Don't pretend competence to avoid asking for help

**Red flag**: >2 hours stuck and you don't know why = need help

---

## Abort Decision Checklist

Before continuing past a signal, verify:

- [ ] Is this in the spec or a scope change?
- [ ] Have I searched for prior art / similar solution in codebase?
- [ ] Have I asked for help or escalated?
- [ ] Is the time spent proportional to problem size?
- [ ] Have I documented my findings so far?
- [ ] Is there a clear next step or am I guessing?

If you answer "no" to any → ABORT or ESCALATE

## Escalation Path

**For clarification questions**:
```
Ask: user, product owner, spec author
Get: written answer (not interpretation)
Time: <15 min response expectation
```

**For technical blocks**:
```
Ask: senior engineer, codebase expert, framework expert
Present: what you tried, what failed, error messages
Get: guidance or hands-on help
Time: <1 hour for advice, schedule deep dive if needed
```

**For architectural decisions**:
```
Ask: architect, tech lead, team lead
Present: the problem, options, tradeoffs
Get: decision + approval
Time: decision should be same day
```

**For conflicts**:
```
Ask: spec owner, product, user
Present: the conflict with evidence (not opinion)
Get: clarification or spec change
Time: before resuming implementation
```

---

## Anti-Patterns

**"Just Push Through"**
- Sunk cost: "I've already spent 8 hours"
- Stubbornness: "I'll figure it out"
- Pride: "I don't want to ask for help"

**Fix**: Abort is not failure. It's knowing when to escalate. Smart > persistent.

**"Ignore the Signal"**
- Keep trying after time budget exceeded
- Don't ask for help despite being stuck
- Hope the error goes away

**Fix**: Signals exist for a reason. Honor them.

**"Change the Definition of Done"**
- Feature becomes 80% done, call it done
- Test failures ignored as "known issues"
- Abort by redefining success

**Fix**: Abort means pause, escalate, replan—not claim victory.

---

**Key Principle**: Abort is not quitting—it's professional judgment. Knowing when to stop, escalate, or replan is more valuable than forcing a bad solution. Smart work > long hours.
