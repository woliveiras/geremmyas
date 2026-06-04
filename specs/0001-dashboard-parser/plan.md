# Plan: Dashboard spec parser

## Approach

Add a new `internal/cli/dashboard/` package with the parser as the first
module. This keeps dashboard logic isolated from existing CLI commands while
sharing the existing frontmatter utilities.

## File structure

```
internal/cli/dashboard/
├── parser.go       # ScanSpecs(), ScanPRDs(), ScanBugfixes()
├── parser_test.go  # unit + integration tests
└── types.go        # DashboardData, Family, Phase, SpecSummary, TaskStats, PRD, Bugfix
```

## Key decisions

1. **New package vs inline in cli.go**: New package. Dashboard is complex enough
   to warrant its own namespace. Other commands do not need these types.

2. **Reuse frontmatter.go**: The existing parser extracts key-value pairs from
   `---` delimited YAML. Extend it to return both frontmatter map and remaining
   body string. Currently it only returns the map.

3. **Error handling**: Per-file errors are collected as warnings in a
   `[]Warning` slice returned alongside the data. Callers decide whether to
   print warnings or fail.

4. **Sorting**: Specs sorted by number (int) within each phase. Families sorted
   alphabetically. Phases sorted by number.

## Dependencies

- `internal/cli/frontmatter.go` — extend to return body
- `os`, `path/filepath`, `regexp`, `strconv`, `strings`, `sort` — stdlib only
