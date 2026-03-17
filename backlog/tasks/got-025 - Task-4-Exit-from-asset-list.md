---
id: GOT-025
title: 'Task 4: Exit from asset list'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 19:58'
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

**Problem Identified**: The current implementation has a bug where Esc/Ctrl+C from the asset list view transitions to the form view instead of exiting the application.

**Solution**: Modify the main model's state transition logic to properly handle exit messages from the asset list view.

**How it will be built**:
1. Update the `StateAssetsView` case in `main.go` model's `Update()` method to detect when an exit command originates from the asset list
2. Return `tea.Quit` directly instead of transitioning to form view
3. The `ViewTransitionMsg` type is currently ambiguous - it's used both for "switch to form" and "exit app"

**Architecture Decision**: Add a separate message type for exit commands or modify the existing logic to check the current state when handling `ViewTransitionMsg`.

**Why this approach**: Minimal changes to existing code, maintains consistency with current message-passing pattern, preserves existing behavior for 'c' key which correctly switches to form.

### 2. Files to Modify

- **cmd/dca/main.go** - Modify the `Update()` method's `StateAssetsView` case to properly handle exit commands
  - Change the `ViewTransitionMsg` handling to check if it's an exit command vs form-switch command
  - Return `tea.Quit` for exit commands from asset list

### 3. Dependencies

- No external dependencies required
- Depends on existing `ViewTransitionMsg` type in `internal/assets/view.go`
- Requires understanding of current state transition logic in `main.go`

### 4. Code Patterns

- Follow existing Bubble Tea message passing pattern
- Use existing `ViewTransitionMsg` type but make the view field more explicit or add exit-specific logic
- Maintain consistency with how form handles `tea.KeyCtrlC` and `tea.KeyEsc` (returns `tea.Quit`)
- Keep error handling minimal for exit paths

### 5. Testing Strategy

- **Unit Tests**: Add/modify tests in `cmd/dca/main_test.go` to verify:
  - Esc key from asset list returns `tea.Quit`
  - Ctrl+C from asset list returns `tea.Quit`
  - 'c' key from asset list still transitions to form view
  - No data loss occurs (entries file unchanged on exit)

- **Existing Tests to Verify**:
  - `internal/assets/view_test.go` already has tests for `UpdateEscape`, `UpdateCtrlC`, `UpdateQuitMsg`
  - These tests verify `ViewTransitionMsg` is returned, which is correct at the component level
  - Need integration tests at the main model level to verify actual quit behavior

- **Test Approach**: Use the same pattern as `main_test.go` - create model, send key messages, verify `tea.Quit` cmd is returned

### 6. Risks and Considerations

**Blocking Issue Identified**: The current code has a bug where Esc/Ctrl+C from asset list doesn't exit the app. This must be fixed.

**Potential Pitfalls**:
- The `ViewTransitionMsg{View: "form"}` is used by both 'c' key (correct) and exit handlers (incorrect)
- Need to distinguish between "user pressed 'c' to open form" vs "user pressed Esc to exit"

**Trade-offs**:
1. **Option A**: Add a new message type like `ExitRequestMsg` from assets view
2. **Option B**: Modify `ViewTransitionMsg` to have an `Action` field (e.g., "form" vs "exit")
3. **Option C**: Handle exit at the assets view level by returning `tea.Quit` directly

**Recommended Approach**: Option C - Have assets view return `tea.Quit` directly for exit commands, since that's the actual desired behavior. The 'c' key can still use `ViewTransitionMsg` for form transition.

**Deployment Considerations**: This is a bug fix, not a new feature. No special deployment steps required. Should be tested with actual app execution (`go run main.go`).
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
