---
id: GOT-079
title: Create rollback plan and verify Phase 2 acceptance criteria
status: To Do
assignee: []
created_date: '2026-03-28 15:12'
labels:
  - testing
  - acceptance
  - phase2
  - req-009
dependencies: []
priority: low
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
Verify Phase 2 acceptance criteria are met with integration and end-to-end testing.

## Implementation Details

### 1. Rollback Plan

**Before making changes:**
```bash
# Create branch for Phase 2
git checkout -b phase2-refactoring

# Create backup
cp internal/form/validation.go internal/form/validation.go.backup
```

**If issues detected:**
```bash
# Revert to previous state
git checkout HEAD -- internal/form/validation.go internal/form/model.go

# Restore backup if needed
cp internal/form/validation.go.backup internal/form/validation.go

# Verify TUI still works
make test
./dca
```

**Recovery steps:**
1. Run `make test` to verify no regressions
2. If tests fail, check git status for changes
3. Revert to backup if necessary
4. Report issue before proceeding

### 2. Acceptance Criteria Verification

**NFA-002: Shared validation logic - no code duplication**

**Verification steps:**
```bash
# 1. Check no duplicate validation logic
grep -c "must be positive" internal/form/validate/validate.go  # Should be 1
grep -c "must be positive" internal/form/validation.go          # Should be 0 after refactor

# 2. Verify single source of truth
grep "func ValidateAmount" internal/form/validate/validate.go  # Should exist
grep "func.*ValidateAmount" internal/form/  # Only in validate.go

# 3. Test identical behavior
go test -v ./internal/form/validate/...
go test -v ./internal/form/... -run TestCalculate
```

**Existing tests pass unchanged:**
```bash
# All tests must pass
make test
# Or: go test -v ./...

# Verify specific test categories
go test -v ./internal/form/... -run "TestFormModel"  # Non-validation tests
go test -v ./internal/form/validate/...              # New validation tests
go test -v ./internal/dca/...                        # Data model tests
go test -v ./cmd/dca/...                             # Integration tests
```

**Validation output identical:**
```bash
# Test with real data
./dca --test-validation
# Or use test helper
go test -v -run TestValidationOutputIdentical
```

### 3. Integration Test Suite

**New test file: internal/form/integration_test.go**

```go
package form

import (
    "testing"
    "github.com/danilo/scripts/github/dca/internal/form/validate"
)

// TestFormValidationIdenticalToValidatePackage
func TestFormValidationIdenticalToValidatePackage(t *testing.T) {
    tests := []struct {
        name     string
        value    string
        validator func(string) error
        pkgFunc  func(string) error
    }{
        {"Amount", "100", (*FormModel).validateAmount, validate.ValidateAmount},
        {"Amount", "0", (*FormModel).validateAmount, validate.ValidateAmount},
        {"Amount", "abc", (*FormModel).validateAmount, validate.ValidateAmount},
        
        {"Date", "2024-01-15T10:30:00Z", (*FormModel).validateDate, validate.ValidateDate},
        {"Date", "invalid", (*FormModel).validateDate, validate.ValidateDate},
        
        {"Asset", "BTC", (*FormModel).validateAsset, validate.ValidateAsset},
        {"Asset", "", (*FormModel).validateAsset, validate.ValidateAsset},
        
        {"Price", "50000", (*FormModel).validatePrice, validate.ValidatePrice},
        {"Price", "0", (*FormModel).validatePrice, validate.ValidatePrice},
    }
    
    entries := &DCAData{Entries: make(map[string][]DCAEntry)}
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            form := NewFormModel(entries, "test.json")
            
            formErr := tt.validator(form, tt.value)
            pkgErr := tt.pkgFunc(tt.value)
            
            if formErr == nil && pkgErr != nil {
                t.Errorf("Form validation passed but package validation failed for %s", tt.name)
            }
            if formErr != nil && pkgErr == nil {
                t.Errorf("Form validation failed but package validation passed for %s", tt.name)
            }
            if formErr != nil && pkgErr != nil && formErr.Error() != pkgErr.Error() {
                t.Errorf("Different error messages for %s: form=%s, pkg=%s", tt.name, formErr, pkgErr)
            }
        })
    }
}

// TestSharesCalculationIdentical tests matches between CalculateSharesFromValues and calculate package
func TestSharesCalculationIdentical(t *testing.T) {
    tests := []struct {
        amount float64
        price  float64
    }{
        {100, 50},
        {100, 33.33333333},
        {0, 100},
        {100, 0},  // division by zero
    }
    
    for _, tt := range tests {
        formResult := CalculateSharesFromValues(tt.amount, tt.price)
        pkgResult := validate.CalculateShares(tt.amount, tt.price)
        
        if formResult != pkgResult {
            t.Errorf("Shares calculation mismatch for amount=%f, price=%f: form=%f, pkg=%f", 
                tt.amount, tt.price, formResult, pkgResult)
        }
    }
}
```

### 4. Phase 2 Final Validation

**Before merging:**
```bash
# 1. All tests pass
make test

# 2. Coverage check
make test-cover

# 3. Check no duplication
grep -r "Amount must be positive" internal/form/validate/ | wc -l  # Should be 1
grep -r "Amount must be positive" internal/form/validation.go | wc -l  # Should be 0

# 4. Build succeeds
make build

# 5. TUI still works (manual check)
./dca
```

**Rollback if:**
- Any existing tests fail
- Validation behavior differs from original
- No Duplication verification fails
- CLI integration introduces regressions

### 5. Rollback checklist

- [ ] All original tests pass
- [ ] TUI functionality verified (form entry, submission, navigation)
- [ ] No validation logic duplicated between files
- [ ] Error messages identical in single source of truth
- [ ] CLI validation matches TUI validation
- [ ] 100% line coverage for validate package

## Acceptance Criteria
<!-- AC:BEGIN -->
- ✅ Rollback plan created and documented
- ✅ All existing tests pass
- ✅ No code duplication between validation.go and validate.go
- ✅ Validation output identical between CLI and TUI
- ✅ 100% line coverage for validate package
<!-- SECTION:DESCRIPTION:END -->

- [ ] #1 Rollback plan documented with specific steps and commands
- [ ] #2 All existing tests pass without modification
- [ ] #3 No validation logic duplicated between files
- [ ] #4 CLI and TUI validation identical (verified by tests)
- [ ] #5 100% line coverage for validate package
- [ ] #6 Phase 2 acceptance criteria verified before merging
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
