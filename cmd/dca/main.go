package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/charmbracelet/bubbletea"

	"github.com/danilo/scripts/github/dca/internal/assets"
	"github.com/danilo/scripts/github/dca/internal/dca"
	"github.com/danilo/scripts/github/dca/internal/form"
)

// Default path for DCA entries JSON file
const defaultEntriesPath = "dca_entries.json"

// AppState represents the current view state of the application
type AppState int

const (
	StateForm AppState = iota
	StateAssetsView
)

// formSubmittedMsg is sent when the form is successfully submitted
type formSubmittedMsg struct{}

type model struct {
	form         *form.FormModel
	assetsView   *assets.AssetsView
	entries      *dca.DCAData
	currentState AppState
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	if m.currentState == StateForm && m.form != nil {
		return m.form.Init()
	}
	if m.currentState == StateAssetsView && m.assetsView != nil {
		return m.assetsView.Init()
	}
	return nil
}

// Update handles state transitions and delegates to current view
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentState {
	case StateForm:
		if m.form != nil {
			newForm, cmd := m.form.Update(msg)
			m.form = newForm.(*form.FormModel)
			// Check for form submission or view transition from form
			if _, ok := msg.(formSubmittedMsg); ok {
				// After form submission, switch to assets view
				m.currentState = StateAssetsView
				m.assetsView = assets.NewAssetsView()
				// Load data into assets view
				vm, err := assets.LoadAndAggregateEntries(defaultEntriesPath)
				if err != nil {
					m.assetsView.Error = err
				} else {
					m.assetsView.Entries = vm.Entries
					m.assetsView.Loaded = true
				}
				return m, nil
			}
			if _, ok := msg.(assets.ViewTransitionMsg); ok {
				// Handle view transition from form (e.g., Ctrl+C during form)
				return m, tea.Quit
			}
			return m, cmd
		}

	case StateAssetsView:
		if m.assetsView != nil {
			newAssetsView, cmd := m.assetsView.Update(msg)
			m.assetsView = newAssetsView.(*assets.AssetsView)
			// On exit (Esc/Ctrl+C), switch back to form
			if _, ok := msg.(assets.ViewTransitionMsg); ok {
				m.currentState = StateForm
				m.form = form.NewFormModel(m.entries, defaultEntriesPath)
				return m, nil
			}
			return m, cmd
		}
	}

	return m, nil
}

// View renders the current state
func (m model) View() string {
	switch m.currentState {
	case StateForm:
		if m.form != nil {
			return m.form.View()
		}
	case StateAssetsView:
		if m.assetsView != nil {
			return m.assetsView.View()
		}
	}
	return "Loading..."
}

func main() {
	// Panic recovery wrapper for the entire application
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Panic recovered: %v\n", r)
			fmt.Fprintf(os.Stderr, "Stack trace:\n%s\n", debug.Stack())
			os.Exit(1)
		}
	}()

	// Load existing entries
	entries, err := dca.LoadEntries(defaultEntriesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading entries: %v\n", err)
		os.Exit(1)
	}

	// Create form model and initialize with entries
	form := form.NewFormModel(entries, defaultEntriesPath)

	// Create initial model with form state
	m := model{
		form:         form,
		entries:      entries,
		currentState: StateForm,
	}

	// Run the program
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Save entries after form submission
	if form.Submitted {
		if err := dca.SaveEntries(defaultEntriesPath, entries); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving entries: %v\n", err)
			os.Exit(1)
		}
	}
}
