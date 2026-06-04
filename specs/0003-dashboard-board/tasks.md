# Tasks: Dashboard board view

## Templates + styles

- [x] Add board section to `family.html` with four status columns
- [x] Add spec card partial: number, title, phase badge, progress bar
- [x] Add deprecated column (hidden by default, toggle controlled)
- [x] Add view toggle (list/board) to family page header
- [x] Add phase filter dropdown to board section
- [x] Add board-specific CSS: column grid, card layout, badge colors, progress
  bar styling, view toggle transitions

## Client-side interaction

- [x] Create `board.js` with phase filter logic + view toggle
- [x] Add `data-phase` attribute to each card
- [x] Implement dropdown change handler: toggle `display:none` on non-matching
  cards
- [x] Implement list/board toggle: swap CSS class on container
- [x] Show "No matches" message when filter produces empty result
- [x] Embed `board.js` via go:embed

## Renderer integration

- [x] Group family specs by status for board column assignment
- [x] Handle unknown status values: place in Draft column with warning icon
- [x] Pass board-grouped data to family template

## Tests

- [x] Unit test: group specs by status (3 Draft, 2 Approved → correct buckets)
- [x] Unit test: spec with no tasks.md → card has "No tasks" label
- [x] Integration test: family page contains board section in output
- [x] Integration test: phase data attributes present on card elements
- [x] Integration test: family with 50 specs → board generated without error
