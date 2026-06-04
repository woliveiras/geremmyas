---
spec: "0000"
title: <Feature Name>
family: <family-slug>
phase: 0
status: Draft
owner: ""
depends_on: []
origin: ""
---

# Spec: <Feature Name>

## Context & Motivation

<!-- Why this feature exists and what problem it solves -->

## Requirements

### Functional

- [ ] Requirement 1
- [ ] Requirement 2
- [ ] Requirement 3

### Non-Functional

- [ ] Performance: ...
- [ ] Security: ...
- [ ] Accessibility: ...

## Test Strategy

<!-- How acceptance criteria will be verified. Per task, prefer one primary type. -->

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Isolated logic, pure functions, single module, no I/O | validators, parsers, domain rules |
| **integration** | Cross-module flows, DB/API, multi-step, external boundaries | HTTP handlers + DB, CLI + filesystem |
| **both** | Feature spans layers; unit for core logic, integration for seams | auth token logic + login API |

Default: **unit** unless acceptance criteria require real I/O or multiple modules.

## Acceptance Criteria

<!-- Each criterion must be testable — maps to at least one test (unit or integration) -->

- [ ] Given [context], when [action], then [expected result]
- [ ] Given [context], when [action], then [expected result]
- [ ] Given [context], when [action], then [expected result]

## Edge Cases

- What happens when input is empty?
- What happens when the user is not authenticated?
- What happens when the external service is unavailable?

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| ... | ... | ... |

## Out of Scope

- Things explicitly NOT included in this spec
