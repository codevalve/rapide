# Git Sync

Rapide can sync your journal across machines using a private Git repository. Your logs stay local-first — Git is an optional bridge, not a requirement.

## Setup

### 1. Create a Private Repository

Create an **empty, private** repository on GitHub (or any Git host). Do not add a README or any files.

### 2. Link It

```bash
rapide sync --setup git@github.com:you/journal.git
```

This configures the remote URL in your `~/.rapide/config.json`.

### 3. First Sync

```bash
rapide sync
```

This will initialize the Git repo in `~/.rapide/`, commit your existing journal, and push to the remote.

## Auto-Sync

Enable automatic sync on every write (both CLI and TUI):

```bash
rapide sync --autosync true
```

When enabled:
- Every `Append` or `saveAll` operation triggers a sync
- The TUI pulls on startup to get the latest entries
- Changes are pushed immediately after each write

To disable:
```bash
rapide sync --autosync false
```

## Manual Sync

Run a one-time sync at any time:

```bash
rapide sync
```

This performs a `git pull --rebase` followed by a `git push`.

## TUI Config

Press `c` inside the TUI to manage sync settings:
- **Remote URL** — change or set the Git remote
- **Autosync** — toggle on/off

## Safety

::: important Rebase Strategy
Rapide uses `git pull --rebase` to avoid merge commits in your journal file. This keeps the JSONL file clean and conflict-free in the common case of single-user access from multiple machines.
:::

::: warning
Git sync is designed for **single-user** journals synced across your own machines. Concurrent writes from multiple users to the same repository may cause conflicts.
:::

## Troubleshooting

### "Not a git repository"

Run `rapide sync` once to initialize the Git repo in `~/.rapide/`.

### Merge conflicts

If auto-sync encounters a conflict (rare with single-user access):
1. Navigate to `~/.rapide/`
2. Resolve the conflict in `entries.jsonl`
3. Run `git add entries.jsonl && git rebase --continue`
4. Run `rapide sync` to push the resolution
