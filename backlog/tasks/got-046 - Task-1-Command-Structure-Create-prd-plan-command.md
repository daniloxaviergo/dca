---
id: GOT-046
title: 'Task 1: Command Structure - Create /prd:plan command'
status: To Do
assignee:
  - Catarina
created_date: '2026-03-18 23:08'
updated_date: '2026-03-18 23:11'
labels: []
dependencies: []
references:
  - 'backlog/docs/doc-010 - PRD: Task Implementation Plan Generator.md'
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
- [ ] #2 Unit tests pass (go test)
- [ ] #3 No new compiler warnings
- [ ] #4 Code follows project style (go fmt)
- [ ] #5 PRD referenced in task
- [ ] #6 Documentation updated (comments)
<!-- DOD:END -->
