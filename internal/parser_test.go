package internal

import (
	"testing"
)

func TestParseEntry(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantMK   string
		wantBlt  string
		wantCnt  string
	}{
		{
			name:     "Standard task",
			args:     []string{"WORK", "|", "•", "Task content"},
			wantMK:   "WORK",
			wantBlt:  "•",
			wantCnt:  "Task content",
		},
		{
			name:     "Note with priority",
			args:     []string{"IDEAS", "|", "-", "Priority note!"},
			wantMK:   "IDEAS",
			wantBlt:  "-",
			wantCnt:  "Priority note!",
		},
		{
			name:     "No margin key",
			args:     []string{"•", "Simple task"},
			wantMK:   "",
			wantBlt:  "•",
			wantCnt:  "Simple task",
		},
		{
			name:     "Multiple pipes",
			args:     []string{"WORK", "|", "PROJ", "|", "•", "Nested pipes"},
			wantMK:   "WORK | PROJ",
			wantBlt:  "•",
			wantCnt:  "Nested pipes",
		},
		{
			name:     "Time margin key",
			args:     []string{"◔ 09:30", "|", "-", "Morning standup"},
			wantMK:   "◔ 09:30",
			wantBlt:  "-",
			wantCnt:  "Morning standup",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseEntry(tt.args)
			if got.MarginKey != tt.wantMK {
				t.Errorf("ParseEntry() MarginKey = %v, want %v", got.MarginKey, tt.wantMK)
			}
			if got.Bullet != tt.wantBlt {
				t.Errorf("ParseEntry() Bullet = %v, want %v", got.Bullet, tt.wantBlt)
			}
			if got.Content != tt.wantCnt {
				t.Errorf("ParseEntry() Content = %v, want %v", got.Content, tt.wantCnt)
			}
		})
	}
}

func TestEntry_String(t *testing.T) {
	// Simple sanity check for String() representation if exists
}
