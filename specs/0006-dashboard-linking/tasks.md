# Tasks: Dashboard artifact linking and dependency visualization

## Dependency graph

- [x] Implement `BuildDependencyGraph(specs []SpecSummary) DependencyGraph`
- [x] Implement `FindBlockedBy(specNum int, graph) []SpecRef` — deps not
  Implemented
- [x] Implement `FindBlocking(specNum int, graph) []SpecRef` — reverse deps
- [x] Detect circular dependencies, collect as warnings
- [x] Handle missing spec numbers in depends_on (label as "missing")

## PRD and bugfix pages

- [x] Create `prd_index.html` template (list of PRDs: date, title, spec count)
- [x] Create `prd.html` template (rendered PRD + "Related specs" section)
- [x] Create `bugfix_index.html` template (list of bugfixes: date, title)
- [x] Create `bugfix.html` template (rendered bugfix content)
- [x] Implement PRD→spec reverse link map (origin field matching)
- [x] Add `RenderPRDs()` and `RenderBugfixes()` to renderer

## Template updates

- [x] Update `spec.html`: add "Blocked by" and "Blocks" sections
- [x] Update `spec.html`: make origin field a clickable link to PRD page
- [x] Update `spec.html`: show "(not found)" for unresolved origin paths
- [x] Update `family.html`: show blocked indicator on spec rows
- [x] Update `layout.html`: add PRDs and Bugfixes nav links

## Renderer integration

- [x] Wire dependency graph computation into `RenderDashboard()` pipeline
- [x] Pass dependency info to spec and family templates
- [x] Generate PRD and bugfix pages in output directory

## Tests

- [x] Unit test: BuildDependencyGraph with known deps → correct adjacency
- [x] Unit test: FindBlockedBy when some deps not Implemented
- [x] Unit test: FindBlocking (reverse lookup)
- [x] Unit test: circular dependency detection
- [x] Unit test: PRD reverse link map (3 specs reference same PRD)
- [x] Integration test: generate dashboard with PRDs + deps → verify links in
  HTML output
- [x] Integration test: missing origin path → "(not found)" in output
- [x] Integration test: no PRDs directory → empty state page
