---
id: GOT-012
title: 'Task 3: Integrate Assets View into Main Application'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-17 00:42'
updated_date: '2026-03-17 10:26'
labels: []
dependencies: []
references:
  - 'PRD: DCA Assets List Table View'
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Integrate assets view into main.go with view state management and keyboard navigation between views, focus on unit test

The first view should be a list of assets
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Assets view accessible from main entry
- [ ] #2 Data consistency maintained across views
- [ ] #3 Changes reflected after save
- [ ] #4 Clean exit from assets view
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
# Implementation Plan: Integrate Assets View into Main Application

## 1. Technical Approach

The integration involves connecting the existing `AssetsView` Bubble Tea component with the main application's state management system. The current `main.go` already has a basic two-view architecture with `StateForm` and `StateAssetsView`, but requires refinement to properly handle:

1. **Initial State**: Application starts in form view by default
2. **Form Submission Flow**: User enters data via form → entry saved → switch to assets view
3. **Assets View Navigation**: User can navigate the assets table and exit back to form
4. **Data Consistency**: Assets view reflects latest saved entries from `dca_entries.json`

### Architecture Decision
- Keep the current `AppState` enum with two states
- Use `tea.Quit` from AssetsView to return to form (not full app exit)
- Reuse the same `DCAData` instance across views for data consistency
- Load asset summaries fresh from file on each transition to assets view

### Why This Approach
- Minimal state changes required to current codebase
- Follows existing Bubble Tea patterns in `dca_form.go`
- Maintains single responsibility: form view for input, assets view for display
- No new dependencies or major refactoring needed

## 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `main.go` | Modify | Fix state transition logic, data loading, and view switching |
| `assets_view.go` | No change | Component already implemented correctly |
| `dca_form.go` | No change | Form submission logic already correct |
| `dca_entry.go` | No change | Data model and I/O already correct |

## 3. Dependencies

- ✅ `AssetsViewModel` and `AssetsView` already implemented in `assets_view.go`
- ✅ `LoadAndAggregateEntries()` function ready for use
- ✅ `DCAData` structure shared across all components
- ✅ Existing `dca_entries.json` file path constant available

### Prerequisites for Integration
1. `main.go` must maintain reference to `entries *DCAData` instance
2. Assets view must reload data from file on each entry (not use cached data)
3. Proper cleanup on form exit to avoid data loss

## 4. Code Patterns

### Follow These Existing Patterns
1. **Bubble Tea Model Pattern**: Same structure as `FormModel`
   - `Init() tea.Cmd` - return nil for static views
   - `Update(msg tea.Msg) (tea.Model, tea.Cmd)` - handle key presses
   - `View() string` - render UI with Lipgloss

2. **State Management**: 
   ```go
   type AppState int
   const ( StateForm AppState = iota; StateAssetsView )
   ```

3. **Error Handling**: Use Lipgloss colored text for errors (foreground 196)

4. **Navigation Patterns**:
   - `tea.KeyEsc` or `tea.KeyCtrlC` to exit view
   - `tea.KeyUp`/`tea.KeyDown` for row selection

5. **Data Consistency**: 
   - Reload from file on transition
   - Use same `DCAData` instance passed between views

### Implementation Steps
1. Fix `main.go` Update loop to properly handle state transitions
2. Ensure assets view reloads fresh data from file on each entry
3. Verify form data is saved before switching to assets view
4. Test clean transitions in both directions

## 5. Testing Strategy

### Unit Tests to Verify
1. **State Transitions**:
   - Form → Assets view on submission
   - Assets → Form on Esc/Ctrl+C
   - Data consistency maintained

2. **Data Loading**:
   - `LoadAndAggregateEntries()` with populated file
   - `LoadAndAggregateEntries()` with empty file
   - Error handling for missing file

3. **Assets View Rendering**:
   - Table displays correctly with entries
   - Empty state message shown when no assets
   - Navigation keys work correctly

### Test Commands
```bash
# Run all tests
go test ./...

# Run specific test file
go test -v -run TestAssetsView ./...

# Verify formatting
gofmt -e *.go

# Verify no vet issues
go vet ./...
```

### Edge Cases to Cover
- Empty `dca_entries.json` file
- Missing file (creates new)
- Multiple entries for same asset
- Single asset with multiple entries
- Rapid state transitions (form → assets → form)

## 6. Risks and Considerations

### Known Issues to Address
1. **Data Loading Timing**: Current code may load data before form submission completes
   - Fix: Ensure save completes before switching views

2. **State Persistence**: Need to verify `entries` pointer remains consistent
   - Both views must reference same `DCAData` instance

3. **Exit Behavior**: Currently uses `tea.Quit` which exits full application
   - Fix: Consider returning to form instead of quitting on Esc from assets view

### Potential Pitfalls
- **File I/O Race Conditions**: If form saves while assets view loads
  - Mitigation: Use atomic writes (already implemented in `SaveEntries`)

- **Memory Leaks**: Repeatedly creating new view instances
  - Mitigation: Reuse `AssetsView` instance, just reload data

- **UI Glitches**: If state transitions are not synchronized
  - Mitigation: Update model synchronously, no async operations

### Rollback Considerations
- Changes are isolated to `main.go`
- Can revert by replacing `main.go` with previous version
- No database or file format changes
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
