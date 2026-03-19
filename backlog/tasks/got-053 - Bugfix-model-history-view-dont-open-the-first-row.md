---
id: GOT-053
title: Bugfix model history view dont open the first row
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-19 17:59'
updated_date: '2026-03-19 18:11'
labels: []
dependencies: []
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
When select the first row dont open model history assets
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Implementation Plan

### 1. Technical Approach

**Problem Analysis:**
The task title "Bugfix model history view dont open the first row" indicates that when the user selects the first data row in the assets list and presses Enter, the modal history view does not open.

**Root Cause Investigation:**
After reviewing the code, I've identified two potential issues:

1. **Missing test coverage**: There are no unit tests for the `handleEnterOnAsset()` function or for pressing Enter to open the modal on the first data row.

2. **Unused SelectedIndex in modal**: The `AssetHistoryModal` struct has a `SelectedIndex` field that is never used in the rendering or navigation logic. This field is initialized to 0 in `NewAssetHistoryModal()` but never updated or used.

**Solution:**
1. Add comprehensive unit tests to verify the modal opens correctly when pressing Enter on the first data row (index 1)
2. Fix any issues found in the `handleEnterOnAsset()` or `Update()` functions
3. Optionally implement modal row selection using `SelectedIndex` if that's the intended behavior

### 2. Files to Modify

**No files to modify - this is a bug investigation and test task.**

If bugs are found, the following files would be modified:
- `internal/assets/view.go` - Fix the modal opening logic if bugs are found
- `internal/assets/view_test.go` - Add tests for modal opening behavior

### 3. Dependencies

- Existing code in `internal/assets/` package
- No external dependencies needed
- Go testing framework available

### 4. Code Patterns

Follow existing patterns in the codebase:
- Use `tea.Cmd` for returning commands from Update
- Use custom message types for state transitions
- Follow the Bubble Tea component pattern
- Use lipgloss for styling
- Use 8 decimal precision for shares, 2 decimals for currency

### 5. Testing Strategy

**Unit Tests to Add:**

1. `TestAssetsView_UpdateKeyEnter_FirstRow` - Verify modal opens when Enter is pressed on first data row (index 1)
2. `TestAssetsView_UpdateKeyEnter_MiddleRow` - Verify modal opens on middle row
3. `TestAssetsView_UpdateKeyEnter_LastRow` - Verify modal opens on last visible row (index 29)
4. `TestAssetsView_UpdateKeyEnter_HeaderRow` - Verify modal does NOT open on header row (index 0)
5. `TestAssetsView_UpdateKeyEnter_EmptyList` - Verify modal does NOT open when list is empty
6. `TestAssetsView_UpdateKeyEnter_NoEntries` - Verify modal does NOT open when no entries match selection

**Test Cases:**
- SelectedIndex = 1 (first data row) with 1+ entries → modal should open
- SelectedIndex = 0 (header) → modal should NOT open
- SelectedIndex > len(entries) → modal should NOT open
- Modal already visible → Enter should trigger LoadMore, not re-open

### 6. Risks and Considerations

- **Testing complexity**: The modal opening involves file I/O (`LoadData`), so tests need proper mocking or temp files
- **Row selection logic**: The wrap-around behavior (indices 0-29) vs actual entries needs careful testing
- **No regression risk**: Adding tests shouldn't break existing functionality

**Next Steps:**
1. Run existing tests to ensure baseline
2. Write tests for modal opening on first data row
3. Fix any issues found during testing
4. Update task with implementation notes
<!-- SECTION:NOTES:END -->
