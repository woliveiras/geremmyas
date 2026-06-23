---
name: go-ci-setup
description: "Set up a Go CI/CD pipeline with formatting, vetting, static analysis, tests, vulnerability scanning, and container build. Use when: creating CI for Go, adding checks to workflow. Do not use: for local development setup, non-CI automation."
---


# Go CI Setup

Multi-step workflow to create a production-grade Go CI pipeline.

## When to Use

- New Go project needs CI from scratch
- Existing project has incomplete or fragile CI
- Adding security scanning, vulnerability checks, or container builds

## Pipeline Layers (in order)

Each layer gates the next. If a layer fails, the pipeline stops.

1. **Format** — `gofmt -l .` (or `gofumpt -l .`) must produce no output
2. **Vet** — `go vet ./...` (built-in, zero-config, no false positives)
3. **Static Analysis** — `staticcheck ./...` (recommended) or `golangci-lint run`
4. **Test** — `go test -race -coverprofile=coverage.out ./...`
5. **Vulnerability Scan** — `govulncheck ./...` (reachable vulns only)
6. **Build** — `go build -trimpath -ldflags="-s -w" ./cmd/...`
7. **Container** (optional) — Multi-stage Docker with distroless/static base

## Steps

### 1. Choose CI Platform

Decide between GitHub Actions, GitLab CI, Cloud Build, or other.
Default to GitHub Actions for GitHub-hosted repositories.

### 2. Create Workflow File

For GitHub Actions: `.github/workflows/ci.yml`

Minimum structure:
```yaml
name: CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  ci:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
      - run: gofmt -l . | tee /dev/stderr | (! read)
      - run: go vet ./...
      - run: go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck ./...
      - run: go test -race -coverprofile=coverage.out ./...
      - run: go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...
      - run: go build -trimpath ./cmd/...
```

### 3. Add Caching

Use `actions/setup-go@v5` which caches `~/go/pkg/mod` automatically via
`go-version-file`. No separate cache step needed.

### 4. Add Container Build (if applicable)

Multi-stage Dockerfile:
```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
```

Key decisions:
- `CGO_ENABLED=0` for static binary (no libc dependency)
- `distroless/static` or `scratch` as final base (no shell, minimal CVE surface)
- Copy only the binary; no source in production image
- Use `nonroot` tag for least privilege

### 5. Add Release Automation (optional)

For CLI tools or libraries:
- **GoReleaser** for cross-compilation, checksums, and multi-platform archives
- **release-please** for changelog and version management
- **cosign** for signing container images and binaries

### 6. Verify Pipeline

- Push a branch with intentional failures (bad format, failing test) to confirm gates work
- Confirm coverage report is generated
- Confirm `govulncheck` output is visible in CI logs

## Checklist

- [ ] Format check (gofmt/gofumpt) runs before anything else
- [ ] `go vet ./...` included
- [ ] Static analysis tool configured (staticcheck or golangci-lint)
- [ ] Tests run with `-race` flag
- [ ] `govulncheck` scans for reachable vulnerabilities
- [ ] Go version from `go.mod` (not hardcoded)
- [ ] Caching configured (setup-go handles this)
- [ ] Container build uses multi-stage + distroless (if applicable)
- [ ] No secrets in build args or logs
