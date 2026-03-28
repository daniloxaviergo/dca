---
id: GOT-081
title: >-
  [doc-013-cli-02] Add CalculateSharesFromValues and RoundTo8Decimals to
  internal/dca/entry.go
status: To Do
assignee: []
created_date: '2026-03-28 15:14'
labels:
  - phase-3
  - calculation
  - dca
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add calculation functions to internal/dca/entry.go to support CLI mode's 8-decimal precision share calculations.

Tasks:
1. Add CalculateSharesFromValues(amount float64, price float64) float64 function
   - Uses math.Round((amount/price)*1e8)/1e8 formula
   - Returns 0 if price is 0 (to avoid division by zero)
2. Add RoundTo8Decimals(value float64) float64 helper function
   - General-purpose rounding utility for 8-decimal precision
3. Update DCAEntry.CalculateShares() to use the new shared function
4. Add comprehensive unit tests for both functions
5. Test edge cases: zero price, negative values, NaN/Inf
6. Verify rounding behavior matches TUI form exactly

Test file: internal/dca/entry_test.go (add new test functions)
Reference: PRD R6 - Share precision with 8 decimal places matching TUI behavior
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 CalculateSharesFromValues function added with 8-decimal precision
- [ ] #2 RoundTo8Decimals helper function added
- [ ] #3 DCAEntry.CalculateShares() uses shared function
- [ ] #4 All edge cases handled (zero price, NaN, Inf)
- [ ] #5 Test coverage includes all rounding scenarios
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
