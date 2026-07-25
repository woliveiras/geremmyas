---
description: "Use when writing or reviewing provider-agnostic LLM integrations in services. Covers structured outputs, retries, limits, secrets, tools, and observability."
applyTo: "**/llm/**/*.py, **/ai/**/*.py, **/agents/**/*.py, **/*llm*.py, **/*openai*.py, **/*anthropic*.py, **/*agent*.py, **/*prompt*.py"
---

# LLM Service Conventions

- Keep provider calls behind an application service interface; do not spread SDK
  calls through route handlers or domain code.
- Use structured outputs for machine-read responses and validate the parsed
  result before using it.
- Set explicit timeout, retry, backoff, and rate-limit behavior at the boundary
  that calls the model.
- Track token limits, model choice, temperature, and cost-relevant settings in
  code or config, not hidden prompt text.
- Never log secrets, raw credentials, full user documents, or unredacted model
  inputs when they may contain private data.
- Treat tool calls as side-effecting operations: authorize them, make them
  idempotent where possible, and record enough context to audit failures.
- Add contract tests for prompts, tool schemas, structured outputs, refusals,
  provider errors, and retry behavior.
