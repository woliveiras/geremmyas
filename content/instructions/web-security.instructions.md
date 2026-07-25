---
description: "Use when writing or reviewing browser UI, route components, forms, client-side auth flows, or user-facing web code."
applyTo: "**/*.tsx, **/*.jsx, **/*.astro, **/*.mdx"
---

# Web Security Conventions

- Treat all user-provided text as untrusted.
- Avoid raw HTML injection. If unavoidable, sanitize first and document why.
- Do not store long-lived secrets or privileged tokens in browser-accessible
  storage.
- Use server-side authorization for protected data; client-side checks are only
  UX hints.
- Use safe URL handling for redirects and links; reject open redirects.
- Keep form validation duplicated at the server boundary; client validation is
  not security.
- Do not expose internal errors, stack traces, or sensitive config in UI states.
