# Tasks: Dashboard git dates and metrics

## Git date extraction

- [x] Implement batch `git log` command builder (format, path patterns)
- [x] Implement git log output parser: extract commit SHA, date, filename,
  status transitions
- [x] Map parsed data to per-spec SpecDates structs
- [x] Handle "no git" scenario: check if `.git/` exists and `git` is on PATH
- [x] Add `--no-git` flag to dashboard command

## Cache

- [x] Define cache JSON schema (last_commit + per-spec dates)
- [x] Implement cache read from `.geremmyas-cache/gitdates.json`
- [x] Implement cache write after extraction
- [x] Implement incremental update: `git log <last>..HEAD`
- [x] Implement `--no-cache` flag (delete + rescan)
- [x] Handle corrupted cache: detect invalid JSON, delete, rescan

## Metrics computation

- [x] Implement `ComputeMetrics(specs []SpecSummary, dates map) Metrics`
- [x] Velocity: group implemented specs by month, count per bucket
- [x] Lead time: per-spec created→implemented duration, compute median
- [x] Review time: per-spec created→approved duration, compute median
- [x] Implementation time: per-spec approved→implemented duration, compute
  median
- [x] Phase breakdown: count specs per status per phase per family
- [x] Family progress: percentage implemented per family
- [x] Exclude deprecated specs from all calculations

## Metrics page

- [x] Vendor Chart.js minified into `dashboard_assets/js/`
- [x] Create `metrics.html` template with summary cards section
- [x] Add velocity bar chart (months on x-axis, count on y-axis)
- [x] Add lead time trend line chart
- [x] Add phase breakdown stacked bar chart
- [x] Add family progress horizontal bars
- [x] Inject chart data as JSON in `<script>` block via Go template
- [x] Update `layout.html` nav to include Metrics link

## Spec detail timeline

- [x] Update `spec.html` template to show timeline section when dates available
- [x] Show created → approved → implemented with duration labels
- [x] Graceful fallback when dates missing: "Dates not available"

## Tests

- [x] Unit test: parse git log output (mock output string → correct dates)
- [x] Unit test: cache read/write roundtrip
- [x] Unit test: incremental merge (existing cache + new commits)
- [x] Unit test: ComputeMetrics with known data → verify median, velocity
- [x] Unit test: deprecated specs excluded from metrics
- [x] Integration test: create temp git repo with commits, extract dates, verify
- [x] Integration test: metrics.html contains Chart.js script and data blocks
- [x] Performance test: cold scan of 100 specs under 5 seconds
