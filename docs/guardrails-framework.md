# Guardrails Framework

Geremmyas keeps safety rules close to the phase where they apply. Mandatory
invariants live in `AGENTS.md`; detailed procedures live as references inside
the workflow that owns them. Only capabilities users invoke directly are
advertised as skills.

## Lifecycle

| Phase | Public workflow | Guardrail |
| --- | --- | --- |
| Ambiguous requirements | `refine` | Explore first, resolve one focused ambiguity, then unload |
| Specification | `spec` | Create spec, plan, tasks, then prove machine readiness |
| Implementation | `tdd` | One observable behavior per red-green-refactor cycle |
| Bugfix | `bugfix` | Reproduce, rank hypotheses, prove red, fix, prove green, document cause |
| Completion | `verify` | Fresh focused and nearby-suite evidence before `[x]` |
| Review | `verify/references/review-contract.md` | Bounded independent runtime subagent after verification |
| Documentation | `docs` | Load only the project docs, glossary, ADR, or RFC mode needed |
| Commit | `git-commit` | Commit verified task-owned slices locally; never infer push or history rewrite |

## Internal References

These former top-level skills are now loaded only by their owning workflow:

| Former skill | Current owner |
| --- | --- |
| Feature readiness | `refine/references/readiness-and-authority.md` and `AGENTS.md` |
| `task-breakdown` | `spec/references/task-breakdown.md` |
| Test generation | `tdd/references/test-generation.md` |
| `abort-criteria` | `tdd/references/abort-criteria.md` |
| `regression-testing` | `bugfix/references/regression-testing.md` |
| Rationalization blocking | `verify/references/rationalization.md` |
| Independent review | `verify/references/review-contract.md` and `AGENTS.md` |

`decision-framework` remains available through the opt-in `decision-support`
pack. `skill-authoring` remains available through `skill-maintenance`.

## Gates

`catalog/workflow-gates.json` is the durable migration inventory. It classifies
each conversational pause as removed, a retained authority boundary, residual
evidence, or deferred release work. `geremmyas lint` scans every catalogued
source plus prompts, target adapters, this document, and `README.md`; an
unclassified match fails lint. Rules marked removed remain visible while their
matching text must stay absent; a match fails lint.

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

### Local commit boundary

After verification and independent review, the primary agent creates one local
Conventional Commit per coherent task-owned slice by default. It stages only
owned paths or hunks, includes the slice's tests and required docs, rereads the
cached diff, and derives the message from that diff. Explicit `no-commits`,
read-only, plan-only, or no-edits instructions override the default.

Local commit authority does not include push, amend, rebase, merge, tag,
release, publication, deployment, or any production mutation.

### Contextual authority

Existing project dependencies and catalogued capabilities are autonomous. A new
uncatalogued direct dependency stops before installation for provenance,
maintenance, security, license, and build-versus-buy evidence plus explicit user choice.
Verified local, disposable, or test mutations may proceed when
rollback or recreation exists. Ambiguous targets remain protected. Production
mutation, deploy, release, publication, and policy changes always require explicit user authorization.
After verifying a non-production target, prefix guarded Terraform, `gcloud`, or
`psql` mutations with `GEREMMYAS_TARGET=local`, `test`, or `disposable`; without
that evidence marker the hook asks for authority. Privileged `sudo` mutations
also require explicit user authorization.
Out-of-scope destructive commands are denied or
replaced with a safe alternative instead of becoming routine prompts.

## Delegation

Create runtime subagents only when specialization, independent context, or
parallelism materially helps. Keep trivial work inline. Give every editor
explicit file, module, or worktree ownership; parallel edits require disjoint
ownership and preserve unrelated work. The primary agent owns integration and
Git. Subagents never stage, commit, rewrite history, push, deploy, or publish.

Describe the bounded task rather than selecting a permanent persona: objective,
scope, ownership, evidence, unknowns, output, and authority. For independent
review, load `verify/references/review-contract.md` after fresh verification and
create a read-only runtime subagent when supported. Repair findings and re-review
automatically. Escalate the same blocker only after three consecutive cycles
with accumulated evidence. On targets without runtime subagents, work inline
with the same bounded contract and report that context isolation was unavailable.
