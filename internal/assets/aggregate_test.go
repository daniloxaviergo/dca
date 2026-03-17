package assets

import (
	"math"
	"os"
	"testing"
	"time"

	"github.com/danilo/scripts/github/dca/internal/dca"
)

// TestAssetSummary_Validate_Pass validates a correct AssetSummary
func TestAssetSummary_Validate_Pass(t *testing.T) {
	summary := AssetSummary{
		Ticker:      "BTC",
		EntryCount:  3,
		TotalShares: 0.023,
		AvgPrice:    55000.0,
		TotalValue:  1500.0,
	}

	if err := summary.Validate(); err != nil {
		t.Errorf("Validate() returned unexpected error: %v", err)
	}
}

// TestAssetSummary_Validate_EmptyTicker rejects empty Ticker
func TestAssetSummary_Validate_EmptyTicker(t *testing.T) {
	summary := AssetSummary{
		Ticker:      "",
		EntryCount:  1,
		TotalShares: 0.01,
		AvgPrice:    50000.0,
		TotalValue:  500.0,
	}

	err := summary.Validate()
	if err == nil {
		t.Error("Validate() should have returned error for empty Ticker")
	}
	if err.Error() != "Ticker is required" {
		t.Errorf("Expected error message 'Ticker is required', got: %v", err)
	}
}

// TestAssetSummary_Validate_NegativeEntryCount rejects negative EntryCount
func TestAssetSummary_Validate_NegativeEntryCount(t *testing.T) {
	summary := AssetSummary{
		Ticker:      "BTC",
		EntryCount:  -1,
		TotalShares: 0.01,
		AvgPrice:    50000.0,
		TotalValue:  500.0,
	}

	err := summary.Validate()
	if err == nil {
		t.Error("Validate() should have returned error for negative EntryCount")
	}
	if err.Error() != "EntryCount cannot be negative" {
		t.Errorf("Expected error message 'EntryCount cannot be negative', got: %v", err)
	}
}

// TestAssetSummary_Validate_NegativeTotalShares rejects negative TotalShares
func TestAssetSummary_Validate_NegativeTotalShares(t *testing.T) {
	summary := AssetSummary{
		Ticker:      "BTC",
		EntryCount:  1,
		TotalShares: -0.01,
		AvgPrice:    50000.0,
		TotalValue:  500.0,
	}

	err := summary.Validate()
	if err == nil {
		t.Error("Validate() should have returned error for negative TotalShares")
	}
	if err.Error() != "TotalShares cannot be negative" {
		t.Errorf("Expected error message 'TotalShares cannot be negative', got: %v", err)
	}
}

// TestAssetSummary_Validate_NegativeTotalValue rejects negative TotalValue
func TestAssetSummary_Validate_NegativeTotalValue(t *testing.T) {
	summary := AssetSummary{
		Ticker:      "BTC",
		EntryCount:  1,
		TotalShares: 0.01,
		AvgPrice:    50000.0,
		TotalValue:  -500.0,
	}

	err := summary.Validate()
	if err == nil {
		t.Error("Validate() should have returned error for negative TotalValue")
	}
	if err.Error() != "TotalValue cannot be negative" {
		t.Errorf("Expected error message 'TotalValue cannot be negative', got: %v", err)
	}
}

// TestCalculateWeightedAverage_Pass calculates weighted average correctly
func TestCalculateWeightedAverage_Pass(t *testing.T) {
	totalAmount := 1500.0
	totalShares := 0.023

	result := CalculateWeightedAverage(totalAmount, totalShares)
	// Expected: 1500 / 0.023 = 65217.3913043478...
	expected := math.Round((1500.0/0.023)*1e8) / 1e8

	if result != expected {
		t.Errorf("CalculateWeightedAverage(%f, %f) = %f, want %f", totalAmount, totalShares, result, expected)
	}
}

// TestCalculateWeightedAverage_ZeroShares returns 0 for zero shares
func TestCalculateWeightedAverage_ZeroShares(t *testing.T) {
	totalAmount := 1500.0
	totalShares := 0.0

	result := CalculateWeightedAverage(totalAmount, totalShares)

	if result != 0 {
		t.Errorf("CalculateWeightedAverage(%f, %f) = %f, want 0", totalAmount, totalShares, result)
	}
}

// TestCalculateWeightedAverage_Precision verifies 8 decimal precision
func TestCalculateWeightedAverage_Precision(t *testing.T) {
	totalAmount := 500.0
	totalShares := 0.00769231

	result := CalculateWeightedAverage(totalAmount, totalShares)
	expected := math.Round((500.0/0.00769231)*1e8) / 1e8

	if result != expected {
		t.Errorf("CalculateWeightedAverage(%f, %f) = %f, want %f", totalAmount, totalShares, result, expected)
	}
}

// TestRoundTo8Decimals_RoundsCorrectly rounds to 8 decimal places
func TestRoundTo8Decimals_RoundsCorrectly(t *testing.T) {
	tests := []struct {
		val  float64
		want float64
	}{
		{val: 3.1415926535, want: 3.14159265},
		{val: 2.718281828, want: 2.71828183},
		{val: 1.0, want: 1.0},
		{val: 0.000000001, want: 0.0},
		{val: 65217.3913043478, want: 65217.39130435},
	}

	for _, tt := range tests {
		got := RoundTo8Decimals(tt.val)
		if got != tt.want {
			t.Errorf("RoundTo8Decimals(%f) = %f, want %f", tt.val, got, tt.want)
		}
	}
}

// TestLoadAndAggregateEntries_EmptyFile handles empty file gracefully
func TestLoadAndAggregateEntries_EmptyFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create empty file
	if err := os.WriteFile(tmpfile.Name(), []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error for empty file: %v", err)
	}

	if len(result.Entries) != 0 {
		t.Errorf("Expected empty entries slice, got %d entries", len(result.Entries))
	}
	if result.Error != nil {
		t.Errorf("Expected Error to be nil, got: %v", result.Error)
	}
}

// TestLoadAndAggregateEntries_MissingFile returns empty data (not an error)
func TestLoadAndAggregateEntries_MissingFile(t *testing.T) {
	result, err := LoadAndAggregateEntries("/nonexistent/file.json")
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error for missing file: %v", err)
	}

	if len(result.Entries) != 0 {
		t.Errorf("Expected empty entries slice, got %d entries", len(result.Entries))
	}
}

// TestLoadAndAggregateEntries_SingleAsset aggregates single asset correctly
func TestLoadAndAggregateEntries_SingleAsset(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	entries := &dca.DCAData{
		Entries: map[string][]dca.DCAEntry{
			"BTC": {
				{
					Amount:        500.0,
					PricePerShare: 65000.0,
					Asset:         "BTC",
					Shares:        0.00769231,
				},
				{
					Amount:        300.0,
					PricePerShare: 60000.0,
					Asset:         "BTC",
					Shares:        0.005,
				},
			},
		},
	}

	if err := dca.SaveEntries(tmpfile.Name(), entries); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error: %v", err)
	}

	if len(result.Entries) != 1 {
		t.Fatalf("Expected 1 asset summary, got %d", len(result.Entries))
	}

	summary := result.Entries[0]
	if summary.Ticker != "BTC" {
		t.Errorf("Expected Ticker 'BTC', got '%s'", summary.Ticker)
	}
	if summary.EntryCount != 2 {
		t.Errorf("Expected EntryCount 2, got %d", summary.EntryCount)
	}
	// Total shares: 0.00769231 + 0.005 = 0.01269231
	expectedShares := 0.01269231
	if summary.TotalShares != expectedShares {
		t.Errorf("Expected TotalShares %.8f, got %.8f", expectedShares, summary.TotalShares)
	}
	// Total amount: 500 + 300 = 800
	// Avg price: 800 / 0.01269231 = 63030.78864373
	expectedAvgPrice := math.Round((800.0/0.01269231)*1e8) / 1e8
	if summary.AvgPrice != expectedAvgPrice {
		t.Errorf("Expected AvgPrice %.8f, got %.8f", expectedAvgPrice, summary.AvgPrice)
	}
	if summary.TotalValue != 800.0 {
		t.Errorf("Expected TotalValue 800.0, got %.8f", summary.TotalValue)
	}
}

// TestLoadAndAggregateEntries_MultipleAssets groups multiple assets correctly
func TestLoadAndAggregateEntries_MultipleAssets(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	entries := &dca.DCAData{
		Entries: map[string][]dca.DCAEntry{
			"BTC": {
				{
					Amount:        500.0,
					PricePerShare: 65000.0,
					Asset:         "BTC",
					Shares:        0.00769231,
				},
			},
			"ETH": {
				{
					Amount:        200.0,
					PricePerShare: 3000.0,
					Asset:         "ETH",
					Shares:        0.06666667,
				},
			},
		},
	}

	if err := dca.SaveEntries(tmpfile.Name(), entries); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error: %v", err)
	}

	if len(result.Entries) != 2 {
		t.Fatalf("Expected 2 asset summaries, got %d", len(result.Entries))
	}

	// Check that both assets are present
	assetTickers := make(map[string]bool)
	for _, s := range result.Entries {
		assetTickers[s.Ticker] = true
	}

	if !assetTickers["BTC"] {
		t.Error("Expected BTC in asset summaries")
	}
	if !assetTickers["ETH"] {
		t.Error("Expected ETH in asset summaries")
	}
}

// TestLoadAndAggregateEntries_MultipleEntriesPerAsset sums shares and calculates weighted average
func TestLoadAndAggregateEntries_MultipleEntriesPerAsset(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	entries := &dca.DCAData{
		Entries: map[string][]dca.DCAEntry{
			"BTC": {
				{
					Amount:        100.0,
					PricePerShare: 50000.0,
					Asset:         "BTC",
					Shares:        0.002,
				},
				{
					Amount:        200.0,
					PricePerShare: 50000.0,
					Asset:         "BTC",
					Shares:        0.004,
				},
				{
					Amount:        300.0,
					PricePerShare: 50000.0,
					Asset:         "BTC",
					Shares:        0.006,
				},
			},
		},
	}

	if err := dca.SaveEntries(tmpfile.Name(), entries); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error: %v", err)
	}

	if len(result.Entries) != 1 {
		t.Fatalf("Expected 1 asset summary, got %d", len(result.Entries))
	}

	summary := result.Entries[0]
	if summary.EntryCount != 3 {
		t.Errorf("Expected EntryCount 3, got %d", summary.EntryCount)
	}
	if summary.TotalShares != 0.012 {
		t.Errorf("Expected TotalShares 0.012, got %.8f", summary.TotalShares)
	}
	// Weighted average: (100+200+300) / (0.002+0.004+0.006) = 600 / 0.012 = 50000
	if summary.AvgPrice != 50000.0 {
		t.Errorf("Expected AvgPrice 50000.0, got %.8f", summary.AvgPrice)
	}
	if summary.TotalValue != 600.0 {
		t.Errorf("Expected TotalValue 600.0, got %.8f", summary.TotalValue)
	}
}

// TestLoadAndAggregateEntries_PopulatedFile loads and aggregates correctly
func TestLoadAndAggregateEntries_PopulatedFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	entries := &dca.DCAData{
		Entries: map[string][]dca.DCAEntry{
			"BTC": {
				{
					Amount:        500.0,
					PricePerShare: 65000.0,
					Asset:         "BTC",
					Date:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					Shares:        0.00769231,
				},
			},
		},
	}

	if err := dca.SaveEntries(tmpfile.Name(), entries); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error: %v", err)
	}

	if len(result.Entries) != 1 {
		t.Errorf("Expected 1 asset summary, got %d", len(result.Entries))
	}

	if result.Error != nil {
		t.Errorf("Expected Error to be nil, got: %v", result.Error)
	}
}

// TestAssetsViewModel_Initialization initializes correctly
func TestAssetsViewModel_Initialization(t *testing.T) {
	vm := &AssetsViewModel{
		Entries: []AssetSummary{},
		Error:   nil,
	}

	if vm.Entries == nil {
		t.Error("Entries should be initialized")
	}
	if vm.Error != nil {
		t.Error("Error should be nil on initialization")
	}
}

// TestAssetSummary_Validate_WithZeroValues handles zero values correctly
func TestAssetSummary_Validate_WithZeroValues(t *testing.T) {
	summary := AssetSummary{
		Ticker:      "BTC",
		EntryCount:  0,
		TotalShares: 0.0,
		AvgPrice:    0.0,
		TotalValue:  0.0,
	}

	err := summary.Validate()
	if err != nil {
		t.Errorf("Validate() should accept zero values for valid asset, got: %v", err)
	}
}

// TestLoadAndAggregateEntries_EmptyEntriesMap handles empty entries map
func TestLoadAndAggregateEntries_EmptyEntriesMap(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	entries := &dca.DCAData{
		Entries: make(map[string][]dca.DCAEntry),
	}

	if err := dca.SaveEntries(tmpfile.Name(), entries); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error: %v", err)
	}

	if len(result.Entries) != 0 {
		t.Errorf("Expected empty entries slice, got %d entries", len(result.Entries))
	}
}

// TestLoadAndAggregateEntries_Calculations_Accurate verifies calculation accuracy
func TestLoadAndAggregateEntries_Calculations_Accurate(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	entries := &dca.DCAData{
		Entries: map[string][]dca.DCAEntry{
			"BTC": {
				{
					Amount:        123.45,
					PricePerShare: 65432.10,
					Asset:         "BTC",
					Shares:        0.00188513,
				},
			},
		},
	}

	if err := dca.SaveEntries(tmpfile.Name(), entries); err != nil {
		t.Fatal(err)
	}

	result, err := LoadAndAggregateEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned error: %v", err)
	}

	if len(result.Entries) != 1 {
		t.Fatalf("Expected 1 asset summary, got %d", len(result.Entries))
	}

	summary := result.Entries[0]

	// Verify calculated fields
	// TotalShares should be 0.00188513
	if summary.TotalShares != 0.00188513 {
		t.Errorf("Expected TotalShares 0.00188513, got %.8f", summary.TotalShares)
	}
	// TotalValue should be 123.45
	if summary.TotalValue != 123.45 {
		t.Errorf("Expected TotalValue 123.45, got %.8f", summary.TotalValue)
	}
	// AvgPrice should be calculated as TotalValue / TotalShares = 123.45 / 0.00188513
	// Note: The calculation uses stored share values, which may have rounding
	expectedAvgPrice := math.Round((123.45/0.00188513)*1e8) / 1e8
	if summary.AvgPrice != expectedAvgPrice {
		t.Errorf("Expected AvgPrice %.8f (123.45/0.00188513), got %.8f", expectedAvgPrice, summary.AvgPrice)
	}
}

// TestLoadAndAggregateEntries_FileNotFound returns empty data (LoadEntries returns empty data for missing files)
func TestLoadAndAggregateEntries_FileNotFound(t *testing.T) {
	// Create a temp file and delete it to get a unique name that doesn't exist
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	tmpName := tmpfile.Name()
	tmpfile.Close()
	os.Remove(tmpName)

	// LoadEntries returns empty data for non-existent files (not an error)
	result, err := LoadAndAggregateEntries(tmpName)
	if err != nil {
		t.Fatalf("LoadAndAggregateEntries() returned unexpected error: %v", err)
	}

	// Should return empty entries slice (not an error)
	if len(result.Entries) != 0 {
		t.Errorf("Expected empty entries slice, got %d entries", len(result.Entries))
	}
}
