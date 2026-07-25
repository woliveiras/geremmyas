---
description: "Use when writing or reviewing Go code that handles input, crypto, templates, file paths, or crosses trust boundaries."
applyTo: "**/*.go"
---

# Go Security

## Input Validation (HTTP)

- Limit request body: `r.Body = http.MaxBytesReader(w, r.Body, maxBytes)`
- Use `json.NewDecoder` with `DisallowUnknownFields()` for strict parsing
- Validate semantic fields after decoding (not just struct tags)
- Check for exactly one JSON value; reject trailing content
- Never use `io.ReadAll` on unbounded input from network or user

## JSON Parsing Pitfalls

- Duplicate keys silently overwrite/merge; validate if payload is from untrusted source
- Field matching is case-insensitive; unknown fields are silently ignored by default
- Always call `Decoder.DisallowUnknownFields()` for API inputs

## Templates

- Use `html/template` for HTML output; it does contextual autoescaping
- Never use `text/template` for user-facing HTML (no autoescape, enables XSS)
- Only use `template.HTML`, `template.JS` etc. for values proven safe
- Template sources from untrusted input are code injection vectors

## File Paths & Traversal

- Use `os.Root` (Go 1.24+) for traversal-resistant file access
- Never construct file paths from user input with `filepath.Join` alone
- Validate that resolved paths stay within the intended directory
- Critical for upload/download, extraction, and multi-tenant file systems

## Cryptography & Secrets

- Use `crypto/rand.Reader` for all security-sensitive random values
- Never use `math/rand` or `math/rand/v2` for keys, tokens, or salts
- Go defaults to TLS 1.3; do not weaken TLS config without objective reason
- Do not hardcode secrets in source; use environment, secret stores, or CI secrets
- Do not log secrets, tokens, PII, or credentials

## Race Detection & Memory Safety

- Pure Go is memory-safe; `cgo` and `unsafe` break that guarantee
- `cgo` allocations are invisible to GC; free manually with `C.free`
- Treat `unsafe` and `cgo` as audit boundaries, not routine code
- Run `go test -race ./...` in CI; data races are undefined behavior
- Use `-asan` build flag (Go 1.25+) for detecting C memory leaks

## Supply Chain & Dependencies

- `go.sum` authenticates downloads via the checksum database; commit it always
- Run `govulncheck ./...` in CI to detect reachable vulnerabilities
- Keep Go and dependencies updated to latest minor/patch
- Do not rewrite or delete published module tags
- Prefer stdlib; add external dependencies only when they add clear value
- Use `GONOSUMCHECK` and `GONOSUMDB` only for genuinely private modules

## Executable Lookup

- Go hardened PATH lookup to avoid running executables from current directory
- Be explicit about paths when invoking external commands with `os/exec`
