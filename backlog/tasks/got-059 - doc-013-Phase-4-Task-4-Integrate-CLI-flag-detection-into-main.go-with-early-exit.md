---
id: GOT-059
title: >-
  [doc-013 Phase 4] Task 4: Integrate CLI flag detection into main.go with early
  exit
status: To Do
assignee: []
created_date: '2026-03-28 17:38'
labels:
  - feature
  - cli
  - integration
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Modify cmd/dca/main.go to detect --add flag and early-exit to CLI mode. The main() function must call runCLI() when --add flag is present, bypassing TUI initialization. All other flag combinations should trigger TUI startup. The change must not affect existing TUI functionality or behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 --add flag detection in main() before TUI init
- [ ] #2 runCLI() called when --add present
- [ ] #3 TUI unchanged for non-CLI invocations
- [ ] #4 No breaking changes to existing behavior
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
