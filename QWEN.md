# DCA Investment Tracker - Project Context

## Project Overview

**DCA Investment Tracker** is a command-line application for tracking Dollar-Cost Averaging (DCA) investment entries with an interactive terminal UI. The application allows users to record investment data (amount, date, asset ticker, price) and automatically calculates shares acquired. It provides an interactive TUI built with Bubble Tea for data entry and asset aggregation visualization.

### Key Technologies

- **Language**: Go 1.25.7
- **TUI Framework**: Bubble Tea v1.3.10
- **Styling**: Lipgloss v1.1.0
- **Data Storage**: JSON file persistence (`dca_entries.json`)
- **Architecture**: Modular structure with clear separation of concerns

### Architecture

```
dca/
├── cmd/dca/                  # Application entry point
│   ├── main.go              # Bubble Tea program, state management
│   └── dca_form.go          # Form validation and submission (deprecated, merged into internal/form)
├── internal/
│   ├── assets/              # Asset aggregation and UI
│   │   ├── aggregate.go     # Asset aggregation logic
│   │   ├── model.go         # Data models (AssetSummary, AssetsViewModel)
│   │   ├── view.go          # Assets view UI component
│   │   └── view_test.go     # View tests
│   ├── dca/                 # Core data model and file I/O
│   │   ├── entry.go         # DCAEntry, DCAData structs, validation
│   │   ├── entry_test.go    # Data model tests
│   │   └── file.go          # File I/O (load/save with atomic writes)
│   └── form/                # Interactive form component
│       ├── model.go         # Form state management, Bubble Tea component
│       ├── validation.go    # Field validation logic
│       └── validation_test.go
├── backlog/                 # Backlog.md task management
├── dca                      # Compiled binary
├── dca_entries.json         # Data persistence file
├── go.mod / go.sum          # Go module dependencies
├── Makefile                 # Development commands
├── coverage.out / coverage_assets.out  # Test coverage reports
└── AGENTS.md                # MCP Backlog workflow guidelines
```

## Building and Running

### Prerequisites

- Go 1.25.7 or higher
- Terminal with UTF-8 support

### Build

```bash
go build -o dca ./cmd/dca
# Or use make:
make build
```

### Run

```bash
./dca
# Or use make:
make run

# Or run directly:
go run ./cmd/dca
```

### Development Commands

| Command | Description |
|---------|-------------|
| `make help` | Display all available commands |
| `make build` | Build the binary |
| `make run` | Run the application |
| `make test` | Run all tests with verbose output (`-v`) |
| `make test-quiet` | Run all tests without verbose output |
| `make test-cover` | Generate coverage report with `go tool cover` |
| `make fmt` | Format all Go files with `go fmt` |
| `make check` | Run fmt, build, and test (CI validation) |
| `make clean` | Remove binary and coverage files |
| `make version` | Show Go version |

## Usage

The application starts in **Assets View** showing aggregated investment data. Use keyboard navigation:

### Assets View (Default)

- `↑` / `↓` - Navigate through asset list (wrap-around)
- `c` - Switch to Form View to enter new data
- `Enter` - Open asset history modal for selected asset
- `Esc` / `Ctrl+C` - Exit application

**Asset List Columns:**
| Column | Description |
|--------|-------------|
| Asset | Ticker symbol |
| Count | Number of entries |
| Total Shares | Sum of shares (8 decimals) |
| Avg Price | Weighted average price (2 decimals) |
| Total Value | Sum of amounts invested (2 decimals) |

### Form View (Press `c`)

1. Fill in fields with `←`/`→` navigation (or `Tab`)
2. Press `Enter` to validate and proceed
3. Shares are auto-calculated with 8 decimal precision
4. Press `Esc` to cancel (no save) and return to assets view
5. Press `Ctrl+C` to exit at any time

**Form Fields:**
| Field | Description | Validation |
|-------|-------------|------------|
| Amount | USD investment | Positive number |
| Date | Investment date | RFC3339 format |
| Asset | Ticker symbol | Non-empty |
| Price | Price per share | Positive number |
| Shares | Auto-calculated | Positive finite number |
| Confirm | Submit (y/n) | y to confirm |

### Asset History Modal (Press `Enter` on asset)

- `↑` / `↓` - Scroll through entries
- `Enter` - Load more entries (paginated, 10 at a time)
- `Esc` - Close modal and return to asset list

**Modal Columns:**
| Column | Description |
|--------|-------------|
| Date | YYYY-MM-DD |
| Avg Price | Weighted average per day (2 decimals) |
| Total Invested | Sum per day (2 decimals) |
| Entry Count | Number of entries per day |

## Data Format

### JSON File Structure

Entries are persisted to `dca_entries.json`:

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
    ]
  }
}
```

### Core Data Models

```go
type DCAEntry struct {
    Amount        float64   `json:"amount"`
    Date          time.Time `json:"date"`
    Asset         string    `json:"asset"`
    PricePerShare float64   `json:"pricePerShare"`
    Shares        float64   `json:"shares"`
}

type DCAData struct {
    Entries map[string][]DCAEntry `json:"entries"`
}
```

### Share Calculation

Shares are calculated as `amount / pricePerShare` with 8 decimal precision using rounding:
```go
shares := math.Round((amount / pricePerShare) * 1e8) / 1e8
```

## Testing

Run all tests with verbose output:
```bash
go test -v ./...
# Or:
make test
```

### Test Coverage

- **Data model** (`internal/dca/entry_test.go`): Entry validation, share calculations, file I/O
- **Form validation** (`internal/form/validation_test.go`): Field validation logic
- **Asset aggregation** (`internal/assets/aggregate_test.go`): Aggregation calculations
- **UI components** (`internal/assets/view_test.go`, `cmd/dca/`): Component state transitions

### Coverage Report

```bash
make test-cover
# Generates coverage.out and displays function coverage
```

## Development Conventions

### Code Style

- **Go fmt**: All code is formatted with `go fmt`
- **Validation**: All validation returns descriptive error messages
- **Error handling**: Use `errors.Is()`, `errors.As()` for specific error checking
- **Floating point**: 8 decimal precision for shares, 2 decimals for display

### State Management

- **Bubble Tea**: Use custom message types for state transitions
- **Form steps**: Explicit step tracking (`StepAmount`, `StepDate`, etc.)
- **View transitions**: Custom messages (`ViewTransitionMsg`, `FormSubmittedMsg`, `FormCancelledMsg`)

### UI Styling

- **Lipgloss**: Consistent styling for borders, colors, alignment
- **Fixed-width columns**: Table columns use exact character widths for alignment
- **30-row table**: Assets view maintains exactly 30 rows (1 header + 29 data)

### File I/O

- **Atomic writes**: Use temp file + rename pattern for safe writes
- **Graceful handling**: Check for missing files, empty files, permission errors
- **JSON indentation**: 2-space indentation for readability

## Project-Specific Notes

### Current State

- **Main branch** is the active development branch
- Application starts in Assets View (not Form View)
- Form submission switches back to Assets View
- Modal displays paginated history with "Load More" functionality

### Common Tasks

**Adding a new feature:**
1. Create/modify files under `internal/` or `cmd/dca/`
2. Add tests for new functionality
3. Run `make check` to validate (fmt, build, test)

**Debugging:**
- Check `dca_entries.json` for data issues
- Use `make test-cover` to verify test coverage
- Review `coverage.out` for untested code paths

### Dependencies

```go
github.com/charmbracelet/bubbletea v1.3.10  // TUI framework
github.com/charmbracelet/lipgloss v1.1.0    // Terminal styling
```

See `go.mod` for full dependency list including indirect dependencies.

## MCP Backlog Workflow

This project uses **Backlog.md** for task management. See `AGENTS.md` for MCP workflow guidelines. Key points:

- Read `backlog://workflow/overview` before creating tasks
- Use search-first workflow to avoid duplicates
- Task creation, execution, and finalization guides available via MCP
