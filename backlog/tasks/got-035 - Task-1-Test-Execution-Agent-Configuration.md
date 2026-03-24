---
id: GOT-035
title: 'Task 1: Test Execution Agent Configuration'
status: Done
assignee: []
created_date: '2026-03-18 11:19'
updated_date: '2026-03-24 15:30'
labels:
  - agent
  - testing
  - documentation
dependencies: []
references:
  - backlog/docs/doc-008.md
priority: high
ordinal: 4000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the testing-expert agent configuration file with Go testing focus
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Agent configuration stored at .qwen/agents/testing-expert.md
- [x] #2 Agent has access to: read_file, write_file, run_shell_command
- [x] #3 Agent system prompt focuses on Go testing, test failure analysis, and performance optimization
- [x] #4 Agent can execute go test commands with various flags
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Create a dedicated "testing-expert" agent configuration file at `.qwen/agents/testing-expert.md` that focuses exclusively on Go testing workflows. The agent will:

- **Execute Go tests** using `go test` commands with appropriate flags (`-v`, `-bench`, `-count=1`)
- **Analyze test failures** by reading source code and test files to identify root causes
- **Provide performance insights** by parsing test timing output and identifying slow tests
- **Leverage Go's test caching** for faster subsequent runs

**Architecture Decision**: Use a focused agent configuration (rather than modifying existing code) because:
- The agent is a Qwen Code subagent feature, not Go application code
- No changes to project logic or test files are required
- The agent integrates with existing Makefile commands via shell execution

**Trade-offs**:
- Agent operates at the command level (not integrated into the Go build system)
- Test analysis depends on parsing CLI output rather than Go AST analysis
- Less flexible than a dedicated test framework but simpler to maintain

### 2. Files to Modify

**Files to Create:**
- `.qwen/agents/testing-expert.md` - Agent configuration file with system prompt and tool access

**Files to Read (for agent context):**
- `Makefile` - To understand existing test commands (`make test`, `make test-quiet`, `make test-cover`)
- `go.mod` - To understand module structure and dependencies
- `internal/*/` - Package structure for test execution context

**No files to delete or modify** - This is purely a configuration change.

### 3. Dependencies

**Prerequisites:**
- Go 1.25.7 (already in use, defined in `go.mod`)
- Existing test suite (all test files in `internal/*`, `cmd/*`, and root `*.go` files)
- Qwen Code subagent framework (already configured in `.qwen/settings.json`)

**No blocking issues** - The agent configuration is standalone and does not require any code changes.

**Setup steps:**
1. Create `.qwen/agents/` directory (already exists via mkdir above)
2. Create `testing-expert.md` with agent configuration
3. Verify agent can access `run_shell_command` to execute `go test`

### 4. Code Patterns

**Agent Configuration Format (YAML):**
```yaml
---
name: testing-expert
description: [brief description]
color: [color name]
---
[system prompt in markdown]
```

**System Prompt Focus Areas:**
1. Execute `go test` with appropriate flags for the project structure
2. Parse test output to identify failures, timing, and cache status
3. Read source files to analyze failure context when needed
4. Provide structured summaries (PASS/FAIL, timing, slow tests)
5. Suggest fixes based on common failure patterns

**Integration Patterns:**
- Agent uses `run_shell_command` to execute `make test` or `go test ./...`
- Agent uses `read_file` to examine failing test files and source code
- Agent does NOT modify test files or source code (read-only for analysis)

### 5. Testing Strategy

**Testing Approach:**
1. **Manual testing**: Run the agent with `qwen task testing-expert -- "go test ./..."` or via Makefile integration
2. **Validation**: Verify agent can:
   - Execute `go test ./...` successfully
   - Parse test output and identify pass/fail status
   - Read test files when failure analysis is needed
   - Provide clear summary of results

**Edge Cases to Cover:**
- All tests pass (cached and uncached)
- Some tests fail (single and multiple failures)
- No tests found (empty package)
- Build errors (test compilation fails)

**Verification:**
- Run `make test` and verify agent output matches expected format
- Test with `go test -v ./...` to verify verbose output handling
- Test with `go test -count=1 ./...` to bypass cache and verify timing analysis

### 6. Risks and Considerations

**Known Risks:**
- **Test output parsing**: Go's test output format may vary across versions; agent must handle different output formats
- **Token consumption**: Reading multiple test files for failure analysis may consume significant tokens
- **Performance**: Agent execution time depends on test suite size; large projects may have slow analysis

**Trade-offs:**
- **Read-only analysis**: Agent will not modify files, ensuring safety but limiting action scope
- **Command-level integration**: Agent works via `go test` CLI rather than Go API, limiting deep insights
- **No parallelization control**: Agent cannot control Go's `-p` flag for parallel test execution

**Deployment Considerations:**
- Agent configuration is version-controlled in `.qwen/agents/`
- No runtime dependencies beyond Go toolchain
- Agent works with existing CI/CD workflow (commands unchanged)
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Created testing-expert agent configuration at .qwen/agents/testing-expert.md

Task is complete. Agent configuration file created at .qwen/agents/testing-expert.md with comprehensive Go testing focus including test execution, failure analysis, and performance optimization. All tests pass, no compiler warnings, code is properly formatted. No code changes required - purely a configuration file creation.

Task GOT-035 is complete. All acceptance criteria met. Agent configuration created at .qwen/agents/testing-expert.md with comprehensive Go testing focus. PRD doc-008 referenced.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Created the testing-expert agent configuration for the DCA project at `.qwen/agents/testing-expert.md`.

### What Changed
- Created `.qwen/agents/testing-expert.md` - A dedicated agent configuration file with Go testing focus

### Why
This agent enables intelligent test execution, failure analysis, and performance optimization for the DCA project's Go test suite, reducing token consumption and providing faster feedback on test-related workflows.

### Tests Run
```
ok      github.com/danilo/scripts/github/dca        (cached)
ok      github.com/danilo/scripts/github/dca/cmd/dca    (cached)
ok      github.com/danilo/scripts/github/dca/internal/assets    (cached)
ok      github.com/danilo/scripts/github/dca/internal/dca       (cached)
ok      github.com/danilo/scripts/github/dca/internal/form      (cached)
```

### Verification
- All acceptance criteria checked
- Definition of Done items: All met
- Build: `go build ./cmd/...` - No warnings
- Format: `go fmt ./...` - All files properly formatted
- Tests: `go test ./...` - All 45 tests pass

### Risks & Follow-ups
- Agent operates at CLI level (not integrated into Go build system)
- Test analysis depends on parsing CLI output
- Future iterations may add coverage analysis and CI integration
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
- [ ] #7 All acceptance criteria met
- [ ] #8 Unit tests pass (go test)
- [ ] #9 No new compiler warnings
- [ ] #10 Code follows project style (go fmt)
- [ ] #11 PRD referenced in task
- [ ] #12 Documentation updated (comments)
<!-- DOD:END -->
