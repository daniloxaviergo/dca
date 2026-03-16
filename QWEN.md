# DCA Project Context

## Project Overview

This is a **task management and project planning workspace** using the [Backlog.md](https://backlog.com) MCP system. The project is named `dca` and is configured for agile software development workflow with structured task management, PRD creation, implementation planning, and task execution.

**Key Technologies:**
- Backlog.md MCP (Model Context Protocol) server for task management
- Local YAML-based configuration
- Markdown-based task and documentation storage
- Git-integrated workflow (auto-commit enabled)

## Directory Structure

```
/home/danilo/scripts/github/dca/
├── .env                      # Environment variables (API keys)
├── .qwen/                    # Qwen Code configuration
│   ├── settings.json         # MCP server configuration
│   ├── agents_current/       # Active agent profiles
│   │   └── backlog-epic-manager.md
│   └── commands/             # Agent command definitions
│       ├── task.md           # Task creation/management
│       ├── plan.md           # Implementation planning
│       ├── prd.md            # Product requirements docs
│       └── exec.md           # Task execution
```

## MCP Commands Available

### Task Management (`/task`)
- Search existing tasks using `task_search` or `task_list`
- Create new tasks with full metadata (description, acceptance criteria, Definition of Done)
- Edit task details (status, plan, notes, acceptance criteria)
- View task details and progress
- Break down large features into atomic tasks
- **DO NOT implement tasks** - only create and structure them

### Implementation Planning (`/plan`)
- Research codebase for task requirements
- Draft implementation plans covering:
  - Technical approach and architecture
  - Files to modify/create
  - Dependencies and integration points
  - Code patterns and conventions
  - Testing strategy
  - Risks and considerations
- **DO NOT implement** - only plan and get approval first

### Product Requirements (`/prd`)
- Create Product Requirements Documents (PRDs)
- Structure PRDs with: Overview, Background, Requirements, Scope, Technical Considerations, Success Metrics, Timeline, Stakeholders
- Use `document_create`, `document_update`, `document_view`, `document_search`
- **DO NOT implement** - only define requirements

### Task Execution (`/exec`)
- Review task details and implementation plan
- Implement approved plans in short loops (code → test → verify)
- Log progress with `task_edit` (notesAppend)
- Check off acceptance criteria as met
- Finalize tasks with PR-style summary and Definition of Done verification
- Update task status to "Done" when complete

## Development Conventions

### Quality Standards
- **Atomic tasks:** Single PR scope
- **Testable AC:** Outcome-focused, not implementation steps
- **Independent tasks:** No future task dependencies
- **Complete:** Sufficient detail for AI agents to implement
- **Verified:** All tests pass before marking complete

### Testing
- Run relevant test suites before finalizing tasks
- Check for regressions
- Verify build success
- No new warnings introduced

## Key Principles

1. **Search first:** Always check for existing work before creating new tasks
2. **Plan before code:** Implementation plans must be approved before coding
3. **Track everything:** All work documented in backlog tasks
4. **Test-driven:** Acceptance criteria verified before completion
5. **Atomic changes:** Each task delivers value independently

## Useful Commands

```bash
# Backlog MCP commands (via Qwen tools)
backlog mcp start
backlog config
backlog task list
backlog task create --title "..." --description "..." --acceptance-criteria "..."

# Standard project commands
git status
git log -n 3
```

## Notes for AI Agents

- The `/task`, `/plan`, `/prd`, and `/exec` commands are MCP tools that map to Backlog.md
- **Always read task details before working** on a task
- **Update task status and progress** in real-time during implementation
