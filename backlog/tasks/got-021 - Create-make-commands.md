---
id: GOT-021
title: Create make commands
status: To Do
assignee:
  - Catarina
created_date: '2026-03-17 17:09'
updated_date: '2026-03-17 17:11'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
create make run, make test, make build, etc...
to simplify the process of development
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

Create a `Makefile` to simplify common development tasks. The Makefile will define standard targets for building, testing, and running the application, following Go best practices.

- **make run**: Execute the application directly (equivalent to `go run ./cmd/dca`)
- **make build**: Compile the binary to `./dca` (equivalent to `go build -o dca ./cmd/dca`)
- **make test**: Run all tests with verbose output (`go test ./... -v`)
- **make test-quiet**: Run tests without verbose output (`go test ./...`)
- **make test-cover**: Generate coverage report (`go test -cover ./...`)
- **make fmt**: Format all Go files (`go fmt ./...`)
- **make check**: Run fmt, build, and test in sequence (CI-friendly validation)
- **make clean**: Remove compiled binary and temporary files

**Architecture decisions:**
- Use GNU Make syntax with `--no-builtin-rules` to avoid unexpected behavior
- Define `go` as a variable for easy customization
- Use `.PHONY` for all targets to ensure they run regardless of file existence
- Keep targets simple and composable (build on existing `go` commands)
- Follow the principle of least surprise - make targets should map directly to common Go commands

### 2. Files to Modify

| File | Action | Purpose |
|------|--------|---------|
| `Makefile` | Create | New file with development task definitions |

### 3. Dependencies

**Prerequisites:**
- GNU Make (standard on Linux/macOS, available on Windows via MSYS2/Cygwin)
- Go 1.25.7+ (already in `go.mod`)

**No additional dependencies required.**

### 4. Code Patterns

**Makefile conventions to follow:**
- Use `$(MAKEFLAGS) --no-builtin-rules` at the top
- Define variables for reusability (`GO ?= go`, `BUILD_DIR ?= .`)
- Use `@` prefix to suppress command echo
- Use `$(shell ...)` for dynamic values (e.g., Go version validation)
- Error handling with `|| exit 1` or separate commands
- Keep commands short and focused on single responsibilities

**Target naming:**
- Use simple, intuitive names (`run`, `build`, `test`, `fmt`, `clean`)
- Use hyphenated alternatives for multi-word names if needed (`test-cover`)
- Follow Go ecosystem conventions (e.g., `test`, not `test-all`)

### 5. Testing Strategy

**Unit testing for Makefile:**
- Verify targets exist and produce expected output
- Test error handling (e.g., build fails with syntax errors)
- Test clean removes files
- Verify help target shows all available commands

**Verification steps:**
1. Run `make build` and verify binary is created
2. Run `make run` and verify application starts (Ctrl+C to exit)
3. Run `make test` and verify all tests pass
4. Run `make fmt` and verify no changes are made to formatted files
5. Run `make clean` and verify binary is removed
6. Run `make check` to verify full validation pipeline

### 6. Risks and Considerations

**No blocking issues or trade-offs.**

**Considerations:**
- Makefile will only work on systems with GNU Make installed
- Windows users may need MSYS2, Cygwin, or WSL
- The `run` target uses `go run` which may be slower than `go build` for large projects
- Coverage reports will be generated in temporary files (no cleanup target for coverage data)

**Future enhancements (not in scope):**
- `make lint` (if linter is added to project)
- `make docs` (if documentation generation is needed)
- `make release` (for versioned releases)
<!-- SECTION:PLAN:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
