---
id: GOT-056
title: 'DCA-CL-01: Extract validation logic from form/validation.go'
status: To Do
assignee: []
created_date: '2026-03-28 14:41'
labels:
  - feature
  - refactoring
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create a new validation package to extract shared validation functions (ValidateAmount, ValidatePrice, ValidateAsset) from internal/form/validation.go to enable code reuse for CLI functionality.

**Context**:
- File: cmd/dca/main.go contains the TUI Bubble Tea application
- File: internal/form/validation.go contains validation utilities
- Need to extract validation functions to shared location for CLI reuse
- PRD location: backlog/docs/doc-013 - Command-Line-Quick-Entry.md

**Task Instructions**:
1. Read internal/form/validation.go to understand current validation logic
2. Create new directory internal/validation/ with validation.go
3. Extract functions:
   - ValidateAmount(amount float64) error
   - ValidatePrice(price float64) error  
   - ValidateAsset(asset string) error
4. Each function should return descriptive errors as specified in PRD:
   - Amount: "Amount must be positive"
   - Price: "Price must be positive"  
   - Asset: "Asset ticker is required"
5. Update internal/form/validation.go to import and use shared functions
6. Add unit tests in internal/validation/validation_test.go
7. Run tests with: make test -v ./internal/validation/

**Acceptance Criteria**:
- All extracted validation functions pass unit tests
- Existing form tests continue to pass (backward compatibility)
- Code coverage ≥90% for validation logic
- No breaking changes to existing functionality

**Dependencies**: None (Phase 1 - must be completed first)
**Priority**: HIGH
**Status**: Todo
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
