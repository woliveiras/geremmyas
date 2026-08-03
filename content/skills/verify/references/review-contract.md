# Independent Review Contract

Load this contract only after fresh verification, when changed behavior or risk
justifies an independent review.

Prefer a runtime subagent created with isolated context. If the runtime has no
subagents, run the same bounded review inline and report that independence was
not available. A permanent custom agent is not required.

## Delegation input

Provide only:

- objective and acceptance criteria;
- changed files or exact diff range;
- relevant spec, plan, bugfix, or documentation artifact;
- fresh verification commands and results;
- known unknowns and explicit review focus;
- read-only ownership unless the primary explicitly assigns disjoint edits.

Do not resend the entire repository or full conversation.

## Reviewer contract

1. Inspect the supplied diff and the nearest implementation, tests, and
   contracts needed to validate it.
2. Prioritize correctness, regressions, security, data loss, concurrency,
   compatibility, and missing tests over style.
3. Verify claims against repository evidence. Do not accept confidence or stale
   output as proof.
4. Return findings ordered by severity with file, tight line range, impact, and
   a concrete remediation direction.
5. Separate actionable findings from residual risks and questions.
6. Return `no findings` explicitly when appropriate, with the checks performed.
7. Never stage, commit, push, merge, publish, deploy, or expand scope.

## Output schema

```text
state: findings | no-findings | blocked
findings:
  - severity: critical | high | medium | low
    location: path:line
    claim: concise defect statement
    evidence: repository or execution evidence
    impact: user or system consequence
unknowns:
  - unavailable evidence or residual risk
checks:
  - command or read-only validation and result
```

The primary agent owns integration. It repairs actionable findings, reruns fresh
verification, and requests another independent review. Escalate only when the
same blocker survives three consecutive cycles with evidence or requires new authority.
