---
id: GOT-031
title: 'Task 2: Optimized Output Logging'
status: To Do
assignee:
  - workflow
created_date: '2026-03-18 00:27'
updated_date: '2026-03-31 09:55'
labels: []
dependencies: []
references:
  - backlog/docs/doc-007.md
priority: high
ordinal: 3000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement optimized output logging to minimize token consumption. The agent should filter test output to only include failures and relevant details, generate JSON for programmatic analysis, and provide color-coded terminal output.

Acceptance Criteria:
- Agent can suppress verbose output when not requested
- Agent generates JSON output for downstream processing
- Agent provides color-coded terminal output when running interactively
- Agent reduces token usage by 60% compared to verbose mode

Technical Notes:
- Create internal/testagent/output.go for output filtering
- Implement tee-style output capture for filtering
- Use lipgloss for color-coded terminal output (existing dependency)
- JSON output should include test name, status, duration, and error messages

References:
- PRD doc-007, Task 2
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 Implementation includes output filtering to reduce token consumption
- [ ] #8 Implementation includes JSON output for programmatic analysis
- [ ] #9 Implementation includes color-coded terminal output using lipgloss
<!-- DOD:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Agent can suppress verbose output when not requested
- [ ] #2 Agent generates JSON output for downstream processing
- [ ] #3 Agent provides color-coded terminal output when running interactively
- [ ] #4 Agent reduces token usage by 60% compared to verbose mode
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Implementation Plan: Optimized Output Logging for Test Agent

### 1. Technical Approach

This task implements output filtering for the test agent to reduce token consumption by 60% while maintaining useful information for developers. The approach uses a **tee-style output capture** pattern combined with **smart filtering** and **lipgloss-based color formatting**.

**Key Components:**
1. **Output Filter** (`internal/testagent/output.go`): Captures `go test` output and filters it based on verbosity settings
2. **JSON Logger**: Generates structured JSON output for programmatic analysis
3. **Terminal Renderer**: Uses lipgloss for color-coded output when running interactively
4. **Context-Aware Filtering**: Only includes failures, summary stats, and relevant details

**Architecture:**
```
go test command
    |
    v
OutputCapture (tee-style)
    |
    +--> FilteredText (to terminal, lipgloss colored)
    +--> JSONOutput (to file/stdout)
    +--> AgentResult (internal struct)
```

**Token Reduction Strategy:**
- Verbose mode: `-v` flag outputs every test start/finish (high token cost)
- Optimized mode: Only output failures + summary (target: 40% of verbose tokens)
- Filter patterns: Remove noise like "PASS", "ok", "ok github.com/..." lines when no failures

---

### 2. Files to Modify

| Action | File | Purpose |
|--------|------|---------|
| **Create** | `internal/testagent/output.go` | New output filtering and formatting module |
| **Modify** | `internal/testagent/agent.go` | Integrate OutputFilter into Run/FormatOutput methods |
| **Modify** | `internal/testagent/flags.go` | Add `Quiet` flag to suppress non-essential output |

**No breaking changes** - all existing APIs remain compatible.

---

### 3. Dependencies

**Existing Dependencies (no new required):**
- `github.com/charmbracelet/bubbletea v1.3.10` - Already in go.mod
- `github.com/charmbracelet/lipgloss v1.1.0` - Already in go.mod (used for terminal styling)

**Prerequisites:**
- Task GOT-030 (Intelligent Test Selection) should be complete for cache integration
- Test output patterns are consistent with `go test` output format

**No blocking issues** - lipgloss is already a project dependency.

---

### 4. Code Patterns

**Follow Existing Patterns:**
1. **Error handling**: Return errors with descriptive messages (already in testagent package)
2. **Atomic writes**: Use temp file + rename pattern (already in cache.go)
3. **lipgloss styling**: Match existing patterns in `internal/form/` and `internal/assets/`

**lipgloss Color Scheme (matching existing code):**
```go
// Success: Green (82)
Foreground(lipgloss.Color("82"))
// Failure: Red (196)  
Foreground(lipgloss.Color("196"))
// Warning: Yellow (220)
Foreground(lipgloss.Color("220"))
// Info: Blue (63) - for active elements
Foreground(lipgloss.Color("63"))
// Dimmed: Gray (240)
Foreground(lipgloss.Color("240"))
```

**Output Format:**
```
=== RUN   TestExample
--- PASS: TestExample (0.01s)
=== RUN   TestFailure
--- FAIL: TestFailure (0.02s)
    example_test.go:42: expected 1, got 2
    stack trace...
FAIL
coverage: 85.0% of statements

Test Results:
  Duration: 1.2s
  Tests: 45 passed, 1 failed, 2 skipped
  Coverage: 85.0%
```

**JSON Output Format:**
```json
{
  "tests": [
    {
      "name": "TestExample",
      "status": "PASS|FAIL|SKIP",
      "duration_ms": 12,
      "error": ""
    }
  ],
  "summary": {
    "total": 45,
    "passed": 44,
    "failed": 1,
    "skipped": 2,
    "coverage": 85.0,
    "duration_ms": 1200
  }
}
```

---

### 5. Testing Strategy

**Unit Tests (internal/testagent/output_test.go):**

| Test | Purpose |
|------|---------|
| `TestOutputFilter_FilterVerbose_None` | Verify no filtering when verbose enabled |
| `TestOutputFilter_FilterQuiet_Success` | Filter PASS lines when quiet enabled |
| `TestOutputFilter_FilterQuiet_Failures保留` | Ensure FAIL lines are never filtered |
| `TestOutputFilter_CalculateTokenSavings` | Verify token reduction target (60%) |
| `TestOutputFilter_RenderJSON` | Verify JSON output format |
| `TestOutputFilter_RenderTerminal` | Verify lipgloss color codes present |
| `TestOutputFilter_EmptyOutput` | Handle empty test output gracefully |
| `TestOutputFilter_MultipleFailures` | Handle multiple test failures |

**Test Approach:**
- Use table-driven tests for filtering logic
- Compare token counts before/after filtering
- Verify lipgloss color codes in rendered output
- Test with actual `go test -v` output samples

**Integration:**
- Existing `testagent_test.go` tests continue to pass
- No changes to public API of `Agent` struct
- Add `OutputFilter` as optional integration point

---

### 6. Risks and Considerations

**Known Risks:**

| Risk | Mitigation |
|------|------------|
| **Token reduction below 60%** | Start with aggressive filtering (10% of verbose), adjust based on actual measurements |
| **Lipgloss on non-TTY terminals** | Detect TTY with `isatty` package, fallback to plain text |
| **go test output format changes** | Parse output line-by-line, handle unknown lines gracefully |
| **JSON output size** | Implement streaming JSON if output becomes too large |

**Trade-offs:**

1. **Verbosity vs. Information**: 
   - Verbose mode: Full test output (preserves `-v` behavior)
   - Quiet mode: Only failures + summary (target for default)
   - User can override with `-v` flag

2. **Performance vs. Accuracy**: 
   - Filter as we go (real-time) vs. post-process (more accurate but memory intensive)
   - Approach: Line-by-line streaming filter (memory efficient)

3. **Styling vs. Compatibility**:
   - Lipgloss provides great styling but requires TTY
   - Fallback: Detect non-TTY and use ANSI color codes or plain text

**Implementation Checklist:**
- [ ] Create `output.go` with `OutputFilter` struct
- [ ] Implement `FilterLines()` method for line-by-line filtering
- [ ] Implement `RenderTerminal()` with lipgloss styling
- [ ] Implement `RenderJSON()` for programmatic output
- [ ] Update `Agent.FormatOutput()` to use new filter
- [ ] Add `Quiet` flag to `TestFlags`
- [ ] Write unit tests for `output.go`
- [ ] Measure token reduction vs. verbose mode
- [ ] Test with real `go test` runs
- [ ] Update documentation (README, task description)
<!-- SECTION:PLAN:END -->
