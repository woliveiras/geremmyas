---
spec: "0005"
title: Dashboard serve and watch mode
family: platform
phase: 3
status: Draft
owner: ""
depends_on: [2]
origin: "docs/prds/2026-06-04-dashboard.md"
---

# Spec: Dashboard serve and watch mode

## Context & Motivation

Without `--serve`, the user must regenerate the dashboard manually after each
spec change and refresh the browser. This friction makes the dashboard less
useful during active development.

A local HTTP server with file watching and automatic browser reload provides a
smooth experience: edit a spec → dashboard updates in the browser.

## Requirements

### Functional

- [ ] `--serve` flag starts an HTTP server serving the generated dashboard
- [ ] `--port` flag sets the server port (default 8080)
- [ ] `--watch` flag watches spec files for changes and regenerates on change
  (implies `--serve`)
- [ ] On file change: re-parse only changed files, regenerate affected HTML
  pages, signal browser to reload
- [ ] Browser reload via Server-Sent Events (SSE) — lightweight, no WebSocket
  library needed
- [ ] Server prints URL to stdout: `Dashboard: http://localhost:8080`
- [ ] Ctrl+C gracefully shuts down the server
- [ ] Server serves files from the output directory (same as static generation)

### Non-Functional

- [ ] File change detection within 500ms of write
- [ ] Page reload completes within 1 second of file change
- [ ] No external dependencies — stdlib `net/http` + polling or `fsnotify`
- [ ] Server handles concurrent requests without race conditions
- [ ] Minimal memory overhead from file watcher (no full repo scan)

## Test Strategy

| Scope | Use when | Examples |
| --- | --- | --- |
| **unit** | File change detection, SSE message format | detect spec.md change, format reload event |
| **integration** | Start server, modify file, verify regeneration | temp repo + server → write file → check output updated |

Default: **integration** for the serve flow, **unit** for watcher logic.

## Acceptance Criteria

- [ ] Given `geremmyas dashboard --serve`, when server starts, then
  `http://localhost:8080` returns index.html with 200 status
- [ ] Given `--port 3000`, when server starts, then listens on port 3000
- [ ] Given `--watch` and a spec.md is modified, when the file is saved, then
  the dashboard regenerates and browser receives reload signal within 2 seconds
- [ ] Given `--watch` and a new spec folder is created, when detected, then the
  new spec appears in the dashboard after regeneration
- [ ] Given the SSE endpoint `/events`, when browser connects, then it receives
  `data: reload` messages on file changes
- [ ] Given Ctrl+C, when pressed, then the server shuts down gracefully (no
  hanging processes)
- [ ] Given `--serve` without `--watch`, when a file changes, then the dashboard
  is NOT regenerated (user must restart to see changes)

## Edge Cases

- Port already in use — print error message suggesting `--port` flag
- File changes in rapid succession (save multiple files) — debounce to one
  regeneration per 300ms window
- Large file change (new family with 20 specs) — full regeneration, not
  incremental
- Spec file deleted — remove from dashboard on next regeneration
- Non-spec file changed in `specs/` (e.g., README.md) — do not trigger
  regeneration
- Watch on macOS/Linux/Windows — use fsnotify for cross-platform support

## Decisions

| Decision | Choice | Reasoning |
|----------|--------|-----------|
| Reload mechanism | Server-Sent Events (SSE) | Simpler than WebSocket; no library needed; `EventSource` is built into browsers |
| File watcher | `fsnotify` library | Cross-platform, low overhead, well-maintained; stdlib has no file watcher |
| Debounce | 300ms window | Fast enough to feel instant; prevents multiple regenerations on batch saves |
| Regeneration scope | Full regeneration on any change | Partial regeneration adds complexity; full regen of 100 specs is < 2s |
| Serve source | Serve from output directory on disk | Simple; no in-memory FS needed |

## Out of Scope

- Hot module replacement (HMR) — full page reload is sufficient
- HTTPS support
- Serving on 0.0.0.0 (localhost only for security)
- Authentication
