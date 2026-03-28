---
id: GOT-061
title: '[doc-013 Phase 6] Update README with CLI usage documentation'
status: To Do
assignee: []
created_date: '2026-03-28 20:50'
labels:
  - documentation
  - cli
dependencies: []
references:
  - 'doc-013 - Phase 6: Update README'
documentation:
  - doc-013
priority: low
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update README.md with a new section documenting the CLI quick entry feature. Include command syntax (./dca --add --asset <ticker> --amount <usd> --price <per-share>), required and optional flags, examples, and behavior notes (silent success, exit codes, time.Now() default). Ensure documentation aligns with PRD specifications and maintains project style.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 README.md updated with CLI quick entry section
- [ ] #2 Command syntax documented
- [ ] #3 All flags explained (required/optional)
- [ ] #4 Example usage provided
- [ ] #5 Behavior notes included (silent success, exit codes)
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
