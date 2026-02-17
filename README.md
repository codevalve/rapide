# Rapide (Go Edition)

A lightweight, dependency-minimal Go binary for Bullet Journal-style rapid logging.

## Installation

```bash
go build -o rapide
```

## Usage

### Adding Entries

Syntax: `rapide [margin-key] | [bullet] content [!]`

- **Note**: `- content`
- **Task**: `• content` or just `content`
- **Event**: `O content`
- **Action Item**: `AI content` or `A content`
- **Priority**: Add `!` at the end

**Examples:**
```bash
rapide work | - Finished the first draft of the Go port!
rapide personal | O Lunch with the team!
rapide "• Fix the login bug!!"
```

### Listing Entries

```bash
rapide list
rapide list --time 3d
rapide list --priority
rapide list --time today
```

## Storage
Entries are stored in `~/.rapide/entries.jsonl`.
