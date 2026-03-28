---
id: GOT-058
title: >-
  [doc-013 Phase 3] Task 3: Implement runCLI() function with validation and
  persistence
status: To Do
assignee: []
created_date: '2026-03-28 17:38'
labels:
  - feature
  - cli
  - validation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement runCLI() function in cmd/dca/cli.go that performs full CLI workflow: validate flags, auto-generate date, calculate shares with 8 decimal precision, and persist to dca_entries.json using internal/dca/file.go functions. The function must return exit codes (0 for success, 1 for errors) and produce no output on success per acceptance criteria.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 runCLI() validates all required flags
- [ ] #2 Date auto-generated using time.Now() in RFC3339
- [ ] #3 Shares calculated with math.Round((amount/price)*1e8)/1e8
- [ ] #4 Persistence uses atomic write pattern from internal/dca/file.go
- [ ] #5 No output on success
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
