---
id: GOT-060
title: '[doc-013 Phase 5] Add unit tests for CLI validation and persistence'
status: To Do
assignee: []
created_date: '2026-03-28 20:50'
updated_date: '2026-03-29 00:54'
labels:
  - testing
  - cli
  - validation
dependencies: []
references:
  - 'doc-013 - Phase 5: Add tests'
documentation:
  - doc-013
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/dca/cli_test.go with comprehensive unit tests covering CLI validation logic, share calculation precision (8 decimals), date auto-generation, and data persistence. Tests must cover all error conditions (missing flags, invalid values) with exit code verification and ensure no regression in existing tests (go test -v ./...).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 cli_test.go created with test file
- [ ] #2 Validation tests for all flag combinations
- [ ] #3 Share calculation precision verified (8 decimals)
- [ ] #4 Date auto-generation tested
- [ ] #5 Exit codes verified for errors
- [ ] #6 All existing tests still pass
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The goal is to add comprehensive unit tests for the CLI validation and persistence layer. The implementation will focus on:

1. **Test the validation functions under `internal/validation/`** - These functions are shared between CLI and TUI forms
2. **Test CLI-specific validation logic** - Tests that ParseFlags properly validates all inputs and returns errors
3. **Test CLI persistence** - Tests that entries are saved correctly to the JSON file
4. **Test share calculation precision** - Verify 8-decimal precision with various inputs
5. **Test date auto-generation** - Verify that when no date is provided, current time is used

The approach will follow existing patterns in the codebase:
- Use `os.Args` manipulation for flag testing
- Use temporary files to avoid polluting actual data
- Test error conditions with exit code verification (via panic recovery for os.Exit calls)
- Test success paths with file content verification

### 2. Files to Modify

**Existing Files (tests to add/enhance):**
- `/home/danilo/scripts/github/dca/cmd/dca/cli_test.go` - Add missing tests
- `/home/danilo/scripts/github/dca/internal/validation/validation_test.go` - Add CLI-specific validation tests
- `/home/danilo/scripts/github/dca/internal/dca/entry_test.go` - Add share calculation precision tests

**Analysis Files (read-only, for understanding):**
- `/home/danilo/scripts/github/dca/cmd/dca/cli.go` - CLI implementation to test
- `/home/danilo/scripts/github/dca/internal/validation/validation.go` - Shared validation to test
- `/home/danilo/scripts/github/dca/internal/dca/entry.go` - Data model with share calculation

### 3. Dependencies

**Prerequisites:**
- `internal/validation/` package must exist (already exists, task GOT-056 completed)
- `cmd/dca/cli.go` must exist (already exists, task GOT-057 completed)
- `cmd/dca/main.go` must integrate CLI (already exists, task GOT-059 completed)
- `internal/dca/` package with file I/O (already exists, tasks GOT-007, GOT-015 completed)

**No new dependencies required** - All needed packages are already in `go.mod`.

### 4. Code Patterns

**Follow existing patterns from the codebase:**

1. **Test structure:**
   ```go
   func TestParseFlags_XYZ(t *testing.T) {
       originalArgs := os.Args
       defer func() { os.Args = originalArgs }()
       
       os.Args = []string{"dca", "-flag", "value", ...}
       
       // Test logic here
   }
   ```

2. **Error exit code verification:**
   ```go
   defer func() {
       if r := recover(); r != nil {
           // Expected: os.Exit() was called
       }
   }()
   ParseFlags() // This calls os.Exit() on error
   ```

3. **Temporary file testing:**
   ```go
   tmpDir, err := os.MkdirTemp("", "dca-test")
   defer os.RemoveAll(tmpDir)
   originalEntriesPath := defaultEntriesPath
   defaultEntriesPath = tmpFile
   defer func() { defaultEntriesPath = originalEntriesPath }()
   ```

4. **Share calculation precision:**
   ```go
   shares := math.Round((amount / price) * 1e8) / 1e8 // 8 decimal rounding
   ```

5. **Date auto-generation:**
   ```go
   if !cliData.HasDate {
       cliData.Date = time.Now().Format(time.RFC3339)
   }
   ```

### 5. Testing Strategy

**Unit Tests to Add:**

**A. Validation Tests (internal/validation/validation_test.go):**
- Test each validation function with valid/invalid inputs
- Verify exact error messages (CLI依赖 specific messages)
- Test edge cases: whitespace-only, empty strings, boundary values

**B. CLI Flag Parsing Tests (cmd/dca/cli_test.go):**
- `TestParseFlags_AddWithAllFields` - All flags provided
- `TestParseFlags_MissingAsset` - Exit code 1
- `TestParseFlags_ZeroAmount` - Exit code 1
- `TestParseFlags_NegativeAmount` - Exit code 1
- `TestParseFlags_ZeroPrice` - Exit code 1
- `TestParseFlags_NegativePrice` - Exit code 1
- `TestParseFlags_InvalidDate` - Exit code 1
- `TestParseFlags_EmptyAsset` - Exit code 1
- `TestParseFlags_WhitespaceAsset` - Exit code 1
- `TestParseFlags_DateAutoSet` - Date defaults to now
- `TestParseFlags_NoAddFlag` - Returns (Add=false, nil)

**C. CLI Entry Creation Tests (cmd/dca/cli_test.go):**
- `TestCreateDCAEntry_Precision` - 8-decimal precision verified
- `TestCreateDCAEntry_DateParsing` - Date string parsed correctly
- `TestCreateDCAEntry_Validation` - Data flows through correctly

**D. CLI Persistence Tests (cmd/dca/cli_test.go):**
- `TestRunCLI_Success` - Entry saved, exit code 0
- `TestRunCLI_Error` - Validation failure, exit code 1
- `TestRunCLI_MissingRequiredFlags` - Missing flags handled
- `TestRunCLI_NonAddMode` - Returns false when no --add
- `TestSaveEntry_AddsToExisting` - Appends to existing entries
- `TestSaveEntry_NewAsset` - Creates new asset array
- `TestRunCLI_SilentSuccess` - No output on success

**E. Share Calculation Precision Tests (internal/dca/entry_test.go):**
- Test various amount/price combinations
- Verify rounding behavior at 8th decimal place
- Test edge cases: very small prices, large amounts

**F. Date Handling Tests (cmd/dca/main_test.go):**
- `TestMain_DateHandling` - Explicit date used
- `TestMain_AutoDate` - Current time when not provided

**G. Integration Tests:**
- `TestMain_MultipleEntries` - Multiple CLI calls work
- `TestMain_ConcurrentEntries` - Multiple assets
- `TestMain_ValidationErrors` - Error messages

### 6. Risks and Considerations

**Potential Issues:**
1. **os.Exit handling** - Tests using panic/recovery for os.Exit() need careful cleanup
2. **Global state** - `defaultEntriesPath` is a package-level variable that needs restoration after tests
3. **Concurrent test execution** - Tests should use unique temp filenames or directories
4. **Date precision** - Auto-generated dates have slight timing variance; tests should allow tolerance
5. **Floating point comparison** - Use exact equality for 8-decimal values (mathematically correct)

**Trade-offs:**
1. **Exit code testing** - Cannot capture exit codes in tests, rely on panic recovery for os.Exit calls
2. **Stderr output** - Capture stderr via os.Pipe for silent success verification
3. **Validation messages** - CLI depends on exact error message format from validation package

**Verification Steps:**
1. Run `go test -v ./...` - All tests pass
2. Run `make test-cover` - Coverage >= 90% for CLI and validation packages
3. Run `make fmt` - No formatting changes
4. Run `make build` - Build succeeds
5. Manual test: `./dca -add -asset BTC -amount 100 -price 50000` - Works correctly
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
