---
id: GOT-062
title: '[doc-019 Phase 1] Update column width and border constants in view.go'
status: To Do
assignee:
  - catarina
created_date: '2026-03-29 12:31'
updated_date: '2026-03-31 11:18'
labels:
  - task
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: high
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update column width constants from current values (10, 8, 14, 13, 13) to new values (12, 8, 16, 14, 16) in internal/assets/view.go. Update ColumnSeparator from "  " to "   " (3 spaces). Add new border style constant using lipgloss.DoubleBorder(). Update modal column widths to accommodate longer tickers and decimal places. Update row padding to 0 horizontal, 1 vertical as specified in validation rules. Ensure all width calculations sum to 82 characters (78 data + 4 border).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Column constants updated to new widths (12, 8, 16, 14, 16)
- [ ] #2 Separator updated to 3 spaces
- [ ] #3 Border style constant defined using lipgloss.DoubleBorder()
- [ ] #4 Row padding set to 0 horizontal, 1 vertical
- [ ] #5 Total width calculation verified as 82 characters
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Update the table layout constants in `internal/assets/view.go` to implement the new column width specifications from PRD doc-019. The changes involve:

1. Update column width constants from current values to new values
2. Change column separator from 2 spaces to 3 spaces
3. Add border style constant using `lipgloss.DoubleBorder()`
4. Verify total width calculation equals 82 characters (78 data + 4 border)
5. Update modal column widths to match new asset column widths

**Architecture decisions:**
- Keep existing constant naming convention (`Column*Width` format)
- Use `const` declarations for all new values in same location as existing constants
- Maintain `Padding(0, 1)` as specified in requirements (already implemented)
- Use lipgloss `DoubleBorder()` style which is available in v1.1.0 dependency

**Why this approach:**
- Minimal code changes - only constants and border style
- No functional logic changes - table rendering will automatically use new widths
- Reusable - border style constant can be used throughout component

### 2. Files to Modify

| File | Changes | Reason |
|------|---------|--------|
| `internal/assets/view.go` | Update constants (lines 19-30) | Update column widths and separator per PRD |
| `internal/assets/view_test.go` | Update test expectations (lines 538, 544-546) | Expect 82 characters instead of 74 |

**Files to create:** None

**Files to delete:** None

### 3. Dependencies

**Prerequisites:**
- Lipgloss v1.1.0 already in `go.mod` (confirmed via `go list -m github.com/charmbracelet/lipgloss`)
- No new dependencies required
- Existing functionality preserved - only visual changes

**Blocking issues:** None identified

**Setup steps required:** None - all changes are code modifications

### 4. Code Patterns

**Conventions to follow:**
- Use `const` for all width and separator values
- Maintain current naming pattern: `Column*Width` and `*Separator`
- Keep comments explaining each constant's purpose
- Use lipgloss border styles consistently with existing code
- Maintain `Padding(0, 1)` which is already implemented in render methods

**Implementation details:**
```go
// OLD (current):
const (
    ColumnAssetWidth      = 10   // Asset: 10 characters, left-aligned
    ColumnCountWidth      = 8    // Count: 8 characters, right-aligned
    ColumnSharesWidth     = 14   // Total Shares: 14 characters, right-aligned with 8 decimal places
    ColumnAvgPriceWidth   = 13   // Avg Price: 13 characters, right-aligned with 2 decimal places
    ColumnTotalValueWidth = 13   // Total Value: 13 characters, right-aligned with 2 decimal places
    ColumnSeparator       = "  " // 2 spaces between columns
)

// NEW (required):
const (
    ColumnAssetWidth      = 12   // Asset: 12 characters, left-aligned
    ColumnCountWidth      = 8    // Count: 8 characters, right-aligned
    ColumnSharesWidth     = 16   // Total Shares: 16 characters, right-aligned with 8 decimal places
    ColumnAvgPriceWidth   = 14   // Avg Price: 14 characters, right-aligned with 2 decimal places
    ColumnTotalValueWidth = 16   // Total Value: 16 characters, right-aligned with 2 decimal places
    ColumnSeparator       = "   " // 3 spaces between columns
)
```

**Modal width updates:**
The PRD states "Update modal column widths to accommodate longer tickers and decimal places" but doesn't specify exact values. I recommend:
- `ModalDateWidth`: 12 → 14 (YYYY-MM-DD + buffer)
- `ModalAvgPriceWidth`: 12 → 14 (2-decimal precision with buffer)
- `ModalTotalInvestedWidth`: 14 → 16 (accommodate larger values)
- `ModalEntryCountWidth`: 10 → 12 (accommodate higher counts)
- `ModalDateSeparator`: Update to 3 spaces for consistency

### 5. Testing Strategy

**Unit tests to update:**

1. **`TestTableLayout_WidthIs100Percent`** - Update expected width from 74 to 82
   - Line 544-546: Change `if len(line) != 74` to `if len(line) != 82`
   - Update comment: "Row width = 82 (includes borders)"

2. **`TestTableLayout_HeaderAlignment`** - Verify column positions match new widths
   - Update parsing logic to match new column positions
   - Verify header and data columns still align vertically

3. **`TestAssetsView_RenderWithEntries`** - No changes needed (tests content, not layout)

4. **Add new test:** `TestTableLayout_NewColumnWidths` - Verify each column width
   - Test each constant value matches new specifications
   - Verify total: 12 + 8 + 16 + 14 + 16 + 4*3 = 82

**Edge cases to cover:**
- Empty table (still 82 chars with borders)
- 30 entries (maximum displayed rows)
- Active row highlighting with new widths
- Modal render with updated column widths

### 6. Risks and Considerations

**Blocking issues:**
- None identified. All changes are constant updates.

**Potential pitfalls:**
- **Modal alignment:** Modal uses different column widths and separator. Need to update to maintain visual consistency with main table
- **Terminal compatibility:** 82-character width should fit standard terminals (80-120 chars), but verify with manual testing
- **Test expectations:** All tests checking row width must be updated from 74 to 82

**Trade-offs:**
- No dynamic width detection (fixed 82 chars as specified in PRD)
- Modal separator could be 2 or 3 spaces - choosing 3 spaces for consistency
- Modal column widths not explicitly specified in PRD - will use incremental increases

**Deployment considerations:**
- No database/schema migrations needed (visual changes only)
- No breaking changes to data structures
- No API changes
- Full backward compatibility maintained

**Verification steps:**
1. Run `go test ./internal/assets/...` - all tests should pass
2. Run `go build -o dca ./cmd/dca` - no compiler warnings
3. Run `make fmt` - verify formatting
4. Manual testing: `./dca` - verify table width and borders look correct
5. Test navigation: ↑/↓/Enter/Esc keys work identically
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
