---
id: GOT-050
title: TestAssetsView_UpdateLoadMore (FAIL)
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-19 09:43'
updated_date: '2026-03-19 09:47'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Error: view_test.go:1128: Expected 20 entries after LoadMore, got 10

Analysis: The test expects LoadMore to load the next batch of 10 days (from 10 to 20), but it's not loading any additional data. The issue is likely in the handleLoadMore method which uses a hardcoded filename dca_entries.json, but the test scenario uses a modal that was pre-populated with 10 entries. The test doesn't actually have a file with more data for the "BTC" asset.

Root Cause: The test creates a modal with 10 entries in memory, then calls Update(LoadMoreMsg{}) which triggers handleLoadMore() using dca_entries.json. However, dca_entries.json may not have additional BTC entries beyond what's already in the modal, OR the modal's state isn't properly initialized with all the data that should be in the file.
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The failing test `TestAssetsView_UpdateLoadMore` creates a temporary test file with 25 BTC entries, loads the first 10 via `LoadData()`, then expects `LoadMore()` to load the next 10. However, `handleLoadMore()` in `view.go` hardcodes `dca_entries.json` instead of using the filename that was passed to `LoadData()`.

The fix involves:
1. Adding a `Filename` field to `AssetHistoryModal` to track which file was loaded
2. Updating `LoadData()` to store the filename in the modal
3. Updating `LoadMore()` to use the stored filename instead of `dca_entries.json`
4. Updating `handleLoadMore()` in `AssetsView` to use the modal's filename

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--|
| `internal/assets/model.go` | Modify | Add `Filename` field to `AssetHistoryModal` struct |
| `internal/assets/view.go` | Modify | Update `handleLoadMore()` to use modal's stored filename |
| `internal/assets/view.go` | Modify | Update `handleOpenModal()` to pass filename to modal |
| `internal/assets/view_test.go` | No change | Test is correct, implementation needs fixing |

### 3. Dependencies

- **No blocking issues** - this is a focused bug fix
- **Prerequisites**: None - standalone fix
- **Setup**: None required before implementation

### 4. Code Patterns

**Modal Structure Enhancement**:
```go
type AssetHistoryModal struct {
    AssetTicker   string
    Filename      string  // NEW: Store the filename used for loading data
    EntriesByDate []EntryByDate
    // ... other fields
}
```

**LoadData Pattern** (existing in `model.go`):
```go
func (m *AssetHistoryModal) LoadData(filename string, assetTicker string) error {
    // ... load data ...
    m.AssetTicker = assetTicker
    m.Filename = filename  // NEW: Store filename
    // ...
}
```

**LoadMore Pattern** (existing in `model.go`):
```go
func (m *AssetHistoryModal) LoadMore(filename string) error {
    // ... use m.Filename instead of filename parameter ...
}
```

**handleLoadMore Pattern** (in `view.go`):
```go
func (a *AssetsView) handleLoadMore() (tea.Model, tea.Cmd) {
    if a.Modal == nil || !a.Modal.Visible {
        return a, nil
    }
    err := a.Modal.LoadMore(a.Modal.Filename)  // Use stored filename
    // ...
}
```

### 5. Testing Strategy

**Unit Tests** (existing in `view_test.go`):
- `TestAssetsView_UpdateLoadMore` - Tests LoadMoreMsg loads next batch
- `TestAssetsView_UpdateLoadMore_EmptyModal` - Tests no-op if modal is nil
- `TestAssetsView_UpdateLoadMore_ModalNotVisible` - Tests no-op if modal not visible
- `TestAssetsView_UpdateLoadMore_Error` - Tests error handling

**Test Coverage**:
- LoadMore with temp file containing more data
- LoadMore with nil modal (no-op)
- LoadMore with invisible modal (no-op)
- LoadMore with non-existent file (graceful handling)

**Verification Steps**:
1. Run `go test -v -run TestAssetsView_UpdateLoadMore ./internal/assets/`
2. Verify all 4 LoadMore-related tests pass
3. Run full test suite: `go test ./internal/assets/`

### 6. Risks and Considerations

**Risk: Breaking existing functionality** - LOW
- The filename field is only set by `LoadData()`, which is already used in production code
- `LoadMore()` method signature accepts filename parameter, but we'll use the stored value
- No changes to public API, only internal state management

**Trade-off: Filename storage** 
- Option A: Store filename in modal (chosen) - clean, self-contained modal
- Option B: Pass filename to `handleLoadMore()` - requires method signature change
- Option C: Store filename separately in AssetsView - duplicated state

**Deployment Considerations**:
- No database migrations required (in-memory only)
- No configuration changes required
- Backward compatible (existing code continues to work)

### Implementation Steps

1. **Add Filename field to `AssetHistoryModal`** in `model.go`
2. **Update `LoadData()`** to store the filename
3. **Update `LoadMore()`** to use the stored filename
4. **Update `handleLoadMore()`** in `view.go` to use modal's filename
5. **Update `handleOpenModal()`** to ensure filename is set (currently already calls LoadData which sets it)
6. **Run tests** to verify fix
<!-- SECTION:PLAN:END -->
