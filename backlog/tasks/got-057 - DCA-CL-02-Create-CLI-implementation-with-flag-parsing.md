---
id: GOT-057
title: 'DCA-CL-02: Create CLI implementation with flag parsing'
status: To Do
assignee: []
created_date: '2026-03-28 14:41'
labels:
  - feature
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/dca/cli.go with CLI-specific logic for processing --add flag and adding investment entries without launching the TUI.

**Context**:
- File: cmd/dca/main.go contains the main TUI application
- Package: internal/form/validation contains extracted validation functions
- File: internal/dca/file.go contains file I/O operations
- PRD location: backlog/docs/doc-013 - Command-Line-Quick-Entry.md

**Task Instructions**:
1. Create cmd/dca/cli.go with proper structure
2. Implement flag parsing for:
   - --add (boolean flag)
   - --asset (string, required)
   - --amount (float64, required)
   - --price (float64, required)
   - --date (string, optional, defaults to time.Now())
3. Implement runCLI() function that:
   - Validates allinput fields using extracted validation functions
   - Calculates shares: math.Round((amount / price) * 1e8) / 1e8
   - Creates DCAEntry struct
   - Persists to file using existing file I/O functions
4. In main.go, detect --add flag and call runCLI() early (before TUI init)
5. Implement error handling with exit codes (1 on error, 0 on success)
6. Silent success: no output on successful entry

**Validation Requirements**:
- --amount: Must be positive (> 0)
- --price: Must be positive (> 0)
- --asset: Must be non-empty
- Shares: Round to 8 decimal places

**Acceptance Criteria**:
- CLI successfully reads flags and validates input
- CLI adds entry to data file (dca_entries.json)
- Correct shares calculated with 8 decimal precision
- Date auto-set to current time
- No output on success
- Exit code 1 on validation or I/O error
- No breaking changes to existing TUI functionality
- Test with: ./dca --add --asset BTC --amount 500 --price 65000

**Dependencies**: DCA-CL-01 (validation extraction must be completed first)
**Priority**: HIGH
**Status**: Todo
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
