---
id: GOT-065
title: '[doc-019 Phase 4] Add comprehensive layout tests for new table structure'
status: To Do
assignee:
  - catarina
created_date: '2026-03-29 12:32'
updated_date: '2026-03-31 13:37'
labels:
  - task
  - testing
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: high
ordinal: 9000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add layout test functions in internal/assets/view_test.go to verify new table specifications. Add TestTableLayout_IncreasedWidth to verify 82-character total width (78 data + 4 border). Add TestTableLayout_BorderStyle to verify double-line rounded borders appear in output. Add TestTableLayout_HeaderAlignment to verify all columns align between headers and data. Add TestTableLayout_Exactly30Rows to verify row count unchanged at 30 (1 header + 29 data/empty). Add TestTableLayout_RenderPerformance to verify rendering completes within 50ms. Update existing tests to expect new column widths and separator format.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 TestTableLayout_IncreasedWidth added to verify 82-character width
- [ ] #2 TestTableLayout_BorderStyle added to verify double-line borders
- [ ] #3 TestTableLayout_HeaderAlignment added to verify column alignment
- [ ] #4 TestTableLayout_Exactly30Rows added to verify row count
- [ ] #5 TestTableLayout_RenderPerformance added to verify <50ms rendering
- [ ] #6 Existing tests updated to expect new column widths
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
