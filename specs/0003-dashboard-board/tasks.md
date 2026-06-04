# Tasks: Dashboard board view

## Templates + styles

- [ ] Add board section to `family.html` with four status columns
- [ ] Add spec card partial: number, title, phase badge, progress bar
- [ ] Add deprecated column (hidden by default, toggle controlled)
- [ ] Add view toggle (list/board) to family page header
- [ ] Add phase filter dropdown to board section
- [ ] Add board-specific CSS: column grid, card layout, badge colors, progress
  bar styling, view toggle transitions

## Client-side interaction

- [ ] Create `board.js` with phase filter logic + view toggle
- [ ] Add `data-phase` attribute to each card
- [ ] Implement dropdown change handler: toggle `display:none` on non-matching
  cards
- [ ] Implement list/board toggle: swap CSS class on container
- [ ] Show "No matches" message when filter produces empty result
- [ ] Embed `board.js` via go:embed

## Renderer integration

- [ ] Group family specs by status for board column assignment
- [ ] Handle unknown status values: place in Draft column with warning icon
- [ ] Pass board-grouped data to family template

## Tests

- [ ] Unit test: group specs by status (3 Draft, 2 Approved → correct buckets)
- [ ] Unit test: spec with no tasks.md → card has "No tasks" label
- [ ] Integration test: family page contains board section in output
- [ ] Integration test: phase data attributes present on card elements
- [ ] Integration test: family with 50 specs → board generated without error
