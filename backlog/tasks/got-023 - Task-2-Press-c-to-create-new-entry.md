---
id: GOT-023
title: 'Task 2: Press ''c'' to create new entry'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 17:50'
labels: []
dependencies:
  - GOT-022
references:
  - internal/assets/view.go
  - cmd/dca/main.go
documentation:
  - doc-005
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add 'c' key handler in asset list to switch to form view. Modify internal/assets/view.go Update() to handle 'c' key press and return ViewTransitionMsg to trigger state change to StateForm.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Pressing 'c' in asset list switches to form view
- [ ] #2 Form fields are reset (amount=empty, asset=empty, price=empty)
- [ ] #3 Date defaults to current timestamp in RFC3339 format
- [ ] #4 User can navigate form with Tab/Enter as before
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
