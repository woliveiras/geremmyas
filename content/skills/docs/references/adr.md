# ADR / MADR 4.0

Generate an Architectural Decision Record using the [MADR](https://adr.github.io/madr/) format.
Use this when recording a justified design choice — technology, library, pattern, or approach.

## When to Use

- You made (or are making) a decision that others will ask "why?" about later
- You're choosing between multiple viable options
- You want to document a decision that is already accepted

For unresolved architecture exploration, create a bounded runtime subagent or
use the RFC mode. For implementation sequencing, use `plan.md` and `tasks.md`
instead of an ADR.

## Before Creating an ADR

Create an ADR only when all three are true:

- The decision is hard to reverse.
- The decision would be surprising without context.
- The decision is the result of a real trade-off between viable options.

Do not create ADRs for obvious, trivial, or easy-to-reverse choices. If the
decision does not meet the bar, explain that no ADR is needed and suggest where
the context should live instead, such as a PRD, spec, or code comment.

## Procedure

1. Infer from the accepted decision and repository evidence:
   - The decision topic (short title)
   - The context / problem being solved
   - The options they considered (at least 2)
   - Which option was chosen and why
   Escalate only if a material decision remains unresolved after architecture review.
2. Determine the next ADR number: list existing files in `docs/decisions/` and increment
3. Fill in the template below — include optional sections only when the user provides enough detail
4. Save to `docs/decisions/NNNN-title-with-dashes.md`

## Template

```markdown
---
status: {proposed | accepted | deprecated | superseded by ADR-NNNN}
date: {YYYY-MM-DD}
---

# {Short title, representative of solved problem and found solution}

## Context and Problem Statement

{Describe the context and problem in 2-3 sentences. Articulate as a question when possible.}

## Decision Drivers

* {Force, concern, or constraint that influenced the decision}
* {Another driver}

## Considered Options

* {Option 1}
* {Option 2}
* {Option 3}

## Decision Outcome

Chosen option: "{Option}", because {justification — reference the decision drivers}.

### Consequences

* Good, because {positive consequence}
* Bad, because {negative consequence}

### Confirmation

{How will compliance with this decision be verified? Code review, automated test, lint rule, etc.}

## Pros and Cons of the Options

### {Option 1}

{Brief description or link}

* Good, because {argument}
* Bad, because {argument}

### {Option 2}

{Brief description or link}

* Good, because {argument}
* Bad, because {argument}

## More Information

{Links to related ADRs, issues, or discussions.}
```

## Rules

- Title must describe the decision, not the problem ("Use PostgreSQL for persistence", not "Database choice")
- Context section: 2-3 sentences max — link to issues for backstory
- Always list at least 2 considered options (even if one is "do nothing" or "status quo")
- Decision Outcome must reference the decision drivers — explain *why*, not just *what*
- Consequences: be honest about trade-offs — every decision has downsides
- Omit optional sections (Decision Drivers, Confirmation, Pros/Cons, More Information) if the decision is straightforward
- Use `status: proposed` for decisions under discussion, `status: accepted` for final decisions, `status: implemented` for decisions that have been implemented, `status: deprecated` for decisions that are no longer recommended, and `status: superseded by ADR-NNNN` for decisions that have been replaced by a newer ADR
- Number files sequentially: `0001-use-fastapi.md`, `0002-adopt-zustand.md`

When writing the decision, follow [the shared writing style](./writing-style.md).
