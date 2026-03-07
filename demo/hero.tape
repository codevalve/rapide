# VHS Tape for Hero Demo
Set WindowBar Colorful
Set FontSize 16
Set Width 1200
Set Height 600
Set Padding 40
Set TypingSpeed 75ms
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
Sleep 1.5s

# Navigate and Pin
Type "jjj"
Sleep 500ms
Type "p"
Sleep 1s
Type "kkk"
Sleep 1s

# Edit an entry
Type "e"
Sleep 500ms
Type " and ship it!"
Sleep 500ms
Enter
Sleep 1s

# Add new entry from TUI
Type "n"
Sleep 500ms
Type "WORK | • Celebrate v2.6 release"
Sleep 500ms
Enter
Sleep 1s

# Filter
Type "/"
Sleep 500ms
Type "work"
Sleep 1.5s
Escape
Sleep 1s

# Quit
Type "q"
Sleep 1s
