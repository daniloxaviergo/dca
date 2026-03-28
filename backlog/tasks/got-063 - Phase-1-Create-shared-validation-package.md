---
id: GOT-063
title: '[Phase 1] Create shared validation package'
status: In Progress
assignee: []
created_date: '2026-03-28 16:49'
labels:
  - feature
  - refactoring
  - phase-1
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create internal/validation package with shared validation functions extracted from internal/form/validation.go. The package should export four validation functions: ValidateAmount, ValidateDate, ValidateAsset, and ValidatePrice. Each function should return descriptive error messages matching the current implementation.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 ValidateAmount validates positive numbers with error 'Amount must be positive'
- [ ] #2 ValidateDate validates RFC3339 format with error 'Use YYYY-MM-DD'
- [ ] #3 ValidateAsset validates non-empty with error 'Asset ticker is required'
- [ ] #4 ValidatePrice validates positive numbers with error 'Price must be positive'
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
