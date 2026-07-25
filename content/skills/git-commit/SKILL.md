---
name: git-commit
description: "Create a safe Git commit with Conventional Commits format. Use when: committing changes, reviewing staged changes. Do not use: for commit history rewriting, interactive rebasing."
---


# Git Commit

Create a local Git commit after reviewing the exact changes that will be
included.

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

1. Inspect the working tree with `git status --short`.
2. Inspect staged changes with `git diff --cached --stat` and
   `git diff --cached`.
3. If there are no staged changes, summarize the unstaged and untracked files
   from `git status --short` and ask the user which exact files should be
   staged.
4. Stage only files explicitly approved by the user. Never run `git add .` as
   the default action.
5. Re-read `git status --short`, `git diff --cached --stat`, and
   `git diff --cached` after staging.
6. Identify the primary type and optional scope from the staged diff.
7. Propose a Conventional Commit message. Add a body only when the "why" is not
   obvious from the subject line.
8. Ask for explicit confirmation before running `git commit`.
9. Run `git commit` only after confirmation, using only the currently staged
   changes.
10. Show the created commit with `git rev-parse --short HEAD` and
    `git show --stat --oneline --no-renames HEAD`.

## Safety Rules

- Do not push.
- Do not amend or rewrite history unless the user explicitly asks for it.
- Do not commit unrelated user changes without explicit approval.
- Do not stage ignored files, secrets, local environment files, or generated
  artifacts unless the user explicitly confirms that exact path.
- If the staged diff includes suspicious secrets or credentials, stop and ask
  the user how to proceed.
- If tests or verification were expected but not run, say that before asking for
  commit confirmation.
