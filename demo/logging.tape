# VHS Tape for Logging Demo
Set WindowBar Colorful
Set FontSize 16
Set Width 1200
Set Height 600
Set Padding 40
Set TypingSpeed 75ms
Output demo/logging.gif

Hide
Type "export RAPIDE_FILE=demo/demo.jsonl"
Enter
Type "clear"
Enter
Show

# Show logging from CLI
Type "./rapide '• Review implementation plan!'"
Enter
Sleep 1s

Type "./rapide 'WORK | O Team Sync @ 2pm'"
Enter
Sleep 1s

# Show the list
Type "./rapide list"
Enter
Sleep 2s

# Show just the filtered margin type list
Type "./rapide list WORK"
Enter
Sleep 5s

# Delete demo item file
Hide
Type "rm ./demo/demo.jsonl"

Show

