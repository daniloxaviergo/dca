---
id: GOT-012
title: 'Task 3: Integrate Assets View into Main Application'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 00:42'
updated_date: '2026-03-17 08:27'
labels: []
dependencies: []
references:
  - 'PRD: DCA Assets List Table View'
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Integrate assets view into main.go with view state management and keyboard navigation between views
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
### 1. Technical Approach

Integrate the existing AssetsView Bubble Tea component into the main application flow by implementing state management and view transitions.

**Architecture:**
- Add `AppState` enum to track which view is active (Form vs AssetsView)
- Modify `model` struct to include `assetsView` and `currentState` fields
- Update `Update()` method to delegate to current view and handle state transitions
- Update `View()` method to render the current view
- Handle data persistence: entries are shared between views, changes in form save to JSON

**Key Design Decisions:**
1. **Shared Data**: `DCAData` is passed to both views; form writes, view reads
2. **Transition Triggers**: Form submission (Enter + Confirm y) switches to AssetsView; AssetsView exit (Esc/Ctrl+C) switches back
3. **Data Loading**: AssetsView loads aggregated data from `LoadAndAggregateEntries()` on transition
4. **No Persistent Selection**: Selected row index resets when switching back to AssetsView

**Why This Approach:**
- Minimal changes to existing code (no restructuring needed)
- Reuses existing `AssetsView` component (already implemented in Task 2)
- Maintains data consistency through shared `DCAData` reference
- Follows existing Bubble Tea pattern from `FormModel`

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `main.go` | Modify | Add AppState enum, update model struct, implement Update/View state delegation |
| `main_test.go` | Add tests | Integration tests for state transitions, data flow |
| `assets_view.go` | No change | Component already exists from Task 2 |
| `dca_entry.go` | No change | Data model and I/O already exist |
| `dca_form.go` | No change | Form component unchanged |

**New Files:** None (all functionality already implemented in main.go)

### 3. Dependencies

**Prerequisites (must be complete first):**
- ✅ Task GOT-010: Assets View Model and Data Aggregation (assets_view.go with LoadAndAggregateEntries)
- ✅ Task GOT-011: Interactive Table UI Component (AssetsView Bubble Tea component)
- ✅ Existing Bubble Tea v1.3.10 and Lipgloss v1.1.0

**No external dependencies required.**

**Blocking issues:** None - existing assets_view.go and dca_form.go are stable.

**Setup steps:** None required - integration uses existing components.

### 4. Code Patterns

**Follow existing patterns from dca_form.go and main.go:**

| Pattern | Implementation |
|---------|----------------|
| **State enum** | `type AppState int` with iota (StateForm, StateAssetsView) |
| **Model struct** | Include all view models as pointers |
| **Update delegation** | Switch on currentState, call current view's Update() |
| **View delegation** | Switch on currentState, call current view's View() |
| **State transition** | Update currentState, create new view instance |
| **Quit handling** | Detect tea.QuitMsg to trigger exit transition |

**Specific conventions:**
- Use type assertions: `newForm.(*FormModel)` after Update()
- Update receiver fields in-place for Bubble Tea model pattern
- Pass `*DCAData` reference to share data between views
- Use `tea.Quit` command to signal exit from AssetsView

**Style requirements:**
- Match existing `dca_form.go` key handling (Esc/Ctrl+C for quit)
- Use lipgloss styling from AssetsView (already implemented)
- Maintain 8-decimal precision for financial values

### 5. Testing Strategy

**Integration tests in main_test.go:**

| Test Case | Purpose |
|-----------|---------|
| `TestAppState_InitState` | Application starts in Form state |
| `TestModel_Update_FormToAssets` | Form submission transitions to AssetsView |
| `TestModel_Update_AssetsToForm` | AssetsView exit transitions back to Form |
| `TestModel_Update_QuitFromAssets` | Ctrl+C in AssetsView exits cleanly |
| `TestModel_Update_KeyNavigation_Form` | Form keyboard navigation works |
| `TestModel_Update_KeyNavigation_Assets` | AssetsView navigation works |
| `TestModel_View_FormRenders` | Form view renders without panic |
| `TestModel_View_AssetsRenders` | AssetsView renders without panic |
| `TestModel_DataConsistency` | Entries modified in form visible in AssetsView |
| `TestMain_EntryRoundtrip` | Full flow: add entry → save → view → exit |

**Test approach:**
- Unit test each state transition with mock keyboard inputs
- Verify `currentState` changes correctly after each message
- Test data consistency: add entry in form, verify it appears in AssetsView
- Test exit behavior: AssetsView quit command returns to Form state
- Use temp file I/O tests for persistence verification

**Edge cases:**
- Empty entries file → AssetsView shows empty state
- Invalid JSON → Form validation error (existing behavior)
- Nil AssetsView → Graceful fallback in Update/View
- Multiple quick transitions → State is replaced correctly

### 6. Risks and Considerations

**Potential issues:**

| Risk | Mitigation |
|------|------------|
| State not updating correctly | Type assertions must match actual types; test thoroughly |
| Data inconsistency | Pass `*DCAData` reference, not copy; verify in tests |
| Memory leaks from old views | Old view discarded after state change; Go GC handles it |
| Race conditions | Single-threaded Bubble Tea event loop; no concurrent access |
| Nil pointer dereference | Check `m.form != nil` and `m.assetsView != nil` in Update/View |

**Design trade-offs:**
- **No "back" from AssetsView to Form**: Transition is one-way (form → view); must exit to return (simpler than full navigation stack)
- **No persistence in AssetsView**: View is read-only; changes only in form (clear separation of concerns)
- **Fresh load on each view switch**: AssetsView reloads data on transition (ensures consistency)

**Implementation checkpoints:**
1. First: Add AppState enum and update model struct with minimal compilation
2. Second: Implement Update() state delegation (form transition works)
3. Third: Implement View() delegation (assets view renders)
4. Fourth: Add exit transition (AssetsView → Form)
5. Fifth: Write integration tests for all state transitions
6. Sixth: Run full test suite and verify no regressions

**Blocking issues:** None identified

**DoD verification:**
- [ ] All acceptance criteria met (4 criteria in task)
- [ ] `go test ./...` passes
- [ ] `go build` completes without warnings
- [ ] `go fmt` applied (no changes)
- [ ] PRD referenced in task
- [ ] Comments updated if code changes
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
