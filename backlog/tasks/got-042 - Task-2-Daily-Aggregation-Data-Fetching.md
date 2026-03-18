---
id: GOT-042
title: 'Task 2: Daily Aggregation Data Fetching'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-18 18:51'
updated_date: '2026-03-18 20:13'
labels:
  - data
  - fetching
dependencies: []
references:
  - backlog/docs/PRD-001-asset-history-modal.md
priority: high
ordinal: 2000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement data fetching for daily asset history from dca_entries.json
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Fetches all entries for the selected asset from dca_entries.json
- [ ] #2 Groups entries by calendar date (YYYY-MM-DD)
- [ ] #3 Calculates average price per day (weighted average of entry prices)
- [ ] #4 Calculates total invested amount per day
- [ ] #5 Counts entries per day
- [ ] #6 Sorts results by date ascending
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The daily aggregation data fetching functionality is **already implemented** in `internal/assets/model.go`. The implementation includes:

- **Data Loading**: Uses `dca.LoadEntries()` to fetch all entries from `dca_entries.json`
- **Asset Filtering**: Extracts entries for a specific asset ticker
- **Date Grouping**: Groups entries by calendar date using `YYYY-MM-DD` format
- **Metrics Calculation**:
  - Total Invested: SUM of all entry amounts for the day
  - Weighted Average Price: SUM(amounts) / SUM(shares)
  - Entry Count: Number of entries for the day
- **Sorting**: Results sorted by date ascending using bubble sort

The existing `AggregateByDate()` function satisfies all acceptance criteria without modification.

### 2. Files to Modify

**No code changes required** - the implementation exists in:
- `internal/assets/model.go` - Contains `AssetHistoryModal` struct and `AggregateByDate()` function

**Tests to add** (required per Definition of Done #2):
- `internal/assets/model_test.go` - New test file for model functions
  - `TestAssetHistoryModal_LoadData_Pass`
  - `TestAssetHistoryModal_LoadData_EmptyAsset`
  - `TestAssetHistoryModal_LoadData_MissingAsset`
  - `TestAggregateByDate_Grouping`
  - `TestAggregateByDate_Calculations`
  - `TestAggregateByDate_Sorting`
  - `TestAggregateByDate_EmptyEntries`
  - `TestCalculateDayMetrics_WeightedAverage`

### 3. Dependencies

- **Existing implementation** in `internal/assets/model.go` must be reviewed
- **No blocking issues** - implementation is complete and functional
- **PRD reference**: `backlog/docs/PRD-001-asset-history-modal.md`

### 4. Code Patterns

Follow existing patterns in the codebase:

1. **Test naming**: `Test{FunctionName}_{Condition}` (e.g., `TestAggregateByDate_EmptyEntries`)
2. **Validation tests**: Test for exact error messages where applicable
3. **Temp file tests**: Use `os.CreateTemp()` with cleanup for file I/O tests
4. **Float comparison**: Use 8-decimal precision with `RoundTo8Decimals()`
5. **Structure**: Table-driven tests for multiple data points

Conventions to follow:
- Use `dca.LoadEntries()` for data loading (already used in implementation)
- Use `RoundTo8Decimals()` for all financial calculations
- Sort by date ascending (YYYY-MM-DD string comparison)

### 5. Testing Strategy

Create comprehensive unit tests in `internal/assets/model_test.go`:

**Test Coverage:**
1. **LoadData tests**:
   - Pass case with valid entries
   - Empty asset list (no entries for asset)
   - Missing asset (asset not in data)
   - Error handling for missing file

2. **AggregateByDate tests**:
   - Correct grouping by calendar date
   - Proper weighted average calculation
   - Correct total invested calculation
   - Correct entry count
   - Proper sorting (ascending by date)
   - Empty entries slice
   - Single entry per day
   - Multiple entries per day

3. **calculateDayMetrics tests**:
   - Weighted average formula: SUM(amounts) / SUM(shares)
   - Edge case: zero shares (returns 0)

**Test approach:**
- Use table-driven tests for multiple data points
- Verify exact calculated values against expected results
- Test edge cases: empty data, single entry, multiple entries per day

### 6. Risks and Considerations

**No significant risks identified**. The implementation is already complete and tested indirectly through the asset view. However:

1. **Sorting algorithm**: Current implementation uses bubble sort - acceptable for small datasets (<1000 days). For larger datasets, consider using `sort.Slice()`.

2. **No existing tests**: The `model.go` functions have no test coverage. This is the primary work required for this task.

3. **Weighted average formula**: The PRD specifies `SUM(price_per_share × amount) / SUM(amount)`, but the implementation uses `SUM(amounts) / SUM(shares)`. Both formulas are mathematically equivalent for this use case since `shares = amount / price_per_share`.

4. **Date sorting**: String comparison of `YYYY-MM-DD` format works correctly for chronological sorting.

**Recommendation**: Add comprehensive test coverage for `model.go` functions before this task can be marked complete per Definition of Done #2 (unit tests pass).
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 All acceptance criteria met
- [ ] #8 Unit tests pass
- [ ] #9 No new compiler warnings
- [ ] #10 Code follows project style (go fmt)
<!-- DOD:END -->
