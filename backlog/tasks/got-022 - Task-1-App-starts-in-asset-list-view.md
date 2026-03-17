---
id: GOT-022
title: 'Task 1: App starts in asset list view'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 17:41'
labels: []
dependencies: []
references:
  - cmd/dca/main.go
documentation:
  - doc-005
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Change app to initialize in asset list view instead of form view. Modify cmd/dca/main.go to set currentState to StateAssetsView and initialize assetsView with loaded data on startup.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 App initializes with StateAssetsView instead of StateForm
- [ ] #2 Asset list displays on first render
- [ ] #3 User can navigate with ↑/↓ arrow keys
- [ ] #4 Esc/Ctrl+C exits the app from asset list
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Change the application's initial state from `StateForm` to `StateAssetsView` so the app opens with the asset list displayed instead of a blank entry form.

**Key Changes:**
- Modify `main()` to initialize `currentState` to `StateAssetsView`
- Load and aggregate existing entries into `AssetsView` on startup
- `AssetsView` already has navigation (↑/↓) and exit (Esc/Ctrl+C) functionality implemented
- The existing `ViewTransitionMsg` system handles transitions between views

**Architecture Decision:**
- Reuse `AssetsView` component (already implemented in `internal/assets/view.go`)
- Reuse `ViewTransitionMsg` for navigation state changes
- Keep form as fallback - user can press 'c' (Task 2) to switch to form
- No new dependencies or complex state management needed

**Why This Approach:**
- Minimal code change (3 lines in main.go + initialization logic)
- Reuses existing tested components
- Maintains consistent state transition patterns
- Aligns with PRD goal: "view assets before entering data"

### 2. Files to Modify

**Primary Changes:**
| File | Action | Reason |
|------|--------|--------|
| `cmd/dca/main.go` | Modify | Change initial state to `StateAssetsView` and initialize `AssetsView` with loaded data |
| `cmd/dca/main.go` | Modify | Move data loading before model initialization |
| `cmd/dca/main.go` | Modify | Update `Init()` to handle `StateAssetsView` case |

**No Changes Required:**
- `internal/assets/view.go` - Already has navigation and exit handling
- `internal/assets/aggregate.go` - Already has data loading/aggregation
- `internal/form/model.go` - Not affected by this task

### 3. Dependencies

**Prerequisites:**
- ✅ `AssetsView` component implemented (`internal/assets/view.go`)
- ✅ Data aggregation logic (`internal/assets/aggregate.go`)
- ✅ `ViewTransitionMsg` message type for view switching
- ✅ Existing `dca.LoadEntries()` function for data loading

**Blocking Issues:**
- None - all dependencies are already in place

**Setup Steps:**
- No external setup required
- Existing `dca_entries.json` file will be loaded (or created empty if missing)

### 4. Code Patterns

**Conventions to Follow:**
- Match existing Bubble Tea pattern (model.Update, model.View)
- Use `tea.Cmd` for initialization (even if nil in this case)
- Preserve error handling pattern from `main()`
- Use lipgloss for styling (already in AssetsView)

**Integration Pattern:**
```go
// Current (form-first):
currentState: StateForm
form: form.NewFormModel(entries, defaultEntriesPath)

// New (assets-first):
currentState: StateAssetsView
assetsView: assets.NewAssetsView()
vm, err := assets.LoadAndAggregateEntries(defaultEntriesPath)
// Populate assetsView with vm.Entries and handle errors
```

### 5. Testing Strategy

**Unit Test Coverage:**
- Test `main()` initializes with `StateAssetsView`
- Test `AssetsView.Update()` handles ↑/↓ arrow keys
- Test `AssetsView.Update()` handles Esc/Ctrl+C (should trigger `tea.Quit`)
- Test empty state display when no entries exist
- Test data loading with existing entries

**Test Files to Check/Update:**
- `cmd/dca/main_test.go` (if exists) - verify initial state
- `internal/assets/view_test.go` - verify navigation and rendering
- `internal/assets/aggregate_test.go` - verify data loading

**Manual Verification:**
1. Run `go run cmd/dca/main.go` - should see asset list immediately
2. Press ↑/↓ - should navigate through rows (or show empty state)
3. Press Esc or Ctrl+C - should exit application
4. If entries exist, verify data displays correctly in table format

### 6. Risks and Considerations

**Known Risks:**
- **Risk**: User expects to enter data immediately on first run
  - **Mitigation**: This is the intended behavior per PRD - users want to view assets first
  
- **Risk**: Asset list shows error if `dca_entries.json` is corrupted
  - **Mitigation**: `LoadEntries()` handles this gracefully, error displayed in view

**Trade-offs:**
- **Trade-off**: Form initialization is deferred until user presses 'c' (Task 2)
  - **Reason**: Reduces initial complexity; form only initialized when needed

- **Trade-off**: Empty asset list shows "No assets yet" message
  - **Reason**: Clear user feedback vs blank screen

**Implementation Notes:**
- `AssetsView.Loaded` field controls whether "Loading data..." or actual data displays
- `AssetsView.Entries` should be populated from `AssetsViewModel.Entries`
- Error state should be displayed if loading fails
- No initialization command (`Init()` returns nil) as AssetsView doesn't need async setup
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
