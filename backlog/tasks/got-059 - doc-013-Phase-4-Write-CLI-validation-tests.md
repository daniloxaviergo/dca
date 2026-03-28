---
id: GOT-059
title: '[doc-013 Phase 4] Write CLI validation tests'
status: To Do
assignee: []
created_date: '2026-03-28 17:46'
labels:
  - testing
  - validation
dependencies: []
references:
  - internal/form/validation_test.go
  - REQ-008
  - ACC-004
  - ACC-005
  - ACC-006
documentation:
  - doc-013
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/dca/cli_test.go with comprehensive unit tests for CLI validation logic. Tests must cover all validation rules: missing required flags (asset, amount, price), invalid values (negative/zero amounts and prices, empty asset ticker). Each test should verify the correct error message and exit behavior (exit code 1 on validation failure). Use table-driven tests for efficiency and coverage.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Test file cmd/dca/cli_test.go created
- [ ] #2 All validation rules covered (positive, zero, negative, empty cases)
- [ ] #3 Error messages match PRD specification
- [ ] #4 Tests pass with go test -v
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
