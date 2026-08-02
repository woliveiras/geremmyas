---
description: "Use when writing or reviewing Terraform configuration, modules, variables, outputs, state changes, imports, and refactors."
applyTo: "**/*.tf, **/*.tfvars, **/*.tfvars.json"
---

# Terraform Conventions

- Run `terraform fmt` and `terraform validate` before proposing a plan.
- Use remote state with locking for shared or production infrastructure.
- Commit `.terraform.lock.hcl` for root modules so provider selections and
  checksums are reproducible.
- Pin provider and module versions with constraints; do not track floating
  branches for production modules.
- Do not put secrets directly in configuration, `tfvars`, outputs, or local
  state. Terraform can persist them in state and plan files.
- Keep root modules small and environment-specific; use child modules for
  reusable infrastructure.
- Use declarative `moved` blocks for resource address refactors and `import`
  blocks for bringing existing resources under management.
- Review `terraform plan` for deletes, replacements, IAM changes, networking,
  data stores, and drift before apply.
- Apply or mutate state autonomously only in a verified local, disposable, or
  test environment with rollback or recreation; prefix the guarded command with
  the matching `GEREMMYAS_TARGET=local`, `test`, or `disposable` marker. Never run `terraform destroy`.
  Treat ambiguous targets as protected; production mutation requires explicit user authorization.
