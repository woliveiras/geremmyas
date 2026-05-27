---
name: rust-release
description: "Plan and execute Rust release engineering. Use when: publishing a crate, setting up release automation, cross-compiling binaries, configuring cargo-dist, or managing changelogs and SemVer."
---

# Rust Release Engineering

Workflow for versioning, publishing, and distributing Rust crates and binaries.

## When to Use

- Publishing a library crate to crates.io
- Setting up automated release pipelines
- Distributing compiled binaries (CLI tools, services)
- Managing changelogs and SemVer bumps
- Cross-compiling for multiple targets

## Library Crate Release

### Tools

| Tool | Purpose |
|------|---------|
| `release-plz` | Automated PR with version bump, changelog, and publish |
| `cargo-release` | Manual release with validation, tagging, and push |
| `cargo-semver-checks` | Detect accidental breaking changes |

### Process

1. Ensure `Cargo.toml` has complete metadata: `description`, `license`, `repository`, `keywords`, `categories`, `rust-version`.
2. Run `cargo semver-checks` to detect unintentional breaking changes.
3. Update version following SemVer:
   - MAJOR: breaking public API changes
   - MINOR: new public functionality, backward-compatible
   - PATCH: bug fixes only
4. Update `CHANGELOG.md` (or let `release-plz` generate it).
5. Run full CI (fmt, clippy, test, doc).
6. Tag the release: `vX.Y.Z`.
7. Publish: `cargo publish --dry-run` then `cargo publish`.

### `release-plz` Automation

```yaml
# .github/workflows/release.yml
name: release

on:
  push:
    branches: [main]

jobs:
  release-plz:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      pull-requests: write
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: dtolnay/rust-toolchain@stable
      - uses: MarcoIeni/release-plz-action@v0.5
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          CARGO_REGISTRY_TOKEN: ${{ secrets.CARGO_REGISTRY_TOKEN }}
```

### `cargo yank`

- Removes a version from the index for new resolves
- Does NOT delete the published artifact
- Use for broken versions, not for "un-publishing"

## Binary Distribution

### Tools

| Tool | Purpose |
|------|---------|
| `cargo-dist` | Build, package, and distribute compiled binaries |
| `cross` | Cross-compilation and cross-testing without local setup |
| `cargo-chef` | Docker layer caching for faster container builds |

### Process

1. Define target platforms in `Cargo.toml`:
   ```toml
   # dist.toml or Cargo.toml [workspace.metadata.dist]
   [dist]
   cargo-dist-version = "0.27.0"
   targets = ["x86_64-unknown-linux-gnu", "aarch64-apple-darwin", "x86_64-pc-windows-msvc"]
   installers = ["shell", "powershell", "homebrew"]
   ```
2. Initialize: `cargo dist init`
3. Build locally: `cargo dist build`
4. CI generates release artifacts on tag push.

### Cross-Compilation

Native:
```bash
rustup target add aarch64-unknown-linux-gnu
cargo build --target aarch64-unknown-linux-gnu
```

With `cross` (zero-setup, uses Docker):
```bash
cross build --target aarch64-unknown-linux-gnu --release
cross test --target aarch64-unknown-linux-gnu
```

### Docker Builds with `cargo-chef`

```dockerfile
FROM rust:1.85 AS chef
RUN cargo install cargo-chef
WORKDIR /app

FROM chef AS planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

FROM chef AS builder
COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json
COPY . .
RUN cargo build --release

FROM debian:bookworm-slim
COPY --from=builder /app/target/release/myapp /usr/local/bin/
ENTRYPOINT ["myapp"]
```

## Cargo.toml Metadata Checklist

```toml
[package]
name = "my-crate"
version = "0.1.0"
edition = "2024"
rust-version = "1.85"
description = "One-line description"
license = "MIT OR Apache-2.0"
repository = "https://github.com/user/repo"
keywords = ["keyword1", "keyword2"]
categories = ["category"]
readme = "README.md"

[package.metadata.docs.rs]
all-features = true
```

## Workspace Release

For workspaces with multiple crates:
- `release-plz` handles independent versioning per crate
- Use `[workspace.dependencies]` for shared dependency versions
- Publish in dependency order (leaf crates first)
- Consider `cargo-workspaces` for coordinated releases

## Pre-Release Validation

Before any release:
1. `cargo fmt --all -- --check`
2. `cargo clippy --workspace --all-targets --all-features -- -D warnings`
3. `cargo test --workspace --all-features`
4. `cargo doc --workspace --all-features --no-deps`
5. `cargo semver-checks` (for libraries)
6. `cargo publish --dry-run`
7. Verify MSRV: `cargo hack check --rust-version --workspace`
