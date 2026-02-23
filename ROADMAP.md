# Rapide Roadmap

This document outlines the planned trajectory for Rapide beyond the initial local CLI release.

## v2.0.0: Interactive TUI (Released 🗿🚀)
The goal is to transition Rapide from a purely command-driven tool to an interactive experience for reviewing and managing logs.

- **Infrastructure**: Implement a Terminal User Interface (TUI) using [Charm Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Charm Lipgloss](https://github.com/charmbracelet/lipgloss).
- **Features**:
  - **Interactive List View**: Scroll through entries, filter in real-time, and toggle "Done" states with hotkeys.
  - **Editor Integration**: Open entries in an interactive multi-line editor for faster journaling.
  - **Visual Board**: A "Dash" view showing a high-level summary of your collections and pending tasks.

## v2.5.0: Git Sync (Bridge)
Bridging the gap between local-only and cloud-hosted logs by leveraging private Git repositories.

- **Infrastructure**: Lightweight Git automation for the `~/.rapide/` directory.
- **Features**:
  - **Single Command Sync**: `rapide sync` to pull, merge (rebase), and push logs to a configured remote.
  - **Optional Autosync**: Configurable setting to automatically sync on every write or TUI exit.
  - **Setup Wizard**: `rapide sync --setup` to easily link a private repository.

## v2.6.0: Polish & Organization
Gleaning the best "quality of life" features from tools like `nb` to make Rapide a power-user's journal.

- **Features**:
  - **Pinning**: `rapide pin <id>` to keep critical entries or current projects at the top of the list.
  - **Inline Tag Support**: Automatic highlighting and filtering for `#hashtags` within entry content.
  - **Interactive Links**: URL detection in the TUI to quickly open reference links in your browser.
  - **Journal Stats**: `rapide stats` to see activity heatmaps and collection distributions.

## v3.0.0: Hosted Rapide (SSH-as-a-Service)
Taking inspiration from [terminal.shop](https://terminal.shop), v3 aims to make your logs accessible from anywhere without requiring local setup.

- **Infrastructure**: Use [Wish](https://github.com/charmbracelet/wish) to build an SSH server that serves the Rapide TUI.
- **Features**:
  - **Remote Access**: Access your journal from any terminal via `ssh rapide.sh`.
  - **Cloud Sync**: Optional hosted storage (JSONL backed by a secure DB) to keep logs in sync across machines.
  - **Webhooks**: Integration with other tools to push alerts or events into your rapid log via API.

---

*Inspired by the philosophy of Rapid Logging and the aesthetics of the Charm toolchain.*
