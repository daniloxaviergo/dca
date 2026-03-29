---
id: GOT-063
title: '[doc-019 Phase 2] Update table rendering methods to new column widths'
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
Modify renderHeaderRow(), renderDataRow(), and renderEmptyDataRow() methods in internal/assets/view.go to use new column width constants (12, 8, 16, 14, 16). Update column separator usage from ColumnSeparator (2 spaces) to new 3-space separator. Modify renderTable() to use double-line rounded borders by replacing lipgloss.RoundedBorder() with lipgloss.NewStyle().Border(lipgloss.DoubleBorder()). Ensure all data formatting maintains decimal precision (8 decimals for shares, 2 decimals for prices/values).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 renderHeaderRow() updated with new column widths
- [ ] #2 renderDataRow() updated with new column widths
- [ ] #3 renderEmptyDataRow() updated with new column widths
- [ ] #4 renderTable() uses lipgloss.DoubleBorder() for double-line borders
- [ ] #5 Data formatting maintains 8 decimal precision for shares and 2 decimals for prices/values
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
