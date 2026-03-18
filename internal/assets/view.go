package assets

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Column width constants for the Assets View table
// Total table width: 10 + 8 + 14 + 13 + 13 + (4 separators × 2) = 74 characters
const (
	ColumnAssetWidth      = 10   // Asset: 10 characters, left-aligned
	ColumnCountWidth      = 8    // Count: 8 characters, right-aligned
	ColumnSharesWidth     = 14   // Total Shares: 14 characters, right-aligned with 8 decimal places
	ColumnAvgPriceWidth   = 13   // Avg Price: 13 characters, right-aligned with 2 decimal places
	ColumnTotalValueWidth = 13   // Total Value: 13 characters, right-aligned with 2 decimal places
	ColumnSeparator       = "  " // 2 spaces between columns
)

// ViewTransitionMsg is a custom message for switching between views
type ViewTransitionMsg struct {
	View string
}

// AssetsView is a Bubble Tea component for displaying asset data in an interactive table
type AssetsView struct {
	Entries          []AssetSummary
	SelectedIndex    int
	Loaded           bool
	Error            error
	TableHeaderStyle lipgloss.Style
	TableRowStyle    lipgloss.Style
	ActiveRowStyle   lipgloss.Style
}

// Init initializes the AssetsView component
func (a *AssetsView) Init() tea.Cmd {
	return nil
}

// Update handles user input for the AssetsView component
func (a *AssetsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return a, tea.Quit
		case tea.KeyRunes:
			if string(msg.Runes) == "c" {
				return a, func() tea.Msg {
					return ViewTransitionMsg{View: "form"}
				}
			}
		case tea.KeyUp:
			return a.handleUp()
		case tea.KeyDown:
			return a.handleDown()
		}
	case tea.QuitMsg:
		return a, tea.Quit
	case ViewTransitionMsg:
		if msg.View == "form" {
			return a, func() tea.Msg {
				return ViewTransitionMsg{View: "form"}
			}
		}
	}
	return a, nil
}

// handleUp moves selection up (with wrap-around from first data row to last visible row)
func (a *AssetsView) handleUp() (tea.Model, tea.Cmd) {
	if len(a.Entries) == 0 {
		return a, nil
	}

	// Total visible rows = 30 (1 header + up to 29 data/empty rows)
	// Data rows are at indices 1-29 (0 is header)
	// Selection index ranges from 0 (header) to 29 (last data/empty row)
	const maxRowIndex = 29 // 30 total rows - 1

	if a.SelectedIndex <= 1 {
		// Wrap from first data row (index 1) to last visible row (index 29)
		// Also handles index 0 (header) wrapping to last row
		a.SelectedIndex = maxRowIndex
	} else {
		a.SelectedIndex--
	}
	return a, nil
}

// handleDown moves selection down (with wrap-around to first data row)
func (a *AssetsView) handleDown() (tea.Model, tea.Cmd) {
	if len(a.Entries) == 0 {
		return a, nil
	}

	// Total visible rows = 30 (1 header + up to 29 data/empty rows)
	// Data rows are at indices 1-29 (0 is header)
	// Selection index ranges from 0 (header) to 29 (last data/empty row)
	const maxRowIndex = 29 // 30 total rows - 1

	if a.SelectedIndex >= maxRowIndex {
		// Wrap to first data row (index 1, skipping header at 0)
		a.SelectedIndex = 1
	} else {
		a.SelectedIndex++
	}
	return a, nil
}

// View renders the AssetsView component
func (a *AssetsView) View() string {
	var sb strings.Builder

	sb.WriteString(a.renderHeader())
	sb.WriteString("\n\n")
	sb.WriteString(a.renderTable())
	sb.WriteString("\n\n")
	sb.WriteString(a.renderFooter())

	return sb.String()
}

// renderHeader displays the view title
func (a *AssetsView) renderHeader() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("159")).
		Background(lipgloss.Color("236")).
		Bold(true).
		Underline(true).
		Render("DCA Assets List")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("View your aggregated investment data")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		subtitle,
	)
}

// renderTable displays the asset data in a table format
func (a *AssetsView) renderTable() string {
	if !a.Loaded {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("Loading data...")
	}

	if a.Error != nil {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Render(fmt.Sprintf("❌ Error loading data: %v", a.Error))
	}

	if len(a.Entries) == 0 {
		return a.renderEmptyState()
	}

	// Cap displayed rows to maintain exactly 30 total rows (1 header + 29 data)
	// Only display first maxDataRows entries
	const maxVisibleRows = 30
	const maxDataRows = maxVisibleRows - 1 // 29 data rows max
	dataRowsToRender := len(a.Entries)
	if dataRowsToRender > maxDataRows {
		dataRowsToRender = maxDataRows
	}

	var rows []string

	// Render header row
	header := a.renderHeaderRow()
	rows = append(rows, header)

	// Render data rows (up to maxDataRows)
	for i := 0; i < dataRowsToRender; i++ {
		entry := a.Entries[i]
		row := a.renderDataRow(i, entry)
		rows = append(rows, row)
	}

	// Pad with empty rows to maintain exactly 30 rows (including header)
	// 30 = 1 header + data rows + empty rows
	// empty rows = 30 - 1 - data rows rendered
	emptyRowsNeeded := maxVisibleRows - 1 - dataRowsToRender
	if emptyRowsNeeded > 0 {
		for i := 0; i < emptyRowsNeeded; i++ {
			row := a.renderEmptyDataRow(dataRowsToRender + i)
			rows = append(rows, row)
		}
	}

	// Create table with border
	tableStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	return tableStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

// renderHeaderRow renders the table header with column names
func (a *AssetsView) renderHeaderRow() string {
	headers := []string{"Asset", "Count", "Total Shares", "Avg Price", "Total Value"}
	formatted := []string{
		fmt.Sprintf("%-*s", ColumnAssetWidth, headers[0]),
		fmt.Sprintf("%*s", ColumnCountWidth, headers[1]),
		fmt.Sprintf("%*s", ColumnSharesWidth, headers[2]),
		fmt.Sprintf("%*s", ColumnAvgPriceWidth, headers[3]),
		fmt.Sprintf("%*s", ColumnTotalValueWidth, headers[4]),
	}

	var styled []string
	for _, f := range formatted {
		styled = append(styled, lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Render(f))
	}

	return strings.Join(styled, ColumnSeparator)
}

// renderDataRow renders a single data row with the specified index
func (a *AssetsView) renderDataRow(index int, entry AssetSummary) string {
	// Format each cell with fixed widths
	// Asset: left-aligned text
	// Count: right-aligned integer
	// Total Shares: right-aligned with 8 decimal places
	// Avg Price: right-aligned with 2 decimal places
	// Total Value: right-aligned with 2 decimal places
	rowData := []string{
		fmt.Sprintf("%-*s", ColumnAssetWidth, entry.Ticker),
		fmt.Sprintf("%*d", ColumnCountWidth, entry.EntryCount),
		fmt.Sprintf("%*.8f", ColumnSharesWidth, entry.TotalShares),
		fmt.Sprintf("%*.2f", ColumnAvgPriceWidth, entry.AvgPrice),
		fmt.Sprintf("%*.2f", ColumnTotalValueWidth, entry.TotalValue),
	}

	rowStr := strings.Join(rowData, ColumnSeparator)

	// Apply active row styling
	if index == a.SelectedIndex {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("63")).
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Padding(0, 1).
			Render(rowStr)
	}

	return lipgloss.NewStyle().
		Padding(0, 1).
		Render(rowStr)
}

// renderEmptyState displays message when no assets exist
func (a *AssetsView) renderEmptyState() string {
	msg := "No assets yet"
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(1).
		Render(msg)
}

// renderEmptyDataRow renders an empty row for padding (when data rows < 30)
func (a *AssetsView) renderEmptyDataRow(index int) string {
	// Format each cell with fixed widths using empty/zeros
	rowData := []string{
		fmt.Sprintf("%-*s", ColumnAssetWidth, ""),
		fmt.Sprintf("%*s", ColumnCountWidth, ""),
		fmt.Sprintf("%*.8f", ColumnSharesWidth, 0.0),
		fmt.Sprintf("%*.2f", ColumnAvgPriceWidth, 0.0),
		fmt.Sprintf("%*.2f", ColumnTotalValueWidth, 0.0),
	}

	rowStr := strings.Join(rowData, ColumnSeparator)

	// Apply active row styling
	if index == a.SelectedIndex {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("63")).
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Padding(0, 1).
			Render(rowStr)
	}

	return lipgloss.NewStyle().
		Padding(0, 1).
		Render(rowStr)
}

// renderFooter displays navigation help text
func (a *AssetsView) renderFooter() string {
	help := []string{"[↑/↓] Navigate", "[Esc/Ctrl+C] Exit"}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(strings.Join(help, "  "))
}

// NewAssetsView creates a new AssetsView component
func NewAssetsView() *AssetsView {
	return &AssetsView{
		SelectedIndex: 0,
		Loaded:        false,
		Error:         nil,
		TableHeaderStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true),
		TableRowStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")),
		ActiveRowStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("63")).
			Foreground(lipgloss.Color("15")).
			Bold(true),
	}
}
