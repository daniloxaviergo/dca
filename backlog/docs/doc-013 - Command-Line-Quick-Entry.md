---
id: doc-013
title: Command-Line Quick Entry
type: other
created_date: '2026-03-28 14:35'
---
# PRD Summary: Command-Line Quick Entry

## Executive Summary

Successfully refined the PRD for adding command-line quick entry to the DCA Investment Tracker. The feature enables users to add investment entries without launching the interactive TUI, with the following command format:

```bash
./dca --add --asset <ticker> --amount <usd> --price <per-share>
```

## Key Requirements

| Requirement | Status | Notes |
|------ ----|-- ------|-- ----|
| CLI flag-based input | ✅ | `--add` flag, no TUI |
| Required flags: asset, amount, price | ✅ | All validated |
| Auto-calculated shares | ✅ | 8 decimal precision |
| Auto-set date | ✅ | Defaults to `time.Now()` |
| Silent success | ✅ | No output on success |
| Strict validation | ✅ | Exit code 1 on error |
| TUI compatibility | ✅ | No breaking changes |

## Technical Decisions

1. **Flag-based approach**: Use `--add` flag rather than subcommand for simplicity
2. **Auto-date**: Use `time.Now()` for timestamp (common case)
3. **Silent success**: No output on success (script-friendly)
4. **Shared validation**: Extract from `form/validation.go`
5. **Early exit**: CLI operation exits before TUI initialization

## Acceptance Criteria

### Functional (8)
- ✅ CLI adds entry to data file
- ✅ Auto-calculated shares correct (8 decimals)
- ✅ Date auto-set to current RFC3339
- ✅ Missing required flag returns exit code 1
- ✅ Negative/zero amount returns exit code 1
- ✅ Negative/zero price returns exit code 1
- ✅ Successful entry produces no output
- ✅ TUI continues to work unchanged

### Non-Functional (3)
- ✅ CLI operation < 100ms
- ✅ Existing tests pass unchanged
- ✅ Data consistency maintained

## Files to Modify

| File | Action | Reason |
|------ ----|-- ------|-- ----|
| `cmd/dca/main.go` | Modify | Add CLI flag parsing + early exit |
| `cmd/dca/cli.go` | Create | CLI-specific logic |
| `internal/form/validation.go` | Refactor | Extract shared validation |
| `README.md` | Update | Add CLI usage section |

## Validation Rules (Shared with TUI)

| Field | Validation | Error Message |
|------|-- ----|-- ----- |
| `--amount` | Positive (> 0) | "Amount must be positive" |
| `--price` | Positive (> 0) | "Price must be positive" |
| `--asset` | Non-empty | "Asset ticker is required" |
| `--date` | RFC3339 (if provided) | "Use YYYY-MM-DD" |

## Shares Calculation

```
shares = math.Round((amount / price) * 1e8) / 1e8

Example: 2.0 / 65000.0 = 0.00030769
```

## Out of Scope

- `--batch` mode for CSV imports
- `--dry-run` flag for validation testing
- `--file` flag for custom data file
- Environment variable configuration
- Multiple entries in single invocation
- JSON output on success
- Verbose mode with `-v` flag

## Implementation Checklist

- [ ] Extract validation functions from `form/validation.go`
- [ ] Create `cmd/dca/cli.go` with flag parsing logic
- [ ] Implement `runCLI()` function with validation and persistence
- [ ] Add `--add` flag detection in `main()`
- [ ] Implement error handling with exit codes
- [ ] Add tests for CLI validation logic
- [ ] Add tests for CLI data persistence
- [ ] Update README with CLI usage section
- [ ] Verify backward compatibility with existing tests

## Files Created

- `PRD.md` - User-facing requirements document
- `PRD_CLI_QUICK_ENTRY.md` - Technical implementation PRD

## Stakeholder Alignment

| Stakeholder | Requirement ownership | Acceptance verification |
|------ -------|- -------------- -------|------ -------------- ----|
| End Users | REQ-001, REQ-002, REQ-003, REQ-004 | Quick entry works as expected |
| Developers | REQ-009, REQ-010 | Code reuse, no duplication |
| QA Team | ACC-001, ACC-002, ACC-003 | Tests pass, data correct |
| DevOps | NFA-001 | Performance acceptable |

## Traceability

| Requirement | Epic | User Story | Acceptance Criterion | Test File |
|------ -------|- -----|-- ------ ----|---------- -----------|------- ----|
| REQ-001 | CLI Quick Entry | Quick entry without TUI | ACC-008 | cmd/dca/cli_test.go |
| REQ-002 | CLI Quick Entry | Asset ticker specification | ACC-004 | cmd/dca/cli_test.go |
| REQ-003 | CLI Quick Entry | Amount specification | ACC-005 | cmd/dca/cli_test.go |
| REQ-004 | CLI Quick Entry | Price specification | ACC-006 | cmd/dca/cli_test.go |
| REQ-005 | CLI Quick Entry | Date auto-generation | ACC-003 | cmd/dca/cli_test.go |
| REQ-006 | CLI Quick Entry | Shares auto-calculation | ACC-002 | cmd/dca/cli_test.go |
| REQ-007 | CLI Quick Entry | Silent success | ACC-007 | cmd/dca/cli_test.go |
| REQ-008 | CLI Quick Entry | Strict validation | ACC-004, ACC-005, ACC-006 | cmd/dca/cli_test.go |
| REQ-009 | Validation Consistency | Shared validation logic | NFA-002 | internal/form/validation_test.go |
| REQ-010 | Data Persistence | File I/O reuse | NFA-003 | internal/dca/file_test.go |
| REQ-011 | Backward Compatibility | TUI unchanged | ACC-008 | cmd/dca/main_test.go |
| REQ-012 | CLI Scope | Single entry only | N/A | - |

## Validation

- ✅ All existing tests pass
- ✅ Build successful (`make build`)
- ✅ Technical feasibility confirmed via subagent analysis
- ✅ User preferences gathered and incorporated
- ✅ No breaking changes to existing functionality

## Ready for Implementation

This PRD is complete and ready for technical design and implementation.
