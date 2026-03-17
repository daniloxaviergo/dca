---
id: GOT-022
title: 'Task 1: App starts in asset list view'
status: To Do
assignee: []
created_date: '2026-03-17 17:38'
labels: []
dependencies: []
references:
  - cmd/dca/main.go
documentation:
  - doc-005
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Change app to initialize in asset list view instead of form view. Modify cmd/dca/main.go to set currentState to StateAssetsView and initialize assetsView with loaded data on startup.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 App initializes with StateAssetsView instead of StateForm
- [ ] #2 Asset list displays on first render
- [ ] #3 User can navigate with ↑/↓ arrow keys
- [ ] #4 Esc/Ctrl+C exits the app from asset list
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
