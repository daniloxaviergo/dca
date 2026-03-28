---
id: GOT-072
title: '[Integration Phase 2] Task 2: Create comprehensive CLI test suite'
status: To Do
assignee: []
created_date: '2026-03-28 17:01'
labels:
  - testing
  - cli
  - quality
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create cmd/dca/cli_test.go with extensive test coverage for CLI mode functionality including flag parsing, validation, entry creation, and file I/O.

WHAT TO IMPLEMENT:

1. Create cmd/dca/cli_test.go with the following test cases:

   a) TestFlagParsingWithAllFlags - Verify all flags parse correctly (--add, --amount, --date, --asset, --price)
   b) TestMissingRequiredFlagReturnsExitCode1 - Test --add without required flags returns exit code 1
   c) TestNegativeAmountReturnsExitCode1 - Test --amount <= 0 returns exit code 1
   d) TestZeroAmountReturnsExitCode1 - Test --amount = 0 returns exit code 1
   e) TestNegativePriceReturnsExitCode1 - Test --price <= 0 returns exit code 1
   f) TestZeroPriceReturnsExitCode1 - Test --price = 0 returns exit code 1
   g) TestMissingAssetFlagReturnsExitCode1 - Test --asset="" returns exit code 1
   h) TestInvalidDateFormatReturnsExitCode1 - Test bad RFC3339 date returns exit code 1
   i) TestSuccessfulAddWithAllFlags - Test valid entry creation with exit code 0
   j) TestAutomaticDateGeneration - Test current RFC3339 date auto-set when --date missing
   k) TestShareCalculationCorrectness - Verify 8-decimal precision: amount=500, price=65000 → shares=0.00769231
   l) TestFilePersistence - Verify entry saved to dca_entries.json with correct structure
   m) TestSilentSuccess - Test no output on successful entry addition
   n) TestTUIIndependence - Verify TUI mode still works when CLI tests complete

2. Each test must verify correct exit code (0 or 1) and handle temporary file cleanup
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 TestFlagParsingWithAllFlags passes
- [ ] #2 TestMissingRequiredFlagReturnsExitCode1 passes
- [ ] #3 TestNegativeAmountReturnsExitCode1 passes
- [ ] #4 TestZeroAmountReturnsExitCode1 passes
- [ ] #5 TestNegativePriceReturnsExitCode1 passes
- [ ] #6 TestZeroPriceReturnsExitCode1 passes
- [ ] #7 TestMissingAssetFlagReturnsExitCode1 passes
- [ ] #8 TestInvalidDateFormatReturnsExitCode1 passes
- [ ] #9 TestSuccessfulAddWithAllFlags passes
- [ ] #10 TestAutomaticDateGeneration passes
- [ ] #11 TestShareCalculationCorrectness passes (500/65000=0.00769231)
- [ ] #12 TestFilePersistence passes
- [ ] #13 TestSilentSuccess passes
- [ ] #14 TestTUIIndependence passes
- [ ] #15 All tests use temp files and clean up
- [ ] #16 Tests verify exit codes 0/1
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 cmd/dca/cli_test.go created with all 14 test cases
- [ ] #8 Uses os/exec to test CLI exit codes
- [ ] #9 Uses temp directory for file I/O tests
- [ ] #10 Cleans up temp files after each test
- [ ] #11 No tests depend on external state (dca_entries.json unaffected)
- [ ] #12 All tests pass with 'go test -v ./...
<!-- DOD:END -->
