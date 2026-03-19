---
id: GOT-050
title: TestAssetsView_UpdateLoadMore (FAIL)
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-19 09:43'
updated_date: '2026-03-19 09:43'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Error: view_test.go:1128: Expected 20 entries after LoadMore, got 10

Analysis: The test expects LoadMore to load the next batch of 10 days (from 10 to 20), but it's not loading any additional data. The issue is likely in the handleLoadMore method which uses a hardcoded filename dca_entries.json, but the test scenario uses a modal that was pre-populated with 10 entries. The test doesn't actually have a file with more data for the "BTC" asset.

Root Cause: The test creates a modal with 10 entries in memory, then calls Update(LoadMoreMsg{}) which triggers handleLoadMore() using dca_entries.json. However, dca_entries.json may not have additional BTC entries beyond what's already in the modal, OR the modal's state isn't properly initialized with all the data that should be in the file.
<!-- SECTION:DESCRIPTION:END -->
