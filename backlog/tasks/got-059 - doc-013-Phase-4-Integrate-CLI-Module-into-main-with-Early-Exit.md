---
id: GOT-059
title: '[doc-013] Phase 4: Integrate CLI Module into main() with Early Exit'
status: To Do
assignee: []
created_date: '2026-03-28 15:18'
labels:
  - cli
  - integration
  - refactoring
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Integrate the CLI module into main() with flag detection and early exit behavior to prevent TUI initialization when CLI mode is active.

## Phase Context
- **Objective**: Add CLI flag detection to main() and implement early exit when --add flag is present
- **Deliverables**: 
  - Flag detection logic in main() before TUI initialization
  - Early exit with runCLI() when CLI mode detected
  - No TUI startup when using --add flag
  - Standard TUI flow for normal execution (no flags)
- **Stakeholders**: Developers (integration), QA Team (functionality)
- **Dependencies**: Phase 2 completed (cli.go available)
- **Constraints**: Must not break existing TUI functionality, maintain backward compatibility

## Task Generation Rules
- Include separate tasks for flag detection and TUI initialization control
- Trace behavior to specific acceptance criteria (ACC-008: TUI unchanged)
- Consider error handling for CLI initialization failures
- Include logging/debugging considerations
- Verify no race conditions in flag parsing

## Acceptance Criteria
1. main() detects --add flag and routes to CLI mode
2. TUI not initialized when CLI mode is active
3. Standard TUI flow preserved when no flags present
4. Early exit does not interfere with existing TUI behavior
5. All existing tests continue to pass
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
