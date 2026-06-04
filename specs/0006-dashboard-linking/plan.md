# Plan: Dashboard artifact linking and dependency visualization

## Approach

Extend the parser (0001) to build a dependency graph and reverse-link maps.
Add new templates for PRD and bugfix detail pages. Update existing spec and
family templates to show dependency and link information.

## File structure

```
internal/cli/dashboard/
├── deps.go                 # BuildDependencyGraph(), FindBlocked(), FindBlocking()
├── deps_test.go
├── types.go                # UPDATE — add DependencyInfo, PRDDetail, BugfixDetail

dashboard_assets/templates/
├── prd.html                # NEW — PRD detail page
├── prd_index.html          # NEW — PRDs list page
├── bugfix.html             # NEW — bugfix detail page
├── bugfix_index.html       # NEW — bugfixes list page
├── spec.html               # UPDATE — add blocked-by/blocks sections, origin link
├── family.html             # UPDATE — add blocked specs indicator
└── layout.html             # UPDATE — add PRDs and Bugfixes nav links
```

## Key decisions

1. **Dependency graph**: Simple adjacency list stored as `map[int][]int`. One
   pass over all specs builds it. Reverse map (blocks) is the transpose.

2. **Blocked detection**: A spec is "blocked" if any of its `depends_on` specs
   has status != "Implemented". Straightforward status check.

3. **PRD rendering**: Reuse the same markdown → HTML pipeline from spec pages.
   PRD pages are simpler (no sidebar metadata, just rendered content + reverse
   links).

## Dependencies

- Spec 0001 (parser, types)
- Spec 0002 (renderer pipeline, templates)
- Spec 0004 (git dates on spec detail page — dependency info adds to the same
  page)
