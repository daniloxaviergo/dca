---
id: GOT-056
title: '[doc-013 Phase 1] Extract validation logic into reusable package'
status: To Do
assignee: []
created_date: '2026-03-28 17:45'
labels:
  - refactor
  - validation
dependencies: []
references:
  - internal/form/validation.go
  - REQ-009
  - NFA-002
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract validation functions from internal/form/validation.go into a reusable validation package. This includes func validateAmount(amount string) error, validatePrice(price string) error, and validateAsset(asset string) error with their corresponding validation rules (positive numbers, non-empty string). The extracted functions must maintain the exact same validation behavior and error messages as the original implementation for consistency across TUI and CLI.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Validation functions are extracted into new internal/form/validation package
- [ ] #2 Error messages remain identical to original implementation
- [ ] #3 All validation edge cases covered (negative, zero, empty)
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
