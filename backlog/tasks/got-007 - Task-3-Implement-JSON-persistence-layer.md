---
id: GOT-007
title: 'Task 3: Implement JSON persistence layer'
status: To Do
assignee: []
created_date: '2026-03-16 21:24'
updated_date: '2026-03-16 23:06'
labels: []
dependencies: []
references:
  - >-
    backlog/docs/doc-002 -
    PRD-DCA-Entry-Form-Persist-Crypto-Investments-in-JSON.md
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement file I/O for JSON data persistence. Create dca_persist.go file with atomic write support.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Load existing dca_entries.json if present (handle file not found gracefully)
- [ ] #2 Create new dca_entries.json if it doesn't exist with proper JSON structure
- [ ] #3 Store entries as map[string][]DCAEntry keyed by asset ticker
- [ ] #4 Write file with 2-space indentation for readability
- [ ] #5 Implement atomic write using temp file + rename pattern
- [ ] #6 Add entry to correct asset array in data structure
- [ ] #7 Display success message: 'Entry saved for [ASSET]' after save
- [ ] #8 Handle file permission errors with clear user message
- [ ] #9 Handle JSON marshal errors with diagnostic message
- [ ] #10 Do not corrupt existing data on write failure
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
# Implementation Plan: Atomic JSON Persistence Layer

## 1. Technical Approach

### Current State
- `dca_entry.go` already contains `DCAEntry`, `DCAData` structs and `LoadEntries()`, `SaveEntries()` functions
- `dca_form.go` uses these functions for form submission
- `main.go` orchestrates loading and saving entries

### What's Missing
The existing `SaveEntries` uses `os.WriteFile` directly, which is **not atomic**. On write failure (disk full, permission error, etc.), the file could be left in a corrupted state. Task 3 requires atomic write using temp file + rename pattern.

### Solution
Update `SaveEntries` to use atomic write pattern:
1. Marshal JSON to bytes
2. Create temp file in same directory with `os.CreateTemp`
3. Write to temp file with `os.WriteFile`
4. Rename temp file to target path with `os.Rename` (atomic on same filesystem)
5. Clean up temp file on error paths

## 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `dca_entry.go` | Modify `SaveEntries` | Add atomic write pattern |
| `dca_entry_test.go` | Add test cases | Verify atomic write and error handling |

## 3. Dependencies

**Prerequisites (already complete):**
- ✓ GOT-005: Data model (`DCAEntry`, `DCAData`) defined
- ✓ GOT-006: CLI input form implemented
- ✓ Existing tests in `dca_entry_test.go`

**No additional dependencies required**

## 4. Code Patterns

**From existing codebase:**
- Error handling: `errors.Is(err, os.ErrNotExist)`, return errors for propagation
- JSON: `json.MarshalIndent` with 2-space indentation (`"  "`)
- Permissions: `0644` (readable by all, writable by owner)

**New atomic write pattern:**
```go
func SaveEntries(filename string, data *DCAData) error {
    // 1. Marshal JSON
    file, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }
    
    // 2. Create temp file in same directory
    tmpfile, err := os.CreateTemp(filepath.Dir(filename), ".dca_entries_*.json")
    if err != nil {
        return err
    }
    
    // 3. Write to temp file
    if _, err := tmpfile.Write(file); err != nil {
        tmpfile.Close()
        os.Remove(tmpfile.Name())
        return err
    }
    
    // 4. Close temp file before rename (required on Windows)
    if err := tmpfile.Close(); err != nil {
        os.Remove(tmpfile.Name())
        return err
    }
    
    // 5. Atomic rename
    if err := os.Rename(tmpfile.Name(), filename); err != nil {
        os.Remove(tmpfile.Name())
        return err
    }
    
    return nil
}
```

## 5. Testing Strategy

### Existing Tests (verify no regression)
- `TestSaveEntries_CreateFile` - New file creation
- `TestSaveEntries_UpdateFile` - File updates
- `TestLoadEntries_*` - All load scenarios

### New Tests to Add
1. `TestSaveEntries_AtomicWrite_Succeeds` - Basic atomic write success
2. `TestSaveEntries_PermissionError_Message` - Clear error message on permission denial
3. `TestSaveEntries_InvalidJSON_Error` - Marshal error handling (AC#9)

### Manual Testing
```bash
go run main.go
# Add an entry, verify success message: "Entry saved for [ASSET]"
# Check dca_entries.json exists with correct structure
```

## 6. Risks and Considerations

### No Blocking Issues
### Design Decisions
- Temp file created in same directory (same filesystem required for atomic rename)
- Temp file pattern: `.dca_entries_*.json` (hidden, easy to identify)
- Cleanup on all error paths to prevent temp file accumulation

### Trade-offs
- Slightly more code complexity (but critical for data integrity)
- Requires `filepath` import (already available in stdlib)

### Deployment
- No migration needed
- Backward compatible - existing files work unchanged
- No config changes required
<!-- SECTION:PLAN:END -->
