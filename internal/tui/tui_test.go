package tui

import (
	"rapide/internal/model"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTUI_Navigation(t *testing.T) {
	m := modelState{
		entries: []model.Entry{
			{Content: "Entry 1"},
			{Content: "Entry 2"},
			{Content: "Entry 3"},
		},
		cursor: 0,
		ready:  true,
		height: 20,
		width:  80,
	}

	// Test Down (j)
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	m = newModel.(modelState)
	if m.cursor != 1 {
		t.Errorf("Expected cursor 1 after 'j', got %d", m.cursor)
	}

	// Test Up (k)
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("k")})
	m = newModel.(modelState)
	if m.cursor != 0 {
		t.Errorf("Expected cursor 0 after 'k', got %d", m.cursor)
	}

	// Test bounds (Up at 0)
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("k")})
	m = newModel.(modelState)
	if m.cursor != 0 {
		t.Errorf("Expected cursor 0 to stay 0, got %d", m.cursor)
	}

	// Test bounds (Down at max)
	m.cursor = 2
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	m = newModel.(modelState)
	if m.cursor != 2 {
		t.Errorf("Expected cursor 2 to stay 2, got %d", m.cursor)
	}
}

func TestTUI_Filtering(t *testing.T) {
	m := modelState{
		entries: []model.Entry{
			{Content: "Apple", Bullet: "•"},
			{Content: "Banana", Bullet: "•"},
			{Content: "Cherry", Bullet: "•"},
		},
		cursor: 0,
		ready:  true,
		height: 20,
	}

	// 1. Trigger filter mode
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("/")})
	m = newModel.(modelState)
	if !m.filtering {
		t.Error("Expected filtering mode to be true")
	}

	// 2. Type "App"
	mModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("A")})
	mModel, _ = mModel.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	mModel, _ = mModel.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	m = mModel.(modelState)

	if m.filterInput != "App" {
		t.Errorf("Expected filterInput 'App', got '%s'", m.filterInput)
	}

	filtered := m.getFilteredEntries()
	if len(filtered) != 1 {
		t.Errorf("Expected 1 filtered entry, got %d", len(filtered))
	}
	if filtered[0].Content != "Apple" {
		t.Errorf("Expected 'Apple', got '%s'", filtered[0].Content)
	}

	// 3. Clear filter with Esc
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc, Runes: []rune("")})
	m = newModel.(modelState)
	if m.filtering {
		t.Error("Expected filtering mode to be false after Esc")
	}

	// 4. Reset filter completely with second Esc (when not filtering)
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc, Runes: []rune("")})
	m = newModel.(modelState)
	if m.filterInput != "" {
		t.Error("Expected filterInput to be empty after second Esc")
	}
}
