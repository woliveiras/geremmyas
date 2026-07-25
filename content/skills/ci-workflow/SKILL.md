---
name: ci-workflow
description: "Create, review, or debug a CI/CD workflow. Use when: editing GitHub Actions, fixing CI failures, adding release/deploy automation. Do not use: for general DevOps, infrastructure as code."
---


# CI Workflow

Build CI/CD workflows that are reliable, scoped, and safe with untrusted input.

## Process

1. Identify the workflow goal: test, lint, build, preview, release, deploy, or
   maintenance.
2. Inspect existing workflow conventions, scripts, package managers, and
   required secrets.
3. Define triggers and concurrency before adding jobs.
4. Set minimal workflow and job permissions.
5. Add jobs in dependency order with small shell steps or reusable scripts.
6. Configure cache only for safe package-manager or build caches.
7. Use OIDC for cloud deploys when available.
8. Add artifact upload only for intentional outputs with explicit retention.
9. Verify with local syntax review and the relevant CI result or log.

## Rules

- Treat event payload text as untrusted input.
- Do not expose secrets in logs, command-line flags, artifacts, or PR comments.
- Pin third-party actions to a commit SHA for sensitive workflows.
- Do not grant broad `write-all` permissions.
- Keep deploy and release workflows protected with environment approvals or
  equivalent human gates.

## Output

- Workflow purpose and trigger summary
- Permission and secret rationale
- Cache/artifact choices
- Verification result or failing CI evidence
