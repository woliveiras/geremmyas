# Tasks: Geremmyas version command

Spec: [spec.md](./spec.md) · Plan: [plan.md](./plan.md)

## Tasks

- [x] **Version command dispatch** (test-type: unit)
  - blocked-by: none
  - summary: Add tests and implementation for `geremmyas version` and
    `geremmyas --version`.
  - desired behavior: Both forms exit 0 and print `geremmyas dev` in local
    builds.
  - acceptance criteria: Unit tests cover output, exit code, and early dispatch
    before catalog-dependent commands.
  - verification: `go test ./internal/cli`

- [x] **Help and README documentation** (test-type: unit)
  - blocked-by: Version command dispatch
  - summary: Add version usage to CLI help and README command docs.
  - desired behavior: Users can discover version support from help and docs.
  - acceptance criteria: Help output lists `geremmyas version`; README quick
    usage and command table include version.
  - verification: `go test ./internal/cli`

- [x] **Release version injection** (test-type: integration)
  - blocked-by: Version command dispatch
  - summary: Update release workflow ldflags to inject release-please tag name.
  - desired behavior: Release binaries print their tag instead of `dev`.
  - acceptance criteria: Workflow build command sets
    `github.com/woliveiras/geremmyas/internal/cli.Version` from release tag.
  - verification: `go build -ldflags="-X github.com/woliveiras/geremmyas/internal/cli.Version=v9.9.9" -o /tmp/geremmyas ./cmd/geremmyas && /tmp/geremmyas version`
