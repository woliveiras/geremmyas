---
description: "Use when writing or reviewing Go SQLite code with database/sql, especially modernc.org/sqlite."
applyTo: "**/database/**/*.go, **/db/**/*.go, **/storage/**/*.go, **/persistence/**/*.go, **/*sqlite*.go, **/*database*.go, **/*repository*.go"
---

# Go SQLite Conventions

- Prefer `modernc.org/sqlite` when the project needs a pure Go SQLite driver
  without CGO.
- Keep the blank driver import in one infrastructure package; do not scatter it
  through handlers or domain code.
- Open SQLite through `database/sql` with driver name `sqlite` when using
  `modernc.org/sqlite`.
- If `modernc.org/libc` is pinned directly, keep it compatible with the
  `modernc.org/sqlite` version in `go.mod`.
- Use DSN pragmas for connection-wide SQLite behavior:
  `file:app.db?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=foreign_keys(ON)&_pragma=synchronous(NORMAL)`.
- Set `db.SetMaxOpenConns(1)` unless the project has a tested write-concurrency
  strategy.
- Use `context.Context` on queries and exec calls that serve requests or jobs.
- Use transactions for multi-statement writes and migrations.
- In tests, use an isolated temp file database or
  `file::memory:?cache=shared&_pragma=foreign_keys(ON)`.
