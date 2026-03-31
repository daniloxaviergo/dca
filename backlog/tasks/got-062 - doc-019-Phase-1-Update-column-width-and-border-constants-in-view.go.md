---
id: GOT-062
title: '[doc-019 Phase 1] Update column width and border constants in view.go'
status: To Do
assignee:
  - catarina
created_date: '2026-03-29 12:31'
updated_date: '2026-03-31 11:20'
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

1. **Update column width constants**: Change from (10, 8, 14, 13, 13) to (12, 8, 16, 14, 16)
2. **Update separator**: Change from 2 spaces ("  ") to 3 spaces ("   ")
3. **Add border style constant**: Define `TableBorder` using `lipgloss.RoundedBorder()` (not `DoubleBorder()` - PRD requires rounded borders for visual distinction)
4. **Update modal column widths**: Align modal column widths with new table specifications
5. **Verify total width**: Calculate new width = 12 + 8 + 16 + 14 + 16 + (4 separators × 3 spaces) = 78 data + 4 border = 82 characters

**Architecture decisions:**
- Keep existing constant naming convention (`Column*Width`, `Modal*Width`)
- Use `const` declarations for all width and separator values
- Maintain existing `Padding(0, 1)` which is already implemented in render methods
- Use `lipgloss.RoundedBorder()` as specified in PRD (not DoubleBorder - rounded provides visual distinction)
- Update modal separator to 3 spaces for consistency with main table

**Why this approach:**
- Minimal code changes - only constants and one border style change
- No functional logic changes - table rendering will automatically use new widths
- Reusable - border style constant can be used throughout component
- Maintains existing padding style for UI consistency

### 2. Files to Modify

| File | Changes | Reason |
|------|---------|--------|
| `internal/assets/view.go` | Update `ColumnAssetWidth` from 10 to 12, `ColumnSharesWidth` from 14 to 16, `ColumnAvgPriceWidth` from 13 to 14, `ColumnTotalValueWidth` from 13 to 16, `ColumnSeparator` from "  " to "   " | Update column widths per PRD acceptance criteria |
| `internal/assets/view.go` | Update `ModalDateWidth` from 12 to 14, `ModalAvgPriceWidth` from 12 to 14, `ModalTotalInvestedWidth` from 14 to 16, `ModalEntryCountWidth` from 10 to 12, `ModalDateSeparator` from "  " to "   " | Align modal with new table column widths |
| `internal/assets/view_test.go` | Update `TestTableLayout_WidthIs100Percent` to expect 82 instead of 74 | Expect new table width with updated constants |
| `internal/assets/view_test.go` | Update comment in `TestTableLayout_WidthIs100Percent` | Clarify row width is 82 characters |

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
- Maintain current naming pattern: `Column*Width` and `Modal*Width`
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

// Modal width constants
const (
    ModalWidth              = 60
    ModalDateWidth          = 12 // YYYY-MM-DD
    ModalAvgPriceWidth      = 12 // Right-aligned with 2 decimals
    ModalTotalInvestedWidth = 14 // Right-aligned with 2 decimals
    ModalEntryCountWidth    = 10 // Right-aligned
    ModalDateSeparator      = "  "
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

// Modal width constants
const (
    ModalWidth              = 60
    ModalDateWidth          = 14 // YYYY-MM-DD + buffer
    ModalAvgPriceWidth      = 14 // Right-aligned with 2 decimals
    ModalTotalInvestedWidth = 16 // Right-aligned with 2 decimals
    ModalEntryCountWidth    = 12 // Right-aligned
    ModalDateSeparator      = "   " // 3 spaces for consistency
)

// Comment in view.go:
// Total table width: 12 + 8 + 16 + 14 + 16 + (4 separators × 3) = 82 characters (78 data + 4 border)
```

### 5. Testing Strategy

**Unit tests to update:**

1. **`TestTableLayout_WidthIs100Percent`** - Update expected width from 74 to 82
   - Change `if len(line) != 74` to `if len(line) != 82`
   - Update comment: "Row width = 82 (includes borders)"

2. **`TestTableLayout_HeaderAlignment`** - Update column position calculations
   - Update `extractColumnsByPosition` function to match new column positions:
     - col0: Asset, starts at 0, width 12
     - col1: Count, starts at 12 + 3 = 15, width 8
     - col2: Shares, starts at 15 + 8 + 3 = 26, width 16
     - col3: AvgPrice, starts at 26 + 16 + 3 = 45, width 14
     - col4: TotalValue, starts at 45 + 14 + 3 = 62, width 16
   - Update calculation comment

3. **`TestTableLayout_ColumnWidthsMatchConstants`** - No changes needed (uses constants directly)

**New test to add:** `TestTableLayout_TableWidthCalculation` - Verify 82-character width
   - Test total width calculation: 12 + 8 + 16 + 14 + 16 + 12 (separators) + 4 (border) = 82
   - Verify with empty data, 1 entry, 29 entries

**Edge cases to cover:**
- Empty table (still 82 chars with borders)
- 30 entries (maximum displayed rows)
- Active row highlighting with new widths
- Modal render with updated column widths

### 6. Risks and Considerations

**Blocking issues:**
- None identified. All changes are constant updates.

**Potential pitfalls:**
- **Column position calculations:** Tests using `extractColumnsByPosition` must be updated to reflect new column positions
- **Modal alignment:** Modal uses different column widths but should align visually with main table
- **Terminal compatibility:** 82-character width should fit standard terminals (80-120 chars), verify with manual testing

**Trade-offs:**
- No dynamic width detection (fixed 82 chars as specified in PRD)
- Modal separator uses 3 spaces for consistency with main table
- Modal column widths not explicitly specified in PRD - will use incremental increases

**Deployment considerations:**
- No database/schema migrations needed (visual changes only)
- No breaking changes to data structures
- No API changes
- Full backward compatibility maintained

**Verification steps:**
1. Run `go test ./internal/assets/... -v` - all tests should pass
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
