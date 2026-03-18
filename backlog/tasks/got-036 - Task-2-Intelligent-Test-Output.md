---
id: GOT-036
title: 'Task 2: Intelligent Test Output'
status: Done
assignee:
  - Thomas
created_date: '2026-03-18 11:19'
updated_date: '2026-03-18 11:47'
labels:
  - agent
  - testing
  - output
dependencies: []
references:
  - backlog/docs/doc-008.md
priority: high
ordinal: 4000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement intelligent test output formatting with timing, cached results, and failure highlighting
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Shows test execution time per package
- [x] #2 Highlights failing tests with error messages
- [x] #3 Indicates which tests used cached results
- [x] #4 Provides summary statistics (total/passed/failed/skipped)
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

The goal is to enhance the `testing-expert.md` agent configuration to provide intelligent test output formatting with timing, cached results, and failure highlighting. This is a configuration/update task, not a code change task.

**Approach:**
- Update `.qwen/agents/testing-expert.md` with enhanced output formatting guidance
- Add logic for parsing Go test output and presenting it in a structured format
- Include templates for summary reports with package timing, cache status, and failure highlights

**Architecture Decisions:**
- Keep the agent as a configuration file (no code changes to Go source)
- Use Go's built-in test flags (`-v`, `-count=1`, `-bench`) for timing data
- Parse `go test` output to extract timing and cache information
- Create structured output format for reduced token consumption

### 2. Files to Modify

- `.qwen/agents/testing-expert.md` - Update agent configuration with:
  - Enhanced test execution commands with timing flags
  - Output parsing logic for timing and cache detection
  - Structured summary format template
  - Failure highlighting and analysis guidance

### 3. Dependencies

- **Existing:** `.qwen/agents/testing-expert.md` (already exists, needs update)
- **Existing:** `go test` command (standard Go tooling)
- **Existing:** Makefile targets (`make test`, `make test-quiet`)
- **No blocking issues** - this is a configuration update

### 4. Code Patterns

Follow existing patterns in the agent file:
- Use `---` YAML frontmatter for metadata
- Use `##` headers for sections
- Include code block examples with backticks
- List commands with `-` bullet points
- Use structured output format with `=== Section ===` markers

### 5. Testing Strategy

**Testing Approach:**
- Run `make test` to verify agent can execute tests
- Run `make test -count=1` to test timing output
- Manually verify the agent can parse and summarize output
- Verify cache detection (cached vs uncached runs)

**What to Verify:**
- Agent can run all Makefile targets
- Agent extracts timing per package
- Agent detects cache status from output
- Agent identifies failures and highlights them
- Agent provides summary statistics

### 6. Risks and Considerations

**No significant risks** - this is a configuration-only change.

**Considerations:**
- Go test output format is stable across versions
- Cache status is indicated by `(cached)` in output
- Timing is shown as `package N.NNs` at the end of each package section
- Test failures show `--- FAIL` with error context on following lines
- Summary lines show `PASS/FAIL package N.NNs` for timing

**Output Format to Implement:**
```markdown
=== Test Execution Summary ===
Packages: 5
Tests: 45 (cached: 40, new: 5)
Duration: 45ms
Status: PASS

=== Package Timing ===
- github.com/danilo/scripts/github/dca: 12ms
- github.com/danilo/scripts/github/dca/cmd/dca: 8ms
- ... (per-package timing)

=== Cache Status ===
Cached: 40 tests (89%)
Uncached: 5 tests (11%)

=== Failure Analysis ===
If failures: List failing tests with errors
If pass: No failures detected.

=== Performance Tips ===
- Suggest parallelization if tests > 100ms
- Note when cache is effective
- Recommend -count=1 for fresh runs
```
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Task review complete: Configuration-only update to testing-expert.md agent

Examined current testing-expert.md agent configuration

Analyzed go test output format with timing and cache indicators

Key findings: Cache status shown as (cached) or timing, verbose output shows per-test timing, summary shows PASS/FAIL package duration, failed tests show error context

Updated .qwen/agents/testing-expert.md with intelligent test output formatting

Added output parsing rules for timing, cache status, and failure detection

Added structured output format template with Test Execution Summary, Package Timing, Cache Status, and Failure Analysis sections

Added parsing guidelines for extracting package timing, cache status, failure detection, count statistics, and timing calculation

All tests pass with make check (fmt, build, test)

No compiler warnings introduced

Updated .qwen/agents/testing-expert.md agent configuration file

No code changes to Go source files required - this was a configuration-only task

PRD referenced: backlog/docs/doc-008.md

All acceptance criteria verified and marked complete

All Definition of Done items satisfied

Final verification: make check passed with 0 warnings and 82 tests passing across 4 packages
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Task GOT-036: Intelligent Test Output Formatting

### Changes Made
Updated `.qwen/agents/testing-expert.md` with intelligent test output formatting capabilities:

1. **Enhanced Test Execution Commands**: Added flags for timing (`-count=1`), race detection (`-race`), and coverage (`-coverprofile`)

2. **Output Parsing Rules**: Documented how to parse Go test output for:
   - Package timing (`package N.NNs`)
   - Cache status (`(cached)` suffix)
   - Test names (`=== RUN TestName`)
   - Failures (`--- FAIL TestName` with error context)

3. **Structured Output Format**: Added template with:
   - Test Execution Summary (packages, tests, duration, status)
   - Package Timing (per-package duration breakdown)
   - Cache Status (cached vs uncached percentages)
   - Failure Analysis (list of failing tests with suggested fixes)
   - Performance Tips (optimization recommendations)

4. **Parsing Guidelines**: Added step-by-step parsing instructions for timing extraction, cache detection, failure detection, and statistics calculation

### Verification
- `make test` - 82 tests pass across 4 packages
- `make check` - fmt, build, test all pass with 0 warnings
- No new compiler warnings introduced
- Code follows project style (go fmt)

### Files Modified
- `.qwen/agents/testing-expert.md` - Updated agent configuration

### Risks
None - configuration-only change to agent documentation
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [x] #5 PRD referenced in task
- [x] #6 Documentation updated (comments)
- [x] #7 All acceptance criteria met
- [x] #8 Unit tests pass (go test)
- [x] #9 No new compiler warnings
- [x] #10 Code follows project style (go fmt)
- [x] #11 PRD referenced in task
- [x] #12 Documentation updated (comments)
<!-- DOD:END -->
