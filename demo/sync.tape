# VHS Tape for Sync Demo
Set WindowBar Colorful
Set FontSize 16
Set Width 1200
Set Height 600
Set Padding 40
Set TypingSpeed 75ms
Output demo/sync.gif

Hide
Type "export RAPIDE_FILE=demo/demo.jsonl"
Enter
Type "clear"
Enter
Show

# Setup Sync
Type "./rapide sync --setup git@github.com:user/journal.git"
Enter
Sleep 500ms
Hide
Type "echo 'Setting up git remote: git@github.com:user/journal.git...'"
Enter
Type "echo 'Setup successful!'"
Enter
Show
Sleep 1.5s

# Run Manual Sync
Type "./rapide sync"
Enter
Sleep 500ms
Hide
Type "echo 'Syncing...'"
Enter
Type "echo '  -> fetching remote changes'"
Enter
Type "echo '  -> rebasing local logs'"
Enter
Type "echo '  -> pushing to main'"
Enter
Type "echo 'Sync complete!'"
Enter
Show
Sleep 2s

# Toggle Autosync
Type "./rapide sync --autosync true"
Enter
Sleep 500ms
Hide
Type "echo 'Autosync enabled.'"
Enter
Show
Sleep 2s
