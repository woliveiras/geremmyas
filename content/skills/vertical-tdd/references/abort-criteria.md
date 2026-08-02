---
name: abort-criteria
description: "Decision gates for when to STOP a task and escalate. Use when: time budget exceeded, circular debugging detected, scope creeping, or architectural mismatch found. Do not use: for normal task breaks, PR reviews, or acceptance testing."
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
- Pause the current approach
- Summarize findings so far
- Compare options with a specialist agent when useful:
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
- Stop the current fix loop
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
- Pause the expanded scope
- Document what changed
- Split into separate tasks
- If the objective and acceptance criteria still determine the choice, update
  the durable artifacts and continue. Escalate only unresolved material product
  ambiguity.

**Red flag**: "While we're in here..." = scope creep, not progress

---

### Signal 4: Unknown Unknowns

**Threshold**: Task blocked on something you don't understand

**Examples**:
- "How does this library work?" (30 min investigation, still unclear)
- "Where is this config?" (searched 20 files, not found)
- "What's the expected behavior?" (spec is ambiguous)

**What to do**:
- Stop guessing and gather evidence:
  - For product ambiguity, inspect specs and existing behavior, then use a
    spec/product specialist before escalating to the user
  - For library questions, read installed documentation/source and delegate to
    a framework specialist (don't spend >1 hour on the same approach)
  - For config, trace repository history and existing examples
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
- Stop trying to patch the symptom
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
- Pause implementation
- Document the conflict with evidence
- Ask an independent spec reviewer to reconcile the evidence first:
  - "Spec says X, but this implementation is impossible because Y"
  - "These two requirements contradict"
  - "Cost is 10x the budget"
- Update the durable spec automatically when intent is still unambiguous.
  Escalate only if the contradiction leaves a material product choice.

**Red flag**: Trying to force contradictory requirements = waste

---

### Signal 7: Architectural Mismatch

**Threshold**: Current architecture doesn't support the feature

**Examples**:
- Feature needs real-time sync but system is request-response only
- Feature needs persistence but current design is stateless
- Feature needs horizontal scaling but system is monolithic

**What to do**:
- Pause feature implementation
- Document the architectural gap
- Options:
  - Option A: Redesign architecture (big effort)
  - Option B: Simplify feature to fit architecture (spec change)
  - Option C: Use workaround (debt, temporary)
  - Option D: Defer feature until architecture ready
- Record the trade-off and use an architect/reviewer to challenge the choice.
  Continue when rollback is proven and the critical harness exists. Escalate
  only when blast radius is high and at least one evidence gap remains:
  rollback is unproven, or critical harness evidence is missing.

**Red flag**: Trying to force feature into incompatible architecture = technical debt debt

---

### Signal 8: You're Not the Right Person

**Threshold**: Task needs expertise you don't have

**Examples**:
- Task requires Kubernetes experience, you've never used it
- Task is security-critical, you're not security-trained
- Task requires domain knowledge you lack

**What to do**:
- Pause and recruit the missing expertise
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
- [ ] Have I used the relevant specialist agent or escalation path?
- [ ] Is the time spent proportional to problem size?
- [ ] Have I documented my findings so far?
- [ ] Is there a clear next step or am I guessing?

If you answer "no" to any → pause, gather evidence, and choose a new approach.

## Escalation Path

**For material product ambiguity**:
```
First: inspect durable artifacts and ask a product/spec specialist
Escalate to user: only when materially different valid outcomes remain
Get: written decision recorded in the spec
```

**For technical blocks**:
```
Ask: specialist subagent, codebase expert, framework expert
Present: what you tried, what failed, error messages
Get: guidance or hands-on help
After 3 failed cycles on the same blocker: escalate with evidence
```

**For architectural decisions**:
```
Ask: architect and independent reviewer agents
Present: the problem, options, tradeoffs
Get: challenged decision, proven rollback, and executable harness evidence
Escalate to user only when blast radius is high and at least one evidence gap
remains: rollback is unproven, or critical harness evidence is missing
```

**For conflicts**:
```
Ask: spec/product specialist first
Present: the conflict with evidence (not opinion)
Get: evidence-backed spec update; ask the user only if material ambiguity remains
Time: before resuming implementation
```

---

## Anti-Patterns

**"Just Push Through"**
- Sunk cost: "I've already spent 8 hours"
- Stubbornness: "I'll figure it out"
- Pride: "I don't want to ask for help"

**Fix**: Abort is not failure. It means changing approach, recruiting expertise,
or escalating at the defined authority boundary. Smart > persistent.

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
