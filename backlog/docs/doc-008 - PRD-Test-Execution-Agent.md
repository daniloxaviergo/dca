---
id: doc-008
title: 'PRD: Test Execution Agent'
type: other
created_date: '2026-03-18 11:10'
---
# PRD: Test Execution Agent

## Overview

### Purpose
Create a specialized agent that executes Go tests in the DCA project to help developers fix issues faster and reduce token consumption during test-related workflows.

### Goals
- Reduce test execution time from ~200ms to sub-50ms for quick feedback
- Lower token consumption by using a focused agent instead of general AI exploration
- Provide intelligent test result summaries with actionable insights
- Enable rapid iteration on test-driven development workflows

## Background

### Problem Statement
The current test execution workflow in the DCA project is difficult and token-expensive:
1. Developers must manually run `go test ./...` or `make test`
2. When tests fail, AI assistants explore the codebase to understand issues (high token cost)
3. No intelligent summary of test failures with root cause analysis
4. Test execution timing and caching behavior is not automatically communicated

### Current State
- Tests run in ~200ms (cached) or ~500ms (uncached)
- Test output is verbose (23+ test functions across 5 packages)
- No automated analysis of test failures or performance issues
- Token-heavy workflow: AI reads files, runs commands, interprets results

### Proposed Solution
Create a specialized "testing-expert" subagent that:
1. Executes tests with optimized output (test timing + failure summary)
2. Analyzes failures and suggests root causes
3. Caches test results and uses smart caching for subsequent runs
4. Provides structured output to reduce token consumption

## Requirements

### User Stories

- **As a developer**, I want to execute tests with a single command so that I can get immediate feedback on code changes
- **As a developer**, I want intelligent test failure summaries so that I can quickly identify and fix issues
- **As a developer**, I want to see test timing information so that I can identify slow tests
- **As a developer**, I want to see which tests were cached so that I understand performance characteristics

### Functional Requirements

#### Task 1: Test Execution Agent Configuration
Create a dedicated "testing-expert" agent configuration for Go test execution.

##### Acceptance Criteria
- [ ] Agent configuration stored at `.qwen/agents/testing-expert.md`
- [ ] Agent has access to: `read_file`, `write_file`, `run_shell_command`
- [ ] Agent system prompt focuses on Go testing, test failure analysis, and performance optimization
- [ ] Agent can execute `go test` commands with various flags

#### Task 2: Intelligent Test Output
The agent should provide intelligent, structured test output.

##### Acceptance Criteria
- [ ] Shows test execution time per package
- [ ] Highlights failing tests with error messages
- [ ] Indicates which tests used cached results
- [ ] Provides summary statistics (total/passed/failed/skipped)

#### Task 3: Failure Analysis
The agent should analyze test failures and provide actionable insights.

##### Acceptance Criteria
- [ ] Identifies the root cause of test failures
- [ ] Suggests potential fixes for common failure patterns
- [ ] Links failing tests to relevant source code locations
- [ ] Provides context for flaky tests

#### Task 4: Performance Optimization
The agent should help improve test execution speed.

##### Acceptance Criteria
- [ ] Identifies slow tests (>100ms) and suggests optimizations
- [ ] Provides caching recommendations
- [ ] Suggests parallel test execution when appropriate

### Non-Functional Requirements

- **Performance**: Test execution agent should complete in <100ms for cached runs
- **Token Efficiency**: Agent output should use 50% fewer tokens than general AI exploration
- **Reliability**: Agent should handle all Go test scenarios (unit, integration, benchmarks)
- **Maintainability**: Agent configuration should be easily updated as project evolves

## Scope

### In Scope
- Test execution agent configuration (`.qwen/agents/testing-expert.md`)
- Go test command execution with intelligent output formatting
- Test failure analysis and root cause identification
- Performance metrics collection and reporting
- Integration with existing `make test` workflow

### Out of Scope
- Test coverage analysis (deferred to future iteration)
- Continuous integration integration (deferred to future iteration)
- Test flakiness detection over multiple runs (deferred to future iteration)
- Test result visualization in UI (deferred to future iteration)

## Technical Considerations

### Existing System Impact
- No changes to existing test code required
- Agent integrates with existing `go test` and `make test` commands
- No changes to project structure or dependencies

### Dependencies
- Go 1.25.7 (already in use)
- `go test` command (standard library)
- Qwen Code subagent framework

### Constraints
- Agent must work with existing project structure
- No external test frameworks (Go standard library only)
- Agent must be cache-aware (use Go's test caching)

## Success Metrics

### Quantitative
- Test execution time: <50ms for cached runs (current: ~200ms)
- Token consumption: 50% reduction vs general AI exploration
- Test failure analysis accuracy: >90% correct root cause identification

### Qualitative
- Developers can quickly identify and fix test failures
- Clear understanding of test performance characteristics
- Reduced cognitive load during test-driven development

## Timeline & Milestones

- [ ] **Design Complete**: Agent configuration and system prompt finalized
- [ ] **Implementation Complete**: Agent saved to `.qwen/agents/testing-expert.md`
- [ ] **Testing Complete**: Agent tested with project's existing test suite
- [ ] **Review**: Stakeholder review of agent output quality

## Stakeholders

### Decision Makers
- **User**: Project owner defining test agent requirements

### Contributors
- **Qwen Code**: AI assistant implementing the agent

## Appendix

### Glossary
- **Test Caching**: Go's feature to skip unchanged tests, improving execution time
- **Test Timing**: Measurement of how long each test takes to execute
- **Root Cause Analysis**: Identification of the underlying issue causing test failures

### References
- [Go Testing Documentation](https://pkg.go.dev/testing): Official Go testing package documentation
- [Qwen Subagents Guide](.qwen/commands/exec.md): Subagent configuration and usage guide
- [Makefile](Makefile): Existing test execution commands (`make test`, `make test-quiet`, `make test-cover`)

## Implementation Notes

### Agent Configuration Structure
```yaml
---
name: testing-expert
description: Executes Go tests with intelligent output, failure analysis, and performance insights
tools:
  - read_file
  - write_file
  - run_shell_command
---
```

### System Prompt Focus Areas
1. Execute `go test` commands with appropriate flags (`-v`, `-bench`, `-count=1`)
2. Parse and summarize test output
3. Identify failure patterns and suggest fixes
4. Provide timing analysis for performance optimization
5. Leverage test caching for faster subsequent runs

### Output Format
```
=== Test Execution Summary ===
Packages: 5
Tests: 45 (cached: 40, new: 5)
Duration: 45ms
Status: PASS

=== Slow Tests (>100ms) ===
TestDCAEntryValidate_NegativeAmount: 2.3ms (not slow)
... (no slow tests in this project)

=== Failure Analysis ===
No failures detected.

=== Performance Tips ===
- All tests cached (run with -count=1 to bypass cache)
- Consider parallel execution for larger test suites
```
