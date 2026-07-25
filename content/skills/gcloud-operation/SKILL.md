---
name: gcloud-operation
description: "Run or prepare a safe Google Cloud CLI operation. Use when: using gcloud, changing GCP resources, deploying to GCP. Do not use: for local CLI exploration, non-GCP operations."
---


# GCloud Operation

Operate Google Cloud with explicit project, account, and permission context.

## Process

1. Identify the target project, region/zone, account, and service account
   impersonation requirements.
2. Inspect the current command or script for implicit `gcloud` config.
3. Prefer read-only commands first, using `--format` for machine-readable
   output when needed.
4. For mutating commands, state the exact resource, project, and consequence.
5. Prefer service account impersonation over downloaded service account keys.
6. Ask for explicit human approval before delete, deploy, IAM, billing,
   service enable/disable, or policy-changing commands.
7. Record the final command, output summary, and verification step.

## Rules

- Do not rely on the developer's active `gcloud` project in scripts.
- Do not commit service account keys, ADC files, access tokens, or kubeconfigs.
- Do not add `--quiet` unless the script already makes destructive choices
  explicit.
- Distinguish gcloud CLI auth from Application Default Credentials.
- Prefer least-privilege service accounts for automation.

## Output

- Target project/account/region summary
- Safe command sequence
- Approval gate for mutations
- Verification command or observed result
