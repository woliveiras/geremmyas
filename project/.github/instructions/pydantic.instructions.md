---
description: "Use when writing or reviewing Pydantic v2 models, settings, validation boundaries, and serialization contracts."
applyTo: "**/schemas/**/*.py, **/models/**/*.py, **/settings/**/*.py, **/*schema*.py, **/*model*.py, **/*settings*.py, **/*dto*.py"
---

# Pydantic Conventions

- Use `BaseModel` for external contracts and validation boundaries, not as the
  default shape for every internal domain object.
- Prefer Pydantic v2 APIs: `model_validate()`, `model_dump()`, and
  `model_dump_json()` instead of v1 parsing or serialization helpers.
- Use `Field()` for constraints, aliases, examples, and descriptions that are
  part of the public contract.
- Use `ConfigDict` only when model-wide behavior is intentional and documented
  by the boundary.
- Separate API, persistence, settings, and domain models when their validation
  rules or serialized shapes diverge.
- Use `pydantic-settings` for environment-backed configuration; validate
  settings at startup before application services read them.
- Avoid accepting or returning raw `dict` values at public boundaries when a
  typed model would make validation, serialization, or tests clearer.
- Test required fields, defaults, aliases, rejected input, and serialized output
  for models used at API, tool, config, or LLM boundaries.
