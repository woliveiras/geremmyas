# Bugfix: CI structure check requires removed agents directory

- Date: 2026-08-03
- Area: GitHub Actions structure validation
- Severity: High (blocks release confidence on `main`)
- Status: Implemented

## Symptom

The `CI` workflow fails in `Structure Check` even though the skills lint,
installer tests, Geremmyas workflow, and local Go suite pass.

## Reproduction

On commit `08abff55899b22acca73d25aa0bf48629b0ab2b1`, GitHub Actions run
`30840230631` executes the expected-file checks with `bash -e` and exits at:

```sh
test -d content/agents
```

The directory does not exist after spec 0009 intentionally removed all bundled
custom-agent profiles.

## Root cause

The lazy workflow harness removed `content/agents/` while retaining generic
target-adapter support for possible future agents. The CI structure inventory
was not reconciled with that architectural change, so it continued treating a
removed canonical directory as required repository structure.

## Hypotheses considered

1. The directory was accidentally omitted from the checkout. Rejected: it is
   absent from `HEAD` by design and the catalog distributes no bundled agents.
2. A generated agent destination should satisfy the check. Rejected: the check
   names the canonical `content/agents/` source, not a materialized target path.
3. The workflow retained a stale structural invariant. Confirmed by the failing
   command and spec 0009's verified no-bundled-agent acceptance criteria.

## Fix

Remove the obsolete `test -d content/agents` assertion. Keep the checks for
canonical prompts, instructions, skills, catalog, CLI, and hook sources.

## Regression test

`TestCIWorkflowDoesNotRequireRemovedBundledAgentsDirectory` reads the checked-in
CI workflow and fails if the obsolete directory assertion is reintroduced.
Existing catalog and materialization tests continue to assert that base packs
distribute no bundled agents while generic adapter support remains available.

## Verification results

- Focused workflow and catalog regressions pass.
- The full Go suite passes with a fresh, uncached run.
- Skills lint passes for all 43 skills.
- Doctor validates the catalog; the checkout intentionally has no local
  `geremmyas.yml`.
- `git diff --check` passes.
- `actionlint` was not available locally. The remote GitHub Actions `CI`
  workflow remains the final platform validation after this commit is pushed.
