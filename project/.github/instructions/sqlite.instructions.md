---
description: "SQLite conventions for schemas, migrations, pragmas, transactions, and query design across languages."
applyTo: "**/*.sql, **/database/**/*, **/databases/**/*, **/db/**/*, **/storage/**/*, **/persistence/**/*, **/migrations/**, **/migrations/**/*, **/*sqlite*.*, **/*database*.*, **/*repository*.*"
---

# SQLite Conventions

- Enable foreign keys explicitly for every application connection.
- Use WAL mode for application databases that need concurrent reads while
  writes happen.
- Configure a busy timeout so lock contention waits briefly instead of failing
  immediately.
- Keep migrations append-only after they have been applied outside local
  development.
- Use transactions for multi-statement writes, migrations, and state changes
  that must be atomic.
- Never concatenate untrusted values into SQL. Use bound parameters from the
  current language or ORM.
- Primary keys: `INTEGER PRIMARY KEY AUTOINCREMENT`
- Timestamps: `TEXT` in ISO 8601
- Booleans: `INTEGER` (0/1) with `NOT NULL DEFAULT 0`
- Define indexes for columns used in `WHERE`, `JOIN`, and `ORDER BY`.
- Choose `ON DELETE CASCADE`, `ON DELETE SET NULL`, or `RESTRICT`
  deliberately; do not leave relationship behavior implicit.
- Back up WAL/SHM files with the main database file when WAL mode is enabled.
