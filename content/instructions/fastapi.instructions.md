---
description: "Use when writing or reviewing FastAPI applications. Covers routers, dependency injection, request/response models, and lifecycle patterns."
applyTo: "**/routers/**/*.py, **/routes/**/*.py, **/api/**/*.py, **/main.py"
---

# FastAPI Conventions

- Routers live in `routers/` or `api/` by domain.
- Use `APIRouter` with explicit `prefix` and `tags`.
- Use Pydantic models for request bodies and responses; avoid raw `dict`
  contracts at API boundaries.
- Declare `response_model` when the public response shape differs from the
  internal object returned by application code.
- Use `Depends()` for dependencies such as database sessions, auth, settings,
  and service objects.
- Use `HTTPException` for expected client-facing errors.
- Use `status.HTTP_*` constants instead of magic numbers.
- Prefer async route handlers for I/O-bound work.
- Use a lifespan handler for startup/shutdown instead of deprecated event
  hooks.
- Keep business logic out of route handlers; route handlers should validate,
  authorize, call application code, and shape the response.
