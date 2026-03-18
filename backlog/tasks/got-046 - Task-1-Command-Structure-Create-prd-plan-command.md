---
id: GOT-046
title: 'Task 1: Command Structure - Create /prd:plan command'
status: In Progress
assignee:
  - Catarina
created_date: '2026-03-18 23:08'
updated_date: '2026-03-18 23:25'
labels: []
dependencies: []
references:
  - 'backlog/docs/doc-010 - PRD: Task Implementation Plan Generator.md'
ordinal: 2000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create the /prd:plan command with proper argument handling. The command should accept a task ID as argument, retrieve the task via task_view, extract PRD document reference from task's references field, retrieve PRD content via document_view, analyze PRD and generate implementation plan, then update task with plan via task_edit.

instructions to command
```markdown
Save frequently used prompts as shortcut commands to improve work efficiency and ensure consistency.

> [!note]
>
> Custom commands now use Markdown format with optional YAML frontmatter. TOML format is deprecated but still supported for backwards compatibility. When TOML files are detected, an automatic migration prompt will be displayed.

### Quick Overview

| Function         | Description                                | Advantages                             | Priority | Applicable Scenarios                                 |
| ---------------- | ------------------------------------------ | -------------------------------------- | -------- | ---------------------------------------------------- |
| Namespace        | Subdirectory creates colon-named commands  | Better command organization            |          |                                                      |
| Global Commands  | `~/.qwen/commands/`                        | Available in all projects              | Low      | Personal frequently used commands, cross-project use |
| Project Commands | `<project root directory>/.qwen/commands/` | Project-specific, version-controllable | High     | Team sharing, project-specific commands              |

Priority Rules: Project commands > User commands (project command used when names are same)

### Command Naming Rules

#### File Path to Command Name Mapping Table

| File Location                            | Generated Command | Example Call          |
| ---------------------------------------- | ----------------- | --------------------- |
| `~/.qwen/commands/test.md`               | `/test`           | `/test Parameter`     |
| `<project>/.qwen/commands/git/commit.md` | `/git:commit`     | `/git:commit Message` |

Naming Rules: Path separator (`/` or `\`) converted to colon (`:`)

### Markdown File Format Specification (Recommended)

Custom commands use Markdown files with optional YAML frontmatter:

```markdown
---
description: Optional description (displayed in /help)
---

Your prompt content here.
Use {{args}} for parameter injection.
```

| Field         | Required | Description                              | Example                                    |
| ------------- | -------- | ---------------------------------------- | ------------------------------------------ |
| `description` | Optional | Command description (displayed in /help) | `description: Code analysis tool`          |
| Prompt body   | Required | Prompt content sent to model             | Any Markdown content after the frontmatter |

### TOML File Format (Deprecated)

> [!warning]
>
> **Deprecated:** TOML format is still supported but will be removed in a future version. Please migrate to Markdown format.

| Field         | Required | Description                              | Example                                    |
| ------------- | -------- | ---------------------------------------- | ------------------------------------------ |
| `prompt`      | Required | Prompt content sent to model             | `prompt = "Please analyze code: {{args}}"` |
| `description` | Optional | Command description (displayed in /help) | `description = "Code analysis tool"`       |

### Parameter Processing Mechanism

| Processing Method            | Syntax             | Applicable Scenarios                 | Security Features                      |
| ---------------------------- | ------------------ | ------------------------------------ | -------------------------------------- |
| Context-aware Injection      | `{{args}}`         | Need precise parameter control       | Automatic Shell escaping               |
| Default Parameter Processing | No special marking | Simple commands, parameter appending | Append as-is                           |
| Shell Command Injection      | `!{command}`       | Need dynamic content                 | Execution confirmation required before |

#### 1. Context-aware Injection (`{{args}}`)

| Scenario         | TOML Configuration                      | Call Method           | Actual Effect            |
| ---------------- | --------------------------------------- | --------------------- | ------------------------ |
| Raw Injection    | `prompt = "Fix: {{args}}"`              | `/fix "Button issue"` | `Fix: "Button issue"`    |
| In Shell Command | `prompt = "Search: !{grep {{args}} .}"` | `/search "hello"`     | Execute `grep "hello" .` |

#### 2. Default Parameter Processing

| Input Situation | Processing Method                                      | Example                                        |
| --------------- | ------------------------------------------------------ | ---------------------------------------------- |
| Has parameters  | Append to end of prompt (separated by two line breaks) | `/cmd parameter` → Original prompt + parameter |
| No parameters   | Send prompt as is                                      | `/cmd` → Original prompt                       |

🚀 Dynamic Content Injection

| Injection Type        | Syntax         | Processing Order    | Purpose                          |
| --------------------- | -------------- | ------------------- | -------------------------------- |
| File Content          | `@{file path}` | Processed first     | Inject static reference files    |
| Shell Commands        | `!{command}`   | Processed in middle | Inject dynamic execution results |
| Parameter Replacement | `{{args}}`     | Processed last      | Inject user parameters           |

#### 3. Shell Command Execution (`!{...}`)

| Operation                       | User Interaction     |
| ------------------------------- | -------------------- |
| 1. Parse command and parameters | -                    |
| 2. Automatic Shell escaping     | -                    |
| 3. Show confirmation dialog     | ✅ User confirmation |
| 4. Execute command              | -                    |
| 5. Inject output to prompt      | -                    |

Example: Git Commit Message Generation

````markdown
---
description: Generate Commit message based on staged changes
---

Please generate a Commit message based on the following diff:

```diff
!{git diff --staged}
```
````

#### 4. File Content Injection (`@{...}`)

| File Type    | Support Status         | Processing Method           |
| ------------ | ---------------------- | --------------------------- |
| Text Files   | ✅ Full Support        | Directly inject content     |
| Images/PDF   | ✅ Multi-modal Support | Encode and inject           |
| Binary Files | ⚠️ Limited Support     | May be skipped or truncated |
| Directory    | ✅ Recursive Injection | Follow .gitignore rules     |

Example: Code Review Command

```markdown
---
description: Code review based on best practices
---

Review {{args}}, reference standards:

@{docs/code-standards.md}
```

### Practical Creation Example

#### "Pure Function Refactoring" Command Creation Steps Table

| Operation                     | Command/Code                              |
| ----------------------------- | ----------------------------------------- |
| 1. Create directory structure | `mkdir -p ~/.qwen/commands/refactor`      |
| 2. Create command file        | `touch ~/.qwen/commands/refactor/pure.md` |
| 3. Edit command content       | Refer to the complete code below.         |
| 4. Test command               | `@file.js` → `/refactor:pure`             |

```markdown
---
description: Refactor code to pure function
---

Please analyze code in current context, refactor to pure function.
Requirements:

1. Provide refactored code
2. Explain key changes and pure function characteristic implementation
3. Maintain function unchanged
```

### Custom Command Best Practices Summary

#### Command Design Recommendations Table

| Practice Points      | Recommended Approach                | Avoid                                       |
| -------------------- | ----------------------------------- | ------------------------------------------- |
| Command Naming       | Use namespaces for organization     | Avoid overly generic names                  |
| Parameter Processing | Clearly use `{{args}}`              | Rely on default appending (easy to confuse) |
| Error Handling       | Utilize Shell error output          | Ignore execution failure                    |
| File Organization    | Organize by function in directories | All commands in root directory              |
| Description Field    | Always provide clear description    | Rely on auto-generated description          |

#### Security Features Reminder Table

| Security Mechanism     | Protection Effect          | User Operation         |
| ---------------------- | -------------------------- | ---------------------- |
| Shell Escaping         | Prevent command injection  | Automatic processing   |
| Execution Confirmation | Avoid accidental execution | Dialog confirmation    |
| Error Reporting        | Help diagnose issues       | View error information |
```
<!-- SECTION:DESCRIPTION:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria met
- [ ] #2 Documentation updated (comments)
- [ ] #3 #1 All acceptance criteria met
- [ ] #4 #2 Documentation updated (comments)
<!-- DOD:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Command accepts task ID as parameter
- [ ] #2 System correctly retrieves task details
- [ ] #3 System extracts PRD document reference from task
- [ ] #4 System retrieves PRD content
- [ ] #5 System updates task with implementation plan
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
<!-- IMPLEMENTATION:BEGIN -->

### 1. Technical Approach

Create the `/prd:plan` command as a Markdown file in `.qwen/commands/prd/plan.md` that leverages the existing MCP tool infrastructure. The command will:

**How it works:**
1. User invokes `/prd:plan <task_id>` with a task ID as argument
2. Command uses `task_view` to retrieve the task details
3. Extracts PRD document reference from the task's `references` field
4. Uses `document_view` to retrieve the PRD content
5. Analyzes the PRD to identify requirements, dependencies, and implementation order
6. Generates a structured implementation plan following the format in `plan.md` template
7. Uses `task_edit` with `planSet` to update the task with the generated plan

**Architecture decisions:**
- Use the existing Markdown command format (same as `plan.md`, `task.md`, etc.)
- Leverage MCP tools (`task_view`, `document_view`, `task_edit`) - no custom code needed
- Follow the same prompt structure as existing commands with YAML frontmatter
- Use `{{args}}` for task ID parameter injection

**Why this approach:**
- Consistent with existing project command patterns
- No additional tooling or dependencies required
- MCP tools provide all necessary data retrieval and update capabilities
- Plan format matches existing task template in `plan.md`

### 2. Files to Modify

| File | Action | Purpose |
|------|--|--|
| `.qwen/commands/prd/plan.md` | Create | New command implementation for `/prd:plan` |

**New file structure:**
```
.qwen/commands/prd/
└── plan.md          # New command for PRD analysis and plan generation
```

**No existing files need modification** - this is a new command file.

### 3. Dependencies

- **Existing**: MCP tools - `task_view`, `document_view`, `task_edit` (built into Qwen Code)
- **Existing**: PRD document must be referenced in task's `references` field
- **Existing**: Task must have valid document reference format (`backlog/docs/doc-XXX - Title.md`)
- **No Go code dependencies** - this is a command definition, not Go code
- **No new dependencies** - uses existing MCP infrastructure

### 4. Code Patterns

**Follow existing command patterns:**
- YAML frontmatter with `description` field
- Use `---` as frontmatter delimiter
- Prompt body with instructions in Markdown format
- Use `{{args}}` for parameter injection

**Command structure:**
```markdown
---
description: Generate implementation plan from PRD
---

<instructions>
```

**Instruction patterns:**
1. Call `task_list`/`task_search` to find existing work
2. Read task details via `task_view`
3. Extract PRD reference from `references` field
4. Retrieve PRD via `document_view`
5. Analyze PRD sections (Overview, Requirements, etc.)
6. Generate plan following the template format
7. Update task via `task_edit` with `planSet`

### 5. Testing Strategy

**Command testing approach:**
- Test with valid task ID that has PRD reference
- Test with task ID that has no PRD reference (should handle gracefully)
- Test with invalid task ID (should show appropriate error)
- Test with PRD that has missing sections (should handle gracefully)

**Test scenarios:**
1. **Happy path**: Valid task with PRD reference → plan generated and saved
2. **No PRD reference**: Task exists but no PRD → graceful error message
3. **Invalid task ID**: Task doesn't exist → graceful error message
4. **Incomplete PRD**: PRD missing sections → handle gracefully

**Verification:**
- Command can be invoked as `/prd:plan got-046`
- Task is updated with implementation plan
- Plan is accessible via `task_view`
- Plan follows the format from `plan.md` template

### 6. Risks and Considerations

| Risk | Mitigation |
|--|--|
| PRD reference format varies | Extract document ID from reference string; handle common formats |
| PRD is incomplete or malformed | Gracefully handle missing sections; note assumptions in plan |
| Task ID is invalid | Show clear error message via `task_edit` notes |
| PRD analysis is incorrect | Plan should be reviewed by user before coding; this is a guidance tool |
| Multiple PRD references | Process first PRD reference; note if others exist |
| Network connectivity issues | MCP tools handle this; command will fail gracefully |

**Trade-offs:**
- First implementation processes single PRD per invocation (can be extended later)
- Plan is read-only (no validation of PRD correctness)
- Only processes tasks with PRD references (other tasks handled gracefully)

**Dependencies:**
- PRD document must follow standard template structure
- Task must have valid `references` field with document reference
- No code changes to project required
<!-- IMPLEMENTATION:END -->
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
**2026-03-18**: Created `/prd:plan` command file at `.qwen/commands/prd/plan.md`

Command implementation details:
- Command accepts task ID as argument via `{{args}}`
- Uses MCP tools: `task_view`, `document_view`, `task_edit`
- Extracts PRD reference from task's `references` field
- Analyzes PRD requirements and generates implementation plan
- Updates task with plan using `task_edit` with `planSet`
- Handles edge cases: no PRD reference, invalid task ID, incomplete PRD
- Follows plan format from `.qwen/commands/plan.md` template
<!-- SECTION:NOTES:END -->
