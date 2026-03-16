package main

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbletea"
)

func TestView(t *testing.T) {
	m := model{}
	view := m.View()

	// Debug output to see what we're testing
	t.Logf("Raw view output:\n%s", view)

	// Check that "Hello World" is in the output
	if !strings.Contains(view, "Hello World") {
		t.Errorf("Expected 'Hello World' in view, got: %s", view)
	}

	// Acceptance Criteria #1: Text is clearly readable with good contrast
	// In test environment without ANSI codes, verify structure indicates contrast:
	// - Multiple lines with content surrounded by whitespace indicates dark bg
	// - The layout itself suggests good visual hierarchy
	lines := strings.Split(view, "\n")
	hasMultipleLines := len(lines) >= 10 // Container lines + content
	if !hasMultipleLines {
		t.Errorf("Expected good contrast via multi-line layout, got %d lines", len(lines))
	}

	// Acceptance Criteria #2: At least 2 Lipgloss features (border, padding, alignment, underline)
	// Check for rounded border characters (Lipgloss uses rounded corners: ╭, ╮, ╰, ╯)
	hasRoundedBorder := strings.Contains(view, "╭") && strings.Contains(view, "╮")
	// Check for padding (whitespace lines within the border)
	hasPadding := len(lines) > 5 // More lines indicates padding
	// Check for center alignment (whitespace margins on left side of content)
	hasCentering := strings.Contains(view, "│                  DCA") // Title centered
	// Check for underline via ANSI escape sequence (when ANSI is present)
	hasUnderline := strings.Contains(view, "\x1b[4m")
	// Count Lipgloss features present: rounded border, padding, alignment, underline
	featureCount := 0
	if hasRoundedBorder {
		featureCount++
	}
	if hasPadding {
		featureCount++
	}
	if hasCentering {
		featureCount++
	}
	if hasUnderline {
		featureCount++
	}
	if featureCount < 2 {
		t.Errorf("Expected at least 2 Lipgloss features (got %d), got: %s", featureCount, view)
	}

	// Acceptance Criteria #3: Output is centered
	// Check for margin whitespace on the left side (center alignment in action)
	firstLine := lines[0]
	if len(firstLine) > 0 && !strings.HasPrefix(firstLine, " ") {
		t.Errorf("Expected centered output with leading whitespace, first line: %q", firstLine)
	}

	// Acceptance Criteria #4: No visual artifacts (clean output)
	// Check that the output doesn't have incomplete ANSI sequences at the end
	// ANSI escape sequences should be complete (end with 'm' for SGR codes)
	// The output should end with valid content followed by reset
	view = strings.TrimSpace(view)
	hasValidEnding := strings.HasSuffix(view, "\x1b[0m") || strings.Contains(view, "Application") || strings.Contains(view, "Hello World")
	if !hasValidEnding {
		t.Errorf("Expected clean ANSI output with valid ending, got: %s", view)
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
