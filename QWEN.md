# DCA Project Context

## Project Overview

This is a **Go command-line application** for tracking Dollar-Cost Averaging (DCA) investment entries. The application provides an interactive terminal UI built with Bubble Tea for entering investment data (amount, date, asset ticker, price per share) and automatically calculates the number of shares acquired.

**Key Technologies:**
- **Language:** Go 1.25.7
- **UI Framework:** Bubble Tea v1.3.10 (TUI framework)
- **Styling:** Lipgloss v1.1.0 (terminal styling)
- **Module:** `github.com/danilo/scripts/github/dca`

**Project Type:** Personal investment tracking tool with atomic, file-based data persistence.

Is MANDATORY execute all/each test using the agent `testing-expert` @.qwen/agents/testing-expert.md
Pass the command `go test -v` to **testing-expert**

## Directory Structure

```
/home/danilo/scripts/github/dca/
├── main.go              # Main application entry point with Bubble Tea program setup
├── dca_entry.go         # Core data structures (DCAEntry, DCAData) and file I/O
├── dca_form.go          # FormModel for interactive TUI with validation
├── dca                  # Compiled binary (if built)
├── *.go                 # Test files for all modules
├── go.mod / go.sum      # Go module dependencies
├── .env                 # Environment variables (API keys - not used by app)
├── backlog/             # Backlog.md MCP task management
├── cmd/
│   └── dca/             # Main application code
│       ├── main.go      # Application entry point
│       ├── dca_form.go  # Form model and UI
│       └── dca_form_test.go
└── internal/
    ├── dca/             # Core data model and I/O
    │   ├── entry.go     # DCAEntry, DCAData, file I/O
    │   └── file.go      # LoadEntries, SaveEntries
    ├── form/            # Form UI and validation
    │   ├── model.go     # FormModel state management
    │   ├── validation.go # Field validation
    │   └── validation_test.go
    └── assets/          # Asset aggregation and view
        ├── aggregate.go # Data aggregation logic
        ├── view.go      # AssetsView UI component
        └── view_test.go
```

## Data Model

### DCAEntry Structure
Represents a single investment entry:

```json
{
  "amount": 500.0,
  "date": "2025-01-01T00:00:00Z",
  "asset": "BTC",
  "pricePerShare": 65000.0,
  "shares": 0.00769231
}
```

**Fields:**
- `amount` (float64): USD investment amount (must be positive)
- `date` (time.Time): Investment date in RFC3339 format
- `asset` (string): Asset ticker symbol (e.g., "BTC", "ETH")
- `pricePerShare` (float64): Price per share at time of purchase (must be positive)
- `shares` (float64): Calculated number of shares (8 decimal precision)

### DCAData Structure
Map of asset tickers to arrays of entries:
```json
{
  "entries": {
    "BTC": [DCAEntry1, DCAEntry2, ...],
    "ETH": [DCAEntry1, ...]
  }
}
```

## Building and Running

### Build
```bash
go build -o dca
```

### Run (Development)
```bash
go run main.go
```

### Run (Compiled Binary)
```bash
./dca
```

### Test

Use the agent @.qwen/agents/testing-expert.md
```bash
go test ./...
```

## Development Conventions

### Code Style
- **Formatting:** Standard Go formatting (`go fmt`)
- **Error Handling:** Explicit error returns with descriptive messages
- **Validation:** Input validation at form level and data model level
- **Testing:** Comprehensive test coverage with TDD approach

### Key Functions

#### dca_entry.go / internal/dca/entry.go
- `LoadEntries(filename string) (*DCAData, error)`: Load entries from JSON file with graceful error handling
- `SaveEntries(filename string, data *DCAData) error`: Atomic write with temp file + rename pattern
- `DCAEntry.Validate() error`: Validates amount/price are positive, shares are finite
- `DCAEntry.CalculateShares() float64`: Computes shares with 8-decimal rounding

#### internal/form/validation.go
- `validateAmount(value string) error`: Validates amount is a positive number
- `validateDate(value string) error`: Validates date in RFC3339 format
- `validateAsset(value string) error`: Validates asset ticker is non-empty
- `validatePrice(value string) error`: Validates price is a positive number

#### internal/form/model.go
- `FormModel`: Bubble Tea model managing form state and navigation
- `FormStep` enum: StepAmount → StepDate → StepAsset → StepPrice → StepShares → StepConfirm → StepDone
- `validateAmount/Date/Asset/Price`: Field-level validation with exact error messages
- `CalculateSharesFromValues(amount, price float64) float64`: Helper for share calculation
- `RoundTo8Decimals(val float64) float64`: Precision helper

#### internal/assets/aggregate.go
- `LoadAndAggregateEntries(filename string) (*AssetsViewModel, error)`: Loads and aggregates entries by asset
- `RoundTo8Decimals(val float64) float64`: Rounds to 8 decimal places
- `CalculateWeightedAverage(totalAmount, totalShares float64) float64`: Weighted average price

### Testing Practices
- **Table-driven tests** for validation functions
- **Temp file tests** with cleanup for file I/O
- **Exact error message** assertions for user-facing messages
- **Edge case coverage:** empty values, negative numbers, zero, invalid formats
- **Test naming:** `Test{Function}_{Condition}` pattern

### UI Conventions (Bubble Tea)
- **Rounded borders** using lipgloss
- **Color-coded fields:** Active field highlighted in blue (color 63)
- **Error display:** ❌ prefix with error message
- **Navigation:** Arrow keys (←/→) for field selection, Enter to proceed
- **Exit:** Ctrl+C or Esc to quit without saving

## Backlog.md MCP Integration

This project uses **Backlog.md** for task management via MCP:

**Available MCP Commands:**
- Task management: create, list, search, edit, view, complete, archive
- Documents: create, update, view, search
- Milestones: list, add, rename, remove, archive
- Definition of Done: get, upsert

**Definition of Done:**
1. All acceptance criteria met
2. Unit tests pass (`go test`)
3. No new compiler warnings
4. Code follows project style (`go fmt`)
5. PRD referenced in task
6. Documentation updated (comments)

## User Workflow

1. **Start Application:** `./dca`
2. **Enter Investment Data:**
   - Amount (USD)
   - Date (YYYY-MM-DDTHH:MM:SSZ)
   - Asset Ticker
   - Price per Share
   - Shares (auto-calculated)
   - Confirm (y/n)
3. **Data Saved:** Entry saved to `dca_entries.json` atomically
4. **Exit:** Ctrl+C to quit (after submission)

## File I/O

- **Default File:** `dca_entries.json` (current directory)
- **Format:** JSON with 2-space indentation
- **Safety:** Atomic writes using temp file + rename pattern
- **Error Handling:** Permission errors, missing files, invalid JSON all handled gracefully

## Common Tasks

### Adding a New Test
1. Create test file (e.g., `dca_entry_test.go`)
2. Test functions: `Test{FunctionName}_{Condition}`
3. Test exact error messages for validation functions
4. Use temp files with cleanup for file I/O tests

### Modifying Form Fields
1. Update `FormStep` enum
2. Update `FormField` in `NewFormModel`
3. Add validation in corresponding `handleEnter` case
4. Update `renderForm` field configuration

## Notes for AI Agents

- Is MANDATORY execute all/each test using the agent `testing-expert` @.qwen/agents/testing-expert.md
- **This is a CLI app** - no web server, no network calls
- **Test rigorously** - validation functions need exact error message tests
- **Follow existing patterns** - match function naming and test structure
- **UI changes** require Bubble Tea model updates and lipgloss styling
- **Data persistence** uses atomic write pattern for safety
- **8-decimal precision** is critical for financial calculations
