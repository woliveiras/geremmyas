# Plan: Dashboard board view

## Approach

Instead of a separate board page, add a kanban board section to the existing
family detail page (from 0002). Users toggle between list view (default) and
board view. This keeps the board scoped to one family at a time.

## File changes

```
dashboard_assets/templates/
├── family.html             # UPDATE — add board section with view toggle
└── layout.html             # no change (no new nav link needed)

internal/cli/dashboard/
├── renderer.go             # UPDATE — pass board-grouped data to family template
└── renderer_test.go        # UPDATE — board tests

dashboard_assets/css/
└── style.css               # UPDATE — board columns, card styles, filter, toggle

dashboard_assets/js/
└── board.js                # NEW — phase filter + list/board toggle (vanilla JS)
```

## Key decisions

1. **Board inside family page, not separate**: Avoids duplicating family
   context. User sees family header, phase info, AND the board in one place.
   Toggle switches the content area.

2. **Data attributes for filtering**: Each card gets `data-phase` attribute.
   JS toggles `display:none` based on dropdown. Family filtering is implicit
   (board is already scoped to one family).

3. **View toggle**: CSS class on a container. `data-view="list"` shows the
   phase-grouped table (existing). `data-view="board"` shows the kanban
   columns. JS toggles the class.

4. **Progress bar**: `<progress value="4" max="6">` with CSS styling.

## Dependencies

- Spec 0002 (family detail page + renderer)
