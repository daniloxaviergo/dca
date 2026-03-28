---
id: GOT-058
title: '[doc-013 Phase 3] Implement CLI run function with validation and persistence'
status: To Do
assignee: []
created_date: '2026-03-28 20:49'
labels:
  - feature
  - cli
  - persistence
dependencies: []
references:
  - 'doc-013 - Phase 3: Implement CLI functionality'
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement runCLI() function in cmd/dca/cli.go that performs validation using shared functions, calculates shares with 8 decimal precision, auto-sets date to time.Now() if not provided, creates DCAEntry struct, and persists to dca_entries.json using internal/dca/file.go. The function must handle errors appropriately with exit code 1 and produce no output on success.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 runCLI() function implemented with full validation
- [ ] #2 Shares calculated with 8 decimal precision
- [ ] #3 Date auto-set to current RFC3339 on missing flag
- [ ] #4 Entry persisted using existing file I/O functions
- [ ] #5 No output on success, exit code 1 on error
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
