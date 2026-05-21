---
description: "Use when writing or reviewing Node.js SQLite code. Covers node:sqlite, better-sqlite3, sqlite3, prepared statements, transactions, and request-path boundaries."
applyTo: "**/database/**/*.ts, **/database/**/*.js, **/db/**/*.ts, **/db/**/*.js, **/storage/**/*.ts, **/storage/**/*.js, **/persistence/**/*.ts, **/persistence/**/*.js, **/*sqlite*.ts, **/*sqlite*.js, **/*database*.ts, **/*database*.js, **/*repository*.ts, **/*repository*.js"
---

# Node SQLite Conventions

- Keep SQLite setup in an infrastructure module; do not open database
  connections inside route handlers.
- Use prepared statements or tagged SQL APIs for values; never interpolate
  untrusted values into SQL strings.
- Wrap multi-statement writes in explicit transactions.
- Keep synchronous drivers such as `node:sqlite` `DatabaseSync` or
  `better-sqlite3` out of latency-sensitive request paths unless the workload
  is small and measured.
- Serialize long-running writes through a queue, job, or worker when they can
  block the event loop.
- Finalize or dispose prepared statements and close database handles during
  application shutdown.
- Run migrations during startup or deployment before serving traffic.
- Test repository behavior with a temp file database when WAL, migrations, or
  file-backed behavior matters.
