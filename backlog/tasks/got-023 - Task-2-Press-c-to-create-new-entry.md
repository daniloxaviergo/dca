---
id: GOT-023
title: 'Task 2: Press ''c'' to create new entry'
status: To Do
assignee:
  - Thomas
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 18:01'
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
- [ ] #1 #1 Pressing 'c' in asset list switches to form view
- [ ] #2 #2 Form fields are reset (amount=empty, asset=empty, price=empty)
- [ ] #3 #3 Date defaults to current timestamp in RFC3339 format
- [ ] #4 #4 User can navigate form with Tab/Enter as before
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

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 All acceptance criteria met
- [ ] #8 Unit tests pass (go test)
- [ ] #9 No new compiler warnings
- [ ] #10 Code follows project style (go fmt)
- [ ] #11 PRD referenced in task
- [ ] #12 Documentation updated (comments)
- [ ] #13 All acceptance criteria met
- [ ] #14 Unit tests pass (go test)
- [ ] #15 No new compiler warnings
- [ ] #16 Code follows project style (go fmt)
- [ ] #17 PRD referenced in task
- [ ] #18 Documentation updated (comments)
<!-- DOD:END -->
