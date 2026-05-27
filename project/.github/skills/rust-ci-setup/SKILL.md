---
name: rust-ci-setup
description: "Set up a CI/CD pipeline for Rust projects. Use when: creating CI for Rust, adding Clippy/fmt/test/audit to CI, configuring GitHub Actions for Rust, or adding Miri/fuzzing jobs."
---

# Rust CI Setup

Configure a complete CI pipeline for a Rust project.

## When to Use

- New Rust project needs CI from scratch
- Adding quality gates (fmt, clippy, test, audit) to an existing project
- Setting up advanced validation (Miri, sanitizers, fuzzing)
- Migrating CI to a new provider

## Pipeline Layers

### Layer 1: Fast Feedback (every PR)

| Step | Command | Purpose |
|------|---------|---------|
| Format | `cargo fmt --all -- --check` | Style consistency |
| Lint | `cargo clippy --workspace --all-targets --all-features -- -D warnings` | Code quality |
| Test | `cargo nextest run --workspace --all-features` (or `cargo test`) | Correctness |
| Doc | `cargo doc --workspace --all-features --no-deps` | Documentation builds |

### Layer 2: Security & Dependencies (every PR or scheduled)

| Step | Command | Purpose |
|------|---------|---------|
| Audit | `cargo audit` | Known vulnerabilities |
| Deny | `cargo deny check` | Licenses, bans, duplicates, advisories |
| MSRV | `cargo hack check --rust-version --workspace` | Minimum supported version |

### Layer 3: Deep Validation (nightly/scheduled)

| Step | Command | Purpose |
|------|---------|---------|
| Miri | `cargo +nightly miri test` | Undefined behavior detection |
| Sanitizers | `RUSTFLAGS="-Zsanitizer=address"` | Memory errors |
| Fuzzing | `cargo fuzz run <target> -- -max_total_time=300` | Input space exploration |
| Coverage | `cargo llvm-cov --workspace --all-features --lcov --output-path lcov.info` | Code coverage |

## GitHub Actions Baseline

```yaml
name: ci

on:
  push:
    branches: [main]
  pull_request:

env:
  CARGO_TERM_COLOR: always
  RUSTFLAGS: "-Dwarnings"

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install Rust toolchain
        uses: dtolnay/rust-toolchain@stable
        with:
          components: rustfmt, clippy

      - name: Cache Cargo
        uses: actions/cache@v4
        with:
          path: |
            ~/.cargo/registry
            ~/.cargo/git
            target
          key: ${{ runner.os }}-cargo-${{ hashFiles('**/Cargo.lock') }}
          restore-keys: ${{ runner.os }}-cargo-

      - name: Format
        run: cargo fmt --all -- --check

      - name: Clippy
        run: cargo clippy --workspace --all-targets --all-features

      - name: Test
        run: cargo test --workspace --all-features

  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: dtolnay/rust-toolchain@stable
      - name: Install cargo-deny
        uses: taiki-e/install-action@cargo-deny
      - name: Install cargo-audit
        uses: taiki-e/install-action@cargo-audit
      - run: cargo audit
      - run: cargo deny check

  miri:
    runs-on: ubuntu-latest
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@v4
      - uses: dtolnay/rust-toolchain@nightly
        with:
          components: miri
      - run: cargo +nightly miri test --workspace
```

## Configuration Files

### `rustfmt.toml`

```toml
style_edition = "2024"
```

### `clippy.toml`

```toml
msrv = "1.85"
```

### `deny.toml`

```toml
[advisories]
vulnerability = "deny"
unmaintained = "warn"
yanked = "warn"

[licenses]
allow = ["MIT", "Apache-2.0", "BSD-2-Clause", "BSD-3-Clause", "ISC", "Unicode-3.0"]

[bans]
multiple-versions = "warn"
wildcards = "deny"

[sources]
unknown-registry = "deny"
unknown-git = "deny"
```

### `Cargo.toml` metadata

```toml
[package]
edition = "2024"
rust-version = "1.85"

[lints.clippy]
all = "warn"
pedantic = "warn"
```

## Process

1. Determine project type (library, binary, workspace).
2. Start with Layer 1 for every PR.
3. Add Layer 2 for security-conscious projects.
4. Add Layer 3 as scheduled jobs for projects with unsafe code, parsers, or FFI.
5. Configure caching for `~/.cargo/registry`, `~/.cargo/git`, and `target`.
6. Set MSRV in `Cargo.toml` and verify in CI.
7. Commit `Cargo.lock` (recommended as starting point for all projects).

## Tooling Installation

For CI, prefer `taiki-e/install-action` for fast binary installs:
- `cargo-nextest`, `cargo-deny`, `cargo-audit`, `cargo-llvm-cov`, `cargo-hack`

For local development:
```bash
cargo install cargo-nextest cargo-deny cargo-audit cargo-llvm-cov cargo-hack
rustup component add rustfmt clippy
```
