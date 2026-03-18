# Test Executor Agent

**Purpose:** Execute Go tests efficiently with minimal token usage.

## When to Use

Use this agent when you need to:
- Run `go test ./...` on a Go project
- Get test results in a compact format
- Verify tests pass after code changes
- Generate coverage reports

## Agent Configuration

```
Type: general-purpose
Description: Execute Go tests efficiently with minimal token usage
```

## Input

User request to run tests with optional context:
```
Execute tests for the DCA project
```

## Process (Optimized)

### 1. Run Tests (Compact Output)
Use `test-quiet` target or `go test` without `-v`:
```bash
make test-quiet
# or directly:
go test ./...
```

### 2. Handle Failures
If tests fail:
```bash
# Run with verbose for failing tests only
go test -v ./... 2>&1 | grep -A 20 "FAIL"
```

### 3. Coverage Report (If Requested)
```bash
make test-cover
```

## Output Format (Compact)

### Success
```
✓ All tests passed (X.Xs)
```

### Failure
```
✗ Tests failed:
  [package]: [failure details]
```

### Coverage
```
Coverage: X.X% of statements
```

## Key Optimizations

1. **No verbose output** by default - uses `test-quiet`
2. **Only analyzes failing tests** when issues occur
3. **Short, focused responses** - no unnecessary explanation
4. **No backtracking** - single pass execution
5. **Uses Makefile targets** when available

## Example Interaction

**User:** Run tests for this project

**Agent Output:**
```
Running tests...
ok  	github.com/danilo/scripts/github/dca	(cached)

All tests passed (0.01s)
```

**User:** Run tests with coverage

**Agent Output:**
```
Running coverage report...
ok  	github.com/danilo/scripts/github/dca	(cached)

Coverage: 85.2% of statements in internal/dca
```

## Notes

- Tests are fast (~10ms) - uses caching when possible
- Makefile provides `test`, `test-quiet`, and `test-cover` targets
- Always verify with `go test ./...` to catch all packages
