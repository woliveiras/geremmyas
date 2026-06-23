---
name: llm-integration-review
description: "Design or review an LLM integration in a service. Use when: adding model calls, tools, structured outputs, retries, rate limits. Do not use: for prompt engineering alone, non-service LLM work."
---


# LLM Integration Review

Review LLM service boundaries before code spreads across handlers and domain
logic.

## Process

1. Identify the user-facing capability, model provider, model, latency target,
   cost risk, and data sensitivity.
2. Find the service boundary that owns model calls, prompts, tools, retries,
   and structured outputs.
3. Verify inputs are validated and private data is minimized or redacted before
   logging/tracing.
4. Prefer structured outputs for machine-read results.
5. Define timeout, retry, backoff, rate limit, and fallback behavior.
6. Treat tool calls as side effects: authorize, make idempotent where possible,
   and record audit context.
7. Add contract tests for prompt inputs, tool schemas, structured outputs,
   refusals, provider errors, and retry behavior.
8. Document operational knobs: model, temperature, token limits, and cost
   controls.

## Rules

- Do not put provider SDK calls directly in route handlers.
- Do not parse critical machine-readable results from free-form prose.
- Do not log secrets, credentials, full private documents, or raw user data.
- Do not let model output authorize itself or choose privileged operations
  without application checks.

## Output

- Boundary and data-flow summary
- Risk checklist
- Required tests
- Operational settings and follow-ups
