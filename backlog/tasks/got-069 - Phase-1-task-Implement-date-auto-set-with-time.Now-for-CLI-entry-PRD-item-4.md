---
id: GOT-069
title: >-
  Phase 1 task: Implement date auto-set with time.Now() for CLI entry (PRD item
  4)
status: To Do
assignee: []
created_date: '2026-03-28 15:06'
labels:
  - date
  - phase-1
  - auto-generation
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement automatic date setting to current time using time.Now() when CLI --add flag is invoked.

**Task Description:**
1. Modify cmd/dca/cli.go to use time.Now() for date when not specified
2. Format date to RFC3339 (YYYY-MM-DDTHH:MM:SSZ)
3. Ensure date matches format used by TUI form
4. Add test for date auto-generation accuracy

**Acceptance Criteria:**
- [ ] Date auto-set to current time when CLI invoked with --add
- [ ] Date formatted to RFC3339 (YYYY-MM-DDTHH:MM:SSZ)
- [ ] Date matches format used by TUI form (internal/form/model.go)
- [ ] Test added to cmd/dca/cli_test.go verifying date generation

**Test Cases:**
- Date should be approximately now() (within 1 second)
- Date format must match RFC3339 parsing in database layer

**Assignee:** Developer
**Priority:** Medium
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
