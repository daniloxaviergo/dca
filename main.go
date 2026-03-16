package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
)

// Default path for DCA entries JSON file
const defaultEntriesPath = "dca_entries.json"

type model struct {
	form   *FormModel
	entries *DCAData
}

func (m model) Init() tea.Cmd {
	if m.form != nil {
		return m.form.Init()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.form != nil {
		newForm, cmd := m.form.Update(msg)
		m.form = newForm.(*FormModel)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.form != nil {
		return m.form.View()
	}
	return "Loading..."
}

func main() {
	// Load existing entries
	entries, err := LoadEntries(defaultEntriesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading entries: %v\n", err)
		os.Exit(1)
	}

	// Create form model
	form := NewFormModel(entries, defaultEntriesPath)

	// Run the program
	p := tea.NewProgram(model{form: form, entries: entries})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Save entries after form submission
	if form.Submitted {
		if err := SaveEntries(defaultEntriesPath, entries); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving entries: %v\n", err)
			os.Exit(1)
		}
	}
}
