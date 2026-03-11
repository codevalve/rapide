# Configuration

Rapide stores its configuration alongside your journal entries.

## Config File

**Location:** `~/.rapide/config.json`

```json
{
  "remote_url": "git@github.com:you/journal.git",
  "auto_sync": true,
  "auto_hide_days": 14
}
```

### Fields

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `remote_url` | string | `""` | Git remote URL for sync |
| `auto_sync` | bool | `false` | Automatically push/pull on every write |
| `auto_hide_days` | int | `14` | Number of days after which completed tasks are hidden |

## Environment Variables

### `RAPIDE_FILE`

Override the default journal file path. Useful for development, testing, or managing multiple journals.

```bash
export RAPIDE_FILE=/path/to/custom.jsonl
rapide list
```

::: tip Development Workflow
Rapide includes a `./dev` script that automatically sets `RAPIDE_FILE=demo/demo.jsonl` for safe local testing:
```bash
./dev tui
./dev init
./dev list
```
:::

## Storage Layout

```
~/.rapide/
├── entries.jsonl     # Your journal (one JSON object per line)
├── config.json       # Settings (Git, auto-hide, etc.)
└── archive_*.jsonl   # Archived entries (created by trim)
```

### JSONL Format

Each line in `entries.jsonl` is a self-contained JSON object:

```json
{"id":"a1b2","timestamp":"2026-03-10T12:00:00Z","margin_key":"work","bullet":"•","content":"Review PR","priority":false,"pinned":true}
```

## Editing Config

### Via TUI

Press `c` in the TUI to open the built-in config editor.

### Via `rapide init`

The setup wizard configures Git sync settings as part of the initial setup.

### Via `rapide sync`

```bash
rapide sync --setup <git-url>    # Set remote URL
rapide sync --autosync true      # Toggle auto-sync
```

### Manual

Edit `~/.rapide/config.json` directly with your preferred editor.
