---
id: GOT-043
title: 'Task 3: Infinite Scroll Implementation'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-18 18:51'
updated_date: '2026-03-18 21:19'
labels:
  - ui
  - infinite-scroll
dependencies: []
references:
  - backlog/docs/PRD-001-asset-history-modal.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add infinite scroll to load historical data in batches of 10 days
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Display initial batch of 10 days of history on modal open
- [ ] #2 Show loading state when fetching more data
- [ ] #3 Load next batch of 10 days when user scrolls to bottom
- [ ] #4 Disable scroll trigger once all data is loaded
- [ ] #5 Handle empty history state gracefully
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
# Implementation Plan: Infinite Scroll for Asset History Modal

### 1. Technical Approach

Add infinite scroll to the asset history modal by:
1. **Batch loading**: Load history in batches of 10 days as configured
2. **Scroll detection**: Track user scroll position and detect when bottom is reached
3. **Lazy loading**: Fetch next batch when user scrolls to bottom of modal
4. **State management**: Track current batch, total available days, and loading state
5. **UI feedback**: Show loading indicator during data fetch

**Architecture decisions:**
- Keep modal separate from main list view (already implemented)
- Reuse existing `AssetHistoryModal` struct with pagination fields
- Append new batches to existing `EntriesByDate` slice
- No data re-fetching on each scroll - cache all loaded data

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/assets/model.go` | Modify | Add pagination fields (offset, loadedCount, allLoaded) to `AssetHistoryModal` |
| `internal/assets/view.go` | Modify | Add scroll navigation, loading state, and batch fetch logic |
| `internal/assets/aggregate.go` | Modify | Add function to fetch next batch of entries |
| `internal/assets/view_test.go` | Add | Test infinite scroll functionality |

### 3. Dependencies

- ✅ Modal UI component (Task 1 - GOT-041): Already implemented
- ✅ Daily aggregation data fetching (Task 2 - GOT-042): Already implemented
- ✅ Data model (`EntryByDate`): Already in place
- **Required before**: Modal must have scrollable content area (already has table layout)

### 4. Code Patterns

**Follow existing patterns:**
- Use Bubble Tea message types for state transitions
- Lipgloss for styling (already used throughout)
- Error handling with `Error` field on modal
- Load data in `LoadData()` method (already pattern)

**New patterns:**
- `LoadMoreMsg` message type to trigger batch fetch
- Scroll position tracking with `SelectedIndex`
- `Visible` field already exists for modal state management

### 5. Testing Strategy

**Unit tests to add:**
- `TestAssetHistoryModal_LoadMore_LoadsNextBatch`: Verify next 10 days load correctly
- `TestAssetHistoryModal_LoadMore_AllLoaded`: Verify trigger disables when all data loaded
- `TestAssetHistoryModal_LoadMore_EmptyHistory`: Handle empty history gracefully
- `TestAssetHistoryModal_NavigateScrollDown`: Verify scroll navigation works
- `TestAssetsView_UpdateLoadMore`: Verify modal sends load more message

**Test edge cases:**
- Modal with 0-9 days (less than batch size)
- Modal with exactly 10 days
- Modal with 11+ days (multiple batches)
- User scrolls when all data already loaded
- Error during batch fetch

### 6. Risks and Considerations

**Potential issues:**
1. **Scroll detection**: Modal is a table with fixed rows; need to track which row user is on
   - *Mitigation*: Use `SelectedIndex` to track position; when near bottom, show "load more" option
2. **Loading state**: Need visual indicator during fetch
   - *Mitigation*: Add `Loading` field to modal, show "Loading more..." message
3. **Memory growth**: All data stays in memory
   - *Mitigation*: Accept trade-off for simplicity; batch size is small (10 days)
4. **No data left to load**: Need to detect when all days loaded
   - *Mitigation*: Compare `len(EntriesByDate)` with total available days from `AggregateByDate()`

**Design trade-offs:**
- **Approach A**: Load all data upfront, paginate in UI
  - Pros: Simpler, no network delay
  - Cons: Slower initial load, higher memory usage
- **Approach B**: Fetch batches from disk/file on demand (current plan)
  - Pros: Faster initial load, lower memory footprint for large histories
  - Cons: More complex state management, potential I/O delay
- **Decision**: Approach B because PRD specifies "infinite scroll" for large datasets

**Acceptance criteria mapping:**
- ✅ #1 Display initial batch of 10 days: Already implemented (line 48-50 in model.go)
- ⏳ #2 Show loading state: Add `Loading` field and UI
- ⏳ #3 Load next batch on scroll: Add scroll-to-bottom detection
- ⏳ #4 Disable trigger when all loaded: Add `AllLoaded` flag
- ⏳ #5 Handle empty history: Already handled (returns empty slice)

**Blocking issues:**
- Modal scroll navigation not yet implemented (table rows don't support scrolling)
- Need to determine scroll trigger mechanism (Enter key? Down arrow at bottom?)
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 All acceptance criteria met
- [ ] #8 Unit tests pass
- [ ] #9 No new compiler warnings
- [ ] #10 Code follows project style (go fmt)
<!-- DOD:END -->
