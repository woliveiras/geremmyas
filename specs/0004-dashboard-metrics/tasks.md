# Tasks: Dashboard git dates and metrics

## Git date extraction

- [ ] Implement batch `git log` command builder (format, path patterns)
- [ ] Implement git log output parser: extract commit SHA, date, filename,
  status transitions
- [ ] Map parsed data to per-spec SpecDates structs
- [ ] Handle "no git" scenario: check if `.git/` exists and `git` is on PATH
- [ ] Add `--no-git` flag to dashboard command

## Cache

- [ ] Define cache JSON schema (last_commit + per-spec dates)
- [ ] Implement cache read from `.geremmyas-cache/gitdates.json`
- [ ] Implement cache write after extraction
- [ ] Implement incremental update: `git log <last>..HEAD`
- [ ] Implement `--no-cache` flag (delete + rescan)
- [ ] Handle corrupted cache: detect invalid JSON, delete, rescan

## Metrics computation

- [ ] Implement `ComputeMetrics(specs []SpecSummary, dates map) Metrics`
- [ ] Velocity: group implemented specs by month, count per bucket
- [ ] Lead time: per-spec created→implemented duration, compute median
- [ ] Review time: per-spec created→approved duration, compute median
- [ ] Implementation time: per-spec approved→implemented duration, compute
  median
- [ ] Phase breakdown: count specs per status per phase per family
- [ ] Family progress: percentage implemented per family
- [ ] Exclude deprecated specs from all calculations

## Metrics page

- [ ] Vendor Chart.js minified into `dashboard_assets/js/`
- [ ] Create `metrics.html` template with summary cards section
- [ ] Add velocity bar chart (months on x-axis, count on y-axis)
- [ ] Add lead time trend line chart
- [ ] Add phase breakdown stacked bar chart
- [ ] Add family progress horizontal bars
- [ ] Inject chart data as JSON in `<script>` block via Go template
- [ ] Update `layout.html` nav to include Metrics link

## Spec detail timeline

- [ ] Update `spec.html` template to show timeline section when dates available
- [ ] Show created → approved → implemented with duration labels
- [ ] Graceful fallback when dates missing: "Dates not available"

## Tests

- [ ] Unit test: parse git log output (mock output string → correct dates)
- [ ] Unit test: cache read/write roundtrip
- [ ] Unit test: incremental merge (existing cache + new commits)
- [ ] Unit test: ComputeMetrics with known data → verify median, velocity
- [ ] Unit test: deprecated specs excluded from metrics
- [ ] Integration test: create temp git repo with commits, extract dates, verify
- [ ] Integration test: metrics.html contains Chart.js script and data blocks
- [ ] Performance test: cold scan of 100 specs under 5 seconds
