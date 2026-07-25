---
description: "Go code quality based on Effective Go, Uber Go Style Guide, and official Go wiki."
applyTo: "**/*.go"
---

# Go Code Quality

Follow [Effective Go](https://go.dev/doc/effective_go),
[Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md),
and the [Go Code Review Comments wiki](https://go.dev/wiki/CodeReviewComments).

## Package & Module Organization

- Use `internal/` for code not exported to consumers; default to private
- Use `cmd/` for executable entry points; keep `main.go` thin
- Do not use `/pkg` as a general rule; normal packages + `internal/` suffice
- Avoid generic package names: `util`, `common`, `misc`, `api`, `types`, `interfaces`
- One module per repository unless you have a strong reason for multi-module
- Use `go` directive as minimum compatibility; `toolchain` as development preference
- Commit `go.sum`; it is part of the supply chain contract

## Naming

- Package names: lowercase, single word, no plural (`book` not `books`)
- No stutter: `chubby.File` not `chubby.ChubbyFile` (package already qualifies)
- No `Get` prefix on getters: `book.Title()` not `book.GetTitle()`
- Initialisms in consistent casing: `URL`, `HTTP`, `ID` (not `Url`, `Http`, `Id`)
- Short, consistent receiver names (1-2 letters matching the type)
- Error variables: `ErrNotFound` (exported), `errNotFound` (unexported)
- Error types: suffix `Error` (`NotFoundError`)
- Error strings: lowercase, no trailing punctuation (they compose in chains)

## Error Handling

- Errors are values; handle, propagate, or translate explicitly
- Wrap with context: `fmt.Errorf("find book %d: %w", id, err)`
- Handle errors once: either log OR return, never both
- Use `%w` for errors callers match with `errors.Is`/`errors.As`; `%v` to obfuscate
- Treat error path first; keep the happy path at the lowest indentation
- Never discard errors with `_` unless explicitly justified
- Sentinel errors for `errors.Is`; custom types for `errors.As`
- Do not use `panic` for normal error control; reserve for broken invariants

## Context

- `context.Context` is always the first parameter, named `ctx`
- Never store `context.Context` in a struct
- Never pass `nil` as context; use `context.TODO()` if unsure
- Propagate context to all I/O, database, HTTP, and queue calls
- Use context for cancellation, deadlines, and tracing metadata

## Interfaces

- Define interfaces at the consumer side, not the implementation
- Return concrete types; accept interfaces only for the surface you need
- Keep interfaces small (1-3 methods); compose for larger contracts
- Do not create interfaces "for mocking" on the producer side
- Verify compliance: `var _ http.Handler = (*MyHandler)(nil)`

## Concurrency

- Every goroutine must have a clear lifetime and shutdown mechanism
- Use `errgroup.WithContext` for parallel subtasks with shared cancellation
- Use `errgroup.SetLimit(n)` for bounded concurrency
- Never fire-and-forget goroutines; they leak if channels block
- Channel size: 0 or 1. Larger buffers need profiling justification
- Prefer passing data by channels or copy over shared memory
- Use `sync.Mutex`/`sync.RWMutex` only when message passing is impractical
- Run tests with `-race`; the race detector must be part of CI

## Functions

- `run()` pattern: `func main() { if err := run(); err != nil { log.Fatal(err) } }`
- `os.Exit` / `log.Fatal` only in `main()`
- Prefer small, focused functions with a single clear responsibility
- Avoid deep nesting; return early on errors

## Testing

- Use stdlib `testing` as base; add `testify` only for assertion productivity
- Table-driven tests when multiple cases exercise the same logic
- `t.Parallel()` where safe
- Don't mock what you don't own; wrap behind interfaces at the consumer
- Fuzzing (`testing.F`) for parsers, validators, and input handling
- Run `go test -race ./...` in CI on every PR
- Use Examples as executable documentation

## Observability

- Use `log/slog` for structured logging (stdlib)
- Use `tracing` with OpenTelemetry for distributed traces
- Expose metrics via `prometheus/client_golang`
- Never log secrets, tokens, PII, or full request bodies

## Linting

- `gofmt` (or `gofumpt`) is non-negotiable
- `go vet` is the minimum static check (built-in, conservative)
- `staticcheck` recommended for deeper analysis with low false positives
- `golangci-lint` optional as orchestrator; curate enabled linters
- In CI: `go vet ./...` + `staticcheck ./...` minimum

## Anti-patterns

- `init()` functions (except driver registration)
- Mutable global variables
- Named returns (except short functions or `defer` error handling)
- `panic` in library code
- Storing `context.Context` in structs
- Interfaces on the implementation side "for mocking"
- `io.ReadAll` on unbounded input without size limits
- Using `math/rand` for security-sensitive values
- Ignoring `go.sum` changes or deleting it
