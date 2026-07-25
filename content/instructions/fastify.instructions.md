---
description: "Use when writing or reviewing Fastify applications. Covers plugins, schemas, hooks, logging, and error handling."
applyTo: "**/routes/**/*.ts, **/routes/**/*.js, **/plugins/**/*.ts, **/plugins/**/*.js, **/server.ts, **/server.js, **/app.ts, **/app.js"
---

# Fastify Conventions

- Model each feature as a plugin when it owns routes, decorators, hooks, or
  dependencies.
- Use JSON Schema for route `body`, `querystring`, `params`, and `response`
  contracts.
- Treat schemas as application code. Never build validation or serialization
  schemas directly from user input.
- Keep async I/O out of schema validation; use hooks such as `preHandler` after
  validation.
- Register shared schemas with stable `$id` values and reuse them with `$ref`.
- Use Fastify's request logger instead of ad hoc console logging.
- Centralize expected errors with `setErrorHandler`; keep 404 handling in
  `setNotFoundHandler`.
- Return plain data from handlers unless the route needs explicit status,
  headers, streaming, or early replies.
