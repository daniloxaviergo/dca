---
id: GOT-078
title: Add CLI support with shared validation at cmd/dca/cli.go
status: To Do
assignee: []
created_date: '2026-03-28 15:11'
labels:
  - cli
  - new-feature
  - phase2
  - req-009
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Objective
Add CLI support at cmd/dca/cli.go that uses the shared validation package for command-line DCA entry creation.

## Implementation Details

### 1. New CLI Command Structure

**New File: cmd/dca/cli.go**
```go
package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    
    "github.com/danilo/scripts/github/dca/internal/dca"
    "github.com/danilo/scripts/github/dca/internal/form/validate"
)

// CLICommand represents the CLI entry format
type CLICommand struct {
    Amount    string
    Date      string
    Asset     string
    Price     string
    FilePath  string
}

// NewCLICommand creates a new CLI command handler
func NewCLICommand(amount, date, asset, price, filePath string) *CLICommand {
    return &CLICommand{
        Amount:   amount,
        Date:     date,
        Asset:    asset,
        Price:    price,
        FilePath: filePath,
    }
}

// Validate runs all validations using the shared validation package
func (c *CLICommand) Validate() error {
    // Validate amount
    if err := validate.ValidateAmount(c.Amount); err != nil {
        return err
    }
    
    // Validate date (optional, use current time if missing)
    if c.Date != "" {
        if err := validate.ValidateDate(c.Date); err != nil {
            return err
        }
    } else {
        c.Date = currentTimeRFC3339()
    }
    
    // Validate asset
    if err := validate.ValidateAsset(c.Asset); err != nil {
        return err
    }
    
    // Validate price
    if err := validate.ValidatePrice(c.Price); err != nil {
        return err
    }
    
    return nil
}

// CalculateShares uses shared validation package
func (c *CLICommand) CalculateShares() float64 {
    amount, _ := strconv.ParseFloat(c.Amount, 64)
    price, _ := strconv.ParseFloat(c.Price, 64)
    return validate.CalculateShares(amount, price)
}

// Execute processes the CLI command
func (c *CLICommand) Execute() error {
    // Validate all inputs
    if err := c.Validate(); err != nil {
        return err
    }
    
    // Calculate shares
    shares := c.CalculateShares()
    
    // Create entry
    entry := dca.DCAEntry{
        Amount:        amount,
        Date:          parseDate(c.Date),
        Asset:         c.Asset,
        PricePerShare: price,
        Shares:        shares,
    }
    
    // Validate entry
    if err := entry.Validate(); err != nil {
        return err
    }
    
    // Load existing entries
    data, err := dca.LoadEntries(c.FilePath)
    if err != nil {
        return err
    }
    
    // Add entry
    if data.Entries == nil {
        data.Entries = make(map[string][]dca.DCAEntry)
    }
    data.Entries[c.Asset] = append(data.Entries[c.Asset], entry)
    
    // Save
    return dca.SaveEntries(c.FilePath, data)
}

// Helper functions
func currentTimeRFC3339() string {
    return time.Now().Format(time.RFC3339)
}

func parseDate(dateStr string) time.Time {
    t, _ := time.Parse(time.RFC3339, dateStr)
    return t
}
```

### 2. CLI Command Usage

**New command: `dca add`**
```bash
# Add entry with all fields
./dca add --amount 100 --date 2024-01-15 --asset BTC --price 50000

# Add entry with current date
./dca add --amount 100 --asset BTC --price 50000

# Help
./dca add --help
```

### 3. Integration with Existing Code

**Update cmd/dca/main.go:**
```go
// Track whether CLI mode or TUI mode
func main() {
    if len(os.Args) > 1 && os.Args[1] == "add" {
        // CLI mode
        runCLI(os.Args[2:])
    } else {
        // TUI mode
        runTUI()
    }
}
```

### 4. CLI Validation Consistency

**Validation must match TUI exactly:**
- Same error messages
- Same validation rules
- Same shares calculation
- Same precision

**Test verification:**
```go
// CLI validation produces same result as TUI validation
func TestCLI_ValidationMatchesTUI(t *testing.T) {
    amount, date, asset, price := "-10", "invalid", "", "0"
    
    tuiErr := form.validateAmount(amount)
    cliErr := cli.ValidateAmount(amount)  // Uses validate package
    
    if tuiErr.Error() != cliErr.Error() {
        t.Errorf("Validation mismatch: TUI=%s CLI=%s", tuiErr, cliErr)
    }
}
```

## Acceptance Criteria
<!-- AC:BEGIN -->
- ✅ CLI at cmd/dca/cli.go created
- ✅ CLI uses shared validation package
- ✅ CLI validation identical to TUI validation
- ✅ CLI shares calculation identical to TUI
- ✅ Error messages match TUI exactly
<!-- SECTION:DESCRIPTION:END -->

- [ ] #1 CLI command at cmd/dca/cli.go created
- [ ] #2 CLI uses shared validation package (validate package)
- [ ] #3 CLI validation produces identical results to TUI validation
- [ ] #4 CLI shares calculation matches TUI exactly
- [ ] #5 Error messages identical between CLI and TUI
- [ ] #6 CLI integration with main.go for mode switching
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
