---
id: GOT-061
title: '[doc-013 Phase 6] Update README with CLI usage documentation'
status: Done
assignee: []
created_date: '2026-03-28 20:50'
updated_date: '2026-03-29 01:21'
labels:
  - documentation
  - cli
dependencies: []
references:
  - 'doc-013 - Phase 6: Update README'
documentation:
  - doc-013
priority: low
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Update README.md with a new section documenting the CLI quick entry feature. Include command syntax (./dca --add --asset <ticker> --amount <usd> --price <per-share>), required and optional flags, examples, and behavior notes (silent success, exit codes, time.Now() default). Ensure documentation aligns with PRD specifications and maintains project style.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 README.md updated with CLI quick entry section
- [x] #2 Command syntax documented
- [x] #3 All flags explained (required/optional)
- [x] #4 Example usage provided
- [x] #5 Behavior notes included (silent success, exit codes)
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task adds CLI usage documentation to the README.md. The CLI quick entry feature allows users to add investment entries from the command line without launching the interactive TUI using the `--add` flag. The documentation will follow the existing README structure and style.

**Key topics to document:**
- CLI entry point syntax: `./dca --add --asset <ticker> --amount <usd> --price <per-share>`
- Required flags: `--add`, `--asset`, `--amount`, `--price`
- Optional flag: `--date` (defaults to current time if omitted)
- Validation rules for each field
- Behavior notes: silent success, exit codes, default date handling

**Documentation location:** Add a new "Command-Line Quick Entry" section in the README between "Getting Started" and "Development Commands".

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `README.md` | Modify | Add CLI quick entry section with usage examples, flags, and behavior notes |

### 3. Dependencies

- **Existing CLI implementation:** The CLI feature is already implemented in `cmd/dca/cli.go`
- **No breaking changes:** The feature is stable and fully tested
- **No new dependencies:** Uses only standard library and existing packages

### 4. Code Patterns

**Follow existing README patterns:**

1. **Section format:** Use `## Heading` for major sections, `### Subheading` for subsections
2. **Code blocks:** Use ```bash for shell commands, ```go for Go code
3. **Tables:** Use markdown tables for flag documentation (like the existing form field table)
4. **Navigation examples:** Include keyboard shortcuts in code format (e.g., `Esc`)
5. **Be explicit about command syntax:** Show the full command with all flags

**Style to follow:**
- Use bold for field names: `**--add**`
- Use angle brackets for placeholders: `<ticker>`, `<usd>`
- Show examples in separate code blocks
- Be concise but complete

### 5. Testing Strategy

**Documentation testing approach:**

1. **Verify CLI works as documented:** Run example commands manually to ensure accuracy
2. **Test error cases:** Verify error messages match documentation
3. **Test flag variations:** Test with and without optional `--date` flag
4. **Verify exit codes:** Confirm 0 on success, 1 on error

**Manual testing to perform:**
```bash
# Build the project
make build

# Test successful entry
./dca --add --asset BTC --amount 100 --price 50000

# Verify exit code
echo $?

# Test with date flag
./dca --add --asset ETH --amount 200 --price 3000 --date "2025-01-01T00:00:00Z"

# Test validation error (missing required flag)
./dca --add --amount 100 --price 50000
echo $?
```

### 6. Risks and Considerations

**No significant risks:**
- The CLI feature is complete and stable (GOT-057, GOT-059, GOT-060 completed)
- No breaking changes to CLI interface
- Documentation aligns with existing code behavior

**Important considerations:**

1. **Flag naming consistency:** Use `--add` flag (not `-add`) as shown in PRD
2. **Date default:** Document that `--date` defaults to `time.Now()` if omitted
3. **Silent success:** Emphasize no output on success (script-friendly)
4. **Validation messages:** Error messages come from validation package
5. **Path to binary:** Use `./dca` in examples (not `dca` or full path)

**Acceptance criteria to verify:**
- [x] README.md updated with CLI quick entry section
- [x] Command syntax documented (show full command with flags)
- [x] All flags explained (required: `--add`, `--asset`, `--amount`, `--price`; optional: `--date`)
- [x] Example usage provided (at least 2 examples)
- [x] Behavior notes included (silent success on exit 0, exit 1 on error, auto-date to current time)
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation notes: README.md updated with CLI quick entry section. All acceptance criteria verified and checked. Tests pass (183/183). Build succeeds with no warnings. Go fmt applied. No code changes needed as this is pure documentation.

Task completed. README.md updated with comprehensive CLI quick entry documentation. All tests pass (183/183). Build succeeds with no warnings. No code changes needed - this is pure documentation task.

Final verification complete. All acceptance criteria met, tests pass, build succeeds. Ready to finalize task.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Updated README.md with comprehensive CLI quick entry documentation section.

## Changes

Added new "Command-Line Quick Entry" section between "Development Commands" and "Usage" in README.md documenting:
- Command syntax: `./dca --add --asset <ticker> --amount <usd> --price <per-share> [--date <rfc3339>]`
- All flags (required: `--add`, `--asset`, `--amount`, `--price`; optional: `--date`)
- Behavior notes (silent success on exit 0, exit 1 on error, auto-date to current time)
- Example usage with and without date flag
- Error handling examples

## Verification

- **Tests**: 183 tests pass
- **Build**: Compiles without warnings
- **Manual testing**: CLI functionality verified (successful entry, error handling)
- **Formatting**: Go files formatted with `go fmt`
- **No code changes**: Pure documentation update
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria met
- [x] #2 Unit tests pass (go test)
- [x] #3 No new compiler warnings
- [x] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
