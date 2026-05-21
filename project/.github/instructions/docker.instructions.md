---
description: "Use when writing or reviewing Dockerfiles, .dockerignore files, and container image build configuration."
applyTo: "**/Dockerfile, **/Dockerfile.*, **/.dockerignore"
---

# Dockerfile Conventions

- Use multi-stage builds when build tools are not needed at runtime.
- Choose the smallest maintained runtime image that supports the app; avoid
  generic distro images when a slimmer runtime works.
- Pin base images with explicit versions; do not use `latest`.
- Include a `.dockerignore` that excludes `.git`, local env files, dependency
  folders, test output, coverage, databases, and build artifacts.
- Do not bake secrets into images with `COPY`, `ARG`, or `ENV`; use BuildKit
  secret mounts or runtime secret injection.
- Create and switch to a non-root user for runtime stages unless the base image
  is intentionally rootless already.
- Copy dependency manifests before source files when it improves build cache
  behavior.
- Prefer `COPY` over `ADD` unless archive extraction or remote source support
  is explicitly required.
- Set an explicit `WORKDIR`.
- Log to stdout/stderr; do not write application logs to files inside the
  container image.
