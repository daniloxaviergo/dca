---
id: GOT-034
title: Create test execution agent configuration
status: To Do
assignee: []
created_date: '2026-03-18 11:10'
labels:
  - agent
  - testing
  - documentation
dependencies: []
references:
  - backlog/docs/doc-008.md
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the testing-expert agent configuration file
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Agent configuration stored at .qwen/agents/testing-expert.md
- [ ] #2 Agent has access to: read_file, write_file, run_shell_command
- [ ] #3 Agent system prompt focuses on Go testing, test failure analysis, and performance optimization
- [ ] #4 Agent can execute go test commands with various flags
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
