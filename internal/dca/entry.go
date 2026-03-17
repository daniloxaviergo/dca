package dca

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

// DCAEntry represents a single DCA investment entry
type DCAEntry struct {
	Amount        float64   `json:"amount"`
	Date          time.Time `json:"date"`
	Asset         string    `json:"asset"`
	PricePerShare float64   `json:"pricePerShare"`
	Shares        float64   `json:"shares"`
}

// DCAData represents a collection of DCA entries keyed by asset ticker
type DCAData struct {
	Entries map[string][]DCAEntry `json:"entries"`
}

// LoadEntries reads DCA entries from a JSON file
func LoadEntries(filename string) (*DCAData, error) {
	data := &DCAData{
		Entries: make(map[string][]DCAEntry),
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return data, nil
		}
		// Handle permission errors with user-friendly message
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			if os.IsPermission(err) {
				return nil, fmt.Errorf("Permission denied: check file permissions")
			}
		}
		return nil, err
	}

	if len(file) == 0 {
		return data, nil
	}

	if err := json.Unmarshal(file, data); err != nil {
		// Handle JSON parse errors gracefully with diagnostic message
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			return nil, fmt.Errorf("JSON parse error at byte %d: invalid data format", jsonErr.Offset)
		}
		return nil, fmt.Errorf("JSON parse error: %v", err)
	}

	return data, nil
}

// SaveEntries writes DCA entries to a JSON file with 2-space indentation using atomic write
func SaveEntries(filename string, data *DCAData) error {
	// 1. Marshal JSON
	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize data: %v", err)
	}

	// 2. Create temp file in same directory
	tmpfile, err := os.CreateTemp(filepath.Dir(filename), ".dca_entries_*.json")
	if err != nil {
		// Handle permission errors with user-friendly message
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			if os.IsPermission(err) {
				return fmt.Errorf("Permission denied: check file permissions")
			}
		}
		return fmt.Errorf("failed to create temp file: %v", err)
	}

	// 3. Write to temp file
	if _, err := tmpfile.Write(fileBytes); err != nil {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
		return fmt.Errorf("failed to write to temp file: %v", err)
	}

	// 4. Close temp file before rename (required on Windows)
	if err := tmpfile.Close(); err != nil {
		os.Remove(tmpfile.Name())
		return fmt.Errorf("failed to close temp file: %v", err)
	}

	// 5. Atomic rename
	if err := os.Rename(tmpfile.Name(), filename); err != nil {
		os.Remove(tmpfile.Name())
		// Handle permission errors with user-friendly message
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			if os.IsPermission(err) {
				return fmt.Errorf("Permission denied: check file permissions")
			}
		}
		return fmt.Errorf("failed to rename temp file: %v", err)
	}

	return nil
}

// Validate checks that the DCAEntry has valid values
func (e *DCAEntry) Validate() error {
	if e.Amount <= 0 {
		return errors.New("Amount must be positive")
	}
	if e.PricePerShare <= 0 {
		return errors.New("Price must be positive")
	}
	// Validate shares is a valid finite positive number
	shares := e.CalculateShares()
	if math.IsNaN(shares) || math.IsInf(shares, 0) || shares <= 0 {
		return errors.New("Shares must be a positive finite number")
	}
	return nil
}

// CalculateShares computes the number of shares based on amount and price per share
// with 8 decimal precision using rounding
func (e *DCAEntry) CalculateShares() float64 {
	if e.PricePerShare == 0 {
		return 0
	}
	shares := e.Amount / e.PricePerShare
	return math.Round(shares*1e8) / 1e8
}
