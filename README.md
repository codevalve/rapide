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
rapide work | - Finished the first draft of the Go port!
rapide personal | O Lunch with the team!
rapide "• Fix the login bug!!"
```

rapide list
rapide list --time 3d
rapide list --priority
rapide list --time today
```

## Storage
Entries are stored in `~/.rapide/entries.jsonl`.
