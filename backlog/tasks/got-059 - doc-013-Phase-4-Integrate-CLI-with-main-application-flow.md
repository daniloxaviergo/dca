---
id: GOT-059
title: '[doc-013 Phase 4] Integrate CLI with main application flow'
status: To Do
assignee: []
created_date: '2026-03-28 20:50'
labels:
  - feature
  - integration
  - cli
dependencies: []
references:
  - 'doc-013 - Phase 4: Integrate with main'
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Modify cmd/dca/main.go to detect the --add flag during initialization and route to CLI mode with early exit before TUI initialization. Ensure the CLI path is executed before Bubble Tea program setup. Maintain backward compatibility by preserving all existing TUI functionality and state transitions unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Main function detects --add flag
- [ ] #2 CLI path executes early exit
- [ ] #3 TUI initialization skipped for CLI mode
- [ ] #4 All existing TUI functionality preserved
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
