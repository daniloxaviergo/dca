---
id: GOT-024
title: 'Task 3: Return to asset list after entry'
status: Done
assignee:
  - Thomas
created_date: '2026-03-17 17:38'
updated_date: '2026-03-17 19:45'
labels: []
dependencies:
  - GOT-023
references:
  - cmd/dca/main.go
  - internal/form/model.go
documentation:
  - doc-005
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
After form submission, return to asset list view with updated data. Modify cmd/dca/main.go Update() to handle formSubmittedMsg by reloading asset data and switching to StateAssetsView.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 After form submission, app switches to asset list view
- [x] #2 Asset data refreshes to include new entry
- [x] #3 Aggregation calculations update correctly
- [x] #4 User can navigate asset list or create another entry
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The current implementation already has the foundation for view transitions in place:
- `formSubmittedMsg` is defined in `internal/form/model.go` and sent on form submission
- `StateAssetsView` and `StateForm` enum states are defined in `cmd/dca/main.go`
- `ViewTransitionMsg` exists in `internal/assets/view.go` for exit transitions

However, the current `main.go` has issues that prevent proper state transitions:

1. **Form submission flow is incomplete**: The `Update()` method receives `formSubmittedMsg` and changes state to `StateAssetsView`, but the `form.FormModel` does not persist the entry data back to the model's `entries` field before switching. The form's `saveEntry()` method modifies `m.Entries` in-place, which should work.

2. **AssetsView doesn't reload data**: When switching from form to assets view, the `AssetsView` is recreated but loaded with stale data (only loaded once at app startup in `main()`). The `LoadAndAggregateEntries()` call happens only during initialization.

3. **Missing data refresh on return**: When returning to asset list after form submission, we need to reload and re-aggregate entries from the JSON file to show the newly added entry.

**Approach**:
- Modify `cmd/dca/main.go` to reload and aggregate entry data when switching to `StateAssetsView` after form submission
- Keep the existing `dca.DCAData` reference in the model and pass it to the form
- After form submission, call `assets.LoadAndAggregateEntries()` to refresh the aggregated view data
- Update the form to not reload entries from file (use the in-memory reference)

**Key decision**: The form already modifies `m.Entries` in-place via `saveEntry()`, so the model's `entries` reference remains valid. We just need to reload the aggregated view data when returning to asset list.

### 2. Files to Modify

| File | Change Type | Description |
|------|-------------|-------------|
| `cmd/dca/main.go` | Modified | Update `Update()` to reload asset data when receiving `formSubmittedMsg` |
| `internal/form/model.go` | Modified | Exported `FormSubmittedMsg` type (was private `formSubmittedMsg`) |
| `internal/assets/view.go` | No change | `AssetsView` already has `Entries` field that can be updated |
| `internal/assets/aggregate.go` | No change | `LoadAndAggregateEntries()` already exists and works correctly |

### 2.5 Code Cleanup

The following duplicate code was identified and removed:
- `cmd/dca/dca_form.go` - Redundant form model implementation
- `cmd/dca/dca_form_test.go` - Redundant tests
- Duplicate `formSubmittedMsg` declaration in `cmd/dca/main.go`

The proper implementation exists in `internal/form/model.go`.

### 3. Dependencies

- **GOT-023 (Task 2)**: Must be complete - the 'c' key handler and form view are prerequisites
- **GOT-022 (Task 1)**: App must start in asset list view - establishes the state management pattern
- Existing code patterns:
  - `formSubmittedMsg` type already defined
  - `ViewTransitionMsg` pattern already exists
  - `LoadAndAggregateEntries()` already implemented

### 4. Code Patterns

Follow existing patterns in the codebase:

1. **State transitions** (from `cmd/dca/main.go`):
   - Use `tea.Cmd` to return custom messages from form
   - Pattern: `return m, func() tea.Msg { return formSubmittedMsg{} }`

2. **Data loading pattern** (from `main()` startup):
   ```go
   vm, err := assets.LoadAndAggregateEntries(defaultEntriesPath)
   if err != nil {
       m.assetsView.Error = err
   } else {
       m.assetsView.Entries = vm.Entries
       m.assetsView.Loaded = true
   }
   ```

3. **Form model pattern** (from `internal/form/model.go`):
   - Form modifies `m.Entries` in-place via `saveEntry()`
   - Entry is saved to file atomically within `saveEntry()`

4. **Error handling**:
   - Display errors in assets view via `m.assetsView.Error`
   - Use `fmt.Errorf("failed to ...: %w", err)` for wrapping

### 5. Testing Strategy

Add tests in `cmd/dca/main_test.go` (to be created):

1. **TestFormSubmittedMsg_TransitionsToAssetsView**
   - Simulate form submission
   - Verify state changes to `StateAssetsView`
   - Verify assets view is populated

2. **TestFormSubmittedMsg_ReloadsData**
   - Add an entry via form
   - Submit form
   - Verify assets view shows updated aggregated data
   - Verify new entry appears in list

3. **TestFormSubmittedMsg_AggregationCorrect**
   - Verify weighted average price calculation updates
   - Verify count increments correctly
   - Verify total shares updates correctly

4. **Integration test pattern** (from `internal/form/model_test.go`):
   - Use temp files for entry storage
   - Verify file is updated after submission
   - Verify data persists correctly

### 6. Risks and Considerations

**Risk 1: Data synchronization**
- The form stores entries in `m.Entries` which is a pointer to the model's `entries`
- After form submission, we reload from file via `LoadAndAggregateEntries()`
- **Mitigation**: The file was just written atomically by `saveEntry()`, so the reload will reflect the new entry

**Risk 2: Double-loading**
- Currently `main()` loads data once at startup for the initial view
- After form submission, we load again
- **Mitigation**: This is acceptable because the file has new data; the initial load is for the empty/initial state

**Risk 3: View transition message conflict**
- `ViewTransitionMsg` is used for exits (Esc/Ctrl+C)
- `formSubmittedMsg` is used for successful submissions
- **Mitigation**: These are distinct message types with distinct handlers; no conflict expected

**Trade-off: Memory vs. File reload**
- Option A: Keep `entries` in memory and pass it to `LoadAndAggregateEntries()` (but this function loads from file)
- Option B: Reload from file after form saves to file (current approach)
- **Chosen**: Option B because it ensures consistency with the persisted state and handles edge cases (e.g., if file write fails)

**Edge case: Form cancellation**
- If user cancels form (Esc/Ctrl+C), no data is saved
- No need to reload assets view
- **Handled**: The form sends `tea.Quit`, which exits without triggering state transition
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
<!-- SECTION:IMPLEMENTATION:BEGIN -->
2026-03-17: Implementation completed with fixes to duplicate type declarations.

**Changes Made:**
1. Renamed `formSubmittedMsg` to `FormSubmittedMsg` (exported) in `internal/form/model.go` (line 431) to enable cross-package usage
2. Removed duplicate files `cmd/dca/dca_form.go` and `cmd/dca/dca_form_test.go` which contained redundant form implementations
3. Updated all references to use `form.FormSubmittedMsg` in `cmd/dca/main.go` and `cmd/dca/main_test.go`
4. Modified `internal/form/model.go` to send `FormSubmittedMsg{}` instead of `formSubmittedMsg{}`

**Verification:**
- All tests pass: `go test ./...` ✓
- Build succeeds with no warnings ✓
- Code formatted: `go fmt ./...` ✓
- Binary created: `./dca` (5MB) ✓

**State Transition Flow:**
When `form.FormSubmittedMsg` is received in `Update()`:
1. State switches to `StateAssetsView`
2. New `AssetsView` created with `assets.NewAssetsView()`
3. Entries reloaded from file via `assets.LoadAndAggregateEntries()`
4. View populated with aggregated data including new entry
<!-- SECTION:IMPLEMENTATION:END -->
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
<!-- DOD:END -->
