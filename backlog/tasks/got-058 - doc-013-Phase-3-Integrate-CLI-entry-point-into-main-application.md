---
id: GOT-058
title: '[doc-013 Phase 3] Integrate CLI entry point into main application'
status: To Do
assignee: []
created_date: '2026-03-28 17:46'
labels:
  - integration
  - flags
dependencies: []
references:
  - cmd/dca/main.go
  - REQ-001
  - REQ-007
  - REQ-008
documentation:
  - doc-013
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Modify cmd/dca/main.go to add CLI flag detection in the main() function. When --add flag is present, call runCLI() and exit immediately before TUI initialization. All other flag combinations and absent flags should proceed to existing TUI flow. Ensure error handling with exit code 1 on validation failures and silent exit with code 0 on success.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 main() checks for --add flag early in execution
- [ ] #2 runCLI() invoked only when --add flag is present
- [ ] #3 TUI initialization skipped for CLI mode
- [ ] #4 Proper exit codes (0 for success, 1 for errors)
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
