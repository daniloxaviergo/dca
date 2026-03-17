---
id: GOT-027
title: 'Task 2: Fix Row Value Alignment'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 20:22'
updated_date: '2026-03-17 20:59'
labels: []
dependencies: []
references:
  - backlog/docs/doc-006.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Fix row value formatting to ensure all values align with their column headers using fmt.Sprintf with width specifiers.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Header row values are left-aligned for text columns (Asset)
- [ ] #2 Numeric columns are right-aligned
- [ ] #3 All column values use fixed-width formatting with fmt.Sprintf
- [ ] #4 Row values match header column width exactly
- [ ] #5 Visual inspection confirms column alignment
- [ ] #6 Unit tests verify alignment
- [ ] #7 go fmt applied
- [ ] #8 go build succeeds
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
