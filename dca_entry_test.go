package main

import (
	"encoding/json"
	"math"
	"os"
	"testing"
	"time"
)

// TestDCAEntryValidate_Pass validates a correct DCAEntry
func TestDCAEntryValidate_Pass(t *testing.T) {
	entry := DCAEntry{
		Amount:         500.0,
		PricePerShare:  65000.0,
		Asset:          "BTC",
		Date:           time.Now(),
		Shares:         0,
	}

	if err := entry.Validate(); err != nil {
		t.Errorf("Validate() returned unexpected error: %v", err)
	}
}

// TestDCAEntryValidate_ZeroAmount rejects Amount = 0
func TestDCAEntryValidate_ZeroAmount(t *testing.T) {
	entry := DCAEntry{
		Amount:         0,
		PricePerShare:  65000.0,
		Asset:          "BTC",
		Date:           time.Now(),
		Shares:         0,
	}

	if err := entry.Validate(); err == nil {
		t.Error("Validate() should have returned error for Amount = 0")
	}
}

// TestDCAEntryValidate_NegativeAmount rejects Amount < 0
func TestDCAEntryValidate_NegativeAmount(t *testing.T) {
	entry := DCAEntry{
		Amount:         -100.0,
		PricePerShare:  65000.0,
		Asset:          "BTC",
		Date:           time.Now(),
		Shares:         0,
	}

	if err := entry.Validate(); err == nil {
		t.Error("Validate() should have returned error for Amount < 0")
	}
}

// TestDCAEntryValidate_ZeroPrice rejects PricePerShare = 0
func TestDCAEntryValidate_ZeroPrice(t *testing.T) {
	entry := DCAEntry{
		Amount:         500.0,
		PricePerShare:  0,
		Asset:          "BTC",
		Date:           time.Now(),
		Shares:         0,
	}

	if err := entry.Validate(); err == nil {
		t.Error("Validate() should have returned error for PricePerShare = 0")
	}
}

// TestDCAEntryValidate_NegativePrice rejects PricePerShare < 0
func TestDCAEntryValidate_NegativePrice(t *testing.T) {
	entry := DCAEntry{
		Amount:         500.0,
		PricePerShare:  -500.0,
		Asset:          "BTC",
		Date:           time.Now(),
		Shares:         0,
	}

	if err := entry.Validate(); err == nil {
		t.Error("Validate() should have returned error for PricePerShare < 0")
	}
}

// TestCalculateShares verifies 8 decimal precision
func TestCalculateShares(t *testing.T) {
	entry := DCAEntry{
		Amount:        500.0,
		PricePerShare: 65000.0,
	}

	shares := entry.CalculateShares()
	expected := math.Round((500.0/65000.0)*1e8) / 1e8

	if shares != expected {
		t.Errorf("CalculateShares() = %v, want %v", shares, expected)
	}

	// Verify it matches 0.00769231
	if shares != 0.00769231 {
		t.Errorf("CalculateShares() = %v, want 0.00769231", shares)
	}
}

// TestCalculateShares_Precision verifies rounding behavior
func TestCalculateShares_Precision(t *testing.T) {
	entry := DCAEntry{
		Amount:        200.0,
		PricePerShare: 32000.0,
	}

	shares := entry.CalculateShares()
	expected := 0.00625

	if shares != expected {
		t.Errorf("CalculateShares() = %v, want %v", shares, expected)
	}
}

// TestSaveEntries_CreateFile creates new file with proper structure
func TestSaveEntries_CreateFile(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	data := &DCAData{
		Entries: map[string][]DCAEntry{
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

	if err := SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatalf("SaveEntries() returned error: %v", err)
	}

	// Verify file exists and contains valid JSON
	fileData, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var result DCAData
	if err := json.Unmarshal(fileData, &result); err != nil {
		t.Fatalf("File is not valid JSON: %v", err)
	}

	if len(result.Entries["BTC"]) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(result.Entries["BTC"]))
	}
}

// TestSaveEntries_UpdateFile appends to existing data
func TestSaveEntries_UpdateFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create initial data
	data := &DCAData{
		Entries: map[string][]DCAEntry{
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

	if err := SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	// Update with new data (overwrite)
	data2 := &DCAData{
		Entries: map[string][]DCAEntry{
			"ETH": {
				{
					Amount:        200.0,
					PricePerShare: 3200.0,
					Asset:         "ETH",
					Date:          time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC),
					Shares:        0.0625,
				},
			},
		},
	}

	if err := SaveEntries(tmpfile.Name(), data2); err != nil {
		t.Fatal(err)
	}

	// Verify file contains updated data
	fileData, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	var result DCAData
	if err := json.Unmarshal(fileData, &result); err != nil {
		t.Fatal(err)
	}

	if _, exists := result.Entries["ETH"]; !exists {
		t.Error("Expected ETH entry to exist")
	}
}

// TestLoadEntries_Populated reads existing JSON correctly
func TestLoadEntries_Populated(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	data := &DCAData{
		Entries: map[string][]DCAEntry{
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

	if err := SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatal(err)
	}

	result, err := LoadEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadEntries() returned error: %v", err)
	}

	if len(result.Entries["BTC"]) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(result.Entries["BTC"]))
	}
}

// TestLoadEntries_EmptyFile handles empty file gracefully
func TestLoadEntries_EmptyFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create empty file
	if err := os.WriteFile(tmpfile.Name(), []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result, err := LoadEntries(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadEntries() returned error for empty file: %v", err)
	}

	if len(result.Entries) != 0 {
		t.Errorf("Expected empty entries map, got %d entries", len(result.Entries))
	}
}

// TestLoadEntries_MissingFile returns empty data (not an error)
func TestLoadEntries_MissingFile(t *testing.T) {
	result, err := LoadEntries("/nonexistent/file.json")
	if err != nil {
		t.Fatalf("LoadEntries() returned error for missing file: %v", err)
	}

	if len(result.Entries) != 0 {
		t.Errorf("Expected empty entries map, got %d entries", len(result.Entries))
	}
}

// TestLoadEntries_InvalidJSON returns error on malformed JSON
func TestLoadEntries_InvalidJSON(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write invalid JSON
	if err := os.WriteFile(tmpfile.Name(), []byte("{invalid json"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err = LoadEntries(tmpfile.Name())
	if err == nil {
		t.Error("LoadEntries() should have returned error for invalid JSON")
	}
}

// TestDCAData_Structure verifies JSON structure matches PRD example
func TestDCAData_JSONStructure(t *testing.T) {
	data := &DCAData{
		Entries: map[string][]DCAEntry{
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

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	// Verify it contains expected keys
	var parsed map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsed); err != nil {
		t.Fatal(err)
	}

	if _, exists := parsed["entries"]; !exists {
		t.Error("JSON should contain 'entries' key")
	}
}
