---
id: GOT-027
title: 'Task 2: Fix Row Value Alignment'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 21:39'
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
- [ ] #1 Header row values are left-aligned for text columns (Asset)
- [ ] #2 Numeric columns are right-aligned
- [ ] #3 All column values use fixed-width formatting with fmt.Sprintf
- [ ] #4 Row values match header column width exactly
- [ ] #5 Visual inspection confirms column alignment
- [ ] #6 Unit tests verify alignment
- [ ] #7 go fmt applied
- [ ] #8 go build succeeds
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

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
