package cmd

import (
	"fmt"
	"rapide/internal/model"
	"rapide/internal/tui"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	timestampStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Faint(true) // SecondaryColor
	marginStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true) // PrimaryColor
	bulletStyle    = lipgloss.NewStyle().Bold(true)
	priorityStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Bold(true) // AccentColor
	noteStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEEEEE"))
	idStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Faint(true) // DimmedIDStyle
	doneStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Strikethrough(true)
	pinnedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true) // HighlightColor
)

func renderEntry(e model.Entry, marginWidth int) {
	if marginWidth < 2 {
		marginWidth = 2
	}
	if marginWidth > 12 {
		marginWidth = 12
	}

	// TUI Format: [ID] 02 Jan 15:04 📌 MarginKey Bullet Content
	rawTS := e.Timestamp.Format("02 Jan 15:04")
	// Indicator column (Pin > Time > Agent)
	icon := tui.GetIcon(e.Pinned, e.MarginKey)
	if icon == "" {
		icon = " "
	}
	indicatorStr := pinnedStyle.Render(icon)

	displayMK := tui.StripIcons(e.MarginKey)
	if len(displayMK) > marginWidth {
		displayMK = displayMK[:marginWidth-1] + "…"
	}
	rawMK := fmt.Sprintf("%-*s", marginWidth, displayMK)
	rawBlt := e.Bullet
	rawCnt := e.Content
	rawID := fmt.Sprintf("[%s]", e.ID)
	rawPrio := ""
	if e.Priority {
		rawPrio = "!"
	}

	// Default colors/styles
	ts := timestampStyle.Render(rawTS)
	mk := "  "
	if displayMK != "" {
		mk = marginStyle.Render(rawMK)
	} else {
		mk = strings.Repeat(" ", marginWidth)
	}
	bltStyle := bulletStyle
	switch e.Bullet {
	case "•":
		bltStyle = bltStyle.Foreground(lipgloss.Color("#FFFFFF"))
	case "O":
		bltStyle = bltStyle.Foreground(lipgloss.Color("#AFEEEE"))
	case "-", "—":
		bltStyle = bltStyle.Foreground(lipgloss.Color("#C0C0C0"))
	case ">":
		bltStyle = bltStyle.Foreground(lipgloss.Color("#D3D3D3"))
	}
	blt := bltStyle.Render(rawBlt)
	cnt := noteStyle.Render(rawCnt)
	id := idStyle.Render(fmt.Sprintf("%-6s", rawID))
	prio := ""
	if e.Priority {
		prio = priorityStyle.Render(rawPrio)
	}

	if e.Bullet == "x" {
		ts = doneStyle.Render(rawTS)
		mk = doneStyle.Render(rawMK)
		blt = doneStyle.Render("x")
		cnt = doneStyle.Render(rawCnt)
		id = doneStyle.Render(fmt.Sprintf("%-6s", rawID))
		prio = doneStyle.Render(rawPrio)
		if e.Pinned {
			indicatorStr = doneStyle.Render("📌")
		}
	}

	fmt.Printf("%s %s %s %s %s %s%s\n", id, ts, indicatorStr, mk, blt, cnt, prio)
}
