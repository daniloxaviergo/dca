---
id: GOT-064
title: >-
  [doc-013-001] Phase 1: Extract validation functions from form/validation.go
  into reusable package
status: To Do
assignee: []
created_date: '2026-03-28 15:20'
labels:
  - refactor
  - validation
  - reusability
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract validation logic from internal/form/validation.go into a new internal/validation/validation.go package that can be used by both CLI and TUI forms. Create clean parameter structs (ValidationParams, ValidateAmountParams, etc.) that encapsulate validation requirements.

Key changes:
- Create internal/validation/validation.go with reusable functions
- Extract amount validation, date validation, asset validation, price validation, share calculation
- Create validation parameter structs for clean API
- Maintain identical error messages and validation behavior
- Ensure 8 decimal precision for share calculations

This task ensures validation is DRY and can be used by both the interactive form and CLI entry points.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Validation functions extracted to internal/validation/validation.go
- [ ] #2 All existing form validation tests pass without modification
- [ ] #3 Validation functions accept parameter structs (ValidateAmountParams, ValidateDateParams, etc.)
- [ ] #4 Error messages remain identical to original format
- [ ] #5 Share calculations use 8 decimal precision with math.Round()
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 Review internal/form/validation.go and identify all reusable validation functions
- [ ] #8 Create internal/validation folder and validation.go file
- [ ] #9 Extract each validation function with clear parameter structs
- [ ] #10 Update internal/form/validation.go to use new package (wrapper functions)
- [ ] #11 Run make test to verify all existing tests pass
- [ ] #12 Document new validation package functions with examples
<!-- DOD:END -->
