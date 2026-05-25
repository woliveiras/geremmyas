---
name: daily
description: "Record daily work progress and prepare scrum standup notes from git commits, PRs, session history, and manual notes. Output lives in me/daily/YYYY-MM-DD.md. Use when: end of day log, preparing for daily meeting, weekly summary, standup, daily standup, scrum daily, what I did today, /daily."
argument-hint: "Mode: log | standup | week"
---

# Daily

Record daily work and prepare standup notes for scrum daily meetings.

The output lives in the repository where the skill runs:

```text
me/daily/
├── 2026-05-25.md
├── 2026-05-24.md
├── 2026-05-23.md
└── ...
```

Each file is a self-contained daily record. Files are kept indefinitely and can
be used as evidence artifacts for `brag-me collect`.

## Modes

| Mode | Use when | Output |
|---|---|---|
| `log` | End of day or anytime during the day to record finished work | New or updated `me/daily/YYYY-MM-DD.md` |
| `standup` | Before daily meeting, preparing what to say | Concise standup summary from yesterday + today |
| `week` | End of week or Monday review | Aggregated summary of the week's daily files |

## Procedure

### Mode: `log`

1. Determine the target date. Default is today's local date.
2. Check if `me/daily/YYYY-MM-DD.md` already exists. If yes, update it
   preserving the `Notes` section and any manual edits. If no, create it from
   [daily-template.md](./assets/daily-template.md).
3. Collect evidence from available sources in this priority order:
   - User-provided notes or context from the conversation
   - Local git log: commits authored today (or since the last daily entry)
   - Session store: Copilot session history from the last 24h (use
     `session_store_sql` standup action when available)
   - GitHub PRs: opened, merged, or reviewed today (use `gh` CLI)
   - Branch status: current branch name, diff stats, uncommitted changes
4. Fill the `Done` section with completed items. Each item should include a
   short description and provenance (commit hash, PR link, or file path).
5. Fill the `Doing` section with in-progress work inferred from current branch,
   open PRs, or user-provided context.
6. Leave `Blockers` empty unless the user mentions them or they are obvious
   from failed CI, merge conflicts, or similar signals.
7. Preserve the `Notes` section content from any previous version of the file.
8. Report: file created/updated, sources checked, items logged, and gaps.

### Mode: `standup`

1. Read today's file (`me/daily/YYYY-MM-DD.md`) and yesterday's file (or the
   most recent previous file if yesterday is missing).
2. Produce a concise verbal summary suitable for a 2-minute standup:
   - **Yesterday**: Done items from the previous daily file
   - **Today**: Doing items from today's file (or carried from yesterday)
   - **Blockers**: Any listed blockers
3. If no daily files exist yet, collect evidence as in `log` mode and produce
   the summary directly.

### Mode: `week`

1. Read all `me/daily/*.md` files from the current week (Monday to today).
2. Aggregate Done items into a single list grouped by day.
3. Highlight recurring blockers or items that stayed in Doing across days.
4. Output the summary to stdout (do not create a separate file unless asked).

## Evidence Collection Rules

- Do not fabricate commits, PRs, issues, or work items.
- Mark unknown or uncertain items as `TBD` with a note on how to verify.
- Include provenance for every item: commit hash, PR number, file path, or
  session reference.
- Use authenticated local tools when available (`git`, `gh`, `gcloud`,
  `session_store_sql`). If a tool is unavailable, note the gap and continue.
- Keep entries concise. Each Done/Doing item should be one line.

## File Rules

- One file per calendar day. Running `log` multiple times updates the same file.
- Never overwrite the `Notes` section during updates.
- Generated sections are wrapped in `<!-- GENERATED: section-name -->` markers.
  Content outside these markers is preserved during updates.
- Files are plain markdown, no HTML.
- Files are kept indefinitely. No automatic pruning.
