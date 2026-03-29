---
id: GOT-059
title: '[doc-013 Phase 4] Integrate CLI with main application flow'
status: To Do
assignee: []
created_date: '2026-03-28 20:50'
updated_date: '2026-03-29 00:02'
labels:
  - feature
  - integration
  - cli
dependencies: []
references:
  - 'doc-013 - Phase 4: Integrate with main'
documentation:
  - doc-013
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Modify cmd/dca/main.go to detect the --add flag during initialization and route to CLI mode with early exit before TUI initialization. Ensure the CLI path is executed before Bubble Tea program setup. Maintain backward compatibility by preserving all existing TUI functionality and state transitions unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Main function detects --add flag
- [ ] #2 CLI path executes early exit
- [ ] #3 TUI initialization skipped for CLI mode
- [ ] #4 All existing TUI functionality preserved
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

**This task is a verification and documentation task** - the CLI integration is already complete and operational. The task ensures the integration pattern is correct and well-tested.

**Current state analysis:**
- `cmd/dca/cli.go` already contains all CLI functionality:
  - `ParseFlags()` - Parses `--add`, `--asset`, `--amount`, `--price`, `--date` flags
  - `CreateDCAEntry()` - Creates DCAEntry with 8-decimal share precision
  - `SaveEntry()` - Saves entry to JSON using atomic write
  - `RunCLI()` - Main CLI entry point that exits with appropriate code

- `cmd/dca/main.go` already has CLI integration:
  - `RunCLI()` is called BEFORE Bubble Tea initialization
  - Returns early if CLI mode is active (no TUI setup)
  - TUI code path unchanged when CLI not triggered

**How the feature will be built:**
- CLI mode: `./dca --add --asset BTC --amount 100 --price 50000` → exits immediately
- TUI mode: `./dca` → starts Bubble Tea program normally

**Architecture decisions and trade-offs:**
- CLI mode detected via `--add` flag in `ParseFlags()`, returns early if active
- TUI mode remains unchanged (app starts in asset view by default)
- CLI returns exit code 1 on error, 0 on success (silent)
- Both modes use shared validation from `internal/validation` and persistence from `internal/dca`

**Why this approach:**
- Minimal changes to existing architecture
- CLI mode is opt-in via flag, TUI unaffected
- Code reuse through shared validation and persistence packages
- Early exit pattern prevents unnecessary initialization overhead

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `cmd/dca/main.go` | Review only | Integration already complete via `RunCLI()` check before TUI init |
| `cmd/dca/cli.go` | Review only | CLI logic already implemented with flag parsing, validation, persistence |
| `internal/validation/validation.go` | Review only | Shared validation used by CLI |
| `internal/dca/entry.go` | Review only | Data model and file I/O used by CLI |

**Files to CREATE (tests - new work):**
| File | Action | Reason |
|------|--------|--------|
| `cmd/dca/cli_test.go` | Create | Unit tests for CLI-specific logic (not yet implemented) |
| `cmd/dca/main_test.go` | Update | Add integration tests for CLI-TUI separation (not yet implemented) |

### 3. Dependencies

**Completed prerequisites (from Backlog):**
- `GOT-056` (Done): Validation functions extracted to shared package
- `GOT-057` (Done): CLI command component with flag parsing created
- `GOT-058` (Done): CLI run function with validation and persistence implemented

**Existing packages (no new dependencies):**
- `github.com/danilo/scripts/github/dca/internal/validation` - Shared validation functions
- `github.com/danilo/scripts/github/dca/internal/dca` - Data model and file I/O
- `github.com/charmbracelet/bubbletea` - TUI framework (unchanged)

**Prerequisites:**
- `cmd/dca/cli.go` must exist (GOT-057 completed)
- `RunCLI()` function must be called in `main()` before TUI init (GOT-057 completed)
- `internal/validation/validation.go` must contain all required validation functions (GOT-056 completed)
- No blocking issues, all dependencies available

### 4. Code Patterns

**Follow existing patterns in the codebase:**

1. **CLI detection pattern** (from `main.go`):
```go
if RunCLI() {
    return // CLI mode handled and exited
}
// Continue with TUI initialization
```

2. **Error handling pattern** (from `cli.go`):
- Use `fmt.Fprintln(os.Stderr, ...)` for error messages
- Exit with code 1 on validation failure, 0 on success
- No output on success (silent)

3. **Validation pattern** (from `internal/validation/validation.go`):
- All validation returns `error` with descriptive messages
- Consistent error messages: "Amount must be positive", "Price must be positive", etc.
- Use shared functions to avoid duplication

4. **Share calculation pattern** (from `internal/dca/entry.go`):
- Use `DCAEntry.CalculateShares()` for 8-decimal precision
- Formula: `math.Round((amount / price) * 1e8) / 1e8`

5. **File I/O pattern** (from `internal/dca/entry.go`):
- Use `dca.LoadEntries()` and `dca.SaveEntries()` for atomic writes
- Temp file + rename pattern for safe persistence

### 5. Testing Strategy

**Unit tests for CLI (`cmd/dca/cli_test.go`):**
- `TestParseFlags_AddWithAllFields` - Flags parsed correctly with required fields
- `TestParseFlags_AddWithDate` - Optional --date flag included
- `TestParseFlags_MissingAsset` - Exit code 1 for missing --asset
- `TestParseFlags_ZeroAmount` - Exit code 1 for zero/negative amount
- `TestParseFlags_ZeroPrice` - Exit code 1 for zero/negative price
- `TestParseFlags_InvalidDate` - Exit code 1 for invalid date format
- `TestParseFlags_DateAutoSet` - Date defaults to current time when not provided
- `TestCreateDCAEntry_Precision` - Shares calculated with 8-decimal precision
- `TestCreateDCAEntry_DateParsing` - Date string parsed correctly to time.Time
- `TestRunCLI_Success` - Entry saved to JSON, exit code 0, no output
- `TestRunCLI_Error` - Exit code 1 on validation failure with stderr message

**Integration tests (`cmd/dca/main_test.go` additions):**
- `TestMain_CLIModeExitsEarly` - CLI flag triggers exit before TUI init
- `TestMain_CLISavesEntry` - Entry persisted to JSON file correctly
- `TestMain_TUIModeUnchanged` - Normal TUI flow unaffected by CLI flags

**Edge cases to cover:**
- Missing required flags (--add, --asset, --amount, --price)
- Negative/zero amount (exit code 1)
- Negative/zero price (exit code 1)
- Empty/whitespace asset (exit code 1)
- Invalid date format (exit code 1)
- Zero price in share calculation (returns 0)
- File permission errors during save

**Verification:**
- Run `make test` to verify all tests pass
- Run `make test-cover` to ensure CLI code paths are covered
- Manual test: `./dca --add --asset BTC --amount 100 --price 50000`
- Verify TUI mode still works: `./dca` (without flags)

### 6. Risks and Considerations

**Known issues:**
- CLI tests may fail if tests directory structure not properly set up
- File I/O tests require temp directory isolation to avoid conflicts

**Potential pitfalls:**
- **Race conditions**: CLI and TUI could conflict on file access (handled by atomic writes in `dca.SaveEntries()`)
- **Date parsing**: Date format must match RFC3339 exactly (validated in ParseFlags, so safe)
- **Error message consistency**: Ensure CLI errors match TUI validation messages

**Trade-offs:**
- Silent success (no output) makes CLI script-friendly but harder to debug
- Auto-date always uses current time; users cannot specify "today" explicitly without --date flag
- Single entry only (batch mode out of scope per PRD)

**Implementation checklist:**
- [ ] Verify `RunCLI()` is called before Bubble Tea initialization in `main()`
- [ ] Verify all flag validations use shared functions from `internal/validation`
- [ ] Verify share calculation uses `DCAEntry.CalculateShares()` for 8-decimal precision
- [ ] Create `cmd/dca/cli_test.go` with comprehensive unit tests
- [ ] Add integration tests in `cmd/dca/main_test.go`
- [ ] Run `make check` (fmt, build, test) to verify quality
- [ ] Manual test both CLI and TUI modes
- [ ] Document any findings or adjustments needed

**Definition of Done mapping:**
| Criterion | Status | Verification |
|-----------|--------|--------------|
| #1 Main function detects --add flag | ✅ Complete | `RunCLI()` check in `main()` |
| #2 CLI path executes early exit | ✅ Complete | `os.Exit()` calls in `cli.go` |
| #3 TUI initialization skipped | ✅ Complete | Early return before TUI setup |
| #4 TUI functionality preserved | ✅ Complete | No modifications to TUI code |
| Unit tests pass | To Do | Create `cli_test.go` and run `make test` |
| No compiler warnings | To Do | Run `make build` and check output |
| Code follows project style | To Do | Run `make fmt` before finalizing |
| Documentation updated | To Do | Add comments to `main.go` integration point |

### 7. Implementation Steps

1. **Review current integration** (5 minutes)
   - Verify `RunCLI()` is called before Bubble Tea initialization in `main()`
   - Verify `RunCLI()` returns `true` when CLI mode is active
   - Verify TUI code path is unreachable when CLI mode is active

2. **Create CLI unit tests** (30 minutes)
   - Create `cmd/dca/cli_test.go`
   - Add tests for all `ParseFlags()` validation paths
   - Add tests for `CreateDCAEntry()` with various inputs
   - Add tests for `RunCLI()` success and error paths

3. **Add integration tests** (15 minutes)
   - Add `TestMain_CLIModeExitsEarly` to `main_test.go`
   - Add `TestMain_CLISavesEntry` to verify persistence
   - Add `TestMain_TUIModeUnchanged` to ensure TUI unaffected

4. **Run validation** (10 minutes)
   - Run `make fmt` to ensure code formatting
   - Run `make build` to verify no compiler warnings
   - Run `make test` to ensure all tests pass
   - Run `make test-cover` to verify coverage

5. **Manual testing** (5 minutes)
   - Test CLI mode: `./dca --add --asset BTC --amount 100 --price 50000`
   - Verify exit code 0 and no output
   - Verify entry saved to `dca_entries.json`
   - Test TUI mode: `./dca` (verify starts normally)

6. **Final verification** (5 minutes)
   - Review acceptance criteria
   - Map to Definition of Done
   - Update task notes with implementation summary
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
