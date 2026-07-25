---
name: terraform-change
description: "Plan and review a Terraform infrastructure change safely. Use when: editing Terraform, importing resources, moving resource addresses. Do not use: for infrastructure design, non-Terraform IaC."
---


# Terraform Change

Change Terraform with a plan-first workflow and explicit human approval for
state or infrastructure mutations.

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
7. Ask for explicit human approval before `apply`, `destroy`, `state rm`,
   `state mv`, or any command that mutates remote state.
8. After apply, capture outputs, follow-up verification, and any rollback notes.

## Rules

- Do not run `terraform apply`, `destroy`, or state mutation commands without
  explicit approval.
- Do not put secrets in `.tf`, `.tfvars`, outputs, plan files, or committed
  state.
- Treat plan files as sensitive because they can contain secret values.
- Prefer small, reviewable infrastructure changes over broad refactors.
- If the backend or workspace is unclear, stop and clarify before planning.

## Output

- Environment and backend summary
- Commands run or proposed
- Plan summary with risk areas
- Approval needed before any mutation
- Post-apply verification or rollback notes
