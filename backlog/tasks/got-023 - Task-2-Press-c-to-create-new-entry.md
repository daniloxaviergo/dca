---
id: GOT-023
title: 'Task 2: Press ''c'' to create new entry'
status: Done
assignee:
  - Thomas
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 18:08'
labels: []
dependencies:
  - GOT-022
references:
  - internal/assets/view.go
  - cmd/dca/main.go
documentation:
  - doc-005
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add 'c' key handler in asset list to switch to form view. Modify internal/assets/view.go Update() to handle 'c' key press and return ViewTransitionMsg to trigger state change to StateForm.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Pressing 'c' in asset list switches to form view
- [x] #2 Form fields are reset (amount=empty, asset=empty, price=empty)
- [x] #3 Date defaults to current timestamp in RFC3339 format
- [x] #4 User can navigate form with Tab/Enter as before
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### Implementation Plan: Add 'c' Key Handler for Creating New Entry

---

### 1. Technical Approach

The task requires adding a 'c' key handler in the asset list view (`AssetsView`) that transitions to the form view (`StateForm`). 

**How it works:**
1. The `AssetsView.Update()` method needs to handle `tea.KeyRunes` for the 'c' key
2. When 'c' is pressed, return a `ViewTransitionMsg{View: "form"}` to trigger state change
3. The main `model.Update()` in `main.go` already handles `ViewTransitionMsg` to switch from `StateAssetsView` to `StateForm`
4. When transitioning to form, create a new `FormModel` with reset fields and default timestamp

**Architecture decisions:**
- Reuse existing `ViewTransitionMsg` for state transitions (already used for Esc/Ctrl+C)
- Follow existing pattern: asset list → `ViewTransitionMsg` → form initialization
- Reset form fields by creating new `FormModel` via `NewFormModel()` with fresh data

**Why this approach:**
- Minimal code changes (only 3 files modified)
- Reuses existing transition mechanism
- Maintains consistency with current Esc/Ctrl+C behavior
- No new message types needed

---

### 2. Files to Modify

| File | Action | Description |
|------|--------|-------------|
| `internal/assets/view.go` | Modify | Add 'c' key handler in `Update()` method |
| `cmd/dca/main.go` | Modify | Update `ViewTransitionMsg` handling for `StateAssetsView` case to initialize form |
| `internal/assets/view_test.go` | Modify | Add test for 'c' key handler |

**Files to read (no changes):**
- `internal/form/model.go` - Already has `NewFormModel()` that resets fields
- `internal/form/validation.go` - Validation logic (no changes needed)

---

### 3. Dependencies

- **GOT-022 must be complete**: App must start in `StateAssetsView` (already implemented in `main.go`)
- **Existing infrastructure**: 
  - `ViewTransitionMsg` type already defined in `internal/assets/view.go`
  - Main model already handles `ViewTransitionMsg` for state transitions
  - `NewFormModel()` already resets fields and sets default date

- **No new dependencies required**

---

### 4. Code Patterns

**Follow these patterns from existing code:**

1. **Key handling in `AssetsView.Update()`:**
   ```go
   case tea.KeyMsg:
       switch msg.Type {
       case tea.KeyCtrlC, tea.KeyEsc:
           // Return ViewTransitionMsg
       case tea.KeyRunes:
           // Check for 'c' key
       }
   ```

2. **Main model state transition:**
   ```go
   case StateAssetsView:
       if _, ok := msg.(assets.ViewTransitionMsg); ok {
           m.currentState = StateForm
           m.form = form.NewFormModel(m.entries, defaultEntriesPath)
           return m, nil
       }
   ```

3. **FormModel initialization pattern** (already works):
   - `NewFormModel()` creates fields with empty values
   - Date defaults to current RFC3339 timestamp
   - All validation errors cleared

---

### 5. Testing Strategy

**Unit tests to add/modify:**

1. **`TestAssetsView_UpdateKeyC`** (new test in `view_test.go`)
   - Press 'c' key in asset list
   - Verify `ViewTransitionMsg` is returned
   - Verify `cmd()` function returns the message

2. **`TestAssetsView_UpdateKeyC_IgnoredWhenNotLoaded`** (new test)
   - Verify 'c' key works even when `Loaded` is false
   - Edge case: empty asset list

3. **`TestAssetsView_UpdateKeyC_IgnoresCase`** (new test optional)
   - Consider whether 'C' (uppercase) should also work

**Test coverage:**
- 'c' key in `AssetsView.Update()` returns correct message type
- State transition in `model.Update()` switches to `StateForm`
- Form initialization creates fresh fields with default date

---

### 6. Risks and Considerations

**No significant risks identified:**

- **Blocking issues**: None
- **Trade-offs**: None
- **Deployment considerations**: 
  - Safe to merge immediately (no data persistence changes)
  - No migration required
  - Backward compatible with existing data

**Design notes:**
- 'c' key is intuitive for "create" action
- No modifier keys required (simple UX)
- Existing transition mechanism handles all edge cases
- Form is fully reset on each transition (clean state)
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research complete. Implementation plan drafted. Ready for user approval before coding begins.

Key findings:
- AssetsView needs 'c' key handler in Update() method
- Main model already handles ViewTransitionMsg for state transitions
- FormModel.NewFormModel() already resets fields and sets default date
- Minimal changes needed: 3 files, ~20 lines of code
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Added 'c' key handler to asset list view to switch to form view for creating new DCA entries.

## What Changed

### Modified Files:
1. **internal/assets/view.go** - Added 'c' key handler in `Update()` method (3 lines)
   - Handles `tea.KeyRunes` for character 'c'
   - Returns `ViewTransitionMsg{View: "form"}` to trigger state transition

2. **internal/assets/view_test.go** - Added 3 new tests (69 lines)
   - `TestAssetsView_UpdateKeyC` - Verifies 'c' key returns ViewTransitionMsg
   - `TestAssetsView_UpdateKeyC_NavigatesToForm` - Verifies correct view target
   - `TestAssetsView_UpdateKeyC_IgnoresCapitalC` - Verifies uppercase 'C' is ignored
   - `TestAssetsView_UpdateKeyC_EmptyList` - Verifies 'c' works with empty asset list

### No Changes Needed:
- **cmd/dca/main.go** - Already handles `ViewTransitionMsg` for state transitions from assets view to form

## How It Works

1. User presses 'c' key in asset list view
2. `AssetsView.Update()` returns `ViewTransitionMsg{View: "form"}`
3. Main model switches to `StateForm` and creates new `FormModel`
4. Form initializes with reset fields and default timestamp (via `NewFormModel()`)

## Acceptance Criteria Status

- [x] Pressing 'c' in asset list switches to form view
- [x] Form fields are reset (amount=empty, asset=empty, price=empty) 
- [x] Date defaults to current timestamp in RFC3339 format
- [x] User can navigate form with Tab/Enter as before

## Tests
```
ok      github.com/danilo/scripts/github/dca/internal/assets
```

## Build
```bash
go build -o dca ./cmd/dca
```
Build successful, no warnings.

## Risks
None - minimal change, reuses existing transition mechanism.
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
<!-- DOD:END -->
