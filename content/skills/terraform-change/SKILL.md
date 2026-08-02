---
name: terraform-change
description: "Plan and review a Terraform infrastructure change safely. Use when: editing Terraform, importing resources, moving resource addresses. Do not use: for infrastructure design, non-Terraform IaC."
---


# Terraform Change

Change Terraform with a plan-first workflow and contextual environment authority.

## Process

1. Identify the workspace, backend, target environment, provider versions, and
   expected blast radius.
2. Read the relevant `.tf`, `.tfvars`, module, backend, and lockfile context.
3. Run or request `terraform fmt` and `terraform validate` before planning.
4. For refactors, prefer `moved` blocks. For existing unmanaged resources,
   prefer `import` blocks or a documented import command.
5. Produce or inspect `terraform plan`.
6. Summarize creates, updates, replacements, deletes, IAM changes, networking,
   data stores, and state changes.
7. Apply or mutate state autonomously only for a verified local, disposable, or
   test environment with rollback or recreation; then prefix the command with
   `GEREMMYAS_TARGET=local`, `test`, or `disposable`. Production requires explicit
   user authorization for every mutation; an ambiguous target is protected.
8. After apply, capture outputs, follow-up verification, and any rollback notes.

## Rules

- Never run `terraform destroy`; use a scoped, recoverable alternative. Do not
  mutate production without explicit user authorization.
- Do not put secrets in `.tf`, `.tfvars`, outputs, plan files, or committed
  state.
- Treat plan files as sensitive because they can contain secret values.
- Prefer small, reviewable infrastructure changes over broad refactors.
- If the backend or workspace is unclear, inspect configuration and continue
  read-only; do not mutate until the environment is proven non-production.

## Output

- Environment and backend summary
- Commands run or proposed
- Plan summary with risk areas
- Authority evidence for local, disposable, or test mutation; authorization for production
- Post-apply verification or rollback notes
