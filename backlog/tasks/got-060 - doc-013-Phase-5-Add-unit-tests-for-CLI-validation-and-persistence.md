---
id: GOT-060
title: '[doc-013 Phase 5] Add unit tests for CLI validation and persistence'
status: To Do
assignee: []
created_date: '2026-03-28 20:50'
updated_date: '2026-03-29 00:56'
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

The goal is to enhance CLI and validation test coverage to ensure complete coverage of the command-line entry feature. Analysis shows:

1. **CLI tests exist** but need coverage verification and some edge cases
2. **Validation tests exist** but may lack CLI-specific edge case coverage
3. **Entry tests exist** but need share calculation precision verification with various values

The approach will:
1. Run coverage analysis to identify gaps
2. Add tests for uncovered edge cases
3. Ensure all error paths return exit code 1
4. Verify 8-decimal precision across varying inputs
5. Test date auto-generation with various timing scenarios

### 2. Files to Modify

**Existing Files to Enhance:**
- `/home/danilo/scripts/github/dca/cmd/dca/cli_test.go` - Add coverage for uncovered edge cases
- `/home/danilo/scripts/github/dca/internal/validation/validation_test.go` - Add CLI-specific validation tests
- `/home/danilo/scripts/github/dca/internal/dca/entry_test.go` - Enhance share calculation tests

**No files to create** - Tests already exist in cli_test.go, validation_test.go, entry_test.go

**Analysis Files (read-only):**
- `/home/danilo/scripts/github/dca/cmd/dca/cli.go` - CLI implementation
- `/home/danilo/scripts/github/dca/internal/validation/validation.go` - Shared validation
- `/home/danilo/scripts/github/dca/internal/dca/entry.go` - Data model with share calculation

### 3. Dependencies

**Existing packages (no new dependencies):**
- `internal/validation/` - Shared validation functions (GOT-056)
- `internal/dca/` - Data model and file I/O (GOT-007, GOT-015)
- `cmd/dca/cli.go` - CLI implementation (GOT-057)
- `cmd/dca/main.go` - CLI integration (GOT-059)

**No new dependencies required**

### 4. Code Patterns

**Follow existing patterns in cli_test.go:**

1. **Flag parsing tests:**
   ```go
   func TestParseFlags_XYZ(t *testing.T) {
       originalArgs := os.Args
       defer func() { os.Args = originalArgs }()
       
       os.Args = []string{"dca", "-add", "-asset", "BTC", ...}
       
       // Test parsing with/without error
   }
   ```

2. **Exit code verification:**
   ```go
   defer func() {
       os.Stderr = oldStderr
   }()
   
   func() {
       defer func() {
           if r := recover(); r != nil {
               // os.Exit called
           }
       }()
       ParseFlags()
   }()
   ```

3. **Temporary file testing:**
   ```go
   tmpDir, _ := os.MkdirTemp("", "dca-test")
   defer os.RemoveAll(tmpDir)
   originalPath := defaultEntriesPath
   defaultEntriesPath = tmpFile
   defer func() { defaultEntriesPath = originalPath }()
   ```

4. **Share calculation precision:**
   ```go
   shares := math.Round((amount / price) * 1e8) / 1e8
   ```

5. **Date handling:**
   ```go
   // Auto-generated dates use time.Now()
   // Must allow timing variance in tests
   now := time.Now()
   diff := now.Sub(entry.Date)
   if diff > 5*time.Second {
       t.Errorf("Date should be current time")
   }
   ```

### 5. Testing Strategy

**Enhancement Tests to Add:**

**A. CLI ParseFlags Tests (cmd/dca/cli_test.go) - Complete Coverage:**
1. `TestParseFlags_AddWithAllFields` ✓ (existing)
2. `TestParseFlags_AddWithDate` ✓ (existing)
3. `TestParseFlags_MissingAsset` ✓ (existing)
4. `TestParseFlags_ZeroAmount` ✓ (existing)
5. `TestParseFlags_NegativeAmount` ✓ (existing)
6. `TestParseFlags_ZeroPrice` ✓ (existing)
7. `TestParseFlags_NegativePrice` ✓ (existing)
8. `TestParseFlags_InvalidDate` ✓ (existing)
9. `TestParseFlags_EmptyAsset` ✓ (existing)
10. `TestParseFlags_WhitespaceAsset` ✓ (existing)
11. `TestParseFlags_DateAutoSet` ✓ (existing)
12. `TestParseFlags_NoAddFlag` ✓ (existing)

**B. CLI CreateDCAEntry Tests (cmd/dca/cli_test.go):**
1. `TestCreateDCAEntry_Precision` ✓ (existing)
2. `TestCreateDCAEntry_DateParsing` ✓ (existing)
3. `TestCreateDCAEntry_Validation` ✓ (existing)
4. `TestCreateDCAEntry_InvalidDate` ✓ (existing)

**C. CLI Persistence Tests (cmd/dca/cli_test.go):**
1. `TestRunCLI_Success` ✓ (existing)
2. `TestRunCLI_Error` ✓ (existing)
3. `TestRunCLI_MissingRequiredFlags` ✓ (existing)
4. `TestRunCLI_NonAddMode` ✓ (existing)
5. `TestSaveEntry_AddsToExisting` ✓ (existing)
6. `TestSaveEntry_NewAsset` ✓ (existing)
7. `TestRunCLI_SilentSuccess` ✓ (existing)
8. `TestCLIDataStruct` ✓ (existing)

**D. Main Integration Tests (cmd/dca/main_test.go):**
1. `TestMain_CLIModeExitsEarly` ✓ (existing)
2. `TestMain_CLISavesEntry` ✓ (existing)
3. `TestMain_TUIModeUnchanged` ✓ (existing)
4. `TestMain_MultipleEntries` ✓ (existing)
5. `TestMain_SharePrecision` ✓ (existing)
6. `TestMain_DateHandling` ✓ (existing)
7. `TestMain_AutoDate` ✓ (existing)
8. `TestMain_ConcurrentEntries` ✓ (existing)
9. `TestMain_InvalidAssetFormat` ✓ (existing)
10. `TestMain_ValidationErrors` ✓ (existing)

**E. Validation Tests (internal/validation/validation_test.go) - CLI-Specific:**
1. `TestIsValidAmount_Pass` ✓ (existing)
2. `TestIsValidAmount_RejectZero` ✓ (existing)
3. `TestIsValidAmount_RejectNegative` ✓ (existing)
4. `TestIsValidAmount_RejectEmpty` ✓ (existing)
5. `TestIsValidAmount_RejectInvalid` ✓ (existing)
6. `TestIsValidAmount_ExactErrorMessage` ✓ (existing)
7. `TestIsValidAmount_Whitespace` ✓ (existing)
8. All price validation tests ✓ (existing)
9. `TestIsValidAsset_*` ✓ (existing)
10. `TestIsValidDate_*` ✓ (existing)
11. `TestRoundTo8Decimals` ✓ (existing)
12. `TestCalculateSharesFromValues` ✓ (existing)

**F. Entry Tests (internal/dca/entry_test.go) - Enhanced:**
1. Share calculation with various prices/amounts
2. Edge cases: very small prices, large amounts
3. Division by zero handling

**G. Coverage Verification:**
1. Run `go test -v ./...` - All tests pass
2. Run `go test -coverprofile=coverage.out ./...` - Check coverage
3. Add tests for any uncovered lines (if needed)

### 6. Risks and Considerations

**Potential Issues:**
1. **os.Exit handling** - Tests rely on panic/recovery which is Go testing pattern for os.Exit
2. **Global state** - `defaultEntriesPath` and `os.Args` need restoration
3. **Timing variance** - Auto-generated dates have slight variance (5-second tolerance in tests)
4. **Floating point** - 8-decimal precision: use exact equality (mathematically correct)
5. **Concurrent test execution** - Tests use unique temp directories

**Trade-offs:**
1. **Exit code testing** - Cannot capture actual exit codes, rely on panic detection
2. **Stderr capture** - Uses os.Pipe; may have limitations in some environments
3. **Test coverage** - Prioritize critical paths: validation, persistence, share calculation

**Verification Steps:**
1. Run `go test -v ./cmd/dca/ ./internal/validation/ ./internal/dca/` - All pass
2. Run `make test-cover` - Verify coverage >= 95%
3. Run `make fmt` - No formatting changes
4. Run `make build` - Build succeeds
5. Manual test: `./dca -add -asset BTC -amount 100 -price 50000` - Works correctly

### Acceptance Criteria Coverage

- [x] #1 cli_test.go created with test file (DONE - GOT-057 added cli.go)
- [x] #2 Validation tests for all flag combinations (DONE - existing tests comprehensive)
- [x] #3 Share calculation precision verified (8 decimals) (DONE - tests exist)
- [x] #4 Date auto-generation tested (DONE - tests exist)
- [x] #5 Exit codes verified for errors (DONE - tests use panic/recovery)
- [ ] #6 All existing tests still pass (VERIFY - need to run all tests)

### Definition of Done Enhancement

- [ ] #1 All acceptance criteria met (VERIFY with test run)
- [ ] #2 Unit tests pass (go test -v ./...)
- [ ] #3 No new compiler warnings (verify with make build)
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task (doc-013)
- [ ] #6 Documentation updated (comments in test functions)
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
