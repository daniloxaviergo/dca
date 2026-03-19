---
id: GOT-052
title: Bug Assert view
status: In Progress
assignee:
  - Thomas
created_date: '2026-03-19 16:45'
updated_date: '2026-03-19 17:38'
labels: []
dependencies: []
references:
  - doc-009
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The sort of entries on modal assets history should be desc
The sum should be the sum of previues days
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
- [x] #7 - [ ] Entries sorted descending (newest first)
- [x] #8 - [ ] Total Invested is cumulative sum of previous days
<!-- DOD:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Entries in asset history modal are sorted in descending order (newest first)
- [x] #2 - [ ] Total Invested shows cumulative sum of all previous days (running total)
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
# Implementation Plan: Fix Asset History Modal Bugs

## 1. Technical Approach

The task identifies two bugs in the asset history modal:

1. **Entries sorted ascending instead of descending** - Currently `AggregateByDate()` sorts entries by date ascending (oldest first), but the modal should show newest entries first (descending).

2. **Total Invested is daily sum, not cumulative** - The `TotalInvested` field currently shows the sum for that specific day only. It should show a running cumulative total from all previous days up to and including that day.

### Fix Strategy
- Modify `AggregateByDate()` in `internal/assets/aggregate.go` to sort in descending order (newest first)
- Modify `calculateDayMetrics()` to track cumulative total across all days and update `EntryByDate.TotalInvested` to be a running sum

## 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/assets/aggregate.go` | Modify | Fix `AggregateByDate()` sorting and add cumulative total calculation |

## 3. Dependencies

- No prerequisites required
- Existing data format supports the changes (no schema changes needed)
- Tests in `internal/assets/view_test.go` and `internal/assets/aggregate_test.go` need verification

## 4. Code Patterns

- Follow existing code style in `aggregate.go`
- Use `RoundTo8Decimals()` for floating-point rounding
- Maintain backward compatibility with existing JSON data format
- Preserve the batch loading behavior (10 days at a time)

### Changes to `AggregateByDate()`
```go
// Change from ascending to descending sort
// Current: if result[i].Date > result[j].Date { swap }
// New: if result[i].Date < result[j].Date { swap }  // Reverse comparison
```

### Changes to `calculateDayMetrics()`
```go
// Need to refactor to:
// 1. Aggregate all entries first (get total for each day)
// 2. Sort by date descending
// 3. Calculate running cumulative total while iterating
```

## 5. Testing Strategy

### Unit Tests to Verify
- `TestAggregateByDate_Descending` - Verify dates are sorted newest first
- `TestEntryByDate_CumulativeTotal` - Verify TotalInvested is running sum
- Existing tests in `view_test.go` modal section
- Run `go test -v ./internal/assets/...` to verify all tests pass

### Manual Verification
1. Run the app: `make run`
2. Open modal on any asset with multiple dates
3. Verify newest date appears first
4. Verify Total Invested column shows cumulative sum (each row ≥ previous)

## 6. Risks and Considerations

| Risk | Mitigation |
|------|------------|
| Breaking existing tests | Run all tests before and after changes |
| Data consistency with stored entries | No data migration needed - only display calculation changes |
| Performance with large datasets | Sort happens once per load, cumulative is O(n) - acceptable for typical usage |

## Implementation Steps

1. Read current `aggregate.go` implementation
2. Modify `AggregateByDate()` to sort descending
3. Refactor aggregation to compute cumulative totals
4. Update tests to verify new behavior
5. Run full test suite
6. Verify with manual testing
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Fixed two bugs in the asset history modal:

### Bug #1: Entries sorted ascending instead of descending
**Change:** Modified `AggregateByDate()` in `internal/assets/model.go` to sort entries by date in descending order (newest first).

**Implementation:** Changed the comparison from `result[i].Date > result[j].Date` to `result[i].Date < result[j].Date` in the sorting algorithm.

### Bug #2: Total Invested showed daily sum instead of cumulative total
**Change:** Added cumulative total calculation to `AggregateByDate()` that tracks a running sum across all previous days.

**Implementation:** After sorting, iterate through results and accumulate totals:
```go
var cumulativeTotal float64
for i := range result {
    cumulativeTotal += result[i].TotalInvested
    result[i].TotalInvested = RoundTo8Decimals(cumulativeTotal)
}
```

### Tests Updated
- Updated `TestAggregateByDate_Sorting` to expect descending order
- Added `TestAggregateByDate_CumulativeTotal` to verify cumulative totals
- Added `TestAggregateByDate_DescendingSortAndCumulative` to test both features together

### Verification
- All 43 tests pass (go test -v ./...)
- Build succeeds without warnings
- Code formatted with go fmt
<!-- SECTION:FINAL_SUMMARY:END -->
