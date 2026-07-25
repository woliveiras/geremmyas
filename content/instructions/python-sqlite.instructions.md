---
description: "Use when writing or reviewing Python SQLite code. Covers sqlite3, SQLAlchemy SQLite use, parameter binding, transactions, row factories, and connection lifecycle."
applyTo: "**/database/**/*.py, **/db/**/*.py, **/storage/**/*.py, **/persistence/**/*.py, **/*sqlite*.py, **/*database*.py, **/*repository*.py"
---

# Python SQLite Conventions

- Use DB-API parameter substitution for values; never build SQL with Python
  string formatting or f-strings for untrusted input.
- Keep connection lifecycle explicit: per request, per job, or owned by a
  repository/service object with clear shutdown.
- Use context managers or explicit `commit()`/`rollback()` around writes.
- Set `row_factory` when call sites need named-column access instead of tuple
  indexes.
- Be deliberate with `check_same_thread`; disabling it requires the project to
  serialize access safely.
- In async web apps, avoid blocking the event loop with long SQLite operations;
  move them to a thread, job, or async-compatible persistence boundary.
- Use SQLAlchemy migrations/settings when the project already standardizes on
  SQLAlchemy, but keep SQLite-specific pragmas explicit.
- Test with isolated temp file databases when transactions, WAL, migrations, or
  concurrent access matter.
