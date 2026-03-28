---
id: GOT-068
title: 'Phase 1 task: Extract validation logic into reusable package (PRD item 1, 6)'
status: To Do
assignee: []
created_date: '2026-03-28 15:06'
labels:
  - validation
  - refactoring
  - phase-1
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extract shared validation functions from internal/form/validation.go into a new validation package for CLI reuse. Create cmd/dca/cli.go file.

**Task Description:**
1. Create new package: internal/validation/
2. Move validation functions from internal/form/validation.go:
   - validateAmount(value string) error
   - validateDate(value string) error (RFC3339)
   - validateAsset(value string) error
   - validatePrice(value string) error
3. Modify internal/form/validation.go to import new package
4. Create cmd/dca/cli.go with:
   - Flag parsing (--add, --asset, --amount, --price)
   - Silent success output (no verbose logging)
   - Share calculation using shared validation package
5. Extract calculation logic from internal/form/model.go:
   - CalculateSharesFromValues(amount, price float64) float64
   - RoundTo8Decimals(val float64) float64

**Acceptance Criteria:**
- [ ] Validation package created with all field validation functions
- [ ] internal/form/validation.go imports and uses new package
- [ ] cli.go created with flag parsing logic
- [ ] Share calculation logic extracted to shared package
- [ ] All validators return consistent error messages

**Dependency Mapping:**
- Phase 2 depends on this task (flag parsing in cli.go)
- Phase 3 depends on this task (CLI usage tests)

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
