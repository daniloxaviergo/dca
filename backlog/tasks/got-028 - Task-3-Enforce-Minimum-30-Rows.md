---
id: GOT-028
title: 'Task 3: Enforce Minimum 30 Rows'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 22:08'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add row padding to maintain exactly 30 rows in the Assets View table regardless of data volume.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 When data rows < 30: pad with empty rows to reach 30
- [ ] #2 When data rows = 30: display all rows without truncation
- [ ] #3 Empty rows use same styling as data rows but with empty values
- [ ] #4 Test verifies exactly 30 rows rendered with 5 assets
- [ ] #5 Test verifies exactly 30 rows rendered with 25 assets
- [ ] #6 go fmt applied
- [ ] #7 go build succeeds
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
