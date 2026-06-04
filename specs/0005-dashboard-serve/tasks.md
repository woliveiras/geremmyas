# Tasks: Dashboard serve and watch mode

## Server

- [ ] Implement `StartServer(outputDir string, port int) error`
- [ ] Serve static files from output directory
- [ ] Implement SSE endpoint at `/events`
- [ ] Add SSE client script to `layout.html` (auto-reload on message)
- [ ] Graceful shutdown on Ctrl+C (signal.NotifyContext)
- [ ] Print server URL to stdout on start
- [ ] Return clear error when port is in use

## File watcher

- [ ] Add `fsnotify` dependency
- [ ] Implement `WatchDirs(dirs []string) (<-chan Event, error)`
- [ ] Watch `specs/` and `docs/` directories recursively
- [ ] Filter events: only `.md` file changes trigger rebuild
- [ ] Implement debounce (300ms timer reset on each event)

## Integration

- [ ] Add `--serve` flag to dashboard command
- [ ] Add `--port` flag (default 8080)
- [ ] Add `--watch` flag (implies --serve)
- [ ] Wire: full dashboard generate (0002 + 0004 when enabled) → start server →
  (if watch) start watcher → on change: rerun full generate + SSE ping

## Tests

- [ ] Unit test: debounce logic (rapid events → single callback)
- [ ] Unit test: SSE message format
- [ ] Integration test: start server, GET /, verify 200 + HTML content
- [ ] Integration test: start server with --watch, modify spec file, verify
  output directory updated (including metrics.html when git metrics enabled)
- [ ] Integration test: port conflict → clear error message
- [ ] Integration test: Ctrl+C → server stops cleanly
