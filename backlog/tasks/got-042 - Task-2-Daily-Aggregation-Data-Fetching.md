---
id: GOT-042
title: 'Task 2: Daily Aggregation Data Fetching'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-18 18:51'
updated_date: '2026-03-18 20:10'
labels:
  - data
  - fetching
dependencies: []
references:
  - backlog/docs/PRD-001-asset-history-modal.md
priority: high
ordinal: 2000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement data fetching for daily asset history from dca_entries.json
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Fetches all entries for the selected asset from dca_entries.json
- [ ] #2 Groups entries by calendar date (YYYY-MM-DD)
- [ ] #3 Calculates average price per day (weighted average of entry prices)
- [ ] #4 Calculates total invested amount per day
- [ ] #5 Counts entries per day
- [ ] #6 Sorts results by date ascending
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 All acceptance criteria met
- [ ] #8 Unit tests pass
- [ ] #9 No new compiler warnings
- [ ] #10 Code follows project style (go fmt)
<!-- DOD:END -->
