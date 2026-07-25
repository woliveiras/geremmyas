---
description: "Use when writing or reviewing general Python code. Keep this language-level; framework-specific rules belong in separate instruction files."
applyTo: "**/*.py"
---

# Python Conventions

Follow PEP 8, the Zen of Python (PEP 20), and prefer stdlib when it expresses
intent clearly. Target Python 3.11+ as minimum for new projects.

## Project Structure

- Use `src/` layout with `pyproject.toml` as center of configuration
- Keep tests outside the importable package tree (`tests/` at root)
- One formatter, one lint policy, shared via `pyproject.toml`
- Use `uv` for project environments; `pipx` for CLI tools
- Declare `requires-python` explicitly in `pyproject.toml`

## Naming & Style

- `snake_case` for functions, modules, variables; `PascalCase` for classes
- Constants: `UPPER_SNAKE_CASE` at module level
- Private: single underscore prefix (`_internal`)
- Imports grouped: stdlib, third-party, local (enforced by `ruff` or `isort`)
- Prefer absolute imports over relative ones
- One class per file when it has significant logic; utility grouping is fine

## Type Hints

- Type all public function signatures and class attributes
- Prefer `str | None` over `Optional[str]` (Python 3.10+)
- Use `from __future__ import annotations` for forward refs or older versions
- Use `Protocol` for structural typing; avoid ABCs when duck typing suffices
- Use `TypedDict` for typed dictionary contracts at boundaries
- Types are design tools and static checks, not runtime validation substitutes
- Run `mypy --strict` or Pyright in CI

## Error Handling

- Raise specific exceptions; never bare `except:` or `except Exception: pass`
- Handle errors at the appropriate level; do not swallow without logging
- Use `contextlib.suppress(SpecificError)` for intentional silent handling
- Error messages: lowercase, descriptive, no trailing punctuation
- Custom exceptions inherit from a domain base exception
- Never use exceptions for normal control flow

## Resources & Context Managers

- Use `with` for any resource needing deterministic cleanup (files, connections,
  locks, transactions)
- Prefer `contextlib.contextmanager` for simple generator-based managers
- Use `contextlib.asynccontextmanager` for async resource lifecycle
- Never `open(...); ...; close()` manually

## Data Structures

- Use `@dataclass(frozen=True, slots=True)` for value objects
- Use `dataclasses` to reduce boilerplate on plain data holders
- Prefer `pathlib.Path` over string manipulation for file paths
- Use generators/iterators for streaming; avoid materializing unbounded data
- Prefer stdlib collections (`deque`, `defaultdict`, `Counter`) over reinvention

## Async

- Use `async`/`await` only for I/O-bound concurrency (network, DB, queues)
- Use `asyncio.TaskGroup` (Python 3.11+) for structured concurrency
- Never call blocking I/O inside the event loop; use `run_in_executor` or
  dedicated async libraries
- Use `multiprocessing` for CPU-bound work; threads for legacy blocking APIs
- Every task must have clear cancellation and error propagation

## Functions & Modules

- Keep functions focused: one function, one job
- Avoid hidden mutable module state; pass dependencies explicitly
- Prefer small, composable functions over large monolithic ones
- Return early on errors; keep happy path at lowest indentation
- Avoid deep nesting; extract helpers when indentation exceeds 3 levels

## Testing

- Use `pytest` as default runner; stdlib `unittest` when zero-dep is required
- Table-driven tests via `@pytest.mark.parametrize`
- Use `Hypothesis` for property-based testing on parsers, validators, numerics
- Measure branch coverage with `coverage.py`; fail CI under threshold
- Don't mock what you don't own; wrap behind protocols/interfaces

## Observability

- Use `logging` (stdlib) with structured messages and correct levels
- Use OpenTelemetry for distributed traces in services
- Never log secrets, tokens, PII, or full request/response bodies
- Log context: request ID, user ID (hashed), operation name

## Linting & Formatting

- `ruff check` as primary linter (or `flake8` if legacy)
- One formatter only: `ruff format` or `black`
- `mypy --strict` or Pyright for type checking
- `bandit` for security scanning
- `pip-audit` for dependency vulnerabilities
- All tools configured in `pyproject.toml`

## Anti-patterns

- `except Exception: pass` (silences real errors and attacks)
- Mutable default arguments (`def f(items=[])`)
- Global mutable state
- `eval()`/`exec()` with external input
- `shell=True` in `subprocess` calls
- Manual `open/close` without `with`
- String formatting for SQL queries
- Importing `*` from modules
- `pickle`/`yaml.load` on untrusted data
- Multiple conflicting formatters or linters
