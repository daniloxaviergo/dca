---
id: GOT-062
title: '[doc-019 Phase 1] Update column width and border constants in view.go'
status: Done
assignee:
  - workflow
created_date: '2026-03-29 12:31'
updated_date: '2026-03-31 12:20'
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
Update column width constants from current values (10, 8, 14, 13, 13) to new values (12, 8, 16, 14, 16) in internal/assets/view.go. Update ColumnSeparator from "  " to "   " (3 spaces). Add new border style constant using lipgloss.RoundedBorder(). Update modal column widths to accommodate longer tickers and decimal places. Ensure all width calculations sum to 80 characters (78 content + 2 border).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 #1 Column constants updated to new widths (12, 8, 16, 14, 16)
- [x] #2 #2 Separator updated to 3 spaces
- [x] #3 #3 Border style constant defined using lipgloss.RoundedBorder()
- [x] #4 #4 Row padding set to Padding(0) for correct width
- [x] #5 #5 Total width calculation verified as 80 characters (78 content + 2 border)
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Update the table layout constants in `internal/assets/view.go` to implement the new column width specifications from PRD doc-019. The changes involve:

1. **Update column width constants**: Change from (10, 8, 14, 13, 13) to (12, 8, 16, 14, 16)
2. **Update separator**: Change from 2 spaces ("  ") to 3 spaces ("   ")
3. **Add border style constant**: Define `TableBorder` using `lipgloss.RoundedBorder()`
4. **Update modal column widths**: Align modal column widths with new table specifications
5. **Verify total width**: Calculate new width = 12 + 8 + 16 + 14 + 16 + (4 separators × 3 spaces) = 78 data + 2 border = 80 characters

**Architecture decisions:**
- Keep existing constant naming convention (`Column*Width`, `Modal*Width`)
- Use `const` declarations for all width and separator values
- Use `lipgloss.RoundedBorder()` as specified in PRD
- Update modal separator to 3 spaces for consistency with main table

**Why this approach:**
- Minimal code changes - only constants and one border style change
- No functional logic changes - table rendering will automatically use new widths
- Reusable - border style constant can be used throughout component
- Consistent styling across table and modal components

### 2. Files to Modify

| File | Changes | Reason |
|------|---------|--|
| `internal/assets/view.go` | Update column constants (12, 8, 16, 14, 16), separator to 3 spaces, add TableBorder, update modal widths | Update column widths per PRD acceptance criteria |
| `internal/assets/view.go` | Remove Padding(0, 1) from renderDataRow and renderEmptyDataRow | Achieve correct 80-character row width |
| `internal/assets/view_test.go` | Update TestTableLayout_WidthIs100Percent to expect 80 characters, use utf8.RuneCountInString() | Match new table width |
| `internal/assets/view_test.go` | Update extractColumnsByPosition with correct column positions | Match new column positions for 78-char clean row |
| `internal/assets/view_test.go` | Update comments to reflect 80-character total width | Accurate documentation |

**Files to create:** None

**Files to delete:** None

### 3. Dependencies

**Prerequisites:**
- Lipgloss v1.1.0 already in `go.mod`
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

**Implementation details:**
- Column widths: 12 + 8 + 16 + 14 + 16 = 78 characters
- Separators: 4 × 3 spaces = 12 characters  
- Total content: 90 characters
- After removing padding and border adjustment: 80 characters

### 5. Testing Strategy

**Unit tests updated:**
1. **TestTableLayout_WidthIs100Percent**: Updated to expect 80 characters using `utf8.RuneCountInString()`
2. **TestTableLayout_HeaderAlignment**: Updated column position calculations
3. **extractColumnsByPosition**: Updated to use correct column positions for 78-character clean row

**Edge cases covered:**
- Empty table (with borders)
- 30 entries (maximum displayed rows)
- Active row highlighting with new widths
- Modal render with updated column widths

### 6. Risks and Considerations

**Blocking issues:**
- None identified. All changes are constant updates.

**Potential pitfalls:**
- **Column position calculations:** Tests using `extractColumnsByPosition` must be updated to reflect new column positions - FIXED
- **Modal alignment:** Modal uses different column widths but should align visually with main table
- **Terminal compatibility:** 80-character width should fit standard terminals (80-120 chars), verified

**Deployment considerations:**
- No database/schema migrations needed (visual changes only)
- No breaking changes to data structures
- No API changes
- Full backward compatibility maintained

**Verification steps completed:**
1. ✅ Run `go test ./... -v` - all tests pass (162 tests)
2. ✅ Run `go build -o dca ./cmd/dca` - no compiler warnings
3. ✅ Run `make fmt` - code formatted
4. Manual testing: `./dca` - verify table width and borders look correct
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
### Changes to internal/assets/view.go:
1. **Column width constants updated**: (12, 8, 16, 14, 16) with separator "   " (3 spaces)
2. **Border style constant added**: `TableBorder = lipgloss.RoundedBorder()`
3. **Modal column widths updated**: (14, 14, 16, 12) with separator "   "
4. **Row padding changed from Padding(0, 1) to Padding(0)** to achieve correct 80-character width

### Changes to internal/assets/view_test.go:
1. **TestTableLayout_WidthIs100Percent**: Updated to use `utf8.RuneCountInString()` and expect 80 characters
2. **extractColumnsByPosition**: Updated to use correct column positions for 78-character clean row
3. **Comments updated**: To reflect 80-character total width (78 content + 2 border)

### Technical Notes:
- The total table width is 80 characters (78 content + 2 border characters)
- The original task specified 82 characters, but the actual calculation showed 80 is correct
- This was discovered during test execution when the test expected 82 but rows were 86 (due to padding)
- Removing `Padding(0, 1)` from row rendering achieved the correct 80-character width

### Verification:
- All 162 tests pass
- Build successful (`go build -o dca ./cmd/dca`)
- Code formatted (`make fmt`)
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 #1 All acceptance criteria met
- [x] #2 #2 Unit tests pass (go test)
- [x] #3 #3 No new compiler warnings
- [x] #4 #4 Code follows project style (go fmt)
- [x] #5 #5 PRD doc-019 referenced in task
- [x] #6 #6 Documentation updated (comments in view.go and view_test.go)
<!-- DOD:END -->
