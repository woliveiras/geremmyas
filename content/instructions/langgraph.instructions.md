---
description: "Use when writing or reviewing LangGraph Python workflows. Covers state schemas, graph nodes, checkpoints, interrupts, and resumable execution."
applyTo: "**/graphs/**/*.py, **/agents/**/*.py, **/workflows/**/*.py, **/*graph*.py, **/*agent*.py, **/*workflow*.py"
---

# LangGraph Conventions

- Define a stable state schema before adding nodes or edges.
- Keep nodes focused: read state, perform one step, and return explicit state
  updates.
- Make node outputs idempotent when a checkpointed run may resume after a
  partial failure.
- Compile long-running or human-in-the-loop graphs with a checkpointer and a
  stable `thread_id`.
- Use `interrupt()` for human-in-the-loop pauses; do not wrap interrupts in a
  broad `except Exception`.
- Keep interrupt calls in deterministic order inside a node so resume values
  map correctly.
- Test graph behavior at node, edge, checkpoint, interrupt, and resume
  boundaries.
