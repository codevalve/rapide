# VHS Tape for Hero Demo
Set FontSize 22
Set Width 1200
Set Height 600
Set Padding 40
Output demo/hero.gif

Hide
Type "export RAPIDE_FILE=demo/demo.jsonl"
Enter
Type "clear"
Enter
Show

# Start Rapide TUI
Type "./rapide tui"
Enter
Sleep 1s

# Navigate a bit
Down 2
Sleep 500ms
Up 1
Sleep 500ms

# Toggle a task
Type "d"
Sleep 1s
Type "d"
Sleep 1s

# Migrate a task
Type "m"
Sleep 1s

# Filter
Type "/"
Sleep 500ms
Type "work"
Sleep 1s
Escape
Sleep 1s

# Quit
Type "q"
Sleep 1s
