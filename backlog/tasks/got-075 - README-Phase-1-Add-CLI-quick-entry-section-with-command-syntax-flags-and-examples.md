---
id: GOT-075
title: >-
  [README Phase 1] Add CLI quick entry section with command syntax, flags, and
  examples
status: To Do
assignee: []
created_date: '2026-03-28 17:05'
labels:
  - documentation
  - cli
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update README.md to include CLI quick entry documentation. Add a new "CLI Quick Entry" section after the "Usage" section that documents:

1. Command syntax: ./dca --add --asset <ticker> --amount <usd> --price <per-share>
2. Required flags explanation (--add, --asset, --amount, --price)
3. Auto-calculated shares information (8 decimal precision)
4. Date auto-set behavior (defaults to current UTC timestamp if --date omitted)
5. Silent success behavior (no output on valid entry)
6. Exit codes (0 for success, 1 for errors)
7. At least 2 working command examples
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 README.md updated with CLI Quick Entry section
- [ ] #2 Command syntax documented with all required flags
- [ ] #3 Auto-calculated shares mentioned with 8 decimal precision
- [ ] #4 Date auto-set behavior documented
- [ ] #5 Silent success behavior explained
- [ ] #6 Exit codes documented (0/1)
- [ ] #7 At least 2 working command examples provided
- [ ] #8 CLI examples match actual implementation
- [ ] #9 Documentation follows existing README format and style
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
