---
id: GOT-014
title: 'Task 1: Create folder structure for modular refactoring'
status: Done
assignee:
  - Thomas
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 11:26'
labels: []
dependencies: []
references:
  - backlog/docs/doc-004.md
priority: high
ordinal: 1000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/ and internal/ directories
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 All directories created: cmd/dca/, internal/dca/, internal/form/, internal/assets/
- [x] #2 No files moved or deleted
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task creates the foundational folder structure for the modular refactoring project. The approach is straightforward:

1. **Create directory structure**: Using `mkdir -p` to create all required directories
2. **No file operations**: This task is purely structural - no code will be moved
3. **Verify structure**: Use `ls -R` to confirm all directories exist
4. **Test compatibility**: Run `go build` to ensure the flat structure still works before proceeding

**Architecture Decision**: 
- Use `cmd/` for application entry points (standard Go layout)
- Use `internal/` for private packages that cannot be imported externally
- Subdirectories under `internal/` align with domain boundaries: `dca`, `form`, `assets`

**Why this approach**:
- Minimal risk: Only creates directories, no file moves or code changes
- Follows Go project layout conventions (see PRD reference)
- Prepares structure for subsequent tasks (GOT-015, GOT-016, GOT-017)
- No breaking changes - existing code continues to work

### 2. Files to Modify

**New directories to create:**
- `cmd/` (parent directory for all command-line entry points)
- `cmd/dca/` (main application entry point location)
- `internal/` (private packages that cannot be imported externally)
- `internal/dca/` (core data model and file I/O)
- `internal/form/` (form UI components and validation)
- `internal/assets/` (asset display and aggregation)

**No files to create or modify** in this task.

### 3. Dependencies

**Prerequisites:**
- None. This is the first task in the refactoring sequence.

**Blocking issues:**
- None identified.

**Required setup:**
- None. The existing codebase is in a flat structure ready for refactoring.

### 4. Code Patterns

**Go project layout conventions to follow:**
- `cmd/` contains application entry points (one directory per binary)
- `internal/` contains packages that cannot be imported by external code
- Package names match directory names (`package dca`, `package form`, `package assets`)
- Tests co-located with source files (e.g., `entry_test.go` in same directory)

**Future task alignment:**
- Task 2 (GOT-015): Move `dca_entry.go` → `internal/dca/entry.go`
- Task 3 (GOT-016): Move `dca_form.go` → `internal/form/` (split into model/view)
- Task 4 (GOT-017): Move `assets_view.go` → `internal/assets/` (split into view/aggregate)
- Task 5 (GOT-018): Move `main.go` → `cmd/dca/main.go`

### 5. Testing Strategy

**Verification approach:**
1. **Directory existence check**: Verify all 6 directories are created using `ls -R`
2. **Build verification**: Run `go build` to ensure the current flat structure is unaffected
3. **No test changes required**: This task has no code changes, so no new tests needed

**Test commands:**
```bash
# Verify directory structure
ls -R cmd internal

# Verify build still works
go build -o dca

# Run existing tests
go test ./...
```

**Edge cases to verify:**
- All directories exist with proper permissions
- No files were accidentally deleted or moved
- The application can still build with the existing flat structure

### 6. Risks and Considerations

**Blocking issues:**
- None identified.

**Potential pitfalls:**
- On Windows: `mkdir -p` may need `New-Item` instead (but project is on Linux)
- Permission issues: `mkdir` may fail if directory exists with restricted permissions (unlikely)

**Trade-offs:**
- None. This is a minimal, low-risk task.

**Deployment considerations:**
- No deployment impact. This task only creates directories.
- The application continues to work with the existing flat structure.
- Subsequent tasks will move files and update imports.

**Definition of Done checklist:**
- [ ] All 6 directories created: `cmd/dca/`, `internal/dca/`, `internal/form/`, `internal/assets/`
- [ ] No files moved or deleted
- [ ] `go build` succeeds
- [ ] `go test ./...` passes (no new failures)
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Created directory structure: cmd/dca/, internal/dca/, internal/form/, internal/assets/

Verified build with go build -o dca (no errors)

Verified tests with go test ./... (0.006s, all pass)

Verified formatting with go fmt ./... (no changes needed)

No files moved or deleted - existing flat structure intact
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
