---
name: agent-rationalization-blocking
description: "Anti-pattern catalog identifying common self-deception in agent workflows. Use when: pre-completion checkpoint, before claiming work is done, validating agent reasoning. Do not use: for external code review or design phases before implementation."
---


# Agent Rationalization Blocking

<HARD-GATE>
Before claiming work is complete, before skipping a gate, before saying
"this is safe to merge"—STOP and check if you're rationalizing.

Rationalization looks like confidence. It is not.
</HARD-GATE>

## When to Use

- Before claiming a task complete
- Before skipping spec review or approval
- Before declaring "no test needed"
- When you feel certain a change is safe
- Before proceeding without verification

## When NOT to Use

- During brainstorming (creative exploration OK)
- For rapid prototyping (label it as such)
- Post-mortem analysis (reflection, not blocking)

## Rationalization Catalog

### "This is simple" Rationalization

**What agent says**: "This is a simple change, doesn't need spec/tests/review"

**Why it's dangerous**:
- "Simple" code + wrong spec = wasted work
- Simple changes compound into complex bug patterns
- "Simple" hides assumptions

**How to block it**:
- Every change needs spec approval, even 1-liners
- "Simple to implement" ≠ "simple to understand impact"
- Run tests anyway. Review anyway.

**Reality check**: Is it simple in isolation, or simple in context? Context is usually wrong.

---

### "I Approved It For Myself" Rationalization

**What agent says**: "I reviewed the spec and it's good"

**Why it's dangerous**:
- Agent is not the user
- Agent cannot represent user intent
- Agent can rationalize away conflicts

**How to block it**:
- STOP. Present spec to actual user.
- Collect explicit "approved" from user.
- Document user approval in commit message or task comment.

**Reality check**: Would I accept this approval if a different agent said it? If no → don't accept from self.

---

### "Probably Works" Rationalization

**What agent says**: "This should work, pretty confident, didn't test it yet"

**Why it's dangerous**:
- Confident bugs are still bugs
- "Should work" escapes to prod as "surprise failure"
- Confidence is not evidence

**How to block it**:
- Run the test. Capture output.
- Show the verification checklist (see verification-checklists skill).
- "Probably" = 0. "Tested and passing" = 100.

**Reality check**: If I said "probably works" about critical code, would I accept that? Then don't accept it from agent.

---

### "I Already Know This Pattern" Rationalization

**What agent says**: "I implemented this before in another project, it'll work here"

**Why it's dangerous**:
- Context changes between projects
- Dependencies differ
- Constraints differ
- "It worked once" ≠ "it works now"

**How to block it**:
- Write it fresh per this context
- Tests prove it works here, not in memory
- Don't copy-paste old solutions without spec review

**Reality check**: Is this project identical to the one where I used this pattern? (Spoiler: no.)

---

### "The Spec Is Obvious" Rationalization

**What agent says**: "The spec doesn't need details, we all know what 'login' means"

**Why it's dangerous**:
- Obvious to whom?
- Obvious in whose experience?
- "Obvious" login = 10 different interpretations

**How to block it**:
- Write the spec anyway
- Detail is not wasted—it's communication
- Obvious ≠ written

**Reality check**: If I hand this spec to someone who's never worked with this codebase, would they understand it? If no → write more.

---

### "Tests Are Overkill" Rationalization

**What agent says**: "The code is so simple, tests aren't needed"

**Why it's dangerous**:
- Simple code breaks in complex systems
- No test = no proof it works
- "Tested by inspection" escapes as regression

**How to block it**:
- Test every behavior from spec
- No exceptions for "simple"
- Tests are insurance, not overhead

**Reality check**: If this code breaks in prod, can we reproduce it fast? (Only if tests exist.)

---

### "I'll Refactor It Later" Rationalization

**What agent says**: "This is a bit hacky but we'll improve it next sprint"

**Why it's dangerous**:
- "Later" rarely happens
- Technical debt compounds
- Next sprint: new priorities, same hack

**How to block it**:
- Don't merge hack code
- Write it right the first time
- If short on time: defer the feature, not the quality

**Reality check**: Has "we'll fix it later" ever happened in this codebase? (Spoiler: rarely.)

---

### "No One Will Notice" Rationalization

**What agent says**: "This edge case is unlikely, we don't need to handle it"

**Why it's dangerous**:
- "Unlikely" happens in prod
- Unhandled edges = bad UX
- "No one will notice" = someone will, in prod

**How to block it**:
- Spec lists edge cases
- Handle per spec
- Can't skip because it's "unlikely"

**Reality check**: If this breaks for a user, how bad is it? (If "bad" → handle it.)

---

### "I'm Confident This Is Secure" Rationalization

**What agent says**: "I validated input once, should be fine"

**Why it's dangerous**:
- One validation point ≠ layered defense
- Confidence != security review
- Attacks evolve faster than memory

**How to block it**:
- Security is not a feeling
- Use checklists (authentication, authorization, injection, secrets, etc.)
- Have security-specialized person review

**Reality check**: If a security expert reviewed this, would they approve? If "maybe" → need review.

---

### "I Changed My Mind, Here's The New Version" Rationalization

**What agent says**: "I decided the approach was wrong, here's a complete rewrite"

**Why it's dangerous**:
- Spec approval was for old approach
- New approach needs spec approval too
- Reviewers don't know what changed

**How to block it**:
- Stop and re-present spec
- New approach = new spec approval gate
- Don't rewrite without telling user

**Reality check**: Is this still solving the same problem per the same spec? If no → gate.

---

## Pre-Completion Checklist

Before marking task complete:

- [ ] Did I get explicit approval for the spec? (not self-approval)
- [ ] Did I run the verification and capture output?
- [ ] Can I defend this approach on technical merit, not confidence?
- [ ] Have I tested the edge cases, not just happy path?
- [ ] Is there a test that proves this works?
- [ ] Would I accept this from another agent, or am I self-rationalizing?

If you answer "no" to any, go back to that phase.

---

## The Rationalization Test

**Ask yourself**: "Am I saying yes because [reason] or because I want this to be done?"

If the answer is "because I want it done" → you're rationalizing. Stop.

**Key Principle**: Rationalization is the leading cause of agent errors. Catch it early. Run gates anyway. Get evidence always. Confidence is the enemy of verification.
