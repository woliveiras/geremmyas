# Bugfix document template

Load this template only when creating a new bugfix document. Preserve an
existing repository format when updating an established artifact.

```markdown
# Bugfix: <user-visible symptom>

**Status:** investigating | fixed | won't fix
**Date opened:** YYYY-MM-DD
**Source:** issue/PR/user report/log link, or "local report"
**Outage:** yes | no

## Summary
<broken and expected behavior>

## Impact
- affected users/systems, scope, frequency, and consequence

## Reproduction
**Environment:** <versions, flags, configuration>
**Steps:** <minimal deterministic sequence>
**Expected:** ...
**Actual:** ...
**Command or loop:** `<command>`

## Hypotheses
| Rank | Hypothesis | Prediction | Result |
| --- | --- | --- | --- |
| 1 | ... | If true, then ... | pending |

## Investigation log
- YYYY-MM-DD HH:MM: observation, command, or evidence

## Regression test
- File and test name:
- Red command, expected reason, and observed output:
- Green command and observed output:

## Fix
- Root cause:
- Smallest change:
- Why it fixes the cause:

## Verification
- [ ] Regression test failed before the fix for the expected reason
- [ ] Regression test passes
- [ ] Original reproduction no longer reproduces
- [ ] Related suite passes
- [ ] Temporary instrumentation removed

## Residual evidence
- unavailable check, reason, owner, remaining risk, and required authority

## Follow-ups
- [ ] Follow-up task, if needed
```
