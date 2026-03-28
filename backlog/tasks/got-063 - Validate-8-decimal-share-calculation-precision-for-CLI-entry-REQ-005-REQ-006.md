---
id: GOT-063
title: >-
  Validate 8-decimal share calculation precision for CLI entry (REQ-005,
  REQ-006)
status: To Do
assignee: []
created_date: '2026-03-28 15:04'
labels:
  - validation
  - precision
  - technical-validation
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Validate that the share calculation uses the same 8-decimal precision algorithm (RoundTo8Decimals) in the CLI entry path as used in the form and data model. Ensure the algorithm rounds correctly: math.Round(shares*1e8)/1e8.

**Task Description:**
1. Verify current share calculation in internal/dca/entry.go uses correct rounding
2. Test 8-decimal precision with known edge cases (e.g., 500/65000=0.00769231)
3. Validate math.Round(float64(int(val*1e8+.5))/1e8) maintains precision without floating point drift
4. Document precision behavior in code comments
5. Add unit tests for boundary cases (very small/large values, zero division)

**Acceptance Criteria:**
- Share calculation returns exactly 8 decimal places for all valid inputs
- Rounding matches standard mathematical rounding (not truncation)
- Edge cases return appropriate values (NaN=0, Inf=0, division by zero=0)
- Test coverage for 5+ boundary cases in internal/form/validation_test.go

**Test Cases:**
- 500.0 / 65000.0 = 0.00769231 (46 decimals → 0.00769231)
- 1000.0 / 333.33 = 3.00003000 (9 decimals → 3.00003)
- 1.0 / 3.0 = 0.33333333 (repeating → 0.33333333)
- 0.01 / 1000000 = 0.00000001 (8 zeros → 0.00000001)

**Failure Modes & Recovery:**
- If precision drift detected: Use float64 rounding instead of int-based approach
- If test failures occur: Review Go float64 representation and add debug logging

**Assignee:** Developer
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
