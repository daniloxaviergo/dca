---
id: GOT-079
title: '[doc-013 Phase 1] GOT-075: Update README.md with CLI quick entry documentation'
status: To Do
assignee: []
created_date: '2026-03-28 17:11'
labels:
  - documentation
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update README.md with CLI Quick Entry section after the "Usage" section. Add documentation for: 1) Command syntax: ./dca --add --asset ticker --amount usd --price per-share 2) Required flags: --add (boolean), --asset, --amount, --price (all required), --date (optional, auto-set to current UTC) 3) Auto-calculated shares with 8 decimal precision (no user input needed) 4) Silent success behavior (no output on valid entry) 5) Exit codes: 0 for success, 1 for errors 6) At least 2 working command examples showing different use cases
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 README.md updated with CLI Quick Entry section after Usage section
- [ ] #2 Command syntax documented with all required flags
- [ ] #3 Auto-calculated shares mentioned with 8 decimal precision
- [ ] #4 Date auto-set behavior documented
- [ ] #5 Silent success behavior explained
- [ ] #6 Exit codes documented (0/1)
- [ ] #7 At least 2 working command examples provided
- [ ] #8 Documentation follows existing README format and style
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
