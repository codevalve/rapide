package cmd

import (
	"fmt"
	"rapide/internal/model"

	"github.com/charmbracelet/lipgloss"
)

var (
	timestampStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#757575"))
	marginStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8")).Bold(true).Width(12)
	bulletStyle    = lipgloss.NewStyle().Bold(true)
	priorityStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	noteStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEEEEE"))
	idStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Italic(true).Width(5)
	doneStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Strikethrough(true)
)

func renderEntry(e model.Entry) {
	rawTS := e.Timestamp.Format("2006-01-02 15:04")
	rawMK := e.MarginKey
	rawBlt := e.Bullet
	rawCnt := e.Content
	rawID := e.ID
	rawPrio := ""
	if e.Priority {
		rawPrio = "!"
	}

	// If done, overwrite components with dimmed style
	if e.Bullet == "x" {
		ts := doneStyle.Render(rawTS)
		mk := doneStyle.Render(fmt.Sprintf("%-12s", rawMK))
		blt := doneStyle.Render(rawBlt)
		cnt := doneStyle.Render(rawCnt)
		id := doneStyle.Render(fmt.Sprintf("%-4s", rawID))
		prio := doneStyle.Render(rawPrio)
		fmt.Printf("%s | %s | %s | %s %s %s\n", id, ts, mk, blt, cnt, prio)
		return
	}

	// Default styling
	ts := timestampStyle.Render(rawTS)
	mk := marginStyle.Render(rawMK)
	blt := bulletStyle.Render(rawBlt)
	cnt := noteStyle.Render(rawCnt)
	id := idStyle.Render(rawID)
	prio := ""
	if e.Priority {
		prio = priorityStyle.Render(rawPrio)
	}

	fmt.Printf("%s | %s | %s | %s %s %s\n", id, ts, mk, blt, cnt, prio)
}
