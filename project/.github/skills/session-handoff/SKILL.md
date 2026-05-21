---
name: session-handoff
description: "Create a concise handoff for another Copilot session or agent. Use when: pausing work, switching agents, summarizing current progress, or preparing the next implementation session."
---

# Session Handoff

Create a short handoff that lets a fresh session continue without rereading the
entire conversation.

## Process

1. Identify the next session's purpose from the user's request.
2. Reference existing artifacts instead of duplicating them:
   - PRD
   - specs
   - `plan.md`
   - `tasks.md`
   - ADRs/RFCs
   - bugfix documents
   - docs
3. Summarize:
   - current state
   - completed work
   - active task
   - commands already run
   - verification evidence
   - known risks and open decisions
   - recommended next action
4. Save outside the repo unless the user asks for a committed artifact.

## Rules

- Redact secrets, tokens, credentials, and personal data.
- Do not duplicate long content from artifacts that already exist.
- Be explicit about what has and has not been verified.
