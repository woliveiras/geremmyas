# Chain Of Evidence Template

Use this matrix to keep claims traceable from research questions to raw evidence and analysis steps.

| RQ | Finding or claim | Supporting evidence | Source type | Raw evidence reference | Analysis step | Rival explanation | Confidence |
|---|---|---|---|---|---|---|---|
| RQ1 | TBD | TBD | direct | TBD | TBD | TBD | Tentative |
| RQ1 | TBD | TBD | archival | TBD | TBD | TBD | Tentative |
| RQ2 | TBD | TBD | indirect | TBD | TBD | TBD | Tentative |

## Confidence Labels

- Strongly supported: multiple independent evidence sources support the claim and rival explanations are weak.
- Moderately supported: more than one evidence source supports the claim, but limitations remain.
- Tentative: evidence is plausible but narrow, indirect, or weakly triangulated.
- Unsupported: proposed claim lacks sufficient evidence.
- Contradicted: available evidence conflicts with the claim.

## Evidence Source Types

- Direct: interviews, focus groups, observations, think-aloud sessions.
- Indirect: telemetry, screen recordings, IDE traces, agent traces, instrumentation.
- Archival: commits, pull requests, issues, Jira tickets, design docs, ADRs, incidents, CI logs, Sentry issues, review comments.

## Claim Check

Before reporting a claim, verify:

- The raw evidence reference can be found by another researcher.
- The analysis step is explicit.
- At least one rival explanation was considered.
- The confidence label matches the strength of triangulation.
- The claim is scoped to the case and does not overgeneralize.
