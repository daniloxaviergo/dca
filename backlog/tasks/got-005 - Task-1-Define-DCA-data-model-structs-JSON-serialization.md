---
id: GOT-005
title: 'Task 1: Define DCA data model (structs, JSON serialization)'
status: To Do
assignee: []
created_date: '2026-03-16 21:24'
labels:
  - data-model
  - core
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define Go structs for DCA entries and implement JSON serialization/deserialization. Create dca_entry.go file with proper data model.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Define DCAEntry struct with fields: Amount (float64), Date (time.Time), Asset (string), PricePerShare (float64), Shares (float64)
- [ ] #2 Define DCAData struct with map[string][]DCAEntry keyed by asset ticker
- [ ] #3 Implement LoadEntries() function to read from dca_entries.json
- [ ] #4 Implement SaveEntries() function to write to dca_entries.json with 2-space indentation
- [ ] #5 Add Validate() method on DCAEntry to check Amount > 0, PricePerShare > 0
- [ ] #6 Add CalculateShares() method to compute Shares = Amount / PricePerShare with 8 decimal precision
<!-- AC:END -->
