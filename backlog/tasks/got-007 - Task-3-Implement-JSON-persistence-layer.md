---
id: GOT-007
title: 'Task 3: Implement JSON persistence layer'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-16 21:24'
updated_date: '2026-03-16 23:03'
labels: []
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement file I/O for JSON data persistence. Create dca_persist.go file with atomic write support.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Load existing dca_entries.json if present (handle file not found gracefully)
- [ ] #2 Create new dca_entries.json if it doesn't exist with proper JSON structure
- [ ] #3 Store entries as map[string][]DCAEntry keyed by asset ticker
- [ ] #4 Write file with 2-space indentation for readability
- [ ] #5 Implement atomic write using temp file + rename pattern
- [ ] #6 Add entry to correct asset array in data structure
- [ ] #7 Display success message: 'Entry saved for [ASSET]' after save
- [ ] #8 Handle file permission errors with clear user message
- [ ] #9 Handle JSON marshal errors with diagnostic message
- [ ] #10 Do not corrupt existing data on write failure
<!-- AC:END -->
