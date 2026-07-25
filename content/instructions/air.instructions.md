---
description: "Use when writing or reviewing Air live reload configuration for Go development environments."
applyTo: "**/.air.toml, **/air.toml, **/air*.toml, **/Dockerfile, **/docker-compose*.yml, **/docker-compose*.yaml, **/compose*.yml, **/compose*.yaml"
---

# Air Conventions

- Treat Air as a development-only hot reload tool; do not make production
  images or runtime scripts depend on it.
- Keep `.air.toml` checked in when the project needs a non-default build
  command, binary path, working directory, or excludes.
- Exclude generated, vendor, build, coverage, tmp, and asset output directories
  from watched paths.
- Build into a disposable `tmp/` or `bin/` path that is ignored by Git.
- In Docker or Compose development, mount the project directory into the
  container path used by Air.
- Pass application runtime args after `--` so they are not parsed as Air flags.
- Prefer a project-local tool command when available, such as `go tool air`, so
  onboarding does not depend on a global binary.
