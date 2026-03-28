---
id: GOT-071
title: '[Integration Phase 1] Task 1: Integrate CLI mode detection into main.go'
status: To Do
assignee: []
created_date: '2026-03-28 17:00'
labels:
  - integration
  - cli
  - bugfix
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement CLI mode detection in main() to allow flag-based entry addition without initializing Bubble Tea TUI.

WHAT TO IMPLEMENT:

1. In cmd/dca/main.go, modify main() to:
   - Parse command flags early (before any Bubble Tea initialization)
   - Check for --add flag presence
   - If --add flag is set: call runCLI() and exit with its return code
   - If --add flag is NOT set: proceed with existing TUI initialization (bubbletea program)

2. Create cmd/dca/cli.go with runCLI() function that:
   - Defines flags: --add, --amount, --date, --asset, --price
   - Parses flags using flag.Parse()
   - Returns -1 if --add not set (TUI mode detected)
   - Validates all required flags and values
   - Auto-generates RFC3339 date if not provided
   - Calculates shares with 8 decimal precision using validation package
   - Creates DCAEntry and saves to defaultEntriesPath
   - Returns exit codes: 0 for success, 1 for errors
   - No output on success

3. Create internal/validation package with CalculateShares() function for 8-decimal precision share computation
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 main() parses flags before Bubble Tea initialization
- [ ] #2 --add flag detection works correctly
- [ ] #3 CLI mode calls runCLI() and exits without TUI
- [ ] #4 TUI mode continues unchanged when --add not present
- [ ] #5 runCLI() uses same defaultEntriesPath (dca_entries.json)
- [ ] #6 No output on successful CLI entry addition
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 main.go modified with flag parsing at start of main()
- [ ] #8 cli.go created with runCLI() and flag definitions
- [ ] #9 internal/validation/validation.go created with CalculateShares()
- [ ] #10 No breaking changes to existing TUI functionality
<!-- DOD:END -->
