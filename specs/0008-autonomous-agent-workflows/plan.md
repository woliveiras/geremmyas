# Plan: Autonomous agent workflows

Spec: [spec.md](./spec.md)

Status: In Progress

## Approach

Change the canonical workflow one independently verifiable behavior at a time.
Start with feature lifecycle and status semantics, then make bugfix delivery
autonomous, make atomic commits the default, expand specialist delegation, and
finish with contextual dependency and environment authority. Each slice updates
the public documentation and target-facing content it affects.

The primary agent owns integration and Git. Subagents may investigate, edit, test,
or review bounded areas with explicit ownership. Each completed slice receives an
independent review, fresh verification, and one Conventional Commit. No push or
production operation is part of this plan.

## Sequencing

1. **Done:** Update feature intake, specification, status, TDD, completion,
   prompt, and template contracts so machine readiness replaces human approval.
2. Convert bugfix approval into an autonomous reproduce-red-fix-green-document
   loop while retaining regression and evidence requirements.
3. Make atomic local commits the default, add explicit opt-out behavior, and
   derive staging and messages from task ownership and the verified diff.
4. Expand subagent contracts for proactive specialist implementation, audit,
   review, safe edit partitioning, automatic repair, and bounded escalation.
5. Make dependency and command authority contextual: catalogued/local/test work
   proceeds, new uncatalogued dependencies and production operations stop.
6. Materialize every applicable pack across supported targets, run the full
   verification matrix, reconcile artifacts, and document migration behavior.

## Main Touch Points

- `content/AGENTS.md`, `content/prompts/sdd.prompt.md`, and
  `content/templates/specs/README.md` for the default lifecycle.
- `content/skills/requirements-interview`, `generate-spec`, `vertical-tdd`,
  `bugfix-loop`, `verification-checklists`, `git-commit`, and related references.
- `content/agents/` and `content/agents/references/subagent-selection.md` for
  specialist orchestration and ownership contracts.
- `content/guardrails/guardrails-rules.txt`, infrastructure and release skills,
  and target hook adapters for contextual authority.
- A durable gate-migration inventory and automated full-catalog scan covering all
  skill/agent reference trees, prompts, instructions, templates, adapters, packs,
  README, and workflow documentation.
- `catalog/packs.json`, `internal/cli` generator/lint tests, `README.md`,
  `docs/guardrails-framework.md`, and target-specific documentation.

## Dependencies

- Builds on spec 0006 workflow consolidation and spec 0007 assistant-neutral
  canonical content.
- Uses the existing Go toolchain, catalog, target generators, hooks, and test
  harness. No new third-party dependency is needed.

## Risks

- Removing conversational gates without preserving evidence could reduce
  auditability. Keep artifacts, red/green tests, review output, and commit scope
  mandatory and test them as contracts.
- Broad autonomy wording could be mistaken for production authority. Repeat the
  local/test versus production boundary in the contract, relevant skills, hooks,
  and public docs.
- A partial migration could leave an old gate in a stack or personal pack. Scan
  all canonical workflow surfaces and catalogue every retained or deferred match,
  not only `core` and `sdd`.
- Parallel edits can overwrite user or agent work. Require explicit ownership,
  isolate overlapping work, and centralize Git operations.
- Default commits can capture unrelated changes. Stage task-owned paths or hunks
  only, re-read the cached diff, and verify atomicity before every commit.
- Existing historical specs use older statuses. Preserve them and change only the
  workflow and template for new or actively migrated work.

## Verification

- Focused policy and generator tests after each slice.
- `go test ./internal/cli`
- `go test ./... -count=1`
- `go build -o geremmyas ./cmd/geremmyas`
- `go run ./cmd/geremmyas lint`
- `go run ./cmd/geremmyas doctor`
- `go run ./cmd/geremmyas context`
- ShellCheck for install and generated hook scripts.
- Temporary project and home materialization for Codex, Copilot, Cursor, Claude
  Code, OpenCode, and a mixed target selection.
- `git diff --check`, `git status --short`, and an atomic-history review.

## Completion

Complete when every acceptance criterion has fresh evidence, all target outputs
agree on the authority model, no stale task marker remains, the spec is
`Verified`, and the local history contains one task-owned commit per delivered
slice. Push and release remain pending separate authorization.
