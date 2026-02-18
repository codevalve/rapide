package tui

import (
	"fmt"
	bujo "rapide/internal"
	"rapide/internal/model"
	"rapide/internal/storage"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	creating    bool
	createInput string
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
		shortID := ""
		if len(e.ID) >= 4 {
			shortID = e.ID[:4]
		}
		if strings.Contains(strings.ToLower(e.Content), query) ||
			strings.Contains(strings.ToLower(e.Bullet), query) ||
			strings.Contains(strings.ToLower(shortID), query) {
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
		if m.creating {
			switch msg.String() {
			case "esc":
				m.creating = false
				m.createInput = ""
			case "enter":
				if m.createInput != "" {
					entry := bujo.ParseEntry([]string{m.createInput})
					s, _ := storage.NewStorage()
					s.Append(entry)
					m.entries, _ = s.List()
					m.creating = false
					m.createInput = ""
					// Move cursor to bottom where new entry is
					m.cursor = len(m.entries) - 1
				} else {
					m.creating = false
				}
			case "backspace":
				if len(m.createInput) > 0 {
					m.createInput = m.createInput[:len(m.createInput)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.createInput += msg.String()
				}
			}
			return m, nil
		}

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
		case "n":
			m.creating = true
			m.createInput = ""
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
		case "d": // Toggle Done
			filtered := m.getFilteredEntries()
			if len(filtered) > 0 {
				entry := filtered[m.cursor]
				if entry.Bullet == "•" {
					entry.Bullet = "x"
				} else if entry.Bullet == "x" {
					entry.Bullet = "•"
				}
				s, _ := storage.NewStorage()
				s.Update(entry.ID, entry)
				m.entries, _ = s.List() // Refresh all
			}
		case "x": // Delete
			filtered := m.getFilteredEntries()
			if len(filtered) > 0 {
				entry := filtered[m.cursor]
				s, _ := storage.NewStorage()
				s.Delete(entry.ID)
				m.entries, _ = s.List()
				if m.cursor >= len(m.getFilteredEntries()) && m.cursor > 0 {
					m.cursor--
				}
			}
		case "m": // Migrate
			filtered := m.getFilteredEntries()
			if len(filtered) > 0 {
				entry := filtered[m.cursor]
				if entry.Bullet == "•" {
					s, _ := storage.NewStorage()
					// 1. Mark current as migrated
					entry.Bullet = ">"
					s.Update(entry.ID, entry)
					// 2. Add new entry for today
					newEntry := model.Entry{
						Content:   entry.Content,
						Bullet:    "•",
						Timestamp: time.Now(),
					}
					s.Append(newEntry)
					m.entries, _ = s.List()
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
	title := TitleStyle.Render("RAPIDE")
	subtitle := DimmedStyle.Strikethrough(false).Render("Project Rapanui")
	header := lipgloss.JoinHorizontal(lipgloss.Bottom, title, "  ", subtitle)

	line := strings.Repeat("─", m.width-4)
	hr := DimmedStyle.Strikethrough(false).Render(line)

	// Calculate visible area for the list
	reservedHeight := 6 // Title line + HR line + Footer block
	visibleHeight := m.height - reservedHeight
	if visibleHeight < 1 {
		visibleHeight = 1
	}

	// List of entries
	filtered := m.getFilteredEntries()
	var contentLines []string
	endIndex := m.startIndex + visibleHeight
	if endIndex > len(filtered) {
		endIndex = len(filtered)
	}

	if len(filtered) == 0 {
		contentLines = append(contentLines, "\n  No entries found matching search query.\n")
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
				contentStr = DimmedStyle.Strikethrough(true).Render(entry.Content)
			}

			// Add short ID
			shortID := ""
			if len(entry.ID) >= 4 {
				shortID = DimmedIDStyle.Render(fmt.Sprintf("[%s]", entry.ID[:4]))
			}

			line := fmt.Sprintf("%-6s %s %s", shortID, bulletStr, contentStr)
			contentLines = append(contentLines, style.Width(m.width-4).Render(line))
		}
	}

	// Padding
	for i := (endIndex - m.startIndex); i < visibleHeight; i++ {
		contentLines = append(contentLines, "")
	}
	content := strings.Join(contentLines, "\n")

	// Dynamic Footer / Status Bar
	var footerStatus string
	if m.creating {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("NEW ENTRY (e.g. • task):"),
			SearchStyle.Render(m.createInput+"_"))
	} else if m.filtering {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("FIND:"),
			SearchStyle.Render(m.filterInput+"_"))
	} else {
		countInfo := fmt.Sprintf("%d entries", len(m.entries))
		if m.filterInput != "" {
			countInfo = fmt.Sprintf("%d/%d found", len(filtered), len(m.entries))
		}
		footerStatus = fmt.Sprintf("%s • %s new • %s filter • %s done • %s migrate • %s delete • %s navigate • %s quit",
			countInfo,
			KeyStyle.Render("n"),
			KeyStyle.Render("/"),
			KeyStyle.Render("d"),
			KeyStyle.Render("m"),
			KeyStyle.Render("x"),
			KeyStyle.Render("j/k"),
			KeyStyle.Render("q"))
	}

	footer := StatusLineStyle.Width(m.width - 4).Render(footerStatus)

	// Combine everything
	return AppStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		header,
		hr,
		content,
		footer,
	))
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
