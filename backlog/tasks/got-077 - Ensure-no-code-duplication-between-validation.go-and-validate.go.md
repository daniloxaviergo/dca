---
id: GOT-077
title: Ensure no code duplication between validation.go and validate.go
status: To Do
assignee: []
created_date: '2026-03-28 15:11'
labels:
  - refactoring
  - duplication-elimination
  - phase2
  - req-009
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
Verify no code duplication exists between validation.go and validate.go during the transition period and create a migration plan.

## Implementation Details

### 1. Temporary Coexistence Strategy

During the transition, both files will exist:
- `validation.go`: Contains the FormModel receiver methods (now delegates)
- `validate/validate.go`: Contains the shared validation logic

**No duplication of actual logic:**
- validation.go: Only contains delegation wrapper methods
- validate/validate.go: Contains actual implementation

### 2. Migration Plan

**Phase 1: Create and test**
- Create validate/validate.go with implementation
- FormModel delegates to validate package
- All tests pass

**Phase 2: Mark deprecated**
- Add deprecation comments to validation.go methods
- Document migration path for future removal

**Phase 3: Remove (future)**
- Can remove validation.go entirely after all consumers use validate package directly
- This is a breaking change - only do when ready for API cleanup

### 3. Validation Method Comparison

**Original validation.go methods:**
```go
func (m *FormModel) validateAmount(value string) error
func (m *FormModel) validateDate(value string) error
func (m *FormModel) validateAsset(value string) error
func (m *FormModel) validatePrice(value string) error
```

**During migration (no duplication):**
```go
// validation.go - only delegation, no business logic
func (m *FormModel) validateAmount(value string) error {
    return ValidateAmount(value)  // delegated
}
// Same pattern for other methods
```

**validate/validate.go - actual implementation**
```go
func ValidateAmount(value string) error {
    // actual validation logic
}
```

### 4. Duplicate Code Detection

**Run after refactoring:**
```bash
# Check for duplicate validation logic
grep -r "must be positive" internal/form/validate/
grep -r "amount must be positive" internal/form/

# Should only find in one location after migration
```

**No duplication when:**
- validation.go contains only delegation methods
- validate.go contains all business logic
- Error messages exist in only one place

## Acceptance Criteria
<!-- AC:BEGIN -->
- ✅ No duplicate validation logic between files
- ✅ validation.go only contains delegation methods
- ✅ validate.go contains all implementation
- ✅ Error messages in single source of truth
- ✅ Migration plan documented
<!-- SECTION:DESCRIPTION:END -->

- [ ] #1 validation.go contains only delegation methods (no business logic)
- [ ] #2 validate.go contains all validation implementations
- [ ] #3 Error messages exist in single source of truth
- [ ] #4 No duplicate validation logic between files
- [ ] #5 Migration plan documented for future validation.go removal
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
