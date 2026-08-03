---
name: game-build-and-release
description: "Use when building or verifying Phaser/Vite bundles, Godot exports, web delivery, CI, versioning, compression, or release smoke tests. Do not use for gameplay implementation or to publish, tag, upload, or deploy without authority."
---

# Game Build And Release

Produce a versioned artifact from a clean, reproducible process and verify the artifact itself.

## Establish the release contract

1. Read repository instructions and inspect engine, exact version, package manager, export presets, CI, target platforms, hosting, and secret boundaries.
2. Define artifact names, version source, environment inputs, supported browsers or platforms, and required acceptance.
3. Separate public build configuration from credentials.
4. Record required tool and template versions.

## Load references

- Read [release-pipeline.md](references/release-pipeline.md) for reproducibility, artifacts, CI stages, smoke tests, and evidence.
- For Phaser, also read [phaser.md](references/phaser.md).
- For Godot, also read [godot.md](references/godot.md).
- Use `$game-testing-2d` for behavior verification and `$game-performance-2d` for production budgets.

## Build and verify

1. Run source checks and focused tests.
2. Build from declared inputs with production configuration.
3. Inspect artifact contents, sizes, hashes, and accidental secrets.
4. Serve or install the artifact using the real delivery path.
5. Run a minimal playable smoke route and inspect logs.
6. Verify cache, base paths, asset loading, save behavior, input, audio, resize, and exit as applicable.
7. Keep deployment or publication separate unless explicitly authorized.

## Report

Provide tool versions, commands, artifact paths and hashes, checks run, failures, warnings, target-specific gaps, and manual acceptance still required.

## Guardrails

- Do not treat development mode as release verification.
- Do not commit credentials or generated secret files.
- Do not rename Godot web-export companions independently.
- Do not claim tester installation, store approval, or production deployment without direct evidence.
- Commit verified local release-engineering changes through the repository's
  default atomic-commit workflow. Do not publish, upload, tag, push, release, or
  deploy without explicit authorization.
