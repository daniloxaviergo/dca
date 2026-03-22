---
id: GOT-054
title: Bug avg price
status: In Progress
assignee: []
created_date: '2026-03-21 11:20'
updated_date: '2026-03-22 23:08'
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

The issue arises because the asset history modal incorrectly calculates the average price for each day by averaging the `pricePerShare` values directly, whereas the main assets table uses the correct weighted average (total amount / total shares). The fix involves updating the daily aggregation logic in the asset history modal to consistently use the weighted average calculation (sum of amounts divided by sum of shares per day).

### 2. Files to Modify

- `internal/assets/view.go`: Modify the asset history data processing function (`GetAssetHistory` or equivalent) to calculate daily weighted average price correctly.
- `internal/assets/aggregate.go`: Refactor the `CalculateDailyAggregation` function to standardize the weighted average calculation logic used in both tables.

### 3. Dependencies

- None. Existing project structure and data models are sufficient.

### 4. Code Patterns

- Follow existing patterns in `internal/assets/aggregate.go` for financial calculations.
- Use consistent rounding rules (8 decimal places for shares, 2 decimals for display).
- Ensure calculations use `float64` with precise arithmetic to avoid floating-point errors.

### 5. Testing Strategy

- Add unit tests in `internal/assets/aggregate_test.go` for daily aggregation scenarios.
- Test edge cases:
  - Single entry per day
  - Multiple entries with same/different prices per day
  - Zero amounts or shares
- Verify that modal's Avg Price matches the main table's calculation for the same asset.

### 6. Risks and Considerations

- No significant risks. The fix is isolated to daily aggregation logic.
- Ensure the main assets table calculation remains unchanged.
- Verify that historical data calculations remain consistent after the fix.
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met (no acceptance criteria defined, so N/A)
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
