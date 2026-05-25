---
name: premortem
description: "Run a premortem on a plan, decision, or launch. Assumes it already failed and works backward to find why. Use when: premortem this, what could kill this, stress test this plan, find the blind spots, what am I missing in this plan, future-proof this decision."
---

# Premortem

Assume the plan already failed 6 months from now. Work backward to find every
reason why. Produce a revised plan with blind spots exposed.

Based on Gary Klein's prospective hindsight method. The frame shift from "what
could go wrong?" to "this already failed, explain why" produces more specific
and honest failure identification.

## When to Use

- A product, feature, or architecture you're about to commit to
- A launch plan, pricing change, or strategy pivot
- Any decision where the cost of being wrong is high and you can still change course

## When NOT to Use

- Vague ideas with no concrete plan (help them plan first)
- Decisions already made and irreversible
- Simple factual questions or routine code reviews
- Requests for creative feedback on a draft

## Minimum Context Threshold

Before running the premortem, you need three things:

1. **What is it?** — The thing being premortemed, describable in one sentence.
2. **Who is affected?** — The audience, customer, team, or stakeholders.
3. **What does success look like?** — The outcome the user hopes for.

Scan the conversation and workspace for existing context first. Only ask for
what's genuinely missing. One question at a time, conversational, not a form.

## Procedure

### 1. Set the frame

State the premise explicitly:

> It's 6 months from now. [The plan] has failed. It's done. We're looking back
> to understand what went wrong.

This framing is the mechanism. Without it, analysis defaults to polite risk
assessment.

### 2. Generate failure reasons

List every genuine reason the plan could have died. Each reason must be:

- Specific to this plan (not generic advice)
- Grounded in actual details provided
- A genuine threat (not a minor inconvenience or extremely unlikely edge case)

Find every real failure mode. Don't stop at 3 if there are 7. Don't pad to 7 if
there are 3.

### 3. Synthesize

For each failure reason, state in 2-3 sentences: the underlying assumption it
exploits and one early warning sign to watch for.

Then produce the synthesis:

- **Most likely failure** — Which scenario is most probable and why.
- **Most dangerous failure** — Which would cause the most damage if it happened.
- **Hidden assumption** — The single biggest thing the user is taking for granted
  that they haven't questioned.
- **Revised plan** — Concrete changes that make the plan more resilient. Each
  revision maps to a specific failure scenario. Be specific: not "consider
  testing pricing" but "run a $47 pilot with 20 people before committing to the
  full workshop."
- **Pre-launch checklist** — 3-5 specific things to verify, test, or put in
  place before executing. Each one prevents or detects a failure mode identified
  above.

### 4. Output

Print the full premortem report in chat as structured markdown. No HTML files,
no separate artifacts unless the user asks.

## Rules

- Always set the premortem frame explicitly before generating failures.
- Be comprehensive but not padded.
- Don't sugarcoat. The point is to tell the user things they don't want to hear
  before reality does.
- The revised plan must be concrete and actionable this week.
- If the user seems to want multiple perspectives on a decision rather than
  failure analysis, suggest a different approach (pros/cons, decision matrix).
