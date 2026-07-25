---
description: "Use when writing or reviewing API handlers, controllers, route modules, service boundaries, or backend input/output code."
applyTo: "**/api/**/*, **/routes/**/*, **/routers/**/*, **/controllers/**/*, **/handlers/**/*, **/middleware/**/*"
---

# API Security Conventions

- Validate and normalize all input at the API boundary.
- Authorize every request before accessing user-owned or tenant-owned data.
- Do not trust IDs, roles, tenant IDs, or ownership claims from the client.
- Use parameterized queries or safe query builders; never concatenate SQL.
- Return generic errors to clients and log diagnostic detail server-side.
- Never log secrets, tokens, passwords, session cookies, or PII.
- Apply rate limits or abuse controls to authentication, write, upload, and
  expensive read endpoints.
- Keep authentication, authorization, validation, and business logic visibly
  separated.
