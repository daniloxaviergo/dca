---
id: GOT-071
title: >-
  Phase 1 task: Verify reuse of existing validation functions matches TUI
  behavior (PRD item 6)
status: To Do
assignee: []
created_date: '2026-03-28 15:06'
labels:
  - validation
  - compatibility
  - phase-1
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Verify that extracted validation functions produce identical results when used by CLI vs TUI form to prevent data inconsistency.

**Task Description:**
1. Compare validation logic in internal/form/validation.go with new internal/validation package
2. Ensure identical error messages for same validation failures
3. Test validation behavior for both CLI and TUI entry paths
4. Add integration test to verify validation produces same results

**Acceptance Criteria:**
- [ ] All validation functions in internal/validation have identical behavior to internal/form
- [ ] Same error messages returned for same invalid inputs
- [ ] Test added to internal/validation/validation_test.go comparing CLI vs TUI behavior
- [ ] Test added to cmd/dca/cli_test.go verifying TUI compatibility

**Validation Comparison Matrix:**
| Input | Expected Error (TUI) | Expected Error (CLI) | Match |
|------ |---- ------ |---- ------- |----- |
| amount="" | "Amount must be positive" | "Amount must be positive" | ✓ |
| amount="abc" | "Amount must be positive" | "Amount must be positive" | ✓ |
| amount="-5" | "Amount must be positive" | "Amount must be positive" | ✓ |
| asset="" | "Asset ticker is required" | "Asset ticker is required" | ✓ |

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
