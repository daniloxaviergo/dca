---
id: GOT-006
title: 'Task 2: Implement CLI input form with prompts'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-16 21:24'
updated_date: '2026-03-16 21:59'
labels: []
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement interactive CLI form using BubbleTea framework for DCA entry collection. Create dca_form.go file.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Display 'Enter DCA Entry' header with visual styling
- [ ] #2 Prompt for amount and accept user input (must be positive number)
- [ ] #3 Prompt for date with default value of today (RFC3339 format)
- [ ] #4 Prompt for asset ticker (e.g., BTC, ETH) and validate as non-empty string
- [ ] #5 Prompt for price per share and accept user input (must be positive number)
- [ ] #6 Calculate and display shares (Amount / Price) with 8 decimal precision
- [ ] #7 Display summary of all entered values for confirmation
- [ ] #8 Support Ctrl+C to cancel without saving
- [ ] #9 Use BubbleTea framework for TUI interaction
- [ ] #10 Handle validation errors with clear error messages and re-prompt
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
# Implementation Plan: CLI Input Form with Prompts

### 1. Technical Approach

Build a BubbleTea-based interactive CLI form using the `teax` package (BubbleTea's higher-level UI primitives) for form input. The form will:

- Create a `formModel` struct to manage form state (current field, values, errors)
- Implement `tea.Model` interface with `Init`, `Update`, `View` methods
- Use `teax/model.Input` for text inputs and `teax/model.DatePicker` or custom date input
- Handle validation in the `Update` loop with error messages displayed inline
- Calculate shares dynamically after amount and price inputs
- Show confirmation screen with all values before final submission
- Implement Ctrl+C handling via `tea.QuitMsg` to exit without saving

**Why this approach:**
- BubbleTea is already a project dependency
- `teax` provides form-specific primitives that reduce boilerplate
- Keep the existing visual style using lipgloss for consistency
- Avoid external survey library to minimize dependencies

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `dca_form.go` | Create | New file containing form model and input handling |
| `main.go` | Modify | Update model to use form flow instead of placeholder UI |

### 3. Dependencies

**Already in go.mod:**
- `github.com/charmbracelet/bubbletea` v1.3.10
- `github.com/charmbracelet/lipgloss` v1.1.0

**May need to add (if teax not available):**
- `github.com/charmbracelet/teax` - Higher-level UI primitives for BubbleTea

### 4. Code Patterns

**Follow existing patterns:**
- Same `model` struct pattern with `tea.Model` interface
- Lipgloss styling for borders, colors, and layout (matching `main.go`)
- Error handling via `errors.New()` and returning errors up the call stack
- 8-decimal precision for shares (from `dca_entry.go`)

**New form-specific patterns:**
- `FormStep` enum to track current input field (amount → date → asset → price → shares → confirm)
- `inputField` struct to manage each input's value, error, and prompt text
- `tea.Batch` command for handling multiple inputs if needed

### 5. Testing Strategy

**Unit tests in `dca_form_test.go`:**
- `TestFormStepForward/Backward` - Verify step transitions
- `TestValidateAmount_Pass/RejectZero/RejectNegative` - Amount validation
- `TestValidateDate_Pass/RejectInvalid` - Date format validation (RFC3339)
- `TestValidateAsset_Pass/RejectEmpty` - Asset ticker validation
- `TestCalculateSharesDisplay` - Verify shares calculation matches entry model

**Integration test in `main_test.go`:**
- `TestFormModel_Init` - Verify initial state
- `TestFormModel_Update_KeyInput` - Test keyboard navigation
- `TestFormModel_Update_Quit` - Verify Ctrl+C exits without save
- `TestFormModel_View_RendersAllFields` - Verify UI shows all prompts

**Manual testing:**
- Run `go run main.go` and verify form appears
- Test each input with valid/invalid data
- Test Ctrl+C during any step
- Verify entry is saved after confirmation

### 6. Risks and Considerations

**Known risks:**
- `teax` may not be available; may need to build custom input components
- BubbleTea's state management for multi-step forms can be verbose
- Date input without picker may require custom validation for RFC3339

**Mitigation strategies:**
- Check if `teax` is available; if not, implement simple input fields with `bufio`-style buffering
- Break form into incremental commits if implementation becomes complex
- Use existing `dca_entry.go` validation functions to avoid duplication

**Trade-offs:**
- Custom date input (user types RFC3339) vs. built-in picker (if available)
- Showing all errors at once vs. one-at-a-time validation
- Single-file implementation vs. splitting form logic across multiple files

**Blocking issues:**
- None identified; implementation is straightforward given existing patterns
<!-- SECTION:PLAN:END -->
