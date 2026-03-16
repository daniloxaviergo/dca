---
id: GOT-008
title: 'Task 4: Add validation and error handling'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-16 21:25'
updated_date: '2026-03-16 23:26'
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
- [ ] #1 Reject negative or zero amounts with message: 'Amount must be positive'
- [ ] #2 Reject invalid date format with helpful example: 'Use YYYY-MM-DD'
- [ ] #3 Reject negative or zero prices with message: 'Price must be positive'
- [ ] #4 Reject empty asset ticker with message: 'Asset ticker is required'
- [ ] #5 Handle file permission errors with user-friendly message: 'Permission denied: check file permissions'
- [ ] #6 Handle JSON parse errors gracefully with diagnostic message
- [ ] #7 Handle file write errors with clear user message
- [ ] #8 Validate that calculated shares is a valid finite number
- [ ] #9 Catch and handle panics from numeric operations
- [ ] #10 Display validation errors inline with the prompt for re-entry
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
