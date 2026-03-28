---
id: GOT-082
title: >-
  [doc-013-cli-03] Update internal/form/validation.go to import and use
  internal/validation package
status: To Do
assignee: []
created_date: '2026-03-28 15:15'
labels:
  - phase-3
  - refactor
  - form
  - refactoring
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Refactor internal/form/validation.go to reuse shared validation package, ensuring TUI and CLI modes have identical validation logic.

Tasks:
1. Add import for internal/validation package
2. Replace inline validation with calls to validation package functions:
   - validateAmount() → validation.ValidateAmount()
   - validatePrice() → validation.ValidatePrice()
   - validateAsset() → validation.ValidateAsset()
   - validateDate() → validation.ValidateDate()
3. Update error return statements to match validation package
4. Run all existing tests to ensure no behavior changes
5. Document the refactoring in code comments

Test file: internal/form/validation_test.go (verify existing tests still pass)
Reference: PRD R3 - Validation reuse through shared package; Ensure no TUI behavior changes
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/form/validation.go imports internal/validation
- [ ] #2 All validation calls use shared package functions
- [ ] #3 All existing TUI tests pass
- [ ] #4 No behavior changes to TUI validation
- [ ] #5 Error messages identical to before refactor
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
