---
name: postgres-query-review
description: "Review a PostgreSQL query, migration, index, or query plan. Use when: optimizing SQL, checking migrations, reading EXPLAIN output. Do not use: for general database design, non-Postgres systems."
---


# Postgres Query Review

Review PostgreSQL changes through correctness, plan shape, and operational
risk.

## Process

1. Identify the query or migration, expected data volume, critical path, and
   latency or safety target.
2. Check correctness first: joins, filters, null handling, constraints,
   transactions, and isolation assumptions.
3. Inspect indexes used by `WHERE`, `JOIN`, `ORDER BY`, uniqueness, and policy
   predicates.
4. Request or read `EXPLAIN` / `EXPLAIN ANALYZE` for important queries on
   realistic data.
5. For migrations, review locks, table rewrites, backfills, concurrent index
   creation, and rollback strategy.
6. Check connection pool impact and transaction duration.
7. Propose the smallest change that addresses the measured issue.

## Rules

- Do not add indexes without a query pattern and expected benefit.
- Do not use production `EXPLAIN ANALYZE` on dangerous writes.
- Do not hold transactions open across network calls or user interaction.
- Treat destructive migrations and large table rewrites as human-approved
  operations.

## Output

- Correctness findings
- Plan/index observations
- Migration/lock risk
- Recommended SQL or migration change
- Verification command or metric
