package assets

import (
	"errors"
	"fmt"

	"github.com/danilo/scripts/github/dca/internal/dca"
)

// AssetSummary represents aggregated data for a single asset
type AssetSummary struct {
	Ticker      string
	EntryCount  int
	TotalShares float64
	AvgPrice    float64
	TotalValue  float64
}

// AssetsViewModel manages the loaded and aggregated asset data
type AssetsViewModel struct {
	Entries []AssetSummary
	Error   error
}

// RoundTo8Decimals rounds a float to 8 decimal places
func RoundTo8Decimals(val float64) float64 {
	return float64(int(val*1e8+.5)) / 1e8
}

// LoadAndAggregateEntries loads entries from a JSON file and aggregates them by asset ticker
func LoadAndAggregateEntries(filename string) (*AssetsViewModel, error) {
	// Load entries from file
	data, err := dca.LoadEntries(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load entries: %w", err)
	}

	// Handle empty entries map gracefully
	if len(data.Entries) == 0 {
		return &AssetsViewModel{
			Entries: []AssetSummary{},
			Error:   nil,
		}, nil
	}

	// Aggregate entries by asset
	var summaries []AssetSummary
	for ticker, entries := range data.Entries {
		summary := aggregateEntries(ticker, entries)
		summaries = append(summaries, summary)
	}

	return &AssetsViewModel{
		Entries: summaries,
		Error:   nil,
	}, nil
}

// aggregateEntries aggregates a slice of DCAEntry for a single asset
func aggregateEntries(ticker string, entries []dca.DCAEntry) AssetSummary {
	var totalShares float64
	var totalAmount float64
	var sumPriceAmount float64

	for _, entry := range entries {
		totalShares += entry.Shares
		totalAmount += entry.Amount
		sumPriceAmount += entry.PricePerShare * entry.Amount
	}

	// Weighted average price (PRD formula): sum(price_per_share × amount) / sum(amounts)
	var avgPrice float64
	if totalAmount > 0 {
		avgPrice = RoundTo8Decimals(sumPriceAmount / totalAmount)
	}

	return AssetSummary{
		Ticker:      ticker,
		EntryCount:  len(entries),
		TotalShares: RoundTo8Decimals(totalShares),
		AvgPrice:    avgPrice,
		TotalValue:  RoundTo8Decimals(totalAmount),
	}
}

// CalculateWeightedAverage calculates the weighted average price using the PRD formula:
// weightedAverage = sum(price_per_share × amount) / sum(amounts)
func CalculateWeightedAverage(totalAmount, sumPriceAmount float64) float64 {
	if totalAmount == 0 {
		return 0
	}
	return RoundTo8Decimals(sumPriceAmount / totalAmount)
}

// Validate validates an AssetSummary
func (s *AssetSummary) Validate() error {
	if s.Ticker == "" {
		return errors.New("Ticker is required")
	}
	if s.EntryCount < 0 {
		return errors.New("EntryCount cannot be negative")
	}
	if s.TotalShares < 0 {
		return errors.New("TotalShares cannot be negative")
	}
	if s.TotalValue < 0 {
		return errors.New("TotalValue cannot be negative")
	}
	return nil
}
