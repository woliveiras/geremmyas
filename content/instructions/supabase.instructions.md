---
description: "Use when writing or reviewing Supabase clients, migrations, RLS policies, Auth, Storage, Edge Functions, generated types, and platform configuration."
applyTo: "supabase/**/*, **/supabase/**/*, **/*supabase*.*, **/database/migrations/**/*.sql, **/migrations/**/*.sql, **/functions/**/*.ts"
---

# Supabase Conventions

- Enable Row Level Security for tables exposed through Supabase APIs and write
  policies before frontend access depends on them.
- Pair client-side publishable or anon keys with least-privilege grants and RLS;
  never treat them as secrets.
- Keep secret keys and service role keys on trusted backends only. They bypass
  RLS and must never be exposed to browser or mobile clients.
- Write both `USING` and `WITH CHECK` policy clauses when reads and writes have
  different rules.
- Index columns used by RLS policies, joins, filters, and ownership checks.
- Put schema, RLS, storage policy, and function changes in migrations or
  checked-in Supabase config, not dashboard-only state.
- Regenerate and commit typed clients when database schema changes affect
  application code.
- Keep Auth user IDs, tenant IDs, and storage object paths aligned with policy
  predicates.
- In Edge Functions, validate caller identity and inputs before using elevated
  keys or direct database credentials.
