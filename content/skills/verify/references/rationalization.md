# Rationalization checks

Load this only when a completion gate is being weakened or skipped.

| Claim | Required response |
| --- | --- |
| "This is simple" or "tests are overkill" | Identify the observable behavior and run a proportionate focused check. |
| "It should work" or "I know this pattern" | Run fresh evidence in this repository and environment. |
| "The build passes" | Exercise the behavior or state why runtime proof is unavailable. |
| "CI passed earlier" | Re-run locally or record the exact current external evidence and limitation. |
| "Review will catch it" | Keep review and execution as separate evidence. |
| "No one will hit the edge case" | Compare it with acceptance criteria and impact; test it when in scope. |

If the check is genuinely unavailable, record reason, owner, residual risk, and
the authority or environment needed. Do not convert unavailable evidence into a
passing result.
