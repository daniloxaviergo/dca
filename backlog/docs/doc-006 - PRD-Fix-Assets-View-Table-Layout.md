---
id: doc-006
title: 'PRD: Fix Assets View Table Layout'
type: other
created_date: '2026-03-17 20:17'
---
# PRD: Fix Assets View Table Layout

## Overview

### Purpose
Fix the table layout in the Assets View to display data in a properly aligned, full-width table with exactly 30 rows, ensuring all column values are correctly aligned with their headers.

### Goals
- **Primary**: Table spans 100% of available width with consistent column widths
- **Secondary**: Table displays exactly 30 rows (showing actual data rows with empty row padding as needed)
- **Success**: All column values (Asset, Count, Total Shares, Avg Price, Total Value) are perfectly vertically aligned with their headers

## Background

### Problem Statement
The current Assets View table has layout issues where:
1. The table does not span the full width of the terminal
2. Row values are not properly aligned with their corresponding column headers
3. No minimum row count is enforced, leading to inconsistent visual presentation

### Current State
The table uses `lipgloss.JoinVertical` with a simple header row and data rows. Column widths are determined dynamically by the content, leading to misalignment when:
- Header names are longer than data values
- Data values have varying number of digits
- No explicit column width constraints are set

Current rendering code:
```go
// renderHeaderRow uses bold styling but no width constraints
return strings.Join(formatted, "  ")

// renderDataRow formats values but doesn't enforce consistent width
rowStr := strings.Join(formatted, "  ")
```

### Proposed Solution
1. Define fixed column widths for each column (Asset, Count, Total Shares, Avg Price, Total Value)
2. Use `fmt.Sprintf` with width specifiers to ensure consistent row value formatting
3. Pad rows with empty data rows to maintain exactly 30 rows total
4. Apply lipgloss width constraints to ensure 100% table width

## Requirements

### User Stories

- **End User**: As a user viewing the Assets View, I want to see a table that fills the entire terminal width with perfectly aligned columns so that I can easily read and compare my asset data.

- **End User**: As a user with few assets, I want to see exactly 30 rows (with empty rows as padding) so that the table maintains a consistent visual height regardless of data volume.

### Functional Requirements

#### Task 1: Define Fixed Column Widths

Define explicit column widths for all 5 columns to ensure 100% table width coverage.

##### Acceptance Criteria
- [ ] Asset column: minimum 10 characters width
- [ ] Count column: 8 characters width (integer values)
- [ ] Total Shares column: 12 characters width (8 decimal precision + formatting)
- [ ] Avg Price column: 12 characters width (2 decimal precision + formatting)
- [ ] Total Value column: 14 characters width (2 decimal precision + formatting)
- [ ] Column separator: 2 spaces between columns
- [ ] Total table width: 100% of terminal width (adjustable based on screen)

#### Task 2: Fix Row Value Alignment

Fix row value formatting to ensure all values align with their column headers.

##### Acceptance Criteria
- [ ] Header row values are left-aligned for text columns (Asset)
- [ ] Numeric columns (Count, Total Shares, Avg Price, Total Value) are right-aligned
- [ ] All column values use fixed-width formatting with `fmt.Sprintf`
- [ ] Row values match header column width exactly

##### User Flows
1. User navigates to Assets View
2. System displays table with headers
3. System displays data rows with aligned values
4. User can navigate rows with up/down arrows

#### Task 3: Enforce Minimum 30 Rows

Ensure the table displays exactly 30 rows (actual data + padding).

##### Acceptance Criteria
- [ ] When data rows < 30: pad with empty rows to reach 30
- [ ] When data rows = 30: display all rows without truncation
- [ ] When data rows > 30: scroll (not in scope for this PRD)
- [ ] Empty rows use same styling as data rows but with empty values

##### User Flows
1. User has 5 assets in their data
2. System displays 5 data rows + 25 empty rows = 30 total rows
3. User can navigate all 30 rows (including empty padded rows)

#### Task 4: Test Table Layout

Add tests to verify table layout and alignment.

##### Acceptance Criteria
- [ ] Test verifies table width is 100% of available width
- [ ] Test verifies header alignment with data columns
- [ ] Test verifies exactly 30 rows are rendered
- [ ] Test verifies empty row padding works correctly

### Non-Functional Requirements

- **Performance**: Table rendering should complete within 100ms for 30 rows
- **Compatibility**: Must work with lipgloss v1.1.0 and bubbletea v1.3.10
- **Maintainability**: Column widths should be defined as constants for easy adjustment
- **Responsiveness**: Table should adapt to terminal width changes

## Scope

### In Scope
- Define fixed column widths for all 5 columns
- Implement `fmt.Sprintf` with width specifiers for row value formatting
- Add row padding to maintain exactly 30 rows
- Update table rendering to ensure 100% width coverage
- Add unit tests for table layout

### Out of Scope
- Dynamic column width adjustment based on terminal size
- Horizontal scrolling for wide terminals
- Virtual scrolling for >30 rows
- Column reordering or resizing

## Technical Considerations

### Existing System Impact
- **Changes**: `internal/assets/view.go` - `renderHeaderRow`, `renderDataRow`, `renderTable` functions
- **No Breaking Changes**: Existing data structures remain unchanged
- **Dependencies**: lipgloss v1.1.0 (already in use)

### Dependencies
- lipgloss for styling
- bubbletea for TUI framework

### Constraints
- Terminal width varies by user setup
- Must maintain backward compatibility with existing data format
- No database or external API dependencies

## Success Metrics

### Quantitative
- Table width: 100% of available terminal width
- Column alignment: All column values vertically aligned with headers (measured by visual inspection)
- Row count: Exactly 30 rows displayed

### Qualitative
- User can easily scan and compare asset data
- Table appears "tight" and professional
- No visual gaps or misalignments between columns

## Timeline & Milestones

### Key Dates
- Design complete: [Date]
- Implementation complete: [Date]
- Testing complete: [Date]
- Launch/Release: [Date]

## Stakeholders

### Decision Makers
- [User]: Product Owner

### Contributors
- [Developer]: Implementation

## Appendix

### Glossary
- **Table Width**: Horizontal span of the table, expressed as percentage of terminal width
- **Column Alignment**: Vertical alignment of row values with their column headers
- **Row Padding**: Empty rows added to maintain minimum row count

### References
- [internal/assets/view.go](file:///home/danilo/scripts/github/dca/internal/assets/view.go): Current table rendering implementation
- [QWEN.md](file:///home/danilo/scripts/github/dca/QWEN.md): Project context and conventions
