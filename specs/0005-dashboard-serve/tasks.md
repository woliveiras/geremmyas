# Tasks: Dashboard serve and watch mode

## Server

- [x] Implement `StartServer(outputDir string, port int) error`
- [x] Serve static files from output directory
- [x] Implement SSE endpoint at `/events`
- [x] Add SSE client script to `layout.html` (auto-reload on message)
- [x] Graceful shutdown on Ctrl+C (signal.NotifyContext)
- [x] Print server URL to stdout on start
- [x] Return clear error when port is in use

## File watcher

- [x] Add `fsnotify` dependency
- [x] Implement `WatchDirs(dirs []string) (<-chan Event, error)`
- [x] Watch `specs/` and `docs/` directories recursively
- [x] Filter events: only `.md` file changes trigger rebuild
- [x] Implement debounce (300ms timer reset on each event)

## Integration

- [x] Add `--serve` flag to dashboard command
- [x] Add `--port` flag (default 8080)
- [x] Add `--watch` flag (implies --serve)
- [x] Wire: full dashboard generate (0002 + 0004 when enabled) → start server →
  (if watch) start watcher → on change: rerun full generate + SSE ping

## Tests

- [x] Unit test: debounce logic (rapid events → single callback)
- [x] Unit test: SSE message format
- [x] Integration test: start server, GET /, verify 200 + HTML content
- [x] Integration test: start server with --watch, modify spec file, verify
  output directory updated (including metrics.html when git metrics enabled)
- [x] Integration test: port conflict → clear error message
- [x] Integration test: Ctrl+C → server stops cleanly
