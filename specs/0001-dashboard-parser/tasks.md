# Tasks: Dashboard spec parser

## Foundation

- [x] Create `internal/cli/dashboard/` package directory
- [x] Define types in `types.go`: DashboardData, Family, Phase, SpecSummary,
  TaskStats, PRD, Bugfix, Warning

## Spec parser

- [x] Extend `frontmatter.go` to return body text after frontmatter delimiter
- [x] Implement `ScanSpecDirs()` — find all `NNNN-<slug>/` directories
- [x] Implement `ParseSpecFrontmatter()` — extract typed fields from frontmatter
  map
- [x] Implement `ParseTaskStats()` — regex count of `[ ]`, `[~]`, `[x]`
  checkboxes
- [x] Implement `ScanSpecs()` — orchestrate: scan dirs → parse each → group by
  family/phase → sort
- [x] Handle missing frontmatter: place in "Ungrouped" family, emit warning
- [x] Handle malformed folders: skip with warning

## PRD and bugfix parsers

- [x] Implement `ScanPRDs()` — scan `docs/prds/*.md`, extract title + date
- [x] Implement `ScanBugfixes()` — scan `docs/bugfixes/*.md`, extract title +
  date
- [x] Implement `ParseDateFromFilename()` — extract YYYY-MM-DD from filename
  pattern

## Tests

- [x] Unit test: ParseSpecFrontmatter with valid frontmatter
- [x] Unit test: ParseSpecFrontmatter with missing/empty frontmatter
- [x] Unit test: ParseTaskStats with mixed checkboxes including indented
- [x] Unit test: ParseDateFromFilename with valid and invalid filenames
- [x] Integration test: ScanSpecs with temp directory tree (multiple families,
  phases, edge cases)
- [x] Integration test: ScanPRDs with temp directory
- [x] Performance test: ScanSpecs with 100 generated spec folders < 1s
