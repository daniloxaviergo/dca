---
id: GOT-056
title: '[doc-013 Phase 1] Extract validation functions into shared package'
status: To Do
assignee: []
created_date: '2026-03-28 20:49'
labels:
  - feature
  - refactoring
  - validation
dependencies: []
references:
  - 'doc-013 - Phase 1: Extract validation logic'
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract validation functions from internal/form/validation.go to enable reuse between TUI form and CLI entry. Create shared validation logic for amount (positive), price (positive), and asset (non-empty) validation with descriptive error messages. Ensure functions are_exported for use across packages and maintain the same validation rules and precision (8 decimals for shares calculation).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Validation functions extracted with unit tests
- [ ] #2 Shared validation logic tested with existing TUI form
- [ ] #3 Functions maintain 8 decimal precision for shares calculation
- [ ] #4 Error messages match PRD specification
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
