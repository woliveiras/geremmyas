---
name: feature-readiness-before-implementation
description: "Automatic readiness and authority boundaries before feature implementation. Use when starting or materially changing a feature."
---


# Readiness Before Implementation

Feature delivery is autonomous by default. Specs remain durable audit evidence,
but readiness is established by executable checks rather than conversational
approval.

## When to Use

- Starting a new feature or significant change
- After writing or materially changing `spec.md`, `plan.md`, or `tasks.md`
- When implementation reveals an in-scope assumption that needs revision
- To distinguish evidence gates from authority boundaries

## When NOT to Use

- During implementation when no artifact or assumption changed
- For bug fixes with clear root cause (use `bugfix`)
- For tiny refactors or obvious syntax fixes
- When the user explicitly limited the session to read-only, plan-only, or
  no-edits work; honor that boundary instead

## Automatic readiness gate

A feature moves from `Draft` to `Ready` only when all checks pass:

- `spec.md`, `plan.md`, and `tasks.md` exist and contain no placeholders.
- Acceptance criteria are testable and map to the test strategy and tasks.
- Contracts, dependencies, assumptions, and error paths are explicit.
- Tasks form verifiable vertical slices with fresh verification commands.
- The requested objective and repository conventions resolve material product
  and technical choices.
- No explicit session override forbids implementation.

## Procedure

1. Build `spec.md`, `plan.md`, and `tasks.md` with status `Draft`.
2. Run the readiness checks and record their evidence in the artifacts.
3. Correct failures and rerun the checks until they pass.
4. Set the spec and index status to `Ready`, then start the first task and set
   the feature to `In Progress`.
5. If an in-scope assumption changes, update the affected artifacts and rerun
   readiness automatically. Continue while the objective remains unchanged.
6. Set the feature to `Verified` only after every acceptance criterion has fresh
   evidence, every finished task is `[x]`, no stale `[~]` remains, required docs
   and plan items are reconciled, and independent review has no actionable
   finding.

## When to escalate

Do not turn ordinary uncertainty into a conversational gate. Escalate before
implementation only when blast radius is high and at least one evidence gap is
present:

1. The plausible production blast radius is high.
2. Rollback is weak, unavailable, or unproven; or
3. the critical harness cannot provide adequate pre-production evidence.

Also ask when unresolved product ambiguity would materially change the stated
objective. Otherwise infer the reversible choice from repository conventions,
record it, and proceed.

## Anti-Patterns

**Evidence shortcuts**
- Marking `Ready` because the design "looks right" without running checks.
- Marking `Verified` from stale, partial, or second-hand evidence.
- Treating a reviewer opinion as a substitute for acceptance tests.

**Authority drift**
- Ignoring a read-only, plan-only, or no-edits override.
- Expanding the objective because an adjacent improvement is convenient.
- Treating autonomous local implementation as production authorization.

---

**Key principle:** keep validation strict and executable. Ask the user only at a
real authority boundary or material objective ambiguity, not as a substitute for
evidence the harness can produce.

## Dependency and environment authority

Existing project dependencies and catalogued capabilities may be used
autonomously. Before adding a new uncatalogued direct dependency, provide a
provenance, maintenance, security, license, and build-versus-buy assessment and
get explicit user choice before installation. Lockfile-selected transitives are
recorded evidence; suspicious provenance or security findings reopen the gate.
A compatible update to an existing direct dependency follows the normal harness
after changelog and compatibility review. Direct dependencies include npm/pnpm
packages, Python packages, Go modules or tools, Rust crates or tools, Gradle
libraries or plugins, GitHub Actions, Terraform providers or modules, CI tools,
and externally operated services.

Mutations may proceed against a verified local, disposable, or test target when
rollback or recreation is available. For guarded Terraform, `gcloud`, or `psql`
mutations, carry that evidence as `GEREMMYAS_TARGET=local`, `test`, or
`disposable`. Treat an ambiguous target as protected
until configuration proves otherwise. Every production mutation, deploy,
release, publication, or policy change requires explicit user authorization.
Dangerous operations outside the objective are denied or replaced with a safe
alternative rather than turned into repeated prompts.
