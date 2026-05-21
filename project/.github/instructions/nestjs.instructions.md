---
description: "Use when writing or reviewing NestJS applications. Covers modules, controllers, providers, validation, and request lifecycle boundaries."
applyTo: "**/*.controller.ts, **/*.service.ts, **/*.module.ts, **/*.resolver.ts, **/*.guard.ts, **/*.interceptor.ts, **/*.filter.ts, **/main.ts"
---

# NestJS Conventions

- Keep controllers thin: validate, authorize, call providers, and shape HTTP
  responses.
- Put business logic in injectable providers; register dependencies through
  modules instead of constructing services manually.
- Keep module boundaries explicit. Export only providers that other modules
  should depend on.
- Use DTOs plus pipes for request validation and transformation; do not parse
  route params, query strings, or bodies inline in controllers.
- Prefer framework lifecycle hooks, guards, interceptors, filters, and pipes
  over ad hoc middleware inside handlers.
- Centralize expected API errors with exception filters or domain-specific
  exceptions.
- Validate configuration at startup before providers read environment values.
- Keep persistence, transport, and domain rules in separate providers when the
  feature has meaningful behavior.
