---
description: "Use when writing or reviewing Python code that handles input, secrets, dependencies, serialization, or crosses trust boundaries."
applyTo: "**/*.py"
---

# Python Security

## Input Validation

- Validate type, format, size, and allowed values at every trust boundary
- Reject invalid input early; do not sanitize and hope for the best
- Use Pydantic, `dataclasses`, or explicit checks for structured validation
- Limit request body size at the framework level (middleware or config)
- Never trust client-supplied content types, filenames, or paths

## SQL Injection

- Never compose SQL by string concatenation or f-string with external data
- Use parameterized queries: `cursor.execute("SELECT ... WHERE id = ?", (id,))`
- ORM queries are generally safe but watch for raw SQL escape hatches
- Audit any use of `.extra()`, `RawSQL`, or `text()` with user input

## Command Injection

- Avoid `subprocess.run(..., shell=True)` with external input
- Even with `shell=False`, validate arguments against an allowlist
- Prefer stdlib or library APIs over shelling out when possible
- `shlex.quote()` is last resort, not first choice

## Deserialization

- Never `pickle.load` from untrusted sources (arbitrary code execution)
- Use `yaml.safe_load()` not `yaml.load()` for external YAML
- Prefer JSON for data exchange with untrusted parties
- `ast.literal_eval()` is not a sandbox; do not use for untrusted input
- `eval()`/`exec()` with external input is always a vulnerability

## Secrets & Credentials

- Never hardcode secrets in source code or version them
- Read secrets from environment or a secret manager at runtime
- Fail explicitly if a required secret is missing: `os.environ["KEY"]`
- Use `secrets.token_urlsafe()` for cryptographically secure random values
- Use `hmac.compare_digest()` for constant-time comparisons of tokens/MACs
- Do not log secrets, tokens, or credentials at any log level

## Dependencies & Supply Chain

- Pin dependencies for reproducible builds
- Run `pip-audit` in CI to detect known vulnerabilities
- Use Trusted Publishing (OIDC) for PyPI; avoid long-lived upload tokens
- Audit new dependencies before adding; prefer well-maintained packages
- Keep `requirements.txt` or lock files committed and up to date

## HTTPS & Network

- Always validate TLS certificates; never `verify=False` in production
- Bandit flags `verify=False` as high severity; treat it as such
- Set timeouts on all network calls; unbounded waits are DoS vectors
- Use `httpx` or `requests` with explicit timeout parameters

## Cryptography

- Use `secrets` module for tokens, keys, and random IDs
- Use `hashlib` with salt for password hashing (or `argon2-cffi`/`bcrypt`)
- Never use `random` module for security-sensitive values
- Do not implement custom crypto; use established libraries

## Code Execution & Sandboxing

- Python has no in-process sandbox; do not trust `exec` restrictions
- Run untrusted code only in isolated containers with resource limits
- Use non-root users, namespaces, cgroups, and minimal filesystem exposure
- Treat `unsafe` annotation evaluation and deep recursion as DoS vectors

## Logging & Exception Safety

- Never `except Exception: pass` without explicit justification and logging
- Log validation failures, auth events, and unexpected exceptions
- Do not log PII, tokens, or raw request bodies
- Structured logging (JSON) enables better alerting and correlation

## Checklist

- [ ] Input validated at boundaries (type, format, size)
- [ ] No SQL by string concatenation
- [ ] No `shell=True` with external arguments
- [ ] No `pickle`/`yaml.load`/`eval` on untrusted data
- [ ] Secrets from environment or secret manager, never in code
- [ ] `pip-audit` or equivalent in CI
- [ ] TLS verification enabled; timeouts set
- [ ] Exceptions handled explicitly with context logged
- [ ] No `verify=False` in production code
