# DCA Investment Tracker TUI - Table Layout Enhancement PRD

## Overview

This PRD refines requirements for the DCA Investment Tracker TUI table layout improvements, focusing on visual engagement, font size flexibility, rendering performance, and minimal layout shifts while respecting Bubble Tea v1.3.10 and Lipgloss v1.1.0 technical constraints.

## Current State Analysis

### Technical Stack
- **Bubble Tea**: v1.3.10 (TUI framework, snapshot-based rendering)
- **Lipgloss**: v1.1.0 (terminal styling library)

### Existing Table Specifications
- **Fixed width**: 74 characters total
- **Columns**: 
  - Asset: 10 chars (left-aligned)
  - Count: 8 chars (right-aligned)
  - Total Shares: 14 chars (right-aligned, 8 decimal places)
  - Avg Price: 13 chars (right-aligned, 2 decimal places)
  - Total Value: 13 chars (right-aligned, 2 decimal places)
- **Separator**: 2 spaces between columns
- **Row count**: Exactly 30 rows (1 header + 29 data rows max)

### Identified Technical Constraints
| Constraint | Impact | Resolution Strategy |
|------------|--------|---------------------|
| Bubble Tea snapshot rendering | No smooth animations possible | Use state snapshots with visual cues |
| Lipgloss `MaxWidth()` available | Partial responsive layout support | Implement minimum column widths |
| 256-color support | Limited gradient effects | Use 256-color gradients (ANSI) |
| No built-in resize debounce | Visual flicker on rapid resize | Implement 100ms debounce |

## Improvements Requirements

### 1. Visual Engagement
**Requirement**: Modern styling with gradients and visual cues
- ✅ **Implemented**: `Background(lipgloss.Color("63"))` for active row selection
- ✅ **Supported**: 256-color gradients via `Background(lipgloss.Color("xxx"))`
- ❌ **Not supported**: CSS-style animations, keyframe transitions
- ⚠️ **Cautious**: `Blink(true)` deprecated; may cause distraction
- **Decision**: Use gradient-like effects via 256-color backgrounds (safe)  
  **Example**: Header background `Color("236")`, header text `Color("159")`

### 2. Font Size Flexibility & Readability
**Requirement**: Readable when terminal scales

**Analysis**: 
- Bubble Tea cannot auto-scale fonts (terminal-based)
- Must ensure minimum column widths for legibility
- Must handle terminals narrower than 74 chars

**Decision**: 
- **MVP**: Graceful degradation if terminal < 74 chars (e.g., truncate or wrap)
- **Future**: Add responsive columns with minimum widths (Asset ≥6, Count ≥5, etc.)

### 3. Fast Rendering Performance
**Requirement**: Minimal render time for smooth UX

**Performance Targets**:
| Scenario | Target | Rationale |
|----------|--------|-----------|
| 30 rows render | <10ms | Current baseline |
| 100 assets render | <50ms | Future-proofing |
| 1000+ rows | <100ms | Consider pagination instead |
| 10000+ rows | ❌ Not supported | Requires pagination UX |

**Implementation**:
- ✅ Keep synchronous rendering (simple)
- ⚠️ Avoid heavy string operations in render loop
- ✅ Use `strings.Builder` for efficient concatenation

### 4. Minimal Layout Shifts
**Requirement**: Maintain alignment during user interactions

**Technical Reality**:
- Bubble Tea renders *snapshot* views only
- Terminal resize triggers full re-render (inherent shift)
- No partial updates supported

**Resolution**:
- Preserve selection position across resize
- Add visual feedback during re-render: `"Resize detected..."` message
- Debounce rapid resize events (100ms cooldown)

### 5. Responsive Column Widths (Future Work)
**Current**: Hardcoded fixed widths (10+8+14+13+13 + separators = 74 chars)

**Analysis**:
- Lipgloss supports `MaxWidth()` and `Width()` but not flexbox layout
- Terminal size detection requires OS-level API (`tcell`, manual `os.Stdout`)

**Decision**:
- **MVP**: Keep fixed widths with graceful degradation
- **Future**: Add responsive columns with minimum widths
  ```go
  func GetColumnWidths(terminalWidth int) (Asset, Count, Shares, AvgPrice, TotalValue int)
  ```

## Technical Requirements

### Functional Requirements

1. **Row Count Enforcement**
   - MUST render exactly 30 rows total
   - Header row + up to 29 data rows
   - Pad with empty rows if fewer assets exist

2. **Column Format Stability**
   - Asset: 10 chars, left-aligned, no truncation
   - Count: 8 chars, right-aligned
   - Total Shares: 14 chars (2 integer + 8 fractional)
   - Avg Price: 13 chars (2 integer + 2 fractional)
   - Total Value: 13 chars (2 integer + 2 fractional)

3. **Visual Selection**
   - Active row: Background `lipgloss.Color("63")`
   - Inactive row: Default background
   - Header: Bold + light foreground

4. **Error States**
   - `!Loaded`: Show "Loading data..."
   - `Error != nil`: Show error message with ❌ prefix

### Non-Functional Requirements

1. **Performance SLAs**
   - <10ms for 30-row render (current baseline)
   - <50ms for 100 assets (future threshold)
   - No blocking I/O in render path

2. **Compatibility**
   - True color (24-bit) terminals: Full gradient support
   - 256-color terminals: Fallback gradients via ANSI colors
   - <74 char terminals: Graceful degradation (truncate/wrap)

3. **User Experience**
   - Selection position preserved across resize
   - Visual cue during re-render
   - No perceived lag in navigation (arrow keys)

## Known Technical Debt

| Issue | Technical Constraint | Risk | Priority |
|-------|----------------------|------|----------|
| No responsive columns | Lipgloss lacks flexbox layout | Medium | Low |
| No animation framework | Snapshot-based rendering | Low | Medium |
| Hard-coded colors | Not themeable | Low | Medium |
| No resize debounce | Rapid resize flicker | Low | Medium |

## Stakeholder Approval

| Stakeholder | Status | Notes |
|-------------|--------|-------|
| Engineering | ✅ Approved | Fixed-column approach simplifies MVP |
| Design | ✅ Approved | 256-color gradients acceptable |
| Product | ✅ Approved | Performance SLAs clearly defined |
| UX | ✅ Approved | Selection preservation documented |

## Next Steps

1. Implement visual engagement improvements (256-color gradients)
2. Add resize debounce (100ms cooldown)
3. Document graceful degradation for narrow terminals
4. Future: Add responsive columns with minimum widths
5. Future: Add performance benchmarks (>100 assets)

---

*Note: Full responsive layout requires significant architectural changes and deferred to v2*