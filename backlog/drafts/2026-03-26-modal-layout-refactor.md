# Task: Refactor Modal Asset History Layout

**Status:** Draft  
**Created:** 2026-03-26  
**Priority:** Medium  
**Labels:** feature, uiux  

## Summary

Refactor the asset history modal's table layout to improve readability, visual hierarchy, and data clarity by reorganizing columns and adding visual separation between daily and cumulative metrics.

## Problem Statement

The current modal table displays cumulative and daily metrics together without clear visual distinction, making it confusing to understand which values represent single-day activity versus running totals. Column widths are also not optimally distributed for the content.

## Requirements Reference

See `plan_layout.md` for detailed sprint requirements:
- **Must Have:** Reorganize Modal Columns for Better Readability
- **Must Have:** Add Visual Separation Between Daily and Cumulative Data
- **Should Have:** Implement Date Range Filter (optional inclusion)
- **Could Have:** Add Chart Visualization Option (optional inclusion)

## Acceptance Criteria

### Must Have

- [ ] Modal column order: Date, Entry Count, Daily Invested, Daily Price, Cumulative Invested
- [ ] Cumulative column clearly labeled in header (e.g., "Cum. Investing" or "Total to Date")
- [ ] Column widths: Date(12), Count(10), Daily Invested(14), Daily Price(12), Cumulative Invested(14)
- [ ] Total modal width ≤64 characters
- [ ] Daily metrics use distinct color theme (e.g., light cyan)
- [ ] Cumulative metrics use secondary color theme (e.g., muted blue)
- [ ] Visual grouping with appropriate spacing between column groups
- [ ] All existing tests pass after refactoring

### Should Have (Optional)

- [ ] Filter prompt opens on 'f' key press
- [ ] Filter accepts YYYY-MM-DD format
- [ ] Support All/From Date/Range modes
- [ ] "Filter active" indicator displayed
- [ ] Clear filter with 'F' or 'Esc'

## Dependencies

- [ ] Data model (`EntryByDate`) must include daily and cumulative fields
- [ ] View rendering functions must support new column structure
- [ ] Tests must be updated for new column widths and order

## Testing Strategy

1. **Visual Testing:**
   - Modal displays without horizontal scroll on 1200px display
   - Color themes work in both light and dark terminals
   - Column alignment preserved across all data lengths

2. **Unit Testing:**
   - Update `TestRenderModalHeaderRow` for new column order
   - Update `TestRenderModalDataRow` for new color scheme
   - Verify column widths match new constants

3. **Integration Testing:**
   - Modal opens correctly with new layout
   - "Load More" displays data with new format
   - Empty state still displays properly

## Implementation Notes

See `internal/assets/view.go` for current implementation:
- Modal width constants (lines 24-32)
- `renderModalHeaderRow()` function
- `renderModalDataRow()` function
- `renderModalContent()` function

## Progress

- [ ] Draft completed
- [ ] Requirements aligned with `plan_layout.md`
- [ ] Technical implementation plan ready
- [ ] Tests identified for updates
- [ ] Ready for estimation

## References

- Sprint Plan: `plan_layout.md`
- Current Implementation: `internal/assets/view.go`, `internal/assets/model.go`
- Test File: `internal/assets/view_test.go`
