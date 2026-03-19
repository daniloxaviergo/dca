---
id: GOT-052
title: Bug Assert view
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-19 16:45'
updated_date: '2026-03-19 17:06'
labels: []
dependencies: []
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The sort of entries on modal assets history should be desc
The sum should be the sum of previues days
<!-- SECTION:DESCRIPTION:END -->

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

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
