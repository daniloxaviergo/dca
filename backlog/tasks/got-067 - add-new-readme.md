---
id: GOT-067
title: add new readme
status: To Do
assignee: []
created_date: '2026-03-31 10:02'
updated_date: '2026-03-31 10:36'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
create a new readme and save in ./new_readme.md
<!-- SECTION:DESCRIPTION:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### 1. Technical Approach

This task creates a new README.md file saved to `./new_readme.md`. The content should be a self-contained, user-facing documentation file that follows the project's documentation standards.

**Implementation approach:**

1. **Analyze existing README.md**: Review current README content, structure, and style to ensure consistency
2. **Determine scope**: Decide whether to:
   - Create a copy of the current README (if "new" means a separate document), or
   - Create a specialized README for a specific feature/component (e.g., CLI-only)
3. **Save file**: Write the new README to `./new_readme.md`

**Decision needed:** Clarify whether "new readme" means:
- A complete copy of the existing README (for distribution/archival)
- A specialized README for CLI usage (since doc-013 added CLI features)
- A new README with updated content reflecting current project state

**Recommended approach:** Create a fully-updated README that includes CLI quick entry documentation (from doc-013) if not already present, then save it to `./new_readme.md` as a separate file.

### 2. Files to Modify

| File | Action | Reason |
|------|--------|--------|
| `./new_readme.md` | Create | New README file as specified in task |
| `README.md` | Read only | Source content for new_readme.md (unless CLI section missing) |

**Potential code file references to read:**

| File | Reason |
|------|-- ------|
| `README.md` | Source content for new README |
| `cmd/dca/cli.go` | CLI implementation details (if CLI section missing from README) |
| `cmd/dca/main.go` | Main application structure |
| `internal/form/validation.go` | Validation logic reference |

### 3. Dependencies

- **Existing README.md**: Requires current README content for base文档
- **CLI implementation** (optional): If CLI section not in README, verify `cmd/dca/cli.go` exists and works
- **Go build verification**: Ensure project builds successfully before finalizing README

**Prerequisites:**
- No acceptance criteria defined (task-level detail)
- Build verification: Run `make check` after creation to verify no issues

### 4. Code Patterns

**README content patterns to follow:**

1. **Header format**: `# DCA Investment Tracker` with emoji/logo if applicable
2. **Sections**: Use `##` for main sections, `###` for subsections
3. **Code blocks**: 
   - Shell commands: ```bash
   - Go code: ```go
   - JSON: ```json
4. **Tables**: Use markdown tables for flag/field documentation
5. **Navigation**: Show keyboard shortcuts in code format (e.g., `Esc`, `Ctrl+C`)
6. **Bold for emphasis**: Use `**text**` for field names, filenames, commands

**Style consistency:**
- Match existing README section headings
- Keep similar code block styling
- Maintain same validation rule format
- Use same example format (e.g., `./dca --add --asset BTC --amount 100 --price 50000`)

### 5. Testing Strategy

**Documentation testing approach:**

Since this is a README file creation, testing focuses on content accuracy and completeness:

1. **Content verification**:
   - CLI quick entry section present (if applicable)
   - All features documented (TUI form, asset list, history modal)
   - All flags documented (required/optional)
   - All examples functional

2. **Build verification**:
   ```bash
   # Verify no breaking changes
   make check
   
   # Check build succeeds
   make build
   
   # Run tests
   make test
   ```

3. **Manual testing** (if CLI section included):
   ```bash
   # Test CLI entry
   ./dca --add --asset BTC --amount 100 --price 50000
   echo $?
   
   # Test error handling
   ./dca --add --asset BTC --amount -10 --price 50000
   echo $?
   ```

4. **File verification**:
   - Confirm `./new_readme.md` exists
   - Verify file is readable and properly formatted
   - Check for missing sections

### 6. Risks and Considerations

**Critical decisions needed before implementation:**

1. **Scope ambiguity**: The task description is minimal ("create a new readme"). Clarify if:
   - It's a copy of current README (for archival/distribution)
   - It's a CLI-focused README (given recent CLI work)
   - It's a new README with updated content

2. **PRD reference needed**: Task definition does not reference a PRD. Should this task reference doc-013 (Command-Line Quick Entry) or another PRD?

3. **Content completeness**: Determine what features should be documented:
   - Full application (TUI + CLI)?
   - CLI-only documentation?
   - New feature-specific documentation?

**Potential pitfalls:**
- **Outdated content**: If README already includes CLI features, ensure new_readme.md is not missing updates
- **Inconsistent style**: Ensure README style matches project standards
- **Build drift**: Verify all documented commands work as written

**Recommendation:**
Before implementing, confirm with the user:
1. What content should the new README contain?
2. Should it reference a specific PRD (e.g., doc-013)?
3. What is the purpose of `./new_readme.md` (copy, specialized, or new document)?
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
