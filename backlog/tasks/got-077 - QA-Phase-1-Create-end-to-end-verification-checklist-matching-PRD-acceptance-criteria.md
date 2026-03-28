---
id: GOT-077
title: >-
  [QA Phase 1] Create end-to-end verification checklist matching PRD acceptance
  criteria
status: To Do
assignee: []
created_date: '2026-03-28 17:05'
labels:
  - quality
  - testing
  - documentation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create a comprehensive end-to-end verification checklist that maps to all PRD doc-013 acceptance criteria:

**Functional Checks (8 items):**
1. ✅ CLI adds entry to data file with correct structure
2. ✅ Auto-calculated shares correct (8 decimals): amount=500, price=65000 → shares=0.00769231
3. ✅ Date auto-set to current RFC3339 format
4. ✅ Missing required flag returns exit code 1
5. ✅ Negative/zero amount returns exit code 1
6. ✅ Negative/zero price returns exit code 1
7. ✅ Successful entry produces no output (silent success)
8. ✅ TUI continues to work unchanged (backward compatibility)

**Non-Functional Checks (3 items):**
1. ✅ CLI operation completes in < 100ms (quick response)
2. ✅ Existing tests pass unchanged (no regressions)
3. ✅ Data consistency maintained (same format across modes)

**Verification Steps:**
1. Test CLI with valid inputs → verify JSON updated correctly
2. Test CLI with invalid inputs → verify exit codes
3. Verify TUI still works (start without --add flag)
4. Run make check → verify no warnings or failures
5. Run go test -v ./... → verify all tests pass
6. Check data file format consistency
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 End-to-end verification checklist created
- [ ] #2 Functional checks: 8/8 mapped from PRD
- [ ] #3 Non-functional checks: 3/3 mapped from PRD
- [ ] #4 Verification steps documented with specific commands
- [ ] #5 Exit code testing documented (0 vs 1)
- [ ] #6 Data file format consistency verified
- [ ] #7 TUI backward compatibility tested
- [ ] #8 make check all steps documented
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
