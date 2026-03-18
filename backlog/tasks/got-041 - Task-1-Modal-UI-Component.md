---
id: GOT-041
title: 'Task 1: Modal UI Component'
status: In Progress
assignee: []
created_date: '2026-03-18 18:51'
updated_date: '2026-03-18 19:02'
labels:
  - ui
  - modal
dependencies: []
references:
  - backlog/docs/PRD-001-asset-history-modal.md
priority: high
ordinal: 1000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create centered modal component for displaying asset history with proper UI controls
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Modal appears centered on screen when Enter is pressed on asset row
- [ ] #2 Modal includes a title showing the asset ticker symbol
- [ ] #3 Modal includes a header row: Date | Avg Price | Total Invested | Entry Count
- [ ] #4 Modal closes when Escape key is pressed
- [ ] #5 Modal includes visual indication of active state (borders, focus)
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
<!-- IMPLEMENTATION:BEGIN -->

### 1. Technical Approach

Create a modal UI component that displays daily asset history aggregated from `dca_entries.json`. The modal will be built using the existing Bubble Tea framework patterns in `internal/form/` and `internal/assets/`.

**Key architectural decisions:**
- Modal will be a separate Bubble Tea component that can be toggled on/off in the main `AssetsView`
- Daily aggregation logic will be added to `internal/assets/` to group entries by calendar date
- Modal state will be managed within `AssetsView` alongside the table view
- Modal will show a fixed initial batch of 10 days (later tasks will add infinite scroll)

**How it works:**
1. User presses Enter on an asset row in `AssetsView` → modal opens with that asset's data
2. Modal loads entries from `dca_entries.json` for the selected asset
3. Entries are grouped by calendar date (YYYY-MM-DD)
4. For each day: weighted avg price, total invested, entry count are calculated
5. Modal displays data in a scrollable table with header row
6. User presses Escape to close modal and return to asset list

### 2. Files to Modify

| File | Action | Purpose |
|------|--------|--|
| `internal/assets/view.go` | Modify | Add modal state, modal open/close handling, modal view rendering |
| `internal/assets/model.go` | Create | New `AssetHistoryModal` struct with modal state management |
| `internal/assets/aggregate.go` | Extend | Add `AggregateByDate()` function to group entries by day |
| `internal/assets/aggregate_test.go` | Extend | Test new date aggregation function |
| `internal/assets/view_test.go` | Extend | Test modal open/close behavior |
| `cmd/dca/main.go` | Modify | Add modal state transition handling (modal open/close) |

### 3. Dependencies

- **Existing**: `dca.LoadEntries()` from `internal/dca/entry.go` for data loading
- **Existing**: `AggregateEntries()` from `internal/assets/aggregate.go` for asset-level aggregation (pattern to follow)
- **Existing**: Lipgloss styling from `internal/form/model.go` for modal borders
- **No new dependencies** - uses existing Bubble Tea + lipgloss

### 4. Code Patterns

**Follow existing patterns:**
- Modal struct with `SelectedIndex`, `Loaded`, `Error` fields like `AssetsView`
- Bubble Tea `Update()` method with key handling for `tea.KeyMsg`
- Lipgloss borders with `RoundedBorder()` and `BorderForeground(lipgloss.Color("240"))`
- Error display with ❌ prefix and error message
- Fixed-width columns with `fmt.Sprintf("%-*s", width, value)` formatting

**New modal patterns:**
- Modal state field in `AssetsView` to toggle between list and modal view
- Modal render method: centered using `lipgloss.JoinVertical(lipgloss.Center, ...)`
- Modal navigation: Up/Down to scroll, Enter to close, Escape to close

### 5. Testing Strategy

**Unit tests for new aggregation:**
- `TestAggregateByDate_SingleDay` - single day with multiple entries
- `TestAggregateByDate_MultipleDays` - multiple days sorted ascending
- `TestAggregateByDate_EmptyEntries` - returns empty slice
- `TestAggregateByDate_Calculations` - verify weighted avg price, total, count

**Integration tests for modal:**
- `TestModal_OpenOnEnter` - modal opens when Enter pressed on asset row
- `TestModal_CloseOnEscape` - modal closes when Escape pressed
- `TestModal_Render` - modal displays title, header row, data rows
- `TestModal_RenderHeader` - header row shows: Date \| Avg Price \| Total Invested \| Entry Count

**Edge cases:**
- Asset with no entries (empty modal)
- Single entry for asset
- Multiple entries on same day
- Very long history (only first 10 days shown)

### 6. Risks and Considerations

| Risk | Mitigation |
|--|--|
| Modal rendering blocks main view | Modal is optional state; list view unchanged when modal closed |
| Performance with large history | First implementation shows only 10 days; infinite scroll (Task 2) handles more |
| Date parsing performance | Simple `time.Time.Format("2006-01-02")` for grouping; acceptable for <10K entries |
| Border styling conflicts | Use existing lipgloss patterns from `internal/form/` |
| State synchronization issues | Modal receives asset ticker, loads fresh data from file each open |
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
