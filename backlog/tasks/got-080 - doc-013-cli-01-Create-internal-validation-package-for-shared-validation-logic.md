---
id: GOT-080
title: >-
  [doc-013-cli-01] Create internal/validation package for shared validation
  logic
status: To Do
assignee: []
created_date: '2026-03-28 15:14'
labels:
  - phase-3
  - validation
  - refactor
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create internal/validation package with shared validation functions extracted from internal/form/validation.go to ensure consistency between TUI and CLI modes.

Tasks:
1. Create directory: internal/validation/
2. Extract and refactor validation functions with exact error messages from PRD:
   - ValidateAmount(value float64) error (err: "Amount must be positive")
   - ValidatePrice(value float64) error (err: "Price must be positive")
   - ValidateAsset(value string) error (err: "Asset ticker is required")
   - ValidateDate(dateStr string) error (err: "Date must be in YYYY-MM-DD format")
3. Add comprehensive unit tests for each validation function
4. Test edge cases: empty string, zero/negative values, invalid formats
5. Verify error messages match PRD requirements exactly
6. Achieve 100% test coverage for validation package

Test file: internal/validation/validation_test.go
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/validation package created with all validation functions
- [ ] #2 All error messages match PRD requirements exactly
- [ ] #3 100% test coverage for validation package
- [ ] #4 No function signature changes to existing TUI code
- [ ] #5 Validation behavior identical to TUI form
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
