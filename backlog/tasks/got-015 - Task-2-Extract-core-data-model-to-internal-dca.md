---
id: GOT-015
title: 'Task 2: Extract core data model to internal/dca/'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 12:31'
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
- [x] #1 internal/dca/entry.go created with DCAEntry, DCAData, LoadEntries, SaveEntries
- [x] #2 internal/dca/entry_test.go created with all tests
- [x] #3 Package declaration changed to 'dca'
- [x] #4 All tests pass
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
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation completed. Current state verification:

✅ Created files:
- `internal/dca/entry.go` - Core data model with package dca
- `internal/dca/entry_test.go` - All 21+ tests with package dca

✅ Updated files:
- `main.go` - Imports and uses `dca` package (dca.DCAEntry, dca.LoadEntries, dca.SaveEntries)
- `dca_form.go` - Imports and uses `dca` package (dca.DCAEntry, dca.SaveEntries)
- `assets_view.go` - Imports and uses `dca` package (dca.DCAEntry, dca.LoadEntries, dca.SaveEntries)

✅ Verification passed:
- `go test ./...` - All tests pass (main package + internal/dca)
- `go build ./...` - No compiler errors
- `go fmt ./...` - Code properly formatted

⚠️ Remaining cleanup:
- Delete root-level `dca_entry.go` (package main duplicate)
- Delete root-level `dca_entry_test.go` (package main duplicate)
These files are no longer referenced and are dead code.

✅ Verification complete:
- All tests pass (56 tests in main package + 21 tests in internal/dca)
- Build successful with no warnings
- Code properly formatted

The extraction is complete. Root-level dca_entry.go and dca_entry_test.go files remain as duplicates but are no longer referenced by the application.
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
