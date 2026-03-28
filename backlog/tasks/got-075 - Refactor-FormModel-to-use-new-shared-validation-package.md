---
id: GOT-075
title: Refactor FormModel to use new shared validation package
status: To Do
assignee: []
created_date: '2026-03-28 15:10'
labels:
  - refactoring
  - form
  - phase2
  - req-009
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
Refactor FormModel to use the new shared validation functions while maintaining backward compatibility with existing tests.

## Implementation Details

### Changes Required

1. **Update FormModel methods to delegate to shared validation**
   - Modify existing methods to call validate package functions:
     - `form.validateAmount()` → `validate.ValidateAmount()`
     - `form.validateDate()` → `validate.ValidateDate()`
     - `form.validateAsset()` → `validate.ValidateAsset()`
     - `form.validatePrice()` → `validate.ValidatePrice()`
   - Keep method signatures identical for backward compatibility
   - Preserve all existing error handling behavior

2. **Update shares calculation usage**
   - Replace `CalculateSharesFromValues()` calls with `validate.CalculateShares()`
   - Replace `RoundTo8Decimals()` calls with `validate.RoundTo8Decimals()`
   - Maintain identical mathematical behavior

3. **Update share calculation within FormModel**
   - In `handleEnter()` for StepPrice, use new validate package
   - In `saveEntry()`, use new validate package
   - In `renderForm()`, use new validate package for dynamic share calculation

## File Changes

**internal/form/model.go:**
```go
// Add import
import (
    // ... existing imports
    "github.com/danilo/scripts/github/dca/internal/form/validate"
)

// Keep existing methods but change implementation
func (m *FormModel) validateAmount(value string) error {
    return validate.ValidateAmount(value)
}
func (m *FormModel) validateDate(value string) error {
    return validate.ValidateDate(value)
}
// ... etc for all validation methods

// Update CalculateSharesFromValues to delegate
func CalculateSharesFromValues(amount, price float64) float64 {
    return validate.CalculateShares(amount, price)
}

// Update RoundTo8Decimals to delegate
func RoundTo8Decimals(val float64) float64 {
    return validate.RoundTo8Decimals(val)
}
```

## Requirements
- **No breaking changes**: All existing tests must pass unchanged
- **Identical behavior**: Validation results must be identical
- **Minimal changes**: Only implementation details change
- **Backward compatibility**: Existing test calls still work

## Acceptance Criteria
<!-- AC:BEGIN -->
- ✅ All existing tests pass without modification
- ✅ Validation output identical to original
- ✅ No API changes to FormModel
- ✅ No changes to external callers of FormModel methods
<!-- SECTION:DESCRIPTION:END -->

- [ ] #1 FormModel validate methods delegate to validate package functions
- [ ] #2 CalculateSharesFromValues and RoundTo8Decimals delegate to validate package
- [ ] #3 All existing tests pass without modification
- [ ] #4 Validation output identical to original implementation
- [ ] #5 No API changes to FormModel or external callers
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
