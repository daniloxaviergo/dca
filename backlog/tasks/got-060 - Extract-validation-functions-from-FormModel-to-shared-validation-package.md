---
id: GOT-060
title: Extract validation functions from FormModel to shared validation package
status: In Progress
assignee: []
created_date: '2026-03-28 14:46'
labels:
  - refactoring
  - validation
  - cli
  - tui
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
 Extract validation functions from internal/form/validation.go to enable shared validation between CLI and TUI components.

## Task Details
 Currently, validation methods are unexported methods on FormModel:
 - `validateAmount(value string) error`
 - `validateDate(value string) error`
 - `validateAsset(value string) error`
 - `validatePrice(value string) error`

 These need to be extracted as exported package-level functions for CLI reuse.

## Requirements
 1. Extract all validation functions to package-level exported functions
 2. Maintain exact error messages from PRD:
    - "Amount must be positive"
    - "Price must be positive"
    - "Asset ticker is required"
    - "Use YYYY-MM-DD"
 3. Update model.go to use extracted functions
 4. Update validation_test.go to test new exported functions
 5. Add validation test coverage for shares field

## Acceptance Criteria
 - All validation functions are package-level exported
 - Exact error messages match PRD specification
 - All existing tests pass after refactoring
 - New validation functions can be called from CLI without FormModel instance

## References
 - PRD: doc-013 - Command-Line Quick Entry
 - REQ-009: Validation Consistency
 - Technical Decision: "Shared validation: Extract from form/validation.go"
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
