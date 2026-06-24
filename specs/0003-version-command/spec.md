---
spec: "0003"
title: Geremmyas version command
family: multi-assistant
phase: 1
status: Implemented
owner: ""
depends_on: []
origin: direct user request
---

# Spec: Geremmyas version command

## Context & Motivation

Users need a reliable way to check which Geremmyas binary is installed. The CLI
currently exposes operational commands (`init`, `sync`, `lint`, `doctor`) but no
version command or flag. Release builds already come from release-please tags,
so the binary can report the installed release tag when CI injects it at build
time, while local builds can clearly identify themselves as development builds.

## Requirements

### Functional

- [ ] Add `geremmyas version`.
- [ ] Add `geremmyas --version`.
- [ ] Print exactly `geremmyas <version>` followed by a newline.
- [ ] Default local-build version to `dev`.
- [ ] Release builds print the release tag injected by CI, e.g.
      `geremmyas v3.1.0`.
- [ ] Version paths must run before catalog loading so they work even if
      embedded catalog validation would fail.
- [ ] Include version usage in CLI help and README command documentation.

### Non-Functional

- [ ] No new third-party dependencies.
- [ ] Keep version source simple and release-compatible with existing GitHub
      Actions release workflow.
- [ ] Preserve existing command behavior and exit codes.

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | Isolated CLI dispatch/output | `Run([]string{"version"})`, `Run([]string{"--version"})` |
| **integration** | Build-time ldflags and binary smoke | `go build -ldflags="-X ...=v9.9.9"` |

Default: **unit** for dispatch behavior, plus one manual integration smoke for
ldflags injection.

## Acceptance Criteria

- [ ] Given a local build, when `geremmyas version` runs, then it exits 0 and
      prints `geremmyas dev`.
- [ ] Given a local build, when `geremmyas --version` runs, then it exits 0 and
      prints `geremmyas dev`.
- [ ] Given a release build with version `v9.9.9` injected via ldflags, when
      `geremmyas version` runs, then it exits 0 and prints `geremmyas v9.9.9`.
- [ ] Given catalog loading would fail, when `geremmyas version` or
      `geremmyas --version` runs, then version output still succeeds.
- [ ] Given `geremmyas --help` runs, when usage is printed, then version usage
      is listed.

## Edge Cases

- Empty injected version is not expected; the default value remains `dev`.
- `-v` is not added.
- Release tags already include the `v` prefix; the command does not add or
  normalize prefixes.

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| CLI UX | Support both `version` and `--version` | Covers discoverable command style and common CLI flag style |
| Version source | Mutable package variable defaulting to `dev`, overridden by ldflags | Simple Go release pattern; no runtime network or file reads |
| Dispatch order | Handle version before catalog loading | Version remains useful for diagnosing broken installs |
| Short flag | Do not add `-v` | Keeps `-v` available for possible future verbose mode |

## Out of Scope

- Online update checks.
- Comparing installed version against latest GitHub release.
- Injecting versions into local checkout installs.
- Changing release-please versioning policy.
