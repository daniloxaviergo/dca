---
id: GOT-064
title: '[doc-019 Phase 3] Update active row highlighting and header styling'
status: To Do
assignee:
  - catarina
created_date: '2026-03-29 12:32'
updated_date: '2026-03-31 13:28'
labels:
  - task
  - ui
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: medium
ordinal: 8000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update active row styling in internal/assets/view.go to use bright cyan background (#63) as specified in validation rules. Update renderHeaderRow() to apply bold + underline styling to headers using lipgloss.Bold(true) and Underline(true) methods. Verify all lipgloss styles apply correctly including foreground white (#15) for headers, bright cyan (#63) for active rows, and consistent padding (0 horizontal, 1 vertical). Ensure header styling matches FC-06 acceptance criteria for column alignment.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Active row uses bright cyan background (#63)
- [ ] #2 Header rows use bold + underline styling
- [ ] #3 Row padding set to 0 horizontal, 1 vertical
- [ ] #4 All lipgloss styles applied consistently
- [ ] #5 FC-06验收 criteria verified
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
