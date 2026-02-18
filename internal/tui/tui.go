package tui

import (
	"fmt"
	"rapide/internal/model"
	"rapide/internal/storage"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type modelState struct {
	entries []model.Entry
	cursor  int
	err     error
	ready   bool
}

func (m modelState) Init() tea.Cmd {
	return nil
}

func (m modelState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m modelState) View() string {
	if m.err != nil {
		return ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	if !m.ready {
		return "Initializing Rapanui..."
	}

	// Header
	header := HeaderStyle.Render(
		TitleStyle.Render("RAPIDE") + "  " + DimmedStyle.Strikethrough(false).Render("Project Rapanui"),
	)

	// List of entries
	var content string
	for i, entry := range m.entries {
		style := EntryStyle
		if i == m.cursor {
			style = SelectedEntryStyle
		}

		// Determine bullet color
		bStyle := BulletStyle
		switch entry.Bullet {
		case "•":
			bStyle = bStyle.Foreground(TaskColor)
		case "O":
			bStyle = bStyle.Foreground(EventColor)
		case "-", "—":
			bStyle = bStyle.Foreground(NoteColor)
		case ">":
			bStyle = bStyle.Foreground(MigratedColor)
		}

		bulletStr := bStyle.Render(entry.Bullet)
		contentStr := entry.Content

		// Priority check
		if strings.HasSuffix(entry.Content, "!") {
			contentStr = PriorityStyle.Render(entry.Content)
		}

		// Completion state (dimmed/strikethrough)
		if entry.Bullet == "x" {
			bulletStr = DimmedStyle.Render("x")
			contentStr = DimmedStyle.Render(entry.Content)
		}

		line := fmt.Sprintf("%s %s", bulletStr, contentStr)
		content += style.Render(line) + "\n"
	}

	// Footer
	footer := StatusLineStyle.Render(
		fmt.Sprintf("%d entries • %s navigate • %s quit",
			len(m.entries),
			KeyStyle.Render("j/k"),
			KeyStyle.Render("q")),
	)

	return AppStyle.Render(header + "\n" + content + "\n" + footer)
}

func InitialModel() modelState {
	s, err := storage.NewStorage()
	if err != nil {
		return modelState{err: err}
	}

	entries, err := s.List()
	if err != nil {
		return modelState{err: err}
	}

	return modelState{
		entries: entries,
		ready:   true,
	}
}
