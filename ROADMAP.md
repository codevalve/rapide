# Rapide Roadmap

This document outlines the planned trajectory for Rapide beyond the initial local CLI release.

## v2.0.0: Interactive TUI (Released 🗿🚀)
The goal is to transition Rapide from a purely command-driven tool to an interactive experience for reviewing and managing logs.

- **Infrastructure**: Implement a Terminal User Interface (TUI) using [Charm Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Charm Lipgloss](https://github.com/charmbracelet/lipgloss).
- **Features**:
  - **Interactive List View**: Scroll through entries, filter in real-time, and toggle "Done" states with hotkeys.
  - **Editor Integration**: Open entries in an interactive multi-line editor for faster journaling.
  - **Visual Board**: A "Dash" view showing a high-level summary of your collections and pending tasks.

## v2.5.0: Git Sync (Bridge) (Released 🌉🚀)
Bridging the gap between local-only and cloud-hosted logs by leveraging private Git repositories.

- **Infrastructure**: Lightweight Git automation for the `~/.rapide/` directory.
- **Features**:
  - **Single Command Sync**: `rapide sync` to pull, merge (rebase), and push logs to a configured remote.
  - **Optional Autosync**: Configurable setting to automatically sync on every write or TUI exit.
  - **Setup Wizard**: `rapide sync --setup` to easily link a private repository.

## v2.6.0: Pinning (Released 📌🚀)
- **Features**:
  - **Pinning**: `rapide pin <id>` to keep critical entries or current projects at the top of the list.

## v2.6.1: Polish & Organization (Released 🛠️🚀)
- **Features**:
  - **TUI Config Editor**: `c` hotkey to manage Git URL and Autosync directly in the UI.
  - **Clipboard Support**: Robust handling for multi-character pastes in the terminal.
  - **Dynamic Layout**: Constrained collection column width with automatic truncation.

## v2.7.2: Security Fix (Released 🔒🚀)
- **Fixes**:
  - **esbuild**: Resolved a security vulnerability in `esbuild` using npm overrides in the documentation site.

## v2.7.1: Build & Security (Released 🛠️🚀)
- **Fixes**:
  - **CI/CD Align**: Updated Go to 1.25 and Node to 22 in GitHub Actions.
  - **Security**: Added Dependabot configuration for automated dependency tracking.

## v2.7.0: Documentation & Onboarding (Released 📖🚀)
The goal is to lower the barrier to entry and ensure Rapide feels accessible to new users while maintaining its professional edge.

- **Infrastructure**: Refined internal documentation and self-documenting CLI help.
- **Features**:
  - **`rapide init`**: Interactive setup wizard to seed your journal.
  - **TUI Help Overlay**: In-app quick reference (press `?`).
  - **VitePress Documentation**: Dedicated docs site on GitHub Pages.

## v3.0.0: Rapide MCP (Model Context Protocol) (Released 🗿🚀)
Bridging the gap between your logs and AI agents by making Rapide a first-class MCP server via a hidden `mcp start` command.

- **Infrastructure**: Implement a Model Context Protocol (MCP) server within the binary.
- **Features**:
  - **Contextual Search**: Allow AI agents to search and retrieve relevant journal entries to inform their tasks.
  - **Automated Logging**: Enable agents to "log a thought" or "record a milestone" directly into Rapide.
  - **Tool Integration**: Expose `rapide` commands (list, done, migrate) as MCP tools.
  - **Privacy First**: Local-first MCP server ensuring your journal stays under your control.

## v3.0.2: CI/CD Maintenance (Released 🛠️🚀)
Maintenance release focusing on fixing build warnings and keeping CI/CD healthy.
- **Infrastructure**:
  - Update GoReleaser configuration to use specified version (`~> v2`).
  - Opt actions runner into Node.js 24 runtime to future-proof workflows.

## v3.0.3: GoReleaser Deprecation Fix (Released 🛠️🚀)
- **Infrastructure**:
  - Address GoReleaser v2 deprecations by renaming `format` to `formats` in `.goreleaser.yaml`.


---

*Inspired by the philosophy of Rapid Logging and the aesthetics of the Charm toolchain.*
