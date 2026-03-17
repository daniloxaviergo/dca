# DCA Investment Tracker

A command-line application for tracking Dollar-Cost Averaging (DCA) investment entries with an interactive terminal UI built using Bubble Tea.

## Overview

DCA Investment Tracker helps you record and analyze your dollar-cost averaging investment strategy. Enter your investment amounts, dates, asset tickers, and prices to automatically calculate shares acquired. View aggregated data for all your assets in an interactive terminal table.

### Features

- Interactive TUI form for entering investment data
- Automatic share calculation with 8 decimal precision
- JSON-based persistence for investment entries
- Assets view showing aggregated statistics per asset
- Keyboard navigation for easy data entry and navigation

## Architecture

The application follows a modular structure with clear separation of concerns:

```
cmd/
└── dca/
    ├── main.go          # Application entry point with Bubble Tea program
    └── dca_form.go      # Form validation and submission logic
internal/
├── assets/
│   ├── aggregate.go     # Asset aggregation logic
│   ├── view.go          # Assets view UI component
│   └── view_test.go     # View tests
├── dca/
│   ├── entry.go         # Core data model (DCAEntry, DCAData)
│   ├── entry_test.go    # Data model tests
│   └── file.go          # File I/O operations (load/save)
└── form/
    ├── model.go         # Form model and state management
    ├── validation.go    # Field validation logic
    └── validation_test.go
```

### Package Dependencies

```
cmd/dca
    ├── internal/form
    │   └── internal/dca
    ├── internal/assets
    │   └── internal/dca
    └── internal/dca
```

## Getting Started

### Prerequisites

- Go 1.25.7 or higher
- Terminal with UTF-8 support

### Build

```bash
go build -o dca ./cmd/dca
```

### Run

```bash
./dca
```

Or run directly:

```bash
go run ./cmd/dca
```

## Usage

### Entering Data

1. Launch the application: `./dca`
2. The form interface will appear with the following fields:

| Field | Description | Format |
|-------|-------------|--------|
| Amount | USD investment amount | Positive number |
| Date | Investment date | RFC3339 (YYYY-MM-DDTHH:MM:SSZ) |
| Asset | Asset ticker symbol | Text (e.g., BTC, ETH) |
| Price | Price per share | Positive number |
| Shares | Calculated shares | Auto-calculated (8 decimals) |

3. Navigate between fields with `←`/`→` or `Tab`
4. Press `Enter` to proceed to the next field or submit
5. Press `Ctrl+C` or `Esc` to cancel/exit at any time

### Viewing Assets

After submitting an entry, the application switches to the Assets view showing:

| Column | Description |
|--------|-------------|
| Asset | Ticker symbol |
| Count | Number of entries for this asset |
| Total Shares | Sum of all shares |
| Avg Price | Weighted average price |
| Total Value | Sum of all amounts invested |

Navigate the list with `↑`/`↓` arrows. Press `Esc` or `Ctrl+C` to return to the form.

## Data Format

### JSON File Structure

Entries are persisted to `dca_entries.json` in the current directory:

```json
{
  "entries": {
    "BTC": [
      {
        "amount": 500.0,
        "date": "2025-01-01T00:00:00Z",
        "asset": "BTC",
        "pricePerShare": 65000.0,
        "shares": 0.00769231
      }
    ],
    "ETH": [
      {
        "amount": 200.0,
        "date": "2025-01-02T00:00:00Z",
        "asset": "ETH",
        "pricePerShare": 3200.0,
        "shares": 0.0625
      }
    ]
  }
}
```

### Data Model

#### DCAEntry

```go
type DCAEntry struct {
	Amount        float64   `json:"amount"`
	Date          time.Time `json:"date"`
	Asset         string    `json:"asset"`
	PricePerShare float64   `json:"pricePerShare"`
	Shares        float64   `json:"shares"`
}
```

#### DCAData

```go
type DCAData struct {
	Entries map[string][]DCAEntry `json:"entries"`
}
```

## Testing

Run all tests:

```bash
go test ./...
```

Test coverage includes:

- **Data model validation** (`internal/dca/`): Validates entry fields, share calculations, file I/O
- **Form validation** (`internal/form/`): Validates user input for all form fields
- **Asset aggregation** (`internal/assets/`): Tests aggregation logic and calculations
- **UI components** (`cmd/dca/`): Tests form state transitions and submission

### Test Commands

```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./internal/dca
go test ./internal/form
go test ./internal/assets
go test ./cmd/dca
```

## Extending the Application

### Adding a New Feature

1. Create a new package under `internal/` or `cmd/`
2. Write tests first (TDD approach)
3. Implement functionality following existing patterns
4. Run tests: `go test ./...`

### Code Patterns

#### Validation

All validation functions return descriptive error messages:

```go
func (m *FormModel) validateAmount(value string) error {
    if value == "" {
        return fmt.Errorf("Amount must be positive")
    }
    // ... validation logic
}
```

#### UI Styling

Use Lipgloss for consistent styling:

```go
titleStyle := lipgloss.NewStyle().
    Foreground(lipgloss.Color("159")).
    Background(lipgloss.Color("236")).
    Bold(true).
    Render("Title")
```

#### State Management

Use custom message types for Bubble Tea state transitions:

```go
type formSubmittedMsg struct{}
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

See `go.mod` for full dependency list.

## License

This project is for personal use.
