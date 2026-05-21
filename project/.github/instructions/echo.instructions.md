---
description: "Use when writing or reviewing Echo applications. Covers handlers, middleware, context handling, centralized errors, and graceful shutdown."
applyTo: "**/handlers/**/*.go, **/routes/**/*.go, **/middleware/**/*.go, **/server/**/*.go, **/cmd/**/*.go, **/*handler*.go, **/*route*.go"
---

# Echo Conventions

- Keep handlers thin: bind input, validate, authorize, call application code,
  and return a response.
- Return errors from handlers and middleware so Echo's centralized HTTP error
  handler can shape the response and logging.
- Use `echo.NewHTTPError` for expected client-facing HTTP errors.
- Do not access `echo.Context` from goroutines after the handler returns; copy
  the request values needed by background work.
- Register custom context middleware before middleware that depends on it.
- Put shared behavior in middleware, not repeated handler code.
- Use request-scoped `context.Context` for downstream database, HTTP, and queue
  calls.
- Start servers with graceful shutdown and explicit HTTP server timeouts in
  production-facing applications.
