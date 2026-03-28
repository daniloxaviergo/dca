---
id: GOT-066
title: '[doc-013 Phase 1] Task 4: Update TUI form to use new shared validation'
status: To Do
assignee: []
created_date: '2026-03-28 16:53'
labels:
  - refactoring
  - validation
  - phase1
  - tui
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update internal/form/validation.go to use the new shared validation functions.

Changes required:
1. Import the new internal/validation package
2. Update existing validation methods to call shared functions
3. Keep existing wrapper methods for backward compatibility
4. Maintain identical error messages and validation behavior
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 TUI form uses shared validation functions
- [ ] #2 Existing tests pass without modification
- [ ] #3 Error messages remain identical
- [ ] #4 No breaking changes to form behavior
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
