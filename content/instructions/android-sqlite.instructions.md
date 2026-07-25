---
description: "Use when writing or reviewing Android SQLite and Room code. Covers entities, DAOs, migrations, async queries, transactions, and migration tests."
applyTo: "**/database/**/*.kt, **/db/**/*.kt, **/storage/**/*.kt, **/persistence/**/*.kt, **/*Dao.kt, **/*Database.kt, **/*Entity.kt, **/*Migration.kt, **/*Repository.kt"
---

# Android SQLite Conventions

- Prefer Room for app database access; use raw SQLite only for cases Room cannot
  model cleanly.
- Keep entities, DAOs, migrations, and database configuration in the data layer,
  not in UI code.
- Use `suspend` DAO functions for one-shot reads/writes and `Flow` for
  observable queries.
- Do not run database I/O on the main thread.
- Write explicit migrations for schema changes; avoid destructive migrations
  outside disposable development data.
- Use `@Transaction` for multi-query reads or writes that must observe a
  consistent database state.
- Keep repository APIs domain-oriented; do not expose Room entities directly if
  domain or UI models diverge.
- Add Room migration tests for every version step that can reach production
  users.
