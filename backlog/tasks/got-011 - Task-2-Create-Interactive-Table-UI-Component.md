---
id: GOT-011
title: 'Task 2: Create Interactive Table UI Component'
status: In Progress
assignee:
  - Qwen Code
created_date: '2026-03-17 00:42'
updated_date: '2026-03-17 01:08'
labels: []
dependencies: []
references:
  - 'PRD: DCA Assets List Table View'
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create assets_view.go with Bubble Tea model pattern for interactive table UI with keyboard navigation
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Table displays with headers: Asset, Count, Total Shares, Avg Price, Total Value
- [ ] #2 Up/Down arrows navigate rows
- [ ] #3 Active row highlighted
- [ ] #4 Esc returns to menu or exits
- [ ] #5 Ctrl+C exits cleanly
- [ ] #6 No assets message displays when list is empty
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Create an `AssetsView` Bubble Tea component that displays aggregated asset data in an interactive table format. The component will:

- **Data Source**: Use existing `LoadAndAggregateEntries()` function from `assets_view.go` to load and aggregate entries from `dca_entries.json`
- **Navigation Model**: Implement keyboard navigation using Arrow Up/Down to move between rows, similar to `FormModel` but with table-specific logic
- **Styling**: Use Lipgloss for table borders, headers, and row highlighting (blue color 63 for active row)
- **State Management**: Track selected row index, handle key events, and render table
- **Exit Behavior**: Esc key exits to previous context, Ctrl+C exits application cleanly

**Architecture Decision**: Create a standalone `AssetsViewModel` that wraps `AssetsViewModel` with Bubble Tea integration methods (`Update`, `View`, `Init`). This keeps separation of concerns while reusing existing aggregation logic.

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `assets_view.go` | Modify | Add `AssetsView` Bubble Tea model type with `Update`, `View`, `Init` methods |
| `main.go` | Modify | Add state transition to assets view; integrate `AssetsView` into main model |
| `assets_view_test.go` | Add tests | Unit tests for `AssetsView` component (navigation, rendering, edge cases) |
| `dca_entries.json` | No change | Existing data file format works as-is |

**Files that need to be created/updated:**
1. **`assets_view.go`** - Add `AssetsView` struct and Bubble Tea methods
2. **`main.go`** - Add view state and navigation logic
3. **`assets_view_test.go`** - Add tests for table rendering and navigation

### 3. Dependencies

**Prerequisites:**
- ✅ Task GOT-010 (Assets View Model) must be complete - `LoadAndAggregateEntries()` function exists
- ✅ Existing Bubble Tea v1.3.10 and Lipgloss v1.1.0 in `go.mod`
- ✅ `dca_entries.json` format is stable (no changes needed)

**No new external dependencies required.**

**Blocking issues:** None

### 4. Code Patterns

**Follow existing patterns from `dca_form.go`:**

| Pattern | Implementation |
|---------|----------------|
| **Model structure** | `type AssetsView struct { rows []AssetSummary; selectedIndex int; loaded bool; error error }` |
| **Key handling** | `Update()` swits on `tea.KeyMsg` types |
| **Navigation** | Arrow keys adjust `selectedIndex`, bounds checking with modulo |
| **Styling** | Lipgloss with rounded borders, color 63 for active element |
| **Render helpers** | Separate `renderTable()`, `renderHeader()`, `renderFooter()` methods |
| **Exit behavior** | `tea.KeyEsc` and `tea.KeyCtrlC` return `tea.Quit` command |

**Specific conventions to follow:**
- 8-decimal precision for financial values (already handled in `assets_view.go`)
- Rounded borders for table using `lipgloss.RoundedBorder()`
- Error display with ❌ prefix (see `dca_form.go`)
- Footer help text for navigation hints
- Table column alignment with fixed widths

### 5. Testing Strategy

**Unit tests to add in `assets_view_test.go`:**

| Test Case | Purpose |
|-----------|---------|
| `TestAssetsView_Render_Empty` | Shows "No assets yet" when list is empty |
| `TestAssetsView_Render_SingleRow` | Renders single asset correctly |
| `TestAssetsView_Render_MultipleRows` | Renders table with multiple assets |
| `TestAssetsView_Navigate_Down` | Arrow Down increments selected index |
| `TestAssetsView_Navigate_Up` | Arrow Up decrements selected index |
| `TestAssetsView_Navigate_Wrap` | Wrap-around at bounds (down from last → first) |
| `TestAssetsView_Navigate_Bounds` | Clamp to valid range (up from first → no change) |
| `TestAssetsView_Update_Escape` | Esc key returns quit command |
| `TestAssetsView_Update_CtrlC` | Ctrl+C returns quit command |
| `TestAssetsView_Init` | Returns nil command on init |
| `TestAssetsView_LoadData` | Successfully loads and aggregates entries |

**Test approach:**
- Use table-driven tests for navigation edge cases
- Test exact output string matching for rendering
- Test key behavior matches Bubble Tea message types
- All error messages must use exact format from validation patterns

### 6. Risks and Considerations

**Potential issues:**

| Risk | Mitigation |
|------|------------|
| Column width calculation | Use Lipgloss `MaxWidth()` or fixed widths for numeric columns |
| Row highlighting style | Match existing blue (color 63) from form fields |
| Empty state rendering | Display centered message with proper styling |
| Data consistency | Reuse `LoadAndAggregateEntries()` - no new file I/O logic |
| Navigation wrapping | Implement wrap-around for intuitive UX (down from last → first) |
| Table rendering order | Sort assets alphabetically for predictable display |

**Design trade-offs:**
- **Wrap-around navigation**: Down from last row wraps to first (more intuitive than clamping)
- **No Enter key action**: Table is view-only; Enter behaves same as other keys (no-op or could select)
- **Simple column alignment**: Use `strings.Join()` with fixed column widths rather than complex table library
- **Single-file component**: Keep `AssetsView` in `assets_view.go` alongside model logic

**Implementation checkpoints:**
1. First: Render empty table with headers and "No assets" message
2. Second: Add row rendering and basic navigation
3. Third: Add row highlighting and styling
4. Fourth: Integrate with main.go and test full flow
5. Fifth: Add comprehensive tests for all acceptance criteria
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
