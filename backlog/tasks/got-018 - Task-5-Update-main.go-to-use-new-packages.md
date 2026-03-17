---
id: GOT-018
title: 'Task 5: Update main.go to use new packages'
status: Done
assignee:
  - Thomas
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 15:36'
labels: []
dependencies:
  - GOT-013
  - GOT-014
  - GOT-015
  - GOT-016
references:
  - backlog/docs/doc-004.md
priority: medium
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move main.go to cmd/ and update imports
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 main.go moved to cmd/dca/main.go
- [x] #2 Import statements added for internal packages
- [x] #3 All references updated to use fully qualified names
- [x] #4 Application behavior unchanged
- [x] #5 All tests pass
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task moves `main.go` from the project root to `cmd/dca/main.go` following Go's standard project layout pattern. No code logic changes are required as the file already uses the `internal/` packages with full module paths.

**Approach:**
1. Move `main.go` → `cmd/dca/main.go`
2. Verify imports (already correct: `internal/dca`, `internal/form`, `internal/assets`)
3. Build and test the application
4. Run all tests to ensure no regressions

**Architecture Decision:**
- Follow Go standard `cmd/` layout pattern where each subdirectory contains a single `main.go`
- Keep `internal/` packages unchanged (already correctly structured)

**Why this approach:**
- Minimal risk - only file movement, no code changes
- Enables multi-command support for future development
- Consistent with previous extraction tasks (GOT-015, GOT-016, GOT-017)

### 2. Files to Modify

| Action | File | Reason |
|--------|------|--------|
| **Move** | `main.go` → `cmd/dca/main.go` | Move entry point to standard location |
| **Create** | (none) | No new files needed |
| **Delete** | (none) | No files to delete |
| **Modify** | (none) | Imports already correct |

### 3. Dependencies

**Prerequisites:**
- `cmd/` directory created (GOT-013)
- `internal/dca/` package (GOT-015)
- `internal/form/` package (GOT-016)
- `internal/assets/` package (GOT-017)

**No additional dependencies required.**

### 4. Code Patterns

**Follow Go standard layout conventions:**
- `cmd/<command>/main.go` - Application entry points
- `internal/<package>/` - Private packages with full module path imports
- `package main` for entry points
- Bubble Tea Model/Update/View patterns

### 5. Testing Strategy

**Verification commands:**
```bash
go build ./cmd/dca
go build ./...
go test ./cmd/dca
go test ./internal/...
go test ./...
go fmt ./cmd/dca/...
```

**Runtime verification:**
- Application starts without errors
- Form entry creation works
- Assets view displays correctly
- State transitions function properly
- Ctrl+C exit works

### 6. Risks and Considerations

**Blocking issues:**
- None identified.

**Potential pitfalls:**
1. Build path confusion - Mitigation: Use `go build ./cmd/dca` explicitly
2. Working directory - Mitigation: Relative paths use current directory (no change)

**Verification checklist:**
- [ ] `go build ./cmd/dca` succeeds
- [ ] `go test ./...` passes
- [ ] Application runs and shows form
- [ ] `go fmt` passes
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Implementation Summary

### Files Moved
- `main.go` → `cmd/dca/main.go`
- `dca_form.go` → `cmd/dca/dca_form.go`
- `dca_form_test.go` → `cmd/dca/dca_form_test.go`

### Reasoning
The form files (`dca_form.go`, `dca_form_test.go`) needed to be moved along with `main.go` because:
1. They use `package main` (same as main.go)
2. They reference `formSubmittedMsg` type which is defined in main.go
3. Keeping them at root would cause compilation errors (undefined type reference)

### Verification Results
- **Build**: `go build ./cmd/dca` ✅
- **Tests**: `go test ./...` ✅ (all 52 tests pass)
- **Format**: `go fmt ./cmd/dca/...` ✅
- **Vet**: `go vet ./cmd/dca/...` ✅

### Project Structure After Changes
```
cmd/dca/
├── main.go              # Entry point
├── dca_form.go          # Form model and UI
└── dca_form_test.go     # Form tests
internal/
├── assets/              # Assets view package
├── dca/                 # DCA data model package
└── form/                # Form components package
```
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Moved main application entry point to `cmd/dca/` directory following Go's standard project layout pattern. All application code now resides in the `cmd/` subdirectory, with internal packages in `internal/`.

## Changes Made

| File | Action | Reason |
|------|--------|--------|
| `main.go` | Moved to `cmd/dca/main.go` | Standard Go layout for entry points |
| `dca_form.go` | Moved to `cmd/dca/dca_form.go` | Uses `package main` and references `formSubmittedMsg` from main.go |
| `dca_form_test.go` | Moved to `cmd/dca/dca_form_test.go` | Tests for form files |

## Verification Results

- ✅ **Build**: `go build ./cmd/dca` succeeds
- ✅ **Tests**: `go test ./...` passes (52 tests)
- ✅ **Format**: `go fmt ./cmd/dca/...` passes
- ✅ **Vet**: `go vet ./cmd/dca/...` no warnings

## New Project Structure

```
cmd/dca/
├── main.go              # Application entry point
├── dca_form.go          # Form model and UI
└── dca_form_test.go     # Form tests
internal/
├── assets/              # Assets view package
├── dca/                 # DCA data model package
└── form/                # Form components package
backlog/                 # Task management
```

## Risk Assessment

**Risk**: Low - File movement only, no code logic changes required. Imports were already using fully qualified paths to internal packages.

**Follow-ups**: None identified. Ready for GOT-019 (future feature work).
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
