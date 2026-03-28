---
id: GOT-057
title: '[doc-013 Phase 2] Create CLI command component with flag parsing'
status: To Do
assignee: []
created_date: '2026-03-28 20:49'
updated_date: '2026-03-28 23:31'
labels:
  - feature
  - cli
  - input
dependencies: []
references:
  - 'doc-013 - Phase 2: Create CLI component'
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/dca/cli.go with CLI-specific logic including flag parsing for --add, --asset, --amount, --price, and optional --date flags. Implement command-line argument parsing using Go's flag package with strict validation (exit code 1 on error). The CLI component must integrate with the shared validation functions and prepare entry data for persistence without requiring TUI initialization.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 cli.go created with flag parsing logic
- [x] #2 All required flags implemented (--add, --asset, --amount, --price)
- [x] #3 Optional --date flag with now() default
- [x] #4 Exit codes implemented for error handling
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The CLI command component will be implemented as a new `cmd/dca/cli.go` file that handles command-line argument parsing and entry creation without requiring the interactive TUI. The approach follows the existing architecture patterns:

1. **Flag-based entry point**: Use Go's `flag` package to parse CLI arguments, with `--add` flag triggering CLI mode (early exit before Bubble Tea initialization)
2. **Shared validation**: Leverage existing `internal/validation` package functions for strict validation (no duplication)
3. **Direct persistence**: Use `internal/dca.SaveEntries()` for atomic file writes
4. **Auto-calculation**: Calculate shares using `validation.CalculateSharesFromValues()` and use `time.Now()` for date

**Key architectural decisions:**
- CLI mode detected via `--add` flag, exits immediately after processing
- TUI mode unchanged (app starts in asset view by default)
- CLI returns exit code 1 on error, 0 on success (silent on success per PRD)
- Uses existing shared validation and persistence layers

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `cmd/dca/cli.go` | Create | New CLI component with flag parsing |
| `cmd/dca/main.go` | Modify | Add CLI flag detection and early exit logic |
| `internal/validation/validation_test.go` | Create | Unit tests for CLI-specific validation scenarios |
| `cmd/dca/main_test.go` | Create | Integration tests for CLI mode |

### 3. Dependencies

**Existing packages (no new dependencies required):**
- `github.com/danilo/scripts/github/dca/internal/validation` - Validation functions
- `github.com/danilo/scripts/github/dca/internal/dca` - Data model and file I/O
- `github.com/charmbracelet/bubbletea` - TUI (unchanged, only used when CLI not triggered)

**Prerequisites:**
- Validation package must exist (DONE - GOT-056 completed)
- DCA file I/O must exist (DONE - internal/dca/file.go in entry.go)
- No blocking issues, all dependencies available

### 4. Code Patterns

**Follow existing patterns in the codebase:**

1. **Error handling**: Use formatted error messages consistent with validation package
2. **Exit codes**: 0 for success, 1 for any error (including validation failures)
3. **File I/O**: Use `dca.SaveEntries()` for atomic writes
4. **Share calculation**: Use `validation.CalculateSharesFromValues()` with 8 decimal precision
5. **Date handling**: Use `time.Now()` with RFC3339 format

**CLI command signature:**
```bash
./dca --add --asset <ticker> --amount <usd> --price <per-share> [--date <rfc3339>]
```

**Required flags:** `--add`, `--asset`, `--amount`, `--price`  
**Optional flags:** `--date` (defaults to `time.Now().Format(time.RFC3339)`)

### 5. Testing Strategy

**Unit tests (`cmd/dca/cli_test.go`):**
- `TestCLI_ParseFlags_Add` - Flags parsed correctly with required fields
- `TestCLI_ParseFlags_AddWithDate` - Optional date flag included
- `TestCLI_ValidateAmount` - Amount validation (positive, parseable)
- `TestCLI_ValidatePrice` - Price validation (positive, parseable)
- `TestCLI_ValidateAsset` - Asset validation (non-empty, non-whitespace)
- `TestCLI_ValidateDate` - Date validation (RFC3339 format)
- `TestCLI_CalculateShares` - Shares computed correctly (8 decimals)
- `TestCLI_ExitOnError` - Exit code 1 on validation failure
- `TestCLI_ExitOnSuccess` - Exit code 0 on successful entry

**Integration tests (`cmd/dca/main_test.go` additions):**
- `TestMain_CLIModeExitsEarly` - CLI flag triggers early exit before TUI init
- `TestMain_CLISavesEntry` - Entry persisted to JSON file correctly
- `TestMain_TUIModeUnchanged` - Normal TUI flow unaffected by CLI flags

**Edge cases:**
- Missing required flags (exit code 1)
- Negative/zero amount (exit code 1)
- Negative/zero price (exit code 1)
- Empty/whitespace asset (exit code 1)
- Invalid date format (exit code 1)
- Zero price in share calculation (returns 0)

### 6. Risks and Considerations

**Low risk items:**
- **No breaking changes**: CLI mode is opt-in via `--add` flag, TUI unchanged
- **Code reuse**: All validation/persistence logic already exists
- **Test coverage**: Existing patterns can be extended

**Potential challenges:**
- **Flag parsing in Go**: Need to handle unknown flags gracefully (Go's `flag` package exits by default on unknown flags, which may need custom handling if stricter control required)
- **Shared state**: Must ensure CLI doesn't interfere with TUI's file locking or concurrent access (atomic writes in `dca.SaveEntries()` handles this)
- **Date format consistency**: Must use `time.RFC3339` for both default and parsed dates

**Trade-offs:**
- CLI exits after single entry (per PRD: "Single entry only" out of scope for batch mode)
- Silent success (no output on success per PRD)
- No verbose mode (`-v` flag out of scope per PRD)

**Deployment considerations:**
- No data migration needed (new feature, no schema changes)
- Backward compatible (existing TUI usage unaffected)
- Binary size increase negligible (standard library only)
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
2026-03-28: Created cmd/dca/cli.go with flag parsing logic using Go's flag package

2026-03-28: Added --add, --asset, --amount, --price, --date flags

2026-03-28: Implemented validation using internal/validation package

2026-03-28: Modified cmd/dca/main.go to check for CLI mode before TUI initialization

2026-03-28: Created internal/dca/entry.go with DCAEntry and DCAData types and file I/O

2026-03-28: CLI exits with code 1 on validation errors, code 0 on success (silent)

2026-03-28: Tested CLI functionality: entry creation, validation, date defaults

2026-03-28: All 148 tests pass, code formatted with go fmt

2026-03-28: CLI flag signature: ./dca --add --asset <ticker> --amount <usd> --price <per-share> [--date <rfc3339>]
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
