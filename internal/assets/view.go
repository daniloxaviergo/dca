package assets

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// handleUp moves selection up (with wrap-around)
func (a *AssetsView) handleUp() (tea.Model, tea.Cmd) {
	if len(a.Entries) == 0 {
		return a, nil
	}
	if a.SelectedIndex <= 0 {
		// Wrap to last row
		a.SelectedIndex = len(a.Entries) - 1
	} else {
		a.SelectedIndex--
	}
	return a, nil
}

// handleDown moves selection down (with wrap-around)
func (a *AssetsView) handleDown() (tea.Model, tea.Cmd) {
	if len(a.Entries) == 0 {
		return a, nil
	}
	if a.SelectedIndex >= len(a.Entries)-1 {
		// Wrap to first row
		a.SelectedIndex = 0
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

	var rows []string

	// Render header row
	header := a.renderHeaderRow()
	rows = append(rows, header)

	// Render data rows
	for i, entry := range a.Entries {
		row := a.renderDataRow(i, entry)
		rows = append(rows, row)
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
	var formatted []string

	for _, h := range headers {
		formatted = append(formatted, lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Render(h))
	}

	return strings.Join(formatted, "  ")
}

// renderDataRow renders a single data row with the specified index
func (a *AssetsView) renderDataRow(index int, entry AssetSummary) string {
	// Format the row data
	rowData := []string{
		entry.Ticker,
		fmt.Sprintf("%d", entry.EntryCount),
		fmt.Sprintf("%.8f", entry.TotalShares),
		fmt.Sprintf("%.2f", entry.AvgPrice),
		fmt.Sprintf("%.2f", entry.TotalValue),
	}

	// Format each cell
	var formatted []string
	for _, cell := range rowData {
		formatted = append(formatted, cell)
	}

	rowStr := strings.Join(formatted, "  ")

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
