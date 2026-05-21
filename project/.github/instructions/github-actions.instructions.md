---
description: "Use when writing or reviewing GitHub Actions workflows, reusable workflows, composite actions, and CI/CD automation."
applyTo: ".github/workflows/**/*.yml, .github/workflows/**/*.yaml, **/action.yml, **/action.yaml"
---

# GitHub Actions Conventions

- Set default `permissions` to read-only and grant write scopes only at the job
  that needs them.
- Prefer OpenID Connect for cloud deployments instead of long-lived cloud
  secrets.
- Pin third-party actions to a full commit SHA when the workflow has deployment,
  release, or secret access.
- Use `concurrency` for deploy, release, preview, and expensive workflows.
- Treat pull request titles, bodies, comments, issue text, branch names, and
  other event payload fields as untrusted input.
- Do not echo secrets or pass them through command-line flags that can appear in
  logs.
- Use caches only for package manager caches or build outputs that are safe to
  restore across runs.
- Upload artifacts with explicit retention and without secrets, credentials, or
  raw production data.
- Keep shell steps small and fail-fast; move repeated logic into scripts or
  composite actions.
