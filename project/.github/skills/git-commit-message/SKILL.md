---
name: git-commit-message
description: "Write a commit message in Conventional Commits format. Use when: writing a commit, commit message, generate commit, conventional commit."
---

# Git Commit Message

Generate a commit message following the Conventional Commits format.

## Format

`type(scope): description`

## Types

- `feat` — new feature or capability
- `fix` — bug fix
- `docs` — documentation changes only
- `test` — adding or updating tests
- `refactor` — code change that neither fixes a bug nor adds a feature
- `chore` — maintenance tasks (deps, CI, configs)
- `style` — formatting, whitespace (no code logic change)
- `perf` — performance improvement

## Rules

- Use lowercase for type and description
- Scope is optional but recommended (e.g., `feat(auth): add login`)
- Description is imperative mood: "add", not "added" or "adds"
- Keep the first line under 72 characters
- Breaking changes: add `!` after type — `feat!: remove legacy API`

## Procedure

1. Read the staged changes (`git diff --cached`)
2. Identify the primary type of change
3. Determine the scope from the affected module/directory
4. Write a concise description in imperative mood
5. Add a body only if the "why" isn't obvious from the description
