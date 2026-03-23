---
id: GOT-054
title: Bug avg price
status: Done
assignee: []
created_date: '2026-03-21 11:20'
updated_date: '2026-03-23 12:21'
labels: []
dependencies: []
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
There are discrepancies between the average price displayed in the Assets table and the average price displayed in the Asset History Modal. The PRD specifies the correct formula should be: `SUM(price_per_share × amount) / SUM(amounts)`. The current implementation uses `sum(amounts) / sum(shares)` which produces different results.

**Example with real data:**
- BTC: Current=66185.37, PRD Formula=66200.00 (difference: 14.63)
- USDT: Current=5.31, PRD Formula=5.32 (difference: 0.0045)

Fix to show the correct average price as specified in PRD-001-Asset-History-Modal.
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

**Root cause:** The `calculateDayMetrics()` function in `internal/assets/aggregate.go` calculates the weighted average price using `sum(amounts) / sum(shares)`, but PRD-001-Asset-History-Modal (Task 4) specifies it should use `SUM(price_per_share × amount) / SUM(amounts)`.

**How the fix will be implemented:**
1. Update `calculateDayMetrics()` to compute: `sum(price_per_share × amount) / sum(amounts)`
2. Update `aggregateEntries()` to use the same formula for table AvgPrice consistency
3. Update `CalculateWeightedAverage()` helper function with the corrected formula
4. Verify the final row of modal's cumulative weighted average matches table AvgPrice (should be identical)

**Why this approach:**
- The PRD explicitly defines the formula as `SUM(price_per_share × amount) / SUM(amounts)`
- This is the correct weighted average when investment amounts are the weights
- Ensures consistency between table display and modal display

### 2. Files to Modify

| File | Change |
|------|--------|
| `internal/assets/aggregate.go` | Update `calculateDayMetrics()` to use PRD formula |
| `internal/assets/aggregate.go` | Update `aggregateEntries()` to use PRD formula |
| `internal/assets/aggregate.go` | Update `CalculateWeightedAverage()` helper function |
| `internal/assets/model_test.go` | Update tests with correct expected values |
| `internal/assets/aggregate_test.go` | Update tests with correct expected values |

### 3. Dependencies

- **PRD Reference**: PRD-001-Asset-History-Modal (doc-009) - Task 4 specifies the formula
- **Current code**: `internal/assets/aggregate.go` - contains the buggy calculation
- **No blocking dependencies** - can proceed with fix

### 4. Code Patterns

Follow existing patterns in `internal/assets/`:
- Use `RoundTo8Decimals()` for rounding to 8 decimal places
- Maintain the same function signatures and return types
- Update comments to reference the PRD formula

**Formula to implement:**
```go
// Weighted average price (PRD formula): sum(price_per_share × amount) / sum(amounts)
var weightedAvgPrice float64
if totalAmount > 0 {
    weightedAvgPrice = RoundTo8Decimals(sumPriceAmount / totalAmount)
}
```

### 5. Testing Strategy

**Unit tests to verify:**
1. `TestCalculateDayMetrics_PRDFormula` - Verify the PRD formula produces correct results
2. `TestAggregateByDate_PRDCalculations` - Verify daily aggregation with PRD formula
3. `TestLoadAndAggregateEntries_PRDFormula` - Verify table aggregation matches PRD formula
4. Update existing tests to use expected values calculated with the PRD formula

**Test cases:**
- Single entry: price=50000, amount=100 → avg should be 50000
- Multiple entries same price: avg should equal that price
- Multiple entries different prices: weighted by investment amount
- Verify table AvgPrice matches modal's final row WeightedAvgPrice

### 6. Risks and Considerations

**Potential issues:**
1. **Breaking change**: The displayed AvgPrice values will change for existing data
   - Mitigation: This is a bug fix; the previous values were incorrect per PRD

2. **Test updates**: All existing tests will need expected value updates
   - Mitigation: Calculate new expected values using PRD formula before updating tests

3. **Modal cumulative average**: The cumulative average in modal should still match table average on final row
   - Mitigation: Verify with integration test after fix

**Performance:**
- No performance impact - same calculation complexity

**Documentation:**
- Update function comments to reference PRD formula
- Add comment explaining why this formula is correct (weighted average by investment amount)
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Fixed weighted average price calculation bug by implementing PRD formula: `SUM(price_per_share × amount) / SUM(amounts)`

**Changes made:**

1. **internal/assets/aggregate.go:**
   - Updated `aggregateEntries()` to calculate weighted average using PRD formula
   - Updated `CalculateWeightedAverage()` helper to accept `(totalAmount, sumPriceAmount)` and return `sumPriceAmount / totalAmount`
   - All weighted average calculations now use the correct PRD formula

2. **internal/assets/model.go:**
   - Updated `calculateDayMetrics()` to use PRD formula for daily weighted average
   - Updated `AggregateByDate()` to use PRD formula for cumulative weighted average calculation

3. **internal/assets/aggregate_test.go:**
   - Updated `TestCalculateWeightedAverage_*` tests to use PRD formula parameters
   - Updated `TestLoadAndAggregateEntries_SingleAsset` expected AvgPrice: 63030.29 → 63125.0
   - Updated `TestLoadAndAggregateEntries_Calculations_Accurate` expected AvgPrice: 65486.20 → 65432.10

4. **internal/assets/model_test.go:**
   - Updated `TestAggregateByDate_Calculations/mixed_prices_same_day` expected AvgPrice: 50595.20 → 50600.0
   - Updated `TestAggregateByDate_MultipleEntriesPerDay` expected AvgPrice: 49873.59 → 49888.89
   - Added `PricePerShare` to tests that were missing it for PRD formula compatibility
   - Updated `TestCalculateDayMetrics_WeightedAverage` documentation

**Verification:**
- All 139 tests pass (67 in assets package)
- Build succeeds with no warnings
- Code formatted with `go fmt`

**Formula change:**
- Old: `sum(amounts) / sum(shares)` 
- New: `sum(price_per_share × amount) / sum(amounts)` (PRD formula)
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met (no acceptance criteria defined, so N/A)
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
