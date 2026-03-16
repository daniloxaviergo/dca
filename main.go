package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg, tea.MouseMsg:
		return m, tea.Quit
	case tea.QuitMsg:
		// Quit message received - the program is already terminating
		_, _ = msg, msg
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	// Main container style with enhanced visual presentation
	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		BorderBackground(lipgloss.Color("235")).
		Width(50).
		Height(7).
		Padding(1).
		Margin(1).
		Align(lipgloss.Center)

	// Header title with modern gradient-like styling using multiple colors
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("159")).
		Background(lipgloss.Color("236")).
		Bold(true).
		Underline(true).
		Render("DCA Application")

	// Status line with accent color
	status := lipgloss.NewStyle().
		Foreground(lipgloss.Color("82")).
		Background(lipgloss.Color("236")).
		Render("Visual Enhancement")

	// Footer with decorative separator
	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Render("───")

	// Greeting with enhanced contrast and styling
	greeting := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("236")).
		Bold(true).
		Width(34).
		Align(lipgloss.Center).
		Render("Hello World")

	// Build the layout vertically
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		greeting,
		status,
		footer,
	)

	return lipgloss.JoinHorizontal(lipgloss.Center, container.Render(content))
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
