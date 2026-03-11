# Getting Started

Get up and running with Rapide in under a minute.

## Installation

### Homebrew (macOS / Linux)

```bash
brew tap codevalve/homebrew-tap
brew install rapide
```

### Go Install

```bash
go install github.com/codevalve/rapide@latest
```

## Quick Setup

Run the interactive setup wizard to configure your journal and seed it with example entries:

```bash
rapide init
```

The wizard will:
1. Ask your **name** (used in the welcome entry)
2. Let you define your **collections** — margin keys like `work`, `health`, `ideas`
3. Optionally link a **private Git repo** for cross-device sync
4. Seed your journal with **example entries** showing every bullet type

## Your First Entry

Log an entry from the command line:

```bash
rapide "• Buy coffee"
```

Add a **collection** (margin key) with a pipe separator:

```bash
rapide "work | O Team sync @ 2pm"
```

Mark something as **priority** with `!`:

```bash
rapide "• Review PR for launch!"
```

## Open the TUI

Launch the interactive terminal interface:

```bash
rapide tui
```

Press **`?`** for a full quick-reference of hotkeys and bullet types.

::: tip First-Run Experience
If your journal is empty when you open the TUI, the help screen opens automatically to get you oriented.
:::

## What's Next?

- Learn the full [Syntax & Bullets](/guide/syntax) system
- Explore the [TUI Interface](/guide/tui) in depth
- See all [CLI Commands](/reference/commands)
- Set up [Git Sync](/reference/git-sync) for cross-device access
