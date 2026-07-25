---
name: approval-gates-before-implementation
description: "HARD GATE: Prevents implementation without explicit spec approval and user confirmation. Use when: starting a feature, before any code. Do not use: during implementation, for tiny refactors, or when approval already granted."
---


# Approval Gates Before Implementation

<HARD-GATE>
Do NOT write production code, scaffold projects, or implement features
until the spec has been presented AND user has explicitly approved it.
No "seems reasonable"—requires explicit approval.
</HARD-GATE>

## When to Use

- Starting a new feature or significant change
- Before you've written a single line of implementation code
- After spec/plan/tasks are ready but spec not yet approved
- To block agent rationalization patterns

## When NOT to Use

- During implementation (too late)
- For bug fixes with clear root cause (use bugfix-loop)
- For tiny refactors or obvious syntax fixes
- When approval has already been granted

## Red Flags — STOP and Present to User

| Excuse | Reality | Fix |
|--------|---------|-----|
| "This is simple, doesn't need a spec" | Simple ≠ obvious design | Write the spec anyway |
| "I'll approve it for myself" | You're not the user | Stop and wait for user sign-off |
| "Probably works" | Confidence ≠ evidence | Show the spec, get explicit "approved" |
| "Just going to scaffold" | Scaffolding is implementation | Present plan first |
| "Obviously the right approach" | Obvious to whom? | Challenge the assumption |
| "This is a refactor, not a feature" | Still changes behavior → still needs approval | Verify with user first |

## Procedure

1. **Build the spec, plan, tasks** (see `generate-spec` skill)
   - Problem statement clear
   - Acceptance criteria written
   - Test strategy documented

2. **Present the complete spec to the user**
   - Include spec.md, plan.md, tasks.md summary
   - Ask: "Does this solve the right problem?"
   - Ask: "Any blockers or unknowns?"

3. **Collect explicit approval**
   - User says: "approved", "looks good", "go ahead", etc.
   - Document: capture the approval in task history or commit context

4. **Only then begin implementation**
   - Execute `vertical-tdd` or equivalent per task
   - No deviations from approved spec without re-approval

5. **If spec changes during implementation**
   - Stop implementation
   - Update spec, plan, tasks
   - Re-present to user and get re-approval
   - Resume implementation

## Anti-Patterns

**Agent Rationalization**
- Skipping spec → "it's just X, obvious"
- Approving for itself → "I've reviewed it, ready to code"
- Self-validating → "This is correct per the code I haven't written yet"

**User Bypass**
- "User didn't object, so it's approved" → Wrong. Silence ≠ approval
- "I inferred what they want" → Present spec, don't infer
- "Similar feature existed once" → That doesn't transfer to this context

**Scope Creep**
- "While we're in here, let's also..." → Only if user approves expanded scope
- "Small polish changes" → Polishing behavior is still a change; document it

---

**Key Principle**: Approval gates prevent wasted work, misaligned solutions, and agent rationalization. The gate is not bureaucracy—it's insurance against building the wrong thing well.
