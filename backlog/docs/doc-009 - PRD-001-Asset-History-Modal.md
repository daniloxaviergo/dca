---
id: doc-009
title: 'PRD-001: Asset History Modal'
type: other
created_date: '2026-03-18 18:46'
---
# PRD: Asset History Modal

## Overview

### Purpose
Add an interactive modal that displays daily aggregated asset history when a user presses Enter on an asset row in the asset view, providing transparency into investment performance over time.

### Goals
- **Goal 1**: Display daily asset aggregation with average price, total invested, and entry count for each day
- **Goal 2**: Implement infinite scroll to load more history as needed
- **Goal 3**: Provide intuitive user experience with clear visual feedback and easy dismissal (Escape key)

## Background

### Problem Statement
Currently, the asset view shows aggregated assets by ticker but provides no visibility into the historical timeline of investments. Users cannot see:
- How daily average prices have changed over time
- The total amount invested per day
- How many entries were made on each day

### Current State
The asset view (`internal/assets/view.go`) displays:
- Asset ticker symbols
- Total shares owned
- Current price
- Weighted average entry price

No historical data beyond the current aggregate view exists in the UI.

### Proposed Solution
Add an interactive modal that:
1. Opens when Enter is pressed on an asset row
2. Displays daily aggregations with average price, total invested, and entry count
3. Supports infinite scroll to load more history as needed
4. Closes when Escape is pressed

## Requirements

### User Stories

- **End User**: 
  - *As a user, I want to press Enter on an asset row to see its daily history so that I can analyze my investment patterns over time*

### Functional Requirements

#### Task 1: Modal UI Component
Create a centered modal component for displaying asset history.

##### Acceptance Criteria
- [ ] Modal appears centered on screen when Enter is pressed on asset row
- [ ] Modal includes a title showing the asset ticker symbol
- [ ] Modal includes a header row: Date | Avg Price | Total Invested | Entry Count
- [ ] Modal closes when Escape key is pressed
- [ ] Modal includes visual indication of active state (borders, focus)

#### Task 2: Daily Aggregation Data Fetching
Implement data fetching for daily asset history.

##### Acceptance Criteria
- [ ] Fetches all entries for the selected asset from `dca_entries.json`
- [ ] Groups entries by calendar date (YYYY-MM-DD)
- [ ] Calculates average price per day (weighted average of entry prices)
- [ ] Calculates total invested amount per day
- [ ] Counts entries per day
- [ ] Sorts results by date ascending

#### Task 3: Infinite Scroll Implementation
Add infinite scroll to load historical data in batches.

##### Acceptance Criteria
- [ ] Display initial batch of 10 days of history on modal open
- [ ] Show loading state when fetching more data
- [ ] Load next batch of 10 days when user scrolls to bottom
- [ ] Disable scroll trigger once all data is loaded
- [ ] Handle empty history state gracefully

#### Task 4: Data Aggregation Logic
Implement aggregation calculations.

##### Acceptance Criteria
- [ ] Average price calculated as: SUM(price_per_share × amount) / SUM(amount)
- [ ] Total invested = SUM(amount) for the day
- [ ] Entry count = number of entries for the day
- [ ] All amounts rounded to 2 decimal places for display
- [ ] All prices rounded to 2 decimal places for display

#### Task 5: Integration with Existing Asset View
Connect modal to asset view.

##### Acceptance Criteria
- [ ] Modal opens when Enter is pressed on asset row
- [ ] Modal receives correct asset ticker as parameter
- [ ] Modal displays data for selected asset only
- [ ] Modal state is isolated from main form state

### Non-Functional Requirements

- **Performance**: Modal should open within 100ms for datasets up to 10,000 entries
- **Scalability**: Infinite scroll must handle 100+ days of history without performance degradation
- **Memory**: Store aggregated data efficiently; avoid duplicate data structures
- **UI Responsiveness**: Modal rendering should not block user input
- **Maintainability**: Code should follow existing patterns in `internal/assets/` and `internal/form/`

## Scope

### In Scope
- Modal UI component using Bubble Tea/TUI framework
- Daily aggregation logic with weighted average price calculation
- Infinite scroll pagination (10 days per batch)
- Integration with existing `dca_entries.json` data source
- Escape key to close modal
- Loading states and empty states

### Out of Scope
- Date range filtering (always shows all history)
- Export to CSV/Excel functionality
- Chart/graph visualization
- Editing or deleting entries from modal
- Real-time data updates (static snapshot)

## Technical Considerations

### Existing System Impact
- **Asset View**: Requires new modal state and event handling
- **Data Layer**: Reuses existing `LoadEntries()` from `internal/dca/file.go`
- **Aggregation**: Builds on existing `internal/assets/aggregate.go` patterns

### Dependencies
- Existing data model (`DCAEntry`, `DCAData`)
- Existing aggregation logic (`internal/assets/aggregate.go`)
- Bubble Tea framework (already in use)

### Constraints
- Modal must work within existing Bubble Tea update/render cycle
- Must use existing lipgloss styling conventions
- Data must be loaded from `dca_entries.json` (no API changes)

## Success Metrics

### Quantitative
- Modal opens in ≤100ms for datasets up to 10,000 entries
- Scroll performance maintains 60fps for 100+ days of history
- Memory usage increase ≤5MB when modal is open

### Qualitative
- Users can easily understand daily investment patterns
- Infinite scroll feels natural and responsive
- Modal dismisses smoothly without UI glitches

## Timeline & Milestones

- **Design Complete**: Modal UI mockup reviewed by team
- **Implementation Complete**: All tasks 1-5 implemented and tested
- **Testing Complete**: Manual testing of edge cases (empty history, large datasets)
- **Release**: Feature ready for production

## Stakeholders

### Decision Makers
- Product Manager: PRD approval
- Tech Lead: Technical implementation review

### Contributors
- Backend Developer: Aggregation logic
- Frontend Developer: Modal UI and event handling
- QA Engineer: Test coverage and edge case validation

## Appendix

### Glossary
- **Weighted Average Price**: The average price per share weighted by the investment amount
- **Daily Aggregation**: Grouping all entries by calendar date with calculated metrics
- **Infinite Scroll**: Loading additional data as user scrolls to bottom of list

### References
- Existing aggregation code: `internal/assets/aggregate.go`
- Data model: `dca_entry.go`
- Bubble Tea docs: https://github.com/charmbracelet/bubbletea
