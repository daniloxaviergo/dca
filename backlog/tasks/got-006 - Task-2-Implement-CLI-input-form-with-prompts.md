---
id: GOT-006
title: 'Task 2: Implement CLI input form with prompts'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-16 21:24'
updated_date: '2026-03-16 21:58'
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
