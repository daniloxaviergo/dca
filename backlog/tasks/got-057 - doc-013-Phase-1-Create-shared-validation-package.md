---
id: GOT-057
title: '[doc-013 Phase 1] Create shared validation package'
status: In Progress
assignee: []
created_date: '2026-03-28 15:30'
labels:
  - refactor
  - validation
  - shared_package
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create internal/validate/ package with shared validation functions for DCA Investment Tracker.

Implement four validation functions:
1. ValidateAmount(value string) error - validates amount is positive
2. ValidateAsset(value string) error - validates asset ticker is non-empty
3. ValidatePrice(value string) error - validates price is positive
4. ValidateDate(value string) error - validates date is RFC3339 format

All functions should return the exact same error messages as the current FormModel validation methods.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/validate/ directory created
- [ ] #2 validate.go with all four validation functions
- [ ] #3 All error messages match current validation.go exactly
- [ ] #4 No external dependencies beyond standard library
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
