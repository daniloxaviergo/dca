---
id: doc-004
title: 'PRD: Modular Code Refactoring - Split into Domain-Focused Folders'
type: other
created_date: '2026-03-17 11:16'
---
# PRD: Modular Code Refactoring - Split into Domain-Focused Folders

## Overview

### Purpose
Refactor the existing monolithic Go codebase into a well-organized, domain-focused folder structure to improve maintainability, testability, and future extensibility.

### Goals
- **Modularity**: Separate concerns into distinct domain areas (data, form, view, main)
- **Maintainability**: Make it easy to locate and modify specific functionality
- **Testability**: Enable focused unit testing of individual components
- **No Breaking Changes**: Maintain full API compatibility with existing code

## Background

### Problem Statement
The current codebase has all Go files (`dca_entry.go`, `dca_form.go`, `assets_view.go`, `main.go`) in a single flat directory. As the application grows, this structure:
- Makes it difficult to locate specific functionality
- Creates unnecessary coupling between unrelated components
- Complicates testing because components cannot be imported independently
- Hinders onboarding for new developers

### Current State
- All source files reside in the project root: `dca_entry.go`, `dca_form.go`, `assets_view.go`, `main.go`, `main_test.go`, `dca_entry_test.go`, `dca_form_test.go`, `assets_view_test.go`
- All types and functions are in the same package (`main`)
- No clear separation between data models, business logic, UI components, and application entry point
- Test files are colocated with source files

### Proposed Solution
Reorganize the codebase into a domain-focused folder structure:
```
cmd/
  dca/
    main.go
internal/
  dca/
    entry.go
    entry_test.go
  form/
    model.go
    model_test.go
    view.go
    view_test.go
  assets/
    view.go
    view_test.go
    aggregate.go
    aggregate_test.go
```

This structure:
- Separates the entry point from domain logic
- Groups related functionality together
- Uses Go's `internal` package to prevent external imports
- Keeps tests co-located with source files for easier maintenance

## Requirements

### User Stories

- **Developer**: As a developer, I want to easily locate the code for a specific feature so that I can make changes quickly
- **Developer**: As a developer, I want to test individual components in isolation so that I can write focused unit tests
- **Developer**: As a developer, I want to understand the codebase structure at a glance so that I can onboard efficiently

### Functional Requirements

#### Task 1: Create New Folder Structure

Create the directory structure for the refactored codebase.

##### User Flows
1. Create `cmd/dca/` directory for the application entry point
2. Create `internal/dca/` directory for core data structures and file I/O
3. Create `internal/form/` directory for form UI components
4. Create `internal/assets/` directory for assets view components
5. Create `internal/` directory if needed for package grouping

##### Acceptance Criteria
- [ ] All new directories are created
- [ ] No files are moved or deleted during this step
- [ ] Structure matches the approved design

#### Task 2: Extract Core Data Model

Move `dca_entry.go` content to `internal/dca/` and refactor.

##### User Flows
1. Create `internal/dca/entry.go` containing:
   - `DCAEntry` struct
   - `DCAData` struct
   - `LoadEntries()` function
   - `SaveEntries()` function
   - `DCAEntry.Validate()` method
   - `DCAEntry.CalculateShares()` method
2. Move `dca_entry_test.go` to `internal/dca/entry_test.go`
3. Update package declaration from `main` to `dca`
4. Ensure all tests pass

##### Acceptance Criteria
- [ ] `DCAEntry` and `DCAData` types moved to `internal/dca/`
- [ ] All file I/O functions in `internal/dca/`
- [ ] All tests in `internal/dca/entry_test.go` pass
- [ ] No breaking changes to public API

#### Task 3: Extract Form Components

Move `dca_form.go` content to `internal/form/` and refactor.

##### User Flows
1. Create `internal/form/model.go` containing:
   - `FormStep` type and constants
   - `FormField` struct
   - `FormModel` struct
   - `FormModel` methods (Update, Init, etc.)
   - `NewFormModel()` function
2. Create `internal/form/validation.go` containing:
   - All validation functions
   - `CalculateSharesFromValues()` function
   - `RoundTo8Decimals()` function
3. Move `dca_form_test.go` to `internal/form/validation_test.go`
4. Update package declaration
5. Update imports in `main.go` and other files

##### Acceptance Criteria
- [ ] `FormModel` and related types in `internal/form/`
- [ ] Validation logic in separate file if appropriate
- [ ] All tests pass
- [ ] Tests for validation functions in `internal/form/`

#### Task 4: Extract Assets View

Move `assets_view.go` content to `internal/assets/` and refactor.

##### User Flows
1. Create `internal/assets/view.go` containing:
   - `AssetSummary` struct
   - `AssetsViewModel` struct
   - `AssetsView` struct
   - `AssetsView` methods (Update, Init, View, etc.)
   - `NewAssetsView()` function
   - Navigation handlers
2. Create `internal/assets/aggregate.go` containing:
   - `LoadAndAggregateEntries()` function
   - `aggregateEntries()` function
   - `CalculateWeightedAverage()` function
   - `ValidateAssetSummary()` method
3. Move `assets_view_test.go` to `internal/assets/`
4. Update package declaration

##### Acceptance Criteria
- [ ] All asset-related types in `internal/assets/`
- [ ] Aggregation logic separated into `aggregate.go`
- [ ] All tests pass
- [ ] No breaking changes

#### Task 5: Update Main Application

Update `main.go` to use the refactored packages.

##### User Flows
1. Move `main.go` to `cmd/dca/main.go`
2. Update package declaration from `main` to `main`
3. Add import statements for internal packages:
   - `github.com/danilo/scripts/github/dca/internal/dca`
   - `github.com/danilo/scripts/github/dca/internal/form`
   - `github.com/danilo/scripts/github/dca/internal/assets`
4. Update references to use fully qualified package names
5. Ensure all tests pass

##### Acceptance Criteria
- [ ] `main.go` in `cmd/dca/`
- [ ] All internal packages imported correctly
- [ ] Application behavior unchanged
- [ ] All tests pass

#### Task 6: Update Imports and Tests

Update all files to use the new package structure.

##### User Flows
1. Update all import statements in test files
2. Run `go test ./...` to verify all tests pass
3. Run `go build ./...` to verify no compilation errors

##### Acceptance Criteria
- [ ] All import statements updated
- [ ] `go test ./...` passes without errors
- [ ] `go build ./...` succeeds

### Non-Functional Requirements

- **Maintainability**: Code should follow Go best practices for package organization
- **Testability**: Each package should be testable in isolation
- **Compatibility**: No breaking changes to the application's external behavior
- **Build Performance**: Should not increase build times significantly

## Scope

### In Scope
- Creating new folder structure (`cmd/`, `internal/`, subdirectories)
- Moving files to new locations
- Updating import statements
- Updating package declarations
- Running all existing tests to verify no regressions

### Out of Scope
- Changing any functionality or business logic
- Adding new features
- Changing the UI or user experience
- Refactoring code logic beyond what's necessary for the new structure
- Creating external libraries or packages for external consumption

## Technical Considerations

### Existing System Impact
- **User-facing behavior**: No changes
- **Configuration**: No changes required
- **Data files**: No changes required

### Dependencies
- Bubble Tea: No changes required
- Lipgloss: No changes required

### Constraints
- All existing tests must continue to pass
- No breaking changes to the application's API
- Must work with Go modules (already in use)

### Package Organization Details
- `cmd/dca/`: Application entry point only
- `internal/dca/`: Core data types and file I/O
- `internal/form/`: Form UI components and validation
- `internal/assets/`: Asset display and aggregation

## Success Metrics

### Quantitative
- [ ] All existing tests pass (100% pass rate)
- [ ] No new compiler warnings
- [ ] Build time remains stable or improves

### Qualitative
- Code is easier to navigate
- New developers can understand the structure
- Tests are easier to write and maintain

## Timeline & Milestones

### Key Dates
- [ ] Design complete: Approval of folder structure
- [ ] Implementation complete: All files moved and tests pass
- [ ] Testing complete: All tests pass, manual verification
- [ ] Launch: Merge to main branch

## Stakeholders

### Decision Makers
- Developer: Approve folder structure and implementation approach

### Contributors
- Developer: Implementation

## Appendix

### Glossary
- **Package**: Go's module for organizing code
- **Internal package**: A Go package that cannot be imported by external code
- **Domain-focused**: Organized around business domains rather than technical layers

### References
- [Go Project Layout](https://github.com/golang-standards/project-layout): Reference for standard Go project structure
- [Go Packages](https://go.dev/doc/code): Official documentation on Go packages

## Definition of Done

- [ ] All new directories created
- [ ] All files moved to new locations
- [ ] All import statements updated
- [ ] All tests pass (`go test ./...`)
- [ ] Application builds successfully (`go build ./...`)
- [ ] No breaking changes to existing functionality
- [ ] Code follows Go formatting standards (`go fmt`)
