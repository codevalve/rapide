package tui

import (
	"fmt"
	"rapide/internal/model"
	"rapide/internal/storage"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type modelState struct {
	entries     []model.Entry
	cursor      int
	startIndex  int
	width       int
	height      int
	err         error
	ready       bool
	filtering   bool
	filterInput string
}

func (m modelState) Init() tea.Cmd {
	return nil
}

func (m modelState) getFilteredEntries() []model.Entry {
	if m.filterInput == "" {
		return m.entries
	}
	var filtered []model.Entry
	query := strings.ToLower(m.filterInput)
	for _, e := range m.entries {
		if strings.Contains(strings.ToLower(e.Content), query) ||
			strings.Contains(strings.ToLower(e.Bullet), query) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func (m modelState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

	case tea.KeyMsg:
		if m.filtering {
			switch msg.String() {
			case "esc", "enter":
				m.filtering = false
			case "backspace":
				if len(m.filterInput) > 0 {
					m.filterInput = m.filterInput[:len(m.filterInput)-1]
					m.cursor = 0
					m.startIndex = 0
				}
			default:
				if len(msg.String()) == 1 {
					m.filterInput += msg.String()
					m.cursor = 0
					m.startIndex = 0
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "/":
			m.filtering = true
			return m, nil
		case "esc":
			m.filterInput = "" // Clear filter
			m.cursor = 0
			m.startIndex = 0
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.startIndex {
					m.startIndex = m.cursor
				}
			}
		case "down", "j":
			filtered := m.getFilteredEntries()
			if m.cursor < len(filtered)-1 {
				m.cursor++
				visibleHeight := m.height - 8
				if visibleHeight < 1 {
					visibleHeight = 1
				}
				if m.cursor >= m.startIndex+visibleHeight {
					m.startIndex = m.cursor - visibleHeight + 1
				}
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

	// Calculate visible area for the list
	reservedHeight := 7
	visibleHeight := m.height - reservedHeight
	if visibleHeight < 1 {
		visibleHeight = 1
	}

	// List of entries
	filtered := m.getFilteredEntries()
	var content string
	endIndex := m.startIndex + visibleHeight
	if endIndex > len(filtered) {
		endIndex = len(filtered)
	}

	if len(filtered) == 0 {
		content = "\n  No entries found matching search query.\n"
	} else {
		for i := m.startIndex; i < endIndex; i++ {
			entry := filtered[i]
			style := EntryStyle
			if i == m.cursor {
				style = SelectedEntryStyle
			}

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

			if strings.HasSuffix(entry.Content, "!") {
				contentStr = PriorityStyle.Render(entry.Content)
			}

			if entry.Bullet == "x" {
				bulletStr = DimmedStyle.Render("x")
				contentStr = DimmedStyle.Render(entry.Content)
			}

			line := fmt.Sprintf("%s %s", bulletStr, contentStr)
			content += style.Width(m.width-4).Render(line) + "\n"
		}
	}

	// Padding
	for i := (endIndex - m.startIndex); i < visibleHeight; i++ {
		content += "\n"
	}

	// Dynamic Footer / Status Bar
	var footerStatus string
	if m.filtering {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("FIND:"),
			SearchStyle.Render(m.filterInput+"_"))
	} else {
		countInfo := fmt.Sprintf("%d entries", len(m.entries))
		if m.filterInput != "" {
			countInfo = fmt.Sprintf("%d/%d found", len(filtered), len(m.entries))
		}
		footerStatus = fmt.Sprintf("%s • %s filter • %s reset • %s navigate • %s quit",
			countInfo,
			KeyStyle.Render("/"),
			KeyStyle.Render("esc"),
			KeyStyle.Render("j/k"),
			KeyStyle.Render("q"))
	}

	footer := StatusLineStyle.Width(m.width - 4).Render(footerStatus)

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
		ready:   false, // Wait for first WindowSizeMsg
	}
}
