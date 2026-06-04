# Plan: Dashboard git dates and metrics

## Approach

Two new modules in `internal/cli/dashboard/`: git date extractor and metrics
calculator. Chart.js vendored as a minified file in `dashboard_assets/js/`.

## File structure

```
internal/cli/dashboard/
├── gitdates.go             # ExtractGitDates(), cache read/write
├── gitdates_test.go
├── metrics.go              # ComputeMetrics()
└── metrics_test.go

dashboard_assets/
├── js/
│   └── chart.min.js        # Chart.js vendored (~200KB minified)
└── templates/
    ├── metrics.html         # NEW — charts + summary cards
    └── spec.html            # UPDATE — add timeline section
```

## Key decisions

1. **Git log parsing**: Run `git log --all -p --format="COMMIT:%H %aI" -- "specs/*/spec.md"`.
   Parse output line by line: `COMMIT:` lines give SHA + date, `diff --git`
   lines give filename, `+status:` / `-status:` lines give transitions.

2. **Cache structure**:
   ```json
   {
     "last_commit": "abc123def",
     "specs": {
       "0001-rls-pillar": {
         "created_at": "2025-11-10T14:30:00+02:00",
         "status_changes": [
           {"from": "", "to": "Draft", "date": "2025-11-10T14:30:00+02:00"},
           {"from": "Draft", "to": "Approved", "date": "2025-11-15T09:00:00+02:00"},
           {"from": "Approved", "to": "Implemented", "date": "2025-12-01T16:45:00+02:00"}
         ]
       }
     }
   }
   ```
   Store all transitions, derive approved_at/implemented_at by finding the
   latest transition to each target status.

3. **Incremental update**: `git log <last_commit>..HEAD` for new commits only.
   Merge new data into existing cache entries.

4. **Chart.js integration**: Metrics page has a `<script>` block with JSON data
   injected by the Go template. Chart.js reads it on page load. No build step.

## Dependencies

- Spec 0002 (renderer framework + templates)
- `git` binary available on PATH
- Chart.js minified file (download once, vendor in repo)
