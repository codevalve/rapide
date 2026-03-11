# Rapide

A lightweight, dependency-minimal Go binary for Bullet Journal-style rapid logging.

![Rapide Demo](demo/hero.gif)

**📖 [Read the full documentation →](https://codevalve.github.io/rapide/)**

## Installation

### Homebrew

```bash
brew tap codevalve/homebrew-tap
brew install rapide
```

### Go Install

```bash
go install github.com/codevalve/rapide@latest
```

## Quick Start

```bash
rapide init                    # Interactive setup wizard
rapide "• Buy coffee"          # Log a task
rapide "work | O Meeting!"     # Log to a collection
rapide tui                     # Open the TUI
```

Press **`?`** inside the TUI for a quick-reference of all hotkeys and bullet types.

## Features

| | |
|---|---|
| 🗿 **Minimalist** | One binary, zero bloat, no heavy dependencies |
| ⚡ **Fast CLI** | Log entries in seconds from your terminal |
| 🖥️ **Interactive TUI** | Full-screen interface with hotkeys, filtering, and inline editing |
| 🌉 **Git Sync** | Sync your journal across machines with a private repo |
| 📌 **Pinning** | Keep critical entries at the top |
| 🔒 **Private** | Everything stays local in `~/.rapide/` |

## Documentation

Visit the **[docs site](https://codevalve.github.io/rapide/)** for:

- [Getting Started](https://codevalve.github.io/rapide/getting-started)
- [Syntax & Bullets](https://codevalve.github.io/rapide/guide/syntax)
- [TUI Guide](https://codevalve.github.io/rapide/guide/tui)
- [CLI Commands](https://codevalve.github.io/rapide/reference/commands)
- [Configuration](https://codevalve.github.io/rapide/reference/configuration)
- [Git Sync](https://codevalve.github.io/rapide/reference/git-sync)

## Shell Autocompletion

```bash
source <(rapide completion zsh)     # Zsh
source <(rapide completion bash)    # Bash
rapide completion fish | source     # Fish
```

## Join the Community 🤝

We are actively looking for testers and early adopters! If you're interested:

1. [Open an issue](https://github.com/codevalve/rapide/issues) with your thoughts.
2. Join the discussion in our [GitHub Discussions](https://github.com/codevalve/rapide/discussions).
3. Star the repo to stay updated!
