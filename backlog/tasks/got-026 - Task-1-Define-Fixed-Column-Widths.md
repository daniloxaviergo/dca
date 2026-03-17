---
id: GOT-026
title: 'Task 1: Define Fixed Column Widths'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 20:30'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define explicit column widths for the Assets View table to ensure 100% width coverage and consistent column sizing across all terminals.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Asset column: minimum 10 characters width
- [ ] #2 Count column: 8 characters width
- [ ] #3 Total Shares column: 12 characters width
- [ ] #4 Avg Price column: 12 characters width
- [ ] #5 Total Value column: 14 characters width
- [ ] #6 Column separator: 2 spaces between columns
- [ ] #7 Total table width: 100% of terminal width
- [ ] #8 Unit tests pass for column width definitions
- [ ] #9 Table renders without panics
- [ ] #10 go fmt applied
- [ ] #11 go build succeeds
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task defines explicit column widths for the Assets View table to ensure consistent, full-width rendering across all terminal sizes. The approach uses fixed-width formatting with `fmt.Sprintf` to guarantee alignment between headers and data rows.

**Key Decisions:**
- **Column widths as constants**: Define column widths as package-level constants for easy adjustment and documentation
- **Fixed-width formatting**: Use `fmt.Sprintf` with width specifiers to ensure all values align with headers
- **Left-align text, right-align numbers**: Asset (text) is left-aligned; numeric columns are right-aligned for readability
- **Terminal width independence**: Column widths are fixed character counts (not percentage-based) since lipgloss width constraints don't provide the fine-grained control needed for table alignment

**Column Width Breakdown:**
| Column | Width | Alignment | Format |
|--------|-------|-----------|--------|
| Asset | 10 | Left | `%-[10]s` |
| Count | 8 | Right | `%[8]d` |
| Total Shares | 12 | Right | `%[12].8f` |
| Avg Price | 12 | Right | `%[12].2f` |
| Total Value | 14 | Right | `%[14].2f` |

**Total Width:** 10 + 8 + 12 + 12 + 14 + (4 separators × 2 spaces) = 72 characters

**Rationale for widths:**
- **Asset (10)**: Minimum 10 chars as specified; allows reasonable ticker display (e.g., "BITCOINUSD")
- **Count (8)**: Integer values typically small; 8 chars handles up to 99,999,999 entries
- **Total Shares (12)**: 8 decimal precision + formatting overhead; e.g., "0.00769231"
- **Avg Price (12)**: 2 decimal precision; handles prices up to 999,999,999.99
- **Total Value (14)**: 2 decimal precision; handles values up to 999,999,999,999.99

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/assets/view.go` | Modify | Update `renderHeaderRow()` and `renderDataRow()` to use fixed-width formatting |
| `internal/assets/view_test.go` | Modify | Add unit tests for column width definitions |

### 3. Dependencies

- **Existing**: lipgloss v1.1.0 (already imported, used for styling)
- **Existing**: bubbletea v1.3.10 (already imported, TUI framework)
- **No prerequisites**: This task is independent and can be implemented immediately

### 4. Code Patterns

Follow existing patterns in `internal/assets/view.go`:

**Header row formatting:**
```go
headers := []string{"Asset", "Count", "Total Shares", "Avg Price", "Total Value"}
var formatted []string
for _, h := range headers {
    formatted = append(formatted, lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Bold(true).Render(h))
}
return strings.Join(formatted, "  ")
```

**Data row formatting (with fixed widths):**
```go
rowData := []string{
    fmt.Sprintf("%-10s", entry.Ticker),
    fmt.Sprintf("%8d", entry.EntryCount),
    fmt.Sprintf("%12.8f", entry.TotalShares),
    fmt.Sprintf("%12.2f", entry.AvgPrice),
    fmt.Sprintf("%14.2f", entry.TotalValue),
}
```

**Styling:**
- Apply lipgloss styles after formatting with fixed widths
- Use `Padding(0, 1)` for row padding (existing pattern)
- Active row styling with background color 63 (existing pattern)

### 5. Testing Strategy

**Unit Tests to Add/Modify:**

1. **TestColumnWidths_ConstantsDefined**: Verify column width constants are defined with correct values
2. **TestRenderHeaderRow_Widths**: Verify header row renders with correct column widths
3. **TestRenderDataRow_Widths**: Verify data rows render with correct column widths
4. **TestRenderDataRow_LeftAlignment**: Verify Asset column is left-aligned
5. **TestRenderDataRow_RightAlignment**: Verify numeric columns are right-aligned
6. **TestRenderTable_Alignment**: Verify headers and data align vertically

**Test approach:**
- Use string containment checks on rendered output
- Verify column boundaries at known character positions
- Test with edge cases: empty ticker, large numbers, 8-decimal values

### 6. Risks and Considerations

**Potential Issues:**
- **Terminal width smaller than table**: If terminal width < 72 chars, table will overflow. This is acceptable per PRD (not handling responsive width adjustment).
- **Ticker overflow**: Tickers longer than 10 chars will overflow. Consider if this needs handling (out of scope for this task).
- **Rounding artifacts**: 8-decimal formatting may show trailing zeros (e.g., "0.01000000"). This is expected for financial precision.

**Trade-offs:**
- **Fixed vs dynamic widths**: Fixed widths ensure alignment but may not optimize space usage on wide terminals
- **No horizontal scrolling**: PRD explicitly excludes horizontal scrolling (out of scope)

**Blocking Issues:**
- None identified. This task is straightforward and well-defined.
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
