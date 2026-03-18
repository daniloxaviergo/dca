package assets

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbletea"
)

// TestAssetsView_RenderEmpty shows "No assets yet" when list is empty
func TestAssetsView_RenderEmpty(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true

	output := av.View()

	if !strings.Contains(output, "No assets yet") {
		t.Errorf("Expected output to contain 'No assets yet', got: %s", output)
	}
}

// TestAssetsView_RenderWithEntries shows table with data
func TestAssetsView_RenderWithEntries(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC", EntryCount: 3, TotalShares: 0.01, AvgPrice: 50000, TotalValue: 500},
		{Ticker: "ETH", EntryCount: 2, TotalShares: 0.02, AvgPrice: 3000, TotalValue: 60},
	}

	output := av.View()

	expectedHeaders := []string{"Asset", "Count", "Total Shares", "Avg Price", "Total Value"}
	for _, h := range expectedHeaders {
		if !strings.Contains(output, h) {
			t.Errorf("Expected output to contain header '%s', got: %s", h, output)
		}
	}

	if !strings.Contains(output, "BTC") {
		t.Errorf("Expected output to contain 'BTC', got: %s", output)
	}
	if !strings.Contains(output, "ETH") {
		t.Errorf("Expected output to contain 'ETH', got: %s", output)
	}
}

// TestAssetsView_NavigateDown increments selection index
func TestAssetsView_NavigateDown(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
		{Ticker: "SOL"},
	}
	av.SelectedIndex = 0

	newAv, _ := av.handleDown()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 1 {
		t.Errorf("Expected selectedIndex to be 1 after down, got %d", av.SelectedIndex)
	}

	newAv, _ = av.handleDown()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 2 {
		t.Errorf("Expected selectedIndex to be 2 after down, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_NavigateUp decrements selection index
func TestAssetsView_NavigateUp(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
		{Ticker: "SOL"},
	}
	av.SelectedIndex = 2

	newAv, _ := av.handleUp()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 1 {
		t.Errorf("Expected selectedIndex to be 1 after up, got %d", av.SelectedIndex)
	}

	// Wrapping from first data row (index 1) to last visible row (index 29)
	newAv, _ = av.handleUp()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 29 {
		t.Errorf("Expected selectedIndex to be 29 (wrapped) after up, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_NavigateWrapDown wraps from last visible row to first data row
func TestAssetsView_NavigateWrapDown(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
	}
	av.SelectedIndex = 29 // last visible row (30 total rows - 1)

	newAv, _ := av.handleDown()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 1 {
		t.Errorf("Expected selectedIndex to wrap to 1 (first data row), got %d", av.SelectedIndex)
	}
}

// TestAssetsView_NavigateWrapUp wraps from first data row to last visible row
func TestAssetsView_NavigateWrapUp(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
	}
	av.SelectedIndex = 1 // first data row

	newAv, _ := av.handleUp()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 29 {
		t.Errorf("Expected selectedIndex to wrap to 29 (last visible row), got %d", av.SelectedIndex)
	}
}

// TestAssetsView_NavigateEmptyList no-op on empty list
func TestAssetsView_NavigateEmptyList(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{}
	av.SelectedIndex = 0

	newAv, _ := av.handleDown()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 0 {
		t.Errorf("Expected selectedIndex to remain 0 on empty list, got %d", av.SelectedIndex)
	}

	newAv, _ = av.handleUp()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 0 {
		t.Errorf("Expected selectedIndex to remain 0 after up on empty list, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_UpdateEscape returns tea.Quit
func TestAssetsView_UpdateEscape(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyEsc})

	if cmd == nil {
		t.Error("Expected non-nil cmd for Esc key (tea.Quit)")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("Expected tea.QuitMsg, got %T", msg)
	}
}

// TestAssetsView_UpdateCtrlC returns tea.Quit
func TestAssetsView_UpdateCtrlC(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyCtrlC})

	if cmd == nil {
		t.Error("Expected non-nil cmd for Ctrl+C key (tea.Quit)")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("Expected tea.QuitMsg, got %T", msg)
	}
}

// TestAssetsView_Init returns nil command
func TestAssetsView_Init(t *testing.T) {
	av := NewAssetsView()

	cmd := av.Init()

	if cmd != nil {
		t.Errorf("Expected Init to return nil, got %v", cmd)
	}
}

// TestAssetsView_UpdateArrowDown navigates down
func TestAssetsView_UpdateArrowDown(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}, {Ticker: "ETH"}}
	av.SelectedIndex = 0

	_, _ = av.Update(tea.KeyMsg{Type: tea.KeyDown})

	if av.SelectedIndex != 1 {
		t.Errorf("Expected selectedIndex to be 1 after KeyDown, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_UpdateArrowUp navigates up
func TestAssetsView_UpdateArrowUp(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}, {Ticker: "ETH"}}
	av.SelectedIndex = 1

	_, _ = av.Update(tea.KeyMsg{Type: tea.KeyUp})

	// Wrapping from first data row (index 1) to last visible row (index 29)
	if av.SelectedIndex != 29 {
		t.Errorf("Expected selectedIndex to be 29 (wrapped) after KeyUp at index 1, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_UpdateOtherKey no-op on other keys
func TestAssetsView_UpdateOtherKey(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}
	av.SelectedIndex = 0

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})

	if cmd != nil {
		t.Errorf("Expected nil cmd for other keys, got %v", cmd)
	}
	if av.SelectedIndex != 0 {
		t.Errorf("Expected selectedIndex to remain 0, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_UpdateQuitMsg returns tea.Quit
func TestAssetsView_UpdateQuitMsg(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}

	_, cmd := av.Update(tea.QuitMsg{})

	if cmd == nil {
		t.Error("Expected non-nil cmd for QuitMsg (tea.Quit)")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("Expected tea.QuitMsg, got %T", msg)
	}
}

// TestAssetsView_UpdateKeyC returns view transition message
func TestAssetsView_UpdateKeyC(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})

	if cmd == nil {
		t.Error("Expected non-nil cmd for 'c' key (ViewTransitionMsg)")
	}
	msg := cmd()
	if _, ok := msg.(ViewTransitionMsg); !ok {
		t.Errorf("Expected ViewTransitionMsg, got %T", msg)
	}
}

// TestAssetsView_UpdateKeyC_NavigatesToForm
func TestAssetsView_UpdateKeyC_NavigatesToForm(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})

	msg := cmd()
	transitionMsg, ok := msg.(ViewTransitionMsg)
	if !ok {
		t.Fatalf("Expected ViewTransitionMsg, got %T", msg)
	}
	if transitionMsg.View != "form" {
		t.Errorf("Expected View='form', got '%s'", transitionMsg.View)
	}
}

// TestAssetsView_UpdateKeyC_IgnoresCapitalC
func TestAssetsView_UpdateKeyC_IgnoresCapitalC(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{{Ticker: "BTC"}}

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'C'}})

	if cmd != nil {
		t.Errorf("Expected nil cmd for 'C' (capital C), got %v", cmd)
	}
}

// TestAssetsView_UpdateKeyC_EmptyList
func TestAssetsView_UpdateKeyC_EmptyList(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{}

	_, cmd := av.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})

	if cmd == nil {
		t.Error("Expected non-nil cmd for 'c' key on empty list (ViewTransitionMsg)")
	}
	msg := cmd()
	if _, ok := msg.(ViewTransitionMsg); !ok {
		t.Errorf("Expected ViewTransitionMsg, got %T", msg)
	}
}

// TestAssetsView_RenderWith5Assets verifies 30 rows (5 data + 25 empty)
func TestAssetsView_RenderWith5Assets(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC", EntryCount: 1, TotalShares: 0.01, AvgPrice: 50000, TotalValue: 500},
		{Ticker: "ETH", EntryCount: 1, TotalShares: 0.02, AvgPrice: 3000, TotalValue: 60},
		{Ticker: "SOL", EntryCount: 1, TotalShares: 0.03, AvgPrice: 100, TotalValue: 3},
		{Ticker: "ADA", EntryCount: 1, TotalShares: 0.04, AvgPrice: 0.5, TotalValue: 2},
		{Ticker: "DOT", EntryCount: 1, TotalShares: 0.05, AvgPrice: 7.5, TotalValue: 1},
	}
	av.SelectedIndex = 0

	output := av.View()

	// Verify headers are present
	expectedHeaders := []string{"Asset", "Count", "Total Shares", "Avg Price", "Total Value"}
	for _, h := range expectedHeaders {
		if !strings.Contains(output, h) {
			t.Errorf("Expected output to contain header '%s', got: %s", h, output)
		}
	}

	// Verify all 5 assets are present
	for _, ticker := range []string{"BTC", "ETH", "SOL", "ADA", "DOT"} {
		if !strings.Contains(output, ticker) {
			t.Errorf("Expected output to contain '%s', got: %s", ticker, output)
		}
	}

	// Count rows: should have exactly 30 rows (1 header + 5 data + 24 empty)
	// Parse the output to count actual rows
	lines := strings.Split(output, "\n")
	dataRowCount := 0
	for _, line := range lines {
		if strings.Contains(line, "BTC") || strings.Contains(line, "ETH") ||
			strings.Contains(line, "SOL") || strings.Contains(line, "ADA") ||
			strings.Contains(line, "DOT") {
			dataRowCount++
		}
	}

	if dataRowCount != 5 {
		t.Errorf("Expected 5 data rows, got %d", dataRowCount)
	}

	// Verify table border is present (indicating padded table)
	// Table uses lipgloss.RoundedBorder() which produces: ╭╮╰╯
	if !strings.Contains(output, "╭") && !strings.Contains(output, "╰") &&
		!strings.Contains(output, "╮") && !strings.Contains(output, "╯") {
		t.Errorf("Expected table border characters, got: %s", output)
	}
}

// TestAssetsView_RenderWith25Assets verifies 30 rows (25 data + 5 empty)
func TestAssetsView_RenderWith25Assets(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true

	// Create 25 asset entries
	assets := []string{"BTC", "ETH", "SOL", "ADA", "DOT", "XRP", "AVAX", "LINK", "MATIC", "UNI",
		"ALGO", "FIL", "AAVE", "ATOM", "ETC", "HBAR", "NEAR", "SOL", "TIA", "INJ",
		"PEPE", "SHIB", "DOGE", "LTC", "TRX"}
	av.Entries = make([]AssetSummary, len(assets))
	for i, ticker := range assets {
		av.Entries[i] = AssetSummary{
			Ticker:      ticker,
			EntryCount:  1,
			TotalShares: float64(i+1) * 0.01,
			AvgPrice:    float64(i+1) * 100,
			TotalValue:  float64(i+1) * 1000,
		}
	}
	av.SelectedIndex = 0

	output := av.View()

	// Verify headers are present
	expectedHeaders := []string{"Asset", "Count", "Total Shares", "Avg Price", "Total Value"}
	for _, h := range expectedHeaders {
		if !strings.Contains(output, h) {
			t.Errorf("Expected output to contain header '%s', got: %s", h, output)
		}
	}

	// Verify all 25 assets are present
	for _, ticker := range assets {
		if !strings.Contains(output, ticker) {
			t.Errorf("Expected output to contain '%s', got: %s", ticker, output)
		}
	}

	// Count data rows
	dataRowCount := 0
	for _, line := range strings.Split(output, "\n") {
		for _, ticker := range assets {
			if strings.Contains(line, ticker) {
				dataRowCount++
				break
			}
		}
	}

	if dataRowCount != 25 {
		t.Errorf("Expected 25 data rows, got %d", dataRowCount)
	}

	// Verify table border is present (rounded corners)
	if !strings.Contains(output, "╭") && !strings.Contains(output, "╰") &&
		!strings.Contains(output, "╮") && !strings.Contains(output, "╯") {
		t.Errorf("Expected table border characters, got: %s", output)
	}
}

// TestAssetsView_NavigateWithPaddedRows wraps through data and empty rows
func TestAssetsView_NavigateWithPaddedRows(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
	}
	av.SelectedIndex = 29 // Last visible row (after padding)

	// Navigate down through empty rows
	newAv, _ := av.handleDown()
	av = newAv.(*AssetsView)
	// Should wrap to first data row (index 1)
	if av.SelectedIndex != 1 {
		t.Errorf("Expected selectedIndex to wrap to 1, got %d", av.SelectedIndex)
	}

	// Navigate up from first data row (index 1) - should wrap to last row (index 29)
	newAv, _ = av.handleUp()
	av = newAv.(*AssetsView)
	// Should wrap to last visible row (index 29) since selectedIndex <= 1
	if av.SelectedIndex != 29 {
		t.Errorf("Expected selectedIndex to wrap to 29, got %d", av.SelectedIndex)
	}
}

// TestTableLayout_WidthIs100Percent verifies table uses defined column widths for full-width display
func TestTableLayout_WidthIs100Percent(t *testing.T) {
	tests := []struct {
		name       string
		numEntries int
	}{
		{"5 entries", 5},
		{"25 entries", 25},
		{"30 entries", 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := NewAssetsView()
			av.Loaded = true
			av.Entries = make([]AssetSummary, tt.numEntries)
			for i := 0; i < tt.numEntries; i++ {
				av.Entries[i] = AssetSummary{
					Ticker:      fmt.Sprintf("ASSET%02d", i),
					EntryCount:  1,
					TotalShares: 0.01,
					AvgPrice:    50000,
					TotalValue:  500,
				}
			}

			output := av.View()
			lines := strings.Split(output, "\n")

			// Find the table rows (between top and bottom borders, after header)
			var tableLines []string
			inTable := false

			for _, line := range lines {
				// Top border starts the table section
				if strings.Contains(line, "╭") && strings.Contains(line, "────") {
					inTable = true
					continue
				}
				// Bottom border ends the table section
				if strings.Contains(line, "╰") && strings.Contains(line, "────") {
					break
				}
				if inTable && strings.Contains(line, "│") {
					tableLines = append(tableLines, line)
				}
			}

			// Total should be exactly 30 rows (1 header + 29 data/empty)
			if len(tableLines) != 30 {
				t.Errorf("Expected exactly 30 rows (including header), got %d", len(tableLines))
			}

			// Verify each row (except empty header row) has correct width
			// Row width = 74 (includes borders)
			for i, line := range tableLines {
				// Skip the header row (index 0) which has centered text
				if i == 0 && strings.Contains(line, "Asset") && strings.Contains(line, "Count") {
					continue
				}
				// All other rows should have exactly 74 characters (with borders)
				if len(line) != 74 {
					t.Errorf("Row %d has width %d, expected 74: %q", i, len(line), line)
				}
			}
		})
	}
}

// TestTableLayout_HeaderAlignment verifies header alignment with data columns
func TestTableLayout_HeaderAlignment(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC", EntryCount: 3, TotalShares: 0.01, AvgPrice: 50000, TotalValue: 500},
	}

	output := av.View()
	lines := strings.Split(output, "\n")

	// Find header row and first data row within table
	var headerRow, dataRow string
	inTable := false

	for _, line := range lines {
		if strings.Contains(line, "╭") && strings.Contains(line, "────") {
			inTable = true
			continue
		}
		if strings.Contains(line, "╰") && strings.Contains(line, "────") {
			break
		}
		if inTable && strings.Contains(line, "│") {
			if headerRow == "" && strings.Contains(line, "Asset") && strings.Contains(line, "Count") {
				headerRow = line
			} else if headerRow != "" && strings.Contains(line, "BTC") {
				dataRow = line
				break
			}
		}
	}

	if headerRow == "" || dataRow == "" {
		t.Fatalf("Could not find header or data row in table")
	}

	// Parse columns by position based on known column widths
	// The table format is consistent: Asset(10) + "  " + Count(8) + "  " + Shares(14) + "  " + AvgPrice(13) + "  " + TotalValue(13) = 66 chars
	// Plus borders = 68 chars

	// Extract header column values by position
	headerCols := extractColumnsByPosition(headerRow)
	dataCols := extractColumnsByPosition(dataRow)

	if len(headerCols) != 5 {
		t.Errorf("Expected 5 columns in header, got %d", len(headerCols))
	}
	if len(dataCols) != 5 {
		t.Errorf("Expected 5 columns in data row, got %d", len(dataCols))
	}

	// Verify each column has same width (use raw width, not trimmed)
	// Raw width includes trailing spaces from fmt.Sprintf width specifiers
	for i := 0; i < len(headerCols); i++ {
		hWidth := len(headerCols[i])
		dWidth := len(dataCols[i])
		if hWidth != dWidth {
			t.Errorf("Column %d width mismatch: header=%d, data=%d", i, hWidth, dWidth)
		}
	}
}

// extractColumnsByPosition extracts column values from a lipgloss row by position
// Based on known column widths: Asset(10) + "  " + Count(8) + "  " + Shares(14) + "  " + AvgPrice(13) + "  " + TotalValue(13)
func extractColumnsByPosition(row string) []string {
	// Remove borders
	clean := strings.Trim(row, "│")

	// Column positions (0-indexed within clean row)
	// Column 0: Asset, starts at 0, width 10
	// Column 1: Count, starts at 10 + 2 = 12, width 8
	// Column 2: Shares, starts at 12 + 8 + 2 = 22, width 14
	// Column 3: AvgPrice, starts at 22 + 14 + 2 = 38, width 13
	// Column 4: TotalValue, starts at 38 + 13 + 2 = 53, width 13

	col0 := clean[0:10]
	col1 := clean[12:20]
	col2 := clean[22:36]
	col3 := clean[38:51]
	col4 := clean[53:66]

	return []string{col0, col1, col2, col3, col4}
}

// TestTableLayout_Exactly30Rows verifies exactly 30 rows are rendered
func TestTableLayout_Exactly30Rows(t *testing.T) {
	tests := []struct {
		name       string
		numEntries int
	}{
		{"1 entry", 1},
		{"5 entries", 5},
		{"25 entries", 25},
		{"29 entries", 29},
		{"30 entries", 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := NewAssetsView()
			av.Loaded = true
			av.Entries = make([]AssetSummary, tt.numEntries)
			for i := 0; i < tt.numEntries; i++ {
				av.Entries[i] = AssetSummary{
					Ticker:      fmt.Sprintf("ASSET%02d", i),
					EntryCount:  1,
					TotalShares: 0.01,
					AvgPrice:    50000,
					TotalValue:  500,
				}
			}

			output := av.View()
			lines := strings.Split(output, "\n")

			// Count table rows (between borders)
			rowCount := 0
			inTable := false

			for _, line := range lines {
				if strings.Contains(line, "╭") && strings.Contains(line, "────") {
					inTable = true
					continue
				}
				if strings.Contains(line, "╰") && strings.Contains(line, "────") {
					break
				}
				if inTable && strings.Contains(line, "│") {
					rowCount++
				}
			}

			// Total should be exactly 30 (1 header + 29 data/empty)
			if rowCount != 30 {
				t.Errorf("Expected exactly 30 rows, got %d", rowCount)
			}
		})
	}
}

// TestTableLayout_EmptyRowPadding verifies empty rows are rendered correctly
func TestTableLayout_EmptyRowPadding(t *testing.T) {
	tests := []struct {
		name        string
		numEntries  int
		expectedPad int
	}{
		{"1 entry", 1, 28},    // 29 - 1 = 28 empty rows
		{"5 entries", 5, 24},  // 29 - 5 = 24 empty rows
		{"25 entries", 25, 4}, // 29 - 25 = 4 empty rows
		{"29 entries", 29, 0}, // 29 - 29 = 0 empty rows
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := NewAssetsView()
			av.Loaded = true
			av.Entries = make([]AssetSummary, tt.numEntries)
			for i := 0; i < tt.numEntries; i++ {
				av.Entries[i] = AssetSummary{
					Ticker:      fmt.Sprintf("ASSET%02d", i),
					EntryCount:  1,
					TotalShares: 0.01,
					AvgPrice:    50000,
					TotalValue:  500,
				}
			}

			output := av.View()
			lines := strings.Split(output, "\n")

			// Count data rows and empty rows
			dataRowCount := 0
			emptyRowCount := 0
			inTable := false

			for _, line := range lines {
				if strings.Contains(line, "╭") && strings.Contains(line, "────") {
					inTable = true
					continue
				}
				if strings.Contains(line, "╰") && strings.Contains(line, "────") {
					break
				}
				if inTable && strings.Contains(line, "│") {
					// Skip header row (contains "Asset" at start)
					if strings.Contains(line, "Asset") && strings.Contains(line, "Count") {
						continue
					}
					// Check if this is an empty row (has zeros)
					if strings.Contains(line, "0.00000000") || strings.Contains(line, "       0 ") {
						emptyRowCount++
					} else {
						dataRowCount++
					}
				}
			}

			if dataRowCount != tt.numEntries {
				t.Errorf("Expected %d data rows, got %d", tt.numEntries, dataRowCount)
			}
			if emptyRowCount != tt.expectedPad {
				t.Errorf("Expected %d empty rows, got %d", tt.expectedPad, emptyRowCount)
			}
		})
	}
}

// TestTableLayout_ColumnWidthsMatchConstants verifies column widths match defined constants
func TestTableLayout_ColumnWidthsMatchConstants(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC", EntryCount: 3, TotalShares: 0.01, AvgPrice: 50000, TotalValue: 500},
	}

	output := av.View()
	lines := strings.Split(output, "\n")

	// Find header row and extract column widths
	var headerRow string
	inTable := false

	for _, line := range lines {
		if strings.Contains(line, "╭") && strings.Contains(line, "────") {
			inTable = true
			continue
		}
		if strings.Contains(line, "╰") && strings.Contains(line, "────") {
			break
		}
		if inTable && strings.Contains(line, "│") {
			if strings.Contains(line, "Asset") && strings.Contains(line, "Count") {
				headerRow = line
				break
			}
		}
	}

	if headerRow == "" {
		t.Fatal("Could not find header row")
	}

	// Extract column values by position
	cols := extractColumnsByPosition(headerRow)

	if len(cols) != 5 {
		t.Fatalf("Expected 5 columns, got %d: %v", len(cols), cols)
	}

	// Verify column widths match constants
	expectedWidths := []int{
		ColumnAssetWidth,
		ColumnCountWidth,
		ColumnSharesWidth,
		ColumnAvgPriceWidth,
		ColumnTotalValueWidth,
	}

	for i, col := range cols {
		// Use raw width (not trimmed) to match the fmt.Sprintf width specifiers
		if len(col) != expectedWidths[i] {
			t.Errorf("Column %d (%s): expected width %d, got %d", i, col, expectedWidths[i], len(col))
		}
	}
}

// TestTableLayout_RowCountCalculation verifies empty row calculation is correct
func TestTableLayout_RowCountCalculation(t *testing.T) {
	const maxVisibleRows = 30 // 1 header + 29 data rows

	tests := []struct {
		name       string
		numEntries int
	}{
		{"0 entries (will render empty state, not table)", 0},
		{"1 entry", 1},
		{"29 entries (no padding)", 29},
		{"30 entries (capped)", 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			av := NewAssetsView()
			av.Loaded = true
			av.Entries = make([]AssetSummary, tt.numEntries)
			for i := 0; i < tt.numEntries; i++ {
				av.Entries[i] = AssetSummary{
					Ticker:      fmt.Sprintf("ASSET%02d", i),
					EntryCount:  1,
					TotalShares: 0.01,
					AvgPrice:    50000,
					TotalValue:  500,
				}
			}

			output := av.View()
			lines := strings.Split(output, "\n")

			// Check if empty state is rendered (for 0 entries)
			if tt.numEntries == 0 {
				if !strings.Contains(output, "No assets yet") {
					t.Errorf("Expected 'No assets yet' message for empty list")
				}
				return
			}

			// Count table rows
			rowCount := 0
			inTable := false

			for _, line := range lines {
				if strings.Contains(line, "╭") && strings.Contains(line, "────") {
					inTable = true
					continue
				}
				if strings.Contains(line, "╰") && strings.Contains(line, "────") {
					break
				}
				if inTable && strings.Contains(line, "│") {
					rowCount++
				}
			}

			if rowCount != maxVisibleRows {
				t.Errorf("Expected %d rows, got %d", maxVisibleRows, rowCount)
			}
		})
	}
}

// parseLipglossRowColumns parses a lipgloss-formatted row into column parts
func parseLipglossRowColumns(row string) []string {
	// Remove border characters
	clean := strings.Trim(row, "│")
	// Split by separator
	return strings.Split(clean, ColumnSeparator)
}
