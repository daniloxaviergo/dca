---
id: GOT-012
title: 'Task 3: Integrate Assets View into Main Application'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-17 00:42'
updated_date: '2026-03-17 08:32'
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

The AssetsView integration is already implemented in `main.go` with state management and transitions. This task's primary goal is to verify and test the complete application flow through integration tests.

**Testing Strategy:**
- Add integration tests to `main_test.go` covering state transitions
- Verify data consistency between views
- Test exit behavior from AssetsView
- Ensure clean application lifecycle

**Key Scenarios to Test:**
1. Application starts in Form state
2. Form submission (Enter + confirm 'y') transitions to AssetsView
3. AssetsView displays aggregated data correctly
4. Esc/Ctrl+C in AssetsView exits application
5. Data entered in Form is reflected in AssetsView after save

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--|
| `main_test.go` | Add tests | Integration tests for state transitions, data flow, and exit behavior |
| `main.go` | No change | Implementation already exists and tested |
| `assets_view.go` | No change | Component already exists from Task 2 |
| `dca_entry.go` | No change | Data model and I/O already exist |
| `dca_form.go` | No change | Form component unchanged |

**New Tests to Add:**

| Test Case | Purpose |
|-----------|---------|
| `TestAppState_InitState` | Application starts in Form state |
| `TestModel_Update_FormToAssets_Transition` | Form submission transitions to AssetsView |
| `TestModel_Update_FormToAssets_DataLoaded` | AssetsView loads data correctly on transition |
| `TestModel_Update_AssetsToForm_Exit` | AssetsView exit transitions back to Form |
| `TestModel_Update_QuitFromAssets` | Ctrl+C in AssetsView returns quit command |
| `TestModel_Update_KeyNavigation_Assets` | AssetsView Up/Down navigation works |
| `TestModel_View_AssetsRenders` | AssetsView renders table correctly |
| `TestModel_DataConsistency` | Form changes reflected in AssetsView after save |
| `TestMain_AppRoundtrip` | Full end-to-end flow test |

### 3. Dependencies

**Prerequisites (must be complete first):**
- ✅ Task GOT-010: Assets View Model and Data Aggregation
- ✅ Task GOT-011: Interactive Table UI Component
- ✅ Task GOT-012: Main.go integration (existing implementation)

**No external dependencies required.**

**Blocking issues:** None

**Setup steps:** None required

### 4. Code Patterns

**Follow existing patterns from main_test.go:**

| Pattern | Implementation |
|---------|----------------|
| **Test setup** | Create model with entries, form, and currentState |
| **Key messages** | Use `tea.KeyMsg{Type: tea.KeyCtrlC}`, `tea.KeyMsg{Type: tea.KeyEsc}` |
| **Update assertions** | Call Update(), verify cmd != nil for quit |
| **View assertions** | Check strings.Contains(view, expectedText) |
| **Test helper** | Use `m := model{form: form, currentState: StateForm}` pattern |

**Specific conventions:**
- Test exact quit command behavior: `if cmd == nil { t.Error(...) }`
- Use temp files for I/O tests (defer os.Remove)
- Verify state changes via `m.currentState`
- Test data consistency with shared `*DCAData` reference

**Style requirements:**
- Follow existing test naming: `Test{Component}_{Action}_{Condition}`
- Use table-driven tests for multiple similar scenarios
- Test error cases with exact error message assertions

### 5. Testing Strategy

**Integration tests in main_test.go:**

```go
// State initialization
func TestAppState_InitState(t *testing.T)

// Form to AssetsView transition
func TestModel_Update_FormToAssets_Transition(t *testing.T)
func TestModel_Update_FormToAssets_DataLoaded(t *testing.T)

// AssetsView to Form exit
func TestModel_Update_AssetsToForm_Exit(t *testing.T)
func TestModel_Update_QuitFromAssets(t *testing.T)

// AssetsView navigation
func TestModel_Update_KeyNavigation_Assets(t *testing.T)

// Rendering
func TestModel_View_AssetsRenders(t *testing.T)

// Data consistency
func TestModel_DataConsistency(t *testing.T)

// End-to-end flow
func TestMain_AppRoundtrip(t *testing.T)
```

**Test approach:**
- Mock form submission by setting `form.Submitted = true` and sending Enter key
- Verify AssetsView loads with data: check `assetsView.Loaded == true`
- Test exit behavior: AssetsView returns `tea.Quit` on Esc/Ctrl+C
- Verify data consistency: modify form, save, then check AssetsView data

**Edge cases:**
- Empty entries file → AssetsView shows empty state message
- No entries in AssetsView → "No assets yet" message displayed
- Multiple entries → All aggregated correctly in table

### 6. Risks and Considerations

**Potential issues:**

| Risk | Mitigation |
|------|------|
| State enum not exported | Use package-private tests (same package) |
| Nil pointer dereference | Check for nil views in Update/View before delegating |
| Data race conditions | Single-threaded Bubble Tea event loop; no concurrent access |

**Design trade-offs:**
- **Current implementation**: Form → AssetsView is one-way; must exit to return (acceptable per PRD)
- **No persistent selection**: Selected index resets on each AssetsView load (simpler state management)

**Implementation checkpoints:**
1. First: Add basic state transition tests (Form → AssetsView)
2. Second: Add AssetsView rendering tests
3. Third: Add exit behavior tests (Esc/Ctrl+C from AssetsView)
4. Fourth: Add data consistency tests
5. Fifth: Run full test suite and verify no regressions

**Blocking issues:** None identified

**DoD verification:**
- [ ] All acceptance criteria met (4 criteria in task)
- [ ] `go test ./...` passes
- [ ] `go build` completes without warnings
- [ ] `go fmt` applied (no changes)
- [ ] PRD referenced in task
- [ ] Test coverage for state transitions
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
