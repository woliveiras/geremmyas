---
description: "Use when writing or reviewing Google Cloud CLI scripts, Cloud Build config, deployment scripts, and GCP automation."
applyTo: "**/cloudbuild*.yml, **/cloudbuild*.yaml, **/.gcloudignore, **/scripts/**/*.sh, **/deploy/**/*.sh, **/*gcloud*.sh, **/*gcp*.sh"
---

# Google Cloud Conventions

- Make project, region, zone, and account explicit in scripts; do not rely on a
  developer's active `gcloud` config.
- Distinguish gcloud CLI auth from Application Default Credentials. Use
  `gcloud auth application-default login` only when client libraries need ADC.
- Prefer service account impersonation over downloaded service account keys.
- Use `--impersonate-service-account` or the matching gcloud config property
  for privileged operations.
- Use least-privilege IAM roles and project-specific service accounts.
- Add `--quiet` only in scripts that have already made all destructive choices
  explicit.
- Emit command output with `--format` when another script or agent will parse
  it.
- Never commit service account JSON keys, ADC files, access tokens, or generated
  kubeconfigs.
- Check the target project before mutating commands such as deploy, delete,
  update, set-iam-policy, or services enable/disable.
