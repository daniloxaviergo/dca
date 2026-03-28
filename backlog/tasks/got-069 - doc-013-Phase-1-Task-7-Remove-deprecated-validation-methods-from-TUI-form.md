---
id: GOT-069
title: '[doc-013 Phase 1] Task 7: Remove deprecated validation methods from TUI form'
status: To Do
assignee: []
created_date: '2026-03-28 16:53'
labels:
  - refactoring
  - cleanup
  - phase1
dependencies: []
priority: low
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Clean up the TUI form validation after migration is complete and stable.

Steps:
1. After successful migration and testing, deprecate and remove wrapper methods
2. Update any remaining direct validation calls to use shared functions
3. Ensure backward compatibility during transition period
4. Update documentation and comments
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Deprecated methods removed after transition
- [ ] #2 All validation uses shared functions
- [ ] #3 Code clean and maintainable
- [ ] #4 No functionality regressions
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
