package main

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbletea"
)

func TestViewWithForm(t *testing.T) {
	entries := &DCAData{Entries: make(map[string][]DCAEntry)}
	form := NewFormModel(entries, "test.json")
	m := model{form: form, entries: entries}
	view := m.View()

	// Debug output to see what we're testing
	t.Logf("Raw view output:\n%s", view)

	// Check that "Enter DCA Entry" header is in the output
	if !strings.Contains(view, "Enter DCA Entry") {
		t.Errorf("Expected 'Enter DCA Entry' header in view, got: %s", view)
	}

	// Check that at least one form field prompt is present
	if !strings.Contains(view, "Amount") {
		t.Errorf("Expected 'Amount' field prompt in view, got: %s", view)
	}

	// Check for lipgloss border characters (rounded corners)
	hasRoundedBorder := strings.Contains(view, "╭") && strings.Contains(view, "╮")
	if !hasRoundedBorder {
		t.Errorf("Expected rounded border in view (lipgloss), got: %s", view)
	}
}

func TestUpdateExitOnCtrlC(t *testing.T) {
	entries := &DCAData{Entries: make(map[string][]DCAEntry)}
	form := NewFormModel(entries, "test.json")
	m := model{form: form, entries: entries}

	// Send Ctrl+C
	msg := tea.KeyMsg{
		Runes: []rune{3}, // Ctrl+C
	}
	msg.Type = tea.KeyCtrlC

	newModel, cmd := m.Update(msg)
	m = newModel.(model)

	// Verify tea.Quit command is returned
	if cmd == nil {
		t.Error("Expected tea.Quit command to be returned on Ctrl+C")
	}
}

func TestUpdateExitOnEscape(t *testing.T) {
	entries := &DCAData{Entries: make(map[string][]DCAEntry)}
	form := NewFormModel(entries, "test.json")
	m := model{form: form, entries: entries}

	// Send Escape key
	msg := tea.KeyMsg{
		Runes: []rune{27}, // Escape
	}
	msg.Type = tea.KeyEsc

	newModel, cmd := m.Update(msg)
	m = newModel.(model)

	// Verify tea.Quit command is returned
	if cmd == nil {
		t.Error("Expected tea.Quit command to be returned on Escape")
	}
}

func TestMainForm_Init(t *testing.T) {
	entries := &DCAData{Entries: make(map[string][]DCAEntry)}
	form := NewFormModel(entries, "test.json")
	m := model{form: form, entries: entries}

	// Verify form is initialized
	if m.form == nil {
		t.Error("Expected form to be initialized")
	}
}

func TestMainForm_KeyInput(t *testing.T) {
	entries := &DCAData{Entries: make(map[string][]DCAEntry)}
	form := NewFormModel(entries, "test.json")
	m := model{form: form, entries: entries}

	// Type "100" into the amount field
	inputs := []string{"1", "0", "0"}
	for _, input := range inputs {
		msg := tea.KeyMsg{
			Runes: []rune(input),
		}
		msg.Type = tea.KeyRunes
		newModel, _ := m.Update(msg)
		m = newModel.(model)
	}

	// Verify the amount field contains the input
	if m.form.Fields["amount"].Value != "100" {
		t.Errorf("Expected amount to be '100', got: %s", m.form.Fields["amount"].Value)
	}
}

func TestMainForm_Quit(t *testing.T) {
	entries := &DCAData{Entries: make(map[string][]DCAEntry)}
	form := NewFormModel(entries, "test.json")
	m := model{form: form, entries: entries}

	// Send Ctrl+C
	msg := tea.KeyMsg{
		Runes: []rune{3}, // Ctrl+C
	}
	msg.Type = tea.KeyCtrlC

	newModel, cmd := m.Update(msg)
	m = newModel.(model)

	// Verify tea.Quit command is returned
	if cmd == nil {
		t.Error("Expected tea.Quit command to be returned on Ctrl+C")
	}
}
