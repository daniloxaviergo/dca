---
id: GOT-015
title: 'Task 2: Extract core data model to internal/dca/'
status: In Progress
assignee:
  - Thomas
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 11:33'
labels: []
dependencies:
  - GOT-013
references:
  - backlog/docs/doc-004.md
priority: high
ordinal: 2000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move dca_entry.go content to new package
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 internal/dca/entry.go created with DCAEntry, DCAData, LoadEntries, SaveEntries
- [ ] #2 internal/dca/entry_test.go created with all tests
- [ ] #3 Package declaration changed to 'dca'
- [ ] #4 All tests pass
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task extracts the core data model from the flat `main` package into a dedicated `internal/dca` package. The approach follows Go's standard project layout patterns:

1. **Create `entry.go`**: Copy current `dca_entry.go` content to `internal/dca/entry.go`
2. **Create `entry_test.go`**: Copy current `dca_entry_test.go` to `internal/dca/entry_test.go`
3. **Update package declaration**: Change from `package main` to `package dca` in both files
4. **Update main.go imports**: Import `internal/dca` and use fully qualified names
5. **Delete old files**: Remove `dca_entry.go` and `dca_entry_test.go` from root
6. **Verify**: Run tests and build to ensure no regressions

**Architecture Decision**: 
- Keep all functionality exactly as-is (no logic changes)
- Use `package dca` to match directory name (Go convention)
- Maintain same function/method signatures for compatibility

**Why this approach**:
- Minimal risk: Straightforward copy + package update
- Follows Go standard layout (PRD reference)
- Enables future separation of form and assets views
- No breaking changes to application behavior

### 2. Files to Modify

**New files to create:**
- `internal/dca/entry.go` - Core data model (DCAEntry, DCAData, LoadEntries, SaveEntries, methods)
- `internal/dca/entry_test.go` - All existing tests

**Files to delete:**
- `dca_entry.go` - After verifying new package works
- `dca_entry_test.go` - After verifying new package works

**Files to modify:**
- `main.go` - Update imports and references to use `dca.` package prefix
- `dca_form.go` - Update references to use `dca.` package prefix (DCAEntry, DCAData)
- `assets_view.go` - Update references to use `dca.` package prefix (DCAEntry, LoadEntries, SaveEntries, RoundTo8Decimals, CalculateWeightedAverage)

### 3. Dependencies

**Prerequisites:**
- ✅ GOT-014 complete: Folder structure created (`internal/dca/` exists)

**Blocking issues:**
- `dca_form.go` uses `DCAEntry`, `DCAData`, `CalculateSharesFromValues`, `RoundTo8Decimals`
- `assets_view.go` uses `DCAEntry`, `LoadEntries`, `SaveEntries`, `RoundTo8Decimals`, `CalculateWeightedAverage`
- All must be updated to use `dca.` package prefix

**Required setup:**
- None. Existing code is ready for refactoring.

### 4. Code Patterns

**Go conventions to follow:**
- Package name matches directory: `package dca`
- Tests co-located: `entry_test.go` in same directory
- Function/method signatures unchanged (external compatibility)
- Same validation messages and error formats

**Integration patterns:**
- Import as: `github.com/danilo/scripts/github/dca/internal/dca`
- Use qualified names: `dca.DCAEntry`, `dca.LoadEntries()`, etc.
- No changes to function behavior or signatures

**Naming:**
- Types: `DCAEntry`, `DCAData` (no prefix needed in internal package)
- Functions: `LoadEntries`, `SaveEntries`, `CalculateSharesFromValues`, `RoundTo8Decimals`
- Methods: `Validate()`, `CalculateShares()` (receiver changes to `*dca.DCAEntry`)

### 5. Testing Strategy

**Verification approach:**
1. **Copy files**: `dca_entry.go` → `internal/dca/entry.go` (with package change)
2. **Copy tests**: `dca_entry_test.go` → `internal/dca/entry_test.go` (with package change)
3. **Update main.go**: Import and use `dca` package for all references
4. **Update dca_form.go**: Import and use `dca` package
5. **Update assets_view.go**: Import and use `dca` package
6. **Delete old root files**: Remove `dca_entry.go` and `dca_entry_test.go`
7. **Run tests**: `go test ./...` to verify all pass
8. **Run build**: `go build ./...` to verify no errors

**Edge cases to cover:**
- Empty file handling (existing test: `TestLoadEntries_EmptyFile`)
- Missing file handling (existing test: `TestLoadEntries_MissingFile`)
- Invalid JSON handling (existing test: `TestLoadEntries_InvalidJSON`)
- Permission errors (existing test: `TestSaveEntries_PermissionErrorMessage`)
- Share calculation precision (existing test: `TestCalculateShares_Precision`)

**Test commands:**
```bash
# Verify internal package tests pass
go test ./internal/dca/...

# Verify full project tests pass
go test ./...

# Verify build succeeds
go build ./...

# Verify formatting
go fmt ./...
```

### 6. Risks and Considerations

**Blocking issues:**
- `main.go`, `dca_form.go`, `assets_view.go` must be updated to use `dca.` prefix
- Circular import risk if internal package tries to import main - will not occur as internal package is bottom of dependency chain

**Potential pitfalls:**
- Forgetting to update one file reference (build will fail with clear error)
- Package declaration not updated (build will fail immediately)
- Test file not copied (tests will fail)

**Trade-offs:**
- None. This is a straightforward refactoring with clear success criteria.

**Deployment considerations:**
- No deployment impact. This is an internal refactoring only.
- Application behavior unchanged.
- Data file format (`dca_entries.json`) unchanged.
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
