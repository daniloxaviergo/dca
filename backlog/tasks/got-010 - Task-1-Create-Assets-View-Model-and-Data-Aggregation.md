---
id: GOT-010
title: 'Task 1: Create Assets View Model and Data Aggregation'
status: To Do
assignee: []
created_date: '2026-03-17 00:42'
labels: []
dependencies: []
references:
  - 'PRD: DCA Assets List Table View'
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create assets_view.go with AssetSummary struct, AssetsViewModel, and data aggregation logic
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Entries loaded from dca_entries.json correctly
- [ ] #2 Grouping by asset ticker works
- [ ] #3 Sum of shares calculated per asset
- [ ] #4 Weighted average entry price calculated correctly
- [ ] #5 Empty file handled gracefully
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
