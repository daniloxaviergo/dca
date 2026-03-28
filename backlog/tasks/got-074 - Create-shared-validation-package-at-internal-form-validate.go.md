---
id: GOT-074
title: Create shared validation package at internal/form/validate.go
status: To Do
assignee: []
created_date: '2026-03-28 15:09'
labels:
  - refactoring
  - validation
  - phase2
  - req-009
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
Extract validation logic from FormModel methods and shares calculation logic to a new shared validation package at `internal/form/validate.go`.

## Implementation Details

### New File: internal/form/validate.go
Create new file with independent validation functions (no FormModel dependency):

**Validation Functions:**
1. `ValidateAmount(value string) error`
   - Validates amount is positive (> 0)
   - Error message: "Amount must be positive"
   - Handles empty strings and invalid numbers

2. `ValidatePrice(value string) error`
   - Validates price is positive (> 0)
   - Error message: "Price must be positive"
   - Handles empty strings and invalid numbers

3. `ValidateAsset(value string) error`
   - Validates asset ticker is non-empty and not whitespace-only
   - Error message: "Asset ticker is required"

4. `ValidateDate(value string) error`
   - Validates RFC3339 date format
   - Error message: "Use YYYY-MM-DD"
   - Handles empty strings

**Shares Calculation Functions:**
5. `CalculateShares(amount, price float64) float64`
   - Implements: `math.Round((amount / price) * 1e8) / 1e8`
   - Handles division by zero (returns 0)
   - Handles NaN/Inf cases (returns 0)

6. `RoundTo8Decimals(val float64) float64`
   - Rounds to 8 decimal places using standard rounding
   - Handles edge cases (zero, negative, very small values)

### Requirements
- Functions must be package-level, not methods
- Must not depend on FormModel or any TUI state
- Must preserve exact behavior and error messages of originals
- All functions should be testable in isolation

## Acceptance Criteria
<!-- AC:BEGIN -->
- ✅ validate.go created with all functions
- ✅ Functions are independent of FormModel
- ✅ Error messages match original exactly
- ✅ Mathematical precision maintained (8 decimals)
- ✅ All edge cases handled identically to original
<!-- SECTION:DESCRIPTION:END -->

- [ ] #1 internal/form/validate.go created with all validation and calculation functions
- [ ] #2 Functions are package-level and independent of FormModel
- [ ] #3 Error messages match original validation.go exactly
- [ ] #4 CalculateShares produces identical results to CalculateSharesFromValues
- [ ] #5 All edge cases (zero, negative, empty, division by zero) handled identically
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
