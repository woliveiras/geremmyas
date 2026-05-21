---
description: "Use when writing or reviewing Go code that embeds static files, templates, migrations, or assets with embed.FS."
applyTo: "**/*.go"
---

# Go Embed Conventions

- Use `//go:embed` only for files that should be compiled into the binary.
- Keep embedded assets in stable package-relative paths; do not depend on the
  process working directory at runtime.
- Use `embed.FS` for directories or multiple files; use `string` or `[]byte`
  only for a single file.
- Import `_ "embed"` when embedding into `string` or `[]byte` and `embed.FS`
  is not referenced directly.
- Use `fs.Sub` to expose a clean subtree to HTTP servers, template parsers, or
  migration runners.
- Treat embedded files as read-only. Write runtime state, uploads, and generated
  files outside `embed.FS`.
- Add tests that read the embedded paths needed by startup, migrations,
  templates, or static handlers.
