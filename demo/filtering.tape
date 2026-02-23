# VHS Tape for Filtering Demo
Set WindowBar Colorful
Set FontSize 22
Set Width 1000
Set Height 500
Set Padding 40
Set TypingSpeed 75ms
Output demo/filtering.gif

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

# Start filtering
Type "/"
Sleep 500ms
Type "work"
Sleep 500ms
Type "!"
Sleep 1s

# Clear and exit
Backspace 5
Sleep 500ms
Escape
Sleep 500ms
Type "q"
Sleep 1s
