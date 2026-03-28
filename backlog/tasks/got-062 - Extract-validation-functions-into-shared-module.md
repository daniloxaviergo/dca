---
id: GOT-062
title: Extract validation functions into shared module
status: In Progress
assignee: []
created_date: '2026-03-28 14:58'
labels:
  - phase-1
  - refactor
  - validation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
### Action Items

1. **Create shared validation module** at `internal/validation/`
   - Create `validation.go` with extracted validation functions
   - Create `validation_test.go` with comprehensive unit tests
   - Functions to extract:
     - `ValidateAmount(amount string) (float64, error)`
     - `ValidatePrice(price string) (float64, error)`
     - `ValidateAsset(asset string) (string, error)`
     - `ValidateDate(date string) (time.Time, error)`
     - `RoundTo8Decimals(val float64) float64`
     - `CalculateShares(amount, price float64) float64`

2. **Refactor TUI form** (`internal/form/validation.go`)
   - Import the shared validation package
   - Update methods to call shared functions
   - Keep existing method signatures for backward compatibility

3. **Run tests** to ensure everything works:

```bash
make test
```

### Technical Requirements

- All validation functions must have identical signatures and behavior
- Error messages must remain exactly as they are (no changes)
- 8 decimal precision for shares calculation: `math.Round((amount / price) * 1e8) / 1e8`
- All existing tests must pass without modification
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All existing tests pass without modification
- [ ] #2 Validation functions are accessible from cmd/dca/cli.go
- [ ] #3 Error messages match existing TUI behavior exactly
- [ ] #4 No breaking changes to existing functionality
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
