# Plan: Geremmyas version command

Spec: [spec.md](./spec.md)

Status: Implemented

## Approach

Add a package-level CLI version variable with default `dev`, route `version` and
`--version` before catalog loading, and inject the release tag in the existing
release workflow through Go ldflags.

## Touch points

- `internal/cli/cli.go` — add version variable/helper, early dispatch, switch
  case, and help usage line.
- `.github/workflows/release.yml` — extend current `-ldflags="-s -w"` with
  `-X github.com/woliveiras/geremmyas/internal/cli.Version=${{ needs.release-please.outputs.tag_name }}`.
- `README.md` — document `geremmyas version` in quick usage and command table.

## Sequencing

1. Add unit tests for `version` and `--version` output and exit code.
2. Add CLI version variable/helper and early dispatch before catalog loading.
3. Update help and README command docs.
4. Update release workflow ldflags.
5. Run focused tests, full tests, and manual ldflags build smoke.

## Dependencies

- No external dependencies.
- Uses existing release-please `tag_name` output from `.github/workflows/release.yml`.

## Risks

- Wrong ldflags import path would silently leave release binaries reporting
  `dev`; manual smoke with `v9.9.9` covers this.
- Handling `--version` after catalog loading would reduce diagnostic value; keep
  early dispatch explicit.

## Verification

- `go test ./internal/cli`
- `go test ./...`
- `go build -ldflags="-X github.com/woliveiras/geremmyas/internal/cli.Version=v9.9.9" -o /tmp/geremmyas ./cmd/geremmyas`
- `/tmp/geremmyas version`
