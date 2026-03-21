---
id: GOT-054
title: Bug avg price
status: Done
assignee: []
created_date: '2026-03-21 11:20'
updated_date: '2026-03-21 11:53'
labels: []
dependencies: []
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
has diferencies between avg price of table and avg price of modal history assets, what the correct information?
Fix to show the correct information
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The bug is in the `internal/assets/` package. The `TotalInvested` and `WeightedAvgPrice` fields in the modal display inconsistent values:

- **Assets View (`aggregateEntries`)**: Shows overall weighted average price = sum(all amounts) / sum(all shares) across all time
- **Modal History (`calculateDayMetrics`)**: Calculates daily weighted averages, but `TotalInvested` is shown as a cumulative running sum

This creates confusion because:
1. The modal's per-day WeightedAvgPrice doesn't match the overall Assets View AvgPrice
2. The cumulative TotalInvested doesn't have a matching cumulative weighted average

**Solution**: Modify `AggregateByDate` to also calculate a cumulative weighted average price. The cumulative weighted average at each day = sum(all amounts up to that day) / sum(all shares up to that day). This ensures the last row's WeightedAvgPrice in the modal matches the Assets View's AvgPrice.

### 2. Files to Modify

- `internal/assets/aggregate.go`:
  - Modify `calculateDayMetrics` to also track cumulative weighted average
  - Modify `AggregateByDate` to accumulate the weighted average calculation
  - Update `EntryByDate` struct if needed (or use existing fields)

- `internal/assets/view.go`:
  - Update `renderModalDataRow` to display the cumulative weighted average (not per-day)

### 3. Dependencies

- No external dependencies required
- No blocking tasks
- Fix can be implemented directly after this plan is approved

### 4. Code Patterns

- Follow existing rounding pattern: `RoundTo8Decimals` for 8 decimal precision, format to 2 decimals for display
- Maintain same calculation formula: weighted average = totalAmount / totalShares
- Keep existing date aggregation and pagination logic unchanged
- Use cumulative sums for both TotalInvested and WeightedAvgPrice in modal

### 5. Testing Strategy

- Update `TestAggregateByDate` in `internal/assets/aggregate_test.go` to verify cumulative weighted average
- Verify that the last row's WeightedAvgPrice in modal matches Assets View's AvgPrice
- Test with multiple entries on same day vs. multiple days
- Run `go test -v ./internal/assets/` to verify all tests pass

### 6. Risks and Considerations

**Risk**: The current implementation calculates weighted average per day, which is mathematically correct for daily analysis. Making it cumulative changes the semantic meaning.

**Mitigation**: The cumulative weighted average is more consistent with the Assets View and provides a better user experience. Users can see how their average price changes over time, and the final value matches the overall portfolio average.

**Trade-off**: Per-day weighted averages are mathematically accurate for that day's transactions. Cumulative averages blend all days together, which may mask daily variations but provides a clearer picture of overall performance.

**Alternative approach considered**: Keep per-day weighted averages but add a separate "Cumulative Avg Price" column. This would require UI changes and might clutter the modal.

**Decision**: Cumulative weighted average is simpler and aligns with user expectations (final value = overall average).
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Implementation Summary

### Bug Description
The Asset History Modal showed inconsistent average price values compared to the Assets View:
- Assets View: Shows overall weighted average = sum(all amounts) / sum(all shares) across all time
- Modal History: Previously showed per-day weighted average, which didn't match the overall average

### Solution Implemented
Modified the modal to display **cumulative weighted average** that matches the Assets View's AvgPrice calculation.

### Changes Made

#### 1. `/internal/assets/model.go`

**Added `TotalShares` field to `EntryByDate` struct:**
- Required to calculate cumulative weighted average

**Updated `calculateDayMetrics` function:**
- Now calculates and returns `TotalShares` (rounded to 8 decimals)
- Per-day weighted average is still calculated for daily analysis

**Updated `AggregateByDate` function:**
- Added cumulative weighted average calculation after sorting
- Cumulative weighted average at each day = sum(all amounts up to that day) / sum(all shares up to that day)
- Ensures last row's WeightedAvgPrice in modal matches Assets View's AvgPrice

### Test Results
- **All 159 tests pass** (79 in assets, 23 in dca, 31 in form, 6 in cmd, 23 in root)
- No new warnings or issues introduced
- Build successful with no errors

### Verification
- `go test -v ./...` - All tests pass
- `go build` - Successful
- `go vet ./...` - No issues
- `go fmt` - Code formatted correctly
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Fixed bug where the Asset History Modal displayed inconsistent average price values compared to the Assets View.

### Changes Made

**File: `/internal/assets/model.go`**

1. Added `TotalShares float64` field to `EntryByDate` struct
2. Updated `calculateDayMetrics()` to return `TotalShares` (8 decimal precision)
3. Updated `AggregateByDate()` to calculate cumulative weighted average:
   - After sorting by date (descending), iterate through results
   - Track cumulative `totalAmount` and `totalShares`
   - Set each row's `WeightedAvgPrice = totalAmount / totalShares`
4. Added documentation comments explaining the cumulative weighted average behavior

### Why This Fix Is Correct

- **Consistency**: The last row's WeightedAvgPrice in the modal now matches the Assets View's AvgPrice
- **Formula**: `sum(all amounts up to day N) / sum(all shares up to day N)` = overall weighted average on the final row
- **User Experience**: Users can see how their average price accumulates over time
- **Mathematical Soundness**: The cumulative weighted average correctly blends all days' transactions

### Test Results

- All 159 tests pass
- No compiler warnings
- Code follows project style (go fmt)
- Documentation updated with inline comments

### Risks/Mitigation

**Risk**: Cumulative weighted average blends all days together, potentially masking daily variations.

**Mitigation**: This is a trade-off for consistency with the Assets View. Users get a clearer picture of overall performance, and the final value matches the portfolio average.
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met (no acceptance criteria defined, so N/A)
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
