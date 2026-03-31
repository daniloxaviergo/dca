---
id: GOT-030
title: 'Task 1: Intelligent Test Selection'
status: To Do
assignee:
  - workflow
created_date: '2026-03-18 00:27'
updated_date: '2026-03-31 09:40'
labels: []
dependencies: []
references:
  - backlog/docs/doc-007.md
priority: high
ordinal: 4000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement intelligent test selection based on file dependencies, test history caching, and modified files analysis. The agent should select only relevant tests to run, cache results, and invalidate cache when source changes.

Acceptance Criteria:
- Agent can identify test files that depend on modified source files
- Agent caches test results (pass/fail/skip) with metadata (timestamp, git commit)
- Agent respects test flags (e.g., -run, -short, -race)
- Cache is invalidated when source files change
- Agent can fallback to full test run if no cache found

Technical Notes:
- Create internal/testagent/selector.go for test selection logic
- Implement dependency analysis using go list or parsing imports
- Use JSON cache format for portability
- Integrate with git for modified files detection

References:
- PRD doc-007, Task 1
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
- [ ] #7 Implementation includes test selection algorithm based on file dependencies
- [ ] #8 Implementation includes test result caching with invalidation logic
- [ ] #9 Implementation includes LLM-based failure analysis with fix suggestions
- [ ] #10 Implementation includes integration with existing Makefile commands
<!-- DOD:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Agent can identify test files that depend on modified source files
- [ ] #2 Agent caches test results (pass/fail/skip) with metadata (timestamp, git commit)
- [ ] #3 Agent respects test flags (e.g., -run, -short, -race)
- [ ] #4 Cache is invalidated when source files change
- [ ] #5 Agent can fallback to full test run if no cache found
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This implementation creates an **Intelligent Test Selection Agent** (`internal/testagent/`) that optimizes test execution by selecting only relevant tests based on file dependencies and change analysis.

**Core Components:**

1. **Dependency Analysis Engine**
   - Uses `go list -json` to parse Go packages and their dependencies
   - Maps source files to their corresponding test files (e.g., `dca_entry.go` → `dca_entry_test.go`)
   - Tracks import relationships between packages
   - Handles table-driven tests by analyzing test file contents

2. **Cache Management System**
   - JSON-based cache format stored in `.dca/test-cache.json`
   - Stores: test name, status (PASS/FAIL/SKIP), timestamp, git commit hash
   - Cache key includes: file modification time, git commit, Go version
   - Automatic cache invalidation when source files change

3. **Modified Files Detection**
   - Uses `git diff --name-only HEAD` to identify modified files
   - Compares against cache to determine which tests need re-running
   - Supports incremental test runs for staged but uncommitted changes

4. **Integration with Go Test Flags**
   - Parses and respects `-run`, `-short`, `-race`, `-v` flags
   - Passes through unsupported flags to `go test`
   - Fallback to full test run when cache is unavailable

**Algorithm:**
1. Detect modified files via git
2. Load cached test results from `.dca/test-cache.json`
3. For each modified file, identify dependent test files
4. Run tests for modified files + any tests that previously failed
5. Update cache with new results
6. Fallback to full run if no cache found or `--full` flag set

### 2. Files to Modify

| Action | File | Purpose |
|--------|------|---------|
| **Create** | `internal/testagent/selector.go` | Core test selection logic with dependency analysis |
| **Create** | `internal/testagent/cache.go` | Cache management (load/save/invalidation) |
| **Create** | `internal/testagent/modified.go` | Git-based modified file detection |
| **Create** | `internal/testagent/flags.go` | Test flag parsing and passthrough |
| **Create** | `internal/testagent/agent.go` | Main agent orchestration |
| **Create** | `internal/testagent/testagent_test.go` | Unit tests for all components |
| **Modify** | `Makefile` | Replace `test`/`test-quiet` targets with agent wrapper |
| **Create** | `.dca/` directory | Cache storage location (auto-created on first run) |

### 3. Dependencies

**Prerequisites:**
- Git must be available for modified file detection
- Go 1.25.7 already in `go.mod`
- No external dependencies required (uses standard library + `go list`)

**Setup Steps:**
1. Create `internal/testagent/` directory structure
2. Initialize cache directory `.dca/` on first run
3. Update `Makefile` to use agent wrapper instead of direct `go test`
4. Ensure all existing tests continue to pass (no behavior changes)

**Blocking Issues:**
- None - all dependencies available

### 4. Code Patterns

**Following Existing Conventions:**
- Use `internal/` package structure (matches `internal/dca/`, `internal/form/`, `internal/assets/`)
- Follow naming: `Test{Function}_{Condition}` for test functions
- Use table-driven tests for validation logic
- Exact error messages for user-facing output
- Atomic file operations for cache (temp file + rename)
- 8-decimal precision for financial calculations (not applicable here, but follows Go style)

**Integration Patterns:**
- Agent wraps `go test` via `exec.Command` for actual test execution
- JSON cache format for portability
- Minimal changes to existing codebase (read-only access to source)
- Makefile targets delegate to agent with flags

**Naming Conventions:**
- Package: `testagent`
- Types: `TestSelector`, `CacheManager`, `ModifiedFilesDetector`
- Functions: `IdentifyDependentTests`, `LoadCache`, `SaveCache`, `DetectModifiedFiles`
- Methods: `Selector.Select()`, `Cache.Get(key)`, `Cache.Set(key, value)`

### 5. Testing Strategy

**Unit Tests:**
- `TestSelector_IdentifyDependentTests_Pass` - verify test file mapping
- `TestSelector_IdentifyDependentTests_NoTests` - handle packages without tests
- `TestCache_Load_EmptyFile` - handle missing cache gracefully
- `TestCache_Save_AtomicWrite` - verify atomic write pattern
- `TestCache_Invalidate_OnFileChange` - verify invalidation logic
- `TestModifiedFiles_DetectGitChanges` - verify git diff parsing
- `TestFlags_ParseFlags` - verify flag parsing and passthrough
- `TestAgent_Run_WithCache` - integration test with cached results
- `TestAgent_Run_FallbackFullRun` - verify fallback behavior

**Edge Cases:**
- Packages with no test files
- Files modified but no tests affected
- Cache file corrupted or unreadable
- Git not available or no git repository
- Large test suites (performance not blocker for v1)

**Verification:**
- All existing tests pass after implementation
- Cache can be invalidated with `--no-cache` flag
- Full test run produces identical results to agent with cache disabled
- Token reduction measured against verbose output

### 6. Risks and Considerations

**Design Trade-offs:**
- **Decision**: Use `go list -json` instead of AST parsing
  - **Reason**: More reliable, handles imports correctly, Go官方 tooling
  - **Risk**: Slower for large codebases (mitigated by caching)

- **Decision**: JSON cache format instead of binary
  - **Reason**: Human-readable, easy to debug, portable
  - **Risk**: Slightly larger files (negligible for test cache)

- **Decision**: Fallback to full run when cache unavailable
  - **Reason**: Ensures correctness, no false negatives
  - **Risk**: No performance benefit on first run (acceptable)

**Known Limitations (v1):**
- Only detects file-level dependencies (not function-level)
- No support for build tags in dependency analysis
- No parallel test selection optimization
- Cache does not persist across machines (by design)

**Future Enhancements:**
- Function-level dependency analysis
- Build tag awareness
- CI/CD cache synchronization
- Historical test execution analysis for flaky test detection

**Implementation Risks:**
- **Low**: Go version compatibility (1.25.7 is current, stable API)
- **Low**: Cache format versioning (include version field in JSON)
- **Medium**: Git command availability (fallback if git not found)
- **Low**: Performance on very large test suites (cache mitigates)
<!-- SECTION:PLAN:END -->
