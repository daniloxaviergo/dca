---
id: GOT-057
title: '[doc-013] Phase 2: Create CLI Module cmd/dca/cli.go'
status: To Do
assignee: []
created_date: '2026-03-28 15:18'
labels:
  - cli
  - cli-module
  - implementation
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the CLI-specific entry point with flag parsing logic for the command-line quick entry feature.

## Phase Context
- **Objective**: Implement a standalone CLI module that handles command-line arguments and orchestrates the quick entry workflow
- **Deliverables**: 
  - cmd/dca/cli.go with struct for CLI config
  - Flag parsing for --add, --asset, --amount, --price, --date
  - Error handling with exit code 1 on validation failures
  - Silent success (no output on valid entry)
- **Stakeholders**: Developers (implementation), DevOps (script-friendly interface)
- **Dependencies**: Phase 1 completed (validation functions available)
- **Constraints**: Must use Go's flag package, maintain consistency with existing codebase

## Task Generation Rules
- Each flag should have its own parsing task
- Include error message consistency with existing TUI validation
- Consider date defaulting to time.Now() behavior
- Ensure the module can be called independently from main()
- Include comprehensive error messages for each validation failure

## Acceptance Criteria
1. CLI module compiles without errors
2. All required flags (--add, --asset, --amount, --price) are properly parsed
3. Optional --date flag defaults to current time when not provided
4. Missing required flags produce clear error messages
5. Validation uses shared functions from Phase 1
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
