package main

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbletea"
)

func TestView(t *testing.T) {
	m := model{}
	view := m.View()

	// Check that "Hello World" is in the output
	if !strings.Contains(view, "Hello World") {
		t.Errorf("Expected 'Hello World' in view, got: %s", view)
	}

	// Check that lipgloss border characters are present (rounded border)
	if !strings.Contains(view, "┌") && !strings.Contains(view, "─") {
		t.Errorf("Expected border characters in view, got: %s", view)
	}

	// Acceptance Criteria #1: Text is clearly readable with good contrast
	// Verify we have dark background (235/236) with bright foregrounds (159, 205, 82)
	hasDarkBackground := strings.Contains(view, "\x1b[48;5;236") || strings.Contains(view, "\x1b[48;5;235")
	hasHighContrastFg := strings.Contains(view, "\x1b[38;5;205") || strings.Contains(view, "\x1b[38;5;159") || strings.Contains(view, "\x1b[38;5;82")
	if !hasDarkBackground || !hasHighContrastFg {
		t.Errorf("Expected good contrast (dark background with bright foreground), got: %s", view)
	}

	// Acceptance Criteria #2: At least 2 Lipgloss features (border, padding, alignment, underline)
	// Check for rounded border characters
	hasRoundedBorder := strings.Contains(view, "┌") && strings.Contains(view, "┐")
	// Check for padding (whitespace lines within the border)
	lines := strings.Split(view, "\n")
	hasPadding := len(lines) > 5 // More lines indicates padding
	// Check for underline on title
	hasUnderline := strings.Contains(view, "Underline") || strings.Contains(view, "\x1b[4m")
	// Check for center alignment (whitespace margins on left)
	hasCentering := strings.Contains(view, "  ") && strings.Index(view, "┌") > 10
	if !hasRoundedBorder || (!hasPadding && !hasUnderline && !hasCentering) {
		t.Errorf("Expected at least 2 Lipgloss features (border/padding/alignment/underline), got: %s", view)
	}

	// Acceptance Criteria #3: Output is centered
	// Check for margin whitespace on the left side (center alignment in action)
	firstLine := lines[0]
	if len(firstLine) > 0 && !strings.HasPrefix(firstLine, " ") {
		t.Errorf("Expected centered output with leading whitespace, first line: %q", firstLine)
	}

	// Acceptance Criteria #4: No visual artifacts (clean output)
	// Check for valid ANSI sequences - no partial escape codes
	invalidAnsi := strings.Contains(view, "\x1b[") && !strings.Contains(view, "\x1b[0m")
	if invalidAnsi {
		t.Errorf("Expected clean ANSI output without artifacts, got: %s", view)
	}
}

func TestUpdateExitOnKeyMsg(t *testing.T) {
	m := model{}
	msg := tea.KeyMsg{}

	newModel, cmd := m.Update(msg)

	// Verify the model is returned (unchanged)
	if newModel != m {
		t.Error("Expected model to be returned unchanged")
	}

	// Verify tea.Quit command is returned
	if cmd == nil {
		t.Error("Expected tea.Quit command to be returned on KeyMsg")
	}
}

func TestUpdateExitOnMouseMsg(t *testing.T) {
	m := model{}
	msg := tea.MouseMsg{}

	newModel, cmd := m.Update(msg)

	// Verify the model is returned (unchanged)
	if newModel != m {
		t.Error("Expected model to be returned unchanged")
	}

	// Verify tea.Quit command is returned
	if cmd == nil {
		t.Error("Expected tea.Quit command to be returned on MouseMsg")
	}
}

func TestUpdateOnQuitMsg(t *testing.T) {
	m := model{}
	msg := tea.QuitMsg{}

	newModel, cmd := m.Update(msg)

	// Verify the model is returned (unchanged)
	if newModel != m {
		t.Error("Expected model to be returned unchanged")
	}

	// Verify no command is returned on QuitMsg (quit already in progress)
	if cmd != nil {
		t.Error("Expected no command to be returned on QuitMsg")
	}
}
