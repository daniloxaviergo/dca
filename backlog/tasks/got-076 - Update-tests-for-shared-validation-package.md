---
id: GOT-076
title: Update tests for shared validation package
status: To Do
assignee: []
created_date: '2026-03-28 15:10'
labels:
  - testing
  - validation
  - phase2
  - req-009
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
Update existing tests to remove redundant validation tests and add new tests for the shared validation package.

## Implementation Details

### 1. Remove Redundant Tests from internal/form/validation_test.go

**Remove these test functions (validation now lives in validate package):**
- `TestFormModel_ValidateAmount_*` (all amount validation tests)
- `TestFormModel_ValidateDate_*` (all date validation tests)
- `TestFormModel_ValidateAsset_*` (all asset validation tests)
- `TestFormModel_ValidatePrice_*` (all price validation tests)

**Keep these tests (not validation-related):**
- `TestCalculateSharesFromValues`
- `TestCalculateSharesFromValues_Precision`
- `TestRoundTo8Decimals`
- `TestFormModel_*` (all other non-validation tests)

### 2. Create New Tests for internal/form/validate/

**New File: internal/form/validate/validate_test.go**

**Test Coverage Requirements:**

```go
// Validation Tests
func TestValidateAmount_Pass(t *testing.T)
func TestValidateAmount_RejectZero(t *testing.T)
func TestValidateAmount_RejectNegative(t *testing.T)
func TestValidateAmount_RejectEmpty(t *testing.T)
func TestValidateAmount_RejectInvalid(t *testing.T)
func TestValidateAmount_ExactErrorMessage(t *testing.T)  // "Amount must be positive"

func TestValidateDate_Pass(t *testing.T)
func TestValidateDate_RejectInvalid(t *testing.T)
func TestValidateDate_ExactErrorMessage(t *testing.T)  // "Use YYYY-MM-DD"

func TestValidateAsset_Pass(t *testing.T)
func TestValidateAsset_RejectEmpty(t *testing.T)
func TestValidateAsset_RejectWhitespace(t *testing.T)
func TestValidateAsset_ExactErrorMessage(t *testing.T)  // "Asset ticker is required"

func TestValidatePrice_Pass(t *testing.T)
func TestValidatePrice_RejectZero(t *testing.T)
func TestValidatePrice_RejectNegative(t *testing.T)
func TestValidatePrice_RejectEmpty(t *testing.T)
func TestValidatePrice_RejectInvalid(t *testing.T)
func TestValidatePrice_ExactErrorMessage(t *testing.T)  // "Price must be positive"

// Shares Calculation Tests
func TestCalculateShares_Pass(t *testing.T)
func TestCalculateShares_DivisionByZero(t *testing.T)
func TestCalculateShares_NaNHandling(t *testing.T)
func TestCalculateShares_InfHandling(t *testing.T)
func TestCalculateShares_Precision(t *testing.T)
func TestCalculateShares_IdenticalToOriginal(t *testing.T)  // Compare to original algorithm

func TestRoundTo8Decimals_Pass(t *testing.T)
func TestRoundTo8Decimals_Negative(t *testing.T)
func TestRoundTo8Decimals_ExactRounding(t *testing.T)
```

### 3. Update Test Coverage Strategy

**Coverage Targets:**
- **Line coverage**: 100% for validate package
- **Edge cases covered**: All error paths, special values
- **Backward compatibility verified**: Existing tests still pass

### 4. Integration Test Strategy

**Verify shared validation works:**
1. Run existing form tests (non-validation) - must pass
2. Run validate package tests - must pass
3. Compare outputs: FormModel validation results == validate package results

## Acceptance Criteria
<!-- AC:BEGIN -->
- ✅ Redundant validation tests removed from validation_test.go
- ✅ New validate tests created with 100% coverage
- ✅ All existing FormModel tests pass
- ✅ Validation output matches between old and new
- ✅ Edge cases covered (empty, zero, negative, NaN, Inf)
<!-- SECTION:DESCRIPTION:END -->

- [ ] #1 Redundant validation tests removed from validation_test.go
- [ ] #2 New validate package tests created with comprehensive coverage
- [ ] #3 All existing FormModel tests pass without modification
- [ ] #4 Validation output identical between original and new implementation
- [ ] #5 100% line coverage target for validate package
- [ ] #6 Edge cases covered (empty, zero, negative, NaN, Inf, division by zero)
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
