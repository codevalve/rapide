# AGENTS.md (Rapide Agent Blueprint)

Welcome, Comrade Agent! You are helping build **Rapide**, a lightweight, dependency-minimal Go binary for Bullet Journal-style rapid logging. 🗿🚀

## 🗿 Project DNA
- **Minimalist**: Avoid heavy external dependencies.
- **Local-First**: Private logs by default, stored in `~/.rapide/entries.jsonl`.
- **Aesthetic**: Uses the **Charmbracelet** toolchain (`bubbletea`, `lipgloss`).
- **Bullet-Journal Syntax**: Content is driven by margin keys and bullets (e.g., `•` for tasks, `O` for events).

---

## 🏗️ Architecture Guide

### 📂 Core Module: `internal/`
- **`model/`**: Defines the `Entry` struct.
- **`parser.go`**: The logic that turns raw strings into journal entries. **CRITICAL**: If adding syntax, update this file and its tests.
- **`storage/`**: Handles JSONL reading/writing and the Git sync bridge.

### 🕹️ TUI: `internal/tui/`
- **`tui.go`**: The `bubbletea.Model` state machine.
- **`styles.go`**: All UI styling must be defined here using `lipgloss`. DO NOT use ad-hoc styles in views.

### 🛠️ Commands: `cmd/`
- Standard Cobra subcommands.
- Keep `root.go` focused on bootstrapping. New features should be subcommands.

---

## 🚦 Guidelines for AI Agents

1.  **Keep it "Charm"**: Follow the design aesthetic of `internal/tui/styles.go`. Use vibrant, terminal-safe colors and clean borders.
2.  **Git Sync Safety**: When modifying sync logic, ensure we preserve user logs. Use `--rebase` for pulls to avoid merge conflicts in the journal repository.
3.  **MCP-First Mindset**: As of v3.0 planning, we are pivoting toward **MCP (Model Context Protocol)**. Any new command should also be considered for potential exposure as an MCP tool.
4.  **CLI Compatibility**: Ensure every feature available in the TUI (like `pin` or `done`) is also available as a standalone CLI command for automation.
5.  **Conventional Commits**: Use the **Conventional Commits** specification for all commit messages (e.g., `feat:`, `fix:`, `chore:`, `docs:`).
6.  **Issue Lifecycle**: When working on tasks, always create or reference a GitHub issue. Close the issue using the GitHub CLI or API once the task is merged and verified.
7.  **Roadmap & Milestone Sync**: Always ensure `ROADMAP.md` and GitHub milestones are kept in sync whenever feature plans change or releases are completed.
8. **Bug Fixes**: When a bug is reported, don't start by trying to fix it. Instead, start by writing a test that reproduces the bug. Then, try to fix the bug and prove it with a passing test. Always use the `./dev` wrapper for test isolation. Create or reference a GitHub issue and use the Conventional Commit `fix:` type. Close the issue once verified.
9. **Context Window Protection**: Follow the mandatory routing rules in `GEMINI.md`. Use `ctx_execute`, `ctx_batch_execute`, and `ctx_execute_file` for large outputs to avoid flooding the context window.

---

## 🧪 Developer Workflow

When developing or testing locally, **never run `./rapide` directly** against your personal journal. Use the `./dev` wrapper instead:

```bash
./dev tui       # Open TUI against demo/demo.jsonl
./dev init      # Test the init wizard
./dev list      # Check demo entries
```

This sets `RAPIDE_FILE=demo/demo.jsonl` automatically, isolating all test writes from your personal `~/.rapide/entries.jsonl`. The `demo/*.jsonl` file is gitignored.

Check `.agents/workflows/` for automated SOPs:
- **`scaffold-command.md`**: How to add a new CLI command.
- **`release-flow.md`**: Steps for bumping version and release.
- **`setup-mcp.md`**: Current blueprint for v3.0 MCP integration.
- **`triage-issues.md`**: The automated GitHub issue triage and engineering workflow.
