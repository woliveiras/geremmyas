---
name: decision-framework
description: "Structured decision-making to prevent emotional/hasty choices. Use before committing to architectural, process, or tool decisions. Enforces trade-off documentation and reversibility assessment."
---

# Decision Framework

## When to Use

- Before architectural decisions (new framework, DB migration, splitting service)
- Major tool or process changes
- Go/no-go decisions on initiatives
- When torn between options and need clarity
- To document the "why" for future you

## When NOT to Use

- Trivial tactical choices (variable naming, minor refactors)
- Time-critical fire-fighting (document later)
- Implementation details already constrained by spec
- Decisions already made by policy/leadership

## Decision Template

**Context**
- What problem are we solving?
- What constraints exist? (time, budget, tech, team skill)
- Who decides? Who is affected?

**Options** (at least 2-3)
- Option A: [approach]
  - Pros: [specific, measurable]
  - Cons: [specific, measurable]
  - Cost: [effort, tokens, infrastructure]
  - Reversibility: [easy/hard/impossible to undo]
  - Timeline: [days/weeks/months]

- Option B: [approach]
  - Pros: ...
  - Cons: ...

**Decision**
- Chosen: [Option X]
- Key reason: [specific factor]
- Trade-off accepted: [what we're giving up]
- Contingency if wrong: [how we'll know + recovery plan]

**Documentation**
- Where: docs/decisions/NNNN-title.md (ADR format)
- Who approved: [decision maker]
- When revisit: [quarterly / after X metric / if condition]

---

## Red Flags — Pause and Reconsider

| Flag | What It Means | Response |
|------|---------------|----------|
| "Everyone does it this way" | Popularity ≠ fit | Ask: fit for OUR constraints? |
| "I just know this is best" | Intuition ≠ reasoning | Force yourself to list trade-offs |
| "No time to evaluate" | Pressure → hasty choice | 30-min structured decision > week of regret |
| "This is reversible" | Usually wrong | Test: can we switch back in 1 day? No? Hard. |
| "We'll fix it later" | Debt, not deferral | Can we? Will we? With what budget? |
| "I'm the expert, trust me" | Expertise ≠ group buy-in | Document so others understand reasoning |
| "Sunk cost already" | Past spend is not future cost | Evaluate on merits NOW, not backward |

## Reversibility Assessment

**High Reversibility** (safe to try)
- Upgrade to new version of library
- Switch linter config
- Refactor internal module structure
- Migration: 1-2 days to revert

**Medium Reversibility** (requires planning)
- Architectural shift (monolith → microservices)
- Database migration (SQL → NoSQL)
- Framework change (Express → Fastify)
- Recovery: weeks, known path

**Low/No Reversibility** (must be right)
- Delete data or drop DB table
- Public API contract change
- Move to closed-source platform
- Recovery: expensive, unclear

**For low-reversibility decisions**:
- Require explicit approval (see approval-gates-before-implementation)
- Run premortem (see premortem skill)
- Have rollback plan documented in advance
- Consider pilot/canary before full commit

## Decision Bias Traps

**Confirmation Bias**
- Seeking info that supports your preferred option
- Fix: Actively list strongest case for each option, write them down

**Sunk Cost**
- "We already invested in X, so X must be right"
- Fix: Evaluate on forward merits only. Ignore past spend.

**Availability Bias**
- "Kubernetes worked at my last job, so it works here"
- Fix: State context constraints explicitly. "Context here = [X]. Last job was [Y]."

**Authority Bias**
- "Senior person said so, must be right"
- Fix: Decision stands on merit, not title. Seniors can be wrong in new context.

**Groupthink**
- "Team consensus means it's right"
- Fix: Assign devil's advocate role. Force articulation of downsides.

**Emotional Urgency**
- "We MUST do this NOW or we'll fail"
- Fix: Distinguish true urgency from anxiety. Anxiety clouds reasoning.

## Collaboration Pattern

**1. Frame the Decision**
```
We need to decide: Should we add Kafka for event streaming?

Context:
- Current: in-process pub/sub, single instance
- Problem: scaling to N instances requires distributed messaging
- Constraint: team has 0 Kafka experience
- Timeline: 3 sprints before we hit single-instance limit

Stakeholders:
- Team lead: owns architecture decision
- Backend team: implements + operates
- DevOps: infrastructure cost/complexity
```

**2. Explore Options (30 min)**
- Kafka vs RabbitMQ vs SQS vs in-memory (Redis)
- Each: pros, cons, reversibility, learning curve

**3. Decide Together**
- Team votes or lead decides
- Capture trade-off: "Choosing Kafka for durability, accepting operations complexity"

**4. Document (5 min)**
- ADR file with decision + reasoning
- Link to ticket/spec that triggered it

**5. Revisit Schedule**
- Mark calendar: "Revisit Kafka decision in Q3"
- Condition: "If single-instance pub/sub becomes bottleneck again"

---

**Key Principle**: Good decisions aren't fast or popular—they're reasoned and trade-off-aware. Document the "why" so future you understands what constraints we optimized for.
