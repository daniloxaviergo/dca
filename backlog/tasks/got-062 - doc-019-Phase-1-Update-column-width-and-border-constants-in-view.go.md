---
id: GOT-062
title: '[doc-019 Phase 1] Update column width and border constants in view.go'
status: To Do
assignee: []
created_date: '2026-03-29 12:31'
labels:
  - task
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update column width constants from current values (10, 8, 14, 13, 13) to new values (12, 8, 16, 14, 16) in internal/assets/view.go. Update ColumnSeparator from "  " to "   " (3 spaces). Add new border style constant using lipgloss.DoubleBorder(). Update modal column widths to accommodate longer tickers and decimal places. Update row padding to 0 horizontal, 1 vertical as specified in validation rules. Ensure all width calculations sum to 82 characters (78 data + 4 border).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Column constants updated to new widths (12, 8, 16, 14, 16)
- [ ] #2 Separator updated to 3 spaces
- [ ] #3 Border style constant defined using lipgloss.DoubleBorder()
- [ ] #4 Row padding set to 0 horizontal, 1 vertical
- [ ] #5 Total width calculation verified as 82 characters
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
