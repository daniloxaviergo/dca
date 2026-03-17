---
id: GOT-017
title: 'Task 4: Extract assets view to internal/assets/'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 11:20'
updated_date: '2026-03-17 15:03'
labels: []
dependencies:
  - GOT-013
references:
  - backlog/docs/doc-004.md
priority: high
ordinal: 5000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move assets_view.go content to new package
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 internal/assets/view.go created with AssetsView
- [x] #2 internal/assets/aggregate.go created with aggregation functions
- [x] #3 internal/assets/aggregate_test.go created with all tests
- [x] #4 Package declaration changed to 'assets'
- [x] #5 All tests pass
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Extract `assets_view.go` into `internal/assets/` package following the established pattern from GOT-015 (DCA) and GOT-016 (Form).

**Key observations:**
- `assets_view.go` contains data aggregation functions AND Bubble Tea UI component
- `RoundTo8Decimals` function is duplicated across packages - NOT consolidating in this task
- Tests in `assets_view_test.go` must be moved and updated

**Approach:**
1. Create `internal/assets/view.go` with `AssetSummary`, `AssetsViewModel`, `AssetsView` structs and Bubble Tea methods
2. Create `internal/assets/aggregate.go` with `LoadAndAggregateEntries`, `aggregateEntries`, `CalculateWeightedAverage`
3. Move tests: `view_test.go` (UI tests) and `aggregate_test.go` (aggregation tests)
4. Update imports in `main.go`

**Trade-offs:**
- Will NOT remove `RoundTo8Decimals` duplication (scope creep)
- Tests split by concern: UI vs data aggregation

### 2. Files to Modify

| Action | File |
|--------|------|
| **Create** | `internal/assets/view.go` - UI component (AssetSummary, AssetsView, Bubble Tea methods) |
| **Create** | `internal/assets/aggregate.go` - Aggregation functions (LoadAndAggregateEntries, CalculateWeightedAverage) |
| **Create** | `internal/assets/view_test.go` - UI/component tests |
| **Create** | `internal/assets/aggregate_test.go` - Aggregation/data tests |
| **Delete** | `assets_view.go` - Remove after successful migration |
| **Delete** | `assets_view_test.go` - Replace with new test files |
| **Modify** | `main.go` - Add import for `internal/assets`, update references |

### 3. Dependencies

**Prerequisites:**
- ✅ GOT-013 (folder structure)
- ✅ GOT-015 (DCA extraction)
- ✅ GOT-016 (Form extraction)

**No additional dependencies required.**

### 4. Code Patterns

**Follow established patterns from `internal/form/` and `internal/dca/`:**

1. **Package declaration**: `package assets`
2. **File naming**: `view.go` (UI), `aggregate.go` (data logic)
3. **Type organization**: Data model → View model → Bubble Tea component
4. **Bubble Tea patterns**: Init/Update/View with handlers
5. **Helper functions**: Exported with `Calculate` prefix
6. **Validation**: `Validate() error` method on structs
7. **Test naming**: `Test{Function}_{Condition}` pattern

**Note**: `RoundTo8Decimals` will be duplicated (acceptable per task scope).

### 5. Testing Strategy

**Test file split:**
- `internal/assets/view_test.go`: UI/component tests (Bubble Tea navigation, rendering)
- `internal/assets/aggregate_test.go`: Aggregation/data tests (file loading, weighted average calculations)

**Testing approach:**
1. Keep exact same test functions with same assertions
2. Update imports to use `internal/dca` package
3. Run `go test ./internal/assets/...` and `go test ./...` to verify

**Coverage areas:**
- Validation tests (empty ticker, negative values)
- Weighted average calculations (precision, zero shares)
- File loading (empty, missing, populated)
- Aggregation (single/multiple assets, multiple entries)
- UI rendering (empty state, table with data)
- Navigation (up/down, wrap-around, empty list)
- Key handling (Esc, Ctrl+C, arrows)

### 6. Risks and Considerations

**Blocking issues:**
- None identified - straightforward file extraction

**Potential pitfalls:**
1. Import path in tests: Must use full path `github.com/danilo/scripts/github/dca/internal/dca`
2. Duplicate `RoundTo8Decimals`: Will exist in both `form` and `assets` packages (acceptable)

**Design decisions:**
1. Split tests: UI in `view_test.go`, aggregation in `aggregate_test.go`
2. No deduplication: `RoundTo8Decimals` remains duplicated (future cleanup task)
3. Package name: `assets` (not `assetview`) to match folder `internal/assets/`

**Post-implementation verification:**
```bash
go test ./internal/assets/...
go test ./...
go build ./...
```
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
