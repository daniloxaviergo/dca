---
id: GOT-024
title: 'Task 3: Return to asset list after entry'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 18:19'
labels: []
dependencies:
  - GOT-023
references:
  - cmd/dca/main.go
  - internal/form/model.go
documentation:
  - doc-005
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
After form submission, return to asset list view with updated data. Modify cmd/dca/main.go Update() to handle formSubmittedMsg by reloading asset data and switching to StateAssetsView.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 After form submission, app switches to asset list view
- [ ] #2 Asset data refreshes to include new entry
- [ ] #3 Aggregation calculations update correctly
- [ ] #4 User can navigate asset list or create another entry
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
