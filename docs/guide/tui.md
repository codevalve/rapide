# TUI Interface

Rapide's interactive Terminal User Interface gives you a full-screen dashboard for managing your journal.

## Launching

```bash
rapide tui
```

::: tip First-Run Help
If your journal is empty, the help overlay opens automatically. Press `?` at any time to toggle it.
:::

## Hotkeys

| Key | Action |
|-----|--------|
| `n` | Create a new entry |
| `e` | Edit the selected entry inline |
| `d` | Toggle task completion (`x`) |
| `m` | Migrate task to today (`>`) |
| `x` | Delete entry |
| `p` | Pin / unpin (pinned items stay at top) |
| `/` | Start real-time filter (search content, collection, or ID) |
| `T` | Trim journal (archive or delete entries before a date) |
| `c` | Open config editor (Git URL, Autosync, Auto-hide days) |
| `?` | Toggle the quick-reference help overlay |
| `esc` | Clear filter / dismiss help |
| `q` | Quit |

## Navigation

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `Home` / `g` | Jump to top |
| `End` / `G` | Jump to bottom |

## Filtering

Press `/` to enter filter mode. Type your query — results update in real-time. Matches are found across:

- **Content** (entry text)
- **Collection** (margin key)
- **ID** (4-character short hash)

Press `esc` to clear the filter and return to the full list.

## Creating Entries

Press `n` to create a new entry. The full [syntax](/guide/syntax) is supported:

```
work | • Review the deployment checklist!
```

Press `Enter` to save or `esc` to cancel.

## Editing Entries

Press `e` on a selected entry to edit it inline. The existing raw text is pre-filled.

## Config Editor

Press `c` to open the in-app config editor. You can set:

- **Remote URL** — Git repository for sync
- **Autosync** — Automatically push/pull on every write
- **Auto-hide days** — Hide completed tasks older than N days

## Trimming

Press `T` to start the trim wizard:

1. Enter a **cutoff date** (e.g., `2025-01-01`)
2. Choose **Archive** (saves to a `.jsonl` file) or **Delete** (permanent removal)
3. **Confirm** the action

## Display Notes

- **Pinned entries** always appear at the top of the list, regardless of timestamp
- **Completed tasks** (`x`) older than the configured auto-hide period are hidden by default — press `a` in list mode or use `--all` in the CLI to reveal them
- **Priority entries** are visually highlighted with distinct colors
