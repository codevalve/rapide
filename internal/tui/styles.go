package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Project Rapanui Colors (Safe ANSI variants)
	PrimaryColor   = lipgloss.Color("6")   // Cyan
	SecondaryColor = lipgloss.Color("8")   // Grey
	HighlightColor = lipgloss.Color("3")   // Yellow/Gold
	ErrorColor     = lipgloss.Color("9")   // Bright Red
	AccentColor    = lipgloss.Color("208") // Orange/Amber

	// Bullet Colors
	TaskColor     = lipgloss.Color("#FFFFFF") // White
	EventColor    = lipgloss.Color("#AFEEEE") // Pale Turquoise
	NoteColor     = lipgloss.Color("#C0C0C0") // Silver
	MigratedColor = lipgloss.Color("#D3D3D3") // Light Grey

	// Styles
	AppStyle = lipgloss.NewStyle().
			Padding(1, 2)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)

	HeaderStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Bold(true)

	TitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	EntryStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	SelectedEntryStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(HighlightColor).
				PaddingLeft(1).
				Foreground(HighlightColor).
				Bold(true)

	BulletStyle   = lipgloss.NewStyle().Bold(true)
	PriorityStyle = lipgloss.NewStyle().Foreground(AccentColor).Bold(true)
	DimmedStyle   = lipgloss.NewStyle().Foreground(SecondaryColor).Strikethrough(true)

	StatusLineStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			PaddingTop(1).
			MarginTop(1)

	KeyStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	DimmedIDStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Faint(true)

	SearchStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			Bold(true)

	SearchPromptStyle = lipgloss.NewStyle().
				Foreground(PrimaryColor).
				Bold(true)
)
