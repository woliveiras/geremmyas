# Guardrails Framework

Geremmyas keeps safety rules close to the phase where they apply. Mandatory
invariants live in `AGENTS.md`; detailed procedures live as references inside
the workflow that owns them. Only capabilities users invoke directly are
advertised as skills.

## Lifecycle

| Phase | Public workflow | Guardrail |
| --- | --- | --- |
| Requirements | `requirements-interview` | Explore first, record commit permission, classify the change |
| Specification | `generate-spec` | Create spec, plan, tasks, then stop for explicit approval |
| Implementation | `vertical-tdd` | One observable behavior per red-green-refactor cycle |
| Bugfix | `bugfix-loop` | Reproduce, rank hypotheses, stop for approval, add regression test |
| Completion | `verification-checklists` | Fresh focused and nearby-suite evidence before `[x]` |
| Review | `code-review-requesting` | Present scope, rationale, tests, risks, and unknowns |
| Commit | `git-commit` | Stage only approved files; no amend or push without permission |

## Internal References

These former top-level skills are now loaded only by their owning workflow:

| Former skill | Current owner |
| --- | --- |
| `approval-gates-before-implementation` | `requirements-interview/references/approval-gates.md` and `AGENTS.md` |
| `task-breakdown` | `generate-spec/references/task-breakdown.md` |
| `generate-tests-from-spec` | `vertical-tdd/references/generate-tests-from-spec.md` |
| `abort-criteria` | `vertical-tdd/references/abort-criteria.md` |
| `regression-testing` | `bugfix-loop/references/regression-testing.md` |
| `agent-rationalization-blocking` | `verification-checklists/references/rationalization.md` |
| `subagent-selection` | `.github/agents/references/subagent-selection.md` and `AGENTS.md` |

`decision-framework` remains available through the opt-in `decision-support`
pack. `skill-authoring` remains available through `skill-maintenance`.

## Gates

### Feature gate

Production code and feature tests wait until `spec.md`, `plan.md`, and
`tasks.md` exist and the user explicitly approves them. A material spec change
reopens the gate.

### Bugfix gate

Every bug has a bugfix document and reproduction. Present hypotheses, proposed
fix, and regression-test boundary, then stop for approval before changing
production code.

### Completion gate

A completion claim needs fresh command output. Run the focused test, the nearest
relevant suite, remove temporary instrumentation, and reconcile task/spec state.
Confidence, compilation alone, or stale CI output is not evidence.

## Delegation

Delegate independent, read-heavy work only when the returned summary will be
smaller than inline exploration. Keep narrow edits and simple searches inline.
Subagents report scope, evidence, unknowns, and a concise result; they do not
edit shared state in parallel.
