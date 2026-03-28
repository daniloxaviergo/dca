---
id: GOT-056
title: '[doc-013] Phase 1: Extract Validation Functions from form/validation.go'
status: To Do
assignee: []
created_date: '2026-03-28 15:18'
labels:
  - refactoring
  - validation
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract shared validation logic from internal/form/validation.go to enable CLI reuse without TUI dependencies.

## Phase Context
- **Objective**: Create reusable validation functions that can be called from both CLI and TUI without circular dependencies
- **Deliverables**: 
  - Validation functions in internal/form/ that accept bare config structs
  - Updated TUI form code that uses the extracted functions
  - No breaking changes to existing form validation behavior
- **Stakeholders**: Developers (code reuse), QA Team (consistent validation)
- **Dependencies**: None - this is a refactoring phase
- **Constraints**: Must maintain backward compatibility with existing form validation

## Task Generation Rules
- Use SMART criteria for each task
- Each task should trace to a specific validation rule from the PRD
- Include tests for all extracted functions
- Consider error handling patterns (errors.Is, errors.As)
- Verify no logic duplication between form and CLI validation

## Acceptance Criteria
1. All validation functions are extracted from form/validation.go into reusable functions
2. Form validation still works exactly as before
3. Validation functions can be called with simple parameter structs (not Bubble Tea components)
4. All existing tests pass
5. New validation test file covers all extracted functions
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
