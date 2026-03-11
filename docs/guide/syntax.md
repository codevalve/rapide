# Syntax & Bullets

Rapide uses **Bullet Journal** notation adapted for the terminal. Every entry follows a simple pattern:

```
[collection | ] [bullet] content [!]
```

## Bullets

Bullets indicate the **type** of each entry:

| Bullet | Meaning | Example |
|--------|---------|---------|
| `•` | Task | `• Review implementation plan` |
| `x` | Done | `x Review implementation plan` |
| `O` | Event | `O Team sync @ 2pm` |
| `-` | Note | `- Interesting article on Go concurrency` |
| `>` | Migrated | `> Moved to next sprint` |
| `<` | Scheduled | `< Follow up next Monday` |
| `AI` | Action Item | `AI Update the deployment scripts` |

::: tip Default Bullet
If you omit the bullet, Rapide defaults to `•` (task):
```bash
rapide "Buy coffee"  # → same as rapide "• Buy coffee"
```
:::

## Collections (Margin Keys)

Collections group entries by topic. Use a **pipe** (`|`) to separate the collection from the entry:

```bash
rapide "work | • Review PR"
rapide "health | O Doctor appointment @ 10am"
rapide "ideas | - App for tracking habits"
```

Collections appear as a dedicated column in both the CLI `list` and the TUI.

### Viewing Collections

```bash
rapide collections        # List all margin keys with counts
rapide list work          # Show only entries in the "work" collection
rapide list -f health     # Same, using the --filter flag
```

## Priority

Append `!` to any entry to mark it as **priority**:

```bash
rapide "• Fix production bug!"
rapide "work | • Deploy hotfix!"
```

Priority entries are visually highlighted in both the CLI and TUI, and can be filtered with:

```bash
rapide list -p            # Show only priority entries
```

## Entry IDs

Every entry gets a unique **4-character short ID** (derived from a SHA-1 hash). These IDs are displayed in `list` and the TUI, and are used to target specific entries:

```bash
rapide done a1b2          # Mark entry a1b2 as done
rapide pin a1b2           # Pin entry a1b2
rapide edit a1b2 "new text"
rapide delete a1b2
```

## Complete Syntax Examples

```bash
# Simple task
rapide "• Buy groceries"

# Task with collection
rapide "home | • Fix leaky faucet"

# Priority event with collection
rapide "work | O Board meeting @ 3pm!"

# Note
rapide "- Remember: deadlines are artificial"

# Action item
rapide "AI Send quarterly report to team"
```
