package cmd

import (
	"fmt"
	"rapide/internal/model"

	"github.com/charmbracelet/lipgloss"
)

var (
	timestampStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#757575"))
	marginStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8")).Bold(true)
	bulletStyle    = lipgloss.NewStyle().Bold(true)
	priorityStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	noteStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEEEEE"))
	idStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Italic(true).Width(5)
	doneStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Strikethrough(true)
	pinnedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true)
)

func renderEntry(e model.Entry, marginWidth int) {
	if marginWidth < 2 {
		marginWidth = 2
	}

	rawTS := e.Timestamp.Format("2006-01-02 15:04")
	rawMK := fmt.Sprintf("%-*s", marginWidth, e.MarginKey)
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
		mk := doneStyle.Render(rawMK)
		blt := doneStyle.Render(rawBlt)
		cnt := doneStyle.Render(rawCnt)
		id := doneStyle.Render(fmt.Sprintf("%-5s", rawID))
		prio := doneStyle.Render(rawPrio)
		pn := "  "
		if e.Pinned {
			pn = doneStyle.Render("📌")
		}
		fmt.Printf("%s | %s | %s | %s %s %s %s\n", id, ts, mk, pn, blt, cnt, prio)
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
	pn := "  "
	if e.Pinned {
		pn = pinnedStyle.Render("📌")
	}

	fmt.Printf("%s | %s | %s | %s %s %s %s\n", id, ts, mk, pn, blt, cnt, prio)
}
