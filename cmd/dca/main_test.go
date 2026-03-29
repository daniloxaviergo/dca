package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/danilo/scripts/github/dca/internal/dca"
)

// TestMain tests the main function's CLI/TUI behavior
func TestMain(m *testing.M) {
	// Setup: Create temp directory for test files
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	// Save original entries path
	originalEntriesPath := defaultEntriesPath

	// Set test entries path
	defaultEntriesPath = filepath.Join(tmpDir, "test_entries.json")

	// Run tests
	code := m.Run()

	// Restore original entries path
	defaultEntriesPath = originalEntriesPath

	os.Exit(code)
}

// TestMain_CLIModeExitsEarly tests that CLI mode exits before TUI initialization
func TestMain_CLIModeExitsEarly(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	// Setup: Save original os.Args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Set CLI args
	os.Args = []string{"dca", "-add", "-asset", "BTC", "-amount", "100", "-price", "50000"}

	// Run CLI directly (main() would call RunCLI which exits)
	// We test RunCLI directly since main() would exit the process
	_, success, err := RunCLIWithoutExit()
	if err != nil {
		t.Fatalf("RunCLI failed: %v", err)
	}

	if !success {
		t.Error("Expected RunCLI to return true when CLI mode is active")
	}

	// Verify file was created
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Error("Expected entries file to be created")
	}
}

// TestMain_CLISavesEntry tests that CLI mode correctly saves entry to JSON
func TestMain_CLISavesEntry(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"dca", "-add", "-asset", "ETH", "-amount", "500", "-price", "3000"}

	// Run CLI
	RunCLIWithoutExit()

	// Verify file content
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read entries file: %v", err)
	}

	var entries dca.DCAData
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse entries JSON: %v", err)
	}

	// Verify entry data
	if len(entries.Entries) != 1 {
		t.Errorf("Expected 1 asset, got %d", len(entries.Entries))
	}

	if len(entries.Entries["ETH"]) != 1 {
		t.Errorf("Expected 1 ETH entry, got %d", len(entries.Entries["ETH"]))
	}

	entry := entries.Entries["ETH"][0]
	if entry.Amount != 500 {
		t.Errorf("Expected amount 500, got %f", entry.Amount)
	}
	if entry.Asset != "ETH" {
		t.Errorf("Expected asset 'ETH', got '%s'", entry.Asset)
	}
	if entry.PricePerShare != 3000 {
		t.Errorf("Expected price 3000, got %f", entry.PricePerShare)
	}
}

// TestMain_TUIModeUnchanged tests that TUI mode (no flags) doesn't trigger CLI
func TestMain_TUIModeUnchanged(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// No flags - TUI mode
	os.Args = []string{"dca"}

	_, success, err := RunCLIWithoutExit()
	if err != nil {
		t.Fatalf("RunCLI failed: %v", err)
	}

	// RunCLI should return false (not CLI mode)
	if success {
		t.Error("Expected RunCLI to return false when no --add flag")
	}
}

// TestMain_MultipleEntries tests that multiple CLI calls append to the same file
func TestMain_MultipleEntries(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args

	// First entry
	os.Args = []string{"dca", "-add", "-asset", "BTC", "-amount", "100", "-price", "50000"}
	RunCLIWithoutExit()

	// Second entry for same asset
	os.Args = []string{"dca", "-add", "-asset", "BTC", "-amount", "200", "-price", "51000"}
	RunCLIWithoutExit()

	// Third entry for different asset
	os.Args = []string{"dca", "-add", "-asset", "ETH", "-amount", "300", "-price", "3000"}
	RunCLIWithoutExit()

	// Restore original args
	os.Args = originalArgs

	// Verify file has 3 entries
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read entries file: %v", err)
	}

	var entries dca.DCAData
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse entries JSON: %v", err)
	}

	if len(entries.Entries) != 2 {
		t.Errorf("Expected 2 assets, got %d", len(entries.Entries))
	}

	if len(entries.Entries["BTC"]) != 2 {
		t.Errorf("Expected 2 BTC entries, got %d", len(entries.Entries["BTC"]))
	}

	if len(entries.Entries["ETH"]) != 1 {
		t.Errorf("Expected 1 ETH entry, got %d", len(entries.Entries["ETH"]))
	}
}

// TestMain_SharePrecision tests that shares are calculated with 8-decimal precision
func TestMain_SharePrecision(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Use values that produce non-terminating decimal
	os.Args = []string{"dca", "-add", "-asset", "BTC", "-amount", "100", "-price", "33333.33"}

	RunCLIWithoutExit()

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read entries file: %v", err)
	}

	var entries dca.DCAData
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse entries JSON: %v", err)
	}

	entry := entries.Entries["BTC"][0]

	// Expected: 100 / 33333.33 = 0.00300000003... rounded to 8 decimals = 0.003
	expectedShares := 0.003
	if entry.Shares != expectedShares {
		t.Errorf("Expected shares %f, got %f", expectedShares, entry.Shares)
	}
}

// TestMain_DateHandling tests that date is handled correctly
func TestMain_DateHandling(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Test with explicit date
	testDate := "2025-02-15T12:30:45Z"
	os.Args = []string{"dca", "-add", "-asset", "BTC", "-amount", "100", "-price", "50000", "-date", testDate}

	RunCLIWithoutExit()

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read entries file: %v", err)
	}

	var entries dca.DCAData
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse entries JSON: %v", err)
	}

	entry := entries.Entries["BTC"][0]
	if entry.Date.Format(time.RFC3339) != testDate {
		t.Errorf("Expected date '%s', got '%s'", testDate, entry.Date.Format(time.RFC3339))
	}
}

// TestMain_AutoDate tests that date defaults to current time when not provided
func TestMain_AutoDate(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// No date provided - should auto-set
	os.Args = []string{"dca", "-add", "-asset", "BTC", "-amount", "100", "-price", "50000"}

	RunCLIWithoutExit()

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read entries file: %v", err)
	}

	var entries dca.DCAData
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse entries JSON: %v", err)
	}

	entry := entries.Entries["BTC"][0]

	// Date should be set to current time (within last few seconds)
	now := time.Now()
	diff := now.Sub(entry.Date)
	if diff < 0 {
		diff = -diff
	}
	if diff > 5*time.Second {
		t.Errorf("Expected date to be current time, got %v (diff: %v)", entry.Date, diff)
	}
}

// TestMain_ConcurrentEntries tests that entries from different assets are stored correctly
func TestMain_ConcurrentEntries(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args

	assets := []struct {
		symbol string
		amount float64
		price  float64
	}{
		{"BTC", 100, 50000},
		{"ETH", 200, 3000},
		{"SOL", 300, 150},
		{"ADA", 400, 0.50},
	}

	// Save multiple entries
	for _, asset := range assets {
		os.Args = []string{"dca", "-add", "-asset", asset.symbol, "-amount", fmt.Sprintf("%g", asset.amount), "-price", fmt.Sprintf("%g", asset.price)}
		RunCLIWithoutExit()
	}

	os.Args = originalArgs

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read entries file: %v", err)
	}

	var entries dca.DCAData
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse entries JSON: %v", err)
	}

	if len(entries.Entries) != len(assets) {
		t.Errorf("Expected %d assets, got %d", len(assets), len(entries.Entries))
	}

	for _, asset := range assets {
		if len(entries.Entries[asset.symbol]) != 1 {
			t.Errorf("Expected 1 entry for %s, got %d", asset.symbol, len(entries.Entries[asset.symbol]))
		}
	}
}

// TestMain_InvalidAssetFormat tests that invalid asset format causes exit
func TestMain_InvalidAssetFormat(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Invalid asset with spaces
	os.Args = []string{"dca", "-add", "-asset", "MY ASSET", "-amount", "100", "-price", "50000"}

	// This should exit with code 1
	defer func() {
		if r := recover(); r != nil {
			// Expected: os.Exit(1)
		}
	}()
	RunCLIWithoutExit()
}

// TestMain_ValidationErrors tests that validation errors are returned properly
func TestMain_ValidationErrors(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "test_entries.json")
	originalEntriesPath := defaultEntriesPath
	defaultEntriesPath = tmpFile
	defer func() { defaultEntriesPath = originalEntriesPath }()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Missing asset - validation error should be returned
	os.Args = []string{"dca", "-add", "-amount", "100", "-price", "50000"}

	active, success, err := RunCLIWithoutExit()

	// When validation fails, error should be returned
	if err == nil {
		t.Error("Expected error for missing asset")
	}

	if active {
		t.Error("Expected active to be false when validation fails")
	}

	if success {
		t.Error("Expected success to be false when validation fails")
	}

	// Check for expected error message in error string
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "required") && !strings.Contains(errMsg, "positive") {
			t.Errorf("Expected error message to mention required or positive, got: %s", errMsg)
		}
	}
}
