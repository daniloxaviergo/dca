package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/charmbracelet/bubbletea"

	"github.com/danilo/scripts/github/dca/internal/dca"
	"github.com/danilo/scripts/github/dca/internal/form"
)

// TestFormSubmittedMsg_TransitionsToAssetsView tests that form submission triggers state transition
func TestFormSubmittedMsg_TransitionsToAssetsView(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_entries.json")

	// Create initial entries
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}

	// Create form model
	formModel := form.NewFormModel(entries, testFile)

	// Set up form for submission
	formModel.Fields["amount"].Value = "100"
	formModel.Fields["date"].Value = "2024-01-15T10:30:00Z"
	formModel.Fields["asset"].Value = "BTC"
	formModel.Fields["price"].Value = "50000"
	formModel.Fields["confirm"].Value = "y"
	formModel.Step = form.StepConfirm

	// Simulate pressing Enter to submit
	_, cmd := formModel.Update(tea.KeyMsg{Type: tea.KeyEnter})

	// Check that command returns FormSubmittedMsg
	if cmd == nil {
		t.Fatal("Expected cmd to return FormSubmittedMsg, got nil")
	}

	// Execute the command and check message type
	msg := cmd()
	if _, ok := msg.(form.FormSubmittedMsg); !ok {
		t.Errorf("Expected FormSubmittedMsg, got %T", msg)
	}

	// Verify form was marked as submitted
	if !formModel.Submitted {
		t.Error("Expected form.Submitted to be true after submission")
	}
}

// TestFormSubmittedMsg_ReloadsData tests that assets view reloads data after form submission
func TestFormSubmittedMsg_ReloadsData(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_entries.json")

	// Create initial entries
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}

	// Create form model
	formModel := form.NewFormModel(entries, testFile)

	// Set up form for submission
	formModel.Fields["amount"].Value = "100"
	formModel.Fields["date"].Value = "2024-01-15T10:30:00Z"
	formModel.Fields["asset"].Value = "BTC"
	formModel.Fields["price"].Value = "50000"
	formModel.Fields["confirm"].Value = "y"
	formModel.Step = form.StepConfirm

	// Simulate pressing Enter to submit
	_, cmd := formModel.Update(tea.KeyMsg{Type: tea.KeyEnter})

	// Execute the command to send FormSubmittedMsg
	msg := cmd()
	if _, ok := msg.(form.FormSubmittedMsg); !ok {
		t.Fatalf("Expected FormSubmittedMsg, got %T", msg)
	}

	// Verify the entry was saved to file
	if len(entries.Entries) != 1 {
		t.Errorf("Expected 1 asset entry, got %d", len(entries.Entries))
	}

	if _, ok := entries.Entries["BTC"]; !ok {
		t.Error("Expected BTC entry to be created")
	}

	if len(entries.Entries["BTC"]) != 1 {
		t.Errorf("Expected 1 BTC entry, got %d", len(entries.Entries["BTC"]))
	}

	// Verify file was created
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Error("Expected test file to be created")
	}
}

// TestFormEntrySavesToSharedReference tests that form saves to shared entries reference
func TestFormEntrySavesToSharedReference(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_entries.json")

	// Create initial entries
	entries := &dca.DCAData{Entries: make(map[string][]dca.DCAEntry)}

	// Create form model with shared reference
	formModel := form.NewFormModel(entries, testFile)

	// Set up form for submission
	formModel.Fields["amount"].Value = "100"
	formModel.Fields["date"].Value = "2024-01-15T10:30:00Z"
	formModel.Fields["asset"].Value = "ETH"
	formModel.Fields["price"].Value = "3000"
	formModel.Fields["confirm"].Value = "y"
	formModel.Step = form.StepConfirm

	// Simulate pressing Enter to submit
	_, _ = formModel.Update(tea.KeyMsg{Type: tea.KeyEnter})

	// Verify entries were modified in-place
	if len(entries.Entries) != 1 {
		t.Errorf("Expected entries to be modified in-place, got %d assets", len(entries.Entries))
	}

	if _, ok := entries.Entries["ETH"]; !ok {
		t.Error("Expected ETH entry to be created in shared reference")
	}

	if len(entries.Entries["ETH"]) != 1 {
		t.Errorf("Expected 1 ETH entry, got %d", len(entries.Entries["ETH"]))
	}
}
