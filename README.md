# Rapide

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

## Interactive TUI 🗿

Rapide v2.0 introduces a state-of-the-art Terminal User Interface. Run it with:

```bash
rapide tui
```

### Key Hotkeys
- **`/`**: Real-time filtering (search by content, collection, or ID)
- **`n`**: Create a new entry without leaving the TUI
- **`d`**: Toggle task completion (`x`)
- **`m`**: Migrate task to today (`>`)
- **`x`**: Delete entry
- **`T`**: Surgical trim (archive/delete entries before a specific date)
- **`q`**: Quit

## Usage

### Logging

Syntax: `rapide [margin-key] | [bullet] content [!]`

- **Task (`•`)**: `rapide "• Buy coffee"` (or just `rapide "Buy coffee"`)
- **Note (`-`)**: `rapide "- Review PR"`
- **Event (`O`)**: `rapide "O Meeting @ 2pm"`
- **Priority**: Append `!` to any entry to mark it important.
- **Short IDs**: Every entry has a unique 4-character ID (displayed in `list` and `tui`).

### Commands

| Command | Usage | Description |
| :--- | :--- | :--- |
| **`tui`** | `rapide tui` | **New!** Enter the interactive terminal interface |
| **`list`** | `rapide list [today/3d/work]` | List entries (filtered by time or collection) |
| **`done`** | `rapide done <id>` | Mark a task as completed (`x`) |
| **`migrate`** | `rapide migrate <id>` | Move a task to today and mark original as migrated (`>`) |
| **`unfinished`**| `rapide unfinished` | List all non-completed tasks (`•`) |
| **`collections`**| `rapide collections` | See list of margin keys and item counts |
| **`search`** | `rapide search <query>` | Search all entries for a keyword or ID |
| **`edit`** | `rapide edit <id> <text>` | Update the content of an existing entry |
| **`delete`** | `rapide delete <id>` | Permanently remove an entry |
| **`trim`** | `rapide trim [--before DATE]` | Archive or delete old logs with confirmation |
| **`version`** | `rapide version` | Show current version |

#### Command Details

- **Filtering Logs**: `rapide list` supports flags for precision tracking:
  - `-t, --time`: `3d`, `today`, `7d`
  - `-f, --filter`: Filter by margin key (e.g., `-f work`)
  - `-b, --bullet`: Filter by symbol (e.g., `-b O` for events, `-b x` for done)
- **Workflow Icons**:
  - `•` Task (standard)
  - `x` Done (struck-through and dimmed)
  - `>` Migrated (moved forward)
  - `O` Event (notable occurrence)
  - `-` Note (simple record)
  - `!` Priority (important)

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

---

## Join the Community 🤝

We are actively looking for testers and early adopters to help shape the future of Rapide! If you're interested in providing feedback, reporting bugs, or suggesting new features, please:

1. [Open an issue](https://github.com/codevalve/rapide/issues) with your thoughts.
2. Join the discussion in our [GitHub Discussions](https://github.com/codevalve/rapide/discussions).
3. Star the repo to stay updated!

