package assets

import (
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

	newAv, _ = av.handleUp()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 0 {
		t.Errorf("Expected selectedIndex to be 0 after up, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_NavigateWrapDown wraps from last to first
func TestAssetsView_NavigateWrapDown(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
	}
	av.SelectedIndex = 1 // last row

	newAv, _ := av.handleDown()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 0 {
		t.Errorf("Expected selectedIndex to wrap to 0, got %d", av.SelectedIndex)
	}
}

// TestAssetsView_NavigateWrapUp wraps from first to last
func TestAssetsView_NavigateWrapUp(t *testing.T) {
	av := NewAssetsView()
	av.Loaded = true
	av.Entries = []AssetSummary{
		{Ticker: "BTC"},
		{Ticker: "ETH"},
	}
	av.SelectedIndex = 0 // first row

	newAv, _ := av.handleUp()
	av = newAv.(*AssetsView)
	if av.SelectedIndex != 1 {
		t.Errorf("Expected selectedIndex to wrap to 1, got %d", av.SelectedIndex)
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

	if av.SelectedIndex != 0 {
		t.Errorf("Expected selectedIndex to be 0 after KeyUp, got %d", av.SelectedIndex)
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
