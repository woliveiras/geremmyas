---
name: subagent-selection
description: "Autonomous specialist orchestration, ownership partitioning, review loops, and fallback for targets without native delegation."
---

# Subagent Selection

Proactively delegate when specialist depth, independent challenge, context
compression, or parallel work improves the result. Keep trivial work inline.

## Delegation Envelope

Every assignment states the objective, governing artifacts, evidence expected,
forbidden areas, commands allowed, and explicit file, module, or worktree
ownership. Parallel editors require disjoint ownership. If agents need the same
file or hunk, serialize the work or let the primary create isolated worktrees.
Preserve user and unrelated agent work.

The primary agent owns integration and Git. Subagents never stage, commit, merge,
rebase, push, tag, release, publish, or deploy. The primary inspects returned
changes, resolves overlaps, verifies the integrated diff, and commits coherent
slices.

## Specialist Matrix

| Need | Role | Default boundary |
| --- | --- | --- |
| Requirements and artifacts | `spec-writer` | Behavior, spec, plan, tasks, index |
| Architecture | `architect` | Module cluster, alternatives, contracts |
| Production code | `implementer` | One behavior and owned implementation paths |
| Tests and harness | `test-engineer` | Acceptance criteria, regression, evidence |
| Security | `security-reviewer` | Trust boundary, diff, dependency, threat surface |
| Performance | `performance-reviewer` | Measured hot path, budget, workload |
| Documentation | `documentation` | Affected API, setup, architecture, operations |
| Spec conformance | `reviewer` | Diff, spec, direct tests, affected boundaries |
| Delivery audit | `auditor` | Scope, ownership, evidence, authority boundaries |
| Codebase mapping | `explorer` | Read-only subsystem map |

Security and performance work should be delegated early when they are acceptance
or risk boundaries, not deferred to the user. Reviewers and auditors remain
independent of the implementation context.

## Parallel Waves

Use parallel specialists when their ownership or review surfaces are independent.
Editors may share a worktree only for disjoint files; otherwise isolate or
serialize them. Never run concurrent Git operations. Do not redo a subagent's
completed investigation without a concrete evidence gap.

Run independent review after each slice or verification wave. Repair findings
and re-review automatically. Escalate only when the same blocker survives three
consecutive cycles; include attempts, evidence, remaining hypothesis, and the
specific decision or authority needed. A new blocker resets the counter.

## Target Fallback

When native subagents exist, invoke roles proactively. When they do not, use the
target's supported delegation mechanism or inline, applying the same role,
ownership, evidence, and independence contract in separate phases.

## Avoid

- Delegating a two-line mechanical edit with no specialist value.
- Giving two editors overlapping ownership.
- Letting a subagent mutate Git history or production.
- Treating disagreement as a user gate before comparing spec and evidence.
- Ignoring a completed specialist result and repeating the same exploration.
