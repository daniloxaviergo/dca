---
id: GOT-060
title: '[doc-013] Phase 5: Add CLI Tests - Validation and Edge Cases'
status: To Do
assignee: []
created_date: '2026-03-28 15:19'
labels:
  - cli
  - testing
  - validation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add comprehensive tests for CLI validation logic and edge case handling.

## Phase Context
- **Objective**: Ensure all validation rules and error cases are properly tested
- **Deliverables**: 
  - Test suite for validation functions
  - Tests for each validation rule (positive, negative, zero, empty)
  - Edge cases (very large numbers, precision limits)
  - Test file: cmd/dca/cli_test.go
- **Stakeholders**: QA Team (coverage), Developers (reliability)
- **Dependencies**: Phase 1 completed (validation functions available), Phase 2 completed (CLI module available)
- **Constraints**: Must match exact error messages from TUI validation

## Task Generation Rules
- Each validation rule should have its own test case group
- Include positive cases (valid inputs) and negative cases (invalid inputs)
- Test error codes (exit code 1 for failures)
- Consider floating-point precision edge cases
- Use table-driven tests for validation rules

## Acceptance Criteria
1. All validation rules tested (amount > 0, price > 0, asset non-empty, date format)
2. Error messages match exact format from validation functions
3. Exit code 1 for all validation failures
4. Edge cases covered (zero, negative, very large numbers)
5. No test duplication with existing form tests
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
