---
id: GOT-008
title: 'Task 4: Add validation and error handling'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-16 21:25'
updated_date: '2026-03-17 00:12'
labels: []
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement comprehensive validation and error handling across all form inputs and file operations.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Reject negative or zero amounts with message: 'Amount must be positive'
- [x] #2 Reject invalid date format with helpful example: 'Use YYYY-MM-DD'
- [x] #3 Reject negative or zero prices with message: 'Price must be positive'
- [x] #4 Reject empty asset ticker with message: 'Asset ticker is required'
- [x] #5 Handle file permission errors with user-friendly message: 'Permission denied: check file permissions'
- [x] #6 Handle JSON parse errors gracefully with diagnostic message
- [x] #7 Handle file write errors with clear user message
- [x] #8 Validate that calculated shares is a valid finite number
- [x] #9 Catch and handle panics from numeric operations
- [x] #10 Display validation errors inline with the prompt for re-entry
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task implements comprehensive input validation and error handling across the DCA form application. The approach focuses on:

- **Form-level validation**: Enhanced validation in `dca_form.go` that rejects invalid inputs and displays errors inline with prompts for immediate re-entry
- **Data model validation**: Extended `Validate()` method in `dca_entry.go` to validate shares (must be finite and positive after calculation)
- **Error message standardization**: Ensure all error messages match the exact format specified in acceptance criteria
- **Graceful error handling**: Handle file permission errors, JSON parse errors, and file write errors with user-friendly messages

The implementation will:
1. Update form validation functions to return exact error messages per AC
2. Add shares validation in `Validate()` method to catch division-by-zero and infinite results
3. Enhance form error display to show inline errors with re-entry capability
4. Add panic recovery in critical numeric operations

### 2. Files to Modify

| File | Changes |
|------|---------|
| `dca_entry.go` | Update `Validate()` to check shares is finite and positive; update error messages to match AC requirements |
| `dca_form.go` | Update `validateAmount()`, `validateDate()`, `validateAsset()`, `validatePrice()` to return exact error messages; enhance error display in `renderForm()` for inline error presentation with re-entry |
| `main.go` | Add panic recovery around form execution; enhance error messages for file operations |

### 3. Dependencies

- **No blocking dependencies**: This task can proceed independently
- **Related tasks**: Task references `doc-002` (PRD) which defines the validation requirements
- **Prerequisites**: Current form structure must be in place (already implemented in GOT-006)

### 4. Code Patterns

Follow existing patterns in the codebase:
- **Error messages**: Use exact format from AC (e.g., "Amount must be positive", "Use YYYY-MM-DD")
- **Validation pattern**: Return `error` from validation functions; set `field.Error` in form
- **Error display**: Show error inline with field value using red foreground color (`lipgloss.Color("196")`)
- **Re-entry support**: Keep field editable when validation fails (don't advance step)
- **Division safety**: Use `math.IsNaN()`, `math.IsInf()` to validate calculated shares

### 5. Testing Strategy

Add unit tests for:
- `TestValidateAmount_RejectZeroAndNegative`: Verify exact error messages
- `TestValidateDate_RejectInvalidFormat`: Verify helpful example in error message
- `TestValidatePrice_RejectZeroAndNegative`: Verify exact error messages
- `TestValidateAsset_RejectEmpty`: Verify exact error message
- `TestValidateShares_IsFinite`: Verify shares validation for NaN/Inf
- `TestFormModel_RenderForm_InlineErrors`: Verify errors display inline with prompts

Verify all existing tests pass and error messages match acceptance criteria.

### 6. Risks and Considerations

| Risk | Mitigation |
|------|------------|
| Date format acceptance | Accept RFC3339 format but provide helpful "YYYY-MM-DD" example in error message |
| Division by zero in shares | Check `PricePerShare == 0` before division; return error or 0 |
| JSON error messages | Use `json.SyntaxError` for parse errors with position info |
| File permission errors | Catch `*os.PathError` and return user-friendly message |
| Panic in numeric operations | Wrap critical sections in `recover()` or use safe math functions |
| Backward compatibility | Error messages are additive; no breaking changes to data model |
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Starting implementation of validation and error handling. Reviewing codebase structure in dca_entry.go, dca_form.go, and main.go. Planning to update all validation functions to match exact error message formats from acceptance criteria.

Implementation complete. All validation functions updated to return exact error messages from acceptance criteria. Error handling for file operations and JSON parsing enhanced. Panic recovery added to main.go. All 54 tests passing. Build successful.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
### Implementation Summary

Implemented comprehensive validation and error handling across all form inputs and file operations:

**dca_entry.go:**
- Updated `Validate()` method to check shares is finite and positive
- Updated error messages to match exact format from acceptance criteria ("Amount must be positive", "Price must be positive", "Shares must be a positive finite number")
- Enhanced `LoadEntries()` to handle permission errors and JSON parse errors with user-friendly messages
- Enhanced `SaveEntries()` to handle permission errors and provide clear file operation errors

**dca_form.go:**
- Updated `validateAmount()` to return exact message "Amount must be positive"
- Updated `validateDate()` to return exact message "Use YYYY-MM-DD"
- Updated `validateAsset()` to return exact message "Asset ticker is required"
- Updated `validatePrice()` to return exact message "Price must be positive"
- Added `math` import for NaN/Inf validation
- Updated `CalculateSharesFromValues()` to validate shares is finite
- Enhanced error display (already inline with red foreground and ❌ indicator)

**main.go:**
- Added panic recovery wrapper around main() with stack trace output

**Tests:**
- Updated existing tests to verify new error message formats
- Added new tests for exact error messages in all validation functions
- Added tests for inline error display
- Added tests for file permission and JSON parse error handling
- All 54 tests passing

**Verification:**
- All tests pass: `go test ./... -v`
- Build succeeds: `go build -o dca .`
- No new warnings or regressions
<!-- SECTION:FINAL_SUMMARY:END -->
