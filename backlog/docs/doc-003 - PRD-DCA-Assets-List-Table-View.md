---
id: doc-003
title: 'PRD: DCA Assets List Table View'
type: other
created_date: '2026-03-17 00:36'
---
# PRD: DCA Assets List Table View

## Overview

### Purpose
Create an interactive terminal table view that aggregates DCA entries by asset, displaying key metrics including average entry price and total shares to help users quickly understand their investment positions.

### Goals
- **Goal 1**: Display all DCA entries grouped by asset ticker in a clean, sortable table format
- **Goal 2**: Calculate and show average entry price and sum of shares per asset
- **Goal 3**: Provide interactive navigation through the list using arrow keys
- **Goal 4**: Maintain consistency with existing Bubble Tea TUI design patterns

## Background

### Problem Statement
Currently, users can enter DCA investment data, but there's no view to quickly see an aggregated summary of all assets. Users must manually track or export data to understand their total position per asset.

### Current State
- The application has `dca_entry.go` with `DCAData` structure storing entries as a map of asset ticker to array of entries
- Users can add new entries via the interactive form (`dca_form.go`)
- Entries are persisted to `dca_entries.json`
- No view exists to summarize entries by asset with calculated metrics

### Proposed Solution
Add a new "Assets" view mode to the Bubble Tea application that:
1. Reads and aggregates entries from `dca_entries.json`
2. Groups entries by asset ticker
3. Calculates average entry price and total shares per asset
4. Displays the data in an interactive terminal table using Lipgloss styling
5. Supports navigation with arrow keys and exit with Esc/Ctrl+C

## Requirements

### User Stories

- **End User**: As a DCA investor, I want to see my assets in a summary table so that I can quickly understand my total position per asset, including average entry price and total shares owned.

- **End User**: As a user, I want to navigate the assets list with keyboard controls so that I can review my investments efficiently without mouse interaction.

### Functional Requirements

#### Task 1: Assets View Model and Data Aggregation
Create a new `assets_view.go` file with the `AssetsViewModel` that reads entries and aggregates them by asset.

**Description:**
- Create `AssetsViewModel` struct to manage the assets list state
- Implement `LoadAndAggregateEntries()` function to read from `dca_entries.json` and group entries by asset
- Calculate per-asset metrics: count of entries, sum of shares, average entry price
- Add validation to handle empty or malformed data

**User Flows**
1. User navigates to Assets view
2. System reads `dca_entries.json` and aggregates entries
3. System calculates metrics for each asset
4. System displays the table

**Acceptance Criteria**
- [ ] System loads entries from `dca_entries.json` without errors
- [ ] Entries are correctly grouped by asset ticker
- [ ] Sum of shares is calculated correctly per asset
- [ ] Average entry price is calculated correctly per asset (weighted average)
- [ ] Empty entries file displays empty table gracefully
- [ ] Malformed JSON is handled with user-friendly error message

#### Task 2: Interactive Table UI Component
Create the table UI using Bubble Tea and Lipgloss.

**Description:**
- Implement `AssetsView` component with Bubble Tea model pattern
- Create table rendering with headers: Asset, Count, Total Shares, Avg Price, Total Value
- Add row selection highlighting (blue background for active row)
- Implement keyboard navigation (Up/Down arrows to move, Enter to select)
- Implement exit navigation (Esc or Ctrl+C to return/quit)

**User Flows**
1. User presses 'a' or navigates to Assets menu option
2. System displays the assets table
3. User uses Up/Down arrows to navigate rows
4. User can press Esc to return to main menu or quit

**Acceptance Criteria**
- [ ] Table displays with proper headers and column alignment
- [ ] Up/Down arrow keys navigate between rows
- [ ] Active row is highlighted with different styling
- [ ] Esc key returns to previous menu or exits application
- [ ] Ctrl+C exits application cleanly
- [ ] Table handles zero assets gracefully (shows "No assets yet" message)

#### Task 3: Integration with Main Application
Integrate the assets view into the existing application flow.

**Description:**
- Add "Assets" option to main menu (if menu exists) or add command-line subcommand
- Add state transition from main view to assets view
- Handle data persistence consistency between views
- Ensure proper cleanup on exit

**User Flows**
1. User runs `./dca` or `go run main.go`
2. User navigates to Assets view
3. User views assets table
4. User exits back to main flow or application

**Acceptance Criteria**
- [ ] Assets view is accessible from main entry point
- [ ] Data remains consistent when switching between views
- [ ] Changes made in form view are reflected in assets view after save
- [ ] Application exits cleanly from assets view

### Non-Functional Requirements

- **Performance**: Load and aggregate 1000+ entries within 500ms
- **Compatibility**: Work with existing Bubble Tea v1.3.10 and Lipgloss v1.1.0
- **Maintainability**: Follow existing code patterns in `dca_form.go` and `dca_entry.go`
- **Error Handling**: Gracefully handle missing file, permission errors, and malformed JSON
- **UI Consistency**: Use same color scheme and styling as existing forms (blue for active field)

## Scope

### In Scope
- Assets table view displaying aggregated entry data
- Keyboard navigation (Up/Down arrows, Esc, Enter)
- Average entry price calculation (weighted by shares)
- Total shares and entry count per asset
- Sum of entry amounts as "Total Value"
- Interactive terminal UI using existing Bubble Tea framework

### Out of Scope
- External price API integration for current market prices
- Export to CSV/Excel functionality
- Sorting columns by clicking headers
- Filtering/sorting assets alphabetically or by value
- Adding new entries from the assets view
- Portfolio value in different currencies
- Historical price charts

## Technical Considerations

### Existing System Impact
- New file `assets_view.go` to be created alongside `dca_entry.go`, `dca_form.go`, `main.go`
- No changes to `DCAEntry` or `DCAData` structures required
- Main application flow needs to support multiple views

### Dependencies
- Existing: Bubble Tea v1.3.10 for TUI framework
- Existing: Lipgloss v1.1.0 for styling
- No new external dependencies required

### Constraints
- Must use existing `dca_entries.json` file format
- Must maintain single-file-per-view architecture
- Keyboard navigation must be intuitive for terminal users

## Success Metrics

### Quantitative
- Table renders 100 assets within 500ms
- Navigation response time < 100ms per keypress
- Zero entries displays within 200ms

### Qualitative
- Table is visually distinct from form view
- Navigation feels smooth and responsive
- Metrics are clearly labeled and easy to understand

## Timeline & Milestones

### Key Dates
- [Date]: Task 1 (Data Aggregation) complete
- [Date]: Task 2 (Table UI) complete
- [Date]: Task 3 (Integration) complete
- [Date]: Testing and bug fixes
- [Date]: Code review and merge

## Stakeholders

### Decision Makers
- Product Manager: Defines feature scope and acceptance criteria

### Contributors
- Backend/Go Developer: Implements data aggregation and TUI integration
- QA/Tester: Validates functionality and edge cases

## Appendix

### Glossary
- **DCA**: Dollar-Cost Averaging - investment strategy of buying fixed dollar amount at regular intervals
- **Avg Price**: Average entry price per share (weighted by shares purchased)
- **Total Shares**: Sum of all shares purchased for an asset
- **Total Value**: Sum of all entry amounts (USD invested)

### References
- `dca_entry.go`: Current data model and file I/O
- `dca_form.go`: Bubble Tea form implementation patterns
- `main.go`: Application entry point structure

### Acceptance Criteria Summary

#### Task 1 Acceptance Criteria
- [ ] Entries loaded from `dca_entries.json` correctly
- [ ] Grouping by asset ticker works
- [ ] Sum of shares calculated per asset
- [ ] Weighted average entry price calculated correctly
- [ ] Empty file handled gracefully

#### Task 2 Acceptance Criteria
- [ ] Table displays with headers: Asset, Count, Total Shares, Avg Price, Total Value
- [ ] Up/Down arrows navigate rows
- [ ] Active row highlighted
- [ ] Esc returns to menu or exits
- [ ] Ctrl+C exits cleanly
- [ ] No assets message displays when list is empty

#### Task 3 Acceptance Criteria
- [ ] Assets view accessible from main entry
- [ ] Data consistency maintained across views
- [ ] Changes reflected after save
- [ ] Clean exit from assets view
