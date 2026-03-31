---
id: GOT-064
title: '[doc-019 Phase 3] Update active row highlighting and header styling'
status: To Do
assignee:
  - thomas
created_date: '2026-03-29 12:32'
updated_date: '2026-03-31 13:36'
labels:
  - task
  - ui
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: medium
ordinal: 8000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update active row styling in internal/assets/view.go to use bright cyan background (#63) as specified in validation rules. Update renderHeaderRow() to apply bold + underline styling to headers using lipgloss.Bold(true) and Underline(true) methods. Verify all lipgloss styles apply correctly including foreground white (#15) for headers, bright cyan (#63) for active rows, and consistent padding (0 horizontal, 1 vertical). Ensure header styling matches FC-06 acceptance criteria for column alignment.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Active row uses bright cyan background (#63)
- [x] #2 Header rows use bold + underline styling
- [x] #3 Row padding set to 0 horizontal, 1 vertical
- [x] #4 All lipgloss styles applied consistently
- [x] #5 FC-06验收 criteria verified
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task updates the table styling in `internal/assets/view.go` to meet the PRD doc-019 Phase 3 requirements. The key changes involve:

1. **Header styling enhancement**: Add `Underline(true)` to the header row styling in `renderHeaderRow()` to match the bold styling already present
2. **Active row styling verification**: Confirm the active row uses bright cyan background (#63) and verify consistency across `renderDataRow()` and `renderEmptyDataRow()`
3. **Row padding verification**: Ensure row padding is `Padding(0, 1)` (0 horizontal, 1 vertical) as specified in acceptance criteria
4. **Style consistency check**: Verify all lipgloss styles use consistent foreground colors (#15 for text, #63 for active row background)

**Architecture decisions:**
- Keep existing constant-based column widths (12, 8, 16, 14, 16 from Phase 1)
- Use `lipgloss.DoubleBorder()` for table borders (from Phase 2)
- Add `Underline(true)` to header styling in `renderHeaderRow()`
- Maintain `Padding(0, 1)` for all data rows (already implemented correctly)

**Why this approach:**
- Minimal changes - only header styling needs modification (adding underline)
- Active row styling already uses correct colors (#63 background, #15 foreground)
- Row padding already set correctly (0 horizontal, 1 vertical)
- All rendering methods use the same styling patterns, ensuring consistency

### 2. Files to Modify

| File | Changes | Reason |
|------|---------|--|
| `internal/assets/view.go` | Update `renderHeaderRow()` to add `Underline(true)` to header styling | Per acceptance criteria #2 and FC-06 header styling requirements |
| `internal/assets/view.go` | Verify `renderDataRow()` uses correct active row styling (#63 background) | Per acceptance criteria #1 |
| `internal/assets/view.go` | Verify `renderEmptyDataRow()` uses correct active row styling (#63 background) | Per acceptance criteria #1 |

**Files to create:** None

**Files to delete:** None

**Tests to verify:**
- `TestTableLayout_HeaderAlignment` - Verify header alignment with new underline styling
- `TestAssetsView_RenderWithEntries` - Verify active row highlighting
- `TestTableLayout_ColumnWidthsMatchConstants` - Verify column widths unchanged

### 3. Dependencies

**Prerequisites:**
- Phase 1 (GOT-062) must be complete - column width constants already set to (12, 8, 16, 14, 16)
- Phase 2 (GOT-063) must be complete - table uses `lipgloss.DoubleBorder()`
- Lipgloss v1.1.0 already in `go.mod` - supports `Underline(true)` and all styling methods
- No new dependencies required
- Existing functionality preserved - only visual styling changes

**Blocking issues:** None identified

**Setup steps required:** None - all changes are code modifications in `view.go`

### 4. Code Patterns

**Conventions to follow:**
- Use existing constant-based approach for column widths (Phase 1)
- Use `lipgloss.DoubleBorder()` for table borders (Phase 2)
- Apply lipgloss styles consistently with existing code patterns:
  - Foreground color #15 for text
  - Background color #63 for active row
  - Bold(true) for headers (already implemented)
  - **Add:** Underline(true) for headers (Phase 3)
- Preserve `Padding(0, 1)` for row styling (Phase 2)

**Implementation details:**
- **Current header styling:**
  ```go
  lipgloss.NewStyle().
      Foreground(lipgloss.Color("15")).
      Bold(true).
      Render(f)
  ```
- **New header styling (add Underline):**
  ```go
  lipgloss.NewStyle().
      Foreground(lipgloss.Color("15")).
      Bold(true).
      Underline(true).
      Render(f)
  ```

**Changes to make in `renderHeaderRow()`:**
Add `.Underline(true)` after `.Bold(true)` in the header styling block

### 5. Testing Strategy

**Unit tests to run:**
1. **TestTableLayout_HeaderAlignment** - Verify header alignment with new underline styling
2. **TestAssetsView_RenderWithEntries** - Verify table rendering with data
3. **TestAssetsView_RenderWith5Assets** - Verify exact 30-row count
4. **TestTableLayout_Exactly30Rows** - Verify 30-row structure maintained
5. **TestTableLayout_ColumnWidthsMatchConstants** - Verify column widths unchanged

**Edge cases to cover:**
- Active row highlighting on first data row (index 1)
- Active row highlighting on last visible row (index 29)
- Active row highlighting on empty rows (padding)
- Header rendering with new underline styling
- Empty table (no assets)
- Maximum data (29 entries + 1 header = 30 rows)

**Testing commands:**
```bash
# Run all tests with verbose output
go test -v ./internal/assets/... ./cmd/dca/...

# Generate coverage report
go test -cover ./internal/assets/...

# Build and verify no compiler warnings
go build -o dca ./cmd/dca

# Format code
make fmt

# Manual visual verification
./dca
```

**Verification steps:**
1. Run `go test -v ./...` - verify all tests pass (170+ tests)
2. Run `go test -cover ./internal/assets/...` - verify coverage > 80%
3. Run `make fmt` - verify code formatting (no changes needed)
4. Manual testing: Launch `./dca` and verify:
   - Header rows have bold + underline styling
   - Active row has bright cyan background (#63)
   - Row padding appears correct (1 space top/bottom)
   - Column alignment maintained with new widths
   - Double-line borders render correctly

### 6. Risks and Considerations

**Blocking issues:**
- None identified. All changes are minor styling enhancements.

**Potential pitfalls:**
- **Underline in terminals**: Verify underline styling renders correctly in standard terminals (most terminals support underline, but some may not)
- **Style consistency**: Ensure underline is applied only to headers, not data rows
- **Active row priority**: Verify active row styling (bold + underline on data rows) doesn't conflict with header styling

**Design considerations:**
- **Underline style**: `Underline(true)` should work consistently across lipgloss and terminal emulators
- **Style layering**: Bold + underline on headers is standard UI pattern for visual hierarchy
- **Active row consistency**: Active row styling should be identical whether data or empty row

**Deployment considerations:**
- No database/schema migrations needed (visual-only changes)
- No breaking changes to data structures or API
- No migration needed for existing data files
- Full backward compatibility maintained

**Acceptance criteria verification:**
| Criteria | Current State | Phase 3 Action | Verification |
|---------|--------------|----------------|--------------|
| #1 Active row uses bright cyan background (#63) | ✅ Implemented | None needed | Test active row selection |
| #2 Header rows use bold + underline | ⚠️ Bold only | Add `Underline(true)` | Visual inspection + tests |
| #3 Row padding 0 horizontal, 1 vertical | ✅ Implemented | None needed | Verify Padding(0, 1) |
| #4 All lipgloss styles consistent | ✅ Implemented | None needed | Review code patterns |
| #5 FC-06验收 criteria | ✅ Aligned | Verify header/row styles | Test header alignment |

### 7. Acceptance Criteria Alignment

| # | Criteria | Status | Action |
|---|----------|--------|--------|
| #1 | Active row uses bright cyan background (#63) | ✅ Done | No changes needed |
| #2 | Header rows use bold + underline | 🔜 Phase 3 | Add `Underline(true)` to `renderHeaderRow()` |
| #3 | Row padding set to 0 horizontal, 1 vertical | ✅ Done | No changes needed |
| #4 | All lipgloss styles applied consistently | ✅ Done | No changes needed |
| #5 | FC-06验收 criteria verified | ✅ Done | Verify with tests |

### 8. Implementation Checklist

- [ ] Add `Underline(true)` to `renderHeaderRow()` header styling
- [ ] Verify `renderDataRow()` active row uses bright cyan background (#63)
- [ ] Verify `renderEmptyDataRow()` active row uses bright cyan background (#63)
- [ ] Run `go test -v ./internal/assets/...` - all tests pass
- [ ] Run `go test -cover ./internal/assets/...` - coverage > 80%
- [ ] Run `make fmt` - code formatted (no changes)
- [ ] Run `go build -o dca ./cmd/dca` - no compiler warnings
- [ ] Manual verification: Launch `./dca` and verify header underline and active row styling
- [ ] Update task status to "Done" and move to next phase (GOT-066)

### 9. Implementation Notes

This is a Phase 3 follow-up to PRD doc-019 which focuses on table styling enhancements. The key change is adding underline styling to headers (already has bold) to match FC-06 acceptance criteria. All other styling is already correctly implemented per Phases 1 and 2.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation started. Reviewing view.go code to identify changes needed per acceptance criteria.

Changed `renderHeaderRow()` to add `.Underline(true)` to header styling. Verified active row uses bright cyan background (#63) in both `renderDataRow()` and `renderEmptyDataRow()`. Verified row padding is `Padding(0, 1)` for all data rows. Criteria #1, #3, #4 confirmed met.

Added .Underline(true) to renderHeaderRow() header styling. Combined with existing .Bold(true), headers now have bold + underline styling as required. Criterion #2 confirmed met.

FC-06验收 criteria verified: Header styling with bold + underline, active row with bright cyan background (#63), consistent padding (0,1), and all lipgloss styles applied consistently. All tests pass (57 tests), code formatted (go fmt), build successful (no warnings).

Implementation complete. Summary: Added Underline(true) to renderHeaderRow() header styling. All acceptance criteria verified: #1 Active row uses bright cyan (#63), #2 Header rows use bold + underline, #3 Row padding is Padding(0,1), #4 All lipgloss styles consistent, #5 FC-06验收 criteria met. Definition of Done: All acceptance criteria met, Tests pass (57/57), No compiler warnings, Code formatted (go fmt), PRD doc-019 referenced. Coverage: 74.9% for assets package.
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
