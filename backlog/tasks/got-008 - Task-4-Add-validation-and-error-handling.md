---
id: GOT-008
title: 'Task 4: Add validation and error handling'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-16 21:25'
updated_date: '2026-03-16 23:25'
labels: []
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement comprehensive validation and error handling across all form inputs and file operations.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Reject negative or zero amounts with message: 'Amount must be positive'
- [ ] #2 Reject invalid date format with helpful example: 'Use YYYY-MM-DD'
- [ ] #3 Reject negative or zero prices with message: 'Price must be positive'
- [ ] #4 Reject empty asset ticker with message: 'Asset ticker is required'
- [ ] #5 Handle file permission errors with user-friendly message: 'Permission denied: check file permissions'
- [ ] #6 Handle JSON parse errors gracefully with diagnostic message
- [ ] #7 Handle file write errors with clear user message
- [ ] #8 Validate that calculated shares is a valid finite number
- [ ] #9 Catch and handle panics from numeric operations
- [ ] #10 Display validation errors inline with the prompt for re-entry
<!-- AC:END -->
