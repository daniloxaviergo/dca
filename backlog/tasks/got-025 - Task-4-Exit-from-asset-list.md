---
id: GOT-025
title: 'Task 4: Exit from asset list'
status: In Progress
assignee: []
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 20:01'
labels: []
dependencies: []
references:
  - internal/assets/view.go
documentation:
  - doc-005
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Users can exit the application from the asset list view. Verify existing Esc/Ctrl+C handlers work from asset list view.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Esc key exits application
- [ ] #2 Ctrl+C exits application
- [ ] #3 No unsaved data loss (entries saved on form submit only)
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

**Problem Identified**: The `ViewTransitionMsg` is used ambiguously:
- 'c' key in asset list should switch to form view ✓ (correct)
- Esc/Ctrl+C in asset list should exit app ✗ (currently switches to form - **bug**)

**Root Cause**: In `cmd/dca/main.go:80-85`, the `StateAssetsView` case treats ALL `ViewTransitionMsg` as "switch to form", but the asset view sends this same message for both 'c' key AND exit commands.

**Solution**: Have the assets view return `tea.Quit` directly for exit commands (Esc/Ctrl+C/QuitMsg), while still using `ViewTransitionMsg` only for the 'c' key form-switch action.

**How it will be built**:
1. Modify `internal/assets/view.go:Update()` to return `tea.Quit` for `tea.KeyCtrlC`, `tea.KeyEsc`, and `tea.QuitMsg`
2. Modify `internal/assets/view.go:Update()` to return `ViewTransitionMsg{View: "form"}` ONLY for 'c' key
3. The main model's `StateAssetsView` case can then be simplified or removed since exit is handled at component level

**Why this approach**: 
- Clean separation of concerns - assets view knows when to exit vs switch views
- Minimal changes - only modify the assets view component
- Consistent with how form handles exit (returns `tea.Quit` directly)

### 2. Files to Modify

- **internal/assets/view.go** - Modify `Update()` method to differentiate exit vs form-switch actions
  - Lines 27-43: Update message handling to return appropriate commands

- **cmd/dca/main.go** - Simplify `StateAssetsView` case
  - Lines 74-85: Remove the `ViewTransitionMsg` handling or keep for 'c' key support

### 3. Dependencies

- No external dependencies required
- Depends on understanding of current `ViewTransitionMsg` usage
- Requires understanding of Bubble Tea message passing patterns

### 4. Code Patterns

- Follow existing Bubble Tea message passing pattern
- Exit commands should return `tea.Quit` directly (as form does in `internal/form/model.go:37-39`)
- View transition messages use `ViewTransitionMsg` for navigation between views
- Keep test patterns consistent with existing tests in `internal/assets/view_test.go`

### 5. Testing Strategy

- **Unit Tests**: Verify in `internal/assets/view_test.go`:
  - `TestAssetsView_UpdateEscape` - should return `tea.Quit`
  - `TestAssetsView_UpdateCtrlC` - should return `tea.Quit`
  - `TestAssetsView_UpdateQuitMsg` - should return `tea.Quit`
  - `TestAssetsView_UpdateKeyC` - should return `ViewTransitionMsg{View: "form"}`

- **Integration Tests**: Verify in `cmd/dca/main_test.go`:
  - Esc from asset list state returns `tea.Quit`
  - Ctrl+C from asset list state returns `tea.Quit`
  - 'c' from asset list state transitions to form state

- **Test Approach**: 
  - Use existing test patterns from both test files
  - Verify actual cmd return values match expected types

### 6. Risks and Considerations

**Blocking Issue**: The current bug means users cannot exit from the asset list view without losing state.

**Potential Pitfalls**:
- Need to ensure 'c' key still works correctly after the fix
- Existing tests in `view_test.go` verify `ViewTransitionMsg` but may need updates to verify `tea.Quit` instead

**Trade-offs Considered**:
1. **Option A**: Add new message type for exit - more verbose, requires more code changes
2. **Option B**: Add action field to `ViewTransitionMsg` - adds complexity to existing pattern
3. **Option C**: Return `tea.Quit` directly in assets view for exit - cleanest, minimal changes

**Recommended**: Option C - simplest and most consistent with form component behavior

**Deployment**: Bug fix only, no special deployment steps. Verify with `go run main.go` and manual testing of all exit paths.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Research Findings

After thoroughly reviewing the codebase, I found that the exit functionality from the asset list view was **already implemented** in previous tasks (GOT-022, GOT-023, GOT-024). The implementation includes:

1. **Esc key handling** - `internal/assets/view.go:29` handles `tea.KeyEsc` in Update()
2. **Ctrl+C handling** - `internal/assets/view.go:28` handles `tea.KeyCtrlC` in Update()
3. **QuitMsg handling** - `internal/assets/view.go:41` handles `tea.QuitMsg`

All three handlers return a `ViewTransitionMsg{View: "form"}`, which the main model then converts to `tea.Quit` in the assets view state.

## Key Code Paths

### Assets View Update() Handler (internal/assets/view.go:27-43)
```go
case tea.KeyCtrlC, tea.KeyEsc:
    return a, func() tea.Msg {
        return ViewTransitionMsg{View: "form"}
    }
...
case tea.QuitMsg:
    return a, func() tea.Msg {
        return ViewTransitionMsg{View: "form"}
    }
```

### Main Model State Transition (cmd/dca/main.go:76-82)
```go
case StateAssetsView:
    if _, ok := msg.(assets.ViewTransitionMsg); ok {
        m.currentState = StateForm
        m.form = form.NewFormModel(m.entries, defaultEntriesPath)
        return m, nil
    }
```

Wait - there's a **BUG**. When asset list receives exit command, it returns `ViewTransitionMsg{View: "form"}`, but the main model interprets this as a signal to switch **back to form**, not to quit the app.

The correct behavior should be: **Esc/Ctrl+C from asset list should exit the app entirely**, not switch to form view.

Looking at the tests (`view_test.go:157-177`), they verify that `ViewTransitionMsg` is returned, but the main model doesn't actually call `tea.Quit` when this message comes from the asset list.

**This is the bug that needs fixing for GOT-025.**
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
