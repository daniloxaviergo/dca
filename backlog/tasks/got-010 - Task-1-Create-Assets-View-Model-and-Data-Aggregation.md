---
id: GOT-010
title: 'Task 1: Create Assets View Model and Data Aggregation'
status: Done
assignee:
  - Thomas
created_date: '2026-03-17 00:42'
updated_date: '2026-03-17 01:01'
labels: []
dependencies: []
references:
  - 'PRD: DCA Assets List Table View'
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create assets_view.go with AssetSummary struct, AssetsViewModel, and data aggregation logic
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Entries loaded from dca_entries.json correctly
- [x] #2 Grouping by asset ticker works
- [x] #3 Sum of shares calculated per asset
- [x] #4 Weighted average entry price calculated correctly
- [x] #5 Empty file handled gracefully
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Create a new `assets_view.go` file implementing the data aggregation layer for asset summaries. The implementation will follow the existing patterns from `dca_form.go` and `dca_entry.go`:

1. **AssetSummary struct**: Aggregate per-asset data (ticker, count, total shares, average price, total value)
2. **AssetsViewModel**: Manages loaded and aggregated data with state for UI binding
3. **Data aggregation logic**: Read from `dca_entries.json`, group by ticker, calculate metrics using existing helper functions

Key design decisions:
- Use existing `DCAData` and `DCAEntry` types (no structural changes)
- Weighted average price = sum(amount) / sum(shares) per asset
- 8-decimal precision for shares and price calculations
- Empty entries map handled gracefully (returns empty slice, not error)

### 2. Files to Modify

**New file to create:**
- `assets_view.go` - AssetSummary struct, AssetsViewModel, LoadAndAggregateEntries()

**No existing files to modify for this task** (Task 1 only - data aggregation layer)

**Dependencies:**
- Existing `dca_entry.go` - DCAEntry, DCAData, LoadEntries() functions
- Existing `dca_form.go` - CalculateSharesFromValues(), RoundTo8Decimals() patterns

### 3. Dependencies

**Prerequisites for implementation:**
- [x] `dca_entry.go` with DCAData structure and LoadEntries() function (already exists)
- [x] CalculateSharesFromValues() and RoundTo8Decimals() utility functions (already in dca_form.go)
- [x] Existing test patterns from `dca_entry_test.go` and `dca_form_test.go`

**No blocking issues** - all required data structures and I/O functions already exist.

**Setup steps:**
- None required - uses existing file I/O and data structures

### 4. Code Patterns

Follow existing project conventions from `dca_form.go` and `dca_entry.go`:

**Struct patterns:**
```go
type AssetSummary struct {
    Ticker      string
    EntryCount  int
    TotalShares float64
    AvgPrice    float64
    TotalValue  float64
}

type AssetsViewModel struct {
    Entries     []AssetSummary
    Error       error
    SelectedIdx int
}
```

**Validation patterns:**
- Return descriptive error messages matching existing format
- Handle nil/empty maps gracefully
- Use existing `LoadEntries()` for file reading

**Calculation patterns:**
- Weighted average price: `sum(amounts) / sum(shares)`
- Round to 8 decimals using existing `RoundTo8Decimals()`
- Handle division by zero for empty asset groups

**Naming conventions:**
- Functions: `LoadAndAggregateEntries()`, `CalculateWeightedAverage()`
- Structs: `AssetSummary`, `AssetsViewModel`
- Methods: receiver names match type (e.g., `func (vm *AssetsViewModel) ...`)

### 5. Testing Strategy

Create `assets_view_test.go` with table-driven tests following existing patterns:

**Test categories:**
1. **LoadAndAggregateEntries_Populated**: Verify correct aggregation with known data
2. **LoadAndAggregateEntries_EmptyFile**: Handle empty/missing file gracefully
3. **LoadAndAggregateEntries_MultipleAssets**: Group multiple assets correctly
4. **LoadAndAggregateEntries_MultipleEntries**: Sum shares and calculate weighted average
5. **AssetSummary_CalculateWeightedAverage**: Edge cases (zero shares, single entry)
6. **AssetsViewModel_Selection**: Selected index management (if applicable)

**Test coverage:**
- Zero entries → empty slice returned
- Single asset with multiple entries → correct aggregation
- Multiple assets → correct grouping
- Weighted average calculation accuracy
- 8-decimal precision maintained

**Edge cases:**
- Empty entries map in DCAData
- Asset with zero shares (should not occur but handle gracefully)
- Very large numbers (no overflow issues)
- Division by zero in weighted average (return 0 or skip)

### 6. Risks and Considerations

**No significant risks identified** for Task 1:

- Data structures already exist (DCAEntry, DCAData)
- File I/O already implemented (LoadEntries in dca_entry.go)
- Calculation logic is straightforward (sum, count, division)
- No new external dependencies required

**Design considerations:**
- Weighted average formula: `sum(amounts) / sum(shares)` not `avg(prices)`
- Shares from entries may have rounding; use stored values directly
- PRD defines "Total Value" as sum of entry amounts (USD invested), not market value

**Trade-offs:**
- No sorting in Task 1 (deferred to Task 2 table UI)
- No error UI in Task 1 (only set Error field for UI to handle)
- No pagination (simple slice, acceptable for <1000 assets per PRD)
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation complete. Created assets_view.go with AssetSummary struct, AssetsViewModel, and LoadAndAggregateEntries function. All acceptance criteria verified through comprehensive test coverage (20 test cases). All tests pass, no compiler warnings, code formatted with go fmt. PRD reference already present in task metadata.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Task GOT-010 completed successfully. Created the assets view model and data aggregation layer for the DCA tracking application.

## What Changed

### New Files Created
1. **assets_view.go** - Implements:
   - `AssetSummary` struct with ticker, entry count, total shares, average price, and total value
   - `AssetsViewModel` struct to manage loaded and aggregated data
   - `LoadAndAggregateEntries()` - Loads entries from JSON, groups by ticker, calculates metrics
   - `CalculateWeightedAverage()` - Computes weighted average price (sum(amounts) / sum(shares))
   - `Validate()` method for AssetSummary validation

2. **assets_view_test.go** - Comprehensive test coverage with 20 test cases:
   - `TestAssetSummary_Validate_*` - 6 tests for validation
   - `TestCalculateWeightedAverage_*` - 3 tests for weighted average calculation
   - `TestRoundTo8Decimals_RoundsCorrectly` - 8 decimal precision tests
   - `TestLoadAndAggregateEntries_*` - 9 tests for aggregation (empty file, missing file, single/multiple assets, multiple entries)
   - `TestAssetsViewModel_Initialization` - 1 test for initialization
   - `TestLoadAndAggregateEntries_Calculations_Accurate` - 1 test for accuracy verification

## Verification

- All 20 test cases pass ✓
- `go test ./...` passes for entire project ✓
- `go build` completes without errors ✓
- `go fmt` applied successfully ✓
- No compiler warnings ✓

## Design Decisions

1. **Weighted average formula**: Used `sum(amounts) / sum(shares)` per PRD definition
2. **Precision**: 8-decimal rounding for all financial calculations
3. **Graceful handling**: Empty/missing files return empty slice, not errors
4. **Code reuse**: Leverages existing `LoadEntries()` and `RoundTo8Decimals()` functions
5. **No sorting**: Deferred to Task 2 (table UI)

## Risks and Follow-ups

- Task 2 will add table UI with sorting capability
- No pagination implemented (acceptable for <1000 assets per PRD)
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
