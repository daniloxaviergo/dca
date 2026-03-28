---
id: GOT-061
title: '[doc-013 Phase 1] Task 1: Extract validation functions to validation_utils.go'
status: In Progress
assignee: []
created_date: '2026-03-28 16:43'
labels:
  - refactor
  - feature
  - validation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract shared validation functions from internal/form/validation.go to enable CLI flag validation without TUI dependencies.

Implement extracted validation functions:
- ValidateAsset(ticker string) error
- ValidateAmount(amount float64) error
- ValidatePrice(price float64) error
- ValidateDate(dateStr string) (time.Time, error)

All functions must be exported (capitalized) and return descriptive error messages matching existing validation logic. The Date validation should parse and return the time.Time value for CLI use.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All validation functions are exported and reusable
- [ ] #2 Functions match existing validation logic exactly
- [ ] #3 Tests cover all validation rules from PRD
- [ ] #4 Exit code 1 on validation errors (CLI requirement)
- [ ] #5 No breaking changes to existing TUI form
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
