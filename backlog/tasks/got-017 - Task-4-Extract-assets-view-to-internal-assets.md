---
id: GOT-017
title: 'Task 4: Extract assets view to internal/assets/'
status: To Do
assignee: []
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 11:22'
labels: []
dependencies:
  - GOT-013
references:
  - backlog/docs/doc-004.md
priority: high
ordinal: 5000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move assets_view.go content to new package
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/assets/view.go created with AssetsView
- [ ] #2 internal/assets/aggregate.go created with aggregation functions
- [ ] #3 internal/assets/aggregate_test.go created with all tests
- [ ] #4 Package declaration changed to 'assets'
- [ ] #5 All tests pass
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
