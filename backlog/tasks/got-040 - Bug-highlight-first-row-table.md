---
id: GOT-040
title: Bug highlight first row table
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-18 15:13'
updated_date: '2026-03-18 15:54'
labels: []
dependencies: []
ordinal: 1000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
the row highlight the table is correct? When go to second row i cant go back to the first row.
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

**Problem:** The table navigation in `AssetsView` has a bug where the user cannot navigate from the second row (index 1) back to the first row (index 0 - the header). The current `handleUp()` logic wraps from index 1 directly to the last row (index 29) instead of decrementing to 0.

**Root Cause:** In `internal/assets/view.go`, the `handleUp()` function uses the condition `if a.SelectedIndex <= 1` to determine wrapping behavior. This causes two issues:
- When `SelectedIndex` is 1 (first data row), it wraps to 29 instead of decrementing to 0 (header)
- The header row (index 0) and first data row (index 1) have different behaviors that create confusion

**Fix:** Modify the `handleUp()` condition so that:
- `SelectedIndex` 1 → decrements to 0 (header)
- `SelectedIndex` 0 → wraps to 29 (last visible row)

This creates a logical navigation flow: header (0) ↔ data rows (1-29) with wrap-around.

**Files to Modify:**
1. `internal/assets/view.go` - Fix the `handleUp()` logic

**Dependencies:**
- None. This is a self-contained bug fix.

**Code Pattern:**
- Match existing code style (lipgloss, Bubble Tea conventions)
- Preserve the wrap-around behavior for header row
- Ensure `handleDown()` logic remains consistent

**Testing Strategy:**
- Existing tests in `view_test.go` cover navigation
- New tests should verify: can go from row 1 to row 0 (header)
- Run `make test` to verify all tests pass

**Risks/Considerations:**
- The current test `TestAssetsView_NavigateUp` expects wrapping from index 1 to 29 - this test will need to be updated to reflect the corrected behavior
- The `handleDown()` logic may need review for symmetry

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `internal/assets/view.go` | Modify `handleUp()` function | Fix the wrapping condition to allow decrementing from index 1 to 0 |

### 3. Dependencies

- No prerequisites. This is a self-contained bug fix.

### 4. Code Patterns

**Current Pattern (buggy):**
```go
if a.SelectedIndex <= 1 {
    a.SelectedIndex = 29  // Wraps from both 0 and 1
} else {
    a.SelectedIndex--
}
```

**Proposed Fix:**
```go
if a.SelectedIndex == 0 {
    a.SelectedIndex = 29  // Only wrap from header
} else {
    a.SelectedIndex--
}
```

This ensures:
- Index 1 → 0 (move to header)
- Index 0 → 29 (wrap to last row)

### 5. Testing Strategy

**Tests to Verify:**
1. `TestAssetsView_NavigateUp` - Will need update; currently expects wrap from 1→29, should expect 1→0
2. `TestAssetsView_UpdateArrowUp` - Should now work correctly from any data row to header
3. `TestAssetsView_NavigateWrapUp` - Update to expect 1→29 when starting from index 1

**Test Commands:**
```bash
make test
```

**Expected Behavior After Fix:**
- Pressing Up at any data row (1-29) should decrement to the previous row
- Pressing Up at header (index 0) should wrap to last row (index 29)
- Pressing Down at last row (index 29) should wrap to first data row (index 1)
- Header row (index 0) should be accessible via Up arrow from row 1

### 6. Risks and Considerations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Existing tests expect current (buggy) behavior | Test failures | Update tests to match corrected behavior |
| User expectations based on current UI | Confusion | Fix improves UX by making navigation intuitive |

**Notes:**
- The table has 30 visible rows: index 0 (header) + indices 1-29 (data/empty rows)
- The bug specifically affects navigation from the first data row to the header
- No changes to rendering logic needed - only input handling

**Implementation Steps:**
1. Edit `handleUp()` in `internal/assets/view.go` to only wrap when `SelectedIndex == 0`
2. Update related tests in `view_test.go` to expect the corrected behavior
3. Run `make test` to verify all tests pass
4. Run `make check` for full CI validation

Ready for your approval. Should I proceed with the implementation?
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
