---
name: subagent-selection
description: "Decision framework for delegate-vs-inline work optimization. Use when: deciding whether to use inline work or spawn specialist agents, cost-benefit analysis needed. Do not use: for simple read-only operations, single-file edits, or direct implementation."
---


# Subagent Selection

## When to Use

- Before running a subagent (any agent call)
- When choosing between inline work vs delegation
- To optimize context window efficiency
- When research/exploration will be expensive inline
- When parallel execution accelerates the overall task

## When NOT to Use

- For trivial changes (just do it inline)
- When you have real-time user input needed (agent can't get clarification)
- For security-sensitive operations (review with user first)
- When context is already bloated (delegate to compress it, not expand it)

## Decision Matrix

Examples of when to delegate vs inline work, and which agent to use:

| Scenario | Delegate? | Agent | Why |
|----------|-----------|-------|-----|
| "Map 14 skill dirs in 3 repos" | YES | Explore | Parallel reads, compressed output |
| "Find all usages of `makeRequest`" | NO | — | 2min grep-search inline |
| "Should we refactor module X?" | YES | architect | Explore + ADR proposal = expensive |
| "Add type hint to 1 var" | NO | — | 30sec edit inline |
| "Analyze 200-line test failure" | YES | Explore/bugfix-loop | Research + hypothesis isolation |
| "Fix typo in README" | NO | — | Just fix it |
| "Design RAG pipeline for new codebase" | YES | Explore → spec-writer | Need pattern inventory first |
| "Write error message" | NO | — | Inline in 20 seconds |
| "Integrate LLM service, need guidance" | YES | runSubagent + llm-integration-review | Multi-file, decision-heavy |

## Cost Analysis

**Inline Costs**
- Direct: Token usage for exploration
- Indirect: Clutters main context
- Parallelization: Sequential (can't parallelize)

**Delegate Costs**
- Direct: Subagent startup + output injection
- Indirect: Less context pollution
- Parallelization: Can run multiple agents in parallel
- Benefit: Compressed output (caveman mode)

**Delegate if**:
- Total tokens (delegation + compressed output) < (inline exploration tokens)
- Exploration is substantial (>100 lines of research)
- Multiple parallel queries would help
- Main context is already >80% of limit

**Inline if**:
- Work is <5 min
- Single, simple query
- No need for context compression
- User needs interactive back-and-forth

## Agent Selection Quick Ref

| Need | Agent | Output | Notes |
|------|-------|--------|-------|
| "Map codebase" | Explore | Structured summary | Quick/medium/thorough modes |
| "Find code patterns" | Explore | Code snippets + locations | Better than grep for semantics |
| "Spec from requirements" | spec-writer | spec.md + plan.md + tasks.md | Wraps interview + generate |
| "Review impl vs spec" | reviewer | Checklist + gaps | Spec-driven only |
| "Architecture options" | architect | ADR + proposals | Deep module analysis |
| "Location of file X" | — | Use vscode_listCodeUsages | Don't delegate, use tool |

## Parallel Execution

**Safe to parallelize** (independent):
- Multiple Explore agents on different repos
- grep-search + file-read + semantic-search (all read-only)
- Multiple agents reviewing different specs

**NOT safe to parallelize**:
- Git operations on same branch
- File edits to same target
- Anything with shared state

## Context Recovery Pattern

If main context is bloated after delegation:
1. Subagent delivers compressed result
2. You read result (low token cost)
3. If need details, ask subagent to expand ONE section
4. Do NOT re-run full exploration

## Anti-Patterns

**Over-Delegating**
- Every 2-line change → subagent
- Exploration that takes 30 seconds → delegated
- Binary decisions → agent analysis (just decide)

**Under-Delegating**
- 200 files to search → manual grep loop
- "Should we refactor this?" → inline pondering (use architect)
- "Map the RTK pattern" → manual file exploration

**Delegation Waste**
- Delegate, ignore output, delegate again
- Delegate for output already in current context
- Delegate then inline re-do the work

---

**Key Principle**: Delegation is not laziness—it's force multiplying. Use it for expensive exploration and parallelization. Inline for quick fixes. Stay honest about token cost vs benefit.
