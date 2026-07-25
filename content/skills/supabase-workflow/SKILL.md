---
name: supabase-workflow
description: "Plan or review a Supabase database/platform change. Use when: changing schema, migrations, RLS policies, Auth, Storage, Edge Functions. Do not use: for general database design, non-Supabase platforms."
---


# Supabase Workflow

Change Supabase with schema, RLS, client access, and generated types kept in
sync.

## Process

1. Identify affected tables, policies, functions, buckets, clients, and user
   roles.
2. Decide whether the change belongs in a migration, checked-in config, Edge
   Function, or application code.
3. For exposed tables, design RLS policies before depending on frontend access.
4. Add indexes for policy predicates, joins, filters, and ownership checks.
5. Keep service role keys on trusted backends only.
6. Regenerate typed clients when schema changes affect application code.
7. Test anon/authenticated/service-role paths separately.
8. Document rollback or follow-up data backfill when migrations are not easily
   reversible.

## Rules

- Do not use dashboard-only changes as the source of truth.
- Do not expose service role keys to browser or mobile code.
- Do not treat anon or publishable keys as secrets; rely on grants and RLS.
- Write both `USING` and `WITH CHECK` when read and write rules differ.

## Output

- Supabase surface affected
- Migration/config/code changes
- RLS and key-safety summary
- Generated type impact
- Verification checklist
