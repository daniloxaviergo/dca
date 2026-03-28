---
id: GOT-065
title: 'Risk mitigation: CLI-TUI integration validation (REQ-001, REQ-002, REQ-003)'
status: To Do
assignee: []
created_date: '2026-03-28 15:05'
labels:
  - risk-mitigation
  - integration
  - concurrency
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Validate CLI entry doesn't interfere with existing TUI functionality. Test edge cases: simultaneous CLI + TUI access, file locking, state consistency.

**Task Description:**
1. Create integration test: CLI entry while TUI has file open (file locking)
2. Test: CLI entry followed by TUI view (data consistency verification)
3. Test: TUI modification followed by CLI entry (concurrent access)
4. Validate atomic write pattern handles CLI + TUI race conditions
5. Verify no partial writes or data corruption in concurrent scenarios

**Acceptance Criteria:**
- [ ] CLI entry works without TUI打开 (no file lock conflicts)
- [ ] TUI can view data immediately after CLI entry (data consistency)
- [ ] No data loss with concurrent CLI + TUI operations
- [ ] Atomic write pattern prevents partial data corruption

**Integration Test Scenarios:**
| Scenario | Expected Behavior |
|------     |------------------|
| CLI entry while TUI is idle | Success, data persisted |
| CLI entry while TUI editing form | TUI continues, CLI succeeds |
| TUI opens file, CLI writes simultaneously | Wait for file closure, then write |
| CLI writes, TUI immediately reads | Reads latest complete data |

**Failure Modes & Recovery:**
- File lock conflict: Implement retry logic with exponential backoff (max 3 retries, 50ms delay)
- Partial write detected: Validate JSON structure before rename, rollback on error
- Data inconsistency: Add version field to JSON, validate on load

**Assignee:** Developer
**Priority:** High
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
