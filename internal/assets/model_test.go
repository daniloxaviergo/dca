package assets

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/danilo/scripts/github/dca/internal/dca"
)

// TestAssetHistoryModal_LoadData_Pass loads valid entries successfully
func TestAssetHistoryModal_LoadData_Pass(t *testing.T) {
	// Create temp file with test data
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	entries := map[string][]dca.DCAEntry{
		"BTC": {
			{
				Amount:        100,
				Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
				Asset:         "BTC",
				PricePerShare: 50000,
				Shares:        0.002,
			},
			{
				Amount:        200,
				Date:          time.Date(2025, 1, 16, 10, 0, 0, 0, time.UTC),
				Asset:         "BTC",
				PricePerShare: 51000,
				Shares:        0.00392157,
			},
		},
	}

	data := &dca.DCAData{Entries: entries}
	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	// Test loading data
	m := NewAssetHistoryModal()
	err = m.LoadData(tmpfile.Name(), "BTC")
	if err != nil {
		t.Fatalf("LoadData() returned error: %v", err)
	}

	if !m.Loaded {
		t.Error("Expected Loaded to be true")
	}
	if len(m.EntriesByDate) != 2 {
		t.Errorf("Expected 2 entries by date, got %d", len(m.EntriesByDate))
	}
	if m.AssetTicker != "BTC" {
		t.Errorf("Expected AssetTicker to be 'BTC', got '%s'", m.AssetTicker)
	}
}

// TestAssetHistoryModal_LoadData_EmptyAsset handles empty asset entries
func TestAssetHistoryModal_LoadData_EmptyAsset(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	entries := map[string][]dca.DCAEntry{
		"BTC": {},
	}

	data := &dca.DCAData{Entries: entries}
	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	m := NewAssetHistoryModal()
	err = m.LoadData(tmpfile.Name(), "BTC")
	if err != nil {
		t.Fatalf("LoadData() returned error for empty asset: %v", err)
	}

	if !m.Loaded {
		t.Error("Expected Loaded to be true")
	}
	if len(m.EntriesByDate) != 0 {
		t.Errorf("Expected 0 entries by date for empty asset, got %d", len(m.EntriesByDate))
	}
}

// TestAssetHistoryModal_LoadData_MissingAsset handles missing asset gracefully
func TestAssetHistoryModal_LoadData_MissingAsset(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	entries := map[string][]dca.DCAEntry{
		"ETH": {
			{
				Amount:        100,
				Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
				Asset:         "ETH",
				PricePerShare: 3000,
				Shares:        0.03333333,
			},
		},
	}

	data := &dca.DCAData{Entries: entries}
	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	m := NewAssetHistoryModal()
	err = m.LoadData(tmpfile.Name(), "BTC") // BTC not in data
	if err != nil {
		t.Fatalf("LoadData() returned error for missing asset: %v", err)
	}

	if !m.Loaded {
		t.Error("Expected Loaded to be true")
	}
	if len(m.EntriesByDate) != 0 {
		t.Errorf("Expected 0 entries by date for missing asset, got %d", len(m.EntriesByDate))
	}
}

// TestAssetHistoryModal_LoadData_FileNotFound handles missing file
func TestAssetHistoryModal_LoadData_FileNotFound(t *testing.T) {
	m := NewAssetHistoryModal()
	err := m.LoadData("/nonexistent/file.json", "BTC")

	// Should return empty data, not an error, for missing file
	if err != nil {
		t.Errorf("LoadData() returned error for missing file: %v", err)
	}

	if !m.Loaded {
		t.Error("Expected Loaded to be true for missing file (empty data)")
	}
	if len(m.EntriesByDate) != 0 {
		t.Errorf("Expected 0 entries by date for missing file, got %d", len(m.EntriesByDate))
	}
}

// TestAggregateByDate_Grouping groups entries by calendar date
func TestAggregateByDate_Grouping(t *testing.T) {
	tests := []struct {
		name     string
		entries  []dca.DCAEntry
		expected int
	}{
		{
			name: "single day",
			entries: []dca.DCAEntry{
				{
					Amount: 100,
					Date:   time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					Asset:  "BTC",
				},
				{
					Amount: 200,
					Date:   time.Date(2025, 1, 15, 14, 0, 0, 0, time.UTC), // Same day
					Asset:  "BTC",
				},
			},
			expected: 1,
		},
		{
			name: "multiple days",
			entries: []dca.DCAEntry{
				{
					Amount: 100,
					Date:   time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					Asset:  "BTC",
				},
				{
					Amount: 200,
					Date:   time.Date(2025, 1, 16, 10, 0, 0, 0, time.UTC),
					Asset:  "BTC",
				},
			},
			expected: 2,
		},
		{
			name: "same day different times",
			entries: []dca.DCAEntry{
				{
					Amount: 100,
					Date:   time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
					Asset:  "BTC",
				},
				{
					Amount: 200,
					Date:   time.Date(2025, 1, 15, 23, 59, 59, 999999999, time.UTC),
					Asset:  "BTC",
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AggregateByDate(tt.entries)

			if len(result) != tt.expected {
				t.Errorf("Expected %d date groups, got %d", tt.expected, len(result))
			}
		})
	}
}

// TestAggregateByDate_Calculations verifies all metrics are calculated correctly
func TestAggregateByDate_Calculations(t *testing.T) {
	tests := []struct {
		name               string
		entries            []dca.DCAEntry
		expectedTotal      float64
		expectedAvgPrice   float64
		expectedEntryCount int
		expectedDate       string
	}{
		{
			name: "single entry",
			entries: []dca.DCAEntry{
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					Asset:         "BTC",
					PricePerShare: 50000,
					Shares:        0.002,
				},
			},
			expectedTotal:      100,
			expectedAvgPrice:   50000,
			expectedEntryCount: 1,
			expectedDate:       "2025-01-15",
		},
		{
			name: "multiple entries same day",
			entries: []dca.DCAEntry{
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					Asset:         "BTC",
					PricePerShare: 50000,
					Shares:        0.002,
				},
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 14, 0, 0, 0, time.UTC),
					Asset:         "BTC",
					PricePerShare: 50000,
					Shares:        0.002,
				},
			},
			expectedTotal:      200,
			expectedAvgPrice:   50000,
			expectedEntryCount: 2,
			expectedDate:       "2025-01-15",
		},
		{
			name: "mixed prices same day",
			entries: []dca.DCAEntry{
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					Asset:         "BTC",
					PricePerShare: 50000,
					Shares:        0.002,
				},
				{
					Amount:        150,
					Date:          time.Date(2025, 1, 15, 14, 0, 0, 0, time.UTC),
					PricePerShare: 51000,
					Shares:        0.00294118,
				},
			},
			expectedTotal:      250,
			expectedAvgPrice:   50600.0, // PRD formula: (50000×100 + 51000×150) / (100+150) = 12650000/250
			expectedEntryCount: 2,
			expectedDate:       "2025-01-15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AggregateByDate(tt.entries)

			if len(result) != 1 {
				t.Fatalf("Expected 1 day result, got %d", len(result))
			}

			day := result[0]
			if day.Date != tt.expectedDate {
				t.Errorf("Expected date '%s', got '%s'", tt.expectedDate, day.Date)
			}

			// Use 8-decimal precision for comparison
			expectedTotal := RoundTo8Decimals(tt.expectedTotal)
			expectedAvgPrice := RoundTo8Decimals(tt.expectedAvgPrice)

			if day.TotalInvested != expectedTotal {
				t.Errorf("Expected total invested %f, got %f", expectedTotal, day.TotalInvested)
			}
			if day.WeightedAvgPrice != expectedAvgPrice {
				t.Errorf("Expected weighted avg price %f, got %f", expectedAvgPrice, day.WeightedAvgPrice)
			}
			if day.EntryCount != tt.expectedEntryCount {
				t.Errorf("Expected entry count %d, got %d", tt.expectedEntryCount, day.EntryCount)
			}
		})
	}
}

// TestAggregateByDate_Sorting sorts entries by date descending (newest first)
func TestAggregateByDate_Sorting(t *testing.T) {
	entries := []dca.DCAEntry{
		{
			Amount: 100,
			Date:   time.Date(2025, 1, 20, 10, 0, 0, 0, time.UTC),
			Asset:  "BTC",
		},
		{
			Amount: 100,
			Date:   time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			Asset:  "BTC",
		},
		{
			Amount: 100,
			Date:   time.Date(2025, 1, 18, 10, 0, 0, 0, time.UTC),
			Asset:  "BTC",
		},
	}

	result := AggregateByDate(entries)

	if len(result) != 3 {
		t.Fatalf("Expected 3 days, got %d", len(result))
	}

	// Verify dates are in descending order (newest first)
	expectedDates := []string{"2025-01-20", "2025-01-18", "2025-01-15"}
	for i, expected := range expectedDates {
		if result[i].Date != expected {
			t.Errorf("Date at position %d: expected '%s', got '%s'", i, expected, result[i].Date)
		}
	}
}

// TestAggregateByDate_EmptyEntries returns empty slice for empty input
func TestAggregateByDate_EmptyEntries(t *testing.T) {
	result := AggregateByDate([]dca.DCAEntry{})

	if len(result) != 0 {
		t.Errorf("Expected empty slice for empty entries, got %d items", len(result))
	}
}

// TestCalculateDayMetrics_WeightedAverage calculates weighted average correctly using PRD formula
// PRD formula: sum(price_per_share × amount) / sum(amounts)
func TestCalculateDayMetrics_WeightedAverage(t *testing.T) {
	tests := []struct {
		name             string
		totalAmount      float64
		totalShares      float64
		expectedAvgPrice float64
	}{
		{
			name:             "standard case",
			totalAmount:      200,
			totalShares:      0.004,
			expectedAvgPrice: 50000,
		},
		{
			name:             "different amounts",
			totalAmount:      350,
			totalShares:      0.007,
			expectedAvgPrice: 50000,
		},
		{
			name:             "edge case - zero shares",
			totalAmount:      100,
			totalShares:      0,
			expectedAvgPrice: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary entries for the test
			dateStr := "2025-01-15"
			entries := []dca.DCAEntry{}

			if tt.totalShares > 0 {
				// Create a single entry with the total values
				// For PRD formula with single entry: weighted avg = price (since amount × price / amount = price)
				entries = []dca.DCAEntry{
					{
						Amount:        tt.totalAmount,
						Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
						PricePerShare: tt.expectedAvgPrice,
						Shares:        tt.totalShares,
					},
				}
			}

			result := calculateDayMetrics(dateStr, entries)

			expectedAvgPrice := RoundTo8Decimals(tt.expectedAvgPrice)
			if result.WeightedAvgPrice != expectedAvgPrice {
				t.Errorf("Expected weighted avg price %f, got %f", expectedAvgPrice, result.WeightedAvgPrice)
			}
		})
	}
}

// TestCalculateDayMetrics_Calculations verifies all metrics in calculateDayMetrics
func TestCalculateDayMetrics_Calculations(t *testing.T) {
	tests := []struct {
		name               string
		entries            []dca.DCAEntry
		expectedDate       string
		expectedTotal      float64
		expectedAvgPrice   float64
		expectedEntryCount int
	}{
		{
			name: "single entry",
			entries: []dca.DCAEntry{
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					Asset:         "BTC",
					PricePerShare: 50000,
					Shares:        0.002,
				},
			},
			expectedDate:       "2025-01-15",
			expectedTotal:      100,
			expectedAvgPrice:   50000,
			expectedEntryCount: 1,
		},
		{
			name: "multiple entries",
			entries: []dca.DCAEntry{
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
					PricePerShare: 50000,
					Shares:        0.002,
				},
				{
					Amount:        100,
					Date:          time.Date(2025, 1, 15, 14, 0, 0, 0, time.UTC),
					PricePerShare: 50000,
					Shares:        0.002,
				},
			},
			expectedDate:       "2025-01-15",
			expectedTotal:      200,
			expectedAvgPrice:   50000,
			expectedEntryCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateDayMetrics(tt.expectedDate, tt.entries)

			expectedTotal := RoundTo8Decimals(tt.expectedTotal)
			expectedAvgPrice := RoundTo8Decimals(tt.expectedAvgPrice)

			if result.Date != tt.expectedDate {
				t.Errorf("Expected date '%s', got '%s'", tt.expectedDate, result.Date)
			}
			if result.TotalInvested != expectedTotal {
				t.Errorf("Expected total invested %f, got %f", expectedTotal, result.TotalInvested)
			}
			if result.WeightedAvgPrice != expectedAvgPrice {
				t.Errorf("Expected weighted avg price %f, got %f", expectedAvgPrice, result.WeightedAvgPrice)
			}
			if result.EntryCount != tt.expectedEntryCount {
				t.Errorf("Expected entry count %d, got %d", tt.expectedEntryCount, result.EntryCount)
			}
		})
	}
}

// TestNewAssetHistoryModal initializes modal with correct defaults
func TestNewAssetHistoryModal(t *testing.T) {
	m := NewAssetHistoryModal()

	if m.SelectedIndex != 0 {
		t.Errorf("Expected SelectedIndex to be 0, got %d", m.SelectedIndex)
	}
	if m.Loaded {
		t.Error("Expected Loaded to be false initially")
	}
	if m.Error != nil {
		t.Error("Expected Error to be nil initially")
	}
	if m.Visible {
		t.Error("Expected Visible to be false initially")
	}
	if m.AssetTicker != "" {
		t.Errorf("Expected empty AssetTicker initially, got '%s'", m.AssetTicker)
	}
	if len(m.EntriesByDate) != 0 {
		t.Errorf("Expected empty EntriesByDate initially, got %d items", len(m.EntriesByDate))
	}
}

// TestAggregateByDate_MultipleEntriesPerDay handles multiple entries on same day
func TestAggregateByDate_MultipleEntriesPerDay(t *testing.T) {
	entries := []dca.DCAEntry{
		{
			Amount:        100,
			Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000,
			Shares:        0.002,
		},
		{
			Amount:        150,
			Date:          time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 51000,
			Shares:        0.00294118,
		},
		{
			Amount:        200,
			Date:          time.Date(2025, 1, 15, 14, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 49000,
			Shares:        0.00408163,
		},
	}

	result := AggregateByDate(entries)

	if len(result) != 1 {
		t.Fatalf("Expected 1 day, got %d", len(result))
	}

	day := result[0]
	if day.Date != "2025-01-15" {
		t.Errorf("Expected date '2025-01-15', got '%s'", day.Date)
	}
	if day.EntryCount != 3 {
		t.Errorf("Expected entry count 3, got %d", day.EntryCount)
	}
	// Total invested = 100 + 150 + 200 = 450
	expectedTotal := RoundTo8Decimals(450.0)
	if day.TotalInvested != expectedTotal {
		t.Errorf("Expected total invested %f, got %f", expectedTotal, day.TotalInvested)
	}
	// Weighted avg (PRD formula) = sum(price × amount) / sum(amounts)
	// = (50000×100 + 51000×150 + 49000×200) / 450 = 22450000 / 450 ≈ 49888.88888889
	expectedAvgPrice := RoundTo8Decimals(22450000.0 / 450.0)
	if day.WeightedAvgPrice != expectedAvgPrice {
		t.Errorf("Expected weighted avg price %f, got %f", expectedAvgPrice, day.WeightedAvgPrice)
	}
}

// TestAssetHistoryModal_LoadData_MultipleAssets loads correct asset data
func TestAssetHistoryModal_LoadData_MultipleAssets(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	entries := map[string][]dca.DCAEntry{
		"BTC": {
			{
				Amount:        100,
				Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
				Asset:         "BTC",
				PricePerShare: 50000,
				Shares:        0.002,
			},
		},
		"ETH": {
			{
				Amount:        200,
				Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
				Asset:         "ETH",
				PricePerShare: 3000,
				Shares:        0.06666667,
			},
		},
	}

	data := &dca.DCAData{Entries: entries}
	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	// Test loading BTC
	m := NewAssetHistoryModal()
	err = m.LoadData(tmpfile.Name(), "BTC")
	if err != nil {
		t.Fatalf("LoadData() for BTC returned error: %v", err)
	}

	if len(m.EntriesByDate) != 1 {
		t.Errorf("Expected 1 entry for BTC, got %d", len(m.EntriesByDate))
	}
	if m.AssetTicker != "BTC" {
		t.Errorf("Expected AssetTicker to be 'BTC', got '%s'", m.AssetTicker)
	}

	// Test loading ETH
	m2 := NewAssetHistoryModal()
	err = m2.LoadData(tmpfile.Name(), "ETH")
	if err != nil {
		t.Fatalf("LoadData() for ETH returned error: %v", err)
	}

	if len(m2.EntriesByDate) != 1 {
		t.Errorf("Expected 1 entry for ETH, got %d", len(m2.EntriesByDate))
	}
	if m2.AssetTicker != "ETH" {
		t.Errorf("Expected AssetTicker to be 'ETH', got '%s'", m2.AssetTicker)
	}
}

// TestAggregateByDate_PreservesDateFormat ensures YYYY-MM-DD format
func TestAggregateByDate_PreservesDateFormat(t *testing.T) {
	entries := []dca.DCAEntry{
		{
			Amount: 100,
			Date:   time.Date(2025, 3, 7, 10, 0, 0, 0, time.UTC),
			Asset:  "BTC",
		},
	}

	result := AggregateByDate(entries)

	if len(result) != 1 {
		t.Fatalf("Expected 1 day, got %d", len(result))
	}

	expectedFormat := "2025-03-07"
	if result[0].Date != expectedFormat {
		t.Errorf("Expected date format 'YYYY-MM-DD', got '%s'", result[0].Date)
	}
}

// TestLoadData_LimitsToMaxDays limits results to 10 days
func TestLoadData_LimitsToMaxDays(t *testing.T) {
	// Create temp file with 15 days of entries
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	var entries []dca.DCAEntry
	for i := 0; i < 15; i++ {
		entries = append(entries, dca.DCAEntry{
			Amount: 100,
			Date:   time.Date(2025, 1, i+1, 10, 0, 0, 0, time.UTC),
			Asset:  "BTC",
		})
	}

	data := &dca.DCAData{Entries: map[string][]dca.DCAEntry{"BTC": entries}}
	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	m := NewAssetHistoryModal()
	err = m.LoadData(tmpfile.Name(), "BTC")
	if err != nil {
		t.Fatalf("LoadData() returned error: %v", err)
	}

	// Should be limited to 10 days
	if len(m.EntriesByDate) != 10 {
		t.Errorf("Expected limited to 10 days, got %d", len(m.EntriesByDate))
	}
}

// TestLoadData_MissingFileReturnsEmpty handles missing data file gracefully
func TestLoadData_MissingFileReturnsEmpty(t *testing.T) {
	m := NewAssetHistoryModal()

	// Test with non-existent file in temp directory
	tmpDir, err := os.MkdirTemp("", "dca_test_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpDir)

	missingFile := filepath.Join(tmpDir, "nonexistent.json")

	err = m.LoadData(missingFile, "BTC")
	if err != nil {
		t.Errorf("LoadData() returned error for missing file: %v", err)
	}

	if !m.Loaded {
		t.Error("Expected Loaded to be true for missing file (returns empty data)")
	}
}

// TestAggregateByDate_CumulativeTotal calculates running total correctly
func TestAggregateByDate_CumulativeTotal(t *testing.T) {
	entries := []dca.DCAEntry{
		{
			Amount:        100,
			Date:          time.Date(2025, 1, 20, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000, // Add price for PRD formula
			Shares:        0.002,
		},
		{
			Amount:        100,
			Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000, // Add price for PRD formula
			Shares:        0.002,
		},
		{
			Amount:        100,
			Date:          time.Date(2025, 1, 18, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000, // Add price for PRD formula
			Shares:        0.002,
		},
	}

	result := AggregateByDate(entries)

	if len(result) != 3 {
		t.Fatalf("Expected 3 days, got %d", len(result))
	}

	// Results should be sorted descending (2025-01-20, 2025-01-18, 2025-01-15)
	// Cumulative totals:
	// Position 0 (2025-01-20): 100
	// Position 1 (2025-01-18): 100 + 100 = 200
	// Position 2 (2025-01-15): 100 + 100 + 100 = 300
	expectedTotals := []float64{100.0, 200.0, 300.0}
	for i, expected := range expectedTotals {
		if result[i].TotalInvested != expected {
			t.Errorf("TotalInvested at position %d: expected %.2f, got %.2f", i, expected, result[i].TotalInvested)
		}
	}

	// Weighted avg price should be 50000 for all (same price each day)
	expectedAvgPrice := 50000.0
	for i := range expectedTotals {
		if result[i].WeightedAvgPrice != expectedAvgPrice {
			t.Errorf("WeightedAvgPrice at position %d: expected %.2f, got %.2f", i, expectedAvgPrice, result[i].WeightedAvgPrice)
		}
	}
}

// TestAggregateByDate_DescendingSortAndCumulative verifies both descending sort and cumulative totals
func TestAggregateByDate_DescendingSortAndCumulative(t *testing.T) {
	entries := []dca.DCAEntry{
		{
			Amount:        100,
			Date:          time.Date(2025, 1, 20, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000,
			Shares:        0.002,
		},
		{
			Amount:        50,
			Date:          time.Date(2025, 1, 20, 12, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000,
			Shares:        0.001,
		},
		{
			Amount:        75,
			Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 50000,
			Shares:        0.0015,
		},
	}

	result := AggregateByDate(entries)

	if len(result) != 2 {
		t.Fatalf("Expected 2 days, got %d", len(result))
	}

	// Should be sorted descending
	if result[0].Date != "2025-01-20" {
		t.Errorf("Expected first date '2025-01-20', got '%s'", result[0].Date)
	}
	if result[1].Date != "2025-01-15" {
		t.Errorf("Expected second date '2025-01-15', got '%s'", result[1].Date)
	}

	// Cumulative totals: 150 (20th), 225 (15th = 150 + 75)
	expectedTotals := []float64{150.0, 225.0}
	for i, expected := range expectedTotals {
		if result[i].TotalInvested != expected {
			t.Errorf("TotalInvested at position %d: expected %.2f, got %.2f", i, expected, result[i].TotalInvested)
		}
	}

	// Weighted avg price should be 50000 for all (same price each day)
	expectedAvgPrice := 50000.0
	for i := range expectedTotals {
		if result[i].WeightedAvgPrice != expectedAvgPrice {
			t.Errorf("WeightedAvgPrice at position %d: expected %.2f, got %.2f", i, expectedAvgPrice, result[i].WeightedAvgPrice)
		}
	}

	// Entry count: 2 on 20th, 1 on 15th
	if result[0].EntryCount != 2 {
		t.Errorf("Expected entry count 2 on 2025-01-20, got %d", result[0].EntryCount)
	}
	if result[1].EntryCount != 1 {
		t.Errorf("Expected entry count 1 on 2025-01-15, got %d", result[1].EntryCount)
	}
}
