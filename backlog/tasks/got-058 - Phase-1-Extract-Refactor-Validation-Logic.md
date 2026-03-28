---
id: GOT-058
title: 'Phase 1: Extract & Refactor Validation Logic'
status: To Do
assignee: []
created_date: '2026-03-28 14:43'
labels:
  - refactoring
  - validation
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract shared validation functions from internal/form/validation.go for CLI reuse

**Phase Context:**
- Objective: Create reusable validation functions that can be used by both TUI form and CLI
- Deliverables: Extracted validation functions in a shared location, unit tests for each validation
- Stakeholders: Developers, QA Team
- Dependencies: None (foundation phase)
- Constraints: Maintain existing TUI functionality, use same validation logic

**Task Requirements:**
1. Create internal/form/validation.go with extracted functions
2. Extract the following validation functions:
   - ValidateAmount(amount string) error
   - ValidatePrice(price string) error
   - ValidateAsset(asset string) error
   - ValidateDate(date string) error
   - CalculateShares(amount, price float64) float64
3. Keep validation_test.go with unit tests for each function
4. Test for edge cases: empty strings, negative numbers, zero values, invalid date formats

**Output Expected:**
- Extracted validation functions with clear error messages matching PRD requirements
- Test coverage of 100% for validation functions
- No breaking changes to existing form package functionality
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
