# Guardrails Framework

Geremmyas keeps safety rules close to the phase where they apply. Mandatory
invariants live in `AGENTS.md`; detailed procedures live as references inside
the workflow that owns them. Only capabilities users invoke directly are
advertised as skills.

## Lifecycle

| Phase | Public workflow | Guardrail |
| --- | --- | --- |
| Requirements | `requirements-interview` | Explore first, record commit permission, classify the change |
| Specification | `generate-spec` | Create spec, plan, tasks, then prove machine readiness |
| Implementation | `vertical-tdd` | One observable behavior per red-green-refactor cycle |
| Bugfix | `bugfix-loop` | Reproduce, rank hypotheses, prove red, fix, prove green, document cause |
| Completion | `verification-checklists` | Fresh focused and nearby-suite evidence before `[x]` |
| Review | `code-review-requesting` | Present scope, rationale, tests, risks, and unknowns |
| Commit | `git-commit` | Stage only approved files; no amend or push without permission |

## Internal References

These former top-level skills are now loaded only by their owning workflow:

| Former skill | Current owner |
| --- | --- |
| Feature readiness | `requirements-interview/references/approval-gates.md` and `AGENTS.md` |
| `task-breakdown` | `generate-spec/references/task-breakdown.md` |
| `generate-tests-from-spec` | `vertical-tdd/references/generate-tests-from-spec.md` |
| `abort-criteria` | `vertical-tdd/references/abort-criteria.md` |
| `regression-testing` | `bugfix-loop/references/regression-testing.md` |
| `agent-rationalization-blocking` | `verification-checklists/references/rationalization.md` |
| `subagent-selection` | `content/agents/references/subagent-selection.md` and `AGENTS.md` |

`decision-framework` remains available through the opt-in `decision-support`
pack. `skill-authoring` remains available through `skill-maintenance`.

## Gates

`catalog/workflow-gates.json` is the durable migration inventory. It classifies
each conversational pause as removed, a retained authority boundary, residual
evidence, or deferred release work. `geremmyas lint` scans every catalogued
source plus prompts, target adapters, this document, and `README.md`; an
unclassified match fails lint. Rules marked removed remain visible while their
matching text must stay absent; a match fails lint. Boundaries scheduled for a
later spec 0008 slice remain classified as retained until that slice changes
both the workflow and inventory.

### Feature readiness gate

Production code and feature tests wait until `spec.md`, `plan.md`, and
`tasks.md` exist and are machine-ready: acceptance criteria are testable,
contracts and dependencies are known, verification commands exist, and no
material product decision remains unresolved. In-scope discoveries update and
revalidate the artifacts automatically. Explicit read-only, plan-only, and
no-edits instructions remain narrower session overrides.

The lifecycle is `Draft` -> `Ready` -> `In Progress` -> `Verified`. The first
`[~]` implementation task starts `In Progress`. Failed evidence keeps or returns
the package there; a newly unresolved material decision returns it to `Draft`.
`Verified` requires evidenced criteria, completed tasks, reconciled docs, and an
independent review with no actionable finding. Commit, push, merge, and release
evidence are recorded separately.

### Bugfix evidence gate

Every bug has a bugfix document, a deterministic reproduction when feasible,
ranked hypotheses, and a regression test that fails for the expected reason
before production code changes. The agent then applies the smallest fix, reruns
the original reproduction plus focused and nearby tests, records the actual
cause, and removes temporary instrumentation. The loop does not pause for
routine proposal acknowledgement. Escalation is limited to a material objective
change, an external or production authority boundary, or the same unresolved
blocker after three consecutive evidence cycles.

### Completion gate

A completion claim needs fresh command output. Run the focused test, the nearest
relevant suite, remove temporary instrumentation, and reconcile task/spec state.
Confidence, compilation alone, or stale CI output is not evidence.

## Delegation

Delegate independent, read-heavy work only when the returned summary will be
smaller than inline exploration. Keep narrow edits and simple searches inline.
Subagents report scope, evidence, unknowns, and a concise result; they do not
edit shared state in parallel.

Use one bounded role per question:

| Role | Boundary |
| --- | --- |
| `explorer` | Requested subsystem, direct flows, sampled conventions |
| `spec-writer` | Requested behavior, affected modules, specs, and tests |
| `reviewer` | Requested diff, governing spec, direct tests, affected boundaries |
| `architect` | One module cluster and its direct callers and dependencies |

Architecture fan-out is conditional. Use up to three parallel alternatives only
when the interface decision is material, hard to reverse, and independently
investigable. Compare routine refactor options inline.
