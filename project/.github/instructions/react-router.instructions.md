---
description: "Use when writing or reviewing React Router v7 routes in framework mode. Covers loaders, actions, typegen, and breaking changes from v6."
applyTo: "**/routes/**/*.tsx, **/routes/**/*.ts, **/app/routes/**/*.tsx"
---

# React Router v7 Conventions

- Fetch route data in `loader` and perform route mutations in `action`; keep
  components focused on rendering and interaction.
- Use generated `Route.*` types from the route's `+types` file. Do not manually
  duplicate loader or action data types.
- Import routing APIs from `react-router`, not `react-router-dom`.
- Follow framework-mode file routing conventions: `.` for URL segments, `$` for
  dynamic params, `_index` for index routes, and `_` prefix for pathless
  layouts.
- Use `<Form>` for standard mutations and `useSubmit` or `useFetcher` for
  programmatic or background mutations.
- Export route-level `ErrorBoundary` for expected loader/action failures.
- Prefer `<Link>` or `<NavLink>` for navigation; use `useNavigate()` only when
  navigation is a side effect of logic.
- For React Router v6 to v7 migrations, use the `migrate-react-router` skill.
