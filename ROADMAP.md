# Rapide Roadmap

This document outlines the planned trajectory for Rapide beyond the initial local CLI release.

## v2.0.0: Interactive TUI (Released 🗿🚀)
The goal is to transition Rapide from a purely command-driven tool to an interactive experience for reviewing and managing logs.

- **Infrastructure**: Implement a Terminal User Interface (TUI) using [Charm Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Charm Lipgloss](https://github.com/charmbracelet/lipgloss).
- **Features**:
  - **Interactive List View**: Scroll through entries, filter in real-time, and toggle "Done" states with hotkeys.
  - **Editor Integration**: Open entries in an interactive multi-line editor for faster journaling.
  - **Visual Board**: A "Dash" view showing a high-level summary of your collections and pending tasks.

## v3.0.0: Hosted Rapide (SSH-as-a-Service)
Taking inspiration from [terminal.shop](https://terminal.shop), v3 aims to make your logs accessible from anywhere without requiring local setup.

- **Infrastructure**: Use [Wish](https://github.com/charmbracelet/wish) to build an SSH server that serves the Rapide TUI.
- **Features**:
  - **Remote Access**: Access your journal from any terminal via `ssh rapide.sh`.
  - **Cloud Sync**: Optional hosted storage (JSONL backed by a secure DB) to keep logs in sync across machines.
  - **Webhooks**: Integration with other tools to push alerts or events into your rapid log via API.

---

*Inspired by the philosophy of Rapid Logging and the aesthetics of the Charm toolchain.*
