---
id: GOT-072
title: >-
  Phase 1 task: Verify data consistency - CLI entry accessible via TUI (PRD item
  7, REQ-007)
status: To Do
assignee: []
created_date: '2026-03-28 15:07'
labels:
  - data-consistency
  - integration
  - phase-1
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Verify that entries added via CLI are immediately visible and correctly aggregated in TUI's asset list view.

**Task Description:**
1. Create integration test: CLI add entry → TUI view → verify entry appears
2. Verify aggregation calculation (total shares, avg price, total value) matches
3. Test: Multiple CLI entries → TUI aggregation
4. Verify data file format compatibility (JSON structure unchanged)

**Acceptance Criteria:**
- [ ] CLI entry immediately visible in TUI after command completes
- [ ] Aggregation metrics correctly calculated (total shares, avg price, total value)
- [ ] Data file format unchanged (backward compatible)
- [ ] Test added to cmd/dca/cli_test.go and internal/assets/view_test.go

**Integration Test Flow:**
1. Clear test database
2. CLI add entry (BTC, $500, $65000, 0.00769231 shares)
3. TUI load entries
4. Verify entry appears in asset list with correct aggregation

**Data Consistency Checks:**
- Entry count matches CLI count
- Total shares calculated correctly (8 decimals)
- Average price formula matches (weighted average)
- File format unchanged (JSON structure identical)

**Assignee:** QA
**Priority:** High
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
