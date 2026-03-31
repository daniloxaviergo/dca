---
id: GOT-066
title: '[doc-019 Phase 5] Verify keyboard navigation and no breaking changes'
status: To Do
assignee:
  - workflow
created_date: '2026-03-29 12:32'
updated_date: '2026-03-31 14:40'
labels:
  - task
  - testing
  - code-quality
dependencies: []
documentation:
  - doc-019
priority: medium
ordinal: 10000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Verify that all existing keyboard navigation functionality remains unchanged in internal/assets/view.go after table layout modifications. Test that ↑/↓ navigation maintains wrap-around behavior (header to last row,反之亦然), Enter key still opens asset history modal, Esc and Ctrl+C still exit application, and 'c' key still switches to form view. Run full test suite to ensure no regressions. Test with various data volumes (0, 5, 29, 30 entries) to verify layout consistency and that all acceptance criteria are met.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Keyboard navigation (↑/↓/Enter/Esc/Ctrl+c) works identically to before
- [ ] #2 Wrap-around behavior preserved for navigation
- [ ] #3 Asset history modal opens correctly on Enter
- [ ] #4 'c' key still switches to form view
- [ ] #5 Tested with data volumes: 0, 5, 29, 30 entries
- [ ] #6 Full test suite passes with make test
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task is a **quality assurance verification task** (not implementation) to ensure the doc-019 PRD Phase 4 changes (table layout improvements) did not introduce any breaking changes to keyboard navigation and UI behavior.

**Approach**:
1. **Review existing navigation code** - Verify `handleUp()`, `handleDown()`, and `Update()` methods for keyboard handling
2. **Test wrap-around behavior** - Confirm header-to-last-row and last-row-to-first-data-row wrapping
3. **Verify modal integration** - Ensure Enter key still opens modal, Esc/Ctrl+C exit correctly
4. **Test form transition** - Confirm 'c' key still triggers `ViewTransitionMsg{View: "form"}`
5. **Run existing test suite** - Execute all tests with `make test` to verify no regressions
6. **Test data volume scenarios** - Verify behavior with 0, 5, 29, and 30 entries

**Files to Review**:
- `internal/assets/view.go` (lines 68-168: Update, handleUp, handleDown)
- `internal/assets/view_test.go` (existing navigation tests)
- `cmd/dca/main.go` (lines 55-67, 90-96: view transition handling)

### 2. Files to Modify

**No code changes required** - This is a verification/testing task.

**Files to Review**:
- `internal/assets/view.go` - Keyboard navigation logic
- `internal/assets/view_test.go` - Existing test coverage
- `cmd/dca/main.go` - View transition handling

**Tests to Run**:
- All existing unit tests in `internal/assets/`
- Verify coverage includes navigation edge cases

### 3. Dependencies

This task has **no prerequisites** as doc-019 Phases 1-4 are already completed (tasks GOT-062, GOT-063, GOT-064, GOT-065 marked as Done).

**Expected state**:
- Table layout rendered with 86-character width (double-line borders)
- All navigation code in place and functional
- Test suite covering existing behavior

### 4. Code Patterns

**Navigate to verify existing patterns**:

1. **Keyboard Handling** (`internal/assets/view.go`, lines 71-107):
```go
case tea.KeyUp:
    return a.handleUp()
case tea.KeyDown:
    return a.handleDown()
case tea.KeyRunes:
    if string(msg.Runes) == "c" {
        return a, func() tea.Msg {
            return ViewTransitionMsg{View: "form"}
        }
    }
case tea.KeyEsc:
    // Quit if modal not visible
case tea.KeyCtrlC:
    return a, tea.Quit
```

2. **Wrap-around Logic** (`internal/assets/view.go`, lines 110-135):
```go
// Header (index 0) to last row (index 29) on Up
// Last row (index 29) to first data (index 1) on Down
const maxRowIndex = 29 // 30 total rows - 1
```

3. **View Transition** (`cmd/dca/main.go`, lines 90-96):
```go
// On ViewTransitionMsg with View="form", switch to StateForm
if transitionMsg, ok := msg.(assets.ViewTransitionMsg); ok && transitionMsg.View == "form" {
    m.currentState = StateForm
    m.form = form.NewFormModel(m.entries, defaultEntriesPath)
    return m, nil
}
```

4. **Modal Opening** (`internal/assets/view.go`, lines 144-172):
```go
// Enter opens modal when not in modal view (selectedIndex > 0)
// Enter loads more when in modal view (a.Modal.Visible = true)
```

**Verification checklist**:
- [ ] All key handlers return `tea.Cmd` correctly
- [ ] Wrap-around uses correct constants (maxRowIndex = 29)
- [ ] Modal state checks before Enter handling
- [ ] View transition message contains `View: "form"`
- [ ] Esc/Ctrl+C return `tea.Quit` cmd
- [ ] Form cancellation returns to assets view

### 5. Testing Strategy

**Test Categories**:

1. **Keyboard Navigation Tests** (existing, run with `make test`):
   - `TestAssetsView_NavigateUp` - Verify up navigation with wrap to index 29
   - `TestAssetsView_NavigateDown` - Verify down navigation with wrap to index 1
   - `TestAssetsView_NavigateWrapUp` - Verify header-to-last-row wrap
   - `TestAssetsView_NavigateWrapDown` - Verify last-row-to-first-data-row wrap
   - `TestAssetsView_UpdateArrowUp/Down` - Integration tests for key messages
   - `TestAssetsView_UpdateEscape` - Esc key exits with `tea.Quit`
   - `TestAssetsView_UpdateCtrlC` - Ctrl+C exits with `tea.Quit`
   - `TestAssetsView_UpdateKeyC` - 'c' key returns `ViewTransitionMsg{View: "form"}`
   - `TestAssetsView_UpdateKeyC_NavigatesToForm` - Verify form view switch

2. **Data Volume Tests** (existing):
   - `TestAssetsView_RenderWith5Assets` - 5 entries (5 data + 24 empty rows)
   - `TestAssetsView_RenderWith25Assets` - 25 entries (25 data + 4 empty rows)
   - `TestTableLayout_Exactly30Rows` - Tests 1, 5, 25, 29, 30 entries
   - `TestTableLayout_EmptyRowPadding` - 1, 5, 25, 29 entries with correct padding

3. **Modal Tests** (existing):
   - `TestAssetsView_UpdateLoadMore` - Modal + LoadMoreMsg integration
   - `TestAssetsView_UpdateLoadMore_EmptyModal` - Nil modal handling
   - `TestAssetsView_UpdateLoadMore_ModalNotVisible` - Not visible handling
   - Modal visibility state tests

**Coverage Requirements**:
- All keyboard handlers tested
- Wrap-around edge cases covered
- Modal state transitions verified
- Form view transition tested

**Verification Steps**:
```bash
# Run all tests
make test

# Check coverage
make test-cover

# Run specific navigation tests
go test -v ./internal/assets/... -run "Navigate"

# Run specific table layout tests
go test -v ./internal/assets/... -run "TableLayout"

# Run form transition tests
go test -v ./internal/assets/... -run "KeyC"
```

**Test Case Mapping**:
| Acceptance Criteria | Test(s) to Verify |
|-------------------|-------------------|
| #1 Keyboard navigation works identically | All existing navigation tests |
| #2 Wrap-around preserved | `TestAssetsView_NavigateWrapUp/Down` |
| #3 Modal opens on Enter | `TestAssetsView_UpdateLoadMore` |
| #4 'c' key switches to form | `TestAssetsView_UpdateKeyC` |
| #5 Tested with 0, 5, 29, 30 entries | `TestTableLayout_Exactly30Rows` |
| #6 Full test suite passes | `make test` exit code 0 |

### 6. Risks and Considerations

**No implementation risks** - This is verification only.

**Verification Risks**:

| Risk | Impact | Mitigation |
|------|--------|-----------|
| Tests pass but runtime behavior differs | Medium | Manual testing with `make run` verifies UI behavior |
| Test coverage gaps | Low | Review coverage report (`coverage.out`) for untested paths |
| PRD changes after implementation | Low | PRD doc-019 marked complete; no pending changes |

**Acceptance Criteria Verification**:

| AC | Verification Method | Status |
|----|--------|-----|
| #1 Keyboard navigation works identically | All existing tests pass |
| #2 Wrap-around preserved | `TestAssetsView_NavigateWrapUp/Down` pass |
| #3 Modal opens on Enter | `TestAssetsView_UpdateLoadMore` and modal tests pass |
| #4 'c' key switches to form | `TestAssetsView_UpdateKeyC` passes |
| #5 Tested with 0, 5, 29, 30 entries | `TestTableLayout_Exactly30Rows` covers all |
| #6 Full test suite passes | `make test` exits with code 0 |

**Definition of Done Verification**:

- [ ] #1 ACs met - All existing navigation tests pass
- [ ] #2 Unit tests pass - `go test ./...` succeeds
- [ ] #3 No compiler warnings - `go build` succeeds silently
- [ ] #4 Code follows style - `make fmt` no changes needed
- [ ] #5 PRD referenced - Task references doc-019
- [ ] #6 Documentation updated - Comments explain navigation logic

**Final Verification Steps**:

1. **Run full test suite**:
   ```bash
   make test
   ```

2. **Check coverage**:
   ```bash
   make test-cover
   ```

3. **Review coverage report** (`coverage.out`) for:
   - All keyboard handlers tested
   - Wrap-around paths verified
   - Modal state transitions covered

4. **Manual UI verification** (if time permits):
   ```bash
   make run
   # Test: ↑/↓ navigation, Enter modal, Esc exit, 'c' form
   ```

5. **Verify no regressions**:
   - All existing tests pass
   - No new compiler warnings
   - No fmt changes needed

**Output Expected**:

- Task record updated with verification results
- If all tests pass: Task marked as Done
- If issues found: Document blockers, create follow-up tasks
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
