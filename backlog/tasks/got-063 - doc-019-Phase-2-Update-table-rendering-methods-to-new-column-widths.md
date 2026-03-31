---
id: GOT-063
title: '[doc-019 Phase 2] Update table rendering methods to new column widths'
status: To Do
assignee:
  - thomas
created_date: '2026-03-29 12:31'
updated_date: '2026-03-31 12:22'
labels:
  - task
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: high
ordinal: 7000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Modify renderHeaderRow(), renderDataRow(), and renderEmptyDataRow() methods in internal/assets/view.go to use new column width constants (12, 8, 16, 14, 16). Update column separator usage from ColumnSeparator (2 spaces) to new 3-space separator. Modify renderTable() to use double-line rounded borders by replacing lipgloss.RoundedBorder() with lipgloss.NewStyle().Border(lipgloss.DoubleBorder()). Ensure all data formatting maintains decimal precision (8 decimals for shares, 2 decimals for prices/values).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 renderHeaderRow() updated with new column widths
- [ ] #2 renderDataRow() updated with new column widths
- [ ] #3 renderEmptyDataRow() updated with new column widths
- [ ] #4 renderTable() uses lipgloss.DoubleBorder() for double-line borders
- [ ] #5 Data formatting maintains 8 decimal precision for shares and 2 decimals for prices/values
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

**Phase 2: Update Table Rendering Methods**

This phase updates the table rendering logic to use the new column width constants defined in Phase 1 (which were already set to 12, 8, 16, 14, 16 with 3-space separators). The key changes are:

1. **Update `renderTable()` border style**: Replace `lipgloss.RoundedBorder()` with `lipgloss.NewStyle().Border(lipgloss.DoubleBorder())` to achieve double-line rounded borders as specified in PRD doc-019.

2. **Verify rendering methods use new constants**: Confirm `renderHeaderRow()`, `renderDataRow()`, and `renderEmptyDataRow()` already use the column width constants from Phase 1 (which they do).

3. **Maintain decimal precision**: Ensure all formatting maintains 8 decimal places for shares (using `%*.8f`) and 2 decimal places for prices/values (using `%*.2f`).

**Architecture decisions:**
- Keep existing constant-based approach - rendering methods use constants that were already updated in Phase 1
- Use `lipgloss.NewStyle().Border(lipgloss.DoubleBorder())` for double-line rounded borders instead of `TableBorder = lipgloss.RoundedBorder()`
- Preserve existing row rendering logic - only border style changes
- Maintain active row highlighting with same coloring (cyan background #63)

**Why this approach:**
- Minimal changes - only border style modification needed since constants were already set in Phase 1
- No functional logic changes - only visual border style improvement
- Consistent with PRD specification for double-line rounded borders
- All existing rendering methods already use the new column width constants

### 2. Files to Modify

| File | Changes | Reason |
|------|---------|--|
| `internal/assets/view.go` | Change `renderTable()` to use `lipgloss.NewStyle().Border(lipgloss.DoubleBorder())` instead of `TableBorder` | Double-line rounded borders per PRD acceptance criteria |
| `internal/assets/view.go` | Add inline `DoubleBorder` style definition or update `TableBorder` to use double-line style | Consistent border styling throughout component |

**Files to create:** None

**Files to delete:** None

### 3. Dependencies

**Prerequisites:**
- Phase 1 (GOT-062) must be complete - column width constants already updated to (12, 8, 16, 14, 16)
- Lipgloss v1.1.0 already in `go.mod` - supports `DoubleBorder()`
- No new dependencies required
- Existing functionality preserved - only visual changes

**Blocking issues:** None identified

**Setup steps required:** None - all changes are code modifications

### 4. Code Patterns

**Conventions to follow:**
- Use existing constant-based approach for column widths (already implemented in Phase 1)
- Apply lipgloss styles consistently with existing code patterns
- Preserve active row highlighting logic with bright cyan background (#63)
- Maintain `Padding(0)` for row styling to achieve correct 80-character width

**Implementation details:**
- **Current border style**: `TableBorder = lipgloss.RoundedBorder()` (single-line rounded)
- **New border style**: `lipgloss.NewStyle().Border(lipgloss.DoubleBorder())` (double-line rounded)
- **Data formatting already correct**:
  - Shares: `%*.8f` (8 decimal places)
  - Prices/Values: `%*.2f` (2 decimal places)
  - Counts: `%*d` (integer)
  - Asset tickers: `%-*s` (left-aligned text)

**Changes to make in `renderTable()`:**
```go
// Change from:
tableStyle := lipgloss.NewStyle().
    Padding(0).
    Border(TableBorder).
    BorderForeground(lipgloss.Color("240"))

// To:
doubleBorder := lipgloss.NewStyle().Border(lipgloss.DoubleBorder())
tableStyle := lipgloss.NewStyle().
    Padding(0).
    Border(doubleBorder).
    BorderForeground(lipgloss.Color("240"))
```

### 5. Testing Strategy

**Existing tests to verify:**
1. **TestTableLayout_WidthIs100Percent**: Verify table width remains correct (80 characters)
2. **TestTableLayout_HeaderAlignment**: Verify column alignment maintained
3. **TestAssetsView_Render**: Verify overall rendering works with new borders

**Additional tests for Phase 2:**
1. **TestTableLayout_DoubleBorder**: Verify double-line borders are applied
2. **TestTableLayout_RowCount**: Verify exactly 30 rows rendered (1 header + 29 data/empty)
3. **TestAssetsView_Render_WithEmptyData**: Verify border rendering with no assets
4. **TestAssetsView_Render_WithMaxData**: Verify border rendering with 29 data rows

**Edge cases covered:**
- Empty table (0 entries + border)
- Minimum data (1 entry + border)
- Maximum data (29 entries + border)
- Active row highlighting with new border style
- Modal rendering (separate border style - should remain single-line rounded)

**Testing commands:**
```bash
# Run all tests with verbose output
go test -v ./internal/assets/...

# Generate coverage report
go test -cover ./internal/assets/...

# Build and verify no compiler warnings
go build -o dca ./cmd/dca

# Manual visual verification
./dca
```

### 6. Risks and Considerations

**Blocking issues:**
- None identified. All changes are visual updates.

**Potential pitfalls:**
- **Border rendering in terminals**: Double-line borders should render correctly in standard terminals; verify with common terminal emulators (Linux, macOS, Windows terminals)
- **Border color consistency**: Ensure border foreground color (#240) provides good contrast with content
- **Modal border style**: Modal uses `TableBorder` - decide if it should also use double-line or remain single-line (PRD suggests only main table changes)

**Design considerations:**
- **Modal border**: PRD only specifies main table border changes; modal currently uses `TableBorder`. Consider keeping modal as single-line or updating both.
- **Active row styling**: Verify highlight (cyan background #63) is visible against double-line borders
- **Terminal compatibility**: Double-line borders are supported by lipgloss and主流 terminals; should work as expected

**Deployment considerations:**
- No database/schema migrations needed (visual changes only)
- No breaking changes to data structures
- No API changes
- Full backward compatibility maintained
- No migration needed for existing data files

**Verification steps:**
1. Run `go test ./... -v` - verify all tests pass
2. Run `go build -o dca ./cmd/dca` - verify no compiler warnings
3. Run `make fmt` - verify code formatting
4. Manual testing: Launch `./dca` and verify:
   - Double-line rounded borders render correctly
   - Column alignment maintained with new widths
   - Active row highlighting visible
   - Exactly 30 rows displayed
   - Decimal precision correct (8 for shares, 2 for prices/values)

### 7. Acceptance Criteria Alignment

| Criteria | Status | Verification |
|----------|--------|--------------|
| #1 `renderHeaderRow()` updated with new column widths | ✅ Already done in Phase 1 | Constants set to (12, 8, 16, 14, 16) |
| #2 `renderDataRow()` updated with new column widths | ✅ Already done in Phase 1 | Uses `Column*Width` constants |
| #3 `renderEmptyDataRow()` updated with new column widths | ✅ Already done in Phase 1 | Uses `Column*Width` constants |
| #4 `renderTable()` uses `lipgloss.DoubleBorder()` | 🔜 **Phase 2 task** | Change `TableBorder` usage |
| #5 Data formatting maintains 8 decimals for shares, 2 for prices | ✅ Already correct | `%*.8f` and `%*.2f` format specifiers |

### 8. Implementation Checklist

- [ ] Update `renderTable()` to use `lipgloss.DoubleBorder()` for double-line borders
- [ ] Verify `renderHeaderRow()` uses new column width constants
- [ ] Verify `renderDataRow()` uses new column width constants
- [ ] Verify `renderEmptyDataRow()` uses new column width constants
- [ ] Run `go test ./internal/assets/... -v` - all tests pass
- [ ] Run `go build -o dca ./cmd/dca` - no compiler warnings
- [ ] Run `make fmt` - code formatted
- [ ] Manual verification: Launch app and verify double-line borders render correctly
- [ ] Update task status to "Done" and move to next phase (GOT-064)
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
