---
description: "Use when writing or reviewing PostgreSQL schema, migrations, SQL, repositories, query plans, and connection handling."
applyTo: "**/*.sql, **/migrations/**, **/migrations/**/*, **/database/**/*, **/db/**/*, **/storage/**/*, **/repositories/**/*, **/*postgres*.*, **/*pgsql*.*, **/*migration*.*, **/*repository*.*"
---

# PostgreSQL Conventions

- Put schema changes in migrations; do not rely on application startup to
  mutate production schema implicitly.
- Model invariants with constraints first: primary keys, foreign keys, unique
  constraints, checks, and `NOT NULL`.
- Add indexes for query patterns, not for every column. Verify important ones
  with `EXPLAIN` or `EXPLAIN ANALYZE` on realistic data.
- Use transactions for multi-statement writes and state transitions that must be
  atomic.
- Keep transactions short; do not hold them open across network calls, user
  interaction, or long background work.
- Use connection pools deliberately. Match pool size to the database and
  deployment limits instead of giving every worker a large pool.
- Prefer parameterized queries or query builders over string interpolation.
- Review deletes, cascading foreign keys, locks, long migrations, and
  concurrent index creation before applying schema changes.
- Keep migration rollback strategy explicit when a migration changes data or
  rewrites large tables.
