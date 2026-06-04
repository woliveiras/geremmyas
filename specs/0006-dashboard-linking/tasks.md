# Tasks: Dashboard artifact linking and dependency visualization

## Dependency graph

- [ ] Implement `BuildDependencyGraph(specs []SpecSummary) DependencyGraph`
- [ ] Implement `FindBlockedBy(specNum int, graph) []SpecRef` — deps not
  Implemented
- [ ] Implement `FindBlocking(specNum int, graph) []SpecRef` — reverse deps
- [ ] Detect circular dependencies, collect as warnings
- [ ] Handle missing spec numbers in depends_on (label as "missing")

## PRD and bugfix pages

- [ ] Create `prd_index.html` template (list of PRDs: date, title, spec count)
- [ ] Create `prd.html` template (rendered PRD + "Related specs" section)
- [ ] Create `bugfix_index.html` template (list of bugfixes: date, title)
- [ ] Create `bugfix.html` template (rendered bugfix content)
- [ ] Implement PRD→spec reverse link map (origin field matching)
- [ ] Add `RenderPRDs()` and `RenderBugfixes()` to renderer

## Template updates

- [ ] Update `spec.html`: add "Blocked by" and "Blocks" sections
- [ ] Update `spec.html`: make origin field a clickable link to PRD page
- [ ] Update `spec.html`: show "(not found)" for unresolved origin paths
- [ ] Update `family.html`: show blocked indicator on spec rows
- [ ] Update `layout.html`: add PRDs and Bugfixes nav links

## Renderer integration

- [ ] Wire dependency graph computation into `RenderDashboard()` pipeline
- [ ] Pass dependency info to spec and family templates
- [ ] Generate PRD and bugfix pages in output directory

## Tests

- [ ] Unit test: BuildDependencyGraph with known deps → correct adjacency
- [ ] Unit test: FindBlockedBy when some deps not Implemented
- [ ] Unit test: FindBlocking (reverse lookup)
- [ ] Unit test: circular dependency detection
- [ ] Unit test: PRD reverse link map (3 specs reference same PRD)
- [ ] Integration test: generate dashboard with PRDs + deps → verify links in
  HTML output
- [ ] Integration test: missing origin path → "(not found)" in output
- [ ] Integration test: no PRDs directory → empty state page
