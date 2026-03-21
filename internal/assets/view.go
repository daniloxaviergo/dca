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

// Modal width constants
const (
	ModalWidth              = 60
	ModalDateWidth          = 12 // YYYY-MM-DD
	ModalAvgPriceWidth      = 12 // Right-aligned with 2 decimals
	ModalTotalInvestedWidth = 14 // Right-aligned with 2 decimals
	ModalEntryCountWidth    = 10 // Right-aligned
	ModalDateSeparator      = "  "
)

// ViewTransitionMsg is a custom message for switching between views
type ViewTransitionMsg struct {
	View string
}

// OpenModalMsg is a message to open the asset history modal
type OpenModalMsg struct {
	AssetTicker string
}

// CloseModalMsg is a message to close the asset history modal
type CloseModalMsg struct{}

// LoadMoreMsg is a message to load the next batch of data for the modal
type LoadMoreMsg struct{}

// AssetsView is a Bubble Tea component for displaying asset data in an interactive table
type AssetsView struct {
	Entries          []AssetSummary
	SelectedIndex    int
	Loaded           bool
	Error            error
	Filename         string // Path to the data file for modal operations
	TableHeaderStyle lipgloss.Style
	TableRowStyle    lipgloss.Style
	ActiveRowStyle   lipgloss.Style
	Modal            *AssetHistoryModal
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
		case tea.KeyCtrlC:
			return a, tea.Quit
		case tea.KeyEsc:
			// Handle Escape key differently based on modal state
			if a.Modal != nil && a.Modal.Visible {
				// Close modal when Escape pressed in modal view
				return a.handleCloseModal()
			}
			// Quit app when Escape pressed in list view
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
		case tea.KeyEnter:
			// Handle Enter differently based on modal state
			if a.Modal != nil && a.Modal.Visible {
				// Load more when Enter is pressed in modal view
				return a.handleLoadMore()
			}
			// Open modal when Enter is pressed on a data row
			// Skip header row (index 0)
			if a.SelectedIndex > 0 && a.SelectedIndex <= len(a.Entries) {
				return a.handleEnterOnAsset()
			}
		}
	case tea.QuitMsg:
		return a, tea.Quit
	case ViewTransitionMsg:
		if msg.View == "form" {
			return a, func() tea.Msg {
				return ViewTransitionMsg{View: "form"}
			}
		}
	case OpenModalMsg:
		return a.handleOpenModal(msg.AssetTicker)
	case CloseModalMsg:
		return a.handleCloseModal()
	case LoadMoreMsg:
		return a.handleLoadMore()
	}
	return a, nil
}

// handleUp moves selection up (with wrap-around from header to last visible row)
func (a *AssetsView) handleUp() (tea.Model, tea.Cmd) {
	if len(a.Entries) == 0 {
		return a, nil
	}

	// Total visible rows = 30 (1 header + up to 29 data/empty rows)
	// Data rows are at indices 1-29 (0 is header)
	// Selection index ranges from 0 (header) to 29 (last data/empty row)
	const maxRowIndex = 29 // 30 total rows - 1

	if a.SelectedIndex == 0 {
		// Wrap from header (index 0) to last visible row (index 29)
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

// handleEnterOnAsset opens the modal for the selected asset
func (a *AssetsView) handleEnterOnAsset() (tea.Model, tea.Cmd) {
	if a.SelectedIndex <= 0 || a.SelectedIndex > len(a.Entries) {
		return a, nil
	}

	asset := a.Entries[a.SelectedIndex-1] // Adjust for header row at index 0

	// Initialize modal if nil
	if a.Modal == nil {
		a.Modal = NewAssetHistoryModal()
	}

	// Load data into modal for the selected asset
	err := a.Modal.LoadData(a.Filename, asset.Ticker)
	if err != nil {
		a.Modal.Error = err
		a.Modal.Loaded = true
	} else {
		a.Modal.Loaded = true
	}

	a.Modal.Visible = true
	return a, nil
}

// handleCloseModal closes the modal and returns to list view
func (a *AssetsView) handleCloseModal() (tea.Model, tea.Cmd) {
	a.Modal.Visible = false
	return a, nil
}

// handleLoadMore loads the next batch of data for the modal
func (a *AssetsView) handleLoadMore() (tea.Model, tea.Cmd) {
	if a.Modal == nil || !a.Modal.Visible {
		return a, nil
	}

	// Load more data using the modal's stored filename
	err := a.Modal.LoadMore(a.Modal.Filename)
	if err != nil {
		a.Modal.Error = err
	}

	return a, nil
}

// handleOpenModal opens the modal for the specified asset
func (a *AssetsView) handleOpenModal(assetTicker string) (tea.Model, tea.Cmd) {
	// Initialize modal if nil
	if a.Modal == nil {
		a.Modal = NewAssetHistoryModal()
	}

	// Load data into modal
	err := a.Modal.LoadData(a.Filename, assetTicker)
	if err != nil {
		a.Modal.Error = err
		a.Modal.Loaded = true
	} else {
		a.Modal.Loaded = true
	}

	a.Modal.Visible = true
	return a, nil
}

// View renders the AssetsView component
func (a *AssetsView) View() string {
	// If modal is visible, render the modal instead of the list
	if a.Modal != nil && a.Modal.Visible {
		return a.renderModal()
	}

	var sb strings.Builder

	sb.WriteString(a.renderHeader())
	sb.WriteString("\n\n")
	sb.WriteString(a.renderTable())
	sb.WriteString("\n\n")
	sb.WriteString(a.renderFooter())

	return sb.String()
}

// renderModal renders the asset history modal
func (a *AssetsView) renderModal() string {
	if a.Modal == nil || !a.Modal.Visible {
		return ""
	}

	// Modal container with centered styling
	modalStyle := lipgloss.NewStyle().
		Width(ModalWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Align(lipgloss.Center).
		Padding(1, 2)

	// Title
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("159")).
		Bold(true).
		Render(fmt.Sprintf("Asset History: %s", a.Modal.AssetTicker))

	var content string
	if !a.Modal.Loaded {
		content = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("Loading history...")
	} else if a.Modal.Error != nil {
		content = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Render(fmt.Sprintf("❌ Error loading history: %v", a.Modal.Error))
	} else if len(a.Modal.EntriesByDate) == 0 {
		content = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("No history for this asset")
	} else {
		content = a.renderModalContent()
	}

	// Combine title and content
	body := lipgloss.JoinVertical(lipgloss.Center, title, content)

	// Add footer instructions based on state
	var footerText []string
	if a.Modal.AllLoaded {
		footerText = append(footerText, "All data loaded")
	} else if a.Modal.Loading {
		footerText = append(footerText, "Loading more...")
	} else {
		footerText = append(footerText, "[Esc] Close Modal")
		footerText = append(footerText, "[Enter] Load More")
	}

	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(strings.Join(footerText, "  "))

	return modalStyle.Render(lipgloss.JoinVertical(lipgloss.Center, body, footer))
}

// renderModalContent renders the modal data table
func (a *AssetsView) renderModalContent() string {
	var rows []string

	// Header row
	header := a.renderModalHeaderRow()
	rows = append(rows, header)

	// Data rows
	for i, entry := range a.Modal.EntriesByDate {
		row := a.renderModalDataRow(i, entry)
		rows = append(rows, row)
	}

	// Create table with border
	tableStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	return tableStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

// renderModalHeaderRow renders the modal header row
func (a *AssetsView) renderModalHeaderRow() string {
	headers := []string{"Date", "Avg Price", "Total Invested", "Entry Count"}
	formatted := []string{
		fmt.Sprintf("%-*s", ModalDateWidth, headers[0]),
		fmt.Sprintf("%*s", ModalAvgPriceWidth, headers[1]),
		fmt.Sprintf("%*s", ModalTotalInvestedWidth, headers[2]),
		fmt.Sprintf("%*s", ModalEntryCountWidth, headers[3]),
	}

	var styled []string
	for _, f := range formatted {
		styled = append(styled, lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Render(f))
	}

	return strings.Join(styled, ModalDateSeparator)
}

// renderModalDataRow renders a single modal data row
func (a *AssetsView) renderModalDataRow(index int, entry EntryByDate) string {
	rowData := []string{
		fmt.Sprintf("%-*s", ModalDateWidth, entry.Date),
		fmt.Sprintf("%*.2f", ModalAvgPriceWidth, entry.WeightedAvgPrice),
		fmt.Sprintf("%*.2f", ModalTotalInvestedWidth, entry.TotalInvested),
		fmt.Sprintf("%*d", ModalEntryCountWidth, entry.EntryCount),
	}

	rowStr := strings.Join(rowData, ModalDateSeparator)

	return lipgloss.NewStyle().
		Padding(0, 1).
		Render(rowStr)
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
		Filename:      "dca_entries.json",
		Modal:         NewAssetHistoryModal(),
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
