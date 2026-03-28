---
id: GOT-059
title: 'Phase 2: CLI Implementation - cmd/dca/cli.go'
status: To Do
assignee: []
created_date: '2026-03-28 14:43'
labels:
  - cli
  - implementation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create CLI implementation file with flag parsing and entry logic

**Phase Context:**
- Objective: Implement CLI quick entry functionality without TUI
- Deliverables: cmd/dca/cli.go with flag handling, validation, persistence
- Stakeholders: Developers
- Dependencies: Phase 1 (validation functions must exist)
- Constraints: Silent success on valid input, exit code 1 on errors, 8 decimal share precision

**Task Requirements:**
1. Create cmd/dca/cli.go with:
   - Flag parsing using flag package
   - Required flags: --add, --asset, --amount, --price
   - Optional flags: --date (defaults to time.Now())
2. Implement runCLI() function that:
   - Validates inputs using shared validation functions
   - Calculates shares with 8 decimal precision
   - Creates DCAEntry struct
   - Persists to dca_entries.json using existing file I/O
3. Error handling:
   - Exit code 1 on validation errors
   - Error messages printed to stderr
   - No output on success

**Output Expected:**
- Working CLI entry point with proper flag parsing
- Integration with existing persistence layer
- Consistent error messages matching PRD specification

**Acceptance Criteria:**
- CLI accepts --add --asset BTC --amount 100 --price 50000
- Shares calculated as math.Round((100/50000)*1e8)/1e8 = 0.00200000
- Date auto-set to current RFC3339
- Missing flag returns exit code 1 with error message
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
