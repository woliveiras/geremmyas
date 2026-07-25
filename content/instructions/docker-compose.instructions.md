---
description: "Use when writing or reviewing Docker Compose files for local development, test environments, or lightweight service orchestration."
applyTo: "**/docker-compose*.yml, **/docker-compose*.yaml, **/compose*.yml, **/compose*.yaml"
---

# Docker Compose Conventions

- Keep Compose focused on local development, tests, or simple environments;
  document when it is not production orchestration.
- Declare service healthchecks when other services depend on readiness.
- Use `depends_on` only for startup ordering; pair it with healthchecks when
  readiness matters.
- Put persistent data in named volumes and disposable caches in tmpfs or
  bind-mounted ignored directories.
- Use service-specific networks instead of exposing every service on the host.
- Store non-secret defaults in `env_file`; pass secrets through Compose secrets
  or environment provided by the caller.
- Avoid `privileged`, host networking, broad volume mounts, and Docker socket
  mounts unless the file documents why they are required.
- Set restart policy deliberately; do not hide crash loops during local
  debugging.
- Prefer explicit image tags or local `build` blocks with stable Dockerfiles.
