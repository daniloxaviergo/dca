---
id: GOT-015
title: 'Task 2: Extract core data model to internal/dca/'
status: To Do
assignee: []
created_date: '2026-03-17 11:20'
labels: []
dependencies:
  - GOT-013
references:
  - backlog/docs/doc-004.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move dca_entry.go content to new package
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/dca/entry.go created with DCAEntry, DCAData, LoadEntries, SaveEntries
- [ ] #2 internal/dca/entry_test.go created with all tests
- [ ] #3 Package declaration changed to 'dca'
- [ ] #4 All tests pass
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
