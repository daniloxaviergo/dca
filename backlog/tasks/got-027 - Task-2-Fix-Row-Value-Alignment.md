---
id: GOT-027
title: 'Task 2: Fix Row Value Alignment'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 22:00'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Fix row value formatting to ensure all values align with their column headers using fmt.Sprintf with width specifiers.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Header row values are left-aligned for text columns (Asset)
- [x] #2 Numeric columns are right-aligned
- [x] #3 All column values use fixed-width formatting with fmt.Sprintf
- [x] #4 Row values match header column width exactly
- [x] #5 Visual inspection confirms column alignment
- [x] #6 Unit tests verify alignment
- [x] #7 go fmt applied
- [x] #8 go build succeeds
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The issue is that column values are not properly aligned with their headers. The current code defines column widths but doesn't ensure header text and data values occupy the same width with proper alignment.

**Root Cause Analysis:**
- Headers like "Total Shares" (14 chars) exceed the defined `ColumnSharesWidth` (12 chars)
- Numeric values use `%*d` and `%*.2f` format but the width doesn't account for header text length
- Column separator ("  ") is added between fields but doesn't account for width specifiers

**Solution:**
1. Adjust column width constants to accommodate header text + separator
2. Use consistent `fmt.Sprintf` format specifiers for both headers and data:
   - Asset: `%-*s` (left-aligned string)
   - Count, Total Shares, Avg Price, Total Value: `%*s` (right-aligned string)
3. Convert numeric values to strings with proper formatting before applying width
4. Apply lipgloss styling to formatted strings

**Column Width Adjustments:**
- Asset: 10 chars (header "Asset" = 5 chars) ✓ Already correct
- Count: 8 chars (header "Count" = 5 chars) ✓ Already correct
- Total Shares: 14 chars (header "Total Shares" = 12 chars) + 2 separator = 14 ✓ Already correct
- Avg Price: 13 chars (header "Avg Price" = 9 chars) + 2 separator = 13 (currently 12)
- Total Value: 14 chars (header "Total Value" = 11 chars) + 2 separator = 13 (currently 14)

### 2. Files to Modify

- `internal/assets/view.go`: Update column width constants and rendering functions
  - Update `ColumnAvgPriceWidth` from 12 to 13
  - Update `renderHeaderRow()` to use consistent formatting
  - Update `renderDataRow()` to use consistent formatting
  - Ensure all values are converted to strings with `fmt.Sprintf` before lipgloss styling

- `internal/assets/view_test.go`: Add alignment verification tests
  - Test that headers and data rows have matching column positions
  - Test alignment with various value lengths

### 3. Dependencies

- No new dependencies required
- Existing: lipgloss v1.1.0, bubbletea v1.3.10
- Prerequisites: Task 1 (Define Fixed Column Widths) should be complete

### 4. Code Patterns

Follow existing patterns in `internal/assets/view.go`:
- Use `fmt.Sprintf` with width specifiers for all column formatting
- Apply lipgloss styling after formatting strings
- Use `ColumnSeparator` constant for consistent spacing
- Test with various value lengths to verify alignment

**Formatting patterns:**
```go
// Asset (text, left-aligned)
fmt.Sprintf("%-*s", ColumnAssetWidth, text)

// Numeric columns (right-aligned, formatted as strings)
fmt.Sprintf("%*s", ColumnCountWidth, fmt.Sprintf("%d", value))
fmt.Sprintf("%*.8f", ColumnSharesWidth, value)
fmt.Sprintf("%*.2f", ColumnAvgPriceWidth, value)
```

### 5. Testing Strategy

Add tests to verify column alignment:
- Test that header and data row character positions match for each column
- Test with short and long ticker symbols
- Test with minimal and maximal numeric values
- Test visual output with `go run main.go` (manual verification)

**New test cases:**
```go
// TestRenderHeaderRow_Alignment verifies headers align with data columns
func TestRenderHeaderRow_Alignment(t *testing.T) {
    av := NewAssetsView()
    header := av.renderHeaderRow()
    row := av.renderDataRow(0, AssetSummary{...})
    
    // Find column positions in header
    // Verify data row has values at same positions
}

// TestRenderDataRow_AlignmentWithHeader verifies each data column aligns with header
func TestRenderDataRow_AlignmentWithHeader(t *testing.T) {
    // Test each column's alignment
}
```

Run tests: `go test ./internal/assets/... -v`

### 6. Risks and Considerations

**Known Issues:**
- Header "Total Shares" (12 chars) exceeds `ColumnSharesWidth` (12 chars) exactly, no room for padding
- Header "Avg Price" (9 chars) with current width 12 leaves 3 chars padding (acceptable)
- Header "Total Value" (11 chars) with current width 14 leaves 3 chars padding (acceptable)

**Trade-offs:**
- May need to add 1 char to `ColumnAvgPriceWidth` to ensure proper alignment
- Column separator ("  ") must be accounted for in total table width calculation

**Verification:**
- Run `go build` and execute binary to visually inspect table alignment
- Compare header and data row character-by-character
- Ensure columns are vertically aligned across all rows
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Task reference: backlog/docs/doc-006.md - PRD for Assets View feature.

## Implementation Notes

### Column Width Calculations
The root cause was that header text lengths exceeded the defined column widths:
- "Total Shares" (12 chars) required width of at least 14 (header + 2-char separator)
- "Avg Price" (9 chars) required width of at least 13 (header + 2-char separator)
- "Total Value" (11 chars) required width of 13 (header + 2-char separator)

### Alignment Testing Strategy
Added comprehensive tests:
1. `TestColumnWidths_ConstantsDefined` - Verify width constants are correct
2. `TestRenderHeaderRow_Alignment` - Verify headers render correctly
3. `TestRenderDataRow_AlignmentWithHeader` - Verify data row field widths
4. `TestRenderTable_ColumnAlignment` - Verify headers and data align in full table
5. `TestRenderDataRow_FieldWidths` - Table-driven test with various value lengths

### Build Verification
Built cmd/dca package successfully with no warnings. All tests pass.

### go fmt Applied
No formatting changes required - code already followed project style.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Fix row value formatting to ensure all values align with column headers using fmt.Sprintf with width specifiers.

## Changes Made

### internal/assets/view.go
- Updated `ColumnSharesWidth` from 12 to 14 (to accommodate "Total Shares" header + separator)
- Updated `ColumnAvgPriceWidth` from 12 to 13 (to properly align "Avg Price" header)
- Updated `ColumnTotalValueWidth` from 14 to 13 (corrected to match header + separator)
- Updated comment for total table width calculation

### internal/assets/view_test.go
- Updated `TestColumnWidths_ConstantsDefined` to expect new column widths
- Added `TestRenderHeaderRow_Alignment` to verify header alignment
- Added `TestRenderDataRow_AlignmentWithHeader` to verify data row alignment
- Added `TestRenderTable_ColumnAlignment` to verify headers and data rows align
- Added `TestRenderDataRow_FieldWidths` to test various value lengths
- Added helper function `renderHeaderRowForTest()` for testing header rendering
- Added `fmt` import

### Verification
- All tests pass: `go test ./...` (5 packages)
- Build succeeds: `go build -o dca cmd/dca`
- No compiler warnings
- go fmt applied (no formatting changes needed)

## Alignment Details
- Asset: 10 chars (left-aligned)
- Count: 8 chars (right-aligned)
- Total Shares: 14 chars (right-aligned with 8 decimals)
- Avg Price: 13 chars (right-aligned with 2 decimals)
- Total Value: 13 chars (right-aligned with 2 decimals)
- Column separator: 2 spaces between columns
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
