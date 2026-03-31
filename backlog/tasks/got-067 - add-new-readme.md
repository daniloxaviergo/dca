---
id: GOT-067
title: add new readme
status: To Do
assignee:
  - catarina
created_date: '2026-03-31 10:02'
updated_date: '2026-03-31 10:45'
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

The task requires creating a new README file at `./new_readme.md`. Based on the task description being minimal ("create a new readme"), I'll create a comprehensive README that serves as the project's main documentation file. This will be a copy/consolidation of the existing `README.md` content into a new file location.

**Approach:**
- Create a new file `new_readme.md` in the project root directory
- Content will mirror the existing README.md but may be reorganized for clarity
- The file will serve as the primary documentation for the DCA Investment Tracker

**Rationale for this approach:**
- The task description is intentionally minimal ("create a new readme")
- The existing README.md already contains comprehensive documentation
- Creating a new file ensures we don't disrupt existing workflows that reference the current README
- The `./new_readme.md` path suggests this is a new location for the README

### 2. Files to Modify

#### Files to Create:
- **`new_readme.md`** - New README file in project root

#### Files to Review (Read-Only):
- **`README.md`** - Existing README content to reference
- **`go.mod`** - Module name and dependencies
- **`Makefile`** - Development commands reference
- **`cmd/dca/main.go`** - Application entry point
- **`internal/dca/entry.go`** - Core data models
- **`internal/validation/validation.go`** - Validation logic

### 3. Dependencies

**Prerequisites:**
- None (this is a documentation-only task)

**Project Requirements:**
- Go 1.25.7 (specified in `go.mod`)
- Bubble Tea v1.3.10 (TUI framework)
- Lipgloss v1.1.0 (terminal styling)

**External References:**
- Project README.md for current documentation structure
- Go documentation for standard library usage
- Bubble Tea and Lipgloss documentation for framework-specific details

### 4. Code Patterns

**Markdown Documentation Patterns:**
- Use H1 for main title: `# DCA Investment Tracker`
- Use H2 for major sections: `## Overview`, `## Getting Started`, etc.
- Use H3 for subsections: `### Build`, `### Run`, etc.
- Use code blocks with language identifiers: ```bash, ```go
- Use tables for reference information (flags, commands, fields)
- Use bullet points and numbered lists for step-by-step instructions
- Use strong emphasis (**text**) for key terms and important notes

**Content Structure:**
- Project title and overview
- Features list
- Architecture overview (folder structure)
- Getting started guide (prerequisites, build, run)
- Usage documentation (assets view, form view, modal)
- Data format specification
- Testing instructions
- Development commands (Makefile reference)
- Extending the application section

**Style Guidelines:**
- Keep descriptions concise but complete
- Use consistent terminology (e.g., "Assets View" vs "assets view")
- Include code examples for all major features
- Include error handling examples where relevant
- Use absolute paths in documentation for clarity (e.g., `/home/danilo/scripts/github/dca/cmd/dca/main.go`)

### 5. Testing Strategy

**This is a documentation-only task with no code changes:**

- **No unit tests required** - README files don't need unit tests
- **No compiler warnings** - No Go code modifications
- **No `go fmt` checks** - Markdown files are not formatted with go fmt
- **Validation approach:**
  - Verify file exists at `new_readme.md`
  - Verify file contains expected sections
  - Verify formatting is correct (proper markdown syntax)
  - Check for broken links (if any are added)

**Documentation Quality Checks:**
- Ensure all code examples are syntactically correct
- Verify command paths match actual implementation
- Confirm dependency versions in README match `go.mod`
- Validate that file paths referenced are accurate

### 6. Risks and Considerations

**Risks:**
- **Low Risk** - This is a straightforward documentation task
- No code changes, so no breaking changes possible
- No dependencies to update or test to modify

**Considerations:**
- **File location**: The task specifies `./new_readme.md` - ensure this is the correct relative path from project root
- **Content completeness**: The existing README.md is comprehensive; ensure new README doesn't miss critical information
- **Version consistency**: Verify Go version, dependency versions, and command examples match current implementation
- **Documentation overlap**: Consider whether this replaces or supplements the existing README.md
- **Future maintenance**: Document who should update the README when features change

**Trade-offs:**
- **Level of detail**: README should be comprehensive but not overwhelming
- **Examples**: Include enough examples to be helpful but avoid redundancy
- **Architecture depth**: Include enough architecture detail for contributors but not for end users
- **Testing documentation**: Decide how much testing detail to include (unit vs integration vs E2E)

**Deployment/ Rollout:**
- No special deployment steps required
- README will be immediately available after creation
- Consider whether to keep the old README.md or remove it after confirming the new one is adequate
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
