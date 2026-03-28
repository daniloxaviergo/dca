---
id: GOT-062
title: '[doc-013] Phase 7: Update README with CLI Usage Section'
status: To Do
assignee: []
created_date: '2026-03-28 15:19'
labels:
  - documentation
  - cli
  - readme
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update the README.md file with comprehensive documentation for the CLI quick entry feature.

## Phase Context
- **Objective**: Provide clear documentation for users on CLI usage
- **Deliverables**: 
  - New CLI Quick Entry section in README
  - Command examples (basic usage, with flags)
  - Usage patterns (script integration)
  - Error handling documentation
- **Stakeholders**: End Users (documentation), DevOps (operational guide)
- **Dependencies**: None - can be done in parallel with implementation
- **Constraints**: Follow existing README format and style

## Task Generation Rules
- Include usage examples for all flag combinations
- Document error messages and exit codes
- Provide script integration examples
- Show comparison between TUI and CLI approaches
- Include practical examples with realistic values

## Acceptance Criteria
1. CLI section clearly documented in README
2. All flags documented (--add, --asset, --amount, --price, --date)
3. Usage examples provided for common scenarios
4. Error handling documented
5. Examples include both valid and invalid usage patterns
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
