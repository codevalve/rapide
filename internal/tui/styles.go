package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Project Rapanui Colors
	PrimaryColor   = lipgloss.Color("#00ADDB") // Rapide Blue
	SecondaryColor = lipgloss.Color("#555555") // Dimmed Grey
	HighlightColor = lipgloss.Color("#FFD700") // Gold for selection
	ErrorColor     = lipgloss.Color("#FF0000")
	AccentColor    = lipgloss.Color("#FF8C00") // Orange for Priority

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
			Foreground(PrimaryColor).
			Bold(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			PaddingBottom(1).
			MarginBottom(1)

	TitleStyle = lipgloss.NewStyle().
			Background(PrimaryColor).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 1).
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

	SearchStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			Bold(true)

	SearchPromptStyle = lipgloss.NewStyle().
				Foreground(PrimaryColor).
				Bold(true)
)
