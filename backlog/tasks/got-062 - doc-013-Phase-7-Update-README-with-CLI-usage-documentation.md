---
id: GOT-062
title: '[doc-013 Phase 7] Update README with CLI usage documentation'
status: To Do
assignee: []
created_date: '2026-03-28 17:46'
labels:
  - documentation
  - cli
dependencies: []
references:
  - README.md
  - PRD Summary
  - User Documentation
documentation:
  - doc-013
priority: low
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update README.md with a new section describing CLI quick entry functionality. Include the command format (./dca --add --asset <ticker> --amount <usd> --price <per-share>), examples of valid commands, description of each flag, and behavior notes (silent success, auto-calculated shares, auto-set date). Maintain consistency with existing documentation style and formatting.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 README.md updated with CLI quick entry section
- [ ] #2 Command format documented with example usage
- [ ] #3 Flag descriptions complete
- [ ] #4 Behaviors explained (silent success, share calculation, date auto-assignment)
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
