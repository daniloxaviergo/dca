package main

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"time"
)

// DCAEntry represents a single DCA investment entry
type DCAEntry struct {
	Amount         float64   `json:"amount"`
	Date           time.Time `json:"date"`
	Asset          string    `json:"asset"`
	PricePerShare  float64   `json:"pricePerShare"`
	Shares         float64   `json:"shares"`
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
		return nil, err
	}

	if len(file) == 0 {
		return data, nil
	}

	if err := json.Unmarshal(file, data); err != nil {
		return nil, err
	}

	return data, nil
}

// SaveEntries writes DCA entries to a JSON file with 2-space indentation
func SaveEntries(filename string, data *DCAData) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0644)
}

// Validate checks that the DCAEntry has valid values
func (e *DCAEntry) Validate() error {
	if e.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if e.PricePerShare <= 0 {
		return errors.New("price per share must be greater than 0")
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
