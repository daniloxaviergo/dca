---
id: GOT-058
title: '[doc-013] Phase 3: Implement CLI runCLI() Function'
status: To Do
assignee: []
created_date: '2026-03-28 15:18'
labels:
  - cli
  - implementation
  - persistence
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the runCLI() function that orchestrates validation, data persistence, and error handling for the command-line quick entry feature.

## Phase Context
- **Objective**: Create the core CLI workflow function that processes arguments and persists the entry
- **Deliverables**: 
  - runCLI() function with full workflow
  - Validation sequence (asset, amount, price, date)
  - Shares calculation using shared logic
  - File I/O using existing DCA persistence layer
  - Proper exit codes (0 for success, 1 for errors)
- **Stakeholders**: Developers (implementation), QA Team (functionality)
- **Dependencies**: Phase 1 (validation) and Phase 2 (flag parsing) completed
- **Constraints**: Must use atomic file writes, maintain 8 decimal precision for shares

## Task Generation Rules
- Break down into logical sub-steps (validate, calculate, persist)
- Each step should have its own task or be grouped logically
- Trace validation to specific validation functions from Phase 1
- Include error handling for file operations (permissions, disk space)
- Consider edge cases (empty entries file, malformed JSON)

## Acceptance Criteria
1. runCLI() function successfully processes valid entries
2. Shares calculated with 8 decimal precision using shared logic
3. File persistence uses atomic write pattern (temp file + rename)
4. Exit code 0 on success, 1 on any failure
5. No output on successful entry (silent success)
6. Error messages match validation format from Phase 1
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
