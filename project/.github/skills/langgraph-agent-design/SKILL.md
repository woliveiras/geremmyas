---
name: langgraph-agent-design
description: "Design or review a LangGraph agent workflow. Use when: creating a graph, choosing state schema, nodes, tools, checkpoints, interrupts. Do not use: for simple chains, non-agent patterns."
---


# LangGraph Agent Design

Design LangGraph workflows around explicit state, resumability, and controlled
side effects.

## Process

1. Define the agent goal, entrypoint, terminal states, and user-visible output.
2. Define the state schema before nodes and edges.
3. List tools and side effects, including authorization and idempotency needs.
4. Split nodes by responsibility: read state, do one step, return state updates.
5. Choose persistence/checkpointer strategy for long-running or resumable work.
6. Add `interrupt()` points for human-in-the-loop decisions.
7. Define retry, resume, cancellation, and failure behavior.
8. Plan tests for node behavior, graph routing, checkpoints, interrupts, and
   tool errors.

## Rules

- Keep interrupt order deterministic inside each node.
- Do not hide mutable state outside the graph state unless it is a real external
  system.
- Do not make non-idempotent tool calls without audit and recovery strategy.
- Prefer simple graphs until the workflow needs branches, persistence, or HITL.

## Output

- State schema outline
- Node and edge map
- Tool and side-effect inventory
- Checkpoint/HITL strategy
- Test plan
