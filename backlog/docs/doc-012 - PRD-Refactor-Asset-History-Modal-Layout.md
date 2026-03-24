---
id: doc-012
title: 'PRD: Refactor Asset History Modal Layout'
type: other
created_date: '2026-03-24 20:37'
---
# PRD: Refactor Asset History Modal Layout

## Overview

### Purpose
Redesign the asset history modal to improve data display with modern UI principles, implementing sorted date display (newest first) and enhanced visual hierarchy within Bubble Tea constraints.

### Goals
- **Goal 1**: Improve information density and readability by 30% through better column alignment and spacing
- **Goal 2**: Implement intuitive date sorting (newest entries first) for quick access to recent investment data
- **Goal 3**: Create a modern, polished UI that maintains consistency with the application's visual identity

## Background

### Problem Statement
The current asset history modal displays investment data with suboptimal visual organization, making it difficult for users to quickly understand their investment patterns. The modal lacks clear visual hierarchy and the date sorting is inconsistent or missing, requiring users to mentally process data rather than having it presented logically.

### Current State
- Modal uses basic lipgloss styling with minimal visual enhancements
- Column widths may not optimize screen real estate
- Date sorting behavior is unclear or inconsistent
- No visual distinction between header, data rows, and interactive elements
- Pagination feedback is minimal
- Modal styling lacks modern UI principles (shadows, rounded corners, consistent spacing)

### Proposed Solution
Refactor the modal layout to:
1. Implement modern UI styling with enhanced borders, colors, and visual hierarchy
2. Ensure date sorting follows a consistent pattern (newest first by default)
3. Optimize column widths for better information density
4. Add visual feedback for interactive states and loading
5. Maintain compatibility with Bubble Tea/TUI constraints

## Requirements

### User Stories

- **Investor**: *As a DCA investor, I want to see my asset history with the most recent entries first so that I can quickly assess my latest investments*
- **Investor**: *As a user, I want the modal to visually distinguish the header from data rows so that I can understand the data structure at a glance*
- **Investor**: *As a user, I want clear pagination indicators so that I know when more data is loading or when all data has been loaded*

### Functional Requirements

#### Task 1: Implement Modern Modal Styling

Refactor the modal's visual appearance using modern UI principles while staying within Bubble Tea's capabilities.

##### User Flows
1. User opens asset history modal by selecting an asset from the main list
2. Modal appears with enhanced visual styling (modern border, colors, alignment)
3. User navigates through entries with clear visual feedback
4. Loading states are indicated with appropriate text

##### Acceptance Criteria
- [ ] Modal uses rounded border with modern foreground color for header (existing: color 63)
- [ ] Header row uses bold text with high-contrast color (existing: color 15)
- [ ] Data rows have consistent padding and alignment
- [ ] Loading state displays "Loading more..." text with appropriate styling
- [ ] Empty state displays "No history for this asset" clearly
- [ ] Error state displays error message with warning color (existing: color 196)

#### Task 2: Implement Consistent Date Sorting (Newest First)

Ensure date entries are always sorted with newest first for predictable user experience.

##### User Flows
1. User opens modal for any asset
2. Entries appear sorted by date descending (most recent first)
3. Load More button adds older entries below current ones
4. All data is eventually displayed with chronological order maintained

##### Acceptance Criteria
- [ ] Date column shows dates in YYYY-MM-DD format
- [ ] Entries are sorted with newest date first (descending order)
- [ ] When "Load More" is pressed, older entries are appended to the bottom
- [ ] No duplicate dates appear in the list
- [ ] Date sorting works correctly with empty or single-entry datasets

#### Task 3: Optimize Column Widths and Spacing

Adjust modal column dimensions to improve information density without sacrificing readability.

##### User Flows
1. User views modal on standard terminal width (60 characters)
2. All columns are visible without wrapping
3. Columns are aligned and easy to scan
4. No horizontal scrolling is required

##### Acceptance Criteria
- [ ] Modal width is 60 characters (existing constant maintained)
- [ ] Date column is wide enough for YYYY-MM-DD format (12 chars)
- [ ] Price and value columns align on decimal points
- [ ] Entry count column is wide enough for at least 4-digit numbers
- [ ] Column spacing is consistent (2 spaces between columns)

#### Task 4: Enhance Visual Feedback

Add visual indicators for interactive elements and loading states.

##### User Flows
1. User sees clear instructions at the bottom of the modal
2. Loading state displays visual feedback
3. "All data loaded" state is clearly indicated
4. Modal can be closed with visible ESC instructions

##### Acceptance Criteria
- [ ] Footer displays clear navigation instructions: "[Esc] Close Modal" and "[Enter] Load More"
- [ ] When all data is loaded, footer shows "All data loaded" instead of "Load More"
- [ ] When loading, footer shows "Loading more..." to indicate activity
- [ ] Footer uses muted color (existing: color 240)

### Non-Functional Requirements

- **Performance**: Modal should load and display first 10 entries within 100ms
- **Compatibility**: Must work within Bubble Tea framework with lipgloss styling
- **Accessibility**: Text contrast must meet terminal readability standards
- **Maintainability**: Code should use constants for width values, not magic numbers
- **Scalability**: Modal should handle up to 1000+ daily entries without performance degradation

## Scope

### In Scope
- Modern modal styling with enhanced visual hierarchy
- Consistent newest-first date sorting
- Column width optimization for 60-char modal
- Visual feedback for loading, empty, and error states
- Footer instructions and state indicators

### Out of Scope
- Adding new data columns to the modal
- Implementing filtering or search functionality
- Changing the pagination batch size (remains at 10 entries)
- Modifying the main assets list view

## Technical Considerations

### Existing System Impact
- Modal layout changes only affect the `AssetHistoryModal` component
- Data aggregation logic (`AggregateByDate`) already handles newest-first sorting
- No changes to data models or JSON file format required
- Main assets list view remains unchanged

### Dependencies
- Bubble Tea v1.3.10 (TUI framework)
- Lipgloss v1.1.0 (styling)
- No external dependencies required

### Constraints
- Modal width is fixed at 60 characters
- Must work within standard terminal sizes
- No mouse input support (keyboard only)
- Color palette limited to standard ANSI colors

## Success Metrics

### Quantitative
- Modal loads and displays first 10 entries in <100ms
- No horizontal scrolling required on 60-char modal
- Date sorting is consistent across all assets

### Qualitative
- Users can identify the most recent entries at a glance
- Visual hierarchy makes it clear which row is the header
- Loading states provide clear feedback
- Overall modal appearance feels modern and polished

## Timeline & Milestones

### Key Dates
- T+1 day: Design complete - modal layout mockup finalized
- T+3 days: Implementation complete - styling, sorting, and feedback enhancements
- T+4 days: Testing complete - verify sorting, loading, and error states
- T+5 days: Launch/Release - merged to main branch

## Stakeholders

### Decision Makers
- Product Owner: Approve PRD scope and timeline

### Contributors
- Developer: Implement modal refactoring
- QA: Test sorting behavior, loading states, and visual consistency

## Appendix

### Glossary
- **DCA**: Dollar-Cost Averaging - investment strategy of regularly investing fixed amounts
- **Modal**: Overlay window displaying detailed asset history
- **TUI**: Terminal User Interface - command-line interface with interactive elements

### References
- [QWEN.md](QWEN.md): Project context and architecture documentation
- [Bubble Tea Documentation](https://bubbletea.chat): TUI framework reference
- [Lipgloss Documentation](https://github.com/charmbracelet/lipgloss): Terminal styling reference
