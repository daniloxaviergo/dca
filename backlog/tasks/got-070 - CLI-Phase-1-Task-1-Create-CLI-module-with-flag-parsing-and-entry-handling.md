---
id: GOT-070
title: '[CLI Phase 1] Task 1: Create CLI module with flag parsing and entry handling'
status: To Do
assignee: []
created_date: '2026-03-28 16:56'
labels:
  - feature
  - cli
dependencies: []
documentation:
  - /home/danilo/scripts/github/dca/internal/form/validation.go
  - /home/danilo/scripts/github/dca/internal/dca/entry.go
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/dca/cli.go with complete CLI implementation:

1. Flag parsing using Go's flag package:
   - --add (boolean, triggers CLI mode)
   - --asset (string, required, ticker symbol)
   - --amount (string, required, USD investment)
   - --price (string, required, price per share)
   - --date (string, optional, defaults to today in RFC3339)

2. Validation using internal/form package functions:
   - validateAsset() - non-empty ticker
   - validateAmount() - positive number
   - validatePrice() - positive number
   - validateDate() - RFC3339 format (if provided)

3. Entry creation with auto-calculated shares:
   - Use math.Round((amount / price) * 1e8) / 1e8 for 8 decimal precision
   - Use time.Now() or parsed --date for entry date
   - Create dca.DCAEntry with all fields populated

4. Data persistence using existing functions:
   - Load data with dca.LoadEntries(defaultEntriesPath)
   - Add entry to asset's entry list
   - Save with dca.SaveEntries(defaultEntriesPath)

5. Error handling:
   - Use log.Fatal() for all error conditions (exit code 1)
   - Use existing validation error message formats
   - Silent success (no output on successful entry)

6. Code style:
   - Format with go fmt
   - Use existing constant defaultEntriesPath = "dca_entries.json"
   - No new dependencies (standard library only)
   - Follow project error message patterns from validation.go
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 CLI runs with ./dca --add --asset BTC --amount 500 --price 65000
- [ ] #2 Shares calculated with 8 decimal precision
- [ ] #3 Date auto-set to today if not provided
- [ ] #4 Silent success (no output) on valid entry
- [ ] #5 Exit code 1 on any error
- [ ] #6 Data persisted to dca_entries.json
- [ ] #7 Existing TUI functionality unchanged
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
