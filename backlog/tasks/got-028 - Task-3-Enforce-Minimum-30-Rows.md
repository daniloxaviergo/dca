---
id: GOT-028
title: 'Task 3: Enforce Minimum 30 Rows'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 22:11'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add row padding to maintain exactly 30 rows in the Assets View table regardless of data volume.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 When data rows < 30: pad with empty rows to reach 30
- [ ] #2 When data rows = 30: display all rows without truncation
- [ ] #3 Empty rows use same styling as data rows but with empty values
- [ ] #4 Test verifies exactly 30 rows rendered with 5 assets
- [ ] #5 Test verifies exactly 30 rows rendered with 25 assets
- [ ] #6 go fmt applied
- [ ] #7 go build succeeds
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The implementation will enforce a minimum of 30 rows in the Assets View table by padding with empty rows when data rows are fewer than 30. The approach:

- **Row padding logic**: In `renderTable()`, calculate how many empty rows are needed (`max(0, 30 - len(a.Entries))`)
- **Empty row styling**: Create a new `renderEmptyDataRow()` function that uses the same styling as data rows but with empty/zero values
- **Selection handling**: Empty rows should be navigable but not selectable for interaction (cursor stops at data rows or wraps)
- **Display**: Only show header + data + padding; do not show footer when table is padded (to avoid visual clutter)

**Why this approach:**
- Simple and predictable: always exactly 30 rows for consistent UI layout
- Reuses existing styling functions to maintain visual consistency
- No changes to data model or aggregation logic needed

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/assets/view.go` | Modify | Add row padding logic to `renderTable()`, create `renderEmptyDataRow()`, update navigation to handle empty rows |
| `internal/assets/view_test.go` | Modify | Add tests for 30-row enforcement with 5 and 25 assets |

### 3. Dependencies

- None - this is a self-contained UI enhancement
- Depends on existing column width constants and styling in `internal/assets/view.go`

### 4. Code Patterns

Follow existing patterns in `view.go`:
- Use `lipgloss.NewStyle().Padding(0, 1).Render(rowStr)` for empty row styling
- Match `renderDataRow()` field formatting but with empty/zeros
- Preserve wrap-around navigation logic in `handleUp()`/`handleDown()`
- Use same `ColumnSeparator` ("  ") between columns

### 5. Testing Strategy

Add table-driven tests in `view_test.go`:
- **Test 5 assets**: Verify 30 rows (5 data + 25 empty), all columns visible
- **Test 25 assets**: Verify 30 rows (25 data + 5 empty), all columns visible
- **Test navigation**: Verify cursor can move through empty rows
- **Test row count**: Parse output to count actual rendered rows

Edge cases:
- 0 assets (30 empty rows)
- Exactly 30 assets (no padding needed)
- More than 30 assets (no padding, display all)

### 6. Risks and Considerations

- **Selection behavior**: Empty rows should be navigable but may need special handling if user tries to interact with them (e.g., pressing 'c' on empty row)
- **Footer visibility**: Consider whether footer should display when table is padded; may create visual imbalance
- **Performance**: Minimal impact - only affects rendering, not data processing
- **Testing complexity**: Counting rows in rendered output requires parsing lipgloss-styled strings; may need to extract row count from unstyled content
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
