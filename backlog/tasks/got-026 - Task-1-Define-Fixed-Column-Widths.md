---
id: GOT-026
title: 'Task 1: Define Fixed Column Widths'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 20:28'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define explicit column widths for the Assets View table to ensure 100% width coverage and consistent column sizing across all terminals.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Asset column: minimum 10 characters width
- [ ] #2 Count column: 8 characters width
- [ ] #3 Total Shares column: 12 characters width
- [ ] #4 Avg Price column: 12 characters width
- [ ] #5 Total Value column: 14 characters width
- [ ] #6 Column separator: 2 spaces between columns
- [ ] #7 Total table width: 100% of terminal width
- [ ] #8 Unit tests pass for column width definitions
- [ ] #9 Table renders without panics
- [ ] #10 go fmt applied
- [ ] #11 go build succeeds
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
