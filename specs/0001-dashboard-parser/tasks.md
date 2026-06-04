# Tasks: Dashboard spec parser

## Foundation

- [ ] Create `internal/cli/dashboard/` package directory
- [ ] Define types in `types.go`: DashboardData, Family, Phase, SpecSummary,
  TaskStats, PRD, Bugfix, Warning

## Spec parser

- [ ] Extend `frontmatter.go` to return body text after frontmatter delimiter
- [ ] Implement `ScanSpecDirs()` — find all `NNNN-<slug>/` directories
- [ ] Implement `ParseSpecFrontmatter()` — extract typed fields from frontmatter
  map
- [ ] Implement `ParseTaskStats()` — regex count of `[ ]`, `[~]`, `[x]`
  checkboxes
- [ ] Implement `ScanSpecs()` — orchestrate: scan dirs → parse each → group by
  family/phase → sort
- [ ] Handle missing frontmatter: place in "Ungrouped" family, emit warning
- [ ] Handle malformed folders: skip with warning

## PRD and bugfix parsers

- [ ] Implement `ScanPRDs()` — scan `docs/prds/*.md`, extract title + date
- [ ] Implement `ScanBugfixes()` — scan `docs/bugfixes/*.md`, extract title +
  date
- [ ] Implement `ParseDateFromFilename()` — extract YYYY-MM-DD from filename
  pattern

## Tests

- [ ] Unit test: ParseSpecFrontmatter with valid frontmatter
- [ ] Unit test: ParseSpecFrontmatter with missing/empty frontmatter
- [ ] Unit test: ParseTaskStats with mixed checkboxes including indented
- [ ] Unit test: ParseDateFromFilename with valid and invalid filenames
- [ ] Integration test: ScanSpecs with temp directory tree (multiple families,
  phases, edge cases)
- [ ] Integration test: ScanPRDs with temp directory
- [ ] Performance test: ScanSpecs with 100 generated spec folders < 1s
