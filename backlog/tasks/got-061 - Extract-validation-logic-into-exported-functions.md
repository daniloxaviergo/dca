---
id: GOT-061
title: Extract validation logic into exported functions
status: In Progress
assignee: []
created_date: '2026-03-28 14:52'
labels:
  - refactor
  - validation
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract all validation functions from internal/form/validation.go into reusable exported functions that can be called from both TUI form and CLI entry.

Current validation functions to extract:
- validateAmount → ValidateAmount
- validateDate → ValidateDate  
- validateAsset → ValidateAsset
- validatePrice → ValidatePrice

Requirements:
- All functions should be exported (capitalized)
- Functions should accept string inputs and return error
- Preserve exact error messages from existing implementation
- Maintain backward compatibility with existing TUI form (add wrapper methods if needed)
- Update validation_test.go to test the exported functions
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All validation functions are properly exported (capitalized)
- [ ] #2 Functions accept string inputs and return error
- [ ] #3 Error messages match exactly those in existing validation.go
- [ ] #4 TUI form continues to work without modification
- [ ] #5 New tests added for exported validation functions
- [ ] #6 Tests pass with make test
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
