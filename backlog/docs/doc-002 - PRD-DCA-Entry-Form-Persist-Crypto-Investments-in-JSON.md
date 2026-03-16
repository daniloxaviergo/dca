---
id: doc-002
title: 'PRD: DCA Entry Form - Persist Crypto Investments in JSON'
type: other
created_date: '2026-03-16 21:16'
---
# PRD: DCA Entry Form - Persist Crypto Investments in JSON

## Overview

### Purpose
Create a CLI-based data entry form to capture and persist DCA (Dollar-Cost Averaging) investment entries in JSON format, organized by cryptocurrency asset, to support portfolio tracking and analysis.

### Goals
- **AC1**: Enable users to input full DCA entry details (amount, date, asset, price, shares) via CLI
- **AC2**: Persist entries in a JSON file with structure indexed by asset ticker
- **AC3**: Support validation, error handling, and user-friendly input prompts
- **AC4**: Enable future analysis of DCA performance metrics

## Background

### Problem Statement
Currently, the DCA application is a visual placeholder with no data persistence. Users cannot record their DCA investments or track their portfolio over time. Without a way to capture investment data, the mathematical benefits of DCA (volatility capture, compound returns) cannot be measured or analyzed.

### Current State
- Existing `main.go` provides only visual UI with no data model
- No data persistence layer exists
- No input mechanisms for user DCA entries
- No file-based storage for investment records

### Proposed Solution
Build a CLI-based form interface that:
1. Prompts users for DCA entry details interactively
2. Validates input data (date format, positive values, valid asset tickers)
3. Persists entries to a JSON file structured by asset ticker
4. Provides feedback on success/failure

## Requirements

### User Stories

- **As a DCA investor**, I want to enter my investment details so that I can track my portfolio over time
  - *I want to input: investment amount, date, asset ticker, price per share, number of shares*
  
- **As a portfolio analyzer**, I want my data stored in a structured format so that I can analyze my DCA performance
  - *I want entries organized by asset ticker for easy lookup*

- **As a daily investor**, I want clear input prompts and validation so that I can enter data quickly without errors
  - *I want helpful error messages and guidance when inputs are invalid*

### Functional Requirements

#### Task 1: Data Model Definition

Define Go structs to represent DCA entries and the data store.

##### Acceptance Criteria
- [ ] Define `DCEntry` struct with fields: Amount, Date, Asset, PricePerShare, Shares
- [ ] Define `DCAData` struct with map keyed by asset ticker
- [ ] Implement JSON serialization/deserialization methods
- [ ] Validate that Amount > 0, PricePerShare > 0, Shares > 0
- [ ] Validate Date is in RFC3339 format

#### Task 2: CLI Input Form

Implement interactive CLI prompts for DCA entry input.

##### User Flows
1. User runs `go run main.go`
2. Application displays "Enter DCA Entry" prompt
3. User enters amount (e.g., "500")
4. System prompts for date (default: today)
5. User enters asset ticker (e.g., "BTC")
6. User enters price per share (e.g., "65000.00")
7. System calculates and displays shares (Amount / Price)
8. User confirms entry or cancels

##### Acceptance Criteria
- [ ] Display prompts in order: amount, date, asset, price, shares
- [ ] Support default date (today) if user presses Enter
- [ ] Calculate shares from amount/price with 8 decimal precision
- [ ] Allow user to confirm or cancel before saving
- [ ] Handle Ctrl+C gracefully (exit without saving)

#### Task 3: JSON Persistence

Implement file I/O for JSON data persistence.

##### User Flows
1. User confirms entry after input
2. Application loads existing JSON file (or creates new)
3. Entry is added to the asset's array in the data structure
4. File is written with proper indentation
5. Success message displays: "Entry saved for BTC"

##### Acceptance Criteria
- [ ] Load existing `dca_entries.json` if present
- [ ] Create new file if it doesn't exist
- [ ] Structure: `{"BTC": [...entries...], "ETH": [...entries...]}` 
- [ ] Write with 2-space indentation for readability
- [ ] Handle file write errors gracefully with clear messages

#### Task 4: Error Handling & Validation

Implement robust validation and error reporting.

##### Acceptance Criteria
- [ ] Reject negative or zero amounts with message: "Amount must be positive"
- [ ] Reject invalid date format with helpful example: "Use YYYY-MM-DD"
- [ ] Reject negative prices with message: "Price must be positive"
- [ ] Handle file permission errors with user-friendly message
- [ ] Handle JSON parse errors gracefully

### Non-Functional Requirements

- **Performance**: Form should load and prompt within 1 second
- **Data Integrity**: JSON file must not be corrupted on write (use atomic write)
- **Compatibility**: Works on Linux, macOS, Windows with Go 1.25+
- **Maintainability**: Code follows Go idioms; structs are properly documented
- **Usability**: Clear prompts, default values, confirmation before save

## Scope

### In Scope
- CLI form with interactive prompts using BubbleTea or standard library
- Data model with proper struct definitions
- JSON file persistence (read/write atomic)
- Input validation with user-friendly error messages
- Entry confirmation before save
- `dca_entries.json` file in project root

### Out of Scope
- Web UI or web form interface
- Import/export to other formats (CSV, XML)
- Multiple file support or database integration
- DCA analytics or visualization (deferred to future PRD)
- User authentication or multi-user support

## Technical Considerations

### Existing System Impact
- Current `main.go` visual UI will be replaced/extended with data entry functionality
- No breaking changes to existing package structure
- New file: `dca_entries.json` (created by application)

### Dependencies
- Current dependencies: `github.com/charmbracelet/bubbletea`, `github.com/charmbracelet/lipgloss`
- May need `github.com/AlecAivazis/survey` for better CLI form UX (if BubbleTea proves too complex)

### Constraints
- Must use Go (project standard)
- Data must persist in human-readable JSON
- CLI-only interface (no GUI)

### Future Enhancements
- Support for recurring DCA schedules (weekly, monthly)
- Portfolio summary dashboard
- Export to CSV/Excel
- Web interface option
- API endpoints for programmatic access

## Success Metrics

### Quantitative
- Input completion rate: >95% of prompts entered successfully
- Form completion time: <30 seconds per entry
- No data loss on save: 100% of saves complete successfully

### Qualitative
- Users can complete an entry without consulting documentation
- Error messages are clear and actionable
- JSON file is readable and editable by hand

## Timeline & Milestones

- **Design**: Define data model and CLI flow (1 day)
- **Implementation**: Task 1-4 (2-3 days)
- **Testing**: Unit tests, manual testing (1 day)
- **Documentation**: README update (0.5 day)
- **Total**: 4-5 days

## Stakeholders

### Decision Makers
- Danilo (Project Owner)

### Contributors
- AI Agent (Implementation)
- Danilo (Testing/Feedback)

## Appendix

### Glossary
- **DCA**: Dollar-Cost Averaging - investment strategy of investing fixed amount at regular intervals
- **Ticker**: Symbol representing a cryptocurrency (e.g., BTC, ETH, SOL)
- **RFC3339**: Date format standard (YYYY-MM-DDTHH:MM:SSZ)

### References
- [BubbleTea Docs](https://github.com/charmbracelet/bubbletea): TUI framework in use
- [Go JSON Package](https://pkg.go.dev/encoding/json): Serialization documentation
- [DCA Investment Strategy](https://www.investopedia.com/dollar-cost-averaging-4695475): Mathematical principles

### Example JSON Structure
```json
{
  "BTC": [
    {
      "Amount": 500.00,
      "Date": "2026-03-15T00:00:00Z",
      "Asset": "BTC",
      "PricePerShare": 65000.00,
      "Shares": 0.00769231
    }
  ],
  "ETH": [
    {
      "Amount": 200.00,
      "Date": "2026-03-10T00:00:00Z",
      "Asset": "ETH",
      "PricePerShare": 3200.00,
      "Shares": 0.0625
    }
  ]
}
```
