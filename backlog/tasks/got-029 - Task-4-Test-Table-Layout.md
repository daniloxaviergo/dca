---
id: GOT-029
title: 'Task 4: Test Table Layout'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 20:22'
updated_date: '2026-03-18 14:25'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add unit tests for table layout and alignment to verify the rendering fixes work correctly.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Test verifies table width is 100% of available width
- [ ] #2 Test verifies header alignment with data columns
- [ ] #3 Test verifies exactly 30 rows are rendered
- [ ] #4 Test verifies empty row padding works correctly
- [ ] #5 All unit tests pass
- [ ] #6 Test coverage for layout functions
- [ ] #7 go test passes with no warnings
- [ ] #8 go fmt applied
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Add comprehensive unit tests for table layout and alignment to verify the rendering fixes work correctly. The tests will verify:

1. **Table Width**: Test that table rendering uses the defined column widths (10 + 8 + 14 + 13 + 13 + 8 separators = 78 characters) to achieve full-width display
2. **Header Alignment**: Test that header row uses same column widths as data rows with `fmt.Sprintf` width specifiers
3. **Row Count**: Test that exactly 30 rows are rendered (1 header + up to 29 data/empty rows)
4. **Empty Row Padding**: Test that empty rows are added when data rows < 30, using same formatting as data rows

**Approach**:
- Use table-driven tests for different asset counts (0, 5, 25, 30+ assets)
- Parse rendered output to verify row counts and alignment
- Test column width consistency by checking formatted string lengths
- Verify empty row padding logic works correctly

### 2. Files to Modify

**New tests to add in `internal/assets/view_test.go`:**
- `TestTableLayout_WidthIs100Percent` - Verify table uses defined column widths
- `TestTableLayout_HeaderAlignment` - Verify header matches data column alignment
- `TestTableLayout_Exactly30Rows` - Verify row count is always 30
- `TestTableLayout_EmptyRowPadding` - Verify empty rows are rendered correctly
- `TestTableLayout_ColumnWidthFormatting` - Verify `fmt.Sprintf` width specifiers work

### 3. Dependencies

**No additional dependencies required:**
- Existing: `internal/assets/view.go` - Contains table rendering logic
- Existing: `internal/assets/aggregate.go` - Data model (AssetSummary)
- Existing: `internal/dca/entry.go` - DCAEntry structure

**Prerequisites:**
- All column widths defined as constants in `view.go` (already in place)
- Empty row padding logic already implemented in `renderTable()`
- Fixed-width formatting already using `fmt.Sprintf` width specifiers

### 4. Code Patterns

**Follow existing test patterns in `view_test.go`:**
- Test naming: `Test{Function}_{Condition}` (e.g., `TestTableLayout_WidthIs100Percent`)
- Table-driven tests for multiple scenarios
- String matching with `strings.Contains()` for partial output verification
- Direct rendering tests (no full Bubble Tea event loop needed)
- Use `t.Errorf()` for assertions with clear error messages

**Conventions to follow:**
- Use `strings.Split(output, "\n")` to parse rendered rows
- Count actual rows by checking for data or border characters
- Verify column separators ("  ") are consistent
- Test edge cases: 0 assets, 1 asset, 29 assets, 30 assets, 31+ assets

### 5. Testing Strategy

**Tests to add:**

1. **`TestTableLayout_WidthIs100Percent`** (Acceptance #1)
   - Create AssetsView with entries
   - Render and extract table rows
   - Verify column widths match constants (10, 8, 14, 13, 13)
   - Verify separator spacing (2 spaces)

2. **`TestTableLayout_HeaderAlignment`** (Acceptance #2)
   - Render header row and data row
   - Split rows by column separator
   - Verify same number of columns
   - Verify same widths per column

3. **`TestTableLayout_Exactly30Rows`** (Acceptance #3)
   - Test with 5 assets (expect 5 data + 24 empty + 1 header = 30)
   - Test with 25 assets (expect 25 data + 4 empty + 1 header = 30)
   - Test with 30+ assets (expect 30 rows max, no extra)
   - Parse output to count actual rows

4. **`TestTableLayout_EmptyRowPadding`** (Acceptance #4)
   - Test with < 30 entries
   - Verify empty rows use same formatting as data rows
   - Verify empty rows contain zeros for numeric fields
   - Verify active row styling works on empty rows

5. **`TestTableLayout_RowCountCalculation`**
   - Test empty row calculation: `maxVisibleRows - 1 - len(entries)`
   - Verify no negative padding

**Coverage targets:**
- All layout functions: `renderTable`, `renderHeaderRow`, `renderDataRow`, `renderEmptyDataRow`
- All column width constants used in tests
- Edge cases: boundary conditions for row counts

### 6. Risks and Considerations

**Potential issues:**
- **Terminal width dependency**: Tests verify column widths but not actual terminal display - consider this a limitation
- **Border characters**: Tests may need to account for lipgloss border style (RoundedBorder)
- **Active row styling**: Empty row padding must maintain active row highlight logic

**Trade-offs:**
- Tests verify format/width consistency but not visual rendering (terminal-specific)
- Row counting may need to account for header, footer, and blank lines in full output
- Column width constants are already defined; tests can reference them directly

**Blocking issues:**
- None identified - code is ready for testing

**Ready to proceed:** Implementation can begin once user approves this plan.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
[2026-03-18] Analysis complete. Two tests are failing: 1) TestTableLayout_HeaderAlignment - Column width mismatches between header and data rows 2) TestTableLayout_ColumnWidthsMatchConstants - Column widths don't match defined constants. Root cause: Tests use len(strings.TrimSpace(...)) which measures trimmed content (e.g., Asset = 5 chars) instead of formatted width (10 chars). The tests need to measure the raw formatted string width before trimming.
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
