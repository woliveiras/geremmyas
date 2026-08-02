---
name: git-commit
description: "Create a safe Git commit with Conventional Commits format. Use when: committing changes, reviewing staged changes. Do not use: for commit history rewriting, interactive rebasing."
---


# Git Commit

Local commits are the default after a coherent task-owned slice passes fresh
tests, docs reconciliation, and independent review. This slice evidence is not
the feature lifecycle state `Verified`; a feature may remain `In Progress`
across several atomic commits. The primary agent owns staging and commits for
integrated work; subagents do not run Git.

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

1. Inspect staged and unstaged state separately with `git status --short`,
   `git diff --cached`, and `git diff` before touching the index. Record whether
   pre-existing staged state exists.
2. Identify the exact task-owned files or hunks for the implementation, tests,
   documentation, and artifacts that form one verified slice. Exclude unrelated
   or ambiguous changes without asking the user to select them.
3. Stage clean task-owned paths explicitly. For a mixed-hunk file, build and
   inspect a task-owned patch and apply only it with `git apply --cached`; never
   stage the whole path.
4. If pre-existing staged state is unrelated or overlaps the slice, do not
   mutate or unstage it. Move the task-owned patch to an isolated worktree with
   its own branch and index, then verify the original index remains unchanged.
   An alternate index may prepare or inspect the patch, but must not advance the
   current branch while the original index is staged. If safe isolation is
   unavailable, leave the index untouched and report the blocker. Never use
   `git add .` or another repository-wide shortcut.
5. Re-read `git status --short`, `git diff --cached --stat`, and
   `git diff --cached` after staging. Remove any unrelated, generated, local,
   secret-bearing, or unsafe change that this procedure staged without
   discarding working-tree content. Never unstage user-owned state.
6. Derive the Conventional Commit type, scope, subject, and optional body from
   the complete cached diff. Re-read the cached diff after deriving the message
   and confirm the message still describes every staged change.
7. Commit without asking the user to approve files, hunks, or the message.
8. Show the created commit with `git rev-parse --short HEAD` and
    `git show --stat --oneline --no-renames HEAD`.
9. If the commit fails, report the failure and leave the reviewed files staged
   safely for diagnosis. Do not broaden scope or mutate history to recover.

## Safety Rules

- Do not push.
- A local commit never implies permission to amend, rebase, merge, tag, release,
  publication, deploy, or change production. Do not push; push always requires
  explicit user authorization.
- Honor read-only, plan-only, no-edits, and no-commits session overrides.
- When an override disables commit, report changed files or, for read-only and
  plan-only work, the proposed files that would have changed.
- Commit only task-owned changes. Preserve dirty and unrelated user work.
- Never stage ignored files, secrets, credentials, local environment files, or
  unrelated generated artifacts. Report suspicious content and exclude it.
- Keep implementation, tests, required docs, and durable workflow artifacts for
  one slice together; split independent behavior into separate atomic commits.
- Commit after fresh verification and review establish slice evidence; do not
  wait for the whole feature to reach lifecycle state `Verified`.
