---
id: GOT-058
title: 'doc-009 Phase 1: Extract validation functions into internal/validation package'
status: In Progress
assignee: []
created_date: '2026-03-28 15:37'
labels:
  - refactor
  - validation
  - cli
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create internal/validation/validation.go with exported functions: ValidateAmount, ValidateDate, ValidateAsset, ValidatePrice. All functions must return exact error messages as specified: "Amount must be positive", "Price must be positive", "Asset ticker is required", "Use YYYY-MM-DD". Shares calculation helper function CalculateSharesFromValues with 8 decimal precision using math.Round.
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
