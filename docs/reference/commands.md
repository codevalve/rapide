# CLI Commands

Complete reference for all Rapide CLI commands.

## Entry Logging

### `rapide [entry]`

The default command — logs a new entry.

```bash
rapide "• Buy coffee"
rapide "work | O Meeting @ 2pm!"
```

See [Syntax & Bullets](/guide/syntax) for the full entry format.

---

## Setup

### `rapide init`

Interactive setup wizard. Prompts for your name, collections, and optional Git sync. Seeds the journal with example entries.

```bash
rapide init
```

::: warning
If your journal already has entries, `init` will ask for confirmation before adding seed entries.
:::

---

## Viewing

### `rapide list`

List journal entries with sorting and filtering.

```bash
rapide list                # Last 30 days
rapide list today          # Today only
rapide list 7d             # Last 7 days
rapide list work           # Filter by collection
```

**Flags:**
| Flag | Short | Description |
|------|-------|-------------|
| `--time` | `-t` | Time filter: `today`, `3d`, `7d` |
| `--filter` | `-f` | Filter by margin key |
| `--bullet` | `-b` | Filter by bullet type: `-`, `O`, `•`, `x` |
| `--priority` | `-p` | Show only priority entries |
| `--all` | `-a` | Include auto-hidden completed tasks |

### `rapide unfinished`

List all non-completed tasks (bullet `•`).

```bash
rapide unfinished
```

### `rapide collections`

Show all margin keys with entry counts.

```bash
rapide collections
```

### `rapide search`

Search all entries by keyword or ID.

```bash
rapide search "deployment"
rapide search a1b2
```

---

## Actions

### `rapide done <id>`

Mark a task as completed (`x`).

```bash
rapide done a1b2
```

### `rapide migrate <id>`

Move a task to today and mark the original as migrated (`>`).

```bash
rapide migrate a1b2
```

### `rapide edit <id> <text>`

Update the content of an existing entry.

```bash
rapide edit a1b2 "Updated task description"
```

### `rapide pin <id>`

Toggle pin status on an entry. Pinned items sort to the top.

```bash
rapide pin a1b2
```

### `rapide delete <id>`

Permanently remove an entry.

```bash
rapide delete a1b2
```

---

## Journal Management

### `rapide trim`

Archive or delete old entries with a guided confirmation.

```bash
rapide trim --before 2025-01-01
```

### `rapide sync`

Sync your journal with a private Git repository.

```bash
rapide sync                    # Pull and push
rapide sync --setup <git-url>  # Link a repo
rapide sync --autosync true    # Enable auto-sync
```

See [Git Sync](/reference/git-sync) for full details.

---

## Interactive

### `rapide tui`

Open the interactive Terminal User Interface.

```bash
rapide tui
```

See [TUI Interface](/guide/tui) for hotkeys and features.

---

## Utility

### `rapide version`

Print the current version.

```bash
rapide version
```

### `rapide completion`

Generate shell autocompletion scripts.

```bash
source <(rapide completion zsh)    # Zsh
source <(rapide completion bash)   # Bash
rapide completion fish | source    # Fish
```
