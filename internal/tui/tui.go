package tui

import (
	"fmt"
	bujo "rapide/internal"
	"rapide/internal/model"
	"rapide/internal/storage"
	"sort"
	"strconv"
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
	trimStep    int       // 0: none, 1: date input, 2: choose a/d, 3: confirm y/n
	trimAction  string    // "archive" or "delete"
	trimInput   string    // raw date input
	trimDate    time.Time // parsed cutoff
	editing     bool
	editInput   string
	configStep  int
	configCfg   *storage.Config
	configInput string
}

func (m modelState) Init() tea.Cmd {
	return nil
}

func (m modelState) getFilteredEntries() []model.Entry {
	var filtered []model.Entry
	query := strings.ToLower(m.filterInput)

	hideCutoff := time.Time{}
	if m.configCfg != nil && m.configCfg.AutoHideDays > 0 {
		hideCutoff = time.Now().Add(-time.Duration(m.configCfg.AutoHideDays) * 24 * time.Hour)
	}

	for _, e := range m.entries {
		// Issue #16: Auto-hide completed items after X days
		if !hideCutoff.IsZero() && e.Bullet == "x" && e.Timestamp.Before(hideCutoff) {
			continue
		}

		if m.filterInput != "" {
			shortID := ""
			if len(e.ID) >= 4 {
				shortID = e.ID[:4]
			}
			if !strings.Contains(strings.ToLower(e.Content), query) &&
				!strings.Contains(strings.ToLower(e.Bullet), query) &&
				!strings.Contains(strings.ToLower(e.MarginKey), query) &&
				!strings.Contains(strings.ToLower(shortID), query) {
				continue
			}
		}
		filtered = append(filtered, e)
	}

	// Sort: Pinned first, then by timestamp (newest first)
	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].Pinned != filtered[j].Pinned {
			return filtered[i].Pinned
		}
		return filtered[i].Timestamp.After(filtered[j].Timestamp)
	})

	return filtered
}

func (m modelState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

	case tea.KeyMsg:
		if m.editing {
			switch msg.String() {
			case "esc":
				m.editing = false
				m.editInput = ""
			case "backspace":
				if len(m.editInput) > 0 {
					m.editInput = m.editInput[:len(m.editInput)-1]
				}
			case "enter":
				if m.editInput != "" {
					filtered := m.getFilteredEntries()
					if len(filtered) > 0 {
						entry := filtered[m.cursor]
						// Use bujo.ParseEntry on the raw input
						newParsed := bujo.ParseEntry([]string{m.editInput})

						// Update fields but preserve ID and Timestamp
						entry.Content = newParsed.Content
						entry.Bullet = newParsed.Bullet
						entry.MarginKey = newParsed.MarginKey
						entry.Priority = newParsed.Priority

						s, _ := storage.NewStorage()
						s.Update(entry.ID, entry)
						m.entries, _ = s.List()
						m.editing = false
						m.editInput = ""
					}
				}
			default:
				if msg.Type == tea.KeyRunes {
					m.editInput += string(msg.Runes)
				} else if msg.Type == tea.KeySpace {
					m.editInput += " "
				}
			}
			return m, nil
		}

		if m.configStep > 0 {
			switch msg.String() {
			case "esc":
				m.configStep = 0
				m.configInput = ""
			case "enter":
				if m.configStep == 1 {
					m.configStep = 2
				} else if m.configStep == 2 {
					m.configStep = 3
					m.configInput = strconv.Itoa(m.configCfg.AutoHideDays)
				} else if m.configStep == 3 {
					days, _ := strconv.Atoi(m.configInput)
					m.configCfg.AutoHideDays = days
					storage.SaveConfig(m.configCfg)
					m.configStep = 0
					m.configInput = ""
				}
			case "backspace":
				if m.configStep == 1 && len(m.configCfg.RemoteURL) > 0 {
					m.configCfg.RemoteURL = m.configCfg.RemoteURL[:len(m.configCfg.RemoteURL)-1]
				} else if m.configStep == 3 && len(m.configInput) > 0 {
					m.configInput = m.configInput[:len(m.configInput)-1]
				}
			case "y", "Y":
				if m.configStep == 2 {
					m.configCfg.AutoSync = true
					m.configStep = 3
					m.configInput = strconv.Itoa(m.configCfg.AutoHideDays)
				} else if m.configStep == 1 {
					m.configCfg.RemoteURL += "y"
				} else if m.configStep == 3 {
					m.configInput += "y"
				}
			case "n", "N":
				if m.configStep == 2 {
					m.configCfg.AutoSync = false
					m.configStep = 3
					m.configInput = strconv.Itoa(m.configCfg.AutoHideDays)
				} else if m.configStep == 1 {
					m.configCfg.RemoteURL += "n"
				} else if m.configStep == 3 {
					m.configInput += "n"
				}
			default:
				if m.configStep == 1 {
					if msg.Type == tea.KeyRunes {
						m.configCfg.RemoteURL += string(msg.Runes)
					} else if msg.Type == tea.KeySpace {
						m.configCfg.RemoteURL += " "
					}
				} else if m.configStep == 3 {
					if msg.Type == tea.KeyRunes {
						m.configInput += string(msg.Runes)
					}
				}
			}
			return m, nil
		}

		if m.trimStep > 0 {
			switch msg.String() {
			case "esc":
				m.trimStep = 0
				m.trimInput = ""
				m.trimAction = ""
			case "backspace":
				if m.trimStep == 1 && len(m.trimInput) > 0 {
					m.trimInput = m.trimInput[:len(m.trimInput)-1]
				}
			case "enter":
				if m.trimStep == 1 {
					if m.trimInput == "" {
						m.trimDate = time.Now().Truncate(24 * time.Hour)
						m.trimStep = 2
					} else {
						d, err := time.Parse("2006-01-02", m.trimInput)
						if err != nil {
							m.err = fmt.Errorf("invalid date format: use YYYY-MM-DD")
							return m, nil
						}
						m.trimDate = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
						m.trimStep = 2
					}
				}
			case "a", "A":
				if m.trimStep == 2 {
					m.trimAction = "archive"
					m.trimStep = 3
				}
			case "d", "D":
				if m.trimStep == 2 {
					m.trimAction = "delete"
					m.trimStep = 3
				}
			case "y", "Y":
				if m.trimStep == 3 {
					s, _ := storage.NewStorage()
					if m.trimAction == "archive" {
						s.ArchiveBefore(m.trimDate)
					} else {
						s.TrimBefore(m.trimDate)
					}
					m.entries, _ = s.List()
					m.trimStep = 0
					m.trimInput = ""
					m.trimAction = ""
					m.cursor = 0
					m.startIndex = 0
				}
			case "n", "N":
				m.trimStep = 0
				m.trimInput = ""
				m.trimAction = ""
			default:
				if m.trimStep == 1 {
					if msg.Type == tea.KeyRunes {
						m.trimInput += string(msg.Runes)
					} else if msg.Type == tea.KeySpace {
						m.trimInput += " "
					}
				}
			}
			return m, nil
		}

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
				if msg.Type == tea.KeyRunes {
					m.createInput += string(msg.Runes)
				} else if msg.Type == tea.KeySpace {
					m.createInput += " "
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
				if msg.Type == tea.KeyRunes {
					m.filterInput += string(msg.Runes)
					m.cursor = 0
					m.startIndex = 0
				} else if msg.Type == tea.KeySpace {
					m.filterInput += " "
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
		case "c":
			cfg, _ := storage.LoadConfig()
			if cfg == nil {
				cfg = &storage.Config{}
			}
			m.configCfg = cfg
			m.configStep = 1
			return m, nil
		case "T":
			m.trimStep = 1
			m.trimInput = ""
			m.err = nil
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
		case "e": // Edit
			filtered := m.getFilteredEntries()
			if len(filtered) > 0 {
				entry := filtered[m.cursor]
				m.editing = true
				// Reconstruct raw string: [Margin | ]Bullet Content
				raw := ""
				if entry.MarginKey != "" {
					raw = entry.MarginKey + " | "
				}
				raw += entry.Bullet + " " + entry.Content
				if entry.Priority {
					raw += "!"
				}
				m.editInput = raw
			}
			return m, nil
		case "d": // Toggle Done
			filtered := m.getFilteredEntries()
			if len(filtered) > 0 {
				entry := filtered[m.cursor]
				switch entry.Bullet {
				case "•":
					entry.Bullet = "x"
				case "x":
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
		case "p": // Toggle Pin
			filtered := m.getFilteredEntries()
			if len(filtered) > 0 {
				entry := filtered[m.cursor]
				s, _ := storage.NewStorage()
				s.TogglePin(entry.ID)
				m.entries, _ = s.List() // Refresh
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

	header := title + "  " + subtitle
	if m.width > 0 && lipgloss.Width(header) > m.width-4 {
		header = title // Simple fallback for very narrow screens
	}

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

	// Dynamic Margin Width calculation (Issue #11)
	maxMargin := 0
	for _, e := range filtered {
		if len(e.MarginKey) > maxMargin {
			maxMargin = len(e.MarginKey)
		}
	}
	if maxMargin > 12 {
		maxMargin = 12
	}
	if maxMargin < 1 {
		maxMargin = 1 // Minimal width for visual separation
	}

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

			// Add Timestamp
			tsStr := TimestampStyle.Render(entry.Timestamp.Format("02 Jan 15:04"))

			// Dynamic MarginKey
			marginStr := ""
			displayMK := entry.MarginKey
			if len(displayMK) > maxMargin {
				displayMK = displayMK[:maxMargin-1] + "…"
			}
			paddedMK := fmt.Sprintf("%-*s", maxMargin, displayMK)
			if entry.MarginKey != "" {
				marginStr = MarginKeyStyle.Render(paddedMK) + " "
			} else {
				marginStr = strings.Repeat(" ", maxMargin+1)
			}

			// Pinned indicator
			pinnedStr := "  "
			if entry.Pinned {
				pinnedStr = PinnedStyle.Render("📌")
			}

			line := fmt.Sprintf("%-6s %s %s %s%s %s", shortID, tsStr, pinnedStr, marginStr, bulletStr, contentStr)
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
	if m.trimStep == 1 {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("TRIM before (YYYY-MM-DD or ENTER for Today):"),
			SearchStyle.Render(m.trimInput+"_"))
	} else if m.trimStep == 2 {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("TRIM ACTION:"),
			SearchStyle.Render("(a)rchive or (d)elete?"))
	} else if m.trimStep == 3 {
		count := 0
		for _, e := range m.entries {
			if e.Timestamp.Before(m.trimDate) {
				count++
			}
		}
		actionStr := "archive"
		if m.trimAction == "delete" {
			actionStr = "DELETE"
		}
		dateStr := m.trimDate.Format("2006-01-02")
		footerStatus = fmt.Sprintf("%s %d entries before %s? (y/n)",
			SearchPromptStyle.Render(fmt.Sprintf("Confirm %s", actionStr)),
			count,
			KeyStyle.Render(dateStr))
	} else if m.editing {
		filtered := m.getFilteredEntries()
		id := "?"
		if len(filtered) > 0 {
			id = filtered[m.cursor].ID
			if len(id) > 4 {
				id = id[:4]
			}
		}
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render(fmt.Sprintf("EDIT [%s]:", id)),
			SearchStyle.Render(m.editInput+"_"))
	} else if m.creating {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("NEW ENTRY (e.g. • task):"),
			SearchStyle.Render(m.createInput+"_"))
	} else if m.filtering {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("FIND:"),
			SearchStyle.Render(m.filterInput+"_"))
	} else if m.configStep == 1 {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("CONFIG Remote URL:"),
			SearchStyle.Render(m.configCfg.RemoteURL+"_"))
	} else if m.configStep == 2 {
		current := "false"
		if m.configCfg.AutoSync {
			current = "true"
		}
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render(fmt.Sprintf("CONFIG AutoSync (y/n) [current: %s]:", current)),
			SearchStyle.Render("_"))
	} else if m.configStep == 3 {
		footerStatus = fmt.Sprintf("%s %s",
			SearchPromptStyle.Render("CONFIG Hide completed items after X days (0 to disable):"),
			SearchStyle.Render(m.configInput+"_"))
	} else {
		countInfo := fmt.Sprintf("%d entries", len(m.entries))
		if m.filterInput != "" {
			countInfo = fmt.Sprintf("%d/%d found", len(filtered), len(m.entries))
		}
		footerStatus = fmt.Sprintf("%s • %s new • %s edit • %s pin • %s filter • %s trim • %s config • %s done • %s quit",
			countInfo,
			KeyStyle.Render("n"),
			KeyStyle.Render("e"),
			KeyStyle.Render("p"),
			KeyStyle.Render("/"),
			KeyStyle.Render("T"),
			KeyStyle.Render("c"),
			KeyStyle.Render("d"),
			KeyStyle.Render("q"))
	}

	footer := StatusLineStyle.Width(m.width - 4).Render(footerStatus)

	// Combine everything
	finalView := lipgloss.JoinVertical(lipgloss.Left,
		header,
		hr,
		content,
		footer,
	)

	return AppStyle.Render(finalView)
}

func InitialModel() modelState {
	s, err := storage.NewStorage()
	if err != nil {
		return modelState{err: err}
	}

	// Pull updates if autosync is enabled
	cfg, _ := storage.LoadConfig()
	if cfg.AutoSync {
		s.Sync()
	}

	entries, err := s.List()
	if err != nil {
		return modelState{err: err}
	}

	return modelState{
		entries:   entries,
		configCfg: cfg,
		ready:     false, // Wait for first WindowSizeMsg
	}
}
