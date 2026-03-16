---
id: GOT-005
title: 'Task 1: Define DCA data model (structs, JSON serialization)'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-16 21:24'
updated_date: '2026-03-16 21:33'
labels:
  - data-model
  - core
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define Go structs for DCA entries and implement JSON serialization/deserialization. Create dca_entry.go file with proper data model.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Define DCAEntry struct with fields: Amount (float64), Date (time.Time), Asset (string), PricePerShare (float64), Shares (float64)
- [ ] #2 Define DCAData struct with map[string][]DCAEntry keyed by asset ticker
- [ ] #3 Implement LoadEntries() function to read from dca_entries.json
- [ ] #4 Implement SaveEntries() function to write to dca_entries.json with 2-space indentation
- [ ] #5 Add Validate() method on DCAEntry to check Amount > 0, PricePerShare > 0
- [ ] #6 Add CalculateShares() method to compute Shares = Amount / PricePerShare with 8 decimal precision
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task establishes the foundation for data persistence by defining the core data structures and JSON serialization layer.

**Design Decisions:**
- **Separate file per domain**: Create `dca_entry.go` to house all data model types and methods, keeping concerns separated from the UI layer (`main.go`)
- **Canonical data model**: Define `DCAEntry` as the source of truth with all required fields; `DCAData` wraps a map for efficient asset lookups
- **Precision handling**: Use `float64` for monetary values (Amount, PricePerShare, Shares) as Go standard library JSON encoder handles this well for financial applications; 8 decimal precision for shares via `math.Round` or string formatting
- **Validation pattern**: Implement `Validate()` method returning an error to allow flexible error aggregation in higher layers
- **File naming**: `dca_entries.json` (plural) aligns with the PRD's example structure

**Why this approach:**
- Separation of concerns: Data model isolated from UI logic
- Direct JSON mapping: No custom marshaling needed for basic structs (using `json` tags)
- Idiomatic Go: Pointer receivers for mutator methods, value receivers for pure functions
- Testability: Simple structs with clear validation logic are easy to unit test

### 2. Files to Modify

| Action | File | Reason |
|--------|------|--------|
| Create | `dca_entry.go` | New file containing `DCAEntry`, `DCAData` structs and all methods |
| Modify | `main.go` | Import and reference new data types (future task integration) |
| Create (runtime) | `dca_entries.json` | Auto-created by `SaveEntries()` if not present |

**No breaking changes** to existing code; this is additive.

### 3. Dependencies

**Existing dependencies (already in go.mod):**
- `github.com/charmbracelet/bubbletea` - UI framework (no direct use in data model)
- `github.com/charmbracelet/lipgloss` - Styling (no direct use in data model)

**No new dependencies required** - standard library packages sufficient:
- `encoding/json` - JSON serialization
- `os` - File I/O operations
- `math` - Rounding for share calculation precision
- `errors` - Error handling
- `time` - RFC3339 date handling

### 4. Code Patterns

**Go conventions to follow:**
- **File naming**: `dca_entry.go` (lowercase, snake_case for multi-word, matches package purpose)
- **Struct field naming**: Exported fields with PascalCase (`Amount`, `Date`, etc.) for JSON serialization
- **JSON tags**: Use `json:"amount"` style for snake_case JSON keys matching PRD example
- **Method receiver naming**: Use single letter (`e *DCAEntry`) for pointer receivers
- **Error handling**: Return `error` from methods that may fail (`LoadEntries`, `Validate`)
- **Comments**: Document exported types and functions with Godoc-style comments
- **Import grouping**: Standard library imports grouped together, then external

**Naming conventions:**
- Structs: `DCAEntry`, `DCAData`
- Functions: `LoadEntries()`, `SaveEntries()`, `CalculateShares()`
- Methods: `Validate()`, `CalculateShares()` (on DCAEntry)

### 5. Testing Strategy

**Unit tests to write in `dca_entry_test.go`:**

| Test | Coverage |
|------|----------|
| `TestDCAEntryValidate_Pass` | Amount > 0, PricePerShare > 0, valid date |
| `TestDCAEntryValidate_ZeroAmount` | Reject Amount = 0 |
| `TestDCAEntryValidate_NegativeAmount` | Reject Amount < 0 |
| `TestDCAEntryValidate_ZeroPrice` | Reject PricePerShare = 0 |
| `TestDCAEntryValidate_NegativePrice` | Reject PricePerShare < 0 |
| `TestCalculateShares` | Verify 8 decimal precision (e.g., 500/65000 = 0.00769231) |
| `TestCalculateShares_Precision` | Verify rounding behavior |
| `TestSaveEntries_CreateFile` | Creates new file with proper structure |
| `TestSaveEntries_UpdateFile` | Appends to existing data |
| `TestLoadEntries_Populated` | Reads existing JSON correctly |
| `TestLoadEntries_EmptyFile` | Handles empty/missing file gracefully |
| `TestLoadEntries_InvalidJSON` | Returns error on malformed JSON |

**Testing approach:**
- Use Go's standard `testing` package
- Create temporary files via `os.CreateTemp()` for file I/O tests
- Clean up temp files in `t.Cleanup()`
- Test with realistic values from PRD examples (BTC: 500/65000, ETH: 200/3200)

### 6. Risks and Considerations

**Blocking issues:** None - this is a standalone data model task.

**Design trade-offs:**
1. **Float precision**: Using `float64` for Shares may introduce floating-point errors; consider if future analysis needs arbitrary precision (could require `math/big.Rat` or string-based storage)
2. **No array support in DCAData**: Using `map[string][]DCAEntry` means no guaranteed order; if insertion order matters, consider `struct{ Ticker string; Entries []DCAEntry }` or separate ordering field
3. **File locking**: `SaveEntries()` does not implement file locking; concurrent writes could corrupt file (defer to future task with multi-user requirements)
4. **No schema versioning**: Future schema changes may require migration logic (not in scope for v1)

**Implementation notes:**
- JSON keys will be lowercase (`amount`, `date`, `asset`, `pricePerShare`, `shares`) matching PRD example
- Date stored as RFC3339 string (not timestamp int) for human readability
- `Validate()` returns error for flexibility; could be extended to return validation results list
- `CalculateShares()` returns `float64` with 8 decimals rounded; consider if caller needs exact precision or if displayed value suffices
<!-- SECTION:PLAN:END -->
