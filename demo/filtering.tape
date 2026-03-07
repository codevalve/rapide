# VHS Tape for Filtering Demo
Set WindowBar Colorful
Set FontSize 16
Set Width 1200
Set Height 600
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
Type "•"
Sleep 1s
Type " work"
Sleep 1s
Type "!"
Sleep 1.5s

# Clear and exit
Backspace 8
Sleep 500ms
Escape
Sleep 500ms
Type "q"
Sleep 1s
