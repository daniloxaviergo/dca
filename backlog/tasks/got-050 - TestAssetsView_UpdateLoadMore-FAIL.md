---
id: GOT-050
title: TestAssetsView_UpdateLoadMore (FAIL)
status: Done
assignee:
  - Catarina
created_date: '2026-03-19 09:43'
updated_date: '2026-03-19 11:04'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Error: view_test.go:1128: Expected 20 entries after LoadMore, got 10

Analysis: The test expects LoadMore to load the next batch of 10 days (from 10 to 20), but it's not loading any additional data. The issue is likely in the handleLoadMore method which uses a hardcoded filename dca_entries.json, but the test scenario uses a modal that was pre-populated with 10 entries. The test doesn't actually have a file with more data for the "BTC" asset.

Root Cause: The test creates a modal with 10 entries in memory, then calls Update(LoadMoreMsg{}) which triggers handleLoadMore() using dca_entries.json. However, dca_entries.json may not have additional BTC entries beyond what's already in the modal, OR the modal's state isn't properly initialized with all the data that should be in the file.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 `TestAssetsView_UpdateLoadMore` passes - LoadMore loads next batch from correct file
- [x] #2 Modal tracks source filename in `Filename` field
- [x] #3 `LoadMore()` uses stored filename to load additional data
- [x] #4 Existing functionality preserved (no regressions)

## PRD Reference

<!-- SECTION:PRD:BEGIN -->
<!-- TODO: Add PRD reference if available -->
<!-- SECTION:PRD:END -->
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
<!-- SECTION:IMPLEMENTATION:BEGIN -->
### 1. Technical Approach

**Problem Analysis**

The failing test `TestAssetsView_UpdateLoadMore` creates a temporary test file with 25 BTC entries:
1. Calls `LoadData(testFile, "BTC")` - loads first 10 days into modal
2. Calls `av.Update(LoadMoreMsg{})` - triggers `handleLoadMore()`
3. Expects 20 entries (10 + next 10)

**Root Cause**

In `handleLoadMore()`, the code calls `a.Modal.LoadMore("dca_entries.json")` using a hardcoded filename. However, the test loads data from a temporary file, and the modal has no way to remember which file was used.

**Solution: Store filename in modal state**

1. Add `Filename` field to `AssetHistoryModal` to track the source file
2. Update `LoadData()` to store the filename when loading data
3. Update `LoadMore()` to use the stored filename (ignore parameter)
4. Update `handleLoadMore()` to call `LoadMore()` with the stored filename

**Design Trade-offs**

Three options were considered:

| Option | Pros | Cons |
|--------|------|------|
| **A: Store filename in modal** (chosen) | Clean, self-contained; modal knows its data source | Adds state to modal |
| B: Pass filename to `handleLoadMore()` | No modal state change | Requires `LoadMoreMsg` to carry filename, more complex |
| C: Store filename in AssetsView | No modal state change | Duplicated state, harder to maintain |

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/assets/model.go` | Modify | Add `Filename` field to `AssetHistoryModal` struct; update `LoadData()` to store filename |
| `internal/assets/view.go` | Modify | Update `handleLoadMore()` to use modal's stored filename; update `LoadMore()` to use stored filename |

### 3. Dependencies

- **No blocking issues** - this is a focused bug fix
- **Prerequisites**: None - standalone fix
- **Setup**: None required before implementation

### 4. Code Patterns

**Modal Structure Enhancement** (in `model.go`):
```go
type AssetHistoryModal struct {
    AssetTicker   string
    Filename      string  // NEW: Store the filename used for loading data
    EntriesByDate []EntryByDate
    SelectedIndex int
    Loaded        bool
    Error         error
    Visible       bool
    Offset        int
    AllLoaded     bool
    Loading       bool
}
```

**Update LoadData()** (in `model.go`):
```go
func (m *AssetHistoryModal) LoadData(filename string, assetTicker string) error {
    // ... existing load logic ...
    
    m.AssetTicker = assetTicker
    m.Filename = filename  // NEW: Store filename for LoadMore
    m.Loaded = true
    m.Error = nil
    return nil
}
```

**Update LoadMore()** (in `model.go`):
```go
func (m *AssetHistoryModal) LoadMore(filename string) error {
    if m.AllLoaded || m.Loading {
        return nil
    }
    
    m.Loading = true
    
    // Use stored filename instead of parameter
    loadFilename := m.Filename
    if loadFilename == "" {
        loadFilename = filename  // Fallback for backward compatibility
    }
    
    // ... rest of load logic using loadFilename ...
}
```

**Update handleLoadMore()** (in `view.go`):
```go
func (a *AssetsView) handleLoadMore() (tea.Model, tea.Cmd) {
    if a.Modal == nil || !a.Modal.Visible {
        return a, nil
    }
    
    // Use modal's stored filename (LoadMore ignores parameter now)
    err := a.Modal.LoadMore(a.Modal.Filename)
    if err != nil {
        a.Modal.Error = err
    }
    
    return a, nil
}
```

**Rationale for ignoring parameter in LoadMore()**

The `filename` parameter in `LoadMore()` is now ignored because:
1. The modal already knows its source file (stored in `Filename` field)
2. Consistent with the pattern used for `AssetTicker` which is also set once in `LoadData()`
3. Simplifies the API - no need to pass filename through message types
4. Backward compatible: if `Filename` is empty, falls back to parameter value

### 5. Testing Strategy

**Unit Tests** (existing in `view_test.go`):
- `TestAssetsView_UpdateLoadMore` - Tests LoadMoreMsg loads next batch (main test, currently failing)
- `TestAssetsView_UpdateLoadMore_EmptyModal` - Tests no-op if modal is nil
- `TestAssetsView_UpdateLoadMore_ModalNotVisible` - Tests no-op if modal not visible
- `TestAssetsView_UpdateLoadMore_Error` - Tests error handling with non-existent file

**Test Coverage**:
- LoadMore with temp file containing more data (test file with 25 entries)
- LoadMore with nil modal (no-op)
- LoadMore with invisible modal (no-op)
- LoadMore with non-existent file (graceful handling, AllLoaded=true)

**Verification Steps**:
1. Run `go test -v -run TestAssetsView_UpdateLoadMore$ ./internal/assets/` (main test)
2. Run `go test -v -run TestAssetsView_UpdateLoadMore ./internal/assets/` (all 4 tests)
3. Run full test suite: `go test ./internal/assets/`
4. Run `make test` to verify all packages pass

### 6. Risks and Considerations

**Risk: Breaking existing functionality** - LOW
- The filename field is only set by `LoadData()`, which is already used in production code
- `LoadMore()` method signature remains unchanged (filename parameter still accepted, just ignored)
- No changes to public API, only internal state management
- Backward compatible: if `Filename` is empty, `LoadMore()` falls back to parameter value

**Trade-off: Filename storage** (see Section 1 for full comparison)
- Option A: Store filename in modal (chosen) - clean, self-contained modal
- Option B: Pass filename to `handleLoadMore()` - requires `LoadMoreMsg` to carry filename
- Option C: Store filename in AssetsView - duplicated state, harder to maintain

**Deployment Considerations**:
- No database migrations required (in-memory only)
- No configuration changes required
- Backward compatible (existing code continues to work)

### 7. Implementation Steps

1. **Add Filename field to `AssetHistoryModal`** in `model.go`:
   ```go
   type AssetHistoryModal struct {
       AssetTicker   string
       Filename      string  // NEW
       // ... existing fields
   }
   ```

2. **Update `LoadData()`** to store the filename after loading:
   ```go
   m.Filename = filename  // After m.AssetTicker = assetTicker
   ```

3. **Update `LoadMore()`** to use stored filename with fallback:
   ```go
   loadFilename := m.Filename
   if loadFilename == "" {
       loadFilename = filename  // Fallback
   }
   // ... use loadFilename instead of filename
   ```

4. **Update `handleLoadMore()`** in `view.go` to use modal's filename:
   ```go
   err := a.Modal.LoadMore(a.Modal.Filename)
   ```

5. **Verify `handleOpenModal()` and `handleEnterOnAsset()`** correctly call `LoadData()` (already do)

6. **Run tests** to verify fix:
   ```bash
   go test -v -run TestAssetsView_UpdateLoadMore$ ./internal/assets/
   make test
   ```
<!-- SECTION:IMPLEMENTATION:END -->
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Fixed GOT-050: TestAssetsView_UpdateLoadMore was failing because LoadMore() was using a hardcoded filename (`dca_entries.json`) instead of the filename stored when `LoadData()` was called with a test file.

## Changes Made

### 1. internal/assets/model.go
- Added `Filename` field to `AssetHistoryModal` struct to track the source filename
- Updated `LoadData()` to store the filename when loading data
- Updated `LoadMore()` to use stored filename with fallback to parameter (backward compatibility)

### 2. internal/assets/view.go  
- Updated `handleLoadMore()` to call `LoadMore()` with modal's stored filename

## Test Results

- `TestAssetsView_UpdateLoadMore`: ✅ PASS
- All 4 LoadMore-related tests: ✅ PASS
- Full test suite: ✅ PASS (59 tests total)
- CI validation (`make check`): ✅ PASS

## Root Cause

The test created a temporary file with 25 BTC entries and loaded initial 10 entries via `LoadData()`. When `LoadMore()` was called, it used a hardcoded filename `dca_entries.json` which only contained 10 BTC entries (not the 25 in the temp file). By storing the filename in modal state, LoadMore now correctly loads from the original file.

## Backward Compatibility

The implementation maintains backward compatibility:
- `LoadMore()` parameter `filename` is still accepted but ignored when `Filename` field is set
- Fallback to parameter value if `Filename` is empty
- No changes to public API signatures
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All implementation steps completed
- [x] #2 All LoadMore-related tests pass (`TestAssetsView_UpdateLoadMore*`)
- [x] #3 Full test suite passes (`make test`)
- [x] #4 Code formatted (`go fmt`)
<!-- END OF TASK FILE -->
<!-- DOD:END -->
