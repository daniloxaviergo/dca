---
id: GOT-039
title: Bug esc on form should go back to list
status: In Progress
assignee:
  - Thomas
created_date: '2026-03-18 14:59'
updated_date: '2026-03-18 15:02'
labels: []
dependencies: []
ordinal: 1000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
When i open 'c' the form if cancel the application exit
I need go back to list assest
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 Unit tests pass (go test)
- [ ] #2 No new compiler warnings
- [ ] #3 Code follows project style (go fmt)
- [ ] #4 PRD referenced in task
- [ ] #5 Documentation updated (comments)
<!-- DOD:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 When user presses ESC on the form, the application should return to the asset list view
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Create a new `FormCancelledMsg` type and modify the form's ESC key handler to return to the assets view instead of quitting the app. The flow will be:

1. User presses ESC on form
2. Form sends `FormCancelledMsg` (instead of `tea.Quit`)
3. Main model receives message and switches `currentState` back to `StateAssetsView`
4. User returns to asset list view

This follows the existing pattern used for `FormSubmittedMsg` and view transitions.

### 2. Files to Modify

**`internal/form/model.go`**
- Add `FormCancelledMsg` struct type
- Modify ESC key handler to send `FormCancelledMsg` instead of `tea.Quit`
- Update footer help text to indicate ESC returns to list

**`cmd/dca/main.go`**
- Add case in model Update to handle `FormCancelledMsg`
- Switch `currentState` back to `StateAssetsView` when form is cancelled
- Reinitialize assets view to refresh display

### 3. Dependencies

No prerequisites. This is a self-contained fix to the form cancellation flow.

### 4. Code Patterns

- Use Bubble Tea message types for view transitions (like `FormSubmittedMsg`)
- Preserve existing error handling patterns
- Follow the state transition pattern already in place
- Keep changes minimal and focused on the specific issue

### 5. Testing Strategy

Add tests in `cmd/dca/main_test.go`:
- `TestFormCancelledMsg_ReturnsToAssetsView` - Verify ESC switches back to assets view
- `TestFormCancelledMsg_DoesNotSaveData` - Verify cancelled form doesn't save entries
- `TestCtrlCStillQuits` - Verify Ctrl+C still exits the app (regression test)

Run with: `go test -v ./...`

### 6. Risks and Considerations

- **Low risk**: This is a bug fix with clear expected behavior
- **No data loss**: Cancel should not save anything (already the case)
- **User expectation**: ESC should behave like "back" not "exit"
- **No breaking changes**: Only fixes incorrect behavior
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
# Implementation Plan for GOT-039

## Problem Statement
When the user presses ESC on the form, the application exits completely instead of returning to the asset list view. The user expects ESC to cancel the form and return to the previous view (assets list).

## Root Cause Analysis
In `internal/form/model.go`, the ESC key handler currently sends `tea.Quit` which terminates the entire Bubble Tea program:

```go
case tea.KeyCtrlC, tea.KeyEsc:
    // Cancel without saving
    return m, tea.Quit
```

This is incorrect because:
1. The form is a modal/view within the app, not the entire app
2. ESC should cancel the form and return to the assets view
3. Only Ctrl+C should quit the entire application

## Technical Approach

### 1. Create a New Message Type
Add a `FormCancelledMsg` in `internal/form/model.go` to signal form cancellation without quitting the app.

### 2. Modify ESC Key Handler
Change the ESC handler to return `FormCancelledMsg` instead of `tea.Quit`:

```go
case tea.KeyEsc:
    // Cancel form and return to assets view
    return m, func() tea.Msg {
        return FormCancelledMsg{}
    }
```

### 3. Handle FormCancellation in Main Model
Update `cmd/dca/main.go` to handle `FormCancelledMsg` by switching back to assets view.

### 4. Keep Ctrl+C for Full Exit
保留 Ctrl+C handler for actual application quit.

## Files to Modify

### `internal/form/model.go`
- Add `FormCancelledMsg` type
- Modify ESC key handler to send `FormCancelledMsg` instead of `tea.Quit`
- Update form footer to indicate ESC returns to list (not exit)

### `cmd/dca/main.go`
- Add case to handle `FormCancelledMsg` in model Update
- When form is cancelled, switch `currentState` back to `StateAssetsView`
- Reinitialize assets view if needed

## Dependencies
- None. This is a localized change within the form component and main state machine.

## Code Patterns to Follow
- Use Bubble Tea message patterns for view transitions
- Follow existing pattern: `FormSubmittedMsg` for successful submission
- Keep error handling minimal (user cancelling is not an error)
- Preserve current field values on cancel (don't save)

## Testing Strategy
Add tests in `cmd/dca/main_test.go`:
1. `TestFormCancelledMsg_ReturnsToAssetsView` - Verify ESC switches back to assets view
2. `TestFormCancelledMsg_DoesNotSaveData` - Verify cancelled form doesn't save entries
3. `TestCtrlCStillQuits` - Verify Ctrl+C still exits the app

## Risks and Considerations
- **No data loss risk**: Cancelling form should not save anything
- **Form state**: Current form state should be discarded on cancel
- **Backward compatibility**: This is a bug fix, no breaking changes
- **User experience**: ESC should feel like "back" not "exit"
<!-- SECTION:NOTES:END -->
