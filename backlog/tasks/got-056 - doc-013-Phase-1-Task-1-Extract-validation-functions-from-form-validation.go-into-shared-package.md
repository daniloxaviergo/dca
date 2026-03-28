---
id: GOT-056
title: >-
  [doc-013 Phase 1] Task 1: Extract validation functions from form/validation.go
  into shared package
status: To Do
assignee: []
created_date: '2026-03-28 17:37'
labels:
  - refactor
  - validation
  - code-quality
dependencies: []
documentation:
  - internal/form/validation.go
  - internal/form/validation_test.go
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract validation logic from internal/form/validation.go into internal/form/validation/shared_validation.go to enable reuse across TUI form and CLI. The extracted functions must validate --amount (positive), --price (positive), --asset (non-empty), and handle date parsing with RFC3339 format, returning descriptive error messages matching the current validation logic without modifying existing TUI behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Validation functions extracted without side effects
- [ ] #2 Shared functions preserve exact error messages from TUI
- [ ] #3 TUI validation tests still pass unchanged
- [ ] #4 New shared package file has clear documentation
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
