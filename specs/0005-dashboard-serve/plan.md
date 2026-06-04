# Plan: Dashboard serve and watch mode

## Approach

Add `server.go` and `watcher.go` to the dashboard package. The serve command
generates the dashboard first (reusing the existing pipeline), then starts an
HTTP server. The watcher monitors `specs/` and `docs/` for changes and triggers
full regeneration with an SSE ping.

## File structure

```
internal/cli/dashboard/
├── server.go           # StartServer(outputDir, port), SSE handler
├── server_test.go
├── watcher.go          # WatchAndRebuild(dirs, onChange)
└── watcher_test.go
```

## Key decisions

1. **SSE implementation**: A simple `/events` endpoint that keeps the
   connection open. On regeneration, write `data: reload\n\n` to all connected
   clients. The HTML layout includes a small `<script>` that connects to
   `/events` and calls `location.reload()` on message.

2. **fsnotify**: New dependency. Lightweight, well-maintained, handles
   cross-platform differences. Alternative (polling) is wasteful.

3. **Debounce**: Collect events for 300ms after the first event, then trigger
   one regeneration. Use a `time.Timer` that resets on each event.

4. **Graceful shutdown**: `signal.NotifyContext` with `os.Interrupt`. Server
   uses `http.Server.Shutdown()`.

## Dependencies

- Spec 0002 (renderer pipeline)
- New dep: `github.com/fsnotify/fsnotify` (file watcher)
