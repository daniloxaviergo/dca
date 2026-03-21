package assets

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/danilo/scripts/github/dca/internal/dca"
)

// createTestFileWithBTC creates a temp file with BTC entries for testing
func createTestFileWithBTC(t *testing.T) string {
	t.Helper()

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
			Date:          time.Date(2025, 1, 16, 10, 0, 0, 0, time.UTC),
			Asset:         "BTC",
			PricePerShare: 51000,
			Shares:        0.00294117,
		},
	}

	data := &dca.DCAData{Entries: map[string][]dca.DCAEntry{"BTC": entries}}
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpfile.Close()

	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatalf("Failed to save test data: %v", err)
	}

	return tmpfile.Name()
}

// createTestFileWithMultipleAssets creates a temp file with multiple asset entries for testing
func createTestFileWithMultipleAssets(t *testing.T) string {
	t.Helper()

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
				Shares:        0.06666666,
			},
		},
		"SOL": {
			{
				Amount:        50,
				Date:          time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
				Asset:         "SOL",
				PricePerShare: 100,
				Shares:        0.5,
			},
		},
	}

	data := &dca.DCAData{Entries: entries}
	tmpfile, err := os.CreateTemp("", "dca_entries_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpfile.Close()

	if err := dca.SaveEntries(tmpfile.Name(), data); err != nil {
		t.Fatalf("Failed to save test data: %v", err)
	}

	return tmpfile.Name()
}

// TestAssetsView_UpdateKeyEnter_FirstRow opens modal on first data row
func TestAssetsView_UpdateKeyEnter_FirstRow(t *testing.T) {
	tmpfile := createTestFileWithBTC(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}
	av.SelectedIndex = 1 // First data row

	fmt.Printf("Before Update: SelectedIndex=%d, Entry ticker=%s\n", av.SelectedIndex, av.Entries[0].Ticker)

	newAv, _ := av.Update(tea.KeyMsg{Type: tea.KeyEnter})

	// Modal should be visible after pressing Enter on first data row
	av = newAv.(*AssetsView)
	fmt.Printf("After Update: Modal=%v, Visible=%v, AssetTicker='%s'\n", av.Modal != nil, av.Modal != nil && av.Modal.Visible, av.Modal.AssetTicker)
	fmt.Printf("Modal Loaded=%v, Error=%v\n", av.Modal.Loaded, av.Modal.Error)

	if av.Modal == nil {
		t.Fatal("Expected modal to be created, got nil")
	}
	if !av.Modal.Visible {
		t.Error("Expected modal to be visible after Enter on first row")
	}
	if av.Modal.AssetTicker != "BTC" {
		t.Errorf("Expected AssetTicker to be 'BTC', got '%s'", av.Modal.AssetTicker)
	}
	if !av.Modal.Loaded {
		t.Error("Expected modal to be loaded")
	}
}

// TestAssetsView_UpdateKeyEnter_MiddleRow opens modal on middle data row
func TestAssetsView_UpdateKeyEnter_MiddleRow(t *testing.T) {
	tmpfile := createTestFileWithMultipleAssets(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
		{Ticker: "SOL"},
	}
	av.SelectedIndex = 2 // Middle data row (ETH)

	newAv, _ := av.Update(tea.KeyMsg{Type: tea.KeyEnter})
	av = newAv.(*AssetsView)

	if av.Modal == nil {
		t.Fatal("Expected modal to be created, got nil")
	}
	if !av.Modal.Visible {
		t.Error("Expected modal to be visible after Enter on middle row")
	}
	if av.Modal.AssetTicker != "ETH" {
		t.Errorf("Expected AssetTicker to be 'ETH', got '%s'", av.Modal.AssetTicker)
	}
}

// TestAssetsView_UpdateKeyEnter_LastRow opens modal on last visible data row
func TestAssetsView_UpdateKeyEnter_LastRow(t *testing.T) {
	tmpfile := createTestFileWithBTC(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}
	av.SelectedIndex = 1 // First data row (only one entry)

	newAv, _ := av.Update(tea.KeyMsg{Type: tea.KeyEnter})
	av = newAv.(*AssetsView)

	if av.Modal == nil {
		t.Fatal("Expected modal to be created, got nil")
	}
	if !av.Modal.Visible {
		t.Error("Expected modal to be visible after Enter on last row")
	}
	if av.Modal.AssetTicker != "BTC" {
		t.Errorf("Expected AssetTicker to be 'BTC', got '%s'", av.Modal.AssetTicker)
	}
}

// TestAssetsView_UpdateKeyEnter_HeaderRow does not open modal on header
func TestAssetsView_UpdateKeyEnter_HeaderRow(t *testing.T) {
	tmpfile := createTestFileWithBTC(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}
	av.SelectedIndex = 0 // Header row

	newAv, cmd := av.Update(tea.KeyMsg{Type: tea.KeyEnter})
	av = newAv.(*AssetsView)

	// Modal should NOT be created when Enter is pressed on header
	if av.Modal != nil && av.Modal.Visible {
		t.Error("Expected modal to NOT be visible after Enter on header row")
	}
	// No command should be returned
	if cmd != nil {
		msg := cmd()
		if _, ok := msg.(OpenModalMsg); ok {
			t.Error("Expected no OpenModalMsg when Enter pressed on header")
		}
	}
}

// TestAssetsView_UpdateKeyEnter_NoEntries does not open modal when list is empty
func TestAssetsView_UpdateKeyEnter_NoEntries(t *testing.T) {
	tmpfile := createTestFileWithBTC(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{}
	av.SelectedIndex = 1

	newAv, cmd := av.Update(tea.KeyMsg{Type: tea.KeyEnter})
	av = newAv.(*AssetsView)

	// Modal should NOT be created when list is empty
	if av.Modal != nil && av.Modal.Visible {
		t.Error("Expected modal to NOT be visible when list is empty")
	}
	if cmd != nil {
		msg := cmd()
		if _, ok := msg.(OpenModalMsg); ok {
			t.Error("Expected no OpenModalMsg when Enter pressed on empty list")
		}
	}
}

// TestAssetsView_UpdateKeyEnter_NoModalOnOutOfBound does not open modal when out of bounds
func TestAssetsView_UpdateKeyEnter_NoModalOnOutOfBound(t *testing.T) {
	tmpfile := createTestFileWithBTC(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}
	av.SelectedIndex = 5 // Out of bounds (only 1 entry, valid range is 1-1)

	newAv, cmd := av.Update(tea.KeyMsg{Type: tea.KeyEnter})
	av = newAv.(*AssetsView)

	// Modal should NOT be created when selection is out of bounds
	if av.Modal != nil && av.Modal.Visible {
		t.Error("Expected modal to NOT be visible when selection is out of bounds")
	}
	if cmd != nil {
		msg := cmd()
		if _, ok := msg.(OpenModalMsg); ok {
			t.Error("Expected no OpenModalMsg when Enter pressed on out of bounds selection")
		}
	}
}

// TestAssetsView_UpdateKeyEnter_MultipleAssets opens modal for each asset
func TestAssetsView_UpdateKeyEnter_MultipleAssets(t *testing.T) {
	tmpfile := createTestFileWithMultipleAssets(t)
	defer os.Remove(tmpfile)

	av := NewAssetsView()
	av.Filename = tmpfile
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
		{Ticker: "SOL"},
	}

	testCases := []struct {
		name        string
		selectedIdx int
		expected    string
	}{
		{"First asset (BTC)", 1, "BTC"},
		{"Second asset (ETH)", 2, "ETH"},
		{"Third asset (SOL)", 3, "SOL"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset modal for each test case
			av.Modal = nil
			av.SelectedIndex = tc.selectedIdx

			newAv, _ := av.Update(tea.KeyMsg{Type: tea.KeyEnter})
			av = newAv.(*AssetsView)

			if av.Modal == nil {
				t.Fatal("Expected modal to be created, got nil")
			}
			if !av.Modal.Visible {
				t.Error("Expected modal to be visible")
			}
			if av.Modal.AssetTicker != tc.expected {
				t.Errorf("Expected AssetTicker to be '%s', got '%s'", tc.expected, av.Modal.AssetTicker)
			}
		})
	}
}
