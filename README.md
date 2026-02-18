# Rapide (Go Edition)

A lightweight, dependency-minimal Go binary for Bullet Journal-style rapid logging.

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

## Usage

### Logging

Syntax: `rapide [margin-key] | [bullet] content [!]`

- **Task (`•`)**: `rapide "• Buy coffee"` (or just `rapide "Buy coffee"`)
- **Note (`-`)**: `rapide "- Review PR"`
- **Event (`O`)**: `rapide "O Meeting @ 2pm"`
- **Priority**: Append `!` to any entry to mark it important.

**Examples:**
```bash
rapide "work | - Finished the first draft of the Go port!"
rapide "personal | O Lunch with the team!"
rapide "• Fix the login bug!!"
```

### Commands

```bash
rapide list
rapide list today
rapide list work
rapide unfinished
rapide collections
```

## Shell Autocompletion

Rapide supports generated autocompletion scripts for various shells.

### Homebrew (macOS/Linux)
Completions are managed automatically. Ensure your shell is configured to load Homebrew completions.

### Manual Installation (macOS, Linux, Windows)

#### Zsh
Add this to your `~/.zshrc`:
```zsh
source <(rapide completion zsh)
```

#### Bash
Add this to your `~/.bashrc`:
```bash
source <(rapide completion bash)
```

#### Fish
Add this to `~/.config/fish/config.fish`:
```fish
rapide completion fish | source
```

#### PowerShell (Windows)
Add this to your PowerShell profile:
```powershell
rapide completion powershell | Out-String | Invoke-Expression
```

---

## Storage
Entries are stored in `~/.rapide/entries.jsonl`.
