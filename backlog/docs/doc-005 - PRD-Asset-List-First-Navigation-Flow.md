---
id: doc-005
title: 'PRD: Asset List First Navigation Flow'
type: other
created_date: '2026-03-17 17:20'
---
# PRD: Asset List First Navigation Flow

## Overview

### Purpose
Change the application's initial state so it opens with the asset list view displayed, and allow users to press 'c' to create a new entry. This prioritizes viewing existing investments over adding new ones.

### Goals
- [G1] User can view all assets and their aggregated data immediately upon app launch
- [G2] User can create a new investment entry by pressing 'c' from the asset list
- [G3] Smooth navigation between asset list and entry form without state loss

## Background

### Problem Statement
The current application opens directly in the entry form, requiring users to complete an entry or cancel before viewing their assets. This workflow is inefficient for users who primarily want to review their portfolio.

### Current State
1. App launches with form view (blank entry form)
2. User must complete or cancel form to access assets
3. No direct path to view assets without entering data
4. Users must press Esc/Ctrl+C to exit form, then restart to see assets

### Proposed Solution
- App starts in asset list view
- 'c' key opens entry form
- After form submission, return to asset list
- Maintains existing navigation for exiting (Esc/Ctrl+C)

## Requirements

### User Stories

- **Investor**: *As an investor, I want to see all my assets and their current status immediately when I open the app so that I can quickly review my portfolio.*

- **Investor**: *As an investor, I want to create a new entry by pressing 'c' from the asset list so that I can easily add new investments without exiting the app.*

### Functional Requirements

#### Task 1: App starts in asset list view

The application should initialize in the assets view state, displaying all aggregated asset data.

##### User Flows
1. User runs the application (`./dca`)
2. System displays asset list table with header, data rows, and footer
3. User sees aggregated data: ticker, count, total shares, avg price, total value
4. User can navigate with arrow keys or exit with Esc/Ctrl+C

##### Acceptance Criteria
- [ ] App initializes with `StateAssetsView` instead of `StateForm`
- [ ] Asset list displays on first render
- [ ] User can navigate with ↑/↓ arrow keys
- [ ] Esc/Ctrl+C exits the app from asset list

#### Task 2: Press 'c' to create new entry

Users can press 'c' from the asset list to switch to the entry form.

##### User Flows
1. User is viewing the asset list
2. User presses 'c' key
3. System switches to form view with all fields reset
4. Form shows current timestamp as default date
5. User can fill in entry details

##### Acceptance Criteria
- [ ] Pressing 'c' in asset list switches to form view
- [ ] Form fields are reset (amount=empty, asset=empty, price=empty)
- [ ] Date defaults to current timestamp in RFC3339 format
- [ ] User can navigate form with Tab/Enter as before

#### Task 3: Return to asset list after entry

After submitting a new entry, the application should return to the asset list view with updated data.

##### User Flows
1. User completes entry form and confirms
2. System saves entry to `dca_entries.json`
3. System switches back to asset list view
4. Asset list refreshes to show updated aggregated data
5. User can continue viewing or create another entry

##### Acceptance Criteria
- [ ] After form submission, app switches to asset list view
- [ ] Asset data refreshes to include new entry
- [ ] Aggregation calculations update correctly
- [ ] User can navigate asset list or create another entry

#### Task 4: Exit from asset list

Users can exit the application from the asset list view.

##### User Flows
1. User is viewing asset list
2. User presses Esc or Ctrl+C
3. Application closes gracefully

##### Acceptance Criteria
- [ ] Esc key exits application
- [ ] Ctrl+C exits application
- [ ] No unsaved data loss (entries saved on form submit only)

### Non-Functional Requirements

- **Performance**: Asset list should load and display within 200ms for datasets up to 1000 entries
- **UI Consistency**: Navigation and styling should match existing form design (lipgloss, rounded borders)
- **Error Handling**: If data file is corrupted, display error message in asset list view
- **Code Quality**: Follow existing project conventions (`go fmt`, test coverage)

## Scope

### In Scope
- Asset list as initial state on app launch
- 'c' key to switch from asset list to entry form
- Return to asset list after form submission
- Asset list updates after new entry is saved

### Out of Scope
- Search/filter functionality for assets
- Edit existing entries from asset list
- Delete entries from asset list
- Export assets to CSV or other formats

## Technical Considerations

### Existing System Impact
- **State Management**: Change `AppState` enum to prioritize `StateAssetsView`
- **Model Initialization**: Swap initial state from form to assets view
- **Entry Saving**: Form submission must return to assets view and refresh data

### Dependencies
- Existing `assets.AssetsView` component (already implemented)
- Existing `form.FormModel` (needs 'c' key handler)
- Existing data loading/aggregation logic (`internal/assets/aggregate.go`)

### Constraints
- Must maintain existing form validation rules
- Must use atomic file writes for data safety
- Must preserve existing keyboard navigation patterns

## Success Metrics

### Quantitative
- Asset list loads in < 200ms for 1000 entries
- Key response time < 50ms from asset list to form
- 100% data persistence success rate

### Qualitative
- Users can view assets before entering any data
- Intuitive 'c' key for "create" action
- Seamless transition between views

## Timeline & Milestones

### Key Dates
- [Date]: Design complete - State transition logic reviewed
- [Date]: Implementation complete - All tasks tested
- [Date]: Testing complete - Integration tests pass
- [Date]: Release - Deploy to users

## Stakeholders

### Decision Makers
- [Name]: Product Manager (PRD approval)

### Contributors
- [Name]: Developer (Implementation)
- [Name]: QA (Testing)

## Appendix

### Glossary
- **Asset**: A unique ticker symbol (e.g., BTC, ETH) with aggregated investment entries
- **Asset Summary**: Aggregated data including count, total shares, avg price, total value
- **Form View**: Entry form for adding new DCA investments
- **Asset List View**: Table displaying aggregated asset data

### References
- [QWEN.md](../QWEN.md): Project context and documentation
- [internal/assets/view.go](../internal/assets/view.go): AssetsView component
- [internal/form/model.go](../internal/form/model.go): FormModel implementation
- [cmd/dca/main.go](../cmd/dca/main.go): Main application entry point
