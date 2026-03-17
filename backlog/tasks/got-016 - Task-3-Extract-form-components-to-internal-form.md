---
id: GOT-016
title: 'Task 3: Extract form components to internal/form/'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 12:53'
labels: []
dependencies:
  - GOT-013
references:
  - backlog/docs/doc-004.md
priority: high
ordinal: 3000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move dca_form.go content to new package
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/form/model.go created with FormModel and related types
- [ ] #2 internal/form/validation.go created with validation functions
- [ ] #3 internal/form/validation_test.go created with all tests
- [ ] #4 Package declaration changed to 'form'
- [ ] #5 All tests pass
- [ ] #6 #1 internal/form/model.go created with FormModel and related types
- [ ] #7 #2 internal/form/validation.go created with validation functions
- [ ] #8 #3 internal/form/validation_test.go created with all tests
- [ ] #9 #4 Package declaration changed to form
- [ ] #10 #5 All tests pass
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task extracts form components from the flat `main` package into a dedicated `internal/form` package, following the same pattern established in task GOT-015 for extracting the DCA data model.

**Extraction Strategy:**
1. **Create `model.go`**: Copy current `dca_form.go` content to `internal/form/model.go` with package `form`
2. **Create `validation.go`**: Extract validation functions (`validateAmount`, `validateDate`, `validateAsset`, `validatePrice`) to `internal/form/validation.go`
3. **Create `validation_test.go`**: Copy and refactor `dca_form_test.go` to test validation functions with package `form`
4. **Update `main.go`**: Import `internal/form` and use `form.FormModel`, `form.NewFormModel`, etc.
5. **Update `dca_form.go`**: Remove after extraction (becomes duplicate/legacy)
6. **Verify**: Run tests and build to ensure no regressions

**Architecture Decisions:**
- Split validation into separate `validation.go` file for better organization (separates validation logic from UI model)
- Keep `model.go` for `FormModel` struct and Bubble Tea methods
- Use `package form` to match directory name (Go convention)
- Preserve all existing functionality without logic changes

**Why this approach:**
- Consistent with GOT-015 pattern (extracted DCA model to `internal/dca`)
- Enables future separation of form from assets view
- Minimal risk: straightforward copy + package update
- Clear file organization: validation functions are naturally separable from model methods

### 2. Files to Modify

**Created:**
- `internal/form/model.go` - FormModel struct, Bubble Tea methods, View rendering (package `form`)
- `internal/form/validation.go` - Validation functions (package `form`)
- `internal/form/validation_test.go` - Unit tests for validation functions (package `form`)

**Modified:**
- `main.go` - Import `internal/form`, replace `FormModel` → `form.FormModel`, `NewFormModel` → `form.NewFormModel`, `FormStep` → `form.FormStep`, etc.
- `dca_form.go` - Remove after verification (becomes duplicate/legacy)

**Deleted (after verification):**
- `dca_form.go` - No longer needed (moved to `internal/form/`)
- `dca_form_test.go` - No longer needed (moved to `internal/form/validation_test.go`)

### 3. Dependencies

**Prerequisites:**
- ✅ `internal/form/` directory must exist (created in GOT-013)
- ✅ `internal/dca/` package must be in place (completed in GOT-015)
- ✅ All existing tests must pass before extraction

**No external dependencies** required for this task.

### 4. Code Patterns

**Follow these conventions from the existing codebase:**

1. **Package declaration**: `package form` in all files
2. **Type naming**: `FormModel`, `FormStep`, `FormField` (same as current)
3. **Validation function signatures**:
   ```go
   func (m *FormModel) validateAmount(value string) error
   func (m *FormModel) validateDate(value string) error
   func (m *FormModel) validateAsset(value string) error
   func (m *FormModel) validatePrice(value string) error
   ```
4. **Helper functions**: Exported as `CalculateSharesFromValues`, `RoundTo8Decimals`
5. **Test naming**: `Test{Function}_{Condition}` pattern
6. **Error messages**: Exact text matching required (e.g., "Amount must be positive")

**Integration patterns:**
- Import path: `github.com/danilo/scripts/github/dca/internal/form`
- Usage: `form.FormModel`, `form.NewFormModel()`, `form.FormStep`
- All Bubble Tea methods remain on `FormModel` (Update, View, Init)

### 5. Testing Strategy

**Test coverage for `internal/form/validation_test.go`:**

1. **Validation function tests** (all existing tests from `dca_form_test.go`):
   - `TestFormModel_ValidateAmount_Pass`
   - `TestFormModel_ValidateAmount_RejectZero`
   - `TestFormModel_ValidateAmount_RejectNegative`
   - `TestFormModel_ValidateAmount_RejectEmpty`
   - `TestFormModel_ValidateAmount_RejectInvalid`
   - `TestFormModel_ValidateDate_Pass`
   - `TestFormModel_ValidateDate_RejectInvalid`
   - `TestFormModel_ValidateAsset_Pass`
   - `TestFormModel_ValidateAsset_RejectEmpty`
   - `TestFormModel_ValidateAsset_RejectWhitespace`
   - `TestFormModel_ValidatePrice_Pass`
   - `TestFormModel_ValidatePrice_RejectZero`
   - `TestFormModel_ValidatePrice_RejectNegative`
   - All exact error message tests

2. **Helper function tests**:
   - `TestCalculateSharesFromValues`
   - `TestCalculateSharesFromValues_Precision`
   - `TestRoundTo8Decimals`

3. **Form model tests**:
   - `TestFormModel_GetFieldFloat64`
   - `TestFormModel_GetFieldFloat64_Empty`
   - `TestFormModel_GetCurrentFieldKey`
   - `TestFormModel_TabForward`
   - `TestFormModel_TabBackward`
   - `TestFormModel_HandleBackspace`
   - `TestFormModel_HandleInput`
   - `TestFormModel_RenderForm`
   - `TestFormModel_InlineErrorDisplay`

**Approach:**
- All existing tests will be adapted to use `form.FormModel` instead of `FormModel`
- All existing tests will be adapted to use `form.NewFormModel` instead of `NewFormModel`
- Table-driven tests preserved as-is
- Exact error message assertions preserved as-is
- Test files use `package form` and import `testing`

**Verification:**
- `go test ./...` - All 46+ tests must pass (46 in main + 21 in internal/dca + 46 in internal/form)
- `go build ./...` - No compiler errors
- `go fmt ./...` - Code properly formatted

### 6. Risks and Considerations

**Blocking issues:**
- ⚠️ **None identified** - Straightforward extraction with no dependencies on other incomplete tasks

**Potential pitfalls:**
- ⚠️ **Import path consistency**: Must use correct import path `github.com/danilo/scripts/github/dca/internal/form` in all files
- ⚠️ **Test file migration**: Need to update all `FormModel` → `form.FormModel` references in test files
- ⚠️ **Build caching**: Run `go clean -cache` if unexpected behavior occurs

**Trade-offs:**
- ⚠️ **dca_form.go cleanup**: Root-level `dca_form.go` will remain as duplicate until manually deleted (follows GOT-015 pattern)
- ⚠️ **No breaking changes**: All functionality preserved; no logic modifications allowed

**Rollout considerations:**
- Task can be completed incrementally (model first, then validation)
- Tests will fail during migration until all references updated
- Must verify `go build ./...` and `go test ./...` pass before marking complete

**Definition of Done verification:**
- [ ] #1 `internal/form/model.go` created with `FormModel` and related types
- [ ] #2 `internal/form/validation.go` created with validation functions
- [ ] #3 `internal/form/validation_test.go` created with all tests
- [ ] #4 Package declaration changed to `form`
- [ ] #5 All tests pass
- [ ] #6 `go fmt` passes without errors
- [ ] #7 `go build ./...` succeeds with no warnings
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
