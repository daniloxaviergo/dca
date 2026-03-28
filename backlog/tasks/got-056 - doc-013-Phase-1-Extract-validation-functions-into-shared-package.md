---
id: GOT-056
title: '[doc-013 Phase 1] Extract validation functions into shared package'
status: To Do
assignee: []
created_date: '2026-03-28 20:49'
updated_date: '2026-03-28 22:29'
labels:
  - feature
  - refactoring
  - validation
dependencies: []
references:
  - 'doc-013 - Phase 1: Extract validation logic'
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract validation functions from internal/form/validation.go to enable reuse between TUI form and CLI entry. Create shared validation logic for amount (positive), price (positive), and asset (non-empty) validation with descriptive error messages. Ensure functions are_exported for use across packages and maintain the same validation rules and precision (8 decimals for shares calculation).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Validation functions extracted with unit tests
- [ ] #2 Shared validation logic tested with existing TUI form
- [ ] #3 Functions maintain 8 decimal precision for shares calculation
- [ ] #4 Error messages match PRD specification
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Extract validation functions from `internal/form/validation.go` into a new shared package `internal/validation` to enable reuse between TUI form and CLI entry. The approach:

- Create new package `internal/validation` with exported validation functions
- Extract 4 validation functions: `IsValidAmount`, `IsValidPrice`, `IsValidAsset`, `IsValidDate`
- Preserve exact error messages from existing validation
- Add tests for all validation functions
- Update TUI form to use shared validation functions
- Maintain 8 decimal precision via existing `RoundTo8Decimals` in form package
- No breaking changes to existing functionality

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/validation/validation.go` | Create | New shared validation package with exported functions |
| `internal/validation/validation_test.go` | Create | Unit tests for all validation functions |
| `internal/form/validation.go` | Modify | Refactor to use shared validation functions |
| `internal/form/model.go` | Modify | Update imports to use shared validation |

### 3. Dependencies

- **No blocking dependencies**: This task can proceed independently
- **Prerequisites**: Current form structure in place (already implemented)
- **Related tasks**: 
  - GOT-057 (CLI flag parsing) will depend on this shared validation
  - GOT-058 (CLI run function) will use shared validation
  - Existing tests must continue to pass

### 4. Code Patterns

Follow existing patterns in the codebase:

**Validation function format**:
```go
func IsValidAmount(value string) error {
    // validation logic
    return fmt.Errorf("Amount must be positive")
}
```

**Naming conventions**:
- Functions: `IsValid{Field}` (e.g., `IsValidAmount`, `IsValidPrice`)
- Files: `validation.go`, `validation_test.go`
- Package: `validation`

**Error messages** (must match exactly):
- Amount: "Amount must be positive"
- Price: "Price must be positive"  
- Asset: "Asset ticker is required"
- Date: "Use YYYY-MM-DD"

**Testing patterns**:
- Test each validation function with valid and invalid inputs
- Test exact error message format
- Test edge cases (zero, negative, empty, whitespace)
- Use table-driven tests where appropriate

### 5. Testing Strategy

**New tests in `internal/validation/validation_test.go`**:

- `TestIsValidAmount_Pass`: Valid positive amount
- `TestIsValidAmount_RejectZero`: Amount = 0
- `TestIsValidAmount_RejectNegative`: Amount < 0
- `TestIsValidAmount_RejectEmpty`: Empty string
- `TestIsValidAmount_RejectInvalid`: Non-numeric string
- `TestIsValidAmount_ExactErrorMessage`: Verify exact error text

- `TestIsValidPrice_Pass`: Valid positive price
- `TestIsValidPrice_RejectZero`: Price = 0
- `TestIsValidPrice_RejectNegative`: Price < 0
- `TestIsValidPrice_RejectEmpty`: Empty string
- `TestIsValidPrice_RejectInvalid`: Non-numeric string
- `TestIsValidPrice_ExactErrorMessage`: Verify exact error text

- `TestIsValidAsset_Pass`: Valid ticker symbol
- `TestIsValidAsset_RejectEmpty`: Empty string
- `TestIsValidAsset_RejectWhitespace`: Whitespace-only string
- `TestIsValidAsset_ExactErrorMessage`: Verify exact error text

- `TestIsValidDate_Pass`: Valid RFC3339 date
- `TestIsValidDate_RejectInvalid`: Invalid format
- `TestIsValidDate_ExactErrorMessage`: Verify exact error text

- `TestRoundTo8Decimals`: Verify 8 decimal precision
- `TestCalculateSharesFromValues`: Verify share calculation

**Existing tests in `internal/form/validation_test.go`**:
- Update tests to use shared validation
- Verify TUI form still produces same results
- All 54+ existing tests must continue to pass

### 6. Risks and Considerations

| Risk | Mitigation |
|--|--|
| Breaking TUI form functionality | Update form to use shared functions; run all existing tests after changes |
| Error message mismatch | Copy error messages exactly from existing validation functions |
| Date validation complexity | Keep RFC3339 parsing as-is; only extract validation logic |
| Shared package versioning | Keep in same module; no external dependencies |
| Test duplication | Remove TUI-specific validation tests after migration |
| Precision drift | Reference existing `RoundTo8Decimals` in form package; document precision in shared package |
| CLI integration | GOT-057, GOT-058 will use this package; ensure exported functions are sufficient |
| Migration complexity | Incremental migration: create shared package first, then update consumers |
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
