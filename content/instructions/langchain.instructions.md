---
description: "Use when writing or reviewing LangChain Python code. Covers agents, tools, runnables, retrievers, structured outputs, and tracing boundaries."
applyTo: "**/chains/**/*.py, **/agents/**/*.py, **/rag/**/*.py, **/retrievers/**/*.py, **/*chain*.py, **/*agent*.py, **/*retriever*.py"
---

# LangChain Conventions

- Prefer explicit runnables, agents, tools, and retrievers over hidden helper
  chains that obscure inputs and outputs.
- Type tool arguments and returns; use Pydantic models when the tool contract is
  more than a primitive value.
- Keep tool descriptions operational and permission-aware; do not expose tools
  that the current user or workflow cannot use.
- Use structured output for machine-read results instead of parsing prose.
- Keep retrieval, prompt construction, model invocation, and post-processing as
  separate steps when debugging or testing would otherwise be difficult.
- Pass callbacks, tracing, runtime config, and metadata through invocation
  boundaries instead of relying on module globals.
- Add narrow tests around prompt inputs, retriever results, tool schemas, and
  structured output parsing.
