---
id: GOT-061
title: '[doc-013 Phase 6] Task 6: Update README.md with CLI usage documentation'
status: To Do
assignee: []
created_date: '2026-03-28 17:38'
labels:
  - documentation
  - cli
  - user-facing
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update README.md to document CLI usage section: add --add flag documentation, example command format, field descriptions, validation rules, and exit codes. Include command examples for adding entries and note silent success behavior. Ensure existing TUI documentation remains unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 CLI section added with command format
- [ ] #2 --add, --asset, --amount, --price flags documented
- [ ] #3 Example commands provided
- [ ] #4 Exit codes documented (0 for success, 1 for errors)
- [ ] #5 TUI section not modified
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
