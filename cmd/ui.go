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
)

func renderEntry(e model.Entry) {
	ts := timestampStyle.Render(e.Timestamp.Format("2006-01-02 15:04"))
	mk := marginStyle.Render(e.MarginKey)
	blt := bulletStyle.Render(e.Bullet)
	cnt := noteStyle.Render(e.Content)
	id := idStyle.Render(e.ID)

	prio := ""
	if e.Priority {
		prio = priorityStyle.Render("!")
	}

	fmt.Printf("%s | %s | %s | %s %s %s\n", id, ts, mk, blt, cnt, prio)
}
